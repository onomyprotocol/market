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
	coinAmsg, err := sdk.ParseCoinNormalized(msg.CoinA)
	if err != nil {
		panic(err)
	}

	coinBmsg, err := sdk.ParseCoinNormalized(msg.CoinB)
	if err != nil {
		panic(err)
	}

	coinPair := sdk.NewCoins(coinAmsg, coinBmsg)

	// CoinA and CoinB after NewCoins sorting
	denom1 := coinPair.GetDenomByIndex(1)
	denom2 := coinPair.GetDenomByIndex(2)

	pair := strings.Join([]string{denom1, denom2}, ",")

	_, found := k.GetPool(ctx, pair, denom1, denom2, msg.Creator)
	if found {
		return nil, sdkerrors.Wrapf(types.ErrPoolAlreadyExists, "%s", pair)
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

	// moduleAcc := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	// Get the borrower address
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)

	// Use the module account as pool account
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, coinPair)
	if sdkError != nil {
		return nil, sdkError
	}

	// Add the loan to the keeper
	k.SetPool(
		ctx,
		pool,
	)

	return &types.MsgCreatePoolResponse{}, nil
}
