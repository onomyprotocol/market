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

func (k Keeper) DropAll(c context.Context, req *types.QueryAllDropRequest) (*types.QueryDropsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var drops []types.Drop
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	dropStore := prefix.NewStore(store, types.KeyPrefix(types.DropKeyPrefix))

	pageRes, err := query.Paginate(dropStore, req.Pagination, func(key []byte, value []byte) error {
		var drop types.Drop
		if err := k.cdc.Unmarshal(value, &drop); err != nil {
			return err
		}

		drops = append(drops, drop)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryDropsResponse{Drops: drops, Pagination: pageRes}, nil
}

func (k Keeper) Drop(c context.Context, req *types.QueryDropRequest) (*types.QueryDropResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetDrop(
		ctx,
		req.Uid,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryDropResponse{Drop: val}, nil
}

func (k Keeper) DropPairs(c context.Context, req *types.QueryDropPairsRequest) (*types.QueryDropPairsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetDropPairs(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryDropPairsResponse{Pairs: val.Pairs}, nil
}

func (k Keeper) DropOwnerPair(c context.Context, req *types.QueryDropOwnerPairRequest) (*types.QueryDropsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	drops, found := k.GetDropOwnerPair(
		ctx,
		req.Address,
		req.Pair,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryDropsResponse{Drops: drops}, nil
}

func (k Keeper) DropAmounts(c context.Context, req *types.QueryDropAmountsRequest) (*types.QueryDropAmountsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	denom1, denom2, amount1, amount2, found := k.GetDropAmounts(
		ctx,
		req.Uid,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryDropAmountsResponse{
		Denom1:  denom1,
		Denom2:  denom2,
		Amount1: amount1.String(),
		Amount2: amount2.String(),
	}, nil
}

func (k Keeper) DropCoin(c context.Context, req *types.QueryDropCoinRequest) (*types.QueryDropCoinResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	err := sdk.ValidateDenom(req.DenomA)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid denomA")
	}

	err = sdk.ValidateDenom(req.DenomB)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid denomB")
	}

	amountA, ok := sdk.NewIntFromString(req.AmountA)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "invalid amountA")
	}
	if amountA.LTE(sdk.ZeroInt()) {
		return nil, status.Error(codes.InvalidArgument, "invalid amountA")
	}

	amountB, drops, found := k.GetDropCoin(
		ctx,
		req.DenomA,
		req.DenomB,
		amountA,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryDropCoinResponse{
		AmountB: amountB.String(),
		Drops:   drops.String(),
	}, nil
}

func (k Keeper) DropsToCoins(c context.Context, req *types.QueryDropsToCoinsRequest) (*types.QueryDropAmountsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	denom1, denom2, amount1, amount2, found := k.GetDropsToCoins(
		ctx,
		req.Pair,
		req.Drops,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryDropAmountsResponse{
		Denom1:  denom1,
		Denom2:  denom2,
		Amount1: amount1.String(),
		Amount2: amount2.String(),
	}, nil
}
