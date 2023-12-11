package keeper

import (
	"context"
	"math/big"
	"strconv"
	"strings"

	"market/x/market/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
