package keeper

import (
	"context"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pendulum-labs/market/x/market/types"
)

func (k msgServer) CreateOrder(goCtx context.Context, msg *types.MsgCreateOrder) (*types.MsgCreateOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	amount, _ := sdk.NewIntFromString(msg.Amount)
	if amount == sdk.NewInt(0) {
		sdkerrors.Wrapf(types.ErrInvalidOrderAmount, "Amount must be greater than zero")
	}

	coinBid := sdk.NewCoin(msg.DenomBid, amount)

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

	rate, err := types.RateStringToInt(msg.Rate)
	if err != nil {
		return nil, err
	}

	prev, _ := strconv.ParseUint(msg.Prev, 10, 64)

	next, _ := strconv.ParseUint(msg.Next, 10, 64)

	// Create the uid
	uid := k.GetUidCount(ctx)

	var order = types.Order{
		Uid:       uid,
		Owner:     msg.Creator,
		Active:    true,
		DenomAsk:  msg.DenomAsk,
		DenomBid:  msg.DenomBid,
		OrderType: msg.OrderType,
		Amount:    amount,
		Rate:      rate,
		Prev:      prev,
		Next:      next,
	}

	// Case 1
	// Only order in book
	if prev == 0 && next == 0 {

		/**********************************************************************
		* THEN Member[AskCoin, BidCoin] stop/limit field must be 0            *
		* Stop / Limit = 0 means that the book is empty                       *
		**********************************************************************/
		if msg.OrderType == "stop" {
			if memberBid.Stop != 0 {
				return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Bid Member stop field not 0")
			}

			// Update MemberBid Stop Head
			memberBid.Stop = uid

			// update memberBid event
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeUpdateMember,
					sdk.NewAttribute(types.AttributeKeyDenomA, memberBid.DenomA),
					sdk.NewAttribute(types.AttributeKeyDenomB, memberBid.DenomB),
					sdk.NewAttribute(types.AttributeKeyStop, strconv.FormatUint(memberBid.Stop, 10)),
				),
			)

		}

		if msg.OrderType == "limit" {
			if memberBid.Limit != 0 {
				return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Bid Member limit field not 0")
			}

			// Update MemberBid Limit Head
			memberBid.Limit = uid

			// update memberBid event
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeUpdateMember,
					sdk.NewAttribute(types.AttributeKeyDenomA, memberBid.DenomA),
					sdk.NewAttribute(types.AttributeKeyDenomB, memberBid.DenomB),
					sdk.NewAttribute(types.AttributeKeyStop, strconv.FormatUint(memberBid.Limit, 10)),
				),
			)

		}

		k.SetMember(ctx, memberBid)

	}

	// Case 2
	// New head of the book
	if order.Prev == 0 && order.Next > 0 {

		nextOrder, _ := k.GetOrder(ctx, next)
		if !nextOrder.Active {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Next order not active")
		}
		if nextOrder.Prev != 0 {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Next order not currently head of book")
		}

		if msg.OrderType == "stop" {

			if types.LTE(order.Rate, nextOrder.Rate) {
				return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Order rate less than or equal Next")
			}

			// Set order as new head of MemberBid Stop
			memberBid.Stop = uid

		}

		if msg.OrderType == "limit" {

			if types.GTE(order.Rate, nextOrder.Rate) {
				return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Order rate greater than or equal Next")
			}

			// Set order as new head of MemberBid Limit
			memberBid.Limit = uid

		}

		// Set nextOrder prev field to order
		nextOrder.Prev = uid

		// Update Next Order
		k.SetOrder(ctx, nextOrder)

		// Update Member Bid
		k.SetMember(ctx, memberBid)
	}

	// Case 3
	// New tail of book
	if order.Prev > 0 && order.Next == 0 {

		prevOrder, _ := k.GetOrder(ctx, prev)

		if !prevOrder.Active {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Prev order not active")
		}
		if prevOrder.Next != 0 {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Prev order not currently tail of book")
		}

		if msg.OrderType == "stop" {

			if types.GT(order.Rate, prevOrder.Rate) {
				return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Order rate greater than Prev")
			}

		}

		if msg.OrderType == "limit" {

			if types.LT(order.Rate, prevOrder.Rate) {
				return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Order rate less than Prev")
			}

		}

		// Set nextOrder Next field to Order
		prevOrder.Next = uid

		// Update Previous Order
		k.SetOrder(ctx, prevOrder)
	}

	// Case 4
	// IF next position and prev position are stated
	if order.Prev > 0 && order.Next > 0 {
		prevOrder, _ := k.GetOrder(ctx, prev)
		nextOrder, _ := k.GetOrder(ctx, next)

		if !prevOrder.Active {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Prev order not active")
		}
		if !nextOrder.Active {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Next order not active")
		}

		if !(nextOrder.Prev == prevOrder.Uid && prevOrder.Next == nextOrder.Uid) {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Prev and Next are not adjacent")
		}

		if msg.OrderType == "stop" {

			if types.GT(order.Rate, prevOrder.Rate) {
				return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Order rate greater than Prev")
			}

			if types.LTE(order.Rate, nextOrder.Rate) {
				return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Order rate less than or equal to Next")
			}

			prevOrder.Next = uid
			nextOrder.Prev = uid

			// Update Previous and Next Orders
			k.SetOrder(ctx, prevOrder)
			k.SetOrder(ctx, nextOrder)
		}

		if msg.OrderType == "limit" {

			if types.LT(order.Rate, prevOrder.Rate) {
				return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Order rate less than Prev")
			}

			if types.GTE(order.Rate, nextOrder.Rate) {
				return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Order rate greater than or equal to Next")
			}

			prevOrder.Next = uid
			nextOrder.Prev = uid

			// Update Previous and Next Orders
			k.SetOrder(ctx, prevOrder)
			k.SetOrder(ctx, nextOrder)
		}
	}

	// Transfer order amount to module
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, coinsBid)
	if sdkError != nil {
		return nil, sdkError
	}

	// Increment UID Counter
	k.SetUidCount(ctx, uid+1)

	k.SetOrder(ctx, order)

	// create order event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreateOrder,
			sdk.NewAttribute(types.AttributeKeyUid, strconv.FormatUint(order.Uid, 10)),
			sdk.NewAttribute(types.AttributeKeyDenomA, memberBid.DenomA),
			sdk.NewAttribute(types.AttributeKeyDenomB, memberBid.DenomB),
			sdk.NewAttribute(types.AttributeKeyOrderType, order.OrderType),
			sdk.NewAttribute(types.AttributeKeyAmount, msg.Amount),
			sdk.NewAttribute(types.AttributeKeyRate, strings.Join(msg.Rate, ",")),
			sdk.NewAttribute(types.AttributeKeyPrev, msg.Prev),
			sdk.NewAttribute(types.AttributeKeyNext, msg.Next),
		),
	)

	if msg.OrderType == "stop" {
		// Execute Ask Limit first which will check stops
		// if there are no Ask Limits enabled.  This is a safe
		// guard in the case there is a stop run.
		// Stop run would potentially take place if
		// stop book is checked first repeatedly during price fall
		_, error := ExecuteLimit(k, ctx, msg.DenomBid, msg.DenomAsk, memberBid, memberAsk)
		if error != nil {
			return nil, error
		}
	} else if msg.OrderType == "limit" {
		_, error := ExecuteLimit(k, ctx, msg.DenomAsk, msg.DenomBid, memberAsk, memberBid)
		if error != nil {
			return nil, error
		}
	}

	return &types.MsgCreateOrderResponse{}, nil
}

