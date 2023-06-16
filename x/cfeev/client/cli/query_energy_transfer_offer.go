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

func CmdAllEnergyTransfeOffers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-energy-transfe-offers",
		Short: "Query all existing energy transfers offers",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllEnergyTransferOffersRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AllEnergyTransferOffers(context.Background(), params)
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

func CmdEnergyTransferOffer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "energy-transfer-offer [id]",
		Short: "Query an energy transfer offer by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryEnergyTransferOfferRequest{
				Id: id,
			}

			res, err := queryClient.EnergyTransferOffer(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdEnergyTransferOffers() *cobra.Command {
	bech32PrefixAddr := sdk.GetConfig().GetBech32AccountAddrPrefix()
	cmd := &cobra.Command{
		Use:   "energy-transfer-offers [owner]",
		Short: "Query for all energy transfer offers of a given CP owner address",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query for all energy transfer offers of a given CP owner address.

Arguments:
  [owner] CP owner account address

Example:
$ %s query %s energy-transfer-offers %se1mydzr5gxtypyjks08nveclwjmjp64trxh4laxj
`,
				version.AppName, types.ModuleName, bech32PrefixAddr,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqOwnerAccAddress := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryEnergyTransferOffersRequest{
				Owner: reqOwnerAccAddress,
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			params.Pagination = pageReq

			res, err := queryClient.EnergyTransferOffers(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
