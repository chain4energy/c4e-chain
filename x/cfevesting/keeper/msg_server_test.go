package keeper_test

import (
	"strconv"
	"time"

	testutils "github.com/chain4energy/c4e-chain/testutil/module/cfevesting"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"

	"testing"

	"github.com/stretchr/testify/require"
)

const (
	vPool1 = "v-pool-1"
	vPool2 = "v-pool-2"
)

func verifyVestingResponse(t *testing.T, response *types.QueryVestingPoolsResponse, accVestings types.AccountVestings, current time.Time, delegationAllowed bool) {
	require.EqualValues(t, len(accVestings.VestingPools), len(response.VestingPools))

	for _, vesting := range accVestings.VestingPools {
		found := false
		for _, vestingInfo := range response.VestingPools {
			if vesting.Id == vestingInfo.Id {
				require.EqualValues(t, vesting.VestingType, vestingInfo.VestingType)
				require.EqualValues(t, vesting.Name, vestingInfo.Name)
				require.EqualValues(t, testutils.GetExpectedWithdrawableForVesting(*vesting, current).String(), response.VestingPools[0].Withdrawable)
				require.EqualValues(t, true, vesting.LockStart.Equal(vestingInfo.LockStart))
				require.EqualValues(t, true, vesting.LockEnd.Equal(vestingInfo.LockEnd))
				require.EqualValues(t, commontestutils.DefaultTestDenom, response.VestingPools[0].Vested.Denom)
				require.EqualValues(t, vesting.Vested, response.VestingPools[0].Vested.Amount)
				require.EqualValues(t, vesting.LastModificationVested.Sub(vesting.LastModificationWithdrawn).String(), response.VestingPools[0].CurrentVestedAmount)
				require.EqualValues(t, vesting.Sent.String(), response.VestingPools[0].SentAmount)

				found = true
			}
		}
		require.True(t, found, "not found vesting nfo with Id: "+strconv.FormatInt(int64(vesting.Id), 10))
	}
}
