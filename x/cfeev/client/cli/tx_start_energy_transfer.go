package cli

import (
	"encoding/json"
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdStartEnergyTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start-energy-transfer [energy-transfer-offer-id] [charger-id] [owner-account-address] [offered-tariff]",
		Short: "Broadcast message start-energy-transfer",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argEnergyTransferOfferId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argChargerId := args[1]
			argOwnerAccountAddress := args[2]
			argOfferedTariff := args[3]
			argEnergyToTransfer, err := cast.ToInt32E(args[5])
			if err != nil {
				return err
			}

			argCollateral := new(sdk.Coin)
			err = json.Unmarshal([]byte(args[4]), argCollateral)
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
				argChargerId,
				argOwnerAccountAddress,
				argOfferedTariff,
				argCollateral,
				argEnergyToTransfer,
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
