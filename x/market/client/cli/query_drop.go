package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/pendulum-labs/market/x/market/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdListDrop() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-drop",
		Short: "list all drop",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllDropRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.DropAll(context.Background(), params)
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

func CmdShowDrop() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-drop [uid] [owner] [pair]",
		Short: "shows a drop",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argUid, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argOwner := args[1]
			argPair := args[2]

			params := &types.QueryGetDropRequest{
				Uid:   argUid,
				Owner: argOwner,
				Pair:  argPair,
			}

			res, err := queryClient.Drop(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
