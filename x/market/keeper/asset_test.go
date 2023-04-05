package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/pendulum-labs/market/testutil/keeper"
	"github.com/pendulum-labs/market/testutil/nullify"
	"github.com/pendulum-labs/market/x/market/keeper"
	"github.com/pendulum-labs/market/x/market/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNAsset(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Asset {
	items := make([]types.Asset, n)
	for i := range items {
		items[i].Active = false
		items[i].Owner = strconv.Itoa(i)
		items[i].AssetType = strconv.Itoa(i)

		keeper.SetAsset(ctx, items[i])
	}
	return items
}

func TestAssetGet(t *testing.T) {
	keeper := keepertest.CreateTestEnvironment(t)
	items := createNAsset(keeper.MarketKeeper, keeper.Context, 10)
	for _, item := range items {
		rst, found := keeper.MarketKeeper.GetAsset(keeper.Context,
			item.Active,
			item.Owner,
			item.AssetType,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAssetRemove(t *testing.T) {
	keeper := keepertest.CreateTestEnvironment(t)
	items := createNAsset(keeper.MarketKeeper, keeper.Context, 10)
	for _, item := range items {
		keeper.MarketKeeper.RemoveAsset(keeper.Context,
			item.Active,
			item.Owner,
			item.AssetType,
		)
		_, found := keeper.MarketKeeper.GetAsset(keeper.Context,
			item.Active,
			item.Owner,
			item.AssetType,
		)
		require.False(t, found)
	}
}

func TestAssetGetAll(t *testing.T) {
	keeper := keepertest.CreateTestEnvironment(t)
	items := createNAsset(keeper.MarketKeeper, keeper.Context, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.MarketKeeper.GetAllAsset(keeper.Context)),
	)
}
