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

type testDataDrop struct {
	RateAstrArray []string
	RateBstrArray []string
}

func TestCreateDrop_case1(t *testing.T) {
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

}

func TestCreateDrop_case2_side1(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "30CoinA", coinBStr: "40CoinB", RateAstrArray: []string{"60", "70"}, RateBstrArray: []string{"80", "90"}}
	coinPair, _ := sample.SampleCoins("100CoinA", "100CoinB")
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
	//Update Next1 value to 1
	drops1.Next1 = uint64(1)
	testdatadrop := testDataDrop{RateAstrArray: []string{"10", "20"}}
	numerator1, _ := sdk.NewIntFromString(testdatadrop.RateAstrArray[0])
	denominator1, _ := sdk.NewIntFromString(testdatadrop.RateAstrArray[1])
	drops1.Rate1 = []sdk.Int{numerator1, denominator1}

	testInput.MarketKeeper.SetDrop(testInput.Context, drops1)

	var d1 = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "40", Rate1: testdata.RateAstrArray, Prev1: "0", Next1: "1", Rate2: testdata.RateBstrArray, Prev2: "0", Next2: "0"}
	createDropResponse, err1 := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d1)
	require.NoError(t, err1)
	require.Contains(t, d.GetCreator(), createDropResponse.String())

}

func TestCreateDrop_case2_side2(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "30CoinA", coinBStr: "40CoinB", RateAstrArray: []string{"60", "70"}, RateBstrArray: []string{"80", "90"}}
	coinPair, _ := sample.SampleCoins("100CoinA", "100CoinB")
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
	//Update Next2 value to 1
	drops1.Next2 = uint64(1)
	testdatadrop := testDataDrop{RateBstrArray: []string{"10", "20"}}
	numerator2, _ := sdk.NewIntFromString(testdatadrop.RateBstrArray[0])
	denominator2, _ := sdk.NewIntFromString(testdatadrop.RateBstrArray[1])
	drops1.Rate2 = []sdk.Int{numerator2, denominator2}

	testInput.MarketKeeper.SetDrop(testInput.Context, drops1)

	var d1 = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "40", Rate1: testdata.RateAstrArray, Prev1: "0", Next1: "0", Rate2: testdata.RateBstrArray, Prev2: "0", Next2: "1"}
	createDropResponse, err1 := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d1)
	require.NoError(t, err1)
	require.Contains(t, d.GetCreator(), createDropResponse.String())

}

func TestCreateDrop_case3_side1(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "30CoinA", coinBStr: "40CoinB", RateAstrArray: []string{"60", "70"}, RateBstrArray: []string{"80", "90"}}
	coinPair, _ := sample.SampleCoins("100CoinA", "100CoinB")
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
	//Update Prev1 value to 1
	drops1.Prev1 = uint64(1)
	testdatadrop := testDataDrop{RateAstrArray: []string{"70", "80"}}
	numerator1, _ := sdk.NewIntFromString(testdatadrop.RateAstrArray[0])
	denominator1, _ := sdk.NewIntFromString(testdatadrop.RateAstrArray[1])
	drops1.Rate1 = []sdk.Int{numerator1, denominator1}

	testInput.MarketKeeper.SetDrop(testInput.Context, drops1)

	var d1 = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "40", Rate1: testdata.RateAstrArray, Prev1: "1", Next1: "0", Rate2: testdata.RateBstrArray, Prev2: "0", Next2: "0"}
	createDropResponse, err1 := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d1)
	require.NoError(t, err1)
	require.Contains(t, d.GetCreator(), createDropResponse.String())

}

