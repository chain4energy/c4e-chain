# Chain4Energy claiming module - cfeclaim

## Abstract

Chain4Energy claiming module provides the functionality of creating campaigns, adding
missions, and claiming tokens based on pre-set conditions. It emphasizes the importance
of transparency and accountability in the token distribution process, aligning with
mission-driven campaigns.

## Contents

1. **[Concept](#concepts)**
2. **[Messages](#messages)**
3. **[Events](#events)**
4. **[Queries](#queries)**
5. **[Invariants](#invariants)**
6. **[Genesis validations](#genesis-initialization-and-validation)**

## Concepts

### Campaign

Each campaign is a structured initiative with a specific goal or target. A campaign
can house multiple missions and will have tokens allotted, which can be claimed upon
the completion of those missions. Each campaign has a specific timeline, a beginning
and an end, within which missions can be added and tokens can be claimed.

Campaigns have the following parameters:
* owner - Address of the campaign creator.
* name - Name of the campaign.
* description - A brief overview of the campaign.
* campaign_type - Type or category of the campaign.
  * VESTING_POOL - Campaign that is created from a vesting pool.
  * DEFAULT - Campaign that is created from an account balance.
* removable_claim_records - Indicates if claim records can be removed after start time of a campaign.
* feegrant_amount - Amount designated for fee grants.
* initial_claim_free_amount - The number of tokens to be released immediately
  (regardless of the lockup and vesting period) in the case of an initial claim.
* free - Percentage of tokens to be released during mission claim (regardless of lockup and vesting period).
* start_time - Timestamp for the start time of the campaign.
* end_time - Timestamp for the end time of the campaign.
* lockup_period - Duration of time during which tokens cannot be withdrawn.
* vesting_period - Duration of vesting period for tokens.
* vesting_pool_name - Name of the vesting pool associated with the campaign (only when campaign is of VESTING_POOL type).

### Mission

A mission is a specific task or objective within a campaign. Upon completion of a mission,
users can claim tokens. Each mission also has a percentage weight field that indicates
how many tokens out of the total amount of tokens assigned to a
given user the user will get for this mission. Example - the user is assigned 100 tokens,
the weight of the mission is 0.4, so the user for completing and claiming this mission
will get 40 tokens.

Missions have the following parameters:
* id - Identifier of the mission.
* campaign_id - ID of the campaign under which the mission resides.
* name - Name of the mission.
* description - Brief about the mission.
* mission_type - Type or category of the mission.
* weight - Mission weight percentage. This is the percentage of tokens that will be sent to the user at the time of claiming the mission.
* claim_start_date - Timestamp for the date from which claims can be initiated for this mission.

Currently, we distinguish following types of missions:
* INITIAL_CLAIM - required mission added at the beginning of the campaign along with adding the campaign.
* DELEGATE - require the user to delegate tokens before claiming these missions
* VOTE - require the user to vote on the proposal before claiming these missions
* CLAIM - the user can immediately redeem it without any additional actions


### User Entry
Every user in the ecosystem is identified by an entry. This entry holds information
related to the user's claims and the records associated with those claims.

User Entries have the following parameters:

* address - User blockchain address.
* claim_records - List of all claim records associated with the user. Each record holds
  information about a specific campaign assigned to user.

### Claim Record
Claim record stores information about the user assigned to a given campaign
(including the number of tokens assigned to him, ids of completed and claimed
missions)

Claim Records have the following parameters:

* campaign_id - Identifier for the campaign related to this claim record.
* address - User address to which the claimed tokens will be sent.
* amount - The amount of tokens assigned to the user for the entire campaign
* completed_missions - List of mission IDs that the user has completed.
* claimed_missions - List of mission IDs for which the user has made a claim.

### Claim Record Entry
While similar to a Claim Record, a Claim Record Entry contains more granular information
about the claim related to a specific campaign. It is used while adding a new claim record.

Claim Record Entries have the following parameters:

* campaign_id - Identifier for the campaign related to this claim entry.
* user_entry_address - The address of the user associated with this claim entry.
* amount -  The amount of tokens assigned to the user for the entire campaign

### Fee granting

Feegrant is a special option in a campaign that allows for the first use of an initial claim for
each user of that campaign. This option was created so that users who did not have a blockchain
account and therefore, did not have tokens to cover transaction fees, could perform the first
mission (initial claim) of the campaign. Feegrant takes on a specific value. If feegrant is set
to a value greater than zero in a given campaign, when adding each user entry, it is checked whether
the user who adds the user entries is also able to cover this feegrant. To prevent situations
in which the founder of a campaign does not have enough tokens in their account to pay for
feegrant for another user, an additional, special module account is created, which stores tokens
to pay for feegrant. If a user correctly uses initial claim (transaction is successful),
feegrant is revoked. Tokens that were not used by users are treated in the same way as
campaign’s leftovers during the closure of the campaign.

### Vesting pool reservation mechanism
Function of the mechanism is to allow vesting pool owners to reserve vesting
pool tokens for other module functionality. In case of a claim module,
tokens are reserved for claiming purposes.
The mechanism will not expose blockchain messages API, it is an internal mechanism only.

Each reservation has its unique id within one vesting pool. Reservation can be
increased and decreased by the owner:
* increased by adding reservation with the same id
* decreased by dedicated method

In the case of the claim module, when user entry of vesting pool campaign is added
reservation is increased, when removed decreased.

## Messages

### Create Campaign

Create a new campaign with a set of specified parameters.

The `MsgCreateCampaign` can be submitted by any valid account to initiate the creation
of a new campaign using the given parameters.

#### Structure

```go
type MsgCreateCampaign struct {
Owner                  string
Name                   string
Description            string
CampaignType           types.CampaignType
RemovableClaimRecords  bool
FeegrantAmount         math.Int
InitialClaimFreeAmount math.Int
Free                   sdk.Dec
StartTime              time.Time
EndTime                time.Time
LockupPeriod           time.Duration
VestingPeriod          time.Duration
VestingPoolName        string
}
```

#### Parameters

| Param                  | Description                                                      |
|------------------------|------------------------------------------------------------------|
| Owner                  | Address of the campaign owner                                    |
| Name                   | Name of the campaign                                             |
| Description            | Description of the campaign                                      |
| CampaignType           | Type of the campaign (e.g., VestingPoolCampaign)                 |
| RemovableClaimRecords  | Indicates if claim records can be removed                        |
| FeegrantAmount         | Amount set for fee grant                                         |
| InitialClaimFreeAmount | Initial free amount to claim                                     |
| Free                   | Decrement value to make claims free                              |
| StartTime              | Starting time of the campaign                                    |
| EndTime                | Ending time of the campaign                                      |
| LockupPeriod           | Duration for which the amount remains locked                     |
| VestingPeriod          | Duration of the vesting                                          |
| VestingPoolName        | Name associated with the vesting pool                            |

#### State Modifications

- Validates various campaign parameters including the start and end times, fee grant amount, among others.
- Confirms that the end time is post the current block time.
- Establishes a new `Campaign` structure using the given parameters.
- Validates the vesting pool if the campaign type happens to be `VestingPoolCampaign`.
- Incorporates the new campaign into the store and allots it a unique ID.
- An `initialClaim` mission is automatically added to the campaign.
- Triggers the `EventNewCampaign` event, signaling the establishment of a new campaign.

## Enable Campaign

Enable an existing campaign with specified start and end times.

The `MsgEnableCampaign` can be used to modify the start and end times of a campaign and to activate it.

### Structure

```go
type MsgEnableCampaign struct {
Owner       string
CampaignId  uint64
StartTime   *time.Time (optional)
EndTime     *time.Time (optional)
}
```

### Parameters

| Param       | Description                                         |
|-------------|-----------------------------------------------------|
| Owner       | Address of the campaign owner                       |
| CampaignId  | ID of the campaign to be enabled                    |
| StartTime   | New starting time of the campaign (optional)        |
| EndTime     | New ending time of the campaign (optional)          |

### State Modifications

- Retrieves the specified campaign by the `CampaignId`.
- Validates the ownership of the campaign.
- Checks and updates the start and end times if provided, ensuring:
  - The start time is before the end time.
  - The start time isn't the same as the end time.
- Validates that the campaign isn't already enabled.
- Sets the campaign's status to `enabled`.
- Updates the campaign in the state with the new details.
- Emits the `EventEnableCampaign` event, signaling the enabling of the specified campaign.

## Close Campaign

Closes an existing campaign based on the specified parameters.

The `MsgCloseCampaign` can be submitted by the campaign owner to signal the end of a
campaign, ensuring all assets and values are properly reconciled.

### Structure

```go
type MsgCloseCampaign struct {
Owner      string
CampaignId uint64
}
```

### Parameters

| Param      | Description                                                  |
|------------|--------------------------------------------------------------|
| Owner      | Address of the campaign owner                                |
| CampaignId | ID of the campaign to be closed                              |

### State Modifications

- Retrieves the specified campaign from the store.
- Validates the parameters for closing the campaign, ensuring the owner is the
  one who initiated the closure and the campaign has not yet ended.
- Returns all remaining assets to the campaign owner.
- Sets the campaign's current amount to zero and disables the campaign.
- Emits an `EventCloseCampaign` event, signaling the closure of a campaign.

## Remove Campaign

Remove an existing campaign based on the specified parameters.

The `MsgRemoveCampaign` can be submitted by the campaign owner to initiate
the removal of a campaign. All the assets within the campaign are returned to the owner before removal.
This message can be only used if the campaign is not enabled.

### Structure

```go
type MsgRemoveCampaign struct {
Owner      string
CampaignId uint64
}
```

### Parameters

| Param      | Description                              |
|------------|------------------------------------------|
| Owner      | Address of the campaign owner            |
| CampaignId | ID of the campaign to be removed         |

### State Modifications

- Retrieves the specified campaign from the store using its ID.
- Validates the removal parameters to ensure that the owner is the one initiating the removal
  and that the campaign is not enabled.
- Returns all current assets and fee grants of the campaign to the owner.
- Removes the campaign and all associated missions from the store.
- Emits an `EventRemoveCampaign` event, signaling the removal of a campaign.

## Add Mission

The `MsgAddMission` message type facilitates the addition of a mission to an existing campaign on
the blockchain. A mission, in the context of this chain, represents a specific task or objective
within a campaign, and has parameters like weight, claim start date, and type, dictating its
behavior. By submitting this message, the campaign owner can diversify the
range of activities within a campaign, giving participants more avenues to engage with the campaign.

### Structure

```go
type MsgAddMission struct {
Owner          string
CampaignId     uint64
Name           string
Description    string
MissionType    types.MissionType
Weight         sdk.Dec
ClaimStartDate *time.Time
}
```

### Parameters
| Param           | Description                                                                                                     |
|-----------------|-----------------------------------------------------------------------------------------------------------------|
| Owner           | Address of the individual or entity responsible for this mission and the campaign it belongs to.                |
| CampaignId      | A unique identifier for the campaign to which this mission will be added.                                       |
| Name            | A concise title for the mission, allowing participants to quickly identify its nature.                          |
| Description     | An elaborate description providing in-depth details about what the mission entails.                             |
| MissionType     | Classifies the mission, e.g., InitialClaim might indicate a task that can only be done at the campaign's start. |
| Weight          | Represents the mission's importance or significance within the campaign.                                        |
| ClaimStartDate  | Specifies the date from which claims related to this mission can begin, ensuring phased or timed participation. |

### State Modifications
- Before any modifications, the system checks mission parameters such as name, mission type,
  and weight for validity. Incorrect or malicious inputs are rejected.
- The specified campaign's existence and the authenticity of the campaign owner are verified.
- All existing missions within the campaign are summed up with the new mission's weight to ensure the
  total doesn't exceed a weight of 1. This ensures no single mission or a combination overshadows the
  rest in significance.
- The state checks if the campaign is still active and not ended. Missions can't be added to an expired
  or inactive campaign.
- Upon successful validation, the new mission is appended to the campaign's list of missions and a
  unique ID is assigned to it.
- An EventAddMission event is emitted to signal the blockchain about the addition of a new mission.
  This can be used by listeners for various purposes like notifications.

## Initial Claim

Allows a user to claim a specified amount from a campaign using its unique ID.

The `MsgInitialClaim` message can be sent by any valid account eligible to claim rewards from
the specified campaign.

### Structure

```go
type MsgInitialClaim struct {
Claimer           string
CampaignId        uint64
DestinationAddress string
}
```

### Parameters
| Param               | Description                                                         |
|---------------------|---------------------------------------------------------------------|
| Claimer             | Address of the user attempting to claim the reward.                 |
| CampaignId          | Unique identifier for the campaign from which the user is claiming. |
| DestinationAddress  | The address where the claimed amount should be deposited.           |

### State Modifications
- The method fetches necessary data such as the campaign details, associated mission, the user's
  entry, and the claim record using the provided CampaignId.
- The method validates the given DestinationAddress ensuring it's not a blocked or invalid address.
- The destination address is assigned to the claim record, and the mission status is updated to reflect
  completion.
- The claimable amount for the user's initial claim is computed. This amount depends on the user's
  interactions and the campaign's configuration.
- For all campaign types, the claimable amount is sent to a periodic continuous vesting account.
  This ensures that the claimed rewards are vested according to the campaign's rules.
- After the claim, the campaign's current balance is decremented by the claimable amount.
- Once the initial claim is processed successfully, an EventInitialClaim event is emitted to indicate
  the completion of the claim.

## Claim

Allows a user to claim rewards based on the completion of a specific mission within a campaign.

The `MsgClaim` can be submitted by any valid account, which has completed the required
mission, to claim the associated rewards.

### Structure

```go
type MsgClaim struct {
Claimer    string
CampaignId uint64
MissionId  uint64
}
```

### Parameters
| Param      | Description                                                                                                   |
|------------|---------------------------------------------------------------------------------------------------------------|
| Claimer    | Address of the user attempting to claim the reward.                                                           |
| CampaignId | Unique identifier for the campaign from which the user is claiming.                                           |
| MissionId  | Unique identifier for the mission within the campaign that the user has completed and is attempting to claim. |

### State Modifications
- Retrieves data related to the campaign, the specific mission, user entry, and the claim
  record using the provided CampaignId and MissionId.
- Confirms if the initial mission associated with the campaign has been claimed.
- If the mission type is specifically a "claim" mission, it's marked as completed.
- Determines the amount eligible to be claimed based on the provided mission.
- For all campaign types, the claimable amount is sent to a periodic continuous vesting account.
  This ensures that the claimed rewards are vested according to the campaign's rules.
- Decreases the current balance of the campaign by the claimed amount.
- Emits the EventClaim event, indicating a successful claim.

## Add Claim Records

This message is responsible for adding claim records to an existing campaign. When adding
claim records, the module takes into consideration the type of the campaign, the provided
claim records, the owner of the campaign, and the associated fees. If the campaign type
is a `VestingPoolCampaign`, special validation and adjustments are made based on the vesting
denomination. If fee grants are associated with the campaign, the module ensures that the
necessary fee grants are given to the appropriate addresses.

The `MsgAddClaimRecords` can be submitted by the campaign owner to add claim records to a given campaign.

### Structure

```go
type MsgAddClaimRecords struct {
Owner              string
CampaignId         uint64
ClaimRecordEntries []*types.ClaimRecordEntry
}
```

### Parameters
| Param              | Description                                                                   |
|--------------------|-------------------------------------------------------------------------------|
| Owner              | Address of the campaign owner.                                                |
| CampaignId         | ID of the campaign to which the claim records are being added.                |
| ClaimRecordEntries | A list of claim record entries that specify the users and amounts for claims. |
### State Modifications
- Validation of Input:
  - Ensures that the campaign associated with CampaignId exists.
  - Validates that the Owner address matches the owner of the campaign.
  - Ensures the campaign has not yet ended.
  - If the campaign is of type VestingPoolCampaign, validates the claim record entries against
    the vesting denomination.
  - Otherwise, validates the claim record entries in general.

- Calculation of Fee Grants:
  - Calculates the total fees required for granting fees based on the fee grant amount and the
    number of claim record entries.
  - Checks if the campaign owner has sufficient balance to cover these fees.

- Adjustment for Vesting Pool Campaigns:
  - If the campaign is of type VestingPoolCampaign, it adds a new reservation to the associated vesting pool.
  - Verifies the owner's balance is sufficient to cover the fees.
  - Verifies the vesting pool balance is sufficient to cover the claim records' amounts.

- Adjustment for Default Campaigns:
  - If the campaign is of a default type, the module calculates the sum of fees and the amounts
    from the claim records.
  - Ensures that the owner has sufficient balance to cover this combined amount.

- Setting Up Fee Grants:
  - If a positive fee grant amount is specified, the module sets up a new fee grant account.
  - Transfers the necessary coins from the owner to this fee grant account.
  - Grants the fee allowance to all claim record users.

- User Entry Creation:
  - For each claim record entry, the module creates or updates the user entry associated with the claim.
  - If a user does not have an account in the system, a new account is created for them.

- An EventAddClaimRecords event is emitted, which details the owner, campaign ID,
  total amount across all claim records, and the number of claim records.

## Delete Claim Record

Allows a campaign owner to remove an existing claim record for a specific user within a
given campaign. The `MsgDeleteClaimRecord` action encompasses several validations, balance
calculations, and state modifications, ensuring that the claim record deletion is properly
recorded and the associated funds are correctly handled.
If the campaign has the removableClaimRecords flag set to true, you can delete claim records at any time during the campaign.
Otherwise, claim records can only be deleted until the campaign is enabled.

The `MsgDeleteClaimRecord` can be submitted by the campaign owner to delete a claim
record for a given user associated with a campaign.

### Structure

```go
type MsgDeleteClaimRecord struct {
Owner       string
CampaignId  uint64
UserAddress string
}
```

#### Parameters

| Param        | Description                                                              |
|--------------|--------------------------------------------------------------------------|
| Owner        | Address of the campaign owner; responsible for initiating the deletion   |
| CampaignId   | Unique identifier of the campaign from which the claim record is deleted |
| UserAddress  | Address of the user whose claim record is to be deleted                  |

### State Modifications

- It ensures that the caller, identified by the Owner address, is indeed the owner of the
  campaign
- It checks the campaign's data to confirm the presence of a claim record associated with the
  provided UserAddress.
- Calculates the total amount linked with the claim record to be deleted. The amount is derived
  from the claim record and any associated missions the user may have claimed.
- If the campaign is of type VestingPoolCampaign, the associated vesting pool reservations are
  adjusted, and necessary funds are released.
- For other types of campaigns, the calculated amount is sent back to the campaign owner's account.
  This ensures that the campaign's balance remains accurate after the claim record deletion.
- If the campaign has set up a FeegrantAmount, any associated fee grant with the user is revoked.
  Any remaining funds from the fee grant are transferred back to the campaign owner.
- The claim record associated with the UserAddress is removed. This step ensures that the user
  no longer has a claim on the campaign after deletion.
- The campaign's current and total amounts are decremented based on the amount associated with
  the deleted claim record. This ensures the campaign's financial data remains accurate.
- At the end of the process, an EventDeleteClaimRecord event is emitted. This event logs the
  details of the deletion, including the involved addresses and the amount associated with the deleted claim record.


## Events

### Handlers for `MsgCreateCampaign`

| Type        | Attribute Key          | Attribute Value                                   |
|-------------|------------------------|---------------------------------------------------|
| NewCampaign | id                     | {campaign_id}                                     |
| NewCampaign | owner                  | {owner_address}                                   |
| NewCampaign | name                   | {campaign_name}                                   |
| NewCampaign | description            | {campaign_description}                            |
| NewCampaign | campaign_type          | {campaign_type_name}                              |
| NewCampaign | feegrant_amount        | {feegrant_amount}                                 |
| NewCampaign | initial_claim_free_amt | {initial_claim_free_amount}                       |
| NewCampaign | start_time             | {start_time}                                      |
| NewCampaign | end_time               | {end_time}                                        |
| NewCampaign | lockup_period          | {lockup_duration}                                 |
| NewCampaign | vesting_period         | {vesting_duration}                                |
| NewCampaign | vesting_pool_name      | {vesting_pool_name}                               |
| message     | action                 | /chain4energy.c4echain.cfeclaim.MsgCreateCampaign |
| message     | sender                 | {sender_address}                                  |

### Handlers for `MsgCloseCampaign`

| Type          | Attribute Key   | Attribute Value                                  |
|---------------|-----------------|--------------------------------------------------|
| CloseCampaign | owner           | {owner_address}                                  |
| CloseCampaign | campaign_id     | {campaign_id}                                    |
| message       | action          | /chain4energy.c4echain.cfeclaim.MsgCloseCampaign |
| message       | sender          | {sender_address}                                 |

### Handlers for `MsgRemoveCampaign`

| Type           | Attribute Key   | Attribute Value                                   |
|----------------|-----------------|---------------------------------------------------|
| RemoveCampaign | owner           | {owner_address}                                   |
| RemoveCampaign | campaign_id     | {campaign_id}                                     |
| message        | action          | /chain4energy.c4echain.cfeclaim.MsgRemoveCampaign |
| message        | sender          | {sender_address}                                  |


### Handlers for `MsgAddMission`

| Type         | Attribute Key    | Attribute Value                               |
|--------------|------------------|-----------------------------------------------|
| AddMission   | id               | {mission_id}                                  |
| AddMission   | owner            | {owner_address}                               |
| AddMission   | campaign_id      | {campaign_id}                                 |
| AddMission   | name             | {mission_name}                                |
| AddMission   | description      | {mission_description}                         |
| AddMission   | mission_type     | {mission_type_name}                           |
| AddMission   | weight           | {mission_weight}                              |
| AddMission   | claim_start_date | {claim_start_date}                            |
| message      | action           | /chain4energy.c4echain.cfeclaim.MsgAddMission |
| message      | sender           | {sender_address}                              |


### Handlers for `MsgInitialClaim`

| Type         | Attribute Key       | Attribute Value                                  |
|--------------|---------------------|--------------------------------------------------|
| InitialClaim | claimer             | {claimer_address}                                |
| InitialClaim | campaign_id         | {campaign_id}                                    |
| InitialClaim | destination_address | {destination_address}                            |
| InitialClaim | amount              | {claimed_amount}                                 |
| message      | action              | /chain4energy.c4echain.cfeclaim.MsgInitialClaim  |
| message      | sender              | {sender_address}                                 |


### Handlers for `MsgClaim`

| Type      | Attribute Key  | Attribute Value                          |
|-----------|----------------|------------------------------------------|
| Claim     | claimer        | {claimer_address}                        |
| Claim     | campaign_id    | {campaign_id}                            |
| Claim     | mission_id     | {mission_id}                             |
| Claim     | amount         | {claimed_amount}                         |
| message   | action         | /chain4energy.c4echain.cfeclaim.MsgClaim |
| message   | sender         | {sender_address}                         |

### Handlers for `MsgAddClaimRecords`

| Type            | Attribute Key           | Attribute Value                                    |
|-----------------|-------------------------|----------------------------------------------------|
| AddClaimRecords | owner                   | {owner_address}                                    |
| AddClaimRecords | campaign_id             | {campaign_id}                                      |
| AddClaimRecords | total_amount            | {total_amount_added}                               |
| AddClaimRecords | number_of_claim_records | {number_of_claim_records}                          |
| message         | action                  | /chain4energy.c4echain.cfeclaim.MsgAddClaimRecords |
| message         | sender                  | {sender_address}                                   |

### Handlers for `MsgDeleteClaimRecord`

| Type              | Attribute Key       | Attribute Value                                      |
|-------------------|---------------------|------------------------------------------------------|
| DeleteClaimRecord | owner               | {owner_address}                                      |
| DeleteClaimRecord | campaign_id         | {campaign_id}                                        |
| DeleteClaimRecord | user_address        | {user_address}                                       |
| DeleteClaimRecord | claim_record_amount | {claim_record_amount}                                |
| message           | action              | /chain4energy.c4echain.cfeclaim.MsgDeleteClaimRecord |
| message           | sender              | {sender_address}                                     |


## Queries

### User Entry Query
Queries a specific user entry by the address.

Endpoint: `/c4e/claim/v1beta1/user_entry/{address}`

```json
{
  "user_entry":{
    "address":"example_address",
    "claim_records":[
      {
        "campaign_id":1,
        "address":"example_address",
        "amount":[
          {
            "denom":"uc4e",
            "amount":"1000"
          }
        ],
        "completedMissions":[
          1,
          2
        ],
        "claimedMissions":[
          2
        ]
      }
    ]
  }
}
```

### Users Entries Query
Queries a list of all user entries.

Endpoint: /c4e/claim/v1beta1/users_entries

```json
{
  "users_entries":[
    {
      "address":"example_address_1",
      "claim_records":[
        {
          "campaign_id":1,
          "address":"example_address",
          "amount":[
            {
              "denom":"uc4e",
              "amount":"1000"
            }
          ],
          "completedMissions":[
            1,
            2
          ],
          "claimedMissions":[
            2
          ]
        }
      ]
    },
    {
      "address":"example_address_2",
      "claim_records":[
        {
          "campaign_id":1,
          "address":"example_address",
          "amount":[
            {
              "denom":"uc4e",
              "amount":"1000"
            }
          ],
          "completedMissions":[
            1,
            2
          ],
          "claimedMissions":[
            2
          ]
        }
      ]
    }
  ],
  "pagination":{
    "next_key": null,
    "total": "2"
  }
}
```

### Mission Query
Queries a specific mission based on campaign ID and mission ID.

Endpoint: `/c4e/claim/v1beta1/mission/{campaign_id}/{mission_id}`

```json
{
  "mission":{
    "id":1,
    "campaign_id":1,
    "name":"Example Mission",
    "description":"Description for example mission",
    "missionType":"INITIAL_CLAIM",
    "weight":"0.3",
    "claim_start_date":"2023-01-01T00:00:00Z"
  }
}
```

### Missions Query
Queries a list of all mission items.

Endpoint: `/c4e/claim/v1beta1/missions`

```json
{
  "missions":[
    {
      "id":1,
      "campaign_id":1,
      "name":"Example Mission 1",
      "description":"Description for example mission 1",
      "missionType":"INITIAL_CLAIM",
      "weight":"0.3",
      "claim_start_date":"2023-01-01T00:00:00Z"
    },
    {
      "id":2,
      "campaign_id":1,
      "name":"Example Mission 2",
      "description":"Description for example mission 2",
      "missionType":"CLAIM",
      "weight":"0.3",
      "claim_start_date":"2023-01-01T00:00:00Z"
    }
  ],
  "pagination":{
    "next_key": null,
    "total": "2"
  }
}
```

### Campaign Missions Query
Queries a list of mission items for a given campaign.

Endpoint: `/c4e/claim/v1beta1/missions/{campaign_id}`

```json
{
  "missions":[
    {
      "id":1,
      "campaign_id":1,
      "name":"Example Mission 1",
      "description":"Description for example mission 1",
      "missionType":"INITIAL_CLAIM",
      "weight":"0.3",
      "claim_start_date":"2023-01-01T00:00:00Z"
    },
    {
      "id":2,
      "campaign_id":1,
      "name":"Example Mission 2",
      "description":"Description for example mission 2",
      "missionType":"CLAIM",
      "weight":"0.3",
      "claim_start_date":"2023-01-01T00:00:00Z"
    }
  ],
  "pagination":{
    "next_key": null,
    "total": "2"
  }
}
```



### Campaign Query
Queries a specific campaign by its ID.

Endpoint: `/c4e/claim/v1beta1/campaign/{campaign_id}`

```json
{
  "campaign": {
    "id": "1",
    "owner": "c4e1p0smw03cwhqn05fkalfpcr0ngqv5jrpnx2cp54",
    "name": "Example campaign",
    "description": "Description for example campaign",
    "campaignType": "VESTING_POOL",
    "removable_claim_records": false,
    "feegrant_amount": "0",
    "initial_claim_free_amount": "0",
    "free": "0.010000000000000000",
    "enabled": false,
    "start_time": "2030-01-01T00:00:00Z",
    "end_time": "2031-01-01T00:00:00Z",
    "lockup_period": "15811200s",
    "vesting_period": "7862400s",
    "vestingPoolName": "Fairdrop",
    "campaign_total_amount": [
      {
        "denom": "uc4e",
        "amount": "10000000000"
      }
    ],
    "campaign_current_amount": [
      {
        "denom": "uc4e",
        "amount": "10000000000"
      }
    ]
  }
}
```

### Campaigns Query
Queries a list of all campaign items.

Endpoint: `/c4e/claim/v1beta1/campaigns`

```json
{
  "campaigns": [
    {
      "id": "0",
      "owner": "c4e1dsm96gwcv35m4rqd93pzcsztpkrqe0ev7getj8",
      "name": "Example Campaign",
      "description": "Example campaign description",
      "campaignType": "VESTING_POOL",
      "removable_claim_records": true,
      "feegrant_amount": "0",
      "initial_claim_free_amount": "0",
      "free": "0.000000000000000000",
      "enabled": false,
      "start_time": "2030-01-01T00:00:00Z",
      "end_time": "2031-01-01T00:00:00Z",
      "lockup_period": "63072000s",
      "vesting_period": "63072000s",
      "vestingPoolName": "Moondrop",
      "campaign_total_amount": [
        {
          "denom": "uc4e",
          "amount": "7280002000000"
        }
      ],
      "campaign_current_amount": [
        {
          "denom": "uc4e",
          "amount": "7280002000000"
        }
      ]
    },
    {
      "id": "1",
      "owner": "c4e1p0smw03cwhqn05fkalfpcr0ngqv5jrpnx2cp54",
      "name": "Example campaign",
      "description": "Example campaign description",
      "campaignType": "VESTING_POOL",
      "removable_claim_records": false,
      "feegrant_amount": "0",
      "initial_claim_free_amount": "0",
      "free": "0.010000000000000000",
      "enabled": false,
      "start_time": "2030-01-01T00:00:00Z",
      "end_time": "2031-01-01T00:00:00Z",
      "lockup_period": "15811200s",
      "vesting_period": "7862400s",
      "vestingPoolName": "Fairdrop",
      "campaign_total_amount": [
        {
          "denom": "uc4e",
          "amount": "8999999989680"
        }
      ],
      "campaign_current_amount": [
        {
          "denom": "uc4e",
          "amount": "8999999989680"
        }
      ]
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "2"
  }
}
```

## Invariants

## Campaign Current Amount Sum Invariant

Invariant validates the campaign's current amount state. This checks if the
sum of campaign claims left is consistent with the `cfeclaim` module account balance.

### Definition

- **Name**: `campaignCurrentAmountSumInvariant`
- **Route**: `"campaigns-current-amount"`
- **Module**: `cfeclaim`

### Logic

1. **Retrieve All Campaigns**: The function fetches all campaigns within the context.
2. **Empty Check**: If there are no campaigns, an invariant message indicates that the campaigns list is empty,
   and the check returns false.
3. **Vesting vs. Default Campaigns**: The function segregates campaigns based on
   their type (`VestingPoolCampaign` vs. `DefaultCampaign`) and performs different calculations
   for each.
- For `VestingPoolCampaign`:
  - The function adds up the current amount of each vesting campaign.
  - It also fetches and sums up the amount locked in reservations for the vesting
    campaign using the `vestingKeeper.MustGetVestingPoolReservation` method.
- For `DefaultCampaign`:
  - The function adds up the current amount of each default campaign.

4. **Module Account Balance Check**:
- Retrieves the coins associated with the `cfeclaim` module account.
- Checks if this balance is equal to the total current amount of default
  campaigns. If not, an invariant error message is returned, and the check returns true.

5. **Vesting Campaign Amount Check**:
- Compares the total current amount of vesting campaigns with the sum of
  tokens locked in vesting pool reservations.
- If they aren't equal, an invariant error message is returned, and the
  check returns true.

6. **Successful Invariant Check**:
- If both checks are successful, an invariant message indicates that
  the claims left sum is equal to the cfeclaim module account balance, and the check
  returns false.

## Genesis Initialization and Validation

The genesis initialization and validation process ensures a consistent and
predefined initial state. This documentation breaks down the purpose and
workings of various functions involved.

`InitGenesis` initializes the module's state from a provided genesis state,
encompassing campaigns, missions, and user entries.

```go
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
setCampaigns(ctx, k, genState.Campaigns, genState.CampaignCount)
setMissions(ctx, k, genState.Missions, genState.MissionCounts)
setUsersEntries(ctx, k, genState.UsersEntries)
}
```

### Set campaigns
It validates each campaign's parameters using ValidateCampaignParams.
Invalid parameters cause a panic. Once validated, the campaign data is set.

### Set missions
It ensures the associated campaign exists for each mission and validates
the mission using ValidateAddMission. Any issues lead to a panic. Valid missions are set.

### Set users entries
Each user entry undergoes validation with Validate. Subsequently,
claim records are verified using validateClaimRecords. Successful validation
results in setting the user entry.

### Helper functions
- validateClaimRecords: Validates claim records, ensuring associated
  campaigns and missions are correct.
- ValidateAddMission: Validates the mission parameters and ensures its
  association with a valid campaign.
- ValidateCampaignParams: Validates the campaign parameters based on campaign type.
- UserEntry.Validate: Validates a user's entry structure, ensuring a valid
  address format and non-duplicacy of campaign and mission IDs.
- The genesis validation functions ensure the integrity and consistency of
  initial data. Any inconsistencies or violations lead to a panic, preventing the
  initialization of an invalid state.