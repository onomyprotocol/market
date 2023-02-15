package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/onomyprotocol/market/x/market/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdListAsset() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-asset",
		Short: "list all asset",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllAssetRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AssetAll(context.Background(), params)
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

func CmdShowAsset() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-asset [active] [owner] [asset-type]",
		Short: "shows a asset",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argActive, err := cast.ToBoolE(args[0])
			if err != nil {
				return err
			}
			argOwner := args[1]
			argAssetType := args[2]

			params := &types.QueryGetAssetRequest{
				Active:    argActive,
				Owner:     argOwner,
				AssetType: argAssetType,
			}

			res, err := queryClient.Asset(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
