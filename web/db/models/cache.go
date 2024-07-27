package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gernest/len64/web/db/schema"
	"golang.org/x/time/rate"
	"gorm.io/gorm"
)

type CachedSite struct {
	ID              uint64
	Domain          string
	StatsStartDate  time.Time
	IngestRateLimit sql.NullFloat64
	UserID          uint64
}

func CacheRateLimit(c *CachedSite) (rate.Limit, int) {
	if !c.IngestRateLimit.Valid {
		return rate.Inf, 0
	}
	return rate.Limit(c.IngestRateLimit.Float64), 10
}

func QuerySitesToCache(db *gorm.DB, fn func(*CachedSite)) (uint64, error) {
	rows, err := db.Model(&schema.Site{}).Select("sites.id, sites.domain,sites.stats_start_date, sites.ingest_rate_limit,site_memberships.user_id").
		Joins("left join  site_memberships on sites.id = site_memberships.site_id and site_memberships.role = 'owner' ").
		Rows()
	if err != nil {
		return 0, fmt.Errorf("loading sites %w", err)
	}
	var site CachedSite
	var count uint64
	for rows.Next() {
		db.ScanRows(rows, &site)
		fn(&site)
		count += 1
	}
	return count, nil
}
