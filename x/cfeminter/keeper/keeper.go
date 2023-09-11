package keeper

import (
	"cosmossdk.io/errors"
	"fmt"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"time"

	"github.com/cometbft/cometbft/libs/log"

	"github.com/chain4energy/c4e-chain/v2/x/cfeminter/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type (
	Keeper struct {
		cdc           codec.BinaryCodec
		storeKey      storetypes.StoreKey
		memKey        storetypes.StoreKey
		bankKeeper    types.BankKeeper
		stakingKeeper types.StakingKeeper
		collectorName string
		authority     string
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	bankKeeper types.BankKeeper,
	stakingKeeper types.StakingKeeper,
	collectorName string,
	authority string,
) *Keeper {
	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		bankKeeper:    bankKeeper,
		stakingKeeper: stakingKeeper,
		collectorName: collectorName,
		authority:     authority,
	}
}

func (k Keeper) GetCollectorName() string {
	return k.collectorName
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetCurrentInflation(ctx sdk.Context) (sdk.Dec, error) {
	minterState := k.GetMinterState(ctx)
	params := k.GetParams(ctx)
	currentMinter, previousMinter := getCurrentAndPreviousMinter(params.Minters, &minterState)
	if currentMinter == nil {
		k.Logger(ctx).Error("minter current sequence id not found error", "SequenceId", minterState.SequenceId)
		return sdk.ZeroDec(), errors.Wrapf(sdkerrors.ErrNotFound, "minter current period for SequenceId %d not found", minterState.SequenceId)
	}

	var startTime time.Time
	if previousMinter == nil {
		startTime = params.StartTime
	} else {
		startTime = *previousMinter.EndTime
	}

	supply := k.bankKeeper.GetSupply(ctx, params.MintDenom)
	result := currentMinter.CalculateInflation(supply.Amount, startTime, ctx.BlockHeader().Time)
	k.Logger(ctx).Debug("get current inflation", "currentMinter", currentMinter.GetMinterJSON(), "previousMinter", previousMinter.GetMinterJSON(), "startTime",
		startTime, "supply", supply, "blockTime", ctx.BlockHeader().Time, "result", result)
	return result, nil
}

// MintCoins implements an alias call to the underlying supply keeper's
// MintCoins to be used in BeginBlocker.
func (k Keeper) MintCoins(ctx sdk.Context, newCoins sdk.Coins) error {
	if newCoins.Empty() {
		// skip as no coins need to be minted
		return nil
	}
	return k.bankKeeper.MintCoins(ctx, types.ModuleName, newCoins)
}

// SendMintedCoins implements an alias call to the underlying supply keeper's
// SendMintedCoins to be used in BeginBlocker.
func (k Keeper) SendMintedCoins(ctx sdk.Context, fees sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.collectorName, fees)
}

// BondedRatio implements an alias call to the underlying staking keeper's
// BondedRatio to be used in BeginBlocker.
func (k Keeper) BondedRatio(ctx sdk.Context) sdk.Dec {
	return k.stakingKeeper.BondedRatio(ctx)
}
