package cli

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"strconv"
	"strings"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAddMission() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-mission [campaign-id] [name] [description] [mission-type] [weight] [optional-claim-start-date]",
		Short: "Add new mission to a campaign",
		Long: strings.TrimSpace(fmt.Sprintf(`Add a new mission to a campaign.
Requirements:
- The campaign can't be enabled
- The campaign can't be closed
- The sum of all missions weight must be lower or equal to 1

Arguments:
  [campaign-id]                ID of the campaign to add the mission to
  [name]                       Name of the mission
  [description]                Description of the mission
  [mission-type]               Type of the mission (delegate/vote/claim)
  [weight]                     Weight of the mission (must be privided as a decimal string)
  [optional-claim-start-date]  Optional claim start date for the mission

Example:
$ %s tx %s add-mission 1 "Mission Name" "Mission Description" "delegate" "0.5"" "2006-01-02 15:04:05 -0700 MST" --from mykey
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(6),
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
			argClaimStartDate, err := parseOptionalTime(args[5])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddMission(
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

func parseOptionalTime(optionalTimeString string) (*time.Time, error) {
	if optionalTimeString == "" {
		return nil, nil
	}
	parsedTime, err := time.Parse(TimeLayout, optionalTimeString)
	if err != nil {
		return nil, err
	}
	return &parsedTime, nil

}
