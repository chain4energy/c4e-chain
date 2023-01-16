package cli

import (
	airdroptypes "github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	"github.com/cosmos/cosmos-sdk/client"
	"io/ioutil"
)

func parseAirdropEntries(clientCtx client.Context, airdropEntriesFile string) ([]*airdroptypes.AirdropEntry, error) {
	var airdropEntries airdroptypes.AirdropEntries

	if airdropEntriesFile == "" {
		return airdropEntries.AirdropEntries, nil
	}

	contents, err := ioutil.ReadFile(airdropEntriesFile)
	if err != nil {
		return nil, err
	}

	err = clientCtx.Codec.UnmarshalJSON(contents, &airdropEntries)
	if err != nil {
		return nil, err
	}

	return airdropEntries.AirdropEntries, nil
}
