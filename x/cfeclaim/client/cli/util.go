package cli

import (
	"encoding/json"
	claimtypes "github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	"os"
)

func parseClaimEntries(campaignId uint64, claimEntriesFile string) ([]*claimtypes.ClaimRecordEntry, error) {
	var claimRecordEntries []*claimtypes.ClaimRecordEntry

	if claimEntriesFile == "" {
		return claimRecordEntries, nil
	}

	contents, err := os.ReadFile(claimEntriesFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &claimRecordEntries)
	if err != nil {
		return nil, err
	}
	for _, claimRecord := range claimRecordEntries {
		claimRecord.CampaignId = campaignId
	}
	return claimRecordEntries, nil
}
