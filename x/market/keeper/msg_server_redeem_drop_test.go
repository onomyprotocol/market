package keeper_test

import (
	"strconv"
	"strings"
	"testing"

	keepertest "market/testutil/keeper"
	"market/testutil/sample"
	"market/x/market/keeper"
	"market/x/market/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestRedeemDrop(t *testing.T) {
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

	// Validate GetPool
	rst1, found := testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, found)
	require.Equal(t, rst1.Pair, pair)
	require.Equal(t, "1200", rst1.Drops.String())
	require.Equal(t, 1, len(rst1.Leaders))
	require.Equal(t, "1200", rst1.Leaders[0].Drops.String())

	// Validate CreateDrop
	var d = types.MsgCreateDrop{Creator: addr2, Pair: pair, Drops: "120"}
	createDropResponse, err := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.NoError(t, err)

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

	// Validate GetPool
	rst, found := testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, found)
	require.Equal(t, rst.Pair, pair)
	require.Equal(t, "1320", rst.Drops.String())
	require.Equal(t, 2, len(rst.Leaders))
	require.Equal(t, addr2, rst.Leaders[1].Address)
	require.Equal(t, "1200", rst.Leaders[0].Drops.String())

	// Validate GetDrop
	drops1, drop1Found := testInput.MarketKeeper.GetDrop(testInput.Context, aftercount)
	require.True(t, drop1Found)
	require.Equal(t, drops1.Pair, pair)
	require.Equal(t, drops1.Drops.String(), d.Drops)
	require.Contains(t, d.GetCreator(), createDropResponse.String())

	// Validate RedeemDrop
	Uid := strconv.FormatUint(drops1.Uid, 10)
	var rd = types.MsgRedeemDrop{Creator: addr2, Uid: Uid}
	createRedeemDropResponse, redeemdropErr := keeper.NewMsgServerImpl(*testInput.MarketKeeper).RedeemDrop(sdk.WrapSDKContext(testInput.Context), &rd)
	require.NoError(t, redeemdropErr)
	require.Contains(t, rd.GetCreator(), createRedeemDropResponse.String())

	// Validate Drop After Redeem Drop
	drops1, drop1Found = testInput.MarketKeeper.GetDrop(testInput.Context, aftercount)
	require.True(t, drop1Found)
	require.Equal(t, drops1.Pair, pair)
	require.Equal(t, drops1.Drops.String(), d.Drops)
	require.Contains(t, d.GetCreator(), createDropResponse.String())
	require.False(t, drops1.Active)

	// Validate GetPool After Redeem Drop
	rst, found = testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, found)
	require.Equal(t, rst.Pair, pair)
	require.Equal(t, "1200", rst.Drops.String())
	require.Equal(t, "1200", rst.Leaders[0].Drops.String())
	require.Equal(t, 1, len(rst.Leaders))
	require.Equal(t, addr, rst.Leaders[0].Address)

	owner, ok := testInput.MarketKeeper.GetDropsOwnerPair(testInput.Context, addr, pair)
	require.True(t, ok)
	require.Truef(t, owner.Sum.Equal(sdk.NewInt(1200)), owner.Sum.String())

	pairs, ok := testInput.MarketKeeper.GetPairs(testInput.Context, addr)
	require.True(t, ok)
	require.Truef(t, pairs.Pairs[0] == pair, pairs.String())

	// Validate GetMember After Redeem Drop
	members, memberfound = testInput.MarketKeeper.GetMember(testInput.Context, denomB, denomA)
	members1, memberfound1 = testInput.MarketKeeper.GetMember(testInput.Context, denomA, denomB)
	require.True(t, memberfound)
	require.Equal(t, members.DenomA, denomB)
	require.Equal(t, members.DenomB, denomA)
	require.Equal(t, members.Balance.String(), "30")
	require.True(t, memberfound1)
	require.Equal(t, members1.DenomA, denomA)
	require.Equal(t, members1.DenomB, denomB)
	require.Equal(t, members1.Balance.String(), "40")

	// Validate RedeemDrop
	Uid2 := strconv.FormatUint(beforecount, 10)
	var rd2 = types.MsgRedeemDrop{Creator: addr, Uid: Uid2}
	createRedeemDropResponse2, redeemdropErr2 := keeper.NewMsgServerImpl(*testInput.MarketKeeper).RedeemDrop(sdk.WrapSDKContext(testInput.Context), &rd2)
	require.NoError(t, redeemdropErr2)
	require.Contains(t, rd2.GetCreator(), createRedeemDropResponse2.String())

	// Validate GetPool After Redeem Drop
	rst, found = testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, found)
	require.Equal(t, rst.Pair, pair)
	require.Equal(t, rst.Drops.String(), "0")
	require.Equal(t, 0, len(rst.Leaders))

	pairs, ok = testInput.MarketKeeper.GetPairs(testInput.Context, addr)
	require.True(t, ok)
	require.Truef(t, len(pairs.Pairs) == 0, pairs.String())
}

func TestRedeemDrop_WithBurnCoin(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)

	require.Equal(t, testInput.MarketKeeper.BurnCoin(testInput.Context), "stake")

	// TestData
	testdata := testData{coinAStr: "100stake", coinBStr: "700CoinB"}
	coinPair, _ := sample.SampleCoins("1000000000stake", "1000000000CoinB")
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

	// Create Drop
	var d = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "123450000"}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.NoError(t, err)

	// redeem the drop
	Uid := strconv.FormatUint(2, 10)
	var rd = types.MsgRedeemDrop{Creator: addr, Uid: Uid}
	_, redeemdropErr := keeper.NewMsgServerImpl(*testInput.MarketKeeper).RedeemDrop(sdk.WrapSDKContext(testInput.Context), &rd)
	require.NoError(t, redeemdropErr)

	rst1, found := testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, found)
	require.Equal(t, rst1.Pair, pair)
	require.Equal(t, "70000", rst1.Drops.String())
}

func TestRedeemDrop_NumericalLimits(t *testing.T) {
	testInput := keepertest.CreateTestEnvironment(t)

	require.Equal(t, testInput.MarketKeeper.BurnCoin(testInput.Context), "stake")

	// TestData
	testdata := testData{coinAStr: keepertest.MaxSupportedCoin("stake"), coinBStr: keepertest.MaxSupportedCoin("CoinB")}
	coinPair, _ := sample.SampleCoins(keepertest.FundMaxSupported("stake"), keepertest.FundMaxSupported("CoinB"))
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

	// Create Drop
	var d = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: keepertest.MaxSupportedDrop("")}
	_, err = keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreateDrop(sdk.WrapSDKContext(testInput.Context), &d)
	require.NoError(t, err)

	// redeem the drop
	Uid := strconv.FormatUint(2, 10)
	var rd = types.MsgRedeemDrop{Creator: addr, Uid: Uid}
	_, redeemdropErr := keeper.NewMsgServerImpl(*testInput.MarketKeeper).RedeemDrop(sdk.WrapSDKContext(testInput.Context), &rd)
	require.NoError(t, redeemdropErr)

	rst1, found := testInput.MarketKeeper.GetPool(testInput.Context, pair)
	require.True(t, found)
	require.Equal(t, rst1.Pair, pair)
	require.Equal(t, keepertest.MaxSupportedDrop(""), rst1.Drops.String())
}
