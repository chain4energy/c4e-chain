package types

import (
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
			if got := CheckAccountType(tt.account); got != tt.want {
				t.Errorf("CheckAccountType() = %v, want %v", got, tt.want)
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPercentShareSumIsGTEThen100(tt.destination); got != tt.want {
				t.Errorf("CheckPercentShareSumIsGTEThen100() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateOrderOfMainSubDistributors(t *testing.T) {
	sourceMain := createSubDistributor(false, true, false, MAIN)
	destinationMain := createSubDistributor(true, false, false, MAIN)
	destinationShareMain := createSubDistributor(false, false, true, MAIN)

	var zeroSubDistributors []SubDistributor

	var onlyOneMainSubdistributor []SubDistributor
	onlyOneMainSubdistributor = append(onlyOneMainSubdistributor, sourceMain)

	var destinationAtTheEnd []SubDistributor
	destinationAtTheEnd = append(destinationAtTheEnd, sourceMain)
	destinationAtTheEnd = append(destinationAtTheEnd, destinationMain)

	var sourceAtTheEnd []SubDistributor
	sourceAtTheEnd = append(sourceAtTheEnd, sourceMain)
	sourceAtTheEnd = append(sourceAtTheEnd, destinationMain)
	sourceAtTheEnd = append(sourceAtTheEnd, sourceMain)

	var destinationShareAtTheEnd []SubDistributor
	destinationShareAtTheEnd = append(destinationShareAtTheEnd, sourceMain)
	destinationShareAtTheEnd = append(destinationShareAtTheEnd, destinationShareMain)

	var destinationShareSourceAtTheEnd []SubDistributor
	destinationShareSourceAtTheEnd = append(destinationShareSourceAtTheEnd, sourceMain)
	destinationShareSourceAtTheEnd = append(destinationShareSourceAtTheEnd, destinationShareMain)
	destinationShareSourceAtTheEnd = append(destinationShareSourceAtTheEnd, destinationShareMain)
	destinationShareSourceAtTheEnd = append(destinationShareSourceAtTheEnd, sourceMain)

	tests := []struct {
		name            string
		subDistributors []SubDistributor
		wantError       bool
	}{
		{"only one main subdistributor", onlyOneMainSubdistributor, false},
		{"zero sub distributors", zeroSubDistributors, true},
		{"wrong order destination main at the end", destinationAtTheEnd, true},
		{"correct order source main at the end", sourceAtTheEnd, false},
		{"wrong order destination main share at the end", destinationShareAtTheEnd, true},
		{"correct order destination main share, source main at the end", destinationShareSourceAtTheEnd, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateOrderOfSubDistributors(tt.subDistributors)
			if tt.wantError == true && err == nil {
				t.Errorf("ValidateOrderOfSubDistributors() wanted error got nil")
			} else if tt.wantError == false && err != nil {
				t.Errorf("ValidateOrderOfSubDistributors() error: %v", err.Error())
			}
		})
	}
}

func TestValidateOrderOfInternalSubDistributors(t *testing.T) {
	sourceMain := createSubDistributor(false, true, false, MAIN)
	sourceInternal := createSubDistributor(false, true, false, INTERNAL_ACCOUNT)
	destinationInternal := createSubDistributor(true, false, false, INTERNAL_ACCOUNT)
	destinationShareInternal := createSubDistributor(false, false, true, INTERNAL_ACCOUNT)

	var onlyOneInternalSubdistributor []SubDistributor
	onlyOneInternalSubdistributor = append(onlyOneInternalSubdistributor, sourceInternal)
	var destinationAtTheEnd []SubDistributor
	var destinationInternalAtTheEnd []SubDistributor
	destinationAtTheEnd = append(destinationAtTheEnd, sourceInternal)
	destinationAtTheEnd = append(destinationAtTheEnd, destinationInternal)

	var sourceInternalAtTheEnd []SubDistributor
	sourceInternalAtTheEnd = append(sourceInternalAtTheEnd, sourceInternal)
	sourceInternalAtTheEnd = append(sourceInternalAtTheEnd, destinationInternal)
	sourceInternalAtTheEnd = append(sourceInternalAtTheEnd, sourceInternal)
	sourceInternalAtTheEnd = append(sourceInternalAtTheEnd, sourceMain)

	var destinationInternalShareAtTheEnd []SubDistributor
	destinationInternalShareAtTheEnd = append(destinationInternalShareAtTheEnd, sourceInternal)
	destinationInternalShareAtTheEnd = append(destinationInternalShareAtTheEnd, destinationShareInternal)
	sourceInternalAtTheEnd = append(sourceInternalAtTheEnd, sourceMain)

	var destinationShareSourceInternalAtTheEndNoSource []SubDistributor
	destinationShareSourceInternalAtTheEndNoSource = append(destinationShareSourceInternalAtTheEndNoSource, sourceInternal)
	destinationShareSourceInternalAtTheEndNoSource = append(destinationShareSourceInternalAtTheEndNoSource, destinationShareInternal)
	destinationShareSourceInternalAtTheEndNoSource = append(destinationShareSourceInternalAtTheEndNoSource, destinationShareInternal)
	destinationShareSourceInternalAtTheEndNoSource = append(destinationShareSourceInternalAtTheEndNoSource, sourceInternal)

	var destinationShareSourceInternalAtTheEndSource []SubDistributor
	destinationShareSourceInternalAtTheEndSource = append(destinationShareSourceInternalAtTheEndSource, sourceInternal)
	destinationShareSourceInternalAtTheEndSource = append(destinationShareSourceInternalAtTheEndSource, destinationShareInternal)
	destinationShareSourceInternalAtTheEndSource = append(destinationShareSourceInternalAtTheEndSource, destinationShareInternal)
	destinationShareSourceInternalAtTheEndSource = append(destinationShareSourceInternalAtTheEndSource, sourceInternal)
	destinationShareSourceInternalAtTheEndSource = append(destinationShareSourceInternalAtTheEndSource, sourceMain)

	tests := []struct {
		name            string
		subDistributors []SubDistributor
		wantError       bool
	}{
		{"only one internal subdistributor", onlyOneInternalSubdistributor, true},
		{"wrong order destination main at the end", destinationAtTheEnd, true},
		{"wrong order destination internal at the end", destinationInternalAtTheEnd, true},
		{"correct order source main at the end", sourceInternalAtTheEnd, false},
		{"wrong order destination internal share at the end", destinationInternalShareAtTheEnd, true},
		{"correct order destination internalshare, source internal at the end, source main at the end", destinationShareSourceInternalAtTheEndSource, false},
		{"correct order destination internal share, source internal at the end but no main source", destinationShareSourceInternalAtTheEndNoSource, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateOrderOfSubDistributors(tt.subDistributors)
			if tt.wantError == true && err == nil {
				t.Errorf("ValidateOrderOfSubDistributors() wanted error got nil")
			} else if tt.wantError == false && err != nil {
				t.Errorf("ValidateOrderOfSubDistributors() error: %v", err.Error())
			}
		})
	}
}

func createSubDistributor(destinationMain bool, sourceMain bool, destinationShareMain bool, accountType string) SubDistributor {
	destinationType, sourceType, destinationShareType := BASE_ACCOUNT, BASE_ACCOUNT, BASE_ACCOUNT
	if destinationMain {
		destinationType = accountType
	}
	if sourceMain {
		sourceType = accountType
	}
	if destinationShareMain {
		destinationShareType = accountType
	}

	return SubDistributor{
		Name: "",
		Destination: Destination{
			Account: Account{
				Id:   "",
				Type: destinationType,
			},
			BurnShare: &BurnShare{
				Percent: sdk.MustNewDecFromStr("0"),
			},
			Share: []*Share{
				{
					Name: "",
					Account: Account{
						Id:   "",
						Type: destinationShareType,
					},
					Percent: sdk.MustNewDecFromStr("0"),
				},
			},
		},
		Sources: []*Account{
			{
				Id:   "",
				Type: sourceType,
			},
		},
	}
}
