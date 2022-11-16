package keeper

import (
	"context"
	"sort"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/onomyprotocol/market/x/market/types"
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

	pool.Drops = pool.Drops.Add(drops)
	k.SetPool(ctx, pool)

	// Create the uid
	count := k.GetUidCount(ctx)

	var drop = types.Drop{
		Uid:    count,
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
	k.SetUidCount(ctx, count+1)

	return &types.MsgCreateDropResponse{}, nil
}

func (k msgServer) validateSenderBalance(ctx sdk.Context, senderAddress sdk.AccAddress, coins sdk.Coins) error {
	for _, coin := range coins {
		balance := k.Keeper.bankKeeper.GetBalance(ctx, senderAddress, coin.Denom)
		if balance.IsLT(coin) {
			return sdkerrors.Wrapf(
				types.ErrInsufficientBalance, "%s is smaller than %s", balance, coin)
		}
	}

	return nil
}
