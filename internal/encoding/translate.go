package encoding

import (
	"encoding/binary"

	"github.com/vinceanalytics/vince/internal/keys"
	"github.com/vinceanalytics/vince/internal/models"
)

func TranslateKey(field models.Field, value []byte) []byte {
	o := make([]byte, 2+1+len(value))
	copy(o, keys.TranslateKeyPrefix)
	o[2] = byte(field)
	copy(o[3:], value)
	return o
}

func TranslateID(field models.Field, id uint64) []byte {
	o := make([]byte, 2+1+8)
	copy(o, keys.TranslateIDPrefix)
	o[2] = byte(field)
	binary.BigEndian.PutUint64(o[3:], id)
	return o
}
