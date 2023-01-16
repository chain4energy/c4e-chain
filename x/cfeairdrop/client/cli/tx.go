package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
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
	cmd.AddCommand(CmdCreateAirdropCampaign())
	cmd.AddCommand(CmdAddMissionToAidropCampaign())
	cmd.AddCommand(CmdCreateAirdropEntry())
	cmd.AddCommand(CmdDeleteAirdropEntry())

	cmd.AddCommand(CmdCloseAirdropCampaign())
	cmd.AddCommand(CmdStartAirdropCampaign())
	// this line is used by starport scaffolding # 1

	return cmd
}
