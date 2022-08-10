package cli

import (
	"strconv"

	"github.com/chain4energy/c4e-chain/x/energybank/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdEnergyTokenUserAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "energy-token-user-address [user-address]",
		Short: "Query energy-token-user-address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqUserAddress := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryEnergyTokenUserAddressRequest{

				UserAddress: reqUserAddress,
			}

			res, err := queryClient.EnergyTokenUserAddress(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
