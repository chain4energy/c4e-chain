// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: c4echain/cfevesting/account_vesting_pool.proto

package types

import (
	fmt "fmt"
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

type VestingPoolReservation struct {
	Id              uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	VestingPoolName string                                 `protobuf:"bytes,2,opt,name=vestingPoolName,proto3" json:"vestingPoolName,omitempty"`
	Amount          github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
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

func (m *VestingPoolReservation) GetVestingPoolName() string {
	if m != nil {
		return m.VestingPoolName
	}
	return ""
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
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x6e, 0xda, 0xae, 0x6b, 0x5d, 0x60, 0xc8, 0x4c, 0x28, 0xea, 0x21, 0x0d, 0x3d, 0xa0, 0x5c,
	0xea, 0x48, 0xa3, 0x77, 0x44, 0x11, 0x48, 0x48, 0xd3, 0x84, 0xc2, 0x84, 0x04, 0x97, 0x28, 0x4d,
	0xde, 0x52, 0x6b, 0x89, 0x1d, 0xc5, 0x6e, 0x4b, 0x4f, 0xfc, 0x85, 0x9d, 0xf9, 0x45, 0x3b, 0xee,
	0x88, 0x76, 0x18, 0xa8, 0xfd, 0x23, 0xc8, 0x4e, 0xba, 0x59, 0x5c, 0xa6, 0xed, 0x64, 0xbf, 0xe7,
	0xf7, 0x7d, 0xef, 0x3d, 0x7d, 0x9f, 0x11, 0x89, 0x27, 0x10, 0xcf, 0x23, 0xca, 0xfc, 0xf8, 0x0c,
	0x96, 0x20, 0x24, 0x65, 0xa9, 0x1f, 0xc5, 0x31, 0x5f, 0x30, 0x19, 0xd6, 0x71, 0x58, 0x70, 0x9e,
	0x91, 0xa2, 0xe4, 0x92, 0x63, 0x57, 0x17, 0x4f, 0x80, 0x41, 0x99, 0xae, 0x6f, 0xc1, 0xe4, 0x0e,
	0x3c, 0x38, 0x4c, 0x79, 0xca, 0x75, 0xb1, 0xaf, 0x6e, 0x15, 0x6e, 0x30, 0x4c, 0x39, 0x4f, 0x33,
	0xf0, 0x75, 0x34, 0x5b, 0x9c, 0xf9, 0x92, 0xe6, 0x20, 0x64, 0x94, 0x17, 0x55, 0xc1, 0xe8, 0x27,
	0x7a, 0xf1, 0xae, 0x6a, 0xfb, 0xb5, 0x22, 0xfa, 0xcc, 0x79, 0x26, 0xf0, 0x21, 0xda, 0xe3, 0x2b,
	0x06, 0xa5, 0x6d, 0xb9, 0x96, 0xd7, 0x0b, 0xaa, 0x00, 0x07, 0xe8, 0xa9, 0x39, 0x9b, 0xb0, 0x9b,
	0x6e, 0xcb, 0xeb, 0x1f, 0x8d, 0xc9, 0x7d, 0xd3, 0x11, 0x83, 0x3c, 0x78, 0xb2, 0x34, 0x3a, 0x8d,
	0xae, 0x5b, 0xa8, 0x6f, 0xbc, 0x62, 0x8c, 0xda, 0x2c, 0xca, 0xa1, 0x6e, 0xac, 0xef, 0xf8, 0x15,
	0xda, 0x61, 0x42, 0xb9, 0x2e, 0xc0, 0x6e, 0xea, 0xb7, 0x7e, 0x9d, 0x3b, 0x5d, 0x17, 0x80, 0xdf,
	0x23, 0x94, 0xf1, 0xf8, 0x3c, 0x14, 0x32, 0x2a, 0xa5, 0xdd, 0x72, 0x2d, 0xaf, 0x7f, 0x34, 0x20,
	0xd5, 0xf6, 0x64, 0xb7, 0x3d, 0x39, 0xdd, 0x6d, 0x3f, 0xed, 0x5e, 0xde, 0x0c, 0x1b, 0x17, 0x7f,
	0x86, 0x56, 0xd0, 0x53, 0xb8, 0x2f, 0x0a, 0x86, 0xdf, 0xa2, 0xae, 0x26, 0x01, 0x96, 0xd8, 0xed,
	0x07, 0x50, 0xec, 0x2b, 0xd4, 0x07, 0x96, 0xe0, 0x6f, 0xe8, 0x39, 0x65, 0x54, 0xd2, 0x28, 0xcb,
	0xd6, 0xa1, 0x4a, 0x42, 0x62, 0xef, 0xa9, 0x61, 0xa7, 0x44, 0x15, 0x5f, 0xdf, 0x0c, 0x5f, 0xa7,
	0x54, 0xce, 0x17, 0x33, 0x12, 0xf3, 0xdc, 0x8f, 0xb9, 0xc8, 0xb9, 0xa8, 0x8f, 0xb1, 0x48, 0xce,
	0x7d, 0xb5, 0x9d, 0x20, 0x9f, 0x98, 0x0c, 0x0e, 0x6e, 0x79, 0x8e, 0x35, 0x0d, 0x3e, 0x46, 0xbd,
	0x15, 0x95, 0xf3, 0xa4, 0x8c, 0x56, 0xcc, 0xee, 0x3c, 0x8a, 0xf3, 0x8e, 0x00, 0x4f, 0x51, 0x5b,
	0x00, 0x93, 0xf6, 0xfe, 0xa3, 0x88, 0x34, 0x56, 0xa9, 0x92, 0x02, 0x03, 0x41, 0x85, 0x76, 0x83,
	0xdd, 0x75, 0x2d, 0xaf, 0x1b, 0xf4, 0xeb, 0x9c, 0x12, 0x73, 0xf4, 0xcb, 0x42, 0x2f, 0x4d, 0xe9,
	0x41, 0x40, 0xb9, 0x8c, 0x24, 0xe5, 0x0c, 0x3f, 0x43, 0x4d, 0x9a, 0x68, 0x95, 0xdb, 0x41, 0x93,
	0x26, 0xd8, 0x43, 0x07, 0x86, 0x2f, 0x4e, 0x94, 0x05, 0x2a, 0x99, 0xff, 0x4f, 0xe3, 0x8f, 0xa8,
	0x13, 0xe5, 0xca, 0xb1, 0x5a, 0xe6, 0x87, 0x4f, 0x5f, 0xa3, 0xa7, 0x27, 0x97, 0x1b, 0xc7, 0xba,
	0xda, 0x38, 0xd6, 0xdf, 0x8d, 0x63, 0x5d, 0x6c, 0x9d, 0xc6, 0xd5, 0xd6, 0x69, 0xfc, 0xde, 0x3a,
	0x8d, 0xef, 0x13, 0x93, 0xc9, 0xb0, 0xb6, 0x1f, 0x4f, 0x60, 0x5c, 0x7d, 0xdb, 0x1f, 0xe6, 0xc7,
	0xd5, 0xdc, 0xb3, 0x8e, 0xf6, 0xc8, 0x9b, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x48, 0xad, 0xa5,
	0xd3, 0xdc, 0x03, 0x00, 0x00,
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
	dAtA[i] = 0x1a
	if len(m.VestingPoolName) > 0 {
		i -= len(m.VestingPoolName)
		copy(dAtA[i:], m.VestingPoolName)
		i = encodeVarintAccountVestingPool(dAtA, i, uint64(len(m.VestingPoolName)))
		i--
		dAtA[i] = 0x12
	}
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
	l = len(m.VestingPoolName)
	if l > 0 {
		n += 1 + l + sovAccountVestingPool(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPoolName", wireType)
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
			m.VestingPoolName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
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
