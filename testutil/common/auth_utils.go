package common

import (
	"time"
	// "testing"

	"github.com/chain4energy/c4e-chain/app"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

type AuthUtils struct {
	// t *testing.T
	helperAccountKeeper *authkeeper.AccountKeeper
	bankUtils           *BankUtils
}

func NewAuthUtils(helperAccountKeeper *authkeeper.AccountKeeper, bankUtils *BankUtils) *AuthUtils {
	return &AuthUtils{helperAccountKeeper: helperAccountKeeper, bankUtils: bankUtils}
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

func CreateVestingAccount(ctx sdk.Context, app *app.App, address string, amount sdk.Int, start time.Time, end time.Time) error {
	to := sdk.MustAccAddressFromBech32(address)

	if acc := app.AccountKeeper.GetAccount(ctx, to); acc != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "account %s already exists", address)
	}

	baseAccount := app.AccountKeeper.NewAccountWithAddress(ctx, to)
	if _, ok := baseAccount.(*authtypes.BaseAccount); !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid account type; expected: BaseAccount, got: %T", baseAccount)
	}

	baseVestingAccount := vestingtypes.NewBaseVestingAccount(baseAccount.(*authtypes.BaseAccount), sdk.NewCoins(sdk.NewCoin(DefaultTestDenom, amount)).Sort(), end.Unix())

	acc := vestingtypes.NewContinuousVestingAccountRaw(baseVestingAccount, start.Unix())

	app.AccountKeeper.SetAccount(ctx, acc)

	AddCoinsToAccountInt(amount, ctx, app, to)
	return nil
}
