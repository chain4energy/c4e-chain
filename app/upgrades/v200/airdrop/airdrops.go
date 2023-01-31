package airdrop

import (
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	cfeairdropkeeper "github.com/chain4energy/c4e-chain/x/cfeairdrop/keeper"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
)

const monthAvgHours = 365 * 24 / 12 * time.Hour
const airdropSource = "fairdrop"

func CreateAirdrops(ctx sdk.Context, airdropKeeper *cfeairdropkeeper.Keeper, accountKeeper *authkeeper.AccountKeeper, bankKeeper *bankkeeper.Keeper) error {
	lockupPeriod := 3 * monthAvgHours
	vestingPeriod := 6 * monthAvgHours
	startTime := time.Now().Add(time.Hour * 100)
	endTime := startTime.Add(time.Hour * 100)
	acc := accountKeeper.GetModuleAccount(ctx, airdropSource)
	ownerAcc := acc.GetAddress().String()
	err := bankkeeper.Keeper.MintCoins(*bankKeeper, ctx, "cfeairdrop", sdk.NewCoins(sdk.NewCoin("uc4e", sdk.NewInt(100000000000))))
	if err != nil {
		return err
	}
	err = airdropKeeper.CreateAidropCampaign(ctx, acc.GetAddress().String(), "stakedrop", "stakedrop", nil, nil, &startTime, &endTime, &lockupPeriod, &vestingPeriod)
	if err != nil {
		return err
	}
	err = airdropKeeper.CreateAidropCampaign(ctx, acc.GetAddress().String(), "teamdrop", "teamdrop", nil, nil, &startTime, &endTime, &lockupPeriod, &vestingPeriod)
	if err != nil {
		return err
	}
	err = airdropKeeper.CreateAidropCampaign(ctx, acc.GetAddress().String(), "santadrop", "santadrop", nil, nil, &startTime, &endTime, &lockupPeriod, &vestingPeriod)
	if err != nil {
		return err
	}
	err = airdropKeeper.CreateAidropCampaign(ctx, acc.GetAddress().String(), "gleamdrop", "gleamdrop", nil, nil, &startTime, &endTime, &lockupPeriod, &vestingPeriod)
	if err != nil {
		return err
	}

	if err = airdropKeeper.StartAirdropCampaign(ctx, ownerAcc, 0); err != nil {
		return err
	}
	if err = airdropKeeper.StartAirdropCampaign(ctx, ownerAcc, 1); err != nil {
		return err
	}
	if err = airdropKeeper.StartAirdropCampaign(ctx, ownerAcc, 2); err != nil {
		return err
	}
	if err = airdropKeeper.StartAirdropCampaign(ctx, ownerAcc, 3); err != nil {
		return err
	}

	if err = airdropKeeper.AddUserAirdropEntries(ctx, ownerAcc, 0, teamdropAirdropEntries); err != nil {
		return err
	}
	if err = airdropKeeper.AddUserAirdropEntries(ctx, ownerAcc, 1, teamdropAirdropEntries); err != nil {
		return err
	}
	if err = airdropKeeper.AddUserAirdropEntries(ctx, ownerAcc, 2, santadropAirdropEntries); err != nil {
		return err
	}
	if err = airdropKeeper.AddUserAirdropEntries(ctx, ownerAcc, 1, teamdropAirdropEntries); err != nil {
		return err
	}

	return nil
}
