package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pendulum-labs/market/x/market/types"
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

	// A(i)*B(i) = A(f)*B(f)
	// A(f) = A(i)*B(i)/B(f)
	// amountAsk = A(i) - A(f) = A(i) - A(i)*B(i)/B(f)
	amountAsk := memberAsk.Balance.Sub((memberAsk.Balance.Mul(memberBid.Balance)).Quo(memberBid.Balance.Add(amountBid)))

	// Market Order Fee
	fee, _ := sdk.NewIntFromString(k.getParams(ctx).MarketFee)
	amountAsk = amountAsk.Sub((amountAsk.Mul(fee)).Quo(sdk.NewInt(10000)))

	// Edge case where strikeAskAmount rounds to 0
	// Rounding favors AMM vs Order
	if amountAsk.Equal(sdk.ZeroInt()) {
		return nil, sdkerrors.Wrapf(types.ErrAmtZero, "amount ask equal to zero")
	}

	// Slippage includes Market Order Fee
	// (A1*B2 - A2*B1)/(A1*B2)
	slippage :=
		((memberAsk.Balance.Mul(amountBid).Sub(amountAsk.Mul(memberBid.Balance))).Mul(sdk.NewInt(10000))).Quo(memberAsk.Balance.Mul(amountBid))
	slipLimit, _ := sdk.NewIntFromString(msg.Slippage)
	if slippage.GT(slipLimit) {
		return nil, sdkerrors.Wrapf(types.ErrSlippageTooGreat, "Slippage %s", slippage)
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
		EndTime:   ctx.BlockHeader().Time.Unix(),
	}

	pool.History = uid

	k.SetPool(ctx, pool)
	k.SetUidCount(ctx, uid+1)
	k.SetOrder(ctx, order)

	return &types.MsgMarketOrderResponse{}, nil

}
