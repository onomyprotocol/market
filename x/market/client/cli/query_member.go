package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/onomyprotocol/market/x/market/types"
	"github.com/spf13/cobra"
)

func CmdListMember() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-member",
		Short: "list all member",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllMemberRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.MemberAll(context.Background(), params)
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

func CmdShowMember() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-member [pair] [denom-a] [denom-b]",
		Short: "shows a member",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argPair := args[0]
			argDenomA := args[1]
			argDenomB := args[2]

			params := &types.QueryGetMemberRequest{
				Pair:   argPair,
				DenomA: argDenomA,
				DenomB: argDenomB,
			}

			res, err := queryClient.Member(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
