package configurer

import (
	"fmt"
	"testing"
)

type setupFn func(configurer Configurer) error

func baseSetup(configurer Configurer) error {
	if err := configurer.RunValidators(); err != nil {
		return err
	}
	return nil
}

func withIBC(setupHandler setupFn) setupFn {
	return func(configurer Configurer) error {
		if err := setupHandler(configurer); err != nil {
			return err
		}

		if err := configurer.RunIBC(); err != nil {
			return err
		}

		return nil
	}
}

func withUpgrade(setupHandler setupFn) setupFn {
	return func(configurer Configurer) error {
		if err := setupHandler(configurer); err != nil {
			return err
		}

		upgradeConfigurer, ok := configurer.(*UpgradeConfigurer)
		if !ok {
			return fmt.Errorf("to run with upgrade, %v must be set during initialization", &UpgradeConfigurer{})
		}

		if err := upgradeConfigurer.RunUpgrade(); err != nil {
			return err
		}

		return nil
	}
}

func StartDockerContainers(t *testing.T, startIBC, isDebugLogEnabled bool, upgradeSettings UpgradeSettings) (Configurer, error) {
	config, err := New(t, startIBC, isDebugLogEnabled, upgradeSettings)
	if err != nil {
		return nil, err
	}
	err = config.ConfigureChains()
	if err != nil {
		return nil, err
	}
	err = config.RunSetup()
	if err != nil {
		return nil, err
	}
	return config, nil
}
