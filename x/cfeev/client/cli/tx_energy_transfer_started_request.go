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

func CmdEnergyTransferStartedRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "energy-transfer-started-request [energy-transfer-id] [charger-id] [info]",
		Short: "Confirm that energy transfer has finally started",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Confirm that energy transfer has finally started.

Arguments:
  [energy-transfer-id] energy transfer identifier
  [charger-id] charger id specified on the charger
  [info] additional info - optional

Example:
$ %s tx %s energy-transfer-started-request 0 EVGC011221122GK0122 started
`,
				version.AppName, types.ModuleName,
			),
		),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argEnergyTransferId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argChargerId := args[1]
			argInfo := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgEnergyTransferStartedRequest(
				clientCtx.GetFromAddress().String(),
				argEnergyTransferId,
				argChargerId,
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
