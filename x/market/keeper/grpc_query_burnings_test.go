package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/onomyprotocol/market/testutil/keeper"
	"github.com/onomyprotocol/market/testutil/nullify"
	"github.com/onomyprotocol/market/x/market/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestBurningsQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNBurnings(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetBurningsRequest
		response *types.QueryGetBurningsResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetBurningsRequest{
				Denom: msgs[0].Denom,
			},
			response: &types.QueryGetBurningsResponse{Burnings: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetBurningsRequest{
				Denom: msgs[1].Denom,
			},
			response: &types.QueryGetBurningsResponse{Burnings: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetBurningsRequest{
				Denom: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Burnings(wctx, tc.request)
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

func TestBurningsQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNBurnings(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllBurningsRequest {
		return &types.QueryAllBurningsRequest{
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
			resp, err := keeper.BurningsAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Burnings), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Burnings),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.BurningsAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Burnings), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Burnings),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.BurningsAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Burnings),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.BurningsAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}