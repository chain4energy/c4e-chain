package encoding

import (
	"github.com/chain4energy/c4e-chain/app"
	params2 "github.com/chain4energy/c4e-chain/tests/e2e/encoding/params"
	"github.com/cosmos/cosmos-sdk/std"
)

// TODO: comment
// MakeEncodingConfig creates an EncodingConfig for testing.
func MakeEncodingConfig() params2.EncodingConfig {
	encodingConfig := params2.MakeEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	app.ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	app.ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
