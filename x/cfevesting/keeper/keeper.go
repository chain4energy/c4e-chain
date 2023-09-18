package keeper

import (
	"fmt"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/cometbft/cometbft/libs/log"

	"github.com/chain4energy/c4e-chain/v2/x/cfevesting/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeKey     storetypes.StoreKey
		memKey       storetypes.StoreKey
		bank         types.BankKeeper
		staking      types.StakingKeeper
		account      types.AccountKeeper
		distribution types.DistributionKeeper
		gov          types.GovKeeper
		authority    string
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	bank types.BankKeeper,
	staking types.StakingKeeper,
	account types.AccountKeeper,
	distribution types.DistributionKeeper,
	gov types.GovKeeper,
	authority string,
) *Keeper {
	return &Keeper{
		cdc:          cdc,
		storeKey:     storeKey,
		memKey:       memKey,
		bank:         bank,
		staking:      staking,
		account:      account,
		distribution: distribution,
		gov:          gov,
		authority:    authority,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
