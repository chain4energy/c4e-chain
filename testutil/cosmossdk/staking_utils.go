package cosmossdk

import (
	"cosmossdk.io/math"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type StakingUtils struct {
	t                   require.TestingT
	helperStakingkeeper stakingkeeper.Keeper
	bankUtils           *BankUtils
}

func NewStakingUtils(t require.TestingT, helperStakingkeeper stakingkeeper.Keeper, bankUtils *BankUtils) StakingUtils {
	return StakingUtils{t: t, helperStakingkeeper: helperStakingkeeper, bankUtils: bankUtils}
}

func (su *StakingUtils) CreateValidator(ctx sdk.Context, addr sdk.ValAddress, pk cryptotypes.PubKey, coin sdk.Coin, commisions stakingtypes.CommissionRates) {
	msg, err := stakingtypes.NewMsgCreateValidator(addr, pk, coin, stakingtypes.Description{}, commisions, sdk.OneInt())
	msgSrvr := stakingkeeper.NewMsgServerImpl(&su.helperStakingkeeper)
	require.NoError(su.t, err)
	res, err := msgSrvr.CreateValidator(sdk.WrapSDKContext(ctx), msg)
	require.NoError(su.t, err)
	require.NotNil(su.t, res)
}

func (su *StakingUtils) SetupValidators(ctx sdk.Context, validators []sdk.ValAddress, delegatePerValidator math.Int) {
	PKs := CreateTestPubKeys(len(validators))
	validatorsLenBefore := len(su.helperStakingkeeper.GetAllValidators(ctx))
	commission := stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(0, 1), sdk.NewDecWithPrec(0, 1), sdk.NewDec(0))
	delCoin := sdk.NewCoin(testenv.DefaultTestDenom, delegatePerValidator)
	for i, valAddr := range validators {
		su.bankUtils.AddCoinsToAccount(ctx, sdk.NewCoins(delCoin), valAddr.Bytes())
		su.CreateValidator(ctx, valAddr, PKs[i], delCoin, commission)
	}
	require.EqualValues(su.t, validatorsLenBefore+len(validators), len(su.helperStakingkeeper.GetAllValidators(ctx)))
}

func (su *StakingUtils) GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator stakingtypes.Validator, found bool) {
	return su.helperStakingkeeper.GetValidator(ctx, addr)
}

func (su *StakingUtils) MessageDelegate(ctx sdk.Context, expectedCurrentAmountOfDelegations int, expectedCurrentAmountOfUnbondingDelegations int, validatorAddress sdk.ValAddress, delegatorAddress sdk.AccAddress, bondAmount math.Int) {
	validator, found := su.GetValidator(ctx, validatorAddress)
	require.True(su.t, found, "validator not found")
	require.Equal(su.t, expectedCurrentAmountOfDelegations, len(su.helperStakingkeeper.GetAllDelegations(ctx)))
	require.Equal(su.t, expectedCurrentAmountOfUnbondingDelegations, len(su.helperStakingkeeper.GetAllUnbondingDelegations(ctx, delegatorAddress)))
	_, delegatorValidatorFoundBefore := su.helperStakingkeeper.GetDelegatorValidator(ctx, delegatorAddress, validatorAddress)
	newShares, err := su.helperStakingkeeper.Delegate(ctx, delegatorAddress, bondAmount, stakingtypes.Unbonded,
		validator, true)
	require.Nil(su.t, err, "Delegation error")
	require.Truef(su.t, newShares.Equal(sdk.NewDecFromInt(bondAmount)), "newShares %s <> bondAmount %s", newShares, bondAmount)
	if delegatorValidatorFoundBefore == nil {
		require.Equal(su.t, expectedCurrentAmountOfDelegations, len(su.helperStakingkeeper.GetAllDelegations(ctx)))
	} else {
		require.Equal(su.t, expectedCurrentAmountOfDelegations+1, len(su.helperStakingkeeper.GetAllDelegations(ctx)))
	}
	require.Equal(su.t, expectedCurrentAmountOfUnbondingDelegations, len(su.helperStakingkeeper.GetAllUnbondingDelegations(ctx, delegatorAddress)))
}

