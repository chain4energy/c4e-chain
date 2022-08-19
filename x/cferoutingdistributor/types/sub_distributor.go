package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s SubDistributor) Validate() error {

	if CheckPercentShareSumIsLowerThen100(s.Destination) {
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

func CheckPercentShareSumIsLowerThen100(destination Destination) bool {
	shares := destination.Share
	percentShareSum := sdk.MustNewDecFromStr("0")
	for _, share := range shares {
		percentShareSum = percentShareSum.Add(share.Percent)
	}

	percentShareSum.Add(destination.BurnShare.Percent)

	if percentShareSum.GTE(sdk.MustNewDecFromStr("100")) {
		return true
	} else {
		return false
	}
}

const (
	INTERNAL_ACCOUNT string = "INTERNAL_ACCOUNT"
	MODULE_ACCOUNT          = "MODULE_ACCOUNT"
	MAIN                    = "MAIN"
	BASE_ACCOUNT            = "BASE_ACCOUNT"
	UNKNOWN_ACCOUNT         = "Unknown"
)

func CheckAccountType(account Account) bool {
	if account.Type == MAIN {
		return true
	} else if account.Type == MODULE_ACCOUNT {
		return true
	} else if account.Type == INTERNAL_ACCOUNT {
		return true
	} else if account.Type == BASE_ACCOUNT {
		return true
	} else {
		return false
	}
}
