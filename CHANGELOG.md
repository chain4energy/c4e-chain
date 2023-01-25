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
