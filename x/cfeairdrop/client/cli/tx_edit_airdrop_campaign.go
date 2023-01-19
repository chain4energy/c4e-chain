package cli

import (
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

func CmdEditAirdropCampaign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit-airdrop-campaign [campaignId] [name] [description] [denom] [start-time] [end-time] [lockup-period] [vesting-period]",
		Short: "Broadcast message CreateAirdropCampaign",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCampaignId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			argName := args[1]
			argDescription := args[2]
			argDenom := args[3]
			timeLayout := "2006-01-02 15:04:05 -0700 MST"

			var argStartTime *time.Time
			if args[4] != "" {
				parsedTime, err := time.Parse(timeLayout, args[4])
				if err != nil {
					return err
				}
				argStartTime = &parsedTime

			} else {
				argStartTime = nil
			}

			var argEndTime *time.Time
			if args[5] != "" {
				parsedTime, err := time.Parse(timeLayout, args[5])
				if err != nil {
					return err
				}
				argEndTime = &parsedTime

			} else {
				argEndTime = nil
			}

			var argLockupPeriod *time.Duration
			if args[6] != "" {
				parsedTime, err := time.ParseDuration(args[6])
				if err != nil {
					return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Expected duration format: e.g. 2h30m40s. Valid time units are “ns”, “us” (or “µs”), “ms”, “s”, “m”, “h”")
				}
				argLockupPeriod = &parsedTime

			} else {
				argLockupPeriod = nil
			}

			var argVestingPeriod *time.Duration
			if args[6] != "" {
				parsedTime, err := time.ParseDuration(args[6])
				if err != nil {
					return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Expected duration format: e.g. 2h30m40s. Valid time units are “ns”, “us” (or “µs”), “ms”, “s”, “m”, “h”")
				}
				argVestingPeriod = &parsedTime

			} else {
				argVestingPeriod = nil
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgEditAirdropCampaign(
				clientCtx.GetFromAddress().String(),
				argCampaignId,
				argName,
				argDescription,
				argDenom,
				argStartTime,
				argEndTime,
				argLockupPeriod,
				argVestingPeriod,
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
