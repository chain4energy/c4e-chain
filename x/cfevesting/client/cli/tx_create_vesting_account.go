package cli

import (
	"fmt"
	"github.com/chain4energy/c4e-chain/v2/app/params"
	"strconv"
	"strings"

	"github.com/chain4energy/c4e-chain/v2/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateVestingAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-vesting-account [to-address] [amount] [start-time] [end-time]",
		Short: "Create a new vesting account funded with an allocation of tokens.",
		Long: strings.TrimSpace(fmt.Sprintf(`Create a new vesting account funded with an allocation of tokens.
The account is a continuous vesting account. The start_time and the end_time must be provided as a UNIX epoch timestamp.

Arguments:
  [to_address] address of a new vesting account
  [amount]     amount of tokens to send to a new vesting account
  [start-time] start time of continuous vesting account vesting. Must be provided as a UNIX epoch timestamp.
  [end-time]   end time of continuous vesting account vesting. Must be provided as a UNIX epoch timestamp.

Example:
$ %s tx %s create-vesting-account %s1gghjut3ccd8ay0zduzj64hwre2fxs9ldmqhffj 123000 1609455 1640991 --from mykey
`, version.AppName, types.ModuleName, params.Bech32PrefixAccAddr),
		),
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argToAddress := args[0]
			argAmount := args[1]
			argStartTime := args[2]
			argEndTime := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinsNormalized(argAmount)
			if err != nil {
				return err
			}

			startTime, err := strconv.ParseInt(argStartTime, 10, 64)
			if err != nil {
				return err
			}

			endTime, err := strconv.ParseInt(argEndTime, 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateVestingAccount(
				clientCtx.GetFromAddress().String(),
				argToAddress,
				amount,
				startTime,
				endTime,
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
