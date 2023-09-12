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

func (k msgServer) RedeemDrop(goCtx context.Context, msg *types.MsgRedeemDrop) (*types.MsgRedeemDropResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	uid, _ := strconv.ParseUint(msg.Uid, 10, 64)

	drop, found := k.GetDrop(ctx, uid)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrDropNotFound, "%s", msg.Uid)
	}

	if drop.Owner != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotDrops, "%s", msg.Uid)
	}

	pair := strings.Split(drop.Pair, ",")

	denom1 := pair[0]
	denom2 := pair[1]

	pool, found := k.GetPool(ctx, drop.Pair)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrPoolNotFound, "%s", drop.Pair)
	}

	member1, found := k.GetMember(ctx, denom2, denom1)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrMemberNotFound, "%s", drop.Pair)
	}

	member2, found := k.GetMember(ctx, denom1, denom2)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrMemberNotFound, "%s", drop.Pair)
	}

	// `total1 = (drop.Drops * member1.Balance) / pool.Drops`
	tmp := big.NewInt(0)
	tmp.Mul(drop.Drops.BigInt(), member1.Balance.BigInt())
	tmp.Quo(tmp, pool.Drops.BigInt())
	total1 := sdk.NewIntFromBigInt(tmp)
	// note: because of https://github.com/cosmos/cosmos-sdk/issues/17342
	// always run this after a call to `NewIntFromBigInt`
	tmp = big.NewInt(0)

	// `total2 = (drop.Drops * member2.Balance) / pool.Drops`
	tmp.Mul(drop.Drops.BigInt(), member2.Balance.BigInt())
	tmp.Quo(tmp, pool.Drops.BigInt())
	total2 := sdk.NewIntFromBigInt(tmp)
	// tmp = big.NewInt(0)

	dropRedeemer, ok := k.GetDropsOwnerPair(ctx, msg.Creator, drop.Pair)
	var sumDropRedeemer sdk.Int
	if ok {
		sumDropRedeemer = dropRedeemer.Sum
	} else {
		return nil, sdkerrors.Wrapf(types.ErrDropSumNotFound, "%s", msg.Creator)
	}

	sumDropRedeemer = sumDropRedeemer.Sub(drop.Drops)

	pool = k.updateLeaders(ctx, pool, msg.Creator, sumDropRedeemer)

	var sdkError error

	// Update Pool Total Drops
	pool.Drops = pool.Drops.Sub(drop.Drops)

	// Withdraw from Pool
	member1.Balance = member1.Balance.Sub(total1)
	member2.Balance = member2.Balance.Sub(total2)

	// moduleAcc := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	// Get the borrower address
	owner, _ := sdk.AccAddressFromBech32(msg.Creator)

	coinOwner1 := sdk.NewCoin(denom1, total1)
	coinOwner2 := sdk.NewCoin(denom2, total2)
	coinsOwner := sdk.NewCoins(coinOwner1, coinOwner2)

	// Payout Owner
	sdkError = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, owner, coinsOwner)
	if sdkError != nil {
		return nil, sdkError
	}

	// Deactivate drop
	drop.Active = false

	// Set Pool Member and Drop
	k.SetDrop(
		ctx,
		drop,
	)

	k.RemoveDropOwner(
		ctx,
		drop,
	)

	k.SetPool(
		ctx,
		pool,
	)

	k.SetMember(
		ctx,
		member1,
	)

	k.SetMember(
		ctx,
		member2,
	)

	return &types.MsgRedeemDropResponse{}, nil
}

func Payout(k msgServer, ctx sdk.Context, profits sdk.Coins, pool types.Pool) (totalEarnings sdk.Coins, error error) {
	earnRatesStringSlice := strings.Split(k.EarnRates(ctx), ",")
	var earnRate sdk.Int
	var earnings sdk.Int
	var coinLeader sdk.Coin
	var coinsLeader sdk.Coins

	for i, v := range pool.Leaders {
		for _, profit := range profits {
			earnRate, _ = sdk.NewIntFromString(earnRatesStringSlice[i])

			earnings = (profit.Amount.Mul(earnRate)).Quo(sdk.NewInt(10000))

			coinLeader = sdk.NewCoin(profit.Denom, earnings)

			coinsLeader = coinsLeader.Add(coinLeader)
		}

		leader, _ := sdk.AccAddressFromBech32(v.Address)

		// Payout Leader
		sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, leader, coinsLeader)
		if sdkError != nil {
			return nil, sdkError
		}

		totalEarnings = totalEarnings.Add(coinsLeader...)
	}

	return
}

func Burn(k msgServer, ctx sdk.Context, profits sdk.Coins) error {

	burnRate, _ := sdk.NewIntFromString(k.BurnRate(ctx))

	denominator, _ := sdk.NewIntFromString("10000")

	for _, profit := range profits {
		burningsAmount := (profit.Amount.Mul(burnRate)).Quo(denominator)

		burnings, found := k.GetBurnings(ctx, profit.Denom)
		if found {
			burnings.Amount = burnings.Amount.Add(burningsAmount)
		} else {
			burnings.Denom = profit.Denom
			burnings.Amount = burningsAmount
		}

		burnings, sdkError := BurnTrade(k, ctx, burnings)
		if sdkError != nil {
			return sdkError
		}

		if found && burnings.Amount == sdk.ZeroInt() {
			k.RemoveBurnings(ctx, burnings.Denom)
			continue
		}

		if burnings.Amount.GT(sdk.ZeroInt()) {
			k.SetBurnings(ctx, burnings)
		}
	}

	return nil
}

// Input Burnings - Output New Burnings
func BurnTrade(k msgServer, ctx sdk.Context, burnings types.Burnings) (types.Burnings, error) {

	burnCoin := k.BurnCoin(ctx)

	coinBurn := sdk.NewCoin(burnCoin, burnings.Amount)

	if burnings.Denom != burnCoin {

		// Ask -> Burn Coin, Bid -> Coin traded for Burn Coin
		amountBid := burnings.Amount

		memberAsk, found := k.GetMember(ctx, burnings.Denom, burnCoin)
		if !found {
			return burnings, nil
		}

		memberBid, found := k.GetMember(ctx, burnCoin, burnings.Denom)
		if !found {
			return burnings, nil
		}

		// Market Order
		// A(i)*B(i) = A(f)*B(f)
		// A(f) = A(i)*B(i)/B(f)
		// strikeAmountAsk = A(i) - A(f) = A(i) - A(i)*B(i)/B(f)
		amountAsk := memberAsk.Balance.Sub((memberAsk.Balance.Mul(memberBid.Balance)).Quo(memberBid.Balance.Add(amountBid)))

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
			Owner:     "system",
			Status:    "filled",
			DenomAsk:  burnCoin,
			DenomBid:  burnings.Denom,
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

		coinBurn = sdk.NewCoin(burnCoin, amountAsk)

	}

	if coinBurn.Amount.GT(sdk.ZeroInt()) {

		coinsBurn := sdk.NewCoins(coinBurn)

		// Burn Ask Amount of Burn Coin
		sdkError := k.bankKeeper.BurnCoins(ctx, types.ModuleName, coinsBurn)
		if sdkError != nil {
			return burnings, sdkError
		}

		k.AddBurned(ctx, coinsBurn.AmountOf(burnCoin))

		burnings.Amount = burnings.Amount.Sub(coinBurn.Amount)

	}

	return burnings, nil
}
