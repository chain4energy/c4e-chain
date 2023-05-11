package cosmossdk

import (
	"cosmossdk.io/errors"
	"fmt"
	"time"

	"cosmossdk.io/math"

	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestexported "github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/stretchr/testify/require"
)

type AuthUtils struct {
	t                   require.TestingT
	helperAccountKeeper *authkeeper.AccountKeeper
	bankUtils           *BankUtils
}

func NewAuthUtils(t require.TestingT, helperAccountKeeper *authkeeper.AccountKeeper, bankUtils *BankUtils) AuthUtils {
	return AuthUtils{t: t, helperAccountKeeper: helperAccountKeeper, bankUtils: bankUtils}
}

func (au *AuthUtils) ModifyVestingAccountOriginalVesting(ctx sdk.Context, address string, newOrignalVestings sdk.Coins) error {
	bechAddress := sdk.MustAccAddressFromBech32(address)
	ownerAccount := au.helperAccountKeeper.GetAccount(ctx, bechAddress)
	if ownerAccount == nil {
		return fmt.Errorf("account %s doesn't exist", address)
	}

	vestingAcc, ok := ownerAccount.(*vestingtypes.ContinuousVestingAccount)
	if !ok {
		return fmt.Errorf("account %s is not ContinuousVestingAccount", address)
	}
	vestingAcc.OriginalVesting = newOrignalVestings
	au.helperAccountKeeper.SetAccount(ctx, vestingAcc)
	return nil
}

func (au *AuthUtils) CreateVestingAccount(ctx sdk.Context, address string, coins sdk.Coins, start time.Time, end time.Time) error {
	to := sdk.MustAccAddressFromBech32(address)

	if acc := au.helperAccountKeeper.GetAccount(ctx, to); acc != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "account %s already exists", address)
	}

	baseAccount := au.helperAccountKeeper.NewAccountWithAddress(ctx, to)
	if _, ok := baseAccount.(*authtypes.BaseAccount); !ok {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid account type; expected: BaseAccount, got: %T", baseAccount)
	}

	baseVestingAccount := vestingtypes.NewBaseVestingAccount(baseAccount.(*authtypes.BaseAccount), coins.Sort(), end.Unix())

	acc := vestingtypes.NewContinuousVestingAccountRaw(baseVestingAccount, start.Unix())

	au.helperAccountKeeper.SetAccount(ctx, acc)

	au.bankUtils.AddCoinsToAccount(ctx, coins, to)
	return nil
}

func (au *AuthUtils) CreateBaseAccount(ctx sdk.Context, address string, coin sdk.Coins) error {
	to := sdk.MustAccAddressFromBech32(address)

	if acc := au.helperAccountKeeper.GetAccount(ctx, to); acc != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "account %s already exists", address)
	}

	baseAccount := au.helperAccountKeeper.NewAccountWithAddress(ctx, to)
	if _, ok := baseAccount.(*authtypes.BaseAccount); !ok {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid account type; expected: BaseAccount, got: %T", baseAccount)
	}

	au.helperAccountKeeper.SetAccount(ctx, baseAccount)

	au.bankUtils.AddCoinsToAccount(ctx, coin, to)
	return nil
}

func (au *AuthUtils) CreateDefaultDenomVestingAccount(ctx sdk.Context, address string, amount math.Int, start time.Time, end time.Time) error {
	return au.CreateVestingAccount(ctx, address, sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, amount)), start, end)
}

func (au *AuthUtils) CreateDefaultDenomBaseAccount(ctx sdk.Context, address string, amount math.Int) error {
	return au.CreateBaseAccount(ctx, address, sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, amount)))
}

func (au *AuthUtils) VerifyIsContinuousVestingAccount(ctx sdk.Context, address sdk.AccAddress) {
	account := au.helperAccountKeeper.GetAccount(ctx, address)
	require.NotNil(au.t, account)
	_, ok := account.(*vestingtypes.ContinuousVestingAccount)
	require.True(au.t, ok)
}

