package cfedistributor_test

import (
	"cosmossdk.io/math"
	"github.com/chain4energy/c4e-chain/v2/testutil/app"
	"testing"

	subdistributortestutils "github.com/chain4energy/c4e-chain/v2/testutil/module/cfedistributor/subdistributor"

	"github.com/chain4energy/c4e-chain/v2/x/cfedistributor/types"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
)

func TestBurningDistributorMainCollectorDes(t *testing.T) {
	BurningDistributorTest(t, subdistributortestutils.MainCollector)
}

func TestBurningDistributorModuleAccountDest(t *testing.T) {
	BurningDistributorTest(t, subdistributortestutils.ModuleAccount)
}

func TestBurningDistributorInternalAccountDest(t *testing.T) {
	BurningDistributorTest(t, subdistributortestutils.InternalAccount)
}

func TestBurningDistributorBaseAccountDest(t *testing.T) {
	BurningDistributorTest(t, subdistributortestutils.BaseAccount)
}

func BurningDistributorTest(t *testing.T, destinationType subdistributortestutils.DestinationType) {
	senderCoin := math.NewInt(1017)

	testHelper := app.SetupTestApp(t)

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(senderCoin, authtypes.FeeCollectorName)
	testHelper.BankUtils.VerifyDefultDenomTotalSupply(testHelper.InitialValidatorsCoin.AddAmount(senderCoin).Amount)

	var subdistributors []types.SubDistributor
	if destinationType != subdistributortestutils.MainCollector {
		subdistributors = append(subdistributors, subdistributortestutils.PreparareMainDefaultDistributor())
	}

	burningSubSistributor := subdistributortestutils.PrepareBurningDistributor(destinationType)
	subdistributors = append(subdistributors, burningSubSistributor)
	if destinationType == subdistributortestutils.MainCollector || destinationType == subdistributortestutils.InternalAccount {
		subdistributors = append(subdistributors, subdistributortestutils.PreparareHelperDistributorForDestination(burningSubSistributor.Destinations.PrimaryShare))
	}

	testHelper.C4eDistributorUtils.SetSubDistributorsParams(subdistributors)
	testHelper.SetContextBlockHeight(int64(2))
	testHelper.BeginBlocker(abci.RequestBeginBlock{})

	//coin on "burnState" should be equal 498, remains: 1 and 0.33 on remains
	testHelper.C4eDistributorUtils.VerifyDefaultDenomBurnStateAmount(sdk.MustNewDecFromStr("0.67"))

	if destinationType == subdistributortestutils.MainCollector {
		testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(types.DistributorMainAccount, math.NewInt(1))
		testHelper.BankUtils.VerifyAccountDefultDenomBalance(subdistributortestutils.HelperDestinationAccountAddress, math.NewInt(498))
		testHelper.C4eDistributorUtils.VerifyNumberOfStates(2)
		testHelper.C4eDistributorUtils.VerifyDefaultDenomStateAmount(subdistributors[1].Destinations.PrimaryShare, sdk.MustNewDecFromStr("0.33"))
	} else if destinationType == subdistributortestutils.ModuleAccount {
		testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(types.DistributorMainAccount, math.NewInt(1))
		testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(subdistributortestutils.C4eDistributorCollectorName, math.NewInt(498))
		testHelper.C4eDistributorUtils.VerifyNumberOfStates(2)
		testHelper.C4eDistributorUtils.VerifyDefaultDenomStateAmount(subdistributors[1].Destinations.PrimaryShare, sdk.MustNewDecFromStr("0.33"))
	} else if destinationType == subdistributortestutils.InternalAccount {
		testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(types.DistributorMainAccount, math.NewInt(1))
		testHelper.BankUtils.VerifyAccountDefultDenomBalance(subdistributortestutils.HelperDestinationAccountAddress, math.NewInt(498))
		testHelper.C4eDistributorUtils.VerifyNumberOfStates(3)
		testHelper.C4eDistributorUtils.VerifyDefaultDenomStateAmount(subdistributors[2].Destinations.PrimaryShare, sdk.MustNewDecFromStr("0.33"))
		testHelper.C4eDistributorUtils.VerifyDefaultDenomStateAmount(subdistributors[1].Destinations.PrimaryShare, sdk.ZeroDec())
	} else {
		testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(types.DistributorMainAccount, math.NewInt(1))
		testHelper.BankUtils.VerifyAccountDefultDenomBalance(subdistributortestutils.BaseAccountAddress, math.NewInt(498))
		testHelper.C4eDistributorUtils.VerifyNumberOfStates(2)
		testHelper.C4eDistributorUtils.VerifyDefaultDenomStateAmount(subdistributors[1].Destinations.PrimaryShare, sdk.MustNewDecFromStr("0.33"))
	}

	testHelper.BankUtils.VerifyDefultDenomTotalSupply(testHelper.InitialValidatorsCoin.AddAmount(math.NewInt(499)).Amount)
	testHelper.C4eDistributorUtils.ValidateGenesisAndInvariants()
}

