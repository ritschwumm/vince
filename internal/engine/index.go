package engine

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/expression"
	"github.com/dolthub/go-mysql-server/sql/types"
	v1 "github.com/vinceanalytics/vince/gen/proto/go/vince/store/v1"
)

type Index struct {
	TableName  string
	Exprs      []sql.Expression
	exprString []string
	PrefixLens []uint16
}

var _ sql.Index = (*Index)(nil)
var _ sql.IndexAddressable = (*Table)(nil)

func (idx *Index) Database() string                    { return "vince" }
func (idx *Index) ColumnExpressions() []sql.Expression { return idx.Exprs }
func (idx *Index) IsGenerated() bool                   { return false }
func (idx *Index) Expressions() []string               { return idx.exprString }
func (idx *Index) ExtendedExpressions() []string       { return idx.exprString }
func (idx *Index) CanSupport(...sql.Range) bool        { return true }
func (idx *Index) IsUnique() bool                      { return false }
func (idx *Index) IsSpatial() bool                     { return false }
func (idx *Index) IsFullText() bool                    { return false }
func (idx *Index) Comment() string                     { return "" }
func (idx *Index) IndexType() string                   { return "BTREE" }
func (idx *Index) Table() string                       { return idx.TableName }
func (idx *Index) ID() string                          { return "vince_fields_idx" }

func (idx *Index) PrefixLengths() []uint16 { return idx.PrefixLens }

func (idx *Index) ColumnExpressionTypes() []sql.ColumnExpressionType {
	cets := make([]sql.ColumnExpressionType, len(idx.Exprs))
	for i, expr := range idx.Exprs {
		cets[i] = sql.ColumnExpressionType{
			Expression: expr.String(),
			Type:       expr.Type(),
		}
	}
	return cets
}

func (idx *Index) ExtendedColumnExpressionTypes() []sql.ColumnExpressionType {
	cets := make([]sql.ColumnExpressionType, 0, len(idx.Exprs))
	cetsInExprs := make(map[string]struct{})
	for _, expr := range idx.Exprs {
		cetsInExprs[strings.ToLower(expr.(*expression.GetField).Name())] = struct{}{}
		cets = append(cets, sql.ColumnExpressionType{
			Expression: expr.String(),
			Type:       expr.Type(),
		})
	}
	return cets
}

