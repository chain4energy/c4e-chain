package cli

import (
	"encoding/json"
	claimtypes "github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	"os"
)

func parseClaimEntries(campaignId uint64, claimEntriesFile string) ([]*claimtypes.ClaimRecord, error) {
	var claimEntries []*claimtypes.ClaimRecord

	if claimEntriesFile == "" {
		return claimEntries, nil
	}

	contents, err := os.ReadFile(claimEntriesFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &claimEntries)
	if err != nil {
		return nil, err
	}
	for _, claimRecord := range claimEntries {
		claimRecord.CampaignId = campaignId
	}
	return claimEntries, nil
}
