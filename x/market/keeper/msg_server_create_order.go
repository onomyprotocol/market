package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/onomyprotocol/market/x/market/types"
)

func (k msgServer) CreateOrder(goCtx context.Context, msg *types.MsgCreateOrder) (*types.MsgCreateOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	/*
			    // TODO: validate orderType is 0 or 1
		        require(coinAsk != coinBid, "Bid and ask coin cannot be the same");
		        require(position.owner == msg.sender, "Position not owned by sender");
		        require(position.amountBid > 0, "Amount of bid must be greater than zero");
	*/

	amount, _ := sdk.NewIntFromString(msg.Amount)

	coinBid := sdk.NewCoin(msg.DenomBid, amount)

	coinsBid := sdk.NewCoins(coinBid)

	// moduleAcc := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	// Get the borrower address
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)

	if err := k.validateSenderBalance(ctx, creator, coinsBid); err != nil {
		return nil, err
	}

	memberAsk, found := k.GetMember(ctx, msg.DenomBid, msg.DenomAsk)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrMemberNotFound, "Member %s", msg.DenomAsk)
	}

	memberBid, found := k.GetMember(ctx, msg.DenomAsk, msg.DenomBid)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrMemberNotFound, "Member %s", msg.DenomBid)
	}

	numerator, _ := sdk.NewIntFromString(msg.Rate[0])

	denominator, _ := sdk.NewIntFromString(msg.Rate[1])

	rate := []sdk.Int{numerator, denominator}

	prev, _ := strconv.ParseUint(msg.Prev, 10, 64)

	next, _ := strconv.ParseUint(msg.Next, 10, 64)

	if prev == 0 && next == 0 {

		/**********************************************************************
		* THEN Member[AskCoin, BidCoin] stop/limit field must be 0            *
		* Stop / Limit = 0 means that the book is empty                       *
		**********************************************************************/
		if msg.OrderType == "stop" {
			if memberBid.Stop != 0 {
				return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Bid Member stop field not 0")
			}

			// Create the uid
			uid := k.GetUidCount(ctx)

			memberBid.Stop = uid

			// Update drop uid count
			k.SetUidCount(ctx, uid+1)

			var order = types.Order{
				Owner:     msg.Creator,
				Active:    true,
				DenomAsk:  msg.DenomAsk,
				DenomBid:  msg.DenomBid,
				OrderType: "stop",
				Amount:    amount,
				Rate:      rate,
				Prev:      0,
				Next:      0,
			}

			_ = order
		}
	}

	_ = memberAsk
	_ = memberBid

	_ = ctx

	return &types.MsgCreateOrderResponse{}, nil
}
