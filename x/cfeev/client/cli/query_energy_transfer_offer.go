package cli

import (
	"context"
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListEnergyTransferOffer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-energy-transfer-offer",
		Short: "list all energyTransferOffer",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllEnergyTransferOfferRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.EnergyTransferOfferAll(context.Background(), params)
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

func CmdShowEnergyTransferOffer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-energy-transfer-offer [id]",
		Short: "shows a energyTransferOffer",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetEnergyTransferOfferRequest{
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
