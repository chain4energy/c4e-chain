package v130

import (
	"cosmossdk.io/math"
	cfeupgradetypes "github.com/chain4energy/c4e-chain/app/upgrades"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"time"
)

const (
	MoondropVestingAccount      = "c4e1dsm96gwcv35m4rqd93pzcsztpkrqe0ev7getj8"
	AirdropModuleAccount        = "fairdrop"
	AirdropModuleAccountAddress = "c4e1dutmadwfernuzmzk8ndtfah254yhrnv34y68ts"
	NewAirdropVestingPoolOwner  = "c4e1p0smw03cwhqn05fkalfpcr0ngqv5jrpnx2cp54"
)

func MigrateAirdropModuleAccount(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) error {
	ctx.Logger().Info("migrating airdrop module account")

	accountVestingPools, found := appKeepers.GetC4eVestingKeeper().GetAccountVestingPools(ctx, NewAirdropVestingPoolOwner)
	if !found {
		ctx.Logger().Info("account vesting pools not found for NewAirdropVestingPoolOwner", "owner", NewAirdropVestingPoolOwner)
		return nil
	}
	vestingDenom := appKeepers.GetC4eVestingKeeper().Denom(ctx)
	airdropAccBalance := bankkeeper.Keeper.GetAllBalances(*appKeepers.GetBankKeeper(), ctx, authtypes.NewModuleAddress(AirdropModuleAccount))
	if len(airdropAccBalance) == 0 {
		ctx.Logger().Info("no coins found for AirdropModuleAccount", "address", AirdropModuleAccountAddress)
		return nil
	}
	coinsAmount := airdropAccBalance.AmountOf(vestingDenom)
	coins := sdk.NewCoins(sdk.NewCoin(vestingDenom, coinsAmount))

	airdropModuleAccAddress, err := sdk.AccAddressFromBech32(AirdropModuleAccountAddress)
	if err != nil {
		return err
	}
	err = bankkeeper.Keeper.SendCoinsFromAccountToModule(*appKeepers.GetBankKeeper(), ctx, airdropModuleAccAddress, cfevestingtypes.ModuleName, coins)
	if err != nil {
		ctx.Logger().Info("migrate airdrop module account error", "err", err)
		return err
	}

	fairdropPool := cfevestingtypes.VestingPool{
		Name:            "Fairdrop",
		VestingType:     FairdropTypeName,
		LockStart:       time.Date(2023, 6, 1, 23, 59, 59, 0, time.UTC),
		LockEnd:         time.Date(2026, 6, 1, 23, 59, 59, 0, time.UTC),
		InitiallyLocked: coinsAmount,
		Withdrawn:       math.ZeroInt(),
		Sent:            math.ZeroInt(),
		GenesisPool:     true,
		Reservations:    nil,
	}
	accountVestingPools.VestingPools = append(accountVestingPools.VestingPools, &fairdropPool)
	appKeepers.GetC4eVestingKeeper().SetAccountVestingPools(ctx, accountVestingPools)
	return nil
}

func MigrateMoondropVestingAccount(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) error {
	ctx.Logger().Info("migrating moondrop module account")

	moondropAccAddress, err := sdk.AccAddressFromBech32(MoondropVestingAccount)
	if err != nil {
		return err
	}
	moondropAccount := appKeepers.GetAccountKeeper().GetAccount(ctx, moondropAccAddress)
	if moondropAccount == nil {
		ctx.Logger().Info("moondrop account not found", "address", MoondropVestingAccount)
		return nil
	}

	vestingAccount, ok := moondropAccount.(*vestingtypes.ContinuousVestingAccount)
	if !ok {
		ctx.Logger().Info("moondrop account is not of *vestingtypes.ContinuousVestingAccount type", "vestingAccount", vestingAccount)
		return nil
	}

	vestingDenom := appKeepers.GetC4eVestingKeeper().Denom(ctx)
	originalVestingAmount := vestingAccount.OriginalVesting
	vestingAccount.OriginalVesting = sdk.NewCoins()
	appKeepers.GetAccountKeeper().SetAccount(ctx, vestingAccount)
	err = bankkeeper.Keeper.SendCoinsFromAccountToModule(*appKeepers.GetBankKeeper(), ctx, moondropAccAddress, cfevestingtypes.ModuleName, originalVestingAmount)
	if err != nil {
		ctx.Logger().Info("migrate moondrop vesting account error", "err", err)
		return err
	}

	accountVestingPools := cfevestingtypes.AccountVestingPools{
		Owner: MoondropVestingAccount,
		VestingPools: []*cfevestingtypes.VestingPool{
			{
				Name:            "Moondrop",
				VestingType:     "Moondrop",
				LockStart:       time.Date(2024, 9, 26, 2, 00, 00, 00, time.UTC),
				LockEnd:         time.Date(2026, 9, 25, 2, 00, 00, 00, time.UTC),
				InitiallyLocked: originalVestingAmount.AmountOf(vestingDenom),
				Withdrawn:       sdk.ZeroInt(),
				Sent:            sdk.ZeroInt(),
				GenesisPool:     true,
				Reservations:    nil,
			},
		},
	}
	appKeepers.GetC4eVestingKeeper().SetAccountVestingPools(ctx, accountVestingPools)
	return nil
}
