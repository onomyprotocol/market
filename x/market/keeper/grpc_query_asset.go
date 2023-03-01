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

func (k Keeper) AssetAll(c context.Context, req *types.QueryAllAssetRequest) (*types.QueryAllAssetResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var assets []types.Asset
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	assetStore := prefix.NewStore(store, types.KeyPrefix(types.AssetKeyPrefix))

	pageRes, err := query.Paginate(assetStore, req.Pagination, func(key []byte, value []byte) error {
		var asset types.Asset
		if err := k.cdc.Unmarshal(value, &asset); err != nil {
			return err
		}

		assets = append(assets, asset)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAssetResponse{Asset: assets, Pagination: pageRes}, nil
}

func (k Keeper) Asset(c context.Context, req *types.QueryGetAssetRequest) (*types.QueryGetAssetResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetAsset(
		ctx,
		req.Active,
		req.Owner,
		req.AssetType,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetAssetResponse{Asset: val}, nil
}