func ExecuteLimit(k msgServer, ctx sdk.Context, denomAsk string, denomBid string, memberAsk types.Member, memberBid types.Member) (bool, error) {
	// IF Limit Head is equal to 0 THEN the Limit Book is EMPTY
	if memberBid.Limit == 0 {
		_, sdkError := ExecuteStop(k, ctx, denomBid, denomAsk, memberBid, memberAsk)
		if sdkError != nil {
			return false, sdkError
		}
		return true, nil
	}

	limitHead, _ := k.GetOrder(ctx, memberBid.Limit)

	if types.LTE([]sdk.Int{memberAsk.Balance, memberBid.Balance}, limitHead.Rate) {
		_, sdkError := ExecuteStop(k, ctx, denomBid, denomAsk, memberBid, memberAsk)
		if sdkError != nil {
			return false, sdkError
		}
		return true, nil
	}

	// Execute Head Limit
	// The AMM Balance of the Bid Coin corresponding to Limit Order Exchange Rate
	maxMemberBidBalance :=
		limitHead.Rate[1].Mul(memberAsk.Balance.Add(memberBid.Balance)).Quo(limitHead.Rate[1].Add(limitHead.Rate[0]))

	// Maximum amountBid of the Bid Coin that the AMM may accept at Limit Order Exchange Rate
	maxAmountBid := maxMemberBidBalance.Sub(memberBid.Balance)

	// Strike Bid Amount: The amountBid of the bid coin exchanged
	var strikeAmountBid sdk.Int

	// Strike Bid Amount given by the user exchange account and received by the
	// Pair AMM Pool B Member is the lesser of maxPoolBid or limit amountBid
	if limitHead.Amount.LTE(maxAmountBid) {
		strikeAmountBid = limitHead.Amount
		limitHead.Active = false
		memberBid.Limit = limitHead.Next
		if limitHead.Next != 0 {
			limitNext, _ := k.GetOrder(ctx, limitHead.Next)
			limitNext.Prev = 0
			k.SetOrder(ctx, limitNext)
		}
	} else {
		strikeAmountBid = maxAmountBid
	}

	// Need limits on Rates.

	// StrikeAmountAsk = StrikeAmountBid * ExchangeRate(A/B)
	// Exchange Rate is held constant at initial AMM balances
	strikeAmountAsk := (strikeAmountBid.Mul(limitHead.Rate[0])).Quo(limitHead.Rate[1])

	// Edge case where strikeAskAmount rounds to 0
	// Rounding favors AMM vs Order
	if strikeAmountAsk.Equal(sdk.NewInt(0)) {
		return false, nil
	}

	// moduleAcc := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	// Get the borrower address
	owner, _ := sdk.AccAddressFromBech32(limitHead.Owner)

	coinAsk := sdk.NewCoin(denomAsk, strikeAmountAsk)
	coinsAsk := sdk.NewCoins(coinAsk)

	// Transfer ask order amount to owner account
	sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, owner, coinsAsk)
	if sdkError != nil {
		return false, sdkError
	}

	limitHead.Amount = limitHead.Amount.Sub(strikeAmountBid)

	k.SetOrder(ctx, limitHead)

	memberBid.Previous = memberBid.Balance
	memberAsk.Previous = memberAsk.Balance

	memberBid.Balance = memberBid.Balance.Add(strikeAmountBid)
	memberAsk.Balance = memberAsk.Balance.Sub(strikeAmountAsk)

	k.SetMember(ctx, memberAsk)
	k.SetMember(ctx, memberBid)
	return true, nil
}

