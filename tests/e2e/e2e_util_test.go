package e2e

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/chain4energy/c4e-chain/tests/e2e/initialization"
	"github.com/chain4energy/c4e-chain/tests/e2e/util"
	"github.com/stretchr/testify/require"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/ory/dockertest/v3/docker"
)

func (s *IntegrationTestSuite) ExecTx(chainId string, validatorIndex int, command []string, success string) (bytes.Buffer, bytes.Buffer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	var containerId string
	if chainId == "" {
		containerId = s.hermesResource.Container.ID
	} else {
		containerId = s.valResources[chainId][validatorIndex].Container.ID
	}

	var (
		outBuf bytes.Buffer
		errBuf bytes.Buffer
	)

	s.Require().Eventually(
		func() bool {
			exec, err := s.dkrPool.Client.CreateExec(docker.CreateExecOptions{
				Context:      ctx,
				AttachStdout: true,
				AttachStderr: true,
				Container:    containerId,
				User:         "root",
				Cmd:          command,
			})
			s.Require().NoError(err)

			err = s.dkrPool.Client.StartExec(exec.ID, docker.StartExecOptions{
				Context:      ctx,
				Detach:       false,
				OutputStream: &outBuf,
				ErrorStream:  &errBuf,
			})
			if err != nil {
				return false
			}

			if success != "" {
				return strings.Contains(outBuf.String(), success) || strings.Contains(errBuf.String(), success)
			}

			return true
		},
		time.Minute,
		time.Second,
		"tx returned a non-zero code; stdout: %s, stderr: %s", outBuf.String(), errBuf.String(),
	)

	return outBuf, errBuf, nil
}

func (s *IntegrationTestSuite) ExecQueryRPC(path string) ([]byte, error) {
	var err error
	var resp *http.Response
	retriesLeft := 5
	for {
		resp, err = http.Get(path)

		if resp.StatusCode == http.StatusServiceUnavailable {
			retriesLeft--
			if retriesLeft == 0 {
				return nil, err
			}
			time.Sleep(10 * time.Second)
		} else {
			break
		}
	}

	if err != nil {
		return nil, fmt.Errorf("failed to execute HTTP request: %w", err)
	}

	defer resp.Body.Close()

	bz, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bz, nil
}

func (s *IntegrationTestSuite) connectIBCChains(chainA *chainConfig, chainB *chainConfig) {
	s.T().Logf("connecting %s and %s chains via IBC", chainA.meta.Id, chainB.meta.Id)
	cmd := []string{"hermes", "create", "channel", chainA.meta.Id, chainB.meta.Id, "--port-a=transfer", "--port-b=transfer"}
	s.ExecTx("", 0, cmd, "successfully opened init channel")
	s.T().Logf("connected %s and %s chains via IBC", chainA.meta.Id, chainB.meta.Id)
}

func (s *IntegrationTestSuite) sendIBC(srcChain *chainConfig, dstChain *chainConfig, recipient string, token sdk.Coin) {
	cmd := []string{"hermes", "tx", "raw", "ft-transfer", dstChain.meta.Id, srcChain.meta.Id, "transfer", "channel-0", token.Amount.String(), fmt.Sprintf("--denom=%s", token.Denom), fmt.Sprintf("--receiver=%s", recipient), "--timeout-height-offset=1000"}
	s.ExecTx("", 0, cmd, "Success")

	s.T().Logf("sending %s from %s to %s (%s)", token, srcChain.meta.Id, dstChain.meta.Id, recipient)
	balancesBPre, err := s.queryBalances(dstChain, 0, recipient)
	s.Require().NoError(err)

	s.Require().Eventually(
		func() bool {
			balancesBPost, err := s.queryBalances(dstChain, 0, recipient)
			s.Require().NoError(err)
			ibcCoin := balancesBPost.Sub(balancesBPre)
			if ibcCoin.Len() == 1 {
				tokenPre := balancesBPre.AmountOfNoDenomValidation(ibcCoin[0].Denom)
				tokenPost := balancesBPost.AmountOfNoDenomValidation(ibcCoin[0].Denom)
				resPre := initialization.OsmoToken.Amount
				resPost := tokenPost.Sub(tokenPre)
				return resPost.Uint64() == resPre.Uint64()
			} else {
				return false
			}
		},
		5*time.Minute,
		time.Second,
		"tx not received on destination chain",
	)

	s.T().Log("successfully sent IBC tokens")
}

