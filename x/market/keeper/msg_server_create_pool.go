package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/onomyprotocol/market/x/market/types"
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

	// NewCoins sorts denoms
	denom1 := coinPair.GetDenomByIndex(1)
	denom2 := coinPair.GetDenomByIndex(2)

	pair := strings.Join([]string{denom1, denom2}, ",")

	_, found := k.GetPool(ctx, pair)
	if found {
		return nil, sdkerrors.Wrapf(types.ErrPoolAlreadyExists, "%s", pair)
	}

	// moduleAcc := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	// Get the borrower address
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)

	// Use the module account as pool account
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, coinPair)
	if sdkError != nil {
		return nil, sdkError
	}

	drops := coinPair.AmountOf(denom1).Add(coinPair.AmountOf(denom2)).String()

	// Create a new Pool with the following user input
	var pool = types.Pool{
		Pair:     pair,
		Leader:   msg.Creator,
		Denom1:   coinPair.GetDenomByIndex(1),
		Denom2:   coinPair.GetDenomByIndex(2),
		Drops:    drops,
		Earnings: "0",
		Burnings: "0",
	}

	var member1 = types.Member{
		Pair:    pair,
		DenomA:  denom2,
		DenomB:  denom1,
		Balance: coinPair.AmountOf(denom1).String(),
		Limit:   0,
		Stop:    0,
	}

	var member2 = types.Member{
		Pair:    pair,
		DenomA:  denom1,
		DenomB:  denom2,
		Balance: coinPair.AmountOf(denom2).String(),
		Limit:   0,
		Stop:    0,
	}

	// Create the uid
	count := k.GetUidCount(ctx)

	var drop = types.Drop{
		Uid:    count,
		Owner:  msg.Creator,
		Pair:   pair,
		Drops:  drops,
		Sum:    drops,
		Active: true,
	}

	// Add the loan to the keeper
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

	// Update drop uid count
	k.SetUidCount(ctx, count+1)

	return &types.MsgCreatePoolResponse{}, nil
}
