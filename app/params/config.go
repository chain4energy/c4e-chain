package params

import (
	"github.com/cosmos/cosmos-sdk/types/address"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	CoinDenom           = "uc4e"
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
}

func RegisterDenoms() {
	err := sdk.RegisterDenom(CoinDenom, sdk.OneDec())
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
			return sdkerrors.Wrap(sdkerrors.ErrUnknownAddress, "addresses cannot be empty")
		}

		if len(bytes) > address.MaxAddrLen {
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "address max length is %d, got %d", address.MaxAddrLen, len(bytes))
		}

		return nil
	})
}

// AuthorityAddress is a gov module authority address
var AuthorityAddress = authtypes.NewModuleAddress(govtypes.ModuleName).String()
