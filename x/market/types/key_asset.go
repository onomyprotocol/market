package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AssetKeyPrefix is the prefix to retrieve all Asset
	AssetKeyPrefix = "Asset/value/"
)

// AssetKey returns the store key to retrieve a Asset from the index fields
func AssetKey(
	active bool,
	owner string,
	assetType string,
) []byte {
	var key []byte

	activeBytes := []byte{0}
	if active {
		activeBytes = []byte{1}
	}
	key = append(key, activeBytes...)
	key = append(key, []byte("/")...)

	ownerBytes := []byte(owner)
	key = append(key, ownerBytes...)
	key = append(key, []byte("/")...)

	assetTypeBytes := []byte(assetType)
	key = append(key, assetTypeBytes...)
	key = append(key, []byte("/")...)

	return key
}
