package cli

import (
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdAddClaimRecords() *cobra.Command { //  TODO opis ja w innych modulach
	cmd := &cobra.Command{
		Use:   "add-claim-records [campaignId] [claim-entries-json-file]",
		Short: "Add new claim records to the campaign",
		Example: `Example claim-entries-json-file
[
	{
		"address": "c4e128dcl5738ffy08kxxgcxyj6sp7zpyu32n4u32k",
		"amount":[
			{
				"denom":"uc4e",
				"amount":"11250000000"
			}
		]
	}
]
`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCampaignId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			argClaimEntries, err := parseClaimEntries(argCampaignId, args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgAddClaimRecords(clientCtx.GetFromAddress().String(), argCampaignId, argClaimEntries)
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
