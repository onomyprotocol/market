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

type testData struct {
	coinAStr      string
	coinBStr      string
	RateAstrArray []string
	RateBstrArray []string
}

var addr string = sample.AccAddress()

func TestCreatePool(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "20CoinA", coinBStr: "20CoinB", RateAstrArray: []string{"10", "20"}, RateBstrArray: []string{"20", "30"}}
	coinPair, _ := sample.SampleCoins(testdata.coinAStr, testdata.coinBStr)
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
	//validate GetPool

	rst, found := testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, found)
	require.Equal(t, rst.Pair, pair)
	//validate GetMember
	members, memberfound := testInput.MarketKeeper.GetMember(testInput.Context, denomB, denomA)
	members1, memberfound1 := testInput.MarketKeeper.GetMember(testInput.Context, denomA, denomB)
	require.True(t, memberfound)
	require.Equal(t, members.DenomA, denomB)
	require.Equal(t, members.DenomB, denomA)
	require.True(t, memberfound1)
	require.Equal(t, members1.DenomA, denomA)
	require.Equal(t, members1.DenomB, denomB)
	//validate GetDrop
	drops, dropFound := testInput.MarketKeeper.GetDrop(testInput.Context, beforecount)
	require.True(t, dropFound)
	require.Equal(t, drops.Pair, pair)

}

func TestCreatePool_PoolAlreadyExist(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	count := 0
	scenarios := []struct {
		coinAStr      string
		coinBStr      string
		RateAstrArray []string
		RateBstrArray []string
	}{
		{coinAStr: "20CoinA", coinBStr: "20CoinB", RateAstrArray: []string{"10", "20"}, RateBstrArray: []string{"20", "30"}},
		{coinAStr: "20CoinA", coinBStr: "20CoinB", RateAstrArray: []string{"10", "20"}, RateBstrArray: []string{"20", "30"}},
	}
	for _, s := range scenarios {
		coinPair, _ := sample.SampleCoins("20CoinA", "20CoinB")

		require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, types.ModuleName, coinPair))
		requestAddress, _ := sdk.AccAddressFromBech32(addr)
		require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, types.ModuleName, requestAddress, coinPair))
		var p = types.MsgCreatePool{CoinA: s.coinAStr, CoinB: s.coinBStr, Creator: addr}
		response, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
		if count == 0 {
			require.NoError(t, err)
			require.Contains(t, p.GetCreator(), response.String())

		} else {
			require.Error(t, err) //Pool Already exists
			require.ErrorContains(t, err, "pool already exists")
			require.NotContains(t, p.GetCreator(), response.String())
		}

		count++

	}

}

func TestCreatePool_Insufficient_Funds(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "15CoinA", coinBStr: "15CoinB", RateAstrArray: []string{"10", "20"}, RateBstrArray: []string{"20", "30"}}
	coinPair, _ := sample.SampleCoins("10CoinA", "10CoinB")

	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, types.ModuleName, coinPair))
	requestAddress, _ := sdk.AccAddressFromBech32(addr)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, types.ModuleName, requestAddress, coinPair))
	var p = types.MsgCreatePool{CoinA: testdata.coinAStr, CoinB: testdata.coinBStr, Creator: addr}
	response, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	require.Error(t, err)
	require.ErrorContains(t, err, "insufficient funds")
	require.NotContains(t, p.GetCreator(), response.String())

}

func TestCreatePool_PoolAlready_Exists_ReSubmit(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	coinPair, _ := sample.SampleCoins("20CoinA", "20CoinB")

	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, types.ModuleName, coinPair))
	requestAddress, _ := sdk.AccAddressFromBech32(addr)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, types.ModuleName, requestAddress, coinPair))
	var p = types.MsgCreatePool{CoinA: "15CoinA", CoinB: "15CoinB", Creator: addr}
	var p1 = types.MsgCreatePool{CoinA: "30CoinA", CoinB: "30CoinB", Creator: addr}
	response, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	response1, err1 := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p1)
	require.NoError(t, err)
	require.Error(t, err1)
	require.ErrorContains(t, err1, "pool already exists")
	require.Contains(t, p.GetCreator(), response.String())
	require.NotContains(t, p.GetCreator(), response1.String())

}

func TestCreatePool_With_New_Creator(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "15CoinA", coinBStr: "15CoinB", RateAstrArray: []string{"10", "20"}, RateBstrArray: []string{"20", "30"}}
	coinPair, _ := sample.SampleCoins("10CoinA", "10CoinB")

	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, types.ModuleName, coinPair))
	requestAddress, _ := sdk.AccAddressFromBech32(addr)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, types.ModuleName, requestAddress, coinPair))
	var p = types.MsgCreatePool{CoinA: testdata.coinAStr, CoinB: testdata.coinBStr, Creator: sample.AccAddress()}
	response, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	require.Error(t, err)
	require.ErrorContains(t, err, "insufficient funds")
	require.NotContains(t, p.GetCreator(), response.String())

}

