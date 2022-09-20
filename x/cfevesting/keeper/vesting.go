package keeper

import (
	// "math"
	"strconv"
	"time"

	metrics "github.com/armon/go-metrics"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// vest

func (k Keeper) CreateVestingPool(ctx sdk.Context, addr string, name string, amount sdk.Int, duration time.Duration, vestingType string) error {
	k.Logger(ctx).Debug("Vest: addr: " + addr + "amount: " + amount.String() + "vestingType: " + vestingType)
	_, err := k.GetVestingType(ctx, vestingType)
	if err != nil {
		k.Logger(ctx).Error(err.Error())

		return sdkerrors.Wrap(sdkerrors.ErrNotFound, err.Error())
	}

	return k.addVestingPool(ctx, name, addr, addr, amount, vestingType, ctx.BlockTime(),
		ctx.BlockTime().Add(duration))

}

func (k Keeper) addVestingPool(
	ctx sdk.Context,
	vestingPoolName string,
	vestingAddr string,
	coinSrcAddr string,
	amount sdk.Int,
	vestingType string,
	lockStart time.Time,
	lockEnd time.Time) error {

	if amount.Equal(sdk.ZeroInt()) {
		return nil
	}

	vestingAccAddress, err := sdk.AccAddressFromBech32(vestingAddr)
	if err != nil {
		k.Logger(ctx).Error(err.Error())
		return err
	}

	denom := k.GetParams(ctx).Denom
	k.Logger(ctx).Debug("denom: " + denom)

	var srcAccAddress sdk.AccAddress
	srcAccAddress, err = sdk.AccAddressFromBech32(coinSrcAddr)
	if err != nil {
		k.Logger(ctx).Error(err.Error())
		return err
	}
	k.Logger(ctx).Debug("vestingAccAddress: " + vestingAccAddress.String())

	balance := k.bank.GetBalance(ctx, srcAccAddress, denom)
	k.Logger(ctx).Debug("balance: " + balance.Amount.String())

	if balance.Amount.LT(amount) {
		k.Logger(ctx).Error("Balance [" + balance.Amount.String() +
			balance.Denom + "] lesser than requested amount: " + amount.String() + denom)
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Balance ["+balance.Amount.String()+
			balance.Denom+"] lesser than requested amount: "+amount.String()+denom)
	}

	accVestings, vestingsFound := k.GetAccountVestings(ctx, vestingAddr)
	k.Logger(ctx).Debug("vestingsFound: " + strconv.FormatBool(vestingsFound))
	var id int32
	if !vestingsFound {
		accVestings = types.AccountVestings{}
		accVestings.Address = vestingAddr
		k.Logger(ctx).Debug("accVestings.Address: " + accVestings.Address)
		id = 1
	} else {
		id = int32(len(accVestings.VestingPools)) + 1

		for _, pool := range accVestings.VestingPools {
			if pool.Name == vestingPoolName {
				k.Logger(ctx).Error("vesting pool name already exists: " + vestingPoolName)
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "vesting pool name already exists: "+vestingPoolName)
			}
		}
	}

	vesting := types.VestingPool{
		Id:                        id,
		Name:                      vestingPoolName,
		VestingType:               vestingType,
		LockStart:                 lockStart,
		LockEnd:                   lockEnd,
		Vested:                    amount,
		Withdrawn:                 sdk.ZeroInt(),
		Sent:                      sdk.ZeroInt(),
		LastModification:          lockStart,
		LastModificationVested:    amount,
		LastModificationWithdrawn: sdk.ZeroInt(),
	}
	accVestings.VestingPools = append(accVestings.VestingPools, &vesting)

	coinToSend := sdk.NewCoin(denom, amount)
	coinsToSend := sdk.NewCoins(coinToSend)
	err = k.bank.SendCoinsFromAccountToModule(ctx, srcAccAddress, types.ModuleName, coinsToSend)
	k.Logger(ctx).Info("after SendCoinsFromAccountToModule: " + coinToSend.Amount.String())

	if err != nil {
		k.Logger(ctx).Error(err.Error())
		return err
	}
	k.SetAccountVestings(ctx, accVestings)
	k.Logger(ctx).Info("Vest exit: " + vestingAddr)
	return nil
}

