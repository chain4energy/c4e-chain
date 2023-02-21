package e2e

import (
	"fmt"
	appparams "github.com/chain4energy/c4e-chain/app/params"
	"github.com/chain4energy/c4e-chain/tests/e2e/configurer/chain"
	"github.com/chain4energy/c4e-chain/tests/e2e/configurer/config"
	"github.com/chain4energy/c4e-chain/tests/e2e/helpers"
	"github.com/chain4energy/c4e-chain/tests/e2e/initialization"
	cfedistributortypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	cfemintertypes "github.com/chain4energy/c4e-chain/x/cfeminter/types"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramsutils "github.com/cosmos/cosmos-sdk/x/params/client/utils"
	"github.com/stretchr/testify/suite"
	"k8s.io/kubernetes/staging/src/k8s.io/apimachinery/pkg/util/json"
	"regexp"
	"testing"
	"time"
)

type ParamsSetupSuite struct {
	BaseSetupSuite
}

func TestParamsChangeSuite(t *testing.T) {
	suite.Run(t, new(ParamsSetupSuite))
}

func (s *ParamsSetupSuite) SetupSuite() {
	s.BaseSetupSuite.SetupSuite(false, false)
}

func (s *ParamsSetupSuite) TestMinterAndDistributorCustom() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	endTime := time.Now().Add(10 * time.Minute).UTC()
	newMinter := cfemintertypes.MinterConfig{
		StartTime: time.Now().UTC(),
		Minters: []*cfemintertypes.Minter{
			{
				SequenceId: 1,
				Type:       cfemintertypes.LinearMintingType,
				LinearMinting: &cfemintertypes.LinearMinting{
					Amount: sdk.NewInt(100000),
				},
				EndTime: &endTime,
			},
			{
				SequenceId: 2,
				Type:       cfemintertypes.NoMintingType,
			},
		},
	}
	newDenom := "newTestDenom"
	s.cfeminterParamsChange(node, chainA, newDenom, newMinter)
	s.validateTotalSupply(node, newDenom, true, 15)

	newSubDistributors := []cfedistributortypes.SubDistributor{
		{
			Name: "New subdistributor",
			Sources: []*cfedistributortypes.Account{
				{
					Id:   cfedistributortypes.DistributorMainAccount,
					Type: cfedistributortypes.Main,
				},
			},
			Destinations: cfedistributortypes.Destinations{
				PrimaryShare: cfedistributortypes.Account{
					Id:   "c4e1q5vgy0r3w9q4cclucr2kl8nwmfe2mgr6g0jlph",
					Type: cfedistributortypes.BaseAccount,
				},
				BurnShare: sdk.ZeroDec(),
			},
		},
	}
	s.cfedistributorParamsChange(node, chainA, newSubDistributors)

	s.validateBalanceOfAccount(node, newDenom, "c4e1q5vgy0r3w9q4cclucr2kl8nwmfe2mgr6g0jlph", true, 15)
}

func (s *ParamsSetupSuite) TestMinterAndDistributorMainnetShort() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)
	newDenom := "MinterAndDistributorMainnetShortDenom2"

	s.cfedistributorParamsChange(node, chainA, helpers.MainnetSubdistributors)
	s.cfeminterParamsChange(node, chainA, newDenom, helpers.MainnetMinterConfigShort)
	totalSupplyBefore, err := node.QuerySupplyOf(appparams.CoinDenom)
	time.Sleep(time.Minute * 1)
	totalSupplyAfter, err := node.QuerySupplyOf(newDenom)
	s.Greater(totalSupplyAfter.Int64(), totalSupplyBefore.Int64())
	greenEnergyBoosterAddress := helpers.GetModuleAccountAddress(cfedistributortypes.GreenEnergyBoosterCollector)
	greenEnergyBoosterBalanceAfter, err := node.QueryBalances(greenEnergyBoosterAddress)
	s.Equal(sdk.NewDec(totalSupplyAfter.Int64()).Mul(sdk.MustNewDecFromStr("0.3")).Mul(sdk.MustNewDecFromStr("0.67")).TruncateInt(), greenEnergyBoosterBalanceAfter.AmountOf(newDenom))
}

func (s *ParamsSetupSuite) TestCfeminterParamsProposalNoMinting() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	newMinter := cfemintertypes.MinterConfig{
		StartTime: time.Now().UTC(),
		Minters: []*cfemintertypes.Minter{
			{
				SequenceId: 1,
				Type:       cfemintertypes.NoMintingType,
			},
		},
	}

	newDenom := "newDenom"
	s.cfeminterParamsChange(node, chainA, newDenom, newMinter)
	s.validateTotalSupply(node, newDenom, false, 25)
}

