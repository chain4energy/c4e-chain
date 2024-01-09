# Chain4Energy minter module - cfeminter

## Abstract

Chain4Energy minter module provides functionality of controlled token emissions. Tokens are emitted according to configuration.

## Contents

1. **[Concept](#concepts)**
2. **[Params](#parameters)**
3. **[State](#state)**
4. **[Messages](#messages)**
5. **[Events](#events)**
6. **[Queries](#queries)**
7. **[Genesis validations](#genesis-validations)**

## Concepts

The purpose of `cfeminter` module is to provide token emission mechanism.

### Token emission mechanism

Token emission mechanism mints calculated amount of tokens per each block. 
Token amount is calculated accordingly to cfeminter module configuration params.
Tokens minting process is divided into separate minters where each minter has
different minting configuration. Token emission mechanism rules are defined within cfeminter 
module configuration params.
Simply, minting process configuration is a list of ordered minters,
where each minter has its own start and end time. Last minter cannot have end time because it
must be defined to work infinitely.

### Periodic continous vesting account

### Minters 

Ordered list of minters defines whole token emission process.
End time of one minter is a start time of the next minter in the minters list.
Each minter has its own minting configuration assigned.

### Minting configuration

Each minter has its specific minting configuration which defines general rules of token emission. 
Currently, cfeminter module supports following minting configs:
- no minting
- linear minting
- exponential step minting

Each minting configuration has its own specific set of parameters modifying token emission.

#### No minting

No minting is a simple minting configuration that mints nothing.
This minting configuration has no parameters.

#### Linear minting

Linear minting is block time based minting configuration. It mints predetermined amount of 
tokens within minter linearly. This minting configuration requires minter with end time 
since given amount of token needs to be minted in finite time period. So this 
minting configuration cannot be set in the last minter.

Linear minting configuration parameters:
* amount - amount of tokens to mint within minter period

#### Exponential step minting

Exponential step minting is block time based minting configuration. It mints predetermined amount
of tokens within minter, where it divides this minter into smaller subminters of 
equal length. Then within each subminter expected amount is minted, linearly. Expected 
amount of subminter minted tokens is equal to tokens minted by previous subminter 
multiplied by configured factor. For example initial minter amount is 40 million, 
multiplying factor set to 0.5 and step duration is four years, then:
* 1st subminter (first 4 years) mints 40 millions linearly 
* 2nd subminter (second 4 years) mints 20 millions linearly
* 3rd subminter (third 4 years) mints 10 millions linearly
* 4th subminter (fourth 4 years) mints 5 millions linearly
and so on.

This minter can mint infinitely.

Exponential step minting configuration parameters: 
* step duration - period of time mint amount is emitted
* amount - amount to mint for the first period
* amount multiplier - amount multiplying factor

## Examples

### Four years halving minting that starts with 40 million tokens and step duration set at 4 years

Minter configuration:
* minting start: now
* Amount of minter Minters: 1
* Minter 1:
    * end time: null
    * config: 
      * type: exponential step minting
      * step duration: 4 years
      * amount: 40 millions
      * amount multiplier: 0.5

Result:
* first 4 years mints 40 millions 
* second 4 years mints 20 millions
* third 4 years mints 10 millions
and so on

### Linear minting of 100 million of token during period of 10 years, next no emission

Minter configuration:

* minting start: now
* Amount of minter Minters: 2
* Minter 1:
    * end time: 10 years from now
    * config:
      * type: linear minting
      * amount: 100 millions
* Minter 2:
    * end time: null
  * config:
      * type: no minting

Result:
* 10 millions yearly for 10 years

## Parameters

The Chain4Energy minter module contains the following configurations parameters:

| Key         | Type            | Description                 |
|-------------|-----------------|-----------------------------|
| mint_denom  | string          | Denom of minting token      |
| start_time  | Time            | Token emission start time   |
| minters     | List of Minters | list of minters             |

### Minter type

| Param                    | Type          | Description                                                                                                                                                    |
|--------------------------|---------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| sequence_id              | uint32        | Minter ordering id                                                                                                                                             |
| end_time                 | Time          | Minter end time                                                                                                                                                |
| config                   | MinterConfigI | Minter configuration type that implements MinterConfigI interface. Allowed configuration types:<br>- NoMinting<br>- LinearMinting <br>- ExponentialStepMinting |                                                                 |

### NoMinting configuration

| Param  | Type      | Description                                  |
|--------|-----------|----------------------------------------------|
| type   | NoMinting | Minter configuration type                    |

### LinearMinting configuration

| Param  | Type          | Description                                  |
|--------|---------------|----------------------------------------------|
| type   | LinearMinting | Minter configuration type                    |
| amount | math.Int      | An amount to mint linearly during the period |

### ExponentialStepMinting configuration

| Param             | Type                     | Description                           |
|-------------------|--------------------------|---------------------------------------|
| type              | ExponentialStepMinting   | Minter configuration type             |
| step_duration     | uint32                   | period of time of token emission      |
| amount            | math.Int                 | amount to mint during "stepDuration"  |
| amount_multiplier | sdk.Dec                  | amount multiplying factor             |

### Example params

1. Four years halving minting that starts with 40 million tokens and step duration set at 4 years

```json
{
  "params": {
    "mint_denom": "uc4e",
    "start_time": "2022-07-05T00:00:00Z",
    "minters": [
      {
        "sequenceId": 1,
        "end_time": null,
        "config": {
          "@type": "/chain4energy.c4echain.cfeminter.ExponentialStepMinting",
          "step_duration": "126144000s",
          "amount": 40000000000000,
          "amount_multiplier": "0.500000000000000000"
        }
      }
    ]
  }
}
```

2. Linear minting of 100 million of token during period of 10 years, next no emission

```json
{
  "params": {
    "mint_denom": "uc4e",
    "start_time": "2022-07-05T00:00:00Z",
    "minters": [
    {
      "sequenceId": 1,
      "end_time": "2023-07-05T00:00:00Z",
      "config": {
        "@type": "/chain4energy.c4echain.cfeminter.LinearMinting",
        "amount": 100000000000000
      }
    },
    {
      "sequenceId": 2,
      "end_time": null,
      "config": {
        "@type": "/chain4energy.c4echain.cfeminter.NoMinting"
      }
    }
    ]
  }
}
```

## State

Chain4Energy minter module state contains information used by current minter.
Module state contains following data:

| Key                | Type                         | Description                   |
|--------------------|------------------------------|-------------------------------|
| minter_state       | MinterState                  | current minter state          |
| state_history      | List of MinterState          | previous minters final states |

### MinterState

| Key                            | Type       | Description                                                        |
|--------------------------------|------------|--------------------------------------------------------------------|
| sequence_id                    | uint32     | current minter sequenceId                                          |
| amount_minted                  | math.Int   | amount minted by current minter                                    |
| remainder_to_mint              | sdk.Dec    | amount that should have been minted in previous block but was not  |
| last_mint_block_time           | sdk.Time   | Time of last mint                                                  |
| remainder_from_previous_period | sdk.Dec    | amount that should have been minted in previous minter but was not | // TODO: rename to remainder_from_previous_minter

### Example state

```json
{
  "minter_state": {
    "sequence_id": 1,
    "amount_minted": "13766330043442",
    "remainder_to_mint": "0.415017757483510908",
    "last_mint_block_time": "2022-11-07T14:49:34.606250Z",
    "remainder_from_previous_period": "0.000000000000000000"
  },
  "state_history": []
}
```

## Messages

### Burn Tokens
Burns a specified amount of tokens from the given address. This process permanently reduces the total supply of the tokens.

MsgBurn Message
Represents a message to burn tokens from an account.

#### Structure
```go
type MsgBurn struct {
Address string
Amount  sdk.Coins
}
```

#### Parameters

#### State Modifications
* Validates if the Address is a valid bech32 account address.
* Validates if the Amount is positive and not nil.
* Ensures that the Address has a balance greater than or equal to the Amount.
* Sends the Amount from the Address to the `types.ModuleName` module account.
* Burns the Amount from the `types.ModuleName` module account, thereby reducing the total supply of those tokens.

## Events

Chain4Energy minter module emits the following events:

### BeginBlockers

#### EventMint

| Type      | Attribute Key | Attribute Value                 |
|-----------|---------------|---------------------------------|
| EventMint | bonded_ratio  | {tokens_boneded_ratio}          |
| EventMint | inflation     | {minting_block_inflation_level} |
| EventMint | amount        | {amount_minted_in_block}        |
[//]: # (TODO remove bonded_ratio )
[//]: # (TODO add minter period id)

### Handlers for `MsgBurn`

| Type               | Attribute Key     | Attribute Value                         |
|--------------------|-------------------|-----------------------------------------|
| burn               | burner            | {sender_address}                        |
| burn               | amount            | {amount}                                |
| message            | action            | /chain4energy.c4echain.cfeclaim.MsgBurn |
| message            | sender            | {sender_address}                        |
| transfer           | recipient         | {module_account}                        |
| transfer           | sender            | {creator}                               |
| transfer           | amount            | {amount}                                |



## Queries

### Params query

Queries the module params.

See example response:

```json
{
  "params": {
    "mint_denom": "uc4e",
    "start_time": "2022-07-05T00:00:00Z",
    "minters": [
      {
        "sequenceId": 1,
        "end_time": null,
        "config": {
          "@type": "/chain4energy.c4echain.cfeminter.ExponentialStepMinting",
          "step_duration": "126144000s",
          "amount": 40000000000000,
          "amount_multiplier": "0.500000000000000000"
        }
      }
    ]
  }
}
```
### State query

Queries the module state.

See example response:

```json
{
  "minter_state": {
    "sequence_id": 1,
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

See example response:

```json
{
  "inflation": "0.102489480201216908"
}
```

## Genesis validations

[//]: # (TODO)
