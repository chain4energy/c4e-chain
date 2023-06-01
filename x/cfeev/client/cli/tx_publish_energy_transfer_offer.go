package cli

import (
	"fmt"
	"strconv"
	"strings"

	"encoding/json"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdPublishEnergyTransferOffer() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "publish-energy-transfer-offer [charger-id] [tariff] [location] [name] [plug-type]",
		Short: "Publish a new energy transfer offer that can be found by EV drivers",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Publish a new energy transfer offer that can be found by EV drivers.

Arguments:
  [charger-id] charger id specified on the charger
  [tariff] tariff at which energy transfer will be billed
  [location] charger location
  [name] charger name
  [plug-type] charger plug type

Example:
$ %s tx %s publish-energy-transfer-offer EVGC011221122GK0122 56 '{"latitude": "34.4", "longitude": "5.2"}' MySuperCharger 2
`,
				version.AppName, types.ModuleName,
			),
		),
		Args: cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argChargerId := args[0]
			argTariff, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}

			argLocation := new(types.Location)
			argName := args[3]
			err = json.Unmarshal([]byte(args[2]), argLocation)
			if err != nil {
				return err
			}

			// handle plug type enum values
			i, err := strconv.ParseInt(args[4], 10, 64)
			if err != nil {
				return err
			}
			i32 := int32(i)
			argPlugType := types.PlugType(i32)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgPublishEnergyTransferOffer(
				clientCtx.GetFromAddress().String(),
				argChargerId,
				argTariff,
				argLocation,
				argName,
				argPlugType,
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