func (s *IntegrationTestSuite) submitUpgradeProposal(c *chainConfig) {
	upgradeHeightStr := strconv.Itoa(c.propHeight)
	s.T().Logf("submitting upgrade proposal on %s container: %s", s.valResources[c.meta.Id][0].Container.Name[1:], s.valResources[c.meta.Id][0].Container.ID)
	cmd := []string{"c4ed", "tx", "gov", "submit-proposal", "software-upgrade", upgradeVersion, fmt.Sprintf("--title=\"%s upgrade\"", upgradeVersion), "--description=\"upgrade proposal submission\"", fmt.Sprintf("--upgrade-height=%s", upgradeHeightStr), "--upgrade-info=\"\"", fmt.Sprintf("--chain-id=%s", c.meta.Id), "--from=val", "-b=block", "--yes", "--keyring-backend=test", "--log_format=json"}
	s.ExecTx(c.meta.Id, 0, cmd, "code: 0")
	s.T().Log("successfully submitted upgrade proposal")
	c.latestProposalNumber = c.latestProposalNumber + 1
}

func (s *IntegrationTestSuite) insertUpgradeProposalToContainer(c *chainConfig) {
	bytes, _ := ioutil.ReadFile("./scripts/update-subdistributors.json")
	proposalString := string(bytes)
	s.T().Logf("inserting params upgrade file to %s container: %s", s.valResources[c.meta.Id][0].Container.Name[1:], s.valResources[c.meta.Id][0].Container.ID)
	cmd := []string{"echo", proposalString, ">", "update-subdistributors.json"}
	s.ExecTx(c.meta.Id, 0, cmd, "code: 0")
	s.T().Log("successfully inserted params upgrade file")
}

func (s *IntegrationTestSuite) submitUpgradeParamsProposal(c *chainConfig, pathToProposalFile string) {
	upgradeHeightStr := strconv.Itoa(c.propHeight)
	s.T().Logf("submitting upgrade params on %s container: %s. Path to proposal file: pathToProposalFile", s.valResources[c.meta.Id][0].Container.Name[1:], s.valResources[c.meta.Id][0].Container.ID)
	cmd := []string{"c4ed", "tx", "gov", "submit-proposal", "param-change", pathToProposalFile, fmt.Sprintf("--title=\"%s upgrade params: \"", pathToProposalFile), "--description=\"upgrade proposal params\"", fmt.Sprintf("--upgrade-height=%s", upgradeHeightStr), "--upgrade-info=\"\"", fmt.Sprintf("--chain-id=%s", c.meta.Id), "--from=val", "-b=block", "--yes", "--keyring-backend=test", "--log_format=json"}
	s.ExecTx(c.meta.Id, 0, cmd, "code: 0")
	s.T().Log("successfully submitted upgrade proposal")
	c.latestProposalNumber = c.latestProposalNumber + 1
}

func (s *IntegrationTestSuite) submitTextProposal(c *chainConfig, text string) {
	s.T().Logf("submitting text proposal on %s container: %s", s.valResources[c.meta.Id][0].Container.Name[1:], s.valResources[c.meta.Id][0].Container.ID)
	cmd := []string{"c4ed", "tx", "gov", "submit-proposal", "--type=text", fmt.Sprintf("--title=\"%s\"", text), "--description=\"test text proposal\"", "--from=val", "-b=block", "--yes", "--keyring-backend=test", "--log_format=json", fmt.Sprintf("--chain-id=%s", c.meta.Id)}
	s.ExecTx(c.meta.Id, 0, cmd, "code: 0")
	s.T().Log("successfully submitted text proposal")
	c.latestProposalNumber = c.latestProposalNumber + 1
}

