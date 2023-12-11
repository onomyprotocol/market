package keeper_test

import (
	"strconv"
	"strings"
	"testing"

	keepertest "market/testutil/keeper"
	"market/testutil/sample"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"market/x/market/keeper"
	"market/x/market/types"

	"github.com/stretchr/testify/require"
)

func TestCreateDrop(t *testing.T) {
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

	// MintCoins
	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, types.ModuleName, coinPair))

	// SendCoinsFromModuleToAccount
	requestAddress2, err := sdk.AccAddressFromBech32(addr2)
	require.NoError(t, err)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, types.ModuleName, requestAddress2, coinPair))

	// MintCoins
	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, types.ModuleName, coinPair))

	// SendCoinsFromModuleToAccount
	requestAddress3, err := sdk.AccAddressFromBech32(addr3)
	require.NoError(t, err)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, types.ModuleName, requestAddress3, coinPair))

	// GetUidCount before CreatePool
	beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)

	// Create Pool
	var p = types.MsgCreatePool{CoinA: testdata.coinAStr, CoinB: testdata.coinBStr, Creator: addr}
	response, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)

	// Validate CreatePool
	require.NoError(t, err)
	require.Contains(t, p.GetCreator(), response.String())
	require.Contains(t, p.GetCoinA(), response.String())
	require.Contains(t, p.GetCoinB(), response.String())

	// Validate SetUidCount function.
	aftercount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)

	// Validate GetDrop
	drops, dropFound := testInput.MarketKeeper.GetDrop(testInput.Context, beforecount)
	require.True(t, dropFound)
	require.Equal(t, drops.Pair, pair)

	owner, ok := testInput.MarketKeeper.GetDropsOwnerPair(testInput.Context, addr, pair)
	require.True(t, ok)
	require.Truef(t, owner.Sum.Equal(sdk.NewInt(1200)), owner.Sum.String())

	// Validate GetPool
	rst1, found := testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, found)
	require.Equal(t, rst1.Pair, pair)
	require.Equal(t, "1200", rst1.Drops.String())
	require.Equal(t, 1, len(rst1.Leaders))
	require.Equal(t, "1200", rst1.Leaders[0].Drops.String())

	beforecount = aftercount

	// Validate CreateDrop
	var d = types.MsgCreateDrop{Creator: addr2, Pair: pair, Drops: "120"}
	createDropResponse, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.NoError(t, err)

	// Validate SetUidCount function.
	aftercount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)

	pairs, ok := testInput.MarketKeeper.GetDropPairs(testInput.Context, addr)
	require.True(t, ok)
	require.Truef(t, pairs.Pairs[0] == pair, pairs.String())

	// Validate GetPool
	rst, found := testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, found)
	require.Equal(t, rst.Pair, pair)
	require.Equal(t, "1320", rst.Drops.String())
	require.Equalf(t, addr, rst.Leaders[0].Address, rst.Leaders[0].Address)
	require.Equalf(t, 2, len(rst.Leaders), rst.Leaders[1].Address)
	require.Equal(t, "1200", rst.Leaders[0].Drops.String())

	// Validate GetMember
	members, memberfound := testInput.MarketKeeper.GetMember(testInput.Context, denomB, denomA)
	members1, memberfound1 := testInput.MarketKeeper.GetMember(testInput.Context, denomA, denomB)
	require.True(t, memberfound)
	require.Equal(t, members.DenomA, denomB)
	require.Equal(t, members.DenomB, denomA)
	require.Equal(t, "33", members.Balance.String())

	require.True(t, memberfound1)
	require.Equal(t, members1.DenomA, denomA)
	require.Equal(t, members1.DenomB, denomB)
	require.Equal(t, "44", members1.Balance.String())

	owner, ok = testInput.MarketKeeper.GetDropsOwnerPair(testInput.Context, addr, pair)
	require.True(t, ok)
	require.Truef(t, owner.Sum.Equal(sdk.NewInt(1200)), owner.Sum.String())

	// Validate GetDrop
	drops, dropFound = testInput.MarketKeeper.GetDrop(testInput.Context, beforecount)
	require.True(t, dropFound)
	require.Equal(t, drops.Pair, pair)
	require.Equal(t, drops.Drops.String(), d.Drops)
	require.Contains(t, d.GetCreator(), createDropResponse.String())

	beforecount = aftercount

	// Validate CreateDrop
	var e = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "120"}
	createDropResponse, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &e)
	require.NoError(t, err)

	// Validate SetUidCount function.
	aftercount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)

	pairs, ok = testInput.MarketKeeper.GetDropPairs(testInput.Context, addr)
	require.True(t, ok)
	require.Truef(t, pairs.Pairs[0] == pair, pairs.String())

	// Validate GetPool
	rst2, found := testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, found)
	require.Equal(t, rst2.Pair, pair)
	require.Equal(t, "1440", rst2.Drops.String())
	require.Equal(t, "1320", rst2.Leaders[0].Drops.String())
	require.Equal(t, rst2.Leaders[0].Address, addr)
	require.Equal(t, rst2.Leaders[1].Address, addr2)
	require.Equal(t, 2, len(rst2.Leaders))
	require.Equal(t, "1320", rst2.Leaders[0].Drops.String())

	// Validate GetDrop
	drops, dropFound = testInput.MarketKeeper.GetDrop(testInput.Context, beforecount)
	require.True(t, dropFound)
	require.Equal(t, drops.Pair, pair)
	require.Equal(t, drops.Drops.String(), e.Drops)
	require.Contains(t, d.GetCreator(), createDropResponse.String())

	beforecount = aftercount

	// Validate CreateDrop
	var f = types.MsgCreateDrop{Creator: addr3, Pair: pair, Drops: "1000"}
	createDropResponse, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &f)
	require.NoError(t, err)

	// Validate SetUidCount function.
	aftercount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)

	pairs, ok = testInput.MarketKeeper.GetDropPairs(testInput.Context, addr3)
	require.True(t, ok)
	require.Truef(t, pairs.Pairs[0] == pair, pairs.String())

	// Validate GetPool
	rst, found = testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, found)
	require.Equal(t, rst.Pair, pair)
	require.Equal(t, "2440", rst.Drops.String())
	require.Equal(t, 3, len(rst.Leaders))
	require.Equal(t, "1320", rst.Leaders[0].Drops.String())

	// Validate GetDrop
	drops, dropFound = testInput.MarketKeeper.GetDrop(testInput.Context, beforecount)
	require.True(t, dropFound)
	require.Equal(t, drops.Pair, pair)
	require.Equal(t, drops.Drops.String(), f.Drops)
	require.Contains(t, d.GetCreator(), createDropResponse.String())

	beforecount = aftercount

	// Validate CreateDrop
	var g = types.MsgCreateDrop{Creator: addr3, Pair: pair, Drops: "400"}
	createDropResponse, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &g)
	require.NoError(t, err)

	// Validate SetUidCount function.
	aftercount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)

	// Validate GetPool
	rst, found = testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, found)
	require.Equal(t, rst.Pair, pair)
	require.Equal(t, "2840", rst.Drops.String())
	require.Equal(t, 3, len(rst.Leaders))
	require.Equal(t, "1400", rst.Leaders[0].Drops.String())
	require.Equal(t, "1320", rst.Leaders[1].Drops.String())
	require.Equal(t, "120", rst.Leaders[2].Drops.String())
	require.Equalf(t, addr3, rst.Leaders[0].Address, rst.Leaders[0].Address)
	require.Equalf(t, addr, rst.Leaders[1].Address, addr3)
	require.Equalf(t, addr2, rst.Leaders[2].Address, rst.Leaders[2].Address)

	// Validate GetDrop
	drops, dropFound = testInput.MarketKeeper.GetDrop(testInput.Context, beforecount)
	require.True(t, dropFound)
	require.Equal(t, drops.Pair, pair)
	require.Equal(t, drops.Drops.String(), g.Drops)
	require.Contains(t, d.GetCreator(), createDropResponse.String())

}

