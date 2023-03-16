package e2e

import (
	"cosmossdk.io/math"
	appparams "github.com/chain4energy/c4e-chain/app/params"
	"github.com/chain4energy/c4e-chain/tests/e2e/configurer/chain"
	"github.com/chain4energy/c4e-chain/tests/e2e/helpers"
	"github.com/chain4energy/c4e-chain/tests/e2e/initialization"
	"github.com/chain4energy/c4e-chain/tests/e2e/util"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	testhelpers "github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	cfedistributortypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	cfemintertypes "github.com/chain4energy/c4e-chain/x/cfeminter/types"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
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

func (s *ParamsSetupSuite) TestMinterAndDistributorCustom() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	endTime := time.Now().Add(10 * time.Minute).UTC()
	newDenom := "newTestDenom"
	linearMinterConfig, _ := codectypes.NewAnyWithValue(&cfemintertypes.LinearMinting{Amount: sdk.NewInt(100000)})
	noMintingConfig, _ := codectypes.NewAnyWithValue(&cfemintertypes.NoMinting{})

	updateMintersParams := cfemintertypes.MsgUpdateParams{
		Authority: appparams.GetAuthority(),
		MintDenom: newDenom,
		StartTime: time.Now().UTC(),
		Minters: []*cfemintertypes.Minter{
			{
				SequenceId: 1,
				EndTime:    &endTime,
				Config:     linearMinterConfig,
			},
			{
				SequenceId: 2,
				Config:     noMintingConfig,
			},
		},
	}

	proposalJSON, err := util.NewProposalJSON([]sdk.Msg{&updateMintersParams})
	s.NoError(err)
	node.SubmitDepositAndVoteOnProposal(proposalJSON, initialization.ValidatorWalletName, chainA)

	var params cfemintertypes.QueryParamsResponse
	node.QueryCfeminterParams(&params)
	s.ValidateNewMinterParams(node, updateMintersParams.Minters, updateMintersParams.StartTime, newDenom, true)

	updateDistributorParams := cfedistributortypes.MsgUpdateParams{
		Authority: appparams.GetAuthority(),
		SubDistributors: []cfedistributortypes.SubDistributor{{
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
				Shares:    []*cfedistributortypes.DestinationShare{},
			},
		},
		},
	}

	proposalJSON, err = util.NewProposalJSON([]sdk.Msg{&updateDistributorParams})
	s.NoError(err)
	node.SubmitDepositAndVoteOnProposal(proposalJSON, initialization.ValidatorWalletName, chainA)
	s.ValidateSubdistributorParams(node, updateDistributorParams.SubDistributors)
	s.validateBalanceOfAccount(node, newDenom, "c4e1q5vgy0r3w9q4cclucr2kl8nwmfe2mgr6g0jlph", true, 15)
}

func (s *ParamsSetupSuite) TestMinterAndDistributorMainnetShort() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)
	newDenom := "MinterAndDistributorMainnetShortDenom2"

	updateMintersParams := cfemintertypes.MsgUpdateParams{
		Authority: appparams.GetAuthority(),
		MintDenom: newDenom,
		StartTime: helpers.MainnetMinterConfigShort.StartTime,
		Minters:   helpers.MainnetMinterConfigShort.Minters,
	}

	updateDistributorParams := cfedistributortypes.MsgUpdateParams{
		Authority:       appparams.GetAuthority(),
		SubDistributors: helpers.MainnetSubdistributors,
	}

	proposalJSON, err := util.NewProposalJSON([]sdk.Msg{&updateDistributorParams})
	s.NoError(err)
	node.SubmitDepositAndVoteOnProposal(proposalJSON, initialization.ValidatorWalletName, chainA)
	s.ValidateSubdistributorParams(node, updateDistributorParams.SubDistributors)

	proposalJSON, err = util.NewProposalJSON([]sdk.Msg{&updateMintersParams})
	s.NoError(err)
	node.SubmitDepositAndVoteOnProposal(proposalJSON, initialization.ValidatorWalletName, chainA)

	totalSupplyBefore, err := node.QuerySupplyOf(appparams.CoinDenom)
	time.Sleep(time.Minute * 1)
	totalSupplyAfter, err := node.QuerySupplyOf(newDenom)
	s.Greater(totalSupplyAfter.Int64(), totalSupplyBefore.Int64())

	greenEnergyBoosterAddress := helpers.GetModuleAccountAddress(cfedistributortypes.GreenEnergyBoosterCollector)
	greenEnergyBoosterBalanceAfter, err := node.QueryBalances(greenEnergyBoosterAddress)
	s.Equal(sdk.NewDec(totalSupplyAfter.Int64()).Mul(sdk.MustNewDecFromStr("0.3")).Mul(sdk.MustNewDecFromStr("0.67")).TruncateInt(), greenEnergyBoosterBalanceAfter.AmountOf(newDenom))
}

