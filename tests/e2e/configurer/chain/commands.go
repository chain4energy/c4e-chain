package chain

import (
	"fmt"
	"github.com/chain4energy/c4e-chain/app/params"
	"github.com/chain4energy/c4e-chain/tests/e2e/configurer/config"
	"github.com/chain4energy/c4e-chain/tests/e2e/initialization"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"os"
	"regexp"
	"strings"
)

func (n *NodeConfig) SubmitDepositAndVoteOnProposal(proposalJson, from string, chainConfig *Config) {
	n.SubmitParamChangeProposal(proposalJson, from)
	chainConfig.LatestProposalNumber += 1
	n.DepositProposal(chainConfig.LatestProposalNumber)
	for _, n := range chainConfig.NodeConfigs {
		n.VoteYesProposal(initialization.ValidatorWalletName, chainConfig.LatestProposalNumber)
	}
}

func (n *NodeConfig) SubmitParamChangeProposal(proposalJson, from string) {
	n.LogActionF("submitting param change proposal %s", proposalJson)
	localProposalFile := n.ConfigDir + "/param_change_proposal.json"
	f, err := os.Create(localProposalFile)
	require.NoError(n.t, err)
	_, err = f.WriteString(proposalJson)
	require.NoError(n.t, err)
	err = f.Close()
	require.NoError(n.t, err)

	cmd := []string{"c4ed", "tx", "gov", "submit-proposal", ".c4e-chain/param_change_proposal.json", formatFromFlag(from)}

	_, _, err = n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)
	err = os.Remove(localProposalFile)
	require.NoError(n.t, err)

	n.LogActionF("successfully submitted param change proposal")
}

func (n *NodeConfig) SubmitParamChangeNotValidProposal(proposalJson, from, errorMessage string) {
	n.LogActionF("submitting not valid param change proposal %s", proposalJson)
	localProposalFile := n.ConfigDir + "/param_change_proposal.json"
	f, err := os.Create(localProposalFile)
	require.NoError(n.t, err)
	_, err = f.WriteString(proposalJson)
	require.NoError(n.t, err)
	err = f.Close()
	require.NoError(n.t, err)

	cmd := []string{"c4ed", "tx", "gov", "submit-proposal", ".c4e-chain/param_change_proposal.json", formatFromFlag(from)}

	_, _, err = n.containerManager.ExecCmdWithResponseString(n.t, n.chainId, n.Name, cmd, errorMessage)
	require.NoError(n.t, err)
	err = os.Remove(localProposalFile)
	require.NoError(n.t, err)
}

func (n *NodeConfig) FailIBCTransfer(from, recipient, amount string) {
	n.LogActionF("IBC sending %s from %s to %s", amount, from, recipient)

	cmd := []string{"c4ed", "tx", "ibc-transfer", "transfer", "transfer", "channel-0", recipient, amount, formatFromFlag(from)}

	_, _, err := n.containerManager.ExecCmdWithResponseString(n.t, n.chainId, n.Name, cmd, "rate limit exceeded")
	require.NoError(n.t, err)

	n.LogActionF("Failed to send IBC transfer (as expected)")
}

func (n *NodeConfig) SubmitUpgradeProposal(upgradeVersion string, upgradeHeight int64, initialDeposit sdk.Coin) {
	n.LogActionF("submitting upgrade proposal %s for height %d", upgradeVersion, upgradeHeight)
	cmd := []string{"c4ed", "tx", "gov", "submit-legacy-proposal", "software-upgrade", upgradeVersion, fmt.Sprintf("--title=\"%s upgrade\"", upgradeVersion), "--description=\"upgrade proposal submission\"", fmt.Sprintf("--upgrade-height=%d", upgradeHeight), "--upgrade-info=\"\"", formatFromFlag("val"), formatDepositFlag(initialDeposit), "--no-validate"}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)
	n.LogActionF("successfully submitted upgrade proposal")
}

func (n *NodeConfig) SubmitLegacyUpgradeProposal(upgradeVersion string, upgradeHeight int64, initialDeposit sdk.Coin) {
	n.LogActionF("submitting upgrade proposal %s for height %d", upgradeVersion, upgradeHeight)
	cmd := []string{"c4ed", "tx", "gov", "submit-proposal", "software-upgrade", upgradeVersion, fmt.Sprintf("--title=\"%s upgrade\"", upgradeVersion), "--description=\"upgrade proposal submission\"", fmt.Sprintf("--upgrade-height=%d", upgradeHeight), "--upgrade-info=\"\"", formatFromFlag("val"), formatDepositFlag(initialDeposit)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)
	n.LogActionF("successfully submitted upgrade proposal")
}

