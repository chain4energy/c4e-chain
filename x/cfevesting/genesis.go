package cfevesting

import (
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type PeriodUnit string

const (
	Day               = "day"
	Hour              = "hour"
	Minute            = "minute"
	Second            = "second"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState, ak types.AccountKeeper) {

	k.Logger(ctx).Info("Init genesis")
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	k.Logger(ctx).Info("Init genesis params: ")
	vestingTypes := types.VestingTypes{}
	for _, gVestingType := range genState.VestingTypes {
		vt := types.VestingType{
			Name:                 gVestingType.Name,
			LockupPeriod:         DurationFromUnits(PeriodUnit(gVestingType.LockupPeriodUnit), gVestingType.LockupPeriod),
			VestingPeriod:        DurationFromUnits(PeriodUnit(gVestingType.VestingPeriodUnit), gVestingType.VestingPeriod),
			TokenReleasingPeriod: DurationFromUnits(PeriodUnit(gVestingType.TokenReleasingPeriodUnit), gVestingType.TokenReleasingPeriod),
			DelegationsAllowed:   gVestingType.DelegationsAllowed,
		}
		vestingTypes.VestingTypes = append(vestingTypes.VestingTypes, &vt)
	}

	k.SetVestingTypes(ctx, vestingTypes)
	allVestings := genState.AccountVestingsList.Vestings

	for _, av := range allVestings {
		k.SetAccountVestings(ctx, *av)
	}
	ak.GetModuleAccount(ctx, types.ModuleName)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	vestingTypes := k.GetVestingTypes(ctx)
	gVestingTypes := []types.GenesisVestingType{}

	for _, vestingType := range vestingTypes.VestingTypes {
		lockupPeriodUnit, lockupPeriod := UnitsFromDuration(vestingType.LockupPeriod)
		vestingPeriodUnit, vestingPeriod := UnitsFromDuration(vestingType.VestingPeriod)
		tokenReleasingPeriodUnit, tokenReleasingPeriod := UnitsFromDuration(vestingType.TokenReleasingPeriod)

		gvt := types.GenesisVestingType{
			Name:                 vestingType.Name,
			LockupPeriod:         lockupPeriod,
			LockupPeriodUnit: 	  string(lockupPeriodUnit),
			VestingPeriod:        vestingPeriod,
			VestingPeriodUnit:    string(vestingPeriodUnit),
			TokenReleasingPeriod: tokenReleasingPeriod,
			TokenReleasingPeriodUnit: string(tokenReleasingPeriodUnit),
			DelegationsAllowed:   vestingType.DelegationsAllowed,
		}
		gVestingTypes = append(gVestingTypes, gvt)
	}

	genesis.VestingTypes = gVestingTypes
	allVestings := k.GetAllAccountVestings(ctx)

	for i := 0; i < len(allVestings); i++ {
		genesis.AccountVestingsList.Vestings = append(genesis.AccountVestingsList.Vestings, &allVestings[i])
	}

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

func DurationFromUnits(unit PeriodUnit, value int64) time.Duration {
	switch unit {
	case Day:
		return 24 * time.Hour * time.Duration(value)
	case Hour:
		return time.Hour * time.Duration(value)
	case Minute:
		return time.Minute * time.Duration(value)
	case Second:
		return time.Second * time.Duration(value)
	}
	panic(sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Unknown PeriodUnit: %s", unit))
}

func UnitsFromDuration(duration time.Duration) (unit PeriodUnit, value int64) {
	if duration%(24*time.Hour) == 0 {
		return Day, int64(duration / (24 * time.Hour))
	}
	if duration%(time.Hour) == 0 {
		return Hour, int64(duration / (time.Hour))
	}
	if duration%(time.Minute) == 0 {
		return Minute, int64(duration / (time.Minute))
	}
	return Second, int64(duration / (time.Second))
}
