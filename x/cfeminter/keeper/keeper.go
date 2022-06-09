package keeper

import (
	"fmt"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
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

		bankKeeper    types.BankKeeper
		collectorName string
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,

	bankKeeper types.BankKeeper,
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
		collectorName: collectorName,
	}
}

func (k Keeper) GetCollectorName() string {
	return k.collectorName
}

// get the minter
func (k Keeper) GetHalvingMinter(ctx sdk.Context) (minter types.HalvingMinter) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.HalvingMinterKey)
	if b == nil {
		panic("stored minter should not have been nil")
	}

	k.cdc.MustUnmarshal(b, &minter)
	return
}

// set the minter
func (k Keeper) SetHalvingMinter(ctx sdk.Context, minter types.HalvingMinter) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&minter)
	store.Set(types.HalvingMinterKey, b)
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) Mint(ctx sdk.Context) (sdk.Int, error) {
	minter := k.GetMinter(ctx)
	minterState := k.GetMinterState(ctx)
	params := k.GetParams(ctx)

	currentId := minterState.CurrentOrderingId
	var currentPeriod *types.MintingPeriod = nil
	var previousPeriod *types.MintingPeriod = nil
	for _, period := range minter.Periods {
		if period.OrderingId == currentId {
			currentPeriod = period
		}
		if previousPeriod == nil {
			if period.OrderingId < currentId {
				previousPeriod = period
			}
		} else {
			if period.OrderingId < currentId && period.OrderingId > previousPeriod.OrderingId {
				previousPeriod = period
			}
		}
	}

	if currentPeriod == nil {
		//  TODO return error
		return sdk.ZeroInt(), nil
	}

	var periodStart time.Time
	if previousPeriod == nil {
		periodStart = minter.Start
	} else {
		periodStart = *previousPeriod.PeriodEnd
	}

	amount := currentPeriod.AmountToMint(&minterState, periodStart, ctx.BlockTime())

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

	if currentPeriod.PeriodEnd == nil || ctx.BlockTime().Before(*currentPeriod.PeriodEnd) {
		minterState.AmountMinted = minterState.AmountMinted.Add(amount)
		k.SetMinterState(ctx, minterState)
		return amount, nil
	} else {
		minterState.CurrentOrderingId++
		minterState.AmountMinted = sdk.ZeroInt()
		k.SetMinterState(ctx, minterState)
		minted, err := k.Mint(ctx)
		if err != nil {
			return minted, err
		}
		return minted.Add(amount), nil
	}
}

// MintCoins implements an alias call to the underlying supply keeper's
// MintCoins to be used in BeginBlocker.
func (k Keeper) MintCoins(ctx sdk.Context, newCoins sdk.Coins) error {
	if newCoins.Empty() {
		// skip as no coins need to be minted
		return nil
	}
	return k.bankKeeper.MintCoins(ctx, types.ModuleName, newCoins)

	//k.bankKeeper.
}

// AddCollectedFees implements an alias call to the underlying supply keeper's
// AddCollectedFees to be used in BeginBlocker.
func (k Keeper) AddCollectedFees(ctx sdk.Context, fees sdk.Coins) error {
	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.collectorName, fees)
}

// func (k Keeper) SendCoinsToCommonAccount(ctx sdk.Context, coins sdk.Coins) error {
// 	k.Logger(ctx).Info("SendCoinsToCommonAccount")
// 	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.InflationCollectorName, coins)

// }

// func (k Keeper) MintSomeCoinEndSendToTest(ctx sdk.Context) {
// 	k.bankKeeper.MintCoins(ctx, "fee_collector", sdk.NewCoins(sdk.NewCoin("uc4e", sdk.NewInt(50))))
// 	k.bankKeeper.MintCoins(ctx, "payment_collector", sdk.NewCoins(sdk.NewCoin("uc4e", sdk.NewInt(30))))
// }
