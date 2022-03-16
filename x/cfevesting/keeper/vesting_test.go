package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/stretchr/testify/require"
)

func TestCalculateWithdrawable(t *testing.T) {
	vesting := types.Vesting{
		Id:                1,
		VestingType:       "test",
		VestingStartBlock: 1000,
		LockEndBlock:      10000,
		VestingEndBlock:   110000,
		Vested:            sdk.NewInt(1000000),
		// Claimable:            0,
		// LastFreeingBlock:     0,
		FreeCoinsBlockPeriod: 10,
		// FreeCoinsPerPeriod:   0,
		DelegationAllowed: true,
		Withdrawn:         sdk.ZeroInt(),
	}

	withdrawable := keeper.CalculateWithdrawable(100, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	withdrawable = keeper.CalculateWithdrawable(1000, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	withdrawable = keeper.CalculateWithdrawable(10000, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	withdrawable = keeper.CalculateWithdrawable(10009, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	withdrawable = keeper.CalculateWithdrawable(10010, vesting)
	require.EqualValues(t, sdk.NewInt(100), withdrawable)

	withdrawable = keeper.CalculateWithdrawable(10100, vesting)
	require.EqualValues(t, sdk.NewInt(1000), withdrawable)

	withdrawable = keeper.CalculateWithdrawable(110000, vesting)
	require.EqualValues(t, sdk.NewInt(1000000), withdrawable)

	withdrawable = keeper.CalculateWithdrawable(109999, vesting)
	require.EqualValues(t, sdk.NewInt(999900), withdrawable)

	vesting.VestingEndBlock = 109999

	withdrawable = keeper.CalculateWithdrawable(10100, vesting)
	require.EqualValues(t, sdk.NewInt(1000), withdrawable)

	withdrawable = keeper.CalculateWithdrawable(109999, vesting)
	require.EqualValues(t, sdk.NewInt(1000000), withdrawable)

	vesting.VestingEndBlock = 110000
	vesting.Vested = sdk.NewInt(1000)

	withdrawable = keeper.CalculateWithdrawable(10010, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	withdrawable = keeper.CalculateWithdrawable(10099, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	withdrawable = keeper.CalculateWithdrawable(10100, vesting)
	require.EqualValues(t, sdk.NewInt(1), withdrawable)

	vesting.Vested = sdk.NewInt(1000000)
	vesting.Withdrawn = sdk.NewInt(500)
	withdrawable = keeper.CalculateWithdrawable(10100, vesting)
	require.EqualValues(t, sdk.NewInt(500), withdrawable)

}
