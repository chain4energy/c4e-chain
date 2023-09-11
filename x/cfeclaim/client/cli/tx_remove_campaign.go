package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/version"
	"strconv"
	"strings"

	"github.com/chain4energy/c4e-chain/v2/x/cfeclaim/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRemoveCampaign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-campaign [campaign-id]",
		Short: "Remove existing campaign (campaign cannot be enabled)",
		Long: strings.TrimSpace(fmt.Sprintf(`Remove an existing campaign.
You can only remove campaigns that are not enabled.

Arguments:
  [campaign-id]    ID of the campaign to remove

Example:
$ %s tx %s remove-campaign 1 --from mykey
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCampaignId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveCampaign(
				clientCtx.GetFromAddress().String(),
				argCampaignId,
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
