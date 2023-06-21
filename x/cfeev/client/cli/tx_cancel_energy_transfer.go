package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCancelEnergyTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel-energy-transfer [energy-transfer-id] [error-info] [error-code]",
		Short: "Cancel energy transfer",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Cancel energy transfer and optionally specify the reason.

Arguments:
  [energy-transfer-id] energy transfer identifier
  [error-info] error info - optional 
  [error-code] error code - optional 

Example:
$ %s tx %s cancel-energy-transfer 0 EVGC011221122GK0122 charger_not_responding 4 --from mykey
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argEnergyTransferId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argErrorInfo := args[1]
			argErrorCode := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCancelEnergyTransfer(
				clientCtx.GetFromAddress().String(),
				argEnergyTransferId,
				argErrorInfo,
				argErrorCode,
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
