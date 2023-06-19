package keeper

import (
	"context"
	"sort"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pendulum-labs/market/x/market/types"
)

func (k msgServer) CreateDrop(goCtx context.Context, msg *types.MsgCreateDrop) (*types.MsgCreateDropResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pairMsg := strings.Split(msg.Pair, ",")
	sort.Strings(pairMsg)

	denom1 := pairMsg[0]
	denom2 := pairMsg[1]

	pair := strings.Join(pairMsg, ",")

	pool, found := k.GetPool(ctx, pair)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrPoolNotFound, "%s", pair)
	}

	member1, found := k.GetMember(ctx, denom2, denom1)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrMemberNotFound, "%s", pair)
	}

	member2, found := k.GetMember(ctx, denom1, denom2)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrMemberNotFound, "%s", pair)
	}

	// Create the uid
	uid := k.GetUidCount(ctx)

	// The Pool Sum is defined as:
	// poolSum == AMM Coin A Balance + AMM Coin B Balance
	poolSum := member1.Balance.Add(member2.Balance)

	drops, _ := sdk.NewIntFromString(msg.Drops)

	// The beginning Drop Sum is defined as:
	// dropSum == Total amount of coinA+coinB needed to create the drop based on pool exchange rate
	// dropSum == poolSum * (Drop.drops / Pool.drops)
	// dropSum == (poolSum * Drop.drops) / Pool.drops
	dropSum := (poolSum.Mul(drops)).Quo(pool.Drops)

	// dropSum == A + B
	// dropSum = B + B * exchrate(A/B)
	// dropSum = B * (1 + exchrate(A/B))
	// B = dropSum / (1 + exchrate(A/B))
	// B = dropSum / (1 + Member1 Balance / Member2 Balance)
	// B = dropSum / ((Member1 + Member2) / Member2)
	// B = dropSum / (poolSum / Member2)
	// B = (dropSum * Member2) / poolSum
	amount1 := (dropSum.Mul(member2.Balance)).Quo(poolSum)

	coin1 := sdk.NewCoin(denom1, amount1)

	// The purchase price of the drop in A coin must be less than Available Balance
	amount2 := dropSum.Sub(amount1)

	coin2 := sdk.NewCoin(denom2, amount2)

	coinPair := sdk.NewCoins(coin1, coin2)

	// moduleAcc := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	// Get the borrower address
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)

	if err := k.validateSenderBalance(ctx, creator, coinPair); err != nil {
		return nil, err
	}

	// Use the module account as pool account
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, coinPair)
	if sdkError != nil {
		return nil, sdkError
	}

	// Deposit into Pool
	member1.Balance = member1.Balance.Add(amount1)
	k.SetMember(ctx, member1)

	// update member1 event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpdateMember,
			sdk.NewAttribute(types.AttributeKeyDenomA, denom2),
			sdk.NewAttribute(types.AttributeKeyDenomB, denom1),
			sdk.NewAttribute(types.AttributeKeyBalance, member1.Balance.String()),
		),
	)

	member2.Balance = member2.Balance.Add(amount2)
	k.SetMember(ctx, member2)

	// update member1 event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpdateMember,
			sdk.NewAttribute(types.AttributeKeyDenomA, denom1),
			sdk.NewAttribute(types.AttributeKeyDenomB, denom2),
			sdk.NewAttribute(types.AttributeKeyBalance, member2.Balance.String()),
		),
	)

	// Get Drop Creator and Pool Leader total drops from all drops owned
	// TODO: Need to double check that database is configured properly
	sumDropsCreator := k.GetDropsSum(ctx, msg.Creator).Add(drops)
	sumDropsLeader := k.GetDropsSum(ctx, pool.Leader)

	// If Creator totaled owned drops is greater than Leader then
	// Creator is new leader
	if sumDropsCreator.GT(sumDropsLeader) {
		pool.Leader = msg.Creator
	}

	pool.Drops = pool.Drops.Add(drops)

	k.SetPool(ctx, pool)

	// update pool event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreatePool,
			sdk.NewAttribute(types.AttributeKeyPair, pair),
			sdk.NewAttribute(types.AttributeKeyLeader, pool.Leader),
			sdk.NewAttribute(types.AttributeKeyAmount, pool.Drops.String()),
		),
	)

	var drop = types.Drop{
		Uid:    uid,
		Owner:  msg.Creator,
		Pair:   pair,
		Drops:  drops,
		Sum:    dropSum,
		Active: true,
	}

	// Add the drop to the keeper
	k.SetDrop(
		ctx,
		drop,
	)

	// create drop event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreateDrop,
			sdk.NewAttribute(types.AttributeKeyUid, strconv.FormatUint(uid, 10)),
			sdk.NewAttribute(types.AttributeKeyPair, pair),
			sdk.NewAttribute(types.AttributeKeyOwner, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyAmount, drops.String()),
			sdk.NewAttribute(types.AttributeKeySum, dropSum.String()),
		),
	)

	// Update drop uid count
	k.SetUidCount(ctx, uid+1)

	return &types.MsgCreateDropResponse{}, nil
}
