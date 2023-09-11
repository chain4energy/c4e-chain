package cli

import (
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/chain4energy/c4e-chain/v2/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdVest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-vesting-pool [name] [amount] [duration] [vesting-type]",
		Short: "Create a new vesting pool for the creator's address",
		Long: strings.TrimSpace(fmt.Sprintf(`Create a new vesting pool with a given name for the creator's address. In the newly created vesting pool a given amount of tokens is locked for a given duration.
During lock time tokens can only be send by the creator to vesting accounts with parameters set according to a given vesting-type.
Tokens can be withdrawn by the creator after lock duration.

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
			amountInt, ok := math.NewIntFromString(argAmount)
			if !ok {
				return errors.Wrap(sdkerrors.ErrInvalidRequest, "amount must be a positive integer")
			}

			duration, err := time.ParseDuration(argDuration)
			if err != nil {
				return errors.Wrap(sdkerrors.ErrInvalidRequest, "Expected duration format: e.g. 2h30m40s. Valid time units are “ns”, “us” (or “µs”), “ms”, “s”, “m”, “h”")
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
