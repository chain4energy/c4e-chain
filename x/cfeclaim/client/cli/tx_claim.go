package cli

import (
	"strconv"

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
		Short: "Broadcast message Claim",
		Args:  cobra.ExactArgs(2),
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
		Use:   "initial-claim [campaign-id] [optional-address-to-claim]",
		Short: "Broadcast message InitialClaim",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCampaignId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			argAddressToClaim := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgInitialClaim(
				clientCtx.GetFromAddress().String(),
				argCampaignId,
				argAddressToClaim,
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
