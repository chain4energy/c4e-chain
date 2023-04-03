# End-to-end Tests
### This module is based on Osmosis E2E testing suite
## Structure

### `e2e` Package

The `e2e` package defines an end-to-end testing suite used for full
 testing functionality. This package is decoupled from
depending on the Chain4Energy codebase. It initializes the chains for testing
via Docker files. As a result, the test suite may provide the desired
Chain4Energy version to Docker containers during the initialization. This
design allows for the opportunity of testing chain upgrades in the
future by providing an older Chain4Energy version to the container,
performing the chain upgrade, and running the latest test suite. When
testing a normal upgrade, the e2e test suite submits an upgrade proposal at
an upgrade height, ensures the upgrade happens at the desired height, and
then checks that operations that worked before still work as intended.

The file e2e_setup.go defines the testing suite and contains the
core bootstrapping logic that creates a testing environment via Docker
containers. A testing network is created dynamically with 4 test validators.

Files `*_test.go` contains the actual end-to-end integration tests
that utilize the testing suite.

Currently, there is a single IBC test in `e2e_test.go`.

Additionally, there is an ability to disable certain components
of the e2e suite. This can be done by setting the environment
variables. See "Environment variables" section below for more details.

## How It Works - Setting up e2e testing environment
The tests that are run can now be run in the following ways:

- Only base logic (`startUpgrade` and `startIbc` are both set to false in the test suite)
  - This is the most basic type of setup where a single chain is created
  - It simply spins up the desired number of validators on a chain (currently 4).
- IBC testing (`startIbc` is set to true in the test suite)
  - 2 chains are created connected by Hermes relayer
  - Upgrade Testing
  - 2 chains of the older Chain4Energy version are created, and
  connected by Hermes relayer
  - An additional full node is created after a chain has started.
  - This node is meant to state sync with the rest of the system.
- Upgrade testing (`startUpgrade` is set to true in the test suite)
  - CLI commands are run to create an upgrade proposal and approve it
  - Old version containers are stopped and the upgrade binary is added
  - Current branch Chain4Energy version is spun up to continue with testing

  This is done in `configurer/setup_runner.go` via function decorator design pattern
where we chain the desired setup components during configurer creation.
[Example](https://github.com/Chain4Energy-labs/Chain4Energy/blob/c5d5c9f0c6b5c7fdf9688057eb78ec793f6dd580/tests/e2e/configurer/configurer.go#L166)

  
## Running from current branch
1. Build current branch c4e-chain image  
```sh
make docker-build-debug
```
2. Run desired test suite by running:
```sh
make test-e2e-vesting
make test-e2e-params-change
make test-e2e-ibc:
```
Or to run all tests run:
```sh
make test-e2e
```

## Running from previous version
1. Build current branch c4e-chain image
```sh
make docker-build-debug
```
2. Build previous version c4e-chain image and init-chain image
```sh
make docker-build-old-chain
```
3. Run desired test suite by running:
```sh
make test-e2e-vesting
make test-e2e-params-change
make test-e2e-ibc:
```
Or to run all tests run:
```sh
make test-e2e
```

### Common Problems

Please note that if the tests are stopped mid-way, the e2e framework might fail to start again due to duplicated containers. Make sure that
containers are removed before running the tests again: `docker containers rm -f $(docker containers ls -a -q)`.

Additionally, Docker networks do not get auto-removed. Therefore, you can manually remove them by running `docker network prune`.


Also, from time to time there may be a problem with the cache.
Clearing the cache may help: `go clean --cache --mod-cache` 


As a last resort you can also clear all docker data (BE CAREFUL)
`docker system prune -af`