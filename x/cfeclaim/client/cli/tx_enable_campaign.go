package cli

import (
	"fmt"
	"github.com/chain4energy/c4e-chain/v2/x/cfeclaim/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var _ = strconv.Itoa(0)

func CmdEnableCampaign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enable-campaign [campaign-id] [optional-start-time] [optional-end-time]",
		Short: "Enable existing campaign",
		Long: strings.TrimSpace(fmt.Sprintf(`Enable an existing campaign.
Remember that after changing the campaign status to enabled, you won't be able to remove or close the campaign until its end time.

Arguments:
  [campaign-id]            ID of the campaign to enable
  [optional-start-time]    Optional start time of the campaign in the format "2006-01-02T15:04:05Z"
  [optional-end-time]      Optional end time of the campaign in the format "2006-01-02T15:04:05Z"

Example:
$ %s tx %s enable-campaign 1 "2006-01-02 15:04:05 -0700 MST" "2006-01-02 15:04:05 -0700 MST" --from mykey
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCampaignId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			argStartTime, err := parseOptionalTime(args[1])
			if err != nil {
				return err
			}

			argEndTime, err := parseOptionalTime(args[2])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgEnableCampaign(
				clientCtx.GetFromAddress().String(),
				argCampaignId,
				argStartTime,
				argEndTime,
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
