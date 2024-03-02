package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	govcli "github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/pendulum-labs/market/x/market/types"
)

type proposalGeneric struct {
	Title       string
	Description string
	Deposit     string
}

type commandGeneric struct {
	MetadataPath string
}

// CmdFundTreasuryProposal implements the command to submit a fund-treasury proposal.
func CmdDenomMetadataProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "denom-metadata",
		Args:  cobra.ExactArgs(0),
		Short: "Submit a denom metadata proposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a denom metadata proposal.
Example:
$ %s tx gov submit-proposal denom-metadata --title="Test Proposal" --description="My awesome proposal" --deposit="10000000000000000000aonex" --from mykey

Must have denom.json in directory containing the denom metadata`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			commandGeneric, err := parseCommandFlags(cmd.Flags())
			if err != nil {
				return err
			}

			path := commandGeneric.MetadataPath

			metadataFile, err := os.Open(path)
			if err != nil {
				return err
			}

			byteMetadata, err := io.ReadAll(metadataFile)
			if err != nil {
				return err
			}

			var metadata banktypes.Metadata

			err = json.Unmarshal(byteMetadata, &metadata)
			if err != nil {
				return err
			}

			err = metadata.Validate()
			if err != nil {
				return err
			}

			proposalGeneric, err := parseSubmitProposalFlags(cmd.Flags())
			if err != nil {
				return err
			}

			deposit, err := sdk.ParseCoinsNormalized(proposalGeneric.Deposit)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()
			content := types.NewDenomMetadataProposal(from, proposalGeneric.Title, proposalGeneric.Description, metadata)

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	addProposalFlags(cmd)

	return cmd
}

func parseSubmitProposalFlags(fs *pflag.FlagSet) (*proposalGeneric, error) {
	title, err := fs.GetString(govcli.FlagTitle)
	if err != nil {
		return nil, err
	}
	description, err := fs.GetString(govcli.FlagDescription)
	if err != nil {
		return nil, err
	}

	deposit, err := fs.GetString(govcli.FlagDeposit)
	if err != nil {
		return nil, err
	}

	return &proposalGeneric{
		Title:       title,
		Description: description,
		Deposit:     deposit,
	}, nil
}

func parseCommandFlags(fs *pflag.FlagSet) (*commandGeneric, error) {
	path, err := fs.GetString("metadata-path")
	if err != nil {
		return nil, err
	}

	return &commandGeneric{
		MetadataPath: path,
	}, nil

}

func addProposalFlags(cmd *cobra.Command) {
	cmd.Flags().String(govcli.FlagTitle, "", "The proposal title")
	cmd.Flags().String(govcli.FlagDescription, "", "The proposal description")
	cmd.Flags().String(govcli.FlagDeposit, "", "The proposal deposit")
}

func addCommandlFlags(cmd *cobra.Command) {
	cmd.Flags().String("metadata-path", "", "The path to metadata json")
}
