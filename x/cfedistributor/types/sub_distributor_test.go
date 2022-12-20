package types_test

import (
	cfedistributortestutils "github.com/chain4energy/c4e-chain/testutil/module/cfedistributor"
	"github.com/chain4energy/c4e-chain/testutil/simulation/helpers"
	"github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

var baseAccount = types.Account{
	Id:   "abc",
	Type: types.BASE_ACCOUNT,
}

func TestCheckAccountType(t *testing.T) {
	cfedistributortestutils.SetTestMaccPerms()
	tests := []struct {
		name         string
		account      types.Account
		expectError  bool
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
			if tt.expectError {
				require.EqualError(t, err, tt.errorMessage)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidateIfFieldsAreNotSetToNil(t *testing.T) {
	twoSubdistributorsUniqueNames := []types.SubDistributor{
		CreateSubDistributor(MAIN_DESTINATION),
		CreateSubDistributor(MAIN_DESTINATION_SHARE),
		CreateSubDistributor(MAIN_SOURCE),
	}

	tests := []struct {
		name            string
		subDistributors []types.SubDistributor
		expectError     bool
		errorMessage    string
	}{
		{"two subdistributors have unique names", twoSubdistributorsUniqueNames, false, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := types.ValidateSubDistributors(tt.subDistributors)
			if tt.expectError {
				require.EqualError(t, err, tt.errorMessage)
				return
			}
			require.NoError(t, err)
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
		expectError bool
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
			err := tt.destination.CheckIfSharesSumIsBetween0And1()
			if tt.expectError {
				require.EqualError(t, err, "share sum must be between 0 and 1")
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidateDestinationsShareSum(t *testing.T) {
	shareEqual10 := &types.DestinationShare{
		Name:        "1",
		Share:       sdk.MustNewDecFromStr("0.10"),
		Destination: baseAccount,
	}

	shareEqual110 := &types.DestinationShare{
		Name:        "1",
		Share:       sdk.MustNewDecFromStr("1"),
		Destination: baseAccount,
	}

	var sharesEqual30 []*types.DestinationShare
	sharesEqual30 = append(sharesEqual30, shareEqual10)

	var sharesEqual110 []*types.DestinationShare
	sharesEqual110 = append(sharesEqual110, shareEqual10)
	sharesEqual110 = append(sharesEqual110, shareEqual110)

	tests := []struct {
		name         string
		destination  types.Destinations
		expectError  bool
		errorMessage string
	}{

		{"Share equal 30", types.Destinations{PrimaryShare: baseAccount, Shares: sharesEqual30, BurnShare: sdk.ZeroDec()}, false, ""},
		{"Share sum equal 110", types.Destinations{PrimaryShare: baseAccount, Shares: sharesEqual110, BurnShare: sdk.ZeroDec()}, true, "share must be between 0 and 1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.destination.Validate("abc")
			if tt.expectError {
				require.EqualError(t, err, tt.errorMessage)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidateDestinationsShares(t *testing.T) {
	primaryShareName := "primary_name"
	corectShares := []*types.DestinationShare{
		{
			Name:        "Name",
			Share:       sdk.MustNewDecFromStr("0.10"),
			Destination: baseAccount,
		},
	}

	sharesEmptyName := []*types.DestinationShare{
		{
			Name:        "",
			Share:       sdk.MustNewDecFromStr("0.10"),
			Destination: baseAccount,
		},
	}

	sharesShareIsNil := []*types.DestinationShare{
		{
			Name:        "abc",
			Share:       sdk.Dec{},
			Destination: baseAccount,
		},
	}

	sharesWrongDestinationAccount := []*types.DestinationShare{
		{
			Name:        "123",
			Share:       sdk.MustNewDecFromStr("0.10"),
			Destination: types.Account{Id: "WrongTypeAccount", Type: "WrongType"},
		},
	}

	sharesShareIsLessThan0 := []*types.DestinationShare{
		{
			Name:        "123",
			Share:       sdk.MustNewDecFromStr("-0.10"),
			Destination: baseAccount,
		},
	}

	sharesShareIsMoreThan1 := []*types.DestinationShare{
		{
			Name:        "123",
			Share:       sdk.MustNewDecFromStr("1.1"),
			Destination: baseAccount,
		},
	}

	sharesPrimaryShareName := []*types.DestinationShare{
		{
			Name:        primaryShareName,
			Share:       sdk.MustNewDecFromStr("0.5"),
			Destination: baseAccount,
		},
	}

	tests := []struct {
		name         string
		destination  types.Destinations
		expectError  bool
		errorMessage string
	}{

		{"DestinationShare empty name", types.Destinations{PrimaryShare: baseAccount, Shares: sharesEmptyName, BurnShare: sdk.ZeroDec()}, true, "destination share name cannot be empty"},
		{"DestinationShare share is nil", types.Destinations{PrimaryShare: baseAccount, Shares: sharesShareIsNil, BurnShare: sdk.ZeroDec()}, true, "share cannot be nil"},
		{"DestinationShare destination is of wrong type", types.Destinations{PrimaryShare: baseAccount, Shares: sharesWrongDestinationAccount, BurnShare: sdk.ZeroDec()}, true, "account \"WrongTypeAccount\" is of the wrong type: WrongType"},
		{"Share is less than 0", types.Destinations{PrimaryShare: baseAccount, Shares: sharesShareIsLessThan0, BurnShare: sdk.ZeroDec()}, true, "share must be between 0 and 1"},
		{"Share is less more than 1", types.Destinations{PrimaryShare: baseAccount, Shares: sharesShareIsMoreThan1, BurnShare: sdk.ZeroDec()}, true, "share must be between 0 and 1"},
		{"Burn share is nil", types.Destinations{PrimaryShare: baseAccount, Shares: corectShares, BurnShare: sdk.Dec{}}, true, "burn share cannot be nil"},
		{"BurnShare is higher than 1", types.Destinations{PrimaryShare: baseAccount, Shares: corectShares, BurnShare: sdk.MustNewDecFromStr("1.1")}, true, "burn share must be between 0 and 1"},
		{"BurnShare is less than 0", types.Destinations{PrimaryShare: baseAccount, Shares: corectShares, BurnShare: sdk.MustNewDecFromStr("-1")}, true, "burn share must be between 0 and 1"},
		{"Share name reserved for primary share", types.Destinations{PrimaryShare: baseAccount, Shares: sharesPrimaryShareName, BurnShare: sdk.ZeroDec()}, true, "share name: " + primaryShareName + " is reserved for primary share"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.destination.Validate(primaryShareName)
			if tt.expectError {
				require.EqualError(t, err, tt.errorMessage)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidateSubDistributor(t *testing.T) {
	subdistributorCorrect := CreateSubDistributor(MAIN_SOURCE)

	subdistributorNoName := CreateSubDistributor(MAIN_SOURCE)
	subdistributorNoName.Name = ""

	subdistributorNoSources := CreateSubDistributor(MAIN_SOURCE)
	subdistributorNoSources.Sources = []*types.Account{}

	subdistributorNilSource := CreateSubDistributor(MAIN_SOURCE)
	subdistributorNilSource.Sources[0] = nil

	subdistributorEmptyDestinations := CreateSubDistributor(MAIN_SOURCE)
	subdistributorEmptyDestinations.Destinations = types.Destinations{}

	tests := []struct {
		name           string
		subDistributor types.SubDistributor
		expectError    bool
		errorMessage   string
	}{
		{"correct subdistributor", subdistributorCorrect, false, ""},
		{"subdistributor has no name", subdistributorNoName, true, "subdistributor name cannot be empty"},
		{"subdistributor has no sources", subdistributorNoSources, true, "subdistributor must have at least one source"},
		{"subdistributor has source with nil type", subdistributorNilSource, true, "source cannot be nil"},
		{"subdistributor has empty destinations", subdistributorEmptyDestinations, true, "destinations cannot be empty"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.subDistributor.Validate()
			if tt.expectError {
				require.EqualError(t, err, tt.errorMessage)
				return
			}
			require.NoError(t, err)
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
		expectError     bool
		errorMessage    string
	}{
		{"only one main subdistributor", onlyOneMainSubdistributor, false, ""},
		{"zero sub distributors", zeroSubDistributors, true, "there must be at least one subdistributor with the source main type"},
		{"wrong order destination main at the end", destinationMainAtTheEnd, true, "wrong order of subdistributors, after each occurrence of a subdistributor with the destination of internal or main account type there must be exactly one occurrence of a subdistributor with the source of internal account type, account id: MAIN"},
		{"correct order source main at the end", sourceMainAtTheEnd, false, ""},
		{"wrong order destination main share at the end", destinationShareMainAtTheEnd, true, "wrong order of subdistributors, after each occurrence of a subdistributor with the destination of internal or main account type there must be exactly one occurrence of a subdistributor with the source of internal account type, account id: MAIN"},
		{"correct order destination main share, source main at the end", destinationShareSourceMainAtTheEnd, false, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := types.ValidateSubDistributors(tt.subDistributors)
			if tt.expectError {
				require.EqualError(t, err, tt.errorMessage)
				return
			}
			require.NoError(t, err)
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
		expectError     bool
		errorMessage    string
	}{
		{"only one internal subdistributor", onlyOneInternalSubdistributor, true, "there must be at least one subdistributor with the source main type"},
		{"wrong order destination main at the end", destinationAtTheEnd, true, "there must be at least one subdistributor with the source main type"},
		{"wrong order destination internal at the end", destinationInternalAtTheEnd, true, "there must be at least one subdistributor with the source main type"},
		{"correct order source main at the end", sourceInternalAtTheEnd, false, ""},
		{"wrong order destination internal share at the end", destinationInternalShareAtTheEnd, true, "wrong order of subdistributors, after each occurrence of a subdistributor with the destination of internal or main account type there must be exactly one occurrence of a subdistributor with the source of internal account type, account id: INTERNAL_ACCOUNT-CUSTOM_ID-INTERNAL_ACCOUNT"},
		{"correct order destination internal share, source internal at the end, source main at the end", destinationShareSourceInternalAtTheEndSource, false, ""},
		{"correct order destination internal share, source internal at the end but no main source", destinationShareSourceInternalAtTheEndNoSource, true, "there must be at least one subdistributor with the source main type"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := types.ValidateSubDistributors(tt.subDistributors)
			if len(tt.errorMessage) > 0 {
				require.Equal(t, err.Error(), tt.errorMessage)
				return
			}
			require.NoError(t, err)
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
		expectError     bool
		errorMessage    string
	}{
		{"two subdistributors have unique names", twoSubdistributorsUniqueNames, false, ""},
		{"two subdistributors have same names", twoSubdistributorsSameNames, true, "subdistributor names must be unique, subdistributor name: " + twoSubdistributorsSameNames[1].Name},
		{"two subdistributors have same share names", twoSubdistributorsSameShareNames, true, "subdistributor names must be unique, subdistributor name: " + twoSubdistributorsSameShareNames[1].Name},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := types.ValidateSubDistributors(tt.subDistributors)
			if tt.expectError {
				require.EqualError(t, err, tt.errorMessage)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidateCorrectModuleAccountInsideSubdistributor(t *testing.T) {
	cfedistributortestutils.SetTestMaccPerms()
	correctPrimaryShareModuleAccount := createSubDistributor(types.MODULE_ACCOUNT, types.BASE_ACCOUNT, types.BASE_ACCOUNT, CUSTOM_ID, false)
	correctSourceModuleAccount := createSubDistributor(types.BASE_ACCOUNT, types.MODULE_ACCOUNT, types.BASE_ACCOUNT, CUSTOM_ID, false)
	correctShareModuleAccount := createSubDistributor(types.BASE_ACCOUNT, types.BASE_ACCOUNT, types.MODULE_ACCOUNT, CUSTOM_ID, false)
	wrongPrimaryShareModuleAccount := createSubDistributor(types.MODULE_ACCOUNT, types.BASE_ACCOUNT, types.BASE_ACCOUNT, CUSTOM_ID, true)
	wrongSourceModuleAccount := createSubDistributor(types.BASE_ACCOUNT, types.MODULE_ACCOUNT, types.BASE_ACCOUNT, CUSTOM_ID, true)
	wrongShareModuleAccount := createSubDistributor(types.BASE_ACCOUNT, types.BASE_ACCOUNT, types.MODULE_ACCOUNT, CUSTOM_ID, true)

	tests := []struct {
		name            string
		subDistributors types.SubDistributor
		expectError     bool
		errorMessage    string
	}{
		{"correct primary share module account", correctPrimaryShareModuleAccount, false, ""},
		{"correct source module account", correctSourceModuleAccount, false, ""},
		{"correct share module account", correctShareModuleAccount, false, ""},
		{"wrong primary share module account", wrongPrimaryShareModuleAccount, true, "module account \"CUSTOM_ID-mainDst\" doesn't exist in maccPerms"},
		{"wrong source module account", wrongSourceModuleAccount, true, "module account \"CUSTOM_ID-src\" doesn't exist in maccPerms"},
		{"wrong share module account", wrongShareModuleAccount, true, "module account \"CUSTOM_ID-shareDst\" doesn't exist in maccPerms"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.subDistributors.Validate()
			if tt.expectError {
				require.EqualError(t, err, tt.errorMessage)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidateUniquenessOfSubdistributors(t *testing.T) {
	type test struct {
		subDistributors []types.SubDistributor
		expectError     bool
		errorMessage    string
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
			accId := accType + "-" + CUSTOM_ID
			if accType == types.MAIN {
				accId = accType
			}
			errorMessage := "same " + accId + " account cannot occur twice within one subdistributor, subdistributor name: " + subDistributorCases[i][0].Name

			tests = append(tests, test{subDistributorCases[i], true, errorMessage})
		}
	}

	for _, tt := range tests {
		t.Run(tt.errorMessage, func(t *testing.T) {
			err := types.ValidateSubDistributors(tt.subDistributors)
			if tt.expectError {
				require.EqualError(t, err, tt.errorMessage)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidateUniquenessOfPrimaryShareNames(t *testing.T) {
	type test struct {
		subDistributors []types.SubDistributor
		expectError     bool
		errorMessage    string
	}
	var tests []test

	for _, accType := range AccountTypes {
		subDistributor1 := createSubDistributor(CUSTOM_ACCOUNT, CUSTOM_ACCOUNT_2, accType, CUSTOM_ID, false)
		subDistributor2 := createSubDistributor(CUSTOM_ACCOUNT, CUSTOM_ACCOUNT_2, accType, CUSTOM_ID, false)

		nameWithPrefix := subDistributor2.Name + "_primary"
		subDistributor2.Destinations.Shares[0].Name = nameWithPrefix

		subDistributors := []types.SubDistributor{
			subDistributor1,
			subDistributor2,
			CreateSubDistributor(MAIN_SOURCE),
			CreateSubDistributor(INTERNAL_SOURCE),
		}

		errorMessage := "subdistributor names must be unique, subdistributor name: " + nameWithPrefix
		tests = append(tests, test{subDistributors, true, errorMessage})
	}

	for _, tt := range tests {
		t.Run(tt.errorMessage, func(t *testing.T) {
			err := types.ValidateSubDistributors(tt.subDistributors)
			if tt.expectError {
				require.EqualError(t, err, tt.errorMessage)
				return
			}
			require.NoError(t, err)
		})
	}
}

var AccountTypes = []string{types.INTERNAL_ACCOUNT, types.MODULE_ACCOUNT, types.MAIN, types.BASE_ACCOUNT}

const (
	CUSTOM_ACCOUNT             = "CUSTOM_ACCOUNT"
	CUSTOM_ACCOUNT_2           = "CUSTOM_ACCOUNT-2"
	CUSTOM_ID                  = "CUSTOM_ID"
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
	primaryShareType string,
	sourceType string,
	destinationShareType string,
	Id string,
	addIdSuffix bool,
) types.SubDistributor {
	return types.SubDistributor{
		Name: helpers.RandStringOfLength(10),
		Destinations: types.Destinations{
			PrimaryShare: types.Account{
				Id:   Id + GetIdSuffix("mainDst", primaryShareType, addIdSuffix),
				Type: primaryShareType,
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
