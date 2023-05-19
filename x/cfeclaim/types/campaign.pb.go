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
	return fileDescriptor_3e0c61f8ccf41838, []int{0}
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
	VestingPeriod   time.Duration `protobuf:"bytes,14,opt,name=vesting_period,json=vestingPeriod,proto3,stdduration" json:"vesting_period"`
	VestingPoolName string        `protobuf:"bytes,15,opt,name=vestingPoolName,proto3" json:"vestingPoolName,omitempty"`
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

type CampaignCurrentAmount struct {
	CampaignId uint64                                   `protobuf:"varint,1,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
	Amount     github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *CampaignCurrentAmount) Reset()         { *m = CampaignCurrentAmount{} }
func (m *CampaignCurrentAmount) String() string { return proto.CompactTextString(m) }
func (*CampaignCurrentAmount) ProtoMessage()    {}
func (*CampaignCurrentAmount) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e0c61f8ccf41838, []int{2}
}
func (m *CampaignCurrentAmount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CampaignCurrentAmount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CampaignCurrentAmount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CampaignCurrentAmount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignCurrentAmount.Merge(m, src)
}
func (m *CampaignCurrentAmount) XXX_Size() int {
	return m.Size()
}
func (m *CampaignCurrentAmount) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignCurrentAmount.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignCurrentAmount proto.InternalMessageInfo

func (m *CampaignCurrentAmount) GetCampaignId() uint64 {
	if m != nil {
		return m.CampaignId
	}
	return 0
}

func (m *CampaignCurrentAmount) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

func init() {
	proto.RegisterEnum("chain4energy.c4echain.cfeclaim.CampaignType", CampaignType_name, CampaignType_value)
	proto.RegisterType((*Campaign)(nil), "chain4energy.c4echain.cfeclaim.Campaign")
	proto.RegisterType((*CampaignTotalAmount)(nil), "chain4energy.c4echain.cfeclaim.CampaignTotalAmount")
	proto.RegisterType((*CampaignCurrentAmount)(nil), "chain4energy.c4echain.cfeclaim.CampaignCurrentAmount")
}

func init() { proto.RegisterFile("c4echain/cfeclaim/campaign.proto", fileDescriptor_3e0c61f8ccf41838) }

var fileDescriptor_3e0c61f8ccf41838 = []byte{
	// 726 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0xcf, 0x4e, 0xdb, 0x4e,
	0x10, 0xc7, 0xe3, 0x10, 0x92, 0xb0, 0x09, 0x21, 0xda, 0x1f, 0xfc, 0x6a, 0x90, 0xea, 0x58, 0x1c,
	0xaa, 0xa8, 0x2a, 0x76, 0xa1, 0xa8, 0xd7, 0x2a, 0x71, 0x02, 0x4d, 0x95, 0x86, 0xc8, 0x84, 0x56,
	0xed, 0xc5, 0x72, 0xec, 0x89, 0x59, 0x11, 0x7b, 0x2d, 0x7b, 0x43, 0xcb, 0x5b, 0x70, 0xac, 0x54,
	0xf5, 0x05, 0xfa, 0x24, 0x1c, 0x39, 0xa2, 0x1e, 0xa0, 0x82, 0x17, 0xa9, 0xbc, 0xb6, 0xd3, 0x94,
	0x4a, 0x15, 0xf4, 0xd4, 0x53, 0xbc, 0xf3, 0xe7, 0x33, 0x93, 0xd9, 0xef, 0x2c, 0x92, 0xad, 0x6d,
	0xb0, 0x0e, 0x4d, 0xe2, 0xa9, 0xd6, 0x08, 0xac, 0xb1, 0x49, 0x5c, 0xd5, 0x32, 0x5d, 0xdf, 0x24,
	0x8e, 0xa7, 0xf8, 0x01, 0x65, 0x14, 0x4b, 0xdc, 0xbd, 0x0d, 0x1e, 0x04, 0xce, 0x89, 0x92, 0x86,
	0x2b, 0x69, 0xf8, 0xda, 0xb2, 0x43, 0x1d, 0xca, 0x43, 0xd5, 0xe8, 0x2b, 0xce, 0x5a, 0xab, 0x39,
	0x94, 0x3a, 0x63, 0x50, 0xf9, 0x69, 0x38, 0x19, 0xa9, 0x8c, 0xb8, 0x10, 0x32, 0xd3, 0xf5, 0x93,
	0x00, 0xe9, 0x76, 0x80, 0x3d, 0x09, 0x4c, 0x46, 0xa8, 0x97, 0xfa, 0x2d, 0x1a, 0xba, 0x34, 0x54,
	0x87, 0x66, 0x08, 0xea, 0xf1, 0xe6, 0x10, 0x98, 0xb9, 0xa9, 0x5a, 0x94, 0x24, 0xfe, 0xf5, 0x8b,
	0x3c, 0x2a, 0x6a, 0x49, 0xa7, 0xb8, 0x82, 0xb2, 0xc4, 0x16, 0x05, 0x59, 0xa8, 0xe7, 0xf4, 0x2c,
	0xb1, 0xf1, 0x32, 0x9a, 0xa7, 0x1f, 0x3c, 0x08, 0xc4, 0xac, 0x2c, 0xd4, 0x17, 0xf4, 0xf8, 0x80,
	0x31, 0xca, 0x79, 0xa6, 0x0b, 0xe2, 0x1c, 0x37, 0xf2, 0x6f, 0x2c, 0xa3, 0x92, 0x0d, 0xa1, 0x15,
	0x10, 0x3f, 0xaa, 0x2d, 0xe6, 0xb8, 0x6b, 0xd6, 0x84, 0xfb, 0xa8, 0x9c, 0x4e, 0x64, 0x70, 0xe2,
	0x83, 0x38, 0x2f, 0x0b, 0xf5, 0xca, 0xd6, 0x13, 0xe5, 0xcf, 0x63, 0x51, 0xb4, 0x99, 0x1c, 0xfd,
	0x17, 0x02, 0x7e, 0x8e, 0x1e, 0x04, 0xe0, 0xd2, 0x63, 0x73, 0x38, 0x06, 0x83, 0x47, 0x1b, 0x01,
	0x58, 0x34, 0xb0, 0x43, 0x31, 0x2f, 0x0b, 0xf5, 0xa2, 0xbe, 0x32, 0x75, 0x6b, 0x91, 0x57, 0x8f,
	0x9d, 0xf8, 0x2d, 0x5a, 0x1a, 0x01, 0x38, 0x81, 0xe9, 0x31, 0xc3, 0x74, 0xe9, 0xc4, 0x63, 0x62,
	0x21, 0xea, 0xb7, 0xa9, 0x9c, 0x5d, 0xd6, 0x32, 0xdf, 0x2e, 0x6b, 0x8f, 0x1c, 0xc2, 0x0e, 0x27,
	0x43, 0xc5, 0xa2, 0xae, 0x9a, 0x8c, 0x2f, 0xfe, 0xd9, 0x08, 0xed, 0x23, 0x95, 0x9d, 0xf8, 0x10,
	0x2a, 0x1d, 0x8f, 0xe9, 0x95, 0x14, 0xd3, 0xe0, 0x14, 0x4c, 0xd0, 0x2a, 0xf1, 0x08, 0x23, 0xe6,
	0x38, 0x69, 0x67, 0x14, 0x00, 0xa4, 0x25, 0x8a, 0x7f, 0x55, 0xe2, 0xff, 0x04, 0xc8, 0xff, 0xc0,
	0x4e, 0x00, 0x90, 0x94, 0x6a, 0xa2, 0x5c, 0x04, 0x17, 0x17, 0xee, 0x4d, 0x6d, 0x81, 0xa5, 0xf3,
	0x5c, 0x2c, 0xa2, 0x02, 0x78, 0xd1, 0x74, 0x6c, 0x11, 0xf1, 0x79, 0xa5, 0x47, 0xac, 0x21, 0x14,
	0x32, 0x33, 0x60, 0x46, 0xa4, 0x36, 0xb1, 0x24, 0x0b, 0xf5, 0xd2, 0xd6, 0x9a, 0x12, 0x2b, 0x4d,
	0x49, 0x95, 0xa6, 0x0c, 0x52, 0x29, 0x36, 0x8b, 0x51, 0xfd, 0xd3, 0xab, 0x9a, 0xa0, 0x2f, 0xf0,
	0xbc, 0xc8, 0x83, 0x5f, 0xa0, 0x22, 0x78, 0x76, 0x8c, 0x28, 0xdf, 0x03, 0x51, 0x00, 0xcf, 0xe6,
	0x80, 0x97, 0x68, 0x71, 0x4c, 0xad, 0xa3, 0x89, 0x6f, 0xf8, 0x10, 0x10, 0x6a, 0x8b, 0x8b, 0x9c,
	0xb2, 0xfa, 0x1b, 0xa5, 0x95, 0x48, 0x3e, 0x86, 0x7c, 0x8a, 0x20, 0xe5, 0x38, 0xb3, 0xcf, 0x13,
	0xf1, 0x2b, 0x54, 0x39, 0x86, 0x90, 0x11, 0xcf, 0x49, 0x51, 0x95, 0xbb, 0xa3, 0x16, 0x93, 0xd4,
	0x84, 0x55, 0x47, 0x4b, 0xa9, 0x81, 0xd2, 0x71, 0x2f, 0x5a, 0x84, 0x25, 0xae, 0xf6, 0xdb, 0xe6,
	0xf5, 0xcf, 0x02, 0xfa, 0x6f, 0x2a, 0x5f, 0xca, 0xcc, 0x71, 0x72, 0x77, 0x35, 0x54, 0x4a, 0x75,
	0x6c, 0x4c, 0xd7, 0x0d, 0xa5, 0xa6, 0x8e, 0x8d, 0x2d, 0x94, 0x4f, 0x44, 0x93, 0x95, 0xe7, 0x78,
	0x9b, 0xf1, 0x2d, 0x2a, 0xd1, 0x12, 0x2b, 0xc9, 0x12, 0x2b, 0x1a, 0x25, 0x5e, 0xf3, 0x69, 0xd4,
	0xe6, 0xd7, 0xab, 0x5a, 0xfd, 0x0e, 0x37, 0x1f, 0x25, 0x84, 0x7a, 0x82, 0x5e, 0xff, 0x22, 0xa0,
	0x95, 0xb4, 0x3b, 0x6d, 0x12, 0x04, 0x30, 0x95, 0xf1, 0x3f, 0xd1, 0xdf, 0xe3, 0x2e, 0x2a, 0xcf,
	0xee, 0x3e, 0x7e, 0x88, 0x56, 0xb5, 0xc6, 0xeb, 0x7e, 0xa3, 0xb3, 0xdb, 0x33, 0x06, 0xef, 0xfa,
	0x6d, 0xe3, 0xa0, 0xb7, 0xdf, 0x6f, 0x6b, 0x9d, 0x9d, 0x4e, 0xbb, 0x55, 0xcd, 0xe0, 0x12, 0x2a,
	0xb4, 0xda, 0x3b, 0x8d, 0x83, 0xee, 0xa0, 0x2a, 0xe0, 0x2a, 0x2a, 0xbf, 0x69, 0xef, 0x0f, 0x3a,
	0xbd, 0x5d, 0xa3, 0xbf, 0xb7, 0xd7, 0xad, 0x66, 0x9b, 0xdd, 0xb3, 0x6b, 0x49, 0x38, 0xbf, 0x96,
	0x84, 0xef, 0xd7, 0x92, 0x70, 0x7a, 0x23, 0x65, 0xce, 0x6f, 0xa4, 0xcc, 0xc5, 0x8d, 0x94, 0x79,
	0xbf, 0x35, 0xdb, 0xd9, 0xcc, 0x5b, 0xa4, 0x5a, 0xdb, 0xb0, 0x11, 0x3f, 0xe9, 0x1f, 0x7f, 0x3e,
	0xea, 0xbc, 0xd3, 0x61, 0x9e, 0xeb, 0xe5, 0xd9, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x05, 0xe4,
	0x28, 0x53, 0xf6, 0x05, 0x00, 0x00,
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
		dAtA[i] = 0x7a
	}
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.VestingPeriod, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.VestingPeriod):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintCampaign(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x72
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.LockupPeriod, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.LockupPeriod):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintCampaign(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x6a
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintCampaign(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x62
	n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
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

func (m *CampaignCurrentAmount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CampaignCurrentAmount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CampaignCurrentAmount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *CampaignCurrentAmount) Size() (n int) {
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
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
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
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
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
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.LockupPeriod, dAtA[iNdEx:postIndex]); err != nil {
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
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.VestingPeriod, dAtA[iNdEx:postIndex]); err != nil {
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
func (m *CampaignCurrentAmount) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: CampaignCurrentAmount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CampaignCurrentAmount: illegal tag %d (wire type %d)", fieldNum, wire)
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
