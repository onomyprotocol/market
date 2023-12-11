package keeper

import (
	"market/x/market/types"

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

// EarnRate - the earning rate of the pool leader
func (k Keeper) EarnRates(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyEarnRates, &res)
	return
}

// BurnRate - the burning rate of the burn coin
func (k Keeper) BurnRate(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyBurnRate, &res)
	return
}

// BurnCoin - the burn coin
func (k Keeper) BurnCoin(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyBurnCoin, &res)
	return
}
