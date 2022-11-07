# Chain4Energy minter module - cfeminter

## Abstract

Chain4Energy minter module provides functionality of controlled token emissions. Tokens are emited according to configuration. 

## Contents

1. **[Concept](#concepts)**
2. **[Params](#parameters)**
3. **[State](#state)**
4. **[Events](#events)**
5. **[Queries](#queries)**
6. **[Invariants](#invariants)**
7. **[Genesis validations](#genesis-validations)**

## Concepts

The purpose of `cfedistributor` module is to provide token emission mechanism.

### Token emission mechanism

Token emission mechanism mints calculated amount of tokens per each block. Token amount is calculated according to cfeminter module configuration params. 
Tokens minting process is divided into periods where each period has fully separated minting rules. Those periods rules are difined within cfeminter module configuration params.
Simply, mintining process configuration is a list of ordered minting periods, where each period has its own start and end time (end time for last period is not required, in that case last minting period works infinitely).

### Miniting periods

Minting period is a period of time when configured minting rules apply. Ordered list of minting periods deifnes whole token emmision process.
End time of one period is a start time of the next period on the periods list.
Each minting pariond has minter type assigned. 
Last minting period on the list must be defined to work indefinitely. (must have no end time)

### Miniter type

Minter type defines general rules of token emission. Each minter type has its specific set of parameters modifying token emission. Parameters are set per minting period.
Currently cfeminter module suppoerts following minter types:
- no miniting
- time linear minter
- periodic reduction minter

#### No minitng

No minting is simple minter type that mints nothing.
This minter type has no parameters.

#### time linear minter

Time linear minter is block time based minter type. It mints configured amount of tokens within minting period lineary.
This minter requires period with end time since given amount of token needs to be minted in finite time period. So this minter type cannot be configured as a type of last period.

Minter type parameters:
* amount - amount of tokens to mint within minter period

#### periodic reduction minter

Periodic reduction minter is block time based minter type. It mints configured amount of tokens within minting period, where it divides this period into smaller sub periods of equal lenght.
Then within one sub period expected amount is minted, lineary. Expected amount of subperiod minted tokens is equal to tokens minted by prevoius subperiod multiplied by configured factor.
For example initial period amount is 40 milions, multiplying factor set to 0.5 and periond length is one year, then:
* 1st subperiod (1st year) mints 40 millions lineary 
* 2nd subperiod (2nd year) mints 20 millions lineary
* 3rd subperiod (3rd year) mints 10 millions lineary
* 4th subperiod (4th year) mints 5 millions lineary
and so on.

This minter can mint infinitely.

Minter type parameters: //TODO better params names
* mint period - period of time mint amount is emitted
* mint amount - amount to mint during mint period
* reduction period length - defines how many mint periods are in subperiod (subperiod = mint period * reduction_period_length)
* reduction factor - amount multiplying factor;

#### Examples

1. Four years halving minting that starts with 10 millions of tokens yearly

Minter configration:

* minting start: now
* Amount of minter periods: 1
* Minter period 1:
    * period end: null
    * minter type: periodic reduction minter
    * periodic reduction minter parameters:
        * mint_period: 1 year
        * mint_amount: 10 millions
        * reduction_period_length: 4
        * reduction_factor: 0.5

Result:
* 1st 4 years mints 10 millions yearly 
* 2nd 4 years mints 5 millions yearly 
* 3rd 4 years mints 2.5 millions yearly
and so on

2. Linaer minting of 100 millions of token during period of 10 years, next no emission

Minter configration:

* minting start: now
* Amount of minter periods: 2
* Minter period 1:
    * period end: 10 years from now
    * minter type: time linaer minter
    * periodic reduction minter parameters:
        * amount: 100 millions
* Minter period 2:
    * period end: null
    * minter type: no minting

Result:
* 10 millions yearly for 10 years

## Parameters

The Chain4Energy distributor module contains the following configurations parameters:

| Key                  | Type                        | Description                     |
| -------------------- | --------------------------- | ------------------------------- |
| sub_distributors     | List of Subdistributor type | list of defined subdistributors |

### Subdistributor type

| Param                | Type                        | Description                     |
| -------------------- | --------------------------- | ------------------------------- |
| name     | string | unique name of the subdistributor |
| sources  | List of Account type | list of source accounts |
| destination  | Destination type | destinations definition |

### Destination type

| Param       | Type               | Description                                                             |
| ----------- | ------------------ | ----------------------------------------------------------------------- |
| account     | Account type       | mian destination - all remaining tokens from ohers shares anr sent here |
| burn_share  | BurnShare type     | share to burn                                                           |
| share       | List of Share type | List of account destinations with share percentage                      |

### Share type

| Param   | Type         | Description              |
| ------- | ------------ | ------------------------ |
| name    | string       | unique name of the share |
| account | Account type | destination account      |
| percent | Dec          | share percentage 0-100   |

### BurnShare type

| Param   | Type | Description                    |
| ------- | ---- | ------------------------------ |
| percent | Dec  | share percentage to burn 0-100 |

### Account type

| Param   | Type | Description                    |
| ------- | ---- | ------------------------------ |
| type    | enum string  | account type:<br />- MAIN - main module account<br />- MODULE_ACCOUNT - module account<br />- BASE_ACCOUNT - base account<br />- INTERNAL_ACCOUNT - cfedistributor internal account |
| id      | string       | account identifier dependant on the type::<br />- MAIN - empty<br />- MODULE_ACCOUNT - module account name<br />- BASE_ACCOUNT - base account address<br />- INTERNAL_ACCOUNT - cfedistributor internal account name |

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
            "destination": {
                "account": {
                    "id": "validators_rewards",
                    "type": "MODULE_ACCOUNT"
                },
                "share": [
                    {
                        "account": {
                            "id": "c4edwijhdhwqu43efvc3543ec34c2erc342dw",
                            "type": "BASE_ACCOUNT"
                        },
                        "name": "inflation_development_fund_share",
                        "percent": "5.0"
                    },
                    {
                        "account": {
                            "id": "incentive_boosters",
                            "type": "INTERNAL_ACCOUNT"
                        },
                        "name": "inflation_incentive_boosters_share",
                        "percent": "35.0"
                    }
                ],
                "burn_share": {
                    "percent": "0.0"
                }
            },

        },
        {
            "name": "transaction fees subdistributor", 
            "sources": [
                {
                "id": "fee_collector",
                "type": "MODULE_ACCOUNT"
                }
            ],
            "destination": {
                "account": {
                    "id": "validators_rewards",
                    "type": "MODULE_ACCOUNT"
                },
                "share": [
                    {
                        "account": {
                            "id": "incentive_boosters",
                            "type": "INTERNAL_ACCOUNT"
                        },
                        "name": "txs_incentive_boosters_share",
                        "percent": "15.0"
                    }
                ],
                "burn_share": {
                    "percent": "5.0"
                }
            },

        },
        {
            "name": "module_fees_subdistributor", 
            "sources": [
                {
                "id": "fictional_module_fee_collector",
                "type": "MODULE_ACCOUNT"
                }
            ],
            "destination": {
                "account": {
                    "id": "validators_rewards",
                    "type": "MODULE_ACCOUNT"
                },
                "share": [
                    {
                        "account": {
                            "id": "c4edwijhdhwqu43efvc3543ec34c2erc342dw",
                            "type": "BASE_ACCOUNT"
                        },
                        "name": "module_development_fund_share",
                        "percent": "30.0"
                    },
                    {
                        "account": {
                            "id": "incentive_boosters",
                            "type": "INTERNAL_ACCOUNT"
                        },
                        "name": "module_incentive_boosters_share",
                        "percent": "20.0"
                    }
                ],
                "burn_share": {
                    "percent": "0.0"
                }
            },

        },
        {
            "name": "incentive boosters subdistributor", 
            "sources": [
                {
                "id": "incentive_boosters",
                "type": "INTERNAL_ACCOUNT"
                }
            ],
            "destination": {
                "account": {
                    "id": "governance_booster",
                    "type": "MODULE_ACCOUNT"
                },
                "share": [
                    {
                        "account": {
                            "id": "weekend_boosters",
                            "type": "INTERNAL_ACCOUNT"
                        },
                        "name": "weekend_boosters_share",
                        "percent": "35.0"
                    }
                ],
                "burn_share": {
                    "percent": "0.0"
                }
            },

        },
        
    ]
}

```

## State

Chain4Energy distributor module state contains decimal amounts left from previouse block that were impossible to send due value less than 1 token.
Module state contains followng data:

| Key                  | Type                        | Description                     |
| -------------------- | --------------------------- | ------------------------------- |
| states     | List of State type | list of states for burn and per each destination account in all subdistributors (except main distributor) |

### State type

Account account = 1         [(gogoproto.nullable) = true];
  bool burn = 2;
  repeated cosmos.base.v1beta1.DecCoin coins_states = 3 [

| Param                | Type                        | Description                     |
| -------------------- | --------------------------- | ------------------------------- |
| account | Account type (see **[Account type](#account-type)**) | destination account or empty in case of burn flag set to true     |
| burn  | bool | specidies if this is burn destination state |
| coins_states  | DecCoin | list of coins to distribute left by previous block |

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
        "coins_states": [
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
        "coins_states": [
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
        "coins_states": []
    },
    {
        "account": {
            "id": "governance_booster",
            "type": "MODULE_ACCOUNT"
        },
        "burn": false,
        "coins_states": [
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
      "coins_states": [
            {
                "denom": "uc4e",
                "amount": "0.359000000000000000"
            }
      ]
    },
    {
        "burn": true,
        "coins_states": [
            {
                "denom": "uc4e",
                "amount": "0.800000000000000000"
            }
        ]
    }
]

```

## Events

Chain4Energy distributor module module emits the following events:

### BeginBlockers

#### Tokens distribution

| Type         | Attribute Key | Attribute Value |
| ------------ | ------------- | --------------- |
| DistributionsResult | DistributionResult | list of DistributionResult type |

##### DistributionResult type

DistributionResult type represents one send operation to one destination in one block

| Param   | Type | Description                    |
| ------- | ---- | ------------------------------ |
| source  | list of Account type (see **[Account type](#account-type)**) | list of sources |
| destination | Account type (see **[Account type](#account-type)**) | destination |
| coinSend | DecCoins | coins sent to destination |

## Queries

### Params query

Queries the module params.

See example reponse:

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
        "destination": {
          "account": {
            "id": "c4e_distributor",
            "type": "MAIN"
          },
          "share": [],
          "burn_share": {
            "percent": "0.000000000000000000"
          }
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
        "destination": {
          "account": {
            "id": "validators_rewards_collector",
            "type": "MODULE_ACCOUNT"
          },
          "share": [
            {
              "name": "development_fund",
              "percent": "5.000000000000000000",
              "account": {
                "id": "c4e10ep2sxpf2kj6jsdcs234edkuf9sf9xqq3sl",
                "type": "BASE_ACCOUNT"
              }
            },
            {
              "name": "usage_incentives",
              "percent": "35.000000000000000000",
              "account": {
                "id": "usage_incentives_collector",
                "type": "INTERNAL_ACCOUNT"
              }
            }
          ],
          "burn_share": {
            "percent": "0.000000000000000000"
          }
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
        "destination": {
          "account": {
            "id": "c4e1q5vgy0r3scsdc32dcewkl8nwmfe2mgr6g0jlph",
            "type": "BASE_ACCOUNT"
          },
          "share": [
            {
              "name": "green_energy_booster",
              "percent": "34.000000000000000000",
              "account": {
                "id": "green_energy_booster_collector",
                "type": "MODULE_ACCOUNT"
              }
            },
            {
              "name": "governance_booster",
              "percent": "33.000000000000000000",
              "account": {
                "id": "governance_booster_collector",
                "type": "MODULE_ACCOUNT"
              }
            }
          ],
          "burn_share": {
            "percent": "0.000000000000000000"
          }
        }
      }
    ]
  }
}
```
### Params states

Queries the module state.

See example reponse:

```json
{
  "states": [
    {
      "account": {
        "id": "c4e10ep2ssdfwefcscaewdedscs9xqqqdwqee3sl",
        "type": "BASE_ACCOUNT"
      },
      "burn": false,
      "coins_states": [
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
      "coins_states": [
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
      "coins_states": [
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
      "coins_states": [
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
      "coins_states": []
    },
    {
      "account": {
        "id": "validators_rewards_collector",
        "type": "MODULE_ACCOUNT"
      },
      "burn": false,
      "coins_states": [
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

### Non Negative Coin State Invariant

Invariant validates module state. Checks if all coins states of all destinations are non negative

### State Sum Balance Check Invariant

Invariant validates module state. Checks sum of all coins states of one denom of all destinations is always intiger value and is equal to cfedistributor module account balance

## Genesis validations

TODO
