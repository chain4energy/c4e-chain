package keeper

import (
	"cosmossdk.io/math"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) Mint(ctx sdk.Context) (math.Int, error) {
	lastBlockTimeForMinter := k.GetMinterState(ctx).LastMintBlockTime
	lastBlockTime := ctx.BlockTime()
	params := k.GetParams(ctx)

	if lastBlockTime.Before(params.StartTime) {
		k.Logger(ctx).Info("minter start in the future", "minterStart", params.StartTime, "currentBlockTime", lastBlockTime)
		return sdk.ZeroInt(), nil
	}
	if lastBlockTimeForMinter.After(lastBlockTime) || lastBlockTimeForMinter.Equal(lastBlockTime) {
		k.Logger(ctx).Info("mint last mint block time is smaller than current block time - possible for first block after genesis init",
			"lastBlockTime", lastBlockTimeForMinter, "currentBlockTime", lastBlockTime)
		return sdk.ZeroInt(), nil
	}
	return k.mint(ctx, &params, 0)
}

func (k Keeper) mint(ctx sdk.Context, params *types.Params, level int) (math.Int, error) {
	minterState := k.GetMinterState(ctx)
	currentMinter, previousMinter := getCurrentAndPreviousMinter(params.Minters, &minterState)

	if currentMinter == nil {
		k.Logger(ctx).Error("mint - current minter not found error", "lev", level, "SequenceId", minterState.SequenceId)
		return sdk.ZeroInt(), sdkerrors.Wrapf(sdkerrors.ErrNotFound, "minter - mint - current minter for sequence id %d not found", minterState.SequenceId)
	}

	var startTime time.Time
	if previousMinter == nil {
		startTime = params.StartTime
	} else {
		startTime = *previousMinter.EndTime
	}

	expectedAmountToMint := currentMinter.AmountToMint(k.Logger(ctx), startTime, ctx.BlockTime())
	expectedAmountToMint = expectedAmountToMint.Add(minterState.RemainderFromPreviousMinter)

	amount := expectedAmountToMint.TruncateInt().Sub(minterState.AmountMinted)
	if amount.IsNegative() {
		k.Logger(ctx).Error("mint negative amount", "lev", level, "minterState", minterState, "startTime", startTime, "currentMinter", currentMinter,
			"previousMinter", previousMinter, "expectedAmountToMint", expectedAmountToMint, "amount", amount)
		return sdk.ZeroInt(), nil
	}
	k.Logger(ctx).Debug("mint", "lev", level, "minterState", minterState, "startTime", startTime, "currentMinter", currentMinter,
		"previousMinter", previousMinter, "expectedAmountToMint", expectedAmountToMint, "amount", amount)

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

	var result math.Int
	if currentMinter.EndTime == nil || ctx.BlockTime().Before(*currentMinter.EndTime) {
		k.SetMinterState(ctx, minterState)
		result = amount
	} else {
		k.SetMinterStateHistory(ctx, minterState)
		k.Logger(ctx).Debug("mint - set minter state history", "lev", level, "minterState", minterState.String())
		minterState = types.MinterState{
			SequenceId:                  minterState.SequenceId + 1,
			AmountMinted:                sdk.ZeroInt(),
			RemainderToMint:             sdk.ZeroDec(),
			RemainderFromPreviousMinter: remainder,
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

func getCurrentAndPreviousMinter(minters []*types.Minter, state *types.MinterState) (currentMinter *types.Minter, previousMinter *types.Minter) {
	currentId := state.SequenceId
	for _, minter := range minters {
		if minter.SequenceId == currentId {
			currentMinter = minter
		}
		if previousMinter == nil {
			if minter.SequenceId < currentId {
				previousMinter = minter
			}
		} else {
			if minter.SequenceId < currentId && minter.SequenceId > previousMinter.SequenceId {
				previousMinter = minter
			}
		}
	}
	return currentMinter, previousMinter
}