func TestCreateDrop_case3_side2(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "30CoinA", coinBStr: "40CoinB", RateAstrArray: []string{"60", "70"}, RateBstrArray: []string{"80", "90"}}
	coinPair, _ := sample.SampleCoins("100CoinA", "100CoinB")
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
	//Update Next2 value to 1
	drops1.Prev2 = uint64(1)
	testdatadrop := testDataDrop{RateBstrArray: []string{"80", "90"}}
	numerator2, _ := sdk.NewIntFromString(testdatadrop.RateBstrArray[0])
	denominator2, _ := sdk.NewIntFromString(testdatadrop.RateBstrArray[1])
	drops1.Rate2 = []sdk.Int{numerator2, denominator2}

	testInput.MarketKeeper.SetDrop(testInput.Context, drops1)

	var d1 = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "40", Rate1: testdata.RateAstrArray, Prev1: "0", Next1: "0", Rate2: testdata.RateBstrArray, Prev2: "1", Next2: "0"}
	createDropResponse, err1 := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d1)
	require.NoError(t, err1)
	require.Contains(t, d.GetCreator(), createDropResponse.String())

}

func TestCreateDrop_case4_side1(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "30CoinA", coinBStr: "40CoinB", RateAstrArray: []string{"60", "70"}, RateBstrArray: []string{"80", "90"}}
	coinPair, _ := sample.SampleCoins("100CoinA", "100CoinB")
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
	//var drops2 = drops1
	//drops2.Next1 = aftercount
	require.True(t, drop1Found)
	require.Equal(t, drops1.Pair, pair)
	require.Equal(t, drops1.Drops.String(), d.Drops)
	require.Contains(t, d.GetCreator(), createDropResponse.String())
	//Update Prev1 value to 1
	drops1.Prev1 = aftercount
	drops1.Next1 = aftercount

	testInput.MarketKeeper.SetDrop(testInput.Context, drops1)
	//testInput.MarketKeeper.SetDrop(testInput.Context, drops2)
	//GetUidCount
	beforedropcount := testInput.MarketKeeper.GetUidCount(testInput.Context)

	var d1 = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "40", Rate1: testdata.RateAstrArray, Prev1: "1", Next1: "1", Rate2: testdata.RateBstrArray, Prev2: "0", Next2: "0"}
	createDropResponse, err1 := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d1)
	require.NoError(t, err1)
	require.Contains(t, d.GetCreator(), createDropResponse.String())

	//Validate GetDrop
	drops2, drop2Found := testInput.MarketKeeper.GetDrop(testInput.Context, beforedropcount)
	require.True(t, drop2Found)
	require.Equal(t, drops2.Pair, pair)
	require.Equal(t, drops2.Drops.String(), d1.Drops)
	require.Equal(t, drops2.Prev1, aftercount)
	require.Equal(t, drops2.Next1, aftercount)
	require.Equal(t, drops2.Next1, drops2.Prev1)
	require.Contains(t, d1.GetCreator(), createDropResponse.String())

}

func TestCreateDrop_case4_side2(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "30CoinA", coinBStr: "40CoinB", RateAstrArray: []string{"60", "70"}, RateBstrArray: []string{"80", "90"}}
	coinPair, _ := sample.SampleCoins("100CoinA", "100CoinB")
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
	//var drops2 = drops1
	//drops2.Next1 = aftercount
	require.True(t, drop1Found)
	require.Equal(t, drops1.Pair, pair)
	require.Equal(t, drops1.Drops.String(), d.Drops)
	require.Contains(t, d.GetCreator(), createDropResponse.String())
	//Update Prev2 value to 1
	drops1.Prev2 = aftercount
	drops1.Next2 = aftercount

	testInput.MarketKeeper.SetDrop(testInput.Context, drops1)
	//GetUidCount
	beforedropcount := testInput.MarketKeeper.GetUidCount(testInput.Context)

	var d1 = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "40", Rate1: testdata.RateAstrArray, Prev1: "0", Next1: "0", Rate2: testdata.RateBstrArray, Prev2: "1", Next2: "1"}
	createDropResponse, err1 := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d1)
	require.NoError(t, err1)
	require.Contains(t, d.GetCreator(), createDropResponse.String())
	//Validate GetDrop
	drops2, drop2Found := testInput.MarketKeeper.GetDrop(testInput.Context, beforedropcount)
	require.True(t, drop2Found)
	require.Equal(t, drops2.Pair, pair)
	require.Equal(t, drops2.Drops.String(), d1.Drops)
	require.Equal(t, drops2.Prev2, aftercount)
	require.Equal(t, drops2.Next2, aftercount)
	require.Equal(t, drops2.Next2, drops2.Prev2)
	require.Contains(t, d1.GetCreator(), createDropResponse.String())

}

