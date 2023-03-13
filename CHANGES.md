## [v1.2.0](https://github.com/chain4energy/c4e-chain/releases/tag/v1.2.0) - [UPDATE_DATE]

### Upgrades
* Bump cosmos-sdk version to [v0.46.10](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.46.10)
* Bump tendermint version to [v0.34.26](https://github.com/informalsystems/tendermint/releases/tag/v0.34.26)
* Bump go version to [v1.19](https://github.com/golang/go/releases/tag/go1.19)
* Bump ibc-go to [v5.2.0](https://github.com/cosmos/ibc-go/releases/tag/v5.2.0)
* Bump ics23 to [v0.9.0](https://github.com/cosmos/ics23/releases/tag/go%2Fv0.9.0)

### Improvements
* Make the app independent of the ignite:
  * create params directory which holds chin encoding, denom and address prefix config
  * remove cosmoscmd App interface
  * create CMD functions for correct chain start and initialization
  * delete module message handlers
* Add end-to-end testing framework that can be used for full testing functionality
* Migrate all modules params from using `x/params` module to using simple KVStore [(cosmos-sdk ADR)](https://docs.cosmos.network/main/architecture/adr-046-module-params ) 
* (x/cfevesting) Rename `address` field of the `AccountVestingPoolsowner` object to owner
* `x/cfeminter` Module params refactoring:
  * cfeminter `params` structure change - removed `MinterConfig` and moved `mint_denom` and `Minter` array directly to `Params`
  * change the configuration logic for individual minters - instead of setting one of the `LinearMinting` or `ExponentialStepMinting`
  fields to a specific value, and the other to null, the configuration now includes one config field that accepts the `MinterConfigI` interface.
  `LinearMinting` and `ExponentialStepMinting` implement this interface which allows to set one specific configuration for minter.
* (x/cfevesting) Vesting cession:
  * add `MsgSplitVesting` to split the vesting and transfer it to the second account
  * add `MsgMoveAvailableVesting` and `MsgMoveAvailableVestingByDenoms` to move available vesting from one account to another
* (x/cfevesting) Vesting pools and accounts migration:
  * migrate 

### Bug fixes
* (x/cfevesting) If there are any vesting pools, changing the vesting `denom` is not possible
* (x/cfeminter) When changing minters via proposal, it checks if there is a minter in the new configuration with a 
`sequence_id` that is currently in `cfeminter` state
