package ro2

import (
	"errors"
	"sync"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/vinceanalytics/vince/internal/encoding"
	"github.com/vinceanalytics/vince/internal/models"
	"github.com/vinceanalytics/vince/internal/roaring"
	"github.com/vinceanalytics/vince/internal/web/query"
)

//go:generate protoc  --go_out=. --go_opt=paths=source_relative pages.proto

type Tx struct {
	tx  *badger.Txn
	it  *badger.Iterator
	enc encoding.Encoding
}

var txPool = &sync.Pool{New: func() any {
	return &Tx{}
}}

func newTx(txn *badger.Txn) *Tx {
	tx := txPool.Get().(*Tx)
	tx.tx = txn
	return tx
}

func (tx *Tx) Iter() *badger.Iterator {
	if tx.it == nil {
		tx.it = tx.tx.NewIterator(badger.IteratorOptions{})
	}
	return tx.it
}

func (tx *Tx) Close() {
	if tx.it != nil {
		tx.it.Close()
	}
	tx.it = nil
}

func (tx *Tx) Release() {
	tx.Close()
	tx.tx = nil
	tx.enc.Reset()
}

func (tx *Tx) Select(domain string, start,
	end time.Time, intrerval query.Interval, filters query.Filters, cb func(shard, view uint64, columns *roaring.Bitmap) error) error {
	shard, ok := tx.ID(models.Field_domain, []byte(domain))
	if !ok {
		return nil
	}
	m := tx.compile(filters)
	return intrerval.Range(start, end, func(t time.Time) error {
		view := uint64(t.UnixMilli())
		match, err := tx.Domain(shard, view)
		if err != nil {
			return err
		}
		if match.IsEmpty() {
			return nil
		}
		columns := m.Apply(tx, shard, view, match)
		if columns.IsEmpty() {
			return nil
		}
		return cb(shard, view, columns)
	})
}

func (tx *Tx) Sum(shard, view uint64, field models.Field, match *roaring.Bitmap) (sum int64, err error) {
	err = tx.Bitmap(shard, view, field, func(bs *roaring.BSI) {
		sum, _ = bs.Sum(match)
	})
	return
}

func (tx *Tx) Unique(shard, view uint64, field models.Field, match *roaring.Bitmap) (uint64, error) {
	bs, err := tx.Transpose(shard, view, field, match)
	if err != nil {
		return 0, err
	}
	return uint64(bs.GetCardinality()), nil
}

func (tx *Tx) Transpose(shard, view uint64, field models.Field, match *roaring.Bitmap) (result *roaring.Bitmap, err error) {
	err = tx.Bitmap(shard, view, field, func(bs *roaring.BSI) {
		tr := bs.Extract(match)
		result = roaring.NewBitmap()
		for _, v := range tr {
			result.Set(uint64(v))
		}
	})
	return
}

func (tx *Tx) TransposeSet(shard, view uint64, field models.Field, match *roaring.Bitmap) (result map[int64][]uint64, err error) {
	err = tx.Bitmap(shard, view, field, func(bs *roaring.BSI) {
		tr := bs.Extract(match)
		result = make(map[int64][]uint64)
		for k, v := range tr {
			result[v] = append(result[v], k)
		}
	})
	return
}

func (tx *Tx) Count(shard, view uint64, field models.Field, match *roaring.Bitmap) (count uint64, err error) {
	err = tx.Bitmap(shard, view, field, func(bs *roaring.BSI) {
		count = bs.GetExistenceBitmap().AndCardinality(match)
	})
	if err != nil {
		return 0, err
	}
	return
}

func (tx *Tx) Bitmap(shard, view uint64, field models.Field, f func(bs *roaring.BSI)) error {
	key := tx.enc.Key(encoding.Key{
		Time:  view,
		Shard: uint32(shard),
		Field: field,
	})
	it, err := tx.tx.Get(key)
	if err != nil {
		if errors.Is(err, badger.ErrKeyNotFound) {
			return nil
		}
		return err
	}

	return it.Value(func(val []byte) error {
		f(roaring.NewBSIFromBuffer(val))
		return nil
	})
}

func (tx *Tx) Domain(shard, view uint64) (*roaring.Bitmap, error) {
	key := tx.enc.Key(encoding.Key{
		Time:  view,
		Shard: uint32(shard),
		Field: models.Field_domain,
	})
	it, err := tx.tx.Get(key)
	if err != nil {
		if errors.Is(err, badger.ErrKeyNotFound) {
			return roaring.NewBitmap(), nil
		}
		return nil, err
	}

	value, err := it.ValueCopy(nil)
	if err != nil {
		return nil, err
	}
	return roaring.FromBuffer(value), nil
}