func (s *ParamsSetupSuite) TestCfevestingNewDenom() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()

	newVestingDenom, err := json.Marshal("newDenom")
	s.NoError(err)

	proposal := paramsutils.ParamChangeProposalJSON{
		Title:       "Cfevesting module params change",
		Description: "Change cfevesting params",
		Changes: paramsutils.ParamChangesJSON{
			paramsutils.ParamChangeJSON{
				Subspace: cfevestingtypes.ModuleName,
				Key:      string(cfevestingtypes.KeyDenom),
				Value:    newVestingDenom,
			},
		},
		Deposit: sdk.NewCoin(appparams.CoinDenom, config.MinDepositValue).String(),
	}

	proposalJSON, err := json.Marshal(proposal)
	s.NoError(err)
	node.SubmitParamChangeProposal(string(proposalJSON), initialization.ValidatorWalletName)
	chainA.LatestProposalNumber += 1

	for _, n := range chainA.NodeConfigs {
		n.VoteYesProposal(initialization.ValidatorWalletName, chainA.LatestProposalNumber)
	}

	s.Eventually(
		func() bool {
			return node.ValidateParams(newVestingDenom, cfevestingtypes.ModuleName, string(cfevestingtypes.KeyDenom))
		},
		time.Minute,
		time.Second*5,
		"C4e node failed to validate params",
	)
}

func (s *ParamsSetupSuite) cfedistributorParamsChange(node *chain.NodeConfig, chainConfig *chain.Config, newSubDistributors []cfedistributortypes.SubDistributor) {
	newSubDistributorsJSON, err := json.Marshal(newSubDistributors)
	s.NoError(err)

	proposal := paramsutils.ParamChangeProposalJSON{
		Title:       "CfeDistributor module params change",
		Description: "Change cfedistributor params",
		Changes: paramsutils.ParamChangesJSON{
			paramsutils.ParamChangeJSON{
				Subspace: cfedistributortypes.ModuleName,
				Key:      string(cfedistributortypes.KeySubDistributors),
				Value:    newSubDistributorsJSON,
			},
		},
		Deposit: sdk.NewCoin(appparams.CoinDenom, config.MinDepositValue).String(),
	}

	proposalJSON, err := json.Marshal(proposal)
	s.NoError(err)
	node.SubmitParamChangeProposal(string(proposalJSON), initialization.ValidatorWalletName)
	chainConfig.LatestProposalNumber += 1

	for _, n := range chainConfig.NodeConfigs {
		n.VoteYesProposal(initialization.ValidatorWalletName, chainConfig.LatestProposalNumber)
	}

	s.Eventually(
		func() bool {
			return node.ValidateParams(newSubDistributorsJSON, cfedistributortypes.ModuleName, string(cfedistributortypes.KeySubDistributors))
		},
		time.Minute,
		time.Second*5,
		"C4e node failed to validate params",
	)
}

func (s *ParamsSetupSuite) cfeminterParamsChange(node *chain.NodeConfig, chainConfig *chain.Config, newDenom string, newMinterConfig cfemintertypes.MinterConfig) {
	newMinterJSON, err := json.Marshal(newMinterConfig)
	s.NoError(err)
	// we need to add double quotes around step_duration value because json.Marshall convert time.Duration to int64 and
	// cosmos sdk AminoJSON requires time.Duration to be in double quotes (same thing happens for example with sdk.Dev type)
	stepDurationRegex := regexp.MustCompile(`(step_duration)":([^,]*)`)
	newMinterJSON = []byte(stepDurationRegex.ReplaceAllString(string(newMinterJSON), "$1\":\"$2\""))

	newDenomJSON, err := json.Marshal(newDenom)
	s.NoError(err)

	proposal := paramsutils.ParamChangeProposalJSON{
		Title:       "Cfeminter module params change",
		Description: "Change cfeminter params",
		Changes: paramsutils.ParamChangesJSON{
			paramsutils.ParamChangeJSON{
				Subspace: cfemintertypes.ModuleName,
				Key:      string(cfemintertypes.KeyMinterConfig),
				Value:    newMinterJSON,
			},
			paramsutils.ParamChangeJSON{
				Subspace: cfemintertypes.ModuleName,
				Key:      string(cfemintertypes.KeyMintDenom),
				Value:    newDenomJSON,
			},
		},
		Deposit: sdk.NewCoin(appparams.CoinDenom, config.MinDepositValue).String(),
	}

	proposalJSON, err := json.Marshal(proposal)
	s.NoError(err)
	node.SubmitParamChangeProposal(string(proposalJSON), initialization.ValidatorWalletName)

	chainConfig.LatestProposalNumber += 1

	for _, n := range chainConfig.NodeConfigs {
		n.VoteYesProposal(initialization.ValidatorWalletName, chainConfig.LatestProposalNumber)
	}

	s.Eventually(
		func() bool {
			return node.ValidateParams(newDenomJSON, cfemintertypes.ModuleName, string(cfemintertypes.KeyMintDenom)) &&
				node.ValidateParams(newMinterJSON, cfemintertypes.ModuleName, string(cfemintertypes.KeyMinterConfig))
		},
		time.Minute,
		time.Second*5,
		"C4e node failed to retrieve params",
	)
}

