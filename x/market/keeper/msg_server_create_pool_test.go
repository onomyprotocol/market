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

var addr string = sample.AccAddress()

func TestCreatePool(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	// CoinAmsg and CoinBmsg pre-sort from raw msg
	coinA, err := sdk.ParseCoinNormalized("20CoinA")
	if err != nil {
		require.Error(t, err)
	}

	coinB, err := sdk.ParseCoinNormalized("20CoinB")
	if err != nil {
		require.Error(t, err)
	}

	coinPair := sdk.NewCoins(coinA, coinB)
	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, types.ModuleName, coinPair))
	requestAddress, _ := sdk.AccAddressFromBech32(addr)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, types.ModuleName, requestAddress, coinPair))
	beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context) // GetUidCount before CreatePool
	var p = types.MsgCreatePool{CoinA: "20CoinA", CoinB: "20CoinB", Creator: addr, RateA: []string{"10", "20"}, RateB: []string{"20", "30"}}
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
	pair := strings.Join([]string{"CoinA", "CoinB"}, ",")
	//pair1 := strings.Join([]string{"CoinB", "CoinA"}, ",")
	rst, found := testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, found)
	require.Equal(t, rst.Pair, pair)
	//validate GetMember
	members, memberfound := testInput.MarketKeeper.GetMemberWithPair(testInput.Context, pair, "CoinB", "CoinA")
	members1, memberfound1 := testInput.MarketKeeper.GetMemberWithPair(testInput.Context, pair, "CoinA", "CoinB")
	require.True(t, memberfound)
	require.Equal(t, members.DenomA, "CoinB")
	require.Equal(t, members.DenomB, "CoinA")
	require.True(t, memberfound1)
	require.Equal(t, members1.DenomA, "CoinA")
	require.Equal(t, members1.DenomB, "CoinB")
	//validate GetDrop
	drops, dropFound := testInput.MarketKeeper.GetDrop(testInput.Context, aftercount)
	require.True(t, dropFound)
	require.Equal(t, drops.Pair, pair)

}

func TestCreatePool_LessCoinPair(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)

	scenarios := []struct {
		coinAStr      string
		coinBStr      string
		RateAstrArray []string
		RateBstrArray []string
	}{
		{coinAStr: "10CoinA", coinBStr: "20CoinB", RateAstrArray: []string{"10", "20"}, RateBstrArray: []string{"20", "30"}},
		{coinAStr: "20CoinA", coinBStr: "10CoinB", RateAstrArray: []string{"10", "20"}, RateBstrArray: []string{"20", "30"}},
	}
	for _, s := range scenarios {
		coinA, err := sdk.ParseCoinNormalized("20CoinA")
		if err != nil {
			require.Error(t, err)
		}

		coinB, err := sdk.ParseCoinNormalized("20CoinB")
		if err != nil {
			require.Error(t, err)
		}

		coinPair := sdk.NewCoins(coinA, coinB)
		require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, types.ModuleName, coinPair))
		requestAddress, _ := sdk.AccAddressFromBech32(addr)
		require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, types.ModuleName, requestAddress, coinPair))
		var p = types.MsgCreatePool{CoinA: s.coinAStr, CoinB: s.coinBStr, Creator: addr, RateA: s.RateAstrArray, RateB: s.RateBstrArray}
		response, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
		require.NoError(t, err)
		require.Contains(t, p.GetCreator(), response.String())

	}

}

func TestCreatePool_Insufficient_Funds(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	coinA, err := sdk.ParseCoinNormalized("10CoinA")
	if err != nil {
		require.Error(t, err)
	}

	coinB, err := sdk.ParseCoinNormalized("10CoinB")
	if err != nil {
		require.Error(t, err)
	}

	coinPair := sdk.NewCoins(coinA, coinB)
	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, types.ModuleName, coinPair))
	requestAddress, _ := sdk.AccAddressFromBech32(addr)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, types.ModuleName, requestAddress, coinPair))
	var p = types.MsgCreatePool{CoinA: "15CoinA", CoinB: "15CoinB", Creator: addr, RateA: []string{"10", "20"}, RateB: []string{"20", "30"}}
	response, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	require.Error(t, err)
	require.NotContains(t, p.GetCreator(), response.String())

}

func TestCreatePool_Insufficient_Funds_ReSubmit(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	coinA, err := sdk.ParseCoinNormalized("20CoinA")
	if err != nil {
		require.Error(t, err)
	}

	coinB, err := sdk.ParseCoinNormalized("20CoinB")
	if err != nil {
		require.Error(t, err)
	}

	coinPair := sdk.NewCoins(coinA, coinB)
	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, types.ModuleName, coinPair))
	requestAddress, _ := sdk.AccAddressFromBech32(addr)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, types.ModuleName, requestAddress, coinPair))
	var p = types.MsgCreatePool{CoinA: "15CoinA", CoinB: "15CoinB", Creator: addr, RateA: []string{"10", "20"}, RateB: []string{"20", "30"}}
	var p1 = types.MsgCreatePool{CoinA: "30CoinA", CoinB: "30CoinB", Creator: addr, RateA: []string{"10", "20"}, RateB: []string{"20", "30"}}
	response, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	response1, err1 := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p1)
	require.NoError(t, err)
	require.Error(t, err1)
	//require.(t, err1, "insufficient funds")
	require.Contains(t, p.GetCreator(), response.String())
	require.NotContains(t, p.GetCreator(), response1.String())

}

/*
func TestCreatePool_Invalid_Coins(t *testing.T) {
	testInput, ctx := keepertest.CreateTestEnvironment(t)

	scenarios := []struct {
		coinAStr      string
		coinBStr      string
		RateAstrArray []string
		RateBstrArray []string
	}{
		{coinAStr: "hsjfs", coinBStr: "20CoinB", RateAstrArray: []string{"10", "20"}, RateBstrArray: []string{"20", "30"}},
		{coinAStr: "20CoinA", coinBStr: "jsfkjsjhf", RateAstrArray: []string{"10", "20"}, RateBstrArray: []string{"20", "30"}},
		{coinAStr: "20CoinA", coinBStr: "20", RateAstrArray: []string{"10", "20"}, RateBstrArray: []string{"20", "30"}},
		{coinAStr: "20Coin", coinBStr: "20CoinB", RateAstrArray: []string{"10", "20"}, RateBstrArray: []string{"20", "30"}},
		{coinAStr: "20CoinA", coinBStr: "20CoinB", RateAstrArray: []string{"awsrrerefefw", "awsrrerefefw"}, RateBstrArray: []string{"20", "30"}},
	}
	for _, s := range scenarios {
		coinA, err := sdk.ParseCoinNormalized("20CoinA")
		if err != nil {
			require.Error(t, err)
		}

		coinB, err := sdk.ParseCoinNormalized("20CoinB")
		if err != nil {
			require.Error(t, err)
		}

		coinPair := sdk.NewCoins(coinA, coinB)
		require.NoError(t, testInput.BankKeeper.MintCoins(ctx, types.ModuleName, coinPair))
		requestAddress, _ := sdk.AccAddressFromBech32(addr)
		require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, requestAddress, coinPair))
		var p = types.MsgCreatePool{CoinA: s.coinAStr, CoinB: s.coinBStr, Creator: addr, RateA: s.RateAstrArray, RateB: s.RateBstrArray}
		response, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(ctx), &p)
		require.Error(t, err)
		require.NotContains(t, p.GetCreator(), response.String())

	}

}
*/
