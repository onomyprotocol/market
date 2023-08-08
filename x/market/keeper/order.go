package keeper

import (
	"strings"

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

// GetOrderOwner returns orders from a single owner
func (k Keeper) GetOrderOwner(
	ctx sdk.Context,
	owner string,
) (orders types.Orders) {
	store1 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersKeyPrefix))

	b := store1.Get(types.OrdersKey(
		owner,
	))
	if b == nil {
		return orders
	}

	// var orders types.Orders
	// var order types.Order

	k.cdc.MustUnmarshal(b, &orders)

	/*
		store2 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderKeyPrefix))


		for _, uid := range orders.Uids {

			b := store2.Get(types.OrderKey(
				uid,
			))

			if b != nil {
				k.cdc.MustUnmarshal(b, &order)
				list = append(list, order)
			}
		}
	*/
	return
}

// SetOrderOwner adds an order to owner's open orders
func (k Keeper) SetOrderOwner(
	ctx sdk.Context,
	owner string,
	uid uint64,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersKeyPrefix))

	var orders types.Orders

	a := store.Get(types.OrdersKey(
		owner,
	))
	if a == nil {
		orders.Uids = []uint64{uid}
		b := k.cdc.MustMarshal(&orders)
		store.Set(types.OrdersKey(owner), b)
		return
	}

	index := 0
	k.cdc.MustUnmarshal(a, &orders)

	for index < len(orders.Uids) {
		if orders.Uids[index] == uid {
			orders.Uids = append(orders.Uids[:index], orders.Uids[index+1:]...)
			break
		}
		index++
	}

	orders.Uids = append(orders.Uids, uid)
	b := k.cdc.MustMarshal(&orders)
	store.Set(types.OrdersKey(owner), b)
}

// RemoveOrderOwner removes an order from owner's open orders
func (k Keeper) RemoveOrderOwner(
	ctx sdk.Context,
	owner string,
	uid uint64,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersKeyPrefix))

	a := store.Get(types.OrdersKey(
		owner,
	))
	if a == nil {
		return
	}

	var orders types.Orders
	index := 0
	k.cdc.MustUnmarshal(a, &orders)

	for index < len(orders.Uids) {
		if orders.Uids[index] == uid {
			orders.Uids = append(orders.Uids[:index], orders.Uids[index+1:]...)
			b := k.cdc.MustMarshal(&orders)
			store.Set(types.OrdersKey(owner), b)
			return
		}
		index++
	}
}

// GetOwnerPairOrders returns orders from a single owner
func (k Keeper) GetOwnerPairOrders(
	ctx sdk.Context,
	owner string,
	pair string,
) (list []types.Order) {
	store1 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersKeyPrefix))

	b := store1.Get(types.OrdersKey(
		owner,
	))
	if b == nil {
		return list
	}

	store2 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderKeyPrefix))

	var orders types.Orders
	var order types.Order

	k.cdc.MustUnmarshal(b, &orders)

	prePair := []string{"", ""}
	orderPair := ""

	for _, uid := range orders.Uids {

		b := store2.Get(types.OrderKey(
			uid,
		))

		if b != nil {
			k.cdc.MustUnmarshal(b, &order)
			prePair = []string{order.DenomAsk, order.DenomBid}
			orderPair = strings.Join(prePair, ",")
			if orderPair == pair {
				list = append(list, order)
			}
		}
	}

	return
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
			Status:    order.Status,
			OrderType: order.OrderType,
			DenomAsk:  order.DenomAsk,
			DenomBid:  order.DenomBid,
			Amount:    order.Amount.String(),
			Rate:      []string{order.Rate[0].String(), order.Rate[1].String()},
			Prev:      order.Prev,
			Next:      order.Next,
			BegTime:   order.BegTime,
		}
		list = append(list, orderResponse)
		uid = order.Next
	}

	return
}

// BookEnds returns adjacent orders determined by rate
func (k Keeper) BookEnds(
	ctx sdk.Context,
	denomA string,
	denomB string,
	orderType string,
	rate []sdk.Int,
) (ends [2]uint64) {

	member, _ := k.GetMember(ctx, denomA, denomB)
	var order types.Order

	if orderType == "limit" {

		if member.Limit == 0 {
			return [2]uint64{0, 0}
		}

		order, _ = k.GetOrder(ctx, member.Limit)

		for types.GTE(rate, order.Rate) {

			if order.Next == 0 {
				break
			}

			order, _ = k.GetOrder(ctx, order.Next)

		}

		if order.Next == 0 {
			if types.GTE(rate, order.Rate) {
				return [2]uint64{order.Uid, 0}
			}
		}

		return [2]uint64{order.Prev, order.Uid}

	} else {

		if member.Stop == 0 {
			return [2]uint64{0, 0}
		}

		order, _ = k.GetOrder(ctx, member.Stop)

		for types.LTE(rate, order.Rate) {

			if order.Next == 0 {
				break
			}

			order, _ = k.GetOrder(ctx, order.Next)

		}

		if order.Next == 0 {
			if types.LTE(rate, order.Rate) {
				return [2]uint64{order.Uid, 0}
			}
		}

		return [2]uint64{order.Prev, order.Uid}
	}
}
