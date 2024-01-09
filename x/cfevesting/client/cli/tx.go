package cli

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/client"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdVest())
	cmd.AddCommand(CmdWithdrawAllAvailable())
	cmd.AddCommand(CmdCreateVestingAccount())
	cmd.AddCommand(CmdSendToVestingAccount())
	cmd.AddCommand(CmdSplitVesting())
	cmd.AddCommand(CmdMoveAvailableVesting())
	cmd.AddCommand(CmdMoveAvailableVestingByDenoms())
	// this line is used by starport scaffolding # 1

	return cmd
}
