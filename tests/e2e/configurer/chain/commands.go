package chain

import (
	"encoding/json"
	"fmt"
	"github.com/chain4energy/c4e-chain/app/params"
	"github.com/chain4energy/c4e-chain/tests/e2e/configurer/config"
	"github.com/chain4energy/c4e-chain/tests/e2e/initialization"
	"github.com/chain4energy/c4e-chain/tests/e2e/util"
	cfedistributormoduletypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	cfemintermoduletypes "github.com/chain4energy/c4e-chain/x/cfeminter/types"
	cfevestingmoduletypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"os"
	"regexp"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func (n *NodeConfig) QueryCfevestingParams(moduleParams *cfevestingmoduletypes.QueryParamsResponse) {
	cmd := []string{"c4ed", "query", "cfevesting", "params", "--output=json"}

	out, _, err := n.containerManager.ExecCmd(n.t, n.Name, cmd, "")
	require.NoError(n.t, err)
	err = json.Unmarshal(out.Bytes(), &moduleParams)
	require.NoError(n.t, err)
}

func (n *NodeConfig) QueryCfeminterParams(moduleParams *cfemintermoduletypes.QueryParamsResponse) {
	cmd := []string{"c4ed", "query", "cfeminter", "params", "--output=json"}

	out, _, err := n.containerManager.ExecCmd(n.t, n.Name, cmd, "")
	require.NoError(n.t, err)
	err = util.Cdc.UnmarshalJSON(out.Bytes(), moduleParams)
	require.NoError(n.t, err)
}

func (n *NodeConfig) QueryCfedistributorParams(moduleParams *cfedistributormoduletypes.QueryParamsResponse) {
	cmd := []string{"c4ed", "query", "cfedistributor", "params", "--output=json"}

	out, _, err := n.containerManager.ExecCmd(n.t, n.Name, cmd, "")
	require.NoError(n.t, err)
	err = json.Unmarshal(out.Bytes(), &moduleParams)
	require.NoError(n.t, err)
}

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

	cmd := []string{"c4ed", "tx", "gov", "submit-proposal", ".c4e-chain/param_change_proposal.json", fmt.Sprintf("--from=%s", from)}

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

	cmd := []string{"c4ed", "tx", "gov", "submit-proposal", ".c4e-chain/param_change_proposal.json", fmt.Sprintf("--from=%s", from)}

	_, _, err = n.containerManager.ExecCmdWithResponseString(n.t, n.chainId, n.Name, cmd, errorMessage)
	require.NoError(n.t, err)
	err = os.Remove(localProposalFile)
	require.NoError(n.t, err)
}

func (n *NodeConfig) FailIBCTransfer(from, recipient, amount string) {
	n.LogActionF("IBC sending %s from %s to %s", amount, from, recipient)

	cmd := []string{"c4ed", "tx", "ibc-transfer", "transfer", "transfer", "channel-0", recipient, amount, fmt.Sprintf("--from=%s", from)}

	_, _, err := n.containerManager.ExecCmdWithResponseString(n.t, n.chainId, n.Name, cmd, "rate limit exceeded")
	require.NoError(n.t, err)

	n.LogActionF("Failed to send IBC transfer (as expected)")
}

func (n *NodeConfig) SubmitUpgradeProposal(upgradeVersion string, upgradeHeight int64, initialDeposit sdk.Coin) {
	n.LogActionF("submitting upgrade proposal %s for height %d", upgradeVersion, upgradeHeight)
	cmd := []string{"c4ed", "tx", "gov", "submit-legacy-proposal", "software-upgrade", upgradeVersion, fmt.Sprintf("--title=\"%s upgrade\"", upgradeVersion), "--description=\"upgrade proposal submission\"", fmt.Sprintf("--upgrade-height=%d", upgradeHeight), "--upgrade-info=\"\"", "--from=val", fmt.Sprintf("--deposit=%s", initialDeposit), "--no-validate"}
	fmt.Println(cmd)
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)
	n.LogActionF("successfully submitted upgrade proposal")
}
func (n *NodeConfig) SubmitTextProposal(text string, initialDeposit sdk.Coin, isExpedited bool) {
	n.LogActionF("submitting text gov proposal")
	cmd := []string{"c4ed", "tx", "gov", "submit-proposal", "--type=text", fmt.Sprintf("--title=\"%s\"", text), "--description=\"test text proposal\"", "--from=val", fmt.Sprintf("--deposit=%s", initialDeposit)}
	if isExpedited {
		cmd = append(cmd, "--is-expedited=true")
	}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)
	n.LogActionF("successfully submitted text gov proposal")
}

func (n *NodeConfig) DepositProposal(proposalNumber int) {
	n.LogActionF("depositing on proposal: %d", proposalNumber)
	deposit := sdk.NewCoin(params.CoinDenom, config.MinDepositValue)
	cmd := []string{"c4ed", "tx", "gov", "deposit", fmt.Sprintf("%d", proposalNumber), deposit.String(), "--from=val"}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)
	n.LogActionF("successfully deposited on proposal %d", proposalNumber)
}

func (n *NodeConfig) VoteYesProposal(from string, proposalNumber int) {
	n.LogActionF("voting yes on proposal: %d", proposalNumber)
	cmd := []string{"c4ed", "tx", "gov", "vote", fmt.Sprintf("%d", proposalNumber), "yes", fmt.Sprintf("--from=%s", from)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)
	n.LogActionF("successfully voted yes on proposal %d", proposalNumber)
}

