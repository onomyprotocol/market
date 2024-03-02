package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/pendulum-labs/market/x/market/types"
)

// FundTreasuryProposal submits the FundTreasuryProposal.
func (k Keeper) DenomMetadataProposal(ctx sdk.Context, request *types.DenomMetadataProposal) error {

	_, exists := k.bankKeeper.GetDenomMetaData(ctx, request.Metadata.Base)

	if exists {
		return sdkerrors.Wrapf(types.ErrDenomExists, "%s", request.Metadata.Base)
	}

	k.bankKeeper.SetDenomMetaData(ctx, *request.Metadata)

	return nil
}
