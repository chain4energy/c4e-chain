package keeper_test

import (
	"testing"
	"time"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	"github.com/chain4energy/c4e-chain/x/cfeairdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestCreateAirdropAccount(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(1, 0)

	moduleAmount := sdk.NewInt(10000)
	amount := sdk.NewInt(1000)
	startTime := commontestutils.TestEnvTime.Add(-24 * 100 * time.Hour).Unix()
	endTime := commontestutils.TestEnvTime.Add(24 * 100 * time.Hour).Unix()
	testHelper.BankUtils.AddDefaultDenomCoinsToModule(moduleAmount, types.ModuleName)

	testHelper.C4eAirdropUtils.SendToAirdropAccount(acountsAddresses[0],
		amount,
		startTime,
		endTime, true,
	)

	testHelper.C4eAirdropUtils.SendToAirdropAccount(acountsAddresses[0],
		amount,
		startTime,
		endTime, false,
	)

	testHelper.C4eAirdropUtils.SendToAirdropAccount(acountsAddresses[0],
		amount,
		startTime,
		endTime, false,
	)
}
