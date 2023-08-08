package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/pendulum-labs/market/x/market/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) OrderAll(c context.Context, req *types.QueryAllOrderRequest) (*types.QueryAllOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var orders []types.Order
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	orderStore := prefix.NewStore(store, types.KeyPrefix(types.OrderKeyPrefix))

	pageRes, err := query.Paginate(orderStore, req.Pagination, func(key []byte, value []byte) error {
		var order types.Order
		if err := k.cdc.Unmarshal(value, &order); err != nil {
			return err
		}

		orders = append(orders, order)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllOrderResponse{Order: orders, Pagination: pageRes}, nil
}

func (k Keeper) Order(c context.Context, req *types.QueryGetOrderRequest) (*types.QueryGetOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetOrder(
		ctx,
		req.Uid,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetOrderResponse{Order: val}, nil
}

func (k Keeper) Book(goCtx context.Context, req *types.QueryBookRequest) (*types.QueryBookResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	book := k.GetBook(ctx, req.DenomA, req.DenomB, req.OrderType)

	return &types.QueryBookResponse{Book: book}, nil
}

func (k Keeper) OrderOwner(c context.Context, req *types.QueryOrderOwnerRequest) (*types.QueryOrderOwnerResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	orders := k.GetOrderOwner(ctx, req.Address)

	return &types.QueryOrderOwnerResponse{Orders: orders}, nil
}

func (k Keeper) OrderOwnerPair(c context.Context, req *types.QueryOrderOwnerPairRequest) (*types.QueryOrderOwnerPairResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	orders := k.GetOwnerPairOrders(ctx, req.Address, req.Pair)

	return &types.QueryOrderOwnerPairResponse{Order: orders}, nil
}
