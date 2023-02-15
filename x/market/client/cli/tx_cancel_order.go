package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/onomyprotocol/market/x/market/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCancelOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel-order [uid]",
		Short: "Broadcast message cancel-order",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argUid := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCancelOrder(
				clientCtx.GetFromAddress().String(),
				argUid,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
