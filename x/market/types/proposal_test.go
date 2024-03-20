package types

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// GenAccountAddress generates random account.
func GenAccountAddress() sdk.AccAddress {
	pk := ed25519.GenPrivKey().PubKey()
	return sdk.AccAddress(pk.Address())
}

func TestDenomMetadataProposal_ValidateBasic(t *testing.T) { //nolint:dupl // test template

	type fields struct {
		Sender      string
		Title       string
		Description string
		Metadata    banktypes.Metadata
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{

		{
			name: "positive",
			fields: fields{
				Sender:      GenAccountAddress().String(),
				Title:       "title",
				Description: "desc",
				Metadata: banktypes.Metadata{
					Name:        "Cosmos Hub Atom",
					Symbol:      "ATOM",
					Description: "The native staking token of the Cosmos Hub.",
					DenomUnits: []*banktypes.DenomUnit{
						{Denom: "uatom", Exponent: uint32(0), Aliases: []string{"microatom"}},
						{Denom: "matom", Exponent: uint32(3), Aliases: []string{"milliatom"}},
						{Denom: "atom", Exponent: uint32(6), Aliases: nil},
					},
					Base:    "uatom",
					Display: "atom",
				},
			},
		},
		{
			name: "negative_invalid_sender",
			fields: fields{
				Sender:      "invalid-sender",
				Title:       "title",
				Description: "desc",
				Metadata: banktypes.Metadata{
					Name:        "Cosmos Hub Atom",
					Symbol:      "ATOM",
					Description: "The native staking token of the Cosmos Hub.",
					DenomUnits: []*banktypes.DenomUnit{
						{Denom: "uatom", Exponent: uint32(0), Aliases: []string{"microatom"}},
						{Denom: "matom", Exponent: uint32(3), Aliases: []string{"milliatom"}},
						{Denom: "atom", Exponent: uint32(6), Aliases: nil},
					},
					Base:    "uatom",
					Display: "atom",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &DenomMetadataProposal{
				Sender:      tt.fields.Sender,
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Metadata:    &tt.fields.Metadata,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
