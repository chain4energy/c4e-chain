package cli

import (
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfeenergybank/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateTokenParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-token-params [name] [trading-company] [burning-time] [burning-type] [send-price]",
		Short: "Broadcast message create-token-params",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argTradingCompany := args[1]
			argBurningTime, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}
			argBurningType := args[3]
			argExchangeRate, err := cast.ToUint64E(args[4])
			argCommissionRate, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateTokenParams(
				clientCtx.GetFromAddress().String(),
				argName,
				argTradingCompany,
				argBurningTime,
				argBurningType,
				argExchangeRate,
				argCommissionRate,
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
