package airdrop

import (
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var gleamContestRecords = []*types.AirdropEntry{
	{
		Address: "c4e1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8fdd9gs",
		Amount:  sdk.NewInt(10000),
	},
	{
		Address: "c4e1uqenz7vu5xyxfsr70was6zzf5g3spsy52lenlc",
		Amount:  sdk.NewInt(20000),
	},
	{
		Address: "c4e17dffs6qldsh30un0jm68ggr40rzkm7tmvh0e78",
		Amount:  sdk.NewInt(30000),
	},
}
