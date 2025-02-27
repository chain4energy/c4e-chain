package cli

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
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
	cmd.AddCommand(CmdBurn())
	// this line is used by starport scaffolding # 1

	return cmd
}
