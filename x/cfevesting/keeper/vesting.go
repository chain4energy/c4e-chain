package keeper

import (
	"math"
	"strconv"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// vest

func (k Keeper) Vest(ctx sdk.Context, addr string, amount sdk.Int, vestingType string) error {
	k.Logger(ctx).Debug("Vest: addr: " + addr + "amount: " + amount.String() + "vestingType: " + vestingType)
	vt, err := k.GetVestingType(ctx, vestingType)
	if err != nil {
		k.Logger(ctx).Error("Error: " + err.Error())

		return sdkerrors.Wrap(sdkerrors.ErrNotFound, err.Error())
	}
	k.Logger(ctx).Debug("vt: DelegationsAllowed: " + strconv.FormatBool(vt.DelegationsAllowed))

	// return k.addVesting(ctx, false, addr, addr, amount, vestingType, vt.DelegationsAllowed, ctx.BlockHeight(),
	// 	vt.LockupPeriod.Nanoseconds()+ctx.BlockHeight(), vt.LockupPeriod.Nanoseconds()+vt.VestingPeriod.Nanoseconds()+ctx.BlockHeight(),
	// 	vt.TokenReleasingPeriod.Nanoseconds())

	return k.addVesting(ctx, false, addr, addr, amount, vestingType, vt.DelegationsAllowed, ctx.BlockTime(),
		ctx.BlockTime().Add(vt.LockupPeriod), ctx.BlockTime().Add(vt.LockupPeriod).Add(vt.VestingPeriod),
		vt.TokenReleasingPeriod)
		
}

func (k Keeper) addVesting(
	ctx sdk.Context,
	isModuleInternalOperation bool,
	vestingAddr string,
	coinSrcAddr string,
	amount sdk.Int,
	vestingType string,
	delegationAllowed bool,
	vestingStart time.Time,
	lockEnd time.Time,
	vestingEnd time.Time,
	releasePeriod time.Duration) error {

	if amount.Equal(sdk.ZeroInt()) {
		return nil
	}

	vestingAccAddress, err := sdk.AccAddressFromBech32(vestingAddr)
	if err != nil {
		k.Logger(ctx).Error("Error: " + err.Error())
		return err
	}

	denom := k.GetParams(ctx).Denom
	k.Logger(ctx).Debug("denom: " + denom)

	var srcAccAddress sdk.AccAddress
	if !isModuleInternalOperation || delegationAllowed {
		srcAccAddress, err = sdk.AccAddressFromBech32(coinSrcAddr)
		if err != nil {
			k.Logger(ctx).Error("Error: " + err.Error())
			return err
		}
		k.Logger(ctx).Debug("vestingAccAddress: " + vestingAccAddress.String())

		balance := k.bank.GetBalance(ctx, srcAccAddress, denom)
		k.Logger(ctx).Debug("balance: " + balance.Amount.String())

		if balance.Amount.LT(amount) {
			k.Logger(ctx).Error("Error: " + "Balance [" + balance.Amount.String() +
				balance.Denom + "]lesser than requested amount: " + amount.String() + denom)
			return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Balance ["+balance.Amount.String()+
				balance.Denom+"]lesser than requested amount: "+amount.String()+denom)
		}
	}

	accVestings, vestingsFound := k.GetAccountVestings(ctx, vestingAddr)
	k.Logger(ctx).Debug("vestingsFound: " + strconv.FormatBool(vestingsFound))
	var id int32
	var delegatableAddress sdk.AccAddress
	if !vestingsFound {
		accVestings = types.AccountVestings{}
		accVestings.Address = vestingAddr
		k.Logger(ctx).Debug("accVestings.Address: " + accVestings.Address)

		if delegationAllowed {
			delegatableAddress = k.CreateModuleAccountForDelegatableVesting(ctx, vestingAddr).GetAddress()
			k.Logger(ctx).Debug("delegatableAddress: " + delegatableAddress.String())

			accVestings.DelegableAddress = delegatableAddress.String()
		}
		id = 1
	} else {
		if delegationAllowed {
			if (accVestings.DelegableAddress == "") {
				delegatableAddress = k.CreateModuleAccountForDelegatableVesting(ctx, vestingAddr).GetAddress()
				k.Logger(ctx).Debug("delegatableAddress: " + delegatableAddress.String())
				accVestings.DelegableAddress = delegatableAddress.String()
			} else {
				delegatableAddress, err = sdk.AccAddressFromBech32(accVestings.DelegableAddress)
				k.Logger(ctx).Debug("delegatableAddress: " + delegatableAddress.String())
				if err != nil {
					k.Logger(ctx).Error("Error: " + err.Error())
					return err
				}
			}
		}
		id = int32(len(accVestings.Vestings)) + 1
	}

	vesting := types.Vesting{
		Id:                id,
		VestingType:       vestingType,
		VestingStart: vestingStart,
		LockEnd:      lockEnd,
		VestingEnd:   vestingEnd,
		Vested:            amount,
		ReleasePeriod: releasePeriod,
		DelegationAllowed:         delegationAllowed,
		Withdrawn:                 sdk.ZeroInt(),
		Sent:                      sdk.ZeroInt(),
		LastModification:     vestingStart,
		LastModificationVested:    amount,
		LastModificationWithdrawn: sdk.ZeroInt(),
	}
	accVestings.Vestings = append(accVestings.Vestings, &vesting)
	

	coinToSend := sdk.NewCoin(denom, amount)
	coinsToSend := sdk.NewCoins(coinToSend)

	if delegationAllowed {

		err = k.bank.SendCoins(ctx, srcAccAddress, delegatableAddress, coinsToSend)
		k.Logger(ctx).Info("after SendCoins: " + coinToSend.Amount.String())

	} else if !isModuleInternalOperation {
		err = k.bank.SendCoinsFromAccountToModule(ctx, srcAccAddress, types.ModuleName, coinsToSend)
		k.Logger(ctx).Info("after SendCoinsFromAccountToModule: " + coinToSend.Amount.String())

	}
	if err != nil {
		k.Logger(ctx).Error("Error: " + err.Error())
		return err
	}
	k.SetAccountVestings(ctx, accVestings)
	k.Logger(ctx).Info("Vest exit: " + vestingAddr)
	return nil
}

func (k Keeper) SendVesting(ctx sdk.Context, fromAddr string, toAddr string, vestingId int32, amount sdk.Int, restartVesting bool) (withdrawn sdk.Coin, returnedError error) {
	if fromAddr == toAddr {
		return withdrawn, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "from address and to address cannot be identical")
	}

	w, err := k.WithdrawAllAvailable(ctx, fromAddr)
	if err != nil {
		return withdrawn, err
	}

	accVestings, vestingsFound := k.GetAccountVestings(ctx, fromAddr)
	if !vestingsFound || len(accVestings.Vestings) == 0 {
		return withdrawn, sdkerrors.Wrap(sdkerrors.ErrNotFound, "no vestings found")
	}
	var vesting *types.Vesting = nil
	for _, vest := range accVestings.Vestings {
		if vest.Id == vestingId {
			vesting = vest
		}
	}
	if vesting == nil {
		return withdrawn, sdkerrors.Wrap(sdkerrors.ErrNotFound, "vesting with id "+strconv.FormatInt(int64(vestingId), 10)+" not found")
	}
	if !vesting.TransferAllowed {
		return withdrawn, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "vesting with id "+strconv.FormatInt(int64(vestingId), 10)+" is not tranferable")
	}
	available := vesting.LastModificationVested.Sub(vesting.LastModificationWithdrawn)
	if available.LT(amount) {
		return withdrawn, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds,
			"vesting available: %s is smaller than %s", available, amount)
	}
	if vesting.DelegationAllowed {
		if len(accVestings.DelegableAddress) == 0 {
			return withdrawn, sdkerrors.Wrap(sdkerrors.ErrLogic, "delegable vesting has no delegable address")
		}
		denom := k.GetParams(ctx).Denom
		from, err := sdk.AccAddressFromBech32(accVestings.DelegableAddress)
		if err != nil {
			return withdrawn, err
		}
		balance := k.bank.GetBalance(ctx, from, denom)
		lockedCoins := k.bank.LockedCoins(ctx, from)
		locked := sdk.NewCoin(denom, lockedCoins.AmountOf(denom))
		spendable := balance.Sub(locked)
		if spendable.Amount.LT(amount) {
			return withdrawn, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds,
				"vesting available: %s is smaller than %s - probably delageted to validator.", spendable.Amount, amount)
		}
	}
	vesting.Sent = amount
	vesting.LastModification = ctx.BlockTime()
	vesting.LastModificationVested = available.Sub(amount)
	vesting.LastModificationWithdrawn = sdk.ZeroInt()

	if restartVesting {
		vt, vErr := k.GetVestingType(ctx, vesting.VestingType)
		if vErr != nil {
			k.Logger(ctx).Error("Error: " + err.Error())

			return withdrawn, sdkerrors.Wrap(sdkerrors.ErrNotFound, err.Error())
		}
		k.Logger(ctx).Debug("vt: DelegationsAllowed: " + strconv.FormatBool(vt.DelegationsAllowed))
		// return w, k.addVesting(ctx, true, toAddr, accVestings.DelegableAddress, amount, vesting.VestingType, vt.DelegationsAllowed, ctx.BlockHeight(),
		// 	vt.LockupPeriod.Nanoseconds()+ctx.BlockHeight(), vt.LockupPeriod.Nanoseconds()+vt.VestingPeriod.Nanoseconds()+ctx.BlockHeight(),
		// 	vt.TokenReleasingPeriod.Nanoseconds())
		err = k.addVesting(ctx, true, toAddr, accVestings.DelegableAddress, amount, vesting.VestingType, vt.DelegationsAllowed, ctx.BlockTime(),
			ctx.BlockTime().Add(vt.LockupPeriod), ctx.BlockTime().Add(vt.LockupPeriod).Add(vt.VestingPeriod),
			vt.TokenReleasingPeriod) 
	} else {
		err = k.addVesting(ctx, true, toAddr, accVestings.DelegableAddress, amount, vesting.VestingType, vesting.DelegationAllowed, ctx.BlockTime(),
			vesting.LockEnd, vesting.VestingEnd,
			vesting.ReleasePeriod)
	}
	if err == nil {
		k.SetAccountVestings(ctx, accVestings)
	}
	return w, err
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

	if len(accVestings.Vestings) == 0 {
		return withdrawn, status.Error(codes.NotFound, "No vestings")
	}

	current := ctx.BlockTime()
	toWithdraw := sdk.ZeroInt()
	toWithdrawDelegable := sdk.ZeroInt()
	for _, vesting := range accVestings.Vestings {
		withdrawable := CalculateWithdrawable(current, *vesting)
		vesting.Withdrawn = vesting.Withdrawn.Add(withdrawable)
		vesting.LastModificationWithdrawn = vesting.LastModificationWithdrawn.Add(withdrawable)
		if vesting.DelegationAllowed {
			toWithdrawDelegable = toWithdrawDelegable.Add(withdrawable)
		} else {
			toWithdraw = toWithdraw.Add(withdrawable)
		}
	}

	k.SetAccountVestings(ctx, accVestings)

	denom := k.GetParams(ctx).Denom
	if toWithdraw.GT(sdk.ZeroInt()) {
		coinToSend := sdk.NewCoin(denom, toWithdraw)
		coinsToSend := sdk.NewCoins(coinToSend)
		err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, accAddress, coinsToSend)
		if err != nil {
			return withdrawn, err
		}
	}
	if toWithdrawDelegable.GT(sdk.ZeroInt()) {
		coinToSend := sdk.NewCoin(denom, toWithdrawDelegable)
		coinsToSend := sdk.NewCoins(coinToSend)
		delegatableAddress, err := sdk.AccAddressFromBech32(accVestings.DelegableAddress)
		if err != nil {
			return withdrawn, err
		}
		err = k.bank.SendCoins(ctx, delegatableAddress, accAddress, coinsToSend)
		if err != nil {
			return withdrawn, err
		}
	}

	return sdk.NewCoin(denom, toWithdraw.Add(toWithdrawDelegable)), nil
}

