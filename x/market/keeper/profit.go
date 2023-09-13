package keeper

import (
	"math/big"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pendulum-labs/market/x/market/types"
)

func (k Keeper) Profit(productBeg sdk.Int, productEnd sdk.Int, memberA types.Member, memberB types.Member) sdk.Coins {
	productDiff := productEnd.Sub(productBeg)

	profitA := big.NewInt(0)
	profitA.Sqrt(((productDiff.Mul(memberA.Balance)).Quo(memberB.Balance)).BigInt())
	profitCoinA := sdk.NewCoin(memberA.DenomB, sdk.NewIntFromBigInt(profitA))

	profitB := big.NewInt(0)
	profitB.Sqrt(((productDiff.Mul(memberB.Balance)).Quo(memberA.Balance)).BigInt())
	profitCoinB := sdk.NewCoin(memberB.DenomB, sdk.NewIntFromBigInt(profitB))
	return sdk.NewCoins(profitCoinA, profitCoinB)
}

func (k Keeper) Payout(ctx sdk.Context, profits sdk.Coins, pool types.Pool) (totalEarnings sdk.Coins, error error) {
	earnRatesStringSlice := strings.Split(k.EarnRates(ctx), ",")
	var earnRate sdk.Int
	var earnings sdk.Int
	var coinLeader sdk.Coin
	var coinsLeader sdk.Coins

	for i, v := range pool.Leaders {
		for _, profit := range profits {
			if profit.Amount == sdk.ZeroInt() {
				continue
			}

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

func (k Keeper) Burn(ctx sdk.Context, profits sdk.Coins) error {

	// burnRate, _ := sdk.NewIntFromString(k.BurnRate(ctx))

	// denominator := sdk.NewInt(10000)

	for _, profit := range profits {
		burningsAmount := profit.Amount.Quo(sdk.NewInt(100))

		burnings, found := k.GetBurnings(ctx, profit.Denom)
		if found {
			burnings.Amount = burnings.Amount.Add(burningsAmount)
		} else {
			burnings.Denom = profit.Denom
			burnings.Amount = burningsAmount
		}

		burnings, err := k.BurnTrade(ctx, burnings)
		if err != nil {
			return err
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
func (k Keeper) BurnTrade(ctx sdk.Context, burnings types.Burnings) (types.Burnings, error) {

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