func (k Keeper) WithdrawAllAvailable(ctx sdk.Context, addr string) (withdrawn sdk.Coin, returnedError error) {
	k.Logger(ctx).Debug("WithdrawAllAvailable: " + addr)

	accAddress, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return withdrawn, err
	}

	accVestings, vestingsFound := k.GetAccountVestings(ctx, addr)
	if !vestingsFound {
		return withdrawn, status.Error(codes.NotFound, "No vestings")
	}

	if len(accVestings.VestingPools) == 0 {
		return withdrawn, status.Error(codes.NotFound, "No vestings")
	}

	current := ctx.BlockTime()
	toWithdraw := sdk.ZeroInt()
	events := make([]types.WithdrawAvailable, 0)
	denom := k.GetParams(ctx).Denom
	for _, vesting := range accVestings.VestingPools {
		withdrawable := CalculateWithdrawable(current, *vesting)
		vesting.Withdrawn = vesting.Withdrawn.Add(withdrawable)
		vesting.LastModificationWithdrawn = vesting.LastModificationWithdrawn.Add(withdrawable)
		toWithdraw = toWithdraw.Add(withdrawable)
		if toWithdraw.IsPositive() {
			events = append(events, types.WithdrawAvailable{
				OwnerAddress:    addr,
				VestingPoolId:   strconv.FormatInt(int64(vesting.Id), 10),
				VestingPoolName: vesting.Name,
				Amount:          toWithdraw.String() + denom,
			})
		}
	}

	if toWithdraw.GT(sdk.ZeroInt()) {
		coinToSend := sdk.NewCoin(denom, toWithdraw)
		coinsToSend := sdk.NewCoins(coinToSend)
		err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, accAddress, coinsToSend)
		if err != nil {
			return withdrawn, err
		}
	}

	k.SetAccountVestings(ctx, accVestings)

	if toWithdraw.IsPositive() && toWithdraw.IsInt64() {
		defer func() {
			telemetry.IncrCounter(1, types.ModuleName, "withdraw_available")
			telemetry.IncrCounterWithLabels(
				[]string{"tx", "msg", types.ModuleName, "withdraw_available"},
				float32(withdrawn.Amount.Int64()),
				[]metrics.Label{telemetry.NewLabel("denom", withdrawn.Denom)},
			)
		}()
	}

	for _, event := range events {
		ctx.EventManager().EmitTypedEvent(&event)
	}

	return sdk.NewCoin(denom, toWithdraw), nil
}

func (k Keeper) SendToNewVestingAccount(ctx sdk.Context, fromAddr string, toAddr string, vestingId int32, amount sdk.Int, restartVesting bool) (withdrawn sdk.Coin, returnedError error) {
	if fromAddr == toAddr {
		return withdrawn, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "from address and to address cannot be identical")
	}

	w, err := k.WithdrawAllAvailable(ctx, fromAddr)
	if err != nil {
		return withdrawn, err
	}

	accVestings, vestingsFound := k.GetAccountVestings(ctx, fromAddr)
	if !vestingsFound || len(accVestings.VestingPools) == 0 {
		return withdrawn, sdkerrors.Wrap(sdkerrors.ErrNotFound, "no vesting pools found")
	}
	var vesting *types.VestingPool = nil
	for _, vest := range accVestings.VestingPools {
		if vest.Id == vestingId {
			vesting = vest
		}
	}
	if vesting == nil {
		return withdrawn, sdkerrors.Wrap(sdkerrors.ErrNotFound, "vesting pool with id "+strconv.FormatInt(int64(vestingId), 10)+" not found")
	}
	available := vesting.LastModificationVested.Sub(vesting.LastModificationWithdrawn)
	if available.LT(amount) {
		return withdrawn, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds,
			"vesting available: %s is smaller than %s", available, amount)
	}
	vesting.Sent = amount
	vesting.LastModification = ctx.BlockTime()
	vesting.LastModificationVested = available.Sub(amount)
	vesting.LastModificationWithdrawn = sdk.ZeroInt()

	if restartVesting {
		vt, vErr := k.GetVestingType(ctx, vesting.VestingType)
		if vErr != nil {
			k.Logger(ctx).Error(err.Error())

			return withdrawn, sdkerrors.Wrap(sdkerrors.ErrNotFound, err.Error())
		}
		err = k.createVestingAccount(ctx, toAddr, amount,
			ctx.BlockTime().Add(vt.LockupPeriod), ctx.BlockTime().Add(vt.LockupPeriod).Add(vt.VestingPeriod))
	} else {
		err = k.createVestingAccount(ctx, toAddr, amount,
			vesting.LockEnd, vesting.LockEnd)
	}
	if err == nil {
		k.SetAccountVestings(ctx, accVestings)
		k.AppendVestingAccount(ctx, types.VestingAccount{Address: toAddr})
		ctx.EventManager().EmitTypedEvent(&types.NewVestingAccountFromVestingPool{
			OwnerAddress:    fromAddr,
			Address:         toAddr,
			VestingPoolId:   strconv.FormatInt(int64(vesting.Id), 10),
			VestingPoolName: vesting.Name,
			Amount:          amount.String() + k.Denom(ctx),
			RestartVesting:  strconv.FormatBool(restartVesting),
		})
	}
	return w, err
}

