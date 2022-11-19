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

	pair := strings.Split(msg.Pair, ",")

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

	dropProfit := dropSumEnd.Sub(drop.Sum)

	// If dropSumEnd is equal to _drop.sum then
	// pool was created and then last drop redeemed
	// before any trade were executed leaving nothing in the pool
	// There-fore pool has no balances as basis for pricing
	// and must be burned.
	if dropSumEnd.Equal(drop.Sum) {

	}

	_ = dropProfit

	return &types.MsgRedeemDropResponse{}, nil
}
