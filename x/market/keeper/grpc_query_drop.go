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
