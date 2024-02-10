package cold

import (
	"os"
	"testing"

	"github.com/apache/arrow/go/v15/arrow/array"
	"github.com/apache/arrow/go/v15/arrow/memory"
	"github.com/stretchr/testify/require"
	"github.com/vinceanalytics/vince/staples"
)

func TestIndex(t *testing.T) {
	schema := staples.NewArrow[staples.Event](memory.NewGoAllocator()).NewRecord().Schema()
	f, err := os.Open("testdata/sample.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	r, _, err := array.RecordFromJSON(memory.NewGoAllocator(), schema, f)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Release()

	seg, err := New(r)
	require.NoError(t, err)
	_ = seg
}
