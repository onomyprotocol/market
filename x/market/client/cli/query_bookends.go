package cli

import (
	"strconv"
	"strings"

	"market/x/market/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdBookends() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bookends [coin-a] [coin-b] [order-type] [rate]",
		Short: "Query bookends",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqCoinA := args[0]
			reqCoinB := args[1]
			reqOrderType := args[2]
			reqRate := strings.Split(args[3], listSeparator)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryBookendsRequest{

				CoinA:     reqCoinA,
				CoinB:     reqCoinB,
				OrderType: reqOrderType,
				Rate:      reqRate,
			}

			res, err := queryClient.Bookends(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
