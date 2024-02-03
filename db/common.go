package db

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
	v1 "github.com/vinceanalytics/staples/staples/gen/go/staples/v1"
)

func Timestamps(r arrow.Record) (lo, hi int64) {
	for i := 0; i < int(r.NumCols()); i++ {
		if r.ColumnName(i) == v1.Filters_Timestamp.String() {
			ts := r.Column(i).(*array.Int64)
			lo = ts.Value(0)
			hi = ts.Value(ts.Len() - 1)
			return
		}
	}
	return
}
