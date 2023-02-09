package encoding

import (
	"github.com/chain4energy/c4e-chain/app"
	params2 "github.com/chain4energy/c4e-chain/tests/e2e/encoding/params"
	"github.com/cosmos/cosmos-sdk/std"
)

// e2e framework requires a new way of declaring parameters and
// encoding for the application. The way it has been added here is identical to how
// these variables and functions are declared in the new cosmos sdk version (0.46.x)

// MakeEncodingConfig creates an EncodingConfig for testing.
func MakeEncodingConfig() params2.EncodingConfig {
	encodingConfig := params2.MakeEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	app.ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	app.ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
