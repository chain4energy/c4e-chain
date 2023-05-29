package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chain4energy/c4e-chain/x/cfefingerprint/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateReferencePayloadLink() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-reference-payload-link [payload-hash]",
		Short: "Create and publish reference payload link based on payload hash",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Create and publish reference payload link based on payload hash.

Arguments:
  [payloadHash] hash for a given payload e.g. some document

Example:
$ %s tx %s create-reference-payload-link 5b3f2de5a5a6a7054d04171dbd692f66f5236f06
`,
				version.AppName, types.ModuleName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPayloadHash := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateReferencePayloadLink(
				clientCtx.GetFromAddress().String(),
				argPayloadHash,
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
