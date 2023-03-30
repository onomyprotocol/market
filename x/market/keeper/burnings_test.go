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

func createNBurnings(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Burnings {
	items := make([]types.Burnings, n)
	for i := range items {
		items[i].Denom = strconv.Itoa(i)

		keeper.SetBurnings(ctx, items[i])
	}
	return items
}

func TestBurningsGet(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNBurnings(keeper.MarketKeeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.MarketKeeper.GetBurnings(ctx,
			item.Denom,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestBurningsRemove(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNBurnings(keeper.MarketKeeper, ctx, 10)
	for _, item := range items {
		keeper.MarketKeeper.RemoveBurnings(ctx,
			item.Denom,
		)
		_, found := keeper.MarketKeeper.GetBurnings(ctx,
			item.Denom,
		)
		require.False(t, found)
	}
}

func TestBurningsGetAll(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNBurnings(keeper.MarketKeeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.MarketKeeper.GetAllBurnings(ctx)),
	)
}
