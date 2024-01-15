package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"market/x/portal/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.ProviderChannel(ctx),
		k.ReserveChannel(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// ProviderChannel returns the ProviderChannel param
func (k Keeper) ProviderChannel(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyProviderChannel, &res)
	return
}

// ReserveChannel returns the ReserveChannel param
func (k Keeper) ReserveChannel(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyReserveChannel, &res)
	return
}
