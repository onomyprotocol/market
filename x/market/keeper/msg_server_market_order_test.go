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

	_, _, denomA, denomB, _ := common(t, testInput)

	// beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)

	//Create Order
	var o = types.MsgMarketOrder{Creator: addr, DenomAsk: denomA, AmountAsk: "8", DenomBid: denomB, AmountBid: "10", Slippage: "700"}

	_, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).MarketOrder(sdk.WrapSDKContext(testInput.Context), &o)
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
	require.Equal(t, "26", memberAsk.Balance.String())
	require.Equal(t, memberAsk.Stop, uint64(0))

}
