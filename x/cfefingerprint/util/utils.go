package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func CalculateHash(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

// concatenate n strings
// e.g. HashConcat("a", "b", "c") => "a:b:c"
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
