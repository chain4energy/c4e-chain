package cfeminter_test

import (
	"github.com/chain4energy/c4e-chain/testutil/app"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfeminter/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const PeriodDuration = time.Duration(345600000000 * 1000000)
const Year = time.Hour * 24 * 365
const NanoSecondsInFourYears = Year * 4

func TestGenesis(t *testing.T) {
	layout := "2006-01-02T15:04:05.000Z"
	str := "2014-11-12T11:45:26.371Z"
	mintTime, _ := time.Parse(layout, str)

	genesisState := types.GenesisState{
		Params: types.NewParams(testenv.DefaultTestDenom, time.Now(), createMinter(time.Now())),
		MinterState: types.MinterState{
			SequenceId:                  9,
			AmountMinted:                sdk.NewInt(12312),
			RemainderToMint:             sdk.MustNewDecFromStr("1233.546"),
			RemainderFromPreviousMinter: sdk.MustNewDecFromStr("7654.423"),
			LastMintBlockTime:           mintTime,
		},

		// this line is used by starport scaffolding # genesis/test/state
	}

	testHelper := app.SetupTestApp(t)
	testHelper.C4eMinterUtils.InitGenesis(genesisState)
	testHelper.C4eMinterUtils.ExportGenesis(genesisState)
	// this line is used by starport scaffolding # genesis/test/assert
}

func TestGenesisWithHistory(t *testing.T) {
	layout := "2006-01-02T15:04:05.000Z"
	str := "2014-11-12T11:45:26.371Z"
	mintTime, _ := time.Parse(layout, str)

	genesisState := types.GenesisState{
		Params: types.NewParams(testenv.DefaultTestDenom, time.Now(), createMinter(time.Now())),
		MinterState: types.MinterState{
			SequenceId:                  9,
			AmountMinted:                sdk.NewInt(12312),
			RemainderToMint:             sdk.MustNewDecFromStr("1233.546"),
			RemainderFromPreviousMinter: sdk.MustNewDecFromStr("7654.423"),
			LastMintBlockTime:           mintTime,
		},
		StateHistory: createHistory(),
		// this line is used by starport scaffolding # genesis/test/state

	}

	testHelper := app.SetupTestApp(t)
	testHelper.C4eMinterUtils.InitGenesis(genesisState)
	testHelper.C4eMinterUtils.ExportGenesis(genesisState)
}

func createHistory() []*types.MinterState {
	layout := "2006-01-02T15:04:05.000Z"
	str := "2014-11-12T11:45:26.371Z"
	mintTime, _ := time.Parse(layout, str)
	history := make([]*types.MinterState, 0)
	state1 := types.MinterState{
		SequenceId:                  0,
		AmountMinted:                sdk.NewInt(324),
		RemainderToMint:             sdk.MustNewDecFromStr("1243.221"),
		LastMintBlockTime:           mintTime,
		RemainderFromPreviousMinter: sdk.MustNewDecFromStr("3124.543"),
	}

	str = "2016-06-12T11:35:46.371Z"
	mintTime, _ = time.Parse(layout, str)
	state2 := types.MinterState{
		SequenceId:                  1,
		AmountMinted:                sdk.NewInt(432),
		RemainderToMint:             sdk.MustNewDecFromStr("12433.221"),
		LastMintBlockTime:           mintTime,
		RemainderFromPreviousMinter: sdk.MustNewDecFromStr("3284.543"),
	}

	return append(history, &state1, &state2)
}

func createMinter(startTime time.Time) []*types.Minter {
	endTime1 := startTime.Add(PeriodDuration)
	endTime2 := endTime1.Add(PeriodDuration)

	linearMinting1 := types.LinearMinting{Amount: sdk.NewInt(1000000)}
	linearMinting2 := types.LinearMinting{Amount: sdk.NewInt(100000)}
	config, _ := codectypes.NewAnyWithValue(&linearMinting1)
	config2, _ := codectypes.NewAnyWithValue(&linearMinting2)

	minter1 := types.Minter{SequenceId: 1, EndTime: &endTime1, Config: config}
	minter2 := types.Minter{SequenceId: 2, EndTime: &endTime2, Config: config2}
	minter3 := types.Minter{SequenceId: 3}

	return []*types.Minter{&minter1, &minter2, &minter3}
}