func TestCreateDrop_Pool_Not_Found(t *testing.T) {
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
	scenarios := []struct {
		coinAStr      string
		coinBStr      string
		RateAstrArray []string
		RateBstrArray []string
		Creator       string
	}{
		{coinAStr: "20CoinC", coinBStr: "20CoinD", RateAstrArray: []string{"10", "20"}, RateBstrArray: []string{"20", "30"}, Creator: addr},
		{coinAStr: "20CoinD", coinBStr: "20CoinA", RateAstrArray: []string{"10", "20"}, RateBstrArray: []string{"20", "30"}, Creator: sample.AccAddress()},
	}
	for _, s := range scenarios {
		coinPair, _ = sample.SampleCoins(s.coinAStr, s.coinBStr)
		denomA, denomB = sample.SampleDenoms(coinPair)
		pair = strings.Join([]string{denomA, denomB}, ",")
		var d = types.MsgCreateDrop{Creator: s.Creator, Pair: pair, Drops: "70", Rate1: testdata.RateAstrArray, Prev1: "0", Next1: "0", Rate2: testdata.RateBstrArray, Prev2: "0", Next2: "0"}
		_, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
		require.Error(t, err)
		require.ErrorContains(t, err, "the pool not found")

	}
}

func TestCreateDrop_Case1_Negative(t *testing.T) {
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
	require.Contains(t, d.GetCreator(), createDropResponse.String())

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
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.Error(t, err)
	require.ErrorContains(t, err, "Member 1 protect field not 0")
	//Set Members1 Protect to 0
	members.Protect = uint64(0)
	testInput.MarketKeeper.SetMember(testInput.Context, members)
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.Error(t, err)
	require.ErrorContains(t, err, "Member 2 protect field not 0")

}

func TestCreateDrop_Case2_Side1_Negative(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "30CoinA", coinBStr: "40CoinB", RateAstrArray: []string{"60", "70"}, RateBstrArray: []string{"80", "90"}}
	coinPair, _ := sample.SampleCoins("100CoinA", "100CoinB")
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
	//validate Next1 drop not active
	var d = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "70", Rate1: testdata.RateAstrArray, Prev1: "0", Next1: "1", Rate2: testdata.RateBstrArray, Prev2: "0", Next2: "0"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.Error(t, err)
	require.ErrorContains(t, err, "Next1 drop not active")
	//Happy Path
	d.Next1 = "0"
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.NoError(t, err)

	//validate GetDrop
	drops1, drop1Found := testInput.MarketKeeper.GetDrop(testInput.Context, aftercount)
	require.True(t, drop1Found)

	//Validate Next1 drop not currently head of book
	drops1.Prev1 = uint64(1)
	testInput.MarketKeeper.SetDrop(testInput.Context, drops1)
	d.Next1 = "1"
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.Error(t, err)
	require.ErrorContains(t, err, "Next1 drop not currently head of book")

	//validate Drop Rate1 less than or equal Next1
	drops1.Prev1 = uint64(0)
	testdatadrop := testDataDrop{RateAstrArray: []string{"70", "80"}}
	numerator1, _ := sdk.NewIntFromString(testdatadrop.RateAstrArray[0])
	denominator1, _ := sdk.NewIntFromString(testdatadrop.RateAstrArray[1])
	drops1.Rate1 = []sdk.Int{numerator1, denominator1}
	//Set Drop
	testInput.MarketKeeper.SetDrop(testInput.Context, drops1)

	var d1 = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "40", Rate1: testdata.RateAstrArray, Prev1: "0", Next1: "1", Rate2: testdata.RateBstrArray, Prev2: "0", Next2: "0"}
	_, err1 := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d1)
	require.Error(t, err1)
	require.ErrorContains(t, err1, "Drop Rate1 less than or equal Next1")

}

