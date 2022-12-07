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
		Use:   "create-airdrop-campaign [owner] [name] [campaign-duration] [lockup-period] [vesting-period] [description]",
		Short: "Broadcast message CreateAirdropCampaign",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argOwner := args[0]
			argName := args[1]
			argCampaignDuration := args[2]
			argLockupPeriod := args[3]
			argVestingPeriod := args[4]
			argDescription := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			campaignDuration, err := time.ParseDuration(argCampaignDuration)
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Expected duration format: e.g. 2h30m40s. Valid time units are “ns”, “us” (or “µs”), “ms”, “s”, “m”, “h”")
			}

			lockupPeriod, err := time.ParseDuration(argLockupPeriod)
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Expected duration format: e.g. 2h30m40s. Valid time units are “ns”, “us” (or “µs”), “ms”, “s”, “m”, “h”")
			}

			vestingPeriod, err := time.ParseDuration(argVestingPeriod)
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Expected duration format: e.g. 2h30m40s. Valid time units are “ns”, “us” (or “µs”), “ms”, “s”, “m”, “h”")
			}

			msg := types.NewMsgCreateAirdropCampaign(
				clientCtx.GetFromAddress().String(),
				argOwner,
				argName,
				campaignDuration,
				lockupPeriod,
				vestingPeriod,
				argDescription,
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
