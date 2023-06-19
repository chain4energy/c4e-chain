package v2

var (
	BurnStateKey   = "burn_state_key"
	StateKeyPrefix = []byte{0x04}
)

func (state State) GetStateKey() string {
	if state.Account != nil && state.Account.Id != "" && state.Account.Type != "" {
		return state.Account.GetAccountKey()
	} else {
		return BurnStateKey
	}
}

func (account Account) GetAccountKey() string {
	return account.Type + "-" + account.Id
}
