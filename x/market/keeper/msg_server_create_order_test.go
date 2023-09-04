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

	// Create Order Msg Type
	beforecount = aftercount
	var s = types.MsgCreateOrder{Creator: addr, DenomAsk: denomA, DenomBid: denomB, Rate: []string{"65", "70"}, OrderType: orderType1, Amount: "10", Prev: "0", Next: "0"}
	rate, err = types.RateStringToInt(s.Rate)
	require.NoError(t, err)

	timeout = time.After(10 * time.Second)
	done = make(chan bool)
	go func() {
		// Get Bookends
		ends = testInput.MarketKeeper.BookEnds(testInput.Context, s.DenomAsk, s.DenomBid, s.OrderType, rate)
		require.NoError(t, err)
		s.Prev = strconv.FormatUint(ends[0], 10)
		s.Next = strconv.FormatUint(ends[1], 10)
		time.Sleep(5 * time.Second)
		done <- true
	}()

	select {
	case <-timeout:
		t.Fatal("Test didn't finish in time")
	case <-done:
	}

	// Create Order
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &s)
	require.NoError(t, err)
	aftercount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)

	// Validate Order
	orders4, orderfound4 := testInput.MarketKeeper.GetOrder(testInput.Context, beforecount)
	require.True(t, orderfound4)
	require.Equal(t, orders4.DenomBid, denomB)
	require.Equal(t, orders4.DenomAsk, denomA)
	require.Equal(t, orders4.Amount.String(), o.Amount)

}

func TestCreateOrder_BothFillOverlap(t *testing.T) {

	testInput := keepertest.CreateTestEnvironment(t)

	_, _, _, _, pair := common(t, testInput)

	// beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)

	scenarios := []types.MsgCreateOrder{
		{Creator: addr, DenomAsk: "CoinA", DenomBid: "CoinB", Rate: []string{"3", "4"}, OrderType: "limit", Amount: "40", Prev: "0", Next: "0"},
		{Creator: addr, DenomAsk: "CoinB", DenomBid: "CoinA", Rate: []string{"4", "3"}, OrderType: "limit", Amount: "30", Prev: "0", Next: "0"},
	}

	var uid uint64

	for _, s := range scenarios {

		orderresponse, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &s)
		require.NoError(t, err)
		uid = orderresponse.Uid
	}

	order, found := testInput.MarketKeeper.GetOrder(testInput.Context, uid)
	require.True(t, found)
	require.True(t, order.Status == "filled")

	allorders := testInput.MarketKeeper.GetAllOrder(testInput.Context)
	require.Truef(t, allorders[0].Uid == 3, strconv.FormatUint(allorders[0].Uid, 10))
	require.Truef(t, allorders[0].Status == "filled", allorders[0].Status)
	require.Equal(t, sdk.NewInt(70).String(), allorders[0].Amount.Add(allorders[1].Amount).String())

	history, _ := testInput.MarketKeeper.GetHistory(testInput.Context, "CoinA,CoinB", "10")
	require.Equal(t, "40", history[0].Amount)
	require.Equal(t, "30", history[1].Amount)

	require.True(t, len(allorders) == 2)

	// Validate Order
	orderowner := testInput.MarketKeeper.GetOrderOwner(testInput.Context, addr)
	require.True(t, len(orderowner) == 0)

	// Validate GetPool
	pool, _ := testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.Equal(t, strconv.FormatUint(pool.History, 10), strconv.FormatUint(allorders[0].Uid, 10))
}

