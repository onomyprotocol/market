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
	keeper := keepertest.CreateTestEnvironment(t)
	items := createNMember(keeper.MarketKeeper, keeper.Context, 10)
	for _, item := range items {
		rst, found := keeper.MarketKeeper.GetMember(keeper.Context,
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
	keeper := keepertest.CreateTestEnvironment(t)
	items := createNMember(keeper.MarketKeeper, keeper.Context, 10)
	for _, item := range items {
		keeper.MarketKeeper.RemoveMember(keeper.Context,
			item.DenomA,
			item.DenomB,
		)
		_, found := keeper.MarketKeeper.GetMember(keeper.Context,
			item.DenomA,
			item.DenomB,
		)
		require.False(t, found)
	}
}

func TestMemberGetAll(t *testing.T) {
	keeper := keepertest.CreateTestEnvironment(t)
	items := createNMember(keeper.MarketKeeper, keeper.Context, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.MarketKeeper.GetAllMember(keeper.Context)),
	)
}
