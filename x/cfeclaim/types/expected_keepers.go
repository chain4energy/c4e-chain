package types

import (
	"cosmossdk.io/math"
	cfevestingtypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	"time"
)

// AccountKeeper defines the expected account keeper interface
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	GetModuleAccount(ctx sdk.Context, moduleName string) authtypes.ModuleAccountI
	SetModuleAccount(ctx sdk.Context, macc authtypes.ModuleAccountI)
	NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	BlockedAddr(addr sdk.AccAddress) bool
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}

// FeeGrantKeeper defines the expected feegrant keeper interface
type FeeGrantKeeper interface {
	GrantAllowance(ctx sdk.Context, granter, grantee sdk.AccAddress, feeAllowance feegrant.FeeAllowanceI) error
	GetAllowance(ctx sdk.Context, granter, grantee sdk.AccAddress) (feegrant.FeeAllowanceI, error)
}

// StakingKeeper defines the expected feegrant keeper interface
type StakingKeeper interface {
	BondDenom(ctx sdk.Context) (res string)
}

// VestingKeeper defines the expected feegrant keeper interface
type CfeVestingKeeper interface {
	Denom(ctx sdk.Context) (res string)
	AddVestingPoolReservation(ctx sdk.Context, owner string, vestingPoolName string, reservationId uint64, amout math.Int) error
	RemoveVestingPoolReservation(ctx sdk.Context, owner string, vestingPoolName string, reservationId uint64, amout math.Int) error
	SendReservedToNewVestingAccount(ctx sdk.Context, owner string, toAddr string, vestingPoolName string, amount math.Int, reservationId uint64, free sdk.Dec, lockupPeriod time.Duration, vestingPeriod time.Duration) error
	SendToPeriodicContinuousVestingAccountFromModule(ctx sdk.Context, moduleName string, userAddress string, amount sdk.Coins, free sdk.Dec, startTime int64, endTime int64) (periodId uint64, periodExists bool, err error)
	MustGetVestingPoolReservation(ctx sdk.Context, owner string, vestingPoolName string, reservationId uint64) (*cfevestingtypes.VestingPoolReservation, error)
	MustGetVestingTypeForVestingPool(ctx sdk.Context, address string, vestingPoolName string) (*cfevestingtypes.VestingType, error)
}
