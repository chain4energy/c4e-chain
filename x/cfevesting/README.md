# Chain4Energy vesting module - cfevesting

## Abstract

Chain4Energy vesting module allows to create and manage vesting pools.
One can create a vesting pool in order to lock a certain amount of tokens for a period of time. During this locking period the locked tokens can only be used for continuous vesting accounts creation. 
The vesting account parameters are calculated from the vesting pool configuration.

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

The purpose of `cfevesting` module is to provide the functionality of locking tokens in a pool for specified amount of time with the ability of sending those tokens to continuous vesting accounts. 
The tokens will be still locked on the vesting accounts, and it will not be possible to sent them further until the locking period ends. This allows to create vesting target groups (e.g. validators vesting pool, investors vesting pools).
The `cfevesting` module keeps track of all the vesting pools and vesting accounts and allows to calculate the total amount of tokens in vesting.

### Vesting Pool

Vesting pool locks a certain amount of tokens for configured period of time. Each vesting pool has its owner (creator account).
A single owner can have multiple vesting pools. Each vesting pool is identified by its unique name (unique among all the pools belonging to a single owner). 

Vesting pools have the following parameters:
* name - unique name (among owner's pools)
* vesting type - vesting type used by vesting pool (see **[Vesting Type](#vesting-type)**)
* lock start - time of pool creation
* lock end - unlocking time (end of lock period)
* initially_locked - amount of tokens locked initially in the pool
* withdrawn - amount of tokens that were already withdrawn from the pool (currently all available (available = initially_locked - sent) tokens can be withdrawn by the owner only after lock end time)
* sent - amount of tokens that were already sent to vesting accounts from the vesting pool
* reservations - amount of tokens that were reserved from the vesting pool (reserved tokens are not available for sending to vesting accounts)

### VestingPoolReservation

Vesting pool reservation defines amount of tokens reserved in a vesting pool.
Vesting pool reservation has the following parameters:
* id - unique id
* amount - amount of reserved tokens

### Vesting Type

Vesting type defines the parameters of continuous vesting accounts that will be created from a given vesting pool.
Vesting type has the following parameters:
* name - unique name 
* lockup period - period of time when all the tokens in the pool are locked
* vesting period - period of time when tokens are linearly vested
* free - the percentage of tokens that will be released at the beginning after sending to a continuous vesting account

New vesting account parameters will be set accordingly:
* continuous vesting account start time = last block time + vesting type lockup period
* continuous vesting account end time = last block time + vesting type lockup period + vesting type vesting period

The vesting types are predefined on genesis.

### Periodic Continuous Vesting Account

A Periodic Continuous Vesting Account is a type of vesting account that unlocks tokens in a continuous manner over specified periods. It introduces the concept of vesting periods within continuous vesting, allowing for periodic unlocking of tokens. This document provides a detailed explanation of the Periodic Continuous Vesting Account and its functionalities, emphasizing the differences from the standard Continuous Vesting Account.

#### Account Structure

The Periodic Continuous Vesting Account extends the `BaseVestingAccount` structure and includes additional fields:

- `StartTime`: The time when vesting starts for the account.
- `VestingPeriods`: A list of `ContinuousVestingPeriod` instances representing individual vesting periods.

#### Additional Functions

#### `AddNewContinousVestingPeriod`

This function allows adding a new vesting period to the account. It takes the `startTime`, `endTime`, and `amount` as parameters and updates the account accordingly.
This function is used to add new vesting periods to the account, contributing to the overall vesting structure.

#### `GetVestingCoinsForSpecificPeriods`

This function calculates the total number of vesting coins for specific periods, based on a given
`blockTime` and an array of period IDs (`periodsToTrace`). It sums up the vesting coins of the 
specified periods to determine the total vesting amount during those periods.

#### Continuous Vesting Period

The `ContinuousVestingPeriod` structure represents an individual vesting period within the
Periodic Continuous Vesting Account. It includes:

- `StartTime`: The starting time of the period.
- `EndTime`: The ending time of the period.
- `Amount`: The amount of coins vesting during this period.

#### `GetVestedCoins`

This function calculates the vested coins for a specific vesting period based
on a given `blockTime`. It considers the vesting scalar and computes the vested 
amount according to the proportional period duration.

#### Validation

The Periodic Continuous Vesting Account and its components include validation checks
to ensure data consistency and accuracy:

- The start time of a vesting period cannot be after its end time.
- The total original vesting amount must match the sum of amounts in all vesting periods.

#### Usage

The Periodic Continuous Vesting Account is designed to manage vesting of tokens in a 
continuous manner with the added feature of multiple vesting periods. Users can add new
vesting periods to the account and calculate vested and vesting coins for specific periods.
This account type provides flexibility in managing token vesting for scenarios requiring
periodic unlocks within a continuous vesting structure.

## Parameters

The Chain4Energy vesting module contains the following configurations parameters:

| Key         | Type     | Description              |
|-------------|----------|--------------------------|
| denom       | string   | Denom of vesting token   |

## State

### Vesting pools state

Chain4Energy vesting module state contains vesting pools lists per owner.

#### AccountVestingPools type

| Key           | Type                      | Description                |
|---------------|---------------------------|----------------------------|
| owner         | string                    | Owner address              |
| vesting_pools | List of VestingPool types | Vesting pools of the owner |

#### VestingPool type

| Key              | Type                     | Description                                                                       |
|------------------|--------------------------|-----------------------------------------------------------------------------------|
| name             | string                   | unique name per owner                                                             |
| vesting_type     | string                   | vesting type used by vesting pool (see **[Vesting Type](#vesting-type)**)         |
| lock_start       | time.Duration            | time of pool creation                                                             |
| lock_end         | time.Duration            | unlocking time (end of lock period)                                               |
| initially_locked | math.Int                 | amount of tokens locked initially in the pool                                     |
| withdrawn        | math.Int                 | amount of tokens that were already withdrawn from the pool                        |
| sent             | math.Int                 | amount of tokens that were already sent to vesting accounts from the vesting pool |
| reservations     | []VestingPoolReservation | array of vesting pool reservations                                                |

#### VestingPoolReservation type

In the context of a vesting pool, the concept of reservations plays a vital 
role in managing the allocation and distribution of tokens. Reservations refer 
to the process of setting aside a specific amount of tokens within a vesting pool 
for designated purposes. Each reservation is defined by its unique identifier (ID) 
and the corresponding amount of tokens it reserves.

| Key            | Type                     | Description                 |
|----------------|--------------------------|-----------------------------|
| id             | string                   | identifier of a reservation |
| amount         | math.Int                 | amount of tokens reserved   |

### Vesting types data dictionary

Vesting types data dictionary contains list of predefined vesting types:

| Key            | Type            | Description                                        |
|----------------|-----------------|----------------------------------------------------|
| name           | string          | unique vesting type name                           |
| lockup_period  | time.Duration   | period of time when all tokens are locked          | 
| vesting_period | time.Duration   | period of time when tokens are are linearly vested |

### Vesting account list

Vesting account list contains address of all vesting accounts created with cfevesting module.

#### VestingAccount type
| Key      | Type     | Description                                  |
|----------|----------|----------------------------------------------|
| id       | uint64   | id of an entity in the vesting accounts list |
| address  | string   | vesting account address                      |

## Messages

### Create Vesting Pool

Creates new vesting pool for the creator account.

`MsgCreateVestingPool` can be submitted by any token holder via a
`MsgCreateVestingPool` transaction.

``` {.go}
type MsgCreateVestingPool struct {
	Creator     string
	Name        string
	Amount      math.Int
	Duration    time.Duration
	VestingType string
}
```

**Params:**

| Param       | Description                        |
|-------------|------------------------------------|
| Creator     | Creator/Owner address              |
| Name        | Vesting pool name                  |
| Amount      | Amount to lock in the vesting pool |
| Duration    | Lock duration                      |
| VestingType | Vesting Type of the pool           |

**State modifications:**

- Validate `Creator` has enough tokens
- Generate new `VestingPool` record for creator/owner
- Save the record in the owner account vesting pools list
- Transfer the tokens from the `Creator` account to cfevesting `ModuleAccount`.

### Send To Vesting Account

Creates a new continuous vesting account and sends tokens from vesting pool to it.

`MsgSendToVestingAccount` can be submitted by any Vesting pool owner via a
`MsgSendToVestingAccount` transaction.

``` {.go}
type MsgSendToVestingAccount struct {
	FromAddress     string
	ToAddress       string
	VestingPoolName string
	Amount          math.Int
	RestartVesting  bool
}
```

**Params:**

| Param           | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
|-----------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| FromAddress     | Vesting pool owner address                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| ToAddress       | New continuous vesting account address                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| VestingPoolName | Vesting pool name                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| Amount          | Amount to lock in the vesting pool                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| RestartVesting  | Defines how time parameters of new vesting account should be calculatad:<br>- true:<br>&nbsp;&nbsp;&nbsp;continuous vesting account start time = last block time + vesting type lockup period<br>&nbsp;&nbsp;&nbsp;continuous vesting account end time = last block time + vesting type lockup period + vesting type vesting period<br>- false:<br>&nbsp;&nbsp;&nbsp;continuous vesting account start time = vesting pool lock end<br>&nbsp;&nbsp;&nbsp;continuous vesting account end time = vesting pool lock end |

**State modifications:**

- Validates if `FromAddress` owner's vesting pool `VestingPoolName` has enough tokens
- If account with `ToAddress` does not exist creates new periodic continuous vesting account with `ToAddress`
- Add a new continuous vesting period to `ToAddress` account with time params calculated according to the pool vesting type
- Sends tokens from cfevesting `ModuleAccount` to `ToAddress`
- Updates the vesting pool state

### Withdraw All Available

Withdraws all available (unlocked) tokens from the vesting pool back to the owner account

`MsgWithdrawAllAvailable` can be submitted by any vesting pool owner via a
`MsgWithdrawAllAvailable` transaction.

``` {.go}
type MsgWithdrawAllAvailable struct {
	Owner string
}
```

| Param | Description                |
|-------|----------------------------|
| Owner | Vesting pool owner address |

**State modifications:**

- Sends unlocked tokens from cfevesting `ModuleAccount` to `Creator` account
- Updates the vesting pool state

### Create Vesting Account

Creates new continuous vesting account and sends token from creator account.

`MsgCreateVestingAccount` can be submitted by any token holder via a
`MsgCreateVestingAccount` transaction.

``` {.go}
type MsgCreateVestingAccount struct {
	FromAddress string
	ToAddress   string
	Amount      sdk.Coins
	StartTime   int64
	EndTime     int64
}
```

**Params:**

| Param            | Description                            |
|------------------|----------------------------------------|
| FromAddress      | Vesting pool owner address             |
| ToAddress        | New continuous vesting account address |
| Amount           | Amount to lock in vesting account      |
| StartTime        | Vesting start time - unix              |
| EndTime          | Vesting end time - unix                |

**State modifications:**

- Validates if `FromAddress` has enough tokens. 
- Creates new continuous vesting account with address equal to ToAddress and time params according to provided data
- Sends tokens from `FromAddress` account to `ToAddress`

### Split vesting

Split tokens that are locked in vesting to new vesting account. Total number of tokens in vesting, vesting times and token release speed are preserved.
This mechanism can also be called as a "vesting cession".

`MsgSplitVesting` can be submitted by any vesting account via a `MsgSplitVesting` transaction.

``` {.go}
type MsgSplitVesting struct {
	FromAddress string
	ToAddress   string
	Amount      sdk.Coins
}
```

**Params:**

| Param            | Description                            |
|------------------|----------------------------------------|
| FromAddress      | Vesting pool owner address             |
| ToAddress        | New continuous vesting account address |
| Amount           | Amount of locked vesting to split      |

**State modifications:**

- Validates if `FromAddress` has enough locked tokens in the vesting
- Creates new continuous vesting account with address equal to ToAddress
  and time parameters set to: 
  - `end time` is set to the vesting account end time
  - in the case when the start time of the vesting account is in the future - `new account start time = from account start time`  
  - in the case when the start time of the vesting account is in the past `new account start time = transaction time`
- Sends locked vesting from `FromAddress` account to `ToAddress`

### Move available vesting

Moves all tokens that are locked in vesting to new vesting account. Total number of tokens in vesting, vesting times and token release speed are preserved.
This mechanism can also be called as a "vesting cession".

`MsgMoveAvailableVesting` can be submitted by any vesting account via a `MsgMoveAvailableVesting` transaction.

``` {.go}
type MsgMoveAvailableVesting struct {
	FromAddress string
	ToAddress   string
}
```

**Params:**

| Param            | Description                            |
|------------------|----------------------------------------|
| FromAddress      | Vesting pool owner address             |
| ToAddress        | New continuous vesting account address |

**State modifications:**

- Validates if `FromAddress` has any locked tokens in the vesting
- Creates new continuous vesting account with address equal to ToAddress
  and time parameters set to:
    - `end time` is set to the vesting account end time
    - in the case when the start time of the vesting account is in the future - `new account start time = from account start time`
    - in the case when the start time of the vesting account is in the past `new account start time = transaction time`
- Sends locked vesting from `FromAddress` account to `ToAddress`

### Move available vesting by denoms

Moves all tokens that are locked in vesting to new vesting account. This message differs from `MsgMoveAvailableVesting` in
that you can additionally provide a list of denominations that are to be taken into account when sending a blocked vesting. 
Total number of tokens in vesting, vesting times and token release speed are preserved.
This mechanism can also be called as a "vesting cession". 

`MsgMoveAvailableVestingByDenoms` can be submitted by any vesting account via a `MsgMoveAvailableVestingByDenoms` transaction.

``` {.go}
type MsgMoveAvailableVestingByDenoms struct {
	FromAddress string
	ToAddress   string
	Denoms      []string
}
```

**Params:**

| Param       | Description                                                           |
|-------------|-----------------------------------------------------------------------|
| FromAddress | Vesting pool owner address                                            |
| ToAddress   | New continuous vesting account address                                |
| Denoms      | List of denominations to be taken into account when unlocking vesting |

**State modifications:**

- Validates if `FromAddress` has any locked tokens (only those highlighted in `denoms`) in the vesting
- Creates new continuous vesting account with address equal to ToAddress and time parameters set to:
    - `end time` is set to the vesting account end time
    - in the case when the start time of the vesting account is in the future - `new account start time = from account start time`
    - in the case when the start time of the vesting account is in the past `new account start time = transaction time`
- Sends locked vesting from `FromAddress` account to `ToAddress`

## Events

Chain4Energy distributor module emits the following events:

### Handlers

#### MsgCreateVestingPool

| Type                | Attribute Key | Attribute Value                                        |
|---------------------|---------------|--------------------------------------------------------|
| EventNewVestingPool | owner         | {owner_address}                                        |
| EventNewVestingPool | name          | {vesting_pool_name}                                    |
| EventNewVestingPool | amount        | {vesting_pool_amount}                                  |
| EventNewVestingPool | duration      | {lock_duration}                                        |
| EventNewVestingPool | vesting_type  | {vesting\_type\_name}                                  |
| message             | action        | /chain4energy.c4echain.cfevesting.MsgCreateVestingPool |
| message             | sender        | {sender_address}                                       |
| transfer            | recipient     | {moduleAccount}                                        |
| transfer            | sender        | {owner_address}                                        |
| transfer            | amount        | {amount}                                               |

#### MsgSendToVestingAccount

| Type                                 | Attribute Key       | Attribute Value                                                                |
|--------------------------------------|---------------------|--------------------------------------------------------------------------------|
| EventNewVestingAccount               | address             | {new\_vesting\_account\_address}                                               |
| EventNewVestingPeriodFromVestingPool | owner               | {owner_address}                                                                |
| EventNewVestingPeriodFromVestingPool | address             | {vesting\_account\_address}                                                    |
| EventNewVestingPeriodFromVestingPool | vesting\_pool\_name | {source\_vesting_pool_name}                                                    |
| EventNewVestingPeriodFromVestingPool | amount              | {amount_to_send_to_new\_vesting\_account}                                      |
| EventNewVestingPeriodFromVestingPool | restart_vesting     | {restart_vesting} see  **[Send To Vesting Account](#send-to-vesting-account)** |
| EventNewVestingPeriodFromVestingPool | period_id           | {new_period_id}                                                                |
| message                              | action              | /chain4energy.c4echain.cfevesting.MsgSendToVestingAccount                      |
| message                              | sender              | {sender_address}                                                               |
| transfer                             | recipient           | {module_account}                                                               |
| transfer                             | sender              | {creator}                                                                      |
| transfer                             | amount              | {amount}                                                                       |

#### MsgWithdrawAllAvailable

| Type                     | Attribute Key       | Attribute Value                                           |
|--------------------------|---------------------|-----------------------------------------------------------|
| EventWithdrawAvailable   | owner               | {owner_address}                                           |
| EventWithdrawAvailable   | vesting\_pool\_name | {source\_vesting_pool_name}                               |
| EventWithdrawAvailable   | amount              | {withdrawn_amount}                                        |
| message                  | action              | /chain4energy.c4echain.cfevesting.MsgWithdrawAllAvailable |
| message                  | sender              | {sender_address}                                          |
| transfer                 | recipient           | {module_account}                                          |
| transfer                 | sender              | {creator}                                                 |
| transfer                 | amount              | {amount}                                                  |

#### MsgCreateVestingAccount

| Type                   | Attribute Key       | Attribute Value                                           |
|------------------------|---------------------|-----------------------------------------------------------|
| EventNewVestingAccount | address             | {new\_vesting\_account\_address}                          |
| message                | action              | /chain4energy.c4echain.cfevesting.MsgCreateVestingAccount |
| message                | sender              | {sender_address}                                          |
| transfer               | recipient           | {module_account}                                          |
| transfer               | sender              | {creator}                                                 |
| transfer               | amount              | {amount}                                                  |

#### MsgSplitVesting

| Type                   | Attribute Key | Attribute Value                                   |
|------------------------|---------------|---------------------------------------------------|
| EventVestingSplit      | source        | {from_account\_address}                           |
| EventVestingSplit      | destination   | {to\_account\_address}                            |
| EventNewVestingAccount | address       | {new\_vesting\_account\_address}                  |
| message                | action        | /chain4energy.c4echain.cfevesting.MsgSplitVesting |
| message                | sender        | {sender_address}                                  |
| transfer               | recipient     | {module_account}                                  |
| transfer               | sender        | {creator}                                         |
| transfer               | amount        | {amount}                                          |

#### MsgMoveAvailableVesting

| Type                   | Attribute Key | Attribute Value                                           |
|------------------------|---------------|-----------------------------------------------------------|
| EventVestingSplit      | source        | {from_account\_address}                                   |
| EventVestingSplit      | destination   | {to\_account\_address}                                    |
| EventNewVestingAccount | address       | {new\_vesting\_account\_address}                          |
| message                | action        | /chain4energy.c4echain.cfevesting.MsgMoveAvailableVesting |
| message                | sender        | {sender_address}                                          |
| transfer               | recipient     | {module_account}                                          |
| transfer               | sender        | {creator}                                                 |
| transfer               | amount        | {amount}                                                  |

#### MsgMoveAvailableVestingByDenoms

| Type                   | Attribute Key | Attribute Value                                                   |
|------------------------|---------------|-------------------------------------------------------------------|
| EventVestingSplit      | source        | {from_account\_address}                                           |
| EventVestingSplit      | destination   | {to\_account\_address}                                            |
| EventNewVestingAccount | address       | {new\_vesting\_account\_address}                                  |
| message                | action        | /chain4energy.c4echain.cfevesting.MsgMoveAvailableVestingByDenoms |
| message                | sender        | {sender_address}                                                  |
| transfer               | recipient     | {module_account}                                                  |
| transfer               | sender        | {creator}                                                         |
| transfer               | amount        | {amount}                                                          |

## Queries

### Params query

Queries the module params.

See example response:

```json
{
  "params": {
    "denom": "uc4e"
  }
}
```
### Summary query

Queries the vesting summary data.

See example response:

```json
{
  "vesting_all_amount": "32500000000000",
  "vesting_in_pools_amount": "32500000000000",
  "vesting_in_accounts_amount": "0",
  "delegated_vesting_amount": "0"
}
```

### Genesis Summary query

Queries the vesting summary data but only for vesting pools with GenesisPool set to true
and accounts that are on a vesting account trace list.

See example response:

```json
{
  "vesting_all_amount": "32500000000000",
  "vesting_in_pools_amount": "32500000000000",
  "vesting_in_accounts_amount": "0",
  "delegated_vesting_amount": "0"
}
```

### Vesting pool query

Queries the vesting pools owned by account with given address.

See example response:

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
      "sent_amount": "0",
      "reservations": [
        {
          "id": "1",
          "amount": "15000000000000"
        }
      ]
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
      "sent_amount": "0",
      "reservations": [
        {
          "id": "1",
          "amount": "15000000000000"
        }
      ]
    }
  ]
}
```

### Vesting types query

Queries the vesting types.

See example response:

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

### Non-Negative Vesting Pool Amounts Invariant

Invariant validates vesting pools state. Checks if all vesting pools amounts are non-negative

### Vesting Pool Consistent Data Invariant

Invariant validates vesting pools state. Checks if all vesting pools amounts are consistent: withdrawn + sent < initially locked

### Module Account Invariant

Invariant validates vesting pools state. Checks if sum of all amounts locked in vesting pools is equal to module account balance.

## Genesis validations

[//]: # (TODO)
