package v131

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
	StrategicReserveAccountNewAmount     = math.NewInt(50_000_000_000_000)
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

	strategicReserveShortTermPoolFound := false
	for _, pool := range accountVestingPools.VestingPools {
		if pool.Name == StrategicReservceShortTermPool {
			pool.InitiallyLocked = pool.InitiallyLocked.Sub(AmountToBurnFromStrategicReservePool)
			if pool.InitiallyLocked.IsNegative() {
				ctx.Logger().Info("after substracting amount to burn from strategic reserve pool initially locked amount is negative")
				return nil
			}
			strategicReserveShortTermPoolFound = true
			break
		}
	}

	if !strategicReserveShortTermPoolFound {
		ctx.Logger().Info("strategic reserve short term pool not found", "owner", StrategicReservceShortTermPoolAccount)
		return nil
	}

	if err := burnCoinsFromAnyModule(ctx, appKeepers, cfevestingtypes.ModuleName, AmountToBurnFromStrategicReservePool); err != nil {
		ctx.Logger().Error("send and burn coins from module error", "err", err)
		return err
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
		ctx.Logger().Info("strategic reserve account is not of *vestingtypes.ContinuousVestingAccount type")
		return nil
	}

	denom := appKeepers.GetC4eVestingKeeper().Denom(ctx)

	if vestingAccount.OriginalVesting.AmountOf(denom).LTE(StrategicReserveAccountNewAmount) {
		ctx.Logger().Info("after substracting amount to burn strategic reserve account and amount to send to liquidity pool owner from vesting account original vesting is negative")
		return nil
	}

	vestingAccount.OriginalVesting = sdk.NewCoins(sdk.NewCoin(denom, StrategicReserveAccountNewAmount))
	appKeepers.GetAccountKeeper().SetAccount(ctx, vestingAccount)

	ctx.Logger().Info("strategic reserve account updated, new original vesting", "originalVesting", vestingAccount.OriginalVesting)
	spendableCoins := bankkeeper.Keeper.SpendableCoins(*appKeepers.GetBankKeeper(), ctx, strategicReserveAccountAddress)

	if AmountToSendToLiquidtyPoolOwner.GT(spendableCoins.AmountOf(denom)) {
		ctx.Logger().Info("spendable coins are less than coins to send to liquidity pool owner",
			"spendableCoins", spendableCoins, "amountToSendToLiquidtyPoolOwner", AmountToSendToLiquidtyPoolOwner)
		return nil
	}

	coinsToSendToLiquidityPoolOwner := sdk.NewCoins(sdk.NewCoin(denom, AmountToSendToLiquidtyPoolOwner))

	err = bankkeeper.Keeper.SendCoins(*appKeepers.GetBankKeeper(), ctx, strategicReserveAccountAddress, liquidityPoolOwnerAccountAddress, coinsToSendToLiquidityPoolOwner)
	if err != nil {
		ctx.Logger().Error("migrate moondrop vesting account error", "err", err)
		return err
	}

	spendableCoins = bankkeeper.Keeper.SpendableCoins(*appKeepers.GetBankKeeper(), ctx, strategicReserveAccountAddress)
	ctx.Logger().Info("spendable coins after sending coins to liquidity pool owner", "spendableCoins", spendableCoins)

	err = bankkeeper.Keeper.SendCoinsFromAccountToModule(
		*appKeepers.GetBankKeeper(),
		ctx,
		strategicReserveAccountAddress,
		cfemintertypes.ModuleName,
		spendableCoins,
	)
	if err != nil {
		ctx.Logger().Error("migrate moondrop vesting account error", "err", err)
		return err
	}

	if err = bankkeeper.Keeper.BurnCoins(*appKeepers.GetBankKeeper(), ctx, cfemintertypes.ModuleName, spendableCoins); err != nil {
		ctx.Logger().Error("burn coins error", "err", err)
		return err
	}

	return nil
}

func UpdateCommunityPool(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) error {
	ctx.Logger().Info("updating community pool")
	feePool := appKeepers.GetDistributionKeeper().GetFeePool(ctx)
	communityPoolBefore := feePool.CommunityPool

	denom := appKeepers.GetC4eVestingKeeper().Denom(ctx)

	if communityPoolBefore.AmountOf(denom).LTE(sdk.NewDecFromInt(CommunityPoolNewAmount)) {
		ctx.Logger().Info("community pool amount before migration was lower or equal to community pool new amount")
		return nil
	}

	communityPool := sdk.NewDecCoins(sdk.NewDecCoin(denom, CommunityPoolNewAmount))
	feePool.CommunityPool = communityPool

	amountToBurn := communityPoolBefore.AmountOf(denom).TruncateInt().Sub(CommunityPoolNewAmount)
	if err := burnCoinsFromAnyModule(ctx, appKeepers, distrtypes.ModuleName, amountToBurn); err != nil {
		ctx.Logger().Error("send and burn coins from module error", "err", err)
		return err
	}

	appKeepers.GetDistributionKeeper().SetFeePool(ctx, feePool)
	return nil
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
