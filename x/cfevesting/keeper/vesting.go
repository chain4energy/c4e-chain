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
	k.Logger(ctx).Debug("create vesting pool", "addr", addr, "amount: ", amount, "vestingType", vestingType)
	_, err := k.GetVestingType(ctx, vestingType)
	if err != nil {
		k.Logger(ctx).Error("create vesting pool get vesting type error", "error", err.Error())
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

	if amount.LTE(sdk.ZeroInt()) {
		k.Logger(ctx).Error("add vesting pool amount <= 0", "vestingPoolName", vestingPoolName, "vestingAddr: ", vestingAddr, "coinSrcAddr", coinSrcAddr, "amount", amount)
		return nil // TODO return error
	}

	_, err := sdk.AccAddressFromBech32(vestingAddr)
	if err != nil {
		k.Logger(ctx).Error("add vesting pool vesting acc address parsing error", 
			"vestingPoolName", vestingPoolName, "vestingAddr: ", vestingAddr, "coinSrcAddr", coinSrcAddr, "error", err.Error())
		return err
	}
	denom := k.GetParams(ctx).Denom

	var srcAccAddress sdk.AccAddress
	srcAccAddress, err = sdk.AccAddressFromBech32(coinSrcAddr)
	if err != nil {
		k.Logger(ctx).Error("add vesting pool source account address parsing error", 
			"vestingPoolName", vestingPoolName, "vestingAddr: ", vestingAddr, "coinSrcAddr", coinSrcAddr, "error", err.Error())
		return err
	}

	balance := k.bank.GetBalance(ctx, srcAccAddress, denom)

	if balance.Amount.LT(amount) {
		k.Logger(ctx).Error("add vesting pool balance lesser than requested amount error", 
			"vestingPoolName", vestingPoolName, "vestingAddr: ", vestingAddr, "coinSrcAddr", coinSrcAddr, "balanceAmount",
			balance.Amount.String(), "balanceDenom", balance.Denom, "reguestedAmount", amount.String()+denom)
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "balance ["+balance.Amount.String()+
			balance.Denom+"] lesser than requested amount: "+amount.String()+denom)
	}

	accVestings, vestingsFound := k.GetAccountVestings(ctx, vestingAddr)
	k.Logger(ctx).Debug("data", "denom", denom, "balance", balance, "vestingsFound", vestingsFound)
	var id int32
	if !vestingsFound {
		accVestings = types.AccountVestings{}
		accVestings.Address = vestingAddr
		id = 1
	} else {
		id = int32(len(accVestings.VestingPools)) + 1

		for _, pool := range accVestings.VestingPools {
			if pool.Name == vestingPoolName {
				k.Logger(ctx).Error("add vesting pool vesting pool name already exists error",
					"vestingPoolName", vestingPoolName, "vestingAddr: ", vestingAddr, "coinSrcAddr")
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "vesting pool name already exists: "+vestingPoolName)
			}
		}
	}

	vestingPool := types.VestingPool{
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
	accVestings.VestingPools = append(accVestings.VestingPools, &vestingPool)

	coinToSend := sdk.NewCoin(denom, amount)
	coinsToSend := sdk.NewCoins(coinToSend)
	err = k.bank.SendCoinsFromAccountToModule(ctx, srcAccAddress, types.ModuleName, coinsToSend)

	k.Logger(ctx).Debug("add vesting pool", "vestingPoolName", vestingPoolName, "vestingAddr: ", vestingAddr, "coinSrcAddr", coinSrcAddr,
		"amount", amount, "vestingType", vestingType, "lockStart", lockStart, "lockEnd", lockEnd,
		"denom", denom, "balance", balance, "vestingsFound", vestingsFound, "vestingPool", vestingPool)
	if err != nil {
		k.Logger(ctx).Error("add vesting pool sendig coins to vesting pool error", "error", err.Error())
		return err
	}
	k.SetAccountVestings(ctx, accVestings)
	return nil
}