func (n *NodeConfig) SubmitTextProposal(text string, initialDeposit sdk.Coin, isExpedited bool) {
	n.LogActionF("submitting text gov proposal")
	cmd := []string{"c4ed", "tx", "gov", "submit-proposal", "--type=text", fmt.Sprintf("--title=\"%s\"", text), "--description=\"test text proposal\"", formatFromFlag("val"), formatDepositFlag(initialDeposit)}
	if isExpedited {
		cmd = append(cmd, "--is-expedited=true")
	}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)
	n.LogActionF("successfully submitted text gov proposal")
}

func (n *NodeConfig) DepositProposal(proposalNumber int) {
	n.LogActionF("depositing on proposal: %d", proposalNumber)
	deposit := sdk.NewCoin(params.MicroC4eUnit, config.MinDepositValue)
	cmd := []string{"c4ed", "tx", "gov", "deposit", fmt.Sprintf("%d", proposalNumber), deposit.String(), formatFromFlag("val")}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)
	n.LogActionF("successfully deposited on proposal %d", proposalNumber)
}

func (n *NodeConfig) VoteYesProposal(from string, proposalNumber int) {
	n.LogActionF("voting yes on proposal: %d", proposalNumber)
	cmd := []string{"c4ed", "tx", "gov", "vote", fmt.Sprintf("%d", proposalNumber), "yes", formatFromFlag(from)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)
	n.LogActionF("successfully voted yes on proposal %d", proposalNumber)
}

func (n *NodeConfig) VoteNoProposal(from string, proposalNumber int) {
	n.LogActionF("voting no on proposal: %d", proposalNumber)
	cmd := []string{"c4ed", "tx", "gov", "vote", fmt.Sprintf("%d", proposalNumber), "no", formatFromFlag(from)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)
	n.LogActionF("successfully voted no on proposal: %d", proposalNumber)
}

func (n *NodeConfig) BankSend(amount string, sendAddress string, receiveAddress string) {
	n.LogActionF("bank sending %s from address %s to %s", amount, sendAddress, receiveAddress)
	cmd := []string{"c4ed", "tx", "bank", "send", sendAddress, receiveAddress, amount, formatFromFlag("val")}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)
	n.LogActionF("successfully sent bank sent %s from address %s to %s", amount, sendAddress, receiveAddress)
}

func (n *NodeConfig) BankSendBaseBalanceFromNode(receiveAddress string) {
	n.LogActionF("bank sending %s from address %s to %s", config.BaseBalance, n.PublicAddress, receiveAddress)
	cmd := []string{"c4ed", "tx", "bank", "send", n.PublicAddress, receiveAddress, config.BaseBalance.String(), formatFromFlag("val")}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)
	n.LogActionF("successfully sent bank sent %s from address %s to %s", config.BaseBalance, n.PublicAddress, receiveAddress)
}

func (n *NodeConfig) CreateWallet(walletName string) string {
	n.LogActionF("creating wallet %s", walletName)
	cmd := []string{"c4ed", "keys", "add", walletName, "--keyring-backend=test"}
	outBuf, _, err := n.containerManager.ExecCmd(n.t, n.Name, cmd, "")
	require.NoError(n.t, err)
	re := regexp.MustCompile("c4e(.{39})")
	walletAddr := fmt.Sprintf("%s\n", re.FindString(outBuf.String()))
	walletAddr = strings.TrimSuffix(walletAddr, "\n")
	n.LogActionF("created wallet %s, waller address - %s", walletName, walletAddr)
	return walletAddr
}

func (n *NodeConfig) GetWallet(walletName string) string {
	n.LogActionF("retrieving wallet %s", walletName)
	cmd := []string{"c4ed", "keys", "show", walletName, "--keyring-backend=test"}
	outBuf, _, err := n.containerManager.ExecCmd(n.t, n.Name, cmd, "")
	require.NoError(n.t, err)
	re := regexp.MustCompile("c4e(.{39})")
	walletAddr := fmt.Sprintf("%s\n", re.FindString(outBuf.String()))
	walletAddr = strings.TrimSuffix(walletAddr, "\n")
	n.LogActionF("wallet %s found, waller address - %s", walletName, walletAddr)
	return walletAddr
}