func TestBurningWithInflationDistributorPassThroughMainCollector(t *testing.T) {
	BurningWithInflationDistributorTest(t, subdistributortestutils.MainCollector, true)
}

func TestBurningWithInflationDistributorPassThroughModuleAccount(t *testing.T) {
	BurningWithInflationDistributorTest(t, subdistributortestutils.ModuleAccount, true)
}

func TestBurningWithInflationDistributorPassInternalAccountAccount(t *testing.T) {
	BurningWithInflationDistributorTest(t, subdistributortestutils.InternalAccount, true)
}

func TestBurningWithInflationDistributorPassBaseAccountAccount(t *testing.T) {
	BurningWithInflationDistributorTest(t, subdistributortestutils.BaseAccount, true)
}

func TestBurningWithInflationDistributorPassThroughMainCollectorNoValidators(t *testing.T) {
	BurningWithInflationDistributorTest(t, subdistributortestutils.MainCollector, false)
}

func TestBurningWithInflationDistributorPassThroughModuleAccountNoValidators(t *testing.T) {
	BurningWithInflationDistributorTest(t, subdistributortestutils.ModuleAccount, false)
}

func TestBurningWithInflationDistributorPassInternalAccountAccountNoValidators(t *testing.T) {
	BurningWithInflationDistributorTest(t, subdistributortestutils.InternalAccount, false)
}

func TestBurningWithInflationDistributorPassBaseAccountAccountNoValidators(t *testing.T) {
	BurningWithInflationDistributorTest(t, subdistributortestutils.BaseAccount, false)
}

