package keeper_test

import (
	"testing"

	testkeeper "market/testutil/keeper"
	"market/x/market/types"

	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k := testkeeper.CreateTestEnvironment(t)
	params := types.DefaultParams()

	k.MarketKeeper.SetParams(k.Context, params)

	require.EqualValues(t, params, k.MarketKeeper.GetParams(k.Context))
}