func (k Keeper) CreateVestingAccount(ctx sdk.Context, fromAddress string, toAddress string,
	amount sdk.Coins, startTime int64, endTime int64) error {
	ak := k.account
	bk := k.bank

	if err := bk.IsSendEnabledCoins(ctx, amount...); err != nil {
		return err
	}

	from, err := sdk.AccAddressFromBech32(fromAddress)
	if err != nil {
		return err
	}
	to, err := sdk.AccAddressFromBech32(toAddress)
	if err != nil {
		return err
	}

	if bk.BlockedAddr(to) {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to receive funds", toAddress)
	}

	if acc := ak.GetAccount(ctx, to); acc != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "account %s already exists", toAddress)
	}

	baseAccount := ak.NewAccountWithAddress(ctx, to)
	if _, ok := baseAccount.(*authtypes.BaseAccount); !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid account type; expected: BaseAccount, got: %T", baseAccount)
	}

	baseVestingAccount := vestingtypes.NewBaseVestingAccount(baseAccount.(*authtypes.BaseAccount), amount.Sort(), endTime)

	acc := vestingtypes.NewContinuousVestingAccountRaw(baseVestingAccount, startTime)

	ak.SetAccount(ctx, acc)
	ctx.EventManager().EmitTypedEvent(&types.NewVestingAccount{
		Address: acc.Address,
	})
	err = bk.SendCoins(ctx, from, to, amount)
	if err != nil {
		return err
	}
	k.AppendVestingAccount(ctx, types.VestingAccount{Address: acc.Address})
	return nil
}

func CalculateWithdrawable(current time.Time, vesting types.VestingPool) sdk.Int {
	if current.Equal(vesting.LockEnd) || current.After(vesting.LockEnd) {
		return vesting.LastModificationVested.Sub(vesting.LastModificationWithdrawn)
	}
	return sdk.ZeroInt()
}

func (k Keeper) createVestingAccount(ctx sdk.Context, toAddress string, amount sdk.Int,
	lockEnd time.Time,
	vestingEnd time.Time) error {

	ak := k.account
	bk := k.bank

	denom := k.GetParams(ctx).Denom
	coinToSend := sdk.NewCoin(denom, amount)
	if err := bk.IsSendEnabledCoins(ctx, coinToSend); err != nil {
		return err
	}

	to, err := sdk.AccAddressFromBech32(toAddress)
	if err != nil {
		return err
	}

	if bk.BlockedAddr(to) {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to receive funds", toAddress)
	}

	if acc := ak.GetAccount(ctx, to); acc != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "account %s already exists", toAddress)
	}

	baseAccount := ak.NewAccountWithAddress(ctx, to)
	if _, ok := baseAccount.(*authtypes.BaseAccount); !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid account type; expected: BaseAccount, got: %T", baseAccount)
	}

	coinsToSend := sdk.NewCoins(coinToSend)

	baseVestingAccount := vestingtypes.NewBaseVestingAccount(baseAccount.(*authtypes.BaseAccount), coinsToSend.Sort(), vestingEnd.Unix())

	var acc authtypes.AccountI

	startTime := lockEnd
	if lockEnd.Before(ctx.BlockTime()) {
		startTime = ctx.BlockTime()
	}

	acc = vestingtypes.NewContinuousVestingAccountRaw(baseVestingAccount, startTime.Unix())

	ak.SetAccount(ctx, acc)
	ctx.EventManager().EmitTypedEvent(&types.NewVestingAccount{
		Address: acc.GetAddress().String(),
	})
	err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, coinsToSend)
	k.Logger(ctx).Info("after SendCoinsFromModuleToAccount: " + coinToSend.Amount.String())

	if err != nil {
		return err
	}

	return nil
}
