package util

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

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

func CalculateHash(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

// extract a single field from input json
func ExtractFieldFromJSON(jsonInput string, field string) (string, error) {
	var input map[string]interface{}
	err := json.Unmarshal([]byte(jsonInput), &input)
	if err != nil {
		// return "", errorcode.Internal.WithMessage("failed to parse signature json, %v", err).LogReturn()
		return "", err
	}

	// try to extract the value
	value, exists := input[field].(string)
	if !exists {
		// return "", errorcode.Internal.WithMessage("failed to parse signature json, missing field %s", field).LogReturn()
		return "", err
	}

	return value, nil
}
