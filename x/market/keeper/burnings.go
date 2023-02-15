package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/market/x/market/types"
)

// SetBurnings set a specific burnings in the store from its index
func (k Keeper) SetBurnings(ctx sdk.Context, burnings types.Burnings) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BurningsKeyPrefix))
	b := k.cdc.MustMarshal(&burnings)
	store.Set(types.BurningsKey(
		burnings.Denom,
	), b)
}

// GetBurnings returns a burnings from its index
func (k Keeper) GetBurnings(
	ctx sdk.Context,
	denom string,

) (val types.Burnings, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BurningsKeyPrefix))

	b := store.Get(types.BurningsKey(
		denom,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBurnings removes a burnings from the store
func (k Keeper) RemoveBurnings(
	ctx sdk.Context,
	denom string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BurningsKeyPrefix))
	store.Delete(types.BurningsKey(
		denom,
	))
}

// GetAllBurnings returns all burnings
func (k Keeper) GetAllBurnings(ctx sdk.Context) (list []types.Burnings) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BurningsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Burnings
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
