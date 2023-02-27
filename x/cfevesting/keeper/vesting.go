package keeper

import (
	"fmt"
	"strconv"
	"time"

	"cosmossdk.io/math"

	metrics "github.com/armon/go-metrics"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

const VESTING_ADDRESS = "vestingAddr: "

func (k Keeper) CreateVestingPool(ctx sdk.Context, addr string, name string, amount math.Int, duration time.Duration, vestingType string) error {
	k.Logger(ctx).Debug("create vesting pool", "addr", addr, "amount: ", amount, "vestingType", vestingType)
	_, err := k.GetVestingType(ctx, vestingType)
	if err != nil {
		k.Logger(ctx).Debug("create vesting pool get vesting type error", "error", err.Error())
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, sdkerrors.Wrap(err, "create vesting pool - get vesting type error").Error())
	}
	if duration <= 0 {
		return sdkerrors.Wrap(types.ErrParam, "add vesting pool - duration is <= 0 or nil")
	}
	return k.addVestingPool(ctx, name, addr, addr, amount, vestingType, ctx.BlockTime(),
		ctx.BlockTime().Add(duration))
}

func (k Keeper) addVestingPool(
	ctx sdk.Context,
	vestingPoolName string,
	vestingAddr string,
	coinSrcAddr string,
	amount math.Int,
	vestingType string,
	lockStart time.Time,
	lockEnd time.Time) error {

	if vestingPoolName == "" {
		k.Logger(ctx).Debug("add vesting pool: empty name ", "vestingPoolName", vestingPoolName, VESTING_ADDRESS, vestingAddr, "coinSrcAddr", coinSrcAddr, "amount", amount)
		return sdkerrors.Wrap(types.ErrParam, "add vesting pool empty name")
	}

	if amount.LTE(sdk.ZeroInt()) {
		k.Logger(ctx).Debug("add vesting pool amount <= 0", "vestingPoolName", vestingPoolName, VESTING_ADDRESS, vestingAddr, "coinSrcAddr", coinSrcAddr, "amount", amount)
		return sdkerrors.Wrap(types.ErrAmount, "add vesting pool - amount is <= 0")
	}

	_, err := sdk.AccAddressFromBech32(vestingAddr)
	if err != nil {
		k.Logger(ctx).Debug("add vesting pool vesting acc address parsing error",
			"vestingPoolName", vestingPoolName, VESTING_ADDRESS, vestingAddr, "coinSrcAddr", coinSrcAddr, "error", err.Error())
		return sdkerrors.Wrap(types.ErrParsing, sdkerrors.Wrap(err, "add vesting pool - vesting acc address error").Error())
	}
	denom := k.GetParams(ctx).Denom

	var srcAccAddress sdk.AccAddress
	srcAccAddress, err = sdk.AccAddressFromBech32(coinSrcAddr)
	if err != nil {
		k.Logger(ctx).Debug("add vesting pool source account address parsing error",
			"vestingPoolName", vestingPoolName, VESTING_ADDRESS, vestingAddr, "coinSrcAddr", coinSrcAddr, "error", err.Error())
		return sdkerrors.Wrap(types.ErrParsing, sdkerrors.Wrap(err, "add vesting pool - source account address error").Error())
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
		accVestingPools.Owner = vestingAddr
	} else {
		for _, pool := range accVestingPools.VestingPools {
			if pool.Name == vestingPoolName {
				k.Logger(ctx).Debug("add vesting pool vesting pool name already exists error",
					"vestingPoolName", vestingPoolName, VESTING_ADDRESS, vestingAddr, "coinSrcAddr")
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
	err = k.bank.SendCoinsFromAccountToModule(ctx, srcAccAddress, types.ModuleName, coinsToSend)

	k.Logger(ctx).Debug("add vesting pool", "vestingPoolName", vestingPoolName, VESTING_ADDRESS, vestingAddr, "coinSrcAddr", coinSrcAddr,
		"amount", amount, "vestingType", vestingType, "lockStart", lockStart, "lockEnd", lockEnd,
		"denom", denom, "balance", balance, "vestingPoolsFound", vestingPoolsFound, "vestingPool", vestingPool)
	if err != nil {
		k.Logger(ctx).Error("add vesting pool sendig coins to vesting pool error", "error", err.Error())
		return sdkerrors.Wrap(types.ErrSendCoins, sdkerrors.Wrap(err, "add vesting pool  - sendig coins to vesting pool error").Error())
	}
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

func (k Keeper) SendToNewVestingAccount(ctx sdk.Context, owner string, toAddr string, vestingPoolName string, amount math.Int, restartVesting bool) (withdrawn sdk.Coin, returnedError error) {
	k.Logger(ctx).Debug("send to new vesting account", "owner", owner, "toAddr", toAddr, "vestingPoolName", vestingPoolName, "amount", amount, "restartVesting", restartVesting)

	if amount.LTE(sdk.ZeroInt()) {
		return withdrawn, sdkerrors.Wrap(types.ErrParam, "send to new vesting account - amount is <= 0")
	}
	if owner == toAddr {
		k.Logger(ctx).Debug("send to new vesting account from address and to address cannot be identical error", "owner", owner, "toAddr", toAddr)
		return withdrawn, sdkerrors.Wrapf(types.ErrIdenticalAccountsAddresses, "send to new vesting account - identical from address (%s) and to address (%s)", owner, toAddr)
	}

	w, err := k.WithdrawAllAvailable(ctx, owner)
	if err != nil {
		k.Logger(ctx).Debug("send to new vesting account withdraw all available error", "owner", owner, "error", err.Error())
		return withdrawn, sdkerrors.Wrap(err, "send to new vesting account - withdraw all available error")
	}

	accVestingPools, vestingPoolsFound := k.GetAccountVestingPools(ctx, owner)
	if !vestingPoolsFound || len(accVestingPools.VestingPools) == 0 {
		k.Logger(ctx).Debug("send to new vesting account no vesting pools found", "owner", owner)
		return withdrawn, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "send to new vesting account - no vesting pools found for address (%s)", owner)
	}
	var vestingPool *types.VestingPool = nil
	for _, vest := range accVestingPools.VestingPools {
		if vest.Name == vestingPoolName {
			vestingPool = vest
		}
	}
	if vestingPool == nil {
		k.Logger(ctx).Debug("send to new vesting account vesting pool id not found", "owner", owner, "vestingPoolName", vestingPoolName)
		return withdrawn, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "send to new vesting account - vesting pool with name %s not found", vestingPoolName)
	}
	available := vestingPool.GetCurrentlyLocked()

	if available.LT(amount) {
		k.Logger(ctx).Debug("send to new vesting account vesting available is smaller than amount", "owner", owner, "vestingPoolName", vestingPoolName, "available", available, "amount", amount)
		return withdrawn, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds,
			"send to new vesting account - vesting available: %s is smaller than requested amount: %s", available, amount)
	}
	vestingPool.Sent = vestingPool.Sent.Add(amount)
	vt, vErr := k.GetVestingType(ctx, vestingPool.VestingType)
	if vErr != nil {
		k.Logger(ctx).Debug("send to new vesting account get vesting type error", "owner", owner, "vestingPool", vestingPool, "error", err.Error())
		return withdrawn, sdkerrors.Wrap(types.ErrGetVestingType, sdkerrors.Wrapf(err, "send to new vesting account - from addr: %s, vestingType %s", owner, vestingPool.VestingType).Error())
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
		k.AppendVestingAccountTrace(ctx, types.VestingAccountTrace{
			Address:            toAddr,
			Genesis:            false,
			FromGenesisPool:    vestingPool.GensisPool,
			FromGenesisAccount: false,
		})

		eventErr := ctx.EventManager().EmitTypedEvent(&types.NewVestingAccountFromVestingPool{
			Owner:           owner,
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
		return sdkerrors.Wrap(types.ErrParam, "create vesting account - negative coin amount")
	}
	if startTime > endTime {
		k.Logger(ctx).Debug("create vesting account start time is after end time", "startTime", startTime, "endTime", endTime)
		return sdkerrors.Wrapf(types.ErrParam, "create vesting account - start time is after end time error (%s > %s)", time.Unix(startTime, 0).String(), time.Unix(endTime, 0).String())
	}

	if err := bk.IsSendEnabledCoins(ctx, amount...); err != nil {
		k.Logger(ctx).Debug("create vesting account send coins disabled", "error", err.Error())
		return sdkerrors.Wrap(err, "create vesting account - send coins disabled")
	}

	from, err := sdk.AccAddressFromBech32(fromAddress)
	if err != nil {
		k.Logger(ctx).Debug("create vesting account from-address parsing error", "fromAddress", fromAddress, "error", err.Error())
		return sdkerrors.Wrap(types.ErrParsing, sdkerrors.Wrapf(err, "create vesting account - from-address parsing error: %s", fromAddress).Error())
	}
	to, err := sdk.AccAddressFromBech32(toAddress)
	if err != nil {
		k.Logger(ctx).Debug("create vesting account to-address parsing error", "toAddress", toAddress, "error", err.Error())
		return sdkerrors.Wrap(types.ErrParsing, sdkerrors.Wrapf(err, "create vesting account - to-address parsing error: %s", toAddress).Error())
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
		return vestingPool.GetCurrentlyLocked()
	}
	return sdk.ZeroInt()
}

func (k Keeper) newVestingAccount(ctx sdk.Context, toAddress string, amount math.Int, free sdk.Dec,
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
		return sdkerrors.Wrap(types.ErrParsing, sdkerrors.Wrapf(err, "new vesting account - to-address parsing error: %s", toAddress).Error())
	}

	if bk.BlockedAddr(to) {
		k.Logger(ctx).Debug("new vesting account is not allowed to receive funds error", "address", toAddress)
		return sdkerrors.Wrapf(types.ErrAccountNotAllowedToReceiveFunds, "new vesting account - account address: %s", toAddress)
	}

	if acc := ak.GetAccount(ctx, to); acc != nil {
		k.Logger(ctx).Debug("new vesting account account already exists error", "toAddress", toAddress)
		return sdkerrors.Wrapf(types.ErrAlreadyExists, "new vesting account - account address: %s", toAddress)
	}

	decimalAmount := sdk.NewDecFromInt(amount)
	originalVestingAmount := decimalAmount.Sub(decimalAmount.Mul(free)).TruncateInt()
	originalVestingCoin := sdk.NewCoin(denom, originalVestingAmount)
	originalVesting := sdk.NewCoins(originalVestingCoin)

	startTime := lockEnd
	if lockEnd.Before(ctx.BlockTime()) {
		startTime = ctx.BlockTime()
	}

	_, err = k.newContinuousVestingAccount(ctx, to, originalVesting, startTime.Unix(), vestingEnd.Unix())
	if err != nil {
		k.Logger(ctx).Debug("new vesting account - to account creation error", "error", err.Error())
		return sdkerrors.Wrap(err, fmt.Sprintf("new vesting account - to account creation error: %s", toAddress))
	}

	coinsToSend := sdk.NewCoins(coinToSend)
	err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, coinsToSend)

	if err != nil {
		k.Logger(ctx).Debug("new vesting account send coins to vesting account error", "error", err.Error())
		return sdkerrors.Wrap(types.ErrSendCoins, sdkerrors.Wrapf(err, "new vesting account - send coins to vesting account error").Error())
	}

	return nil
}

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
