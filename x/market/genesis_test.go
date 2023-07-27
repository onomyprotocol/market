package market_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/pendulum-labs/market/testutil/keeper"
	"github.com/pendulum-labs/market/testutil/nullify"
	"github.com/pendulum-labs/market/x/market"
	"github.com/pendulum-labs/market/x/market/types"
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
				Leaders: []*types.Leader{
					{
						Address: "0",
						Drops:   sdk.NewIntFromUint64(uint64(0)),
					},
				},
				Drops: sdk.NewIntFromUint64(uint64(0)),
			},
			{
				Pair:   "1",
				Denom1: "1",
				Denom2: "1",
				Leaders: []*types.Leader{
					{
						Address: "1",
						Drops:   sdk.NewIntFromUint64(uint64(1)),
					},
				},
				Drops: sdk.NewIntFromUint64(uint64(1)),
			},
		},
		DropList: []types.Drop{
			{
				Uid:     0,
				Owner:   "0",
				Pair:    "0",
				Drops:   sdk.NewIntFromUint64(uint64(0)),
				Product: sdk.NewIntFromUint64(uint64(0)),
				Active:  true,
			},
			{
				Uid:     1,
				Owner:   "1",
				Pair:    "1",
				Drops:   sdk.NewIntFromUint64(uint64(1)),
				Product: sdk.NewIntFromUint64(uint64(0)),
				Active:  true,
			},
		},
		MemberList: []types.Member{
			{
				Pair:     "0",
				DenomA:   "0",
				DenomB:   "0",
				Balance:  sdk.NewIntFromUint64(uint64(0)),
				Previous: sdk.NewIntFromUint64(uint64(0)),
				Limit:    uint64(0),
				Stop:     uint64(0),
			},
			{
				Pair:     "1",
				DenomA:   "1",
				DenomB:   "1",
				Balance:  sdk.NewIntFromUint64(uint64(1)),
				Previous: sdk.NewIntFromUint64(uint64(1)),
				Limit:    uint64(1),
				Stop:     uint64(1),
			},
		},
		BurningsList: []types.Burnings{
			{
				Denom:  "0",
				Amount: sdk.NewIntFromUint64(uint64(0)),
			},
			{
				Denom:  "1",
				Amount: sdk.NewIntFromUint64(uint64(1)),
			},
		},
		OrderList: []types.Order{
			{
				Uid:       0,
				Owner:     "0",
				Active:    true,
				OrderType: "0",
				DenomAsk:  "0",
				DenomBid:  "0",
				Amount:    sdk.NewIntFromUint64(uint64(0)),
				Rate:      []sdk.Int{sdk.NewInt(int64(0)), sdk.NewInt(int64(0))},
				Prev:      uint64(0),
				Next:      uint64(0),
			},
			{
				Uid:       1,
				Owner:     "1",
				Active:    false,
				OrderType: "1",
				DenomAsk:  "1",
				DenomBid:  "1",
				Amount:    sdk.NewIntFromUint64(uint64(1)),
				Rate:      []sdk.Int{sdk.NewInt(int64(1)), sdk.NewInt(int64(1))},
				Prev:      uint64(1),
				Next:      uint64(1),
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k := keepertest.CreateTestEnvironment(t)
	market.InitGenesis(k.Context, *k.MarketKeeper, genesisState)
	got := market.ExportGenesis(k.Context, *k.MarketKeeper)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PoolList, got.PoolList)
	require.ElementsMatch(t, genesisState.DropList, got.DropList)
	require.ElementsMatch(t, genesisState.MemberList, got.MemberList)
	require.ElementsMatch(t, genesisState.BurningsList, got.BurningsList)
	require.ElementsMatch(t, genesisState.OrderList, got.OrderList)
	// this line is used by starport scaffolding # genesis/test/assert
}
