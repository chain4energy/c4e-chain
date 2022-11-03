# Chain4Energy distributor module - cfedistributor

## Abstract

Chain4Energy distributor module provides functionality of tokens distribution mechanism. Tokens distribition mechanism sends tokens from source accounts to destination accounts according to configuration. 

## Contents

1. **[Concept](#concepts)**
2. **[State](#state)**
4. **[Events](#events)**
6. **[Params](#parameters)**
8. **[Queries](#queries)**

## Concepts

The purpose of `cfedistributor` module is to provide functionality of flexible token distribution mechanism.

### Tokens distribition mechanism

Tokens distribition mechanism is based on the list of configured subdistributors.
Tokens distribition mechanism iterates through all subdistributors in predefined order and executes each subdistributor per each block. 

### Subdistributor

Subdistributor is responsible for sending coins from it's source accounts to it's destination accounts or for burning tokens according to configured share percentage.

See the grapical represenation of subdistributor.

![Subdistributor](./../../docs/modules/cfedistributor/subdistributor.svg?raw=1)

Subdistributors are executed in predefined order so destinations of one subdistributor can became sources of another subdistributor.

#### Subdistributor sources and destinations

Subdistributor supports several types of sources:
* main module account
* base account
* module account
* internal account

Subdistributor supports several types of destinations:
* main module account
* base account
* module account
* internal account
* Burn destination

where:
* Main module account - Chain4Energy distributor module account. Some other modules can send tokens directly to Chain4Energy distributor module. It is required to define one subdistributor where main module account is the source to prevent token accumulation.
* Base account - Standard Cosmos SDK account configured as account address.
* Module account - Cosmos SDK module account configured as module account name.
* Internal account - helper virtual account internal to Chain4Energy distributor module. Useful in case of designing more complicateed token flow.
* Burn destination - tokens burner.

#### Example

Let's consider an example of tokens distribution. In our example we have following token sources:
* inflation
* transaction fees
* some module functionality fees. Our fictional module provides some functionlity for which user is required to pay some fee and part of this fee is shared with the blockchain.

Inflation distribution:
* 60% to distribution module for standard validators rewarding funcionality
* 5% to development fund
* 35% to incentive booster pools

Transaction fees distribution:
* 80% to distribution module for standard validators rewarding funcionality
* 5% to burn
* 15% to incentive booster pools

Fictional module functionality fees distribution:
* 50% to distribution module for standard validators rewarding funcionality
* 30% to development fund
* 20% to incentive booster pools

We also have following incentive booster pools distribution:
* 65% to governance booster pool
* 35% to weekend booster pool

See the grapical represenation of this distribution.

![ExampleDistribution](./../../docs/modules/cfedistributor/example-distribution.svg?raw=1)

Let's also assume that:
* inflation is minted directly to cfedistributor module account
* cosmos sdk distribution module fetches tokens from module account named "validators_rewards"
* development fund is base account with address "c4edwijhdhwqu43efvc3543ec34c2erc342dw"
* governance booster pool is module account named "governance_booster"
* weekend booster pool is module account named "weekend_booster"
* fictional module fees are stored in module account named "fictional_module_fee_collector"

Than we can model our distribution flow with following subdistributors:
* inflation subdistributor

![inflation subdistributor](./../../docs/modules/cfedistributor/example-inflation-subdistributor.svg?raw=1)

* transaction fees subdistributor

![transaction fees subdistributor](./../../docs/modules/cfedistributor/example-tx-fees-subdistributor.svg?raw=1)

* module fees subdistributor

![module fees subdistributor](./../../docs/modules/cfedistributor/example-module-fees-subdistributor.svg?raw=1)

* incentive boosters subdistributor

![incentive boosters subdistributor](./../../docs/modules/cfedistributor/example-boosters-subdistributor.svg?raw=1)


## State

## Events

The incentives module emits the following events:

### Handlers

### BeginBlockers

#### Tokens distribution

| Type         | Attribute Key | Attribute Value |
| ------------ | ------------- | --------------- |

## Parameters

The Chain4Energy distributor module contains the following parameters:

| Key                  | Type   | Example  |
| -------------------- | ------ | -------- |


## Queries

