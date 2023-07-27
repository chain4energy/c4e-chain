package types

import "fmt"

func (u UserCertificates) GetUserCertificate(certificateId uint64) (*Certificate, error) {
	for _, certificate := range u.Certificates {
		if certificate.Id == certificateId {
			return certificate, nil
		}
	}
	return nil, fmt.Errorf("certificate not found")
}

func (u Certificate) ValidateAuthorizer(authorizerAddress string) bool {
	for _, allowedAuthority := range u.AllowedAuthorities {
		if allowedAuthority == authorizerAddress {
			return true
		}
	}
	return false
}
