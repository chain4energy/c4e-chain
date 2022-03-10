package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"

	"github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	"github.com/stretchr/testify/require"
)

func TestCalculateWithdrawable(t *testing.T) {
	vesting := types.Vesting{"test", 1000, 10000, 110000, 1000000, 0, 0, 10, 0, true, 0}

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
	vesting.Vested = 1000

	withdrawable = keeper.CalculateWithdrawable(10010, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	withdrawable = keeper.CalculateWithdrawable(10099, vesting)
	require.EqualValues(t, sdk.ZeroInt(), withdrawable)

	withdrawable = keeper.CalculateWithdrawable(10100, vesting)
	require.EqualValues(t, sdk.NewInt(1), withdrawable)

	vesting.Vested = 1000000
	vesting.Withdrawn = 500
	withdrawable = keeper.CalculateWithdrawable(10100, vesting)
	require.EqualValues(t, sdk.NewInt(500), withdrawable)

}
