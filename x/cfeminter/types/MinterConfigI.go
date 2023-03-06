package types

import (
	"cosmossdk.io/math"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/libs/log"
	"gopkg.in/yaml.v2"
	"time"
)

var (
	_ MinterConfigI                      = &LinearMinting{}
	_ MinterConfigI                      = &NoMinting{}
	_ MinterConfigI                      = &ExponentialStepMinting{}
	_ codectypes.UnpackInterfacesMessage = (*Minter)(nil)
	_ codectypes.UnpackInterfacesMessage = (*Params)(nil)
	_ codectypes.UnpackInterfacesMessage = (*MsgUpdateMinters)(nil)
)

type MinterConfigI interface {
	codec.ProtoMarshaler
	proto.Message
	Validate() error
	CalculateInflation(totalSupply math.Int, startTime time.Time, endTime *time.Time, blockTime time.Time) sdk.Dec
	AmountToMint(logger log.Logger, startTime time.Time, endTime *time.Time, blockTime time.Time) sdk.Dec
	String() string
}

func (m *Minter) GetMinterConfig() (MinterConfigI, error) {
	if m.Config == nil {
		return nil, fmt.Errorf("minter config is nil")
	}
	minterConfigI, ok := m.Config.GetCachedValue().(MinterConfigI)
	if !ok {
		return nil, fmt.Errorf("expected %T, got %T", (MinterConfigI)(nil), m.Config.GetCachedValue())
	}
	return minterConfigI, nil
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
	minterConfig, _ := m.GetMinterConfig()
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

func (m *NoMinting) Validate() error {
	return nil
}

func (m *NoMinting) CalculateInflation(totalSupply math.Int, startTime time.Time, endTime *time.Time, blockTime time.Time) sdk.Dec {
	return sdk.ZeroDec()
}

func (m *NoMinting) AmountToMint(logger log.Logger, startTime time.Time, endTime *time.Time, blockTime time.Time) sdk.Dec {
	return sdk.ZeroDec()
}

func (m *MsgUpdateMinters) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	for _, minter := range m.Minters {
		if err := minter.UnpackInterfaces(unpacker); err != nil {
			return err
		}
	}
	return nil
}
