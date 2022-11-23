package keeper

import (
	"context"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/onomyprotocol/market/x/market/types"
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

	denom1 := pair[1]
	denom2 := pair[2]

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

	// (dropSumFinal * bigNum) / ( poolSum * bigNum / member2.balance )
	profit2 := (dropProfit.Mul(sdk.NewInt(10 ^ 18))).Quo((poolSum.Mul(sdk.NewInt(10 ^ 18))).Quo(member2.Balance))
	earn2 := (profit2.Mul(earnRate[0])).Quo(earnRate[1])
	burn2 := (profit2.Mul(burnRate[0])).Quo(burnRate[1])

	// Redemption value in coin 2
	drop2 := total2.Sub(earn2.Add(burn2))

	profit1 := dropProfit.Sub(profit2)
	earn1 := (profit1.Mul(earnRate[0])).Quo(earnRate[1])
	burn1 := (profit1.Mul(burnRate[0])).Quo(burnRate[1])

	// Redemption value in coin 1
	drop1 := total1.Sub(earn1.Add(burn1))

	// Update Pool Total Drops
	pool.Drops = pool.Drops.Sub(drop.Drops)

	// Withdraw from Pool
	member1.Balance = member1.Balance.Sub(total1)
	member2.Balance = member2.Balance.Sub(total2)

	// moduleAcc := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	// Get the borrower address
	owner, _ := sdk.AccAddressFromBech32(msg.Creator)

	coin1 := sdk.NewCoin(denom1, drop1)
	coin2 := sdk.NewCoin(denom2, drop2)
	coins := sdk.NewCoins(coin1, coin2)

	// Increment Owner
	sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, owner, coins)
	if sdkError != nil {
		return nil, sdkError
	}

	drop.Active = false

	// Set Pool and Drop

	return &types.MsgRedeemDropResponse{}, nil
}
