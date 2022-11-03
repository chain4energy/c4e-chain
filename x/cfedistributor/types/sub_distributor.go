package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const maxShareSum = 100

func (s SubDistributor) Validate() error {

	if s.Destination.CheckPercentShareSumIsBetween0And100() {
		return fmt.Errorf("share sum is greater or equal 100")
	}

	for _, source := range s.Sources {
		if !source.Validate() {
			return fmt.Errorf("the source account is of the wrong type: " + source.String())
		}
	}

	for _, share := range s.Destination.Share {
		if !share.Account.Validate() {
			return fmt.Errorf("the destination account is of the wrong type: " + share.Account.String())
		}
	}

	return nil
}

func (destination Destination) CheckPercentShareSumIsBetween0And100() bool {
	shares := destination.Share
	percentShareSum := sdk.ZeroDec()
	for _, share := range shares {
		percentShareSum = percentShareSum.Add(share.Percent)
	}

	if destination.BurnShare != nil {

		percentShareSum = percentShareSum.Add(destination.BurnShare.Percent)
	}

	return percentShareSum.GTE(sdk.NewDec(maxShareSum)) || percentShareSum.IsNegative()
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

func (account Account) Validate() bool {
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
			if err := setOccurence(lastOccurrence, lastOccurrenceIndex, subDistributorName, subDistributors[i].Sources[j], i, SOURCE); err != nil {
				return err
			}
		}

		if err := setOccurence(lastOccurrence, lastOccurrenceIndex, subDistributorName, &subDistributors[i].Destination.Account, i, DESTINATION); err != nil {
			return err
		}

		for j := 0; j < len(subDistributors[i].Destination.Share); j++ {
			shareName := subDistributors[i].Destination.Share[j].Name
			if err = validateUniquenessOfNames(shareName, &shareNameOccured); err != nil {
				return err
			}

			if err := setOccurence(lastOccurrence, lastOccurrenceIndex, subDistributorName, &subDistributors[i].Destination.Share[j].Account, i, DESTINATION); err != nil {
				return err
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

func getId(account *Account) string {
	if account.Type == MAIN {
		return MAIN
	}
	return account.Type +"-"+account.Id
}

func isAccountPositionValidatable(accType string) bool {
	return accType == INTERNAL_ACCOUNT || accType == MAIN
}

func setOccurence(lastOccurrence map[string]string, lastOccurrenceIndex map[string]int, subDistributorName string, account *Account, position int, occuranceType string) error {
	id := getId(account)
	if lastOccurrenceIndex[id] == position + 1 {
		return fmt.Errorf("same %s account cannot occur twice within one subdistributor, subdistributor name: %s",
			id, subDistributorName)
	}
	if isAccountPositionValidatable(account.Type) {
		lastOccurrence[id] = occuranceType
	}
	lastOccurrenceIndex[id] = position + 1
	return nil
}

func validateUniquenessOfNames(subDistributorName string, nameOccured *map[string]bool) error {
	if (*nameOccured)[subDistributorName] {
		return fmt.Errorf("subdistributor names must be unique, subdistributor name: " + subDistributorName)
	}
	(*nameOccured)[subDistributorName] = true

	return nil
}
