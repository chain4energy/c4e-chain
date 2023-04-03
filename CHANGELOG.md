<!--
Guiding Principles:

Changelogs are for humans, not machines.
There should be an entry for every single version.
The same types of changes should be grouped.
Versions and sections should be linkable.
The latest version comes first.
The release date of each version is displayed.
Mention whether you follow Semantic Versioning.

Usage:

Change log entries are to be added to the Unreleased section under the
appropriate stanza (see below). Each entry should ideally include a tag and
the Github issue reference in the following format:

* (<tag>) \#<issue-number> message

The issue numbers will later be link-ified during the release process so you do
not have to worry about including a link manually, but you can if you wish.

Types of changes (Stanzas):

"Features" for new features.
"Improvements" for changes in existing functionality.
"Deprecated" for soon-to-be removed features.
"Bug Fixes" for any bug fixes.
"Client Breaking" for breaking CLI commands and REST routes used by end-users.
"API Breaking" for breaking exported APIs used by developers building on SDK.
"State Machine Breaking" for any changes that result in a different AppState 
given same genesisState and txList.
Ref: https://keepachangelog.com/en/1.0.0/
-->

# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Unreleased

### Features

### Bug fixes

### Misc Improvements

## [v1.2.0](https://github.com/chain4energy/c4e-chain/releases/tag/v1.2.0) - 2023-04-03

**Upgrades**

 - Bumped cosmos-sdk version to v0.46.10
- Bumped tendermint version to v0.34.26
- Bumped go version to v1.19
- Bumped ibc-go to v5.2.0
- Bumped ics23 to v0.9.0

**Improvements**

- Made the app independent of the ignite:
- created params directory which holds chin encoding, denom and address prefix config
- removed cosmoscmd App interface
- created CMD functions for correct chain start and initialization
- deleted module message handlers
- Added [end-to-end testing framework](https://github.com/chain4energy/c4e-chain/tree/master-1.2.0/tests/e2e) that can be used for full testing functionality (this framework is based on Osmosis E2E testing suite)
- Migrated all modules params from using x/params module to using simple KVStore [(cosmos-sdk ADR)](https://docs.cosmos.network/main/architecture/adr-046-module-params)
- (x/cfevesting) Renamed address field of the AccountVestingPools object to owner
- x/cfeminter Module params refactoring:
- cfeminter params structure changed - removed MinterConfig and moved mint\_denom and Minter array directly to Params
- changed the configuration logic for individual minters - instead of setting one of the LinearMinting or ExponentialStepMinting fields to a specific value, and the other to null, the configuration now includes one config field that accepts the MinterConfigI interface. LinearMinting and ExponentialStepMinting implement this interface which allows to set one specific configuration for minter.
- (x/cfevesting) Vesting cession:
- added MsgSplitVesting to split the vesting and transfer it to the second account
- added MsgMoveAvailableVesting and MsgMoveAvailableVestingByDenoms to move available vesting from one account to another
- (x/cfevesting) Vesting pools and accounts migration:
- Founders accounts vesting period extension by one year
   - from: lockup: 1 year, vesting: 2 years
   - to: lockup: 2years, vesting: 2 years
- Splited ValidatorsVestingPool into 5 smaller pools:
- Validator round pool
- initially locked - 8498690 C4E
- lock end - 2024-12-26 00:00
- vesting type** (parameters of vesting account created from this vesting pool) - Validator round
   - lockup period - 274 days (~9 months)
   - vesting period - 548 days (~18 months)
   - initial free tokens percentage - 5%
- VC round pool
- initially locked - 15000000 C4E
- lock end - 2025-09-22 14:00
- vesting type (parameters of vesting account created from this vesting pool) - VC round
   - lockup period - 548 days (~18 months)
   - vesting period - 548 days (~18 months)
   - initial free tokens percentage - 5%
- Early-bird (private) round pool
- initially locked - 8000000 C4E
- lock end - 2024-12-22 14:00
- vesting type (parameters of vesting account created from this vesting pool) - Early-bird round
   - lockup period - 456 days (~15 months)
   - vesting period - 365 days (~12 months)
   - initial free tokens percentage - 10%
- Public round pool
- initially locked - 9000000 C4E
- lock end - 2024-03-22 14:00
- vesting type (parameters of vesting account created from this vesting pool) - Public round
   - lockup period - 274 days (~9 months)
   - vesting period - 274 days (~9 months)
   - initial free tokens percentage - 15%
- Strategic reserve short term round pool
- initially locked - 40000000 C4E
- lock end - 2024-09-22 14:00
- vesting type (parameters of vesting account created from this vesting pool) - Strategic reserve short term round
   - lockup period - 365 days (~12 months)
   - vesting period - 365 days (~12 months)
   - initial free tokens percentage - 20%
- Updated genesis vesting pools and accounts tracking for more accurate circulating supply calculation

**Bug fixes**

- (x/cfevesting) If there are any vesting pools, changing the vesting denom is not possible
- (x/cfeminter) When changing minters via proposal, it checks if there is a minter in the new configuration with a sequence\_id that is currently in cfeminter state

**Tests that have been carried out**

- Simulation tests
- Performance/stability tests
- Manual E2E tests
- Automatic E2E tests
- Unit tests

## [v1.1.0](https://github.com/chain4energy/c4e-chain/releases/tag/v1.1.0) - 2023-01-24
### Misc Improvements
1. Distribution
   - 1st token distribution ready version
   - tokens distribution mechanism based on the list of configured subdistributors [README](https://github.com/chain4energy/c4e-chain/blob/master/x/cfedistributor/README.md).
   - new params structure
   - new state structure
   - extended validation
   - new emit events types
2. Minting
   - 1st minting (inflation) ready version
   - new params structure
       * Linear Minting type
       * Exponential Ste pMinting type
   - new state structure
   - extended validation
   - new emit events types
3. Vesting
   - vesting pool params changed
   - extended validation
   - new emit events types
     vesting type percentage of tokens that are released initially
4. Simulation tests
5. Performance/stability tested
6. Other
   - rest api versioning

### Bug fixes
- vesting pool sent tokens calculation bug
- cfeminter init genesis time e’rror


## [v1.0.1](https://github.com/chain4energy/c4e-chain/releases/tag/v1.0.1) - 2022-11-24

* **Upgrade Cosmos SDK for [Dragonbarry patch](https://forum.cosmos.network/t/ibc-security-advisory-dragonberry/7702)**

## [v1.0.0](https://github.com/chain4energy/c4e-chain/releases/tag/v1.0.0) - 2022-09-22

Initial Release!
