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

	denom1 := pairMsg[1]
	denom2 := pairMsg[2]

	pair := strings.Join(pairMsg, ",")

	pool, found := k.GetPool(ctx, pair)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrPoolNotFound, "%s", pair)
	}

	member1, found := k.GetMember(ctx, denom2, denom1)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrPoolNotFound, "%s", pair)
	}

	member2, found := k.GetMember(ctx, denom1, denom2)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrPoolNotFound, "%s", pair)
	}

	// Create the uid
	uid := k.GetUidCount(ctx)

	prev1, _ := strconv.ParseUint(msg.Prev1, 10, 64)
	next1, _ := strconv.ParseUint(msg.Next1, 10, 64)

	numerator1, _ := sdk.NewIntFromString(msg.Rate1[0])
	denominator1, _ := sdk.NewIntFromString(msg.Rate1[1])
	rate1 := []sdk.Int{numerator1, denominator1}

	prev2, _ := strconv.ParseUint(msg.Prev2, 10, 64)
	next2, _ := strconv.ParseUint(msg.Next2, 10, 64)

	numerator2, _ := sdk.NewIntFromString(msg.Rate1[0])
	denominator2, _ := sdk.NewIntFromString(msg.Rate1[1])
	rate2 := []sdk.Int{numerator2, denominator2}

	// Case 1
	// Only drop in pool
	if prev1 == 0 && next1 == 0 && prev2 == 0 && next2 == 0 {
		if member1.Protect != 0 {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Member 1 protect field not 0")
		}

		if member2.Protect != 0 {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Member 2 protect field not 0")
		}

		member1.Protect = uid
		member2.Protect = uid

	}

	// Drop Protection Book modelled as a Stop - Decreasing exchange rates
	// Justification: Exchange rate decreases while trades are executed

	// Case 2 Side 1
	// New head of the book
	if prev1 == 0 && next1 > 0 {
		nextDrop1, _ := k.GetDrop(ctx, next1)
		if !nextDrop1.Active {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Next1 drop not active")
		}
		if nextDrop1.Prev1 != 0 {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Next1 drop not currently head of book")
		}

		if types.LTE(rate1, nextDrop1.Rate1) {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Drop Rate1 less than or equal Next1")
		}

		// Set drop as new head of Member1 Protect
		member1.Protect = uid

		// Set nextOrder prev field to order
		nextDrop1.Prev1 = uid

		k.SetDrop(
			ctx,
			nextDrop1,
		)
	}

	// Case 2 Side 2
	// New head of the book
	if prev2 == 0 && next2 > 0 {
		nextDrop2, _ := k.GetDrop(ctx, next2)
		if !nextDrop2.Active {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Next2 drop not active")
		}
		if nextDrop2.Prev1 != 0 {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Next2 drop not currently head of book")
		}

		if types.LTE(rate2, nextDrop2.Rate2) {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Drop Rate2 less than or equal Next1")
		}

		// Set drop as new head of Member1 Protect
		member2.Protect = uid

		// Set nextOrder prev field to order
		nextDrop2.Prev2 = uid

		k.SetDrop(
			ctx,
			nextDrop2,
		)
	}

	// Case 3 Side 1
	// New tail of book
	if prev1 > 0 && next1 == 0 {

		prevDrop1, _ := k.GetDrop(ctx, prev1)

		if !prevDrop1.Active {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Prev1 drop not active")
		}
		if prevDrop1.Next1 != 0 {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Prev1 drop not currently tail of book")
		}

		if types.GT(rate1, prevDrop1.Rate1) {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Order rate greater than Prev")
		}

		// Set nextDrop1 Next1 field to Drop UID
		prevDrop1.Next1 = uid

		k.SetDrop(
			ctx,
			prevDrop1,
		)
	}

	// Case 3 Side 2
	// New tail of book
	if prev2 > 0 && next2 == 0 {

		prevDrop2, _ := k.GetDrop(ctx, prev2)

		if !prevDrop2.Active {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Prev2 drop not active")
		}
		if prevDrop2.Next2 != 0 {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Prev2 drop not currently tail of book")
		}

		if types.GT(rate1, prevDrop2.Rate2) {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Drop rate2 greater than Prev2 rate2")
		}

		// Set nextDrop1 Next1 field to Drop UID
		prevDrop2.Next2 = uid

		k.SetDrop(
			ctx,
			prevDrop2,
		)
	}

	// Case 4 Side 1
	// IF next position and prev position are stated
	if prev1 > 0 && next1 > 0 {

		prevDrop1, _ := k.GetDrop(ctx, prev1)
		nextDrop1, _ := k.GetDrop(ctx, next1)

		if !prevDrop1.Active {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Prev1 drop not active")
		}
		if !nextDrop1.Active {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Next1 drop not active")
		}

		if !(nextDrop1.Prev1 == prevDrop1.Uid && prevDrop1.Next1 == nextDrop1.Uid) {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Prev1 and Next1 are not adjacent")
		}

		k.SetDrop(
			ctx,
			prevDrop1,
		)

		k.SetDrop(
			ctx,
			nextDrop1,
		)
	}

	// Case 4 Side 2
	// IF next position and prev position are stated
	if prev2 > 0 && next2 > 0 {

		prevDrop2, _ := k.GetDrop(ctx, prev2)
		nextDrop2, _ := k.GetDrop(ctx, next2)

		if !prevDrop2.Active {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Prev2 drop not active")
		}
		if !nextDrop2.Active {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Next2 drop not active")
		}

		if !(nextDrop2.Prev2 == prevDrop2.Uid && prevDrop2.Next1 == nextDrop2.Uid) {
			return nil, sdkerrors.Wrapf(types.ErrInvalidOrder, "Prev2 and Next2 are not adjacent")
		}

		k.SetDrop(
			ctx,
			prevDrop2,
		)

		k.SetDrop(
			ctx,
			nextDrop2,
		)
	}

	// The Pool Sum current is defined as:
	// poolSum == AMM A Coin Balance + AMM B Coin Balance
	poolSum := member1.Balance.Add(member2.Balance)

	drops, _ := sdk.NewIntFromString(msg.Drops)

	// The beginning Drop Sum is defined as:
	// dropSum == Total amount of coinA+coinB needed to create the drop based on pool exchange rate
	dropSum := ((drops.Mul(poolSum).Mul(sdk.NewInt(10 ^ 18))).Quo(pool.Drops)).Quo(sdk.NewInt(10 ^ 18))

	// dropSum == A + B
	// dropSum = B + B * exchrate(A/B)
	// dropSum = B * (1 + exchrate(A/B))
	// B = dropSum / (1 + exchrate(A/B))
	// 1 + exchrate(A/B) = 1 + AMM A Balance / AMM B Balance
	// = AMM B Balance / AMM B Balance + AMM A Balance / AMM B Balance
	// = (AMM B Balance + AMM A Balance)/AMM B Balance
	// B = dropSum / [(AMM B Balance + AMM A Balance)/AMM B Balance]
	amount1 := dropSum.Mul(sdk.NewInt(10 ^ 18)).Quo((poolSum.Mul(sdk.NewInt(10 ^ 18))).Quo(member2.Balance))

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

	member2.Balance = member2.Balance.Add(amount2)
	k.SetMember(ctx, member2)

	// Get Drop Creator and Pool Leader total drops from all drops owned
	sumDropsCreator := k.GetOwnerDropsInt(ctx, msg.Creator).Add(drops)
	sumDropsLeader := k.GetOwnerDropsInt(ctx, pool.Leader)

	// If Creator totaled owned drops is greater than Leader then
	// Creator is new leader
	if sumDropsCreator.GT(sumDropsLeader) {
		pool.Leader = msg.Creator
	}

	pool.Drops = pool.Drops.Add(drops)
	k.SetPool(ctx, pool)

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

	// Update drop uid count
	k.SetUidCount(ctx, uid+1)

	return &types.MsgCreateDropResponse{}, nil
}
