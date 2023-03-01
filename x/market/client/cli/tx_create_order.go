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

func CmdCreateOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-order [denom-ask] [denom-bid] [order-type] [amount] [rate] [prev] [next]",
		Short: "Broadcast message create-order",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenomAsk := args[0]
			argDenomBid := args[1]
			argOrderType := args[2]
			argAmount := args[3]
			argRate := strings.Split(args[4], listSeparator)
			argPrev := args[5]
			argNext := args[6]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateOrder(
				clientCtx.GetFromAddress().String(),
				argDenomAsk,
				argDenomBid,
				argOrderType,
				argAmount,
				argRate,
				argPrev,
				argNext,
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
