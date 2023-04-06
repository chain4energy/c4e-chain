package claim

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	cfeclaimkeeper "github.com/chain4energy/c4e-chain/x/cfeclaim/keeper"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
)

//go:embed stakedrop.json santadrop.json gleamdrop.json teamdrop.json
var f embed.FS

const monthAvgHours = 365 * 24 / 12 * time.Hour
const claimSource = "cfeminter"

func Creates(ctx sdk.Context, claimKeeper *cfeclaimkeeper.Keeper, accountKeeper *authkeeper.AccountKeeper, bankKeeper *bankkeeper.Keeper) error {
	lockupPeriod := 3 * monthAvgHours
	vestingPeriod := 6 * monthAvgHours
	startTime := time.Now().Add(time.Hour * 2)
	endTime := startTime.Add(time.Hour * 100)
	acc := accountKeeper.GetModuleAccount(ctx, claimSource)
	if acc == nil {
		claimKeeper.Logger(ctx).Error("source module account not found", "name", claimSource)
		return fmt.Errorf("source module account not found: %s", claimSource)
	}
	ownerAcc := acc.GetAddress().String()
	err := bankkeeper.Keeper.MintCoins(*bankKeeper, ctx, claimSource, sdk.NewCoins(sdk.NewCoin("uc4e", sdk.NewInt(100000000000000))))
	if err != nil {
		return err
	}
	zeroInt := sdk.ZeroInt()
	err = claimKeeper.CreateCampaign(ctx, acc.GetAddress().String(), "stakedrop", "stakedrop", types.CampaignDefault, &zeroInt, &zeroInt, &startTime, &endTime, &lockupPeriod, &vestingPeriod)
	if err != nil {
		return err
	}
	// TODO: change campaign type to teamdrop
	err = claimKeeper.CreateCampaign(ctx, acc.GetAddress().String(), "teamdrop", "teamdrop", types.CampaignDefault, &zeroInt, &zeroInt, &startTime, &endTime, &lockupPeriod, &vestingPeriod)
	if err != nil {
		return err
	}
	err = claimKeeper.CreateCampaign(ctx, acc.GetAddress().String(), "santadrop", "santadrop", types.CampaignDefault, &zeroInt, &zeroInt, &startTime, &endTime, &lockupPeriod, &vestingPeriod)
	if err != nil {
		return err
	}
	err = claimKeeper.CreateCampaign(ctx, acc.GetAddress().String(), "gleamdrop", "gleamdrop", types.CampaignDefault, &zeroInt, &zeroInt, &startTime, &endTime, &lockupPeriod, &vestingPeriod)
	if err != nil {
		return err
	}

	if err = claimKeeper.StartCampaign(ctx, ownerAcc, 0); err != nil {
		return err
	}
	if err = claimKeeper.StartCampaign(ctx, ownerAcc, 1); err != nil {
		return err
	}
	if err = claimKeeper.StartCampaign(ctx, ownerAcc, 2); err != nil {
		return err
	}
	if err = claimKeeper.StartCampaign(ctx, ownerAcc, 3); err != nil {
		return err
	}

	stakedropEntries, err := readEntriesFromJson("stakedrop.json")
	if err != nil {
		return err
	}
	if err = claimKeeper.AddUsersEntries(ctx, ownerAcc, 0, stakedropEntries); err != nil {
		return err
	}

	teamdropEntries, err := readEntriesFromJson("teamdrop.json")
	if err != nil {
		return err
	}
	if err = claimKeeper.AddUsersEntries(ctx, ownerAcc, 1, teamdropEntries); err != nil {
		return err
	}

	santadropEntries, err := readEntriesFromJson("santadrop.json")
	if err != nil {
		return err
	}
	if err = claimKeeper.AddUsersEntries(ctx, ownerAcc, 2, santadropEntries); err != nil {
		return err
	}

	gleamdropEntries, err := readEntriesFromJson("gleamdrop.json")
	if err != nil {
		return err
	}
	if err = claimKeeper.AddUsersEntries(ctx, ownerAcc, 3, gleamdropEntries); err != nil {
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
