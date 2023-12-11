package keeper

import (
	"context"

	"market/x/market/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) OrderAll(c context.Context, req *types.QueryAllOrderRequest) (*types.QueryOrdersResponse, error) {
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

	return &types.QueryOrdersResponse{Orders: orders, Pagination: pageRes}, nil
}

func (k Keeper) Order(c context.Context, req *types.QueryOrderRequest) (*types.QueryOrderResponse, error) {
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

	return &types.QueryOrderResponse{Order: val}, nil
}

func (k Keeper) Book(goCtx context.Context, req *types.QueryBookRequest) (*types.QueryBookResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	book := k.GetBook(ctx, req.DenomA, req.DenomB, req.OrderType)

	return &types.QueryBookResponse{Book: book}, nil
}

func (k Keeper) OrderOwner(c context.Context, req *types.QueryOrderOwnerRequest) (*types.QueryOrdersResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	orders := k.GetOrderOwner(ctx, req.Address)

	return &types.QueryOrdersResponse{Orders: orders}, nil
}

func (k Keeper) OrderOwnerUids(c context.Context, req *types.QueryOrderOwnerRequest) (*types.QueryOrderOwnerUidsResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	orders := k.GetOrderOwnerUids(ctx, req.Address)

	return &types.QueryOrderOwnerUidsResponse{Orders: orders}, nil
}

func (k Keeper) Quote(goCtx context.Context, req *types.QueryQuoteRequest) (*types.QueryQuoteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if req.DenomAsk != req.DenomAmount && req.DenomBid != req.DenomAmount {
		return nil, sdkerrors.Wrapf(types.ErrDenomMismatch, "Denom %s not ask or bid", req.DenomAmount)
	}

	amount, ok := sdk.NewIntFromString(req.Amount)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid amount integer")
	}

	memberAsk, found := k.GetMember(ctx, req.DenomBid, req.DenomAsk)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrMemberNotFound, "Member %s", req.DenomAsk)
	}

	memberBid, found := k.GetMember(ctx, req.DenomAsk, req.DenomBid)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrMemberNotFound, "Member %s", req.DenomBid)
	}

	denomResp, amountResp, error := k.GetQuote(ctx, memberAsk, memberBid, req.DenomAmount, amount)
	if error != nil {
		return nil, error
	}

	return &types.QueryQuoteResponse{Denom: denomResp, Amount: amountResp.String()}, nil
}
