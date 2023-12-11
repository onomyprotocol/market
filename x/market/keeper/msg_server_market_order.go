package keeper

import (
	"context"

	"market/x/market/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) MarketOrder(goCtx context.Context, msg *types.MsgMarketOrder) (*types.MsgMarketOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	amountBid, _ := sdk.NewIntFromString(msg.AmountBid)

	coinBid := sdk.NewCoin(msg.DenomBid, amountBid)

	coinsBid := sdk.NewCoins(coinBid)

	trader, _ := sdk.AccAddressFromBech32(msg.Creator)

	// Check if order creator has available balance
	if err := k.validateSenderBalance(ctx, trader, coinsBid); err != nil {
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

	productBeg := memberAsk.Balance.Mul(memberBid.Balance)

	// A(i)*B(i) = A(f)*B(f)
	// A(f) = A(i)*B(i)/B(f)
	// amountAsk = A(i) - A(f) = A(i) - A(i)*B(i)/B(f)
	// Compensate for rounding: strikeAmountAsk = A(i) - A(f) = A(i) - [A(i)*B(i)/B(f)+1]
	amountAsk := memberAsk.Balance.Sub(((memberAsk.Balance.Mul(memberBid.Balance)).Quo(memberBid.Balance.Add(amountBid))).Add(sdk.NewInt(1)))

	// Market Order Fee
	fee, _ := sdk.NewIntFromString(k.getParams(ctx).MarketFee)
	amountAsk = amountAsk.Sub((amountAsk.Mul(fee)).Quo(sdk.NewInt(10000)))

	// Edge case where strikeAskAmount rounds to 0
	// Rounding favors AMM vs Order
	if amountAsk.LTE(sdk.ZeroInt()) {
		return nil, sdkerrors.Wrapf(types.ErrAmtZero, "amount ask equal or less than zero")
	}

	// Slippage is initialized at zero
	slippage := sdk.ZeroInt()

	amountAskExpected, _ := sdk.NewIntFromString(msg.AmountAsk)

	// Slippage is only updated if amount expected is greater than received
	if amountAskExpected.GT(amountAsk) {
		slippage = ((amountAskExpected.Sub(amountAsk)).Mul(sdk.NewInt(10000))).Quo(amountAskExpected)

		slipLimit, _ := sdk.NewIntFromString(msg.Slippage)

		if slippage.GT(slipLimit) {
			return nil, sdkerrors.Wrapf(types.ErrSlippageTooGreat, "Slippage %s", slippage)
		}
	}

	// Transfer bid amount from trader account to module
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, trader, types.ModuleName, coinsBid)
	if sdkError != nil {
		return nil, sdkError
	}

	coinAsk := sdk.NewCoin(msg.DenomAsk, amountAsk)
	coinsAsk := sdk.NewCoins(coinAsk)

	// Transfer ask amount from module to trader account
	sdkError = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, trader, coinsAsk)
	if sdkError != nil {
		return nil, sdkError
	}

	memberAsk.Balance = memberAsk.Balance.Sub(amountAsk)
	memberBid.Balance = memberBid.Balance.Add(amountBid)

	k.SetMember(ctx, memberAsk)
	k.SetMember(ctx, memberBid)

	uid := k.GetUidCount(ctx)

	pool, _ := k.GetPool(ctx, memberBid.Pair)
	prevOrder, _ := k.GetOrder(ctx, pool.History)

	prevOrder.Prev = uid

	var order = types.Order{
		Uid:       uid,
		Owner:     msg.Creator,
		Status:    "filled",
		DenomAsk:  msg.DenomAsk,
		DenomBid:  msg.DenomBid,
		OrderType: "market",
		Amount:    amountBid,
		Rate:      []sdk.Int{amountAsk, amountBid},
		Prev:      0,
		Next:      pool.History,
		BegTime:   ctx.BlockHeader().Time.Unix(),
		UpdTime:   ctx.BlockHeader().Time.Unix(),
	}

	pool.History = uid

	if pool.Denom1 == msg.DenomBid {
		pool.Volume1.Amount = pool.Volume1.Amount.Add(amountBid)
		pool.Volume2.Amount = pool.Volume2.Amount.Add(amountAsk)
	} else {
		pool.Volume1.Amount = pool.Volume1.Amount.Add(amountAsk)
		pool.Volume2.Amount = pool.Volume2.Amount.Add(amountBid)
	}

	k.IncVolume(ctx, msg.DenomBid, amountBid)
	k.IncVolume(ctx, msg.DenomAsk, amountAsk)

	k.SetPool(ctx, pool)
	k.SetUidCount(ctx, uid+1)
	k.SetOrder(ctx, order)

	memberBid, memberAsk, error := ExecuteLimit(k, ctx, coinBid.Denom, coinAsk.Denom, memberBid, memberAsk)
	if error != nil {
		return nil, error
	}
	memberBid, memberAsk, error = ExecuteLimit(k, ctx, coinAsk.Denom, coinBid.Denom, memberAsk, memberBid)
	if error != nil {
		return nil, error
	}

	if memberAsk.Balance.Mul(memberBid.Balance).Equal(productBeg) {
		return &types.MsgMarketOrderResponse{AmountBid: msg.AmountBid, AmountAsk: amountAsk.String(), Slippage: slippage.String()}, nil
	}

	if memberAsk.Balance.Mul(memberBid.Balance).LT(productBeg) {
		return nil, sdkerrors.Wrapf(types.ErrProductInvalid, "Pool product lower after Trade %s", memberAsk.Pair)
	}

	profitAsk, profitBid := k.Profit(productBeg, memberAsk, memberBid)

	memberAsk, error = k.Payout(ctx, profitAsk, memberAsk, pool)
	if error != nil {
		return nil, error
	}

	memberBid, error = k.Payout(ctx, profitBid, memberBid, pool)
	if error != nil {
		return nil, error
	}

	if memberAsk.Balance.Mul(memberBid.Balance).LT(productBeg) {
		return nil, sdkerrors.Wrapf(types.ErrProductInvalid, "Pool product lower after Payout %s", memberAsk.Pair)
	}

	memberAsk, error = k.Burn(ctx, profitAsk, memberAsk)
	if error != nil {
		return nil, error
	}

	memberBid, error = k.Burn(ctx, profitBid, memberBid)
	if error != nil {
		return nil, error
	}

	if memberAsk.Balance.Mul(memberBid.Balance).LT(productBeg) {
		return nil, sdkerrors.Wrapf(types.ErrProductInvalid, "Pool product lower after Burn %s", memberAsk.Pair)
	}

	k.SetMember(ctx, memberAsk)
	k.SetMember(ctx, memberBid)

	return &types.MsgMarketOrderResponse{AmountBid: msg.AmountBid, AmountAsk: amountAsk.String(), Slippage: slippage.String()}, nil
}