func TestCreateDrop_Case2_Side2_Negative(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "30CoinA", coinBStr: "40CoinB", RateAstrArray: []string{"60", "70"}, RateBstrArray: []string{"80", "90"}}
	coinPair, _ := sample.SampleCoins("100CoinA", "100CoinB")
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
	//validate Next1 drop not active
	var d = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "70", Rate1: testdata.RateAstrArray, Prev1: "0", Next1: "0", Rate2: testdata.RateBstrArray, Prev2: "0", Next2: "1"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.Error(t, err)
	require.ErrorContains(t, err, "Next2 drop not active")
	//validate CreateDrop
	d.Next2 = "0"
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.NoError(t, err)

	//validate GetDrop
	drops1, drop1Found := testInput.MarketKeeper.GetDrop(testInput.Context, aftercount)
	require.True(t, drop1Found)

	//Validate Next2 drop not currently head of book
	drops1.Prev2 = uint64(1)
	testInput.MarketKeeper.SetDrop(testInput.Context, drops1)
	d.Next2 = "1"
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.Error(t, err)
	require.ErrorContains(t, err, "Next2 drop not currently head of book")

	//validate Drop Rate2 less than or equal Next1
	drops1.Prev2 = uint64(0)
	testdatadrop := testDataDrop{RateBstrArray: []string{"100", "110"}}
	numerator2, _ := sdk.NewIntFromString(testdatadrop.RateBstrArray[0])
	denominator2, _ := sdk.NewIntFromString(testdatadrop.RateBstrArray[1])
	drops1.Rate2 = []sdk.Int{numerator2, denominator2}
	//Set Drop
	testInput.MarketKeeper.SetDrop(testInput.Context, drops1)

	var d1 = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "40", Rate1: testdata.RateAstrArray, Prev1: "0", Next1: "0", Rate2: testdata.RateBstrArray, Prev2: "0", Next2: "1"}
	_, err1 := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d1)
	require.Error(t, err1)
	require.ErrorContains(t, err1, "Drop Rate2 less than or equal Next2")

}

func TestCreateDrop_Case3_Side1_Negative(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "30CoinA", coinBStr: "40CoinB", RateAstrArray: []string{"60", "70"}, RateBstrArray: []string{"80", "90"}}
	coinPair, _ := sample.SampleCoins("100CoinA", "100CoinB")
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
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	//validate CreatePool
	require.NoError(t, err)

	//validate SetUidCount function.
	aftercount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)

	//validate GetDrop
	drops, dropFound := testInput.MarketKeeper.GetDrop(testInput.Context, beforecount)
	require.True(t, dropFound)
	require.Equal(t, drops.Pair, pair)
	//validate CreateDrop
	var d = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "70", Rate1: testdata.RateAstrArray, Prev1: "1", Next1: "0", Rate2: testdata.RateBstrArray, Prev2: "0", Next2: "0"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.Error(t, err)
	require.ErrorContains(t, err, "Prev1 drop not active")

	//Happy Path
	d.Prev1 = "0"
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.NoError(t, err)

	//validate GetDrop
	drops1, drop1Found := testInput.MarketKeeper.GetDrop(testInput.Context, aftercount)
	require.True(t, drop1Found)
	//Validate Prev1 drop not currently head of book
	drops1.Next1 = uint64(1)
	testInput.MarketKeeper.SetDrop(testInput.Context, drops1)
	d.Prev1 = "1"
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.Error(t, err)
	require.ErrorContains(t, err, "Prev1 drop not currently tail of book")
	//Validate Order rate greater than Prev
	drops1.Next1 = uint64(0)
	testdatadrop := testDataDrop{RateAstrArray: []string{"20", "30"}}
	numerator1, _ := sdk.NewIntFromString(testdatadrop.RateAstrArray[0])
	denominator1, _ := sdk.NewIntFromString(testdatadrop.RateAstrArray[1])
	drops1.Rate1 = []sdk.Int{numerator1, denominator1}

	testInput.MarketKeeper.SetDrop(testInput.Context, drops1)

	var d1 = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "40", Rate1: testdata.RateAstrArray, Prev1: "1", Next1: "0", Rate2: testdata.RateBstrArray, Prev2: "0", Next2: "0"}
	_, err1 := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d1)
	require.Error(t, err1)
	require.ErrorContains(t, err1, "Order rate greater than Prev")

}

