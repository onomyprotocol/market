package keeper

import (
	"market/x/market/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

// GetBurned get the amount of NOM Burned by ONEX
func (k Keeper) GetBurned(ctx sdk.Context) (burned types.Burned) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.BurnedKey)
	a := store.Get(byteKey)

	// Burned doesn't exist return zero int
	if a == nil {
		return types.Burned{
			Amount: sdk.ZeroInt(),
		}
	}

	k.cdc.MustUnmarshal(a, &burned)

	return
}

// SetBurned set the amount of NOM Burned by ONEX
func (k Keeper) AddBurned(ctx sdk.Context, amount sdk.Int) (burned types.Burned) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.BurnedKey)

	a := store.Get(byteKey)

	// Burned doesn't exist then initialize with amount
	if a == nil {
		burned = types.Burned{
			Amount: amount,
		}
	} else {
		k.cdc.MustUnmarshal(a, &burned)
		burned.Amount = burned.Amount.Add(amount)
	}

	b := k.cdc.MustMarshal(&burned)
	store.Set(byteKey, b)

	return
}
