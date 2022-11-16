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
		return nil, sdkerrors.Wrapf(types.ErrPoolDoesNotExist, "%s", pair)
	}

	member1, found := k.GetMember(ctx, denom2, denom1)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrPoolDoesNotExist, "%s", pair)
	}

	member2, found := k.GetMember(ctx, denom1, denom2)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrPoolDoesNotExist, "%s", pair)
	}

	// The Pool Sum current is defined as:
	// poolSum == AMM A Coin Balance + AMM B Coin Balance
	poolSum := member1.Balance.Add(member2.Balance)

	// The beginning Drop Sum is defined as:
	// dropSum == Total amount of coinA+coinB needed to create the drop based on pool exchange rate

	_ = poolSum
	_ = pool
	_ = ctx

	return &types.MsgCreateDropResponse{}, nil
}
