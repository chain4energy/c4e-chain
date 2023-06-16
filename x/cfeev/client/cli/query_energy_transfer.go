package cli

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"strconv"
	"strings"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdAllEnergyTransfers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-energy-transfers",
		Short: "Query all existing energy transfers",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllEnergyTransfersRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AllEnergyTransfers(context.Background(), params)
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

func CmdEnergyTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "energy-transfer [id]",
		Short: "Query an energy transfer by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryEnergyTransferRequest{
				Id: id,
			}

			res, err := queryClient.EnergyTransfer(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdEnergyTransfers() *cobra.Command {
	bech32PrefixAddr := sdk.GetConfig().GetBech32AccountAddrPrefix()
	cmd := &cobra.Command{
		Use:   "energy-transfers [owner-acc-address]",
		Short: "Query for all energy transfers of a given CP owner address",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query for all energy transfers of a given CP owner address.

Arguments:
  [ownerAccAddress] CP owner account address

Example:
$ %s query %s energy-transfers %se1mydzr5gxtypyjks08nveclwjmjp64trxh4laxj
`,
				version.AppName, types.ModuleName, bech32PrefixAddr,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqOwner := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryEnergyTransfersRequest{
				Owner: reqOwner,
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			params.Pagination = pageReq

			res, err := queryClient.EnergyTransfers(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
