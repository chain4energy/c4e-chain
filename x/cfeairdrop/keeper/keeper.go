package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		accountKeeper  types.AccountKeeper
		bankKeeper     types.BankKeeper
		feeGrantKeeper types.FeeGrantKeeper
		stakingKeeper  types.StakingKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,

	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	feeGrantKeeper types.FeeGrantKeeper,
	stakingKeeper types.StakingKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:            cdc,
		storeKey:       storeKey,
		memKey:         memKey,
		paramstore:     ps,
		accountKeeper:  accountKeeper,
		bankKeeper:     bankKeeper,
		feeGrantKeeper: feeGrantKeeper,
		stakingKeeper:  stakingKeeper,
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
