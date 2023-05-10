package cli

import (
	"strconv"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdStartCampaign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start-campaign [campaign-id] [optional-start-time] [optional-end-time]",
		Short: "Broadcast message StartCampaign",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCampaignId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var argStartTime *time.Time
			if args[1] != "" {
				parsedTime, err := time.Parse(TimeLayout, args[1])
				if err != nil {
					return err
				}
				argStartTime = &parsedTime

			}
			var argEndTime *time.Time
			if args[2] != "" {
				parsedTime, err := time.Parse(TimeLayout, args[2])
				if err != nil {
					return err
				}
				argEndTime = &parsedTime

			}

			msg := types.NewMsgStartCampaign(
				clientCtx.GetFromAddress().String(),
				argCampaignId,
				argStartTime,
				argEndTime,
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