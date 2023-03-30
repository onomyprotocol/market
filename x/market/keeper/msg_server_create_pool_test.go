package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/pendulum-labs/market/testutil/keeper"
	"github.com/pendulum-labs/market/testutil/sample"

	"github.com/pendulum-labs/market/x/market/keeper"
	"github.com/pendulum-labs/market/x/market/types"
	"github.com/stretchr/testify/require"
)

var _ = strconv.IntSize
var addr string = sample.AccAddress()

func TestCreatePool(t *testing.T) {
	testInput, ctx := keepertest.CreateTestEnvironment(t)
	// CoinAmsg and CoinBmsg pre-sort from raw msg
	coinA, err := sdk.ParseCoinNormalized("20CoinA")
	if err != nil {
		panic(err)
	}

	coinB, err := sdk.ParseCoinNormalized("20CoinB")
	if err != nil {
		panic(err)
	}

	coinPair := sdk.NewCoins(coinA, coinB)
	require.NoError(t, testInput.BankKeeper.MintCoins(ctx, types.ModuleName, coinPair))
	requestAddress, _ := sdk.AccAddressFromBech32(addr)
	require.NoError(t, testInput.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, requestAddress, coinPair))
	var p = types.MsgCreatePool{CoinA: "20CoinA", CoinB: "20CoinB", Creator: addr, RateA: []string{"10", "20"}, RateB: []string{"20", "30"}}
	response, _ := keeper.NewMsgServerImpl(*testInput.MarketKeeper).CreatePool(sdk.WrapSDKContext(ctx), &p)
	require.NoError(t, err)
	require.Contains(t, p.GetCreator(), response.String())

}
