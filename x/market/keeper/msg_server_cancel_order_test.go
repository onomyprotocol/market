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
	require.True(t, orders.Status == "canceled")
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
	require.True(t, orders.Status == "canceled")
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
	require.True(t, orders.Status == "canceled")
	require.Equal(t, orders.DenomBid, denomB)
	require.Equal(t, orders.DenomAsk, denomA)
	require.Equal(t, orders.Amount.String(), o.Amount)
	require.Equal(t, orders.OrderType, "stop")
}

func TestCancelOrderEmptyPool(t *testing.T) {
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

	// Validate RedeemDrop
	Uid := strconv.FormatUint(1, 10)
	var rd = types.MsgRedeemDrop{Creator: addr, Uid: Uid}
	createRedeemDropResponse, redeemdropErr := keeper.NewMsgServerImpl(*testInput.MarketKeeper).RedeemDrop(sdk.WrapSDKContext(testInput.Context), &rd)
	require.NoError(t, redeemdropErr)
	require.Contains(t, rd.GetCreator(), createRedeemDropResponse.String())

	// Validate RedeemDrop
	Uid = strconv.FormatUint(2, 10)
	rd = types.MsgRedeemDrop{Creator: addr, Uid: Uid}
	createRedeemDropResponse, redeemdropErr = keeper.NewMsgServerImpl(*testInput.MarketKeeper).RedeemDrop(sdk.WrapSDKContext(testInput.Context), &rd)
	require.NoError(t, redeemdropErr)
	require.Contains(t, rd.GetCreator(), createRedeemDropResponse.String())

	// Create Pool
	var p = types.MsgCreatePool{CoinA: testdata.coinAStr, CoinB: testdata.coinBStr, Creator: addr}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	require.NoError(t, err)

	// Cancel Order
	Uid = strconv.FormatUint(beforecount, 10)
	var co = types.MsgCancelOrder{Creator: addr, Uid: Uid}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CancelOrder(sdk.WrapSDKContext(testInput.Context), &co)
	require.NoError(t, err)

}
