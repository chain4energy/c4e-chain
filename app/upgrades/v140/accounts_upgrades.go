package v140

import (
	"cosmossdk.io/math"
	cfeupgradetypes "github.com/chain4energy/c4e-chain/app/upgrades"
	cfemintertypes "github.com/chain4energy/c4e-chain/x/cfeminter/types"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
)

const (
	StrategicReservceShortTermPoolAccount = "c4e1p0smw03cwhqn05fkalfpcr0ngqv5jrpnx2cp54"
	StrategicReservceShortTermPool        = "Strategic reserve short term round pool"

	StrategicReserveAccount = "c4e1hcfjejmxzl8d95xka5j8cjegmf32u2lee3q422"
	LiquidityPoolOwner      = "c4e16n7yweagu3fxfzvay6cz035hddda7z3ntdxq3l"
)

var (
	AmountToBurnFromStrategicReservePool = math.NewInt(20_000_000_000_000)
	AmountToBurnStrategicReserveAccount  = math.NewInt(30_000_000_000_000)
	AmountToSendToLiquidtyPoolOwner      = math.NewInt(10_000_000_000_000)
	CommunityPoolNewAmount               = math.NewInt(40_000_000_000_000)
)

func UpdateStrategicReserveShortTermPool(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) error {
	ctx.Logger().Info("updating strategic reserve short term pool")

	accountVestingPools, found := appKeepers.GetC4eVestingKeeper().GetAccountVestingPools(ctx, StrategicReservceShortTermPoolAccount)
	if !found {
		ctx.Logger().Info("account vesting pools not found for StrategicReservceShortTermPoolAccount", "owner", StrategicReservceShortTermPoolAccount)
		return nil
	}

	if err := burnCoinsFromAnyModule(ctx, appKeepers, cfevestingtypes.ModuleName, AmountToBurnFromStrategicReservePool); err != nil {
		ctx.Logger().Info("send and burn coins from module error", "err", err)
		return nil
	}

	for _, pool := range accountVestingPools.VestingPools {
		if pool.Name == StrategicReservceShortTermPool {
			pool.InitiallyLocked = pool.InitiallyLocked.Sub(AmountToBurnFromStrategicReservePool)
		}
	}

	appKeepers.GetC4eVestingKeeper().SetAccountVestingPools(ctx, accountVestingPools)
	return nil
}

func UpdateStrategicReserveAccount(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) error {
	ctx.Logger().Info("updating strategic reserve account")

	strategicReserveAccountAddress, err := sdk.AccAddressFromBech32(StrategicReserveAccount)
	if err != nil {
		return err
	}

	liquidityPoolOwnerAccountAddress, err := sdk.AccAddressFromBech32(LiquidityPoolOwner)
	if err != nil {
		return err
	}

	strategicReserveAccount := appKeepers.GetAccountKeeper().GetAccount(ctx, strategicReserveAccountAddress)
	if strategicReserveAccount == nil {
		ctx.Logger().Info("strategic reserve account not found", "address", strategicReserveAccountAddress)
		return nil
	}

	vestingAccount, ok := strategicReserveAccount.(*vestingtypes.ContinuousVestingAccount)
	if !ok {
		ctx.Logger().Info("strategic reserve account is not of *vestingtypes.ContinuousVestingAccount type", "vestingAccount", vestingAccount)
		return nil
	}

	denom := appKeepers.GetC4eVestingKeeper().Denom(ctx)
	vestingAccount.OriginalVesting = vestingAccount.OriginalVesting.Sub(sdk.NewCoin(denom, AmountToBurnStrategicReserveAccount))
	appKeepers.GetAccountKeeper().SetAccount(ctx, vestingAccount)

	coinsToSendToLiquidityPoolOwner := sdk.NewCoins(sdk.NewCoin(denom, AmountToSendToLiquidtyPoolOwner))
	err = bankkeeper.Keeper.SendCoins(*appKeepers.GetBankKeeper(), ctx, strategicReserveAccountAddress, liquidityPoolOwnerAccountAddress, coinsToSendToLiquidityPoolOwner)
	if err != nil {
		ctx.Logger().Info("migrate moondrop vesting account error", "err", err)
		return err
	}

	spendableCoins := bankkeeper.Keeper.SpendableCoins(*appKeepers.GetBankKeeper(), ctx, strategicReserveAccountAddress)
	err = bankkeeper.Keeper.SendCoinsFromAccountToModule(
		*appKeepers.GetBankKeeper(),
		ctx,
		strategicReserveAccountAddress,
		cfemintertypes.ModuleName,
		spendableCoins,
	)
	if err != nil {
		ctx.Logger().Info("migrate moondrop vesting account error", "err", err)
		return err
	}

	if err = bankkeeper.Keeper.BurnCoins(*appKeepers.GetBankKeeper(), ctx, cfemintertypes.ModuleName, spendableCoins); err != nil {
		ctx.Logger().Info("burn coins error", "err", err)
		return nil
	}

	return nil
}

func UpdateCommunityPool(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) error {
	feePool := appKeepers.GetDistributionKeeper().GetFeePool(ctx)
	communityPoolBefore := feePool.CommunityPool

	denom := appKeepers.GetC4eVestingKeeper().Denom(ctx)
	communityPool := sdk.NewDecCoinsFromCoins(sdk.NewCoin(denom, CommunityPoolNewAmount))
	feePool.CommunityPool = communityPool

	amountToBurn := communityPoolBefore.AmountOf(denom).TruncateInt().Sub(CommunityPoolNewAmount)
	if err := burnCoinsFromAnyModule(ctx, appKeepers, distrtypes.ModuleName, amountToBurn); err != nil {
		ctx.Logger().Info("send and burn coins from module error", "err", err)
		return nil
	}

	appKeepers.GetDistributionKeeper().SetFeePool(ctx, feePool)
	return nil
}

func burnCoinsFromAnyModule(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers, fromModule string, amount math.Int) error {
	denom := appKeepers.GetC4eVestingKeeper().Denom(ctx)
	coins := sdk.NewCoins(sdk.NewCoin(denom, amount))

	if err := bankkeeper.Keeper.SendCoinsFromModuleToModule(*appKeepers.GetBankKeeper(), ctx, fromModule, cfemintertypes.ModuleName, coins); err != nil {
		ctx.Logger().Info("send coins from module to module error", "err", err)
		return nil
	}
	if err := bankkeeper.Keeper.BurnCoins(*appKeepers.GetBankKeeper(), ctx, cfemintertypes.ModuleName, coins); err != nil {
		ctx.Logger().Info("burn coins error", "err", err)
		return nil
	}
	return nil
}
