package cfeminter

import (
	"testing"

	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/stretchr/testify/require"
)

func CompareMinters(t *testing.T, m1 types.Minter, m2 types.Minter) {
	require.True(t, m1.Start.Equal(m2.Start))
	for i, p1 := range m1.Periods {
		p2 := m2.Periods[i]
		if p1.PeriodEnd == nil {
			require.Nil(t, p2.PeriodEnd)
		} else {
			require.True(t, p1.PeriodEnd.Equal(*p2.PeriodEnd))
		}
		require.EqualValues(t, p1.OrderingId, p2.OrderingId)
		require.EqualValues(t, p1.TimeLinearMinter, p2.TimeLinearMinter)
		require.EqualValues(t, p1.Type, p2.Type)
	}
}
