package cfedistributor

import (
	sdkrand "github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math/rand"
)

type DestinationType int64

const (
	MainCollector DestinationType = iota
	ModuleAccount
	InternalAccount
	BaseAccount
)

func RandomCollectorName(r *rand.Rand) DestinationType {
	randVal := sdkrand.RandIntBetween(r, 0, 3)
	switch randVal {
	case 0:
		return MainCollector
	case 1:
		return ModuleAccount
	case 2:
		return InternalAccount
	case 3:
		return BaseAccount
	}
	return MainCollector
}

func PrepareBurningDistributor(destinationType DestinationType) types.SubDistributor {
	var address string
	if destinationType == BaseAccount {
		address = "cosmos13zg4u07ymq83uq73t2cq3dj54jj37zzgr3hlck"
	} else {
		address = "c4e_distributor"
	}

	var destAccount = types.Account{}
	destAccount.Id = address

	if destinationType == ModuleAccount {
		destAccount.Type = "MODULE_ACCOUNT"
	} else if destinationType == InternalAccount {
		destAccount.Type = "INTERNAL_ACCOUNT"
	} else {
		destAccount.Type = "BASE_ACCOUNT"
	}

	if destinationType == MainCollector {
		destAccount.Type = "MAIN"
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
		Sources:     []*types.Account{{Id: "fee_collector", Type: "MODULE_ACCOUNT"}},
		Destination: destination,
	}

	return distributor1
}

func PrepareInflationToPassAcoutSubDistr(passThroughAccoutType DestinationType) types.SubDistributor {
	source := types.Account{
		Id:   "c4e",
		Type: "MAIN",
	}

	var address string
	if passThroughAccoutType == BaseAccount {
		address = "cosmos13zg4u07ymq83uq73t2cq3dj54jj37zzgr3hlck"
	} else {
		address = "c4e_distributor"
	}

	var destAccount = types.Account{
		Id: address,
	}

	if passThroughAccoutType == ModuleAccount {
		destAccount.Type = "MODULE_ACCOUNT"
	} else if passThroughAccoutType == InternalAccount {
		destAccount.Type = "INTERNAL_ACCOUNT"
	} else {
		destAccount.Type = "BASE_ACCOUNT"
	}

	if passThroughAccoutType == MainCollector {
		destAccount.Type = "MAIN"
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
		address = "cosmos13zg4u07ymq83uq73t2cq3dj54jj37zzgr3hlck"
	} else {
		address = "c4e_distributor"
	}

	var source = types.Account{
		Id: address,
	}

	if sourceAccoutType == ModuleAccount {
		source.Type = "MODULE_ACCOUNT"
	} else if sourceAccoutType == InternalAccount {
		source.Type = "INTERNAL_ACCOUNT"
	} else {
		source.Type = "BASE_ACCOUNT"
	}

	if sourceAccoutType == MainCollector {
		source.Type = "MAIN"
	}

	// source := types.Account{IsMainCollector: true, IsModuleAccount: false, IsInternalAccount: false}

	burnShare := types.BurnShare{
		Percent: sdk.MustNewDecFromStr("0"),
	}

	var destName string
	if toValidators {
		destName = types.ValidatorsRewardsCollector
	} else {
		destName = "no_validators"
	}

	var destAccount = types.Account{
		Id:   destName,
		Type: "MODULE_ACCOUNT",
	}

	var shareDevelopmentFundAccount = types.Account{
		Id:   "cosmos1p20lmfzp4g9vywl2jxwexwh6akvkxzpa6hdrag",
		Type: "BASE_ACCOUNT",
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
