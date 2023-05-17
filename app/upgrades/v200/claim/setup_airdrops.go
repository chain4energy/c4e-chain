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
	TeamdropVestingPoolOwner = "c4e1dsm96gwcv35m4rqd93pzcsztpkrqe0ev7getj8"
	AirdropVestingPoolOwner  = "c4e1p0smw03cwhqn05fkalfpcr0ngqv5jrpnx2cp54"
)

func Creates(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) error {
	airdropLockupPeriod := 183 * 24 * time.Hour
	airdropVestingPeriod := 91 * 24 * time.Hour
	teamdropLockupPeriod := 730 * 24 * time.Hour
	teamdropVestingPeriod := 730 * 24 * time.Hour
	startTime := ctx.BlockTime()
	endTime := startTime.Add(time.Hour * 100)
	zeroInt := math.ZeroInt()
	zeroDec := sdk.ZeroDec()

	_, found := appKeepers.GetC4eVestingKeeper().GetAccountVestingPools(ctx, AirdropVestingPoolOwner)
	if !found {
		ctx.Logger().Info("account vesting pools not found for NewAirdropVestingPoolOwner", "owner", AirdropVestingPoolOwner)
		return nil
	}

	_, found = appKeepers.GetC4eVestingKeeper().GetAccountVestingPools(ctx, TeamdropVestingPoolOwner)
	if !found {
		ctx.Logger().Info("account vesting pools not found for TeamdropVestingPoolOwner", "owner", TeamdropVestingPoolOwner)
		return nil
	}

	_, err := appKeepers.GetC4eClaimKeeper().CreateCampaign(ctx, TeamdropVestingPoolOwner, "teamdrop", "teamdrop",
		types.VestingPoolCampaign, true, &zeroInt, &zeroInt, &zeroDec, &startTime, &endTime, &teamdropLockupPeriod, &teamdropVestingPeriod, "Teamdrop")
	if err != nil {
		return err
	}

	_, err = appKeepers.GetC4eClaimKeeper().CreateCampaign(ctx, AirdropVestingPoolOwner, "stakedrop", "stakedrop",
		types.VestingPoolCampaign, false, &zeroInt, &zeroInt, &zeroDec, &startTime, &endTime, &airdropLockupPeriod, &airdropVestingPeriod, "Fairdrop")
	if err != nil {
		return err
	}

	_, err = appKeepers.GetC4eClaimKeeper().CreateCampaign(ctx, AirdropVestingPoolOwner, "santadrop", "santadrop",
		types.VestingPoolCampaign, false, &zeroInt, &zeroInt, &zeroDec, &startTime, &endTime, &airdropLockupPeriod, &airdropVestingPeriod, "Fairdrop")
	if err != nil {
		return err
	}

	_, err = appKeepers.GetC4eClaimKeeper().CreateCampaign(ctx, AirdropVestingPoolOwner, "gleamdrop", "gleamdrop",
		types.VestingPoolCampaign, false, &zeroInt, &zeroInt, &zeroDec, &startTime, &endTime, &airdropLockupPeriod, &airdropVestingPeriod, "Fairdrop")
	if err != nil {
		return err
	}

	stakedropEntries, err := readEntriesFromJson("stakedrop.json")
	if err != nil {
		return err
	}
	if err = appKeepers.GetC4eClaimKeeper().AddClaimRecords(ctx, AirdropVestingPoolOwner, 0, stakedropEntries); err != nil {
		return err
	}

	dynamicEntries, err := readEntriesFromJson("teamdrop.json")
	if err != nil {
		return err
	}
	if err = appKeepers.GetC4eClaimKeeper().AddClaimRecords(ctx, TeamdropVestingPoolOwner, 1, dynamicEntries); err != nil {
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