func BurningWithInflationDistributorTest(t *testing.T, passThroughAccoutType subdistributortestutils.DestinationType, toValidators bool) {
	testHelper := app.SetupTestApp(t)

	//prepare module account with coin to distribute fee_collector 1017
	cointToMint := math.NewInt(1017)

	testHelper.BankUtils.AddDefaultDenomCoinsToModule(cointToMint, authtypes.FeeCollectorName)

	cointToMintFromInflation := math.NewInt(5044)
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(cointToMintFromInflation, types.DistributorMainAccount)

	initialCoinAmount := testHelper.InitialValidatorsCoin.AddAmount(cointToMint).AddAmount(cointToMintFromInflation)

	testHelper.BankUtils.VerifyDefultDenomTotalSupply(initialCoinAmount.Amount)

	var subDistributors []types.SubDistributor
	subDistributors = append(subDistributors, subdistributortestutils.PrepareBurningDistributor(subdistributortestutils.MainCollector))
	if passThroughAccoutType != subdistributortestutils.MainCollector {
		subDistributors = append(subDistributors, subdistributortestutils.PrepareInflationToPassAcoutSubDistr(passThroughAccoutType))
	}
	subDistributors = append(subDistributors, subdistributortestutils.PrepareInflationSubDistributor(passThroughAccoutType, toValidators))

	lastDistributorIndex := len(subDistributors) - 1
	testHelper.C4eDistributorUtils.SetSubDistributorsParams(subDistributors)
	testHelper.SetContextBlockHeight(int64(2))
	testHelper.BeginBlocker(abci.RequestBeginBlock{})

	if passThroughAccoutType == subdistributortestutils.MainCollector {
		testHelper.C4eDistributorUtils.VerifyNumberOfStates(3)
	} else {
		testHelper.C4eDistributorUtils.VerifyNumberOfStates(4)
	}

	// coins flow:
	// fee 1017*51% = 518.67 to burn, so 518 burned - and burn remains 0.67

	testHelper.BankUtils.VerifyDefultDenomTotalSupply(initialCoinAmount.SubAmount(math.NewInt(518)).Amount)
	testHelper.C4eDistributorUtils.VerifyDefaultDenomBurnStateAmount(sdk.MustNewDecFromStr("0.67"))

	// added 499 to main collector
	// main collector state = 499 + 5044 = 5543, but 5543 - 0,67 = 5542.33 to distribute

	if passThroughAccoutType == subdistributortestutils.ModuleAccount || passThroughAccoutType == subdistributortestutils.InternalAccount {
		// 5542.33 moved to c4e_distributor module or internal account
		// and all is distributed further, and 0 in remains
		testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(subdistributortestutils.C4eDistributorCollectorName, math.NewInt(0))
		testHelper.C4eDistributorUtils.VerifyDefaultDenomStateAmount(subDistributors[0].Destinations.PrimaryShare, sdk.ZeroDec())
	} else if passThroughAccoutType == subdistributortestutils.BaseAccount {
		// 5542.33 moved to cosmos13zg4u07ymq83uq73t2cq3dj54jj37zzgr3hlck account
		// and all is distributed further, and 0 in remains
		testHelper.BankUtils.VerifyAccountDefultDenomBalance(subdistributortestutils.BaseAccountAddress, math.NewInt(0))
		testHelper.C4eDistributorUtils.VerifyDefaultDenomStateAmount(subDistributors[0].Destinations.PrimaryShare, sdk.ZeroDec())
	}

	// 5542.33*10.345% = 573.3540385 to cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag, so
	// 573 on cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag and 0.3540385 on its distributor state

	testHelper.BankUtils.VerifyAccountDefultDenomBalance(subdistributortestutils.ShareDevelopmentFundAccountAddress, math.NewInt(573))
	testHelper.C4eDistributorUtils.VerifyDefaultDenomStateAmount(subDistributors[lastDistributorIndex].Destinations.Shares[0].Destination, sdk.MustNewDecFromStr("0.3540385"))

	// 5542.33 - 573.3540385 = 4968.9759615 to validators_rewards_collector, so
	// 4968 on validators_rewards_collector or no_validators module account and 0.9759615 on its distributor state

	lastSubDistributorPrimaryShare := subDistributors[lastDistributorIndex].Destinations.PrimaryShare
	if toValidators {
		// validators_rewards_collector coins sent to vaalidator distribition so amount is 0,
		testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(types.ValidatorsRewardsCollector, math.NewInt(0))
		// still 0.9759615 on its distributor state remains
		testHelper.C4eDistributorUtils.VerifyDefaultDenomStateAmount(lastSubDistributorPrimaryShare, sdk.MustNewDecFromStr("0.9759615"))
		// and 4968 to validators rewards
		testHelper.BankUtils.VerifyAccountDefultDenomBalance(testHelper.App.DistrKeeper.GetDistributionAccount(testHelper.Context).GetAddress(), math.NewInt(4968))
	} else {
		// no_validators module account coins amount is 4968,
		// and remains 0.9759615 on its distributor state
		testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(subdistributortestutils.NoValidatorsCollectorName, math.NewInt(4968))
		testHelper.C4eDistributorUtils.VerifyDefaultDenomStateAmount(lastSubDistributorPrimaryShare, sdk.MustNewDecFromStr("0.9759615"))
	}

	// 5543 - 573 - 4968 = 2 (its ramains 0,67 + 0.3540385 + 0.9759615 = 2) on main collector
	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(types.DistributorMainAccount, math.NewInt(2))
	testHelper.C4eDistributorUtils.ValidateGenesisAndInvariants()
}

