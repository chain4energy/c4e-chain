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

`MsgCreateVestingAccount` can be submitted by any token holder via a
`MsgCreateVestingAccount` transaction.

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

### Handlers

#### MsgCreateVestingPool

|  Type          | Attribute Key     | Attribute Value  |
|  --------------| ------------------| -----------------|
|  NewVestingPool  | creator  | {creator_owner_address}   |
|  NewVestingPool  | name             | {vesting_pool_name}          |
|  NewVestingPool  | amount            | {vesting_pool_amount}         |
|  NewVestingPool  | duration          | {lock_duration}       |
|  NewVestingPool  | vesting\_type      | {vesting\_type\_name}     |
|  message       | action            | ??     |
|  message       | sender            | ??       |
|  transfer      | recipient         | {moduleAccount}  |
|  transfer      | sender            | {creator}         |
|  transfer      | amount            | {amount}    |

// TODO verify

#### MsgSendToVestingAccount

|  Type          | Attribute Key     | Attribute Value  |
|  --------------| ------------------| -----------------|
|  WithdrawAvailable  | owner\_address  | {owner_address}   |
|  WithdrawAvailable  | vesting\_pool\_name | {source\_vesting_pool_name}         |
|  WithdrawAvailable  | amount          | {withdrawn_amount}       |
|  NewVestingAccountFromVestingPool  | owner\_address  | {owner_address}   |
|  NewVestingAccountFromVestingPool  | address             | {new\_vesting\_account\_address}          |
|  NewVestingAccountFromVestingPool  | vesting\_pool\_name | {source\_vesting_pool_name}         |
|  NewVestingAccountFromVestingPool  | amount          | {amount_to_send_to_new\_vesting\_account}       |
|  NewVestingAccountFromVestingPool  | restart_vesting      | {restart_vesting} see  **[Send To Vesting Account](#send-to-vesting-account)**   |
|  message       | action            | ??     |
|  message       | sender            | ??       |
|  transfer      | recipient         | {moduleAccount}  |
|  transfer      | sender            | {creator}         |
|  transfer      | amount            | {amount}    |

// TODO verify

#### MsgWithdrawAllAvailable

|  Type          | Attribute Key     | Attribute Value  |
|  --------------| ------------------| -----------------|
|  WithdrawAvailable  | owner\_address  | {owner_address}   |
|  WithdrawAvailable  | vesting\_pool\_name | {source\_vesting_pool_name}         |
|  WithdrawAvailable  | amount          | {withdrawn_amount}       |
|  message       | action            | ??     |
|  message       | sender            | ??       |
|  transfer      | recipient         | {moduleAccount}  |
|  transfer      | sender            | {creator}         |
|  transfer      | amount            | {amount}    |

// TODO verify

#### MsgCreateVestingAccount

|  Type          | Attribute Key     | Attribute Value  |
|  --------------| ------------------| -----------------|
|  NewVestingAccount  | address  | {new\_vesting\_account\_address}   |
|  message       | action            | ??     |
|  message       | sender            | ??       |
|  transfer      | recipient         | {moduleAccount}  |
|  transfer      | sender            | {creator}         |
|  transfer      | amount            | {amount}    |

// TODO verify - some more params

## Queries

### Params query

Queries the module params.

See example reponse:

```json
{
  "params": {
    "denom": "uc4e"
  }
}
```
### Summary query

Queries the vesting summary data.

See example reponse:

```json
{
  "vesting_all_amount": "32500000000000",
  "vesting_in_pools_amount": "32500000000000",
  "vesting_in_accounts_amount": "0",
  "delegated_vesting_amount": "0"
}
```

### Vesting pool query

Queries the vesting pools of owner address.

See example reponse:

```json
{
  "vesting_pools": [
    {
      "name": "Advisors pool",
      "vesting_type": "Advisors pool",
      "lock_start": "2022-03-30T00:00:00Z",
      "lock_end": "2025-03-30T00:00:00Z",
      "withdrawable": "0",
      "initially_locked": {
        "denom": "uc4e",
        "amount": "15000000000000"
      },
      "currently_locked": "15000000000000",
      "sent_amount": "0"
    },
    {
      "name": "Validators pool",
      "vesting_type": "Validators pool",
      "lock_start": "2022-03-30T00:00:00Z",
      "lock_end": "2024-03-30T00:00:00Z",
      "withdrawable": "0",
      "initially_locked": {
        "denom": "uc4e",
        "amount": "17500000000000"
      },
      "currently_locked": "17500000000000",
      "sent_amount": "0"
    }
  ]
}
```

### Vesting types query

Queries the vesting types lisy.

See example reponse:

```json
{
  "vesting_types": [
    {
      "name": "Advisors pool",
      "lockup_period": "5",
      "lockup_period_unit": "minute",
      "vesting_period": "5",
      "vesting_period_unit": "day"
    },
    {
      "name": "Validators pool",
      "lockup_period": "10",
      "lockup_period_unit": "minute",
      "vesting_period": "10",
      "vesting_period_unit": "day"
    }
  ]
}
```

## Invariants

### Non Negative Vesting Pool Amounts Invariant

Invariant validates vesting pools state. Checks if all vesting pools amounts are non negative

### Vesting Pool Consistent Data Invariant

Invariant validates vesting pools state. Checks if all vesting pools amounts are consistent: wothdrawn + sent < initially locked

### Module Account Invariant

Invariant validates vesting pools state. Checks if sum of all amounts locked in vesting pools is equal to module account balance.

## Genesis validations

TODO
