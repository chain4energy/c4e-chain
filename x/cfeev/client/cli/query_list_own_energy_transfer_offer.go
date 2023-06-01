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

func CmdListOwnEnergyTransferOffer() *cobra.Command {
	bech32PrefixAddr := sdk.GetConfig().GetBech32AccountAddrPrefix()
	cmd := &cobra.Command{
		Use:   "list-own-energy-transfer-offer [owner-acc-address]",
		Short: "Query for all energy transfer offers of a given CP owner address",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query for all energy transfer offers of a given CP owner address.

Arguments:
  [ownerAccAddress] CP owner account address

Example:
$ %s query %s list-own-energy-transfer-offer %se1mydzr5gxtypyjks08nveclwjmjp64trxh4laxj
`,
				version.AppName, types.ModuleName, bech32PrefixAddr,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqOwnerAccAddress := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryListOwnEnergyTransferOfferRequest{

				OwnerAccAddress: reqOwnerAccAddress,
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			params.Pagination = pageReq

			res, err := queryClient.ListOwnEnergyTransferOffer(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