func TestCreateOrder_OneSide1FillOverlap(t *testing.T) {

	testInput := keepertest.CreateTestEnvironment(t)

	_, _, _, _, pair := common(t, testInput)

	// beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)

	scenarios := []types.MsgCreateOrder{
		{Creator: addr, DenomAsk: "CoinA", DenomBid: "CoinB", Rate: []string{"3", "4"}, OrderType: "limit", Amount: "50", Prev: "0", Next: "0"},
		{Creator: addr, DenomAsk: "CoinB", DenomBid: "CoinA", Rate: []string{"4", "3"}, OrderType: "limit", Amount: "30", Prev: "0", Next: "0"},
	}

	var uid uint64

	for _, s := range scenarios {

		orderresponse, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &s)
		require.NoError(t, err)
		uid = orderresponse.Uid
	}

	order, found := testInput.MarketKeeper.GetOrder(testInput.Context, uid)
	require.True(t, found)
	require.True(t, order.Status == "filled")

	allorders := testInput.MarketKeeper.GetAllOrder(testInput.Context)
	require.Truef(t, allorders[0].Uid == 3, strconv.FormatUint(allorders[0].Uid, 10))
	require.Truef(t, allorders[0].Status == "active", allorders[0].Status)
	require.Equal(t, sdk.NewInt(10).String(), allorders[0].Amount.String())

	history, _ := testInput.MarketKeeper.GetHistory(testInput.Context, "CoinA,CoinB", "10")
	require.Equal(t, "40", history[0].Amount)
	require.Equal(t, "30", history[1].Amount)

	require.Equal(t, 3, len(allorders))

	// Validate Order
	orderowner := testInput.MarketKeeper.GetOrderOwner(testInput.Context, addr)
	require.True(t, len(orderowner) == 1)

	// Validate GetPool
	pool, _ := testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.Equal(t, strconv.FormatUint(pool.History, 10), strconv.FormatUint(allorders[2].Uid, 10))

	member, found := testInput.MarketKeeper.GetMember(testInput.Context, "CoinA", "CoinB")
	require.True(t, found)
	require.Equal(t, member.Limit, allorders[0].Uid)
}

func TestCreateOrder_OneSide2FillOverlap(t *testing.T) {

	testInput := keepertest.CreateTestEnvironment(t)

	_, _, _, _, pair := common(t, testInput)

	// beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)

	scenarios := []types.MsgCreateOrder{
		{Creator: addr, DenomAsk: "CoinA", DenomBid: "CoinB", Rate: []string{"3", "4"}, OrderType: "limit", Amount: "40", Prev: "0", Next: "0"},
		{Creator: addr, DenomAsk: "CoinB", DenomBid: "CoinA", Rate: []string{"4", "3"}, OrderType: "limit", Amount: "40", Prev: "0", Next: "0"},
	}

	var uid uint64

	for _, s := range scenarios {

		orderresponse, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &s)
		require.NoError(t, err)
		uid = orderresponse.Uid
	}

	order, found := testInput.MarketKeeper.GetOrder(testInput.Context, uid)
	require.True(t, found)
	require.True(t, order.Status == "active")

	allorders := testInput.MarketKeeper.GetAllOrder(testInput.Context)
	require.Truef(t, allorders[0].Uid == 3, strconv.FormatUint(allorders[0].Uid, 10))
	require.Truef(t, allorders[1].Status == "active", allorders[1].Status)
	require.Equal(t, sdk.NewInt(10).String(), allorders[1].Amount.String())

	history, _ := testInput.MarketKeeper.GetHistory(testInput.Context, "CoinA,CoinB", "10")
	require.Equal(t, "30", history[0].Amount)
	require.Equal(t, "40", history[1].Amount)

	require.Equal(t, 3, len(allorders))

	// Validate Order
	orderowner := testInput.MarketKeeper.GetOrderOwner(testInput.Context, addr)
	require.True(t, len(orderowner) == 1)

	// Validate GetPool
	pool, _ := testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.Equal(t, strconv.FormatUint(pool.History, 10), strconv.FormatUint(allorders[2].Uid, 10))

	member, found := testInput.MarketKeeper.GetMember(testInput.Context, "CoinB", "CoinA")
	require.True(t, found)
	require.Equal(t, member.Limit, allorders[1].Uid)

	bookends := testInput.MarketKeeper.BookEnds(testInput.Context, "CoinB", "CoinA", "limit", []sdk.Int{sdk.NewInt(3), sdk.NewInt(3)})
	require.Equal(t, strconv.FormatUint(uint64(0), 10), strconv.FormatUint(bookends[0], 10))
}
