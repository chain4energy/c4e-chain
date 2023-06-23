package types

import (
	"cosmossdk.io/errors"
	"fmt"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	MaxLatitude  = sdk.NewDec(90)
	MinLatitude  = sdk.NewDec(-90)
	MaxLongitude = sdk.NewDec(180)
	MinLongitude = sdk.NewDec(-180)
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
	case "Type1":
		return Plug1.String()
	case "Type2":
		return Plug2.String()
	default:
		return option
	}
}

func (l Location) Validate() error {
	if l.Latitude == nil {
		return errors.Wrapf(c4eerrors.ErrParam, "latitude cannot be nil")
	}
	if l.Longitude == nil {
		return errors.Wrapf(c4eerrors.ErrParam, "longitude cannot be nil")
	}
	if l.Latitude.GT(MaxLatitude) || l.Latitude.LT(MinLatitude) {
		return errors.Wrapf(c4eerrors.ErrParam, "latitude must be between %s and %s", MaxLatitude, MinLatitude)
	}
	if l.Longitude.GT(MaxLongitude) || l.Longitude.LT(MinLongitude) {
		return errors.Wrapf(c4eerrors.ErrParam, "longitude must be between %s and %s", MaxLongitude, MinLongitude)
	}
	return nil
}

func (plugType PlugType) Validate() error {
	switch plugType {
	case Plug1, Plug2, PlugCHAdeMO, PlugCSS:
		return nil
	default:
		return errors.Wrapf(c4eerrors.ErrParam, "invalid plug type (%s)", plugType)
	}
}
