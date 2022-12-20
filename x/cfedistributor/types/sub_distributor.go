package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const maxShare = 1
const primaryShareNameSuffix = "_primary"

func (s SubDistributor) Validate() error {
	if s.Name == "" {
		return fmt.Errorf("subdistributor name cannot be empty")
	}
	if &s.Destinations == nil {
		return fmt.Errorf("subdistributor destinaions cannot be nil")
	}
	if err := s.Destinations.Validate(s.GetPrimaryShareName()); err != nil {
		return err
	}
	if err := s.Destinations.PrimaryShare.Validate(); err != nil {
		return err
	}
	if len(s.Sources) < 1 {
		return fmt.Errorf("subdistributor must have at least one source")
	}
	for _, source := range s.Sources {
		if source == nil {
			return fmt.Errorf("source cannot be nil")
		}
		if err := source.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (s SubDistributor) GetPrimaryShareName() string {
	return s.Name + primaryShareNameSuffix
}

func (destinations Destinations) Validate(primaryShareName string) error {
	if destinations.BurnShare.IsNil() {
		return fmt.Errorf("burn share cannot be nil")
	}
	if destinations.BurnShare.GTE(sdk.NewDec(maxShare)) || destinations.BurnShare.IsNegative() {
		return fmt.Errorf("burn share must be between 0 and 1")
	}

	for _, destinationShare := range destinations.Shares {
		if err := destinationShare.Validate(primaryShareName); err != nil {
			return err
		}
	}

	if err := destinations.CheckPercentShareSumIsBetween0And1(); err != nil {
		return err
	}
	return nil
}

func (destinations Destinations) CheckPercentShareSumIsBetween0And1() error {
	shareSum := destinations.BurnShare

	for _, share := range destinations.Shares {
		shareSum = shareSum.Add(share.Share)
	}

	if shareSum.GTE(sdk.NewDec(maxShare)) || shareSum.IsNegative() {
		return fmt.Errorf("share sum must be between 0 and 1")
	}

	return nil
}

func (destinationShare *DestinationShare) Validate(primaryShareName string) error {
	if destinationShare == nil {
		return fmt.Errorf("destination share cannot be nil")
	}
	if destinationShare.Share.IsNil() {
		return fmt.Errorf("share cannot be nil")
	}
	if destinationShare.Share.GTE(sdk.NewDec(maxShare)) || destinationShare.Share.IsNegative() {
		return fmt.Errorf("share must be between 0 and 1")
	}
	if err := destinationShare.Destination.Validate(); err != nil {
		return err
	}
	if destinationShare.Name == "" {
		return fmt.Errorf("destination share name cannot be empty")
	}
	if destinationShare.Name == primaryShareName {
		return fmt.Errorf("share name: %s is reserved for primary share", destinationShare.Name)
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

func (account Account) Validate() error {
	switch account.Type {
	case INTERNAL_ACCOUNT, MAIN, BASE_ACCOUNT:
		return nil
	case MODULE_ACCOUNT:
		if !accountExistInMacPerms(account.Id) {
			return fmt.Errorf("module account \"%s\" doesn't exist in maccPerms", account.Id)
		}
		return nil
	default:
		return fmt.Errorf("account \"%s\" is of the wrong type: %s", account.Id, account.Type)
	}
}

func (account Account) GetAccounteKey() string {
	return account.Type + "-" + account.Id
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
			if err := setOccurrence(lastOccurrence, lastOccurrenceIndex, subDistributorName, subDistributors[i].Sources[j], i, SOURCE); err != nil {
				return err
			}
		}

		if err := setOccurrence(lastOccurrence, lastOccurrenceIndex, subDistributorName, &subDistributors[i].Destinations.PrimaryShare, i, DESTINATION); err != nil {
			return err
		}

		if err = validateUniquenessOfNames(subDistributors[i].GetPrimaryShareName(), &shareNameOccured); err != nil {
			return err
		}

		for j := 0; j < len(subDistributors[i].Destinations.Shares); j++ {
			shareName := subDistributors[i].Destinations.Shares[j].Name
			if err = validateUniquenessOfNames(shareName, &shareNameOccured); err != nil {
				return err
			}

			if err := setOccurrence(lastOccurrence, lastOccurrenceIndex, subDistributorName, &subDistributors[i].Destinations.Shares[j].Destination, i, DESTINATION); err != nil {
				return err
			}
		}

	}

	if lastOccurrence[MAIN] == "" {
		return fmt.Errorf("there must be at least one subdistributor with the source main type")
	}
	for accountId := range lastOccurrence {
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
	return account.Type + "-" + account.Id
}

func isAccountPositionValidatable(accType string) bool {
	return accType == INTERNAL_ACCOUNT || accType == MAIN
}

func setOccurrence(lastOccurrence map[string]string, lastOccurrenceIndex map[string]int, subDistributorName string, account *Account, position int, occuranceType string) error {
	id := getId(account)
	currentPosition := position + 1
	if lastOccurrenceIndex[id] == currentPosition {
		return fmt.Errorf("same %s account cannot occur twice within one subdistributor, subdistributor name: %s",
			id, subDistributorName)
	}
	if isAccountPositionValidatable(account.Type) {
		lastOccurrence[id] = occuranceType
	}
	lastOccurrenceIndex[id] = currentPosition
	return nil
}

func validateUniquenessOfNames(subDistributorName string, nameOccured *map[string]bool) error {
	if (*nameOccured)[subDistributorName] {
		return fmt.Errorf("subdistributor names must be unique, subdistributor name: " + subDistributorName)
	}
	(*nameOccured)[subDistributorName] = true

	return nil
}

func accountExistInMacPerms(accountId string) bool {
	_, found := maccPerms[accountId]
	return found
}

//
//Name nie może być  "",
//Source musi być co najmniej jeden i żaden nie może być nullem
//Destinations ma jeden
//Shares może być puste ale nie moze mieć nuli
//Name nie pusty,
//Burn share i share między 0 i 1 osobno
// wszystkie inty i dec sprawdzać czy isNil
// mintery czy null w tablicy
