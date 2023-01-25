package keeper

import (
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	"strconv"
	"time"

	metrics "github.com/armon/go-metrics"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

const VESTING_ADDRESS = "vestingAddr: "

func (k Keeper) CreateVestingPool(ctx sdk.Context, addr string, name string, amount sdk.Int, duration time.Duration, vestingType string) error {
	k.Logger(ctx).Debug("create vesting pool", "addr", addr, "amount: ", amount, "vestingType", vestingType)
	_, err := k.GetVestingType(ctx, vestingType)
	if err != nil {
		k.Logger(ctx).Debug("create vesting pool get vesting type error", "error", err.Error())
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, sdkerrors.Wrap(err, "create vesting pool - get vesting type error").Error())
	}
	if duration <= 0 {
		return sdkerrors.Wrap(c4eerrors.ErrParam, "add vesting pool - duration is <= 0 or nil")
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

	if vestingPoolName == "" {
		k.Logger(ctx).Debug("add vesting pool: empty name ", "vestingPoolName", vestingPoolName, VESTING_ADDRESS, vestingAddr, "coinSrcAddr", coinSrcAddr, "amount", amount)
		return sdkerrors.Wrap(c4eerrors.ErrParam, "add vesting pool empty name")
	}

	if amount.LTE(sdk.ZeroInt()) {
		k.Logger(ctx).Debug("add vesting pool amount <= 0", "vestingPoolName", vestingPoolName, VESTING_ADDRESS, vestingAddr, "coinSrcAddr", coinSrcAddr, "amount", amount)
		return sdkerrors.Wrap(c4eerrors.ErrAmount, "add vesting pool - amount is <= 0")
	}

	_, err := sdk.AccAddressFromBech32(vestingAddr)
	if err != nil {
		k.Logger(ctx).Debug("add vesting pool vesting acc address parsing error",
			"vestingPoolName", vestingPoolName, VESTING_ADDRESS, vestingAddr, "coinSrcAddr", coinSrcAddr, "error", err.Error())
		return sdkerrors.Wrap(c4eerrors.ErrParsing, sdkerrors.Wrap(err, "add vesting pool - vesting acc address error").Error())
	}
	denom := k.GetParams(ctx).Denom

	var srcAccAddress sdk.AccAddress
	srcAccAddress, err = sdk.AccAddressFromBech32(coinSrcAddr)
	if err != nil {
		k.Logger(ctx).Debug("add vesting pool source account address parsing error",
			"vestingPoolName", vestingPoolName, VESTING_ADDRESS, vestingAddr, "coinSrcAddr", coinSrcAddr, "error", err.Error())
		return sdkerrors.Wrap(c4eerrors.ErrParsing, sdkerrors.Wrap(err, "add vesting pool - source account address error").Error())
	}

	balance := k.bank.GetBalance(ctx, srcAccAddress, denom)

	if balance.Amount.LT(amount) {
		k.Logger(ctx).Debug("add vesting pool balance less than requested amount error",
			"vestingPoolName", vestingPoolName, VESTING_ADDRESS, vestingAddr, "coinSrcAddr", coinSrcAddr, "balanceAmount",
			balance.Amount.String(), "balanceDenom", balance.Denom, "reguestedAmount", amount.String()+denom)
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "add vesting pool - balance [%s%s] less than requested amount: %s%s",
			balance.Amount.String(), balance.Denom, amount.String(), denom)
	}

	accVestingPools, vestingPoolsFound := k.GetAccountVestingPools(ctx, vestingAddr)
	k.Logger(ctx).Debug("data", "denom", denom, "balance", balance, "vestingPoolsFound", vestingPoolsFound)
	if !vestingPoolsFound {
		accVestingPools = types.AccountVestingPools{}
		accVestingPools.Address = vestingAddr
	} else {
		for _, pool := range accVestingPools.VestingPools {
			if pool.Name == vestingPoolName {
				k.Logger(ctx).Debug("add vesting pool vesting pool name already exists error",
					"vestingPoolName", vestingPoolName, VESTING_ADDRESS, vestingAddr, "coinSrcAddr")
				return sdkerrors.Wrapf(c4eerrors.ErrAlreadyExists, "add vesting pool - vesting pool name: %s", vestingPoolName)
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
	err = k.bank.SendCoinsFromAccountToModule(ctx, srcAccAddress, types.ModuleName, coinsToSend)

	k.Logger(ctx).Debug("add vesting pool", "vestingPoolName", vestingPoolName, VESTING_ADDRESS, vestingAddr, "coinSrcAddr", coinSrcAddr,
		"amount", amount, "vestingType", vestingType, "lockStart", lockStart, "lockEnd", lockEnd,
		"denom", denom, "balance", balance, "vestingPoolsFound", vestingPoolsFound, "vestingPool", vestingPool)
	if err != nil {
		k.Logger(ctx).Error("add vesting pool sendig coins to vesting pool error", "error", err.Error())
		return sdkerrors.Wrap(c4eerrors.ErrSendCoins, sdkerrors.Wrap(err, "add vesting pool  - sendig coins to vesting pool error").Error())
	}
	k.SetAccountVestingPools(ctx, accVestingPools)
	return nil
}

func (k Keeper) WithdrawAllAvailable(ctx sdk.Context, address string) (withdrawn sdk.Coin, returnedError error) {
	accAddress, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		k.Logger(ctx).Debug("withdraw all available address parsing error", "address", address, "error", err.Error())
		return withdrawn, sdkerrors.Wrap(c4eerrors.ErrParsing, sdkerrors.Wrapf(err, "withdraw all available address parsing error: %s", address).Error())
	}

	accVestingPools, vestingPoolsFound := k.GetAccountVestingPools(ctx, address)
	if !vestingPoolsFound {
		k.Logger(ctx).Debug("withdraw all available no vesting pools found error", "address", address)
		return withdrawn, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "withdraw all available - no vesting pools found error: address: %s", address)
	}

	if len(accVestingPools.VestingPools) == 0 {
		k.Logger(ctx).Debug("withdraw all available no vesting pools in array error", "address", address)
		return withdrawn, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "withdraw all available - no vesting pools in array error: address: %s", address)
	}

	current := ctx.BlockTime()
	toWithdraw := sdk.ZeroInt()
	events := make([]types.WithdrawAvailable, 0)
	denom := k.GetParams(ctx).Denom
	for _, vestingPool := range accVestingPools.VestingPools {
		withdrawable := CalculateWithdrawable(current, *vestingPool)
		vestingPool.Withdrawn = vestingPool.Withdrawn.Add(withdrawable)
		toWithdraw = toWithdraw.Add(withdrawable)
		k.Logger(ctx).Debug("withdraw all available data", "address", address, "vestingPool", vestingPool, "withdrawable", withdrawable,
			"toWithdraw", toWithdraw)
		if toWithdraw.IsPositive() {
			events = append(events, types.WithdrawAvailable{
				OwnerAddress:    address,
				VestingPoolName: vestingPool.Name,
				Amount:          toWithdraw.String() + denom,
			})
		}
	}

	if toWithdraw.GT(sdk.ZeroInt()) {
		coinToSend := sdk.NewCoin(denom, toWithdraw)
		coinsToSend := sdk.NewCoins(coinToSend)
		err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, accAddress, coinsToSend)
		if err != nil {
			k.Logger(ctx).Error("withdraw all available sending coins to vesting account error", "address", address, "error", err.Error())
			return withdrawn, sdkerrors.Wrap(c4eerrors.ErrSendCoins, sdkerrors.Wrapf(err, "withdraw all available - send coins to vesting account error: address: %s", address).Error())
		}
	}

	k.SetAccountVestingPools(ctx, accVestingPools)
	k.Logger(ctx).Debug("set account vesting pools", "accAddress", accVestingPools.Address, "newVestingPools", accVestingPools.VestingPools)
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
	k.Logger(ctx).Debug("withdraw all available ret", "address", address, "result", result)
	return result, nil
}

