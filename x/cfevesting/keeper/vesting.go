package keeper

import (
	"cosmossdk.io/errors"
	"fmt"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"strconv"
	"time"

	"cosmossdk.io/math"

	"github.com/armon/go-metrics"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

const VestingAddress = "vestingAddr: "

func (k Keeper) CreateVestingPool(ctx sdk.Context, addr string, name string, amount math.Int, duration time.Duration, vestingType string) error {
	k.Logger(ctx).Debug("create vesting pool", "addr", addr, "amount: ", amount, "vestingType", vestingType)
	_, err := k.GetVestingType(ctx, vestingType)
	if err != nil {
		k.Logger(ctx).Debug("create vesting pool get vesting type error", "error", err.Error())
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, sdkerrors.Wrap(err, "create vesting pool - get vesting type error").Error())
	}
	accAddress, err := types.ValidateCreateVestingPool(addr, name, amount, duration)
	if err != nil {
		k.Logger(ctx).Debug("create vesting pool validation error", "error", err.Error())
		return err
	}
	return k.addVestingPool(ctx, name, accAddress, amount, vestingType, ctx.BlockTime(),
		ctx.BlockTime().Add(duration))
}

func (k Keeper) addVestingPool(
	ctx sdk.Context,
	vestingPoolName string,
	accAddress sdk.AccAddress,
	amount math.Int,
	vestingType string,
	lockStart time.Time,
	lockEnd time.Time) error {

	denom := k.GetParams(ctx).Denom

	balance := k.bank.GetBalance(ctx, accAddress, denom)

	if balance.Amount.LT(amount) {
		k.Logger(ctx).Debug("add vesting pool balance less than requested amount error",
			"vestingPoolName", vestingPoolName, VestingAddress, accAddress, "balanceAmount",
			balance.Amount.String(), "balanceDenom", balance.Denom, "reguestedAmount", amount.String()+denom)
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "add vesting pool - balance [%s%s] less than requested amount: %s%s",
			balance.Amount.String(), balance.Denom, amount.String(), denom)
	}

	accVestingPools, vestingPoolsFound := k.GetAccountVestingPools(ctx, accAddress.String())
	k.Logger(ctx).Debug("data", "denom", denom, "balance", balance, "vestingPoolsFound", vestingPoolsFound)
	if !vestingPoolsFound {
		accVestingPools = types.AccountVestingPools{}
		accVestingPools.Owner = accAddress.String()
	} else {
		for _, pool := range accVestingPools.VestingPools {
			if pool.Name == vestingPoolName {
				k.Logger(ctx).Debug("add vesting pool vesting pool name already exists error",
					"vestingPoolName", vestingPoolName, VestingAddress, accAddress)
				return sdkerrors.Wrapf(types.ErrAlreadyExists, "add vesting pool - vesting pool name: %s", vestingPoolName)
			}
		}
	}

	vestingPool := types.VestingPool{
		Name:            vestingPoolName,
		VestingType:     vestingType,
		LockStart:       lockStart,
		LockEnd:         lockEnd,
		InitiallyLocked: amount,
		Withdrawn:       sdk.ZeroInt(),
		Sent:            sdk.ZeroInt(),
	}
	accVestingPools.VestingPools = append(accVestingPools.VestingPools, &vestingPool)

	coinToSend := sdk.NewCoin(denom, amount)
	coinsToSend := sdk.NewCoins(coinToSend)
	err := k.bank.SendCoinsFromAccountToModule(ctx, accAddress, types.ModuleName, coinsToSend)
	if err != nil {
		k.Logger(ctx).Error("add vesting pool sendig coins to vesting pool error", "error", err.Error())
		return sdkerrors.Wrap(types.ErrSendCoins, sdkerrors.Wrap(err, "add vesting pool  - sendig coins to vesting pool error").Error())
	}

	k.Logger(ctx).Debug("add vesting pool", "vestingPoolName", vestingPoolName, VestingAddress, accAddress,
		"amount", amount, "vestingType", vestingType, "lockStart", lockStart, "lockEnd", lockEnd,
		"denom", denom, "balance", balance, "vestingPoolsFound", vestingPoolsFound, "vestingPool", vestingPool)

	k.SetAccountVestingPools(ctx, accVestingPools)
	return nil
}