func (s *ParamsSetupSuite) TestCfevestingEmptyDenom() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()

	newVestingDenom, err := json.Marshal("")
	s.NoError(err)

	proposal := paramsutils.ParamChangeProposalJSON{
		Title:       "Cfevesting module params change",
		Description: "Change cfevesting params",
		Changes: paramsutils.ParamChangesJSON{
			paramsutils.ParamChangeJSON{
				Subspace: cfevestingtypes.ModuleName,
				Key:      string(cfevestingtypes.KeyDenom),
				Value:    newVestingDenom,
			},
		},
		Deposit: sdk.NewCoin(appparams.CoinDenom, config.MinDepositValue).String(),
	}

	proposalJSON, err := json.Marshal(proposal)
	s.NoError(err)
	node.SubmitParamChangeNotValidProposal(string(proposalJSON), initialization.ValidatorWalletName, "invalid parameter value: denom cannot be empty")
	node.QueryFailedProposal(chainA.LatestProposalNumber + 1)
}

func (s *ParamsSetupSuite) TestCfeminterEmptyDenom() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	endTime := time.Now().Add(10 * time.Minute).UTC()
	newMinterConfig := cfemintertypes.MinterConfig{
		StartTime: time.Now().UTC(),
		Minters: []*cfemintertypes.Minter{
			{
				SequenceId: 1,
				Type:       cfemintertypes.LinearMintingType,
				LinearMinting: &cfemintertypes.LinearMinting{
					Amount: sdk.NewInt(100000),
				},
				EndTime: &endTime,
			},
			{
				SequenceId: 2,
				Type:       cfemintertypes.NoMintingType,
			},
		},
	}
	newDenomJSON, err := json.Marshal("")
	newMinterJSON, err := json.Marshal(newMinterConfig)
	s.NoError(err)

	proposal := paramsutils.ParamChangeProposalJSON{
		Title:       "Cfeminter module params change - empty denom",
		Description: "Change cfeminter params",
		Changes: paramsutils.ParamChangesJSON{
			paramsutils.ParamChangeJSON{
				Subspace: cfemintertypes.ModuleName,
				Key:      string(cfemintertypes.KeyMinterConfig),
				Value:    newMinterJSON,
			},
			paramsutils.ParamChangeJSON{
				Subspace: cfemintertypes.ModuleName,
				Key:      string(cfemintertypes.KeyMintDenom),
				Value:    newDenomJSON,
			},
		},
		Deposit: sdk.NewCoin(appparams.CoinDenom, config.MinDepositValue).String(),
	}

	proposalJSON, err := json.Marshal(proposal)
	node.SubmitParamChangeNotValidProposal(string(proposalJSON), initialization.ValidatorWalletName, "invalid parameter value: denom cannot be empty")
	node.QueryFailedProposal(chainA.LatestProposalNumber + 1)
}

func (s *ParamsSetupSuite) TestCfeminterNoMinters() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	newMinterConfig := cfemintertypes.MinterConfig{
		StartTime: time.Now().UTC(),
	}
	newMinterConfig.Minters = make([]*cfemintertypes.Minter, 1)

	newDenomJSON, err := json.Marshal("newDenom")
	newMinterJSON, err := json.Marshal(newMinterConfig)

	fmt.Println(string(newMinterJSON))

	s.NoError(err)

	proposal := paramsutils.ParamChangeProposalJSON{
		Title:       "Cfeminter module params change - empty denom",
		Description: "Change cfeminter params",
		Changes: paramsutils.ParamChangesJSON{
			paramsutils.ParamChangeJSON{
				Subspace: cfemintertypes.ModuleName,
				Key:      string(cfemintertypes.KeyMinterConfig),
				Value:    newMinterJSON,
			},
			paramsutils.ParamChangeJSON{
				Subspace: cfemintertypes.ModuleName,
				Key:      string(cfemintertypes.KeyMintDenom),
				Value:    newDenomJSON,
			},
		},
		Deposit: sdk.NewCoin(appparams.CoinDenom, config.MinDepositValue).String(),
	}

	proposalJSON, err := json.Marshal(proposal)
	node.SubmitParamChangeNotValidProposal(string(proposalJSON), initialization.ValidatorWalletName, "invalid parameter value: minter on position 1 cannot be nil")
	node.QueryFailedProposal(chainA.LatestProposalNumber + 1)
}

func (s *ParamsSetupSuite) TestCfedistributorNoSubdistributors() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	newSubDistributors := []cfedistributortypes.SubDistributor{}
	newSubDistributorsJSON, err := json.Marshal(newSubDistributors)

	s.NoError(err)
	proposal := paramsutils.ParamChangeProposalJSON{
		Title:       "CfeDistributor module params change",
		Description: "Change cfedistributor params",
		Changes: paramsutils.ParamChangesJSON{
			paramsutils.ParamChangeJSON{
				Subspace: cfedistributortypes.ModuleName,
				Key:      string(cfedistributortypes.KeySubDistributors),
				Value:    newSubDistributorsJSON,
			},
		},
		Deposit: sdk.NewCoin(appparams.CoinDenom, config.MinDepositValue).String(),
	}

	proposalJSON, err := json.Marshal(proposal)
	node.SubmitParamChangeNotValidProposal(string(proposalJSON), initialization.ValidatorWalletName, "invalid parameter value: there must be at least one subdistributor with the source main type")
	node.QueryFailedProposal(chainA.LatestProposalNumber + 1)
}
