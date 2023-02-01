package airdrop

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	cfeairdropkeeper "github.com/chain4energy/c4e-chain/x/cfeairdrop/keeper"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
)

//go:embed stakedrop.json santadrop.json gleamdrop.json teamdrop.json
var f embed.FS

const monthAvgHours = 365 * 24 / 12 * time.Hour
const airdropSource = "cfeminter"

func CreateAirdrops(ctx sdk.Context, airdropKeeper *cfeairdropkeeper.Keeper, accountKeeper *authkeeper.AccountKeeper, bankKeeper *bankkeeper.Keeper) error {
	lockupPeriod := 3 * monthAvgHours
	vestingPeriod := 6 * monthAvgHours
	startTime := time.Now().Add(time.Hour * 2)
	endTime := startTime.Add(time.Hour * 100)
	acc := accountKeeper.GetModuleAccount(ctx, airdropSource)
	if acc == nil {
		airdropKeeper.Logger(ctx).Error("source module account not found", "name", airdropSource)
		return fmt.Errorf("source module account not found: %s", airdropSource)
	}
	ownerAcc := acc.GetAddress().String()
	err := bankkeeper.Keeper.MintCoins(*bankKeeper, ctx, airdropSource, sdk.NewCoins(sdk.NewCoin("uc4e", sdk.NewInt(100000000000000))))
	if err != nil {
		return err
	}
	zeroInt := sdk.ZeroInt()
	err = airdropKeeper.CreateAidropCampaign(ctx, acc.GetAddress().String(), "stakedrop", "stakedrop", &zeroInt, &zeroInt, &startTime, &endTime, &lockupPeriod, &vestingPeriod)
	if err != nil {
		return err
	}
	err = airdropKeeper.CreateAidropCampaign(ctx, acc.GetAddress().String(), "teamdrop", "teamdrop", &zeroInt, &zeroInt, &startTime, &endTime, &lockupPeriod, &vestingPeriod)
	if err != nil {
		return err
	}
	err = airdropKeeper.CreateAidropCampaign(ctx, acc.GetAddress().String(), "santadrop", "santadrop", &zeroInt, &zeroInt, &startTime, &endTime, &lockupPeriod, &vestingPeriod)
	if err != nil {
		return err
	}
	err = airdropKeeper.CreateAidropCampaign(ctx, acc.GetAddress().String(), "gleamdrop", "gleamdrop", &zeroInt, &zeroInt, &startTime, &endTime, &lockupPeriod, &vestingPeriod)
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

	stakedropAirdropEntries, err := readAirdropEntriesFromJson("stakedrop.json")
	if err != nil {
		return err
	}
	if err = airdropKeeper.AddUsersEntries(ctx, ownerAcc, 0, stakedropAirdropEntries); err != nil {
		return err
	}

	teamdropAirdropEntries, err := readAirdropEntriesFromJson("teamdrop.json")
	if err != nil {
		return err
	}
	if err = airdropKeeper.AddUsersEntries(ctx, ownerAcc, 1, teamdropAirdropEntries); err != nil {
		return err
	}

	santadropAirdropEntries, err := readAirdropEntriesFromJson("santadrop.json")
	if err != nil {
		return err
	}
	if err = airdropKeeper.AddUsersEntries(ctx, ownerAcc, 2, santadropAirdropEntries); err != nil {
		return err
	}

	gleamdropAirdropEntries, err := readAirdropEntriesFromJson("gleamdrop.json")
	if err != nil {
		return err
	}
	if err = airdropKeeper.AddUsersEntries(ctx, ownerAcc, 3, gleamdropAirdropEntries); err != nil {
		return err
	}

	return nil
}

func readAirdropEntriesFromJson(fileName string) ([]*types.ClaimRecord, error) {
	data, err := f.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var airdropEntires []*types.ClaimRecord
	err = json.Unmarshal(data, &airdropEntires)
	if err != nil {
		return nil, err
	}
	return airdropEntires, nil
}
