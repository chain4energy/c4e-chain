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

func CmdCreateAirdropCampaign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-airdrop-campaign [name] [description] [start-time] [end-time] [lockup-period] [vesting-period]",
		Short: "Broadcast message CreateAirdropCampaign",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argDescription := args[1]
			argStartTime, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				return err
			}
			argEndTime, err := strconv.ParseInt(args[3], 10, 64)
			if err != nil {
				return err
			}
			argLockupPeriod, err := time.ParseDuration(args[4])
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Expected duration format: e.g. 2h30m40s. Valid time units are “ns”, “us” (or “µs”), “ms”, “s”, “m”, “h”")
			}
			argVestingPeriod, err := time.ParseDuration(args[5])
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Expected duration format: e.g. 2h30m40s. Valid time units are “ns”, “us” (or “µs”), “ms”, “s”, “m”, “h”")
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateAirdropCampaign(
				clientCtx.GetFromAddress().String(),
				argName,
				argDescription,
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
