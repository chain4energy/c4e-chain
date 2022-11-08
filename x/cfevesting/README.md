# Chain4Energy vesting module - cfevesting

## Abstract

Chain4Energy vesting module provides functionality creation and manegement of vesting pools. 
Vesting pool locks  configured amount of tokens for configured period of time. Those locked tokens can further be sent form vesting pool, but only to newly created continuous vesting account. Vesting account parameters are calculted from vesting pool configuration.

## Contents

1. **[Concept](#concepts)**
2. **[Params](#parameters)**
3. **[State](#state)**
3. **[Messages](#messages)**
4. **[Events](#events)**
5. **[Queries](#queries)**
6. **[Invariants](#invariants)**
7. **[Genesis validations](#genesis-validations)**

## Concepts

The purpose of `cfevesting` module is to provide functionality of locking tokens in a pool for specified amount of time, but with the ability to send those tokens to continuous vesting accounts. So, tokens are still locked but a cannot be sent further from vesting account. This allows to create vesting target groups (e.g. validators vesting pool, investors vesting pools).
`Cfevesting` module keeps list of all vesting accounts cretated form vesting pools. It is use for calculation of current toens amount in vesting.

### Vesting Pool

Vesting pool locks some amount of tokens to configured period of time. Each vesting pool has its owner (account). One owner can have multiple vasting pools. Each vesting pool is identified by its name. The name is unique among all vesitng pools of one owner. Vesting pools have fillowing parameters:
* name - unique name per owner
* vesting type - vesting type used by vesting pool (see **[Vesting Type](#vesting-type)**)
* lock start - time of pool creation
* lock end - time of tokens unlocking
* initially_locked - amount locked initially in the pool
* withdrawn - amount withdrawn from the pool (currently tokens can be withdrawn only after lock end time)
* sent - amount sent to vesting accounts from the vesting pool

### Vesting Type

Vesting type defines how continuous vesting account time values are calculated during creation: 
* continuous vesting account start time - last block time + vesting type lockup period
* continuous vesting account end time - last block time + vesting type lockup period + vesting type vesting period
where:
* vesting type lockup period - is period of time when all tokens are locked
* vesting type vesting period - is period of time when tokens are are lieary vested

The list of vesting types is predifined on genesis.

## Parameters

The Chain4Energy vesting module contains the following configurations parameters:

| Key                  | Type                        | Description                     |
| -------------------- | --------------------------- | ------------------------------- |
| denom     | string | Denom of vesting token |

## State

### Vesting pools state

Chain4Energy vesting module state of account vesting pools stores vesting pools lists per owners.
Vesting pools state contains followng data:

#### AccountVestingPool type

| Key                  | Type                        | Description                     |
| -------------------- | --------------------------- | ------------------------------- |
| address     | string | Owner address |
| vesting_pools     | List of VestingPool type | Vesting pools of the owner |

#### VestingPool type

| Key                  | Type                        | Description                     |
| -------------------- | --------------------------- | ------------------------------- |
| name     | string | unique name per owner |
| vesting_type     | string | vesting type used by vesting pool (see **[Vesting Type](#vesting-type)**) |
| lock_start     | time.Duration | time of pool creation |
| lock_end     | time.Duration  | time of tokens unlockin |
| initially_locked     | sdk.Int | amount locked initially in the pool |
| withdrawn     | sdk.Int | amount withdrawn from the pool (currently tokens can be withdrawn only after lock end time) |
| sent     | sdk.Int | amount sent to vesting accounts from the vesting pool |

### Vesting types data dictionary

Vesting types data dictionary contains list of predefined vesting types:

| Key                  | Type                        | Description                     |
| -------------------- | --------------------------- | ------------------------------- |
| name     | string | unique vestoing type name |
| lockup_period     | time.Duration | is period of time when all tokens are locked |
| vesting_period     | time.Duration | is period of time when tokens are are lieary vested |

## Messages

### Create Vesting Pool

Creates new vesting pool for the creator account.

`MsgCreateVestingPool` can be submitted by any token holder via a
`MsgCreateVestingPool` transaction.

``` {.go}
type MsgCreateVestingPool struct {
	Creator     string
	Name        string
	Amount      sdk.Int
	Duration    time.Duration
	VestingType string
}
```

**Params:**

| Param                  | Description                     |
| -------------------- | ------------------------------- |
| Creator     | Creator address |
| Name     | Vesting pool name |
| Amount     | Amount to lock in vesting pool |
| Duration     | Lock duration |
| VestingType     | Vesting Type of niew vesting pool |

**State modifications:**

- Validate `Creator` has enough tokens
- Generate new `VestingPool` record for creator
- Save the record inside the keeper's Account Vesting Pools
- Transfer the tokens from the `Creator` account to cfevesting `ModuleAccount`.

### Send To Vesting Account

Creates new continuous vesting account and sends token from vestoing pool.

`MsgSendToVestingAccount` can be submitted by any Vesting pool owner via a
`MsgSendToVestingAccount` transaction.

``` {.go}
type MsgSendToVestingAccount struct {
	FromAddress     string
	ToAddress       string
	VestingPoolName string
	Amount          sdk.Int
	RestartVesting  bool
}
```

**Params:**

| Param                  | Description                     |
| -------------------- | ------------------------------- |
| FromAddress     | Vesting pool owenr address |
| ToAddress     | New continuous vesting account address |
| VestingPoolName     | Vesting pool name |
| Amount     | Amount to lock in vesting pool |
| RestartVesting     | Defines how time parameters of new vesting account should be calculatad:<br>- true:<br>&nbsp;&nbsp;&nbsp;continuous vesting account start time - last block time + vesting type lockup period<br>&nbsp;&nbsp;&nbsp;continuous vesting account end time - last block time + vesting type lockup period + vesting type vesting period<br>- false:<br>&nbsp;&nbsp;&nbsp;continuous vesting account start time - vesting pool lock end<br>&nbsp;&nbsp;&nbsp;continuous vesting account end time - vesting pool lock end |

**State modifications:**

- Validate `FromAddress` vesting pool with name has enough tokens
- Creates new continuous vesting account with address equal to ToAddress and time params according vesting type of VestingPool with name equal to VestingPoolName
- Sends tokens from cfevesting `ModuleAccount` to ToAddress
- Updates Vesting pool state

### Withdraw All Available

Withdraws all available (unlocked) tokens from vesting pool back to owner account

`MsgWithdrawAllAvailable` can be submitted by any Vesting pool owner via a
`MsgWithdrawAllAvailable` transaction.

``` {.go}
type MsgWithdrawAllAvailable struct {
	Creator string
}
```

| Param                  | Description                     |
| -------------------- | ------------------------------- |
| Creator     | Vestign pool owner address |   // TODO change to owner

**State modifications:**

- Sends unlocked tokens from cfevesting `ModuleAccount` to `Creator` account
- Updates Vesting pool state

### Create Vesting Account

Creates new continuous vesting account and sends token from creator account.

`MsgCreateVestingPool` can be submitted by any token holder via a
`MsgCreateVestingPool` transaction.

``` {.go}
type MsgCreateVestingAccount struct {
	FromAddress string
	ToAddress   string
	Amount      sdk.Coins`
	StartTime   int64
	EndTime     int64
}
```

**Params:**

| Param                  | Description                     |
| -------------------- | ------------------------------- |
| FromAddress     | Vesting pool owenr address |
| ToAddress     | New continuous vesting account address |
| Amount     | Amount to lock in vesting account |
| StartTime     | Vesting start time - unix |
| EndTime     | Vesting end time - unix |

**State modifications:**

- Validate `FromAddress` enough tokens
- Creates new continuous vesting account with address equal to ToAddress and time params according to provided data
- Sends tokens from `FromAddress` account to ToAddress

## Events

Chain4Energy distributor module emits the following events:

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
### States query

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
