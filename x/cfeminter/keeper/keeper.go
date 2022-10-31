package keeper

import (
	"fmt"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		bankKeeper    types.BankKeeper
		stakingKeeper types.StakingKeeper
		collectorName string
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,

	bankKeeper types.BankKeeper,
	stakingKeeper types.StakingKeeper,
	collectorName string,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		bankKeeper:    bankKeeper,
		stakingKeeper: stakingKeeper,
		collectorName: collectorName,
	}
}

func (k Keeper) GetCollectorName() string {
	return k.collectorName
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) Mint(ctx sdk.Context) (sdk.Int, error) {
	lastBlockTimeForMinter := k.GetMinterState(ctx).LastMintBlockTime
	lastBlockTime := ctx.BlockTime()

	if lastBlockTimeForMinter.After(lastBlockTime) || lastBlockTimeForMinter.Equal(lastBlockTime) {
		k.Logger(ctx).Info("mint last mint block time is smaller than current block time - possible for first block after genesis init",
			"lastBlockTime", lastBlockTimeForMinter, "currentBlockTime", lastBlockTime)
		return sdk.ZeroInt(), nil
	}
	return k.mint(ctx, 0)
}

func (k Keeper) mint(ctx sdk.Context, level int) (sdk.Int, error) {
	minterState := k.GetMinterState(ctx)
	params := k.GetParams(ctx)
	minter := params.Minter

	currentPeriod, previousPeriod := getCurrentAndPreviousPeriod(minter, &minterState)

	if currentPeriod == nil {
		k.Logger(ctx).Error("mint - current period not found error", "lev", level, "position", minterState.Position)
		return sdk.ZeroInt(), sdkerrors.Wrapf(sdkerrors.ErrNotFound, "minter - mint - current period for position %d not found", minterState.Position)
	}

	var periodStart time.Time
	if previousPeriod == nil {
		periodStart = minter.Start
	} else {
		periodStart = *previousPeriod.PeriodEnd
	}

	expectedAmountToMint := currentPeriod.AmountToMint(k.Logger(ctx), &minterState, periodStart, ctx.BlockTime())
	expectedAmountToMint = expectedAmountToMint.Add(minterState.RemainderFromPreviousPeriod)

	amount := expectedAmountToMint.TruncateInt().Sub(minterState.AmountMinted)
	k.Logger(ctx).Debug("mint", "lev", level, "minterState", minterState, "periodStart", periodStart, "currentPeriod", currentPeriod,
		"previousPeriod", previousPeriod, "expectedAmountToMint", expectedAmountToMint, "amount", amount)

	remainder := expectedAmountToMint.Sub(expectedAmountToMint.TruncateDec())

	coin := sdk.NewCoin(params.MintDenom, amount)
	coins := sdk.NewCoins(coin)

	err := k.MintCoins(ctx, coins)
	if err != nil {
		k.Logger(ctx).Error("mint - mint coins error", "lev", level, "error", err.Error())
		return sdk.ZeroInt(), sdkerrors.Wrap(err, "minter mint coins error")
	}

	err = k.AddCollectedFees(ctx, coins)
	if err != nil {
		k.Logger(ctx).Error("mint - add collected fees error", "lev", level, "error", err.Error())
		return sdk.ZeroInt(), sdkerrors.Wrap(err, "minter - mint - add collected fees error")
	}

	minterState.AmountMinted = minterState.AmountMinted.Add(amount)
	minterState.LastMintBlockTime = ctx.BlockTime()
	minterState.RemainderToMint = remainder

	var result sdk.Int
	if currentPeriod.PeriodEnd == nil || ctx.BlockTime().Before(*currentPeriod.PeriodEnd) {
		k.SetMinterState(ctx, minterState)
		result = amount
	} else {
		k.SetMinterStateHistory(ctx, minterState)
		k.Logger(ctx).Debug("mint - set minter state history", "lev", level, "minterState", minterState.String())
		minterState = types.MinterState{
			Position:                    minterState.Position + 1,
			AmountMinted:                sdk.ZeroInt(),
			RemainderToMint:             sdk.ZeroDec(),
			RemainderFromPreviousPeriod: remainder,
			LastMintBlockTime:           ctx.BlockTime(),
		}
		k.SetMinterState(ctx, minterState)
		minted, err := k.mint(ctx, level+1)
		if err != nil {
			k.Logger(ctx).Error("mint - sub mint error", "lev", level, "error", err.Error())
			return minted, err
		}
		result = minted.Add(amount)
	}

	k.Logger(ctx).Debug("mint ret", "lev", level, "result", result, "minterState", minterState)
	return result, nil
}

func (k Keeper) GetCurrentInflation(ctx sdk.Context) (sdk.Dec, error) {
	minterState := k.GetMinterState(ctx)
	params := k.GetParams(ctx)
	minter := params.Minter

	currentPeriod, previousPeriod := getCurrentAndPreviousPeriod(minter, &minterState)

	if currentPeriod == nil {
		k.Logger(ctx).Error("minter current period not found error", "position", minterState.Position)
		return sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrNotFound, "minter current period for position %d not found", minterState.Position)
	}

	var periodStart time.Time
	if previousPeriod == nil {
		periodStart = minter.Start
	} else {
		periodStart = *previousPeriod.PeriodEnd
	}

	supply := k.bankKeeper.GetSupply(ctx, params.MintDenom)
	result := currentPeriod.CalculateInfation(supply.Amount, periodStart, ctx.BlockHeader().Time)
	k.Logger(ctx).Debug("get current inflation", "currentPeriod", currentPeriod, "previousPeriod", previousPeriod, "periodStart",
		periodStart, "supply", supply, "blockTime", ctx.BlockHeader().Time, "result", result)
	return result, nil
}

func getCurrentAndPreviousPeriod(minter types.Minter, state *types.MinterState) (currentPeriod *types.MintingPeriod, previousPeriod *types.MintingPeriod) {
	currentId := state.Position
	for _, period := range minter.Periods {
		if period.Position == currentId {
			currentPeriod = period
		}
		if previousPeriod == nil {
			if period.Position < currentId {
				previousPeriod = period
			}
		} else {
			if period.Position < currentId && period.Position > previousPeriod.Position {
				previousPeriod = period
			}
		}
	}
	return currentPeriod, previousPeriod
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

// AddCollectedFees implements an alias call to the underlying supply keeper's
// AddCollectedFees to be used in BeginBlocker.
func (k Keeper) AddCollectedFees(ctx sdk.Context, fees sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.collectorName, fees)
}

// BondedRatio implements an alias call to the underlying staking keeper's
// BondedRatio to be used in BeginBlocker.
func (k Keeper) BondedRatio(ctx sdk.Context) sdk.Dec {
	return k.stakingKeeper.BondedRatio(ctx)
}
