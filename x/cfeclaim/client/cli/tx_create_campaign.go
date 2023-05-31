package cli

import (
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/version"
	"strconv"
	"strings"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateCampaign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-campaign [name] [description] [campaign-type] [removable-claim-records] [feegrant-amount] [initial_claim_free_amount] [free] [start-time] [end-time] [lockup-period] [vesting-period] [optional-vesting-pool-name]",
		Short: "Create new campaign",
		Long: strings.TrimSpace(fmt.Sprintf(`Create a new campaign.

Arguments:
  [name]                       Name of the campaign
  [description]                Description of the campaign
  [campaign-type]              Type of the campaign (default/vesting_pool)
  [removable-claim-records]    Indicates if claim records can be removed after campaign is enabled
  [feegrant-amount]            Amount of feegrant allocated for the campaign
  [initial_claim_free_amount]  Initial claim free amount for the campaign
  [free]                       Free ratio for the campaign
  [start-time]                 Start time of the campaign
  [end-time]                   End time of the campaign
  [lockup-period]              Lockup period for the campaign
  [vesting-period]             Vesting period for the campaign
  [optional-vesting-pool-name] Optional name of the vesting pool

Example:
$ %s tx %s create-campaign "My Campaign" "Campaign description" "default" true "1000000"" "5000000"" "0.5" "2006-01-02 15:04:05 -0700 MST" "2006-01-02 15:04:05 -0700 MST" 720h 30h "" --from mykey
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(12),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argDescription := args[1]
			argCampaignType, err := types.CampaignTypeFromString(types.NormalizeCampaignType(args[2]))
			if err != nil {
				return err
			}
			argRemovableClaimRecords, err := strconv.ParseBool(args[3])
			if err != nil {
				return errors.Wrap(err, "Wrong [removable-claim-records] value")
			}
			argFeegrantAmount, ok := math.NewIntFromString(args[4])
			if !ok {
				return errors.Wrap(sdkerrors.ErrInvalidRequest, "Wrong [initial_claim_free_amount] value")
			}
			if err != nil {
				return err
			}
			argInitialClaimFreeAmount, ok := math.NewIntFromString(args[5])
			if !ok {
				return errors.Wrap(sdkerrors.ErrInvalidRequest, "Wrong [initial_claim_free_amount] value")
			}
			argFree, err := sdk.NewDecFromStr(args[6])
			if err != nil {
				return errors.Wrapf(sdkerrors.ErrInvalidRequest, "Wrong [initial_claim_free_amount] value, error: %s", err.Error())
			}
			argStartTime, err := time.Parse(TimeLayout, args[7])
			if err != nil {
				return err
			}
			argEndTime, err := time.Parse(TimeLayout, args[8])
			if err != nil {
				return err
			}
			if err != nil {
				return err
			}
			argLockupPeriod, err := time.ParseDuration(args[9])
			if err != nil {
				return errors.Wrap(sdkerrors.ErrInvalidRequest, "Expected duration format: e.g. 2h30m40s. Valid time units are “ns”, “us” (or “µs”), “ms”, “s”, “m”, “h”")
			}
			argVestingPeriod, err := time.ParseDuration(args[10])
			if err != nil {
				return errors.Wrap(sdkerrors.ErrInvalidRequest, "Expected duration format: e.g. 2h30m40s. Valid time units are “ns”, “us” (or “µs”), “ms”, “s”, “m”, “h”")
			}
			argVestingPoolName := args[11]
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateCampaign(
				clientCtx.GetFromAddress().String(),
				argName,
				argDescription,
				argCampaignType,
				argRemovableClaimRecords,
				&argFeegrantAmount,
				&argInitialClaimFreeAmount,
				&argFree,
				&argStartTime,
				&argEndTime,
				&argLockupPeriod,
				&argVestingPeriod,
				argVestingPoolName,
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
