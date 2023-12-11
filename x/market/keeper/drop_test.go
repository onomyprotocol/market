package keeper_test

import (
	"strconv"
	"testing"

	keepertest "market/testutil/keeper"
	"market/testutil/nullify"
	"market/x/market/keeper"
	"market/x/market/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNDrop(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Drop {
	items := make([]types.Drop, n)
	for i := range items {
		items[i].Uid = uint64(i)
		items[i].Owner = strconv.Itoa(i)
		items[i].Pair = strconv.Itoa(i)
		items[i].Drops = sdk.NewIntFromUint64(uint64(i))
		items[i].Product = sdk.NewIntFromUint64(uint64(i))

		keeper.SetDrop(ctx, items[i])
	}
	return items
}

func TestDropGet(t *testing.T) {
	keeper := keepertest.CreateTestEnvironment(t)
	items := createNDrop(keeper.MarketKeeper, keeper.Context, 10)
	for _, item := range items {
		rst, found := keeper.MarketKeeper.GetDrop(keeper.Context,
			item.Uid,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDropRemove(t *testing.T) {
	keeper := keepertest.CreateTestEnvironment(t)
	items := createNDrop(keeper.MarketKeeper, keeper.Context, 10)
	for _, item := range items {
		keeper.MarketKeeper.RemoveDrop(keeper.Context,
			item.Uid,
		)
		_, found := keeper.MarketKeeper.GetDrop(keeper.Context,
			item.Uid,
		)
		require.False(t, found)
	}
}

func TestDropGetAll(t *testing.T) {
	keeper := keepertest.CreateTestEnvironment(t)
	items := createNDrop(keeper.MarketKeeper, keeper.Context, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.MarketKeeper.GetAllDrop(keeper.Context)),
	)
}
