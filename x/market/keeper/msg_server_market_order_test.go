package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/pendulum-labs/market/testutil/keeper"
	"github.com/pendulum-labs/market/x/market/keeper"
	"github.com/pendulum-labs/market/x/market/types"
	"github.com/stretchr/testify/require"
)

func TestMarketOrder(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	wctx := sdk.WrapSDKContext(testInput.Context)

	_, _, denomA, denomB, pair := common(t, testInput)

	// beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)

	//Create Order
	var o = types.MsgMarketOrder{Creator: addr, DenomAsk: denomA, AmountAsk: "15", DenomBid: denomB, AmountBid: "10", Slippage: "700"}

	quoteBid, error := testInput.MarketKeeper.Quote(wctx, &types.QueryQuoteRequest{
		DenomAsk:    o.DenomAsk,
		DenomBid:    o.DenomBid,
		DenomAmount: o.DenomAsk,
		Amount:      o.AmountAsk,
	})

	o.AmountBid = quoteBid.Amount

	require.NoError(t, error)

	_, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).MarketOrder(wctx, &o)
	require.NoError(t, err)

	/*
		aftercount := testInput.MarketKeeper.GetUidCount(testInput.Context)
		require.Equal(t, beforecount+1, aftercount)

		//Validate Order
		orders, orderfound := testInput.MarketKeeper.GetOrder(testInput.Context, beforecount)
		require.True(t, orderfound)
		require.Equal(t, orders.DenomBid, denomB)
		require.Equal(t, orders.DenomAsk, denomA)
		require.Equal(t, orders.Amount.String(), o.Amount)

	*/

	// Validate GetMember
	memberAsk, memberAskfound := testInput.MarketKeeper.GetMember(testInput.Context, denomB, denomA)

	require.True(t, memberAskfound)
	require.Equal(t, memberAsk.DenomA, denomB)
	require.Equal(t, memberAsk.DenomB, denomA)
	require.Equal(t, "18", memberAsk.Balance.String())
	require.Equal(t, memberAsk.Stop, uint64(0))

	// Validate order estimation

	pool, poolFound := testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, poolFound)
	order, orderFound := testInput.MarketKeeper.GetOrder(testInput.Context, pool.History)
	require.True(t, orderFound)
	amountAskInt, ok := sdk.NewIntFromString(o.AmountAsk)
	require.True(t, ok)
	amountBidInt, ok := sdk.NewIntFromString(o.AmountBid)
	require.True(t, ok)
	require.True(t, types.EQ(order.Rate, []sdk.Int{amountAskInt, amountBidInt}))

}
