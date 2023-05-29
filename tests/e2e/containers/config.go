package containers

// ImageConfig contains all images and their respective tags
// needed for running e2e tests.
type ImageConfig struct {
	InitRepository string
	InitTag        string

	C4eRepository string
	C4eTag        string

	RelayerRepository string
	RelayerTag        string
}

//nolint:deadcode
const (
	// Current Git branch c4e-chain repo/version. It is meant to be built locally.
	// It is used when skipping upgrade by setting startUpgrade in test suite to true).
	// This image should be pre-built with `make docker-build-debug` either in CI or locally.
	CurrentBranchC4eRepository = "chain4energy"
	CurrentBranchC4eTag        = "debug"
	// Pre-upgrade c4e-chain repo/tag to pull.
	// It must be built from previous branch.  startUpgrade in test suite should be unset
	// for this functionality to be used.
	previousVersionC4eRepository = "chain4energy-old-dev"
	previousVersionC4eTag        = "v1.2.0"
	// Pre-upgrade repo/tag for c4e-chain initialization (this should be one version below upgradeVersion)
	previousVersionInitRepository = "chain4energy-old-chain-init"
	previousVersionInitTag        = "v1.2.0"
	// If migration chain is used, we need to set the repo/tag for migration chain initialization to the penultimate version
	penultimateVersionC4eRepository = "chain4energy-old-dev"
	penultimateVersionC4eTag        = "v1.1.0"
	// Pre-upgrade repo/tag for c4e-chain initialization (this should be one version below upgradeVersion)
	penultimateVersionInitRepository = "chain4energy-old-chain-init"
	penultimateVersionInitTag        = "v1.1.0"
	// Hermes repo/version for relayer
	relayerRepository = "osmolabs/hermes"
	relayerTag        = "0.13.0"
)

// NewImageConfig returns ImageConfig needed for running e2e test.
// If startUpgrade is true, returns images for running the upgrade
func NewImageConfig(startUpgrade bool, migrationChaining bool) ImageConfig {
	config := ImageConfig{
		RelayerRepository: relayerRepository,
		RelayerTag:        relayerTag,
	}

	if !startUpgrade {
		// If upgrade is not tested, we do not need InitRepository and InitTag
		// because we directly call the initialization logic without
		// the need for Docker.
		config.C4eRepository = CurrentBranchC4eRepository
		config.C4eTag = CurrentBranchC4eTag
		return config
	}
	// If migration chain is used, we need to set the repo/tag for migration chain initialization to the penultimate version.
	// This is because the migration chain is initialized with the penultimate version and then upgraded to the current version
	if migrationChaining {
		config.C4eRepository = penultimateVersionC4eRepository
		config.C4eTag = penultimateVersionC4eTag
		config.InitRepository = penultimateVersionInitRepository
		config.InitTag = penultimateVersionInitTag
		return config
	}

	// If upgrade is tested, we need to utilize InitRepository and InitTag
	// to initialize older state with Docker
	config.InitRepository = previousVersionInitRepository
	config.InitTag = previousVersionInitTag

	// Upgrades are run at the time when upgrade height is reached
	// and are submitted via a governance proposal. Thefore, we
	// must start running the previous Chain4Energy version. Then, the node
	// should auto-upgrade, at which point we can restart the updated
	// Chain4Energy validator container.
	config.C4eRepository = previousVersionC4eRepository
	config.C4eTag = previousVersionC4eTag

	return config
}
