package types

import (
	"cosmossdk.io/math"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

var (
	_ MinterConfigI                      = &LinearMinting{}
	_ MinterConfigI                      = &NoMinting{}
	_ MinterConfigI                      = &ExponentialStepMinting{}
	_ codectypes.UnpackInterfacesMessage = (*Minter)(nil)
	_ codectypes.UnpackInterfacesMessage = (*Params)(nil)
	_ codectypes.UnpackInterfacesMessage = (*MsgUpdateParams)(nil)
	_ codectypes.UnpackInterfacesMessage = (*MsgUpdateMintersParams)(nil)
)

type MinterConfigI interface {
	codec.ProtoMarshaler
	Validate() error
	CalculateInflation(totalSupply math.Int, startTime time.Time, endTime *time.Time, blockTime time.Time) sdk.Dec
	AmountToMint(logger log.Logger, startTime time.Time, endTime *time.Time, blockTime time.Time) sdk.Dec
	String() string
}

func (m *MsgUpdateParams) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	for _, minter := range m.Minters {
		if err := minter.UnpackInterfaces(unpacker); err != nil {
			return err
		}
	}
	return nil
}

func (m *MsgUpdateMintersParams) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	for _, minter := range m.Minters {
		if err := minter.UnpackInterfaces(unpacker); err != nil {
			return err
		}
	}
	return nil
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
