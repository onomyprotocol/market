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
		Use:   "show-drop [uid]",
		Short: "shows a drop",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argUid, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryDropRequest{
				Uid: argUid,
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

func CmdShowDropPairs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-drop-pairs [address]",
		Short: "show pairs owner has drops",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argAddr := args[0]

			params := &types.QueryDropPairsRequest{
				Address: argAddr,
			}

			res, err := queryClient.DropPairs(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdDropOwnerPair() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "drop-owner-pair [address] [pair]",
		Short: "shows all drops owned for pair",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			argAddress, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}

			argPair, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryDropOwnerPairRequest{
				Address:    argAddress,
				Pair:       argPair,
				Pagination: pageReq,
			}

			res, err := queryClient.DropOwnerPair(context.Background(), params)
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

func CmdDropCoin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "drop-coin [denomA] [denomB] [amountA]",
		Short: "calculate drops and amountB for given denomA and amountA",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			argDenomA, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}

			argDenomB, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}

			argAmountA, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryDropCoinRequest{
				DenomA:  argDenomA,
				DenomB:  argDenomB,
				AmountA: argAmountA,
			}

			res, err := queryClient.DropCoin(context.Background(), params)
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
