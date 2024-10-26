package aggregates

import (
	"context"
	"math"
	"time"

	"github.com/bits-and-blooms/bitset"
	"github.com/vinceanalytics/vince/internal/models"
	"github.com/vinceanalytics/vince/internal/roaring"
	"github.com/vinceanalytics/vince/internal/timeseries"
	"github.com/vinceanalytics/vince/internal/web/query"
)

type Stats struct {
	uid            *roaring.Bitmap
	Visitors       float64
	Visits         float64
	PageViews      float64
	ViewsPerVisits float64
	BounceRate     float64
	VisitDuration  float64
	Events         float64
}

func Reduce(metrics []string) func(*Stats, map[string]any) {
	ls := make([]func(*Stats, map[string]any), len(metrics))
	for i := range metrics {
		ls[i] = reduce(metrics[i])
	}
	return func(s *Stats, m map[string]any) {
		for i := range ls {
			ls[i](s, m)
		}
	}
}

func reduce(metric string) func(s *Stats, o map[string]any) {
	switch metric {
	case "visitors":
		return func(s *Stats, o map[string]any) { o[metric] = s.Visitors }
	case "visits":
		return func(s *Stats, o map[string]any) { o[metric] = s.Visits }
	case "pageviews":
		return func(s *Stats, o map[string]any) { o[metric] = s.PageViews }
	case "bounce_rate":
		return func(s *Stats, o map[string]any) { o[metric] = s.BounceRate }
	case "views_per_visit":
		return func(s *Stats, o map[string]any) { o[metric] = s.ViewsPerVisits }
	case "visit_duration":
		return func(s *Stats, o map[string]any) { o[metric] = s.VisitDuration }
	case "events":
		return func(s *Stats, o map[string]any) { o[metric] = s.Events }
	default:
		return func(s *Stats, o map[string]any) {}
	}
}

func StatToValue(metric string) func(s *Stats) float64 {
	switch metric {
	case "visitors":
		return func(s *Stats) float64 { return s.Visitors }
	case "visits":
		return func(s *Stats) float64 { return s.Visits }
	case "pageviews":
		return func(s *Stats) float64 { return s.PageViews }
	case "bounce_rate":
		return func(s *Stats) float64 { return s.BounceRate }
	case "views_per_visit":
		return func(s *Stats) float64 { return s.ViewsPerVisits }
	case "visit_duration":
		return func(s *Stats) float64 { return s.VisitDuration }
	case "events":
		return func(s *Stats) float64 { return s.Events }
	default:
		return func(s *Stats) float64 { return 0 }
	}
}

func (s *Stats) Compute() {
	if s.uid != nil {
		s.Visitors = float64(s.uid.GetCardinality())
	}
	if s.VisitDuration != 0 {
		s.VisitDuration = time.Duration(s.VisitDuration).Seconds()
	}
	s.ViewsPerVisits = s.PageViews
	if s.Visits != 0 {
		s.ViewsPerVisits /= s.Visits
		s.ViewsPerVisits = math.Round(s.ViewsPerVisits)
		s.BounceRate /= s.Visits
		s.BounceRate = math.Floor(s.BounceRate * 100)
		s.VisitDuration /= s.Visits
		s.VisitDuration = math.Floor(s.VisitDuration)
	}

	//avoid negative bounce rates
	s.BounceRate = max(s.BounceRate, 0)
}

func (d *Stats) Read(ctx context.Context, ts *timeseries.Timeseries, fields *bitset.BitSet, shard, view uint64, match *roaring.Bitmap, data timeseries.FieldsData) {
	models.EachField(fields, func(f models.Field) {
		switch f {
		case models.Field_view:
			count := data[f].True(shard, match).GetCardinality()
			d.PageViews += float64(count)
		case models.Field_session:
			count := data[f].True(shard, match).GetCardinality()
			d.Visits += float64(count)
		case models.Field_bounce:
			sum := data[f].BSISum(shard, match)
			d.BounceRate += float64(sum)
		case models.Field_duration:
			sum := data[f].BSISum(shard, match)
			d.VisitDuration += float64(sum)
		case models.Field_id:
			if d.uid == nil {
				d.uid = roaring.NewBitmap()
			}
			data[f].ExtractBSI(shard, match, func(id uint64, value int64) {
				d.uid.Set(uint64(value))
			})
		case models.Field_event:
			d.Events += float64(match.GetCardinality())
		}
		return
	})
}

func Aggregates(ctx context.Context, ts *timeseries.Timeseries, domain string, start, end time.Time, interval query.Interval, filters query.Filters, metrics []string) *Stats {
	fields := models.DataForMetrics(metrics...)
	r := new(Stats)
	ts.Select(ctx, fields, domain, start, end, interval, filters, func(shard, view uint64, columns *roaring.Bitmap, data timeseries.FieldsData) {
		r.Read(ctx, ts, fields, shard, view, columns, data)
	})
	return r
}
