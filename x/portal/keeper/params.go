package keeper

import (
	"market/x/portal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return k.getParams(ctx)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) getParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// ReserveChannel - the IBC channel of the Reserve Chain
func (k Keeper) ReserveChannel(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyReserveChannel, &res)
	return
}

// OnomyChannel - the IBC channel of the Onomy Chain
func (k Keeper) OnomyChannel(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyOnomyChannel, &res)
	return
}
