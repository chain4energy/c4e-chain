package subdistributor

import (
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
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

var accAdresses, _ = commontestutils.CreateAccounts(2, 0)

var BaseAccountAddress = accAdresses[0]
var ShareDevelopmentFundAccountAddress = accAdresses[1]

var BaseAccountAddressString = BaseAccountAddress.String()
var ShareDevelopmentFundAccountAddressString = ShareDevelopmentFundAccountAddress.String()

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

	burnShare := types.BurnShare{
		Percent: sdk.MustNewDecFromStr("0"),
	}

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