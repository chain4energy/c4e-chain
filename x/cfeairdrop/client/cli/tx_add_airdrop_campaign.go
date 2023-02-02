package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateCampaign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-campaign [name] [description] [campaign-type] [feegrant-amount] [initial_claim_free_amount] [start-time] [end-time] [lockup-period] [vesting-period]",
		Short: "Broadcast message CreateCampaign",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argDescription := args[1]
			argCampaignType, err := types.CampaignTypeFromString(types.NormalizeCampaignType(args[2]))
			if err != nil {
				return err
			}
			argFeegrantAmount, ok := sdk.NewIntFromString(args[3])
			if !ok {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Wrong [initial_claim_free_amount] value")
			}
			if err != nil {
				return err
			}
			argInitialClaimFreeAmount, ok := sdk.NewIntFromString(args[4])
			if !ok {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Wrong [initial_claim_free_amount] value")
			}
			timeLayout := "2006-01-02 15:04:05 -0700 MST"
			argStartTime, err := time.Parse(timeLayout, args[5])
			if err != nil {
				return err
			}
			argEndTime, err := time.Parse(timeLayout, args[6])
			if err != nil {
				return err
			}
			if err != nil {
				return err
			}
			argLockupPeriod, err := time.ParseDuration(args[7])
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Expected duration format: e.g. 2h30m40s. Valid time units are “ns”, “us” (or “µs”), “ms”, “s”, “m”, “h”")
			}
			argVestingPeriod, err := time.ParseDuration(args[8])
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Expected duration format: e.g. 2h30m40s. Valid time units are “ns”, “us” (or “µs”), “ms”, “s”, “m”, “h”")
			}
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
				&argStartTime,
				&argEndTime,
				&argLockupPeriod,
				&argVestingPeriod,
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
