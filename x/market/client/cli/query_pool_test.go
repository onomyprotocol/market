package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"market/testutil/network"
	"market/testutil/nullify"
	"market/x/market/client/cli"
	"market/x/market/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithPoolObjects(t *testing.T, n int) (*network.Network, []types.Pool) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		pool := types.Pool{
			Pair:   strconv.Itoa(i),
			Denom1: strconv.Itoa(i),
			Denom2: strconv.Itoa(i),
			Leaders: []*types.Leader{
				{
					Address: strconv.Itoa(i),
					Drops:   sdk.NewIntFromUint64(uint64(i)),
				},
			},
			Drops: sdk.NewIntFromUint64(uint64(i)),
		}
		nullify.Fill(&pool)
		state.PoolList = append(state.PoolList, pool)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.PoolList
}

func TestShowPool(t *testing.T) {
	net, objs := networkWithPoolObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc      string
		idPair    string
		idDenom1  string
		idDenom2  string
		idLeaders []*types.Leader
		idDrops   sdk.Int
		args      []string
		err       error
		obj       types.Pool
	}{
		{
			desc:      "found",
			idPair:    objs[0].Pair,
			idDenom1:  objs[0].Denom1,
			idDenom2:  objs[0].Denom2,
			idLeaders: objs[0].Leaders,
			idDrops:   objs[0].Drops,
			args:      common,
			obj:       objs[0],
		},
		{
			desc:     "not found",
			idPair:   strconv.Itoa(100000),
			idDenom1: strconv.Itoa(100000),
			idDenom2: strconv.Itoa(100000),
			idLeaders: []*types.Leader{
				{
					Address: strconv.Itoa(100000),
					Drops:   sdk.NewInt(100000),
				},
			},
			idDrops: sdk.NewInt(100000),
			args:    common,
			err:     status.Error(codes.InvalidArgument, "not found"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idPair,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdPool(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetPoolResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.Pool)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.Pool),
				)
			}
		})
	}
}

func TestListPool(t *testing.T) {
	net, objs := networkWithPoolObjects(t, 5)

	ctx := net.Validators[0].ClientCtx
	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListPool(), args)
			require.NoError(t, err)
			var resp types.QueryAllPoolResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Pool), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.Pool),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListPool(), args)
			require.NoError(t, err)
			var resp types.QueryAllPoolResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Pool), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.Pool),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListPool(), args)
		require.NoError(t, err)
		var resp types.QueryAllPoolResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.Pool),
		)
	})
}
