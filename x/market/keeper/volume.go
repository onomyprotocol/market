package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pendulum-labs/market/x/market/types"
)

// SetVolume set a specific volume in the store from its index
func (k Keeper) SetVolume(ctx sdk.Context, volume types.Volume) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VolumeKeyPrefix))
	b := k.cdc.MustMarshal(&volume)
	store.Set(types.VolumeKey(
		volume.Denom,
	), b)
}

// GetVolume returns a volume from its index
func (k Keeper) GetVolume(
	ctx sdk.Context,
	denom string,
) (val types.Volume, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VolumeKeyPrefix))

	b := store.Get(types.VolumeKey(
		denom,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetVolume returns a volume from its index
func (k Keeper) IncVolume(
	ctx sdk.Context,
	denom string,
	amount sdk.Int,
) types.Volume {
	volume, found := k.GetVolume(ctx, denom)

	if found {
		volume.Amount = volume.Amount.Add(amount)
	} else {
		volume.Denom = denom
		volume.Amount = amount
	}

	k.SetVolume(ctx, volume)

	return volume
}

// RemoveVolume removes a volume from the store
func (k Keeper) RemoveVolume(
	ctx sdk.Context,
	denom string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VolumeKeyPrefix))
	store.Delete(types.VolumeKey(
		denom,
	))
}

// GetAllVolume returns all volumes
func (k Keeper) GetAllVolumes(ctx sdk.Context) (list []types.Volume) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VolumeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Volume
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
