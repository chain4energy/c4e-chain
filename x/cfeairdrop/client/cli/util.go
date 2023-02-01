package cli

import (
	"encoding/json"
	airdroptypes "github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	"github.com/cosmos/cosmos-sdk/client"
	"io/ioutil"
)

func parseAirdropEntries(clientCtx client.Context, campaignId uint64, airdropEntriesFile string) ([]*airdroptypes.ClaimRecord, error) {
	var airdropEntries []*airdroptypes.ClaimRecord

	if airdropEntriesFile == "" {
		return airdropEntries, nil
	}

	contents, err := ioutil.ReadFile(airdropEntriesFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, airdropEntries)
	if err != nil {
		return nil, err
	}
	for _, claimRecord := range airdropEntries {
		claimRecord.CampaignId = campaignId
	}
	return airdropEntries, nil
}
