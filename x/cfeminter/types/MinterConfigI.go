package types

import (
	"cosmossdk.io/math"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/libs/log"
	"gopkg.in/yaml.v2"
	"time"
)

var (
	_ MinterConfigI                      = (*LinearMinting)(nil)
	_ MinterConfigI                      = (*ExponentialStepMinting)(nil)
	_ codectypes.UnpackInterfacesMessage = (*Minter)(nil)
	_ codectypes.UnpackInterfacesMessage = (*Params)(nil)
)

type MinterConfigI interface {
	proto.Message
	Validate() error
	CalculateInflation(totalSupply math.Int, startTime time.Time, endTime *time.Time, blockTime time.Time) sdk.Dec
	AmountToMint(logger log.Logger, startTime time.Time, endTime *time.Time, blockTime time.Time) sdk.Dec
	String() string
}

func (m Minter) GetMinterConfig() MinterConfigI {
	if m.Config == nil {
		return nil
	}
	minterConfigI, ok := m.Config.GetCachedValue().(MinterConfigI)
	if !ok {
		return nil
	}
	return minterConfigI
}

func (m Minter) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var minterConfig MinterConfigI
	if m.Config != nil {
		if err := unpacker.UnpackAny(m.Config, &minterConfig); err != nil {
			return err
		}
	}
	return nil
}

func (params Params) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	for _, minter := range params.Minters {
		if err := minter.UnpackInterfaces(unpacker); err != nil {
			return err
		}
	}
	return nil
}

func (acc *Minter) String() string {
	out, _ := yaml.Marshal(acc)
	return string(out)
}

type MinterJSON struct {
	SequenceId uint32     `json:"sequence_id"`
	EndTime    *time.Time `json:"end_time"`

	// custom fields based on concrete vesting type which can be omitted
	Config string `json:"config,omitempty"`
	Type   string `json:"type,omitempty"`
}

func (m *Minter) GetMinterJSON() MinterJSON {
	if m == nil {
		return MinterJSON{}
	}
	minterConfig := m.GetMinterConfig()
	var config string
	if minterConfig != nil {
		config = minterConfig.String()
	}
	return MinterJSON{
		SequenceId: m.SequenceId,
		EndTime:    m.EndTime,
		Type:       m.Config.GetTypeUrl(),
		Config:     config,
	}
}
