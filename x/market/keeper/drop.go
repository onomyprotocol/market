package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pendulum-labs/market/x/market/types"
)

// SetDrop set a specific drop in the store from its index
func (k Keeper) SetDrop(ctx sdk.Context, drop types.Drop) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropKeyPrefix))
	a := k.cdc.MustMarshal(&drop)
	store.Set(types.DropKey(
		drop.Uid,
	), a)
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

// GetDrop returns a drop from its index
func (k Keeper) GetDropPairs(
	ctx sdk.Context,
	address string,
) (val types.DropPairs, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropPairsKeyPrefix))

	a := store.Get(types.DropPairsKey(
		address,
	))
	if a == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(a, &val)
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

	store1.Delete(types.DropKey(
		uid,
	))
}

// SetDrop set a specific drop in the store from its index
func (k Keeper) SetDropOwner(
	ctx sdk.Context,
	drop types.Drop,
) {
	store1 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropsKeyPrefix))

	var drops types.Drops

	a := store1.Get(types.DropsKey(
		drop.Owner,
		drop.Pair,
	))
	if a == nil {
		drops.Sum = drop.Drops
		drops.Uids = []uint64{drop.Uid}
	} else {
		k.cdc.MustUnmarshal(a, &drops)

		uids, _ := addUid(drops.Uids, drop.Uid)

		drops = types.Drops{
			Uids: uids,
			Sum:  drops.Sum.Add(drop.Drops),
		}
	}

	b := k.cdc.MustMarshal(&drops)

	store1.Set(types.DropsKey(
		drop.Owner,
		drop.Pair,
	), b)

	// Add drop pair to owner
	store2 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropPairsKeyPrefix))

	var dropPairs types.DropPairs

	c := store2.Get(types.DropPairsKey(
		drop.Owner,
	))

	if c == nil {
		dropPairs.Pairs = []string{drop.Pair}
	} else {
		k.cdc.MustUnmarshal(c, &dropPairs)
		dropPairs.Pairs, _ = addPair(dropPairs.Pairs, drop.Pair)
	}

	d := k.cdc.MustMarshal(&dropPairs)

	store2.Set(types.DropPairsKey(
		drop.Owner,
	), d)
}

// RemoveDrop removes a drop from the store
func (k Keeper) RemoveDropOwner(
	ctx sdk.Context,
	drop types.Drop,
) {
	// Remove uid from owner drop list
	store1 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropsKeyPrefix))

	var drops types.Drops

	a := store1.Get(types.DropsKey(
		drop.Owner,
		drop.Pair,
	))
	if a == nil {
		return
	}

	k.cdc.MustUnmarshal(a, &drops)

	drops.Uids, _ = removeUid(drops.Uids, drop.Uid)
	drops.Sum = drops.Sum.Sub(drop.Drops)

	b := k.cdc.MustMarshal(&drops)

	store1.Set(types.DropsKey(
		drop.Owner,
		drop.Pair,
	), b)

	if !(len(drops.Uids) > 0) {
		// Remove uid from owner drop list
		store2 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropPairsKeyPrefix))

		var dropPairs types.DropPairs

		c := store2.Get(types.DropPairsKey(
			drop.Owner,
		))

		k.cdc.MustUnmarshal(c, &dropPairs)

		dropPairs.Pairs, _ = removePair(dropPairs.Pairs, drop.Pair)

		d := k.cdc.MustMarshal(&dropPairs)

		store2.Set(types.DropPairsKey(
			drop.Owner,
		), d)
	}
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

// GetOwnerDrops returns drops from a single owner
func (k Keeper) GetPairs(
	ctx sdk.Context,
	owner string,
) (pairs types.DropPairs, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropPairsKeyPrefix))

	b := store.Get(types.DropPairsKey(
		owner,
	))
	if b == nil {
		return pairs, false
	}

	k.cdc.MustUnmarshal(b, &pairs)

	return pairs, true
}
