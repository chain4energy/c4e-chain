package cli

import (
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCloseAirdropCampaign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "close-airdrop-campaign [campaign-id] [burn] [community-pool-send]",
		Short: "Broadcast message CloseAirdropCampaign",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCampaignId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argBurn, err := cast.ToBoolE(args[1])
			if err != nil {
				return err
			}
			argCommunityPoolSend, err := cast.ToBoolE(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCloseAirdropCampaign(
				clientCtx.GetFromAddress().String(),
				argCampaignId,
				argBurn,
				argCommunityPoolSend,
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