func ExecuteStop(k msgServer, ctx sdk.Context, denomAsk string, denomBid string, memberAsk types.Member, memberBid types.Member) (bool, error) {
	// Checking for existence of stop order at the memberBid head
	if memberBid.Stop == 0 {
		return true, nil
	}

	stopHead, _ := k.GetOrder(ctx, memberBid.Stop)

	if types.GTE([]sdk.Int{memberAsk.Balance, memberBid.Balance}, stopHead.Rate) {
		return true, nil
	}

	// Execute Stop Order

	maxMemberBidBal := memberAsk.Balance.Add(memberBid.Balance).Sub(memberAsk.Balance.Quo(sdk.NewInt(2)))
	maxMemberBidAmount := maxMemberBidBal.Sub(memberBid.Balance)

	// Strike Bid Amount: The amountBid of the bid coin exchanged
	strikeAmountBid := stopHead.Amount
	var strikeAmountAsk sdk.Int

	if stopHead.Amount.GT(maxMemberBidAmount) {
		strikeAmountBid = maxMemberBidAmount
		strikeAmountAsk = (strikeAmountBid.Mul(memberAsk.Balance.Sub(strikeAmountBid))).Quo(memberBid.Balance.Add(strikeAmountBid))

		// Edge case where strikeAskAmount rounds to 0
		// Rounding favors AMM vs Order
		if strikeAmountAsk.Equal(sdk.NewInt(0)) {
			return false, nil
		}

	} else {
		strikeAmountAsk = strikeAmountBid.Mul(memberAsk.Balance.Sub(strikeAmountBid)).Quo(memberBid.Balance.Add(strikeAmountBid))

		// Edge case where strikeAskAmount rounds to 0
		// Rounding favors AMM vs Order
		if strikeAmountAsk.Equal(sdk.NewInt(0)) {
			return false, nil
		}

		// THEN set Head(Stop) active to false as entire order will be filled
		stopHead.Active = false
		// Set Next Position as Head of Stop Book
		memberBid.Stop = stopHead.Next

		if stopHead.Next != 0 {
			stopNext, _ := k.GetOrder(ctx, stopHead.Next)
			stopNext.Prev = 0
			k.SetOrder(ctx, stopNext)
		}

		// At this point the Head(Stop) position has been deactivated and the Next
		// Stop position has been set as the Head Stop
	}

	// moduleAcc := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	// Get the borrower address
	owner, _ := sdk.AccAddressFromBech32(stopHead.Owner)

	coinAsk := sdk.NewCoin(denomAsk, strikeAmountAsk)
	coinsAsk := sdk.NewCoins(coinAsk)

	// Transfer ask order amount to owner account
	sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, owner, coinsAsk)
	if sdkError != nil {
		return false, sdkError
	}

	stopHead.Amount = stopHead.Amount.Sub(strikeAmountBid)

	k.SetOrder(ctx, stopHead)

	memberBid.Previous = memberBid.Balance
	memberAsk.Previous = memberAsk.Balance

	memberBid.Balance = memberBid.Balance.Add(strikeAmountBid)
	memberAsk.Balance = memberAsk.Balance.Sub(strikeAmountAsk)

	if sdkError != nil {
		return false, sdkError
	}

	k.SetMember(ctx, memberAsk)
	k.SetMember(ctx, memberBid)
	return true, nil
}
