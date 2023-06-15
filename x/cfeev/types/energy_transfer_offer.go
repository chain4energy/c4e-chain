package types

import (
	"fmt"
)

// Campaign types
const (
	UnspecifiedPlugType = PlugType_PLUG_TYPE_UNSPECIFIED
	PlugCSS             = PlugType_CCS
	PlugCHAdeMO         = PlugType_CHAdeMO
	Plug1               = PlugType_Type1
	Plug2               = PlugType_Type2
)

func PlugTypeFromString(str string) (PlugType, error) {
	option, ok := PlugType_value[str]
	if !ok {
		return UnspecifiedPlugType, fmt.Errorf("'%s' is not a valid plug type, available options: CSS/CHAdeMO/1/2", str)
	}
	return PlugType(option), nil
}

// NormalizePlugType - normalize plug type
func NormalizePlugType(option string) string {
	switch option {
	case "CSS", "css":
		return PlugCSS.String()

	case "CHAdeMO", "chadem0":
		return PlugCHAdeMO.String()
	case "1":
		return Plug1.String()
	case "2":
		return Plug2.String()

	default:
		return option
	}
}
