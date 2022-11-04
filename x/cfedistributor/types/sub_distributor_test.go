package types

import (
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"
)

func TestCheckAccountType(t *testing.T) {
	tests := []struct {
		name    string
		account Account
		want    bool
	}{
		{"Check base account", Account{Id: "c4e1avc7vz3khvlf6fgd3a2exnaqnhhk0sxzzgxc4n", Type: BASE_ACCOUNT}, true},
		{"Check module account", Account{Id: "sample", Type: MODULE_ACCOUNT}, true},
		{"Check internal account", Account{Id: "sample", Type: INTERNAL_ACCOUNT}, true},
		{"Check main account", Account{Id: "sample", Type: MAIN}, true},
		{"Check wrong account", Account{Id: "test", Type: "test"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.account.Validate(); got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckPercentShareSumIsGTEThen100(t *testing.T) {

	shareEqual30 := DestinationShare{
		Name:        "1",
		Share:       sdk.MustNewDecFromStr("30"),
		Destination: Account{},
	}

	shareEqual31 := DestinationShare{
		Name:        "2",
		Share:       sdk.MustNewDecFromStr("31"),
		Destination: Account{},
	}

	shareEqual50 := DestinationShare{
		Name:        "3",
		Share:       sdk.MustNewDecFromStr("50"),
		Destination: Account{},
	}

	shareEqual19 := DestinationShare{
		Name:        "4",
		Share:       sdk.MustNewDecFromStr("19"),
		Destination: Account{},
	}

	shareEqual20 := DestinationShare{
		Name:        "5",
		Share:       sdk.MustNewDecFromStr("20"),
		Destination: Account{},
	}

	shareEqualMinus20 := DestinationShare{
		Name:        "5",
		Share:       sdk.MustNewDecFromStr("-20"),
		Destination: Account{},
	}

	burnShare := sdk.NewDec(50)

	var sharesEqual30 []*DestinationShare
	sharesEqual30 = append(sharesEqual30, &shareEqual30)

	var sharesEqual50 []*DestinationShare
	sharesEqual50 = append(sharesEqual50, &shareEqual30)
	sharesEqual50 = append(sharesEqual50, &shareEqual50)

	var sharesEqual81 []*DestinationShare
	sharesEqual81 = append(sharesEqual81, &shareEqual30)
	sharesEqual81 = append(sharesEqual81, &shareEqual50)

	var sharesEqual100 []*DestinationShare
	sharesEqual100 = append(sharesEqual100, &shareEqual31)
	sharesEqual100 = append(sharesEqual100, &shareEqual50)
	sharesEqual100 = append(sharesEqual100, &shareEqual19)

	var sharesEqual101 []*DestinationShare
	sharesEqual101 = append(sharesEqual101, &shareEqual31)
	sharesEqual101 = append(sharesEqual101, &shareEqual50)
	sharesEqual101 = append(sharesEqual101, &shareEqual20)

	var sharesEqualMinus10 []*DestinationShare
	sharesEqualMinus10 = append(sharesEqual101, &shareEqualMinus20)
	sharesEqualMinus10 = append(sharesEqual101, &shareEqualMinus20)
	sharesEqualMinus10 = append(sharesEqual101, &shareEqual30)

	tests := []struct {
		name        string
		destination Destinations
		want        bool
	}{

		{"Share equal 30", Destinations{PrimaryShare: Account{}, Shares: sharesEqual30, BurnShare: sdk.ZeroDec()}, false},
		{"Share equal 80 with burn", Destinations{PrimaryShare: Account{}, Shares: sharesEqual30, BurnShare: burnShare}, false},
		{"Share equal 50", Destinations{PrimaryShare: Account{}, Shares: sharesEqual50, BurnShare: sdk.ZeroDec()}, false},
		{"Share equal 100 with burn", Destinations{PrimaryShare: Account{}, Shares: sharesEqual50, BurnShare: burnShare}, true},
		{"Share equal 81", Destinations{PrimaryShare: Account{}, Shares: sharesEqual81, BurnShare: sdk.ZeroDec()}, false},
		{"Share equal 100", Destinations{PrimaryShare: Account{}, Shares: sharesEqual100, BurnShare: sdk.ZeroDec()}, true},
		{"Share equal 101", Destinations{PrimaryShare: Account{}, Shares: sharesEqual101, BurnShare: sdk.ZeroDec()}, true},
		{"Share equal -10", Destinations{PrimaryShare: Account{}, Shares: sharesEqualMinus10, BurnShare: sdk.ZeroDec()}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.destination.CheckPercentShareSumIsBetween0And100(); got != tt.want {
				t.Errorf("CheckPercentShareSumIsBetween0And100() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateOrderOfMainSubDistributors(t *testing.T) {
	twoMainTypesWithinOneSubdistributor := []SubDistributor{
		createSubDistributor(BASE_ACCOUNT, MAIN, MAIN, ""),
	}

	var zeroSubDistributors []SubDistributor

	onlyOneMainSubdistributor := []SubDistributor{
		CreateSubDistributor(MAIN_SOURCE),
		createSubDistributor(BASE_ACCOUNT, MAIN, BASE_ACCOUNT, ""),
	}
	destinationMainAtTheEnd := []SubDistributor{
		CreateSubDistributor(MAIN_SOURCE),
		CreateSubDistributor(MAIN_DESTINATION),
	}

	sourceMainAtTheEnd := []SubDistributor{
		CreateSubDistributor(MAIN_SOURCE),
		CreateSubDistributor(MAIN_DESTINATION),
		CreateSubDistributor(MAIN_SOURCE),
	}

	destinationShareMainAtTheEnd := []SubDistributor{
		CreateSubDistributor(MAIN_SOURCE),
		CreateSubDistributor(MAIN_DESTINATION_SHARE),
	}

	destinationShareSourceMainAtTheEnd := []SubDistributor{
		CreateSubDistributor(MAIN_DESTINATION),
		CreateSubDistributor(MAIN_DESTINATION_SHARE),
		CreateSubDistributor(MAIN_DESTINATION_SHARE),
		CreateSubDistributor(MAIN_SOURCE),
	}

	tests := []struct {
		name            string
		subDistributors []SubDistributor
		wantError       bool
	}{
		{"only one main subdistributor", onlyOneMainSubdistributor, false},
		{"two main types within one subdistributor", twoMainTypesWithinOneSubdistributor, true},
		{"zero sub distributors", zeroSubDistributors, true},
		{"wrong order destination main at the end", destinationMainAtTheEnd, true},
		{"correct order source main at the end", sourceMainAtTheEnd, false},
		{"wrong order destination main share at the end", destinationShareMainAtTheEnd, true},
		{"correct order destination main share, source main at the end", destinationShareSourceMainAtTheEnd, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := ValidateSubDistributors(tt.subDistributors)
			if tt.wantError == true && err == nil {
				t.Errorf("ValidateOrderOfSubDistributors() wanted error got nil")
			} else if tt.wantError == false && err != nil {
				t.Errorf("ValidateOrderOfSubDistributors() error: %v", err.Error())
			}
		})
	}
}

func TestValidateOrderOfInternalSubDistributors(t *testing.T) {
	onlyOneInternalSubdistributor := []SubDistributor{
		CreateSubDistributor(INTERNAL_SOURCE),
	}
	twoInternalTypesWithinOneSubdistributor := []SubDistributor{
		createSubDistributor(BASE_ACCOUNT, INTERNAL_ACCOUNT, INTERNAL_ACCOUNT, "custom_id"),
	}

	var destinationAtTheEnd []SubDistributor

	destinationInternalAtTheEnd := []SubDistributor{
		CreateSubDistributor(INTERNAL_SOURCE),
		CreateSubDistributor(INTERNAL_DESTINATION),
	}

	sourceInternalAtTheEnd := []SubDistributor{
		CreateSubDistributor(INTERNAL_SOURCE),
		CreateSubDistributor(INTERNAL_DESTINATION),
		CreateSubDistributor(INTERNAL_SOURCE),
		CreateSubDistributor(MAIN_SOURCE),
	}

	destinationInternalShareAtTheEnd := []SubDistributor{
		createSubDistributor(BASE_ACCOUNT, INTERNAL_ACCOUNT, BASE_ACCOUNT, "custom_id"),
		CreateSubDistributor(INTERNAL_DESTINATION_SHARE),
		CreateSubDistributor(MAIN_SOURCE),
	}

	destinationShareSourceInternalAtTheEndNoSource := []SubDistributor{
		CreateSubDistributor(INTERNAL_SOURCE),
		CreateSubDistributor(INTERNAL_DESTINATION_SHARE),
		CreateSubDistributor(INTERNAL_DESTINATION_SHARE),
		CreateSubDistributor(INTERNAL_SOURCE),
	}

	destinationShareSourceInternalAtTheEndSource := []SubDistributor{
		CreateSubDistributor(INTERNAL_SOURCE),
		CreateSubDistributor(INTERNAL_DESTINATION_SHARE),
		CreateSubDistributor(INTERNAL_DESTINATION_SHARE),
		CreateSubDistributor(INTERNAL_DESTINATION_SHARE),
		CreateSubDistributor(INTERNAL_SOURCE),
		CreateSubDistributor(MAIN_SOURCE),
	}

	tests := []struct {
		name            string
		subDistributors []SubDistributor
		wantError       bool
	}{
		{"only one internal subdistributor", onlyOneInternalSubdistributor, true},
		{"two internal types within one subdistributor", twoInternalTypesWithinOneSubdistributor, true},
		{"wrong order destination main at the end", destinationAtTheEnd, true},
		{"wrong order destination internal at the end", destinationInternalAtTheEnd, true},
		{"correct order source main at the end", sourceInternalAtTheEnd, false},
		{"wrong order destination internal share at the end", destinationInternalShareAtTheEnd, true},
		{"correct order destination internal share, source internal at the end, source main at the end", destinationShareSourceInternalAtTheEndSource, false},
		{"correct order destination internal share, source internal at the end but no main source", destinationShareSourceInternalAtTheEndNoSource, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSubDistributors(tt.subDistributors)
			if tt.wantError == true && err == nil {
				t.Errorf("ValidateOrderOfSubDistributors() wanted error got nil")
			} else if tt.wantError == false && err != nil {
				t.Errorf("ValidateOrderOfSubDistributors() error: %v", err.Error())
			}
		})
	}
}

func TestValidateUniquenessOfNames(t *testing.T) {
	twoSubdistributorsUniqueNames := []SubDistributor{
		CreateSubDistributor(MAIN_DESTINATION),
		CreateSubDistributor(MAIN_DESTINATION_SHARE),
		CreateSubDistributor(MAIN_SOURCE),
	}

	subDistributorMainDestination := CreateSubDistributor(MAIN_DESTINATION)
	subDistributorMainSource := CreateSubDistributor(MAIN_SOURCE)
	subDistributorMainDestinationShare := CreateSubDistributor(MAIN_DESTINATION_SHARE)
	twoSubdistributorsSameNames := []SubDistributor{
		subDistributorMainDestination,
		subDistributorMainDestination,
		subDistributorMainSource,
	}

	twoSubdistributorsSameShareNames := []SubDistributor{
		subDistributorMainDestinationShare,
		subDistributorMainDestinationShare,
		subDistributorMainSource,
	}

	tests := []struct {
		name            string
		subDistributors []SubDistributor
		wantError       bool
	}{
		{"two subdistributors have unique names", twoSubdistributorsUniqueNames, false},
		{"two subdistributors have same names", twoSubdistributorsSameNames, true},
		{"two subdistributors have same share names", twoSubdistributorsSameShareNames, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSubDistributors(tt.subDistributors)
			if tt.wantError == true && err == nil {
				t.Errorf("ValidateSubDistributors() wanted error got nil")
			} else if tt.wantError == false && err != nil {
				t.Errorf("ValidateSubDistributors() error: %v", err.Error())
			}
		})
	}
}

const (
	MAIN_SOURCE                = "MAIN_SOURCE"
	MAIN_DESTINATION           = "MAIN_DESTINATION"
	MAIN_DESTINATION_SHARE     = "MAIN_DESTINATION_SHARE"
	INTERNAL_SOURCE            = "INTERNAL_SOURCE"
	INTERNAL_DESTINATION       = "INTERNAL_DESTINATION"
	INTERNAL_DESTINATION_SHARE = "INTERNAL_DESTINATION_SHARE"
)

func CreateSubDistributor(accType string) SubDistributor {
	switch accType {
	case MAIN_SOURCE:
		return createSubDistributor(BASE_ACCOUNT, MAIN, BASE_ACCOUNT, "custom_id")
	case MAIN_DESTINATION:
		return createSubDistributor(MAIN, BASE_ACCOUNT, BASE_ACCOUNT, "custom_id")
	case MAIN_DESTINATION_SHARE:
		return createSubDistributor(BASE_ACCOUNT, BASE_ACCOUNT, MAIN, "custom_id")
	case INTERNAL_SOURCE:
		return createSubDistributor(BASE_ACCOUNT, INTERNAL_ACCOUNT, BASE_ACCOUNT, "custom_id")
	case INTERNAL_DESTINATION:
		return createSubDistributor(INTERNAL_ACCOUNT, BASE_ACCOUNT, BASE_ACCOUNT, "custom_id")
	case INTERNAL_DESTINATION_SHARE:
		return createSubDistributor(BASE_ACCOUNT, BASE_ACCOUNT, INTERNAL_ACCOUNT, "custom_id")
	}
	return SubDistributor{}
}

func createSubDistributor(
	destinationType string,
	sourceType string,
	destinationShareType string,
	Id string,
) SubDistributor {
	return SubDistributor{
		Name: helpers.RandStringOfLength(10),
		Destinations: Destinations{
			PrimaryShare: Account{
				Id:   Id + getIdSuffix("mainDst", destinationType),
				Type: destinationType,
			},
			BurnShare: sdk.MustNewDecFromStr("0"),
			Shares: []*DestinationShare{
				{
					Name: helpers.RandStringOfLength(10),
					Destination: Account{
						Id:   Id + getIdSuffix("shareDst", destinationShareType),
						Type: destinationShareType,
					},
					Share: sdk.MustNewDecFromStr("0"),
				},
			},
		},
		Sources: []*Account{
			{
				Id:   Id + getIdSuffix("src", sourceType),
				Type: sourceType,
			},
		},
	}
}

func getIdSuffix(suffix string, accType string) string {
	if accType == INTERNAL_ACCOUNT || accType == MAIN {
		return "-" + accType
	}
	return "-" + suffix
}
