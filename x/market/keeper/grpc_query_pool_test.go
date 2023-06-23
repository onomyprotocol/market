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

func TestPoolQuerySingle(t *testing.T) {
	keeper := keepertest.CreateTestEnvironment(t)
	wctx := sdk.WrapSDKContext(keeper.Context)
	msgs := createNPool(keeper.MarketKeeper, keeper.Context, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetPoolRequest
		response *types.QueryGetPoolResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetPoolRequest{
				Pair:   msgs[0].Pair,
				Denom1: msgs[0].Denom1,
				Denom2: msgs[0].Denom2,
				Leader: msgs[0].Leader,
			},
			response: &types.QueryGetPoolResponse{Pool: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetPoolRequest{
				Pair:   msgs[1].Pair,
				Denom1: msgs[1].Denom1,
				Denom2: msgs[1].Denom2,
				Leader: msgs[1].Leader,
			},
			response: &types.QueryGetPoolResponse{Pool: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetPoolRequest{
				Pair:   strconv.Itoa(100000),
				Denom1: strconv.Itoa(100000),
				Denom2: strconv.Itoa(100000),
				Leader: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.MarketKeeper.Pool(wctx, tc.request)
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

func TestPoolQueryPaginated(t *testing.T) {
	keeper := keepertest.CreateTestEnvironment(t)
	wctx := sdk.WrapSDKContext(keeper.Context)
	msgs := createNPool(keeper.MarketKeeper, keeper.Context, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllPoolRequest {
		return &types.QueryAllPoolRequest{
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
			resp, err := keeper.MarketKeeper.PoolAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Pool), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Pool),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.MarketKeeper.PoolAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Pool), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Pool),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.MarketKeeper.PoolAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Pool),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.MarketKeeper.PoolAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
