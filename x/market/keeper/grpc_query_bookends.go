package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/pendulum-labs/market/x/market/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Bookends(goCtx context.Context, req *types.QueryBookendsRequest) (*types.QueryBookendsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	rate, err := types.RateStringToInt(req.Rate)
	if err != nil {
		return nil, err
	}

	ends := k.BookEnds(ctx, req.GetCoinA(), req.GetCoinB(), req.GetOrderType(), rate)

	// TODO: Process the query
	_ = ctx

	return &types.QueryBookendsResponse{CoinA: req.CoinA, CoinB: req.CoinB, OrderType: req.OrderType, Rate: req.Rate, Prev: ends[0], Next: ends[1]}, nil
}
