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
	denom1 string,
	denom2 string,
	leader string,
) []byte {
	var key []byte

	pairBytes := []byte(pair)
	key = append(key, pairBytes...)
	key = append(key, []byte("/")...)

	denom1Bytes := []byte(denom1)
	key = append(key, denom1Bytes...)
	key = append(key, []byte("/")...)

	denom2Bytes := []byte(denom2)
	key = append(key, denom2Bytes...)
	key = append(key, []byte("/")...)

	leaderBytes := []byte(leader)
	key = append(key, leaderBytes...)
	key = append(key, []byte("/")...)

	return key
}
