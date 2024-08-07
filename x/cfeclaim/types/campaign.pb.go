// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chain4energy/c4echain/cfeclaim/campaign.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type CampaignType int32

const (
	CampaignType_CAMPAIGN_TYPE_UNSPECIFIED CampaignType = 0
	CampaignType_DEFAULT                   CampaignType = 1
	CampaignType_VESTING_POOL              CampaignType = 2
)

var CampaignType_name = map[int32]string{
	0: "CAMPAIGN_TYPE_UNSPECIFIED",
	1: "DEFAULT",
	2: "VESTING_POOL",
}

var CampaignType_value = map[string]int32{
	"CAMPAIGN_TYPE_UNSPECIFIED": 0,
	"DEFAULT":                   1,
	"VESTING_POOL":              2,
}

func (x CampaignType) String() string {
	return proto.EnumName(CampaignType_name, int32(x))
}

func (CampaignType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_533e85cec41311cf, []int{0}
}

type Campaign struct {
	Id                     uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner                  string                                 `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Name                   string                                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description            string                                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	CampaignType           CampaignType                           `protobuf:"varint,5,opt,name=campaignType,proto3,enum=chain4energy.c4echain.cfeclaim.CampaignType" json:"campaignType,omitempty"`
	RemovableClaimRecords  bool                                   `protobuf:"varint,6,opt,name=removable_claim_records,json=removableClaimRecords,proto3" json:"removable_claim_records,omitempty"`
	FeegrantAmount         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=feegrant_amount,json=feegrantAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"feegrant_amount"`
	InitialClaimFreeAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,8,opt,name=initial_claim_free_amount,json=initialClaimFreeAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"initial_claim_free_amount"`
	Free                   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,9,opt,name=free,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"free"`
	Enabled                bool                                   `protobuf:"varint,10,opt,name=enabled,proto3" json:"enabled,omitempty"`
	StartTime              time.Time                              `protobuf:"bytes,11,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time"`
	EndTime                time.Time                              `protobuf:"bytes,12,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time"`
	// period of locked coins from claim
	LockupPeriod time.Duration `protobuf:"bytes,13,opt,name=lockup_period,json=lockupPeriod,proto3,stdduration" json:"lockup_period"`
	// period of vesting coins after lockup period
	VestingPeriod         time.Duration                            `protobuf:"bytes,14,opt,name=vesting_period,json=vestingPeriod,proto3,stdduration" json:"vesting_period"`
	VestingPoolName       string                                   `protobuf:"bytes,15,opt,name=vestingPoolName,proto3" json:"vestingPoolName,omitempty"`
	CampaignTotalAmount   github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,16,rep,name=campaign_total_amount,json=campaignTotalAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"campaign_total_amount"`
	CampaignCurrentAmount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,17,rep,name=campaign_current_amount,json=campaignCurrentAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"campaign_current_amount"`
}

func (m *Campaign) Reset()         { *m = Campaign{} }
func (m *Campaign) String() string { return proto.CompactTextString(m) }
func (*Campaign) ProtoMessage()    {}
func (*Campaign) Descriptor() ([]byte, []int) {
	return fileDescriptor_533e85cec41311cf, []int{0}
}
func (m *Campaign) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Campaign) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Campaign.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Campaign) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Campaign.Merge(m, src)
}
func (m *Campaign) XXX_Size() int {
	return m.Size()
}
func (m *Campaign) XXX_DiscardUnknown() {
	xxx_messageInfo_Campaign.DiscardUnknown(m)
}

var xxx_messageInfo_Campaign proto.InternalMessageInfo

func (m *Campaign) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Campaign) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Campaign) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Campaign) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Campaign) GetCampaignType() CampaignType {
	if m != nil {
		return m.CampaignType
	}
	return CampaignType_CAMPAIGN_TYPE_UNSPECIFIED
}

func (m *Campaign) GetRemovableClaimRecords() bool {
	if m != nil {
		return m.RemovableClaimRecords
	}
	return false
}

