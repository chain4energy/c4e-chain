package cfedistributor_test

import (
	"testing"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

type DestinationType int64

const (
	MainCollector DestinationType = iota
	ModuleAccount
	InternalAccount
	BaseAccount
)

const c4eDistributorCollectorName = types.GreenEnergyBoosterCollector
const noValidatorsCollectorName = types.GovernanceBoosterCollector

var accAdresses, _ = commontestutils.CreateAccounts(2, 0)

var baseAccountAddress = accAdresses[0]
var shareDevelopmentFundAccountAddress = accAdresses[1]

var baseAccountAddressString = baseAccountAddress.String()
var shareDevelopmentFundAccountAddressString = shareDevelopmentFundAccountAddress.String()

func prepareBurningDistributor(destinationType DestinationType) types.SubDistributor {
	var address string
	if destinationType == BaseAccount {
		address = baseAccountAddressString
	} else {
		address = c4eDistributorCollectorName
	}

	var destAccount = types.Account{}
	destAccount.Id = address

	if destinationType == ModuleAccount {
		destAccount.Type = types.MODULE_ACCOUNT
	} else if destinationType == InternalAccount {
		destAccount.Type = types.INTERNAL_ACCOUNT
	} else {
		destAccount.Type = types.BASE_ACCOUNT
	}

	if destinationType == MainCollector {
		destAccount.Type = types.MAIN
	}

	burnShare := types.BurnShare{
		Percent: sdk.MustNewDecFromStr("51"),
	}

	destination := types.Destination{
		Account:   destAccount,
		Share:     nil,
		BurnShare: &burnShare,
	}

	distributor1 := types.SubDistributor{
		Name:        "tx_fee_distributor",
		Sources:     []*types.Account{{Id: authtypes.FeeCollectorName, Type: types.MODULE_ACCOUNT}},
		Destination: destination,
	}

	return distributor1
}

func prepareInflationToPassAcoutSubDistr(passThroughAccoutType DestinationType) types.SubDistributor {
	source := types.Account{
		Id:   "c4e",
		Type: types.MAIN,
	}

	var address string
	if passThroughAccoutType == BaseAccount {
		address = baseAccountAddressString
	} else {
		address = c4eDistributorCollectorName
	}

	var destAccount = types.Account{
		Id: address,
	}

	if passThroughAccoutType == ModuleAccount {
		destAccount.Type = types.MODULE_ACCOUNT
	} else if passThroughAccoutType == InternalAccount {
		destAccount.Type = types.INTERNAL_ACCOUNT
	} else {
		destAccount.Type = types.BASE_ACCOUNT
	}

	if passThroughAccoutType == MainCollector {
		destAccount.Type = types.MAIN
	}

	burnShare := types.BurnShare{
		Percent: sdk.MustNewDecFromStr("0"),
	}

	destination := types.Destination{
		Account:   destAccount,
		Share:     nil,
		BurnShare: &burnShare,
	}
	return types.SubDistributor{
		Name:        "pass_distributor",
		Sources:     []*types.Account{&source},
		Destination: destination,
	}
}

func prepareInflationSubDistributor(sourceAccoutType DestinationType, toValidators bool) types.SubDistributor {

	var address string
	if sourceAccoutType == BaseAccount {
		address = baseAccountAddressString
	} else {
		address = c4eDistributorCollectorName
	}

	var source = types.Account{
		Id: address,
	}

	if sourceAccoutType == ModuleAccount {
		source.Type = types.MODULE_ACCOUNT
	} else if sourceAccoutType == InternalAccount {
		source.Type = types.INTERNAL_ACCOUNT
	} else {
		source.Type = types.BASE_ACCOUNT
	}

	if sourceAccoutType == MainCollector {
		source.Type = types.MAIN
	}

	burnShare := types.BurnShare{
		Percent: sdk.MustNewDecFromStr("0"),
	}

	var destName string
	if toValidators {
		destName = types.ValidatorsRewardsCollector
	} else {
		destName = noValidatorsCollectorName
	}

	var destAccount = types.Account{
		Id:   destName,
		Type: types.MODULE_ACCOUNT,
	}

	var shareDevelopmentFundAccount = types.Account{
		Id:   shareDevelopmentFundAccountAddressString,
		Type: types.BASE_ACCOUNT,
	}

	shareDevelopmentFund := types.Share{
		Name:    "development_fund",
		Percent: sdk.MustNewDecFromStr("10.345"),
		Account: shareDevelopmentFundAccount,
	}

	destination := types.Destination{
		Account:   destAccount,
		Share:     []*types.Share{&shareDevelopmentFund},
		BurnShare: &burnShare,
	}

	return types.SubDistributor{
		Name:        "tx_fee_distributor",
		Sources:     []*types.Account{&source},
		Destination: destination,
	}
}

func TestBurningDistributorMainCollectorDes(t *testing.T) {
	BurningDistributorTest(t, MainCollector)
}

func TestBurningDistributorModuleAccountDest(t *testing.T) {
	BurningDistributorTest(t, ModuleAccount)
}

func TestBurningDistributorInternalAccountDest(t *testing.T) {
	BurningDistributorTest(t, InternalAccount)
}

func TestBurningDistributorBaseAccountDest(t *testing.T) {
	BurningDistributorTest(t, BaseAccount)
}

func BurningDistributorTest(t *testing.T, destinationType DestinationType) {
	senderCoin := sdk.NewInt(1017)

	testHelper := testapp.SetupTestApp(t)

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(senderCoin, authtypes.FeeCollectorName)
	testHelper.BankUtils.VerifyDefultDenomTotalSupply(testHelper.InitialValidatorsCoin.AddAmount(senderCoin).Amount)

	var subdistributors []types.SubDistributor
	subdistributors = append(subdistributors, prepareBurningDistributor(destinationType))

	testHelper.App.CfedistributorKeeper.SetParams(testHelper.Context, types.NewParams(subdistributors))
	testHelper.SetContextBlockHeight(int64(2))
	testHelper.BeginBlocker(abci.RequestBeginBlock{})

	//coin on "burnState" should be equal 498, remains: 1 and 0.33 on remains
	burnState, _ := testHelper.App.CfedistributorKeeper.GetState(testHelper.Context, "burn_state_key")

	coinRemains := burnState.CoinsStates
	require.EqualValues(t, sdk.MustNewDecFromStr("0.67"), coinRemains.AmountOf(commontestutils.DefaultTestDenom))

	if destinationType == MainCollector {
		testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(types.DistributorMainAccount, sdk.NewInt(499))
		require.EqualValues(t, 1, len(testHelper.App.CfedistributorKeeper.GetAllStates(testHelper.Context)))
	} else if destinationType == ModuleAccount {
		testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(types.DistributorMainAccount, sdk.NewInt(1))
		testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(c4eDistributorCollectorName, sdk.NewInt(498))

		require.EqualValues(t, 2, len(testHelper.App.CfedistributorKeeper.GetAllStates(testHelper.Context)))
		c4eDistrState, _ := testHelper.App.CfedistributorKeeper.GetState(testHelper.Context, c4eDistributorCollectorName)
		coinRemains := c4eDistrState.CoinsStates
		require.EqualValues(t, sdk.MustNewDecFromStr("0.33"), coinRemains.AmountOf(commontestutils.DefaultTestDenom))

	} else if destinationType == InternalAccount {
		testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(types.DistributorMainAccount, sdk.NewInt(499))

		require.EqualValues(t, 2, len(testHelper.App.CfedistributorKeeper.GetAllStates(testHelper.Context)))
		c4eDistrState, _ := testHelper.App.CfedistributorKeeper.GetState(testHelper.Context, c4eDistributorCollectorName)
		coinRemains := c4eDistrState.CoinsStates
		require.EqualValues(t, sdk.MustNewDecFromStr("498.33"), coinRemains.AmountOf(commontestutils.DefaultTestDenom))
	} else {
		testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(types.DistributorMainAccount, sdk.NewInt(1))
		testHelper.BankUtils.VerifyAccountDefultDenomBalance(baseAccountAddress, sdk.NewInt(498))

		require.EqualValues(t, 2, len(testHelper.App.CfedistributorKeeper.GetAllStates(testHelper.Context)))

		c4eDistrState, _ := testHelper.App.CfedistributorKeeper.GetState(testHelper.Context, baseAccountAddressString)
		coinRemains := c4eDistrState.CoinsStates
		require.EqualValues(t, sdk.MustNewDecFromStr("0.33"), coinRemains.AmountOf(commontestutils.DefaultTestDenom))
	}

	testHelper.BankUtils.VerifyDefultDenomTotalSupply(testHelper.InitialValidatorsCoin.AddAmount(sdk.NewInt(499)).Amount)
}

func TestBurningWithInflationDistributorPassThroughMainCollector(t *testing.T) {
	BurningWithInflationDistributorTest(t, MainCollector, true)
}

func TestBurningWithInflationDistributorPassThroughModuleAccount(t *testing.T) {
	BurningWithInflationDistributorTest(t, ModuleAccount, true)
}

func TestBurningWithInflationDistributorPassInternalAccountAccount(t *testing.T) {
	BurningWithInflationDistributorTest(t, InternalAccount, true)
}

func TestBurningWithInflationDistributorPassBaseAccountAccount(t *testing.T) {
	BurningWithInflationDistributorTest(t, BaseAccount, true)
}

func TestBurningWithInflationDistributorPassThroughMainCollectorNoValidators(t *testing.T) {
	BurningWithInflationDistributorTest(t, MainCollector, false)
}

func TestBurningWithInflationDistributorPassThroughModuleAccountNoValidators(t *testing.T) {
	BurningWithInflationDistributorTest(t, ModuleAccount, false)
}

func TestBurningWithInflationDistributorPassInternalAccountAccountNoValidators(t *testing.T) {
	BurningWithInflationDistributorTest(t, InternalAccount, false)
}

func TestBurningWithInflationDistributorPassBaseAccountAccountNoValidators(t *testing.T) {
	BurningWithInflationDistributorTest(t, BaseAccount, false)
}

func BurningWithInflationDistributorTest(t *testing.T, passThroughAccoutType DestinationType, toValidators bool) {

	testHelper := testapp.SetupTestApp(t)

	//prepare module account with coin to distribute fee_collector 1017
	cointToMint := sdk.NewInt(1017)

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(cointToMint, authtypes.FeeCollectorName)

	cointToMintFromInflation := sdk.NewInt(5044)
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(cointToMintFromInflation, types.DistributorMainAccount)

	initialCoinAmount := testHelper.InitialValidatorsCoin.AddAmount(cointToMint).AddAmount(cointToMintFromInflation)

	testHelper.BankUtils.VerifyDefultDenomTotalSupply(initialCoinAmount.Amount)

	var subDistributors []types.SubDistributor

	subDistributors = append(subDistributors, prepareBurningDistributor(MainCollector))
	if passThroughAccoutType != MainCollector {
		subDistributors = append(subDistributors, prepareInflationToPassAcoutSubDistr(passThroughAccoutType))
	}
	subDistributors = append(subDistributors, prepareInflationSubDistributor(passThroughAccoutType, toValidators))

	testHelper.App.CfedistributorKeeper.SetParams(testHelper.Context, types.NewParams(subDistributors))
	testHelper.SetContextBlockHeight(int64(2))
	testHelper.BeginBlocker(abci.RequestBeginBlock{})

	if passThroughAccoutType == MainCollector {
		require.EqualValues(t, 3, len(testHelper.App.CfedistributorKeeper.GetAllStates(testHelper.Context)))
	} else if passThroughAccoutType == ModuleAccount {
		require.EqualValues(t, 4, len(testHelper.App.CfedistributorKeeper.GetAllStates(testHelper.Context)))
	} else if passThroughAccoutType == InternalAccount {
		require.EqualValues(t, 4, len(testHelper.App.CfedistributorKeeper.GetAllStates(testHelper.Context)))
	} else {
		require.EqualValues(t, 4, len(testHelper.App.CfedistributorKeeper.GetAllStates(testHelper.Context)))
	}

	// coins flow:
	// fee 1017*51% = 518.67 to burn, so 518 burned - and burn remains 0.67
	testHelper.BankUtils.VerifyDefultDenomTotalSupply(initialCoinAmount.SubAmount(sdk.NewInt(518)).Amount)

	burnState, _ := testHelper.App.CfedistributorKeeper.GetBurnState(testHelper.Context)
	coinRemains := burnState.CoinsStates
	require.EqualValues(t, sdk.MustNewDecFromStr("0.67"), coinRemains.AmountOf(commontestutils.DefaultTestDenom))

	// added 499 to main collector
	// main collector state = 499 + 5044 = 5543, but 5543 - 0,67 = 5542.33 to distribute

	if passThroughAccoutType == ModuleAccount || passThroughAccoutType == InternalAccount {
		// 5542.33 moved to c4e_distributor module or internal account
		// and all is distributed further, and 0 in remains

		testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(c4eDistributorCollectorName, sdk.NewInt(0))

		remains, _ := testHelper.App.CfedistributorKeeper.GetState(testHelper.Context, c4eDistributorCollectorName)

		coinRemainsDevelopmentFund := remains.CoinsStates
		require.EqualValues(t, sdk.MustNewDecFromStr("0"), coinRemainsDevelopmentFund.AmountOf(commontestutils.DefaultTestDenom))
	} else if passThroughAccoutType == BaseAccount {
		// 5542.33 moved to cosmos13zg4u07ymq83uq73t2cq3dj54jj37zzgr3hlck account
		// and all is distributed further, and 0 in remains
		testHelper.BankUtils.VerifyAccountDefultDenomBalance(baseAccountAddress, sdk.NewInt(0))

		remains, _ := testHelper.App.CfedistributorKeeper.GetState(testHelper.Context, baseAccountAddressString)

		coinRemainsDevelopmentFund := remains.CoinsStates
		require.EqualValues(t, sdk.MustNewDecFromStr("0"), coinRemainsDevelopmentFund.AmountOf(commontestutils.DefaultTestDenom))
	}

	// 5542.33*10.345% = 573.3540385 to cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag, so
	// 573 on cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag and 0.3540385 on its distributor state

	testHelper.BankUtils.VerifyAccountDefultDenomBalance(shareDevelopmentFundAccountAddress, sdk.NewInt(573))

	remains, _ := testHelper.App.CfedistributorKeeper.GetState(testHelper.Context, shareDevelopmentFundAccountAddressString)
	coinRemainsDevelopmentFund := remains.CoinsStates
	require.EqualValues(t, sdk.MustNewDecFromStr("0.3540385"), coinRemainsDevelopmentFund.AmountOf(commontestutils.DefaultTestDenom))

	// 5542.33 - 573.3540385 = 4968.9759615 to validators_rewards_collector, so
	// 4968 on validators_rewards_collector or no_validators module account and 0.9759615 on its distributor state

	if toValidators {
		// validators_rewards_collector coins sent to vaalidator distribition so amount is 0,

		testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(types.ValidatorsRewardsCollector, sdk.NewInt(0))

		// still 0.9759615 on its distributor state remains
		remains, _ = testHelper.App.CfedistributorKeeper.GetState(testHelper.Context, types.ValidatorsRewardsCollector)
		coinRemainsValidatorsReward := remains.CoinsStates
		require.EqualValues(t, sdk.MustNewDecFromStr("0.9759615"), coinRemainsValidatorsReward.AmountOf(commontestutils.DefaultTestDenom))
		// and 4968 to validators rewards

		testHelper.BankUtils.VerifyAccountDefultDenomBalance(testHelper.App.DistrKeeper.GetDistributionAccount(testHelper.Context).GetAddress(), sdk.NewInt(4968))
	} else {
		// no_validators module account coins amount is 4968,
		// and remains 0.9759615 on its distributor state

		testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(noValidatorsCollectorName, sdk.NewInt(4968))

		remains, _ = testHelper.App.CfedistributorKeeper.GetState(testHelper.Context, noValidatorsCollectorName)
		coinRemainsValidatorsReward := remains.CoinsStates
		require.EqualValues(t, sdk.MustNewDecFromStr("0.9759615"), coinRemainsValidatorsReward.AmountOf(commontestutils.DefaultTestDenom))
	}

	// 5543 - 573 - 4968 = 2 (its ramains 0,67 + 0.3540385 + 0.9759615 = 2) on main collector
	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(types.DistributorMainAccount, sdk.NewInt(2))

}

func TestBurningWithInflationDistributorAfter3001Blocks(t *testing.T) {

	testHelper := testapp.SetupTestApp(t)

	var subdistributors []types.SubDistributor
	subdistributors = append(subdistributors, prepareBurningDistributor(MainCollector))
	subdistributors = append(subdistributors, prepareInflationSubDistributor(MainCollector, true))
	testHelper.App.CfedistributorKeeper.SetParams(testHelper.Context, types.NewParams(subdistributors))

	for i := int64(1); i <= 3001; i++ {

		cointToMint := sdk.NewInt(1017)

		testHelper.BankUtils.AddDefaultDenomCoinsToModule(cointToMint, authtypes.FeeCollectorName)

		cointToMintFromInflation := sdk.NewInt(5044)

		testHelper.BankUtils.AddDefaultDenomCoinsToModule(cointToMintFromInflation, types.DistributorMainAccount)

		testHelper.SetContextBlockHeight(int64(i))
		testHelper.BeginBlocker(abci.RequestBeginBlock{})
		testHelper.EndBlocker(abci.RequestEndBlock{})
		burn, _ := sdk.NewDecFromStr("518.67")
		burn = burn.MulInt64(i)
		burn.GT(burn.TruncateDec())
		totalExpected := sdk.NewDec(i * (1017 + 5044)).Sub(burn)

		totalExpectedTruncated := totalExpected.TruncateInt()

		if burn.GT(burn.TruncateDec()) {
			totalExpectedTruncated = totalExpectedTruncated.AddRaw(1)
		}
		require.EqualValues(t, testHelper.InitialValidatorsCoin.AddAmount(totalExpectedTruncated), testHelper.App.BankKeeper.GetSupply(testHelper.Context, commontestutils.DefaultTestDenom))
	}
	testHelper.SetContextBlockHeight(int64(3002))
	testHelper.BeginBlocker(abci.RequestBeginBlock{})
	testHelper.EndBlocker(abci.RequestEndBlock{})

	// coins flow:
	// fee 3001*1017*51% = 1556528.67 to burn, so 1556528 burned - and burn remains 0.67
	// fee 3001*(1017*51%) = 3001*518.67 = 1556528.67 to burn, so 1556528 burned - and burn remains 0.67
	require.EqualValues(t, testHelper.InitialValidatorsCoin.AddAmount(sdk.NewInt(3001*(1017+5044)-1556528)), testHelper.App.BankKeeper.GetSupply(testHelper.Context, commontestutils.DefaultTestDenom))

	burnState, _ := testHelper.App.CfedistributorKeeper.GetBurnState(testHelper.Context)
	coinRemains := burnState.CoinsStates
	require.EqualValues(t, sdk.MustNewDecFromStr("0.67"), coinRemains.AmountOf("uc4e"))

	// added 3001*1017 - 1556528 = 1495489 to main collector
	// main collector state = 1495489 + 3001*5044 = 16632533, but 16632533 - 0.67 (burning remains) = 16632532.33 to distribute

	// 16632532.33*10.345% = 1720635.4695385 to cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag, so
	// 1720635 on cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag and 0.4695385 on its distributor state

	testHelper.BankUtils.VerifyAccountDefultDenomBalance(shareDevelopmentFundAccountAddress, sdk.NewInt(1720635))

	remains, _ := testHelper.App.CfedistributorKeeper.GetState(testHelper.Context, shareDevelopmentFundAccountAddressString)
	coinRemainsDevelopmentFund := remains.CoinsStates
	require.EqualValues(t, sdk.MustNewDecFromStr("0.4695385"), coinRemainsDevelopmentFund.AmountOf(commontestutils.DefaultTestDenom))

	// 16632532.33- 1720635.4695385 = 14911896.8604615 to validators_rewards_collector, so
	// 14911896 on validators_rewards_collector or no_validators module account and 0.8604615 on its distributor state

	// validators_rewards_collector coins sent to vaalidator distribition so amount is 0,

	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(types.ValidatorsRewardsCollector, sdk.ZeroInt())

	// still 0.8845 on its distributor state remains
	remains, _ = testHelper.App.CfedistributorKeeper.GetState(testHelper.Context, types.ValidatorsRewardsCollector)
	coinRemainsValidatorsReward := remains.CoinsStates
	require.EqualValues(t, sdk.MustNewDecFromStr("0.8604615"), coinRemainsValidatorsReward.AmountOf(commontestutils.DefaultTestDenom))
	// and 14906927 to validators rewards

	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(distrtypes.ModuleName, sdk.NewInt(14911896))

	// 16632533 - 1720635 - 14911896 = 1 (its ramains 0.67 + 0.4695385 + 0.8604615 = 2) on main collector
	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(types.DistributorMainAccount, sdk.NewInt(2))

}
