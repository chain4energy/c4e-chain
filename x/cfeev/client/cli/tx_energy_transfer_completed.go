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

func CmdEnergyTransferCompleted() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "energy-transfer-completed [energy-transfer-id] [used-service-units] [info]",
		Short: "Indicate that energy transfer has been completed",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Indicate that energy transfer has been completed.

Arguments:
  [energy-transfer-id] energy transfer identifier
  [used-service-units] the number of service units (kWh) that were transferred
  [info] additional info - optional

Example:
$ %s tx %s energy-transfer-completed 0 EVGC011221122GK0122 22 completed --from mykey
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argEnergyTransferId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argUsedServiceUnits, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argInfo := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgEnergyTransferCompleted(
				clientCtx.GetFromAddress().String(),
				argEnergyTransferId,
				argUsedServiceUnits,
				argInfo,
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
