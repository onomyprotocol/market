package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/pendulum-labs/market/testutil/keeper"
	"github.com/pendulum-labs/market/testutil/nullify"
	"github.com/pendulum-labs/market/x/market/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestDropQuerySingle(t *testing.T) {
	keeper := keepertest.CreateTestEnvironment(t)
	wctx := sdk.WrapSDKContext(keeper.Context)
	msgs := createNDrop(keeper.MarketKeeper, keeper.Context, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryDropRequest
		response *types.QueryDropResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryDropRequest{
				Uid: msgs[0].Uid,
			},
			response: &types.QueryDropResponse{Drop: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryDropRequest{
				Uid: msgs[1].Uid,
			},
			response: &types.QueryDropResponse{Drop: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryDropRequest{
				Uid: 100000,
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.MarketKeeper.Drop(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestDropQueryPaginated(t *testing.T) {
	keeper := keepertest.CreateTestEnvironment(t)
	wctx := sdk.WrapSDKContext(keeper.Context)
	msgs := createNDrop(keeper.MarketKeeper, keeper.Context, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllDropRequest {
		return &types.QueryAllDropRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.MarketKeeper.DropAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Drops), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Drops),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.MarketKeeper.DropAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Drops), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Drops),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.MarketKeeper.DropAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Drops),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.MarketKeeper.DropAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