func (m *Campaign) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Campaign) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

func (m *Campaign) GetEndTime() time.Time {
	if m != nil {
		return m.EndTime
	}
	return time.Time{}
}

func (m *Campaign) GetLockupPeriod() time.Duration {
	if m != nil {
		return m.LockupPeriod
	}
	return 0
}

func (m *Campaign) GetVestingPeriod() time.Duration {
	if m != nil {
		return m.VestingPeriod
	}
	return 0
}

func (m *Campaign) GetVestingPoolName() string {
	if m != nil {
		return m.VestingPoolName
	}
	return ""
}

func (m *Campaign) GetCampaignTotalAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.CampaignTotalAmount
	}
	return nil
}

func (m *Campaign) GetCampaignCurrentAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.CampaignCurrentAmount
	}
	return nil
}

func init() {
	proto.RegisterEnum("chain4energy.c4echain.cfeclaim.CampaignType", CampaignType_name, CampaignType_value)
	proto.RegisterType((*Campaign)(nil), "chain4energy.c4echain.cfeclaim.Campaign")
}

func init() {
	proto.RegisterFile("chain4energy/c4echain/cfeclaim/campaign.proto", fileDescriptor_533e85cec41311cf)
}

var fileDescriptor_533e85cec41311cf = []byte{
	// 723 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x41, 0x4f, 0xdb, 0x48,
	0x14, 0x8e, 0x43, 0x20, 0x61, 0x12, 0x42, 0x76, 0x16, 0x16, 0x83, 0xb4, 0x4e, 0xb4, 0x87, 0x55,
	0xb4, 0x5a, 0xec, 0x42, 0x51, 0xaf, 0x55, 0xe2, 0x04, 0x9a, 0x2a, 0x0d, 0x91, 0x09, 0xad, 0xda,
	0x8b, 0xe5, 0xd8, 0x13, 0x33, 0x22, 0x9e, 0xb1, 0xc6, 0x13, 0x5a, 0x4e, 0x3d, 0xf4, 0x0f, 0x70,
	0xec, 0x6f, 0xe8, 0x2f, 0xe1, 0xc8, 0xb1, 0xea, 0x01, 0x2a, 0xf8, 0x23, 0xd5, 0x8c, 0xed, 0x90,
	0x52, 0xa9, 0x82, 0xaa, 0x27, 0x7b, 0xe6, 0xbd, 0xf7, 0x7d, 0xdf, 0xbc, 0xf9, 0xde, 0x80, 0x4d,
	0xf7, 0xc8, 0xc1, 0x64, 0x07, 0x11, 0xc4, 0xfc, 0x53, 0xc3, 0xdd, 0x41, 0x72, 0x6d, 0xb8, 0x23,
	0xe4, 0x8e, 0x1d, 0x1c, 0x18, 0xae, 0x13, 0x84, 0x0e, 0xf6, 0x89, 0x1e, 0x32, 0xca, 0x29, 0xd4,
	0x66, 0xd3, 0xf5, 0x34, 0x5d, 0x4f, 0xd3, 0x37, 0x56, 0x7c, 0xea, 0x53, 0x99, 0x6a, 0x88, 0xbf,
	0xb8, 0x6a, 0xa3, 0xea, 0x53, 0xea, 0x8f, 0x91, 0x21, 0x57, 0xc3, 0xc9, 0xc8, 0xe0, 0x38, 0x40,
	0x11, 0x77, 0x82, 0x30, 0x49, 0xd0, 0xee, 0x26, 0x78, 0x13, 0xe6, 0x70, 0x4c, 0x49, 0x1a, 0x77,
	0x69, 0x14, 0xd0, 0xc8, 0x18, 0x3a, 0x11, 0x32, 0x4e, 0xb6, 0x86, 0x88, 0x3b, 0x5b, 0x86, 0x4b,
	0x71, 0x12, 0xff, 0xe7, 0xaa, 0x00, 0x0a, 0x66, 0xa2, 0x14, 0x96, 0x41, 0x16, 0x7b, 0xaa, 0x52,
	0x53, 0xea, 0x39, 0x2b, 0x8b, 0x3d, 0xb8, 0x02, 0xe6, 0xe9, 0x5b, 0x82, 0x98, 0x9a, 0xad, 0x29,
	0xf5, 0x45, 0x2b, 0x5e, 0x40, 0x08, 0x72, 0xc4, 0x09, 0x90, 0x3a, 0x27, 0x37, 0xe5, 0x3f, 0xac,
	0x81, 0xa2, 0x87, 0x22, 0x97, 0xe1, 0x50, 0x70, 0xab, 0x39, 0x19, 0x9a, 0xdd, 0x82, 0x7d, 0x50,
	0x4a, 0x3b, 0x32, 0x38, 0x0d, 0x91, 0x3a, 0x5f, 0x53, 0xea, 0xe5, 0xed, 0xff, 0xf5, 0x9f, 0xb7,
	0x45, 0x37, 0x67, 0x6a, 0xac, 0xef, 0x10, 0xe0, 0x13, 0xb0, 0xc6, 0x50, 0x40, 0x4f, 0x9c, 0xe1,
	0x18, 0xd9, 0x32, 0xdb, 0x66, 0xc8, 0xa5, 0xcc, 0x8b, 0xd4, 0x85, 0x9a, 0x52, 0x2f, 0x58, 0xab,
	0xd3, 0xb0, 0x29, 0xa2, 0x56, 0x1c, 0x84, 0xaf, 0xc0, 0xf2, 0x08, 0x21, 0x9f, 0x39, 0x84, 0xdb,
	0x4e, 0x40, 0x27, 0x84, 0xab, 0x79, 0xa1, 0xb7, 0xa9, 0x9f, 0x5f, 0x56, 0x33, 0x5f, 0x2e, 0xab,
	0xff, 0xfa, 0x98, 0x1f, 0x4d, 0x86, 0xba, 0x4b, 0x03, 0x23, 0x69, 0x5f, 0xfc, 0xd9, 0x8c, 0xbc,
	0x63, 0x83, 0x9f, 0x86, 0x28, 0xd2, 0x3b, 0x84, 0x5b, 0xe5, 0x14, 0xa6, 0x21, 0x51, 0x20, 0x06,
	0xeb, 0x98, 0x60, 0x8e, 0x9d, 0x71, 0x22, 0x67, 0xc4, 0x10, 0x4a, 0x29, 0x0a, 0xbf, 0x44, 0xf1,
	0x57, 0x02, 0x28, 0x0f, 0xb0, 0xcb, 0x10, 0x4a, 0xa8, 0x9a, 0x20, 0x27, 0xc0, 0xd5, 0xc5, 0x07,
	0xa3, 0xb6, 0x90, 0x6b, 0xc9, 0x5a, 0xa8, 0x82, 0x3c, 0x22, 0xa2, 0x3b, 0x9e, 0x0a, 0x64, 0xbf,
	0xd2, 0x25, 0x34, 0x01, 0x88, 0xb8, 0xc3, 0xb8, 0x2d, 0xdc, 0xa6, 0x16, 0x6b, 0x4a, 0xbd, 0xb8,
	0xbd, 0xa1, 0xc7, 0x4e, 0xd3, 0x53, 0xa7, 0xe9, 0x83, 0xd4, 0x8a, 0xcd, 0x82, 0xe0, 0x3f, 0xbb,
	0xaa, 0x2a, 0xd6, 0xa2, 0xac, 0x13, 0x11, 0xf8, 0x14, 0x14, 0x10, 0xf1, 0x62, 0x88, 0xd2, 0x03,
	0x20, 0xf2, 0x88, 0x78, 0x12, 0xe0, 0x19, 0x58, 0x1a, 0x53, 0xf7, 0x78, 0x12, 0xda, 0x21, 0x62,
	0x98, 0x7a, 0xea, 0x92, 0x44, 0x59, 0xff, 0x01, 0xa5, 0x95, 0x58, 0x3e, 0x06, 0xf9, 0x28, 0x40,
	0x4a, 0x71, 0x65, 0x5f, 0x16, 0xc2, 0xe7, 0xa0, 0x7c, 0x82, 0x22, 0x8e, 0x89, 0x9f, 0x42, 0x95,
	0xef, 0x0f, 0xb5, 0x94, 0x94, 0x26, 0x58, 0x75, 0xb0, 0x9c, 0x6e, 0x50, 0x3a, 0xee, 0x89, 0x41,
	0x58, 0x96, 0x6e, 0xbf, 0xbb, 0x0d, 0xdf, 0x83, 0xd5, 0xd4, 0xaf, 0x36, 0xa7, 0xdc, 0x19, 0xa7,
	0x56, 0xa8, 0xd4, 0xe6, 0x24, 0x79, 0x7c, 0x37, 0xba, 0x18, 0x4d, 0x3d, 0x19, 0x4d, 0xdd, 0xa4,
	0x98, 0x34, 0x1f, 0x09, 0xf2, 0x4f, 0x57, 0xd5, 0xfa, 0x3d, 0xee, 0x53, 0x14, 0x44, 0xd6, 0x9f,
	0xd3, 0xc9, 0x10, 0x44, 0x89, 0x49, 0x3e, 0x28, 0x60, 0x6d, 0xaa, 0xc0, 0x9d, 0x30, 0x86, 0x6e,
	0x1d, 0xff, 0xc7, 0xef, 0xd7, 0x30, 0x3d, 0xad, 0x19, 0x53, 0xc5, 0x2a, 0xfe, 0xeb, 0x82, 0xd2,
	0xec, 0x10, 0xc3, 0xbf, 0xc1, 0xba, 0xd9, 0x78, 0xd1, 0x6f, 0x74, 0xf6, 0x7a, 0xf6, 0xe0, 0x75,
	0xbf, 0x6d, 0x1f, 0xf6, 0x0e, 0xfa, 0x6d, 0xb3, 0xb3, 0xdb, 0x69, 0xb7, 0x2a, 0x19, 0x58, 0x04,
	0xf9, 0x56, 0x7b, 0xb7, 0x71, 0xd8, 0x1d, 0x54, 0x14, 0x58, 0x01, 0xa5, 0x97, 0xed, 0x83, 0x41,
	0xa7, 0xb7, 0x67, 0xf7, 0xf7, 0xf7, 0xbb, 0x95, 0x6c, 0xb3, 0x7b, 0x7e, 0xad, 0x29, 0x17, 0xd7,
	0x9a, 0xf2, 0xf5, 0x5a, 0x53, 0xce, 0x6e, 0xb4, 0xcc, 0xc5, 0x8d, 0x96, 0xf9, 0x7c, 0xa3, 0x65,
	0xde, 0x6c, 0xcf, 0x0a, 0xbd, 0xf3, 0x34, 0xc7, 0x6f, 0xb5, 0xf1, 0xee, 0xf6, 0x75, 0x96, 0xc2,
	0x87, 0x0b, 0xf2, 0xe2, 0x1f, 0x7f, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xff, 0x21, 0xd4, 0x5c, 0xcc,
	0x05, 0x00, 0x00,
}

