package keeper

import (
	"math/big"
	"sort"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pendulum-labs/market/x/market/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

// GetOrderOwner returns orders from a single owner
func (k Keeper) GetDropOwnerPair(
	ctx sdk.Context,
	owner string,
	pair string,
) (list []types.Drop, found bool) {
	store1 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropsKeyPrefix))

	a := store1.Get(types.DropsKey(
		owner,
		pair,
	))
	if a == nil {
		return list, false
	}

	var drops types.Drops

	k.cdc.MustUnmarshal(a, &drops)

	store2 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropKeyPrefix))

	for _, uid := range drops.Uids {
		var drop types.Drop

		b := store2.Get(types.DropKey(
			uid,
		))

		if b != nil {
			k.cdc.MustUnmarshal(b, &drop)
			list = append(list, drop)
		}
	}

	return list, true
}

// GetOrderOwner returns orders from a single owner
func (k Keeper) GetDropAmounts(
	ctx sdk.Context,
	uid uint64,
) (denom1 string, denom2 string, amount1 sdk.Int, amount2 sdk.Int, found bool) {
	dropStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropKeyPrefix))

	a := dropStore.Get(types.DropKey(
		uid,
	))
	if a == nil {
		return denom1, denom2, amount1, amount2, false
	}

	var drop types.Drop
	k.cdc.MustUnmarshal(a, &drop)

	pair := strings.Split(drop.Pair, ",")

	denom1 = pair[0]
	denom2 = pair[1]

	memberStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MemberKeyPrefix))

	b := memberStore.Get(types.MemberKey(
		denom2,
		denom1,
	))
	if b == nil {
		return denom1, denom2, amount1, amount2, false
	}

	var member1 types.Member
	k.cdc.MustUnmarshal(b, &member1)

	c := memberStore.Get(types.MemberKey(
		denom1,
		denom2,
	))
	if c == nil {
		return denom1, denom2, amount1, amount2, false
	}

	var member2 types.Member
	k.cdc.MustUnmarshal(c, &member2)

	poolStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKeyPrefix))

	d := poolStore.Get(types.PoolKey(
		drop.Pair,
	))
	if d == nil {
		return denom1, denom2, amount1, amount2, false
	}

	var pool types.Pool
	k.cdc.MustUnmarshal(d, &pool)

	amount1, amount2, error := dropAmounts(drop.Drops, pool, member1, member2)
	if error != nil {
		return denom1, denom2, amount1, amount2, false
	}

	found = true

	return
}

func dropAmounts(drops sdk.Int, pool types.Pool, member1 types.Member, member2 types.Member) (sdk.Int, sdk.Int, error) {
	// see `msg_server_redeem_drop` for our bigint strategy
	// `dropAmtMember1 = (drops * member1.Balance) / pool.Drops`
	tmp := big.NewInt(0)
	tmp.Mul(drops.BigInt(), member1.Balance.BigInt())
	tmp.Quo(tmp, pool.Drops.BigInt())
	dropAmtMember1 := sdk.NewIntFromBigInt(tmp)
	tmp = big.NewInt(0)

	if dropAmtMember1.LTE(sdk.ZeroInt()) {
		return sdk.ZeroInt(), sdk.ZeroInt(), sdkerrors.Wrapf(types.ErrAmtZero, "%s", member1.DenomB)
	}

	// `dropAmtMember2 = (drops * member2.Balance) / pool.Drops`
	tmp.Mul(drops.BigInt(), member2.Balance.BigInt())
	tmp.Quo(tmp, pool.Drops.BigInt())
	dropAmtMember2 := sdk.NewIntFromBigInt(tmp)
	//tmp = big.NewInt(0)

	if dropAmtMember2.LTE(sdk.ZeroInt()) {
		return sdk.ZeroInt(), sdk.ZeroInt(), sdkerrors.Wrapf(types.ErrAmtZero, "%s", member2.DenomB)
	}

	return dropAmtMember1, dropAmtMember2, nil
}

// GetOrderOwner returns orders from a single owner
func (k Keeper) GetDropCoin(
	ctx sdk.Context,
	denomA string,
	denomB string,
	amountA sdk.Int,
) (amountB sdk.Int, drops sdk.Int, found bool) {

	memberStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MemberKeyPrefix))

	b := memberStore.Get(types.MemberKey(
		denomB,
		denomA,
	))
	if b == nil {
		return amountB, drops, false
	}

	var member1 types.Member
	k.cdc.MustUnmarshal(b, &member1)

	c := memberStore.Get(types.MemberKey(
		denomA,
		denomB,
	))
	if c == nil {
		return amountB, drops, false
	}

	var member2 types.Member
	k.cdc.MustUnmarshal(c, &member2)

	poolStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKeyPrefix))

	prePair := []string{denomA, denomB}
	sort.Strings(prePair)
	pair := strings.Join(prePair, ",")

	d := poolStore.Get(types.PoolKey(pair))
	if d == nil {
		return amountB, drops, false
	}

	var pool types.Pool
	k.cdc.MustUnmarshal(d, &pool)

	amountB, drops, error := dropCoin(amountA, pool, member1, member2)
	if error != nil {
		return amountB, drops, false
	}

	found = true

	return
}

func dropCoin(amountA sdk.Int, pool types.Pool, memberA types.Member, memberB types.Member) (sdk.Int, sdk.Int, error) {
	// see `msg_server_redeem_drop` for our bigint strategy
	// `dropAmtMember1 = (drops * member1.Balance) / pool.Drops`
	tmp := big.NewInt(0)
	tmp.Mul(amountA.BigInt(), pool.Drops.BigInt())
	tmp.Quo(tmp, memberA.Balance.BigInt())
	drops := sdk.NewIntFromBigInt(tmp)
	tmp2 := big.NewInt(0)
	tmp2.Mul(tmp, memberB.Balance.BigInt())
	tmp2.Quo(tmp2, pool.Drops.BigInt())
	amountB := sdk.NewIntFromBigInt(tmp2)

	return amountB, drops, nil
}

// GetOrderOwner returns orders from a single owner
func (k Keeper) GetDropsToCoins(
	ctx sdk.Context,
	denom1 string,
	denom2 string,
	drops string,
) (amount1 sdk.Int, amount2 sdk.Int, err error) {

	dropsInt, ok := sdk.NewIntFromString(drops)
	if !ok {
		return amount1, amount2, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "drops not a valid integer")
	}

	memberStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MemberKeyPrefix))

	b := memberStore.Get(types.MemberKey(
		denom2,
		denom1,
	))
	if b == nil {
		return amount1, amount2, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "member not found")
	}

	var member1 types.Member
	k.cdc.MustUnmarshal(b, &member1)

	c := memberStore.Get(types.MemberKey(
		denom1,
		denom2,
	))
	if c == nil {
		return amount1, amount2, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "member not found")
	}

	var member2 types.Member
	k.cdc.MustUnmarshal(c, &member2)

	poolStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKeyPrefix))

	prePair := []string{denom1, denom2}
	sort.Strings(prePair)
	pair := strings.Join(prePair, ",")

	d := poolStore.Get(types.PoolKey(
		pair,
	))
	if d == nil {
		return amount1, amount2, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "pair not a valid denom pair")
	}

	var pool types.Pool
	k.cdc.MustUnmarshal(d, &pool)

	if denom1 == prePair[0] {
		amount1, amount2, error := dropAmounts(dropsInt, pool, member1, member2)
		if error != nil {
			return amount1, amount2, error
		}
	} else {
		amount2, amount1, error := dropAmounts(dropsInt, pool, member2, member1)
		if error != nil {
			return amount2, amount1, error
		}
	}

	return
}
