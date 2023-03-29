package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/pendulum-labs/market/testutil/keeper"
	"github.com/pendulum-labs/market/testutil/sample"

	//"github.com/pendulum-labs/market/testutil/nullify"
	"github.com/pendulum-labs/market/x/market/keeper"
	"github.com/pendulum-labs/market/x/market/types"
	"github.com/stretchr/testify/require"
)

var _ = strconv.IntSize

func TestCreatePool(t *testing.T) {
	k, ctx := keepertest.MarketKeeper(t)
	var p = types.MsgCreatePool{CoinA: "10CoinA", CoinB: "20CoinB", Creator: sample.AccAddress(), RateA: []string{"10", "20"}, RateB: []string{"20", "30"}}
	response, err := keeper.NewMsgServerImpl(*k).CreatePool(sdk.WrapSDKContext(ctx), &p)
	require.NoError(t, err)
	require.Contains(t, p.GetCreator(), response.String())

}
