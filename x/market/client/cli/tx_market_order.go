package cli

import (
	"strconv"

	"market/x/market/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdMarketOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "market-order [denom-ask] [amount-ask] [denom-bid] [amount-bid] [slippage]",
		Short: "Broadcast message market-order",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenomAsk := args[0]
			argAmountAsk := args[1]
			argDenomBid := args[2]
			argAmountBid := args[3]
			argSlippage := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMarketOrder(
				clientCtx.GetFromAddress().String(),
				argDenomAsk,
				argAmountAsk,
				argDenomBid,
				argAmountBid,
				argSlippage,
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