func (s *ParamsSetupSuite) TestCfevestingEmptyDenom() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	proposalMessage := cfevestingtypes.MsgUpdateDenomParam{
		Authority: appparams.GetAuthority(),
		Denom:     "",
	}
	proposalJSON, err := util.NewProposalJSON([]sdk.Msg{&proposalMessage})
	s.NoError(err)

	node.SubmitParamChangeNotValidProposal(proposalJSON, initialization.ValidatorWalletName, "denom cannot be empty: invalid proposal message")
	node.QueryFailedProposal(chainA.LatestProposalNumber + 1)
}

func (s *ParamsSetupSuite) TestCfevestingNewDenom() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()

	newVestingDenom := "newDenom"
	s.NoError(err)

	proposalMessage := cfevestingtypes.MsgUpdateDenomParam{
		Authority: appparams.GetAuthority(),
		Denom:     newVestingDenom,
	}
	proposalJSON, err := util.NewProposalJSON([]sdk.Msg{&proposalMessage})
	s.NoError(err)

	node.SubmitDepositAndVoteOnProposal(proposalJSON, initialization.ValidatorWalletName, chainA)

	s.Eventually(
		func() bool {
			var params cfevestingtypes.QueryParamsResponse
			node.QueryCfevestingParams(&params)
			return s.Equal(params.Params.Denom, newVestingDenom)
		},
		time.Minute,
		time.Second*5,
		"C4e node failed to validate params",
	)
}

func (s *ParamsSetupSuite) TestCfevestingNewDenomVestingPoolsExist() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()

	// set previous denom
	proposalMessage := cfevestingtypes.MsgUpdateDenomParam{
		Authority: appparams.GetAuthority(),
		Denom:     appparams.CoinDenom,
	}
	proposalJSON, err := util.NewProposalJSON([]sdk.Msg{&proposalMessage})
	s.NoError(err)

	node.SubmitDepositAndVoteOnProposal(proposalJSON, initialization.ValidatorWalletName, chainA)

	// transfer funds and create vesting pool
	creatorWalletName := testhelpers.RandStringOfLength(10)
	creatorAddress := node.CreateWallet(creatorWalletName)
	node.BankSend(sdk.NewCoin(appparams.CoinDenom, sdk.NewInt(baseBalance)).String(), chainA.NodeConfigs[0].PublicAddress, creatorAddress)
	balanceBefore, err := node.QueryBalances(creatorAddress)

	vestingTypes := node.QueryVestingTypes()
	s.NoError(err)
	balanceBeforeAmount := balanceBefore.AmountOf(appparams.CoinDenom)
	randVestingPoolName := testhelpers.RandStringOfLength(5)
	vestingAmount := balanceBeforeAmount.Quo(sdk.NewInt(4))
	node.CreateVestingPool(randVestingPoolName, vestingAmount.String(), (10 * time.Minute).String(), vestingTypes[0].Name, creatorWalletName)

	// submit new proposal
	proposalMessage = cfevestingtypes.MsgUpdateDenomParam{
		Authority: appparams.GetAuthority(),
		Denom:     "abcNewDenom",
	}
	proposalJSON, err = util.NewProposalJSON([]sdk.Msg{&proposalMessage})
	s.NoError(err)
	node.SubmitParamChangeProposal(proposalJSON, initialization.ValidatorWalletName)
	chainA.LatestProposalNumber += 1
	node.DepositProposal(chainA.LatestProposalNumber)
	for _, n := range chainA.NodeConfigs {
		n.VoteYesProposal(initialization.ValidatorWalletName, chainA.LatestProposalNumber)
	}

	s.Eventually(
		func() bool {
			status, _ := node.QueryPropStatus(chainA.LatestProposalNumber)
			return status == "PROPOSAL_STATUS_FAILED"
		},
		time.Minute,
		time.Second*5,
		"C4e node failed to validate params",
	)
}

func (s *ParamsSetupSuite) TestCfeminterParamsProposalNoMinting() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)
	config, _ := codectypes.NewAnyWithValue(&cfemintertypes.NoMinting{})
	updateMintersParams := cfemintertypes.MsgUpdateMintersParams{
		Authority: appparams.GetAuthority(),
		StartTime: time.Now().UTC(),
		Minters: []*cfemintertypes.Minter{
			{
				SequenceId: 1,
				Config:     config,
			},
		},
	}

	proposalJSON, err := util.NewProposalJSON([]sdk.Msg{&updateMintersParams})
	s.NoError(err)
	node.SubmitDepositAndVoteOnProposal(proposalJSON, initialization.ValidatorWalletName, chainA)

	var params cfemintertypes.QueryParamsResponse
	node.QueryCfeminterParams(&params)
	s.ValidateNewMinterParams(node, updateMintersParams.Minters, updateMintersParams.StartTime, params.Params.MintDenom, false)
}

