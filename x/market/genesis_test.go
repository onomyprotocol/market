package market_test

import (
	"testing"

	keepertest "github.com/onomyprotocol/market/testutil/keeper"
	"github.com/onomyprotocol/market/testutil/nullify"
	"github.com/onomyprotocol/market/x/market"
	"github.com/onomyprotocol/market/x/market/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PoolList: []types.Pool{
			{
				Pair:   "0",
				Denom1: "0",
				Denom2: "0",
				Leader: "0",
			},
			{
				Pair:   "1",
				Denom1: "1",
				Denom2: "1",
				Leader: "1",
			},
		},
		DropList: []types.Drop{
			{
				Uid:   0,
				Owner: "0",
				Pair:  "0",
			},
			{
				Uid:   1,
				Owner: "1",
				Pair:  "1",
			},
		},
		MemberList: []types.Member{
			{
				Pair:   "0",
				DenomA: "0",
				DenomB: "0",
			},
			{
				Pair:   "1",
				DenomA: "1",
				DenomB: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MarketKeeper(t)
	market.InitGenesis(ctx, *k, genesisState)
	got := market.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PoolList, got.PoolList)
	require.ElementsMatch(t, genesisState.DropList, got.DropList)
	require.ElementsMatch(t, genesisState.MemberList, got.MemberList)
	// this line is used by starport scaffolding # genesis/test/assert
}