func (k Keeper) WithdrawAllAvailable(ctx sdk.Context, address string) (withdrawn sdk.Coin, returnedError error) {
	accAddress, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		k.Logger(ctx).Error("withdraw all available address parsing error", "address", address, "error", err.Error())
		return withdrawn, err
	}

	accVestings, vestingsFound := k.GetAccountVestings(ctx, address)
	if !vestingsFound {
		k.Logger(ctx).Error("withdraw all available no vestings found error", "address", address)
		return withdrawn, status.Error(codes.NotFound, "No vestings")
	}

	if len(accVestings.VestingPools) == 0 {
		k.Logger(ctx).Error("withdraw all available no vestings in array error", "address", address)
		return withdrawn, status.Error(codes.NotFound, "No vestings")
	}

	current := ctx.BlockTime()
	toWithdraw := sdk.ZeroInt()
	events := make([]types.WithdrawAvailable, 0)
	denom := k.GetParams(ctx).Denom
	for _, vestingPool := range accVestings.VestingPools {
		withdrawable := CalculateWithdrawable(current, *vestingPool)
		vestingPool.Withdrawn = vestingPool.Withdrawn.Add(withdrawable)
		vestingPool.LastModificationWithdrawn = vestingPool.LastModificationWithdrawn.Add(withdrawable)
		toWithdraw = toWithdraw.Add(withdrawable)
		k.Logger(ctx).Debug("withdraw all available data", "address", address, "vestingPool", vestingPool, "withdrawable", withdrawable,
			"toWithdraw", toWithdraw)
		if toWithdraw.IsPositive() {
			events = append(events, types.WithdrawAvailable{
				OwnerAddress:    address,
				VestingPoolId:   strconv.FormatInt(int64(vestingPool.Id), 10),
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
			return withdrawn, err
		}
	}

	k.SetAccountVestings(ctx, accVestings)
	k.Logger(ctx).Debug("set account vesting pools", "accAddress", accVestings.Address, "newVestingPools", accVestings.VestingPools)
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

func (k Keeper) SendToNewVestingAccount(ctx sdk.Context, fromAddr string, toAddr string, vestingPoolId int32, amount sdk.Int, restartVesting bool) (withdrawn sdk.Coin, returnedError error) {
	k.Logger(ctx).Debug("send to new vesting account", "fromAddr", fromAddr, "toAddr", toAddr, "vestingPoolId", vestingPoolId, "amount", amount, "restartVesting", restartVesting)
	if fromAddr == toAddr {
		k.Logger(ctx).Error("send to new vesting account from address and to address cannot be identical error", "address", fromAddr)
		return withdrawn, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "from address and to address cannot be identical")
	}

	w, err := k.WithdrawAllAvailable(ctx, fromAddr)
	if err != nil {
		k.Logger(ctx).Error("send to new vesting account withdraw all available error", "fromAddr", fromAddr, "error", err.Error())
		return withdrawn, err
	}

	accVestings, vestingsFound := k.GetAccountVestings(ctx, fromAddr)
	if !vestingsFound || len(accVestings.VestingPools) == 0 {
		k.Logger(ctx).Error("send to new vesting no vesting pools found", "fromAddr", fromAddr)
		return withdrawn, sdkerrors.Wrap(sdkerrors.ErrNotFound, "no vesting pools found")
	}
	var vestingPool *types.VestingPool = nil
	for _, vest := range accVestings.VestingPools {
		if vest.Id == vestingPoolId {
			vestingPool = vest
		}
	}
	if vestingPool == nil {
		k.Logger(ctx).Error("send to new vesting vesting pool id not found", "fromAddr", fromAddr, "vestingPoolId", vestingPoolId)
		return withdrawn, sdkerrors.Wrap(sdkerrors.ErrNotFound, "vesting pool with id "+strconv.FormatInt(int64(vestingPoolId), 10)+" not found")
	}
	available := vestingPool.LastModificationVested.Sub(vestingPool.LastModificationWithdrawn)
	if available.LT(amount) {
		k.Logger(ctx).Error("vesting available is smaller than amount", "fromAddr", fromAddr, "vestingPoolId", vestingPoolId, "available", available, "amount", amount)
		return withdrawn, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds,
			"vesting available: %s is smaller than %s", available, amount)
	}
	vestingPool.Sent = amount
	vestingPool.LastModification = ctx.BlockTime()
	vestingPool.LastModificationVested = available.Sub(amount)
	vestingPool.LastModificationWithdrawn = sdk.ZeroInt()

	if restartVesting {
		vt, vErr := k.GetVestingType(ctx, vestingPool.VestingType)
		if vErr != nil {
			k.Logger(ctx).Error("send to new vesting account get vesting type error", "fromAddr", fromAddr, "vestingPool", vestingPool, "error", err.Error())
			return withdrawn, sdkerrors.Wrap(sdkerrors.ErrNotFound, err.Error())
		}
		err = k.createVestingAccount(ctx, toAddr, amount,
			ctx.BlockTime().Add(vt.LockupPeriod), ctx.BlockTime().Add(vt.LockupPeriod).Add(vt.VestingPeriod))
	} else {
		err = k.createVestingAccount(ctx, toAddr, amount,
			vestingPool.LockEnd, vestingPool.LockEnd)
	}
	if err == nil {
		k.SetAccountVestings(ctx, accVestings)
		k.AppendVestingAccount(ctx, types.VestingAccount{Address: toAddr})

		err := ctx.EventManager().EmitTypedEvent(&types.NewVestingAccountFromVestingPool{
			OwnerAddress:    fromAddr,
			Address:         toAddr,
			VestingPoolId:   strconv.FormatInt(int64(vestingPool.Id), 10),
			VestingPoolName: vestingPool.Name,
			Amount:          amount.String() + k.Denom(ctx),
			RestartVesting:  strconv.FormatBool(restartVesting),
		})
		if err != nil {
			k.Logger(ctx).Error("new vesting account from vesting pool emit event error", "error", err.Error())
		}
	}
	k.Logger(ctx).Debug("send to new vesting account ret", "withdrawn", w, "error", err, "accVestings", accVestings)
	return w, err
}

