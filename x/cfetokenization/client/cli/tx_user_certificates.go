package cli

import (
	"encoding/json"
	"github.com/chain4energy/c4e-chain/x/cfetokenization/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdCreateUserCertificates() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-user-certificates [certificates]",
		Short: "Create a new UserCertificates",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCertificates := new(types.Certificate)
			err = json.Unmarshal([]byte(args[0]), argCertificates)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateUserCertificates(clientCtx.GetFromAddress().String(), argCertificates)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
