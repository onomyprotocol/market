package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/pendulum-labs/market/x/market/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdMarketOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "market-order [denom-ask] [denom-bid] [amount-bid] [slippage]",
		Short: "Broadcast message market-order",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenomAsk := args[0]
			argDenomBid := args[1]
			argAmountBid := args[2]
			argSlippage := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMarketOrder(
				clientCtx.GetFromAddress().String(),
				argDenomAsk,
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
