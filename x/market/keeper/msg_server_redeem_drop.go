package keeper

import (
	"context"
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
		return nil, sdkerrors.Wrapf(types.ErrNotDropOwner, "%s", msg.Uid)
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

	poolSum := member1.Balance.Add(member2.Balance)

	// Each Drop is a proportional rite to the AMM balances
	// % total drops in pool = dropAmount(drop)/dropAmount(pool)*100%
	// drop_ratio = dropAmount(drop)/dropAmount(pool)
	// dropSum(end) = (AMM Bal A + AMM Bal B) * drop_ratio
	// Profit = dropSum(end) - dropSum(begin)
	dropSumEnd := (poolSum.Mul(drop.Drops)).Quo(pool.Drops)
	total2 := (dropSumEnd.Mul(sdk.NewInt(10 ^ 18))).Quo((poolSum.Mul(sdk.NewInt(10 ^ 18))).Quo(member2.Balance))
	total1 := dropSumEnd.Sub(total2)

	dropProfit := dropSumEnd.Sub(drop.Sum)

	earnRate := k.EarnRate(ctx)
	burnRate := k.BurnRate(ctx)

	// (dropProfit * bigNum) / ( poolSum * bigNum / member2.balance )
	profit2 := (dropProfit.Mul(sdk.NewInt(10 ^ 18))).Quo((poolSum.Mul(sdk.NewInt(10 ^ 18))).Quo(member2.Balance))
	earn2 := (profit2.Mul(earnRate[0])).Quo(earnRate[1])
	burn2 := (profit2.Mul(burnRate[0])).Quo(burnRate[1])

	// Redemption value in coin 2
	dropOwner2 := total2.Sub(earn2.Add(burn2))

	profit1 := dropProfit.Sub(profit2)
	earn1 := (profit1.Mul(earnRate[0])).Quo(earnRate[1])
	burn1 := (profit1.Mul(burnRate[0])).Quo(burnRate[1])

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

	// Redemption value in coin 1
	dropOwner1 := total1.Sub(earn1.Add(burn1))

	// Update Pool Total Drops
	pool.Drops = pool.Drops.Sub(drop.Drops)

	// Withdraw from Pool
	member1.Balance = member1.Balance.Sub(total1)
	member2.Balance = member2.Balance.Sub(total2)

	// moduleAcc := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	// Get the borrower address
	owner, _ := sdk.AccAddressFromBech32(msg.Creator)

	coinOwner1 := sdk.NewCoin(denom1, dropOwner1)
	coinOwner2 := sdk.NewCoin(denom2, dropOwner2)
	coinsOwner := sdk.NewCoins(coinOwner1, coinOwner2)

	// Payout Owner
	sdkError = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, owner, coinsOwner)
	if sdkError != nil {
		return nil, sdkError
	}

	coinLeader1 := sdk.NewCoin(denom1, earn1)
	coinLeader2 := sdk.NewCoin(denom2, earn2)
	coinsLeader := sdk.NewCoins(coinLeader1, coinLeader2)

	leader, _ := sdk.AccAddressFromBech32(pool.Leader)

	// Payout Leader
	sdkError = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, leader, coinsLeader)
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

	maxMemberBidBal := memberAsk.Balance.Add(memberBid.Balance).Sub(memberAsk.Balance.Quo(sdk.NewInt(2)))
	maxMemberBidAmount := maxMemberBidBal.Sub(memberBid.Balance)

	// Partial order may consume only half of memberAsk pool amount
	if amountBid.GT(maxMemberBidAmount) {
		amountBid = maxMemberBidAmount
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

	coinAsk := sdk.NewCoin(burnCoin, amountAsk)
	coinsAsk := sdk.NewCoins(coinAsk)

	// Burn Ask Amount of Stake Coin
	sdkError := k.bankKeeper.BurnCoins(ctx, "mint", coinsAsk)
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
