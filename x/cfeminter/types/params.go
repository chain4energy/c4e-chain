package types

import (
	"time"

	"gopkg.in/yaml.v2"
)

var (
	KeyMintDenom     = []byte("MintDenom")
	KeyMinterConfig  = []byte("MinterConfig")
	DefaultMintDenom = "uc4e"
	DefaultStartTime = time.Now()
	DefaultMinters   = []*Minter{
		{
			SequenceId: 1,
			Type:       NoMintingType,
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
