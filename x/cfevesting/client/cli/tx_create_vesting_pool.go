package cli

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdVest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-vesting-pool [name] [amount] [duration] [vesting-type]",
		Short: "Createsa a new vesting pool for creator address",
		Long: strings.TrimSpace(fmt.Sprintf(`Creates a new vesting pool with a given name for the creater address. In newly created vesting pool given amount of tokens is locked for given duration.
During lock time tokens can only be send by creator to vesting accounts with parameters set according to given vesting-type.
Tokens can be withdrawn by creator after lock duration.

Arguments:
  [name]         unique name per creator of new vesting pool
  [amount]       amount of tokens to lock
  [duration]     tokens lock duration. Expected duration format: e.g. 2h30m40s. Valid time units are “ns”, “us” (or “µs”), “ms”, “s”, “m”, “h”
  [vesting-type] name of predefined vesting type

Example:
$ %s tx %s create-vesting-pool my_vesting_pool 1000000 8760h my_vesting_type --from mykey
`,
			version.AppName, types.ModuleName,
		),
		),
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argAmount := args[1]
			argDuration := args[2]
			argVestingType := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			amountInt, ok := sdk.NewIntFromString(argAmount)
			if !ok {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount must be a positive integer")
			}

			duration, err := time.ParseDuration(argDuration)
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Expected duration format: e.g. 2h30m40s. Valid time units are “ns”, “us” (or “µs”), “ms”, “s”, “m”, “h”")
			}

			msg := types.NewMsgCreateVestingPool(
				clientCtx.GetFromAddress().String(),
				argName,
				amountInt,
				duration,
				argVestingType,
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