func (n *NodeConfig) VoteNoProposal(from string, proposalNumber int) {
	n.LogActionF("voting no on proposal: %d", proposalNumber)
	cmd := []string{"c4ed", "tx", "gov", "vote", fmt.Sprintf("%d", proposalNumber), "no", fmt.Sprintf("--from=%s", from)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)
	n.LogActionF("successfully voted no on proposal: %d", proposalNumber)
}

func (n *NodeConfig) BankSend(amount string, sendAddress string, receiveAddress string) {
	n.LogActionF("bank sending %s from address %s to %s", amount, sendAddress, receiveAddress)
	cmd := []string{"c4ed", "tx", "bank", "send", sendAddress, receiveAddress, amount, "--from=val"}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)
	n.LogActionF("successfully sent bank sent %s from address %s to %s", amount, sendAddress, receiveAddress)
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

func (n *NodeConfig) QueryPropStatusTimed(proposalNumber int, desiredStatus string, totalTime chan time.Duration) {
	start := time.Now()
	require.Eventually(
		n.t,
		func() bool {
			status, err := n.QueryPropStatus(proposalNumber)
			if err != nil {
				return false
			}

			return status == desiredStatus
		},
		1*time.Minute,
		10*time.Millisecond,
		"C4e node failed to retrieve prop tally",
	)
	elapsed := time.Since(start)
	totalTime <- elapsed
}

func (n *NodeConfig) CreateVestingPool(vestingPoolName, amount, duration, vestinType, from string) {
	n.LogActionF("creating vesting pool")
	cmd := []string{"c4ed", "tx", "cfevesting", "create-vesting-pool", vestingPoolName, amount, duration, vestinType, fmt.Sprintf("--from=%s", from)}
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
	cmd := []string{"c4ed", "tx", "cfevesting", "withdraw-all-available", fmt.Sprintf("--from=%s", from)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)

	n.LogActionF("successfully withdrew all avaliable vestings")
}

func (n *NodeConfig) CreateVestingAccount(toAddress string, amount string, startTime, endTime, from string) {
	n.LogActionF("creating vesting account")
	cmd := []string{"c4ed", "tx", "cfevesting", "create-vesting-account", toAddress, amount, startTime, endTime, fmt.Sprintf("--from=%s", from)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)

	n.LogActionF("successfully created vesting account %s", toAddress)
}

func (n *NodeConfig) SplitVesting(toAddress string, amount string, from string) {
	n.LogActionF("split vesting")
	cmd := []string{"c4ed", "tx", "cfevesting", "split-vesting", toAddress, amount, fmt.Sprintf("--from=%s", from)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)

	n.LogActionF("successfully splitted vesting to account %s", toAddress)
}

func (n *NodeConfig) SplitVestingError(toAddress string, amount string, from, errorString string) {
	n.LogActionF("split vesting")
	cmd := []string{"c4ed", "tx", "cfevesting", "split-vesting", toAddress, amount, fmt.Sprintf("--from=%s", from)}
	_, _, err := n.containerManager.ExecCmdWithResponseString(n.t, n.chainId, n.Name, cmd, errorString)
	require.NoError(n.t, err)

	n.LogActionF("successfully splitted vesting to account %s", toAddress)
}

func (n *NodeConfig) MoveAvailableVesting(toAddress string, from string) {
	n.LogActionF("creating vesting account")
	cmd := []string{"c4ed", "tx", "cfevesting", "move-available-vesting", toAddress, fmt.Sprintf("--from=%s", from)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)

	n.LogActionF("successfully moved all available vesting to account %s", toAddress)
}

func (n *NodeConfig) MoveAvailableVestingError(toAddress, from, errorString string) {
	n.LogActionF("creating vesting account")
	cmd := []string{"c4ed", "tx", "cfevesting", "move-available-vesting", toAddress, fmt.Sprintf("--from=%s", from)}
	_, _, err := n.containerManager.ExecCmdWithResponseString(n.t, n.chainId, n.Name, cmd, errorString)
	require.NoError(n.t, err)

	n.LogActionF("successfully moved all available vesting to account %s", toAddress)
}

func (n *NodeConfig) MoveAvailableVestingByDenoms(toAddress string, denoms string, from string) {
	n.LogActionF("move available vesting by denoms")
	cmd := []string{"c4ed", "tx", "cfevesting", "move-available-vesting-by-denoms", toAddress, denoms, fmt.Sprintf("--from=%s", from)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)

	n.LogActionF("successfully moved all available vesting by denoms to account %s", toAddress)
}

func (n *NodeConfig) MoveAvailableVestingByDenomsError(toAddress string, denoms string, from, errorString string) {
	n.LogActionF("move available vesting by denoms")
	cmd := []string{"c4ed", "tx", "cfevesting", "move-available-vesting-by-denoms", toAddress, denoms, fmt.Sprintf("--from=%s", from)}
	_, _, err := n.containerManager.ExecCmdWithResponseString(n.t, n.chainId, n.Name, cmd, errorString)
	require.NoError(n.t, err)

	n.LogActionF("successfully moved all available vesting by denoms to account %s", toAddress)
}

func (n *NodeConfig) CreateCampaign(vestingPoolName, amount, duration, vestinType, from string) {
	n.LogActionF("creating campaign")
	cmd := []string{"c4ed", "tx", "cfevesting", "create-vesting-pool", vestingPoolName, amount, duration, vestinType, fmt.Sprintf("--from=%s", from)}
	_, _, err := n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)

	n.LogActionF("successfully created vesting pool %s", vestingPoolName)
}