func (k Keeper) WithdrawAllAvailable(ctx sdk.Context, owner string) (withdrawn sdk.Coin, returnedError error) {
	ownerAddress, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		k.Logger(ctx).Debug("withdraw all available owner parsing error", "owner", owner, "error", err.Error())
		return withdrawn, sdkerrors.Wrap(types.ErrParsing, sdkerrors.Wrapf(err, "withdraw all available owner parsing error: %s", owner).Error())
	}

	accVestingPools, vestingPoolsFound := k.GetAccountVestingPools(ctx, owner)
	if !vestingPoolsFound {
		k.Logger(ctx).Debug("withdraw all available no vesting pools found error", "owner", owner)
		return withdrawn, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "withdraw all available - no vesting pools found error: owner: %s", owner)
	}

	if len(accVestingPools.VestingPools) == 0 {
		k.Logger(ctx).Debug("withdraw all available no vesting pools in array error", "owner", owner)
		return withdrawn, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "withdraw all available - no vesting pools in array error: owner: %s", owner)
	}

	current := ctx.BlockTime()
	toWithdraw := sdk.ZeroInt()
	events := make([]types.WithdrawAvailable, 0)
	denom := k.GetParams(ctx).Denom
	for _, vestingPool := range accVestingPools.VestingPools {
		withdrawable := CalculateWithdrawable(current, *vestingPool)
		vestingPool.Withdrawn = vestingPool.Withdrawn.Add(withdrawable)
		toWithdraw = toWithdraw.Add(withdrawable)
		k.Logger(ctx).Debug("withdraw all available data", "owner", owner, "vestingPool", vestingPool, "withdrawable", withdrawable,
			"toWithdraw", toWithdraw)
		if toWithdraw.IsPositive() {
			events = append(events, types.WithdrawAvailable{
				Owner:           owner,
				VestingPoolName: vestingPool.Name,
				Amount:          toWithdraw.String() + denom,
			})
		}
	}

	if toWithdraw.GT(sdk.ZeroInt()) {
		coinToSend := sdk.NewCoin(denom, toWithdraw)
		coinsToSend := sdk.NewCoins(coinToSend)
		err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, ownerAddress, coinsToSend)
		if err != nil {
			k.Logger(ctx).Error("withdraw all available sending coins to vesting account error", "owner", owner, "error", err.Error())
			return withdrawn, sdkerrors.Wrap(types.ErrSendCoins, sdkerrors.Wrapf(err, "withdraw all available - send coins to vesting account error: owner: %s", owner).Error())
		}
	}

	k.SetAccountVestingPools(ctx, accVestingPools)
	k.Logger(ctx).Debug("set account vesting pools", "ownerAddress", accVestingPools.Owner, "newVestingPools", accVestingPools.VestingPools)
	if toWithdraw.IsPositive() && toWithdraw.IsInt64() {
		defer func() {
			telemetry.SetGaugeWithLabels(
				[]string{"tx", "msg", types.ModuleName, "withdraw_available"},
				float32(withdrawn.Amount.Int64()),
				[]metrics.Label{telemetry.NewLabel("denom", withdrawn.Denom)},
			)
		}()
	}

	for _, event := range events {
		err := ctx.EventManager().EmitTypedEvent(&event)
		if err != nil {
			k.Logger(ctx).Error("withdraw all available emit event error", "error", err.Error())
		}
	}
	result := sdk.NewCoin(denom, toWithdraw)
	k.Logger(ctx).Debug("withdraw all available ret", "owner", owner, "result", result)
	return result, nil
}

