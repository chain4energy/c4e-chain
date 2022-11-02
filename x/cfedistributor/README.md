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

http://www.plantuml.com/plantuml/proxy?src=(../../docs/modules/cfedistributor/subdistributor.puml)

-------------                                     x% share     -----------------
|  Source 1 |----->---                        ---->------------| Destination 1 |
-------------        |                        |                -----------------
                     |                        |       
-------------        |                        |   y% share     -----------------
|  Source 2 |----->--|                        |--->------------| Destination 3 |
-------------        |                        |                -----------------
                     |------->------->------>-|
      .              |                        |                        .
      .              |                        |                        .
      .              |                        |                        .
                     |                        |
-------------        |                        |   z% share     -----------------
|  Source n |----->---                        |--->------------| Destination 1 |
-------------                                 |                -----------------
                                              |
                                              |   v% share     -----------------
                                              ---->------------|     Burn      |
                                                               -----------------

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

