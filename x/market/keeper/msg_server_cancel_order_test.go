package keeper_test

import (
	"strconv"
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/pendulum-labs/market/testutil/keeper"
	"github.com/pendulum-labs/market/testutil/sample"
	"github.com/pendulum-labs/market/x/market/keeper"
	"github.com/pendulum-labs/market/x/market/types"
	"github.com/stretchr/testify/require"
)

func TestCancelOrder_case1_stop(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	// TestData
	testdata := testData{coinAStr: "30CoinA", coinBStr: "40CoinB", RateAstrArray: []string{"60", "70"}, RateBstrArray: []string{"80", "90"}}
	coinPair, _ := sample.SampleCoins("70CoinA", "70CoinB")
	denomA, denomB := sample.SampleDenoms(coinPair)
	pair := strings.Join([]string{denomA, denomB}, ",")

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

	beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)

	// Create Order
	var o = types.MsgCreateOrder{Creator: addr, DenomAsk: denomA, DenomBid: denomB, Rate: testdata.RateAstrArray, OrderType: "stop", Amount: "0", Prev: "0", Next: "0"}
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

	// Validate GetMember
	memberA, memberfound := testInput.MarketKeeper.GetMember(testInput.Context, orders.DenomBid, orders.DenomAsk)

	require.True(t, memberfound)
	require.Equal(t, memberA.DenomA, denomB)
	require.Equal(t, memberA.DenomB, denomA)
	require.Equal(t, "33", memberA.Balance.String())
	require.Equal(t, memberA.Stop, uint64(0))

	// Cancel Order
	Uid := strconv.FormatUint(orders.Uid, 10)
	var co = types.MsgCancelOrder{Creator: addr, Uid: Uid}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CancelOrder(sdk.WrapSDKContext(testInput.Context), &co)
	require.NoError(t, err)

	// Validate GetMember
	memberA, memberfound = testInput.MarketKeeper.GetMember(testInput.Context, orders.DenomBid, orders.DenomAsk)
	require.True(t, memberfound)
	require.Equal(t, memberA.DenomA, denomB)
	require.Equal(t, memberA.DenomB, denomA)
	require.Equal(t, "33", memberA.Balance.String())
	require.Equal(t, memberA.Stop, uint64(0))

	//Validate Order
	orders, orderfound = testInput.MarketKeeper.GetOrder(testInput.Context, beforecount)
	require.True(t, orderfound)
	require.False(t, orders.Active)
	require.Equal(t, orders.DenomBid, denomB)
	require.Equal(t, orders.DenomAsk, denomA)
	require.Equal(t, orders.Amount.String(), o.Amount)
	require.Equal(t, o.OrderType, "stop")

}

func TestCancelOrder_case1_limit(t *testing.T) {

	testInput := keepertest.CreateTestEnvironment(t)
	// TestData
	testdata := testData{coinAStr: "30CoinA", coinBStr: "40CoinB", RateAstrArray: []string{"60", "70"}, RateBstrArray: []string{"80", "90"}}
	coinPair, _ := sample.SampleCoins("70CoinA", "70CoinB")
	denomA, denomB := sample.SampleDenoms(coinPair)
	pair := strings.Join([]string{denomA, denomB}, ",")

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

	beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)

	// Create Order
	var o = types.MsgCreateOrder{Creator: addr, DenomAsk: denomA, DenomBid: denomB, Rate: testdata.RateAstrArray, OrderType: "limit", Amount: "0", Prev: "0", Next: "0"}
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

	//Validate GetMember
	memberBid, memberfoundBid := testInput.MarketKeeper.GetMember(testInput.Context, orders.DenomAsk, orders.DenomBid)
	require.True(t, memberfoundBid)
	require.Equal(t, memberBid.DenomA, denomA)
	require.Equal(t, memberBid.DenomB, denomB)
	require.Equal(t, "44", memberBid.Balance.String())
	require.Equal(t, memberBid.Stop, uint64(0))

	memberBid.Stop = orders.Uid
	testInput.MarketKeeper.SetMember(testInput.Context, memberBid)

	//Cancel Order
	Uid := strconv.FormatUint(orders.Uid, 10)
	var co = types.MsgCancelOrder{Creator: addr, Uid: Uid}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CancelOrder(sdk.WrapSDKContext(testInput.Context), &co)
	require.NoError(t, err)

	memberBid, memberfoundBid = testInput.MarketKeeper.GetMember(testInput.Context, orders.DenomAsk, orders.DenomBid)
	require.True(t, memberfoundBid)
	require.Equal(t, memberBid.DenomA, denomA)
	require.Equal(t, memberBid.DenomB, denomB)
	require.Equal(t, "44", memberBid.Balance.String())
	require.Equal(t, memberBid.Stop, uint64(orders.Uid))

	//Validate Order
	orders, orderfound = testInput.MarketKeeper.GetOrder(testInput.Context, beforecount)
	require.True(t, orderfound)
	require.False(t, orders.Active)
	require.Equal(t, orders.DenomBid, denomB)
	require.Equal(t, orders.DenomAsk, denomA)
	require.Equal(t, orders.Amount.String(), o.Amount)
	require.Equal(t, orders.OrderType, "limit")

}

func TestCancelOrder_case2_stop(t *testing.T) {

	testInput := keepertest.CreateTestEnvironment(t)
	// TestData
	testdata := testData{coinAStr: "30CoinA", coinBStr: "40CoinB", RateAstrArray: []string{"60", "70"}, RateBstrArray: []string{"80", "90"}}
	coinPair, _ := sample.SampleCoins("70CoinA", "70CoinB")
	denomA, denomB := sample.SampleDenoms(coinPair)
	pair := strings.Join([]string{denomA, denomB}, ",")

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

	beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)

	//Create Order
	var o = types.MsgCreateOrder{Creator: addr, DenomAsk: denomA, DenomBid: denomB, Rate: testdata.RateAstrArray, OrderType: "stop", Amount: "0", Prev: "0", Next: "0"}
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

	// Validate GetMember
	memberAsk, memberAskfound := testInput.MarketKeeper.GetMember(testInput.Context, orders.DenomBid, orders.DenomAsk)

	require.True(t, memberAskfound)
	require.Equal(t, memberAsk.DenomA, denomB)
	require.Equal(t, memberAsk.DenomB, denomA)
	require.Equal(t, "33", memberAsk.Balance.String())
	require.Equal(t, memberAsk.Stop, uint64(0))

	o.Next = strconv.FormatUint(beforecount, 10)
	o.Rate = []string{"70", "80"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &o)
	require.NoError(t, err)

	orders, orderfound = testInput.MarketKeeper.GetOrder(testInput.Context, aftercount)
	require.True(t, orderfound)
	require.Equal(t, orders.Next, beforecount)

	// Cancel Order
	Uid := strconv.FormatUint(orders.Uid, 10)
	var co = types.MsgCancelOrder{Creator: addr, Uid: Uid}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CancelOrder(sdk.WrapSDKContext(testInput.Context), &co)
	require.NoError(t, err)

	memberBid, memberBidfound := testInput.MarketKeeper.GetMember(testInput.Context, orders.DenomAsk, orders.DenomBid)
	require.True(t, memberBidfound)
	require.Equal(t, memberBid.DenomA, denomA)
	require.Equal(t, memberBid.DenomB, denomB)
	require.Equal(t, "44", memberBid.Balance.String())
	require.Equal(t, memberBid.Stop, beforecount)

	// Validate Order
	orders, orderfound = testInput.MarketKeeper.GetOrder(testInput.Context, aftercount)
	require.True(t, orderfound)
	require.False(t, orders.Active)
	require.Equal(t, orders.DenomBid, denomB)
	require.Equal(t, orders.DenomAsk, denomA)
	require.Equal(t, orders.Amount.String(), o.Amount)
	require.Equal(t, orders.OrderType, "stop")
}