func (k Keeper) SendToNewVestingAccountFromReservation(ctx sdk.Context, owner string, toAddr string, vestingPoolName string, amount math.Int, reservationId uint64, startTime time.Time, endTime time.Time) error {
	k.Logger(ctx).Debug("send to new vesting account", "owner", owner, "toAddr", toAddr, "vestingPoolName", vestingPoolName, "amount", amount, "reservationId", reservationId)

	accVestingPools, vestingPool, vestingPoolsFound := k.GetAccountVestingPool(ctx, owner, vestingPoolName)
	if !vestingPoolsFound || len(accVestingPools.VestingPools) == 0 || vestingPool == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "send to new vesting account - no vesting pools found for address (%s)", owner)
	}

	available := vestingPool.GetCurrentlyLockedInReservations()

	if available.LT(amount) {
		return errors.Wrapf(sdkerrors.ErrInsufficientFunds, "send to new vesting account - vesting available: %s is smaller than requested amount: %s", available, amount)
	}
	if err := vestingPool.SubstractFromReservation(reservationId, amount); err != nil {
		return err
	}

	periodId, err := k.SendToPeriodicContinuousVestingAccountFromModule(ctx, types.ModuleName, toAddr, sdk.NewCoins(sdk.NewCoin("uc4e", amount)), startTime.Unix(), endTime.Unix())
	if err != nil {
		return err
	}

	k.SetAccountVestingPools(ctx, accVestingPools)
	vestingAccountTrace, found := k.GetVestingAccountTrace(ctx, toAddr)
	if !found {
		k.AppendVestingAccountTrace(ctx, types.VestingAccountTrace{
			Address:            toAddr,
			Genesis:            false,
			FromGenesisPool:    vestingPool.GenesisPool,
			FromGenesisAccount: false,
			PeriodsToTrace:     []uint64{periodId},
		})
	} else {
		vestingAccountTrace.PeriodsToTrace = append(vestingAccountTrace.PeriodsToTrace, periodId)
	}

	eventErr := ctx.EventManager().EmitTypedEvent(&types.NewVestingAccountFromVestingPool{ // TODO : modify
		Owner:           owner,
		Address:         toAddr,
		VestingPoolName: vestingPool.Name,
		Amount:          amount.String() + k.Denom(ctx),
		RestartVesting:  "false",
	})
	if eventErr != nil {
		k.Logger(ctx).Error("new vesting account from vesting pool emit event error", "error", eventErr.Error())
	}

	return nil
}

