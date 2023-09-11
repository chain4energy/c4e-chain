package types_test

import (
	"cosmossdk.io/math"
	testenv "github.com/chain4energy/c4e-chain/v2/testutil/env"
	"github.com/chain4energy/c4e-chain/v2/x/cfeminter/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
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
				MinterState: types.MinterState{
					SequenceId:                  2,
					AmountMinted:                math.NewInt(123),
					RemainderToMint:             sdk.MustNewDecFromStr("123.221"),
					LastMintBlockTime:           time.Now(),
					RemainderFromPreviousMinter: sdk.MustNewDecFromStr("324.543"),
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid:        false,
			errorMassage: "denom cannot be empty",
		},
		{
			desc: "no minters in params",
			genState: &types.GenesisState{
				Params: types.NewParams("myc4e", time.Now(), []*types.Minter{}),
				MinterState: types.MinterState{
					SequenceId:                  2,
					AmountMinted:                math.NewInt(123),
					RemainderToMint:             sdk.MustNewDecFromStr("123.221"),
					LastMintBlockTime:           time.Now(),
					RemainderFromPreviousMinter: sdk.MustNewDecFromStr("324.543"),
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid:        false,
			errorMassage: "no minters defined",
		},
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				Params: types.NewParams("myc4e", createOkCfeminterParams().StartTime, createOkCfeminterParams().Minters),
				MinterState: types.MinterState{
					SequenceId:                  2,
					AmountMinted:                math.NewInt(123),
					RemainderToMint:             sdk.MustNewDecFromStr("123.221"),
					LastMintBlockTime:           time.Now(),
					RemainderFromPreviousMinter: sdk.MustNewDecFromStr("324.543"),
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "invalid genesis state - wrong minter",
			genState: &types.GenesisState{
				Params: types.NewParams("myc4e", createNotOkCfeminterParams().StartTime, createNotOkCfeminterParams().Minters),
				MinterState: types.MinterState{
					SequenceId:                  2,
					AmountMinted:                math.NewInt(123),
					RemainderToMint:             sdk.MustNewDecFromStr("123.221"),
					LastMintBlockTime:           time.Now(),
					RemainderFromPreviousMinter: sdk.MustNewDecFromStr("324.543"),
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid:        false,
			errorMassage: "missing minter with sequence id 3",
		},
		{
			desc: "invalid genesis state - wrong minter state - amount",
			genState: &types.GenesisState{
				Params: types.NewParams("myc4e", createOkCfeminterParams().StartTime, createOkCfeminterParams().Minters),
				MinterState: types.MinterState{
					SequenceId:                  2,
					AmountMinted:                math.NewInt(-123),
					RemainderToMint:             sdk.MustNewDecFromStr("123.221"),
					LastMintBlockTime:           time.Now(),
					RemainderFromPreviousMinter: sdk.MustNewDecFromStr("324.543"),
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid:        false,
			errorMassage: "minter state validation error: amountMinted cannot be less than 0",
		},
		{
			desc: "invalid genesis state - wrong minter state - reminder to mint",
			genState: &types.GenesisState{
				Params: types.NewParams("myc4e", createOkCfeminterParams().StartTime, createOkCfeminterParams().Minters),
				MinterState: types.MinterState{
					SequenceId:                  2,
					AmountMinted:                math.NewInt(123),
					RemainderToMint:             sdk.MustNewDecFromStr("-123.221"),
					LastMintBlockTime:           time.Now(),
					RemainderFromPreviousMinter: sdk.MustNewDecFromStr("324.543"),
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid:        false,
			errorMassage: "minter state validation error: remainderToMint cannot be less than 0",
		},
		{
			desc: "invalid genesis state - wrong minter state - remainder from previous minter",
			genState: &types.GenesisState{
				Params: types.NewParams("myc4e", createOkCfeminterParams().StartTime, createOkCfeminterParams().Minters),
				MinterState: types.MinterState{
					SequenceId:                  2,
					AmountMinted:                math.NewInt(123),
					RemainderToMint:             sdk.MustNewDecFromStr("123.221"),
					LastMintBlockTime:           time.Now(),
					RemainderFromPreviousMinter: sdk.MustNewDecFromStr("-324.543"),
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid:        false,
			errorMassage: "minter state validation error: remainderFromPreviousMinter cannot be less than 0",
		},
		{
			desc: "invalid genesis state - wrong minter state SequenceId",
			genState: &types.GenesisState{
				Params: types.NewParams("myc4e", createOkCfeminterParams().StartTime, createOkCfeminterParams().Minters),
				MinterState: types.MinterState{
					SequenceId:                  6,
					AmountMinted:                math.NewInt(123),
					RemainderToMint:             sdk.MustNewDecFromStr("123.221"),
					LastMintBlockTime:           time.Now(),
					RemainderFromPreviousMinter: sdk.MustNewDecFromStr("324.543"),
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid:        false,
			errorMassage: "cfeminter genesis validation error: minter state sequence id 6 not found in minters",
		},
		{
			desc: "valid genesis state with history",
			genState: &types.GenesisState{
				Params: types.NewParams("myc4e", createOkCfeminterParams().StartTime, createOkCfeminterParams().Minters),
				MinterState: types.MinterState{
					SequenceId:                  2,
					AmountMinted:                math.NewInt(123),
					RemainderToMint:             sdk.MustNewDecFromStr("123.221"),
					LastMintBlockTime:           time.Now(),
					RemainderFromPreviousMinter: sdk.MustNewDecFromStr("324.543"),
				},
				StateHistory: createHistory(),
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "invalid genesis state - wrong minter state - RemainderFromPreviousMinter is nil",
			genState: &types.GenesisState{
				Params: types.NewParams("myc4e", createOkCfeminterParams().StartTime, createOkCfeminterParams().Minters),
				MinterState: types.MinterState{
					SequenceId:                  2,
					AmountMinted:                math.NewInt(100),
					RemainderToMint:             sdk.MustNewDecFromStr("324.543"),
					LastMintBlockTime:           time.Now(),
					RemainderFromPreviousMinter: sdk.Dec{},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid:        false,
			errorMassage: "minter state validation error: remainderFromPreviousMinter cannot be nil",
		},
		{
			desc: "invalid genesis state - wrong minter state - RemainderToMint is nil",
			genState: &types.GenesisState{
				Params: types.NewParams("myc4e", createOkCfeminterParams().StartTime, createOkCfeminterParams().Minters),
				MinterState: types.MinterState{
					SequenceId:                  2,
					AmountMinted:                math.NewInt(100),
					RemainderToMint:             sdk.Dec{},
					LastMintBlockTime:           time.Now(),
					RemainderFromPreviousMinter: sdk.MustNewDecFromStr("324.543"),
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid:        false,
			errorMassage: "minter state validation error: remainderToMint cannot be nil",
		},
		{
			desc: "invalid genesis state - wrong minter state - AmountMinted is nil",
			genState: &types.GenesisState{
				Params: types.NewParams("myc4e", createOkCfeminterParams().StartTime, createOkCfeminterParams().Minters),
				MinterState: types.MinterState{
					SequenceId:                  2,
					AmountMinted:                math.Int{},
					RemainderToMint:             sdk.MustNewDecFromStr("324.543"),
					LastMintBlockTime:           time.Now(),
					RemainderFromPreviousMinter: sdk.MustNewDecFromStr("324.543"),
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid:        false,
			errorMassage: "minter state validation error: amountMinted cannot be nil",
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

func createHistory() []*types.MinterState {
	history := make([]*types.MinterState, 0)
	state1 := types.MinterState{
		SequenceId:                  0,
		AmountMinted:                math.NewInt(324),
		RemainderToMint:             sdk.MustNewDecFromStr("1243.221"),
		LastMintBlockTime:           time.Now(),
		RemainderFromPreviousMinter: sdk.MustNewDecFromStr("3124.543"),
	}
	state2 := types.MinterState{
		SequenceId:                  1,
		AmountMinted:                math.NewInt(432),
		RemainderToMint:             sdk.MustNewDecFromStr("12433.221"),
		LastMintBlockTime:           time.Now(),
		RemainderFromPreviousMinter: sdk.MustNewDecFromStr("3284.543"),
	}
	return append(history, &state1, &state2)
}

func createOkCfeminterParams() types.Params {
	startTime := time.Now()

	endTime1 := startTime.Add(PeriodDuration)
	endTime2 := endTime1.Add(PeriodDuration)

	linearMinting1 := types.LinearMinting{Amount: math.NewInt(1000000)}
	linearMinting2 := types.LinearMinting{Amount: math.NewInt(100000)}
	config, _ := codectypes.NewAnyWithValue(&linearMinting1)
	config2, _ := codectypes.NewAnyWithValue(&linearMinting2)

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Config: config}
	minter2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Config: config2}
	minter3 := types.Minter{SequenceId: 3, Config: testenv.NoMintingConfig}

	minters := []*types.Minter{&minter1, &minter2, &minter3}
	return types.Params{
		StartTime: startTime,
		Minters:   minters,
	}
}

func createNotOkCfeminterParams() types.Params {
	startTime := time.Now()

	endTime1 := startTime.Add(PeriodDuration)
	endTime2 := endTime1.Add(PeriodDuration)

	linearMinting1 := types.LinearMinting{Amount: math.NewInt(1000000)}
	linearMinting2 := types.LinearMinting{Amount: math.NewInt(100000)}
	config, _ := codectypes.NewAnyWithValue(&linearMinting1)
	config2, _ := codectypes.NewAnyWithValue(&linearMinting2)

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Config: config}
	minter2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Config: config2}
	minter3 := types.Minter{SequenceId: 5, Config: testenv.NoMintingConfig}

	minters := []*types.Minter{&minter1, &minter2, &minter3}

	return types.Params{
		StartTime: startTime,
		Minters:   minters,
	}
}
