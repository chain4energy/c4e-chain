package types_test

import (
	"github.com/chain4energy/c4e-chain/testutil/module/cfedistributor"
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCheckAccountType(t *testing.T) {
	types.SetMaccPerms(cfedistributor.TestMaccPerms)
	tests := []struct {
		name         string
		account      types.Account
		want         bool
		errorMessage string
	}{
		{"Check base account", types.Account{Id: "c4e1avc7vz3khvlf6fgd3a2exnaqnhhk0sxzzgxc4n", Type: types.BASE_ACCOUNT}, false, ""},
		{"Check module account - account doesn't exist in maccPerms", types.Account{Id: "sample", Type: types.MODULE_ACCOUNT}, true, "module account \"sample\" doesn't exist in maccPerms"},
		{"Check module account - account exists in maccPerms", types.Account{Id: "CUSTOM_ID", Type: types.MODULE_ACCOUNT}, false, ""},
		{"Check internal account", types.Account{Id: "sample", Type: types.INTERNAL_ACCOUNT}, false, ""},
		{"Check main account", types.Account{Id: "sample", Type: types.MAIN}, false, ""},
		{"Check wrong account", types.Account{Id: "test", Type: "wrong_type"}, true, "account \"test\" is of the wrong type: wrong_type"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.account.Validate()
			if tt.want {
				if err == nil {
					t.Errorf("Wanted error got nil")
				}
				if err.Error() != tt.errorMessage {
					t.Errorf("Wanted error: %s , got %s", tt.errorMessage, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Didn't want error got %v", err.Error())
				}
			}
		})
	}
}

func TestCheckPercentShareSumIsGTEThen100(t *testing.T) {

	shareEqual30 := types.DestinationShare{
		Name:        "1",
		Share:       sdk.MustNewDecFromStr("0.30"),
		Destination: types.Account{},
	}

	shareEqual31 := types.DestinationShare{
		Name:        "2",
		Share:       sdk.MustNewDecFromStr("0.31"),
		Destination: types.Account{},
	}

	shareEqual50 := types.DestinationShare{
		Name:        "3",
		Share:       sdk.MustNewDecFromStr("0.50"),
		Destination: types.Account{},
	}

	shareEqual19 := types.DestinationShare{
		Name:        "4",
		Share:       sdk.MustNewDecFromStr("0.19"),
		Destination: types.Account{},
	}

	shareEqual20 := types.DestinationShare{
		Name:        "5",
		Share:       sdk.MustNewDecFromStr("0.20"),
		Destination: types.Account{},
	}

	shareEqualMinus20 := types.DestinationShare{
		Name:        "5",
		Share:       sdk.MustNewDecFromStr("-0.20"),
		Destination: types.Account{},
	}

	burnShare := sdk.MustNewDecFromStr("0.50")

	var sharesEqual30 []*types.DestinationShare
	sharesEqual30 = append(sharesEqual30, &shareEqual30)

	var sharesEqual50 []*types.DestinationShare
	sharesEqual50 = append(sharesEqual50, &shareEqual30)
	sharesEqual50 = append(sharesEqual50, &shareEqual50)

	var sharesEqual81 []*types.DestinationShare
	sharesEqual81 = append(sharesEqual81, &shareEqual30)
	sharesEqual81 = append(sharesEqual81, &shareEqual50)

	var sharesEqual100 []*types.DestinationShare
	sharesEqual100 = append(sharesEqual100, &shareEqual31)
	sharesEqual100 = append(sharesEqual100, &shareEqual50)
	sharesEqual100 = append(sharesEqual100, &shareEqual19)

	var sharesEqual101 []*types.DestinationShare
	sharesEqual101 = append(sharesEqual101, &shareEqual31)
	sharesEqual101 = append(sharesEqual101, &shareEqual50)
	sharesEqual101 = append(sharesEqual101, &shareEqual20)

	var sharesEqualMinus10 []*types.DestinationShare
	sharesEqualMinus10 = append(sharesEqual101, &shareEqualMinus20)
	sharesEqualMinus10 = append(sharesEqual101, &shareEqualMinus20)
	sharesEqualMinus10 = append(sharesEqual101, &shareEqual30)

	tests := []struct {
		name        string
		destination types.Destinations
		want        bool
	}{

		{"Share equal 30", types.Destinations{PrimaryShare: types.Account{}, Shares: sharesEqual30, BurnShare: sdk.ZeroDec()}, false},
		{"Share equal 80 with burn", types.Destinations{PrimaryShare: types.Account{}, Shares: sharesEqual30, BurnShare: burnShare}, false},
		{"Share equal 50", types.Destinations{PrimaryShare: types.Account{}, Shares: sharesEqual50, BurnShare: sdk.ZeroDec()}, false},
		{"Share equal 100 with burn", types.Destinations{PrimaryShare: types.Account{}, Shares: sharesEqual50, BurnShare: burnShare}, true},
		{"Share equal 81", types.Destinations{PrimaryShare: types.Account{}, Shares: sharesEqual81, BurnShare: sdk.ZeroDec()}, false},
		{"Share equal 100", types.Destinations{PrimaryShare: types.Account{}, Shares: sharesEqual100, BurnShare: sdk.ZeroDec()}, true},
		{"Share equal 101", types.Destinations{PrimaryShare: types.Account{}, Shares: sharesEqual101, BurnShare: sdk.ZeroDec()}, true},
		{"Share equal -10", types.Destinations{PrimaryShare: types.Account{}, Shares: sharesEqualMinus10, BurnShare: sdk.ZeroDec()}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.destination.CheckPercentShareSumIsBetween0And1(); got != tt.want {
				t.Errorf("CheckPercentShareSumIsBetween0And100() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateOrderOfMainSubDistributors(t *testing.T) {

	var zeroSubDistributors []types.SubDistributor

	onlyOneMainSubdistributor := []types.SubDistributor{
		CreateSubDistributor(MAIN_SOURCE),
	}
	destinationMainAtTheEnd := []types.SubDistributor{
		CreateSubDistributor(MAIN_SOURCE),
		CreateSubDistributor(MAIN_DESTINATION),
	}

	sourceMainAtTheEnd := []types.SubDistributor{
		CreateSubDistributor(MAIN_SOURCE),
		CreateSubDistributor(MAIN_DESTINATION),
		CreateSubDistributor(MAIN_SOURCE),
	}

	destinationShareMainAtTheEnd := []types.SubDistributor{
		CreateSubDistributor(MAIN_SOURCE),
		CreateSubDistributor(MAIN_DESTINATION_SHARE),
	}

	destinationShareSourceMainAtTheEnd := []types.SubDistributor{
		CreateSubDistributor(MAIN_DESTINATION),
		CreateSubDistributor(MAIN_DESTINATION_SHARE),
		CreateSubDistributor(MAIN_DESTINATION_SHARE),
		CreateSubDistributor(MAIN_SOURCE),
	}

	tests := []struct {
		name            string
		subDistributors []types.SubDistributor
		wantError       bool
	}{
		{"only one main subdistributor", onlyOneMainSubdistributor, false},
		{"zero sub distributors", zeroSubDistributors, true},
		{"wrong order destination main at the end", destinationMainAtTheEnd, true},
		{"correct order source main at the end", sourceMainAtTheEnd, false},
		{"wrong order destination main share at the end", destinationShareMainAtTheEnd, true},
		{"correct order destination main share, source main at the end", destinationShareSourceMainAtTheEnd, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := types.ValidateSubDistributors(tt.subDistributors)
			if tt.wantError == true && err == nil {
				t.Errorf("TestValidateOrderOfMainSubDistributors() wanted error got nil")
			} else if tt.wantError == false && err != nil {
				t.Errorf("TestValidateOrderOfMainSubDistributors() error: %v", err.Error())
			}
		})
	}
}

func TestValidateOrderOfInternalSubDistributors(t *testing.T) {
	onlyOneInternalSubdistributor := []types.SubDistributor{
		CreateSubDistributor(INTERNAL_SOURCE),
	}

	var destinationAtTheEnd []types.SubDistributor

	destinationInternalAtTheEnd := []types.SubDistributor{
		CreateSubDistributor(INTERNAL_SOURCE),
		CreateSubDistributor(INTERNAL_DESTINATION),
	}

	sourceInternalAtTheEnd := []types.SubDistributor{
		CreateSubDistributor(INTERNAL_SOURCE),
		CreateSubDistributor(INTERNAL_DESTINATION),
		CreateSubDistributor(INTERNAL_SOURCE),
		CreateSubDistributor(MAIN_SOURCE),
	}

	destinationInternalShareAtTheEnd := []types.SubDistributor{
		CreateSubDistributor(INTERNAL_DESTINATION_SHARE),
		CreateSubDistributor(INTERNAL_DESTINATION_SHARE),
		CreateSubDistributor(MAIN_SOURCE),
	}

	destinationShareSourceInternalAtTheEndNoSource := []types.SubDistributor{
		CreateSubDistributor(INTERNAL_SOURCE),
		CreateSubDistributor(INTERNAL_DESTINATION_SHARE),
		CreateSubDistributor(INTERNAL_DESTINATION_SHARE),
		CreateSubDistributor(INTERNAL_SOURCE),
	}

	destinationShareSourceInternalAtTheEndSource := []types.SubDistributor{
		CreateSubDistributor(INTERNAL_SOURCE),
		CreateSubDistributor(INTERNAL_DESTINATION_SHARE),
		CreateSubDistributor(INTERNAL_DESTINATION_SHARE),
		CreateSubDistributor(INTERNAL_DESTINATION_SHARE),
		CreateSubDistributor(INTERNAL_SOURCE),
		CreateSubDistributor(MAIN_SOURCE),
	}

	tests := []struct {
		name            string
		subDistributors []types.SubDistributor
		wantError       bool
	}{
		{"only one internal subdistributor", onlyOneInternalSubdistributor, true},
		{"wrong order destination main at the end", destinationAtTheEnd, true},
		{"wrong order destination internal at the end", destinationInternalAtTheEnd, true},
		{"correct order source main at the end", sourceInternalAtTheEnd, false},
		{"wrong order destination internal share at the end", destinationInternalShareAtTheEnd, true},
		{"correct order destination internal share, source internal at the end, source main at the end", destinationShareSourceInternalAtTheEndSource, false},
		{"correct order destination internal share, source internal at the end but no main source", destinationShareSourceInternalAtTheEndNoSource, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := types.ValidateSubDistributors(tt.subDistributors)
			if tt.wantError == true && err == nil {
				t.Errorf("TestValidateOrderOfInternalSubDistributors() wanted error got nil")

			} else if tt.wantError == false && err != nil {
				t.Errorf("TestValidateOrderOfInternalSubDistributors() error: %v", err.Error())
			}
		})
	}
}

func TestValidateUniquenessOfNames(t *testing.T) {
	twoSubdistributorsUniqueNames := []types.SubDistributor{
		CreateSubDistributor(MAIN_DESTINATION),
		CreateSubDistributor(MAIN_DESTINATION_SHARE),
		CreateSubDistributor(MAIN_SOURCE),
	}

	subDistributorMainDestination := CreateSubDistributor(MAIN_DESTINATION)
	subDistributorMainSource := CreateSubDistributor(MAIN_SOURCE)
	subDistributorMainDestinationShare := CreateSubDistributor(MAIN_DESTINATION_SHARE)
	twoSubdistributorsSameNames := []types.SubDistributor{
		subDistributorMainDestination,
		subDistributorMainDestination,
		subDistributorMainSource,
	}

	twoSubdistributorsSameShareNames := []types.SubDistributor{
		subDistributorMainDestinationShare,
		subDistributorMainDestinationShare,
		subDistributorMainSource,
	}

	tests := []struct {
		name            string
		subDistributors []types.SubDistributor
		wantError       bool
	}{
		{"two subdistributors have unique names", twoSubdistributorsUniqueNames, false},
		{"two subdistributors have same names", twoSubdistributorsSameNames, true},
		{"two subdistributors have same share names", twoSubdistributorsSameShareNames, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := types.ValidateSubDistributors(tt.subDistributors)
			if tt.wantError == true && err == nil {
				t.Errorf("TestValidateUniquenessOfNames() wanted error got nil")
			} else if tt.wantError == false && err != nil {
				t.Errorf("TestValidateUniquenessOfNames() error: %v", err.Error())
			}
		})
	}
}

func TestValidateUniquenessOfSubdistributors(t *testing.T) {
	type test struct {
		expectedError   string
		subDistributors []types.SubDistributor
	}
	var tests []test

	for _, accType := range AccountTypes {
		subDistributorCases := make(map[int][]types.SubDistributor)
		subDistributorCases[0] = []types.SubDistributor{
			createSubDistributor(CUSTOM_ACCOUNT, accType, accType, CUSTOM_ID, false),
		}
		subDistributorCases[1] = []types.SubDistributor{
			createSubDistributor(accType, CUSTOM_ACCOUNT, accType, CUSTOM_ID, false),
		}
		subDistributorCases[2] = []types.SubDistributor{
			createSubDistributor(accType, accType, CUSTOM_ACCOUNT, CUSTOM_ID, false),
		}
		subDistributorCases[3] = []types.SubDistributor{
			createSubDistributor(accType, accType, accType, CUSTOM_ID, false),
		}

		sameShares := createSubDistributor(CUSTOM_ACCOUNT, CUSTOM_ACCOUNT_2, accType, CUSTOM_ID, false)
		copiedShare := *sameShares.Destinations.Shares[0]
		copiedShare.Name = helpers.RandStringOfLength(10)
		sameShares.Destinations.Shares = append(sameShares.Destinations.Shares, &copiedShare)
		subDistributorCases[4] = []types.SubDistributor{
			sameShares,
		}

		for i := 0; i < 5; i++ {
			subDistributorCases[i] = append(subDistributorCases[i], CreateSubDistributor(MAIN_SOURCE))
			subDistributorCases[i] = append(subDistributorCases[i], CreateSubDistributor(INTERNAL_SOURCE))
			expectedError := "same " + types.GetId(&types.Account{Type: accType, Id: CUSTOM_ID}) +
				" account cannot occur twice within one subdistributor, subdistributor name: " + subDistributorCases[i][0].Name

			tests = append(tests, test{expectedError, subDistributorCases[i]})
		}
	}

	for _, tt := range tests {
		t.Run(tt.expectedError, func(t *testing.T) {
			err := types.ValidateSubDistributors(tt.subDistributors)
			if err == nil {
				t.Errorf("TestValidateUniquenessOfSubdistributors() wanted error got nil")
			}
			require.EqualValues(t, tt.expectedError, err.Error())
		})
	}
}

func TestValidateUniquenessOfPrimaryShareNames(t *testing.T) {
	type test struct {
		expectedError   string
		subDistributors []types.SubDistributor
	}
	var tests []test

	for _, accType := range AccountTypes {
		subDistributor1 := createSubDistributor(CUSTOM_ACCOUNT, CUSTOM_ACCOUNT_2, accType, CUSTOM_ID, false)
		subDistributor2 := createSubDistributor(CUSTOM_ACCOUNT, CUSTOM_ACCOUNT_2, accType, CUSTOM_ID, false)

		nameWithPrefix := subDistributor2.Name + types.PrimaryShareNameSuffix
		subDistributor2.Destinations.Shares[0].Name = nameWithPrefix

		subDistributors := []types.SubDistributor{
			subDistributor1,
			subDistributor2,
			CreateSubDistributor(MAIN_SOURCE),
			CreateSubDistributor(INTERNAL_SOURCE),
		}

		expectedError := "subdistributor names must be unique, subdistributor name: " + nameWithPrefix
		tests = append(tests, test{expectedError, subDistributors})
	}

	for _, tt := range tests {
		t.Run(tt.expectedError, func(t *testing.T) {
			err := types.ValidateSubDistributors(tt.subDistributors)
			if err == nil {
				t.Errorf("TestValidateUniquenessOfPrimaryShareNames() wanted error got nil")
			}
			require.EqualValues(t, tt.expectedError, err.Error())
		})
	}
}

var AccountTypes = []string{types.INTERNAL_ACCOUNT, types.MODULE_ACCOUNT, types.MAIN, types.BASE_ACCOUNT}

const (
	CUSTOM_ACCOUNT             = "CUSTOM_ACCOUNT"
	CUSTOM_ACCOUNT_2           = "CUSTOM_ACCOUNT-2"
	CUSTOM_ID                  = "custom_id"
	MAIN_SOURCE                = "MAIN_SOURCE"
	MAIN_DESTINATION           = "MAIN_DESTINATION"
	MAIN_DESTINATION_SHARE     = "MAIN_DESTINATION_SHARE"
	INTERNAL_SOURCE            = "INTERNAL_SOURCE"
	INTERNAL_DESTINATION       = "INTERNAL_DESTINATION"
	INTERNAL_DESTINATION_SHARE = "INTERNAL_DESTINATION_SHARE"
)

func CreateSubDistributor(accType string) types.SubDistributor {
	switch accType {
	case MAIN_SOURCE:
		return createSubDistributor(types.BASE_ACCOUNT, types.MAIN, types.BASE_ACCOUNT, CUSTOM_ID, true)
	case MAIN_DESTINATION:
		return createSubDistributor(types.MAIN, types.BASE_ACCOUNT, types.BASE_ACCOUNT, CUSTOM_ID, true)
	case MAIN_DESTINATION_SHARE:
		return createSubDistributor(types.BASE_ACCOUNT, types.BASE_ACCOUNT, types.MAIN, CUSTOM_ID, true)
	case INTERNAL_SOURCE:
		return createSubDistributor(types.BASE_ACCOUNT, types.INTERNAL_ACCOUNT, types.BASE_ACCOUNT, CUSTOM_ID, true)
	case INTERNAL_DESTINATION:
		return createSubDistributor(types.INTERNAL_ACCOUNT, types.BASE_ACCOUNT, types.BASE_ACCOUNT, CUSTOM_ID, true)
	case INTERNAL_DESTINATION_SHARE:
		return createSubDistributor(types.BASE_ACCOUNT, types.BASE_ACCOUNT, types.INTERNAL_ACCOUNT, CUSTOM_ID, true)
	}
	return types.SubDistributor{}
}

func createSubDistributor(
	destinationType string,
	sourceType string,
	destinationShareType string,
	Id string,
	addIdSuffix bool,
) types.SubDistributor {
	return types.SubDistributor{
		Name: helpers.RandStringOfLength(10),
		Destinations: types.Destinations{
			PrimaryShare: types.Account{
				Id:   Id + GetIdSuffix("mainDst", destinationType, addIdSuffix),
				Type: destinationType,
			},
			BurnShare: sdk.ZeroDec(),
			Shares: []*types.DestinationShare{
				{
					Name: helpers.RandStringOfLength(10),
					Destination: types.Account{
						Id:   Id + GetIdSuffix("shareDst", destinationShareType, addIdSuffix),
						Type: destinationShareType,
					},
					Share: sdk.ZeroDec(),
				},
			},
		},
		Sources: []*types.Account{
			{
				Id:   Id + GetIdSuffix("src", sourceType, addIdSuffix),
				Type: sourceType,
			},
		},
	}
}

func GetIdSuffix(suffix string, accType string, addIdSuffix bool) string {
	if !addIdSuffix {
		return ""
	}
	if accType == types.INTERNAL_ACCOUNT || accType == types.MAIN {
		return "-" + accType
	}
	return "-" + suffix
}