func (k Keeper) SendToNewVestingAccountFromLocked(ctx sdk.Context, owner string, toAddr string, vestingPoolName string, amount math.Int, restartVesting bool) (withdrawn sdk.Coin, returnedError error) {
	k.Logger(ctx).Debug("send to new vesting account", "owner", owner, "toAddr", toAddr, "vestingPoolName", vestingPoolName, "amount", amount, "restartVesting", restartVesting)

	w, err := k.WithdrawAllAvailable(ctx, owner) // TODO: maybe to delete?
	if err != nil {
		k.Logger(ctx).Debug("send to new vesting account withdraw all available error", "owner", owner, "error", err.Error()) // TODO: logs strategy?
		return withdrawn, sdkerrors.Wrap(err, "send to new vesting account - withdraw all available error")                   // TODO: maybe make errors easier?
	}

	accVestingPools, vestingPool, vestingPoolsFound := k.GetAccountVestingPool(ctx, owner, vestingPoolName)
	if !vestingPoolsFound || len(accVestingPools.VestingPools) == 0 || vestingPool == nil {
		return w, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "send to new vesting account - no vesting pools found for address (%s)", owner)
	}

	available := vestingPool.GetCurrentlyLockedWithoutReservations()

	if available.LT(amount) {
		return w, errors.Wrapf(sdkerrors.ErrInsufficientFunds,
			"send to new vesting account - vesting available: %s is smaller than requested amount: %s", available, amount)
	}
	vestingPool.Sent = vestingPool.Sent.Add(amount)
	vt, vErr := k.GetVestingType(ctx, vestingPool.VestingType)
	if vErr != nil {
		k.Logger(ctx).Debug("send to new vesting account get vesting type error", "owner", owner, "vestingPool", vestingPool, "error", err.Error())
		return withdrawn, sdkerrors.Wrap(types.ErrGetVestingType, sdkerrors.Wrapf(err, "send to new vesting account - from addr: %s, vestingType %s", owner, vestingPool.VestingType).Error())
	}
	coinsToSend := sdk.NewCoins(sdk.NewCoin(k.Denom(ctx), amount))
	var periodId uint64
	if restartVesting {
		periodId, err = k.SendToPeriodicContinuousVestingAccountFromModule(ctx, types.ModuleName, toAddr, coinsToSend,
			ctx.BlockTime().Add(vt.LockupPeriod).Unix(), ctx.BlockTime().Add(vt.LockupPeriod).Add(vt.VestingPeriod).Unix())
	} else {
		periodId, err = k.SendToPeriodicContinuousVestingAccountFromModule(ctx, types.ModuleName, toAddr, coinsToSend,
			vestingPool.LockEnd.Unix(), vestingPool.LockEnd.Unix())
	}
	if err == nil {
		k.SetAccountVestingPools(ctx, accVestingPools)
		vestingAccountTrace, found := k.GetVestingAccountTrace(ctx, toAddr)
		if !found {
			k.AppendVestingAccountTrace(ctx, types.VestingAccountTrace{
				Address:            toAddr,
				Genesis:            false,
				FromGenesisPool:    vestingPool.GenesisPool,
				FromGenesisAccount: false,
				PeriodsToTrace:     []uint64{periodId},
			})
		} else {
			vestingAccountTrace.PeriodsToTrace = append(vestingAccountTrace.PeriodsToTrace, periodId)
		}

		eventErr := ctx.EventManager().EmitTypedEvent(&types.NewVestingAccountFromVestingPool{
			Owner:           owner,
			Address:         toAddr,
			VestingPoolName: vestingPool.Name,
			Amount:          amount.String() + k.Denom(ctx),
			RestartVesting:  strconv.FormatBool(restartVesting),
		})
		if eventErr != nil {
			k.Logger(ctx).Error("new vesting account from vesting pool emit event error", "error", eventErr.Error())
		}
	}
	k.Logger(ctx).Debug("send to new vesting account ret", "withdrawn", w, "error", err, "accVestingPools", accVestingPools)
	return w, err
}

func (k Keeper) CreateVestingAccount(ctx sdk.Context, fromAddress string, toAddress string,
	amount sdk.Coins, startTime int64, endTime int64) error {
	k.Logger(ctx).Debug("create vesting account", "fromAddress", fromAddress, "toAddress", toAddress,
		"amount", amount, "startTime", startTime, "endTime", endTime)
	ak := k.account
	bk := k.bank

	from, to, err := types.ValidateCreateVestingAccount(fromAddress, toAddress, amount, startTime, endTime)
	if err != nil {
		k.Logger(ctx).Debug("create vesting account validation error", "error", err.Error())
		return err
	}

	if err = bk.IsSendEnabledCoins(ctx, amount...); err != nil {
		k.Logger(ctx).Debug("create vesting account send coins disabled", "error", err.Error())
		return sdkerrors.Wrap(err, "create vesting account - send coins disabled")
	}

	if bk.BlockedAddr(to) {
		k.Logger(ctx).Debug("create vesting account account is not allowed to receive funds error", "toAddress", toAddress)
		return sdkerrors.Wrapf(types.ErrAccountNotAllowedToReceiveFunds, "create vesting account - account address: %s", toAddress)
	}

	if acc := ak.GetAccount(ctx, to); acc != nil {
		k.Logger(ctx).Debug("create vesting account account already exists error", "toAddress", toAddress)
		return sdkerrors.Wrapf(types.ErrAlreadyExists, "create vesting account - account address: %s", toAddress)
	}

	acc, err := k.newContinuousVestingAccount(ctx, to, amount.Sort(), startTime, endTime)
	if err != nil {
		k.Logger(ctx).Debug("create vesting account - to account creation error", "error", err.Error())
		return sdkerrors.Wrap(err, fmt.Sprintf("new vesting account - to account creation error: %s", toAddress))
	}

	err = bk.SendCoins(ctx, from, to, amount)
	if err != nil {
		k.Logger(ctx).Debug("create vesting account send coins to vesting account error", "fromAddress", fromAddress, "toAddress", toAddress,
			"amount", amount, "error", err.Error())
		return sdkerrors.Wrap(types.ErrSendCoins, sdkerrors.Wrapf(err,
			"create vesting account - send coins to vesting account error (from: %s, to: %s, amount: %s)", fromAddress, toAddress, amount).Error())
	}

	k.Logger(ctx).Debug("append vesting account", "address", acc.Address)
	return nil
}

