package keeper

import (
	"strconv"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// get the vesting types
func (k Keeper) Vest(ctx sdk.Context, addr string, amount uint64, vestingType string) error {
	k.Logger(ctx).Debug("Vest: addr: " + addr + "amount: " + strconv.FormatUint(amount, 10) + "vestingType: " + vestingType)
	vt, err := k.GetVestingType(ctx, vestingType)
	k.Logger(ctx).Debug("vt: DelegationsAllowed: " + strconv.FormatBool(vt.DelegationsAllowed))

	if err != nil {
		k.Logger(ctx).Error("Error: " + err.Error())

		return sdkerrors.Wrap(sdkerrors.ErrNotFound, err.Error())
	}
	if amount == 0 {
		return nil
	}

	accAddress, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		k.Logger(ctx).Error("Error: " + err.Error())
		return err
	}
	k.Logger(ctx).Debug("accAddress: " + accAddress.String())

	denom := k.GetParams(ctx).Denom
	k.Logger(ctx).Debug("denom: " + denom)

	balance := k.bank.GetBalance(ctx, accAddress, denom)
	k.Logger(ctx).Debug("balance: " + balance.Amount.String())

	if balance.Amount.LT(sdk.NewIntFromUint64(amount)) {
		k.Logger(ctx).Error("Error: " + "Balance ["+balance.Amount.String()+balance.Denom+"]lesser than requested amount: "+strconv.FormatUint(amount, 10)+denom)
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Balance ["+balance.Amount.String()+balance.Denom+"]lesser than requested amount: "+strconv.FormatUint(amount, 10)+denom)
	}

	// return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Balance: " + balance.Amount.String())

	accVestings, vestingsFound := k.GetAccountVestings(ctx, addr)
	k.Logger(ctx).Debug("vestingsFound: " + strconv.FormatBool(vestingsFound))
	var id int32
	var delegatableAddress sdk.AccAddress
	if !vestingsFound {
		accVestings = types.AccountVestings{}
		accVestings.Address = addr
		k.Logger(ctx).Debug("accVestings.Address: " + accVestings.Address)

		if vt.DelegationsAllowed {
			delegatableAddress = k.CreateModuleAccountForDelegatableVesting(ctx, addr).GetAddress()
			k.Logger(ctx).Debug("delegatableAddress: " + delegatableAddress.String())

			accVestings.DelegableAddress = delegatableAddress.String()
		}
		id = 1
	} else {
		if vt.DelegationsAllowed {
			delegatableAddress, err = sdk.AccAddressFromBech32(accVestings.DelegableAddress)
			k.Logger(ctx).Debug("delegatableAddress: " + delegatableAddress.String())
			if err != nil {
				k.Logger(ctx).Error("Error: " + err.Error())
				return err
			}
		}
		id = int32(len(accVestings.Vestings)) + 1
	}
	// numOfPeriods := vt.VestingPeriod / vt.TokenReleasingPeriod
	// amountPerPeriod := amount / uint64(numOfPeriods)
	vesting := types.Vesting{VestingType: vestingType,
		Id: id,
		VestingStartBlock:    ctx.BlockHeight(),
		LockEndBlock:         vt.LockupPeriod + ctx.BlockHeight(),
		VestingEndBlock:      vt.LockupPeriod + vt.VestingPeriod + ctx.BlockHeight(),
		Vested:               amount,
		// Claimable:            0,
		// LastFreeingBlock:     0,
		FreeCoinsBlockPeriod: vt.TokenReleasingPeriod,
		// FreeCoinsPerPeriod:   amountPerPeriod,
		DelegationAllowed:    vt.DelegationsAllowed,
		Withdrawn:            0}
	accVestings.Vestings = append(accVestings.Vestings, &vesting)
	k.SetAccountVestings(ctx, accVestings)

	coinToSend := sdk.NewCoin(denom, sdk.NewIntFromUint64(amount))
	coinsToSend := sdk.NewCoins(coinToSend)

	if vt.DelegationsAllowed {
		err = k.bank.SendCoins(ctx, accAddress, delegatableAddress, coinsToSend)
		k.Logger(ctx).Info("after SendCoins: " + coinToSend.Amount.String())

	} else {
		err = k.bank.SendCoinsFromAccountToModule(ctx, accAddress, types.ModuleName, coinsToSend)
		k.Logger(ctx).Info("after SendCoinsFromAccountToModule: " + coinToSend.Amount.String())

	}
	if err != nil {
		k.Logger(ctx).Error("Error: " + err.Error())
		return err
	}
	k.Logger(ctx).Info("Vest exit: " + addr)
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

	if len(accVestings.Vestings) == 0 {
		return withdrawn, status.Error(codes.NotFound, "No vestings")
	}

	height := ctx.BlockHeight()
	toWithdraw := sdk.ZeroInt()
	toWithdrawDelegable := sdk.ZeroInt()
	for _, vesting := range accVestings.Vestings {
		withdrawable := CalculateWithdrawable(height, *vesting)
		vesting.Withdrawn += withdrawable.Uint64()
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

func CalculateWithdrawable(height int64, vesting types.Vesting) sdk.Int {
	if height <= vesting.LockEndBlock {
		return sdk.ZeroInt()
	}
	if height >= vesting.VestingEndBlock {
		return sdk.NewIntFromUint64(vesting.Vested)
	}

	allVestingBlocks := vesting.VestingEndBlock - vesting.LockEndBlock
	blocksFromStart := height - vesting.LockEndBlock
	numOfPeriodsFromStart := blocksFromStart / vesting.FreeCoinsBlockPeriod
	numOfPeriods := allVestingBlocks / vesting.FreeCoinsBlockPeriod
	rest := allVestingBlocks - numOfPeriods*vesting.FreeCoinsBlockPeriod

	if blocksFromStart == 0 {
		return sdk.ZeroInt()
	}
	if rest > 0 {
		numOfPeriods++
	}

	vested := sdk.NewIntFromUint64(vesting.Vested)

	withdrawableFromStart := vested.MulRaw(numOfPeriodsFromStart).QuoRaw(numOfPeriods)
	return withdrawableFromStart.Sub(sdk.NewIntFromUint64(vesting.Withdrawn))

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
