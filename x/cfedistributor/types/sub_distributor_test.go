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

	shareEqual30 := Share{
		Name:    "1",
		Percent: sdk.MustNewDecFromStr("30"),
		Account: Account{},
	}

	shareEqual31 := Share{
		Name:    "2",
		Percent: sdk.MustNewDecFromStr("31"),
		Account: Account{},
	}

	shareEqual50 := Share{
		Name:    "3",
		Percent: sdk.MustNewDecFromStr("50"),
		Account: Account{},
	}

	shareEqual19 := Share{
		Name:    "4",
		Percent: sdk.MustNewDecFromStr("19"),
		Account: Account{},
	}

	shareEqual20 := Share{
		Name:    "5",
		Percent: sdk.MustNewDecFromStr("20"),
		Account: Account{},
	}

	shareEqualMinus20 := Share{
		Name:    "5",
		Percent: sdk.MustNewDecFromStr("-20"),
		Account: Account{},
	}

	burnShare := BurnShare{Percent: sdk.MustNewDecFromStr("50")}

	var sharesEqual30 []*Share
	sharesEqual30 = append(sharesEqual30, &shareEqual30)

	var sharesEqual50 []*Share
	sharesEqual50 = append(sharesEqual50, &shareEqual30)
	sharesEqual50 = append(sharesEqual50, &shareEqual50)

	var sharesEqual81 []*Share
	sharesEqual81 = append(sharesEqual81, &shareEqual30)
	sharesEqual81 = append(sharesEqual81, &shareEqual50)

	var sharesEqual100 []*Share
	sharesEqual100 = append(sharesEqual100, &shareEqual31)
	sharesEqual100 = append(sharesEqual100, &shareEqual50)
	sharesEqual100 = append(sharesEqual100, &shareEqual19)

	var sharesEqual101 []*Share
	sharesEqual101 = append(sharesEqual101, &shareEqual31)
	sharesEqual101 = append(sharesEqual101, &shareEqual50)
	sharesEqual101 = append(sharesEqual101, &shareEqual20)

	var sharesEqualMinus10 []*Share
	sharesEqualMinus10 = append(sharesEqual101, &shareEqualMinus20)
	sharesEqualMinus10 = append(sharesEqual101, &shareEqualMinus20)
	sharesEqualMinus10 = append(sharesEqual101, &shareEqual30)

	tests := []struct {
		name        string
		destination Destination
		want        bool
	}{

		{"Share equal 30", Destination{Account: Account{}, Share: sharesEqual30, BurnShare: nil}, false},
		{"Share equal 80 with burn", Destination{Account: Account{}, Share: sharesEqual30, BurnShare: &burnShare}, false},
		{"Share equal 50", Destination{Account: Account{}, Share: sharesEqual50, BurnShare: nil}, false},
		{"Share equal 100 with burn", Destination{Account: Account{}, Share: sharesEqual50, BurnShare: &burnShare}, true},
		{"Share equal 81", Destination{Account: Account{}, Share: sharesEqual81, BurnShare: nil}, false},
		{"Share equal 100", Destination{Account: Account{}, Share: sharesEqual100, BurnShare: nil}, true},
		{"Share equal 101", Destination{Account: Account{}, Share: sharesEqual101, BurnShare: nil}, true},
		{"Share equal -10", Destination{Account: Account{}, Share: sharesEqualMinus10, BurnShare: nil}, true},
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
		Destination: Destination{
			Account: Account{
				Id:   Id,
				Type: destinationType,
			},
			BurnShare: &BurnShare{
				Percent: sdk.MustNewDecFromStr("0"),
			},
			Share: []*Share{
				{
					Name: helpers.RandStringOfLength(10),
					Account: Account{
						Id:   Id,
						Type: destinationShareType,
					},
					Percent: sdk.MustNewDecFromStr("0"),
				},
			},
		},
		Sources: []*Account{
			{
				Id:   Id,
				Type: sourceType,
			},
		},
	}
}