func (su *StakingUtils) MessageUndelegate(ctx sdk.Context, expectedCurrentAmountOfDelegations int, expectedCurrentAmountOfUnbondingDelegations int, validatorAddress sdk.ValAddress, delegatorAddress sdk.AccAddress, unbondAmount math.Int) {

	require.Equal(su.t, expectedCurrentAmountOfDelegations, len(su.helperStakingkeeper.GetAllDelegations(ctx)))
	require.Equal(su.t, expectedCurrentAmountOfUnbondingDelegations, len(su.helperStakingkeeper.GetAllUnbondingDelegations(ctx, delegatorAddress)))
	_, err := su.helperStakingkeeper.Undelegate(ctx, delegatorAddress, validatorAddress, sdk.NewDecFromInt(unbondAmount))
	require.Nil(su.t, err, "Delegation error")

	//require.Equal(su.t, expectedCurrentAmountOfDelegations, len(su.helperStakingkeeper.GetAllDelegations(ctx))) // TODO check why expectedCurrentAmountOfDelegations not decrements  by 1
	require.Equal(su.t, expectedCurrentAmountOfUnbondingDelegations+1, len(su.helperStakingkeeper.GetAllUnbondingDelegations(ctx, delegatorAddress)))
}

func (su *StakingUtils) VerifyNumberOfUnbondingDelegations(ctx sdk.Context, expectedNumberOfUnbondingDelegations int, delegatorAddress sdk.AccAddress) {
	require.Equal(su.t, expectedNumberOfUnbondingDelegations, len(su.helperStakingkeeper.GetAllUnbondingDelegations(ctx, delegatorAddress)))
}

func (su *StakingUtils) GetStakingDenom(ctx sdk.Context) string {
	return su.helperStakingkeeper.GetParams(ctx).BondDenom
}

type ContextStakingUtils struct {
	StakingUtils
	testContext testenv.TestContext
}

func NewContextStakingUtils(t require.TestingT, testContext testenv.TestContext, helperStakingkeeper stakingkeeper.Keeper, bankUtils *BankUtils) *ContextStakingUtils {
	stakingUtils := NewStakingUtils(t, helperStakingkeeper, bankUtils)
	return &ContextStakingUtils{StakingUtils: stakingUtils, testContext: testContext}
}

func (su *ContextStakingUtils) CreateValidator(addr sdk.ValAddress, pk cryptotypes.PubKey, coin sdk.Coin, commisions stakingtypes.CommissionRates) {
	su.StakingUtils.CreateValidator(su.testContext.GetContext(), addr, pk, coin, commisions)
}

func (su *ContextStakingUtils) SetupValidators(validators []sdk.ValAddress, delegatePerValidator math.Int) {
	su.StakingUtils.SetupValidators(su.testContext.GetContext(), validators, delegatePerValidator)
}

func (su *ContextStakingUtils) GetValidator(addr sdk.ValAddress) (validator stakingtypes.Validator, found bool) {
	return su.StakingUtils.GetValidator(su.testContext.GetContext(), addr)
}

func (su *ContextStakingUtils) MessageDelegate(expectedCurrentAmountOfDelegations int, expectedCurrentAmountOfUnbondingDelegations int,
	validatorAddress sdk.ValAddress, delegatorAddress sdk.AccAddress, bondAmount math.Int) {
	su.StakingUtils.MessageDelegate(su.testContext.GetContext(), expectedCurrentAmountOfDelegations, expectedCurrentAmountOfUnbondingDelegations,
		validatorAddress, delegatorAddress, bondAmount)

}

func (su *ContextStakingUtils) MessageUndelegate(expectedCurrentAmountOfDelegations int, expectedCurrentAmountOfUnbondingDelegations int, validatorAddress sdk.ValAddress, delegatorAddress sdk.AccAddress, unbondAmount math.Int) {
	su.StakingUtils.MessageUndelegate(su.testContext.GetContext(), expectedCurrentAmountOfDelegations, expectedCurrentAmountOfUnbondingDelegations,
		validatorAddress, delegatorAddress, unbondAmount)
}

func (su *ContextStakingUtils) VerifyNumberOfUnbondingDelegations(expectedNumberOfUnbondingDelegations int, delegatorAddress sdk.AccAddress) {
	su.StakingUtils.VerifyNumberOfUnbondingDelegations(su.testContext.GetContext(), expectedNumberOfUnbondingDelegations, delegatorAddress)
}

func (su *ContextStakingUtils) GetValidators() []stakingtypes.Validator {
	return su.StakingUtils.GetValidators(su.testContext.GetContext())
}

func (su *StakingUtils) GetValidators(ctx sdk.Context) []stakingtypes.Validator {
	return su.helperStakingkeeper.GetValidators(ctx, 100)
}

func (su *ContextStakingUtils) GetStakingDenom() string {
	return su.StakingUtils.GetStakingDenom(su.testContext.GetContext())
}

func (su *ContextStakingUtils) ExportGenesisAndValidate() {
	su.StakingUtils.ExportGenesisAndValidate(su.testContext.GetContext())
}

func (su *StakingUtils) ExportGenesisAndValidate(ctx sdk.Context) {
	exportedGenesis := su.helperStakingkeeper.ExportGenesis(ctx)
	err := staking.ValidateGenesis(exportedGenesis)
	require.NoError(su.t, err)
}