func (s *ParamsSetupSuite) TestCfeminterEmptyDenom() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)
	noMintingConfig, _ := codectypes.NewAnyWithValue(&cfemintertypes.NoMinting{})
	linearMintingConfig, _ := codectypes.NewAnyWithValue(&cfemintertypes.LinearMinting{Amount: math.NewInt(100000)})
	endTime := time.Now().Add(10 * time.Minute).UTC()
	startTime := time.Now().UTC()
	minters := []*cfemintertypes.Minter{
		{
			SequenceId: 1,
			Config:     linearMintingConfig,
			EndTime:    &endTime,
		},
		{
			SequenceId: 2,
			Config:     noMintingConfig,
		},
	}

	proposalMessage := cfemintertypes.MsgUpdateParams{
		Authority: appparams.GetAuthority(),
		MintDenom: "",
		StartTime: startTime,
		Minters:   minters,
	}
	proposalJSON, err := util.NewProposalJSON([]sdk.Msg{&proposalMessage})
	s.NoError(err)

	node.SubmitParamChangeNotValidProposal(proposalJSON, initialization.ValidatorWalletName, "validation error: denom cannot be empty: invalid proposal content: invalid proposal message")
	node.QueryFailedProposal(chainA.LatestProposalNumber + 1)
}

func (s *ParamsSetupSuite) TestCfeminterNoMinters() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	startTime := time.Now().UTC()
	var minters []*cfemintertypes.Minter

	proposalMessage := cfemintertypes.MsgUpdateParams{
		Authority: appparams.GetAuthority(),
		MintDenom: testenv.DefaultTestDenom,
		StartTime: startTime,
		Minters:   minters,
	}
	proposalJSON, err := util.NewProposalJSON([]sdk.Msg{&proposalMessage})
	s.NoError(err)

	node.SubmitParamChangeNotValidProposal(proposalJSON, initialization.ValidatorWalletName, "validation error: no minters defined: invalid proposal content: invalid proposal message")
	node.QueryFailedProposal(chainA.LatestProposalNumber + 1)
}

func (s *ParamsSetupSuite) TestCfedistributorNoSubdistributors() {
	chainA := s.configurer.GetChainConfig(0)
	node, err := chainA.GetDefaultNode()
	s.NoError(err)

	updateDistributorParams := cfedistributortypes.MsgUpdateParams{
		Authority:       appparams.GetAuthority(),
		SubDistributors: []cfedistributortypes.SubDistributor{},
	}

	proposalJSON, err := util.NewProposalJSON([]sdk.Msg{&updateDistributorParams})
	s.NoError(err)

	node.SubmitParamChangeNotValidProposal(proposalJSON, initialization.ValidatorWalletName, "validation error: there must be at least one subdistributor with the source main type: invalid proposal content: invalid proposal message")
	node.QueryFailedProposal(chainA.LatestProposalNumber + 1)
}

func (s *ParamsSetupSuite) ValidateNewMinterParams(node *chain.NodeConfig, minters []*cfemintertypes.Minter, startTime time.Time, mintDenom string, totalSupplyIncreasing bool) {
	var params cfemintertypes.QueryParamsResponse
	s.validateTotalSupply(node, mintDenom, totalSupplyIncreasing, 25)
	s.Eventually(
		func() bool {
			node.QueryCfeminterParams(&params)
			paramsMinters := params.Params.Minters
			if !(len(minters) == len(paramsMinters)) {
				return false
			}
			if !startTime.Equal(params.Params.StartTime) {
				return false
			}
			if mintDenom != params.Params.MintDenom {
				return false
			}
			for i, minter := range minters {
				minterFromParams := paramsMinters[i]
				if minter.EndTime == nil {
					if minterFromParams.EndTime != nil {
						return false
					}
				} else {
					if !minter.EndTime.Equal(*minterFromParams.EndTime) {
						return false
					}
				}
				if minter.SequenceId != minterFromParams.SequenceId {
					return false
				}
				if !minter.Config.Equal(minterFromParams.Config) {
					return false
				}
			}
			return true
		},
		time.Minute,
		time.Second*5,
		"C4e node failed to validate params",
	)
}

func (s *ParamsSetupSuite) ValidateSubdistributorParams(node *chain.NodeConfig, subDistributors []cfedistributortypes.SubDistributor) {
	var params cfedistributortypes.QueryParamsResponse
	s.Eventually(
		func() bool {
			node.QueryCfedistributorParams(&params)
			return assert.ObjectsAreEqualValues(subDistributors, params.Params.SubDistributors)
		},
		time.Minute,
		time.Second*5,
		"C4e node failed to validate params",
	)
}