func (k Keeper) SendToNewVestingAccount(ctx sdk.Context, fromAddr string, toAddr string, vestingPoolName string, amount sdk.Int, restartVesting bool) (withdrawn sdk.Coin, returnedError error) {
	k.Logger(ctx).Debug("send to new vesting account", "fromAddr", fromAddr, "toAddr", toAddr, "vestingPoolName", vestingPoolName, "amount", amount, "restartVesting", restartVesting)

	if amount.LTE(sdk.ZeroInt()) {
		return withdrawn, sdkerrors.Wrap(c4eerrors.ErrParam, "send to new vesting account - amount is <= 0")
	}
	if fromAddr == toAddr {
		k.Logger(ctx).Debug("send to new vesting account from address and to address cannot be identical error", "fromAddr", fromAddr, "toAddr", toAddr)
		return withdrawn, sdkerrors.Wrapf(types.ErrIdenticalAccountsAddresses, "send to new vesting account - identical from address (%s) and to address (%s)", fromAddr, toAddr)
	}

	w, err := k.WithdrawAllAvailable(ctx, fromAddr)
	if err != nil {
		k.Logger(ctx).Debug("send to new vesting account withdraw all available error", "fromAddr", fromAddr, "error", err.Error())
		return withdrawn, sdkerrors.Wrap(err, "send to new vesting account - withdraw all available error")
	}

	accVestingPools, vestingPoolsFound := k.GetAccountVestingPools(ctx, fromAddr)
	if !vestingPoolsFound || len(accVestingPools.VestingPools) == 0 {
		k.Logger(ctx).Debug("send to new vesting account no vesting pools found", "fromAddr", fromAddr)
		return withdrawn, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "send to new vesting account - no vesting pools found for address (%s)", fromAddr)
	}
	var vestingPool *types.VestingPool = nil
	for _, vest := range accVestingPools.VestingPools {
		if vest.Name == vestingPoolName {
			vestingPool = vest
		}
	}
	if vestingPool == nil {
		k.Logger(ctx).Debug("send to new vesting account vesting pool id not found", "fromAddr", fromAddr, "vestingPoolName", vestingPoolName)
		return withdrawn, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "send to new vesting account - vesting pool with name %s not found", vestingPoolName)
	}
	available := vestingPool.GetCurrentlyLocked()

	if available.LT(amount) {
		k.Logger(ctx).Debug("send to new vesting account vesting available is smaller than amount", "fromAddr", fromAddr, "vestingPoolName", vestingPoolName, "available", available, "amount", amount)
		return withdrawn, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds,
			"send to new vesting account - vesting available: %s is smaller than requested amount: %s", available, amount)
	}
	vestingPool.Sent = vestingPool.Sent.Add(amount)
	vt, vErr := k.GetVestingType(ctx, vestingPool.VestingType)
	if vErr != nil {
		k.Logger(ctx).Debug("send to new vesting account get vesting type error", "fromAddr", fromAddr, "vestingPool", vestingPool, "error", err.Error())
		return withdrawn, sdkerrors.Wrap(types.ErrGetVestingType, sdkerrors.Wrapf(err, "send to new vesting account - from addr: %s, vestingType %s", fromAddr, vestingPool.VestingType).Error())
	}
	if restartVesting {
		err = k.newVestingAccount(ctx, toAddr, amount, vt.Free,
			ctx.BlockTime().Add(vt.LockupPeriod), ctx.BlockTime().Add(vt.LockupPeriod).Add(vt.VestingPeriod))
	} else {
		err = k.newVestingAccount(ctx, toAddr, amount, vt.Free,
			vestingPool.LockEnd, vestingPool.LockEnd)
	}
	if err == nil {
		k.SetAccountVestingPools(ctx, accVestingPools)
		k.AppendVestingAccount(ctx, types.VestingAccount{Address: toAddr})

		eventErr := ctx.EventManager().EmitTypedEvent(&types.NewVestingAccountFromVestingPool{
			OwnerAddress:    fromAddr,
			Address:         toAddr,
			VestingPoolName: vestingPool.Name,
			Amount:          amount.String() + k.Denom(ctx),
			RestartVesting:  strconv.FormatBool(restartVesting),
		})
		if eventErr != nil {
			k.Logger(ctx).Error("new vesting account from vesting pool emit event error", "error", err.Error())
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

	if amount.IsAnyNegative() {
		return sdkerrors.Wrap(c4eerrors.ErrParam, "create vesting account - negative coin amount")
	}
	if startTime > endTime {
		k.Logger(ctx).Debug("create vesting account start time is after end time", "startTime", startTime, "endTime", endTime)
		return sdkerrors.Wrapf(c4eerrors.ErrParam, "create vesting account - start time is after end time error (%s > %s)", time.Unix(startTime, 0).String(), time.Unix(endTime, 0).String())
	}

	if err := bk.IsSendEnabledCoins(ctx, amount...); err != nil {
		k.Logger(ctx).Debug("create vesting account send coins disabled", "error", err.Error())
		return sdkerrors.Wrap(err, "create vesting account - send coins disabled")
	}

	from, err := sdk.AccAddressFromBech32(fromAddress)
	if err != nil {
		k.Logger(ctx).Debug("create vesting account from-address parsing error", "fromAddress", fromAddress, "error", err.Error())
		return sdkerrors.Wrap(c4eerrors.ErrParsing, sdkerrors.Wrapf(err, "create vesting account - from-address parsing error: %s", fromAddress).Error())
	}
	to, err := sdk.AccAddressFromBech32(toAddress)
	if err != nil {
		k.Logger(ctx).Debug("create vesting account to-address parsing error", "toAddress", toAddress, "error", err.Error())
		return sdkerrors.Wrap(c4eerrors.ErrParsing, sdkerrors.Wrapf(err, "create vesting account - to-address parsing error: %s", toAddress).Error())
	}

	if bk.BlockedAddr(to) {
		k.Logger(ctx).Debug("create vesting account account is not allowed to receive funds error", "toAddress", toAddress)
		return sdkerrors.Wrapf(c4eerrors.ErrAccountNotAllowedToReceiveFunds, "create vesting account - account address: %s", toAddress)
	}

	if acc := ak.GetAccount(ctx, to); acc != nil {
		k.Logger(ctx).Debug("create vesting account account already exists error", "toAddress", toAddress)
		return sdkerrors.Wrapf(c4eerrors.ErrAlreadyExists, "create vesting account - account address: %s", toAddress)
	}

	baseAccount := ak.NewAccountWithAddress(ctx, to)
	if _, ok := baseAccount.(*authtypes.BaseAccount); !ok {
		k.Logger(ctx).Debug("create vesting account invalid account type; expected: BaseAccount", "notExpectedAccount", baseAccount)
		return sdkerrors.Wrapf(c4eerrors.ErrInvalidAccountType, "create vesting account - expected BaseAccount, got: %T", baseAccount)
	}

	baseVestingAccount := vestingtypes.NewBaseVestingAccount(baseAccount.(*authtypes.BaseAccount), amount.Sort(), endTime)

	acc := vestingtypes.NewContinuousVestingAccountRaw(baseVestingAccount, startTime)

	ak.SetAccount(ctx, acc)
	k.Logger(ctx).Debug("create vesting account", "baseAccount", baseVestingAccount.BaseAccount, "originalVesting",
		baseVestingAccount.OriginalVesting, "delegatedFree", baseVestingAccount.DelegatedFree, "delegatedVesting",
		baseVestingAccount.DelegatedVesting, "endTime", baseVestingAccount.EndTime, "startTime", startTime)
	err = ctx.EventManager().EmitTypedEvent(&types.NewVestingAccount{
		Address: acc.Address,
	})
	if err != nil {
		k.Logger(ctx).Error("new vestig account emit event error", "error", err.Error())
	}

	err = bk.SendCoins(ctx, from, to, amount)
	if err != nil {
		k.Logger(ctx).Debug("create vesting account send coins to vesting account error", "fromAddress", fromAddress, "toAddress", toAddress,
			"amount", amount, "error", err.Error())
		return sdkerrors.Wrap(c4eerrors.ErrSendCoins, sdkerrors.Wrapf(err,
			"create vesting account - send coins to vesting account error (from: %s, to: %s, amount: %s)", fromAddress, toAddress, amount).Error())
	}

	k.AppendVestingAccount(ctx, types.VestingAccount{Address: acc.Address})
	k.Logger(ctx).Debug("append vesting account", "address", acc.Address)
	return nil
}

func CalculateWithdrawable(current time.Time, vestingPool types.VestingPool) sdk.Int {
	if current.Equal(vestingPool.LockEnd) || current.After(vestingPool.LockEnd) {
		return vestingPool.GetCurrentlyLocked()
	}
	return sdk.ZeroInt()
}

func (k Keeper) newVestingAccount(ctx sdk.Context, toAddress string, amount sdk.Int, free sdk.Dec,
	lockEnd time.Time,
	vestingEnd time.Time) error {
	denom := k.GetParams(ctx).Denom
	k.Logger(ctx).Debug("new vesting account", "toAddress", toAddress, "amount", amount, "lockEnd", lockEnd,
		"vestingEnd", vestingEnd, "denom", denom)

	ak := k.account
	bk := k.bank
	coinToSend := sdk.NewCoin(denom, amount)
	if err := bk.IsSendEnabledCoins(ctx, coinToSend); err != nil {
		k.Logger(ctx).Debug("new vesting account is send coins disabled error", "error", err.Error())
		return sdkerrors.Wrapf(err, "new vesting account - is send coins disabled")
	}

	to, err := sdk.AccAddressFromBech32(toAddress)
	if err != nil {
		k.Logger(ctx).Debug("new vesting account parsing error", "error", err.Error())
		return sdkerrors.Wrap(c4eerrors.ErrParsing, sdkerrors.Wrapf(err, "new vesting account - to-address parsing error: %s", toAddress).Error())
	}

	if bk.BlockedAddr(to) {
		k.Logger(ctx).Debug("new vesting account is not allowed to receive funds error", "address", toAddress)
		return sdkerrors.Wrapf(c4eerrors.ErrAccountNotAllowedToReceiveFunds, "new vesting account - account address: %s", toAddress)
	}

	if acc := ak.GetAccount(ctx, to); acc != nil {
		k.Logger(ctx).Debug("new vesting account account already exists error", "toAddress", toAddress)
		return sdkerrors.Wrapf(c4eerrors.ErrAlreadyExists, "new vesting account - account address: %s", toAddress)
	}

	baseAccount := ak.NewAccountWithAddress(ctx, to)
	if _, ok := baseAccount.(*authtypes.BaseAccount); !ok {
		k.Logger(ctx).Debug("new vesting account invalid account type; expected: BaseAccount", "toAddress", toAddress, "notExpectedAccount", baseAccount)
		return sdkerrors.Wrapf(c4eerrors.ErrInvalidAccountType, "new vesting account - expected BaseAccount, got: %T", baseAccount)
	}

	decimalAmount := amount.ToDec()
	originalVestingAmount := decimalAmount.Sub(decimalAmount.Mul(free)).TruncateInt()
	originalVestingCoin := sdk.NewCoin(denom, originalVestingAmount)
	originalVesting := sdk.NewCoins(originalVestingCoin)

	baseVestingAccount := vestingtypes.NewBaseVestingAccount(baseAccount.(*authtypes.BaseAccount), originalVesting.Sort(), vestingEnd.Unix())

	var acc authtypes.AccountI

	startTime := lockEnd
	if lockEnd.Before(ctx.BlockTime()) {
		startTime = ctx.BlockTime()
	}

	acc = vestingtypes.NewContinuousVestingAccountRaw(baseVestingAccount, startTime.Unix())

	ak.SetAccount(ctx, acc)
	k.Logger(ctx).Debug("new vesting account", "baseAccount", baseVestingAccount.BaseAccount, "baseVestingAccount",
		baseVestingAccount, "startTime", startTime.Unix())

	err = ctx.EventManager().EmitTypedEvent(&types.NewVestingAccount{
		Address: acc.GetAddress().String(),
	})
	if err != nil {
		k.Logger(ctx).Error("new vesting account emit event error", "error", err.Error())
	}

	coinsToSend := sdk.NewCoins(coinToSend)
	err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, coinsToSend)

	if err != nil {
		k.Logger(ctx).Debug("new vesting account send coins to vesting account error", "error", err.Error())
		return sdkerrors.Wrap(c4eerrors.ErrSendCoins, sdkerrors.Wrapf(err, "new vesting account - send coins to vesting account error").Error())
	}

	return nil
}
