package keeper

import (
	"context"
	"math/big"
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

	if pool.Drops.Equal(sdk.NewInt(0)) {
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

	// The Pool Sum is defined as:
	// poolProduct == AMM Coin A Balance * AMM Coin B Balance
	poolProduct := member1.Balance.Mul(member2.Balance)

	drops, _ := sdk.NewIntFromString(msg.Drops)

	// The beginning Drop Sum is defined as:
	// dropProduct == Total amount of coinA+coinB needed to create the drop based on pool exchange rate
	// dropProduct == poolProduct * (Drop.drops / Pool.drops)
	// dropProduct == (poolSum * Drop.drops) / Pool.drops
	dropProduct := (poolProduct.Mul(drops)).Quo(pool.Drops)

	// dropProduct == A * B
	// dropProduct = B * B * exchrate(A/B)
	// dropProduct = B^2 * exchrate(A/B)
	// B^2 = dropProduct / exchrate(A/B)
	// B^2 = dropProduct / (Member1 Balance / Member2 Balance)
	// B^2 = (dropProduct * Member2 Balance) / Member1
	// B = SQRT((dropProduct * Member2 Balance) / Member1)
	bigInt := &big.Int{}
	amount2 :=
		sdk.NewIntFromBigInt(
			bigInt.Sqrt(
				sdk.Int.BigInt(
					(dropProduct.Mul(member2.Balance)).Quo(member1.Balance),
				),
			),
		)

	coin2 := sdk.NewCoin(denom2, amount2)

	amount1 := dropProduct.Quo(amount2)

	coin1 := sdk.NewCoin(denom1, amount1)

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

	// update member2 event
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
	sumDropCreator := k.GetDropsSum(ctx, msg.Creator).Add(drops)

	numLeaders := len(pool.Leaders)
	maxLeaders := len(strings.Split(k.EarnRates(ctx), ","))

	index := numLeaders

	// Check if Drop Creator is already on leader board
	// If so, make index = drop creator position
	for i := 0; i < numLeaders; i++ {
		if pool.Leaders[i].Address == msg.Creator {
			index = i
			break
		}
	}

	if index == 0 {
		// If drop creator is already top of leader board
		// Only update number of drops
		pool.Leaders[0].Drops = sumDropCreator
	} else {
		for i := index - 1; i >= 0; i-- {
			if sumDropCreator.GT(pool.Leaders[i].Drops) {
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
							Address: msg.Creator,
							Drops:   sumDropCreator,
						}
					}
				}
			} else {
				if index == numLeaders && numLeaders < maxLeaders {
					pool.Leaders = append(pool.Leaders, &types.Leader{
						Address: msg.Creator,
						Drops:   sumDropCreator,
					})
				} else {
					pool.Leaders[i+1].Address = msg.Creator
					pool.Leaders[i+1].Drops = sumDropCreator
				}
				break
			}
		}
	}

	pool.Drops = pool.Drops.Add(drops)

	k.SetPool(ctx, pool)

	var leaders []string

	for i := 0; i < numLeaders; i++ {
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
