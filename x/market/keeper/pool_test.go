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

func createNPool(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Pool {
	items := make([]types.Pool, n)
	for i := range items {
		items[i].Pair = strconv.Itoa(i)
		items[i].Denom1 = strconv.Itoa(i)
		items[i].Denom2 = strconv.Itoa(i)
		items[i].Leaders = []*types.Leader{
			{
				Address: strconv.Itoa(i),
				Drops:   sdk.NewInt(int64(i)),
			},
		}
		items[i].Drops = sdk.NewIntFromUint64(uint64(0))

		keeper.SetPool(ctx, items[i])
	}
	return items
}

func TestPoolGet(t *testing.T) {
	keeper := keepertest.CreateTestEnvironment(t)
	items := createNPool(keeper.MarketKeeper, keeper.Context, 10)
	for _, item := range items {
		rst, found := keeper.MarketKeeper.GetPool(keeper.Context,
			item.Pair,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPoolRemove(t *testing.T) {
	keeper := keepertest.CreateTestEnvironment(t)
	items := createNPool(keeper.MarketKeeper, keeper.Context, 10)
	for _, item := range items {
		keeper.MarketKeeper.RemovePool(keeper.Context,
			item.Pair,
		)
		_, found := keeper.MarketKeeper.GetPool(keeper.Context,
			item.Pair,
		)
		require.False(t, found)
	}
}

func TestPoolGetAll(t *testing.T) {
	keeper := keepertest.CreateTestEnvironment(t)
	items := createNPool(keeper.MarketKeeper, keeper.Context, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.MarketKeeper.GetAllPool(keeper.Context)),
	)
}
