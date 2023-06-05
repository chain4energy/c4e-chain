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

func CmdCancelEnergyTransferRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel-energy-transfer-request [energy-transfer-id] [charger-id] [error-info] [error-code]",
		Short: "Cancel energy transfer",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Cancel energy transfer and optionally specify the reason.

Arguments:
  [energy-transfer-id] energy transfer identifier
  [charger-id] charger id specified on the charger
  [error-info] error info - optional 
  [error-code] error code - optional 

Example:
$ %s tx %s cancel-energy-transfer-request 0 EVGC011221122GK0122 charger_not_responding 4
`,
				version.AppName, types.ModuleName,
			),
		),
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argEnergyTransferId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argChargerId := args[1]
			argErrorInfo := args[2]
			argErrorCode := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCancelEnergyTransferRequest(
				clientCtx.GetFromAddress().String(),
				argEnergyTransferId,
				argChargerId,
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
