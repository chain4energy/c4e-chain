package cfeminter_test

import (
	"github.com/cosmos/btcutil/bech32"
	"testing"
	"time"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"

	"github.com/chain4energy/c4e-chain/x/cfeminter/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const iterationText = "iterarion %d"

const PeriodDuration = time.Duration(345600000000 * 1000000)
const Year = time.Hour * 24 * 365
const SecondsInYear = int32(3600 * 24 * 365)

func convertAddressToC4E(address string) string {
	_, decoded, _ := bech32.Decode(address, 100)
	encoded, _ := bech32.Encode("c4e", decoded)
	return encoded
}

func TestGenesis(t *testing.T) {

	convertAddressToC4E("cosmos1ryams4f7af0p6yrj59hztu668sjvp6rhhzene5")

	//layout := "2006-01-02T15:04:05.000Z"
	//str := "2014-11-12T11:45:26.371Z"
	//mintTime, _ := time.Parse(layout, str)
	//genesisState := types.GenesisState{
	//	Params: types.NewParams("myc4e", createMinter(time.Now())),
	//	MinterState: types.MinterState{
	//		Position:                    9,
	//		AmountMinted:                sdk.NewInt(12312),
	//		RemainderToMint:             sdk.MustNewDecFromStr("1233.546"),
	//		RemainderFromPreviousPeriod: sdk.MustNewDecFromStr("7654.423"),
	//		LastMintBlockTime:           mintTime,
	//	},
	//
	//	// this line is used by starport scaffolding # genesis/test/state
	//
	//}
	//testHelper := testapp.SetupTestApp(t)
	//
	//testHelper.C4eMinterUtils.InitGenesis(genesisState)
	//testHelper.C4eMinterUtils.ExportGenesis(genesisState)
	// this line is used by starport scaffolding # genesis/test/assert
}

func TestGenesisWithHistory(t *testing.T) {
	layout := "2006-01-02T15:04:05.000Z"
	str := "2014-11-12T11:45:26.371Z"
	mintTime, _ := time.Parse(layout, str)
	genesisState := types.GenesisState{
		Params: types.NewParams("myc4e", createMinter(time.Now())),
		MinterState: types.MinterState{
			Position:                    9,
			AmountMinted:                sdk.NewInt(12312),
			RemainderToMint:             sdk.MustNewDecFromStr("1233.546"),
			RemainderFromPreviousPeriod: sdk.MustNewDecFromStr("7654.423"),
			LastMintBlockTime:           mintTime,
		},
		StateHistory: createHistory(),
		// this line is used by starport scaffolding # genesis/test/state

	}

	testHelper := testapp.SetupTestApp(t)

	testHelper.C4eMinterUtils.InitGenesis(genesisState)
	testHelper.C4eMinterUtils.ExportGenesis(genesisState)
}

func createHistory() []*types.MinterState {
	layout := "2006-01-02T15:04:05.000Z"
	str := "2014-11-12T11:45:26.371Z"
	mintTime, _ := time.Parse(layout, str)

	history := make([]*types.MinterState, 0)
	state1 := types.MinterState{
		Position:                    0,
		AmountMinted:                sdk.NewInt(324),
		RemainderToMint:             sdk.MustNewDecFromStr("1243.221"),
		LastMintBlockTime:           mintTime,
		RemainderFromPreviousPeriod: sdk.MustNewDecFromStr("3124.543"),
	}
	str = "2016-06-12T11:35:46.371Z"
	mintTime, _ = time.Parse(layout, str)
	state2 := types.MinterState{
		Position:                    1,
		AmountMinted:                sdk.NewInt(432),
		RemainderToMint:             sdk.MustNewDecFromStr("12433.221"),
		LastMintBlockTime:           mintTime,
		RemainderFromPreviousPeriod: sdk.MustNewDecFromStr("3284.543"),
	}
	return append(history, &state1, &state2)
}

func createMinter(startTime time.Time) types.Minter {
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
