package keeper

import (
	"fmt"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/cometbft/cometbft/libs/log"

	"github.com/chain4energy/c4e-chain/v2/x/cfeclaim/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey storetypes.StoreKey
		memKey   storetypes.StoreKey

		accountKeeper  types.AccountKeeper
		bankKeeper     types.BankKeeper
		feeGrantKeeper types.FeeGrantKeeper
		stakingKeeper  types.StakingKeeper
		vestingKeeper  types.CfeVestingKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,

	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	feeGrantKeeper types.FeeGrantKeeper,
	stakingKeeper types.StakingKeeper,
	vestingKeeper types.CfeVestingKeeper,
) *Keeper {
	return &Keeper{
		cdc:            cdc,
		storeKey:       storeKey,
		memKey:         memKey,
		accountKeeper:  accountKeeper,
		bankKeeper:     bankKeeper,
		feeGrantKeeper: feeGrantKeeper,
		stakingKeeper:  stakingKeeper,
		vestingKeeper:  vestingKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetAccountCoins(ctx sdk.Context, account sdk.AccAddress) sdk.Coins {
	return k.bankKeeper.GetAllBalances(ctx, account)
}

func (k Keeper) GetAccountAddressModuleAccount(ctx sdk.Context, accountName string) sdk.AccAddress {
	return k.accountKeeper.GetModuleAccount(ctx, accountName).GetAddress()
}

func (k Keeper) GetAccountCoinsForModuleAccount(ctx sdk.Context, accountName string) sdk.Coins {
	accAddress := k.GetAccountAddressModuleAccount(ctx, accountName)
	return k.GetAccountCoins(ctx, accAddress)
}
