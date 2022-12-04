package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// OrderKeyPrefix is the prefix to retrieve all Order
	OrderKeyPrefix = "Order/value/"
)

// OrderKey returns the store key to retrieve a Order from the index fields
func OrderSetKey(
	uid uint64,
	owner string,
	active bool,
	orderType string,
	denomAsk string,
	denomBid string,
) []byte {
	var key []byte

	uidBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(uidBytes, uid)
	key = append(key, uidBytes...)
	key = append(key, []byte("/")...)

	ownerBytes := []byte(owner)
	key = append(key, ownerBytes...)
	key = append(key, []byte("/")...)

	activeBytes := []byte{0}
	if active {
		activeBytes = []byte{1}
	}
	key = append(key, activeBytes...)
	key = append(key, []byte("/")...)

	orderTypeBytes := []byte(orderType)
	key = append(key, orderTypeBytes...)
	key = append(key, []byte("/")...)

	denomAskBytes := []byte(denomAsk)
	key = append(key, denomAskBytes...)
	key = append(key, []byte("/")...)

	denomBidBytes := []byte(denomBid)
	key = append(key, denomBidBytes...)
	key = append(key, []byte("/")...)

	return key
}

func OrderGetKey(
	uid uint64,
) []byte {
	var key []byte

	uidBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(uidBytes, uid)
	key = append(key, uidBytes...)
	key = append(key, []byte("/")...)

	return key
}
