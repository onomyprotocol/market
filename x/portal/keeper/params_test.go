package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "market/testutil/keeper"
	"market/x/portal/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.PortalKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.ProviderChannel, k.ProviderChannel(ctx))
	require.EqualValues(t, params.ReserveChannel, k.ReserveChannel(ctx))
}
