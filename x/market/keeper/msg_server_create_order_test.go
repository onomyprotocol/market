package keeper_test

import (
	"strconv"
	"strings"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/pendulum-labs/market/testutil/keeper"
	"github.com/pendulum-labs/market/testutil/sample"
	"github.com/pendulum-labs/market/x/market/keeper"
	"github.com/pendulum-labs/market/x/market/types"
	"github.com/stretchr/testify/require"
)

func common(t *testing.T, testInput keepertest.TestInput) (
	testdata testData,
	coinPair sdk.Coins,
	denomA string,
	denomB string,
	pair string,
) {

	// TestData
	testdata = testData{coinAStr: "30CoinA", coinBStr: "40CoinB", RateAstrArray: []string{"60", "70"}, RateBstrArray: []string{"80", "90"}}
	coinPair, _ = sample.SampleCoins("140CoinA", "140CoinB")
	denomA, denomB = sample.SampleDenoms(coinPair)
	pair = strings.Join([]string{denomA, denomB}, ",")

	// MintCoins
	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, types.ModuleName, coinPair))

	// SendCoinsFromModuleToAccount
	requestAddress, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, types.ModuleName, requestAddress, coinPair))

	// Create Pool
	var p = types.MsgCreatePool{CoinA: testdata.coinAStr, CoinB: testdata.coinBStr, Creator: addr}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	require.NoError(t, err)

	// CreateDrop
	var d = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "120"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.NoError(t, err)

	return
}

func TestCreateOrder(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)

	testdata, _, denomA, denomB, _ := common(t, testInput)

	beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)

	//Create Order
	var o = types.MsgCreateOrder{Creator: addr, DenomAsk: denomA, DenomBid: denomB, Rate: testdata.RateAstrArray, OrderType: "stop", Amount: "0", Prev: "0", Next: "0"}
	rate, _ := types.RateStringToInt(o.Rate)
	bookends := testInput.MarketKeeper.BookEnds(testInput.Context, o.DenomAsk, o.DenomBid, o.OrderType, rate)
	o.Prev = strconv.FormatUint(bookends[0], 10)
	o.Next = strconv.FormatUint(bookends[1], 10)
	_, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &o)
	require.NoError(t, err)

	aftercount := testInput.MarketKeeper.GetUidCount(testInput.Context)

	require.Equal(t, beforecount+1, aftercount)

	//Validate Order
	orders, orderfound := testInput.MarketKeeper.GetOrder(testInput.Context, beforecount)
	require.True(t, orderfound)
	require.Equal(t, orders.DenomBid, denomB)
	require.Equal(t, orders.DenomAsk, denomA)
	require.Equal(t, orders.Amount.String(), o.Amount)

	// Validate GetMember
	memberAsk, memberAskfound := testInput.MarketKeeper.GetMember(testInput.Context, orders.DenomBid, orders.DenomAsk)

	require.True(t, memberAskfound)
	require.Equal(t, memberAsk.DenomA, denomB)
	require.Equal(t, memberAsk.DenomB, denomA)
	require.Equal(t, "33", memberAsk.Balance.String())
	require.Equal(t, memberAsk.Stop, uint64(0))

}