func (au *AuthUtils) VerifyAccountDoesNotExist(ctx sdk.Context, address sdk.AccAddress) {
	account := au.helperAccountKeeper.GetAccount(ctx, address)
	require.Nil(au.t, account)
}

func (au *AuthUtils) VerifyVestingAccount(ctx sdk.Context, address sdk.AccAddress, lockedAmount sdk.Coins, startTime time.Time, endTime time.Time) {
	account := au.helperAccountKeeper.GetAccount(ctx, address)

	vacc, ok := account.(vestexported.VestingAccount)
	require.True(au.t, ok, ok)
	locked := vacc.LockedCoins(ctx.BlockTime())
	require.True(au.t, locked.IsEqual(lockedAmount))

	require.Equal(au.t, endTime.Unix(), vacc.GetEndTime())
	require.Equal(au.t, startTime.Unix(), vacc.GetStartTime())
}

func (au *AuthUtils) VerifyDefaultDenomVestingAccount(ctx sdk.Context, address sdk.AccAddress, lockedAmount math.Int, startTime time.Time, endTime time.Time) {
	au.VerifyVestingAccount(ctx, address, sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, lockedAmount)), startTime, endTime)
}

type ContextAuthUtils struct {
	AuthUtils
	testContext testenv.TestContext
}

func NewContextAuthUtils(t require.TestingT, testContext testenv.TestContext, helperAccountKeeper *authkeeper.AccountKeeper, bankUtils *BankUtils) *ContextAuthUtils {
	authUtils := NewAuthUtils(t, helperAccountKeeper, bankUtils)
	return &ContextAuthUtils{AuthUtils: authUtils, testContext: testContext}
}

func (au *ContextAuthUtils) CreateVestingAccount(address string, coins sdk.Coins, start time.Time, end time.Time) error {
	return au.AuthUtils.CreateVestingAccount(au.testContext.GetContext(), address, coins, start, end)
}

func (au *ContextAuthUtils) CreateDefaultDenomVestingAccount(address string, amount math.Int, start time.Time, end time.Time) error {
	return au.AuthUtils.CreateDefaultDenomVestingAccount(au.testContext.GetContext(), address, amount, start, end)
}

func (au *ContextAuthUtils) VerifyVestingAccount(address sdk.AccAddress, lockedAmount sdk.Coins, startTime time.Time, endTime time.Time) {
	au.AuthUtils.VerifyVestingAccount(au.testContext.GetContext(), address, lockedAmount, startTime, endTime)
}

func (au *ContextAuthUtils) VerifyDefaultDenomVestingAccount(address sdk.AccAddress, lockedAmount math.Int, startTime time.Time, endTime time.Time) {
	au.AuthUtils.VerifyDefaultDenomVestingAccount(au.testContext.GetContext(), address, lockedAmount, startTime, endTime)
}

func (au *ContextAuthUtils) ModifyVestingAccountOriginalVesting(address string, newOrignalVestings sdk.Coins) error {
	return au.AuthUtils.ModifyVestingAccountOriginalVesting(au.testContext.GetContext(), address, newOrignalVestings)
}

func (au *ContextAuthUtils) CreateDefaultDenomBaseAccount(address string, amount sdk.Int) error {
	return au.AuthUtils.CreateDefaultDenomBaseAccount(au.testContext.GetContext(), address, amount)
}

func (au *ContextAuthUtils) VerifyIsContinuousVestingAccount(address sdk.AccAddress) {
	au.AuthUtils.VerifyIsContinuousVestingAccount(au.testContext.GetContext(), address)
}

func (au *ContextAuthUtils) VerifyAccountDoesNotExist(address sdk.AccAddress) {
	au.AuthUtils.VerifyAccountDoesNotExist(au.testContext.GetContext(), address)

}
