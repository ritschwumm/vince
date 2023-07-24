package neo

import (
	"io"
	"time"

	"github.com/parquet-go/parquet-go"
)

type Sys struct {
	Timestamp time.Time         `parquet:"timestamp,timestamp,zstd"`
	Labels    map[string]string `parquet:"labels" parquet-key:",dict,zstd" parquet-value:",dict,zstd"`
	Name      string            `parquet:"name,zstd"`
	Value     float64           `parquet:"value,zstd"`
}

func SysWriter(w io.Writer) *parquet.GenericWriter[Sys] {
	return parquet.NewGenericWriter[Sys](w,
		parquet.BloomFilters(
			parquet.SplitBlockFilter(FilterBitsPerValue, "labels", "key_value", "key"),
			parquet.SplitBlockFilter(FilterBitsPerValue, "labels", "key_value", "value"),
			parquet.SplitBlockFilter(FilterBitsPerValue, "name"),
		),
	)
}
