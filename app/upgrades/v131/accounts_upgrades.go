package v131

import (
	"cosmossdk.io/math"
	cfeupgradetypes "github.com/chain4energy/c4e-chain/app/upgrades"
	cfemintertypes "github.com/chain4energy/c4e-chain/x/cfeminter/types"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

const (
	StrategicReservcePoolOwnerAccount = "c4e1p0smw03cwhqn05fkalfpcr0ngqv5jrpnx2cp54"
	StrategicReservceShortTermPool    = "Strategic reserve short term round pool"

	StrategicReserveAccount = "c4e1hcfjejmxzl8d95xka5j8cjegmf32u2lee3q422"
	LiquidityPoolOwner      = "c4e16n7yweagu3fxfzvay6cz035hddda7z3ntdxq3l"
	EarlyBirdRoundPoolName  = "Early-bird round pool"
	PublicRoundPoolName     = "Public round pool"
)

var (
	liquidityPoolOwnerAccountAddress, _ = sdk.AccAddressFromBech32(LiquidityPoolOwner)
	strategicReserveAccountAddress, _   = sdk.AccAddressFromBech32(StrategicReserveAccount)
	strategicReservePoolOwnerAddress, _ = sdk.AccAddressFromBech32(StrategicReservcePoolOwnerAccount)
)

var (
	AmountToBurnFromStrategicReservePool = math.NewInt(20_000_000_000_000)

	AmountToBurnFromCommunityPool = math.NewInt(60_000_000_000_000)

	AmountToUnlockFromStrategicReserveAccount = math.NewInt(34_200_000_000_000)
	AmountToSendToLiquidtyPoolOwner           = math.NewInt(10_000_000_000_000)
	AmountToBurnFromStrategicReserveAccount   = math.NewInt(24_200_000_000_000)
)

func UpdateStrategicReserveShortTermPool(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) error {
	ctx.Logger().Info("updating strategic reserve short term pool")

	accountVestingPools, found := appKeepers.GetC4eVestingKeeper().GetAccountVestingPools(ctx, StrategicReservcePoolOwnerAccount)
	if !found {
		ctx.Logger().Info("account vesting pools not found for StrategicReservcePoolOwnerAccount", "owner", StrategicReservcePoolOwnerAccount)
		return nil
	}

	strategicReservePoolFound := false
	earlyBirdPoolFound := false
	publicRoundPoolFound := false

	mergedPoolsAmountSum := math.ZeroInt()
	for _, pool := range accountVestingPools.VestingPools {
		switch pool.Name {
		case StrategicReservceShortTermPool:
			if pool.GetLockedNotReserved().Sub(AmountToBurnFromStrategicReservePool).IsNegative() {
				ctx.Logger().Info("after substracting amount to burn from strategic reserve pool initially currently locked amount is negative")
				return nil
			}

			pool.InitiallyLocked = pool.InitiallyLocked.Sub(AmountToBurnFromStrategicReservePool)
			strategicReservePoolFound = true
			break

		case EarlyBirdRoundPoolName:
			lockedNotReserved := pool.GetLockedNotReserved()
			pool.InitiallyLocked = pool.InitiallyLocked.Sub(lockedNotReserved)
			mergedPoolsAmountSum = mergedPoolsAmountSum.Add(lockedNotReserved)
			earlyBirdPoolFound = true
			break

		case PublicRoundPoolName:
			lockedNotReserved := pool.GetLockedNotReserved()
			pool.InitiallyLocked = pool.InitiallyLocked.Sub(lockedNotReserved)
			mergedPoolsAmountSum = mergedPoolsAmountSum.Add(lockedNotReserved)
			publicRoundPoolFound = true
			break
		}
	}
	if !strategicReservePoolFound || !earlyBirdPoolFound || !publicRoundPoolFound {
		ctx.Logger().Info("one of the pools not found", "strategicReservePoolFound", strategicReservePoolFound,
			"earlyBirdPoolFound", earlyBirdPoolFound, "publicRoundPoolFound", publicRoundPoolFound)
		return nil
	}

	if err := burnCoinsFromAnyModule(ctx, appKeepers, cfevestingtypes.ModuleName, AmountToBurnFromStrategicReservePool); err != nil {
		ctx.Logger().Error("send and burn coins from module error", "err", err)
		return err
	}

	denom := appKeepers.GetC4eVestingKeeper().Denom(ctx)
	coins := sdk.NewCoins(sdk.NewCoin(denom, mergedPoolsAmountSum))
	if err := bankkeeper.Keeper.SendCoinsFromModuleToAccount(*appKeepers.GetBankKeeper(), ctx, cfevestingtypes.ModuleName, strategicReservePoolOwnerAddress, coins); err != nil {
		ctx.Logger().Error("send coins from module to module error", "err", err)
		return err
	}

	appKeepers.GetC4eVestingKeeper().SetAccountVestingPools(ctx, accountVestingPools)
	ctx.Logger().Info("strategic reserve short term pool updated")
	return nil
}

func UpdateStrategicReserveAccount(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) error {
	ctx.Logger().Info("updating strategic reserve account")

	strategicReserveAccount := appKeepers.GetAccountKeeper().GetAccount(ctx, strategicReserveAccountAddress)
	if strategicReserveAccount == nil {
		ctx.Logger().Info("strategic reserve account not found", "address", strategicReserveAccountAddress)
		return nil
	}

	vestingAccount, ok := strategicReserveAccount.(*vestingtypes.ContinuousVestingAccount)
	if !ok {
		ctx.Logger().Info("strategic reserve account is not of *vestingtypes.ContinuousVestingAccount type")
		return nil
	}

	denom := appKeepers.GetC4eVestingKeeper().Denom(ctx)
	coinsToUnlockFromStrategicReserveAccount := sdk.NewCoins(sdk.NewCoin(denom, AmountToUnlockFromStrategicReserveAccount))
	lockedCoins := vestingAccount.LockedCoins(ctx.BlockTime())
	if !coinsToUnlockFromStrategicReserveAccount.IsAllLTE(lockedCoins) {
		ctx.Logger().Debug("unlock unbonded continuous vesting account coins - not enough to unlock", "lockedCoins", lockedCoins, "coinsToUnlockFromStrategicReserveAccount", coinsToUnlockFromStrategicReserveAccount)
		return nil
	}

	_, err := appKeepers.GetC4eVestingKeeper().UnlockUnbondedContinuousVestingAccountCoins(ctx,
		strategicReserveAccountAddress,
		coinsToUnlockFromStrategicReserveAccount,
	)
	if err != nil {
		ctx.Logger().Error("unlock unbonded continuous vesting account coins error", "err", err)
		return err
	}

	coinsToSendToLiquidityPoolOwner := sdk.NewCoins(sdk.NewCoin(denom, AmountToSendToLiquidtyPoolOwner))
	err = bankkeeper.Keeper.SendCoins(*appKeepers.GetBankKeeper(), ctx, strategicReserveAccountAddress, liquidityPoolOwnerAccountAddress, coinsToSendToLiquidityPoolOwner)
	if err != nil {
		ctx.Logger().Error("send coins error", "err", err)
		return err
	}

	return burnCoinsFromAccount(ctx, appKeepers, strategicReserveAccountAddress, AmountToBurnFromStrategicReserveAccount)
}

func UpdateCommunityPool(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) error {
	ctx.Logger().Info("updating community pool")
	feePool := appKeepers.GetDistributionKeeper().GetFeePool(ctx)
	communityPoolBefore := feePool.CommunityPool

	denom := appKeepers.GetC4eVestingKeeper().Denom(ctx)

	if communityPoolBefore.AmountOf(denom).LTE(sdk.NewDecFromInt(AmountToBurnFromCommunityPool)) {
		ctx.Logger().Info("community pool amount before migration was lower or equal to community pool new amount")
		return nil
	}

	coinsToBurnFromCommunityPool := sdk.NewCoins(sdk.NewCoin(appKeepers.GetC4eVestingKeeper().Denom(ctx), AmountToBurnFromCommunityPool))
	if err := appKeepers.GetDistributionKeeper().DistributeFromFeePool(ctx, coinsToBurnFromCommunityPool, liquidityPoolOwnerAccountAddress); err != nil {
		ctx.Logger().Error("distribute from fee pool error", "err", err)
		return err
	}

	return burnCoinsFromAccount(ctx, appKeepers, liquidityPoolOwnerAccountAddress, AmountToBurnFromCommunityPool)
}

func burnCoinsFromAccount(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers, accountAddress sdk.AccAddress, amount math.Int) error {
	denom := appKeepers.GetC4eVestingKeeper().Denom(ctx)
	coins := sdk.NewCoins(sdk.NewCoin(denom, amount))
	if err := bankkeeper.Keeper.SendCoinsFromAccountToModule(
		*appKeepers.GetBankKeeper(),
		ctx,
		accountAddress,
		cfemintertypes.ModuleName,
		coins,
	); err != nil {
		ctx.Logger().Error("send coins from account to module error", "err", err)
		return err
	}

	return bankkeeper.Keeper.BurnCoins(*appKeepers.GetBankKeeper(), ctx, cfemintertypes.ModuleName, coins)
}

func burnCoinsFromAnyModule(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers, fromModule string, amount math.Int) error {
	denom := appKeepers.GetC4eVestingKeeper().Denom(ctx)
	coins := sdk.NewCoins(sdk.NewCoin(denom, amount))

	if err := bankkeeper.Keeper.SendCoinsFromModuleToModule(*appKeepers.GetBankKeeper(), ctx, fromModule, cfemintertypes.ModuleName, coins); err != nil {
		ctx.Logger().Error("send coins from module to module error", "err", err)
		return err
	}
	if err := bankkeeper.Keeper.BurnCoins(*appKeepers.GetBankKeeper(), ctx, cfemintertypes.ModuleName, coins); err != nil {
		ctx.Logger().Error("burn coins error", "err", err)
		return err
	}
	return nil
}
