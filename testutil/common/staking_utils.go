package common

import (
	"testing"
	"github.com/stretchr/testify/require"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type StakingUtils struct {
	t *testing.T
	helperStakingkeeper stakingkeeper.Keeper
	bankUtils           *BankUtils
}

func NewStakingUtils(t *testing.T, helperStakingkeeper stakingkeeper.Keeper, bankUtils *BankUtils) *StakingUtils {
	return &StakingUtils{t: t, helperStakingkeeper: helperStakingkeeper, bankUtils: bankUtils}
}

func (su *StakingUtils) CreateValidator(ctx sdk.Context, sk stakingkeeper.Keeper, addr sdk.ValAddress, pk cryptotypes.PubKey, coin sdk.Coin, commisions stakingtypes.CommissionRates) {
	msg, err := stakingtypes.NewMsgCreateValidator(addr, pk, coin, stakingtypes.Description{}, commisions, sdk.OneInt())
	msgSrvr := stakingkeeper.NewMsgServerImpl(sk)
	require.NoError(su.t, err)
	res, err := msgSrvr.CreateValidator(sdk.WrapSDKContext(ctx), msg)
	require.NoError(su.t, err)
	require.NotNil(su.t, res)
}

func (su *StakingUtils) SetupValidators(ctx sdk.Context, validators []sdk.ValAddress, delegatePerValidator sdk.Int) {
	PKs := CreateTestPubKeys(len(validators))
	commission := stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(0, 1), sdk.NewDecWithPrec(0, 1), sdk.NewDec(0))
	delCoin := sdk.NewCoin(DefaultTestDenom, delegatePerValidator)
	for i, valAddr := range validators {
		su.bankUtils.AddCoinsToAccount(ctx, delCoin, valAddr.Bytes())
		su.CreateValidator(ctx, su.helperStakingkeeper, valAddr, PKs[i], delCoin, commission)
	}
	require.EqualValues(su.t, len(validators)+1, len(su.helperStakingkeeper.GetAllValidators(ctx)))
}