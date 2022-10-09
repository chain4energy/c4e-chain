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
	k.Logger(ctx).Debug("add vesting pool", "vestingPoolName", vestingPoolName, "vestingAddr: ", vestingAddr, "coinSrcAddr", coinSrcAddr,
		"amount", amount, "vestingType", vestingType, "lockStart", lockStart, "lockEnd", lockEnd)

	if amount.Equal(sdk.ZeroInt()) {
		k.Logger(ctx).Error("add vesting pool amount equals zero error")
		return nil
	}

	_, err := sdk.AccAddressFromBech32(vestingAddr)
	if err != nil {
		k.Logger(ctx).Error("add vesting pool vestingAccAddress parsing error", "error", err.Error())
		return err
	}
	denom := k.GetParams(ctx).Denom

	var srcAccAddress sdk.AccAddress
	srcAccAddress, err = sdk.AccAddressFromBech32(coinSrcAddr)
	if err != nil {
		k.Logger(ctx).Error("add vesting pool srcAccAddress parsing error", "error", err.Error())
		return err
	}

	balance := k.bank.GetBalance(ctx, srcAccAddress, denom)

	if balance.Amount.LT(amount) {
		k.Logger(ctx).Error("add vesting pool balance lesser than requested amount error", "balanceAmount",
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
		k.Logger(ctx).Debug("account vesting pools", "address", accVestings.Address)
		id = 1
	} else {
		id = int32(len(accVestings.VestingPools)) + 1

		for _, pool := range accVestings.VestingPools {
			if pool.Name == vestingPoolName {
				k.Logger(ctx).Error("add vesting pool vesting pool name already exists error", "vestingPoolName", vestingPoolName)
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
	k.Logger(ctx).Debug("coins to send", "amount", coinToSend)

	if err != nil {
		k.Logger(ctx).Error("add vesting pool sendig coins to vesting pool error", "error", err.Error())
		return err
	}
	k.SetAccountVestings(ctx, accVestings)
	k.Logger(ctx).Debug("vesting pool added", "address", accVestings.Address,
		"id", vestingPool.Id,
		"name", vestingPool.Name,
		"vestingType", vestingPool.VestingType,
		"lockStart", vestingPool.LockStart,
		"lockEnd", vestingPool.LockEnd,
		"vested", vestingPool.Vested,
		"withdrawn", vestingPool.Withdrawn,
		"sent", vestingPool.Sent,
		"lastModification", vestingPool.LastModification,
		"lastModificationVested", vestingPool.LastModificationVested,
		"lastModificationWithdrawn", vestingPool.LastModificationWithdrawn,
	)

	return nil
}

func (k Keeper) WithdrawAllAvailable(ctx sdk.Context, address string) (withdrawn sdk.Coin, returnedError error) {
	k.Logger(ctx).Debug("withdraw all available", "address", address)

	accAddress, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		k.Logger(ctx).Error("withdraw all available parsing error", "error", err.Error())
		return withdrawn, err
	}

	accVestings, vestingsFound := k.GetAccountVestings(ctx, address)
	if !vestingsFound {
		k.Logger(ctx).Error("withdraw all available no vestings found error")
		return withdrawn, status.Error(codes.NotFound, "No vestings")
	}

	if len(accVestings.VestingPools) == 0 {
		k.Logger(ctx).Error("withdraw all available no vestings in array error")
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
		k.Logger(ctx).Debug("withdrawable data", "withdrawable", withdrawable, "LastModificationWithdrawn", vesting.LastModificationWithdrawn,
			"toWithdraw", toWithdraw)
		if toWithdraw.IsPositive() {
			events = append(events, types.WithdrawAvailable{
				OwnerAddress:    address,
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
		k.Logger(ctx).Debug("send coins from module to account", "accAddress", accAddress, "coinsToSend", coinsToSend)
		if err != nil {
			k.Logger(ctx).Error("withdraw all available sending coins to vesting account error", "error", err.Error())
			return withdrawn, err
		}
	}

	k.SetAccountVestings(ctx, accVestings)
	k.Logger(ctx).Debug("set account vestings", "accAddress", accVestings.Address, "newVestingPools", accVestings.VestingPools)
	if toWithdraw.IsPositive() && toWithdraw.IsInt64() {
		defer func() {
			telemetry.IncrCounter(1, types.ModuleName, "withdraw_available")
			telemetry.SetGaugeWithLabels(
				[]string{"tx", "msg", types.ModuleName, "withdraw_available"},
				float32(withdrawn.Amount.Int64()),
				[]metrics.Label{telemetry.NewLabel("denom", withdrawn.Denom)},
			)
		}()
	}

	for _, event := range events {
		ctx.EventManager().EmitTypedEvent(&event)
	}
	k.Logger(ctx).Debug("withdraw all available ret", "denom", denom, "toWithdraw", toWithdraw)
	return sdk.NewCoin(denom, toWithdraw), nil
}

func (k Keeper) SendToNewVestingAccount(ctx sdk.Context, fromAddr string, toAddr string, vestingId int32, amount sdk.Int, restartVesting bool) (withdrawn sdk.Coin, returnedError error) {
	k.Logger(ctx).Debug("send to new vesting account", "fromAddr", fromAddr, "toAddr", toAddr, "vestingId", vestingId, "amount", amount, "restartVesting", restartVesting)
	if fromAddr == toAddr {
		k.Logger(ctx).Error("send to new vesting account from address and to address cannot be identical error", "address", fromAddr)
		return withdrawn, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "from address and to address cannot be identical")
	}

	w, err := k.WithdrawAllAvailable(ctx, fromAddr)
	if err != nil {
		k.Logger(ctx).Error("send to new vesting account withdraw all available error", "error", err.Error())
		return withdrawn, err
	}

	accVestings, vestingsFound := k.GetAccountVestings(ctx, fromAddr)
	if !vestingsFound || len(accVestings.VestingPools) == 0 {
		k.Logger(ctx).Error("send to new vesting no vesting pools found")
		return withdrawn, sdkerrors.Wrap(sdkerrors.ErrNotFound, "no vesting pools found")
	}
	var vesting *types.VestingPool = nil
	for _, vest := range accVestings.VestingPools {
		if vest.Id == vestingId {
			vesting = vest
		}
	}
	if vesting == nil {
		k.Logger(ctx).Error("send to new vesting vesting pool id not found", "vestingId", strconv.FormatInt(int64(vestingId), 10))
		return withdrawn, sdkerrors.Wrap(sdkerrors.ErrNotFound, "vesting pool with id "+strconv.FormatInt(int64(vestingId), 10)+" not found")
	}
	available := vesting.LastModificationVested.Sub(vesting.LastModificationWithdrawn)
	if available.LT(amount) {
		k.Logger(ctx).Error("vesting available is smaller than amount", "available", available, "amount", amount)
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
			k.Logger(ctx).Error("send to new vesting account get vesting type error", "error", err.Error())
			return withdrawn, sdkerrors.Wrap(sdkerrors.ErrNotFound, err.Error())
		}
		err = k.createVestingAccount(ctx, toAddr, amount,
			ctx.BlockTime().Add(vt.LockupPeriod), ctx.BlockTime().Add(vt.LockupPeriod).Add(vt.VestingPeriod))
	} else {
		err = k.createVestingAccount(ctx, toAddr, amount,
			vesting.LockEnd, vesting.LockEnd)
	}
	if err == nil {
		k.SetAccountVestings(ctx, accVestings) //TODO: why?
		k.Logger(ctx).Debug("set account vestings", "accVestingsAddress", accVestings.Address)
		k.AppendVestingAccount(ctx, types.VestingAccount{Address: toAddr})
		k.Logger(ctx).Debug("append vesting account", "address", toAddr)

		ctx.EventManager().EmitTypedEvent(&types.NewVestingAccountFromVestingPool{
			OwnerAddress:    fromAddr,
			Address:         toAddr,
			VestingPoolId:   strconv.FormatInt(int64(vesting.Id), 10),
			VestingPoolName: vesting.Name,
			Amount:          amount.String() + k.Denom(ctx),
			RestartVesting:  strconv.FormatBool(restartVesting),
		})
	}
	k.Logger(ctx).Debug("send to new vesting account ret", "withdrawn", w, "error", err)
	return w, err
}

func (k Keeper) CreateVestingAccount(ctx sdk.Context, fromAddress string, toAddress string,
	amount sdk.Coins, startTime int64, endTime int64) error {
	k.Logger(ctx).Debug("create vesting account", "fromAddress", fromAddress, "toAddress", toAddress,
		"amount", amount, "startTime", startTime, "endTime", endTime)
	ak := k.account
	bk := k.bank

	if err := bk.IsSendEnabledCoins(ctx, amount...); err != nil {
		k.Logger(ctx).Error("create vesting account is send enabled coins error", "error", err.Error())
		return err
	}

	from, err := sdk.AccAddressFromBech32(fromAddress)
	if err != nil {
		k.Logger(ctx).Error("create vesting account parsing error", "error", err.Error())
		return err
	}
	to, err := sdk.AccAddressFromBech32(toAddress)
	if err != nil {
		k.Logger(ctx).Error("create vesting account parsing error", "error", err.Error())
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
	ctx.EventManager().EmitTypedEvent(&types.NewVestingAccount{
		Address: acc.Address,
	})

	err = bk.SendCoins(ctx, from, to, amount)
	k.Logger(ctx).Debug("send coins", "addressFrom", from, "addressTo", to, "amount", amount)

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
	k.Logger(ctx).Debug("create vesting account", "toAddress", toAddress, "amount", amount, "lockEnd", lockEnd,
		"vestingEnd", vestingEnd)

	ak := k.account
	bk := k.bank

	denom := k.GetParams(ctx).Denom
	coinToSend := sdk.NewCoin(denom, amount)
	k.Logger(ctx).Debug("create vesting account data", "denom", denom, "coinToSend", coinToSend.String())
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
		k.Logger(ctx).Error("create vesting account is not allowed to receive funds error", "address", to)
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to receive funds", toAddress)
	}

	if acc := ak.GetAccount(ctx, to); acc != nil {
		k.Logger(ctx).Error("create vesting account account already exists error", "toAddress", toAddress)
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "account %s already exists", toAddress)
	}

	baseAccount := ak.NewAccountWithAddress(ctx, to)
	if _, ok := baseAccount.(*authtypes.BaseAccount); !ok {
		k.Logger(ctx).Error("create vesting account invalid account type; expected: BaseAccoun", "notExpectedAccount", baseAccount)
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
	k.Logger(ctx).Debug("create vesting account", "baseAccount", baseVestingAccount.BaseAccount, "originalVesting",
		baseVestingAccount.OriginalVesting, "delegatedFree", baseVestingAccount.DelegatedFree, "delegatedVesting",
		baseVestingAccount.DelegatedVesting, "endTime", baseVestingAccount.EndTime, "startTime", startTime.Unix())

	ctx.EventManager().EmitTypedEvent(&types.NewVestingAccount{
		Address: acc.GetAddress().String(),
	})
	err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, coinsToSend)
	k.Logger(ctx).Debug("send coins from module to account", "moduleName", types.ModuleName, "addressTo", to,
		"coinsToSend", coinsToSend)

	if err != nil {
		k.Logger(ctx).Error("create vesting account send coins from module to account error", "error", err.Error())
		return err
	}

	return nil
}
