package cli

import (
	"strconv"

	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/pendulum-labs/market/x/market/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateDrop() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-drop [pair] [drops]",
		Short: "Broadcast message create-drop",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPair := args[0]
			argDrops := args[1]
			argRate1 := strings.Split(args[2], listSeparator)
			argPrev1 := args[3]
			argNext1 := args[4]
			argRate2 := strings.Split(args[5], listSeparator)
			argPrev2 := args[6]
			argNext2 := args[7]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateDrop(
				clientCtx.GetFromAddress().String(),
				argPair,
				argDrops,
				argRate1,
				argPrev1,
				argNext1,
				argRate2,
				argPrev2,
				argNext2,
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
