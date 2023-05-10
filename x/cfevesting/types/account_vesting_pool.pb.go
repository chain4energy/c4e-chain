// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: c4echain/cfevesting/account_vesting_pool.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
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

type AccountVestingPools struct {
	Owner        string         `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	VestingPools []*VestingPool `protobuf:"bytes,2,rep,name=vesting_pools,json=vestingPools,proto3" json:"vesting_pools,omitempty"`
}

func (m *AccountVestingPools) Reset()         { *m = AccountVestingPools{} }
func (m *AccountVestingPools) String() string { return proto.CompactTextString(m) }
func (*AccountVestingPools) ProtoMessage()    {}
func (*AccountVestingPools) Descriptor() ([]byte, []int) {
	return fileDescriptor_b42467bcdc7cf50e, []int{0}
}
func (m *AccountVestingPools) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountVestingPools) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountVestingPools.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountVestingPools) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountVestingPools.Merge(m, src)
}
func (m *AccountVestingPools) XXX_Size() int {
	return m.Size()
}
func (m *AccountVestingPools) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountVestingPools.DiscardUnknown(m)
}

var xxx_messageInfo_AccountVestingPools proto.InternalMessageInfo

func (m *AccountVestingPools) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *AccountVestingPools) GetVestingPools() []*VestingPool {
	if m != nil {
		return m.VestingPools
	}
	return nil
}

type VestingPool struct {
	Name            string                                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	VestingType     string                                 `protobuf:"bytes,2,opt,name=vesting_type,json=vestingType,proto3" json:"vesting_type,omitempty"`
	LockStart       time.Time                              `protobuf:"bytes,3,opt,name=lock_start,json=lockStart,proto3,stdtime" json:"lock_start"`
	LockEnd         time.Time                              `protobuf:"bytes,4,opt,name=lock_end,json=lockEnd,proto3,stdtime" json:"lock_end"`
	InitiallyLocked github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=initially_locked,json=initiallyLocked,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"initially_locked"`
	Withdrawn       github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=withdrawn,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"withdrawn"`
	Sent            github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=sent,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"sent"`
	GenesisPool     bool                                   `protobuf:"varint,8,opt,name=genesis_pool,json=genesisPool,proto3" json:"genesis_pool,omitempty"`
	Reservations    []*VestingPoolReservation              `protobuf:"bytes,9,rep,name=reservations,proto3" json:"reservations,omitempty"`
}

func (m *VestingPool) Reset()         { *m = VestingPool{} }
func (m *VestingPool) String() string { return proto.CompactTextString(m) }
func (*VestingPool) ProtoMessage()    {}
func (*VestingPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_b42467bcdc7cf50e, []int{1}
}
func (m *VestingPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VestingPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VestingPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VestingPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VestingPool.Merge(m, src)
}
func (m *VestingPool) XXX_Size() int {
	return m.Size()
}
func (m *VestingPool) XXX_DiscardUnknown() {
	xxx_messageInfo_VestingPool.DiscardUnknown(m)
}

var xxx_messageInfo_VestingPool proto.InternalMessageInfo

func (m *VestingPool) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VestingPool) GetVestingType() string {
	if m != nil {
		return m.VestingType
	}
	return ""
}

func (m *VestingPool) GetLockStart() time.Time {
	if m != nil {
		return m.LockStart
	}
	return time.Time{}
}

func (m *VestingPool) GetLockEnd() time.Time {
	if m != nil {
		return m.LockEnd
	}
	return time.Time{}
}

func (m *VestingPool) GetGenesisPool() bool {
	if m != nil {
		return m.GenesisPool
	}
	return false
}

func (m *VestingPool) GetReservations() []*VestingPoolReservation {
	if m != nil {
		return m.Reservations
	}
	return nil
}

type VestingPoolReservation struct {
	Id     uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
}

func (m *VestingPoolReservation) Reset()         { *m = VestingPoolReservation{} }
func (m *VestingPoolReservation) String() string { return proto.CompactTextString(m) }
func (*VestingPoolReservation) ProtoMessage()    {}
func (*VestingPoolReservation) Descriptor() ([]byte, []int) {
	return fileDescriptor_b42467bcdc7cf50e, []int{2}
}
func (m *VestingPoolReservation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VestingPoolReservation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VestingPoolReservation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VestingPoolReservation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VestingPoolReservation.Merge(m, src)
}
func (m *VestingPoolReservation) XXX_Size() int {
	return m.Size()
}
func (m *VestingPoolReservation) XXX_DiscardUnknown() {
	xxx_messageInfo_VestingPoolReservation.DiscardUnknown(m)
}

