package cli

import (
	"cosmossdk.io/errors"
	"fmt"
	"strconv"
	"strings"

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

func CmdSendToVestingAccount() *cobra.Command {
	bech32PrefixAddr := sdk.GetConfig().GetBech32AccountAddrPrefix()

	cmd := &cobra.Command{
		Use:   "send-to-vesting-account [to-address] [vesting-pool-name] [amount] [restart-vesting]",
		Short: "Create a new vesting account funded with an allocation of tokens from a given vesting pool.",
		Long: strings.TrimSpace(fmt.Sprintf(`Create a new vesting account funded with an allocation of tokens from a given vesting pool.
The account is a continuous vesting account. The start_time and the end_time are calculated according to the vesting type of a given vesting-pool.
If [restart-vesting] is set to true than
  - start_time = committed block's time + vesting's type of vesting pool lock time
  - end_time = committed block's time + vesting's type of vesting pool lock time + vesting's type of vesting pool vesting time
If [restart-vesting] is set to false than
  - start_time = given vesting pool lock end time
  - end_time = given vesting pool lock end time
Before new vesting account creation all available tokens are withdrawn from the vesting pool. If current time is after vesting pool lock end new vesting account creation fails.

Arguments:
  [to_address]        address of a new vesting account
  [vesting-pool-name] token source vesting pool name 
  [amount]            amount of tokens to send to a new vesting account from the vesting pool
  [restart-vesting]   true or false. Specifies if vesting account params should be calculted accordign to the vesting type

Example:
$ %s tx %s create-vesting-account %s1gghjut3ccd8ay0zduzj64hwre2fxs9ldmqhffj my_vesting_pool 120000 true --from mykey
`, version.AppName, types.ModuleName, bech32PrefixAddr,
		),
		),
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argToAddress := args[0]
			argVestingPoolName := args[1]
			argAmount := args[2]
			argRestartVesting := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			if argVestingPoolName == "" {
				return errors.Wrap(sdkerrors.ErrInvalidRequest, "vesting-pool name cannot be empty")
			}
			amountInt, ok := sdk.NewIntFromString(argAmount)
			if !ok || amountInt.LTE(sdk.ZeroInt()) {
				return errors.Wrap(sdkerrors.ErrInvalidRequest, "amount must be a positive integer")
			}
			restartVestingBool, err := strconv.ParseBool(argRestartVesting)
			if err != nil {
				return errors.Wrap(sdkerrors.ErrInvalidRequest, "restart-vesting must be a boolean [1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False]")
			}

			msg := types.NewMsgSendToVestingAccount(
				clientCtx.GetFromAddress().String(),
				argToAddress,
				argVestingPoolName,
				amountInt,
				restartVestingBool,
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
