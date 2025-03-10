package cli

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
)

const (
	TimeLayout = "2006-01-02 15:04:05 -0700 MST"
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

	cmd.AddCommand(CmdClaim())
	cmd.AddCommand(CmdInitialClaim())
	cmd.AddCommand(CmdCreateCampaign())
	cmd.AddCommand(CmdAddMission())
	cmd.AddCommand(CmdAddClaimRecords())
	cmd.AddCommand(CmdDeleteClaimRecord())
	cmd.AddCommand(CmdCloseCampaign())
	cmd.AddCommand(CmdEnableCampaign())
	cmd.AddCommand(CmdRemoveCampaign())
	// this line is used by starport scaffolding # 1

	return cmd
}
