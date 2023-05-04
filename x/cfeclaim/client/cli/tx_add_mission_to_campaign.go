package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAddMissionToCampaign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-mission-to-campaign [campaign-id] [name] [description] [mission-type] [weight] [optional-claim-start-date]",
		Short: "Broadcast message AddMissionToAidropCampaign",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCampaignId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			argName := args[1]
			argDescription := args[2]

			argMissionType, err := types.MissionTypeFromString(types.NormalizeMissionType(args[3]))
			if err != nil {
				return err
			}
			argWeight, err := sdk.NewDecFromStr(args[4])
			if err != nil {
				return err
			}
			var argClaimStartDate *time.Time
			if args[5] == "" {
				argClaimStartDate = nil
			} else {
				parsedTime, err := time.Parse(timeLayout, args[5])
				if err != nil {
					return err
				}
				argClaimStartDate = &parsedTime
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddMissionToCampaign(
				clientCtx.GetFromAddress().String(),
				argCampaignId,
				argName,
				argDescription,
				argMissionType,
				&argWeight,
				argClaimStartDate,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
