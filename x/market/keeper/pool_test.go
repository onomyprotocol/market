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

func createNPool(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Pool {
	items := make([]types.Pool, n)
	for i := range items {
		items[i].Pair = strconv.Itoa(i)
		items[i].Denom1 = strconv.Itoa(i)
		items[i].Denom2 = strconv.Itoa(i)
		items[i].Leader = strconv.Itoa(i)

		keeper.SetPool(ctx, items[i])
	}
	return items
}

func TestPoolGet(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNPool(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPool(ctx,
			item.Pair,
			item.Denom1,
			item.Denom2,
			item.Leader,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPoolRemove(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNPool(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePool(ctx,
			item.Pair,
			item.Denom1,
			item.Denom2,
			item.Leader,
		)
		_, found := keeper.GetPool(ctx,
			item.Pair,
			item.Denom1,
			item.Denom2,
			item.Leader,
		)
		require.False(t, found)
	}
}

func TestPoolGetAll(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNPool(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPool(ctx)),
	)
}
