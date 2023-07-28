package keeper

import (
	"context"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
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

	poolProduct := member1.Balance.Mul(member2.Balance)

	// Each Drop is a proportional rite to the AMM balances
	// % total drops in pool = dropAmount(drop)/dropAmount(pool)*100%
	// drop_ratio = dropAmount(drop)/dropAmount(pool)
	// dropProduct(end) = (AMM Bal A * AMM Bal B) * drop_ratio
	// Profit = total*[(dropProduct(end) - dropProduct(begin))/dropProduct(end)]
	dropProductEnd := (poolProduct.Mul(drop.Drops)).Quo(pool.Drops)

	total1 := (drop.Drops.Mul(member1.Balance)).Quo(pool.Drops)
	profit1 := (total1.Mul(dropProductEnd)).Quo(dropProductEnd).Sub((total1.Mul(drop.Product)).Quo(dropProductEnd))

	total2 := (drop.Drops.Mul(member2.Balance)).Quo(pool.Drops)
	profit2 := (total2.Mul(dropProductEnd)).Quo(dropProductEnd).Sub((total1.Mul(drop.Product)).Quo(dropProductEnd))

	earnRatesStringSlice := strings.Split(k.EarnRates(ctx), ",")
	var earnRate sdk.Int
	var earnings1 sdk.Int
	earnings1Total := sdk.NewInt(0)
	var earnings2 sdk.Int
	earnings2Total := sdk.NewInt(0)
	var coinLeader1 sdk.Coin
	var coinLeader2 sdk.Coin
	var coinsLeader sdk.Coins

	for i, v := range pool.Leaders {
		earnRate, _ = sdk.NewIntFromString(earnRatesStringSlice[i])

		earnings1 = (profit1.Mul(earnRate)).Quo(sdk.NewInt(10000))
		earnings1Total = earnings1Total.Add(earnings1)

		earnings2 = (profit2.Mul(earnRate)).Quo(sdk.NewInt(10000))
		earnings2Total = earnings2Total.Add(earnings2)

		coinLeader1 = sdk.NewCoin(denom1, earnings1)
		coinLeader2 = sdk.NewCoin(denom2, earnings2)
		coinsLeader = sdk.NewCoins(coinLeader1, coinLeader2)

		leader, _ := sdk.AccAddressFromBech32(v.Address)

		// Payout Leader
		sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, leader, coinsLeader)
		if sdkError != nil {
			return nil, sdkError
		}
	}

	sumDropRedeemer, ok := k.GetDropsSum(ctx, msg.Creator, drop.Pair)

	if ok {
		sumDropRedeemer = sumDropRedeemer.Sub(drop.Drops)
	} else {
		return nil, sdkerrors.Wrapf(types.ErrDropSumNotFound, "%s", msg.Creator)
	}

	numLeaders := len(pool.Leaders)

	var index int
	flag := false

	// Check if Drop Redeemer is on leader board
	// If so, make index = drop creator position
	for i := 0; i < numLeaders; i++ {
		if pool.Leaders[i].Address == msg.Creator {
			index = i
			flag = true
			break
		}
	}

	// Re-order leader board to reflect new rankings
	if flag {
		if sumDropRedeemer.GT(sdk.NewInt(0)) {
			if numLeaders > 1 {
				for i := index + 1; i < numLeaders; i++ {
					if sumDropRedeemer.LT(pool.Leaders[i].Drops) {
						// If drop reedemer has less total drops move
						// this position down the leader board
						pool.Leaders[i-1].Address = pool.Leaders[i].Address
						pool.Leaders[i-1].Drops = pool.Leaders[i].Drops
						// Remove redeemer from bottom of list
						// Prevents someone from liquidating all drops
						// yet still receive leader rewards
						if i == numLeaders-1 {
							break
						}
						pool.Leaders[i].Address = msg.Creator
						pool.Leaders[i].Drops = sumDropRedeemer
					} else {
						break
					}
				}
			} else {
				pool.Leaders[0].Drops = sumDropRedeemer
			}
		} else {
			// If drop redeemer is the only leader and has zero drops
			// set pool.Leaders to nil
			if numLeaders == 1 {
				pool.Leaders = nil
			}
			for i := index + 1; i < numLeaders; i++ {
				pool.Leaders[i-1].Address = pool.Leaders[i].Address
				pool.Leaders[i-1].Drops = pool.Leaders[i].Drops
				// Remove redeemer from bottom of list
				// Prevents someone from liquidating all drops
				// yet still receive leader rewards
				if i == numLeaders-1 {
					pool.Leaders[i] = nil
					break
				}
			}
		}
	}

	burnRate, _ := sdk.NewIntFromString(k.BurnRate(ctx))

	burn1 := (profit1.Mul(burnRate)).Quo(sdk.NewInt(10000))

	// Redemption value in coin 1
	redeem1 := total1.Sub(earnings1Total.Add(burn1))

	burn2 := (profit2.Mul(burnRate)).Quo(sdk.NewInt(10000))

	// Redemption value in coin 2
	redeem2 := total2.Sub(earnings2Total.Add(burn2))

	var sdkError error

	// Update burnings
	burnings1, found := k.GetBurnings(ctx, denom1)
	if !found {
		burnings1 = types.Burnings{
			Denom:  denom1,
			Amount: burn1,
		}
		burnings1, sdkError = Burn(k, ctx, burnings1)
		if sdkError != nil {
			return nil, sdkError
		}
	} else {
		burnings1.Amount = burnings1.Amount.Add(burn1)
		burnings1, sdkError = Burn(k, ctx, burnings1)
		if sdkError != nil {
			return nil, sdkError
		}
	}
	k.SetBurnings(ctx, burnings1)

	burnings2, found := k.GetBurnings(ctx, denom2)
	if !found {
		burnings2 = types.Burnings{
			Denom:  denom2,
			Amount: burn2,
		}
		burnings2, sdkError = Burn(k, ctx, burnings2)
		if sdkError != nil {
			return nil, sdkError
		}
	} else {
		burnings2.Amount = burnings2.Amount.Add(burn2)
		burnings2, sdkError = Burn(k, ctx, burnings2)
		if sdkError != nil {
			return nil, sdkError
		}
	}

	k.SetBurnings(ctx, burnings2)

	// Update Pool Total Drops
	pool.Drops = pool.Drops.Sub(drop.Drops)

	// Withdraw from Pool
	member1.Balance = member1.Balance.Sub(total1)
	member2.Balance = member2.Balance.Sub(total2)

	// moduleAcc := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	// Get the borrower address
	owner, _ := sdk.AccAddressFromBech32(msg.Creator)

	coinOwner1 := sdk.NewCoin(denom1, redeem1)
	coinOwner2 := sdk.NewCoin(denom2, redeem2)
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

	k.RemoveDropFromList(
		ctx,
		drop.Uid,
		drop.Owner,
		drop.Pair,
	)

	k.SetDropsSum(
		ctx,
		drop.Owner,
		drop.Pair,
		sumDropRedeemer,
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

func Burn(k msgServer, ctx sdk.Context, burnings types.Burnings) (types.Burnings, error) {

	amountBid := burnings.Amount

	// Coin that will be burned
	burnCoin := k.BurnCoin(ctx)

	memberAsk, found := k.GetMember(ctx, burnings.Denom, burnCoin)
	if !found {
		return burnings, nil
	}

	memberBid, found := k.GetMember(ctx, burnCoin, burnings.Denom)
	if !found {
		return burnings, nil
	}

	// A(i)*B(i) = A(f)*B(f)
	// A(f) = A(i)*B(i)/B(f)
	// strikeAmountAsk = A(i) - A(f) = A(i) - A(i)*B(i)/B(f)
	amountAsk := memberAsk.Balance.Sub((memberAsk.Balance.Mul(memberBid.Balance)).Quo(memberBid.Balance.Add(amountBid)))

	coinAsk := sdk.NewCoin(burnCoin, amountAsk)
	coinsAsk := sdk.NewCoins(coinAsk)

	// Burn Ask Amount of Stake Coin
	sdkError := k.bankKeeper.BurnCoins(ctx, minttypes.ModuleName, coinsAsk)
	if sdkError != nil {
		return burnings, sdkError
	}

	memberAsk.Balance = memberAsk.Balance.Sub(amountAsk)
	memberBid.Balance = memberBid.Balance.Add(amountBid)

	k.SetMember(ctx, memberAsk)
	k.SetMember(ctx, memberBid)

	burnings.Amount = burnings.Amount.Sub(amountBid)

	return burnings, nil
}
