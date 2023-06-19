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
	store1.Set(types.DropSetKey(
		drop.Uid,
	), b)

	store2 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropOwnerKeyPrefix))

	c := store2.Get(types.DropOwnerKey(
		drop.Owner,
	))

	var ownedDrops types.DropOwner

	if c == nil {
		ownedDrops = types.DropOwner{
			Owner: drop.Owner,
			Uids:  []uint64{drop.Uid},
		}
	} else {
		k.cdc.MustUnmarshal(c, &ownedDrops)
		ownedDrops.Uids = append(ownedDrops.Uids, drop.Uid)
	}

	d := k.cdc.MustMarshal(&ownedDrops)
	store2.Set(types.DropOwnerKey(
		drop.Owner,
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
func (k Keeper) GetOwnerDrops(
	ctx sdk.Context,
	owner string,
) (list []types.Drop) {
	store1 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropOwnerKeyPrefix))

	b := store1.Get(types.DropOwnerKey(
		owner,
	))
	if b == nil {
		return list
	}

	store2 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropKeyPrefix))

	var ownedDrops types.DropOwner
	var drop types.Drop

	k.cdc.MustUnmarshal(b, &ownedDrops)

	for _, uid := range ownedDrops.Uids {

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
func (k Keeper) GetOwnerDropsInt(
	ctx sdk.Context,
	owner string,
) (drops sdk.Int) {
	store1 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropOwnerKeyPrefix))

	b := store1.Get(types.DropOwnerKey(
		owner,
	))
	if b == nil {
		return sdk.NewInt(0)
	}

	store2 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropKeyPrefix))

	var ownedDrops types.DropOwner
	var drop types.Drop

	k.cdc.MustUnmarshal(b, &ownedDrops)

	for _, uid := range ownedDrops.Uids {

		b := store2.Get(types.DropKey(
			uid,
		))

		if b != nil {
			k.cdc.MustUnmarshal(b, &drop)
			drops = drops.Add(drop.Drops)
		}
	}

	return
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

	// Remove uid from owner drop list
	store2 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropOwnerKeyPrefix))

	var ownedDrops types.DropOwner

	c := store2.Get(types.DropOwnerKey(
		drop.Owner,
	))
	if c == nil {
		return
	}

	k.cdc.MustUnmarshal(c, &ownedDrops)

	var list []uint64

	for _, uid := range ownedDrops.Uids {

		if uid != drop.Uid {
			list = append(list, uid)
		}
	}

	ownedDrops = types.DropOwner{
		Owner: drop.Owner,
		Uids:  list,
	}

	d := k.cdc.MustMarshal(&ownedDrops)

	store2.Set(types.DropOwnerKey(
		drop.Owner,
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
