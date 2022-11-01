package types

import (
	"fmt"
	"strings"

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

func ValidateSubDistributors(subDistributors []SubDistributor) error {
	lastOccurrence := make(map[string]string)
	lastOccurrenceIndex := make(map[string]int)
	subDistributorNameOccured := make(map[string]bool)
	shareNameOccured := make(map[string]bool)

	for i := 0; i < len(subDistributors); i++ {
		subDistributorName := subDistributors[i].Name
		err := validateUniquenessOfNames(subDistributorName, &subDistributorNameOccured)
		if err != nil {
			return err
		}

		for j := 0; j < len(subDistributors[i].Sources); j++ {
			sourceType := subDistributors[i].Sources[j].Type
			if sourceType == INTERNAL_ACCOUNT || sourceType == MAIN {
				id := setId(sourceType, subDistributors[i].Sources[j].Id)
				if lastOccurrenceIndex[id] == i && lastOccurrence[id] != "" {
					return fmt.Errorf("same %s account cannot occur twice within one subdistributor, subdistributor name: %s",
						strings.ToLower(sourceType), subDistributorName)
				}
				lastOccurrence[id] = SOURCE
				lastOccurrenceIndex[id] = i
			}
		}

		destinationType := subDistributors[i].Destination.Account.Type
		if destinationType == INTERNAL_ACCOUNT || destinationType == MAIN {
			id := setId(destinationType, subDistributors[i].Destination.Account.Id)
			lastOccurrence[id] = DESTINATION
			lastOccurrenceIndex[id] = i
		}

		for j := 0; j < len(subDistributors[i].Destination.Share); j++ {
			shareName := subDistributors[i].Destination.Share[j].Name
			if err = validateUniquenessOfNames(shareName, &shareNameOccured); err != nil {
				return err
			}

			destinationShareType := subDistributors[i].Destination.Share[j].Account.Type
			if destinationShareType == INTERNAL_ACCOUNT || destinationShareType == MAIN {
				id := setId(destinationShareType, subDistributors[i].Destination.Share[j].Account.Id)
				if lastOccurrenceIndex[id] == i && lastOccurrence[id] != "" {
					return fmt.Errorf("same %s account cannot occur twice within one subdistributor, subdistributor name: %s",
						strings.ToLower(destinationShareType), subDistributorName)
				}
				lastOccurrence[id] = DESTINATION
				lastOccurrenceIndex[id] = i
			}
		}

	}
	if lastOccurrence[MAIN] == "" {
		return fmt.Errorf("there must be at least one subdistributor with the source main type")
	}
	for accountId, _ := range lastOccurrence {
		if lastOccurrence[accountId] != SOURCE {
			return fmt.Errorf("wrong order of subdistributors, after each occurrence of a subdistributor with the " +
				"destination of internal or main account type there must be exactly one occurrence of a subdistributor with the " +
				"source of internal account type, account id: " + accountId)
		}
	}

	return nil
}

func setId(accType string, id string) string {
	if accType == INTERNAL_ACCOUNT {
		return id
	}
	return MAIN
}

func validateUniquenessOfNames(subDistributorName string, nameOccured *map[string]bool) error {
	if (*nameOccured)[subDistributorName] {
		return fmt.Errorf("subdistributor names must be unique, subdistributor name: " + subDistributorName)
	}
	(*nameOccured)[subDistributorName] = true
	return nil
}