func (s *IntegrationTestSuite) SubmitParamChangeProposal(proposalJson, from string) {
	s.T().Logf("submitting param change proposal %s", proposalJson)
	wd, err := os.Getwd()
	require.NoError(n.t, err)
	localProposalFile := wd + "/scripts/param_change_proposal.json"
	f, err := os.Create(localProposalFile)
	require.NoError(n.t, err)
	_, err = f.WriteString(proposalJson)
	require.NoError(n.t, err)
	err = f.Close()
	require.NoError(n.t, err)

	cmd := []string{"osmosisd", "tx", "gov", "submit-proposal", "param-change", "/osmosis/param_change_proposal.json", fmt.Sprintf("--from=%s", from)}

	_, _, err = n.containerManager.ExecTxCmd(n.t, n.chainId, n.Name, cmd)
	require.NoError(n.t, err)

	err = os.Remove(localProposalFile)
	require.NoError(n.t, err)

	n.LogActionF("successfully submitted param change proposal")
}

func (s *IntegrationTestSuite) depositProposal(c *chainConfig) {
	propStr := strconv.Itoa(c.latestProposalNumber)
	s.T().Logf("depositing to proposal from %s container: %s", s.valResources[c.meta.Id][0].Container.Name[1:], s.valResources[c.meta.Id][0].Container.ID)
	cmd := []string{"c4ed", "tx", "gov", "deposit", propStr, "500000000uc4e", "--from=val", fmt.Sprintf("--chain-id=%s", c.meta.Id), "-b=block", "--yes", "--keyring-backend=test"}
	s.ExecTx(c.meta.Id, 0, cmd, "code: 0")
	s.T().Log("successfully deposited to proposal")
}

func (s *IntegrationTestSuite) voteProposal(c *chainConfig) {
	propStr := strconv.Itoa(c.latestProposalNumber)
	s.T().Logf("voting yes on proposal for chain-id: %s", c.meta.Id)
	cmd := []string{"c4ed", "tx", "gov", "vote", propStr, "yes", "--from=val", fmt.Sprintf("--chain-id=%s", c.meta.Id), "-b=block", "--yes", "--keyring-backend=test"}
	for i := range c.validators {
		if _, ok := c.skipRunValidatorIndexes[i]; ok {
			continue
		}
		s.ExecTx(c.meta.Id, i, cmd, "code: 0")
		s.T().Logf("successfully voted yes on proposal from %s container: %s", s.valResources[c.meta.Id][i].Container.Name[1:], s.valResources[c.meta.Id][i].Container.ID)
	}
}

func (s *IntegrationTestSuite) voteNoProposal(c *chainConfig, i int, from string) {
	propStr := strconv.Itoa(c.latestProposalNumber)
	s.T().Logf("voting no on proposal for chain-id: %s", c.meta.Id)
	cmd := []string{"c4ed", "tx", "gov", "vote", propStr, "no", fmt.Sprintf("--from=%s", from), fmt.Sprintf("--chain-id=%s", c.meta.Id), "-b=block", "--yes", "--keyring-backend=test"}
	s.ExecTx(c.meta.Id, i, cmd, "code: 0")
	s.T().Logf("successfully voted no for proposal from %s container: %s", s.valResources[c.meta.Id][i].Container.Name[1:], s.valResources[c.meta.Id][i].Container.ID)
}

func (s *IntegrationTestSuite) chainStatus(c *chainConfig, i int) []byte {
	cmd := []string{"c4ed", "status"}
	outBuff, _, err := s.ExecTx(c.meta.Id, i, cmd, "")
	s.Require().NoError(err)
	return outBuff.Bytes()
}

