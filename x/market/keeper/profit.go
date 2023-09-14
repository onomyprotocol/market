package keeper

import (
	"math/big"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pendulum-labs/market/x/market/types"
)

func (k Keeper) Profit(productBeg sdk.Int, memberA types.Member, memberB types.Member) (profitA sdk.Int, profitB sdk.Int) {

	principalA := big.NewInt(0)
	principalA.Mul(productBeg.BigInt(), memberA.Balance.BigInt())
	principalA.Quo(principalA, memberB.Balance.BigInt())
	principalA.Sqrt(principalA)
	profitA = memberA.Balance.Sub(sdk.NewIntFromBigInt(principalA))

	principalB := big.NewInt(0)
	principalB.Mul(productBeg.BigInt(), memberB.Balance.BigInt())
	principalB.Quo(principalB, memberA.Balance.BigInt())
	principalB.Sqrt(principalB)
	profitB = memberB.Balance.Sub(sdk.NewIntFromBigInt(principalB))

	return
}

func (k Keeper) Payout(ctx sdk.Context, profit sdk.Int, member types.Member, pool types.Pool) (types.Member, error) {
	if profit == sdk.ZeroInt() {
		return member, nil
	}

	earnRatesStringSlice := strings.Split(k.EarnRates(ctx), ",")
	var earnRate sdk.Int
	var earnings sdk.Int
	var earningsCoin sdk.Coin
	var earningsCoins sdk.Coins

	for i, v := range pool.Leaders {

		earnRate, _ = sdk.NewIntFromString(earnRatesStringSlice[i])

		earnings = (profit.Mul(earnRate)).Quo(sdk.NewInt(10000))

		earningsCoin = sdk.NewCoin(member.DenomB, earnings)

		earningsCoins = sdk.NewCoins(earningsCoin)

		leader, _ := sdk.AccAddressFromBech32(v.Address)

		// Payout Leader
		sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, leader, earningsCoins)
		if sdkError != nil {
			return member, sdkError
		}

		member.Balance = member.Balance.Sub(earnings)
	}

	return member, nil
}

func (k Keeper) Burn(ctx sdk.Context, profit sdk.Int, member types.Member) (types.Member, error) {

	if profit == sdk.ZeroInt() {
		return member, nil
	}

	burnRate, _ := sdk.NewIntFromString(k.BurnRate(ctx))

	denominator := sdk.NewInt(10000)

	burningsAmount := (profit.Mul(burnRate)).Quo(denominator)

	member.Balance = member.Balance.Sub(burningsAmount)

	burnings, found := k.GetBurnings(ctx, member.DenomB)
	if found {
		burnings.Amount = burnings.Amount.Add(burningsAmount)
	} else {
		burnings.Denom = member.DenomB
		burnings.Amount = burningsAmount
	}

	burnings, err := k.BurnTrade(ctx, burnings)
	if err != nil {
		return member, err
	}

	if found && burnings.Amount == sdk.ZeroInt() {
		k.RemoveBurnings(ctx, burnings.Denom)
		return member, nil
	}

	if burnings.Amount.GT(sdk.ZeroInt()) {
		k.SetBurnings(ctx, burnings)
	}

	return member, nil
}

// Input Burnings - Output New Burnings
func (k Keeper) BurnTrade(ctx sdk.Context, burnings types.Burnings) (types.Burnings, error) {

	burnDenom := k.BurnCoin(ctx)

	burnCoin := sdk.NewCoin(burnDenom, burnings.Amount)

	if burnings.Denom != burnDenom {

		// Ask -> Burn Coin, Bid -> Coin traded for Burn Coin
		amountBid := burnings.Amount

		memberAsk, found := k.GetMember(ctx, burnings.Denom, burnDenom)
		if !found {
			return burnings, nil
		}

		memberBid, found := k.GetMember(ctx, burnDenom, burnings.Denom)
		if !found {
			return burnings, nil
		}

		// Market Order
		// A(i)*B(i) = A(f)*B(f)
		// A(f) = A(i)*B(i)/B(f)
		// strikeAmountAsk = A(i) - A(f) = A(i) - A(i)*B(i)/B(f)
		amountAsk := memberAsk.Balance.Sub((memberAsk.Balance.Mul(memberBid.Balance)).Quo(memberBid.Balance.Add(amountBid)))

		// Market Order Fee
		marketRate, _ := sdk.NewIntFromString(k.getParams(ctx).MarketFee)

		// Burn trades still payout fees
		fee := (amountAsk.Mul(marketRate)).Quo(sdk.NewInt(10000))
		amountAsk = amountAsk.Sub(fee)

		pool, found := k.GetPool(ctx, memberAsk.Pair)
		if !found {
			return burnings, nil
		}

		// Distribute to Leaders portion of the Burn Trade
		// Burn Trade does not Burn further or will create loop
		memberAsk, err := k.Payout(ctx, fee, memberAsk, pool)
		if err != nil {
			return burnings, err
		}

		memberAsk.Balance = memberAsk.Balance.Sub(amountAsk)
		memberBid.Balance = memberBid.Balance.Add(amountBid)

		k.SetMember(ctx, memberAsk)
		k.SetMember(ctx, memberBid)

		uid := k.GetUidCount(ctx)

		prevOrder, _ := k.GetOrder(ctx, pool.History)

		prevOrder.Prev = uid

		var order = types.Order{
			Uid:       uid,
			Owner:     "system",
			Status:    "filled",
			DenomAsk:  burnDenom,
			DenomBid:  burnings.Denom,
			OrderType: "burn",
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

		burnCoin = sdk.NewCoin(burnDenom, amountAsk)

	}

	if burnCoin.Amount.GT(sdk.ZeroInt()) {

		burnCoins := sdk.NewCoins(burnCoin)

		// Burn Ask Amount of Burn Coin
		sdkError := k.bankKeeper.BurnCoins(ctx, types.ModuleName, burnCoins)
		if sdkError != nil {
			return burnings, sdkError
		}

		k.AddBurned(ctx, burnCoins.AmountOf(burnDenom))

		burnings.Amount = burnings.Amount.Sub(burnCoin.Amount)

	}

	return burnings, nil
}