func (k Keeper) CreateVestingAccount(ctx sdk.Context, fromAddress string, toAddress string,
	amount sdk.Coins, startTime int64, endTime int64) error {
	k.Logger(ctx).Debug("create vesting account", "fromAddress", fromAddress, "fromAddress", fromAddress,
		"amount", amount, "startTime", startTime, "endTime", endTime)
	ak := k.account
	bk := k.bank

	if err := bk.IsSendEnabledCoins(ctx, amount...); err != nil {
		k.Logger(ctx).Error("create vesting account send coins disabled", "error", err.Error())
		return err
	}

	from, err := sdk.AccAddressFromBech32(fromAddress)
	if err != nil {
		k.Logger(ctx).Error("create vesting account from-address parsing error", "fromAddress", fromAddress, "error", err.Error())
		return err
	}
	to, err := sdk.AccAddressFromBech32(toAddress)
	if err != nil {
		k.Logger(ctx).Error("create vesting account to-address parsing error", "fromAddress", fromAddress, "error", err.Error())
		return err
	}

	if bk.BlockedAddr(to) {
		k.Logger(ctx).Error("create vesting account account is not allowed to receive funds error", "toAddress", toAddress)
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to receive funds", toAddress)
	}

	if acc := ak.GetAccount(ctx, to); acc != nil {
		k.Logger(ctx).Error("create vesting account account already exists error", "toAddress", toAddress)
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "account %s already exists", toAddress)
	}

	baseAccount := ak.NewAccountWithAddress(ctx, to)
	if _, ok := baseAccount.(*authtypes.BaseAccount); !ok {
		k.Logger(ctx).Error("create vesting account invalid account type; expected: BaseAccount", "notExpectedAccount", baseAccount)
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid account type; expected: BaseAccount, got: %T", baseAccount)
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
		k.Logger(ctx).Debug("create vesting account send coins from module to account error", "error", err.Error())
		return err
	}

	k.AppendVestingAccount(ctx, types.VestingAccount{Address: acc.Address})
	k.Logger(ctx).Debug("append vesting account", "address", acc.Address)
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
	denom := k.GetParams(ctx).Denom
	k.Logger(ctx).Debug("create vesting account", "toAddress", toAddress, "amount", amount, "lockEnd", lockEnd,
		"vestingEnd", vestingEnd, "denom", denom)

	ak := k.account
	bk := k.bank
	coinToSend := sdk.NewCoin(denom, amount)
	if err := bk.IsSendEnabledCoins(ctx, coinToSend); err != nil {
		k.Logger(ctx).Error("create vesting account is send enabled coins error", "error", err.Error())
		return err
	}

	to, err := sdk.AccAddressFromBech32(toAddress)
	if err != nil {
		k.Logger(ctx).Error("create vesting account parsing error", "error", err.Error())
		return err
	}

	if bk.BlockedAddr(to) {
		k.Logger(ctx).Error("create vesting account is not allowed to receive funds error", "address", toAddress)
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to receive funds", toAddress)
	}

	if acc := ak.GetAccount(ctx, to); acc != nil {
		k.Logger(ctx).Error("create vesting account account already exists error", "toAddress", toAddress)
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "account %s already exists", toAddress)
	}

	baseAccount := ak.NewAccountWithAddress(ctx, to)
	if _, ok := baseAccount.(*authtypes.BaseAccount); !ok {
		k.Logger(ctx).Error("create vesting account invalid account type; expected: BaseAccount", "toAddress", toAddress, "notExpectedAccount", baseAccount)
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
	k.Logger(ctx).Debug("create vesting account", "baseAccount", baseVestingAccount.BaseAccount, "baseVestingAccount",
		baseVestingAccount, "startTime", startTime.Unix())

	err = ctx.EventManager().EmitTypedEvent(&types.NewVestingAccount{
		Address: acc.GetAddress().String(),
	})
	if err != nil {
		k.Logger(ctx).Error("create vesting account emit event error", "error", err.Error())
	}
	err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, coinsToSend)

	if err != nil {
		k.Logger(ctx).Error("create vesting account send coins to vesting account error", "error", err.Error())
		return err
	}

	return nil
}
