package types

import (
	"cosmossdk.io/math"
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
	return unpacker.UnpackAny(m.Config, &minterConfig)
}

func (params Params) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	for _, minter := range params.Minters {
		if minter.Config != nil {
			var minterConfig MinterConfigI
			err := unpacker.UnpackAny(minter.Config, &minterConfig)
			return err
		}
	}
	return nil
}
