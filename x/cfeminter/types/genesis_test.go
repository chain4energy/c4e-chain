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
				MinterState: types.MinterState{
					SequenceId:                  2,
					AmountMinted:                sdk.NewInt(123),
					RemainderToMint:             sdk.MustNewDecFromStr("123.221"),
					LastMintBlockTime:           time.Now(),
					RemainderFromPreviousPeriod: sdk.MustNewDecFromStr("324.543"),
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
					AmountMinted:                sdk.NewInt(123),
					RemainderToMint:             sdk.MustNewDecFromStr("123.221"),
					LastMintBlockTime:           time.Now(),
					RemainderFromPreviousPeriod: sdk.MustNewDecFromStr("324.543"),
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
				Params: types.NewParams("myc4e", time.Now(), createOkMinters()),
				MinterState: types.MinterState{
					SequenceId:                  2,
					AmountMinted:                sdk.NewInt(123),
					RemainderToMint:             sdk.MustNewDecFromStr("123.221"),
					LastMintBlockTime:           time.Now(),
					RemainderFromPreviousPeriod: sdk.MustNewDecFromStr("324.543"),
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "invalid genesis state - wrong minter",
			genState: &types.GenesisState{
				Params: types.NewParams("myc4e", time.Now(), createNotOkMinters()),
				MinterState: types.MinterState{
					SequenceId:                  2,
					AmountMinted:                sdk.NewInt(123),
					RemainderToMint:             sdk.MustNewDecFromStr("123.221"),
					LastMintBlockTime:           time.Now(),
					RemainderFromPreviousPeriod: sdk.MustNewDecFromStr("324.543"),
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid:        false,
			errorMassage: "missing minter with sequence id 3",
		},
		{
			desc: "invalid genesis state - wrong minter state - amount",
			genState: &types.GenesisState{
				Params: types.NewParams("myc4e", time.Now(), createOkMinters()),
				MinterState: types.MinterState{
					SequenceId:                  2,
					AmountMinted:                sdk.NewInt(-123),
					RemainderToMint:             sdk.MustNewDecFromStr("123.221"),
					LastMintBlockTime:           time.Now(),
					RemainderFromPreviousPeriod: sdk.MustNewDecFromStr("324.543"),
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid:        false,
			errorMassage: "minter state amount cannot be less than 0",
		},
		{
			desc: "invalid genesis state - wrong minter state - reminder to mint",
			genState: &types.GenesisState{
				Params: types.NewParams("myc4e", time.Now(), createOkMinters()),
				MinterState: types.MinterState{
					SequenceId:                  2,
					AmountMinted:                sdk.NewInt(123),
					RemainderToMint:             sdk.MustNewDecFromStr("-123.221"),
					LastMintBlockTime:           time.Now(),
					RemainderFromPreviousPeriod: sdk.MustNewDecFromStr("324.543"),
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid:        false,
			errorMassage: "minter remainder to mint amount cannot be less than 0",
		},
		{
			desc: "invalid genesis state - wrong minter state - remainder from previous minter",
			genState: &types.GenesisState{
				Params: types.NewParams("myc4e", time.Now(), createOkMinters()),
				MinterState: types.MinterState{
					SequenceId:                  2,
					AmountMinted:                sdk.NewInt(123),
					RemainderToMint:             sdk.MustNewDecFromStr("123.221"),
					LastMintBlockTime:           time.Now(),
					RemainderFromPreviousPeriod: sdk.MustNewDecFromStr("-324.543"),
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid:        false,
			errorMassage: "minter remainder from previous period amount cannot be less than 0",
		},
		{
			desc: "invalid genesis state - wrong minter state SequenceId",
			genState: &types.GenesisState{
				Params: types.NewParams("myc4e", time.Now(), createOkMinters()),
				MinterState: types.MinterState{
					SequenceId:                  6,
					AmountMinted:                sdk.NewInt(123),
					RemainderToMint:             sdk.MustNewDecFromStr("123.221"),
					LastMintBlockTime:           time.Now(),
					RemainderFromPreviousPeriod: sdk.MustNewDecFromStr("324.543"),
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid:        false,
			errorMassage: "minter state current sequence id not found in minters",
		},
		{
			desc: "valid genesis state with history",
			genState: &types.GenesisState{
				Params: types.NewParams("myc4e", time.Now(), createOkMinters()),
				MinterState: types.MinterState{
					SequenceId:                  2,
					AmountMinted:                sdk.NewInt(123),
					RemainderToMint:             sdk.MustNewDecFromStr("123.221"),
					LastMintBlockTime:           time.Now(),
					RemainderFromPreviousPeriod: sdk.MustNewDecFromStr("324.543"),
				},
				StateHistory: createHistory(),
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
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
		AmountMinted:                sdk.NewInt(324),
		RemainderToMint:             sdk.MustNewDecFromStr("1243.221"),
		LastMintBlockTime:           time.Now(),
		RemainderFromPreviousPeriod: sdk.MustNewDecFromStr("3124.543"),
	}
	state2 := types.MinterState{
		SequenceId:                  1,
		AmountMinted:                sdk.NewInt(432),
		RemainderToMint:             sdk.MustNewDecFromStr("12433.221"),
		LastMintBlockTime:           time.Now(),
		RemainderFromPreviousPeriod: sdk.MustNewDecFromStr("3284.543"),
	}
	return append(history, &state1, &state2)
}

func createOkMinters() []*types.Minter {
	startTime := time.Now()

	endTime1 := startTime.Add(PeriodDuration)
	endTime2 := endTime1.Add(PeriodDuration)

	LinearMinting1 := types.LinearMinting{Amount: sdk.NewInt(1000000)}
	LinearMinting2 := types.LinearMinting{Amount: sdk.NewInt(100000)}

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Type: types.LINEAR_MINTING, LinearMinting: &LinearMinting1}
	minter2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Type: types.LINEAR_MINTING, LinearMinting: &LinearMinting2}

	minter3 := types.Minter{SequenceId: 3, Type: types.NO_MINTING}
	minters := []*types.Minter{&minter1, &minter2, &minter3}
	return minters
}

func createNotOkMinters() []*types.Minter {
	startTime := time.Now()

	endTime1 := startTime.Add(PeriodDuration)
	endTime2 := endTime1.Add(PeriodDuration)

	LinearMinting1 := types.LinearMinting{Amount: sdk.NewInt(1000000)}
	LinearMinting2 := types.LinearMinting{Amount: sdk.NewInt(100000)}

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Type: types.LINEAR_MINTING, LinearMinting: &LinearMinting1}
	minter2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Type: types.LINEAR_MINTING, LinearMinting: &LinearMinting2}

	minter3 := types.Minter{SequenceId: 5, Type: types.NO_MINTING}
	minters := []*types.Minter{&minter1, &minter2, &minter3}

	return minters
}
