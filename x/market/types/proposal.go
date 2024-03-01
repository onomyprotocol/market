package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	// ProposalTypeDenomMetadataProposal defines the type for a DenomMetadataProposal.
	ProposalTypeDenomMetadataProposal = "DenomMetadataProposal"
)

var (
	_ govtypes.Content = &DenomMetadataProposal{}
)

func init() { // nolint:gochecknoinits // cosmos sdk style
	govtypes.RegisterProposalType(ProposalTypeDenomMetadataProposal)
	govtypes.RegisterProposalTypeCodec(&DenomMetadataProposal{}, fmt.Sprintf("%s/%s", ModuleName, ProposalTypeDenomMetadataProposal))
}

// NewDenomMetadataProposal creates a new fund treasury proposal.
func NewDenomMetadataProposal(sender sdk.AccAddress, title string, description string, metadata banktypes.Metadata, rate []sdk.Uint) *DenomMetadataProposal {
	return &DenomMetadataProposal{sender.String(), title, description, &metadata, rate}
}

// GetTitle returns the title of a fund treasury proposal.
func (m *DenomMetadataProposal) GetTitle() string { return m.Title }

// GetDescription returns the description of a fund treasury proposal.
func (m *DenomMetadataProposal) GetDescription() string { return m.Description }

// ProposalRoute returns the routing key of a fund treasury proposal.
func (m *DenomMetadataProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of the fund treasury proposal.
func (m *DenomMetadataProposal) ProposalType() string { return ProposalTypeDenomMetadataProposal }

// ValidateBasic runs basic stateless validity checks.
func (m *DenomMetadataProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(m)
	if err != nil {
		return err
	}
	sender, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}
	if err := sdk.VerifyAddressFormat(sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address: %s", err)
	}

	return nil
}

// GetProposer returns the proposer from the proposal struct.
func (m *DenomMetadataProposal) GetProposer() string { return m.Sender }
