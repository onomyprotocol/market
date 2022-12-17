package types_test

import (
	"testing"

	"github.com/onomyprotocol/market/x/market/types"
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
				BurningsList: []types.Burnings{
					{
						Denom: "0",
					},
					{
						Denom: "1",
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
					},
					{
						Uid:       1,
						Owner:     "1",
						Active:    false,
						OrderType: "1",
						DenomAsk:  "1",
						DenomBid:  "1",
					},
				},
				AssetList: []types.Asset{
					{
						Active:    true,
						Owner:     "0",
						AssetType: "0",
					},
					{
						Active:    false,
						Owner:     "1",
						AssetType: "1",
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
					},
					{
						Pair:   "0",
						Denom1: "0",
						Denom2: "0",
						Leader: "0",
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
					},
					{
						Uid:   0,
						Owner: "0",
						Pair:  "0",
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
						Pair:   "0",
						DenomA: "0",
						DenomB: "0",
					},
					{
						Pair:   "0",
						DenomA: "0",
						DenomB: "0",
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
						Denom: "0",
					},
					{
						Denom: "0",
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
					},
					{
						Uid:       0,
						Owner:     "0",
						Active:    true,
						OrderType: "0",
						DenomAsk:  "0",
						DenomBid:  "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated asset",
			genState: &types.GenesisState{
				AssetList: []types.Asset{
					{
						Active:    true,
						Owner:     "0",
						AssetType: "0",
					},
					{
						Active:    true,
						Owner:     "0",
						AssetType: "0",
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
