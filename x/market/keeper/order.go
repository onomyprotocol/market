package keeper

import (
	"strconv"
	"strings"

	"market/x/market/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// SetOrder set a specific order in the store from its index
func (k Keeper) SetOrder(ctx sdk.Context, order types.Order) {

	// order event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeOrder,
			sdk.NewAttribute(types.AttributeKeyUid, strconv.FormatUint(order.Uid, 10)),
			sdk.NewAttribute(types.AttributeKeyOwner, order.Owner),
			sdk.NewAttribute(types.AttributeKeyStatus, order.Status),
			sdk.NewAttribute(types.AttributeKeyOrderType, order.OrderType),
			sdk.NewAttribute(types.AttributeKeyDenomAsk, order.DenomAsk),
			sdk.NewAttribute(types.AttributeKeyDenomBid, order.DenomBid),
			sdk.NewAttribute(types.AttributeKeyAmount, order.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyRate, strings.Join([]string{order.Rate[0].String(), order.Rate[1].String()}, ",")),
			sdk.NewAttribute(types.AttributeKeyPrev, strconv.FormatUint(order.Prev, 10)),
			sdk.NewAttribute(types.AttributeKeyNext, strconv.FormatUint(order.Next, 10)),
			sdk.NewAttribute(types.AttributeKeyBeginTime, strconv.FormatInt(order.BegTime, 10)),
			sdk.NewAttribute(types.AttributeKeyUpdateTime, strconv.FormatInt(order.UpdTime, 10)),
		),
	)

	store1 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderKeyPrefix))

	b := k.cdc.MustMarshal(&order)
	store1.Set(types.OrderKey(
		order.Uid,
	), b)
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

// SetOrderOwner adds an order to owner's open orders
func (k Keeper) SetOrderOwner(
	ctx sdk.Context,
	owner string,
	uid uint64,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderOwnerKeyPrefix))

	var orders types.Orders

	a := store.Get(types.OrderOwnerKey(
		owner,
	))
	if a == nil {
		orders.Uids = []uint64{uid}
		b := k.cdc.MustMarshal(&orders)
		store.Set(types.OrderOwnerKey(owner), b)
		return
	}

	k.cdc.MustUnmarshal(a, &orders)

	// First remove uid if present
	// Allows the order, if changed, to be at top of list
	orders.Uids, _ = removeUid(orders.Uids, uid)

	// Append uid in the front
	orders.Uids = append(orders.Uids, uid)
	b := k.cdc.MustMarshal(&orders)
	store.Set(types.OrderOwnerKey(owner), b)
}

// GetOrderOwner returns order uids from a single owner
func (k Keeper) GetOrderOwnerUids(
	ctx sdk.Context,
	owner string,
) (orders types.Orders) {
	store1 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderOwnerKeyPrefix))

	a := store1.Get(types.OrderOwnerKey(
		owner,
	))
	if a == nil {
		return orders
	}

	k.cdc.MustUnmarshal(a, &orders)

	return orders
}

// GetOrderOwner returns orders from a single owner
func (k Keeper) GetOrderOwner(
	ctx sdk.Context,
	owner string,
) (list []types.Order) {
	store1 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderOwnerKeyPrefix))

	a := store1.Get(types.OrderOwnerKey(
		owner,
	))
	if a == nil {
		return list
	}

	var orders types.Orders

	k.cdc.MustUnmarshal(a, &orders)

	store2 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderKeyPrefix))

	for _, uid := range orders.Uids {
		var order types.Order

		b := store2.Get(types.OrderKey(
			uid,
		))

		if b != nil {
			k.cdc.MustUnmarshal(b, &order)
			list = append(list, order)
		}
	}

	return
}

// RemoveOrderOwner removes an order from owner's open orders
func (k Keeper) RemoveOrderOwner(
	ctx sdk.Context,
	owner string,
	uid uint64,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderOwnerKeyPrefix))

	a := store.Get(types.OrderOwnerKey(
		owner,
	))
	if a == nil {
		return
	}

	var orders types.Orders
	k.cdc.MustUnmarshal(a, &orders)

	orders.Uids, _ = removeUid(orders.Uids, uid)

	b := k.cdc.MustMarshal(&orders)
	store.Set(types.OrderOwnerKey(owner), b)
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

	i := 0

	for uid > 0 && i < 100 {
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
		i++
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

// BookEnds returns adjacent orders determined by rate
func (k Keeper) GetQuote(
	ctx sdk.Context,
	memberAsk types.Member,
	memberBid types.Member,
	denomAmount string,
	amount sdk.Int,
) (string, sdk.Int, error) {

	denom := memberAsk.DenomB
	var amountResp sdk.Int

	if denomAmount == memberBid.DenomB {

		// A(i)*B(i) = A(f)*B(f)
		// A(f) = A(i)*B(i)/B(f)
		// amountAsk = A(i) - A(f) = A(i) - A(i)*B(i)/B(f)
		amountResp = memberAsk.Balance.Sub(((memberAsk.Balance.Mul(memberBid.Balance)).Quo(memberBid.Balance.Add(amount))).Add(sdk.NewInt(1)))

		// Market Order Fee
		fee, _ := sdk.NewIntFromString(k.getParams(ctx).MarketFee)
		amountResp = amountResp.Sub((amountResp.Mul(fee)).Quo(sdk.NewInt(10000)))

		// Edge case where strikeAskAmount rounds to 0
		// Rounding favors AMM vs Order
		if amountResp.Equal(sdk.ZeroInt()) {
			return denom, sdk.ZeroInt(), sdkerrors.Wrapf(types.ErrAmtZero, "amount ask equal to zero")
		}

	} else {
		denom = memberBid.DenomB

		// Market Order Fee
		fee, _ := sdk.NewIntFromString(k.getParams(ctx).MarketFee)
		amountPlusFee := amount.Add((amount.Mul(fee)).Quo(sdk.NewInt(10000))).Add(sdk.NewInt(1))

		// A(i)*B(i) = A(f)*B(f)
		// B(f) = A(i)*B(i)/A(f)
		// amountBid = B(f) - B(i) = A(i)*B(i)/A(f) - B(i) = A(i)*B(i)/(A(i) - amountAskPlusFee) - B(i)
		amountResp = ((memberAsk.Balance.Mul(memberBid.Balance)).Quo(memberAsk.Balance.Sub(amountPlusFee))).Sub(memberBid.Balance)

		// Edge case where strikeAskAmount rounds to 0
		// Rounding favors AMM vs Order
		if amountResp.LTE(sdk.ZeroInt()) {
			return denom, sdk.ZeroInt(), sdkerrors.Wrapf(types.ErrLiquidityLow, "not enough liquidity")
		}

	}

	return denom, amountResp, nil
}