func CalculateWithdrawable(current time.Time, vestingPool types.VestingPool) math.Int {
	if current.Equal(vestingPool.LockEnd) || current.After(vestingPool.LockEnd) {
		return vestingPool.GetCurrentlyLockedWithoutReservations()
	}
	return sdk.ZeroInt()
}

//func (k Keeper) newVestingAccount(ctx sdk.Context, toAddress sdk.AccAddress, amount math.Int, free sdk.Dec, // TODO: probably to delte
//	lockEnd time.Time,
//	vestingEnd time.Time) error {
//	denom := k.GetParams(ctx).Denom
//	k.Logger(ctx).Debug("new vesting account", "toAddress", toAddress, "amount", amount, "lockEnd", lockEnd,
//		"vestingEnd", vestingEnd, "denom", denom)
//
//	ak := k.account
//	bk := k.bank
//	coinToSend := sdk.NewCoin(denom, amount)
//	if err := bk.IsSendEnabledCoins(ctx, coinToSend); err != nil {
//		k.Logger(ctx).Debug("new vesting account is send coins disabled error", "error", err.Error())
//		return sdkerrors.Wrapf(err, "new vesting account - is send coins disabled")
//	}
//
//	if bk.BlockedAddr(toAddress) {
//		k.Logger(ctx).Debug("new vesting account is not allowed to receive funds error", "address", toAddress)
//		return sdkerrors.Wrapf(types.ErrAccountNotAllowedToReceiveFunds, "new vesting account - account address: %s", toAddress)
//	}
//
//	if acc := ak.GetAccount(ctx, toAddress); acc != nil {
//		k.Logger(ctx).Debug("new vesting account account already exists error", "toAddress", toAddress)
//		return sdkerrors.Wrapf(types.ErrAlreadyExists, "new vesting account - account address: %s", toAddress)
//	}
//
//	decimalAmount := sdk.NewDecFromInt(amount)
//	originalVestingAmount := decimalAmount.Sub(decimalAmount.Mul(free)).TruncateInt()
//	originalVestingCoin := sdk.NewCoin(denom, originalVestingAmount)
//	originalVesting := sdk.NewCoins(originalVestingCoin)
//
//	startTime := lockEnd
//	if lockEnd.Before(ctx.BlockTime()) {
//		startTime = ctx.BlockTime()
//	}
//
//	_, err := k.newContinuousVestingAccount(ctx, toAddress, originalVesting, startTime.Unix(), vestingEnd.Unix())
//	if err != nil {
//		k.Logger(ctx).Debug("new vesting account - to account creation error", "error", err.Error())
//		return sdkerrors.Wrap(err, fmt.Sprintf("new vesting account - to account creation error: %s", toAddress))
//	}
//
//	coinsToSend := sdk.NewCoins(coinToSend)
//	err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, toAddress, coinsToSend)
//
//	if err != nil {
//		k.Logger(ctx).Debug("new vesting account send coins to vesting account error", "error", err.Error())
//		return sdkerrors.Wrap(types.ErrSendCoins, sdkerrors.Wrapf(err, "new vesting account - send coins to vesting account error").Error())
//	}
//
//	return nil
//}

