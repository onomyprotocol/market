package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/pendulum-labs/market/x/market/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Bookends(goCtx context.Context, req *types.QueryBookendsRequest) (*types.QueryBookendsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var rateUint64 [2]uint64
	var err error

	// Rate[0] needs to fit into uint64 to avoid numerical errors
	rateUint64[0], err = strconv.ParseUint(req.Rate[0], 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid rate")
	}

	// Rate[1] needs to fit into uint64 to avoid numerical errors
	rateUint64[1], err = strconv.ParseUint(req.Rate[1], 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid rate")
	}

	var rate []sdk.Int

	rate[0] = sdk.NewIntFromUint64(rateUint64[0])
	rate[1] = sdk.NewIntFromUint64(rateUint64[1])

	ends := k.GetBookEnds(ctx, req.GetCoinA(), req.GetCoinB(), req.GetOrderType(), rate)

	// TODO: Process the query
	_ = ctx

	return &types.QueryBookendsResponse{CoinA: req.CoinA, CoinB: req.CoinB, OrderType: req.OrderType, Rate: req.Rate, Prev: ends[0], Next: ends[1]}, nil
}
