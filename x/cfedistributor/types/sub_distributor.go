package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s SubDistributor) Validate() error {

	if CheckPercentShareSumIsGTEThen100(s.Destination) {
		return fmt.Errorf("share sum is greater or equal 100")
	}

	for _, source := range s.Sources {
		if !CheckAccountType(*source) {
			return fmt.Errorf("the source account is of the wrong type: " + source.String())
		}
	}

	for _, share := range s.Destination.Share {
		if !CheckAccountType(share.Account) {
			return fmt.Errorf("the destination account is of the wrong type: " + share.Account.String())
		}
	}

	return nil
}

func (s State) StateIdString() string {

	if s.Burn {
		return "BURN"
	} else if s.Account != nil && s.Account.Type == MAIN {
		return MAIN
	} else if s.Account != nil {
		return s.Account.Type + "-" + s.Account.Id
	} else {
		return UNKNOWN_ACCOUNT
	}
}

func CheckPercentShareSumIsGTEThen100(destination Destination) bool {
	shares := destination.Share
	percentShareSum := sdk.MustNewDecFromStr("0")
	for _, share := range shares {
		percentShareSum = percentShareSum.Add(share.Percent)
	}

	if destination.BurnShare != nil {

		percentShareSum = percentShareSum.Add(destination.BurnShare.Percent)
	}

	return percentShareSum.GTE(sdk.MustNewDecFromStr("100"))
}

const (
	INTERNAL_ACCOUNT string = "INTERNAL_ACCOUNT"
	MODULE_ACCOUNT          = "MODULE_ACCOUNT"
	MAIN                    = "MAIN"
	BASE_ACCOUNT            = "BASE_ACCOUNT"
	UNKNOWN_ACCOUNT         = "Unknown"
)

func CheckAccountType(account Account) bool {
	switch account.Type {
	case INTERNAL_ACCOUNT, MODULE_ACCOUNT, MAIN, BASE_ACCOUNT:
		return true
	default:
		return false
	}
}
