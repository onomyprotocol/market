package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DropKeyPrefix is the prefix to retrieve all Drop
	DropKeyPrefix = "Drop/value/"
	// DropsKeyPrefix is the prefix to retrieve all Owner of Drops
	DropsKeyPrefix = "Drop/Owner/"
	// DropsSumKeyPrefix is the prefix to retrieve all Owner of Drops
	DropsSumKeyPrefix = "Drop/Sum/"
)

// DropKey returns the store key to retrieve a Drop from the index fields
func DropKey(
	uid uint64,
) []byte {
	var key []byte

	uidBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(uidBytes, uid)
	key = append(key, uidBytes...)
	key = append(key, []byte("/")...)

	return key
}

// DropKey returns the store key to retrieve a Drop from the index fields
func DropsKey(
	owner string,
	pair string,
) []byte {
	var key []byte

	ownerBytes := []byte(owner)
	pairBytes := []byte(pair)
	key = append(key, ownerBytes...)
	key = append(key, pairBytes...)
	key = append(key, []byte("/")...)

	return key
}

// DropKey returns the store key to retrieve a Drop from the index fields
func DropsSumKey(
	owner string,
	pair string,
) []byte {
	var key []byte

	ownerBytes := []byte(owner)
	pairBytes := []byte(pair)
	key = append(key, ownerBytes...)
	key = append(key, pairBytes...)
	key = append(key, []byte("/")...)

	return key
}
