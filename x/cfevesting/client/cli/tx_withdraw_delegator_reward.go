package cli

import (
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	sdk "github.com/cosmos/cosmos-sdk/types"

)

var _ = strconv.Itoa(0)

func CmdWithdrawDelegatorReward() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-delegator-reward [validator-address]",
		Short: "Broadcast message withdrawDelegatorReward",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argValidatorAddress := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			_, err = sdk.ValAddressFromBech32(argValidatorAddress)
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdrawDelegatorReward(
				clientCtx.GetFromAddress().String(),
				argValidatorAddress,
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