func (n *NodeConfig) CreateVestingPool(vestingPoolName, amount, duration, vestinType, from string) {
	n.LogActionF("creating vesting pool")
	cmd := []string{"c4ed", "tx", "cfevesting", "create-vesting-pool", vestingPoolName, amount, duration, vestinType, formatFromFlag(from)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)

	n.LogActionF("successfully created vesting pool %s", vestingPoolName)
}

func (n *NodeConfig) SendToVestingAccount(fromAddress, toAddress, vestingPoolName, amount, restartVesting string) {
	n.LogActionF("creating vesting pool")
	cmd := []string{"c4ed", "tx", "cfevesting", "send-to-vesting-account", toAddress, vestingPoolName, amount, restartVesting, fmt.Sprintf("--from=%s", fromAddress)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)

	n.LogActionF("successfully send vesting pool %s to vesting account %s", vestingPoolName, toAddress)
}

func (n *NodeConfig) WithdrawAllAvailable(from string) {
	n.LogActionF("withdraw all avaliable")
	cmd := []string{"c4ed", "tx", "cfevesting", "withdraw-all-available", formatFromFlag(from)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)

	n.LogActionF("successfully withdrew all avaliable vestings")
}

func (n *NodeConfig) CreateVestingAccount(toAddress string, amount string, startTime, endTime, from string) {
	n.LogActionF("creating vesting account")
	cmd := []string{"c4ed", "tx", "cfevesting", "create-vesting-account", toAddress, amount, startTime, endTime, formatFromFlag(from)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)

	n.LogActionF("successfully created vesting account %s", toAddress)
}

func (n *NodeConfig) SplitVesting(toAddress string, amount string, from string) {
	n.LogActionF("split vesting")
	cmd := []string{"c4ed", "tx", "cfevesting", "split-vesting", toAddress, amount, formatFromFlag(from)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)

	n.LogActionF("successfully splitted vesting to account %s", toAddress)
}

func (n *NodeConfig) SplitVestingError(toAddress string, amount string, from, errorString string) {
	n.LogActionF("split vesting")
	cmd := []string{"c4ed", "tx", "cfevesting", "split-vesting", toAddress, amount, formatFromFlag(from)}
	_, _, err := n.containerManager.ExecCmdWithResponseString(n.t, n.chainId, n.Name, cmd, errorString)
	require.NoError(n.t, err)
}

func (n *NodeConfig) MoveAvailableVesting(toAddress string, from string) {
	n.LogActionF("move available vesting")
	cmd := []string{"c4ed", "tx", "cfevesting", "move-available-vesting", toAddress, formatFromFlag(from)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)

	n.LogActionF("successfully moved all available vesting to account %s", toAddress)
}

func (n *NodeConfig) MoveAvailableVestingError(toAddress, from, errorString string) {
	n.LogActionF("move available vesting error")
	cmd := []string{"c4ed", "tx", "cfevesting", "move-available-vesting", toAddress, formatFromFlag(from)}
	_, _, err := n.containerManager.ExecCmdWithResponseString(n.t, n.chainId, n.Name, cmd, errorString)
	require.NoError(n.t, err)
}

func (n *NodeConfig) MoveAvailableVestingByDenoms(toAddress string, denoms string, from string) {
	n.LogActionF("move available vesting by denoms")
	cmd := []string{"c4ed", "tx", "cfevesting", "move-available-vesting-by-denoms", toAddress, denoms, formatFromFlag(from)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)

	n.LogActionF("successfully moved all available vesting by denoms to account %s", toAddress)
}

func (n *NodeConfig) MoveAvailableVestingByDenomsError(toAddress string, denoms string, from, errorString string) {
	n.LogActionF("move available vesting by denoms")
	cmd := []string{"c4ed", "tx", "cfevesting", "move-available-vesting-by-denoms", toAddress, denoms, formatFromFlag(from)}
	_, _, err := n.containerManager.ExecCmdWithResponseString(n.t, n.chainId, n.Name, cmd, errorString)
	require.NoError(n.t, err)
}

func (n *NodeConfig) CreateCampaign(name, description, campaignType, removableClaimRecords, feegrantAmount, initialClaimFreeAmount, free, startTime, endtime, lockupPeriod, vestingPeriod, vestingPoolName, from string) (campaignId uint64) {
	n.LogActionF("creating campaign")
	cmd := []string{"c4ed", "tx", "cfeclaim", "create-campaign", name, description, campaignType, removableClaimRecords, feegrantAmount, initialClaimFreeAmount, free, startTime, endtime, lockupPeriod, vestingPeriod, vestingPoolName, formatFromFlag(from)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)
	n.LogActionF("successfully created campaign %s", name)
	return n.QueryLastCampaignsId()
}

