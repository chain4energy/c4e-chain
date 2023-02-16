package cosmossdk

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func CheckInvariant(t require.TestingT, ctx sdk.Context, invariant sdk.Invariant, failed bool, message string) {
	msg, wasFailed := invariant(ctx)
	require.EqualValues(t, failed, wasFailed)
	require.EqualValues(t, message, msg)
}

func ValidateManyInvariants(t require.TestingT, ctx sdk.Context, invariants []sdk.Invariant) {
	for i := 0; i < len(invariants); i++ {
		msg, failed := invariants[i](ctx)
		require.False(t, failed, "Invariant failed - "+msg)
	}
}
