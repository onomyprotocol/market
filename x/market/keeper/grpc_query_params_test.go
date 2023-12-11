package keeper_test

import (
	"testing"

	testkeeper "market/testutil/keeper"
	"market/x/market/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	keeper := testkeeper.CreateTestEnvironment(t)
	wctx := sdk.WrapSDKContext(keeper.Context)
	params := types.DefaultParams()
	keeper.MarketKeeper.SetParams(keeper.Context, params)

	response, err := keeper.MarketKeeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