func (n *NodeConfig) CreateCampaignError(name, description, campaignType, removableClaimRecords, feegrantAmount, initialClaimFreeAmount, free, startTime, endtime, lockupPeriod, vestingPeriod, vestingPoolName, from, errorString string) {
	n.LogActionF("creating campaign")
	cmd := []string{"c4ed", "tx", "cfeclaim", "create-campaign", name, description, campaignType, removableClaimRecords, feegrantAmount, initialClaimFreeAmount, free, startTime, endtime, lockupPeriod, vestingPeriod, vestingPoolName, formatFromFlag(from)}
	_, _, err := n.containerManager.ExecCmdWithResponseString(n.t, n.chainId, n.Name, cmd, errorString)
	require.NoError(n.t, err)
}

func (n *NodeConfig) AddClaimRecords(campaignId, claimRecordsJsonFile, from string) {
	n.LogActionF("add user entries %s", claimRecordsJsonFile)
	localProposalFile := n.ConfigDir + "/user_entries.json"
	f, err := os.Create(localProposalFile)
	require.NoError(n.t, err)
	_, err = f.WriteString(claimRecordsJsonFile)
	require.NoError(n.t, err)
	err = f.Close()
	require.NoError(n.t, err)

	cmd := []string{"c4ed", "tx", "cfeclaim", "add-claim-records", campaignId, ".c4e-chain/user_entries.json", formatFromFlag(from)}

	_, _, err = n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)
	err = os.Remove(localProposalFile)
	require.NoError(n.t, err)

	n.LogActionF("successfully addedclaim records")
}

func (n *NodeConfig) AddClaimRecordsError(campaignId, claimRecordsJsonFile, from, errorString string) {
	n.LogActionF("add user entries %s", claimRecordsJsonFile)
	localProposalFile := n.ConfigDir + "/user_entries.json"
	f, err := os.Create(localProposalFile)
	require.NoError(n.t, err)
	_, err = f.WriteString(claimRecordsJsonFile)
	require.NoError(n.t, err)
	err = f.Close()
	require.NoError(n.t, err)

	cmd := []string{"c4ed", "tx", "cfeclaim", "add-claim-records", campaignId, ".c4e-chain/user_entries.json", formatFromFlag(from)}

	_, _, err = n.containerManager.ExecCmdWithResponseString(n.t, n.chainId, n.Name, cmd, errorString)
	require.NoError(n.t, err)
	err = os.Remove(localProposalFile)
	require.NoError(n.t, err)
}

func (n *NodeConfig) AddMission(campaignId, name, description, missionType, weight, claimStartDate, from string) {
	n.LogActionF("add mission to campaign")
	cmd := []string{"c4ed", "tx", "cfeclaim", "add-mission", campaignId, name, description, missionType, weight, claimStartDate, formatFromFlag(from)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)

	n.LogActionF("successfully add new mission %s to campaign %s", name, campaignId)
}

func (n *NodeConfig) AddMissionError(campaignId, name, description, missionType, weight, claimStartDate, from, errorString string) {
	n.LogActionF("add mission to campaign")
	cmd := []string{"c4ed", "tx", "cfeclaim", "add-mission", campaignId, name, description, missionType, weight, claimStartDate, formatFromFlag(from)}
	_, _, err := n.containerManager.ExecCmdWithResponseString(n.t, n.chainId, n.Name, cmd, errorString)
	require.NoError(n.t, err)
}

func (n *NodeConfig) EnableCampaign(campaignId, optionalStartTime, optionalEndTime, from string) {
	n.LogActionF("start campaign")
	cmd := []string{"c4ed", "tx", "cfeclaim", "enable-campaign", campaignId, optionalStartTime, optionalEndTime, formatFromFlag(from)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)

	n.LogActionF("successfully started campaign %s", campaignId)
}

func (n *NodeConfig) EnableCampaignError(campaignId, optionalStartTime, optionalEndTime, from, errorString string) {
	n.LogActionF("start campaign")
	cmd := []string{"c4ed", "tx", "cfeclaim", "enable-campaign", campaignId, optionalStartTime, optionalEndTime, formatFromFlag(from)}
	_, _, err := n.containerManager.ExecCmdWithResponseString(n.t, n.chainId, n.Name, cmd, errorString)
	require.NoError(n.t, err)
}