func CalculateWithdrawable(current time.Time, vesting types.Vesting) sdk.Int {
	if current.Equal(vesting.VestingStart) || current.Before(vesting.VestingStart) {
		return sdk.ZeroInt()
	}
	if current.Equal(vesting.LockEnd) || current.Before(vesting.LockEnd) {
		return sdk.ZeroInt()
	}
	if current.Equal(vesting.VestingEnd) || current.After(vesting.VestingEnd) {
		return vesting.LastModificationVested.Sub(vesting.LastModificationWithdrawn)
	}

	var lockEnd time.Time
	if vesting.VestingStart.After(vesting.LockEnd) {
		lockEnd = vesting.VestingStart
	} else {
		lockEnd = vesting.LockEnd
	}
	if vesting.GetLastModification().After(lockEnd) {
		lockEnd = vesting.LastModification
	}
	wholeVestingPariod := vesting.VestingEnd.Sub(lockEnd).Nanoseconds()
	fromStart := current.Sub(lockEnd).Nanoseconds()
	numOfPeriodsFromStart := fromStart / vesting.ReleasePeriod.Nanoseconds()
	numOfPeriods := int64(math.Ceil(float64(wholeVestingPariod) / float64(vesting.ReleasePeriod)))

	vested := vesting.LastModificationVested

	withdrawableFromStart := vested.MulRaw(numOfPeriodsFromStart).QuoRaw(numOfPeriods)
	return withdrawableFromStart.Sub(vesting.LastModificationWithdrawn)

}

func (k Keeper) CreateModuleAccountForDelegatableVesting(ctx sdk.Context, address string) authtypes.ModuleAccountI {
	perms := []string{}
	macc := authtypes.NewEmptyModuleAccount(createAccounNameForDelegatableVesting(address), perms...)
	maccI := (k.account.NewAccount(ctx, macc)).(authtypes.ModuleAccountI) // set the account number
	k.account.SetModuleAccount(ctx, maccI)
	return maccI
}

func (k Keeper) GetModuleAccountForDelegableVesting(ctx sdk.Context, address sdk.AccAddress) authtypes.ModuleAccountI {
	maccI := k.account.GetAccount(ctx, address).(authtypes.ModuleAccountI)
	return maccI
}

func createAccounNameForDelegatableVesting(address string) string {
	return types.ModuleName + address
}
