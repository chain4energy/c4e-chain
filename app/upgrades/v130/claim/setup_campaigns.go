package claim

import (
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	"embed"
	"encoding/json"
	"fmt"
	cfeupgradetypes "github.com/chain4energy/c4e-chain/app/upgrades"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

//go:embed stakedrop.json santadrop.json greendrop.json moondrop.json zealydrop.json amadrop.json
var f embed.FS

const (
	FairdropVestingPoolName  = "Fairdrop"
	MoondropVestingPoolName  = "Moondrop"
	MoondropVestingPoolOwner = "c4e1dsm96gwcv35m4rqd93pzcsztpkrqe0ev7getj8"
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

const (
	stakedropName        = "Stake Drop"
	stakedropDescription = "Stake Drop is the airdrop aimed to spread knowledge about the C4E ecosystem among the Cosmos $ATOM " +
		"stakers community. The airdrop snapshot has been taken on September 28th, 2022 at 9:30 PM " +
		"UTC (during the ATOM 2.0 roadmap announcement at the Cosmoverse Conference."

	moondropName = "Moon Drop"

	santadropName        = "Santa Drop"
	santadropDescription = "Santa Drop prize pool for was 10.000 C4E Tokens, with 10 lucky winners getting 1000 tokens per each. The participants had to complete the tasks to get a chance to be among lucky winners."

	greendropName        = "Green Drop"
	greendropDescription = "It was the first airdrop competition aimed at spreading knowledge about the C4E ecosystem. The Prize Pool was 1.000.000 C4E tokens and what is best — all the participants who completed the tasks are eligible for the c4e tokens from it!"

	zealydropName        = "Incentived Testnet I"
	zealydropDescription = "Incentivized Testnet Zealy campaign, is innovative approach designed to foster engagement and bolster network security. Community members are rewarded for participating in testnet and marketing tasks, receiving C4E tokens as a result of their contributions."

	amadropName        = "AMA Drop"
	amadropDescription = "Have you been active at our AMA sessions and won C4E prizes? This Drop belongs to you."
)

var (
	airdropStartTime      = time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC)
	airdropEndTime        = time.Date(2031, 1, 1, 0, 0, 0, 0, time.UTC)
	airdropLockupPeriod   = 183 * 24 * time.Hour
	airdropVestingPeriod  = 91 * 24 * time.Hour
	moondropLockupPeriod  = 730 * 24 * time.Hour
	moondropVestingPeriod = 730 * 24 * time.Hour
)

func setupCampaigns(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) error {
	ctx.Logger().Info("setup campaigns")

	zeroInt := math.ZeroInt()
	onePercentDec := sdk.MustNewDecFromStr("0.01")

	_, err := appKeepers.GetC4eClaimKeeper().CreateCampaign(ctx, MoondropVestingPoolOwner, moondropName, "",
		types.VestingPoolCampaign, true, zeroInt, zeroInt, sdk.ZeroDec(), airdropStartTime, airdropEndTime, moondropLockupPeriod, moondropVestingPeriod, MoondropVestingPoolName)
	if err != nil {
		return err
	}

	_, err = appKeepers.GetC4eClaimKeeper().CreateCampaign(ctx, AirdropVestingPoolOwner, stakedropName, stakedropDescription,
		types.VestingPoolCampaign, false, zeroInt, zeroInt, onePercentDec, airdropStartTime, airdropEndTime, airdropLockupPeriod, airdropVestingPeriod, FairdropVestingPoolName)
	if err != nil {
		return err
	}

	_, err = appKeepers.GetC4eClaimKeeper().CreateCampaign(ctx, AirdropVestingPoolOwner, santadropName, santadropDescription,
		types.VestingPoolCampaign, false, zeroInt, zeroInt, onePercentDec, airdropStartTime, airdropEndTime, airdropLockupPeriod, airdropVestingPeriod, FairdropVestingPoolName)
	if err != nil {
		return err
	}

	_, err = appKeepers.GetC4eClaimKeeper().CreateCampaign(ctx, AirdropVestingPoolOwner, greendropName, greendropDescription,
		types.VestingPoolCampaign, false, zeroInt, zeroInt, onePercentDec, airdropStartTime, airdropEndTime, airdropLockupPeriod, airdropVestingPeriod, FairdropVestingPoolName)
	if err != nil {
		return err
	}

	_, err = appKeepers.GetC4eClaimKeeper().CreateCampaign(ctx, AirdropVestingPoolOwner, zealydropName, zealydropDescription,
		types.VestingPoolCampaign, false, zeroInt, zeroInt, onePercentDec, airdropStartTime, airdropEndTime, airdropLockupPeriod, airdropVestingPeriod, FairdropVestingPoolName)
	if err != nil {
		return err
	}

	_, err = appKeepers.GetC4eClaimKeeper().CreateCampaign(ctx, AirdropVestingPoolOwner, amadropName, amadropDescription,
		types.VestingPoolCampaign, false, zeroInt, zeroInt, onePercentDec, airdropStartTime, airdropEndTime, airdropLockupPeriod, airdropVestingPeriod, FairdropVestingPoolName)

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

	moondropVestingPools, found := appKeepers.GetC4eVestingKeeper().GetAccountVestingPools(ctx, MoondropVestingPoolOwner)
	if !found {
		return errors.Wrapf(sdkerrors.ErrNotFound, "account vesting pools not found for MoondropVestingPoolOwner %s", MoondropVestingPoolOwner)
	}
	found = false
	for _, vestingPool := range moondropVestingPools.VestingPools {
		if vestingPool.Name == MoondropVestingPoolName {
			found = true
			break
		}
	}
	if !found {
		return errors.Wrapf(sdkerrors.ErrNotFound, "moondrop vesting pool not found fo for MoondropVestingPoolOwner %s", MoondropVestingPoolOwner)
	}
	return nil
}

func addClaimRecordsToCampaigns(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) error {
	moondropEntries, err := readClaimRecordEntriesFromJson("moondrop.json")
	if err != nil {
		return err
	}
	if err = appKeepers.GetC4eClaimKeeper().AddClaimRecords(ctx, MoondropVestingPoolOwner, 0, moondropEntries); err != nil {
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

	greendropEntries, err := readClaimRecordEntriesFromJson("greendrop.json")
	if err != nil {
		return err
	}
	if err = appKeepers.GetC4eClaimKeeper().AddClaimRecords(ctx, AirdropVestingPoolOwner, 3, greendropEntries); err != nil {
		return err
	}

	zealaydropEntries, err := readClaimRecordEntriesFromJson("zealydrop.json")
	foundEntries := 0
	totalEntries := len(zealaydropEntries)
	fmt.Println("totalEntries: ", totalEntries)
	for _, entry := range zealaydropEntries {
		_, found := appKeepers.GetC4eClaimKeeper().GetUserEntry(ctx, entry.UserEntryAddress)
		if found {
			foundEntries++
		}
	}
	fmt.Print("foundEntries: ", foundEntries, "\n")
	if err != nil {
		return err
	}
	if err = appKeepers.GetC4eClaimKeeper().AddClaimRecords(ctx, AirdropVestingPoolOwner, 4, zealaydropEntries); err != nil {
		return err
	}

	amadropEntries, err := readClaimRecordEntriesFromJson("amadrop.json")
	if err != nil {
		return err
	}

	return appKeepers.GetC4eClaimKeeper().AddClaimRecords(ctx, AirdropVestingPoolOwner, 5, amadropEntries)
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
