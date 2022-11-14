package cli

import (
	"context"

	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListInitialClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-initial-claim",
		Short: "list all initialClaim",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllInitialClaimRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.InitialClaimAll(context.Background(), params)
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

func CmdShowInitialClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-initial-claim [campaign-id]",
		Short: "shows a initialClaim",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argCampaignId := args[0]

			params := &types.QueryGetInitialClaimRequest{
				CampaignId: argCampaignId,
			}

			res, err := queryClient.InitialClaim(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
