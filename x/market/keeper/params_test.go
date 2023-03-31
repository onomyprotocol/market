package keeper_test

import (
	"testing"

	testkeeper "github.com/pendulum-labs/market/testutil/keeper"
	"github.com/pendulum-labs/market/x/market/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k := testkeeper.CreateTestEnvironment(t)
	params := types.DefaultParams()

	k.MarketKeeper.SetParams(k.Context, params)

	require.EqualValues(t, params, k.MarketKeeper.GetParams(k.Context))
}
