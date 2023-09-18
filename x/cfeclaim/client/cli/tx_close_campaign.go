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

func CmdCloseCampaign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "close-campaign [campaign-id]",
		Short: "Close a campaign",
		Long: strings.TrimSpace(fmt.Sprintf(`Close an existing campaign.
Requirements:
- The campaign must be in an active state
- The campaign must be owned by the sender
- The campaign must be over the end time

Arguments:
  [campaign-id]          ID of the campaign to close

Example:
$ %s tx %s close-campaign 1 --from mykey
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

			msg := types.NewMsgCloseCampaign(
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
