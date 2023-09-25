package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pendulum-labs/market/x/market/types"
)

func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// CoinAmsg and CoinBmsg pre-sort from raw msg
	coinA, err := sdk.ParseCoinNormalized(msg.CoinA)
	if err != nil {
		panic(err)
	}

	coinB, err := sdk.ParseCoinNormalized(msg.CoinB)
	if err != nil {
		panic(err)
	}

	coinPair := sdk.NewCoins(coinA, coinB)

	// NewCoins sorts denoms.
	// The sorted pair joined by "," is used as the key for the pool.
	denom1 := coinPair.GetDenomByIndex(0)
	denom2 := coinPair.GetDenomByIndex(1)
	pair := strings.Join([]string{denom1, denom2}, ",")

	// Test if pool either exists and active or exists and inactive
	// Inactive pool will be dry or have no drops
	member1, _ := k.GetMember(ctx, denom2, denom1)

	member2, _ := k.GetMember(ctx, denom1, denom2)

	pool, found := k.GetPool(ctx, pair)
	if found {
		if !member1.Balance.Equal(sdk.ZeroInt()) {
			return nil, sdkerrors.Wrapf(types.ErrPoolAlreadyExists, "%s", pair)
		}

		if !member2.Balance.Equal(sdk.ZeroInt()) {
			return nil, sdkerrors.Wrapf(types.ErrPoolAlreadyExists, "%s", pair)
		}

		if !pool.Drops.Equal(sdk.ZeroInt()) {
			return nil, sdkerrors.Wrapf(types.ErrPoolAlreadyExists, "%s", pair)
		}
	}

	// Get the borrower address
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)

	// All coins added to pools are deposited into the module account until redemption
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, coinPair)
	if sdkError != nil {
		return nil, sdkError
	}

	// Drops define proportional ownership to the liquidity in the pool
	drops := coinPair.AmountOf(denom1).Mul(coinPair.AmountOf(denom2))

	leader := types.Leader{
		Address: msg.Creator,
		Drops:   drops,
	}

	if found {
		pool.Drops = drops
		pool.Leaders = []*types.Leader{&leader}
		member1.Balance = coinPair.AmountOf(denom1)
		member2.Balance = coinPair.AmountOf(denom2)
	} else {
		pool = types.Pool{
			Pair:    pair,
			Leaders: []*types.Leader{&leader},
			Denom1:  coinPair.GetDenomByIndex(0),
			Denom2:  coinPair.GetDenomByIndex(1),
			Drops:   drops,
			History: uint64(0),
		}

		member1 = types.Member{
			Pair:    pair,
			DenomA:  denom2,
			DenomB:  denom1,
			Balance: coinPair.AmountOf(denom1),
			Limit:   0,
			Stop:    0,
		}

		member2 = types.Member{
			Pair:    pair,
			DenomA:  denom1,
			DenomB:  denom2,
			Balance: coinPair.AmountOf(denom2),
			Limit:   0,
			Stop:    0,
		}
	}

	// Create the uid
	count := k.GetUidCount(ctx)

	var drop = types.Drop{
		Uid:     count,
		Owner:   msg.Creator,
		Pair:    pair,
		Drops:   drops,
		Product: drops,
		Active:  true,
	}

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
	k.SetUidCount(ctx, count+1)

	return &types.MsgCreatePoolResponse{}, nil
}
