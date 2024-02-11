package stats

import (
	"net/http"
	"slices"
	"time"

	"github.com/RoaringBitmap/roaring"
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
	"github.com/apache/arrow/go/v15/arrow/compute"
	"github.com/vinceanalytics/vince/filters"
	v1 "github.com/vinceanalytics/vince/gen/go/staples/v1"
	"github.com/vinceanalytics/vince/logger"
	"github.com/vinceanalytics/vince/request"
	"github.com/vinceanalytics/vince/session"
)

func BreakDown(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := r.URL.Query()
	req := v1.BreakDown_Request{
		SiteId:  query.Get("site_id"),
		Period:  ParsePeriod(ctx, query),
		Metrics: ParseMetrics(ctx, query),
		Filters: ParseFilters(ctx, query),
	}
	if !request.Validate(ctx, w, &req) {
		return
	}
	period := req.Period
	if period == nil {
		period = &v1.TimePeriod{
			Value: &v1.TimePeriod_Base_{
				Base: v1.TimePeriod__30d,
			},
		}
	}
	filter := &v1.Filters{
		List: append(req.Filters, &v1.Filter{
			Property: v1.Property_domain,
			Op:       v1.Filter_equal,
			Value:    req.SiteId,
		}),
	}
	slices.Sort(req.Metrics)
	slices.Sort(req.Property)
	selectedColumns := metricsToProjection(filter, req.Metrics, req.Property...)
	from, to := PeriodToRange(ctx, time.Now, period, r.URL.Query())
	scannedRecord, err := session.Get(ctx).Scan(ctx, from.UnixMilli(), to.UnixMilli(), filter)
	if err != nil {
		logger.Get(ctx).Error("Failed scanning", "err", err)
		request.Internal(ctx, w)
		return
	}
	defer scannedRecord.Release()
	mapping := map[string]arrow.Array{}
	for i := 0; i < int(scannedRecord.NumCols()); i++ {
		mapping[scannedRecord.ColumnName(i)] = scannedRecord.Column(i)
	}
	defer clear(mapping)
	// build key mapping
	b := array.NewUint32Builder(compute.GetAllocator(ctx))
	defer b.Release()
	var result []*v1.BreakDown_Response_Result
	// TODO: run this concurrently
	xc := &Compute{
		mapping: make(map[string]arrow.Array),
	}
	defer xc.Release()

	for _, prop := range req.Property {
		var groups []*v1.BreakDown_Response_Group
		for key, bitmap := range hashProp(mapping[filters.Column(prop)]) {
			b.AppendValues(bitmap.ToArray(), nil)
			idx := b.NewUint32Array()

			clear(xc.mapping)
			xc.view = nil
			xc.visit = nil

			for _, name := range selectedColumns {
				a, err := compute.TakeArray(ctx, mapping[name], idx)
				if err != nil {
					idx.Release()
					logger.Get(ctx).Error("Failed taking array for breaking down", "column", name)
					request.Internal(ctx, w)
					return
				}
				xc.mapping[name] = a
			}

			var values []*v1.Value
			for _, metric := range req.Metrics {
				value, err := xc.Metric(ctx, metric)
				if err != nil {
					logger.Get(ctx).Error("Failed computing metric", "metric", metric)
					request.Internal(ctx, w)
					return
				}
				values = append(values, &v1.Value{
					Metric: metric,
					Value:  value,
				})
			}
			groups = append(groups, &v1.BreakDown_Response_Group{
				Key:    key,
				Values: values,
			})
			idx.Release()
		}
		result = append(result, &v1.BreakDown_Response_Result{
			Property: prop,
			Groups:   groups,
		})
	}
	request.Write(ctx, w, &v1.BreakDown_Response{Results: result})
}

func hashProp(a arrow.Array) map[string]*roaring.Bitmap {
	o := make(map[string]*roaring.Bitmap)
	d := a.(*array.Dictionary)
	s := d.Dictionary().(*array.String)
	for i := 0; i < a.Len(); i++ {
		if d.IsNull(i) {
			continue
		}
		x := s.Value(d.GetValueIndex(i))
		b, ok := o[x]
		if !ok {
			b = new(roaring.Bitmap)
			o[x] = b
		}
		b.Add(uint32(i))
	}
	return o
}