func TestCreateDrop_Case3_Side2_Negative(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "30CoinA", coinBStr: "40CoinB", RateAstrArray: []string{"60", "70"}, RateBstrArray: []string{"80", "90"}}
	coinPair, _ := sample.SampleCoins("100CoinA", "100CoinB")
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
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	//validate CreatePool
	require.NoError(t, err)

	//validate SetUidCount function.
	aftercount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)

	//validate GetDrop
	drops, dropFound := testInput.MarketKeeper.GetDrop(testInput.Context, beforecount)
	require.True(t, dropFound)
	require.Equal(t, drops.Pair, pair)
	//validate CreateDrop
	var d = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "70", Rate1: testdata.RateAstrArray, Prev1: "0", Next1: "0", Rate2: testdata.RateBstrArray, Prev2: "1", Next2: "0"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.Error(t, err)
	require.ErrorContains(t, err, "Prev2 drop not active")

	//Happy Path
	d.Prev2 = "0"
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.NoError(t, err)

	//validate GetDrop
	drops1, drop1Found := testInput.MarketKeeper.GetDrop(testInput.Context, aftercount)
	require.True(t, drop1Found)
	//Validate Prev1 drop not currently head of book
	drops1.Next2 = uint64(1)
	testInput.MarketKeeper.SetDrop(testInput.Context, drops1)
	d.Prev2 = "1"
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.Error(t, err)
	require.ErrorContains(t, err, "Prev2 drop not currently tail of book")
	//Validate Order rate greater than Prev
	drops1.Next2 = uint64(0)
	testdatadrop := testDataDrop{RateBstrArray: []string{"20", "30"}}
	numerator2, _ := sdk.NewIntFromString(testdatadrop.RateBstrArray[0])
	denominator2, _ := sdk.NewIntFromString(testdatadrop.RateBstrArray[1])
	drops1.Rate2 = []sdk.Int{numerator2, denominator2}

	testInput.MarketKeeper.SetDrop(testInput.Context, drops1)

	var d1 = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "40", Rate1: testdata.RateAstrArray, Prev1: "0", Next1: "0", Rate2: testdata.RateBstrArray, Prev2: "1", Next2: "0"}
	_, err1 := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d1)
	require.Error(t, err1)
	require.ErrorContains(t, err1, "Drop rate2 greater than Prev2 rate2")

}

func TestCreateDrop_Case4_Side1_Negative(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "30CoinA", coinBStr: "40CoinB", RateAstrArray: []string{"60", "70"}, RateBstrArray: []string{"80", "90"}}
	coinPair, _ := sample.SampleCoins("100CoinA", "100CoinB")
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
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	//validate CreatePool
	require.NoError(t, err)

	//validate SetUidCount function.
	aftercount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)

	//validate GetDrop
	drops, dropFound := testInput.MarketKeeper.GetDrop(testInput.Context, beforecount)
	require.True(t, dropFound)
	require.Equal(t, drops.Pair, pair)
	//validate CreateDrop
	var d = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "70", Rate1: testdata.RateAstrArray, Prev1: "1", Next1: "1", Rate2: testdata.RateBstrArray, Prev2: "0", Next2: "0"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.Error(t, err)
	require.ErrorContains(t, err, "Prev1 drop not active")

	//HappyPath
	d.Prev1 = "0"
	d.Next1 = "0"
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.NoError(t, err)
	//validate GetDrop
	drops1, drop1Found := testInput.MarketKeeper.GetDrop(testInput.Context, aftercount)

	require.True(t, drop1Found)
	//Update Prev1 and Next1 value to 1
	drops1.Prev1 = aftercount
	drops1.Next1 = aftercount

	testInput.MarketKeeper.SetDrop(testInput.Context, drops1)

	var d1 = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "40", Rate1: testdata.RateAstrArray, Prev1: "1", Next1: "2", Rate2: testdata.RateBstrArray, Prev2: "0", Next2: "0"}
	_, err1 := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d1)
	require.Error(t, err1)
	require.ErrorContains(t, err1, "Next1 drop not active")

	drops1.Next1 = uint64(3)
	testInput.MarketKeeper.SetDrop(testInput.Context, drops1)

	d1 = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "40", Rate1: testdata.RateAstrArray, Prev1: "1", Next1: "1", Rate2: testdata.RateBstrArray, Prev2: "0", Next2: "0"}
	_, err1 = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d1)
	require.Error(t, err1)
	require.ErrorContains(t, err1, "Prev1 and Next1 are not adjacent")

}

