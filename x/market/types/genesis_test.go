package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pendulum-labs/market/x/market/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				PoolList: []types.Pool{
					{
						Pair:   "0",
						Denom1: "0",
						Denom2: "0",
						Leader: "0",
						Drops:  sdk.NewIntFromUint64(uint64(0)),
					},
					{
						Pair:   "1",
						Denom1: "1",
						Denom2: "1",
						Leader: "1",
						Drops:  sdk.NewIntFromUint64(uint64(1)),
					},
				},
				DropList: []types.Drop{
					{
						Uid:    0,
						Owner:  "0",
						Pair:   "0",
						Drops:  sdk.NewIntFromUint64(uint64(0)),
						Sum:    sdk.NewIntFromUint64(uint64(0)),
						Active: true,
					},
					{
						Uid:    1,
						Owner:  "1",
						Pair:   "1",
						Drops:  sdk.NewIntFromUint64(uint64(1)),
						Sum:    sdk.NewIntFromUint64(uint64(0)),
						Active: true,
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
						OrderType: "stop",
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
						OrderType: "limit",
						DenomAsk:  "1",
						DenomBid:  "1",
						Amount:    sdk.NewIntFromUint64(uint64(1)),
						Rate:      []sdk.Int{sdk.NewInt(int64(1)), sdk.NewInt(int64(1))},
						Prev:      uint64(1),
						Next:      uint64(1),
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated pool",
			genState: &types.GenesisState{
				PoolList: []types.Pool{
					{
						Pair:   "0",
						Denom1: "0",
						Denom2: "0",
						Leader: "0",
						Drops:  sdk.NewIntFromUint64(uint64(0)),
					},
					{
						Pair:   "0",
						Denom1: "0",
						Denom2: "0",
						Leader: "0",
						Drops:  sdk.NewIntFromUint64(uint64(0)),
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated drop",
			genState: &types.GenesisState{
				DropList: []types.Drop{
					{
						Uid:   0,
						Owner: "0",
						Pair:  "0",
						Drops: sdk.NewIntFromUint64(uint64(0)),
					},
					{
						Uid:   0,
						Owner: "0",
						Pair:  "0",
						Drops: sdk.NewIntFromUint64(uint64(0)),
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated member",
			genState: &types.GenesisState{
				MemberList: []types.Member{
					{
						Pair:    "0",
						DenomA:  "0",
						DenomB:  "0",
						Balance: sdk.NewIntFromUint64(uint64(0)),
					},
					{
						Pair:    "0",
						DenomA:  "0",
						DenomB:  "0",
						Balance: sdk.NewIntFromUint64(uint64(0)),
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated burnings",
			genState: &types.GenesisState{
				BurningsList: []types.Burnings{
					{
						Denom:  "0",
						Amount: sdk.NewIntFromUint64(uint64(0)),
					},
					{
						Denom:  "0",
						Amount: sdk.NewIntFromUint64(uint64(0)),
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated order",
			genState: &types.GenesisState{
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
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