/*
func TestCancelOrder_case2_limit(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "30CoinA", coinBStr: "40CoinB", RateAstrArray: []string{"60", "70"}, RateBstrArray: []string{"80", "90"}}
	coinPair, _ := sample.SampleCoins("70CoinA", "70CoinB")
	denomA, denomB := sample.SampleDenoms(coinPair)
	pair := strings.Join([]string{denomA, denomB}, ",")

	//MintCoins
	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, types.ModuleName, coinPair))
	//SendCoinsFromModuleToAccount
	requestAddress, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, types.ModuleName, requestAddress, coinPair))
	// GetUidCount before CreatePool
	beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	//Create Pool
	var p = types.MsgCreatePool{CoinA: testdata.coinAStr, CoinB: testdata.coinBStr, Creator: addr}
	response, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	//validate CreatePool
	require.NoError(t, err)
	require.Contains(t, p.GetCreator(), response.String())
	require.Contains(t, p.GetCoinA(), response.String())
	require.Contains(t, p.GetCoinB(), response.String())
	//validate SetUidCount function.
	aftercount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)

	//validate GetDrop
	drops, dropFound := testInput.MarketKeeper.GetDrop(testInput.Context, beforecount)
	require.True(t, dropFound)
	require.Equal(t, drops.Pair, pair)
	//validate CreateDrop
	var d = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "70"}
	createDropResponse, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.NoError(t, err)

	//validate GetMember
	members, memberfound := testInput.MarketKeeper.GetMember(testInput.Context, denomB, denomA)
	members1, memberfound1 := testInput.MarketKeeper.GetMember(testInput.Context, denomA, denomB)
	require.True(t, memberfound)
	require.Equal(t, members.DenomA, denomB)
	require.Equal(t, members.DenomB, denomA)
	//require.Equal(t, members.Balance.String(), "70")
	require.True(t, memberfound1)
	require.Equal(t, members1.DenomA, denomA)
	require.Equal(t, members1.DenomB, denomB)
	//require.Equal(t, members1.Balance.String(), "70")
	//Validate GetPool
	rst, found := testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, found)
	require.Equal(t, rst.Pair, pair)
	//require.Equal(t, rst.Drops.String(), "140")
	//validate GetDrop
	drops1, drop1Found := testInput.MarketKeeper.GetDrop(testInput.Context, aftercount)
	require.True(t, drop1Found)
	require.Equal(t, drops1.Pair, pair)
	require.Equal(t, drops1.Drops.String(), d.Drops)
	require.Contains(t, d.GetCreator(), createDropResponse.String())
	// GetUidCount before CreatePool
	beforecount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	//Create Order
	var o = types.MsgCreateOrder{Creator: addr, DenomAsk: members1.DenomA, DenomBid: members1.DenomB, Rate: testdata.RateAstrArray, OrderType: "limit", Amount: "0", Prev: "0", Next: "0"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &o)
	require.NoError(t, err)
	aftercount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)
	//Validate Order
	orders, orderfound := testInput.MarketKeeper.GetOrder(testInput.Context, beforecount)
	require.True(t, orderfound)
	require.Equal(t, orders.DenomBid, denomB)
	require.Equal(t, orders.DenomAsk, denomA)
	require.Equal(t, orders.Amount.String(), o.Amount)

	//validate GetMember
	members, memberfound = testInput.MarketKeeper.GetMember(testInput.Context, orders.DenomBid, orders.DenomAsk)

	require.True(t, memberfound)
	require.Equal(t, members.DenomA, denomB)
	require.Equal(t, members.DenomB, denomA)
	//require.Equal(t, members.Balance.String(), "70")
	require.Equal(t, members.Stop, uint64(0))
	require.Equal(t, members.Limit, uint64(0))

	o.Next = "0"
	o.Rate = []string{"50", "60"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &o)
	require.NoError(t, err)

	orders, orderfound = testInput.MarketKeeper.GetOrder(testInput.Context, aftercount)
	require.True(t, orderfound)
	require.Equal(t, orders.Next, uint64(0))

	members1.Stop = orders.Uid
	testInput.MarketKeeper.SetMember(testInput.Context, members1)

	//Cancel Order
	Uid := strconv.FormatUint(orders.Uid, 10)
	var co = types.MsgCancelOrder{Creator: addr, Uid: Uid}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CancelOrder(sdk.WrapSDKContext(testInput.Context), &co)
	require.NoError(t, err)

	members1, memberfound1 = testInput.MarketKeeper.GetMember(testInput.Context, orders.DenomAsk, orders.DenomBid)
	require.True(t, memberfound1)
	require.Equal(t, members1.DenomA, denomA)
	require.Equal(t, members1.DenomB, denomB)
	require.Equal(t, members1.Balance.String(), "70")
	require.Equal(t, members1.Stop, orders.Uid)
	require.Equal(t, members1.Limit, uint64(0))

	//Validate Order
	orders, orderfound = testInput.MarketKeeper.GetOrder(testInput.Context, aftercount)
	require.True(t, orderfound)
	require.False(t, orders.Active)
	require.Equal(t, orders.DenomBid, denomB)
	require.Equal(t, orders.DenomAsk, denomA)
	require.Equal(t, orders.Amount.String(), o.Amount)
	require.Equal(t, orders.OrderType, "limit")

}

func TestCancelOrder_case3_stop(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "30CoinA", coinBStr: "40CoinB", RateAstrArray: []string{"60", "70"}, RateBstrArray: []string{"80", "90"}}
	coinPair, _ := sample.SampleCoins("70CoinA", "70CoinB")
	denomA, denomB := sample.SampleDenoms(coinPair)
	pair := strings.Join([]string{denomA, denomB}, ",")

	//MintCoins
	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, types.ModuleName, coinPair))
	//SendCoinsFromModuleToAccount
	requestAddress, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, types.ModuleName, requestAddress, coinPair))
	// GetUidCount before CreatePool
	beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	//Create Pool
	var p = types.MsgCreatePool{CoinA: testdata.coinAStr, CoinB: testdata.coinBStr, Creator: addr}
	response, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	//validate CreatePool
	require.NoError(t, err)
	require.Contains(t, p.GetCreator(), response.String())
	require.Contains(t, p.GetCoinA(), response.String())
	require.Contains(t, p.GetCoinB(), response.String())
	//validate SetUidCount function.
	aftercount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)

	//validate GetDrop
	drops, dropFound := testInput.MarketKeeper.GetDrop(testInput.Context, beforecount)
	require.True(t, dropFound)
	require.Equal(t, drops.Pair, pair)
	//validate CreateDrop
	var d = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "70"}
	createDropResponse, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.NoError(t, err)

	//validate GetMember
	members, memberfound := testInput.MarketKeeper.GetMember(testInput.Context, denomB, denomA)
	members1, memberfound1 := testInput.MarketKeeper.GetMember(testInput.Context, denomA, denomB)
	require.True(t, memberfound)
	require.Equal(t, members.DenomA, denomB)
	require.Equal(t, members.DenomB, denomA)
	require.Equal(t, members.Balance.String(), "70")
	require.True(t, memberfound1)
	require.Equal(t, members1.DenomA, denomA)
	require.Equal(t, members1.DenomB, denomB)
	require.Equal(t, members1.Balance.String(), "70")

	//Validate GetPool
	rst, found := testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, found)
	require.Equal(t, rst.Pair, pair)
	require.Equal(t, rst.Drops.String(), "140")
	//validate GetDrop
	drops1, drop1Found := testInput.MarketKeeper.GetDrop(testInput.Context, aftercount)
	require.True(t, drop1Found)
	require.Equal(t, drops1.Pair, pair)
	require.Equal(t, drops1.Drops.String(), d.Drops)
	require.Contains(t, d.GetCreator(), createDropResponse.String())
	// GetUidCount before CreatePool
	beforecount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	//Create Order
	var o = types.MsgCreateOrder{Creator: addr, DenomAsk: members1.DenomA, DenomBid: members1.DenomB, Rate: testdata.RateAstrArray, OrderType: "stop", Amount: "0", Prev: "0", Next: "0"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &o)
	require.NoError(t, err)
	aftercount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)
	//Validate Order
	orders, orderfound := testInput.MarketKeeper.GetOrder(testInput.Context, beforecount)
	require.True(t, orderfound)
	require.Equal(t, orders.DenomBid, denomB)
	require.Equal(t, orders.DenomAsk, denomA)
	require.Equal(t, orders.Amount.String(), o.Amount)

	//validate GetMember
	members, memberfound = testInput.MarketKeeper.GetMember(testInput.Context, orders.DenomBid, orders.DenomAsk)

	require.True(t, memberfound)
	require.Equal(t, members.DenomA, denomB)
	require.Equal(t, members.DenomB, denomA)
	require.Equal(t, members.Balance.String(), "70")

	require.Equal(t, members.Stop, uint64(0))

	o.Prev = strconv.FormatUint(beforecount, 10)
	o.Rate = []string{"50", "60"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &o)
	require.NoError(t, err)

	orders, orderfound = testInput.MarketKeeper.GetOrder(testInput.Context, aftercount)
	require.True(t, orderfound)
	require.Equal(t, orders.Next, uint64(0))
	//Cancel Order
	Uid := strconv.FormatUint(orders.Uid, 10)
	var co = types.MsgCancelOrder{Creator: addr, Uid: Uid}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CancelOrder(sdk.WrapSDKContext(testInput.Context), &co)
	require.NoError(t, err)

	members1, memberfound1 = testInput.MarketKeeper.GetMember(testInput.Context, orders.DenomAsk, orders.DenomBid)
	require.True(t, memberfound1)
	require.Equal(t, members1.DenomA, denomA)
	require.Equal(t, members1.DenomB, denomB)
	require.Equal(t, members1.Balance.String(), "70")

	require.Equal(t, members1.Stop, beforecount)

	//Validate Order
	orders, orderfound = testInput.MarketKeeper.GetOrder(testInput.Context, aftercount)
	require.True(t, orderfound)
	require.False(t, orders.Active)
	require.Equal(t, orders.DenomBid, denomB)
	require.Equal(t, orders.DenomAsk, denomA)
	require.Equal(t, orders.Amount.String(), o.Amount)
	require.Equal(t, orders.OrderType, "stop")
}

func TestCancelOrder_case3_limit(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "50CoinA", coinBStr: "60CoinB", RateAstrArray: []string{"60", "70"}, RateBstrArray: []string{"80", "90"}}
	coinPair, _ := sample.SampleCoins("80CoinA", "80CoinB")
	denomA, denomB := sample.SampleDenoms(coinPair)
	pair := strings.Join([]string{denomA, denomB}, ",")

	//MintCoins
	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, types.ModuleName, coinPair))
	//SendCoinsFromModuleToAccount
	requestAddress, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, types.ModuleName, requestAddress, coinPair))
	// GetUidCount before CreatePool
	beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	//Create Pool
	var p = types.MsgCreatePool{CoinA: testdata.coinAStr, CoinB: testdata.coinBStr, Creator: addr}
	response, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	//validate CreatePool
	require.NoError(t, err)
	require.Contains(t, p.GetCreator(), response.String())
	require.Contains(t, p.GetCoinA(), response.String())
	require.Contains(t, p.GetCoinB(), response.String())
	//validate SetUidCount function.
	aftercount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)

	//validate GetDrop
	drops, dropFound := testInput.MarketKeeper.GetDrop(testInput.Context, beforecount)
	require.True(t, dropFound)
	require.Equal(t, drops.Pair, pair)
	require.Equal(t, "110", drops.Drops.String())

	//validate GetMember
	membersA, memberfound := testInput.MarketKeeper.GetMember(testInput.Context, denomB, denomA)
	require.True(t, memberfound)
	require.Equal(t, membersA.DenomA, denomB)
	require.Equal(t, membersA.DenomB, denomA)
	require.Equal(t, "50", membersA.Balance.String())

	//validate CreateDrop
	var d = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "20"}
	createDropResponse, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.NoError(t, err)

	//validate GetMember
	members, memberfound := testInput.MarketKeeper.GetMember(testInput.Context, denomB, denomA)
	members1, memberfound1 := testInput.MarketKeeper.GetMember(testInput.Context, denomA, denomB)
	require.True(t, memberfound)
	require.Equal(t, members.DenomA, denomB)
	require.Equal(t, members.DenomB, denomA)
	require.Equal(t, "60", members.Balance.String())

	require.True(t, memberfound1)
	require.Equal(t, members1.DenomA, denomA)
	require.Equal(t, members1.DenomB, denomB)
	require.Equal(t, members1.Balance.String(), "70")

	//Validate GetPool
	rst, found := testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, found)
	require.Equal(t, rst.Pair, pair)
	require.Equal(t, rst.Drops.String(), "130")
	//validate GetDrop
	drops1, drop1Found := testInput.MarketKeeper.GetDrop(testInput.Context, aftercount)
	require.True(t, drop1Found)
	require.Equal(t, drops1.Pair, pair)
	require.Equal(t, drops1.Drops.String(), d.Drops)
	require.Contains(t, d.GetCreator(), createDropResponse.String())
	// GetUidCount before CreatePool
	beforecount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	//Create Order
	var o = types.MsgCreateOrder{Creator: addr, DenomAsk: members1.DenomA, DenomBid: members1.DenomB, Rate: testdata.RateAstrArray, OrderType: "limit", Amount: "1", Prev: "0", Next: "0"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &o)
	require.NoError(t, err)
	aftercount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)
	//Validate Order
	orders, orderfound := testInput.MarketKeeper.GetOrder(testInput.Context, beforecount)
	require.True(t, orderfound)
	require.Equal(t, orders.DenomBid, denomB)
	require.Equal(t, orders.DenomAsk, denomA)
	require.Equal(t, orders.Amount.String(), o.Amount)

	//validate GetMember
	members, memberfound = testInput.MarketKeeper.GetMember(testInput.Context, orders.DenomBid, orders.DenomAsk)

	require.True(t, memberfound)
	require.Equal(t, members.DenomA, denomB)
	require.Equal(t, members.DenomB, denomA)
	require.Equal(t, "60", members.Balance.String())

	require.Equal(t, members.Stop, uint64(0))

	o.Prev = strconv.FormatUint(beforecount, 10)
	o.Rate = []string{"70", "80"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &o)
	require.NoError(t, err)

	orders, orderfound = testInput.MarketKeeper.GetOrder(testInput.Context, aftercount)
	require.True(t, orderfound)
	require.Equal(t, orders.Next, uint64(0))

	//members1.Stop = orders.Uid
	//testInput.MarketKeeper.SetMember(testInput.Context, members1)

	//Cancel Order
	Uid := strconv.FormatUint(orders.Uid, 10)
	var co = types.MsgCancelOrder{Creator: addr, Uid: Uid}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CancelOrder(sdk.WrapSDKContext(testInput.Context), &co)
	require.NoError(t, err)

	members1, memberfound1 = testInput.MarketKeeper.GetMember(testInput.Context, orders.DenomAsk, orders.DenomBid)
	require.True(t, memberfound1)
	require.Equal(t, members1.DenomA, denomA)
	require.Equal(t, members1.DenomB, denomB)
	require.Equal(t, members1.Balance.String(), "70")

	require.Equal(t, members1.Stop, uint64(0))
	require.Equal(t, beforecount, members1.Limit)

	//Validate Order
	orders, orderfound = testInput.MarketKeeper.GetOrder(testInput.Context, aftercount)
	require.True(t, orderfound)
	require.False(t, orders.Active)
	require.Equal(t, orders.DenomBid, denomB)
	require.Equal(t, orders.DenomAsk, denomA)
	require.Equal(t, orders.Amount.String(), o.Amount)
	require.Equal(t, orders.OrderType, "limit")

}


func TestCancelOrder_case4_stop(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "30CoinA", coinBStr: "40CoinB", RateAstrArray: []string{"70", "80"}, RateBstrArray: []string{"90", "100"}}
	coinPair, _ := sample.SampleCoins("70CoinA", "70CoinB")
	denomA, denomB := sample.SampleDenoms(coinPair)
	pair := strings.Join([]string{denomA, denomB}, ",")

	//MintCoins
	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, types.ModuleName, coinPair))
	//SendCoinsFromModuleToAccount
	requestAddress, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, types.ModuleName, requestAddress, coinPair))
	// GetUidCount before CreatePool
	beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	//Create Pool
	var p = types.MsgCreatePool{CoinA: testdata.coinAStr, CoinB: testdata.coinBStr, Creator: addr, RateA: testdata.RateAstrArray, RateB: testdata.RateBstrArray}
	response, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	//validate CreatePool
	require.NoError(t, err)
	require.Contains(t, p.GetCreator(), response.String())
	require.Contains(t, p.GetCoinA(), response.String())
	require.Contains(t, p.GetCoinB(), response.String())
	//validate SetUidCount function.
	aftercount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)

	//validate GetDrop
	drops, dropFound := testInput.MarketKeeper.GetDrop(testInput.Context, beforecount)
	require.True(t, dropFound)
	require.Equal(t, drops.Pair, pair)
	//validate CreateDrop
	var d = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "70", Rate1: testdata.RateAstrArray, Prev1: "0", Next1: "0", Rate2: testdata.RateBstrArray, Prev2: "0", Next2: "0"}
	createDropResponse, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.NoError(t, err)

	//validate GetMember
	members, memberfound := testInput.MarketKeeper.GetMember(testInput.Context, denomB, denomA)
	members1, memberfound1 := testInput.MarketKeeper.GetMember(testInput.Context, denomA, denomB)
	require.True(t, memberfound)
	require.Equal(t, members.DenomA, denomB)
	require.Equal(t, members.DenomB, denomA)
	require.Equal(t, members.Balance.String(), "70")

	require.True(t, memberfound1)
	require.Equal(t, members1.DenomA, denomA)
	require.Equal(t, members1.DenomB, denomB)
	require.Equal(t, members1.Balance.String(), "70")

	//Validate GetPool
	rst, found := testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, found)
	require.Equal(t, rst.Pair, pair)
	require.Equal(t, rst.Drops.String(), "140")
	//validate GetDrop
	drops1, drop1Found := testInput.MarketKeeper.GetDrop(testInput.Context, aftercount)
	require.True(t, drop1Found)
	require.Equal(t, drops1.Pair, pair)
	require.Equal(t, drops1.Drops.String(), d.Drops)
	require.Contains(t, d.GetCreator(), createDropResponse.String())
	// GetUidCount before CreatePool
	beforecount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	//Create Order Prev 0 Next 0
	var o = types.MsgCreateOrder{Creator: addr, DenomAsk: members1.DenomA, DenomBid: members1.DenomB, Rate: testdata.RateAstrArray, OrderType: "stop", Amount: "0", Prev: "0", Next: "0"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &o)
	require.NoError(t, err)
	aftercount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)
	//Validate Order
	orders, orderfound := testInput.MarketKeeper.GetOrder(testInput.Context, beforecount)
	require.True(t, orderfound)
	require.Equal(t, orders.DenomBid, denomB)
	require.Equal(t, orders.DenomAsk, denomA)
	require.Equal(t, orders.Amount.String(), o.Amount)
	require.Equal(t, orders.Prev, uint64(0))
	require.Equal(t, orders.Next, uint64(0))

	//Create Order Prev 0 Next 2
	var o1 = types.MsgCreateOrder{Creator: addr, DenomAsk: members1.DenomA, DenomBid: members1.DenomB, Rate: []string{"80", "90"}, OrderType: "stop", Amount: "0", Prev: "0", Next: "2"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &o1)
	require.NoError(t, err)
	aftercount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+2, aftercount)
	//Validate Order
	orders1, orderfound1 := testInput.MarketKeeper.GetOrder(testInput.Context, beforecount+1)
	require.True(t, orderfound1)
	require.Equal(t, orders1.DenomBid, denomB)
	require.Equal(t, orders1.DenomAsk, denomA)
	require.Equal(t, orders1.Amount.String(), o1.Amount)
	require.Equal(t, orders1.Prev, uint64(0))
	require.Equal(t, orders1.Next, uint64(2))

	//Create Order Prev 3 Next 0
	orders1.Next = uint64(0)
	testInput.MarketKeeper.SetOrder(testInput.Context, orders1)
	var o2 = types.MsgCreateOrder{Creator: addr, DenomAsk: members1.DenomA, DenomBid: members1.DenomB, Rate: []string{"40", "50"}, OrderType: "stop", Amount: "0", Prev: "3", Next: "0"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &o2)
	require.NoError(t, err)
	aftercount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+3, aftercount)
	//Validate Order
	orders2, orderfound2 := testInput.MarketKeeper.GetOrder(testInput.Context, beforecount+2)
	require.True(t, orderfound2)
	require.Equal(t, orders2.DenomBid, denomB)
	require.Equal(t, orders2.DenomAsk, denomA)
	require.Equal(t, orders2.Amount.String(), o2.Amount)
	require.Equal(t, orders2.Prev, uint64(3))
	require.Equal(t, orders2.Next, uint64(0))
	orders1.Prev = uint64(4)
	testdatadrop := testData{RateAstrArray: []string{"30", "40"}}
	numerator1, _ := sdk.NewIntFromString(testdatadrop.RateAstrArray[0])
	denominator1, _ := sdk.NewIntFromString(testdatadrop.RateAstrArray[1])
	orders1.Rate = []sdk.Int{numerator1, denominator1}
	testInput.MarketKeeper.SetOrder(testInput.Context, orders1)
	//Create Order Prev 4 Next 3
	orders2.Next = uint64(3)
	testInput.MarketKeeper.SetOrder(testInput.Context, orders2)
	var o3 = types.MsgCreateOrder{Creator: addr, DenomAsk: members1.DenomA, DenomBid: members1.DenomB, Rate: []string{"40", "50"}, OrderType: "stop", Amount: "0", Prev: "4", Next: "3"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &o3)
	require.NoError(t, err)
	aftercount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+4, aftercount)
	//Validate Order
	orders3, orderfound3 := testInput.MarketKeeper.GetOrder(testInput.Context, beforecount+3)
	require.True(t, orderfound3)
	require.Equal(t, orders3.DenomBid, denomB)
	require.Equal(t, orders3.DenomAsk, denomA)
	require.Equal(t, orders3.Amount.String(), o2.Amount)
	require.Equal(t, orders3.Prev, uint64(4))
	require.Equal(t, orders3.Next, uint64(3))
	require.Equal(t, orders3.OrderType, "stop")

	//Cancel Order
	Uid := strconv.FormatUint(orders3.Uid, 10)
	var co = types.MsgCancelOrder{Creator: addr, Uid: Uid}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CancelOrder(sdk.WrapSDKContext(testInput.Context), &co)
	require.NoError(t, err)

	members1, memberfound1 = testInput.MarketKeeper.GetMember(testInput.Context, orders.DenomAsk, orders.DenomBid)
	require.True(t, memberfound1)
	require.Equal(t, members1.DenomA, denomA)
	require.Equal(t, members1.DenomB, denomB)
	require.Equal(t, members1.Balance.String(), "70")

	require.Equal(t, members1.Stop, beforecount+1)
	//next order validation
	orders, orderfound = testInput.MarketKeeper.GetOrder(testInput.Context, orders1.Uid)
	require.True(t, orderfound)
	require.True(t, orders.Active)
	require.Equal(t, orders.DenomBid, denomB)
	require.Equal(t, orders.DenomAsk, denomA)
	require.Equal(t, orders.Amount.String(), o.Amount)
	require.Equal(t, orders.OrderType, "stop")
	require.Equal(t, orders.Prev, orders3.Prev)
	require.Equal(t, orders.Next, uint64(0))
	//prev order validation
	orders, orderfound = testInput.MarketKeeper.GetOrder(testInput.Context, orders2.Uid)
	require.True(t, orderfound)
	require.True(t, orders.Active)
	require.Equal(t, orders.DenomBid, denomB)
	require.Equal(t, orders.DenomAsk, denomA)
	require.Equal(t, orders.Amount.String(), o.Amount)
	require.Equal(t, orders.OrderType, "stop")
	require.Equal(t, orders.Prev, orders1.Uid)
	require.Equal(t, orders.Next, orders3.Next)

	//Validate Order
	orders, orderfound = testInput.MarketKeeper.GetOrder(testInput.Context, orders3.Uid)
	require.True(t, orderfound)
	require.False(t, orders.Active)
	require.Equal(t, orders.DenomBid, denomB)
	require.Equal(t, orders.DenomAsk, denomA)
	require.Equal(t, orders.Amount.String(), o.Amount)
	require.Equal(t, orders.OrderType, "stop")
	require.Equal(t, orders.Prev, orders3.Prev)
	require.Equal(t, orders.Next, orders3.Next)
}

func TestCancelOrder_case4_limit(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "30CoinA", coinBStr: "40CoinB", RateAstrArray: []string{"70", "80"}, RateBstrArray: []string{"90", "100"}}
	coinPair, _ := sample.SampleCoins("70CoinA", "70CoinB")
	denomA, denomB := sample.SampleDenoms(coinPair)
	pair := strings.Join([]string{denomA, denomB}, ",")

	//MintCoins
	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, types.ModuleName, coinPair))
	//SendCoinsFromModuleToAccount
	requestAddress, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, types.ModuleName, requestAddress, coinPair))
	// GetUidCount before CreatePool
	beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	//Create Pool
	var p = types.MsgCreatePool{CoinA: testdata.coinAStr, CoinB: testdata.coinBStr, Creator: addr, RateA: testdata.RateAstrArray, RateB: testdata.RateBstrArray}
	response, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	//validate CreatePool
	require.NoError(t, err)
	require.Contains(t, p.GetCreator(), response.String())
	require.Contains(t, p.GetCoinA(), response.String())
	require.Contains(t, p.GetCoinB(), response.String())
	//validate SetUidCount function.
	aftercount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)

	//validate GetDrop
	drops, dropFound := testInput.MarketKeeper.GetDrop(testInput.Context, beforecount)
	require.True(t, dropFound)
	require.Equal(t, drops.Pair, pair)
	//validate CreateDrop
	var d = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "70", Rate1: testdata.RateAstrArray, Prev1: "0", Next1: "0", Rate2: testdata.RateBstrArray, Prev2: "0", Next2: "0"}
	createDropResponse, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.NoError(t, err)

	//validate GetMember
	members, memberfound := testInput.MarketKeeper.GetMember(testInput.Context, denomB, denomA)
	members1, memberfound1 := testInput.MarketKeeper.GetMember(testInput.Context, denomA, denomB)
	require.True(t, memberfound)
	require.Equal(t, members.DenomA, denomB)
	require.Equal(t, members.DenomB, denomA)
	require.Equal(t, members.Balance.String(), "70")

	require.True(t, memberfound1)
	require.Equal(t, members1.DenomA, denomA)
	require.Equal(t, members1.DenomB, denomB)
	require.Equal(t, members1.Balance.String(), "70")

	//Validate GetPool
	rst, found := testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, found)
	require.Equal(t, rst.Pair, pair)
	require.Equal(t, rst.Drops.String(), "140")
	//validate GetDrop
	drops1, drop1Found := testInput.MarketKeeper.GetDrop(testInput.Context, aftercount)
	require.True(t, drop1Found)
	require.Equal(t, drops1.Pair, pair)
	require.Equal(t, drops1.Drops.String(), d.Drops)
	require.Contains(t, d.GetCreator(), createDropResponse.String())
	// GetUidCount before CreatePool
	beforecount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	//Create Order Prev 0 Next 0
	var o = types.MsgCreateOrder{Creator: addr, DenomAsk: members1.DenomA, DenomBid: members1.DenomB, Rate: testdata.RateAstrArray, OrderType: "limit", Amount: "0", Prev: "0", Next: "0"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &o)
	require.NoError(t, err)
	aftercount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)
	//Validate Order
	orders, orderfound := testInput.MarketKeeper.GetOrder(testInput.Context, beforecount)
	require.True(t, orderfound)
	require.Equal(t, orders.DenomBid, denomB)
	require.Equal(t, orders.DenomAsk, denomA)
	require.Equal(t, orders.Amount.String(), o.Amount)
	require.Equal(t, orders.Prev, uint64(0))
	require.Equal(t, orders.Next, uint64(0))

	orders.Active = true
	testInput.MarketKeeper.SetOrder(testInput.Context, orders)
	//Create Order Prev 0 Next 2
	var o1 = types.MsgCreateOrder{Creator: addr, DenomAsk: members1.DenomA, DenomBid: members1.DenomB, Rate: []string{"40", "50"}, OrderType: "limit", Amount: "0", Prev: "0", Next: "2"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &o1)
	require.NoError(t, err)
	aftercount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+2, aftercount)
	//Validate Order
	orders1, orderfound1 := testInput.MarketKeeper.GetOrder(testInput.Context, beforecount+1)
	require.True(t, orderfound1)
	require.Equal(t, orders1.DenomBid, denomB)
	require.Equal(t, orders1.DenomAsk, denomA)
	require.Equal(t, orders1.Amount.String(), o1.Amount)
	require.Equal(t, orders1.Prev, uint64(0))
	require.Equal(t, orders1.Next, uint64(2))

	//Create Order Prev 3 Next 0
	orders1.Next = uint64(0)
	orders1.Active = true
	testInput.MarketKeeper.SetOrder(testInput.Context, orders1)
	var o2 = types.MsgCreateOrder{Creator: addr, DenomAsk: members1.DenomA, DenomBid: members1.DenomB, Rate: []string{"40", "50"}, OrderType: "limit", Amount: "0", Prev: "3", Next: "0"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &o2)
	require.NoError(t, err)
	aftercount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+3, aftercount)
	//Validate Order
	orders2, orderfound2 := testInput.MarketKeeper.GetOrder(testInput.Context, beforecount+2)
	require.True(t, orderfound2)
	require.Equal(t, orders2.DenomBid, denomB)
	require.Equal(t, orders2.DenomAsk, denomA)
	require.Equal(t, orders2.Amount.String(), o2.Amount)
	require.Equal(t, orders2.Prev, uint64(3))
	require.Equal(t, orders2.Next, uint64(0))
	orders1.Prev = uint64(4)
	testdatadrop := testData{RateAstrArray: []string{"60", "70"}}
	numerator1, _ := sdk.NewIntFromString(testdatadrop.RateAstrArray[0])
	denominator1, _ := sdk.NewIntFromString(testdatadrop.RateAstrArray[1])
	orders1.Rate = []sdk.Int{numerator1, denominator1}
	testInput.MarketKeeper.SetOrder(testInput.Context, orders1)
	//Create Order Prev 4 Next 3
	orders2.Next = uint64(3)
	testInput.MarketKeeper.SetOrder(testInput.Context, orders2)
	var o3 = types.MsgCreateOrder{Creator: addr, DenomAsk: members1.DenomA, DenomBid: members1.DenomB, Rate: []string{"50", "60"}, OrderType: "limit", Amount: "0", Prev: "4", Next: "3"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &o3)
	require.NoError(t, err)
	aftercount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+4, aftercount)
	//Validate Order
	orders3, orderfound3 := testInput.MarketKeeper.GetOrder(testInput.Context, beforecount+3)
	require.True(t, orderfound3)
	require.Equal(t, orders3.DenomBid, denomB)
	require.Equal(t, orders3.DenomAsk, denomA)
	require.Equal(t, orders3.Amount.String(), o2.Amount)
	require.Equal(t, orders3.Prev, uint64(4))
	require.Equal(t, orders3.Next, uint64(3))
	require.Equal(t, orders3.OrderType, "limit")

	//Cancel Order
	Uid := strconv.FormatUint(orders3.Uid, 10)
	var co = types.MsgCancelOrder{Creator: addr, Uid: Uid}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CancelOrder(sdk.WrapSDKContext(testInput.Context), &co)
	require.NoError(t, err)

	members1, memberfound1 = testInput.MarketKeeper.GetMember(testInput.Context, orders.DenomAsk, orders.DenomBid)
	require.True(t, memberfound1)
	require.Equal(t, members1.DenomA, denomA)
	require.Equal(t, members1.DenomB, denomB)
	require.Equal(t, members1.Balance.String(), "70")

	//next order validation
	orders, orderfound = testInput.MarketKeeper.GetOrder(testInput.Context, orders1.Uid)
	require.True(t, orderfound)
	require.True(t, orders.Active)
	require.Equal(t, orders.DenomBid, denomB)
	require.Equal(t, orders.DenomAsk, denomA)
	require.Equal(t, orders.Amount.String(), o.Amount)
	require.Equal(t, orders.OrderType, "limit")
	require.Equal(t, orders.Prev, orders3.Prev)
	require.Equal(t, orders.Next, uint64(0))
	//prev order validation
	orders, orderfound = testInput.MarketKeeper.GetOrder(testInput.Context, orders2.Uid)
	require.True(t, orderfound)
	require.True(t, orders.Active)
	require.Equal(t, orders.DenomBid, denomB)
	require.Equal(t, orders.DenomAsk, denomA)
	require.Equal(t, orders.Amount.String(), o.Amount)
	require.Equal(t, orders.OrderType, "limit")
	require.Equal(t, orders.Prev, orders1.Uid)
	require.Equal(t, orders.Next, orders3.Next)

	//Validate Order
	orders, orderfound = testInput.MarketKeeper.GetOrder(testInput.Context, orders3.Uid)
	require.True(t, orderfound)
	require.False(t, orders.Active)
	require.Equal(t, orders.DenomBid, denomB)
	require.Equal(t, orders.DenomAsk, denomA)
	require.Equal(t, orders.Amount.String(), o.Amount)
	require.Equal(t, orders.OrderType, "limit")
	require.Equal(t, orders.Prev, orders3.Prev)
	require.Equal(t, orders.Next, orders3.Next)

}

// Order not found
func TestCancelOrder_case1_stop_negative(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "30CoinA", coinBStr: "40CoinB", RateAstrArray: []string{"60", "70"}, RateBstrArray: []string{"80", "90"}}
	coinPair, _ := sample.SampleCoins("70CoinA", "70CoinB")
	denomA, denomB := sample.SampleDenoms(coinPair)
	pair := strings.Join([]string{denomA, denomB}, ",")

	//MintCoins
	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, types.ModuleName, coinPair))
	//SendCoinsFromModuleToAccount
	requestAddress, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, types.ModuleName, requestAddress, coinPair))
	// GetUidCount before CreatePool
	beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	//Create Pool
	var p = types.MsgCreatePool{CoinA: testdata.coinAStr, CoinB: testdata.coinBStr, Creator: addr, RateA: testdata.RateAstrArray, RateB: testdata.RateBstrArray}
	response, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	//validate CreatePool
	require.NoError(t, err)
	require.Contains(t, p.GetCreator(), response.String())
	require.Contains(t, p.GetCoinA(), response.String())
	require.Contains(t, p.GetCoinB(), response.String())
	//validate SetUidCount function.
	aftercount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)

	//validate GetDrop
	drops, dropFound := testInput.MarketKeeper.GetDrop(testInput.Context, beforecount)
	require.True(t, dropFound)
	require.Equal(t, drops.Pair, pair)
	//validate CreateDrop
	var d = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "70", Rate1: testdata.RateAstrArray, Prev1: "0", Next1: "0", Rate2: testdata.RateBstrArray, Prev2: "0", Next2: "0"}
	createDropResponse, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.NoError(t, err)

	//validate GetMember
	members, memberfound := testInput.MarketKeeper.GetMember(testInput.Context, denomB, denomA)
	members1, memberfound1 := testInput.MarketKeeper.GetMember(testInput.Context, denomA, denomB)
	require.True(t, memberfound)
	require.Equal(t, members.DenomA, denomB)
	require.Equal(t, members.DenomB, denomA)
	require.Equal(t, members.Balance.String(), "70")

	require.True(t, memberfound1)
	require.Equal(t, members1.DenomA, denomA)
	require.Equal(t, members1.DenomB, denomB)
	require.Equal(t, members1.Balance.String(), "70")

	//Validate GetPool
	rst, found := testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, found)
	require.Equal(t, rst.Pair, pair)
	require.Equal(t, rst.Drops.String(), "140")
	//validate GetDrop
	drops1, drop1Found := testInput.MarketKeeper.GetDrop(testInput.Context, aftercount)
	require.True(t, drop1Found)
	require.Equal(t, drops1.Pair, pair)
	require.Equal(t, drops1.Drops.String(), d.Drops)
	require.Contains(t, d.GetCreator(), createDropResponse.String())
	// GetUidCount before CreatePool
	beforecount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	//Create Order
	var o = types.MsgCreateOrder{Creator: addr, DenomAsk: members1.DenomA, DenomBid: members1.DenomB, Rate: testdata.RateAstrArray, OrderType: "stop", Amount: "0", Prev: "0", Next: "0"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &o)
	require.NoError(t, err)
	aftercount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)
	//Validate Order
	orders, orderfound := testInput.MarketKeeper.GetOrder(testInput.Context, beforecount)
	require.True(t, orderfound)
	require.Equal(t, orders.DenomBid, denomB)
	require.Equal(t, orders.DenomAsk, denomA)
	require.Equal(t, orders.Amount.String(), o.Amount)

	//validate GetMember
	members, memberfound = testInput.MarketKeeper.GetMember(testInput.Context, orders.DenomBid, orders.DenomAsk)

	require.True(t, memberfound)
	require.Equal(t, members.DenomA, denomB)
	require.Equal(t, members.DenomB, denomA)
	require.Equal(t, members.Balance.String(), "70")

	require.Equal(t, members.Stop, uint64(0))

	//Create Order
	Uid := strconv.FormatUint(drops1.Uid, 10)
	var co = types.MsgCancelOrder{Creator: addr, Uid: Uid}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CancelOrder(sdk.WrapSDKContext(testInput.Context), &co)
	require.Error(t, err)
	require.ErrorContains(t, err, "order not found")

}

func TestCancelOrder_case1_limit_invalid_Order(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "30CoinA", coinBStr: "40CoinB", RateAstrArray: []string{"60", "70"}, RateBstrArray: []string{"80", "90"}}
	coinPair, _ := sample.SampleCoins("70CoinA", "70CoinB")
	denomA, denomB := sample.SampleDenoms(coinPair)
	pair := strings.Join([]string{denomA, denomB}, ",")

	//MintCoins
	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, types.ModuleName, coinPair))
	//SendCoinsFromModuleToAccount
	requestAddress, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, types.ModuleName, requestAddress, coinPair))
	// GetUidCount before CreatePool
	beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	//Create Pool
	var p = types.MsgCreatePool{CoinA: testdata.coinAStr, CoinB: testdata.coinBStr, Creator: addr, RateA: testdata.RateAstrArray, RateB: testdata.RateBstrArray}
	response, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	//validate CreatePool
	require.NoError(t, err)
	require.Contains(t, p.GetCreator(), response.String())
	require.Contains(t, p.GetCoinA(), response.String())
	require.Contains(t, p.GetCoinB(), response.String())
	//validate SetUidCount function.
	aftercount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)

	//validate GetDrop
	drops, dropFound := testInput.MarketKeeper.GetDrop(testInput.Context, beforecount)
	require.True(t, dropFound)
	require.Equal(t, drops.Pair, pair)
	//validate CreateDrop
	var d = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "70", Rate1: testdata.RateAstrArray, Prev1: "0", Next1: "0", Rate2: testdata.RateBstrArray, Prev2: "0", Next2: "0"}
	createDropResponse, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.NoError(t, err)

	//validate GetMember
	members, memberfound := testInput.MarketKeeper.GetMember(testInput.Context, denomB, denomA)
	members1, memberfound1 := testInput.MarketKeeper.GetMember(testInput.Context, denomA, denomB)
	require.True(t, memberfound)
	require.Equal(t, members.DenomA, denomB)
	require.Equal(t, members.DenomB, denomA)
	//require.Equal(t, members.Balance.String(), "70")

	require.True(t, memberfound1)
	require.Equal(t, members1.DenomA, denomA)
	require.Equal(t, members1.DenomB, denomB)
	//require.Equal(t, members1.Balance.String(), "70")

	//Validate GetPool
	rst, found := testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, found)
	require.Equal(t, rst.Pair, pair)
	//require.Equal(t, rst.Drops.String(), "140")
	//validate GetDrop
	drops1, drop1Found := testInput.MarketKeeper.GetDrop(testInput.Context, aftercount)
	require.True(t, drop1Found)
	require.Equal(t, drops1.Pair, pair)
	require.Equal(t, drops1.Drops.String(), d.Drops)
	require.Contains(t, d.GetCreator(), createDropResponse.String())
	// GetUidCount before CreatePool
	beforecount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	//Create Order
	var o = types.MsgCreateOrder{Creator: addr, DenomAsk: members1.DenomA, DenomBid: members1.DenomB, Rate: testdata.RateAstrArray, OrderType: "limit", Amount: "0", Prev: "0", Next: "0"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &o)
	require.NoError(t, err)
	aftercount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)
	//Validate Order
	orders, orderfound := testInput.MarketKeeper.GetOrder(testInput.Context, beforecount)
	require.True(t, orderfound)
	require.Equal(t, orders.DenomBid, denomB)
	require.Equal(t, orders.DenomAsk, denomA)
	require.Equal(t, orders.Amount.String(), o.Amount)

	//validate GetMember
	members, memberfound = testInput.MarketKeeper.GetMember(testInput.Context, orders.DenomBid, orders.DenomAsk)

	require.True(t, memberfound)
	require.Equal(t, members.DenomA, denomB)
	require.Equal(t, members.DenomB, denomA)
	//require.Equal(t, members.Balance.String(), "70")

	require.Equal(t, members.Stop, uint64(0))
	require.Equal(t, members.Limit, uint64(0))

	//Cancel Order
	Uid := strconv.FormatUint(orders.Uid, 10)
	var co = types.MsgCancelOrder{Creator: addr, Uid: Uid}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CancelOrder(sdk.WrapSDKContext(testInput.Context), &co)
	require.Error(t, err)
	require.ErrorContains(t, err, "invalid order")

}

func TestCancelOrder_case2_invalid_order(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "30CoinA", coinBStr: "40CoinB", RateAstrArray: []string{"60", "70"}, RateBstrArray: []string{"80", "90"}}
	coinPair, _ := sample.SampleCoins("70CoinA", "70CoinB")
	denomA, denomB := sample.SampleDenoms(coinPair)
	pair := strings.Join([]string{denomA, denomB}, ",")

	//MintCoins
	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, types.ModuleName, coinPair))
	//SendCoinsFromModuleToAccount
	requestAddress, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, types.ModuleName, requestAddress, coinPair))
	// GetUidCount before CreatePool
	beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	//Create Pool
	var p = types.MsgCreatePool{CoinA: testdata.coinAStr, CoinB: testdata.coinBStr, Creator: addr, RateA: testdata.RateAstrArray, RateB: testdata.RateBstrArray}
	response, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	//validate CreatePool
	require.NoError(t, err)
	require.Contains(t, p.GetCreator(), response.String())
	require.Contains(t, p.GetCoinA(), response.String())
	require.Contains(t, p.GetCoinB(), response.String())
	//validate SetUidCount function.
	aftercount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)

	//validate GetDrop
	drops, dropFound := testInput.MarketKeeper.GetDrop(testInput.Context, beforecount)
	require.True(t, dropFound)
	require.Equal(t, drops.Pair, pair)
	//validate CreateDrop
	var d = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "70", Rate1: testdata.RateAstrArray, Prev1: "0", Next1: "0", Rate2: testdata.RateBstrArray, Prev2: "0", Next2: "0"}
	createDropResponse, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.NoError(t, err)

	//validate GetMember
	members, memberfound := testInput.MarketKeeper.GetMember(testInput.Context, denomB, denomA)
	members1, memberfound1 := testInput.MarketKeeper.GetMember(testInput.Context, denomA, denomB)
	require.True(t, memberfound)
	require.Equal(t, members.DenomA, denomB)
	require.Equal(t, members.DenomB, denomA)
	//require.Equal(t, members.Balance.String(), "70")

	require.True(t, memberfound1)
	require.Equal(t, members1.DenomA, denomA)
	require.Equal(t, members1.DenomB, denomB)
	//require.Equal(t, members1.Balance.String(), "70")

	//Validate GetPool
	rst, found := testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, found)
	require.Equal(t, rst.Pair, pair)
	//require.Equal(t, rst.Drops.String(), "140")
	//validate GetDrop
	drops1, drop1Found := testInput.MarketKeeper.GetDrop(testInput.Context, aftercount)
	require.True(t, drop1Found)
	require.Equal(t, drops1.Pair, pair)
	require.Equal(t, drops1.Drops.String(), d.Drops)
	require.Contains(t, d.GetCreator(), createDropResponse.String())
	// GetUidCount before CreatePool
	beforecount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	//Create Order
	var o = types.MsgCreateOrder{Creator: addr, DenomAsk: members1.DenomA, DenomBid: members1.DenomB, Rate: testdata.RateAstrArray, OrderType: "limit", Amount: "0", Prev: "0", Next: "0"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &o)
	require.NoError(t, err)
	aftercount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)
	//Validate Order
	orders, orderfound := testInput.MarketKeeper.GetOrder(testInput.Context, beforecount)
	require.True(t, orderfound)
	require.Equal(t, orders.DenomBid, denomB)
	require.Equal(t, orders.DenomAsk, denomA)
	require.Equal(t, orders.Amount.String(), o.Amount)

	//validate GetMember
	members, memberfound = testInput.MarketKeeper.GetMember(testInput.Context, orders.DenomBid, orders.DenomAsk)

	require.True(t, memberfound)
	require.Equal(t, members.DenomA, denomB)
	require.Equal(t, members.DenomB, denomA)
	//require.Equal(t, members.Balance.String(), "70")

	require.Equal(t, members.Stop, uint64(0))
	require.Equal(t, members.Limit, uint64(0))

	//SetOrder to Next value as 1
	orders.Next = uint64(1)
	orders.Active = true
	testInput.MarketKeeper.SetOrder(testInput.Context, orders)

	//Verify Set Operation is succesfull

	orders, orderfound = testInput.MarketKeeper.GetOrder(testInput.Context, beforecount)
	require.True(t, orderfound)
	require.Equal(t, orders.Next, uint64(1))

	o.Next = "2"
	o.Rate = []string{"50", "60"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &o)
	require.NoError(t, err)

	orders, orderfound = testInput.MarketKeeper.GetOrder(testInput.Context, aftercount)
	require.True(t, orderfound)
	require.Equal(t, orders.Next, uint64(2))

	//Cancel Order
	Uid := strconv.FormatUint(orders.Uid, 10)
	var co = types.MsgCancelOrder{Creator: addr, Uid: Uid}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CancelOrder(sdk.WrapSDKContext(testInput.Context), &co)
	require.Error(t, err)
	require.ErrorContains(t, err, "invalid order")

}

func TestCancelOrder_case3_order_not_found(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "30CoinA", coinBStr: "40CoinB", RateAstrArray: []string{"60", "70"}, RateBstrArray: []string{"80", "90"}}
	coinPair, _ := sample.SampleCoins("70CoinA", "70CoinB")
	denomA, denomB := sample.SampleDenoms(coinPair)
	pair := strings.Join([]string{denomA, denomB}, ",")

	//MintCoins
	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, types.ModuleName, coinPair))
	//SendCoinsFromModuleToAccount
	requestAddress, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, types.ModuleName, requestAddress, coinPair))
	// GetUidCount before CreatePool
	beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	//Create Pool
	var p = types.MsgCreatePool{CoinA: testdata.coinAStr, CoinB: testdata.coinBStr, Creator: addr, RateA: testdata.RateAstrArray, RateB: testdata.RateBstrArray}
	response, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	//validate CreatePool
	require.NoError(t, err)
	require.Contains(t, p.GetCreator(), response.String())
	require.Contains(t, p.GetCoinA(), response.String())
	require.Contains(t, p.GetCoinB(), response.String())
	//validate SetUidCount function.
	aftercount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)

	//validate GetDrop
	drops, dropFound := testInput.MarketKeeper.GetDrop(testInput.Context, beforecount)
	require.True(t, dropFound)
	require.Equal(t, drops.Pair, pair)
	//validate CreateDrop
	var d = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "70", Rate1: testdata.RateAstrArray, Prev1: "0", Next1: "0", Rate2: testdata.RateBstrArray, Prev2: "0", Next2: "0"}
	createDropResponse, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.NoError(t, err)

	//validate GetMember
	members, memberfound := testInput.MarketKeeper.GetMember(testInput.Context, denomB, denomA)
	members1, memberfound1 := testInput.MarketKeeper.GetMember(testInput.Context, denomA, denomB)
	require.True(t, memberfound)
	require.Equal(t, members.DenomA, denomB)
	require.Equal(t, members.DenomB, denomA)
	require.Equal(t, members.Balance.String(), "70")

	require.True(t, memberfound1)
	require.Equal(t, members1.DenomA, denomA)
	require.Equal(t, members1.DenomB, denomB)
	require.Equal(t, members1.Balance.String(), "70")

	//Validate GetPool
	rst, found := testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, found)
	require.Equal(t, rst.Pair, pair)
	require.Equal(t, rst.Drops.String(), "140")
	//validate GetDrop
	drops1, drop1Found := testInput.MarketKeeper.GetDrop(testInput.Context, aftercount)
	require.True(t, drop1Found)
	require.Equal(t, drops1.Pair, pair)
	require.Equal(t, drops1.Drops.String(), d.Drops)
	require.Contains(t, d.GetCreator(), createDropResponse.String())
	// GetUidCount before CreatePool
	beforecount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	//Create Order
	var o = types.MsgCreateOrder{Creator: addr, DenomAsk: members1.DenomA, DenomBid: members1.DenomB, Rate: testdata.RateAstrArray, OrderType: "stop", Amount: "0", Prev: "0", Next: "0"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &o)
	require.NoError(t, err)
	aftercount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)
	//Validate Order
	orders, orderfound := testInput.MarketKeeper.GetOrder(testInput.Context, beforecount)
	require.True(t, orderfound)
	require.Equal(t, orders.DenomBid, denomB)
	require.Equal(t, orders.DenomAsk, denomA)
	require.Equal(t, orders.Amount.String(), o.Amount)

	//validate GetMember
	members, memberfound = testInput.MarketKeeper.GetMember(testInput.Context, orders.DenomBid, orders.DenomAsk)

	require.True(t, memberfound)
	require.Equal(t, members.DenomA, denomB)
	require.Equal(t, members.DenomB, denomA)
	require.Equal(t, members.Balance.String(), "70")

	require.Equal(t, members.Stop, uint64(0))

	//SetOrder to Next value as 1
	orders.Prev = uint64(1)
	testInput.MarketKeeper.SetOrder(testInput.Context, orders)

	//Verify Set Operation is succesfull

	orders, orderfound = testInput.MarketKeeper.GetOrder(testInput.Context, beforecount)
	require.True(t, orderfound)
	require.Equal(t, orders.Prev, uint64(1))

	o.Prev = "2"
	o.Rate = []string{"50", "60"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &o)
	require.NoError(t, err)

	orders, orderfound = testInput.MarketKeeper.GetOrder(testInput.Context, aftercount)
	require.True(t, orderfound)
	require.Equal(t, orders.Next, uint64(0))

	//SetOrder to Prev value as 6 to get order not found
	orders.Prev = uint64(6)
	testInput.MarketKeeper.SetOrder(testInput.Context, orders)

	//Cancel Order
	Uid := strconv.FormatUint(orders.Uid, 10)
	var co = types.MsgCancelOrder{Creator: addr, Uid: Uid}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CancelOrder(sdk.WrapSDKContext(testInput.Context), &co)
	require.Error(t, err)
	require.ErrorContains(t, err, "order not found")

}

func TestCancelOrder_case4_order_not_found(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "30CoinA", coinBStr: "40CoinB", RateAstrArray: []string{"70", "80"}, RateBstrArray: []string{"90", "100"}}
	coinPair, _ := sample.SampleCoins("70CoinA", "70CoinB")
	denomA, denomB := sample.SampleDenoms(coinPair)
	pair := strings.Join([]string{denomA, denomB}, ",")

	//MintCoins
	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, types.ModuleName, coinPair))
	//SendCoinsFromModuleToAccount
	requestAddress, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, types.ModuleName, requestAddress, coinPair))
	// GetUidCount before CreatePool
	beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	//Create Pool
	var p = types.MsgCreatePool{CoinA: testdata.coinAStr, CoinB: testdata.coinBStr, Creator: addr, RateA: testdata.RateAstrArray, RateB: testdata.RateBstrArray}
	response, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	//validate CreatePool
	require.NoError(t, err)
	require.Contains(t, p.GetCreator(), response.String())
	require.Contains(t, p.GetCoinA(), response.String())
	require.Contains(t, p.GetCoinB(), response.String())
	//validate SetUidCount function.
	aftercount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)

	//validate GetDrop
	drops, dropFound := testInput.MarketKeeper.GetDrop(testInput.Context, beforecount)
	require.True(t, dropFound)
	require.Equal(t, drops.Pair, pair)
	//validate CreateDrop
	var d = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "70", Rate1: testdata.RateAstrArray, Prev1: "0", Next1: "0", Rate2: testdata.RateBstrArray, Prev2: "0", Next2: "0"}
	createDropResponse, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.NoError(t, err)

	//validate GetMember
	members, memberfound := testInput.MarketKeeper.GetMember(testInput.Context, denomB, denomA)
	members1, memberfound1 := testInput.MarketKeeper.GetMember(testInput.Context, denomA, denomB)
	require.True(t, memberfound)
	require.Equal(t, members.DenomA, denomB)
	require.Equal(t, members.DenomB, denomA)
	require.Equal(t, members.Balance.String(), "70")

	require.True(t, memberfound1)
	require.Equal(t, members1.DenomA, denomA)
	require.Equal(t, members1.DenomB, denomB)
	require.Equal(t, members1.Balance.String(), "70")

	//Validate GetPool
	rst, found := testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, found)
	require.Equal(t, rst.Pair, pair)
	require.Equal(t, rst.Drops.String(), "140")
	//validate GetDrop
	drops1, drop1Found := testInput.MarketKeeper.GetDrop(testInput.Context, aftercount)
	require.True(t, drop1Found)
	require.Equal(t, drops1.Pair, pair)
	require.Equal(t, drops1.Drops.String(), d.Drops)
	require.Contains(t, d.GetCreator(), createDropResponse.String())
	// GetUidCount before CreatePool
	beforecount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	//Create Order Prev 0 Next 0
	var o = types.MsgCreateOrder{Creator: addr, DenomAsk: members1.DenomA, DenomBid: members1.DenomB, Rate: testdata.RateAstrArray, OrderType: "limit", Amount: "0", Prev: "0", Next: "0"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &o)
	require.NoError(t, err)
	aftercount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)
	//Validate Order
	orders, orderfound := testInput.MarketKeeper.GetOrder(testInput.Context, beforecount)
	require.True(t, orderfound)
	require.Equal(t, orders.DenomBid, denomB)
	require.Equal(t, orders.DenomAsk, denomA)
	require.Equal(t, orders.Amount.String(), o.Amount)
	require.Equal(t, orders.Prev, uint64(0))
	require.Equal(t, orders.Next, uint64(0))

	orders.Active = true
	testInput.MarketKeeper.SetOrder(testInput.Context, orders)
	//Create Order Prev 0 Next 2
	var o1 = types.MsgCreateOrder{Creator: addr, DenomAsk: members1.DenomA, DenomBid: members1.DenomB, Rate: []string{"40", "50"}, OrderType: "limit", Amount: "0", Prev: "0", Next: "2"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &o1)
	require.NoError(t, err)
	aftercount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+2, aftercount)
	//Validate Order
	orders1, orderfound1 := testInput.MarketKeeper.GetOrder(testInput.Context, beforecount+1)
	require.True(t, orderfound1)
	require.Equal(t, orders1.DenomBid, denomB)
	require.Equal(t, orders1.DenomAsk, denomA)
	require.Equal(t, orders1.Amount.String(), o1.Amount)
	require.Equal(t, orders1.Prev, uint64(0))
	require.Equal(t, orders1.Next, uint64(2))

	//Create Order Prev 3 Next 0
	orders1.Next = uint64(0)
	orders1.Active = true
	testInput.MarketKeeper.SetOrder(testInput.Context, orders1)
	var o2 = types.MsgCreateOrder{Creator: addr, DenomAsk: members1.DenomA, DenomBid: members1.DenomB, Rate: []string{"40", "50"}, OrderType: "limit", Amount: "0", Prev: "3", Next: "0"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &o2)
	require.NoError(t, err)
	aftercount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+3, aftercount)
	//Validate Order
	orders2, orderfound2 := testInput.MarketKeeper.GetOrder(testInput.Context, beforecount+2)
	require.True(t, orderfound2)
	require.Equal(t, orders2.DenomBid, denomB)
	require.Equal(t, orders2.DenomAsk, denomA)
	require.Equal(t, orders2.Amount.String(), o2.Amount)
	require.Equal(t, orders2.Prev, uint64(3))
	require.Equal(t, orders2.Next, uint64(0))
	orders1.Prev = uint64(4)
	testdatadrop := testData{RateAstrArray: []string{"60", "70"}}
	numerator1, _ := sdk.NewIntFromString(testdatadrop.RateAstrArray[0])
	denominator1, _ := sdk.NewIntFromString(testdatadrop.RateAstrArray[1])
	orders1.Rate = []sdk.Int{numerator1, denominator1}
	testInput.MarketKeeper.SetOrder(testInput.Context, orders1)
	//Create Order Prev 4 Next 3
	orders2.Next = uint64(3)
	testInput.MarketKeeper.SetOrder(testInput.Context, orders2)
	var o3 = types.MsgCreateOrder{Creator: addr, DenomAsk: members1.DenomA, DenomBid: members1.DenomB, Rate: []string{"50", "60"}, OrderType: "limit", Amount: "0", Prev: "4", Next: "3"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateOrder(sdk.WrapSDKContext(testInput.Context), &o3)
	require.NoError(t, err)
	aftercount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+4, aftercount)
	//Validate Order
	orders3, orderfound3 := testInput.MarketKeeper.GetOrder(testInput.Context, beforecount+3)
	require.True(t, orderfound3)
	require.Equal(t, orders3.DenomBid, denomB)
	require.Equal(t, orders3.DenomAsk, denomA)
	require.Equal(t, orders3.Amount.String(), o2.Amount)
	require.Equal(t, orders3.Prev, uint64(4))
	require.Equal(t, orders3.Next, uint64(3))
	require.Equal(t, orders3.OrderType, "limit")
	//Assign Next value as 7 to get order not found
	orders3.Next = uint64(7)
	testInput.MarketKeeper.SetOrder(testInput.Context, orders3)
	//Cancel Order
	Uid := strconv.FormatUint(orders3.Uid, 10)
	var co = types.MsgCancelOrder{Creator: addr, Uid: Uid}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CancelOrder(sdk.WrapSDKContext(testInput.Context), &co)
	require.Error(t, err)
	require.ErrorContains(t, err, "order not found")

}
*/
