package cli

import (
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	"fmt"
	"github.com/chain4energy/c4e-chain/app/params"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
		Use:   "start-energy-transfer [energy-transfer-offer-id] [charger-id] [owner-account-address] [offered-tariff] [energy-to-transfer] [collateral]",
		Short: "Start energy transfer",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Start energy transfer

Arguments:
  [energy-transfer-offer-id] energy transfer offer identifier
  [charger-id] charger id specified on the charger
  [info] additional info - optional
  [owner-account-address] owner account address
  [offered-tariff] offered tariff
  [energy-to-transfer] energy to transfer
  [collateral] collateral

Example:
$ %s tx %s start-energy-transfer 0 EVGC011221122GK0122 %s1gghjut3ccd8ay0zduzj64hwre2fxs9ldmqhffj 100 100 "100" --from mykey
`, version.AppName, types.ModuleName, params.Bech32PrefixAccAddr)),
		Args: cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argEnergyTransferOfferId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argChargerId := args[1]
			argOwnerAccountAddress := args[2]
			argOfferedTariff, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}
			argEnergyToTransfer, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}

			argCollateral, ok := math.NewIntFromString(args[5])
			if !ok {
				return errors.Wrap(sdkerrors.ErrInvalidRequest, "Wrong [collateral] value")
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgStartEnergyTransfer(
				clientCtx.GetFromAddress().String(),
				argEnergyTransferOfferId,
				argChargerId,
				argOwnerAccountAddress,
				argOfferedTariff,
				&argCollateral,
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
