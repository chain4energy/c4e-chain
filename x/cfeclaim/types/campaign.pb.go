// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: c4echain/cfeclaim/campaign.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
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
	CampaignType_DYNAMIC                   CampaignType = 2
	CampaignType_VESTING_POOL              CampaignType = 3
)

var CampaignType_name = map[int32]string{
	0: "CAMPAIGN_TYPE_UNSPECIFIED",
	1: "DEFAULT",
	2: "DYNAMIC",
	3: "VESTING_POOL",
}

var CampaignType_value = map[string]int32{
	"CAMPAIGN_TYPE_UNSPECIFIED": 0,
	"DEFAULT":                   1,
	"DYNAMIC":                   2,
	"VESTING_POOL":              3,
}

func (x CampaignType) String() string {
	return proto.EnumName(CampaignType_name, int32(x))
}

func (CampaignType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3e0c61f8ccf41838, []int{0}
}

type Campaign struct {
	Id                     uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner                  string                                 `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Name                   string                                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description            string                                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	CampaignType           CampaignType                           `protobuf:"varint,5,opt,name=campaignType,proto3,enum=chain4energy.c4echain.cfeclaim.CampaignType" json:"campaignType,omitempty"`
	FeegrantAmount         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=feegrant_amount,json=feegrantAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"feegrant_amount"`
	InitialClaimFreeAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=initial_claim_free_amount,json=initialClaimFreeAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"initial_claim_free_amount"`
	Free                   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=free,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"free"`
	Enabled                bool                                   `protobuf:"varint,9,opt,name=enabled,proto3" json:"enabled,omitempty"`
	StartTime              time.Time                              `protobuf:"bytes,10,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time"`
	EndTime                time.Time                              `protobuf:"bytes,11,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time"`
	// period of locked coins from claim
	LockupPeriod time.Duration `protobuf:"bytes,12,opt,name=lockup_period,json=lockupPeriod,proto3,stdduration" json:"lockup_period"`
	// period of vesting coins after lockup period
	VestingPeriod   time.Duration `protobuf:"bytes,13,opt,name=vesting_period,json=vestingPeriod,proto3,stdduration" json:"vesting_period"`
	VestingPoolName string        `protobuf:"bytes,14,opt,name=vestingPoolName,proto3" json:"vestingPoolName,omitempty"`
}

func (m *Campaign) Reset()         { *m = Campaign{} }
func (m *Campaign) String() string { return proto.CompactTextString(m) }
func (*Campaign) ProtoMessage()    {}
func (*Campaign) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e0c61f8ccf41838, []int{0}
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

type CampaignTotalAmount struct {
	CampaignId uint64                                   `protobuf:"varint,1,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
	Amount     github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *CampaignTotalAmount) Reset()         { *m = CampaignTotalAmount{} }
func (m *CampaignTotalAmount) String() string { return proto.CompactTextString(m) }
func (*CampaignTotalAmount) ProtoMessage()    {}
func (*CampaignTotalAmount) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e0c61f8ccf41838, []int{1}
}
func (m *CampaignTotalAmount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CampaignTotalAmount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CampaignTotalAmount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CampaignTotalAmount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignTotalAmount.Merge(m, src)
}
func (m *CampaignTotalAmount) XXX_Size() int {
	return m.Size()
}
func (m *CampaignTotalAmount) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignTotalAmount.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignTotalAmount proto.InternalMessageInfo

func (m *CampaignTotalAmount) GetCampaignId() uint64 {
	if m != nil {
		return m.CampaignId
	}
	return 0
}

func (m *CampaignTotalAmount) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

type CampaignAmountLeft struct {
	CampaignId uint64                                   `protobuf:"varint,1,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
	Amount     github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *CampaignAmountLeft) Reset()         { *m = CampaignAmountLeft{} }
func (m *CampaignAmountLeft) String() string { return proto.CompactTextString(m) }
func (*CampaignAmountLeft) ProtoMessage()    {}
func (*CampaignAmountLeft) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e0c61f8ccf41838, []int{2}
}
func (m *CampaignAmountLeft) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CampaignAmountLeft) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CampaignAmountLeft.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CampaignAmountLeft) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignAmountLeft.Merge(m, src)
}
func (m *CampaignAmountLeft) XXX_Size() int {
	return m.Size()
}
func (m *CampaignAmountLeft) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignAmountLeft.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignAmountLeft proto.InternalMessageInfo

