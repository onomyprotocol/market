package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/pendulum-labs/market/x/market/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group market queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListPool())
	cmd.AddCommand(CmdPool())
	cmd.AddCommand(CmdListDrop())
	cmd.AddCommand(CmdShowDrop())
	cmd.AddCommand(CmdShowDropPairs())
	cmd.AddCommand(CmdListMember())
	cmd.AddCommand(CmdShowMember())
	cmd.AddCommand(CmdListBurnings())
	cmd.AddCommand(CmdShowBurnings())
	cmd.AddCommand(CmdListOrder())
	cmd.AddCommand(CmdShowOrder())
	cmd.AddCommand(CmdBook())
	cmd.AddCommand(CmdBookends())
	cmd.AddCommand(CmdHistory())
	cmd.AddCommand(CmdOrderOwner())
	cmd.AddCommand(CmdOrderOwnerUids())

	// this line is used by starport scaffolding # 1

	return cmd
}
