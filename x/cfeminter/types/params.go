package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"time"

	"gopkg.in/yaml.v2"
)

var (
	DefaultMintDenom = "uc4e"
	DefaultStartTime = time.Now()
	DefaultMinters   = []*Minter{
		{
			SequenceId: 1,
			Config:     LinearMintingConfig(),
		},
	}
)

// NewParams creates a new Params instance
func NewParams(denom string, startTime time.Time, minters []*Minter) Params {
	return Params{MintDenom: denom, StartTime: startTime, Minters: minters}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultMintDenom, DefaultStartTime, DefaultMinters)
}

// String implements the Stringer interface.
func (params Params) String() string {
	out, _ := yaml.Marshal(params)
	return string(out)
}

func LinearMintingConfig() *codectypes.Any {
	config, _ := codectypes.NewAnyWithValue(&NoMinting{})
	config.TypeUrl = "/c4echain.cfeminter.NoMinting"
	return config
}
