package util

import (
	c4eapp "github.com/chain4energy/c4e-chain/v2/app"
	"github.com/chain4energy/c4e-chain/v2/app/params"
	cfemintermoduletypes "github.com/chain4energy/c4e-chain/v2/x/cfeminter/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var (
	EncodingConfig params.EncodingConfig
	Cdc            codec.Codec
)

func init() {
	EncodingConfig, Cdc = initEncodingConfigAndCdc()
	_ = cfemintermoduletypes.Amino
}

func initEncodingConfigAndCdc() (params.EncodingConfig, codec.Codec) {
	encodingConfig := c4eapp.MakeEncodingConfig()

	encodingConfig.InterfaceRegistry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&stakingtypes.MsgCreateValidator{},
	)
	encodingConfig.InterfaceRegistry.RegisterImplementations(
		(*cryptotypes.PubKey)(nil),
		&secp256k1.PubKey{},
		&ed25519.PubKey{},
	)

	cdc := encodingConfig.Codec

	return encodingConfig, cdc
}