var xxx_messageInfo_VestingPoolReservation proto.InternalMessageInfo

func (m *VestingPoolReservation) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*AccountVestingPools)(nil), "chain4energy.c4echain.cfevesting.AccountVestingPools")
	proto.RegisterType((*VestingPool)(nil), "chain4energy.c4echain.cfevesting.VestingPool")
	proto.RegisterType((*VestingPoolReservation)(nil), "chain4energy.c4echain.cfevesting.VestingPoolReservation")
}

func init() {
	proto.RegisterFile("c4echain/cfevesting/account_vesting_pool.proto", fileDescriptor_b42467bcdc7cf50e)
}

var fileDescriptor_b42467bcdc7cf50e = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x4d, 0x6b, 0xdb, 0x30,
	0x18, 0xc7, 0xe3, 0x34, 0x4d, 0x13, 0xa5, 0x7b, 0x41, 0x2b, 0xc3, 0xe4, 0xe0, 0x64, 0x39, 0x8c,
	0x5c, 0x22, 0xd1, 0x2e, 0x87, 0xdd, 0xc6, 0x32, 0x36, 0x18, 0x94, 0x31, 0xbc, 0x32, 0xd8, 0x18,
	0x04, 0xc5, 0x7e, 0xea, 0x88, 0xda, 0x92, 0xb1, 0x94, 0x64, 0x39, 0xed, 0x2b, 0xf4, 0xe3, 0xec,
	0x23, 0xf4, 0xd8, 0xe3, 0xd8, 0xa1, 0x1b, 0xc9, 0x17, 0x19, 0x92, 0x9d, 0xc6, 0x83, 0xc1, 0x48,
	0x4f, 0xb6, 0xa4, 0xe7, 0xff, 0xd3, 0xf3, 0xf2, 0x47, 0x88, 0x04, 0x43, 0x08, 0xa6, 0x8c, 0x0b,
	0x1a, 0x9c, 0xc3, 0x1c, 0x94, 0xe6, 0x22, 0xa2, 0x2c, 0x08, 0xe4, 0x4c, 0xe8, 0x71, 0xb1, 0x1e,
	0xa7, 0x52, 0xc6, 0x24, 0xcd, 0xa4, 0x96, 0xb8, 0x6b, 0x83, 0x87, 0x20, 0x20, 0x8b, 0x96, 0xb7,
	0x62, 0xb2, 0x15, 0xb7, 0x8f, 0x22, 0x19, 0x49, 0x1b, 0x4c, 0xcd, 0x5f, 0xae, 0x6b, 0x77, 0x22,
	0x29, 0xa3, 0x18, 0xa8, 0x5d, 0x4d, 0x66, 0xe7, 0x54, 0xf3, 0x04, 0x94, 0x66, 0x49, 0x5a, 0x04,
	0x78, 0x81, 0x54, 0x89, 0x54, 0x74, 0xc2, 0x14, 0xd0, 0xf9, 0xf1, 0x04, 0x34, 0x3b, 0xa6, 0x81,
	0xe4, 0x22, 0x3f, 0xef, 0x7d, 0x43, 0x8f, 0x5e, 0xe6, 0x69, 0x7d, 0xcc, 0x2f, 0x7a, 0x2f, 0x65,
	0xac, 0xf0, 0x11, 0xda, 0x97, 0x0b, 0x01, 0x99, 0xeb, 0x74, 0x9d, 0x7e, 0xd3, 0xcf, 0x17, 0xd8,
	0x47, 0xf7, 0xca, 0xb9, 0x2b, 0xb7, 0xda, 0xdd, 0xeb, 0xb7, 0x4e, 0x06, 0xe4, 0x7f, 0xd9, 0x93,
	0x12, 0xdc, 0x3f, 0x9c, 0x97, 0x6e, 0xea, 0x7d, 0xaf, 0xa1, 0x56, 0xe9, 0x14, 0x63, 0x54, 0x13,
	0x2c, 0x81, 0xe2, 0x62, 0xfb, 0x8f, 0x9f, 0xa0, 0x8d, 0x66, 0xac, 0x97, 0x29, 0xb8, 0x55, 0x7b,
	0xd6, 0x2a, 0xf6, 0xce, 0x96, 0x29, 0xe0, 0x57, 0x08, 0xc5, 0x32, 0xb8, 0x18, 0x2b, 0xcd, 0x32,
	0xed, 0xee, 0x75, 0x9d, 0x7e, 0xeb, 0xa4, 0x4d, 0xf2, 0xee, 0x90, 0x4d, 0x77, 0xc8, 0xd9, 0xa6,
	0x3b, 0xa3, 0xc6, 0xd5, 0x4d, 0xa7, 0x72, 0xf9, 0xab, 0xe3, 0xf8, 0x4d, 0xa3, 0xfb, 0x60, 0x64,
	0xf8, 0x05, 0x6a, 0x58, 0x08, 0x88, 0xd0, 0xad, 0xed, 0x80, 0x38, 0x30, 0xaa, 0xd7, 0x22, 0xc4,
	0x9f, 0xd0, 0x43, 0x2e, 0xb8, 0xe6, 0x2c, 0x8e, 0x97, 0x63, 0xb3, 0x09, 0xa1, 0xbb, 0x6f, 0x92,
	0x1d, 0x11, 0x13, 0xfc, 0xf3, 0xa6, 0xf3, 0x34, 0xe2, 0x7a, 0x3a, 0x9b, 0x90, 0x40, 0x26, 0xb4,
	0x18, 0x4d, 0xfe, 0x19, 0xa8, 0xf0, 0x82, 0x9a, 0xea, 0x14, 0x79, 0x2b, 0xb4, 0xff, 0xe0, 0x96,
	0x73, 0x6a, 0x31, 0xf8, 0x14, 0x35, 0x17, 0x5c, 0x4f, 0xc3, 0x8c, 0x2d, 0x84, 0x5b, 0xbf, 0x13,
	0x73, 0x0b, 0xc0, 0x23, 0x54, 0x53, 0x20, 0xb4, 0x7b, 0x70, 0x27, 0x90, 0xd5, 0x9a, 0xa9, 0x44,
	0x20, 0x40, 0x71, 0x65, 0xdd, 0xe0, 0x36, 0xba, 0x4e, 0xbf, 0xe1, 0xb7, 0x8a, 0x3d, 0x3b, 0xcc,
	0x2f, 0xe8, 0x30, 0x03, 0x05, 0xd9, 0x9c, 0x69, 0x2e, 0x85, 0x72, 0x9b, 0xd6, 0x2f, 0xcf, 0x77,
	0xf3, 0xcb, 0x16, 0xe0, 0xff, 0x45, 0xeb, 0xa5, 0xe8, 0xf1, 0xbf, 0xe3, 0xf0, 0x7d, 0x54, 0xe5,
	0xa1, 0xb5, 0x50, 0xcd, 0xaf, 0xf2, 0x10, 0xbf, 0x41, 0x75, 0x96, 0x18, 0x93, 0xe7, 0xd6, 0xd9,
	0xb9, 0xe0, 0x42, 0x3d, 0x7a, 0x77, 0xb5, 0xf2, 0x9c, 0xeb, 0x95, 0xe7, 0xfc, 0x5e, 0x79, 0xce,
	0xe5, 0xda, 0xab, 0x5c, 0xaf, 0xbd, 0xca, 0x8f, 0xb5, 0x57, 0xf9, 0x3c, 0x2c, 0x93, 0x4a, 0xd5,
	0xd1, 0x60, 0x08, 0x83, 0xfc, 0x25, 0xf8, 0x5a, 0x7e, 0x0b, 0x2c, 0x7b, 0x52, 0xb7, 0xb6, 0x7a,
	0xf6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x14, 0x89, 0x75, 0x1c, 0x2f, 0x04, 0x00, 0x00,
}

