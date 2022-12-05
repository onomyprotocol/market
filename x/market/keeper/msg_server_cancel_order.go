package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/onomyprotocol/market/x/market/types"
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

	if order.Prev == 0 {
		memberBid, found := k.GetMember(ctx, order.DenomAsk, order.DenomBid)
		if !found {
			return nil, sdkerrors.Wrapf(types.ErrMemberNotFound, "%s", order.DenomBid)
		}

		if memberBid.Stop != order.Uid {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "%s", order.Uid)
		}

		if order.Next == 0 {
			memberBid.Stop = 0
		}

		if order.Next > 0 {
			memberBid.Stop = order.Next
		}
	}

	if order.Next == 0 {
		stops[coinAsk][coinBid][position.prev].next = 0
	}

	// TODO: I think this needs to be two conditionals?
	if order.Prev > 0 || order.Next > 0 {
		stops[coinAsk][coinBid][position.next].prev = position.prev
		stops[coinAsk][coinBid][position.prev].next = position.next
	}

	_ = ctx

	return &types.MsgCancelOrderResponse{}, nil
}
