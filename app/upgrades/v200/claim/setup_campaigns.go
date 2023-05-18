package claim

import (
	"cosmossdk.io/math"
	"embed"
	"encoding/json"
	cfeupgradetypes "github.com/chain4energy/c4e-chain/app/upgrades"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

//go:embed stakedrop.json santadrop.json gleamdrop.json teamdrop.json
var f embed.FS

const (
	FairdropVestingPoolName  = "Fairdrop"
	TeamdropVestingPoolName  = "Teamdrop"
	TeamdropVestingPoolOwner = "c4e1dsm96gwcv35m4rqd93pzcsztpkrqe0ev7getj8"
	AirdropVestingPoolOwner  = "c4e1p0smw03cwhqn05fkalfpcr0ngqv5jrpnx2cp54"
)

func SetupCampaigns(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) error {
	airdropLockupPeriod := 183 * 24 * time.Hour
	airdropVestingPeriod := 91 * 24 * time.Hour
	teamdropLockupPeriod := 730 * 24 * time.Hour
	teamdropVestingPeriod := 730 * 24 * time.Hour
	startTime := ctx.BlockTime()
	endTime := startTime.Add(time.Hour * 100)
	zeroInt := math.ZeroInt()
	zeroDec := sdk.ZeroDec()
	inititalClaimOneC4E := math.NewInt(1000000)

	airdropVestingPools, found := appKeepers.GetC4eVestingKeeper().GetAccountVestingPools(ctx, AirdropVestingPoolOwner)
	if !found {
		ctx.Logger().Info("account vesting pools not found for NewAirdropVestingPoolOwner", "owner", AirdropVestingPoolOwner)
		return nil
	}
	found = false
	for _, vestingPool := range airdropVestingPools.VestingPools {
		if vestingPool.Name == FairdropVestingPoolName {
			found = true
			break
		}
	}
	if !found {
		ctx.Logger().Info("fairdrop vesting pool not found fo for NewAirdropVestingPoolOwner", "owner", AirdropVestingPoolOwner)
		return nil
	}

	teamdropVestingPools, found := appKeepers.GetC4eVestingKeeper().GetAccountVestingPools(ctx, TeamdropVestingPoolOwner)
	if !found {
		ctx.Logger().Info("account vesting pools not found for TeamdropVestingPoolOwner", "owner", TeamdropVestingPoolOwner)
		return nil
	}
	found = false
	for _, vestingPool := range teamdropVestingPools.VestingPools {
		if vestingPool.Name == TeamdropVestingPoolName {
			found = true
			break
		}
	}
	if !found {
		ctx.Logger().Info("teamdrop vesting pool not found fo for NewAirdropVestingPoolOwner", "owner", TeamdropVestingPoolOwner)
		return nil
	}

	_, err := appKeepers.GetC4eClaimKeeper().CreateCampaign(ctx, TeamdropVestingPoolOwner, "teamdrop", "teamdrop",
		types.VestingPoolCampaign, true, &zeroInt, &inititalClaimOneC4E, &zeroDec, &startTime, &endTime, &teamdropLockupPeriod, &teamdropVestingPeriod, TeamdropVestingPoolName)
	if err != nil {
		return err
	}

	_, err = appKeepers.GetC4eClaimKeeper().CreateCampaign(ctx, AirdropVestingPoolOwner, "stakedrop", "stakedrop",
		types.VestingPoolCampaign, false, &zeroInt, &inititalClaimOneC4E, &zeroDec, &startTime, &endTime, &airdropLockupPeriod, &airdropVestingPeriod, FairdropVestingPoolName)
	if err != nil {
		return err
	}

	_, err = appKeepers.GetC4eClaimKeeper().CreateCampaign(ctx, AirdropVestingPoolOwner, "santadrop", "santadrop",
		types.VestingPoolCampaign, false, &zeroInt, &inititalClaimOneC4E, &zeroDec, &startTime, &endTime, &airdropLockupPeriod, &airdropVestingPeriod, FairdropVestingPoolName)
	if err != nil {
		return err
	}

	_, err = appKeepers.GetC4eClaimKeeper().CreateCampaign(ctx, AirdropVestingPoolOwner, "gleamdrop", "gleamdrop",
		types.VestingPoolCampaign, false, &zeroInt, &inititalClaimOneC4E, &zeroDec, &startTime, &endTime, &airdropLockupPeriod, &airdropVestingPeriod, FairdropVestingPoolName)
	if err != nil {
		return err
	}

	teamdropEntries, err := readEntriesFromJson("teamdrop.json")
	if err != nil {
		return err
	}
	if err = appKeepers.GetC4eClaimKeeper().AddClaimRecords(ctx, TeamdropVestingPoolOwner, 0, teamdropEntries); err != nil {
		return err
	}

	stakedropEntries, err := readEntriesFromJson("stakedrop.json")
	if err != nil {
		return err
	}
	if err = appKeepers.GetC4eClaimKeeper().AddClaimRecords(ctx, AirdropVestingPoolOwner, 1, stakedropEntries); err != nil {
		return err
	}

	santadropEntries, err := readEntriesFromJson("santadrop.json")
	if err != nil {
		return err
	}
	if err = appKeepers.GetC4eClaimKeeper().AddClaimRecords(ctx, AirdropVestingPoolOwner, 2, santadropEntries); err != nil {
		return err
	}

	gleamdropEntries, err := readEntriesFromJson("gleamdrop.json")
	if err != nil {
		return err
	}
	if err = appKeepers.GetC4eClaimKeeper().AddClaimRecords(ctx, AirdropVestingPoolOwner, 3, gleamdropEntries); err != nil {
		return err
	}

	return nil
}

func readEntriesFromJson(fileName string) ([]*types.ClaimRecord, error) {
	data, err := f.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var claimEntires []*types.ClaimRecord
	err = json.Unmarshal(data, &claimEntires)
	if err != nil {
		return nil, err
	}
	return claimEntires, nil
}
