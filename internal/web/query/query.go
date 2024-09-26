package query

import (
	"encoding/json"
	"net/url"
	"time"
)

type Query struct {
	period Period
	cmp    *Period
	filter Filters
	metric string
	all    bool
}

func New(u url.Values) *Query {
	var fs Filters
	json.Unmarshal([]byte(u.Get("filters")), &fs)

	period := period(u.Get("period"), u.Get("date"))
	if i := u.Get("interval"); i != "" {
		switch i {
		case "minute":
			period.Interval = Minute
		case "hour":
			period.Interval = Hour
		case "date":
			period.Interval = Date
		case "week":
			period.Interval = Week
		case "month":
			period.Interval = Month
		}
	}

	var cmp *Period
	switch u.Get("period") {
	case "all", "realtime":
	default:
		now := time.Now().UTC()
		switch u.Get("comparison") {
		case "previous_period":
			diff := period.End.Sub(period.Start)
			cmp = &Period{Start: period.Start.Add(-diff), End: period.End.Add(-diff)}
		case "year_over_year":
			start := period.Start.AddDate(-1, 0, 0)
			end := earliest(period.End, now).AddDate(-1, 0, 0)
			cmp = &Period{Start: start, End: end}
		case "custom":
		}
	}

	return &Query{
		period: period,
		cmp:    cmp,
		metric: u.Get("metric"),
		all:    u.Get("period") == "all",
	}
}

func (q *Query) All() bool { return q.all }

func earliest(a, b time.Time) time.Time {
	if a.Before(b) {
		return a
	}
	return b
}

func (q *Query) Start() time.Time   { return q.period.Start }
func (q *Query) From() string       { return q.period.Start.Format(time.DateOnly) }
func (q *Query) End() time.Time     { return q.period.End }
func (q *Query) To() string         { return q.period.End.Format(time.DateOnly) }
func (q *Query) Interval() Interval { return q.period.Interval }
func (q *Query) Filter() Filters    { return q.filter }
func (q *Query) Metric() string     { return q.metric }
func (q *Query) Compare() *Period   { return q.cmp }
