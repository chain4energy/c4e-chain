package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const maxShare = 1
const primaryShareNameSuffix = "_primary"

func (subdistributor SubDistributor) Validate() error {
	if subdistributor.Name == "" {
		return fmt.Errorf("subdistributor name cannot be empty")
	}
	if err := subdistributor.Destinations.Validate(subdistributor.GetPrimaryShareName()); err != nil {
		return fmt.Errorf("subdistributor %s destinations validation error: %w", subdistributor.Name, err)
	}
	if len(subdistributor.Sources) < 1 {
		return fmt.Errorf("subdistributor %s must have at least one source", subdistributor.Name)
	}
	for i, source := range subdistributor.Sources {
		if source == nil {
			return fmt.Errorf("subdistributor %s source on position %d cannot be nil", subdistributor.Name, i+1)
		}
		if err := source.Validate(); err != nil {
			return fmt.Errorf("subdistributor %s source with id \"%s\" validation error: %w", subdistributor.Name, source.Id, err)
		}
	}

	return nil
}

func (subdistributor SubDistributor) GetPrimaryShareName() string {
	return subdistributor.Name + primaryShareNameSuffix
}

func (destinations Destinations) Validate(primaryShareName string) error {
	if destinations.BurnShare.IsNil() {
		return fmt.Errorf("burn share cannot be nil")
	}
	if destinations.BurnShare.GTE(sdk.NewDec(maxShare)) || destinations.BurnShare.IsNegative() {
		return fmt.Errorf("burn share must be between 0 and 1")
	}

	for i, share := range destinations.Shares {
		if share == nil {
			return fmt.Errorf("destination share on position %d cannot be nil", i+1)
		}
		if err := share.Validate(primaryShareName); err != nil {
			return err
		}
	}
	if err := destinations.PrimaryShare.Validate(); err != nil {
		return fmt.Errorf("primary share validation error: %w", err)
	}
	if err := destinations.CheckIfSharesSumIsBetween0And1(); err != nil {
		return err
	}
	return nil
}

func (destinations Destinations) CheckIfSharesSumIsBetween0And1() error {
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
	if destinationShare.Name == "" {
		return fmt.Errorf("destination share name cannot be empty")
	}
	if destinationShare.Name == primaryShareName {
		return fmt.Errorf("destination share name: %s is reserved for primary share", destinationShare.Name)
	}
	if destinationShare.Share.IsNil() {
		return fmt.Errorf("destination share %s share cannot be nil", destinationShare.Name)
	}
	if destinationShare.Share.GTE(sdk.NewDec(maxShare)) || destinationShare.Share.IsNegative() {
		return fmt.Errorf("destination share %s share must be between 0 and 1", destinationShare.Name)
	}
	if err := destinationShare.Destination.Validate(); err != nil {
		return fmt.Errorf("destination share %s destination account validation error: %w", destinationShare.Name, err)
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
	case MAIN:
		return nil
	case INTERNAL_ACCOUNT:
		if account.Id == "" {
			return fmt.Errorf("internal account id cannot be empty")
		}
	case BASE_ACCOUNT:
		if _, err := sdk.AccAddressFromBech32(account.Id); err != nil {
			return fmt.Errorf("base account id \"%s\" is not a valid bech32 address: %w", account.Id, err)
		}
	case MODULE_ACCOUNT:
		if !accountExistInMacPerms(account.Id) {
			return fmt.Errorf("module account \"%s\" doesn't exist in maccPerms", account.Id)
		}
	default:
		return fmt.Errorf("account \"%s\" is of the wrong type: %s", account.Id, account.Type)
	}
	return nil
}

func (account Account) GetAccounteKey() string {
	return account.Type + "-" + account.Id
}

func ValidateSubDistributors(subDistributors []SubDistributor) error {
	lastOccurrence := make(map[string]string)
	lastOccurrenceIndex := make(map[string]int)
	subDistributorNameOccurred := make(map[string]bool)
	shareNameOccurred := make(map[string]bool)

	for i := 0; i < len(subDistributors); i++ {
		subDistributorName := subDistributors[i].Name
		if err := validateUniquenessOfNames(subDistributorName, &subDistributorNameOccurred); err != nil {
			return err
		}
		if err := validateSources(subDistributors[i].Sources, i, lastOccurrence, lastOccurrenceIndex, subDistributorName, SOURCE); err != nil {
			return err
		}

		if err := setOccurrence(lastOccurrence, lastOccurrenceIndex, subDistributorName, &subDistributors[i].Destinations.PrimaryShare, i, DESTINATION); err != nil {
			return err
		}

		if err := validateUniquenessOfNames(subDistributors[i].GetPrimaryShareName(), &shareNameOccurred); err != nil {
			return err
		}
		if err := validateDestinationsShares(subDistributors[i].Destinations.Shares, i, lastOccurrence, lastOccurrenceIndex, shareNameOccurred, subDistributorName, DESTINATION); err != nil {
			return err
		}
	}

	return validateLastOccurrence(lastOccurrence)
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

func setOccurrence(lastOccurrence map[string]string, lastOccurrenceIndex map[string]int, subDistributorName string, account *Account, position int, accountType string) error {
	id := getId(account)
	currentPosition := position + 1
	if lastOccurrenceIndex[id] == currentPosition {
		return fmt.Errorf("same %s account cannot occur twice within one subdistributor, subdistributor name: %s",
			id, subDistributorName)
	}
	if isAccountPositionValidatable(account.Type) {
		lastOccurrence[id] = accountType
	}
	lastOccurrenceIndex[id] = currentPosition
	return nil
}

func validateUniquenessOfNames(subDistributorName string, nameOccurred *map[string]bool) error {
	if (*nameOccurred)[subDistributorName] {
		return fmt.Errorf("subdistributor names must be unique, subdistributor name: %s", subDistributorName)
	}
	(*nameOccurred)[subDistributorName] = true

	return nil
}

func accountExistInMacPerms(accountId string) bool {
	_, found := maccPerms[accountId]
	return found
}

func validateSources(accounts []*Account, subDistributorIndex int, lastOccurrence map[string]string, lastOccurrenceIndex map[string]int, subDistributorName string, accountType string) error {
	for _, source := range accounts {
		if err := setOccurrence(lastOccurrence, lastOccurrenceIndex, subDistributorName, source, subDistributorIndex, accountType); err != nil {
			return err
		}
	}

	return nil
}
func validateDestinationsShares(shares []*DestinationShare, subDistributorIndex int, lastOccurrence map[string]string, lastOccurrenceIndex map[string]int, shareNameOccurred map[string]bool, subDistributorName string, accountType string) error {
	for _, share := range shares {
		if err := validateUniquenessOfNames(share.Name, &shareNameOccurred); err != nil {
			return err
		}
		if err := setOccurrence(lastOccurrence, lastOccurrenceIndex, subDistributorName, &share.Destination, subDistributorIndex, accountType); err != nil {
			return err
		}
	}

	return nil
}

func validateLastOccurrence(lastOccurrence map[string]string) error {
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
