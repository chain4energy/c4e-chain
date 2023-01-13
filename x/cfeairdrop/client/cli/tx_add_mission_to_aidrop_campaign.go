package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAddMissionToAidropCampaign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-mission-to-aidrop-campaign [name] [campaign-id] [description] [mission-type] [weight]",
		Short: "Broadcast message AddMissionToAidropCampaign",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argCampaignId := args[1]
			argDescription := args[2]
			argMissionType := args[3]
			argWeight := args[4]

			campaignId, err := strconv.ParseUint(argCampaignId, 10, 64)
			if err != nil {
				return err
			}

			weight, err := sdk.NewDecFromStr(argWeight)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddMissionToAidropCampaign(
				clientCtx.GetFromAddress().String(),
				argName,
				campaignId,
				argDescription,
				argMissionType,
				weight,
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
