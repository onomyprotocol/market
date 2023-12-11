package keeper_test

import (
	"context"
	"testing"

	keepertest "market/testutil/keeper"
	"market/x/market/keeper"
	"market/x/market/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k := keepertest.CreateTestEnvironment(t)
	return keeper.NewMsgServerImpl(*k.MarketKeeper), sdk.WrapSDKContext(k.Context)
}