func TestCreateDrop_Pool_Not_Found(t *testing.T) {
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

	// GetUidCount before CreatePool
	beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)

	// Create Pool
	var p = types.MsgCreatePool{CoinA: testdata.coinAStr, CoinB: testdata.coinBStr, Creator: addr}
	response, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)

	// Validate CreatePool
	require.NoError(t, err)
	require.Contains(t, p.GetCreator(), response.String())
	require.Contains(t, p.GetCoinA(), response.String())
	require.Contains(t, p.GetCoinB(), response.String())

	// Validate SetUidCount function.
	aftercount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)

	// Validate GetDrop
	drops, dropFound := testInput.MarketKeeper.GetDrop(testInput.Context, beforecount)
	require.True(t, dropFound)
	require.Equal(t, drops.Pair, pair)

	// Validate CreateDrop
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
		var d = types.MsgCreateDrop{Creator: s.Creator, Pair: pair, Drops: "70"}
		_, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
		require.Error(t, err)
		require.ErrorContains(t, err, "the pool not found")

	}
}

func TestCreateDrop_Pool_Not_Active(t *testing.T) {
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

	//Validate RedeemDrop
	Uid := strconv.FormatUint(drops.Uid, 10)
	var rd = types.MsgRedeemDrop{Creator: addr, Uid: Uid}
	createRedeemDropResponse, redeemdropErr := keeper.NewMsgServerImpl(*testInput.MarketKeeper).RedeemDrop(sdk.WrapSDKContext(testInput.Context), &rd)
	require.NoError(t, redeemdropErr)
	require.Contains(t, rd.GetCreator(), createRedeemDropResponse.String())

	//validate CreateDrop (Inactive Pool)
	scenarios := []struct {
		coinAStr      string
		coinBStr      string
		RateAstrArray []string
		RateBstrArray []string
		Creator       string
	}{
		{coinAStr: "20CoinA", coinBStr: "20CoinB", RateAstrArray: []string{"10", "20"}, RateBstrArray: []string{"20", "30"}, Creator: addr},
		{coinAStr: "20CoinB", coinBStr: "20CoinA", RateAstrArray: []string{"10", "20"}, RateBstrArray: []string{"20", "30"}, Creator: sample.AccAddress()},
	}
	for _, s := range scenarios {
		coinPair, _ = sample.SampleCoins(s.coinAStr, s.coinBStr)
		denomA, denomB = sample.SampleDenoms(coinPair)
		pair = strings.Join([]string{denomA, denomB}, ",")
		var d = types.MsgCreateDrop{Creator: s.Creator, Pair: pair, Drops: "70"}
		_, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
		require.Error(t, err)
		require.ErrorContains(t, err, "the pool is inactive")

	}

	// GetUidCount before CreatePool
	beforecount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	//Create Pool
	p = types.MsgCreatePool{CoinA: testdata.coinAStr, CoinB: testdata.coinBStr, Creator: addr}
	response, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)
	//validate CreatePool
	require.NoError(t, err)
	require.Contains(t, p.GetCreator(), response.String())
	require.Contains(t, p.GetCoinA(), response.String())
	require.Contains(t, p.GetCoinB(), response.String())
	//validate SetUidCount function.
	aftercount = testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)

	//validate CreateDrop (Active Pool)
	scenarios = []struct {
		coinAStr      string
		coinBStr      string
		RateAstrArray []string
		RateBstrArray []string
		Creator       string
	}{
		{coinAStr: "20CoinA", coinBStr: "20CoinB", RateAstrArray: []string{"10", "20"}, RateBstrArray: []string{"20", "30"}, Creator: addr},
	}
	for _, s := range scenarios {
		coinPair, _ = sample.SampleCoins(s.coinAStr, s.coinBStr)
		denomA, denomB = sample.SampleDenoms(coinPair)
		pair = strings.Join([]string{denomA, denomB}, ",")
		var d = types.MsgCreateDrop{Creator: s.Creator, Pair: pair, Drops: "70"}
		_, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)

		require.NoError(t, err)
	}
}

