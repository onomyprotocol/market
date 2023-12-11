package types

import (
	"encoding/binary"
	//github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
)

var _ binary.ByteOrder

const (
	// MemberKeyPrefix is the prefix to retrieve all Member
	MemberKeyPrefix = "Member/value/"
)

// MemberKey returns the store key to retrieve a Member from the index fields
func MemberSetKey(
	denomA string,
	denomB string,
	//balance github_com_cosmos_cosmos_sdk_types.Int,
	//previous github_com_cosmos_cosmos_sdk_types.Int,
	//limit uint64,
	//stop uint64,
	//protect uint64,

) []byte {
	var key []byte

	denomABytes := []byte(denomA)
	key = append(key, denomABytes...)
	key = append(key, []byte("/")...)

	denomBBytes := []byte(denomB)
	key = append(key, denomBBytes...)
	key = append(key, []byte("/")...)
	/*
		limitBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(limitBytes, limit)
		key = append(key, limitBytes...)
		key = append(key, []byte("/")...)

		stopBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(stopBytes, stop)
		key = append(key, stopBytes...)
		key = append(key, []byte("/")...)

		protectBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(protectBytes, protect)
		key = append(key, stopBytes...)
		key = append(key, []byte("/")...)
	*/
	return key
}

// MemberKey returns the store key to retrieve a Member from the index fields
func MemberKey(
	denomA string,
	denomB string,
) []byte {
	var key []byte

	denomABytes := []byte(denomA)
	key = append(key, denomABytes...)
	key = append(key, []byte("/")...)

	denomBBytes := []byte(denomB)
	key = append(key, denomBBytes...)
	key = append(key, []byte("/")...)

	return key
}

// MemberKey returns the store key to retrieve a Member from the index fields
func MemberKeyPair(
	pair string,
	denomA string,
	denomB string,
) []byte {
	var key []byte

	pairBytes := []byte(pair)
	key = append(key, pairBytes...)
	key = append(key, []byte("/")...)

	denomABytes := []byte(denomA)
	key = append(key, denomABytes...)
	key = append(key, []byte("/")...)

	denomBBytes := []byte(denomB)
	key = append(key, denomBBytes...)
	key = append(key, []byte("/")...)
	return key
}
