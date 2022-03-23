package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingmodule "github.com/cosmos/cosmos-sdk/x/staking"

	"github.com/chain4energy/c4e-chain/x/cfevesting/internal/testutils"
)

func TestDelegate(t *testing.T) {
	addHelperModuleAccountPerms()
	const vested = 1000000
	app, ctx := setupApp(0)
	setupStakingBondDenom(ctx, app)

	acountsAddresses, validatorsAddresses := testutils.CreateAccounts(2, 1)

	setupValidators(t, ctx, app, validatorsAddresses, vested/2)

	accAddr := acountsAddresses[0]
	delegableAccAddr := acountsAddresses[1]

	// adds coind to delegable account - means that coins in vesting for accAddr
	addCoinsToAccount(vested, ctx, app, delegableAccAddr)
	// adds some coins to distibutor account - to allow test to process
	addCoinsToModuleByName(100000000, distrtypes.ModuleName, ctx, app)

	valAddr := validatorsAddresses[0]

	setupAccountsVestings(ctx, app, accAddr.String(), delegableAccAddr.String(), vested, true)

	delegate(t, ctx, app, accAddr, delegableAccAddr, valAddr, vested/2, 0, vested, 0, vested/2)
	verifyDelegations(t, ctx, app, delegableAccAddr,  []sdk.ValAddress{valAddr}, []int64{vested/2})

	verifyQueryRewards(t, ctx, app, delegableAccAddr, valAddr, false, 0)

	stakingmodule.EndBlocker(ctx, app.StakingKeeper)
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)

	validatorRewards := uint64(10000)
	allocateRewardsToValidator(ctx, app, validatorRewards, valAddr)

	verifyQueryRewards(t, ctx, app, delegableAccAddr, valAddr, true, validatorRewards/2)

	delegate(t, ctx, app, accAddr, delegableAccAddr, valAddr, vested/2, 0, vested/2, int64(validatorRewards/2), 0)
	verifyDelegations(t, ctx, app, delegableAccAddr,  []sdk.ValAddress{valAddr}, []int64{vested})

}


