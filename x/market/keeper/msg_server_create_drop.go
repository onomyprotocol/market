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

	if pool.Drops.Equal(sdk.ZeroInt()) {
		return nil, sdkerrors.Wrapf(types.ErrPoolInactive, "%s", pair)
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

	drops, _ := sdk.NewIntFromString(msg.Drops)

	dropAmtMember1, dropAmtMember2, error := dropAmounts(drops, pool, member1, member2)
	if error != nil {
		return nil, error
	}

	dropProduct := dropAmtMember1.Mul(dropAmtMember2)

	coin1 := sdk.NewCoin(denom1, dropAmtMember1)
	coin2 := sdk.NewCoin(denom2, dropAmtMember2)

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
	member1.Balance = member1.Balance.Add(dropAmtMember1)
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

	member2.Balance = member2.Balance.Add(dropAmtMember2)
	k.SetMember(ctx, member2)

	// update member2 event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpdateMember,
			sdk.NewAttribute(types.AttributeKeyDenomA, denom1),
			sdk.NewAttribute(types.AttributeKeyDenomB, denom2),
			sdk.NewAttribute(types.AttributeKeyBalance, member2.Balance.String()),
		),
	)

	dropCreatorSum := drops
	dropOwner, ok := k.GetDropsOwnerPair(ctx, msg.Creator, pair)

	if ok {
		dropCreatorSum = dropCreatorSum.Add(dropOwner.Sum)
	}

	pool = k.updateLeaders(ctx, pool, msg.Creator, dropCreatorSum)

	pool.Drops = pool.Drops.Add(drops)

	k.SetPool(ctx, pool)

	var leaders []string

	for i := 0; i < len(pool.Leaders); i++ {
		leaders = append(leaders, "{"+strings.Join([]string{pool.Leaders[i].Address, pool.Leaders[i].Drops.String()}, ", ")+"}")
	}

	// update pool event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpdatePool,
			sdk.NewAttribute(types.AttributeKeyPair, pair),
			sdk.NewAttribute(types.AttributeKeyLeaders, strings.Join(leaders, ", ")),
			sdk.NewAttribute(types.AttributeKeyAmount, pool.Drops.String()),
		),
	)

	var drop = types.Drop{
		Uid:     uid,
		Owner:   msg.Creator,
		Pair:    pair,
		Drops:   drops,
		Product: dropProduct,
		Active:  true,
	}

	// Add the drop to the keeper
	k.SetDrop(
		ctx,
		drop,
	)

	k.SetDropOwner(
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
			sdk.NewAttribute(types.AttributeKeyProduct, dropProduct.String()),
		),
	)

	// Update drop uid count
	k.SetUidCount(ctx, uid+1)

	return &types.MsgCreateDropResponse{}, nil
}

func (k msgServer) updateLeaders(ctx sdk.Context, pool types.Pool, dropCreator string, dropCreatorSum sdk.Int) types.Pool {
	numLeaders := len(pool.Leaders)
	maxLeaders := len(strings.Split(k.EarnRates(ctx), ","))

	index := numLeaders

	// Check if Drop Creator is already on leader board
	// If so, make index = drop creator position
	for i := 0; i < numLeaders; i++ {
		if pool.Leaders[i].Address == dropCreator {
			index = i
			break
		}
	}

	if index == 0 {
		// If drop creator is already top of leader board
		// Only update number of drops
		pool.Leaders[0].Drops = dropCreatorSum
	} else {
		for i := index - 1; i >= 0; i-- {
			if dropCreatorSum.GT(pool.Leaders[i].Drops) {
				// Append
				if i == index-1 && index == numLeaders && numLeaders < maxLeaders {
					pool.Leaders = append(pool.Leaders, pool.Leaders[i])
				} else {
					// If drop creator has more total drops move
					// this position down the leader board
					pool.Leaders[i+1] = pool.Leaders[i]
					// If at top of the list then place drop creator
					// as top leader
					if i == 0 {
						pool.Leaders[0] = &types.Leader{
							Address: dropCreator,
							Drops:   dropCreatorSum,
						}
					}
				}
			} else {
				if index == numLeaders && numLeaders < maxLeaders {
					pool.Leaders = append(pool.Leaders, &types.Leader{
						Address: dropCreator,
						Drops:   dropCreatorSum,
					})
				} else {
					pool.Leaders[i+1].Address = dropCreator
					pool.Leaders[i+1].Drops = dropCreatorSum
				}
				break
			}
		}
	}
	return pool
}
