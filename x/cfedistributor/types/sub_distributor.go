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
		return BURN
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
	// Account Types
	INTERNAL_ACCOUNT = "INTERNAL_ACCOUNT"
	MODULE_ACCOUNT   = "MODULE_ACCOUNT"
	MAIN             = "MAIN"
	BASE_ACCOUNT     = "BASE_ACCOUNT"
	// Other consts
	UNKNOWN_ACCOUNT = "Unknown"
	BURN            = "BURN"
	DESTINATION     = "DESTINATION"
	SOURCE          = "SOURCE"
)

func CheckAccountType(account Account) bool {
	switch account.Type {
	case INTERNAL_ACCOUNT, MODULE_ACCOUNT, MAIN, BASE_ACCOUNT:
		return true
	default:
		return false
	}
}

func ValidateOrderOfSubDistributors(subDistributors *[]SubDistributor) error {
	err := validateOrderOfMainSubDistributors(*subDistributors)
	if err != nil {
		return err
	}
	err = validateOrderOfInternalSubDistributors(*subDistributors)
	return err
}

func validateOrderOfMainSubDistributors(subDistributors []SubDistributor) error {
	lastOccurrence := ""
	lastOccurrenceIndex := -1

	for i := 0; i < len(subDistributors); i++ {
		if subDistributors[i].Destination.Account.Type == MAIN {
			lastOccurrence = DESTINATION
			lastOccurrenceIndex = i
		}
		for j := 0; j < len(subDistributors[i].Destination.Share); j++ {
			if subDistributors[i].Destination.Share[j].Account.Type == MAIN {
				if lastOccurrenceIndex == i {
					return fmt.Errorf("main type cannot occur twice within one subdistributor, " +
						"subdistributor name: " + subDistributors[i].Name)
				}
				lastOccurrence = DESTINATION
				lastOccurrenceIndex = i
			}
		}

		for j := 0; j < len(subDistributors[i].Sources); j++ {
			if subDistributors[i].Sources[j].Type == MAIN {
				if lastOccurrenceIndex == i {
					return fmt.Errorf("main type cannot occur twice within one subdistributor, " +
						"subdistributor name: " + subDistributors[i].Name)
				}
				lastOccurrence = SOURCE
			}
		}
	}

	if lastOccurrence == "" {
		return fmt.Errorf("there must be at least one subdistributor with the source main type")
	}
	if lastOccurrence != SOURCE {
		return fmt.Errorf("wrong order of subdistributors, after each occurrence of a subdistributor with the " +
			"destination main type there must be exactly one occurrence of a subdistributor with the source main type")
	}

	return nil
}

func validateOrderOfInternalSubDistributors(subDistributors []SubDistributor) error {
	lastOccurrence := make(map[string]string)
	lastOccurrenceIndex := make(map[string]int)

	for i := 0; i < len(subDistributors); i++ {
		if subDistributors[i].Destination.Account.Type == INTERNAL_ACCOUNT {
			id := subDistributors[i].Destination.Account.Id
			lastOccurrence[id] = DESTINATION
			lastOccurrenceIndex[id] = i
		}

		for j := 0; j < len(subDistributors[i].Destination.Share); j++ {
			if subDistributors[i].Destination.Share[j].Account.Type == INTERNAL_ACCOUNT {
				id := subDistributors[i].Destination.Share[j].Account.Id
				if lastOccurrenceIndex[id] == i && lastOccurrence[id] != "" {
					return fmt.Errorf("same internal account cannot occur twice within one subdistributor, "+
						"subdistributor name: %s", subDistributors[i].Name)
				}
				lastOccurrence[id] = DESTINATION
			}
		}

		for j := 0; j < len(subDistributors[i].Sources); j++ {
			if subDistributors[i].Sources[j].Type == INTERNAL_ACCOUNT {
				id := subDistributors[i].Sources[j].Id
				if lastOccurrenceIndex[id] == i && lastOccurrence[id] != "" {
					return fmt.Errorf("same internal account cannot occur twice within one subdistributor, " +
						"subdistributor name: " + subDistributors[i].Name)
				}
				lastOccurrence[id] = SOURCE
			}
		}
	}

	for internalAccountId, _ := range lastOccurrence {
		if lastOccurrence[internalAccountId] != SOURCE {
			return fmt.Errorf("wrong order of subdistributors, after each occurrence of a subdistributor with the " +
				"destination of internal account type there must be exactly one occurrence of a subdistributor with the " +
				"source of internal account type, internal account id: " + internalAccountId)
		}
	}

	return nil
}