func (s *IntegrationTestSuite) getCurrentChainHeight(c *chainConfig, i int) int {
	var block syncInfo
	s.Require().Eventually(
		func() bool {
			out := s.chainStatus(c, i)
			err := json.Unmarshal(out, &block)
			if err != nil {
				return false
			}
			return true
		},
		1*time.Minute,
		time.Second,
		"Osmosis node failed to retrieve height info",
	)
	currentHeight, err := strconv.Atoi(block.SyncInfo.LatestHeight)
	s.Require().NoError(err)
	return currentHeight
}

func (s *IntegrationTestSuite) queryBalances(c *chainConfig, i int, addr string) (sdk.Coins, error) {
	cmd := []string{"c4ed", "query", "bank", "balances", addr, "--output=json"}
	outBuf, _, err := s.ExecTx(c.meta.Id, i, cmd, "")
	s.Require().NoError(err)

	var balancesResp banktypes.QueryAllBalancesResponse
	err = util.Cdc.UnmarshalJSON(outBuf.Bytes(), &balancesResp)
	s.Require().NoError(err)

	return balancesResp.GetBalances(), nil
}

func (s *IntegrationTestSuite) queryPropTally(endpoint, addr string) (sdk.Int, sdk.Int, sdk.Int, sdk.Int, error) {
	path := fmt.Sprintf(
		"%s/cosmos/gov/v1beta1/proposals/%s/tally",
		endpoint, addr,
	)
	bz, err := s.ExecQueryRPC(path)
	s.Require().NoError(err)

	var balancesResp govtypes.QueryTallyResultResponse
	if err := util.Cdc.UnmarshalJSON(bz, &balancesResp); err != nil {
		return sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt(), err
	}
	noTotal := balancesResp.Tally.No
	yesTotal := balancesResp.Tally.Yes
	noWithVetoTotal := balancesResp.Tally.NoWithVeto
	abstainTotal := balancesResp.Tally.Abstain

	return noTotal, yesTotal, noWithVetoTotal, abstainTotal, nil
}

func (s *IntegrationTestSuite) sendTx(c *chainConfig, i int, amount string, sendAddress string, receiveAddress string) {
	s.T().Logf("sending %s from %s to %s on chain-id: %s", amount, sendAddress, receiveAddress, c.meta.Id)
	cmd := []string{"c4ed", "tx", "bank", "send", sendAddress, receiveAddress, amount, fmt.Sprintf("--chain-id=%s", c.meta.Id), "--from=val", "-b=block", "--yes", "--keyring-backend=test"}
	s.ExecTx(c.meta.Id, i, cmd, "code: 0")
	s.T().Logf("successfully sent tx from %s container: %s", s.valResources[c.meta.Id][i].Container.Name[1:], s.valResources[c.meta.Id][i].Container.ID)
}

func (s *IntegrationTestSuite) extractValidatorOperatorAddresses(config *chainConfig) {
	for i, val := range config.validators {
		if _, ok := config.skipRunValidatorIndexes[i]; ok {
			s.T().Logf("skipping %s validator with index %d from running...", val.validator.Name, i)
			continue
		}
		cmd := []string{"c4ed", "debug", "addr", val.validator.PublicKey}
		s.T().Logf("extracting validator operator addresses for chain-id: %s", config.meta.Id)
		_, errBuf, err := s.ExecTx(config.meta.Id, i, cmd, "")
		s.Require().NoError(err)
		re := regexp.MustCompile("osmovaloper(.{39})")
		operAddr := fmt.Sprintf("%s\n", re.FindString(errBuf.String()))
		config.validators[i].operatorAddress = strings.TrimSuffix(operAddr, "\n")
	}
}

func (s *IntegrationTestSuite) createWallet(c *chainConfig, index int, walletName string) string {
	cmd := []string{"c4ed", "keys", "add", walletName, "--keyring-backend=test"}
	outBuf, _, err := s.ExecTx(c.meta.Id, index, cmd, "")
	s.Require().NoError(err)
	re := regexp.MustCompile("osmo1(.{38})")
	walletAddr := fmt.Sprintf("%s\n", re.FindString(outBuf.String()))
	walletAddr = strings.TrimSuffix(walletAddr, "\n")
	return walletAddr
}
