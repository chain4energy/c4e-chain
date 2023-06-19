package params

import (
	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/types/address"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	C4eUnit             = "c4e"
	MicroC4eUnit        = "uc4e"
	c4eExponent         = 6
	Bech32PrefixAccAddr = "c4e"
)

var (
	Bech32PrefixAccPub   = Bech32PrefixAccAddr + "pub"
	Bech32PrefixValAddr  = Bech32PrefixAccAddr + "valoper"
	Bech32PrefixValPub   = Bech32PrefixAccAddr + "valoperpub"
	Bech32PrefixConsAddr = Bech32PrefixAccAddr + "valcons"
	Bech32PrefixConsPub  = Bech32PrefixAccAddr + "valconspub"
)

func init() {
	SetAddressPrefixes()
	RegisterDenoms()
	SetAuthorityAddress()
}

func RegisterDenoms() {
	err := sdk.RegisterDenom(C4eUnit, sdk.OneDec())
	if err != nil {
		panic(err)
	}
	err = sdk.RegisterDenom(MicroC4eUnit, sdk.NewDecWithPrec(1, c4eExponent))
	if err != nil {
		panic(err)
	}
}

func SetAddressPrefixes() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)

	// This is copied from the cosmos sdk v0.43.0-beta1
	// source: https://github.com/cosmos/cosmos-sdk/blob/v0.43.0-beta1/types/address.go#L141

	config.SetAddressVerifier(func(bytes []byte) error {
		if len(bytes) == 0 {
			return errors.Wrap(sdkerrors.ErrUnknownAddress, "addresses cannot be empty")
		}

		if len(bytes) > address.MaxAddrLen {
			return errors.Wrapf(sdkerrors.ErrUnknownAddress, "address max length is %d, got %d", address.MaxAddrLen, len(bytes))
		}

		return nil
	})
}

var authorityAddress string

// GetAuthority returns gov module authority address
func GetAuthority() string {
	return authorityAddress
}

// SetAuthorityAddress set gov module authority address
func SetAuthorityAddress() {
	authorityAddress = authtypes.NewModuleAddress(govtypes.ModuleName).String()
}
