package claim

import (
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	"embed"
	"encoding/json"
	cfeupgradetypes "github.com/chain4energy/c4e-chain/app/upgrades"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

//go:embed stakedrop.json santadrop.json gleamdrop.json teamdrop.json
var f embed.FS

const (
	FairdropVestingPoolName  = "Fairdrop"
	TeamdropVestingPoolName  = "Teamdrop"
	TeamdropVestingPoolOwner = "c4e1dsm96gwcv35m4rqd93pzcsztpkrqe0ev7getj8"
	AirdropVestingPoolOwner  = "c4e1p0smw03cwhqn05fkalfpcr0ngqv5jrpnx2cp54"
)

func SetupAirdrops(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) error {
	if err := validateSetupCampaigns(ctx, appKeepers); err != nil {
		ctx.Logger().Error("validateSetupCampaigns", "err", err)
		return nil
	}
	if err := setupCampaigns(ctx, appKeepers); err != nil {
		ctx.Logger().Error("validateSetupCampaigns", "err", err)
		return nil
	}
	if err := addClaimRecordsToCampaigns(ctx, appKeepers); err != nil {
		ctx.Logger().Error("validateSetupCampaigns", "err", err)
		return nil
	}
	ctx.Logger().Info("setup campaigns finished",
		"campaignsLen", len(appKeepers.GetC4eClaimKeeper().GetAllCampaigns(ctx)),
		"userEntriesLen", len(appKeepers.GetC4eClaimKeeper().GetAllUsersEntries(ctx)),
	)
	return nil
}

func setupCampaigns(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) error {
	ctx.Logger().Info("setup campaigns")

	airdropLockupPeriod := 183 * 24 * time.Hour
	airdropVestingPeriod := 91 * 24 * time.Hour
	teamdropLockupPeriod := 730 * 24 * time.Hour
	teamdropVestingPeriod := 730 * 24 * time.Hour
	startTime := ctx.BlockTime()
	endTime := startTime.Add(time.Hour * 100)
	zeroInt := math.ZeroInt()
	zeroDec := sdk.ZeroDec()
	inititalClaimOneC4E := math.NewInt(1000000)
	_, err := appKeepers.GetC4eClaimKeeper().CreateCampaign(ctx, TeamdropVestingPoolOwner, "teamdrop", "teamdrop",
		types.VestingPoolCampaign, true, zeroInt, inititalClaimOneC4E, zeroDec, startTime, endTime, teamdropLockupPeriod, teamdropVestingPeriod, TeamdropVestingPoolName)
	if err != nil {
		return err
	}

	_, err = appKeepers.GetC4eClaimKeeper().CreateCampaign(ctx, AirdropVestingPoolOwner, "stakedrop", "stakedrop",
		types.VestingPoolCampaign, false, zeroInt, inititalClaimOneC4E, zeroDec, startTime, endTime, airdropLockupPeriod, airdropVestingPeriod, FairdropVestingPoolName)
	if err != nil {
		return err
	}

	_, err = appKeepers.GetC4eClaimKeeper().CreateCampaign(ctx, AirdropVestingPoolOwner, "santadrop", "santadrop",
		types.VestingPoolCampaign, false, zeroInt, inititalClaimOneC4E, zeroDec, startTime, endTime, airdropLockupPeriod, airdropVestingPeriod, FairdropVestingPoolName)
	if err != nil {
		return err
	}

	_, err = appKeepers.GetC4eClaimKeeper().CreateCampaign(ctx, AirdropVestingPoolOwner, "gleamdrop", "gleamdrop",
		types.VestingPoolCampaign, false, zeroInt, inititalClaimOneC4E, zeroDec, startTime, endTime, airdropLockupPeriod, airdropVestingPeriod, FairdropVestingPoolName)

	return err
}

func validateSetupCampaigns(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) error {
	airdropVestingPools, found := appKeepers.GetC4eVestingKeeper().GetAccountVestingPools(ctx, AirdropVestingPoolOwner)
	if !found {
		return errors.Wrapf(sdkerrors.ErrNotFound, "account vesting pools not found for NewAirdropVestingPoolOwner %s", AirdropVestingPoolOwner)
	}
	found = false
	for _, vestingPool := range airdropVestingPools.VestingPools {
		if vestingPool.Name == FairdropVestingPoolName {
			found = true
			break
		}
	}
	if !found {
		return errors.Wrapf(sdkerrors.ErrNotFound, "fairdrop vesting pool not found for NewAirdropVestingPoolOwner %s", AirdropVestingPoolOwner)
	}

	teamdropVestingPools, found := appKeepers.GetC4eVestingKeeper().GetAccountVestingPools(ctx, TeamdropVestingPoolOwner)
	if !found {
		return errors.Wrapf(sdkerrors.ErrNotFound, "account vesting pools not found for TeamdropVestingPoolOwner %s", TeamdropVestingPoolOwner)
	}
	found = false
	for _, vestingPool := range teamdropVestingPools.VestingPools {
		if vestingPool.Name == TeamdropVestingPoolName {
			found = true
			break
		}
	}
	if !found {
		return errors.Wrapf(sdkerrors.ErrNotFound, "teamdrop vesting pool not found fo for TeamdropVestingPoolOwner %s", TeamdropVestingPoolOwner)
	}
	return nil
}

func addClaimRecordsToCampaigns(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) error {
	teamdropEntries, err := readClaimRecordEntriesFromJson("teamdrop.json")
	if err != nil {
		return err
	}
	if err = appKeepers.GetC4eClaimKeeper().AddClaimRecords(ctx, TeamdropVestingPoolOwner, 0, teamdropEntries); err != nil {
		return err
	}

	stakedropEntries, err := readClaimRecordEntriesFromJson("stakedrop.json")
	if err != nil {
		return err
	}
	if err = appKeepers.GetC4eClaimKeeper().AddClaimRecords(ctx, AirdropVestingPoolOwner, 1, stakedropEntries); err != nil {
		return err
	}

	santadropEntries, err := readClaimRecordEntriesFromJson("santadrop.json")
	if err != nil {
		return err
	}
	if err = appKeepers.GetC4eClaimKeeper().AddClaimRecords(ctx, AirdropVestingPoolOwner, 2, santadropEntries); err != nil {
		return err
	}

	gleamdropEntries, err := readClaimRecordEntriesFromJson("gleamdrop.json")
	if err != nil {
		return err
	}
	return appKeepers.GetC4eClaimKeeper().AddClaimRecords(ctx, AirdropVestingPoolOwner, 3, gleamdropEntries)
}

func readClaimRecordEntriesFromJson(fileName string) ([]*types.ClaimRecordEntry, error) {
	data, err := f.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var claimRecordEntries []*types.ClaimRecordEntry
	err = json.Unmarshal(data, &claimRecordEntries)
	if err != nil {
		return nil, err
	}
	return claimRecordEntries, nil
}