func (idx *Index) build(ctx *sql.Context,
	ranges ...sql.Range) (idxFilters []IndexFilter, valFilters []ValueFilter, rangeCollectionExpr sql.Expression, err error) {
	if len(ranges) == 0 {
		return
	}
	exprs := idx.Exprs

	for rangIdx, rang := range ranges {
		if len(exprs) < len(rang) {
			err = fmt.Errorf("expected different key count: exprs(%d) < (ranges[%d])(%d)", len(exprs), rangIdx, len(rang))
			return
		}
		var rangeExpr sql.Expression
		for i, rce := range rang {
			field := indexedField(exprs[i])
			lo, hi := bounds(rce)
			var rangeColumnExpr sql.Expression
			switch rce.Type() {
			// Both Empty and All may seem like strange inclusions, but if only one range is given we need some
			// expression to evaluate, otherwise our expression would be a nil expression which would panic.
			case sql.RangeType_Empty:
				rangeColumnExpr = expression.NewEquals(expression.NewLiteral(1, types.Int8), expression.NewLiteral(2, types.Int8))
			case sql.RangeType_All:
				rangeColumnExpr = expression.NewEquals(expression.NewLiteral(1, types.Int8), expression.NewLiteral(1, types.Int8))
			case sql.RangeType_EqualNull:
				rangeColumnExpr = expression.NewIsNull(exprs[i])
			case sql.RangeType_GreaterThan:
				if sql.RangeCutIsBinding(rce.LowerBound) {
					idxFilters = append(idxFilters,
						buildIndex(field, lo, hi, Gt),
					)
					rangeColumnExpr = expression.NewGreaterThan(exprs[i], expression.NewLiteral(sql.GetRangeCutKey(rce.LowerBound), rce.Typ.Promote()))
				} else {
					rangeColumnExpr = expression.NewNot(expression.NewIsNull(exprs[i]))
				}
			case sql.RangeType_GreaterOrEqual:
				idxFilters = append(idxFilters,
					buildIndex(field, lo, hi, GtEg),
				)
				rangeColumnExpr = expression.NewGreaterThanOrEqual(exprs[i], expression.NewLiteral(sql.GetRangeCutKey(rce.LowerBound), rce.Typ.Promote()))
			case sql.RangeType_LessThanOrNull:
				idxFilters = append(idxFilters,
					buildIndex(field, lo, hi, Lt),
				)
				rangeColumnExpr = expression.JoinOr(
					expression.NewLessThan(exprs[i], expression.NewLiteral(sql.GetRangeCutKey(rce.UpperBound), rce.Typ.Promote())),
					expression.NewIsNull(exprs[i]),
				)
			case sql.RangeType_LessOrEqualOrNull:
				idxFilters = append(idxFilters,
					buildIndex(field, lo, hi, LtEq),
				)
				rangeColumnExpr = expression.JoinOr(
					expression.NewLessThanOrEqual(exprs[i], expression.NewLiteral(sql.GetRangeCutKey(rce.UpperBound), rce.Typ.Promote())),
					expression.NewIsNull(exprs[i]),
				)
			case sql.RangeType_ClosedClosed:
				idxFilters = append(idxFilters,
					buildIndex(field, lo, hi, Eq),
				)
				rangeColumnExpr = expression.JoinAnd(
					expression.NewGreaterThanOrEqual(exprs[i], expression.NewLiteral(sql.GetRangeCutKey(rce.LowerBound), rce.Typ.Promote())),
					expression.NewLessThanOrEqual(exprs[i], expression.NewLiteral(sql.GetRangeCutKey(rce.UpperBound), rce.Typ.Promote())),
				)
			case sql.RangeType_OpenOpen:
				if sql.RangeCutIsBinding(rce.LowerBound) {
					idxFilters = append(idxFilters,
						buildIndex(field, lo, hi, Gt),
					)
					rangeColumnExpr = expression.JoinAnd(
						expression.NewGreaterThan(exprs[i], expression.NewLiteral(sql.GetRangeCutKey(rce.LowerBound), rce.Typ.Promote())),
						expression.NewLessThan(exprs[i], expression.NewLiteral(sql.GetRangeCutKey(rce.UpperBound), rce.Typ.Promote())),
					)
				} else {
					// Lower bound is (NULL, ...)
					rangeColumnExpr = expression.NewLessThan(exprs[i], expression.NewLiteral(sql.GetRangeCutKey(rce.UpperBound), rce.Typ.Promote()))
				}
			case sql.RangeType_OpenClosed:
				if sql.RangeCutIsBinding(rce.LowerBound) {
					rangeColumnExpr = expression.JoinAnd(
						expression.NewGreaterThan(exprs[i], expression.NewLiteral(sql.GetRangeCutKey(rce.LowerBound), rce.Typ.Promote())),
						expression.NewLessThanOrEqual(exprs[i], expression.NewLiteral(sql.GetRangeCutKey(rce.UpperBound), rce.Typ.Promote())),
					)
				} else {
					// Lower bound is (NULL, ...]
					rangeColumnExpr = expression.NewLessThanOrEqual(exprs[i], expression.NewLiteral(sql.GetRangeCutKey(rce.UpperBound), rce.Typ.Promote()))
				}
			case sql.RangeType_ClosedOpen:
				rangeColumnExpr = expression.JoinAnd(
					expression.NewGreaterThanOrEqual(exprs[i], expression.NewLiteral(sql.GetRangeCutKey(rce.LowerBound), rce.Typ.Promote())),
					expression.NewLessThan(exprs[i], expression.NewLiteral(sql.GetRangeCutKey(rce.UpperBound), rce.Typ.Promote())),
				)
			}
			rangeExpr = expression.JoinAnd(rangeExpr, rangeColumnExpr)
		}
		rangeCollectionExpr = expression.JoinOr(rangeCollectionExpr, rangeExpr)
	}
	return
}

