package types

// staking module event types
const (
	EventTypeDelegate        = "vesting_delegate"
	EventTypeUnbond          = "vesting_unbond"
	EventTypeRedelegate      = "vesting_redelegate"
	EventTypeWithdrawRewards = "vesting_withdraw_rewards"
	EventTypeWithdrawVesting = "vesting_withdraw_vesting"
	EventTypeVest            = "vesting_vest"
	EventTypeTransfer        = "vesting_transfer"

	AttributeKeyValidator   = "validator"
	AttributeKeyVestingType = "vesting_type"

	AttributeKeySrcValidator     = "source_validator"
	AttributeKeyDstValidator     = "destination_validator"
	AttributeKeyDelegator        = "delegator"
	AttributeKeyDelegableAddress = "delegable_address"
	AttributeKeyCompletionTime   = "completion_time"
	AttributeKeyRecipient  = "recipient"
	AttributeKeySender     = "sender"
	AttributeKeyWithdrawn  = "withdrawn"
	AttributeValueCategory = ModuleName
)
