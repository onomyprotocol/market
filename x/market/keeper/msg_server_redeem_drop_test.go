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
	var d = types.MsgCreateDrop{Creator: addr, Pair: pair, Drops: "120"}
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
	require.Equal(t, 1, len(rst.Leaders))
	require.Equal(t, "1320", rst.Leaders[0].Drops.String())

	// Validate GetDrop
	drops1, drop1Found := testInput.MarketKeeper.GetDrop(testInput.Context, aftercount)
	require.True(t, drop1Found)
	require.Equal(t, drops1.Pair, pair)
	require.Equal(t, drops1.Drops.String(), d.Drops)
	require.Contains(t, d.GetCreator(), createDropResponse.String())

	// Validate RedeemDrop
	Uid := strconv.FormatUint(drops1.Uid, 10)
	var rd = types.MsgRedeemDrop{Creator: addr, Uid: Uid}
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
	require.Equal(t, "1200", rst.Leaders[0].Drops.String(), rst)

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
}
