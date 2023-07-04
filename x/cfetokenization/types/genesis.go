package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		CertificateTypeList:  []CertificateType{},
		UserDevicesList:      []UserDevices{},
		UserCertificatesList: []UserCertificates{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in certificateType
	certificateTypeIdMap := make(map[uint64]bool)
	certificateTypeCount := gs.GetCertificateTypeCount()
	for _, elem := range gs.CertificateTypeList {
		if _, ok := certificateTypeIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for certificateType")
		}
		if elem.Id >= certificateTypeCount {
			return fmt.Errorf("certificateType id should be lower or equal than the last id")
		}
		certificateTypeIdMap[elem.Id] = true
	}
	// Check for duplicated ID in userDevices
	userDevicesIdMap := make(map[uint64]bool)
	userDevicesCount := gs.GetUserDevicesCount()
	for _, elem := range gs.UserDevicesList {
		if _, ok := userDevicesIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for userDevices")
		}
		if elem.Id >= userDevicesCount {
			return fmt.Errorf("userDevices id should be lower or equal than the last id")
		}
		userDevicesIdMap[elem.Id] = true
	}
	// Check for duplicated ID in userCertificates
	userCertificatesIdMap := make(map[uint64]bool)
	userCertificatesCount := gs.GetUserCertificatesCount()
	for _, elem := range gs.UserCertificatesList {
		if _, ok := userCertificatesIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for userCertificates")
		}
		if elem.Id >= userCertificatesCount {
			return fmt.Errorf("userCertificates id should be lower or equal than the last id")
		}
		userCertificatesIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
