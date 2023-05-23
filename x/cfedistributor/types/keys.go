package types

var (
	ParamsKey      = []byte{0x00}
	StateKeyPrefix = []byte{0x04}
)

const (
	// ModuleName defines the module name
	ModuleName = "cfedistributor"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cfedistributor"

	DistributorMainAccount = "distributor_main_account"

	ValidatorsRewardsCollector = "validators_rewards_collector"

	GreenEnergyBoosterCollector = "green_energy_booster_collector"

	GovernanceBoosterCollector = "governance_booster_collector"

	BurnStateKey = "burn_state_key"

	BurnDestination = "burn_destination"

	DenomToTrace = "uc4e" // TODO should be in module params - do omowienia - mozna tez wziac to z jakiegos constanta globalnego albo pobrane z konfiguracji mintingu
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