func TestBookEnds(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)

	testdata, _, denomA, denomB, _ := common(t, testInput)

	beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)

	orderType1 := "limit"

	//Create Order
	var o = types.MsgCreateOrder{Creator: addr, DenomAsk: denomA, DenomBid: denomB, Rate: testdata.RateAstrArray, OrderType: orderType1, Amount: "10", Prev: "0", Next: "0"}
	rate, err := types.RateStringToInt(o.Rate)
	require.NoError(t, err)
	ends := testInput.MarketKeeper.BookEnds(testInput.Context, o.DenomAsk, o.DenomBid, o.OrderType, rate)
	require.NoError(t, err)
	o.Prev = strconv.FormatUint(ends[0], 10)
	o.Next = strconv.FormatUint(ends[1], 10)
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &o)
	require.NoError(t, err)
	aftercount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)

	// Validate Order
	orders, orderfound := testInput.MarketKeeper.GetOrder(testInput.Context, beforecount)
	require.True(t, orderfound)
	require.Equal(t, orders.DenomBid, denomB)
	require.Equal(t, orders.DenomAsk, denomA)
	require.Equal(t, orders.Amount.String(), o.Amount)

	// Create Order Msg Type
	beforecount = aftercount
	var q = types.MsgCreateOrder{Creator: addr, DenomAsk: denomA, DenomBid: denomB, Rate: testdata.RateAstrArray, OrderType: orderType1, Amount: "10", Prev: "0", Next: "0"}
	rate, err = types.RateStringToInt(q.Rate)
	require.NoError(t, err)

	// Get Bookends
	ends = testInput.MarketKeeper.BookEnds(testInput.Context, q.DenomAsk, q.DenomBid, q.OrderType, rate)
	require.NoError(t, err)
	q.Prev = strconv.FormatUint(ends[0], 10)
	q.Next = strconv.FormatUint(ends[1], 10)

	// Create Order
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &q)
	require.NoError(t, err)
	aftercount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)

	// Validate Order
	orders2, orderfound2 := testInput.MarketKeeper.GetOrder(testInput.Context, beforecount)
	require.True(t, orderfound2)
	require.Equal(t, orders2.DenomBid, denomB)
	require.Equal(t, orders2.DenomAsk, denomA)
	require.Equal(t, orders2.Amount.String(), o.Amount)

	// Create Order Msg Type
	beforecount = aftercount
	var r = types.MsgCreateOrder{Creator: addr, DenomAsk: denomA, DenomBid: denomB, Rate: []string{"1", "1000"}, OrderType: orderType1, Amount: "10", Prev: "0", Next: "0"}
	rate, err = types.RateStringToInt(r.Rate)
	require.NoError(t, err)

	timeout := time.After(10 * time.Second)
	done := make(chan bool)
	go func() {
		// Get Bookends
		ends = testInput.MarketKeeper.BookEnds(testInput.Context, r.DenomAsk, r.DenomBid, r.OrderType, rate)
		require.NoError(t, err)
		r.Prev = strconv.FormatUint(ends[0], 10)
		r.Next = strconv.FormatUint(ends[1], 10)
		time.Sleep(5 * time.Second)
		done <- true
	}()

	select {
	case <-timeout:
		t.Fatal("Test didn't finish in time")
	case <-done:
	}

	// Get Bookends
	ends = testInput.MarketKeeper.BookEnds(testInput.Context, r.DenomAsk, r.DenomBid, r.OrderType, rate)
	require.NoError(t, err)
	r.Prev = strconv.FormatUint(ends[0], 10)
	r.Next = strconv.FormatUint(ends[1], 10)

	// Create Order
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &r)
	require.NoError(t, err)
	aftercount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)

	// Validate Order
	orders3, orderfound3 := testInput.MarketKeeper.GetOrder(testInput.Context, beforecount)
	require.True(t, orderfound3)
	require.Equal(t, orders3.DenomBid, denomB)
	require.Equal(t, orders3.DenomAsk, denomA)
	require.Equal(t, orders3.Amount.String(), o.Amount)

	// Validate GetMember
	memberAsk, memberAskfound := testInput.MarketKeeper.GetMember(testInput.Context, orders.DenomBid, orders.DenomAsk)

	require.True(t, memberAskfound)
	require.Equal(t, memberAsk.DenomA, denomB)
	require.Equal(t, memberAsk.DenomB, denomA)
	require.Equal(t, "33", memberAsk.Balance.String())
	require.Equal(t, memberAsk.Stop, uint64(0))

}

func TestCreateOrder_Scenarios(t *testing.T) {

	testInput := keepertest.CreateTestEnvironment(t)

	common(t, testInput)

	// beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)

	scenarios := []types.MsgCreateOrder{
		{Creator: addr, DenomAsk: "CoinA", DenomBid: "CoinB", Rate: []string{"60", "70"}, OrderType: "stop", Amount: "10", Prev: "0", Next: "0"},
		{Creator: addr, DenomAsk: "CoinA", DenomBid: "CoinB", Rate: []string{"60", "70"}, OrderType: "limit", Amount: "10", Prev: "0", Next: "0"},
	}

	for _, s := range scenarios {

		_, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &s)
		require.NoError(t, err)
		// require.ErrorContains(t, err, "insufficient balance")
	}

}
