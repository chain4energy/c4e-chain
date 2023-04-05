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

func CmdAddClaimRecords() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-claim-records [campaignId] [airdrop-entries-json-file]",
		Short: "Create a new ClaimRecord",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCampaignId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			argAirdropEntries, err := parseAirdropEntries(clientCtx, argCampaignId, args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgAddClaimRecords(clientCtx.GetFromAddress().String(), argCampaignId, argAirdropEntries)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteClaimRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-claim-record [campaign-id] [user-address]",
		Short: "Delete a ClaimRecord by id",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			campaignId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			userAddress := args[1]
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteClaimRecord(clientCtx.GetFromAddress().String(), campaignId, userAddress)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
