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

	"github.com/pendulum-labs/market/testutil/network"
	"github.com/pendulum-labs/market/testutil/nullify"
	"github.com/pendulum-labs/market/x/market/client/cli"
	"github.com/pendulum-labs/market/x/market/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithMemberObjects(t *testing.T, n int) (*network.Network, []types.Member) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		member := types.Member{
			Pair:     strconv.Itoa(i),
			DenomA:   strconv.Itoa(i),
			DenomB:   strconv.Itoa(i),
			Balance:  sdk.NewIntFromUint64(uint64(i)),
			Previous: sdk.NewIntFromUint64(uint64(i)),
		}
		nullify.Fill(&member)
		state.MemberList = append(state.MemberList, member)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.MemberList
}

func TestShowMember(t *testing.T) {
	net, objs := networkWithMemberObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc     string
		idDenomA string
		idDenomB string

		args []string
		err  error
		obj  types.Member
	}{
		{
			desc:     "found",
			idDenomA: objs[0].DenomA,
			idDenomB: objs[0].DenomB,

			args: common,
			obj:  objs[0],
		},
		{
			desc:     "not found",
			idDenomA: strconv.Itoa(100000),
			idDenomB: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.InvalidArgument, "not found"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idDenomA,
				tc.idDenomB,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowMember(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetMemberResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.Member)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.Member),
				)
			}
		})
	}
}

func TestListMember(t *testing.T) {
	net, objs := networkWithMemberObjects(t, 5)

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
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListMember(), args)
			require.NoError(t, err)
			var resp types.QueryAllMemberResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Member), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.Member),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListMember(), args)
			require.NoError(t, err)
			var resp types.QueryAllMemberResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Member), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.Member),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListMember(), args)
		require.NoError(t, err)
		var resp types.QueryAllMemberResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.Member),
		)
	})
}
