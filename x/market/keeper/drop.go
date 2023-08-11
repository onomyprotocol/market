package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pendulum-labs/market/x/market/types"
)

// SetDrop set a specific drop in the store from its index
func (k Keeper) SetDrop(ctx sdk.Context, drop types.Drop) {
	store1 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropKeyPrefix))
	b := k.cdc.MustMarshal(&drop)
	store1.Set(types.DropKey(
		drop.Uid,
	), b)

	store2 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropsKeyPrefix))

	c := store2.Get(types.DropsKey(
		drop.Owner,
		drop.Pair,
	))

	var drops types.Drops

	if c == nil {
		drops = types.Drops{
			Uids: []uint64{drop.Uid},
		}
	} else {
		k.cdc.MustUnmarshal(c, &drops)
		drops.Uids = append(drops.Uids, drop.Uid)
	}

	d := k.cdc.MustMarshal(&drops)
	store2.Set(types.DropsKey(
		drop.Owner,
		drop.Pair,
	), d)
}

// GetDrop returns a drop from its index
func (k Keeper) GetDrop(
	ctx sdk.Context,
	uid uint64,
) (val types.Drop, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropKeyPrefix))

	b := store.Get(types.DropKey(
		uid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetOwnerDrops returns drops from a single owner
func (k Keeper) GetDropsOwnerPairDetail(
	ctx sdk.Context,
	owner string,
	pair string,
) (list []types.Drop) {
	store1 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropsKeyPrefix))

	b := store1.Get(types.DropsKey(
		owner,
		pair,
	))
	if b == nil {
		return list
	}

	store2 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropKeyPrefix))

	var drops types.Drops
	var drop types.Drop

	k.cdc.MustUnmarshal(b, &drops)

	for _, uid := range drops.Uids {

		b := store2.Get(types.DropKey(
			uid,
		))

		if b != nil {
			k.cdc.MustUnmarshal(b, &drop)
			list = append(list, drop)
		}
	}

	return
}

// GetOwnerDrops returns drops from a single owner
func (k Keeper) GetDropsOwnerPair(
	ctx sdk.Context,
	owner string,
	pair string,
) (drops types.Drops, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropsKeyPrefix))

	b := store.Get(types.DropsKey(
		owner,
		pair,
	))
	if b == nil {
		return drops, false
	}

	k.cdc.MustUnmarshal(b, &drops)

	return drops, true
}

// SetDrop set a specific drop in the store from its index
func (k Keeper) SetDrops(ctx sdk.Context, owner string, pair string, sum sdk.Int, uids []uint64) {
	store1 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropsKeyPrefix))
	drops := types.Drops{
		Uids: uids,
		Sum:  sum,
	}
	b := k.cdc.MustMarshal(&drops)
	store1.Set(types.DropsKey(
		owner,
		pair,
	), b)
}

// RemoveDrop removes a drop from the store
func (k Keeper) RemoveDrop(
	ctx sdk.Context,
	uid uint64,
) {
	store1 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropKeyPrefix))

	b := store1.Get(types.DropKey(
		uid,
	))

	if b == nil {
		return
	}

	var drop types.Drop

	k.cdc.MustUnmarshal(b, &drop)

	store1.Delete(types.DropKey(
		uid,
	))
}

// RemoveDrop removes a drop from the store
func (k Keeper) RemoveDropOwner(
	ctx sdk.Context,
	dropUid uint64,
	owner string,
	pair string,
) {
	// Remove uid from owner drop list
	store2 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropsKeyPrefix))

	var drops types.Drops

	c := store2.Get(types.DropsKey(
		owner,
		pair,
	))
	if c == nil {
		return
	}

	k.cdc.MustUnmarshal(c, &drops)

	var list []uint64

	for _, uid := range drops.Uids {
		if uid != dropUid {
			list = append(list, uid)
		}
	}

	drops = types.Drops{
		Uids: list,
	}

	d := k.cdc.MustMarshal(&drops)

	store2.Set(types.DropsKey(
		owner,
		pair,
	), d)
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