func (k Keeper) newContinuousVestingAccount(ctx sdk.Context, to sdk.AccAddress, originalVesting sdk.Coins, startTime int64, vestingEnd int64) (*vestingtypes.ContinuousVestingAccount, error) {
	baseAccount := k.account.NewAccountWithAddress(ctx, to)
	if _, ok := baseAccount.(*authtypes.BaseAccount); !ok {
		k.Logger(ctx).Debug("new continuous vesting account invalid account type; expected: BaseAccount", "toAddress", to, "notExpectedAccount", baseAccount)
		return nil, sdkerrors.Wrapf(types.ErrInvalidAccountType, "new continuous vesting account - expected BaseAccount, got: %T", baseAccount)
	}
	baseVestingAccount := vestingtypes.NewBaseVestingAccount(baseAccount.(*authtypes.BaseAccount), originalVesting.Sort(), vestingEnd)

	acc := vestingtypes.NewContinuousVestingAccountRaw(baseVestingAccount, startTime)
	defer telemetry.IncrCounter(1, "new", "account")
	k.account.SetAccount(ctx, acc)
	k.Logger(ctx).Debug("new continuous vesting account", "baseAccount", baseVestingAccount.BaseAccount, "baseVestingAccount",
		baseVestingAccount, "startTime", startTime)
	err := ctx.EventManager().EmitTypedEvent(&types.NewVestingAccount{
		Address: acc.Address,
	})
	if err != nil {
		k.Logger(ctx).Error("new vestig account emit event error", "error", err.Error())
	}
	return acc, nil
}

func (k Keeper) AddVestingPoolReservation(ctx sdk.Context, owner string, vestingPoolName string, reservationId uint64, amout math.Int) error {
	vestingDenom := k.Denom(ctx)
	accountVestingPools, found := k.GetAccountVestingPools(ctx, owner)
	if !found {
		return errors.Wrapf(c4eerrors.ErrNotExists, "vesting pools not found for address %s", owner)
	}

	found = false
	var vestingPool *types.VestingPool
	for _, vestPool := range accountVestingPools.VestingPools {
		if vestPool.Name == vestingPoolName {
			vestingPool = vestPool
			currentlyLocked := vestingPool.GetCurrentlyLockedWithoutReservations()
			if currentlyLocked.LT(amout) {
				return errors.Wrapf(sdkerrors.ErrInsufficientFunds, "%s%s is smaller than %s%s", currentlyLocked, vestingDenom, amout, vestingDenom)
			}
			found = true
			break
		}
	}
	if !found {
		return errors.Wrapf(c4eerrors.ErrNotExists, "vesting pool %s not found for address %s", vestingPoolName, owner)
	}
	// TODO : add event
	vestingPool.AddReservation(reservationId, amout)

	k.SetAccountVestingPools(ctx, accountVestingPools)
	return nil
}

func (k Keeper) RemoveVestingPoolReservation(ctx sdk.Context, owner string, vestingPoolName string, reservationId uint64, amout math.Int) error {
	accountVestingPools, found := k.GetAccountVestingPools(ctx, owner)
	if !found {
		return errors.Wrapf(c4eerrors.ErrNotExists, "vesting pools not found for address %s", owner)
	}

	found = false
	var vestingPool *types.VestingPool
	for _, vestPool := range accountVestingPools.VestingPools {
		if vestPool.Name == vestingPoolName {
			vestingPool = vestPool
			found = true
			break
		}
	}
	if !found {
		return errors.Wrapf(c4eerrors.ErrNotExists, "vesting pool %s not found for address %s", vestingPoolName, owner)
	}
	// TODO : add event
	if err := vestingPool.SubstractFromReservation(reservationId, amout); err != nil {
		return err
	}

	k.SetAccountVestingPools(ctx, accountVestingPools)
	return nil
}
