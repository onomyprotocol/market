package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/onomyprotocol/market/testutil/keeper"
	"github.com/onomyprotocol/market/testutil/nullify"
	"github.com/onomyprotocol/market/x/market/keeper"
	"github.com/onomyprotocol/market/x/market/types"
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
	items := createNBurnings(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetBurnings(ctx,
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
	items := createNBurnings(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveBurnings(ctx,
			item.Denom,
		)
		_, found := keeper.GetBurnings(ctx,
			item.Denom,
		)
		require.False(t, found)
	}
}

func TestBurningsGetAll(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNBurnings(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllBurnings(ctx)),
	)
}