func TestBurningWithInflationDistributorAfter3001Blocks(t *testing.T) {
	testHelper := app.SetupTestApp(t)

	var subdistributors []types.SubDistributor
	subdistributors = append(subdistributors, subdistributortestutils.PrepareBurningDistributor(subdistributortestutils.MainCollector))
	subdistributors = append(subdistributors, subdistributortestutils.PrepareInflationSubDistributor(subdistributortestutils.MainCollector, true))
	testHelper.C4eDistributorUtils.SetSubDistributorsParams(subdistributors)

	for i := int64(1); i <= 3001; i++ {
		cointToMint := math.NewInt(1017)

		testHelper.BankUtils.AddDefaultDenomCoinsToModule(cointToMint, authtypes.FeeCollectorName)

		cointToMintFromInflation := math.NewInt(5044)

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
		testHelper.BankUtils.VerifyDefultDenomTotalSupply(testHelper.InitialValidatorsCoin.AddAmount(totalExpectedTruncated).Amount)
		testHelper.C4eDistributorUtils.ValidateGenesisAndInvariants()
	}
	testHelper.SetContextBlockHeight(int64(3002))
	testHelper.BeginBlocker(abci.RequestBeginBlock{})
	testHelper.EndBlocker(abci.RequestEndBlock{})

	// coins flow:
	// fee 3001*1017*51% = 1556528.67 to burn, so 1556528 burned - and burn remains 0.67
	// fee 3001*(1017*51%) = 3001*518.67 = 1556528.67 to burn, so 1556528 burned - and burn remains 0.67
	testHelper.BankUtils.VerifyDefultDenomTotalSupply(testHelper.InitialValidatorsCoin.AddAmount(math.NewInt(3001*(1017+5044) - 1556528)).Amount)

	testHelper.C4eDistributorUtils.VerifyDefaultDenomBurnStateAmount(sdk.MustNewDecFromStr("0.67"))

	// added 3001*1017 - 1556528 = 1495489 to main collector
	// main collector state = 1495489 + 3001*5044 = 16632533, but 16632533 - 0.67 (burning remains) = 16632532.33 to distribute

	// 16632532.33*10.345% = 1720635.4695385 to cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag, so
	// 1720635 on cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag and 0.4695385 on its distributor state

	testHelper.BankUtils.VerifyAccountDefultDenomBalance(subdistributortestutils.ShareDevelopmentFundAccountAddress, math.NewInt(1720635))
	testHelper.C4eDistributorUtils.VerifyDefaultDenomStateAmount(subdistributors[1].Destinations.Shares[0].Destination, sdk.MustNewDecFromStr("0.4695385"))

	// 16632532.33- 1720635.4695385 = 14911896.8604615 to validators_rewards_collector, so
	// 14911896 on validators_rewards_collector or no_validators module account and 0.8604615 on its distributor state

	// validators_rewards_collector coins sent to vaalidator distribition so amount is 0,

	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(types.ValidatorsRewardsCollector, math.ZeroInt())

	// still 0.8845 on its distributor state remains
	testHelper.C4eDistributorUtils.VerifyDefaultDenomStateAmount(subdistributors[1].Destinations.PrimaryShare, sdk.MustNewDecFromStr("0.8604615"))

	// and 14906927 to validators rewards
	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(distrtypes.ModuleName, math.NewInt(14911896))
	// 16632533 - 1720635 - 14911896 = 1 (its ramains 0.67 + 0.4695385 + 0.8604615 = 2) on main collector
	testHelper.BankUtils.VerifyModuleAccountDefultDenomBalance(types.DistributorMainAccount, math.NewInt(2))

	testHelper.C4eDistributorUtils.ValidateGenesisAndInvariants()
}
