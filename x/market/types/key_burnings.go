package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// BurningsKeyPrefix is the prefix to retrieve all Burnings
	BurningsKeyPrefix = "Burnings/value/"
)

// BurningsKey returns the store key to retrieve a Burnings from the index fields
func BurningsKey(
	denom string,
) []byte {
	var key []byte

	denomBytes := []byte(denom)
	key = append(key, denomBytes...)
	key = append(key, []byte("/")...)

	return key
}