func (m *Campaign) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Campaign) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Campaign) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CampaignCurrentAmount) > 0 {
		for iNdEx := len(m.CampaignCurrentAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CampaignCurrentAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCampaign(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0x8a
		}
	}
	if len(m.CampaignTotalAmount) > 0 {
		for iNdEx := len(m.CampaignTotalAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CampaignTotalAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCampaign(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0x82
		}
	}
	if len(m.VestingPoolName) > 0 {
		i -= len(m.VestingPoolName)
		copy(dAtA[i:], m.VestingPoolName)
		i = encodeVarintCampaign(dAtA, i, uint64(len(m.VestingPoolName)))
		i--
		dAtA[i] = 0x7a
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.VestingPeriod, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.VestingPeriod):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintCampaign(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x72
	n2, err2 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.LockupPeriod, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.LockupPeriod):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintCampaign(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x6a
	n3, err3 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.EndTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintCampaign(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x62
	n4, err4 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StartTime):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintCampaign(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x5a
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x50
	}
	{
		size := m.Free.Size()
		i -= size
		if _, err := m.Free.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCampaign(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.InitialClaimFreeAmount.Size()
		i -= size
		if _, err := m.InitialClaimFreeAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCampaign(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.FeegrantAmount.Size()
		i -= size
		if _, err := m.FeegrantAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCampaign(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.RemovableClaimRecords {
		i--
		if m.RemovableClaimRecords {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.CampaignType != 0 {
		i = encodeVarintCampaign(dAtA, i, uint64(m.CampaignType))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintCampaign(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintCampaign(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintCampaign(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintCampaign(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCampaign(dAtA []byte, offset int, v uint64) int {
	offset -= sovCampaign(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Campaign) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovCampaign(uint64(m.Id))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovCampaign(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCampaign(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovCampaign(uint64(l))
	}
	if m.CampaignType != 0 {
		n += 1 + sovCampaign(uint64(m.CampaignType))
	}
	if m.RemovableClaimRecords {
		n += 2
	}
	l = m.FeegrantAmount.Size()
	n += 1 + l + sovCampaign(uint64(l))
	l = m.InitialClaimFreeAmount.Size()
	n += 1 + l + sovCampaign(uint64(l))
	l = m.Free.Size()
	n += 1 + l + sovCampaign(uint64(l))
	if m.Enabled {
		n += 2
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovCampaign(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovCampaign(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.LockupPeriod)
	n += 1 + l + sovCampaign(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.VestingPeriod)
	n += 1 + l + sovCampaign(uint64(l))
	l = len(m.VestingPoolName)
	if l > 0 {
		n += 1 + l + sovCampaign(uint64(l))
	}
	if len(m.CampaignTotalAmount) > 0 {
		for _, e := range m.CampaignTotalAmount {
			l = e.Size()
			n += 2 + l + sovCampaign(uint64(l))
		}
	}
	if len(m.CampaignCurrentAmount) > 0 {
		for _, e := range m.CampaignCurrentAmount {
			l = e.Size()
			n += 2 + l + sovCampaign(uint64(l))
		}
	}
	return n
}

func sovCampaign(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCampaign(x uint64) (n int) {
	return sovCampaign(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Campaign) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCampaign
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Campaign: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Campaign: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCampaign
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCampaign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCampaign
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCampaign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCampaign
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCampaign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CampaignType", wireType)
			}
			m.CampaignType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CampaignType |= CampaignType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemovableClaimRecords", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.RemovableClaimRecords = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeegrantAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCampaign
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCampaign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FeegrantAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialClaimFreeAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCampaign
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCampaign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InitialClaimFreeAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Free", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCampaign
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCampaign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Free.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCampaign
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCampaign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCampaign
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCampaign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockupPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCampaign
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCampaign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.LockupPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCampaign
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCampaign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.VestingPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPoolName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCampaign
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCampaign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingPoolName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CampaignTotalAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCampaign
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCampaign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CampaignTotalAmount = append(m.CampaignTotalAmount, types.Coin{})
			if err := m.CampaignTotalAmount[len(m.CampaignTotalAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CampaignCurrentAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCampaign
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCampaign
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CampaignCurrentAmount = append(m.CampaignCurrentAmount, types.Coin{})
			if err := m.CampaignCurrentAmount[len(m.CampaignCurrentAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCampaign(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCampaign
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCampaign(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCampaign
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCampaign
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCampaign
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCampaign
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCampaign
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCampaign
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCampaign        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCampaign          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCampaign = fmt.Errorf("proto: unexpected end of group")
)
