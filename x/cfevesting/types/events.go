package types

// staking module event types
const (
	// EventTypeCompleteUnbonding    = "complete_unbonding"
	// EventTypeCompleteRedelegation = "complete_redelegation"
	// EventTypeCreateValidator      = "create_validator"
	// EventTypeEditValidator        = "edit_validator"
	EventTypeDelegate        = "vesting_delegate"
	EventTypeUnbond          = "vesting_unbond"
	EventTypeRedelegate      = "vesting_redelegate"
	EventTypeWithdrawRewards = "vesting_withdraw_rewards"
	EventTypeWithdrawVesting = "vesting_withdraw_vesting"
	EventTypeVest            = "vesting_vest"

	AttributeKeyValidator   = "validator"
	AttributeKeyVestingType = "vesting_type"

	// AttributeKeyCommissionRate    = "commission_rate"
	// AttributeKeyMinSelfDelegation = "min_self_delegation"
	AttributeKeySrcValidator     = "source_validator"
	AttributeKeyDstValidator     = "destination_validator"
	AttributeKeyDelegator        = "delegator"
	AttributeKeyDelegableAddress = "delegable_address"
	AttributeKeyCompletionTime   = "completion_time"
	// AttributeKeyNewShares         = "new_shares"
	AttributeValueCategory = ModuleName
)
