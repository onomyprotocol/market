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

func TestOrderQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNOrder(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetOrderRequest
		response *types.QueryGetOrderResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetOrderRequest{
				Uid:       msgs[0].Uid,
				Owner:     msgs[0].Owner,
				Active:    msgs[0].Active,
				OrderType: msgs[0].OrderType,
				DenomAsk:  msgs[0].DenomAsk,
				DenomBid:  msgs[0].DenomBid,
			},
			response: &types.QueryGetOrderResponse{Order: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetOrderRequest{
				Uid:       msgs[1].Uid,
				Owner:     msgs[1].Owner,
				Active:    msgs[1].Active,
				OrderType: msgs[1].OrderType,
				DenomAsk:  msgs[1].DenomAsk,
				DenomBid:  msgs[1].DenomBid,
			},
			response: &types.QueryGetOrderResponse{Order: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetOrderRequest{
				Uid:       100000,
				Owner:     strconv.Itoa(100000),
				Active:    false,
				OrderType: strconv.Itoa(100000),
				DenomAsk:  strconv.Itoa(100000),
				DenomBid:  strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Order(wctx, tc.request)
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

func TestOrderQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNOrder(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllOrderRequest {
		return &types.QueryAllOrderRequest{
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
			resp, err := keeper.OrderAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Order), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Order),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.OrderAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Order), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Order),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.OrderAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Order),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.OrderAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}