func TestCreateDrop_Negative(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)
	//TestData
	testdata := testData{coinAStr: "30CoinA", coinBStr: "40CoinB"}
	coinPair, _ := sample.SampleCoins("140CoinA", "140CoinB")
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
	var d = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "120"}
	createDropResponse, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.NoError(t, err)
	require.Contains(t, d.GetCreator(), createDropResponse.String())

	//validate GetMember
	members, memberfound := testInput.MarketKeeper.GetMember(testInput.Context, denomB, denomA)
	members1, memberfound1 := testInput.MarketKeeper.GetMember(testInput.Context, denomA, denomB)
	require.True(t, memberfound)
	require.Equal(t, members.DenomA, denomB)
	require.Equal(t, members.DenomB, denomA)
	require.Equal(t, "33", members.Balance.String())

	require.True(t, memberfound1)
	require.Equal(t, members1.DenomA, denomA)
	require.Equal(t, members1.DenomB, denomB)
	require.Equal(t, "44", members1.Balance.String())

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
	var d = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "2000"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.Error(t, err)
	require.ErrorContains(t, err, "insufficient balance")

}

func TestCreateDrop_InvalidDrop(t *testing.T) {

	coinPair, _ := sample.SampleCoins("35CoinA", "45CoinB")
	denomA, denomB := sample.SampleDenoms(coinPair)
	pair := strings.Join([]string{denomA, denomB}, ",")

	// Validate CreateDrop
	dropTest := types.NewMsgCreateDrop(addr, pair, "-1")
	err := dropTest.ValidateBasic()
	require.Error(t, err)

}

