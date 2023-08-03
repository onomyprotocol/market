package keeper

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pendulum-labs/market/x/market/types"
)

// SetPool set a specific pool in the store from its index
func (k Keeper) SetPool(ctx sdk.Context, pool types.Pool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKeyPrefix))
	b := k.cdc.MustMarshal(&pool)
	store.Set(types.PoolSetKey(
		pool.Pair,
	), b)
}

// GetPool returns a pool from its index
func (k Keeper) GetPool(
	ctx sdk.Context,
	pair string,
) (val types.Pool, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKeyPrefix))

	b := store.Get(types.PoolKey(
		pair,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePool removes a pool from the store
func (k Keeper) RemovePool(
	ctx sdk.Context,
	pair string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKeyPrefix))
	store.Delete(types.PoolKey(
		pair,
	))
}

// GetAllPool returns all pool
func (k Keeper) GetAllPool(ctx sdk.Context) (list []types.Pool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Pool
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetHistory returns history of pool trades
func (k Keeper) GetHistory(
	ctx sdk.Context,
	pair string,
	length string,
) (list []types.OrderResponse, found bool) {

	len, err := strconv.ParseUint(length, 10, 64)
	if err != nil {
		len = 0
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderKeyPrefix))

	pool, found := k.GetPool(ctx, pair)

	if !found {
		return nil, found
	}

	uid := pool.History

	if uid == 0 {
		return nil, found
	}

	counter := uint64(0)

	for uid > 0 && (counter < len || len == 0) {
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
			EndTime:   order.EndTime,
		}
		list = append(list, orderResponse)
		counter = counter + 1
		uid = order.Next
	}

	return
}
