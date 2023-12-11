package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PoolKeyPrefix is the prefix to retrieve all Pool
	PoolKeyPrefix = "Pool/value/"
)

// PoolKey returns the store key to retrieve a Pool
func PoolKey(
	pair string,
) []byte {
	var key []byte

	pairBytes := []byte(pair)
	key = append(key, pairBytes...)
	key = append(key, []byte("/")...)

	return key
}

// PoolSetKey returns the store key to set a Pool with the index fields
func PoolSetKey(
	pair string,

) []byte {
	var key []byte

	pairBytes := []byte(pair)
	key = append(key, pairBytes...)
	key = append(key, []byte("/")...)

	return key
}