func (m *CampaignAmountLeft) GetCampaignId() uint64 {
	if m != nil {
		return m.CampaignId
	}
	return 0
}

func (m *CampaignAmountLeft) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

func init() {
	proto.RegisterEnum("chain4energy.c4echain.cfeclaim.CampaignType", CampaignType_name, CampaignType_value)
	proto.RegisterType((*Campaign)(nil), "chain4energy.c4echain.cfeclaim.Campaign")
	proto.RegisterType((*CampaignTotalAmount)(nil), "chain4energy.c4echain.cfeclaim.CampaignTotalAmount")
	proto.RegisterType((*CampaignAmountLeft)(nil), "chain4energy.c4echain.cfeclaim.CampaignAmountLeft")
}

func init() { proto.RegisterFile("c4echain/cfeclaim/campaign.proto", fileDescriptor_3e0c61f8ccf41838) }

var fileDescriptor_3e0c61f8ccf41838 = []byte{
	// 703 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0x4f, 0x4f, 0xdb, 0x3c,
	0x1c, 0xc7, 0x9b, 0x52, 0xda, 0xe2, 0x96, 0x52, 0xf9, 0x41, 0x8f, 0x02, 0xd2, 0xd2, 0x88, 0xc3,
	0x54, 0x4d, 0x23, 0x19, 0x8c, 0xfb, 0xd4, 0xa6, 0x85, 0x75, 0x2a, 0xa5, 0x0a, 0x65, 0x13, 0xbb,
	0x44, 0x6e, 0xe2, 0x06, 0x8b, 0xc4, 0x8e, 0x12, 0x97, 0x8d, 0x77, 0xc1, 0x71, 0xda, 0x2e, 0x3b,
	0xef, 0x95, 0x70, 0xe4, 0x38, 0xed, 0x00, 0x13, 0xbc, 0x91, 0x29, 0x4e, 0xc2, 0x3a, 0x26, 0x4d,
	0xb0, 0xd3, 0x4e, 0x89, 0x7f, 0x7f, 0x3e, 0xbf, 0xaf, 0xec, 0xaf, 0x0d, 0x54, 0x7b, 0x0b, 0xdb,
	0x47, 0x88, 0x50, 0xdd, 0x9e, 0x60, 0xdb, 0x43, 0xc4, 0xd7, 0x6d, 0xe4, 0x07, 0x88, 0xb8, 0x54,
	0x0b, 0x42, 0xc6, 0x19, 0x54, 0x44, 0x7a, 0x0b, 0x53, 0x1c, 0xba, 0xa7, 0x5a, 0x56, 0xae, 0x65,
	0xe5, 0xab, 0xcb, 0x2e, 0x73, 0x99, 0x28, 0xd5, 0xe3, 0xbf, 0xa4, 0x6b, 0xb5, 0xe1, 0x32, 0xe6,
	0x7a, 0x58, 0x17, 0xab, 0xf1, 0x74, 0xa2, 0x73, 0xe2, 0xe3, 0x88, 0x23, 0x3f, 0x48, 0x0b, 0x94,
	0xbb, 0x05, 0xce, 0x34, 0x44, 0x9c, 0x30, 0x9a, 0xe5, 0x6d, 0x16, 0xf9, 0x2c, 0xd2, 0xc7, 0x28,
	0xc2, 0xfa, 0xc9, 0xc6, 0x18, 0x73, 0xb4, 0xa1, 0xdb, 0x8c, 0xa4, 0xf9, 0xb5, 0xcf, 0x45, 0x50,
	0x36, 0x52, 0xa5, 0xb0, 0x06, 0xf2, 0xc4, 0x91, 0x25, 0x55, 0x6a, 0x16, 0xcc, 0x3c, 0x71, 0xe0,
	0x32, 0x98, 0x67, 0xef, 0x28, 0x0e, 0xe5, 0xbc, 0x2a, 0x35, 0x17, 0xcc, 0x64, 0x01, 0x21, 0x28,
	0x50, 0xe4, 0x63, 0x79, 0x4e, 0x04, 0xc5, 0x3f, 0x54, 0x41, 0xc5, 0xc1, 0x91, 0x1d, 0x92, 0x20,
	0x9e, 0x2d, 0x17, 0x44, 0x6a, 0x36, 0x04, 0x87, 0xa0, 0x9a, 0xed, 0xc8, 0xe8, 0x34, 0xc0, 0xf2,
	0xbc, 0x2a, 0x35, 0x6b, 0x9b, 0x4f, 0xb5, 0x3f, 0x6f, 0x8b, 0x66, 0xcc, 0xf4, 0x98, 0xbf, 0x10,
	0xe0, 0x1b, 0xb0, 0x34, 0xc1, 0xd8, 0x0d, 0x11, 0xe5, 0x16, 0xf2, 0xd9, 0x94, 0x72, 0xb9, 0x18,
	0xcf, 0x6d, 0x6b, 0xe7, 0x97, 0x8d, 0xdc, 0xb7, 0xcb, 0xc6, 0x63, 0x97, 0xf0, 0xa3, 0xe9, 0x58,
	0xb3, 0x99, 0xaf, 0xa7, 0xdb, 0x90, 0x7c, 0xd6, 0x23, 0xe7, 0x58, 0xe7, 0xa7, 0x01, 0x8e, 0xb4,
	0x1e, 0xe5, 0x66, 0x2d, 0xc3, 0xb4, 0x04, 0x05, 0x12, 0xb0, 0x42, 0x28, 0xe1, 0x04, 0x79, 0x96,
	0x10, 0x61, 0x4d, 0x42, 0x8c, 0xb3, 0x11, 0xa5, 0xbf, 0x1a, 0xf1, 0x7f, 0x0a, 0x34, 0x62, 0xde,
	0x76, 0x88, 0x71, 0x3a, 0xaa, 0x0d, 0x0a, 0x31, 0x5c, 0x2e, 0x3f, 0x98, 0xda, 0xc1, 0xb6, 0x29,
	0x7a, 0xa1, 0x0c, 0x4a, 0x98, 0xa2, 0xb1, 0x87, 0x1d, 0x79, 0x41, 0x95, 0x9a, 0x65, 0x33, 0x5b,
	0x42, 0x03, 0x80, 0x88, 0xa3, 0x90, 0x5b, 0xb1, 0x6b, 0x64, 0xa0, 0x4a, 0xcd, 0xca, 0xe6, 0xaa,
	0x96, 0x38, 0x46, 0xcb, 0x1c, 0xa3, 0x8d, 0x32, 0x4b, 0xb5, 0xcb, 0xf1, 0xfc, 0xb3, 0xab, 0x86,
	0x64, 0x2e, 0x88, 0xbe, 0x38, 0x03, 0x5f, 0x80, 0x32, 0xa6, 0x4e, 0x82, 0xa8, 0x3c, 0x00, 0x51,
	0xc2, 0xd4, 0x11, 0x80, 0x97, 0x60, 0xd1, 0x63, 0xf6, 0xf1, 0x34, 0xb0, 0x02, 0x1c, 0x12, 0xe6,
	0xc8, 0x55, 0x41, 0x59, 0xf9, 0x8d, 0xd2, 0x49, 0xad, 0x9b, 0x40, 0x3e, 0xc4, 0x90, 0x6a, 0xd2,
	0x39, 0x14, 0x8d, 0xf0, 0x15, 0xa8, 0x9d, 0xe0, 0x88, 0x13, 0xea, 0x66, 0xa8, 0xc5, 0xfb, 0xa3,
	0x16, 0xd3, 0xd6, 0x94, 0xd5, 0x04, 0x4b, 0x59, 0x80, 0x31, 0x6f, 0x10, 0x1b, 0xba, 0x26, 0x5c,
	0x7b, 0x37, 0xbc, 0xf6, 0x49, 0x02, 0xff, 0xdd, 0xda, 0x90, 0x71, 0xe4, 0xa5, 0x67, 0xd7, 0x00,
	0x95, 0xcc, 0x8f, 0xd6, 0xed, 0xb5, 0x01, 0x59, 0xa8, 0xe7, 0x40, 0x1b, 0x14, 0x53, 0xd3, 0xe4,
	0xd5, 0x39, 0x21, 0x33, 0x39, 0x45, 0x2d, 0xbe, 0x8c, 0x5a, 0x7a, 0x19, 0x35, 0x83, 0x11, 0xda,
	0x7e, 0x16, 0xcb, 0xfc, 0x72, 0xd5, 0x68, 0xde, 0xe3, 0xe4, 0xe3, 0x86, 0xc8, 0x4c, 0xd1, 0x6b,
	0x1f, 0x25, 0x00, 0x33, 0x75, 0x89, 0xb0, 0x3e, 0x9e, 0xfc, 0x23, 0xe2, 0x9e, 0x1c, 0x82, 0xea,
	0xec, 0x05, 0x86, 0x8f, 0xc0, 0x8a, 0xd1, 0xda, 0x1d, 0xb6, 0x7a, 0x3b, 0x03, 0x6b, 0x74, 0x38,
	0xec, 0x5a, 0x07, 0x83, 0xfd, 0x61, 0xd7, 0xe8, 0x6d, 0xf7, 0xba, 0x9d, 0x7a, 0x0e, 0x56, 0x40,
	0xa9, 0xd3, 0xdd, 0x6e, 0x1d, 0xf4, 0x47, 0x75, 0x49, 0x2c, 0x0e, 0x07, 0xad, 0xdd, 0x9e, 0x51,
	0xcf, 0xc3, 0x3a, 0xa8, 0xbe, 0xee, 0xee, 0x8f, 0x7a, 0x83, 0x1d, 0x6b, 0xb8, 0xb7, 0xd7, 0xaf,
	0xcf, 0xb5, 0xfb, 0xe7, 0xd7, 0x8a, 0x74, 0x71, 0xad, 0x48, 0xdf, 0xaf, 0x15, 0xe9, 0xec, 0x46,
	0xc9, 0x5d, 0xdc, 0x28, 0xb9, 0xaf, 0x37, 0x4a, 0xee, 0xed, 0xe6, 0xac, 0xcc, 0x99, 0xd7, 0x45,
	0xb7, 0xb7, 0xf0, 0x7a, 0xf2, 0x48, 0xbf, 0xff, 0xf9, 0x4c, 0x0b, 0xd9, 0xe3, 0xa2, 0x70, 0xce,
	0xf3, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x96, 0xec, 0xc8, 0xc8, 0x05, 0x00, 0x00,
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
	if len(m.VestingPoolName) > 0 {
		i -= len(m.VestingPoolName)
		copy(dAtA[i:], m.VestingPoolName)
		i = encodeVarintCampaign(dAtA, i, uint64(len(m.VestingPoolName)))
		i--
		dAtA[i] = 0x72
	}
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.VestingPeriod, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.VestingPeriod):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintCampaign(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x6a
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.LockupPeriod, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.LockupPeriod):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintCampaign(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x62
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintCampaign(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x5a
	n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintCampaign(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x52
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
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
	dAtA[i] = 0x42
	{
		size := m.InitialClaimFreeAmount.Size()
		i -= size
		if _, err := m.InitialClaimFreeAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCampaign(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.FeegrantAmount.Size()
		i -= size
		if _, err := m.FeegrantAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCampaign(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
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

func (m *CampaignTotalAmount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CampaignTotalAmount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CampaignTotalAmount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCampaign(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.CampaignId != 0 {
		i = encodeVarintCampaign(dAtA, i, uint64(m.CampaignId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CampaignAmountLeft) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CampaignAmountLeft) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CampaignAmountLeft) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCampaign(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.CampaignId != 0 {
		i = encodeVarintCampaign(dAtA, i, uint64(m.CampaignId))
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
	l = m.FeegrantAmount.Size()
	n += 1 + l + sovCampaign(uint64(l))
	l = m.InitialClaimFreeAmount.Size()
	n += 1 + l + sovCampaign(uint64(l))
	l = m.Free.Size()
	n += 1 + l + sovCampaign(uint64(l))
	if m.Enabled {
		n += 2
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovCampaign(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovCampaign(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.LockupPeriod)
	n += 1 + l + sovCampaign(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.VestingPeriod)
	n += 1 + l + sovCampaign(uint64(l))
	l = len(m.VestingPoolName)
	if l > 0 {
		n += 1 + l + sovCampaign(uint64(l))
	}
	return n
}

func (m *CampaignTotalAmount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CampaignId != 0 {
		n += 1 + sovCampaign(uint64(m.CampaignId))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovCampaign(uint64(l))
		}
	}
	return n
}

func (m *CampaignAmountLeft) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CampaignId != 0 {
		n += 1 + sovCampaign(uint64(m.CampaignId))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovCampaign(uint64(l))
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
		case 7:
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
		case 8:
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
		case 9:
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
		case 10:
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
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
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
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
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
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.LockupPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
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
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.VestingPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
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
func (m *CampaignTotalAmount) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: CampaignTotalAmount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CampaignTotalAmount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CampaignId", wireType)
			}
			m.CampaignId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CampaignId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *CampaignAmountLeft) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: CampaignAmountLeft: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CampaignAmountLeft: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CampaignId", wireType)
			}
			m.CampaignId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCampaign
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CampaignId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
