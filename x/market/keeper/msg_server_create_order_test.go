package keeper_test

import (
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/pendulum-labs/market/testutil/keeper"
	"github.com/pendulum-labs/market/testutil/sample"
	"github.com/pendulum-labs/market/x/market/keeper"
	"github.com/pendulum-labs/market/x/market/types"
	"github.com/stretchr/testify/require"
)

func TestCreateOrder_case1_stop(t *testing.T) {
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
	require.Equal(t, members.Protect, uint64(1))
	require.True(t, memberfound1)
	require.Equal(t, members1.DenomA, denomA)
	require.Equal(t, members1.DenomB, denomB)
	require.Equal(t, members1.Balance.String(), "70")
	require.Equal(t, members1.Protect, uint64(1))
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
	require.Equal(t, members.Protect, uint64(1))
	require.Equal(t, members.Stop, uint64(0))

}

func TestCreateOrder_case1_limit(t *testing.T) {
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
	var d = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "50", Rate1: testdata.RateAstrArray, Prev1: "0", Next1: "0", Rate2: testdata.RateBstrArray, Prev2: "0", Next2: "0"}
	createDropResponse, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.NoError(t, err)

	//validate GetMember
	members, memberfound := testInput.MarketKeeper.GetMember(testInput.Context, denomB, denomA)
	members1, memberfound1 := testInput.MarketKeeper.GetMember(testInput.Context, denomA, denomB)
	require.True(t, memberfound)
	require.Equal(t, members.DenomA, denomB)
	require.Equal(t, members.DenomB, denomA)
	//require.Equal(t, members.Balance.String(), "70")
	require.Equal(t, members.Protect, uint64(1))
	require.True(t, memberfound1)
	require.Equal(t, members1.DenomA, denomA)
	require.Equal(t, members1.DenomB, denomB)
	//require.Equal(t, members1.Balance.String(), "70")
	require.Equal(t, members1.Protect, uint64(1))
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
	require.Equal(t, members.Protect, uint64(1))
	require.Equal(t, members.Stop, uint64(0))

}
