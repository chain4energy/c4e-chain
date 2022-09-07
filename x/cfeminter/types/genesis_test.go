package types_test

import (
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc         string
		genState     *types.GenesisState
		valid        bool
		errorMassage string
	}{
		{
			desc: "no params",
			genState: &types.GenesisState{
				MinterState: types.MinterState{2, sdk.NewInt(123)},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid:        false,
			errorMassage: "denom cannot be empty",
		},
		{
			desc: "no periods",
			genState: &types.GenesisState{
				Params:      types.NewParams("myc4e", types.Minter{}),
				MinterState: types.MinterState{2, sdk.NewInt(123)},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid:        false,
			errorMassage: "no minter periods defined",
		},
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				Params:      types.NewParams("myc4e", createOkMinter()),
				MinterState: types.MinterState{2, sdk.NewInt(123)},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "invalid genesis state - wrong minter",
			genState: &types.GenesisState{
				Params:      types.NewParams("myc4e", createNotOkMinter()),
				MinterState: types.MinterState{2, sdk.NewInt(123)},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid:        false,
			errorMassage: "missing period with ordering id 3",
		},
		{
			desc: "invalid genesis state - wrong minter state",
			genState: &types.GenesisState{
				Params:      types.NewParams("myc4e", createOkMinter()),
				MinterState: types.MinterState{2, sdk.NewInt(-123)},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid:        false,
			errorMassage: "minter state amount cannot be less than 0",
		},
		{
			desc: "invalid genesis state - wrong minter state ordering id",
			genState: &types.GenesisState{
				Params:      types.NewParams("myc4e", createOkMinter()),
				MinterState: types.MinterState{6, sdk.NewInt(123)},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid:        false,
			errorMassage: "minter state Current Ordering Id not found in minter periods",
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tc.errorMassage)
			}
		})
	}
}

func createOkMinter() types.Minter {
	startTime := time.Now()

	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}
	linearMinter2 := types.TimeLinearMinter{Amount: sdk.NewInt(100000)}

	period1 := types.MintingPeriod{Position: 1, PeriodEnd: &endTime1, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{Position: 2, PeriodEnd: &endTime2, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{Position: 3, Type: types.NO_MINTING}
	periods := []*types.MintingPeriod{&period1, &period2, &period3}
	minter := types.Minter{Start: startTime, Periods: periods}
	return minter
}

func createNotOkMinter() types.Minter {
	startTime := time.Now()

	endTime1 := startTime.Add(time.Duration(PeriodDuration))
	endTime2 := endTime1.Add(time.Duration(PeriodDuration))

	linearMinter1 := types.TimeLinearMinter{Amount: sdk.NewInt(1000000)}
	linearMinter2 := types.TimeLinearMinter{Amount: sdk.NewInt(100000)}

	period1 := types.MintingPeriod{Position: 1, PeriodEnd: &endTime1, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter1}
	period2 := types.MintingPeriod{Position: 2, PeriodEnd: &endTime2, Type: types.TIME_LINEAR_MINTER, TimeLinearMinter: &linearMinter2}

	period3 := types.MintingPeriod{Position: 5, Type: types.NO_MINTING}
	periods := []*types.MintingPeriod{&period1, &period2, &period3}
	minter := types.Minter{Start: startTime, Periods: periods}
	return minter
}
