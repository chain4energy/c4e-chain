package cli

import (
	"context"
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdMissions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "missions",
		Short: "Query all existing missions",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryMissionsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.Missions(context.Background(), params)
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

func CmdMission() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mission [campaign-id] [mission-id]",
		Short: "Query a specific mission",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argCampaignId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			argMissionId, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryMissionRequest{
				CampaignId: argCampaignId,
				MissionId:  argMissionId,
			}

			res, err := queryClient.Mission(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdCampaignMissions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "missions [campaign-id]",
		Short: "Query all existing missions of a campaign",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argCampaignId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			params := &types.QueryCampaignMissionsRequest{
				CampaignId: argCampaignId,
			}

			res, err := queryClient.CampaignMissions(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
