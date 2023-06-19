package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pendulum-labs/market/x/market/types"
)

// SetOrder set a specific order in the store from its index
func (k Keeper) SetOrder(ctx sdk.Context, order types.Order) {
	store1 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderKeyPrefix))

	b := k.cdc.MustMarshal(&order)
	store1.Set(types.OrderKey(
		order.Uid,
	), b)

	store2 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DropsKeyPrefix))

	c := store2.Get(types.DropsKey(
		order.Owner,
	))

	var orders types.Orders

	if c == nil {
		orders = types.Orders{
			Uids: []uint64{order.Uid},
		}
	} else {
		k.cdc.MustUnmarshal(c, &orders)
		orders.Uids = append(orders.Uids, order.Uid)
	}

	d := k.cdc.MustMarshal(&orders)
	store2.Set(types.OrdersKey(
		order.Owner,
	), d)
}

// GetOrder returns a order from its index
func (k Keeper) GetOrder(
	ctx sdk.Context,
	uid uint64,
) (val types.Order, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderKeyPrefix))

	b := store.Get(types.OrderKey(
		uid,
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
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderKeyPrefix))
	store.Delete(types.OrderKey(
		uid,
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

// GetOrder returns a order from its index
func (k Keeper) GetBook(
	ctx sdk.Context,
	denomA string,
	denomB string,
	orderType string,
) (list []types.OrderResponse) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderKeyPrefix))

	member, _ := k.GetMember(ctx, denomA, denomB)

	var uid uint64

	if orderType == "limit" {
		uid = member.Limit
	} else {
		uid = member.Stop
	}

	if uid == 0 {
		return nil
	}

	for uid > 0 {
		b := store.Get(types.OrderKey(
			uid,
		))
		var order types.Order
		k.cdc.MustUnmarshal(b, &order)
		orderResponse := types.OrderResponse{
			Uid:       order.Uid,
			Owner:     order.Owner,
			Active:    order.Active,
			OrderType: order.OrderType,
			DenomAsk:  order.DenomAsk,
			DenomBid:  order.DenomBid,
			Amount:    order.Amount.String(),
			Rate:      []string{order.Rate[0].String(), order.Rate[1].String()},
			Prev:      order.Prev,
			Next:      order.Next,
		}
		list = append(list, orderResponse)
		uid = order.Next
	}

	return
}

// GetOrder returns a order from its index
func (k Keeper) GetBookEnds(
	ctx sdk.Context,
	denomA string,
	denomB string,
	orderType string,
) (ends []types.OrderResponse) {
	return nil
}
