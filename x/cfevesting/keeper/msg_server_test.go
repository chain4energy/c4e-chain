package keeper_test

import (
	"time"

	testutils "github.com/chain4energy/c4e-chain/v2/testutil/module/cfevesting"

	testenv "github.com/chain4energy/c4e-chain/v2/testutil/env"

	"github.com/chain4energy/c4e-chain/v2/x/cfevesting/types"

	"testing"

	"github.com/stretchr/testify/require"
)

const (
	vPool1 = "v-pool-1"
	vPool2 = "v-pool-2"
)

func verifyVestingResponse(t *testing.T, response *types.QueryVestingPoolsResponse, accVestingPools types.AccountVestingPools, current time.Time, delegationAllowed bool) {
	require.EqualValues(t, len(accVestingPools.VestingPools), len(response.VestingPools))

	for _, vesting := range accVestingPools.VestingPools {
		found := false
		for _, vestingInfo := range response.VestingPools {
			if vesting.Name == vestingInfo.Name {
				require.EqualValues(t, vesting.VestingType, vestingInfo.VestingType)
				require.EqualValues(t, vesting.Name, vestingInfo.Name)
				require.EqualValues(t, testutils.GetExpectedWithdrawableForVesting(*vesting, current).String(), response.VestingPools[0].Withdrawable)
				require.EqualValues(t, true, vesting.LockStart.Equal(vestingInfo.LockStart))
				require.EqualValues(t, true, vesting.LockEnd.Equal(vestingInfo.LockEnd))
				require.EqualValues(t, testenv.DefaultTestDenom, response.VestingPools[0].InitiallyLocked.Denom)
				require.EqualValues(t, vesting.InitiallyLocked, response.VestingPools[0].InitiallyLocked.Amount)
				require.EqualValues(t, vesting.GetLockedNotReserved().String(), response.VestingPools[0].CurrentlyLocked)
				require.EqualValues(t, vesting.Sent.String(), response.VestingPools[0].SentAmount)

				found = true
			}
		}
		require.True(t, found, "not found vesting info with Name: "+vesting.Name)
	}
}
