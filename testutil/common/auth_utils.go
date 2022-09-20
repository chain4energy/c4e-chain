package common

import (
	"time"

	"github.com/chain4energy/c4e-chain/app"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

func CreateVestingAccount(ctx sdk.Context, app *app.App, address string, amount sdk.Int, start time.Time, end time.Time) error {
	to := sdk.MustAccAddressFromBech32(address)

	if acc := app.AccountKeeper.GetAccount(ctx, to); acc != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "account %s already exists", address)
	}

	baseAccount := app.AccountKeeper.NewAccountWithAddress(ctx, to)
	if _, ok := baseAccount.(*authtypes.BaseAccount); !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid account type; expected: BaseAccount, got: %T", baseAccount)
	}

	baseVestingAccount := vestingtypes.NewBaseVestingAccount(baseAccount.(*authtypes.BaseAccount), sdk.NewCoins(sdk.NewCoin(Denom, amount)).Sort(), end.Unix())

	acc := vestingtypes.NewContinuousVestingAccountRaw(baseVestingAccount, start.Unix())

	app.AccountKeeper.SetAccount(ctx, acc)

	AddCoinsToAccountInt(amount, ctx, app, to)
	return nil
}
