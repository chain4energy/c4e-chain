package cli

import (
	"fmt"
	"github.com/chain4energy/c4e-chain/app/params"
	"github.com/cosmos/cosmos-sdk/version"
	"strconv"
	"strings"

	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim [campaign-id] [mission-id]",
		Short: "Claim from a campaign",
		Args:  cobra.ExactArgs(2),
		Long: strings.TrimSpace(fmt.Sprintf(`Claim rewards from a campaign.

Arguments:
  [campaign-id] ID of the campaign to claim rewards from
  [mission-id]  ID of the mission to claim rewards for

Example:
$ %s tx %s claim 1 2 --from mykey
`, version.AppName, types.ModuleName)),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCampaignId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			argMissionId, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgClaim(
				clientCtx.GetFromAddress().String(),
				argCampaignId,
				argMissionId,
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

func CmdInitialClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "initial-claim [campaign-id] [destination-address]",
		Short: "Initial claim from a campaign",
		Long: strings.TrimSpace(fmt.Sprintf(`Make an initial claim from a campaign.

Arguments:
  [campaign-id]               ID of the campaign to make the initial claim from
  [destination-address] 	  Destination address to claim the rewards to. This address will be used as a claming address for another missions of the campaign.

Example:
$ %s tx %s initial-claim 1 %s1gghjut3ccd8ay0zduzj64hwre2fxs9ldmqhffj --from mykey
`, version.AppName, types.ModuleName, params.Bech32PrefixAccAddr)),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCampaignId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			argDestinationAddress := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgInitialClaim(
				clientCtx.GetFromAddress().String(),
				argCampaignId,
				argDestinationAddress,
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