func (m *AccountVestingPools) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountVestingPools) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountVestingPools) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VestingPools) > 0 {
		for iNdEx := len(m.VestingPools) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VestingPools[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAccountVestingPool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintAccountVestingPool(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *VestingPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VestingPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VestingPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Reservations) > 0 {
		for iNdEx := len(m.Reservations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Reservations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAccountVestingPool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.GenesisPool {
		i--
		if m.GenesisPool {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	{
		size := m.Sent.Size()
		i -= size
		if _, err := m.Sent.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAccountVestingPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.Withdrawn.Size()
		i -= size
		if _, err := m.Withdrawn.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAccountVestingPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.InitiallyLocked.Size()
		i -= size
		if _, err := m.InitiallyLocked.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAccountVestingPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LockEnd, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LockEnd):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintAccountVestingPool(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LockStart, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LockStart):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintAccountVestingPool(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	if len(m.VestingType) > 0 {
		i -= len(m.VestingType)
		copy(dAtA[i:], m.VestingType)
		i = encodeVarintAccountVestingPool(dAtA, i, uint64(len(m.VestingType)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintAccountVestingPool(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *VestingPoolReservation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VestingPoolReservation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VestingPoolReservation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAccountVestingPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Id != 0 {
		i = encodeVarintAccountVestingPool(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccountVestingPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccountVestingPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccountVestingPools) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovAccountVestingPool(uint64(l))
	}
	if len(m.VestingPools) > 0 {
		for _, e := range m.VestingPools {
			l = e.Size()
			n += 1 + l + sovAccountVestingPool(uint64(l))
		}
	}
	return n
}

func (m *VestingPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAccountVestingPool(uint64(l))
	}
	l = len(m.VestingType)
	if l > 0 {
		n += 1 + l + sovAccountVestingPool(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LockStart)
	n += 1 + l + sovAccountVestingPool(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LockEnd)
	n += 1 + l + sovAccountVestingPool(uint64(l))
	l = m.InitiallyLocked.Size()
	n += 1 + l + sovAccountVestingPool(uint64(l))
	l = m.Withdrawn.Size()
	n += 1 + l + sovAccountVestingPool(uint64(l))
	l = m.Sent.Size()
	n += 1 + l + sovAccountVestingPool(uint64(l))
	if m.GenesisPool {
		n += 2
	}
	if len(m.Reservations) > 0 {
		for _, e := range m.Reservations {
			l = e.Size()
			n += 1 + l + sovAccountVestingPool(uint64(l))
		}
	}
	return n
}

func (m *VestingPoolReservation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAccountVestingPool(uint64(m.Id))
	}
	l = m.Amount.Size()
	n += 1 + l + sovAccountVestingPool(uint64(l))
	return n
}

func sovAccountVestingPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccountVestingPool(x uint64) (n int) {
	return sovAccountVestingPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccountVestingPools) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccountVestingPool
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
			return fmt.Errorf("proto: AccountVestingPools: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountVestingPools: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVestingPool
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
				return ErrInvalidLengthAccountVestingPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountVestingPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPools", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVestingPool
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
				return ErrInvalidLengthAccountVestingPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccountVestingPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingPools = append(m.VestingPools, &VestingPool{})
			if err := m.VestingPools[len(m.VestingPools)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccountVestingPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccountVestingPool
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
func (m *VestingPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccountVestingPool
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
			return fmt.Errorf("proto: VestingPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VestingPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVestingPool
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
				return ErrInvalidLengthAccountVestingPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountVestingPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVestingPool
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
				return ErrInvalidLengthAccountVestingPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountVestingPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockStart", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVestingPool
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
				return ErrInvalidLengthAccountVestingPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccountVestingPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LockStart, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockEnd", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVestingPool
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
				return ErrInvalidLengthAccountVestingPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccountVestingPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LockEnd, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitiallyLocked", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVestingPool
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
				return ErrInvalidLengthAccountVestingPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountVestingPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InitiallyLocked.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Withdrawn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVestingPool
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
				return ErrInvalidLengthAccountVestingPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountVestingPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Withdrawn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVestingPool
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
				return ErrInvalidLengthAccountVestingPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountVestingPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Sent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisPool", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVestingPool
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
			m.GenesisPool = bool(v != 0)
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reservations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVestingPool
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
				return ErrInvalidLengthAccountVestingPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccountVestingPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reservations = append(m.Reservations, &VestingPoolReservation{})
			if err := m.Reservations[len(m.Reservations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccountVestingPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccountVestingPool
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
func (m *VestingPoolReservation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccountVestingPool
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
			return fmt.Errorf("proto: VestingPoolReservation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VestingPoolReservation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVestingPool
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
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountVestingPool
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
				return ErrInvalidLengthAccountVestingPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountVestingPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccountVestingPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccountVestingPool
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
func skipAccountVestingPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccountVestingPool
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
					return 0, ErrIntOverflowAccountVestingPool
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
					return 0, ErrIntOverflowAccountVestingPool
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
				return 0, ErrInvalidLengthAccountVestingPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccountVestingPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccountVestingPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccountVestingPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccountVestingPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccountVestingPool = fmt.Errorf("proto: unexpected end of group")
)
