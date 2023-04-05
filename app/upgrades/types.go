package upgrades

import (
	cfeairdropkeeper "github.com/chain4energy/c4e-chain/x/cfeairdrop/keeper"
	cfevestingkeeper "github.com/chain4energy/c4e-chain/x/cfevesting/keeper"
	store "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type AppKeepers interface {
	GetAccountKeeper() *authkeeper.AccountKeeper
	GetBankKeeper() *bankkeeper.Keeper
	GetParamKeeper() *paramsKeeper.Keeper
	GetC4eVestingKeeper() *cfevestingkeeper.Keeper
}

// BaseAppParamManager defines an interrace that BaseApp is expected to fullfil
// that allows upgrade handlers to modify BaseApp parameters.
type BaseAppParamManager interface {
	GetConsensusParams(ctx sdk.Context) *abci.ConsensusParams
	StoreConsensusParams(ctx sdk.Context, cp *abci.ConsensusParams)
}

type AppKeepers interface {
	GetKeeper() *cfeairdropkeeper.Keeper
	GetAccountKeeper() *authkeeper.AccountKeeper
	GetBankKeeper() *bankkeeper.Keeper
}

// Upgrade defines a struct containing necessary fields that a SoftwareUpgradeProposal
// must have written, in order for the state migration to go smoothly.
// An upgrade must implement this struct, and then set it in the app.go.
// The app.go will then define the handler.
type Upgrade struct {
	// Upgrade version name, for the upgrade handler, e.g. `v2`
	UpgradeName string

	// CreateUpgradeHandler defines the function that creates an upgrade handler
	CreateUpgradeHandler func(*module.Manager, module.Configurator, BaseAppParamManager, AppKeepers) upgradetypes.UpgradeHandler

	// Store upgrades, should be used for any new modules introduced, new modules deleted, or store names renamed.
	StoreUpgrades store.StoreUpgrades
}
