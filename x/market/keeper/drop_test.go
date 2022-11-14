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

func createNDrop(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Drop {
	items := make([]types.Drop, n)
	for i := range items {
		items[i].Uid = uint64(i)
		items[i].Owner = strconv.Itoa(i)
		items[i].Pair = strconv.Itoa(i)

		keeper.SetDrop(ctx, items[i])
	}
	return items
}

func TestDropGet(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNDrop(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDrop(ctx,
			item.Uid,
			item.Owner,
			item.Pair,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDropRemove(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNDrop(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDrop(ctx,
			item.Uid,
			item.Owner,
			item.Pair,
		)
		_, found := keeper.GetDrop(ctx,
			item.Uid,
			item.Owner,
			item.Pair,
		)
		require.False(t, found)
	}
}

func TestDropGetAll(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNDrop(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDrop(ctx)),
	)
}
