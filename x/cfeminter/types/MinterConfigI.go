package types

import (
	"cosmossdk.io/math"
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/libs/log"
	"time"
)

var (
	_ MinterConfigI = (*LinearMinting)(nil)
	_ MinterConfigI = (*ExponentialStepMinting)(nil)
	_ proto.Message = (*LinearMinting)(nil)
	_ proto.Message = (*ExponentialStepMinting)(nil)

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

func (m *Minter) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var minterConfig MinterConfigI
	if m.Config != nil {
		return unpacker.UnpackAny(m.Config, &minterConfig)
	}
	return nil
}

func (params Params) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	for _, minter := range params.Minters {
		if minter.Config != nil {
			var minterConfig MinterConfigI
			if err := unpacker.UnpackAny(minter.Config, &minterConfig); err != nil {
				return err
			}
		}
	}
	return nil
}

// MarshalJSON returns the JSON representation of a ModuleAccount.
func (ma ExponentialStepMinting) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExponentialStepMinting{
		StepDuration:     ma.StepDuration,
		Amount:           ma.Amount,
		AmountMultiplier: ma.AmountMultiplier,
	})
}

// UnmarshalJSON unmarshals raw JSON bytes into a ModuleAccount.
func (ma *ExponentialStepMinting) UnmarshalJSON(bz []byte) error {
	var alias ExponentialStepMinting
	if err := json.Unmarshal(bz, &alias); err != nil {
		return err
	}
	return nil
}

// MarshalJSON returns the JSON representation of a ModuleAccount.
func (ma LinearMinting) MarshalJSON() ([]byte, error) {
	return json.Marshal(LinearMinting{
		Amount: ma.Amount,
	})
}

// UnmarshalJSON unmarshals raw JSON bytes into a ModuleAccount.
func (ma *LinearMinting) UnmarshalJSON(bz []byte) error {
	var alias LinearMinting
	if err := json.Unmarshal(bz, &alias); err != nil {
		return err
	}
	return nil
}

func (acc ExponentialStepMinting) MarshalYAML() (interface{}, error) {
	bz, err := codec.MarshalYAML(codec.NewProtoCodec(codectypes.NewInterfaceRegistry()), &acc)
	if err != nil {
		return nil, err
	}
	return string(bz), err
}

func (acc LinearMinting) MarshalYAML() (interface{}, error) {
	bz, err := codec.MarshalYAML(codec.NewProtoCodec(codectypes.NewInterfaceRegistry()), &acc)
	if err != nil {
		return nil, err
	}
	return string(bz), err
}

func (acc ExponentialStepMinting) String() string {
	out, _ := acc.MarshalYAML()
	return out.(string)
}

func (acc LinearMinting) String() string {
	out, _ := acc.MarshalYAML()
	return out.(string)
}
