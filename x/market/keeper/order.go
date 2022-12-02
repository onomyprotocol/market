package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/market/x/market/types"
)

// SetOrder set a specific order in the store from its index
func (k Keeper) SetOrder(ctx sdk.Context, order types.Order) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderKeyPrefix))
	b := k.cdc.MustMarshal(&order)
	store.Set(types.OrderKey(
		order.Uid,
		order.Owner,
		order.Active,
		order.OrderType,
		order.DenomAsk,
		order.DenomBid,
	), b)
}

// GetOrder returns a order from its index
func (k Keeper) GetOrder(
	ctx sdk.Context,
	uid uint64,
	owner string,
	active bool,
	orderType string,
	denomAsk string,
	denomBid string,

) (val types.Order, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderKeyPrefix))

	b := store.Get(types.OrderKey(
		uid,
		owner,
		active,
		orderType,
		denomAsk,
		denomBid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOrder removes a order from the store
func (k Keeper) RemoveOrder(
	ctx sdk.Context,
	uid uint64,
	owner string,
	active bool,
	orderType string,
	denomAsk string,
	denomBid string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderKeyPrefix))
	store.Delete(types.OrderKey(
		uid,
		owner,
		active,
		orderType,
		denomAsk,
		denomBid,
	))
}

// GetAllOrder returns all order
func (k Keeper) GetAllOrder(ctx sdk.Context) (list []types.Order) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Order
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
