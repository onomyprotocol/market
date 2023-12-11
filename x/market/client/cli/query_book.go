package cli

import (
	"strconv"

	"market/x/market/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdBook() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "book [denom-a] [denom-b] [order-type]",
		Short: "Query book",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenomA := args[0]
			reqDenomB := args[1]
			reqOrderType := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryBookRequest{
				DenomA:    reqDenomA,
				DenomB:    reqDenomB,
				OrderType: reqOrderType,
			}

			res, err := queryClient.Book(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
