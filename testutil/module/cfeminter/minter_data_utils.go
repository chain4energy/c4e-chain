package cfeminter

import (
	"testing"

	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/stretchr/testify/require"
)

func CompareCfeminterParams(t *testing.T, m1 types.Params, m2 types.Params) {
	require.True(t, m1.StartTime.Equal(m2.StartTime))
	require.True(t, m1.MintDenom == m2.MintDenom)
	for i, p1 := range m1.Minters {
		p2 := m2.Minters[i]
		if p1.EndTime == nil {
			require.Nil(t, p2.EndTime)
		} else {
			require.True(t, p1.EndTime.Equal(*p2.EndTime))
		}
		require.EqualValues(t, p1.SequenceId, p2.SequenceId)
		require.EqualValues(t, p1.GetMinterConfig(), p2.GetMinterConfig())
	}
}

func CompareMinterStates(t *testing.T, expected types.MinterState, state types.MinterState) {
	require.EqualValues(t, expected.SequenceId, state.SequenceId)
	require.Truef(t, expected.AmountMinted.Equal(state.AmountMinted), "expected.AmountMinted %s <> state.AmountMinted %s", expected.AmountMinted, state.AmountMinted)
	require.Truef(t, expected.RemainderToMint.Equal(state.RemainderToMint), "expected.RemainderToMint %s <> state.RemainderToMint %s", expected.RemainderToMint, state.RemainderToMint)
	require.Truef(t, expected.LastMintBlockTime.Equal(state.LastMintBlockTime), "expected.LastMintBlockTime %s <> state.LastMintBlockTime %s", expected.LastMintBlockTime.Local(), state.LastMintBlockTime.Local())
	require.Truef(t, expected.RemainderFromPreviousMinter.Equal(state.RemainderFromPreviousMinter), "expected.RemainderFromPreviousMinter %s <> state.RemainderFromPreviousMinter %s", expected.RemainderFromPreviousMinter, state.RemainderFromPreviousMinter)
}
