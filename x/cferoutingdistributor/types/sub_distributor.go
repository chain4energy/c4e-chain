package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s SubDistributor) Validate() error {

	if CheckPercentShareSumIsLowerThen100(s.Destination) {
		return fmt.Errorf("share sum is greater or equal 100")
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
