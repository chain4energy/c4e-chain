package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/version"
	"strconv"
	"strings"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdStartEnergyTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start-energy-transfer [energy-transfer-offer-id] [offered-tariff] [energy-to-transfer]",
		Short: "Start energy transfer",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Start energy transfer

Arguments:
  [energy-transfer-offer-id] energy transfer offer identifier
  [offered-tariff] offered tariff
  [energy-to-transfer] energy to transfer

Example:
$ %s tx %s start-energy-transfer 0 100 100 --from mykey
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argEnergyTransferOfferId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argOfferedTariff, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argEnergyToTransfer, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgStartEnergyTransfer(
				clientCtx.GetFromAddress().String(),
				argEnergyTransferOfferId,
				argOfferedTariff,
				argEnergyToTransfer,
			)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
