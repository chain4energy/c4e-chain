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
		return errors.Wrap(sdkerrors.ErrNotFound, errors.Wrap(err, "create vesting pool - get vesting type error").Error())
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
		return errors.Wrapf(sdkerrors.ErrInsufficientFunds, "add vesting pool - balance [%s%s] less than requested amount: %s%s",
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
				return errors.Wrapf(types.ErrAlreadyExists, "add vesting pool - vesting pool name: %s", vestingPoolName)
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
		return errors.Wrap(types.ErrSendCoins, errors.Wrap(err, "add vesting pool  - sendig coins to vesting pool error").Error())
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
		return withdrawn, errors.Wrap(types.ErrParsing, errors.Wrapf(err, "withdraw all available owner parsing error: %s", owner).Error())
	}

	accVestingPools, vestingPoolsFound := k.GetAccountVestingPools(ctx, owner)
	if !vestingPoolsFound {
		k.Logger(ctx).Debug("withdraw all available no vesting pools found error", "owner", owner)
		return withdrawn, errors.Wrapf(sdkerrors.ErrNotFound, "withdraw all available - no vesting pools found error: owner: %s", owner)
	}

	if len(accVestingPools.VestingPools) == 0 {
		k.Logger(ctx).Debug("withdraw all available no vesting pools in array error", "owner", owner)
		return withdrawn, errors.Wrapf(sdkerrors.ErrNotFound, "withdraw all available - no vesting pools in array error: owner: %s", owner)
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
			return withdrawn, errors.Wrap(types.ErrSendCoins, errors.Wrapf(err, "withdraw all available - send coins to vesting account error: owner: %s", owner).Error())
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

func (k Keeper) getVestingPoolAndType(ctx sdk.Context, owner string, vestingPoolName string) (*types.AccountVestingPools, *types.VestingPool, *types.VestingType, error) {
	accVestingPools, vestingPool, vestingPoolsFound := k.GetAccountVestingPool(ctx, owner, vestingPoolName)
	if !vestingPoolsFound || len(accVestingPools.VestingPools) == 0 || vestingPool == nil {
		return nil, nil, nil, errors.Wrapf(sdkerrors.ErrNotFound, "no vesting pool %s found for address %s", vestingPoolName, owner)
	}

	vestingType, err := k.GetVestingType(ctx, vestingPool.VestingType)
	if err != nil {
		return nil, nil, nil, errors.Wrap(types.ErrGetVestingType, errors.Wrapf(err, "from addr: %s, vestingType %s", owner, vestingPool.VestingType).Error())
	}
	return &accVestingPools, vestingPool, &vestingType, nil
}

// The SendReservedToNewVestingAccount function sends reserved tokens from the vesting pool to a new vesting account.
// The function validates whether the lockupPeriod and vestingPeriod are greater than or equal to the vesting periods
// specified in the vesting type. Additionally, the free (sdk.Dec) is validated, and if it exceeds the value specified
// in the vesting type, instead of returning an error, it is set to the free value specified in the vesting type.
func (k Keeper) SendReservedToNewVestingAccount(ctx sdk.Context, owner string, toAddr string, vestingPoolName string, amount math.Int, reservationId uint64,
	free sdk.Dec, lockupPeriod time.Duration, vestingPeriod time.Duration) error {
	k.Logger(ctx).Debug("send reserved to new vesting account", "owner", owner, "toAddr", toAddr, "vestingPoolName", vestingPoolName, "amount", amount, "reservationId", reservationId)

	accVestingPools, vestingPool, vestingType, err := k.getVestingPoolAndType(ctx, owner, vestingPoolName)
	if err != nil {
		k.Logger(ctx).Debug("send reserved to new vesting account get vesting type error", "owner", owner, "vestingPool", vestingPool, "error", err.Error())
		return errors.Wrap(err, "send reserved to new vesting account")
	}

	if err = vestingType.ValidateVestingPeriods(lockupPeriod, vestingPeriod); err != nil {
		return err
	}

	if free.GT(vestingType.Free) {
		free = vestingType.Free
	}

	if err = vestingPool.SendFromReservedTokens(reservationId, amount); err != nil {
		return err
	}

	denom := k.Denom(ctx)
	periodId, err := k.SendToPeriodicContinuousVestingAccountFromModule(ctx, types.ModuleName, toAddr, sdk.NewCoins(sdk.NewCoin(denom, amount)), free,
		ctx.BlockTime().Add(lockupPeriod).Unix(), ctx.BlockTime().Add(lockupPeriod).Add(vestingPeriod).Unix())
	if err != nil {
		return err
	}

	k.SetAccountVestingPoolsAndVestingAccountTrace(ctx, owner, toAddr, amount, periodId, false, accVestingPools, vestingPool)
	return nil
}

func (k Keeper) SendToNewVestingAccount(ctx sdk.Context, owner string, toAddr string, vestingPoolName string, amount math.Int, restartVesting bool) error {
	k.Logger(ctx).Debug("send to new vesting account", "owner", owner, "toAddr", toAddr, "vestingPoolName", vestingPoolName, "amount", amount, "restartVesting", restartVesting)

	accVestingPools, vestingPool, vestingType, err := k.getVestingPoolAndType(ctx, owner, vestingPoolName)
	if err != nil {
		k.Logger(ctx).Debug("send locked to new vesting account get vesting type error", "owner", owner, "vestingPool", vestingPool, "error", err.Error())
		return errors.Wrap(err, "send locked to new vesting account")
	}

	if err = vestingPool.SendFromLockedTokens(amount); err != nil {
		return err
	}

	denom := k.Denom(ctx)
	var periodId uint64
	if restartVesting {
		periodId, err = k.SendToPeriodicContinuousVestingAccountFromModule(ctx, types.ModuleName, toAddr, sdk.NewCoins(sdk.NewCoin(denom, amount)), vestingType.Free,
			ctx.BlockTime().Add(vestingType.LockupPeriod).Unix(), ctx.BlockTime().Add(vestingType.LockupPeriod).Add(vestingType.VestingPeriod).Unix())
	} else {
		periodId, err = k.SendToPeriodicContinuousVestingAccountFromModule(ctx, types.ModuleName, toAddr, sdk.NewCoins(sdk.NewCoin(denom, amount)), vestingType.Free,
			vestingPool.LockEnd.Unix(), vestingPool.LockEnd.Unix())
	}
	if err != nil {
		return err
	}

	k.SetAccountVestingPoolsAndVestingAccountTrace(ctx, owner, toAddr, amount, periodId, restartVesting, accVestingPools, vestingPool)
	k.Logger(ctx).Debug("send to new vesting account ret", "error", err, "accVestingPools", accVestingPools)
	return nil
}

func (k Keeper) SetAccountVestingPoolsAndVestingAccountTrace(ctx sdk.Context, owner string, toAddr string, amount math.Int, periodId uint64,
	restartVesting bool, accVestingPools *types.AccountVestingPools, vestingPool *types.VestingPool) {
	k.SetAccountVestingPools(ctx, *accVestingPools)
	vestingAccountTrace, found := k.GetVestingAccountTrace(ctx, toAddr)

	if !found {
		vestingAccountTrace = types.VestingAccountTrace{
			Address:            toAddr,
			Genesis:            false,
			FromGenesisPool:    vestingPool.GenesisPool,
			FromGenesisAccount: false,
			PeriodsToTrace:     []uint64{},
		}
		if vestingPool.GenesisPool {
			vestingAccountTrace.PeriodsToTrace = []uint64{periodId}
		}
		k.AppendVestingAccountTrace(ctx, vestingAccountTrace)
	} else if vestingPool.GenesisPool {
		vestingAccountTrace.PeriodsToTrace = append(vestingAccountTrace.PeriodsToTrace, periodId)
		k.SetVestingAccountTrace(ctx, vestingAccountTrace)
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
		return errors.Wrap(err, "create vesting account - send coins disabled")
	}

	if bk.BlockedAddr(to) {
		k.Logger(ctx).Debug("create vesting account account is not allowed to receive funds error", "toAddress", toAddress)
		return errors.Wrapf(types.ErrAccountNotAllowedToReceiveFunds, "create vesting account - account address: %s", toAddress)
	}

	if acc := ak.GetAccount(ctx, to); acc != nil {
		k.Logger(ctx).Debug("create vesting account account already exists error", "toAddress", toAddress)
		return errors.Wrapf(types.ErrAlreadyExists, "create vesting account - account address: %s", toAddress)
	}

	acc, err := k.newContinuousVestingAccount(ctx, to, amount.Sort(), startTime, endTime)
	if err != nil {
		k.Logger(ctx).Debug("create vesting account - to account creation error", "error", err.Error())
		return errors.Wrap(err, fmt.Sprintf("new vesting account - to account creation error: %s", toAddress))
	}

	err = bk.SendCoins(ctx, from, to, amount)
	if err != nil {
		k.Logger(ctx).Debug("create vesting account send coins to vesting account error", "fromAddress", fromAddress, "toAddress", toAddress,
			"amount", amount, "error", err.Error())
		return errors.Wrap(types.ErrSendCoins, errors.Wrapf(err,
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

func (k Keeper) newContinuousVestingAccount(ctx sdk.Context, to sdk.AccAddress, originalVesting sdk.Coins, startTime int64, vestingEnd int64) (*vestingtypes.ContinuousVestingAccount, error) {
	baseAccount := k.account.NewAccountWithAddress(ctx, to)
	if _, ok := baseAccount.(*authtypes.BaseAccount); !ok {
		k.Logger(ctx).Debug("new continuous vesting account invalid account type; expected: BaseAccount", "toAddress", to, "notExpectedAccount", baseAccount)
		return nil, errors.Wrapf(types.ErrInvalidAccountType, "new continuous vesting account - expected BaseAccount, got: %T", baseAccount)
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
	accountVestingPools, vestingPool, found := k.GetAccountVestingPool(ctx, owner, vestingPoolName)
	if !found {
		return errors.Wrapf(c4eerrors.ErrNotExists, "vesting pools not found for address %s", owner)
	}

	currentlyLocked := vestingPool.GetCurrentlyLockedWithoutReservations()
	if currentlyLocked.LT(amout) {
		return errors.Wrapf(sdkerrors.ErrInsufficientFunds, "%s%s is smaller than %s%s", currentlyLocked, vestingDenom, amout, vestingDenom)
	}

	vestingPool.AddReservation(reservationId, amout)

	k.SetAccountVestingPools(ctx, accountVestingPools)
	return nil
}

func (k Keeper) RemoveVestingPoolReservation(ctx sdk.Context, owner string, vestingPoolName string, reservationId uint64, amout math.Int) error {
	accountVestingPools, vestingPool, found := k.GetAccountVestingPool(ctx, owner, vestingPoolName)
	if !found {
		return errors.Wrapf(c4eerrors.ErrNotExists, "vesting pools not found for address %s", owner)
	}
	if err := vestingPool.SubstractFromReservation(reservationId, amout); err != nil {
		return err
	}

	k.SetAccountVestingPools(ctx, accountVestingPools)
	return nil
}
