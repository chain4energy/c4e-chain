package cli

import (
	"fmt"
	"github.com/chain4energy/c4e-chain/v2/app/params"
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

func CmdAddClaimRecords() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-claim-records [campaignId] [claim-entries-json-file]",
		Short: "Add new claim records to the campaign",
		Long: strings.TrimSpace(fmt.Sprintf(`Add new claim records to a campaign.
The command takes a campaign ID and a JSON file containing the claim entries to be added.

Example claim-entries-json-file:
[
	{
		"user_entry_address": "c4e128dcl5738ffy08kxxgcxyj6sp7zpyu32n4u32k",
		"amount": [
			{
				"denom": "uc4e",
				"amount": "11250000000"
			}
		]
	}
]

Arguments:
  [campaignId]               ID of the campaign to add claim records to
  [claim-entries-json-file]  Path to a JSON file containing the claim entries

Example:
$ %s tx %s add-claim-records 1 claim_entries.json --from mykey
`, version.AppName, types.ModuleName)),
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
		Long: strings.TrimSpace(fmt.Sprintf(`Delete a claim record by ID.
This command allows you to delete a specific claim record from a campaign by providing the campaign ID and the user address associated with the claim record.

If a campaign has removableClaimRecords flag set to true you can remove claim records after campaign is enabled.
If a campaign has removableClaimRecords flag set to false you can remove claim records only before campaign is enabled.

Arguments:
  [campaign-id]    ID of the campaign that contains the claim record
  [user-address]   Address of the user associated with the claim record

Example:
$ %s tx %s delete-claim-record 1 %s1gghjut3ccd8ay0zduzj64hwre2fxs9ldmqhffj --from mykey
`, version.AppName, types.ModuleName, params.Bech32PrefixAccAddr)),
		Args: cobra.ExactArgs(2),
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
