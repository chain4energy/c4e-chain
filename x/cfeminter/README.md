# Chain4Energy minter module - cfeminter

## Abstract

Chain4Energy minter module provides functionality of controlled token emissions. Tokens are emited according to configuration. 

## Contents

1. **[Concept](#concepts)**
2. **[Params](#parameters)**
3. **[State](#state)**
4. **[Events](#events)**
5. **[Queries](#queries)**
6. **[Genesis validations](#genesis-validations)**

## Concepts

The purpose of `cfeminter` module is to provide token emission mechanism.

### Token emission mechanism

Token emission mechanism mints calculated amount of tokens per each block. Token amount is calculated according to cfeminter module configuration params. 
Tokens minting process is divided into Minters where each period has fully separated minting rules. Those Minters rules are difined within cfeminter module configuration params.
Simply, mintining process configuration is a list of ordered minting Minters, where each period has its own start and end time (end time for last period is not required, in that case last minting period works infinitely).

### Miniting Minters

Minting period is a period of time when configured minting rules apply. Ordered list of minting Minters deifnes whole token emmision process.
End time of one period is a start time of the next period on the Minters list.
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

Periodic reduction minter is block time based minter type. It mints configured amount of tokens within minting period, where it divides this period into smaller sub Minters of equal lenght.
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
* reduction period length - defines how many mint Minters are in subperiod (subperiod = mint period * reduction_period_length)
* reduction factor - amount multiplying factor;

#### Examples

1. Four years halving minting that starts with 10 millions of tokens yearly

Minter configration:

* minting start: now
* Amount of minter Minters: 1
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

2. Linear minting of 100 millions of token during period of 10 years, next no emission

Minter configration:

* minting start: now
* Amount of minter Minters: 2
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

The Chain4Energy minter module contains the following configurations parameters:

| Key                  | Type                        | Description                     |
| -------------------- | --------------------------- | ------------------------------- |
| mintDenom     | string | Denom of minting token |
| minter     | Minter | Token emission configuration |

### Minter type

| Param                | Type                        | Description                     |
| -------------------- | --------------------------- | ------------------------------- |
| start     | Time | Token emission start time |
| Minters  | List of MintingPeriod | list of minting Minters |

### MintingPeriod type

| Param       | Type               | Description                                                             |
| ----------- | ------------------ | ----------------------------------------------------------------------- |
| SequenceId    | int32       | Minter period ordering SequenceId |
| period_end     | Time | Minter period end time |
| types     | Enum string | Minter period type. Allowed values:<br>- NO_MINTING<br>- LINEAR_MINTING<br>- EXPONENTIAL_STEP_MINTING;|
| LINEAR_MINTING  | LinearMinting    | Time linear minter configuration|
| EXPONENTIAL_STEP_MINTING | ExponentialStepMinting | Periodic reduction minter configuration |

### LinearMinting type

| Param   | Type         | Description              |
| ------- | ------------ | ------------------------ |
| amount    | sdk.Int       | An smount to mint lieary during the period |

### ExponentialStepMinting type

| Param   | Type | Description                    |
| ------- | ---- | ------------------------------ |
| mint_period | int32  | period of time of "mint_amount" token emission |
| mint_amount | sdk.Int   | amount to mint during "mint_period" |
| reduction_period_length | int32  | defines how many mint Minters are in subperiod (see **[periodic reduction minter](#periodic-reduction-minter)**) (subperiod = mint period * reduction_period_length) |
| reduction_factor | sdk.Dec   | amount multiplying factor |

### Example params

#### periodic reduction minter

Periodic reduction minter is block time based minter type. It mints configured amount of tokens within minting period, where it divides this period into smaller sub Minters of equal lenght.
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
* reduction period length - defines how many mint Minters are in subperiod (subperiod = mint period * reduction_period_length)
* reduction factor - amount multiplying factor;

#### Examples

See the configuration params for **[examples](#examples)** from **[Concept](#concepts)** section

1. Four years halving minting that starts with 10 millions of tokens yearly

```json

{
  "params": {
    "mint_denom": "uc4e",
    "minter": {
      "start": "2022-07-05T00:00:00Z",
      "Minters": [
        {
          "SequenceId": 1,
          "period_end": null,
          "type": "EXPONENTIAL_STEP_MINTING",
          "LINEAR_MINTING": null,
          "EXPONENTIAL_STEP_MINTING": {
            "mint_period": 31536000,
            "mint_amount": "10000000000000",
            "reduction_period_length": 4,
            "reduction_factor": "0.500000000000000000"
          }
        }
      ]
    }
  }
}

```

2. Linear minting of 100 millions of token during period of 10 years, next no emission

```json

{
  "params": {
    "mint_denom": "uc4e",
    "minter": {
      "start": "2022-07-05T00:00:00Z",
      "Minters": [
        {
          "SequenceId": 1,
          "period_end": "2023-07-05T00:00:00Z",
          "type": "EXPONENTIAL_STEP_MINTING",
          "LINEAR_MINTING": {
            "amount": "100000000000000",
          },
          "EXPONENTIAL_STEP_MINTING": null
        },
        {
          "SequenceId": 1,
          "period_end": null,
          "type": "NO_MINTING",
          "LINEAR_MINTING": null,
          "EXPONENTIAL_STEP_MINTING": null
        }
      ]
    }
  }
}

```

## State

Chain4Energy minter module state contains informations used  by current minting period.
Module state contains followng data:

| Key                  | Type                        | Description                     |
| -------------------- | --------------------------- | ------------------------------- |
| minter_state     | MinterState | current minting period state |
| state_history     | List of MinterState | previuos minting Minters final states |

### MinterState

| Key                  | Type                        | Description                     |
| -------------------- | --------------------------- | ------------------------------- |
| SequenceId     | int32 | current minting period SequenceId |
| amount_minted     | sdk.Int | amount minted by current minting period |
| remainder_to_mint     | sdk.Dec | decimal remainder - decimal amount that should be minted but was not Int. |
| last_mint_block_time     | sdk.Time | Time of last mint |
| remainder_from_previous_period     | sdk.Dec | decimal remainder left by previous minting period |

### Example state

```json

{
  "minter_state": {
    "SequenceId": 1,
    "amount_minted": "13766330043442",
    "remainder_to_mint": "0.415017757483510908",
    "last_mint_block_time": "2022-11-07T14:49:34.606250Z",
    "remainder_from_previous_period": "0.000000000000000000"
  },
  "state_history": []
}

```

## Events

Chain4Energy minter module emits the following events:

### BeginBlockers

#### Mint

| Type         | Attribute Key | Attribute Value |
| ------------ | ------------- | --------------- |
| bonded_ratio | Dec | |
| inflation | sdk.Dec | Minting block inflation level |
| amount | sdk.Int | Amount minted in block |

TODO remove bonded_ratio 
TODO add minter period id

## Queries

### Params query

Queries the module params.

See example reponse:

```json
{
  "params": {
    "mint_denom": "uc4e",
    "minter": {
      "start": "2022-07-05T00:00:00Z",
      "Minters": [
        {
          "SequenceId": 1,
          "period_end": null,
          "type": "EXPONENTIAL_STEP_MINTING",
          "LINEAR_MINTING": null,
          "EXPONENTIAL_STEP_MINTING": {
            "mint_period": 31536000,
            "mint_amount": "40000000000000",
            "reduction_period_length": 4,
            "reduction_factor": "0.500000000000000000"
          }
        }
      ]
    }
  }
}
```
### State query

Queries the module state.

See example reponse:

```json
{
  "minter_state": {
    "SequenceId": 1,
    "amount_minted": "13766330043442",
    "remainder_to_mint": "0.415017757483510908",
    "last_mint_block_time": "2022-11-07T14:49:34.606250Z",
    "remainder_from_previous_period": "0.000000000000000000"
  },
  "state_history": []
}
```

### Inflation query

Queries current inflation.

See example reponse:

```json
{
  "inflation": "0.102489480201216908"
}
```

## Genesis validations

TODO
