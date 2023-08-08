package keeper

import (
	"context"
	"math/big"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pendulum-labs/market/x/market/types"
)

func (k msgServer) CreateOrder(goCtx context.Context, msg *types.MsgCreateOrder) (*types.MsgCreateOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	amount, _ := sdk.NewIntFromString(msg.Amount)
	if amount == sdk.ZeroInt() {
		sdkerrors.Wrapf(types.ErrInvalidOrderAmount, "Amount must be greater than zero")
	}

	coinBid := sdk.NewCoin(msg.DenomBid, amount)

	coinsBid := sdk.NewCoins(coinBid)

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
		Status:    "active",
		DenomAsk:  msg.DenomAsk,
		DenomBid:  msg.DenomBid,
		OrderType: msg.OrderType,
		Amount:    amount,
		Rate:      rate,
		Prev:      prev,
		Next:      next,
		BegTime:   ctx.BlockHeader().Time.Unix(),
		EndTime:   0,
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
		if !(nextOrder.Status == "active") {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Next order not active")
		}
		if !(nextOrder.DenomAsk == order.DenomAsk && nextOrder.DenomBid == order.DenomBid) {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Incorrect book")
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

		if !(prevOrder.Status == "active") {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Prev order not active")
		}
		if prevOrder.Next != 0 {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Prev order not currently tail of book")
		}
		if !(prevOrder.DenomAsk == order.DenomAsk && prevOrder.DenomBid == order.DenomBid) {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Incorrect book")
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

		if !(prevOrder.Status == "active") {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Prev order not active")
		}
		if !(prevOrder.DenomAsk == order.DenomAsk && prevOrder.DenomBid == order.DenomBid) {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Incorrect book")
		}

		if !(nextOrder.Status == "active") {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Next order not active")
		}
		if !(nextOrder.DenomAsk == order.DenomAsk && nextOrder.DenomBid == order.DenomBid) {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Incorrect book")
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
	k.SetOrderOwner(ctx, order.Owner, order.Uid)

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

	return &types.MsgCreateOrderResponse{Uid: order.Uid}, nil
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

	// Max Member(Bid) Balance B(f)
	// The AMM Balance of the Bid Coin corresponding to Limit Order Exchange Rate
	// Model: Constant Product
	// A(i): Initial Balance of Ask Coin in AMM Pool
	// B(f): Final Balance of Bid Coin in AMM Pool
	// Exch(f):
	// A(i)*B(i)=A(f)*B(f)
	// Exch(f)=A(f)/B(f) -> B(f)*Exch(f)=A(f)
	// A(i)*B(i)=B(f)*Exch(f)*B(f) -> A(i)*B(i)=Exch(f)*B(f)^2
	// (A(i)*B(i))/Exch(f)=B(f)^2
	// B(f)=SQRT((A(i)*B(i))/Exch(f))
	bigInt := &big.Int{}
	maxMemberBidBalance :=
		sdk.NewIntFromBigInt(
			bigInt.Sqrt(
				sdk.Int.BigInt(
					(limitHead.Rate[1].Mul(memberAsk.Balance.Mul(memberBid.Balance))).Quo(limitHead.Rate[0]),
				),
			),
		)

	// Maximum amountBid of the Bid Coin that the AMM may accept at Limit Order Exchange Rate
	maxAmountBid := maxMemberBidBalance.Sub(memberBid.Balance)

	// Strike Bid Amount: The amountBid of the bid coin exchanged
	var strikeAmountBid sdk.Int

	// Strike Bid Amount given by the user exchange account and received by the
	// Pair AMM Pool B Member is the lesser of maxPoolBid or limit amountBid
	if limitHead.Amount.LTE(maxAmountBid) {
		strikeAmountBid = limitHead.Amount
		memberBid.Limit = limitHead.Next
		if limitHead.Next != 0 {
			limitNext, _ := k.GetOrder(ctx, limitHead.Next)
			limitNext.Prev = 0
			k.SetOrder(ctx, limitNext)
		}
	} else {
		strikeAmountBid = maxAmountBid
	}

	// StrikeAmountAsk = StrikeAmountBid * ExchangeRate(A/B)
	// Exchange Rate is held constant at initial AMM balances
	strikeAmountAsk := (strikeAmountBid.Mul(limitHead.Rate[0])).Quo(limitHead.Rate[1])

	// Edge case where strikeAskAmount rounds to 0
	// Rounding favors AMM vs Order
	if strikeAmountAsk.Equal(sdk.ZeroInt()) {
		return false, nil
	}

	pool, _ := k.GetPool(ctx, memberAsk.Pair)

	if limitHead.Amount.Equal(strikeAmountBid) {
		limitHead.Status = "filled"
		limitHead.Prev = 0
		k.RemoveOrderOwner(ctx, limitHead.Owner, limitHead.Uid)

		if pool.History == 0 {
			limitHead.Next = 0
		} else {
			prevFilledOrder, _ := k.GetOrder(ctx, pool.History)
			prevFilledOrder.Prev = limitHead.Uid
			limitHead.Next = prevFilledOrder.Uid
			k.SetOrder(ctx, prevFilledOrder)
		}

		pool.History = limitHead.Uid

	} else {
		// Add partially filled order to history
		// Keep remainder of order into book
		partialFillOrder := limitHead

		partialFillOrder.Uid = k.GetUidCount(ctx)
		k.SetUidCount(ctx, partialFillOrder.Uid+1)

		partialFillOrder.Amount = strikeAmountBid
		partialFillOrder.Status = "filled"

		limitHead.Amount = limitHead.Amount.Sub(strikeAmountBid)

		if pool.History == 0 {
			pool.History = partialFillOrder.Uid
			partialFillOrder.Prev = 0
			partialFillOrder.Next = 0
		} else {
			prevFilledOrder, _ := k.GetOrder(ctx, pool.History)
			prevFilledOrder.Prev = partialFillOrder.Uid
			k.SetOrder(ctx, prevFilledOrder)
			partialFillOrder.Prev = 0
			partialFillOrder.Next = prevFilledOrder.Uid
		}

		k.SetOrder(ctx, partialFillOrder)
		k.SetOrderOwner(ctx, limitHead.Owner, limitHead.Uid)
	}

	limitHead.EndTime = ctx.BlockHeader().Time.Unix()

	k.SetOrder(ctx, limitHead)
	k.SetPool(ctx, pool)

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

	// Strike Bid Amount: The amountBid of the bid coin exchanged
	strikeAmountBid := stopHead.Amount

	// A(i)*B(i) = A(f)*B(f)
	// A(f) = A(i)*B(i)/B(f)
	// strikeAmountAsk = A(i) - A(f) = A(i) - A(i)*B(i)/B(f)
	strikeAmountAsk := memberAsk.Balance.Sub((memberAsk.Balance.Mul(memberBid.Balance)).Quo(memberBid.Balance.Add(strikeAmountBid)))

	// Edge case where strikeAskAmount rounds to 0
	// Rounding favors AMM vs Order
	if strikeAmountAsk.Equal(sdk.ZeroInt()) {
		return false, nil
	}

	// THEN set Head(Stop).Status to filled as entire order will be filled
	stopHead.Status = "filled"
	k.RemoveOrderOwner(ctx, stopHead.Owner, stopHead.Uid)

	// Set Next Position as Head of Stop Book
	memberBid.Stop = stopHead.Next

	if stopHead.Next != 0 {
		stopNext, _ := k.GetOrder(ctx, stopHead.Next)
		stopNext.Prev = 0
		k.SetOrder(ctx, stopNext)
	}

	// At this point the Head(Stop) position has been deactivated and the Next
	// Stop position has been set as the Head Stop

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

	// Update pool order history
	pool, _ := k.GetPool(ctx, memberAsk.Pair)
	pool.History = stopHead.Uid

	// order filled
	// just add the order to history
	if pool.History == 0 {
		stopHead.Prev = 0
		stopHead.Next = 0
	} else {
		prevFilledOrder, _ := k.GetOrder(ctx, pool.History)
		prevFilledOrder.Prev = stopHead.Uid
		stopHead.Prev = 0
		stopHead.Next = prevFilledOrder.Uid
		k.SetOrder(ctx, prevFilledOrder)
		k.SetOrderOwner(ctx, stopHead.Owner, stopHead.Uid)
	}

	stopHead.EndTime = ctx.BlockHeader().Time.Unix()

	k.SetOrder(ctx, stopHead)
	k.SetPool(ctx, pool)

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
