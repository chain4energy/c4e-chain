package subdistributor

import (
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type DestinationType int64

const (
	MainCollector DestinationType = iota
	ModuleAccount
	InternalAccount
	BaseAccount
)

const C4eDistributorCollectorName = types.GreenEnergyBoosterCollector
const NoValidatorsCollectorName = types.GovernanceBoosterCollector

var accAdresses, _ = commontestutils.CreateAccounts(3, 0)

var BaseAccountAddress = accAdresses[0]
var ShareDevelopmentFundAccountAddress = accAdresses[1]
var HelperDestinationAccountAddress = accAdresses[2]

var BaseAccountAddressString = BaseAccountAddress.String()
var ShareDevelopmentFundAccountAddressString = ShareDevelopmentFundAccountAddress.String()
var HelperDestinationAccountAddressString = HelperDestinationAccountAddress.String()

func PreparareMainDefaultDistributor() types.SubDistributor {
	helperDestination := types.Destinations{
		PrimaryShare: types.Account{Id: HelperDestinationAccountAddressString, Type: types.BASE_ACCOUNT},
		Shares:       nil,
		BurnShare:    sdk.ZeroDec(),
	}
	distributor1 := types.SubDistributor{
		Name:         "default_main_distributor",
		Sources:      []*types.Account{{Id: "", Type: types.MAIN}},
		Destinations: helperDestination,
	}

	return distributor1
}

func PreparareHelperDistributorForDestination(destination types.Account) types.SubDistributor {
	helperDestination := types.Destinations{
		PrimaryShare: types.Account{Id: HelperDestinationAccountAddressString, Type: types.BASE_ACCOUNT},
		Shares:       nil,
		BurnShare:    sdk.ZeroDec(),
	}
	distributor1 := types.SubDistributor{
		Name:         "test_helper_distributor",
		Sources:      []*types.Account{&destination},
		Destinations: helperDestination,
	}

	return distributor1
}

func PrepareBurningDistributor(destinationType DestinationType) types.SubDistributor {
	var address string
	if destinationType == BaseAccount {
		address = BaseAccountAddressString
	} else {
		address = C4eDistributorCollectorName
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

	burnShare := sdk.MustNewDecFromStr("0.51")

	destination := types.Destinations{
		PrimaryShare: destAccount,
		Shares:       nil,
		BurnShare:    burnShare,
	}

	distributor1 := types.SubDistributor{
		Name:         helpers.RandStringOfLength(10),
		Sources:      []*types.Account{{Id: authtypes.FeeCollectorName, Type: types.MODULE_ACCOUNT}},
		Destinations: destination,
	}

	return distributor1
}

func PrepareInflationToPassAcoutSubDistr(passThroughAccoutType DestinationType) types.SubDistributor {
	source := types.Account{
		Id:   "c4e",
		Type: types.MAIN,
	}

	var address string
	if passThroughAccoutType == BaseAccount {
		address = BaseAccountAddressString
	} else {
		address = C4eDistributorCollectorName
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

	burnShare := sdk.ZeroDec()

	destination := types.Destinations{
		PrimaryShare: destAccount,
		Shares:       nil,
		BurnShare:    burnShare,
	}
	return types.SubDistributor{
		Name:         helpers.RandStringOfLength(10),
		Sources:      []*types.Account{&source},
		Destinations: destination,
	}
}

func PrepareInflationSubDistributor(sourceAccoutType DestinationType, toValidators bool) types.SubDistributor {

	var address string
	if sourceAccoutType == BaseAccount {
		address = BaseAccountAddressString
	} else {
		address = C4eDistributorCollectorName
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

	burnShare := sdk.ZeroDec()

	var destName string
	if toValidators {
		destName = types.ValidatorsRewardsCollector
	} else {
		destName = NoValidatorsCollectorName
	}

	var destAccount = types.Account{
		Id:   destName,
		Type: types.MODULE_ACCOUNT,
	}

	var shareDevelopmentFundAccount = types.Account{
		Id:   ShareDevelopmentFundAccountAddressString,
		Type: types.BASE_ACCOUNT,
	}

	shareDevelopmentFund := types.DestinationShare{
		Name:        helpers.RandStringOfLength(10),
		Share:       sdk.MustNewDecFromStr("0.10345"),
		Destination: shareDevelopmentFundAccount,
	}

	destination := types.Destinations{
		PrimaryShare: destAccount,
		Shares:       []*types.DestinationShare{&shareDevelopmentFund},
		BurnShare:    burnShare,
	}

	return types.SubDistributor{
		Name:         helpers.RandStringOfLength(10),
		Sources:      []*types.Account{&source},
		Destinations: destination,
	}
}
