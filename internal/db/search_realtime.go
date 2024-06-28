package db

import (
	"context"
	"time"

	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/bufbuild/protovalidate-go"
	"github.com/gernest/rows"
	v1 "github.com/vinceanalytics/vince/gen/go/vince/v1"
	"github.com/vinceanalytics/vince/internal/defaults"
	"github.com/vinceanalytics/vince/internal/logger"
)

var validate *protovalidate.Validator

func init() {
	var err error
	validate, err = protovalidate.New(protovalidate.WithFailFast(true))
	if err != nil {
		logger.Fail("Failed setting up validator", "err", err)
	}
}

func (db *DB) Realtime(ctx context.Context, req *v1.Realtime_Request) (*v1.Realtime_Response, error) {
	defaults.Set(req)
	err := validate.Validate(req)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	firstTime := now.Add(-5 * time.Minute)
	result := new(realtimeQuery)
	err = db.Search(firstTime, now, []*v1.Filter{
		{Property: v1.Property_domain, Op: v1.Filter_equal, Value: req.SiteId},
		{Property: v1.Property_tenant_id, Op: v1.Filter_equal, Value: req.TenantId},
	}, result)
	if err != nil {
		return nil, err
	}
	return &v1.Realtime_Response{Visitors: result.Visitors()}, nil
}

type realtimeQuery struct {
	roaring64.Bitmap
	fmt ViewFmt
}

func (r *realtimeQuery) View(_ time.Time) View {
	return r
}

func (r *realtimeQuery) Apply(tx *Tx, columns *rows.Row) error {
	view := r.fmt.Format(tx.View, "uid")
	return transpose(tx.Tx, &r.Bitmap, view, tx.Shard, columns)
}

func (r *realtimeQuery) Visitors() uint64 {
	return r.GetCardinality()
}