func (n *NodeConfig) CloseCampaign(campaignId, from string) {
	n.LogActionF("close campaign")
	cmd := []string{"c4ed", "tx", "cfeclaim", "close-campaign", campaignId, formatFromFlag(from)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)

	n.LogActionF("successfully closed campaign %s", campaignId)
}

func (n *NodeConfig) CloseCampaignError(campaignId, from, errorString string) {
	n.LogActionF("close campaign")
	cmd := []string{"c4ed", "tx", "cfeclaim", "close-campaign", campaignId, formatFromFlag(from)}
	_, _, err := n.containerManager.ExecCmdWithResponseString(n.t, n.chainId, n.Name, cmd, errorString)
	require.NoError(n.t, err)
}

func (n *NodeConfig) ClaimMission(campaignId, missionId, from string) {
	n.LogActionF("claim mission")
	cmd := []string{"c4ed", "tx", "cfeclaim", "claim", campaignId, missionId, formatFromFlag(from)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)

	n.LogActionF("successfully claimed mission %s from campaign %s", missionId, campaignId)
}

func (n *NodeConfig) ClaimMissionError(campaignId, missionId, from, errorString string) {
	n.LogActionF("claim mission")
	cmd := []string{"c4ed", "tx", "cfeclaim", "claim", campaignId, missionId, formatFromFlag(from)}
	_, _, err := n.containerManager.ExecCmdWithResponseString(n.t, n.chainId, n.Name, cmd, errorString)
	require.NoError(n.t, err)
}

func (n *NodeConfig) ClaimInitialMission(campaignId, destinationAddress, from string) {
	n.LogActionF("claim initial mission")
	cmd := []string{"c4ed", "tx", "cfeclaim", "initial-claim", campaignId, destinationAddress, formatFromFlag(from)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)

	n.LogActionF("successfully claimed initial mission from campaign %s with optional address %s", campaignId, destinationAddress)
}

func (n *NodeConfig) ClaimInitialMissionError(campaignId, destinationAddress, from, errorString string) {
	n.LogActionF("claim initial mission error")
	cmd := []string{"c4ed", "tx", "cfeclaim", "initial-claim", campaignId, destinationAddress, formatFromFlag(from)}
	_, _, err := n.containerManager.ExecCmdWithResponseString(n.t, n.chainId, n.Name, cmd, errorString)
	require.NoError(n.t, err)
}

func (n *NodeConfig) DeleteClaimRecord(campaignId, userAddress, from string) {
	n.LogActionF("delete claim record")
	cmd := []string{"c4ed", "tx", "cfeclaim", "delete-claim-record", campaignId, userAddress, formatFromFlag(from)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)

	n.LogActionF("successfully deleted claim %s record from campaign with id %d", userAddress, campaignId)
}

func (n *NodeConfig) DeleteClaimRecordError(campaignId, userEntryAddress, from, errorString string) {
	n.LogActionF("delete claim record error")
	cmd := []string{"c4ed", "tx", "cfeclaim", "delete-claim-record", campaignId, userEntryAddress, formatFromFlag(from)}
	_, _, err := n.containerManager.ExecCmdWithResponseString(n.t, n.chainId, n.Name, cmd, errorString)
	require.NoError(n.t, err)
}

func (n *NodeConfig) RemoveCampaign(campaignId, from string) {
	n.LogActionF("remove campaign")
	cmd := []string{"c4ed", "tx", "cfeclaim", "remove-campaign", campaignId, formatFromFlag(from)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)

	n.LogActionF("successfully removed campaign with id %d", campaignId)
}

func (n *NodeConfig) RemoveCampaignError(campaignId, from, errorString string) {
	n.LogActionF("remove campaign error")
	cmd := []string{"c4ed", "tx", "cfeclaim", "remove-campaign", campaignId, formatFromFlag(from)}
	_, _, err := n.containerManager.ExecCmdWithResponseString(n.t, n.chainId, n.Name, cmd, errorString)
	require.NoError(n.t, err)
}

func (n *NodeConfig) DelegateToValidator(validatorAddress, amount, from string) {
	n.LogActionF("delegate to validator")
	cmd := []string{"c4ed", "tx", "staking", "delegate", validatorAddress, amount, formatFromFlag(from)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)

	n.LogActionF("successfully delegated %s to validator %s", amount, validatorAddress)
}

func formatFromFlag(from string) string {
	return fmt.Sprintf("--from=%s", from)
}

func formatDepositFlag(desposit sdk.Coin) string {
	return fmt.Sprintf("--deposit=%s", desposit)
}
