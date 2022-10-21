package common

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestexported "github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/stretchr/testify/require"
)

type AuthUtils struct {
	t                   *testing.T
	helperAccountKeeper *authkeeper.AccountKeeper
	bankUtils           *BankUtils
}

func NewAuthUtils(t *testing.T, helperAccountKeeper *authkeeper.AccountKeeper, bankUtils *BankUtils) AuthUtils {
	return AuthUtils{t: t, helperAccountKeeper: helperAccountKeeper, bankUtils: bankUtils}
}

func (au *AuthUtils) CreateVestingAccount(ctx sdk.Context, address string, coin sdk.Coin, start time.Time, end time.Time) error {
	to := sdk.MustAccAddressFromBech32(address)

	if acc := au.helperAccountKeeper.GetAccount(ctx, to); acc != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "account %s already exists", address)
	}

	baseAccount := au.helperAccountKeeper.NewAccountWithAddress(ctx, to)
	if _, ok := baseAccount.(*authtypes.BaseAccount); !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid account type; expected: BaseAccount, got: %T", baseAccount)
	}

	baseVestingAccount := vestingtypes.NewBaseVestingAccount(baseAccount.(*authtypes.BaseAccount), sdk.NewCoins(coin).Sort(), end.Unix())

	acc := vestingtypes.NewContinuousVestingAccountRaw(baseVestingAccount, start.Unix())

	au.helperAccountKeeper.SetAccount(ctx, acc)

	au.bankUtils.AddCoinsToAccount(ctx, coin, to)
	return nil
}

func (au *AuthUtils) CreateDefaultDenomVestingAccount(ctx sdk.Context, address string, amount sdk.Int, start time.Time, end time.Time) error {
	return au.CreateVestingAccount(ctx, address, sdk.NewCoin(DefaultTestDenom, amount), start, end)
}

func (au *AuthUtils) VerifyVestingAccount(ctx sdk.Context, address sdk.AccAddress, lockedDenom string, lockedAmount sdk.Int, startTime time.Time, endTime time.Time) {
	account := au.helperAccountKeeper.GetAccount(ctx, address)

	vacc, ok := account.(vestexported.VestingAccount)
	require.True(au.t, ok, ok)
	locked := vacc.LockedCoins(ctx.BlockTime())
	require.Equal(au.t, 1, len(locked))

	require.Equal(au.t, lockedDenom, locked[0].Denom)
	require.Equal(au.t, lockedAmount, locked[0].Amount)

	require.Equal(au.t, endTime.Unix(), vacc.GetEndTime())
	require.Equal(au.t, startTime.Unix(), vacc.GetStartTime())
}

func (au *AuthUtils) VerifyDefaultDenomVestingAccount(ctx sdk.Context, address sdk.AccAddress, lockedAmount sdk.Int, startTime time.Time, endTime time.Time) {
	au.VerifyVestingAccount(ctx, address, DefaultTestDenom, lockedAmount, startTime, endTime)
}

type ContextAuthUtils struct {
	AuthUtils
	testContext TestContext
}

func NewContextAuthUtils(t *testing.T, testContext TestContext, helperAccountKeeper *authkeeper.AccountKeeper, bankUtils *BankUtils) *ContextAuthUtils {
	authUtils := NewAuthUtils(t, helperAccountKeeper, bankUtils)
	return &ContextAuthUtils{AuthUtils: authUtils, testContext: testContext}
}

func (au *ContextAuthUtils) CreateVestingAccount(address string, coin sdk.Coin, start time.Time, end time.Time) error {
	return au.AuthUtils.CreateVestingAccount(au.testContext.GetContext(), address, coin, start, end)
}

func (au *ContextAuthUtils) CreateDefaultDenomVestingAccount(address string, amount sdk.Int, start time.Time, end time.Time) error {
	return au.AuthUtils.CreateDefaultDenomVestingAccount(au.testContext.GetContext(), address, amount, start, end)
}

func (au *ContextAuthUtils) VerifyVestingAccount(address sdk.AccAddress, lockedDenom string, lockedAmount sdk.Int, startTime time.Time, endTime time.Time) {
	au.AuthUtils.VerifyVestingAccount(au.testContext.GetContext(), address, lockedDenom, lockedAmount, startTime, endTime)
}

func (au *ContextAuthUtils) VerifyDefaultDenomVestingAccount(address sdk.AccAddress, lockedAmount sdk.Int, startTime time.Time, endTime time.Time) {
	au.AuthUtils.VerifyDefaultDenomVestingAccount(au.testContext.GetContext(), address, lockedAmount, startTime, endTime)
}
