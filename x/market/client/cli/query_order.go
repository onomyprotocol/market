package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/pendulum-labs/market/x/market/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdListOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-order",
		Short: "list all order",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllOrderRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.OrderAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-order [uid]",
		Short: "shows a order",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argUid, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetOrderRequest{
				Uid: argUid,
			}

			res, err := queryClient.Order(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
