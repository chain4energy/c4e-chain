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
			return fmt.Errorf("source account is undefined: " + source.String())
		}
	}

	for _, share := range s.Destination.Share {
		if !CheckAccountType(share.Account) {
			return fmt.Errorf("destination account is undefined: " + share.Account.String())
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

func CheckAccountType(account Account) bool {
	accountExist := false
	if &account.Address != nil && account.Address != "" {
		accountExist = true
		if &account.ModuleName != nil && account.ModuleName != "" {
			return false
		}

		if &account.InternalName != nil && account.InternalName != "" {
			return false
		}
	}

	if &account.ModuleName != nil && account.ModuleName != "" {
		accountExist = true
		if &account.Address != nil && account.Address != "" {
			return false
		}

		if &account.InternalName != nil && account.InternalName != "" {
			return false
		}
	}

	if &account.InternalName != nil && account.InternalName != "" {
		accountExist = true
		if &account.ModuleName != nil && account.ModuleName != "" {
			return false
		}

		if &account.Address != nil && account.Address != "" {
			return false
		}
	}
	return accountExist
}
