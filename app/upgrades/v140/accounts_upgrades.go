package v140

import (
	"cosmossdk.io/math"
	cfeupgradetypes "github.com/chain4energy/c4e-chain/app/upgrades"
	cfedistributormoduletypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

const (
	StrategicReservceShortTermPoolAccount = "c4e1p0smw03cwhqn05fkalfpcr0ngqv5jrpnx2cp54"
	StrategicReservceShortTermPool        = "Strategic reserve short term round pool"

	StrategicReservceAccount = "c4e1hcfjejmxzl8d95xka5j8cjegmf32u2lee3q422"
	LiquidityPoolOwner       = "c4e16n7yweagu3fxfzvay6cz035hddda7z3ntdxq3l"
)

var (
	AmountToBurnFromStrategicReservePool = math.NewInt(10000000000000)
	AmountToBurnStrategicReserveAccount  = math.NewInt(30000000000000)
	AmountToSendToLiquidtyPoolOwner      = math.NewInt(10000000000000)

	CommunityPoolNewAmount = math.NewInt(40000000000000)
)

func UpdateStrategicReserveShortTermPool(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) error {
	ctx.Logger().Info("migrating airdrop module account")

	accountVestingPools, found := appKeepers.GetC4eVestingKeeper().GetAccountVestingPools(ctx, StrategicReservceShortTermPoolAccount)
	if !found {
		ctx.Logger().Info("account vesting pools not found for StrategicReservceShortTermPoolAccount", "owner", StrategicReservceShortTermPoolAccount)
		return nil
	}
	vestingDenom := appKeepers.GetC4eVestingKeeper().Denom(ctx)
	if err := bankkeeper.Keeper.BurnCoins(*appKeepers.GetBankKeeper(), ctx, cfevestingtypes.ModuleName, sdk.NewCoins(sdk.NewCoin(vestingDenom, AmountToBurnFromStrategicReservePool))); err != nil {
		ctx.Logger().Info("burn coins error", "err", err)
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
	ctx.Logger().Info("migrating moondrop module account")

	strategicReserveAccountAddress, err := sdk.AccAddressFromBech32(StrategicReservceAccount)
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
		ctx.Logger().Info("moondrop account is not of *vestingtypes.ContinuousVestingAccount type", "vestingAccount", vestingAccount)
		return nil
	}

	vestingDenom := appKeepers.GetC4eVestingKeeper().Denom(ctx)
	vestingAccount.OriginalVesting = vestingAccount.OriginalVesting.Sub(sdk.NewCoin(vestingDenom, AmountToBurnStrategicReserveAccount))
	appKeepers.GetAccountKeeper().SetAccount(ctx, vestingAccount)

	err = bankkeeper.Keeper.SendCoins(*appKeepers.GetBankKeeper(), ctx, strategicReserveAccountAddress, liquidityPoolOwnerAccountAddress, sdk.NewCoins(sdk.NewCoin(vestingDenom, AmountToSendToLiquidtyPoolOwner)))
	if err != nil {
		ctx.Logger().Info("migrate moondrop vesting account error", "err", err)
		return err
	}

	spendableCoins := bankkeeper.Keeper.SpendableCoins(*appKeepers.GetBankKeeper(), ctx, strategicReserveAccountAddress)

	err = bankkeeper.Keeper.SendCoinsFromAccountToModule(*appKeepers.GetBankKeeper(), ctx, strategicReserveAccountAddress, cfevestingtypes.ModuleName, sdk.NewCoins(sdk.NewCoin(vestingDenom, AmountToSendToLiquidtyPoolOwner)))
	if err != nil {
		ctx.Logger().Info("migrate moondrop vesting account error", "err", err)
		return err
	}

	if err = bankkeeper.Keeper.BurnCoins(*appKeepers.GetBankKeeper(), ctx, cfevestingtypes.ModuleName, spendableCoins); err != nil {
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

	amountToBurn := communityPoolBefore.AmountOf(denom).TruncateInt().Sub(AmountToBurnFromStrategicReservePool)
	coinsToBurn := sdk.NewCoins(sdk.NewCoin(denom, amountToBurn))
	if err := bankkeeper.Keeper.BurnCoins(*appKeepers.GetBankKeeper(), ctx, cfedistributormoduletypes.ModuleName, coinsToBurn); err != nil {
		ctx.Logger().Info("burn coins error", "err", err)
		return nil
	}
	appKeepers.GetDistributionKeeper().SetFeePool(ctx, feePool)
	return nil
}
