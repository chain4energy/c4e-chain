# Chain4Energy distributor module - cfedistributor

## Abstract

Chain4Energy distributor module provides functionality of tokens distribution mechanism. 
Tokens distribution mechanism sends tokens from source accounts to destination accounts according to configuration. 

## Contents

1. **[Concept](#concepts)**
2. **[Params](#parameters)**
3. **[State](#state)**
4. **[Events](#events)**
5. **[Queries](#queries)**
6. **[Invariants](#invariants)**
7. **[Genesis validations](#genesis-validations)**

## Concepts

The purpose of `cfedistributor` module is to provide functionality of flexible token distribution mechanism.

### Tokens distribition mechanism

Tokens distribution mechanism is based on the list of configured subdistributors.
Tokens distribution mechanism iterates through all subdistributors in predefined order and executes each subdistributor per each block. 

### Subdistributor

Subdistributor is responsible for sending coins from its source accounts to its destination accounts or for burning tokens accordingly to configured share percentage.

See the grapical represenation of subdistributor.

![Subdistributor](./../../docs/modules/cfedistributor/subdistributor.svg?raw=1)

Subdistributors are executed in predefined order so destinations of one subdistributor can become sources of another subdistributor.

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
* Main module account - Chain4Energy distributor module account. 
Some other modules can send tokens directly to Chain4Energy distributor module. 
It is required to define one subdistributor where main module account is the source to prevent token accumulation.
* Base account - Standard Cosmos SDK account configured as account address.
* Module account - Cosmos SDK module account configured as module account name.
* Internal account - helper virtual account internal to Chain4Energy distributor module. Useful in case of designing more complicated token flow.
* Burn destination - tokens burner.

#### Example

Let's consider an example of tokens distribution. In our example we have following token sources:
* inflation
* transaction fees
* some module functionality fees. Our fictional module provides some functionality for which 
user is required to pay some fee and part of this fee is shared with the blockchain.

Inflation distribution:
* 60% to distribution module for standard validators rewarding functionality
* 5% to development fund
* 35% to incentive booster pools

Transaction fees distribution:
* 80% to distribution module for standard validators rewarding functionality
* 5% to burn
* 15% to incentive booster pools

Fictional module functionality fees distribution:
* 50% to distribution module for standard validators rewarding functionality
* 30% to development fund
* 20% to incentive booster pools

We also have following incentive booster pools distribution:
* 65% to governance booster pool
* 35% to weekend booster pool

See the graphical representation of this distribution.

![ExampleDistribution](./../../docs/modules/cfedistributor/example-distribution.svg?raw=1)

Let's also assume that:
* inflation is minted directly to cfedistributor module account
* cosmos sdk distribution module fetches tokens from module account named "validators_rewards"
* development fund is base account with address "c4edwijhdhwqu43efvc3543ec34c2erc342dw"
* governance booster pool is module account named "governance_booster"
* weekend booster pool is module account named "weekend_booster"
* fictional module fees are stored in module account named "fictional_module_fee_collector"

Then we can model our distribution flow with following subdistributors:
* inflation subdistributor

![inflation subdistributor](./../../docs/modules/cfedistributor/example-inflation-subdistributor.svg?raw=1)

* transaction fees subdistributor

![transaction fees subdistributor](./../../docs/modules/cfedistributor/example-tx-fees-subdistributor.svg?raw=1)

* module fees subdistributor

![module fees subdistributor](./../../docs/modules/cfedistributor/example-module-fees-subdistributor.svg?raw=1)

* incentive boosters subdistributor

![incentive boosters subdistributor](./../../docs/modules/cfedistributor/example-boosters-subdistributor.svg?raw=1)

## Parameters

The Chain4Energy distributor module contains the following configurations parameters:

| Key                    | Type                          | Description                     |
|------------------------|-------------------------------|---------------------------------|
| sub_distributors       | List of Subdistributor type   | list of defined subdistributors |

### Subdistributor type

| Param        | Type                 | Description                       |
|--------------|----------------------|-----------------------------------|
| name         | string               | unique name of the subdistributor |
| sources      | List of Account type | list of source accounts           |
| destinations | Destinations type    | destinations definition           |

### Destinations type

| Param         | Type                          | Description                                                                 |
|---------------|-------------------------------|-----------------------------------------------------------------------------|
| primary_share | Account type                  | primary destination - all remaining tokens from others shares are sent here |
| burn_share    | sdk.Dec type                  | share to burn (0-1)                                                         |
| shares        | List of DestinationShare type | List of destination accounts with share percentage                          |

### DestinationShare type

| Param       | Type           | Description              |
|-------------|----------------|--------------------------|
| name        | string         | unique name of the share |
| destination | Account type   | destination account      |
| share       | sdk.Dec        | share percentage (0-1)   |

### Account type

| Param     | Type          | Description                                                                                                                                                                                                          |
|-----------|---------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| type      | enum string   | account type:<br />- MAIN - main module account<br />- MODULE_ACCOUNT - module account<br />- BASE_ACCOUNT - base account<br />- INTERNAL_ACCOUNT - cfedistributor internal account                                  |
| id        | string        | account identifier dependant on the type::<br />- MAIN - empty<br />- MODULE_ACCOUNT - module account name<br />- BASE_ACCOUNT - base account address<br />- INTERNAL_ACCOUNT - cfedistributor internal account name |

### Example params

See the configuration params for **[example](#example)** from **[Concept](#concepts)** section

```json

{
    "sub_distributors": [
        {
            "name": "inflation_subdistributor", 
            "sources": [
                {
                "id": "",
                "type": "MAIN"
                }
            ],
            "destinations": {
                "primary_share": {
                    "id": "validators_rewards",
                    "type": "MODULE_ACCOUNT"
                },
                "shares": [
                    {
                        "destination": {
                            "id": "c4edwijhdhwqu43efvc3543ec34c2erc342dw",
                            "type": "BASE_ACCOUNT"
                        },
                        "name": "inflation_development_fund_share",
                        "share": "0.05"
                    },
                    {
                        "destination": {
                            "id": "incentive_boosters",
                            "type": "INTERNAL_ACCOUNT"
                        },
                        "name": "inflation_incentive_boosters_share",
                        "share": "0.35"
                    }
                ],
                "burn_share": "0.0"
            }
        },
        {
            "name": "transaction fees subdistributor", 
            "sources": [
                {
                "id": "fee_collector",
                "type": "MODULE_ACCOUNT"
                }
            ],
            "destinations": {
                "primary_share": {
                    "id": "validators_rewards",
                    "type": "MODULE_ACCOUNT"
                },
                "shares": [
                    {
                        "destination": {
                            "id": "incentive_boosters",
                            "type": "INTERNAL_ACCOUNT"
                        },
                        "name": "txs_incentive_boosters_share",
                        "share": "0.15"
                    }
                ],
                "burn_share": "0.05"
            }
        },
        {
            "name": "module_fees_subdistributor", 
            "sources": [
                {
                "id": "fictional_module_fee_collector",
                "type": "MODULE_ACCOUNT"
                }
            ],
            "destinations": {
                "primary_share": {
                    "id": "validators_rewards",
                    "type": "MODULE_ACCOUNT"
                },
                "shares": [
                    {
                        "destination": {
                            "id": "c4edwijhdhwqu43efvc3543ec34c2erc342dw",
                            "type": "BASE_ACCOUNT"
                        },
                        "name": "module_development_fund_share",
                        "share": "0.3"
                    },
                    {
                        "destination": {
                            "id": "incentive_boosters",
                            "type": "INTERNAL_ACCOUNT"
                        },
                        "name": "module_incentive_boosters_share",
                        "share": "0.2"
                    }
                ],
                "burn_share": "0.0"
            }
        },
        {
            "name": "incentive boosters subdistributor", 
            "sources": [
                {
                "id": "incentive_boosters",
                "type": "INTERNAL_ACCOUNT"
                }
            ],
            "destinations": {
                "primary_share": {
                    "id": "governance_booster",
                    "type": "MODULE_ACCOUNT"
                },
                "shares": [
                    {
                        "destination": {
                            "id": "weekend_boosters",
                            "type": "INTERNAL_ACCOUNT"
                        },
                        "name": "weekend_boosters_share",
                        "share": "0.35"
                    }
                ],
                "burn_share": "0.0"
            }
        }
    ]
}

```

## State

Chain4Energy distributor module state contains decimal amounts left from previous block that were impossible to send due value less than 1 token.
Module state contains following data:

| Key          | Type                 | Description                                                                                               |
|--------------|----------------------|-----------------------------------------------------------------------------------------------------------|
| states       | List of State type   | list of states for burn and per each destination account in all subdistributors (except main distributor) |

### State type

| Param     | Type                                                 | Description                                                   |
|-----------|------------------------------------------------------|---------------------------------------------------------------|
| account   | Account type (see **[Account type](#account-type)**) | destination account or empty in case of burn flag set to true |
| burn      | bool                                                 | specifies if this is burn destination state                   |
| remains   | DecCoin                                              | list of coins to distribute left by previous block            |

### Example state

See the state for **[example](#example)** from **[Concept](#concepts)** section

```json

[
    {
        "account": {
            "id": "validators_rewards",
            "type": "MODULE_ACCOUNT"
        },
        "burn": false,
        "remains": [
            {
                "denom": "uc4e",
                "amount": "0.900000000000000000"
            }
        ]
    },
    {
        "account": {
            "id": "c4edwijhdhwqu43efvc3543ec34c2erc342dw",
            "type": "BASE_ACCOUNT"
        },
        "burn": false,
        "remains": [
            {
                "denom": "uc4e",
                "amount": "0.359000000000000000"
            }
        ]
    },
    {
        "account": {
            "id": "incentive_boosters",
            "type": "INTERNAL_ACCOUNT"
        },
        "burn": false,
        "remains": []
    },
    {
        "account": {
            "id": "governance_booster",
            "type": "MODULE_ACCOUNT"
        },
        "burn": false,
        "remains": [
            {
                "denom": "uc4e",
                "amount": "0.582000000000000000"
            }
        ]
    },
    {
      "account": {
            "id": "weekend_booster",
            "type": "MODULE_ACCOUNT"
      },
      "burn": false,
      "remains": [
            {
                "denom": "uc4e",
                "amount": "0.359000000000000000"
            }
      ]
    },
    {
        "burn": true,
        "remains": [
            {
                "denom": "uc4e",
                "amount": "0.800000000000000000"
            }
        ]
    }
]

```

## Events

Chain4Energy distributor module emits the following events:

### BeginBlockers

#### Tokens distribution

| Type             | Attribute Key         | Description       |
|------------------|-----------------------|-------------------|
| Distribution     | Distribution type     | Distribution data |
| DistributionBurn | DistributionBurn type | Burn data         |

##### Distribution type

Distribution type represents one send operation to one destination in one block

| Param          | Type                                                         | Description                                                                            |
|----------------|--------------------------------------------------------------|----------------------------------------------------------------------------------------|
| subdistributor | string                                                       | Name of the subdisributor                                                              | 
| share_name     | string                                                       | Name of the DestinationShare (see **[DestinationShare type](#destinationshare-type)**) | 
| sources        | list of Account type (see **[Account type](#account-type)**) | list of sources                                                                        |
| destination    | Account type (see **[Account type](#account-type)**)         | destination                                                                            |
| amount         | DecCoins                                                     | coins sent to destination                                                              |

##### DistributionBurn type

DistributionBurn type represents one burn operation

| Param          | Type                                                         | Description                     |
|----------------|--------------------------------------------------------------|---------------------------------|
| subdistributor | string                                                       | Name of the subdisributor       | 
| sources        | list of Account type (see **[Account type](#account-type)**) | list of sources                 |
| amount         | DecCoins                                                     | coins burned                    |

## Queries

### Params query

Queries the module params.

See example response:

```json
{
  "params": {
    "sub_distributors": [
      {
        "name": "tx_fee_distributor",
        "sources": [
          {
            "id": "fee_collector",
            "type": "MODULE_ACCOUNT"
          }
        ],
        "destinations": {
          "primary_share": {
            "id": "c4e_distributor",
            "type": "MAIN"
          },
          "shares": [],
          "burn_share": "0.000000000000000000"
        }
      },
      {
        "name": "inflation_and_fee_distributor",
        "sources": [
          {
            "id": "c4e_distributor",
            "type": "MAIN"
          }
        ],
        "destinations": {
          "primary_share": {
            "id": "validators_rewards_collector",
            "type": "MODULE_ACCOUNT"
          },
          "shares": [
            {
              "name": "development_fund",
              "share": "0.050000000000000000",
              "destination": {
                "id": "c4e10ep2sxpf2kj6jsdcs234edkuf9sf9xqq3sl",
                "type": "BASE_ACCOUNT"
              }
            },
            {
              "name": "usage_incentives",
              "share": "0.350000000000000000",
              "destination": {
                "id": "usage_incentives_collector",
                "type": "INTERNAL_ACCOUNT"
              }
            }
          ],
          "burn_share":  "0.000000000000000000"
        }
      },
      {
        "name": "usage_incentives_distributor",
        "sources": [
          {
            "id": "usage_incentives_collector",
            "type": "INTERNAL_ACCOUNT"
          }
        ],
        "destinations": {
          "primary_share": {
            "id": "c4e1q5vgy0r3scsdc32dcewkl8nwmfe2mgr6g0jlph",
            "type": "BASE_ACCOUNT"
          },
          "shares": [
            {
              "name": "green_energy_booster",
              "share": "0.340000000000000000",
              "destination": {
                "id": "green_energy_booster_collector",
                "type": "MODULE_ACCOUNT"
              }
            },
            {
              "name": "governance_booster",
              "share": "0.330000000000000000",
              "destination": {
                "id": "governance_booster_collector",
                "type": "MODULE_ACCOUNT"
              }
            }
          ],
          "burn_share": "0.000000000000000000"
        }
      }
    ]
  }
}
```
### States query

Queries the module state.

See example response:

```json
{
  "states": [
    {
      "account": {
        "id": "c4e10ep2ssdfwefcscaewdedscs9xqqqdwqee3sl",
        "type": "BASE_ACCOUNT"
      },
      "burn": false,
      "remains": [
        {
          "denom": "uc4e",
          "amount": "0.900000000000000000"
        }
      ]
    },
    {
      "account": {
        "id": "c4e1q5vgy0r3w9q4ccsdcds23422mgr6g0jlph",
        "type": "BASE_ACCOUNT"
      },
      "burn": false,
      "remains": [
        {
          "denom": "uc4e",
          "amount": "0.359000000000000000"
        }
      ]
    },
    {
      "account": {
        "id": "governance_booster_collector",
        "type": "MODULE_ACCOUNT"
      },
      "burn": false,
      "remains": [
        {
          "denom": "uc4e",
          "amount": "0.359000000000000000"
        }
      ]
    },
    {
      "account": {
        "id": "green_energy_booster_collector",
        "type": "MODULE_ACCOUNT"
      },
      "burn": false,
      "remains": [
        {
          "denom": "uc4e",
          "amount": "0.582000000000000000"
        }
      ]
    },
    {
      "account": {
        "id": "usage_incentives_collector",
        "type": "INTERNAL_ACCOUNT"
      },
      "burn": false,
      "remains": []
    },
    {
      "account": {
        "id": "validators_rewards_collector",
        "type": "MODULE_ACCOUNT"
      },
      "burn": false,
      "remains": [
        {
          "denom": "uc4e",
          "amount": "0.800000000000000000"
        }
      ]
    }
  ],
  "coins_on_distributor_account": [
    {
      "denom": "uc4e",
      "amount": "3"
    }
  ]
}
```

## Invariants

### Non-Negative Coin State Invariant

Invariant validates module state. Checks if all coins states of all destinations are non-negative

### State Sum Balance Check Invariant

Invariant validates module state. Checks sum of all coins states of one denom of all destinations is always 
integer value and is equal to cfedistributor module account balance

## Genesis validations

[//]: # (TODO)
