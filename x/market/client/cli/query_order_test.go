package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
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

func networkWithOrderObjects(t *testing.T, n int) (*network.Network, []types.Order) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		order := types.Order{
			Uid:       uint64(i),
			Owner:     strconv.Itoa(i),
			Active:    false,
			OrderType: strconv.Itoa(i),
			DenomAsk:  strconv.Itoa(i),
			DenomBid:  strconv.Itoa(i),
			Amount:    sdk.NewInt(int64(i)),
			Rate:      []sdk.Int{sdk.NewInt(int64(i)), sdk.NewInt(int64(i))},
		}
		nullify.Fill(&order)
		state.OrderList = append(state.OrderList, order)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.OrderList
}

func TestShowOrder(t *testing.T) {
	net, objs := networkWithOrderObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc        string
		idUid       uint64
		idOwner     string
		idActive    bool
		idOrderType string
		idDenomAsk  string
		idDenomBid  string
		idAmount    sdk.Int
		idRate      []sdk.Int

		args []string
		err  error
		obj  types.Order
	}{
		{
			desc:        "found",
			idUid:       objs[0].Uid,
			idOwner:     objs[0].Owner,
			idActive:    objs[0].Active,
			idOrderType: objs[0].OrderType,
			idDenomAsk:  objs[0].DenomAsk,
			idDenomBid:  objs[0].DenomBid,
			idAmount:    objs[0].Amount,
			idRate:      objs[0].Rate,

			args: common,
			obj:  objs[0],
		},
		{
			desc:        "not found",
			idUid:       100000,
			idOwner:     strconv.Itoa(100000),
			idActive:    false,
			idOrderType: strconv.Itoa(100000),
			idDenomAsk:  strconv.Itoa(100000),
			idDenomBid:  strconv.Itoa(100000),
			idAmount:    sdk.NewInt(int64(100000)),
			idRate:      []sdk.Int{sdk.NewInt(int64(100000)), sdk.NewInt(int64(100000))},

			args: common,
			err:  status.Error(codes.InvalidArgument, "not found"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				strconv.Itoa(int(tc.idUid)),
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowOrder(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetOrderResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.Order)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.Order),
				)
			}
		})
	}
}

func TestListOrder(t *testing.T) {
	net, objs := networkWithOrderObjects(t, 5)

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
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListOrder(), args)
			require.NoError(t, err)
			var resp types.QueryAllOrderResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Order), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.Order),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListOrder(), args)
			require.NoError(t, err)
			var resp types.QueryAllOrderResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Order), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.Order),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListOrder(), args)
		require.NoError(t, err)
		var resp types.QueryAllOrderResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.Order),
		)
	})
}
