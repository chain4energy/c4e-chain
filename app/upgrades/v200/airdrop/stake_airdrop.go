package airdrop

import (
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var stakeRecords = []*types.AirdropEntry{
	{
		Address:      "c4e1yyjfd5cj5nd0jrlvrhc5p3mnkcn8v9q8fdd9gs",
		AirdropCoins: sdk.NewCoins(sdk.NewCoin("uc4e", sdk.NewInt(4000))),
	},
	{
		Address:      "c4e1uqenz7vu5xyxfsr70was6zzf5g3spsy52lenlc",
		AirdropCoins: sdk.NewCoins(sdk.NewCoin("uc4e", sdk.NewInt(6000))),
	},
	{
		Address:      "c4e17dffs6qldsh30un0jm68ggr40rzkm7tmvh0e78",
		AirdropCoins: sdk.NewCoins(sdk.NewCoin("uc4e", sdk.NewInt(8000))),
	},
}
