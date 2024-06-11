package cli

import (
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCampaigns() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "campaigns",
		Short: "Query all existing campaigns",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryCampaignsRequest{}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			params.Pagination = pageReq

			res, err := queryClient.Campaigns(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdCampaign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "campaign [campaign-id]",
		Short: "Query a specific campaign",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			argCampaignId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			params := &types.QueryCampaignRequest{CampaignId: argCampaignId}

			res, err := queryClient.Campaign(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
