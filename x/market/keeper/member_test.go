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

func createNMember(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Member {
	items := make([]types.Member, n)
	for i := range items {
		items[i].Pair = strconv.Itoa(i)
		items[i].DenomA = strconv.Itoa(i)
		items[i].DenomB = strconv.Itoa(i)

		keeper.SetMember(ctx, items[i])
	}
	return items
}

func TestMemberGet(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNMember(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetMember(ctx,
			item.DenomA,
			item.DenomB,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestMemberRemove(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNMember(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMember(ctx,
			item.DenomA,
			item.DenomB,
		)
		_, found := keeper.GetMember(ctx,
			item.DenomA,
			item.DenomB,
		)
		require.False(t, found)
	}
}

func TestMemberGetAll(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNMember(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMember(ctx)),
	)
}
