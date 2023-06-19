package types

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/rand"
)

func CalculateHash(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

func CalculatePayloadLink(referenceId, payloadHash string) string {
	return CalculateHash(HashConcat(referenceId, payloadHash))
}

// HashConcat concatenate n strings e.g. HashConcat("a", "b", "c") => "a:b:c"
func HashConcat(s ...string) string {
	result := ""
	for index, value := range s {
		if index == 0 {
			result = value
		} else {
			result = result + ":" + value
		}
	}
	return result
}

func CreateReferenceID(length int, txHash string) (string, error) {
	dataBytes := []byte(txHash)
	seed := binary.BigEndian.Uint64(dataBytes)

	rand.Seed(int64(seed))
	b := make([]byte, length+2)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b)[2 : length+2], nil
}

func (payloadLink GenesisPayloadLink) Validate() error {
	if payloadLink.ReferenceValue == "" {
		return fmt.Errorf("referance value cannot be empty")
	}
	if payloadLink.ReferenceKey == "" {
		return fmt.Errorf("referance key cannot be empty")
	}
	return nil
}
