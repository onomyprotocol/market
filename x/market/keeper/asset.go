package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pendulum-labs/market/x/market/types"
)

// SetAsset set a specific asset in the store from its index
func (k Keeper) SetAsset(ctx sdk.Context, asset types.Asset) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssetKeyPrefix))
	b := k.cdc.MustMarshal(&asset)
	store.Set(types.AssetKey(
		asset.Active,
		asset.Owner,
		asset.AssetType,
	), b)
}

// GetAsset returns a asset from its index
func (k Keeper) GetAsset(
	ctx sdk.Context,
	active bool,
	owner string,
	assetType string,

) (val types.Asset, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssetKeyPrefix))

	b := store.Get(types.AssetKey(
		active,
		owner,
		assetType,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAsset removes a asset from the store
func (k Keeper) RemoveAsset(
	ctx sdk.Context,
	active bool,
	owner string,
	assetType string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssetKeyPrefix))
	store.Delete(types.AssetKey(
		active,
		owner,
		assetType,
	))
}

// GetAllAsset returns all asset
func (k Keeper) GetAllAsset(ctx sdk.Context) (list []types.Asset) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssetKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Asset
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