func TestCreatePool_With_Empty_Rates(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "15CoinA", coinBStr: "15CoinB", RateAstrArray: []string{"0", "0"}, RateBstrArray: []string{"0", "0"}}
	coinPair, _ := sample.SampleCoins("20CoinA", "20CoinB")
	denomA, denomB := sample.SampleDenoms(coinPair)
	pair := strings.Join([]string{denomA, denomB}, ",")
	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, types.ModuleName, coinPair))
	requestAddress, _ := sdk.AccAddressFromBech32(addr)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, types.ModuleName, requestAddress, coinPair))
	//validate SetUidCount function.
	beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)

	var p = types.MsgCreatePool{CoinA: testdata.coinAStr, CoinB: testdata.coinBStr, Creator: addr}
	response, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	require.NoError(t, err)
	require.Contains(t, p.GetCreator(), response.String())
	//validate SetUidCount function.
	aftercount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)
	rst, found := testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, found)
	require.Equal(t, rst.Pair, pair)
	//validate GetMember
	members, memberfound := testInput.MarketKeeper.GetMember(testInput.Context, denomB, denomA)
	members1, memberfound1 := testInput.MarketKeeper.GetMember(testInput.Context, denomA, denomB)
	require.True(t, memberfound)
	require.Equal(t, members.DenomA, denomB)
	require.Equal(t, members.DenomB, denomA)
	require.True(t, memberfound1)
	require.Equal(t, members1.DenomA, denomA)
	require.Equal(t, members1.DenomB, denomB)
	//validate GetDrop
	drops, dropFound := testInput.MarketKeeper.GetDrop(testInput.Context, beforecount)
	require.True(t, dropFound)
	require.Equal(t, drops.Pair, pair)

}

func TestCreatePool_With_Swap_Coins(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "15CoinB", coinBStr: "15CoinA", RateAstrArray: []string{"0", "0"}, RateBstrArray: []string{"0", "0"}}
	coinPair, _ := sample.SampleCoins("20CoinA", "20CoinB")
	denomA, denomB := sample.SampleDenoms(coinPair)
	pair := strings.Join([]string{denomA, denomB}, ",")
	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, types.ModuleName, coinPair))
	requestAddress, _ := sdk.AccAddressFromBech32(addr)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, types.ModuleName, requestAddress, coinPair))
	//validate SetUidCount function.
	beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)

	var p = types.MsgCreatePool{CoinA: testdata.coinAStr, CoinB: testdata.coinBStr, Creator: addr}
	response, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	require.NoError(t, err)
	require.Contains(t, p.GetCreator(), response.String())
	//validate SetUidCount function.
	aftercount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)
	rst, found := testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, found)
	require.Equal(t, rst.Pair, pair)
	//validate GetMember
	members, memberfound := testInput.MarketKeeper.GetMember(testInput.Context, denomB, denomA)
	members1, memberfound1 := testInput.MarketKeeper.GetMember(testInput.Context, denomA, denomB)
	require.True(t, memberfound)
	require.Equal(t, members.DenomA, denomB)
	require.Equal(t, members.DenomB, denomA)
	require.True(t, memberfound1)
	require.Equal(t, members1.DenomA, denomA)
	require.Equal(t, members1.DenomB, denomB)
	//validate GetDrop
	drops, dropFound := testInput.MarketKeeper.GetDrop(testInput.Context, beforecount)
	require.True(t, dropFound)
	require.Equal(t, drops.Pair, pair)

}

func TestCreatePool_Invalid_Coins(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)

	scenarios := []struct {
		coinAStr      string
		coinBStr      string
		RateAstrArray []string
		RateBstrArray []string
	}{
		{coinAStr: "20Coin", coinBStr: "20CoinB", RateAstrArray: []string{"10", "20"}, RateBstrArray: []string{"20", "30"}},
		{coinAStr: "20CoinA", coinBStr: "20Coin", RateAstrArray: []string{"10", "20"}, RateBstrArray: []string{"20", "30"}},
		//{coinAStr: "20CoinA", coinBStr: "20", RateAstrArray: []string{"10", "20"}, RateBstrArray: []string{"20", "30"}},
	}
	for _, s := range scenarios {
		coinPair, _ := sample.SampleCoins("20CoinA", "20CoinB")
		require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, types.ModuleName, coinPair))
		requestAddress, _ := sdk.AccAddressFromBech32(addr)
		require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, types.ModuleName, requestAddress, coinPair))
		var p = types.MsgCreatePool{CoinA: s.coinAStr, CoinB: s.coinBStr, Creator: addr}
		response, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
		require.Error(t, err)
		require.NotContains(t, p.GetCreator(), response.String())

	}

}
