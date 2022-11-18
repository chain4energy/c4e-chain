package keeper

import (
	"time"

	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) Mint(ctx sdk.Context) (sdk.Int, error) {
	lastBlockTimeForMinter := k.GetMinterState(ctx).LastMintBlockTime
	lastBlockTime := ctx.BlockTime()
	params := k.GetParams(ctx)

	if lastBlockTime.Before(params.Minter.Start) {
		k.Logger(ctx).Info("minter start in the future. %d > %d", "minterStart", params.Minter.Start, "currentBlockTime", lastBlockTime)
		return sdk.ZeroInt(), nil
	}
	if lastBlockTimeForMinter.After(lastBlockTime) || lastBlockTimeForMinter.Equal(lastBlockTime) {
		k.Logger(ctx).Info("mint last mint block time is smaller than current block time - possible for first block after genesis init",
			"lastBlockTime", lastBlockTimeForMinter, "currentBlockTime", lastBlockTime)
		return sdk.ZeroInt(), nil
	}
	return k.mint(ctx, &params, 0)
}

func (k Keeper) mint(ctx sdk.Context, params *types.Params, level int) (sdk.Int, error) {
	minterState := k.GetMinterState(ctx)

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

	err = k.SendMintedCoins(ctx, coins)
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
		minted, err := k.mint(ctx, params, level+1)
		if err != nil {
			k.Logger(ctx).Error("mint - sub mint error", "lev", level, "error", err.Error())
			return minted, err
		}
		result = minted.Add(amount)
	}

	k.Logger(ctx).Debug("mint ret", "lev", level, "result", result, "minterState", minterState)
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
