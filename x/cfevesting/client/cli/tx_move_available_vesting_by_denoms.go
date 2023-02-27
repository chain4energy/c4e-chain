package cli

import (
	"strconv"
	"strings"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdMoveAvailableVestingByDenoms() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "move-available-vesting-by-denoms [to-address] [denoms]",
		Short: "Broadcast message MoveAvailableVestingByDenoms",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argToAddress := args[0]
			argDenoms := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			denoms, err := parseDenoms(argDenoms)
			if err != nil {
				return err
			}

			msg := types.NewMsgMoveAvailableVestingByDenoms(
				clientCtx.GetFromAddress().String(),
				argToAddress,
				denoms,
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

func parseDenoms(denomsStr string) (denoms []string, err error) {
	denomsStr = strings.TrimSpace(denomsStr)
	if len(denomsStr) == 0 {
		return nil, nil
	}

	denomStrs := strings.Split(denomsStr, ",")
	for _, denomStr := range denomStrs {
		coin := strings.TrimSpace(denomStr)
		if len(coin) > 0 {
			denoms = append(denoms, coin)
		}
		// TODO check if duplications
	}
	return
	// return newDecCoins, nil
}
