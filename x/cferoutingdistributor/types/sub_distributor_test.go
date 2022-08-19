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