func TestZeroAmtPaid(t *testing.T) {

	testInput := keepertest.CreateTestEnvironment(t)

	// TestData
	testdata := testData{coinAStr: "1CoinA", coinBStr: "4000CoinB", RateAstrArray: []string{"60", "70"}, RateBstrArray: []string{"80", "90"}}
	coinPair, _ := sample.SampleCoins("70CoinA", "7000CoinB")
	denomA, denomB := sample.SampleDenoms(coinPair)
	pair := strings.Join([]string{denomA, denomB}, ",")

	// MintCoins
	require.NoError(t, testInput.BankKeeper.MintCoins(testInput.Context, types.ModuleName, coinPair))

	// SendCoinsFromModuleToAccount
	requestAddress, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(testInput.Context, types.ModuleName, requestAddress, coinPair))

	// GetUidCount before CreatePool
	beforecount := testInput.MarketKeeper.GetUidCount(testInput.Context)

	// Create Pool
	var p = types.MsgCreatePool{CoinA: testdata.coinAStr, CoinB: testdata.coinBStr, Creator: addr}
	response, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(testInput.Context), &p)

	// Validate CreatePool
	require.NoError(t, err)
	require.Contains(t, p.GetCreator(), response.String())
	require.Contains(t, p.GetCoinA(), response.String())
	require.Contains(t, p.GetCoinB(), response.String())

	// Validate SetUidCount function.
	aftercount := testInput.MarketKeeper.GetUidCount(testInput.Context)
	require.Equal(t, beforecount+1, aftercount)

	// Validate GetDrop
	drops, dropFound := testInput.MarketKeeper.GetDrop(testInput.Context, beforecount)
	require.True(t, dropFound)
	require.Equal(t, drops.Pair, pair)

	// Validate GetPool
	rst1, found := testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, found)
	require.Equal(t, rst1.Pair, pair)
	require.Equal(t, "4000", rst1.Drops.String())
	require.Equal(t, 1, len(rst1.Leaders))
	require.Equal(t, "4000", rst1.Leaders[0].Drops.String())

	// Validate CreateDrop
	var d = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "1"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.Error(t, err)
}
