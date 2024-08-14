package ro2

import (
	"bytes"
	"errors"

	"github.com/dgraph-io/badger/v4"
	"github.com/vinceanalytics/vince/internal/ro"
	"github.com/vinceanalytics/vince/internal/roaring"
	"github.com/vinceanalytics/vince/internal/roaring/roaring64"
)

//go:generate protoc  --go_out=. --go_opt=paths=source_relative pages.proto

type Tx struct {
	tx   *badger.Txn
	keys Keys
}

func (tx *Tx) release() {
	tx.keys.Release()
}

func (tx *Tx) Add(date, shard, field uint64, keys []uint32, values []string, r *roaring64.Bitmap) error {
	return r.Save(&txWrite{
		date:  date,
		shard: shard,
		field: field,
		tx:    tx,
	})
}

type txWrite struct {
	date, shard, field uint64
	tx                 *Tx
}

var _ roaring64.Context = (*txWrite)(nil)

func (t *txWrite) Value(key uint32, cKey uint16, value func(uint8, []byte) error) error {
	xk := t.tx.keys.Get().
		SetDate(t.date).
		SetShard(t.shard).
		SetField(t.field).
		SetKey(key).
		SetContainer(cKey)
	it, err := t.tx.tx.Get(xk[:])
	if err != nil {
		if !errors.Is(err, badger.ErrKeyNotFound) {
			return err
		}
		return nil
	}
	return it.Value(func(val []byte) error {
		return value(it.UserMeta(), val)
	})
}

func (t *txWrite) Write(key uint32, cKey uint16, typ uint8, value []byte) error {
	xk := t.tx.keys.Get().
		SetDate(t.date).
		SetKind(ROAR).
		SetShard(t.shard).
		SetField(t.field).
		SetKey(key).
		SetContainer(cKey)
	return t.tx.tx.SetEntry(badger.NewEntry(xk[:], value).WithMeta(typ))
}

// we explicitly accept bit depth because city is a uint32 and we stole bounce
// as -1,0,1 meaning we can control how much we iterate for valid bits.
func (tx *Tx) ExtractBSI(date, shard, field uint64, bitDepth uint64, match *roaring64.Bitmap, f func(row uint64, c int64)) {
	exists := tx.Row(date, shard, field, 0)
	exists.And(match)
	if exists.IsEmpty() {
		return
	}
	data := make(map[uint64]uint64)
	mergeBits(exists, 0, data)

	sign := tx.Row(date, shard, field, 1)
	mergeBits(sign, 1<<63, data)

	for i := uint64(0); i < bitDepth; i++ {
		bits := tx.Row(date, shard, field, 2+uint64(i))
		if bits.IsEmpty() {
			continue
		}
		bits.And(exists)
		if bits.IsEmpty() {
			continue
		}
		mergeBits(bits, 1<<i, data)
	}
	for columnID, val := range data {
		val = uint64((2*(int64(val)>>63) + 1) * int64(val&^(1<<63)))
		f(columnID, int64(val))
	}
}

func mergeBits(bits *roaring64.Bitmap, mask uint64, out map[uint64]uint64) {
	it := bits.Iterator()
	for it.HasNext() {
		out[it.Next()] |= mask
	}
}

func (tx *Tx) ExtractMutex(date, shard, field uint64, match *roaring64.Bitmap, f func(row uint64, c *roaring.Container)) {
	filter := make([]*roaring.Container, 1<<ro.ShardVsContainerExponent)
	match.Each(func(key uint32, cKey uint16, value *roaring.Container) error {
		if value.IsEmpty() {
			return nil
		}
		idx := cKey % (1 << ro.ShardVsContainerExponent)
		filter[idx] = value
		return nil
	})
	opts := badger.IteratorOptions{
		Prefix: tx.keys.Get().
			SetDate(date).
			SetShard(shard).
			SetField(field).
			FieldPrefix(),
	}
	itr := tx.tx.NewIterator(opts)
	defer itr.Close()
	prevRow := ^uint64(0)
	seenThisRow := false
	var ac roaring.Container
	for itr.Seek(nil); itr.Valid(); itr.Next() {
		item := itr.Item()
		k := ReadKey(item.Key())
		row := uint64(k) >> ro.ShardVsContainerExponent
		if row == prevRow {
			if seenThisRow {
				continue
			}
		} else {
			seenThisRow = false
			prevRow = row
		}
		idx := k % (1 << ro.ShardVsContainerExponent)
		if filter[idx] == nil {
			continue
		}
		item.Value(func(val []byte) error {
			return ac.From(item.UserMeta(), val)
		})
		if ac.Intersects(filter[idx]) {
			f(row, &ac)
		}
	}
}

func (tx *Tx) Row(date, shard, field uint64, rowID uint64) *roaring64.Bitmap {
	from, to := tx.keyRange(date, shard, field, rowID)
	opts := badger.IteratorOptions{
		Prefix: from.KeyPrefix(),
	}
	itr := tx.tx.NewIterator(opts)
	defer itr.Close()
	valid := func() bool {
		return itr.Valid() &&
			bytes.Compare(itr.Item().Key(), to[:]) == -1
	}
	b := roaring.New()
	for itr.Seek(from[:]); valid(); itr.Next() {
		item := itr.Item()
		item.Value(func(val []byte) error {
			b.FromWire(0, item.UserMeta(), val)
			return nil
		})
	}
	return roaring64.NewFromMap(b)
}

func (tx *Tx) keyRange(date, shard, field uint64, rowID uint64) (from, to *Key) {
	b := roaring64.New()
	b.Add(rowID * ro.ShardWidth)
	b.Add((rowID + 1) * ro.ShardWidth)
	b.Each(func(key uint32, cKey uint16, value *roaring.Container) error {
		k := tx.keys.Get().SetDate(date).SetShard(shard).
			SetField(field).SetKey(key).SetContainer(cKey)
		if from == nil {
			from = k
			return nil
		}
		to = k
		return nil
	})
	return
}
