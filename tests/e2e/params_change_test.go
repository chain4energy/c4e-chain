package e2e

import (
	"encoding/json"
	"github.com/chain4energy/c4e-chain/tests/e2e/configurer/config"
	appparams "github.com/chain4energy/c4e-chain/tests/e2e/encoding/params"
	"github.com/chain4energy/c4e-chain/tests/e2e/initialization"
	cfedistributortypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	cfemintertypes "github.com/chain4energy/c4e-chain/x/cfeminter/types"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramsutils "github.com/cosmos/cosmos-sdk/x/params/client/utils"
	"github.com/stretchr/testify/suite"
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

func (s *ParamsSetupSuite) TestCfedistributorParamsProposal() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()

	newSubDistributors := []cfedistributortypes.SubDistributor{
		{
			Name: "New subdistributor",
			Sources: []*cfedistributortypes.Account{
				{
					Id:   cfedistributortypes.GreenEnergyBoosterCollector,
					Type: cfedistributortypes.MAIN,
				},
			},
			Destinations: cfedistributortypes.Destinations{
				PrimaryShare: cfedistributortypes.Account{
					Id:   cfedistributortypes.ValidatorsRewardsCollector,
					Type: cfedistributortypes.MODULE_ACCOUNT,
				},
				BurnShare: sdk.ZeroDec(),
			},
		},
	}

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
	chainA.LatestProposalNumber += 1

	for _, n := range chainA.NodeConfigs {
		n.VoteYesProposal(initialization.ValidatorWalletName, chainA.LatestProposalNumber)
	}

	s.Eventually(
		func() bool {
			return node.ValidateParams(newSubDistributorsJSON, cfedistributortypes.ModuleName, string(cfedistributortypes.KeySubDistributors))
		},
		time.Minute,
		time.Second*5,
		"C4e node failed to validate params",
	)
	totalSupplyBefore, err := node.QueryBalances(cfedistributortypes.ValidatorsRewardsCollector)
	s.NoError(err)
	time.Sleep(time.Second * 15)
	totalSupplyAfter, err := node.QueryBalances(cfedistributortypes.ValidatorsRewardsCollector)
	s.Equal(totalSupplyBefore, totalSupplyAfter)
	s.NoError(err)
}

func (s *ParamsSetupSuite) TestCfeminterParamsProposalNoMinting() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()

	newMinter := cfemintertypes.Minter{
		Start: time.Now().UTC(),
		Periods: []*cfemintertypes.MintingPeriod{
			{
				Position: 1,
				Type:     cfemintertypes.NO_MINTING,
			},
		},
	}

	newMinterJSON, err := json.Marshal(newMinter)
	s.NoError(err)
	newDenom := "newDenom"
	newMintDenomJSON, err := json.Marshal(newDenom)
	s.NoError(err)

	proposal := paramsutils.ParamChangeProposalJSON{
		Title:       "Cfeminter module params change",
		Description: "Change cfeminter params",
		Changes: paramsutils.ParamChangesJSON{
			paramsutils.ParamChangeJSON{
				Subspace: cfemintertypes.ModuleName,
				Key:      string(cfemintertypes.KeyMinter),
				Value:    newMinterJSON,
			},
			paramsutils.ParamChangeJSON{
				Subspace: cfemintertypes.ModuleName,
				Key:      string(cfemintertypes.KeyMintDenom),
				Value:    newMintDenomJSON,
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
			return node.ValidateParams(newMintDenomJSON, cfemintertypes.ModuleName, string(cfemintertypes.KeyMintDenom)) &&
				node.ValidateParams(newMinterJSON, cfemintertypes.ModuleName, string(cfemintertypes.KeyMinter))
		},
		time.Minute,
		time.Second*5,
		"C4e node failed to retrieve params",
	)

	totalSupplyBefore, err := node.QuerySupplyOf(newDenom)
	s.NoError(err)
	time.Sleep(time.Second * 15)
	totalSupplyAfter, err := node.QuerySupplyOf(newDenom)
	s.Equal(totalSupplyBefore, totalSupplyAfter)
	s.NoError(err)
}

func (s *ParamsSetupSuite) TestCfeminterParamsProposalLinearMinter() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	periodEnd := time.Now().Add(10 * time.Minute).UTC()
	newMinter := cfemintertypes.Minter{
		Start: time.Now().UTC(),
		Periods: []*cfemintertypes.MintingPeriod{
			{
				Position: 1,
				Type:     cfemintertypes.TIME_LINEAR_MINTER,
				TimeLinearMinter: &cfemintertypes.TimeLinearMinter{
					Amount: sdk.NewInt(100000),
				},
				PeriodEnd: &periodEnd,
			},
			{
				Position: 2,
				Type:     cfemintertypes.NO_MINTING,
			},
		},
	}

	newMinterJSON, err := json.Marshal(newMinter)
	s.NoError(err)

	newDenom := "newDenom2"
	newMintDenomJSON, err := json.Marshal(newDenom)
	s.NoError(err)

	proposal := paramsutils.ParamChangeProposalJSON{
		Title:       "Cfeminter module params change",
		Description: "Change cfeminter params",
		Changes: paramsutils.ParamChangesJSON{
			paramsutils.ParamChangeJSON{
				Subspace: cfemintertypes.ModuleName,
				Key:      string(cfemintertypes.KeyMinter),
				Value:    newMinterJSON,
			},
			paramsutils.ParamChangeJSON{
				Subspace: cfemintertypes.ModuleName,
				Key:      string(cfemintertypes.KeyMintDenom),
				Value:    newMintDenomJSON,
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
			return node.ValidateParams(newMintDenomJSON, cfemintertypes.ModuleName, string(cfemintertypes.KeyMintDenom)) &&
				node.ValidateParams(newMinterJSON, cfemintertypes.ModuleName, string(cfemintertypes.KeyMinter))
		},
		time.Minute,
		time.Second*5,
		"C4e node failed to retrieve params",
	)

	totalSupplyBefore, err := node.QuerySupplyOf(newDenom)
	s.NoError(err)
	time.Sleep(time.Second * 15)
	totalSupplyAfter, err := node.QuerySupplyOf(newDenom)
	s.True(totalSupplyAfter.GT(totalSupplyBefore))
	s.NoError(err)
}

func (s *ParamsSetupSuite) TestCfeVestingProposal() {
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
