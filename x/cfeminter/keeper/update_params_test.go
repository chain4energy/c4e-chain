package keeper_test

import (
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	"testing"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestCorrectUpdateParams(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	testHelper := prepareApp(t, startTime, startTime, createLinearMintings(startTime))
	testHelper.C4eMinterUtils.SetMinterState(1, sdk.ZeroInt(), sdk.ZeroDec(), startTime, sdk.ZeroDec())

	params := types.Params{
		MintDenom: testenv.DefaultTestDenom,
		StartTime: testenv.TestEnvTime,
		Minters:   createLinearMintings(testenv.TestEnvTime),
	}
	testHelper.C4eMinterUtils.UpdateParams(testenv.GetAuthority(), params)
}

func TestUpdateParamsWrongMInter(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	testHelper := prepareApp(t, startTime, startTime, createLinearMintings(startTime))
	testHelper.C4eMinterUtils.SetMinterState(1, sdk.ZeroInt(), sdk.ZeroDec(), startTime, sdk.ZeroDec())

	params := types.Params{
		MintDenom: testenv.DefaultTestDenom,
		StartTime: testenv.TestEnvTime,
		Minters:   []*types.Minter{},
	}
	testHelper.C4eMinterUtils.UpdateParamsError(testenv.GetAuthority(), params, "minter state sequence id 1 not found in minters: invalid proposal content")
}

func TestUpdateParamsWrongMinterSequenceId(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	testHelper := prepareApp(t, startTime, startTime, createLinearMintings(startTime))
	testHelper.C4eMinterUtils.SetMinterState(10, sdk.ZeroInt(), sdk.ZeroDec(), startTime, sdk.ZeroDec())

	params := types.Params{
		MintDenom: testenv.DefaultTestDenom,
		StartTime: testenv.TestEnvTime,
		Minters:   createLinearMintings(testenv.TestEnvTime),
	}
	testHelper.C4eMinterUtils.UpdateParamsError(testenv.GetAuthority(), params, "minter state sequence id 10 not found in minters: invalid proposal content")
}

func TestUpdateParamsWrongAuthority(t *testing.T) {
	startTime := time.Date(2022, 2, 3, 0, 0, 0, 0, time.UTC)
	testHelper := prepareApp(t, startTime, startTime, createLinearMintings(startTime))
	testHelper.C4eMinterUtils.SetMinterState(2, sdk.ZeroInt(), sdk.ZeroDec(), startTime, sdk.ZeroDec())

	params := types.Params{
		MintDenom: testenv.DefaultTestDenom,
		StartTime: testenv.TestEnvTime,
		Minters:   createLinearMintings(testenv.TestEnvTime),
	}
	testHelper.C4eMinterUtils.UpdateParamsError("abcd", params, "expected gov account as only signer for proposal message: invalid proposal content")
}
