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
	minterState := k.GetMinterState(ctx)
	params := k.GetParams(ctx)
	minter := params.Minter

	currentPeriod, previousPeriod := getCurrentAndPreviousPeriod(minter, &minterState)

	if currentPeriod == nil {
		return sdk.ZeroInt(), sdkerrors.Wrapf(sdkerrors.ErrNotFound, "minter current period for position %d not found", minterState.Position)

	}

	var periodStart time.Time
	if previousPeriod == nil {
		periodStart = minter.Start
	} else {
		periodStart = *previousPeriod.PeriodEnd
	}
	k.Logger(ctx).Debug("Mint minterState start",
		"AmountMinted", minterState.AmountMinted.String(),
		"Position", minterState.Position,
		"RemainderFromPreviousPeriod", minterState.RemainderFromPreviousPeriod,
		"RemainderToMint", minterState.RemainderToMint,
		"LastMintBlockTime", minterState.LastMintBlockTime,
	)

	expectedAmountToMint := currentPeriod.AmountToMint(k.Logger(ctx), &minterState, periodStart, ctx.BlockTime())
	expectedAmountToMint = expectedAmountToMint.Add(minterState.RemainderFromPreviousPeriod)

	amount := expectedAmountToMint.TruncateInt().Sub(minterState.AmountMinted)
	k.Logger(ctx).Debug("Mint", "periodStart", periodStart, "periodEnd", currentPeriod.PeriodEnd, "currentPeriodType", currentPeriod.Type, "expectedAmountToMint", expectedAmountToMint, "amount", amount)

	remainder := expectedAmountToMint.Sub(expectedAmountToMint.TruncateDec())
	if amount.IsNegative() {
		k.Logger(ctx).Info("Mint negative amount - possible for first block after genesis init", "amount", amount)
		return sdk.ZeroInt(), nil
	}

	coin := sdk.NewCoin(params.MintDenom, amount)
	coins := sdk.NewCoins(coin)

	err := k.MintCoins(ctx, coins)
	if err != nil {
		return sdk.ZeroInt(), err
	}

	err = k.AddCollectedFees(ctx, coins)
	if err != nil {
		return sdk.ZeroInt(), err
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
		minterState = types.MinterState{
			Position:                    minterState.Position + 1,
			AmountMinted:                sdk.ZeroInt(),
			RemainderToMint:             sdk.ZeroDec(),
			RemainderFromPreviousPeriod: remainder,
			LastMintBlockTime:           ctx.BlockTime(),
		}
		k.SetMinterState(ctx, minterState)
		minted, err := k.Mint(ctx)
		if err != nil {
			return minted, err
		}
		result = minted.Add(amount)
	}

	k.Logger(ctx).Debug("Mint minterState end",
		"AmountMinted", minterState.AmountMinted.String(),
		"Position", minterState.Position,
		"RemainderFromPreviousPeriod", minterState.RemainderFromPreviousPeriod,
		"RemainderToMint", minterState.RemainderToMint,
		"LastMintBlockTime", minterState.LastMintBlockTime,
	)
	return result, nil
}

func (k Keeper) GetCurrentInflation(ctx sdk.Context) (sdk.Dec, error) {
	minterState := k.GetMinterState(ctx)
	params := k.GetParams(ctx)
	minter := params.Minter

	currentPeriod, previousPeriod := getCurrentAndPreviousPeriod(minter, &minterState)

	if currentPeriod == nil {
		return sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrNotFound, "minter current period for position %d not found", minterState.Position)

	}

	var periodStart time.Time
	if previousPeriod == nil {
		periodStart = minter.Start
	} else {
		periodStart = *previousPeriod.PeriodEnd
	}

	supply := k.bankKeeper.GetSupply(ctx, params.MintDenom)

	return currentPeriod.CalculateInfation(supply.Amount, periodStart, ctx.BlockHeader().Time), nil
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
