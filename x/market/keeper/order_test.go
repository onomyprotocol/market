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

func createNOrder(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Order {
	items := make([]types.Order, n)
	for i := range items {
		items[i].Uid = uint64(i)
		items[i].Owner = strconv.Itoa(i)
		items[i].Active = false
		items[i].OrderType = strconv.Itoa(i)
		items[i].DenomAsk = strconv.Itoa(i)
		items[i].DenomBid = strconv.Itoa(i)
		items[i].Amount = sdk.NewInt(int64(i))
		items[i].Rate = []sdk.Int{sdk.NewInt(int64(i)), sdk.NewInt(int64(i))}

		keeper.SetOrder(ctx, items[i])
	}
	return items
}

func TestOrderGet(t *testing.T) {
	keeper := keepertest.CreateTestEnvironment(t)
	items := createNOrder(keeper.MarketKeeper, keeper.Context, 10)
	for _, item := range items {
		rst, found := keeper.MarketKeeper.GetOrder(keeper.Context,
			item.Uid,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestOrderRemove(t *testing.T) {
	keeper := keepertest.CreateTestEnvironment(t)
	items := createNOrder(keeper.MarketKeeper, keeper.Context, 10)
	for _, item := range items {
		keeper.MarketKeeper.RemoveOrder(keeper.Context,
			item.Uid,
		)
		_, found := keeper.MarketKeeper.GetOrder(keeper.Context,
			item.Uid,
		)
		require.False(t, found)
	}
}

func TestOrderGetAll(t *testing.T) {
	keeper := keepertest.CreateTestEnvironment(t)
	items := createNOrder(keeper.MarketKeeper, keeper.Context, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.MarketKeeper.GetAllOrder(keeper.Context)),
	)
}
