package keeper

import (
	"context"

	"market/x/market/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) BurningsAll(c context.Context, req *types.QueryAllBurningsRequest) (*types.QueryAllBurningsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var burningss []types.Burnings
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	burningsStore := prefix.NewStore(store, types.KeyPrefix(types.BurningsKeyPrefix))

	pageRes, err := query.Paginate(burningsStore, req.Pagination, func(key []byte, value []byte) error {
		var burnings types.Burnings
		if err := k.cdc.Unmarshal(value, &burnings); err != nil {
			return err
		}

		burningss = append(burningss, burnings)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBurningsResponse{Burnings: burningss, Pagination: pageRes}, nil
}

func (k Keeper) Burnings(c context.Context, req *types.QueryGetBurningsRequest) (*types.QueryGetBurningsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetBurnings(
		ctx,
		req.Denom,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetBurningsResponse{Burnings: val}, nil
}

func (k Keeper) Burned(c context.Context, req *types.QueryBurnedRequest) (*types.QueryBurnedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	// Coin that will be burned
	burnCoin := k.BurnCoin(ctx)

	val := k.GetBurned(
		ctx,
	)

	return &types.QueryBurnedResponse{Denom: burnCoin, Amount: val.Amount.String()}, nil
}