func TestCreateDrop_Case4_Side2_Negative(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "30CoinA", coinBStr: "40CoinB", RateAstrArray: []string{"60", "70"}, RateBstrArray: []string{"80", "90"}}
	coinPair, _ := sample.SampleCoins("100CoinA", "100CoinB")
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
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	//validate CreatePool
	require.NoError(t, err)

	//validate SetUidCount function.
	aftercount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)

	//validate GetDrop
	drops, dropFound := testInput.MarketKeeper.GetDrop(testInput.Context, beforecount)
	require.True(t, dropFound)
	require.Equal(t, drops.Pair, pair)
	//validate CreateDrop
	var d = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "70", Rate1: testdata.RateAstrArray, Prev1: "0", Next1: "0", Rate2: testdata.RateBstrArray, Prev2: "1", Next2: "1"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.Error(t, err)
	require.ErrorContains(t, err, "Prev2 drop not active")

	//HappyPath
	d.Prev2 = "0"
	d.Next2 = "0"
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.NoError(t, err)
	//validate GetDrop
	drops1, drop1Found := testInput.MarketKeeper.GetDrop(testInput.Context, aftercount)

	require.True(t, drop1Found)
	//Update Prev2 and Next2 value to 1
	drops1.Prev2 = aftercount
	drops1.Next2 = aftercount

	testInput.MarketKeeper.SetDrop(testInput.Context, drops1)

	var d1 = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "40", Rate1: testdata.RateAstrArray, Prev1: "0", Next1: "0", Rate2: testdata.RateBstrArray, Prev2: "1", Next2: "2"}
	_, err1 := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d1)
	require.Error(t, err1)
	require.ErrorContains(t, err1, "Next2 drop not active")

	drops1.Next2 = uint64(3)
	testInput.MarketKeeper.SetDrop(testInput.Context, drops1)

	d1 = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "40", Rate1: testdata.RateAstrArray, Prev1: "0", Next1: "0", Rate2: testdata.RateBstrArray, Prev2: "1", Next2: "1"}
	_, err1 = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d1)
	require.Error(t, err1)
	require.ErrorContains(t, err1, "Prev2 and Next2 are not adjacent")

}

func TestCreateDrop_ValidateSenderBalance(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "30CoinA", coinBStr: "40CoinB", RateAstrArray: []string{"30", "40"}, RateBstrArray: []string{"50", "60"}}
	coinPair, _ := sample.SampleCoins("35CoinA", "45CoinB")
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
	var d = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "20", Rate1: testdata.RateAstrArray, Prev1: "0", Next1: "0", Rate2: testdata.RateBstrArray, Prev2: "0", Next2: "0"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.Error(t, err)
	require.ErrorContains(t, err, "insufficient balance")

}

/*
func TestCreateDrop_InvalidDrop(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "30CoinA", coinBStr: "40CoinB", RateAstrArray: []string{"30", "40"}, RateBstrArray: []string{"50", "60"}}
	coinPair, _ := sample.SampleCoins("35CoinA", "45CoinB")
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
	var d = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "-1", Rate1: testdata.RateAstrArray, Prev1: "0", Next1: "0", Rate2: testdata.RateBstrArray, Prev2: "0", Next2: "0"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.Error(t, err)
	//require.ErrorContains(t, err, "insufficient balance")

}*/
