package chain

import "fmt"

func (n *NodeConfig) ValidateCampaignIsNotOverYet(campaignIdString, creatorWalletName string) {
	n.CloseCampaignError(campaignIdString, creatorWalletName, "campaign is not over yet")
}

func (n *NodeConfig) ValidateDeleteClaimerNotFound(campaignIdString, userEntryAddress, creatorWalletName string) {
	n.DeleteClaimRecordError(campaignIdString, userEntryAddress, creatorWalletName, fmt.Sprintf("claim record with campaign id %s not found for address %s", campaignIdString, userEntryAddress))
}

func (n *NodeConfig) ValidateClaimInitialClaimerNotFound(campaignIdString, destinationAddress, claimer string) {
	n.ClaimInitialMissionError(campaignIdString, destinationAddress, claimer, "not found: key not found")
}

func (n *NodeConfig) ValidateClaimInitialCampaignNotStartedYet(campaignIdString, destinationAddress, claimer string) {
	n.ClaimInitialMissionError(campaignIdString, destinationAddress, claimer, fmt.Sprintf("campaign %s not started yet", campaignIdString))
}
