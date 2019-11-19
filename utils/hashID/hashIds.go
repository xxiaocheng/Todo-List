package hashID

import (
	"strconv"
	"todoList/config"

	"github.com/speps/go-hashids"
)

// inital the `hashids`
func initHashids() (*hashids.HashID, error) {
	hd := hashids.NewData()
	hd.Salt = config.Config.Salt
	hd.MinLength = config.Config.MinLength
	h, error := hashids.NewWithData(hd)
	return h, error
}

// Encode `id` to hash
func EncodeIDToHash(id uint) string {
	h, _ := initHashids()
	e, _ := h.EncodeHex(strconv.FormatUint(uint64(id), 10))
	return e
}

// Decode hash to `id`
func DecodeHashToID(hash string) uint {
	h, _ := initHashids()
	e, _ := h.DecodeHex(hash)
	id, _ := strconv.ParseUint(e, 10, 32)
	return uint(id)
}
