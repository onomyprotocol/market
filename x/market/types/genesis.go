package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PoolList:     []Pool{},
		DropList:     []Drop{},
		MemberList:   []Member{},
		BurningsList: []Burnings{},
		OrderList:    []Order{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in pool
	poolIndexMap := make(map[string]struct{})

	for _, elem := range gs.PoolList {
		index := string(PoolKey(elem.Pair))
		if _, ok := poolIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for pool")
		}
		poolIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in drop
	dropIndexMap := make(map[string]struct{})

	for _, elem := range gs.DropList {
		index := string(DropKey(elem.Uid))
		if _, ok := dropIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for drop")
		}
		dropIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in member
	memberIndexMap := make(map[string]struct{})

	for _, elem := range gs.MemberList {
		index := string(MemberKey(elem.DenomA, elem.DenomB))
		if _, ok := memberIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for member")
		}
		memberIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in burnings
	burningsIndexMap := make(map[string]struct{})

	for _, elem := range gs.BurningsList {
		index := string(BurningsKey(elem.Denom))
		if _, ok := burningsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for burnings")
		}
		burningsIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in order
	orderIndexMap := make(map[string]struct{})

	for _, elem := range gs.OrderList {
		index := string(OrderKey(elem.Uid))
		if _, ok := orderIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for order")
		}
		orderIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
