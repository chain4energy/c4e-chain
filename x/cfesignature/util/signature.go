package util

import (
	"crypto"
	"crypto/x509"
	"encoding/pem"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Strings reflect the x509 standard and are distinct from golang's x509 library string
var signatureAlgorithmDetails = []struct {
	algo       x509.SignatureAlgorithm
	name       string
	pubKeyAlgo x509.PublicKeyAlgorithm
	hash       crypto.Hash
}{
	{x509.DSAWithSHA256, "dsaWithSha256", x509.DSA, crypto.SHA256},
	{x509.ECDSAWithSHA256, "ecdsaWithSha256", x509.ECDSA, crypto.SHA256},
	{x509.SHA256WithRSA, "sha256WithRsaEncryption", x509.RSA, crypto.SHA256},
}

func GetSignatureAlgorithmFromString(name string) (x509.SignatureAlgorithm, error) {
	for _, details := range signatureAlgorithmDetails {
		if details.name == name {
			return details.algo, nil
		}
	}
	return -1, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "signature algorithm not supported")

}

func GetUserCertificateFromString(inputCert []byte) (*x509.Certificate, error) {

	block, _ := pem.Decode(inputCert)
	if block == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "failed to find certificate - PEM formatted block")

	}
	// try to extract a single cert from ASN.1 DER data
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "failed to parse certificate")

	}

	return cert, nil
}
