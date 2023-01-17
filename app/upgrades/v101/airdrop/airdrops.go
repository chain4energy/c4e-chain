package airdrop

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	cfeairdropkeeper "github.com/chain4energy/c4e-chain/x/cfeairdrop/keeper"
	cfeairdroptypes "github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
)

const monthAvgHours = 365 * 24 / 12 * time.Hour
const airdropSource = "fairdrop"

func CreateAirdrops(ctx sdk.Context, airdropKeeper *cfeairdropkeeper.Keeper, accountKeeper *authkeeper.AccountKeeper) error {
	lockupPeriod := 3 * monthAvgHours
	vestingPeriod := 6 * monthAvgHours
	gleamCamapaignId := uint64(1)
	stakeCamapaignId := uint64(2)

	gleamStart, err := time.Parse(time.RFC3339Nano, "2022-11-17T18:21:58.952129766Z")
	if err != nil {
		airdropKeeper.Logger(ctx).Error("error parsing gleam start time", "error", err)
		return err
	}
	gleamEnd, err := time.Parse(time.RFC3339Nano, "2022-11-17T18:21:58.952129766Z")
	if err != nil {
		airdropKeeper.Logger(ctx).Error("error parsing gleam end time", "error", err)
		return err
	}

	stakeStart, err := time.Parse(time.RFC3339Nano, "2022-11-17T18:21:58.952129766Z")
	if err != nil {
		airdropKeeper.Logger(ctx).Error("error parsing gleam start time", "error", err)
		return err
	}
	stakeEnd, err := time.Parse(time.RFC3339Nano, "2022-11-17T18:21:58.952129766Z")
	if err != nil {
		airdropKeeper.Logger(ctx).Error("error parsing gleam end time", "error", err)
		return err
	}

	gleamCampaign := cfeairdroptypes.Campaign{
		Id:            gleamCamapaignId,
		Enabled:       true,
		StartTime:     &gleamStart,
		EndTime:       &gleamEnd,
		LockupPeriod:  lockupPeriod,
		VestingPeriod: vestingPeriod,
		Description:   "Gleam contest airdrop ??????", // TODO description + add name to campaign
	}

	stakeCampaign := cfeairdroptypes.Campaign{
		Id:            stakeCamapaignId,
		Enabled:       true,
		StartTime:     &stakeStart,
		EndTime:       &stakeEnd,
		LockupPeriod:  lockupPeriod,
		VestingPeriod: vestingPeriod,
		Description:   "ATOM stake airdrop ??????", // TODO description + add name to campaign
	}
	airdropKeeper.SetCampaign(ctx, gleamCampaign)
	airdropKeeper.SetCampaign(ctx, stakeCampaign)

	// Gleam contests missions
	weight := sdk.NewDec(1)
	airdropKeeper.SetMission(
		ctx,
		cfeairdroptypes.Mission{
			Id:          0,
			CampaignId:  gleamCamapaignId,
			Weight:      &weight,
			Description: "Claim gleam contest airdrop", // TODO description ??
		},
	)
	//airdropKeeper.SetInitialClaim(ctx, cfeairdroptypes.InitialClaim{CampaignId: gleamCamapaignId, MissionId: 0})

	// ATOM stakers missions
	weight = sdk.MustNewDecFromStr("0.2")
	airdropKeeper.SetMission(
		ctx,
		cfeairdroptypes.Mission{
			CampaignId:  stakeCamapaignId,
			Id:          0,
			Weight:      &weight,
			Description: "Claim initial stakers airdrop", // TODO description ??
		},
	)

	weight = sdk.MustNewDecFromStr("0.4")
	airdropKeeper.SetMission(
		ctx,
		cfeairdroptypes.Mission{
			CampaignId:  stakeCamapaignId,
			Id:          1,
			Weight:      &weight,
			Description: "Claim delegtion stakers airdrop", // TODO description ??
		},
	)

	airdropKeeper.SetMission(
		ctx,
		cfeairdroptypes.Mission{
			CampaignId:  stakeCamapaignId,
			Id:          2,
			Weight:      &weight,
			Description: "Claim voting stakers airdrop", // TODO description ??
		},
	)

	//airdropKeeper.SetInitialClaim(ctx, cfeairdroptypes.InitialClaim{CampaignId: stakeCamapaignId, MissionId: 0})

	acc := accountKeeper.GetModuleAccount(ctx, airdropSource)
	if acc == nil {
		airdropKeeper.Logger(ctx).Error("source module account not found", "name", airdropSource)
		return fmt.Errorf("source module account not found: %s", airdropSource)
	}
	if err = airdropKeeper.AddUserAirdropEntries(ctx, acc.GetAddress().String(), gleamCamapaignId, gleamContestRecords); err != nil {
		return err
	}
	if err = airdropKeeper.AddUserAirdropEntries(ctx, acc.GetAddress().String(), stakeCamapaignId, stakeRecords); err != nil {
		return err
	}
	return nil
}
