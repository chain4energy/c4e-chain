package types

import (
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"time"
)

var (
	KeyMintDenom    = []byte("MintDenom")
	KeyMinterConfig = []byte("MinterConfig")
) //

var _ paramtypes.ParamSet = (*LegacyParams)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&LegacyParams{})
}

// ParamSetPairs get the params.ParamSet
func (params *LegacyParams) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMintDenom, &params.MintDenom, emptyValidation),
		paramtypes.NewParamSetPair(KeyMinterConfig, &params.MinterConfig, emptyValidation),
	}
}

type LegacyParams struct {
	MintDenom    string       `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	MinterConfig MinterConfig `protobuf:"bytes,2,opt,name=minter_config,json=minterConfig,proto3" json:"minter_config"`
}

type MinterConfig struct {
	StartTime time.Time       `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time"`
	Minters   []*LegacyMinter `protobuf:"bytes,3,rep,name=minters,proto3" json:"minters,omitempty"`
}

type LegacyMinter struct {
	SequenceId uint32     `protobuf:"varint,1,opt,name=sequence_id,json=sequenceId,proto3" json:"sequence_id,omitempty"`
	EndTime    *time.Time `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time,omitempty"`
	// types:
	//
	//	NO_MINTING;
	//	LINEAR_MINTING;
	//	EXPONENTIAL_STEP_MINTING;
	Type                   string                  `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	LinearMinting          *LinearMinting          `protobuf:"bytes,4,opt,name=linear_minting,json=linearMinting,proto3" json:"linear_minting,omitempty"`
	ExponentialStepMinting *ExponentialStepMinting `protobuf:"bytes,5,opt,name=exponential_step_minting,json=exponentialStepMinting,proto3" json:"exponential_step_minting,omitempty"`
}

type LegacyMinterState struct {
	SequenceId                  uint32                                 `protobuf:"varint,1,opt,name=sequence_id,json=sequenceId,proto3" json:"sequence_id,omitempty"`
	AmountMinted                github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=amount_minted,json=amountMinted,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount_minted"`
	RemainderToMint             github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=remainder_to_mint,json=remainderToMint,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"remainder_to_mint"`
	LastMintBlockTime           time.Time                              `protobuf:"bytes,4,opt,name=last_mint_block_time,json=lastMintBlockTime,proto3,stdtime" json:"last_mint_block_time"`
	RemainderFromPreviousPeriod github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=remainder_from_previous_period,json=remainderFromPreviousPeriod,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"remainder_from_previous_period"`
}

func emptyValidation(v interface{}) error {
	return nil
}
func DefaultLegacyParams() LegacyParams {
	return NewLegacyParams(DefaultMintDenom, DefaultLegacyMinters)
}
func NewLegacyParams(denom string, minterConfig MinterConfig) LegacyParams {
	return LegacyParams{MintDenom: denom, MinterConfig: minterConfig}
}

var (
	DefaultLegacyMinters = MinterConfig{
		StartTime: time.Now(),
		Minters: []*LegacyMinter{
			{
				SequenceId: 1,
				Type:       "NO_MINTING",
			},
		},
	}
)