func bounds(rce sql.RangeColumnExpr) (lo, hi any) {
	switch e := rce.LowerBound.(type) {
	case sql.Below:
		lo = e.Key
	case sql.Above:
		lo = e.Key
	}
	switch e := rce.UpperBound.(type) {
	case sql.Below:
		hi = e.Key
	case sql.Above:
		hi = e.Key
	}
	return
}

func indexedField(expr sql.Expression) v1.Column {
	f := expr.(*expression.GetField)
	return Indexed[f.Name()]
}

func (t *Table) GetIndexes(ctx *sql.Context) ([]sql.Index, error) {
	return []sql.Index{t.createIndex()}, nil
}

func (t *Table) createIndex() sql.Index {
	exprs := make([]sql.Expression, 0, len(t.schema.sql))
	exprsString := make([]string, 0, len(t.schema.sql))
	for _, column := range t.schema.sql {
		if _, ok := Indexed[column.Name]; !ok {
			continue
		}
		idx, field := t.getField(column.Name)
		ex := expression.NewGetFieldWithTable(idx, field.Type, t.name, field.Name, field.Nullable)
		exprs = append(exprs, ex)
		exprsString = append(exprsString, ex.String())
	}
	return &Index{
		TableName:  t.name,
		Exprs:      exprs,
		exprString: exprsString,
	}
}

func (t *Table) getField(col string) (int, *sql.Column) {
	i := t.schema.sql.IndexOf(col, t.name)
	if i == -1 {
		return -1, nil
	}
	return i, t.schema.sql[i]
}

type IndexedTable struct {
	*Table
	Lookup sql.IndexLookup
}

func (t *IndexedTable) LookupPartitions(ctx *sql.Context, lookup sql.IndexLookup) (sql.PartitionIter, error) {
	idxFilters, valFilters, filterExpr, err := lookup.Index.(*Index).build(ctx, lookup.Ranges...)
	if err != nil {
		return nil, err
	}
	child, err := t.Table.Partitions(ctx)
	if err != nil {
		return nil, err
	}

	return &rangePartitionIter{
		child: child.(*partitionIter),
		part: Partition{
			Expr:   filterExpr,
			Index:  idxFilters,
			Values: valFilters,
			Range:  true,
		},
	}, nil
}

func (t *Table) IndexedAccess(i sql.IndexLookup) sql.IndexedTable {
	return &IndexedTable{Table: t, Lookup: i}
}

// PartitionRows implements the sql.PartitionRows interface.
func (t *IndexedTable) PartitionRows(ctx *sql.Context, partition sql.Partition) (sql.RowIter, error) {
	iter, err := t.Table.PartitionRows(ctx, partition)
	if err != nil {
		return nil, err
	}
	return iter, nil
}

// rangePartitionIter returns a partition that has range and table data access
type rangePartitionIter struct {
	child *partitionIter
	part  Partition
}

var _ sql.PartitionIter = (*rangePartitionIter)(nil)

func (i *rangePartitionIter) Close(ctx *sql.Context) error {
	return i.child.Close(ctx)
}

func (i *rangePartitionIter) Next(ctx *sql.Context) (sql.Partition, error) {
	part, err := i.child.Next(ctx)
	if err != nil {
		return nil, err
	}
	p := part.(*Partition)
	i.part.Info.Id = p.Info.Id
	i.part.Info.Domain = p.Info.Domain
	i.part.Info.Min = p.Info.Min
	i.part.Info.Max = p.Info.Max
	rs, err := buildIndexFilter(ctx, i.part.Index, i.child.readIndex)
	if err != nil {
		if errors.Is(err, ErrSkipBlock) {
			return i.Next(ctx)
		}
		if !errors.Is(err, ErrNoFilter) {
			ctx.GetLogger().WithError(err).Warn("failed to build index filter")
		}
	} else {
		i.part.RowGroups = rs.RowGroups
		i.part.Pages = rs.Pages
	}
	return &i.part, nil
}
