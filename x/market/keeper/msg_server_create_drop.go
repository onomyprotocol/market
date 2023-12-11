package keeper

import (
	"context"
	"sort"
	"strings"

	"market/x/market/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

	if member1.Balance.Equal(sdk.ZeroInt()) {
		return nil, sdkerrors.Wrapf(types.ErrMemberBalanceZero, "Member %s", member1.DenomB)
	}

	if member2.Balance.Equal(sdk.ZeroInt()) {
		return nil, sdkerrors.Wrapf(types.ErrMemberBalanceZero, "Member %s", member2.DenomB)
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

	member2.Balance = member2.Balance.Add(dropAmtMember2)
	k.SetMember(ctx, member2)

	dropCreatorSum := drops
	dropOwner, ok := k.GetDropsOwnerPair(ctx, msg.Creator, pair)

	if ok {
		dropCreatorSum = dropCreatorSum.Add(dropOwner.Sum)
	}

	pool = k.updateLeaders(ctx, pool, msg.Creator, dropCreatorSum)

	pool.Drops = pool.Drops.Add(drops)

	k.SetPool(ctx, pool)

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

	// Update drop uid count
	k.SetUidCount(ctx, uid+1)

	return &types.MsgCreateDropResponse{}, nil
}

func (k msgServer) updateLeaders(ctx sdk.Context, pool types.Pool, dropCreator string, dropCreatorSum sdk.Int) types.Pool {
	maxLeaders := len(strings.Split(k.EarnRates(ctx), ","))

	for i := 0; i < len(pool.Leaders); i++ {
		if pool.Leaders[i].Address == dropCreator {
			pool.Leaders = pool.Leaders[:i+copy(pool.Leaders[i:], pool.Leaders[i+1:])]
		}
	}

	if dropCreatorSum.Equal(sdk.ZeroInt()) {
		return pool
	}

	if len(pool.Leaders) == 0 {
		pool.Leaders = append(pool.Leaders, &types.Leader{
			Address: dropCreator,
			Drops:   dropCreatorSum,
		})
	} else {
		for i := 0; i < len(pool.Leaders); i++ {
			if dropCreatorSum.GT(pool.Leaders[i].Drops) {
				if len(pool.Leaders) < maxLeaders {
					pool.Leaders = append(pool.Leaders, pool.Leaders[len(pool.Leaders)-1])
				}
				copy(pool.Leaders[i+1:], pool.Leaders[i:])
				pool.Leaders[i] = &types.Leader{
					Address: dropCreator,
					Drops:   dropCreatorSum,
				}
				break
			} else {
				if (i == len(pool.Leaders)-1) && len(pool.Leaders) < maxLeaders {
					pool.Leaders = append(pool.Leaders, &types.Leader{
						Address: dropCreator,
						Drops:   dropCreatorSum,
					})
					break
				}
			}
		}
	}
	return pool
}
