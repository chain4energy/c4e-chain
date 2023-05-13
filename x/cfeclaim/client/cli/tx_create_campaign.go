package cli

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"
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
		Use:   "create-campaign [name] [description] [campaign-type] [feegrant-amount] [initial_claim_free_amount] [free] [start-time] [end-time] [lockup-period] [vesting-period] [optional-vesting-pool-name]",
		Short: "Broadcast message CreateCampaign",
		Args:  cobra.ExactArgs(11),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argDescription := args[1]
			argCampaignType, err := types.CampaignTypeFromString(types.NormalizeCampaignType(args[2]))
			if err != nil {
				return err
			}
			argFeegrantAmount, ok := sdk.NewIntFromString(args[3])
			if !ok {
				return errors.Wrap(sdkerrors.ErrInvalidRequest, "Wrong [initial_claim_free_amount] value")
			}
			if err != nil {
				return err
			}
			argInitialClaimFreeAmount, ok := sdk.NewIntFromString(args[4])
			if !ok {
				return errors.Wrap(sdkerrors.ErrInvalidRequest, "Wrong [initial_claim_free_amount] value")
			}
			argFree, err := sdk.NewDecFromStr(args[5])
			if err != nil {
				return errors.Wrapf(sdkerrors.ErrInvalidRequest, "Wrong [initial_claim_free_amount] value, error: %s", err.Error())
			}
			argStartTime, err := time.Parse(TimeLayout, args[6])
			if err != nil {
				return err
			}
			argEndTime, err := time.Parse(TimeLayout, args[7])
			if err != nil {
				return err
			}
			if err != nil {
				return err
			}
			argLockupPeriod, err := time.ParseDuration(args[8])
			if err != nil {
				return errors.Wrap(sdkerrors.ErrInvalidRequest, "Expected duration format: e.g. 2h30m40s. Valid time units are “ns”, “us” (or “µs”), “ms”, “s”, “m”, “h”")
			}
			argVestingPeriod, err := time.ParseDuration(args[9])
			if err != nil {
				return errors.Wrap(sdkerrors.ErrInvalidRequest, "Expected duration format: e.g. 2h30m40s. Valid time units are “ns”, “us” (or “µs”), “ms”, “s”, “m”, “h”")
			}
			argVestingPoolName := args[10]
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateCampaign(
				clientCtx.GetFromAddress().String(),
				argName,
				argDescription,
				argCampaignType,
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
