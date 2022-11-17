package cfeairdrop

import (
	"testing"

	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
	cfeairdropmodulekeeper "github.com/chain4energy/c4e-chain/x/cfeairdrop/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
)

type ContextC4eAirdropUtils struct {
	C4eAirdropUtils
	testContext commontestutils.TestContext
}

func NewContextC4eAirdropUtils(t *testing.T, testContext commontestutils.TestContext, helpeCfeairdropmodulekeeper *cfeairdropmodulekeeper.Keeper,
	helperAccountKeeper *authkeeper.AccountKeeper,
	bankUtils *commontestutils.BankUtils) *ContextC4eAirdropUtils {
	c4eAirdropUtils := NewC4eAirdropUtils(t, helpeCfeairdropmodulekeeper, helperAccountKeeper, bankUtils)
	return &ContextC4eAirdropUtils{C4eAirdropUtils: c4eAirdropUtils, testContext: testContext}
}

func (h *ContextC4eAirdropUtils) SendToAirdropAccount(toAddress sdk.AccAddress,
	amount sdk.Int, startTime int64, endTime int64, createAccount bool) {
	h.C4eAirdropUtils.SendToAirdropAccount(h.testContext.GetContext(), toAddress, amount, startTime, endTime, createAccount)
}
