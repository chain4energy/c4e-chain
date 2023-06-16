package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chain4energy/c4e-chain/x/cfeev/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdOwnEnergyTransfers() *cobra.Command {
	bech32PrefixAddr := sdk.GetConfig().GetBech32AccountAddrPrefix()
	cmd := &cobra.Command{
		Use:   "own-energy-transfers [driver] [transfer-status]",
		Short: "Query for all energy transfers of a given EV driver address",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query for all energy transfers of a given EV driver address.

Arguments:
  [driver] EV driver account address
  [transfer-status] energy transfer status (0: pending, 1: completed, 2: cancelled)
Example:
$ %s query %s list-owner-energy-transfer-offer %se1mydzr5gxtypyjks08nveclwjmjp64trxh4laxj
`,
				version.AppName, types.ModuleName, bech32PrefixAddr,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			driver := args[0]

			status, err := strconv.ParseUint(args[1], 10, 32)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryOwnEnergyTransfersRequest{
				Driver:         driver,
				TransferStatus: int32(status),
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			params.Pagination = pageReq

			res, err := queryClient.OwnEnergyTransfers(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
