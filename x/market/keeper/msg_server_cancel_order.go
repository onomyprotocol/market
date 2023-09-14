package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pendulum-labs/market/x/market/types"
)

func (k msgServer) CancelOrder(goCtx context.Context, msg *types.MsgCancelOrder) (*types.MsgCancelOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	uid, _ := strconv.ParseUint(msg.Uid, 10, 64)

	order, found := k.GetOrder(ctx, uid)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrOrderNotFound, "%s", msg.Uid)
	}

	if order.Owner != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotOrderOwner, "%s", msg.Uid)
	}

	memberBid, found := k.GetMember(ctx, order.DenomAsk, order.DenomBid)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrMemberNotFound, "%s", order.DenomBid)
	}

	if order.Prev == 0 {
		if memberBid.Stop != order.Uid && memberBid.Limit != order.Uid {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "%c", order.Uid)
		}

		if order.Next == 0 {
			if order.OrderType == "stop" {
				memberBid.Stop = 0
			}

			if order.OrderType == "limit" {
				memberBid.Limit = 0
			}

			k.SetMember(ctx, memberBid)
		} else {
			nextOrder, found := k.GetOrder(ctx, order.Next)
			if !found {
				return nil, sdkerrors.Wrapf(types.ErrOrderNotFound, "%c", order.Next)
			}

			nextOrder.Prev = 0

			if order.OrderType == "stop" {
				memberBid.Stop = order.Next
			}

			if order.OrderType == "limit" {
				memberBid.Limit = order.Next
			}

			k.SetMember(ctx, memberBid)
			k.SetOrder(ctx, nextOrder)
		}
	} else {
		prevOrder, found := k.GetOrder(ctx, order.Prev)
		if !found {
			return nil, sdkerrors.Wrapf(types.ErrOrderNotFound, "%c", order.Prev)
		}

		if order.Next == 0 {
			prevOrder.Next = 0
			k.SetOrder(ctx, prevOrder)
		} else {
			nextOrder, found := k.GetOrder(ctx, order.Next)
			if !found {
				return nil, sdkerrors.Wrapf(types.ErrOrderNotFound, "%c", order.Next)
			}

			nextOrder.Prev = order.Prev
			prevOrder.Next = order.Next

			k.SetOrder(ctx, prevOrder)
			k.SetOrder(ctx, nextOrder)
		}
	}

	order.Status = "canceled"
	order.EndTime = ctx.BlockHeader().Time.Unix()
	k.RemoveOrderOwner(ctx, order.Owner, order.Uid)
	k.SetOrder(ctx, order)

	return &types.MsgCancelOrderResponse{}, nil
}
