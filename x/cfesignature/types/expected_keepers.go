package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// AccountKeeper defines the expected account keeper (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI

	// Return a new account with the next account number and the specified address. Does not save the new account to the store.
	NewAccountWithAddress(sdk.Context, sdk.AccAddress) types.AccountI

	// Set an account in the store.
	SetAccount(sdk.Context, types.AccountI)

	GetModuleAddress(name string) sdk.AccAddress
	GetModuleAccount(ctx sdk.Context, name string) types.ModuleAccountI

	// TODO remove with genesis 2-phases refactor https://github.com/cosmos/cosmos-sdk/issues/2862
	SetModuleAccount(sdk.Context, types.ModuleAccountI)
}
