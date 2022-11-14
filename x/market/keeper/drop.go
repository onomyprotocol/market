package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/market/x/market/types"
)

// SetDrop set a specific drop in the store from its index
func (k Keeper) SetDrop(ctx sdk.Context, drop types.Drop) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropKeyPrefix))
	b := k.cdc.MustMarshal(&drop)
	store.Set(types.DropKey(
		drop.Uid,
		drop.Owner,
		drop.Pair,
	), b)
}

// GetDrop returns a drop from its index
func (k Keeper) GetDrop(
	ctx sdk.Context,
	uid uint64,
	owner string,
	pair string,

) (val types.Drop, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropKeyPrefix))

	b := store.Get(types.DropKey(
		uid,
		owner,
		pair,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDrop removes a drop from the store
func (k Keeper) RemoveDrop(
	ctx sdk.Context,
	uid uint64,
	owner string,
	pair string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropKeyPrefix))
	store.Delete(types.DropKey(
		uid,
		owner,
		pair,
	))
}

// GetAllDrop returns all drop
func (k Keeper) GetAllDrop(ctx sdk.Context) (list []types.Drop) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Drop
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
