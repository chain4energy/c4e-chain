// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cfevesting/account_vesting_pool.proto

package v1

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_gogo_protobuf_types "github.com/cosmos/gogoproto/types"
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
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// string delegable_address = 2;
	VestingPools []*VestingPool `protobuf:"bytes,3,rep,name=vesting_pools,json=vestingPools,proto3" json:"vesting_pools,omitempty"`
}

func (m *AccountVestingPools) Reset()         { *m = AccountVestingPools{} }
func (m *AccountVestingPools) String() string { return proto.CompactTextString(m) }
func (*AccountVestingPools) ProtoMessage()    {}
func (*AccountVestingPools) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7f071dc6551c365, []int{0}
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

func (m *AccountVestingPools) GetAddress() string {
	if m != nil {
		return m.Address
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
	Id                        int32                                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                      string                                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	VestingType               string                                 `protobuf:"bytes,3,opt,name=vesting_type,json=vestingType,proto3" json:"vesting_type,omitempty"`
	LockStart                 time.Time                              `protobuf:"bytes,4,opt,name=lock_start,json=lockStart,proto3,stdtime" json:"lock_start"`
	LockEnd                   time.Time                              `protobuf:"bytes,5,opt,name=lock_end,json=lockEnd,proto3,stdtime" json:"lock_end"`
	Vested                    github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=vested,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"vested"`
	Withdrawn                 github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=withdrawn,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"withdrawn"`
	Sent                      github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,8,opt,name=sent,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"sent"`
	LastModification          time.Time                              `protobuf:"bytes,9,opt,name=last_modification,json=lastModification,proto3,stdtime" json:"last_modification"`
	LastModificationVested    github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,10,opt,name=last_modification_vested,json=lastModificationVested,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"last_modification_vested"`
	LastModificationWithdrawn github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,11,opt,name=last_modification_withdrawn,json=lastModificationWithdrawn,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"last_modification_withdrawn"`
}

func (m *VestingPool) Reset()         { *m = VestingPool{} }
func (m *VestingPool) String() string { return proto.CompactTextString(m) }
func (*VestingPool) ProtoMessage()    {}
func (*VestingPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7f071dc6551c365, []int{1}
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

func (m *VestingPool) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

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

func (m *VestingPool) GetLastModification() time.Time {
	if m != nil {
		return m.LastModification
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*AccountVestingPools)(nil), "chain4energy.c4echain.cfevesting.migrations.v100.AccountVestingPools")
	proto.RegisterType((*VestingPool)(nil), "chain4energy.c4echain.cfevesting.migrations.v100.VestingPool")
}

func init() {
	proto.RegisterFile("cfevesting/account_vesting_pool.proto", fileDescriptor_e7f071dc6551c365)
}

var fileDescriptor_e7f071dc6551c365 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xcf, 0x6f, 0xd3, 0x30,
	0x1c, 0xc5, 0x9b, 0xfe, 0xae, 0x03, 0x08, 0x0c, 0x42, 0xa6, 0x48, 0x69, 0x99, 0x04, 0xea, 0xa5,
	0x8e, 0x34, 0x7a, 0x47, 0x14, 0x81, 0x84, 0x04, 0x08, 0xc2, 0x34, 0x24, 0x2e, 0x51, 0x6a, 0xbb,
	0xa9, 0xb5, 0xc4, 0x8e, 0x62, 0x77, 0xa3, 0x67, 0xfe, 0x81, 0xfd, 0x59, 0x3b, 0xee, 0x88, 0x38,
	0x0c, 0x68, 0xff, 0x11, 0x14, 0x27, 0x51, 0xa3, 0xed, 0x30, 0xb5, 0xa7, 0xf8, 0x6b, 0xfb, 0x7d,
	0xde, 0xcb, 0x93, 0x12, 0xf0, 0x9c, 0xcc, 0xd9, 0x29, 0x53, 0x9a, 0x8b, 0xd0, 0x0d, 0x08, 0x91,
	0x4b, 0xa1, 0xfd, 0x62, 0xf6, 0x13, 0x29, 0x23, 0x9c, 0xa4, 0x52, 0x4b, 0x38, 0x24, 0x8b, 0x80,
	0x8b, 0x09, 0x13, 0x2c, 0x0d, 0x57, 0x98, 0x4c, 0x98, 0x99, 0xf1, 0x56, 0xdc, 0x7f, 0x14, 0xca,
	0x50, 0x9a, 0xcb, 0x6e, 0xb6, 0xca, 0x75, 0xfd, 0x41, 0x28, 0x65, 0x18, 0x31, 0xd7, 0x4c, 0xb3,
	0xe5, 0xdc, 0xd5, 0x3c, 0x66, 0x4a, 0x07, 0x71, 0x92, 0x5f, 0x38, 0xf8, 0x69, 0x81, 0x87, 0xaf,
	0x73, 0xdf, 0xe3, 0x9c, 0xf4, 0x59, 0xca, 0x48, 0x41, 0x04, 0x3a, 0x01, 0xa5, 0x29, 0x53, 0x0a,
	0x59, 0x43, 0x6b, 0xd4, 0xf3, 0xca, 0x11, 0x7a, 0xe0, 0x6e, 0x35, 0xa0, 0x42, 0x8d, 0x61, 0x63,
	0x64, 0x1f, 0x8e, 0xf1, 0x6d, 0x11, 0x71, 0xc5, 0xc0, 0xbb, 0x73, 0x5a, 0x71, 0x3b, 0xf8, 0xd7,
	0x02, 0x76, 0xe5, 0x14, 0xde, 0x03, 0x75, 0x4e, 0x8d, 0x71, 0xcb, 0xab, 0x73, 0x0a, 0x21, 0x68,
	0x8a, 0x20, 0x66, 0xa8, 0x6e, 0xa2, 0x98, 0x35, 0x7c, 0x06, 0x4a, 0x86, 0xaf, 0x57, 0x09, 0x43,
	0x0d, 0x73, 0x66, 0x17, 0x7b, 0x47, 0xab, 0x84, 0xc1, 0x37, 0x00, 0x44, 0x92, 0x9c, 0xf8, 0x4a,
	0x07, 0xa9, 0x46, 0xcd, 0xa1, 0x35, 0xb2, 0x0f, 0xfb, 0x38, 0xaf, 0x04, 0x97, 0x95, 0xe0, 0xa3,
	0xb2, 0x92, 0x69, 0xf7, 0xe2, 0x6a, 0x50, 0x3b, 0xff, 0x33, 0xb0, 0xbc, 0x5e, 0xa6, 0xfb, 0x9a,
	0xc9, 0xe0, 0x2b, 0xd0, 0x35, 0x10, 0x26, 0x28, 0x6a, 0xed, 0x80, 0xe8, 0x64, 0xaa, 0xb7, 0x82,
	0xc2, 0x77, 0xa0, 0x9d, 0x85, 0x62, 0x14, 0xb5, 0xb3, 0x88, 0x53, 0x9c, 0x5d, 0xf9, 0x7d, 0x35,
	0x78, 0x11, 0x72, 0xbd, 0x58, 0xce, 0x30, 0x91, 0xb1, 0x4b, 0xa4, 0x8a, 0xa5, 0x2a, 0x1e, 0x63,
	0x45, 0x4f, 0xdc, 0xec, 0x9d, 0x14, 0x7e, 0x2f, 0xb4, 0x57, 0xa8, 0xe1, 0x07, 0xd0, 0x3b, 0xe3,
	0x7a, 0x41, 0xd3, 0xe0, 0x4c, 0xa0, 0xce, 0x5e, 0xa8, 0x2d, 0x00, 0x4e, 0x41, 0x53, 0x31, 0xa1,
	0x51, 0x77, 0x2f, 0x90, 0xd1, 0xc2, 0x2f, 0xe0, 0x41, 0x14, 0x28, 0xed, 0xc7, 0x92, 0xf2, 0x39,
	0x27, 0x81, 0xe6, 0x52, 0xa0, 0xde, 0x0e, 0x1d, 0xdd, 0xcf, 0xe4, 0x1f, 0x2b, 0x6a, 0xb8, 0x00,
	0xe8, 0x06, 0xd2, 0x2f, 0xea, 0x03, 0x7b, 0x45, 0x7d, 0x7c, 0xdd, 0xe3, 0x38, 0xaf, 0x53, 0x80,
	0xa7, 0x37, 0x9d, 0xb6, 0x05, 0xdb, 0x7b, 0x99, 0x3d, 0xb9, 0x6e, 0xf6, 0xad, 0x04, 0x4e, 0x3f,
	0x5d, 0xac, 0x1d, 0xeb, 0x72, 0xed, 0x58, 0x7f, 0xd7, 0x8e, 0x75, 0xbe, 0x71, 0x6a, 0x97, 0x1b,
	0xa7, 0xf6, 0x6b, 0xe3, 0xd4, 0xbe, 0x4f, 0xaa, 0xf0, 0xca, 0x47, 0xe4, 0x92, 0x09, 0x1b, 0x9b,
	0x0d, 0xf7, 0x87, 0x5b, 0xf9, 0x4f, 0x18, 0xbb, 0x59, 0xdb, 0x34, 0xfb, 0xf2, 0x7f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xe3, 0xc3, 0xe0, 0xf2, 0x42, 0x04, 0x00, 0x00,
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
			dAtA[i] = 0x1a
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAccountVestingPool(dAtA, i, uint64(len(m.Address)))
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
	{
		size := m.LastModificationWithdrawn.Size()
		i -= size
		if _, err := m.LastModificationWithdrawn.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAccountVestingPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size := m.LastModificationVested.Size()
		i -= size
		if _, err := m.LastModificationVested.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAccountVestingPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LastModification, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LastModification):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintAccountVestingPool(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x4a
	{
		size := m.Sent.Size()
		i -= size
		if _, err := m.Sent.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAccountVestingPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.Withdrawn.Size()
		i -= size
		if _, err := m.Withdrawn.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAccountVestingPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.Vested.Size()
		i -= size
		if _, err := m.Vested.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAccountVestingPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LockEnd, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LockEnd):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintAccountVestingPool(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x2a
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LockStart, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LockStart):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintAccountVestingPool(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x22
	if len(m.VestingType) > 0 {
		i -= len(m.VestingType)
		copy(dAtA[i:], m.VestingType)
		i = encodeVarintAccountVestingPool(dAtA, i, uint64(len(m.VestingType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintAccountVestingPool(dAtA, i, uint64(len(m.Name)))
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
	l = len(m.Address)
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
	if m.Id != 0 {
		n += 1 + sovAccountVestingPool(uint64(m.Id))
	}
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
	l = m.Vested.Size()
	n += 1 + l + sovAccountVestingPool(uint64(l))
	l = m.Withdrawn.Size()
	n += 1 + l + sovAccountVestingPool(uint64(l))
	l = m.Sent.Size()
	n += 1 + l + sovAccountVestingPool(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastModification)
	n += 1 + l + sovAccountVestingPool(uint64(l))
	l = m.LastModificationVested.Size()
	n += 1 + l + sovAccountVestingPool(uint64(l))
	l = m.LastModificationWithdrawn.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
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
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
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
		case 3:
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
		case 4:
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
		case 5:
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
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vested", wireType)
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
			if err := m.Vested.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
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
		case 8:
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
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastModification", wireType)
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
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LastModification, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastModificationVested", wireType)
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
			if err := m.LastModificationVested.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastModificationWithdrawn", wireType)
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
			if err := m.LastModificationWithdrawn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
