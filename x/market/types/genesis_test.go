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
