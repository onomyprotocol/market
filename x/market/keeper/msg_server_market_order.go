package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/onomyprotocol/market/x/market/types"
)

func (k msgServer) MarketOrder(goCtx context.Context, msg *types.MsgMarketOrder) (*types.MsgMarketOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	/*
	   // TODO: validate orderType is 0 or 1
	   require(coinAsk != coinBid, "Bid and ask coin cannot be the same");
	   require(position.owner == msg.sender, "Position not owned by sender");
	   require(position.amountBid > 0, "Amount of bid must be greater than zero");
	*/

	amountBid, _ := sdk.NewIntFromString(msg.AmountBid)

	coinBid := sdk.NewCoin(msg.DenomBid, amountBid)

	coinsBid := sdk.NewCoins(coinBid)

	// moduleAcc := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	// Get the borrower address
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)

	// Check if order creator has available balance
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

	maxMemberBidBal := memberAsk.Balance.Add(memberBid.Balance).Sub(memberAsk.Balance.Quo(sdk.NewInt(2)))
	maxMemberBidAmount := maxMemberBidBal.Sub(memberBid.Balance)

	if amountBid.GT(maxMemberBidAmount) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidOrderAmount, "Max bid amount %s", maxMemberBidAmount.String())
	}

	// Summation Invariant
	// A(i) + B(i) = A(f) + B(f)

	// Derivation
	// A(f) = A(i) + B(i) - B(f)
	// A(f) = A(i) - amountBid
	// Exch(f) = A(f) / B(f)
	// Exch(f) = (A(i) - amountBid) / B(f)
	// B(f) = B(i) + amountBid
	// Exch(f) =  (A(i) - amountBid) / (B(i) + amountBid)
	// amountAsk = amountBid * Exch(f) = [amountBid * (A(i) - amountBid)] / (B(i) + amountBid)
	amountAsk := (amountBid.Mul(memberAsk.Balance.Sub(amountBid))).Quo(memberBid.Balance.Add(amountBid))

	quoteAsk, _ := sdk.NewIntFromString(msg.QuoteAsk)

	// If quote of ask coin is greater than strike ask amount then check slippage
	// Market order without slippage has quoteAsk set to zero
	if quoteAsk.GT(amountAsk) {
		strikeSlippage := ((quoteAsk.Sub(amountAsk)).Mul(sdk.NewInt(10000))).Quo(quoteAsk)
		slippage, _ := sdk.NewIntFromString(msg.Slippage)
		if strikeSlippage.GT(slippage) {
			return nil, sdkerrors.Wrapf(types.ErrSlippageTooGreat, "Slippage %s", strikeSlippage)
		}
	}
	return &types.MsgMarketOrderResponse{}, nil

}
