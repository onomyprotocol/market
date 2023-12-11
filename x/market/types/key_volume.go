package types

import (
	"encoding/binary"
	//github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
)

var _ binary.ByteOrder

const (
	// MemberKeyPrefix is the prefix to retrieve all Volumes
	VolumeKeyPrefix = "Volume/value/"
)

// MemberKey returns the store key to retrieve a Volume
func VolumeKey(
	denom string,
) []byte {
	var key []byte

	denomBytes := []byte(denom)
	key = append(key, denomBytes...)
	key = append(key, []byte("/")...)

	return key
}
