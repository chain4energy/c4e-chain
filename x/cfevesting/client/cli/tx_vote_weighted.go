package cli

import (
	"fmt"
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	govutils "github.com/cosmos/cosmos-sdk/x/gov/client/utils"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdVoteWeighted() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vote-weighted [proposal-id] [options]",
		Short: "Broadcast message voteWeighted",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argProposalId := args[0]
			argOptions := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			proposalID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("proposal-id %s not a valid int, please input a valid proposal-id", argProposalId)
			}

			// Figure out which vote options user chose
			options, err := govtypes.WeightedVoteOptionsFromString(govutils.NormalizeWeightedVoteOptions(argOptions))
			if err != nil {
				return err
			}

			msg := types.NewMsgVoteWeighted(
				clientCtx.GetFromAddress().String(),
				proposalID,
				options,
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
