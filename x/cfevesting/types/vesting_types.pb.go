// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chain4energy/c4echain/cfevesting/vesting_types.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

type VestingTypes struct {
	VestingTypes []*VestingType `protobuf:"bytes,1,rep,name=vesting_types,json=vestingTypes,proto3" json:"vesting_types,omitempty"`
}

func (m *VestingTypes) Reset()         { *m = VestingTypes{} }
func (m *VestingTypes) String() string { return proto.CompactTextString(m) }
func (*VestingTypes) ProtoMessage()    {}
func (*VestingTypes) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aecefaec2bee0ce, []int{0}
}
func (m *VestingTypes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VestingTypes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VestingTypes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VestingTypes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VestingTypes.Merge(m, src)
}
func (m *VestingTypes) XXX_Size() int {
	return m.Size()
}
func (m *VestingTypes) XXX_DiscardUnknown() {
	xxx_messageInfo_VestingTypes.DiscardUnknown(m)
}

var xxx_messageInfo_VestingTypes proto.InternalMessageInfo

func (m *VestingTypes) GetVestingTypes() []*VestingType {
	if m != nil {
		return m.VestingTypes
	}
	return nil
}

type VestingType struct {
	// vesting type name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// period of locked coins (minutes) from vesting start
	LockupPeriod time.Duration `protobuf:"bytes,2,opt,name=lockup_period,json=lockupPeriod,proto3,stdduration" json:"lockup_period"`
	// period of vesting coins (minutes) from lockup period end
	VestingPeriod time.Duration `protobuf:"bytes,3,opt,name=vesting_period,json=vestingPeriod,proto3,stdduration" json:"vesting_period"`
	// the percentage of tokens that are released initially
	Free github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=free,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"free"`
}

func (m *VestingType) Reset()         { *m = VestingType{} }
func (m *VestingType) String() string { return proto.CompactTextString(m) }
func (*VestingType) ProtoMessage()    {}
func (*VestingType) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aecefaec2bee0ce, []int{1}
}
func (m *VestingType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VestingType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VestingType.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VestingType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VestingType.Merge(m, src)
}
func (m *VestingType) XXX_Size() int {
	return m.Size()
}
func (m *VestingType) XXX_DiscardUnknown() {
	xxx_messageInfo_VestingType.DiscardUnknown(m)
}

var xxx_messageInfo_VestingType proto.InternalMessageInfo

func (m *VestingType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VestingType) GetLockupPeriod() time.Duration {
	if m != nil {
		return m.LockupPeriod
	}
	return 0
}

func (m *VestingType) GetVestingPeriod() time.Duration {
	if m != nil {
		return m.VestingPeriod
	}
	return 0
}

func init() {
	proto.RegisterType((*VestingTypes)(nil), "chain4energy.c4echain.cfevesting.VestingTypes")
	proto.RegisterType((*VestingType)(nil), "chain4energy.c4echain.cfevesting.VestingType")
}

func init() {
	proto.RegisterFile("chain4energy/c4echain/cfevesting/vesting_types.proto", fileDescriptor_3aecefaec2bee0ce)
}

var fileDescriptor_3aecefaec2bee0ce = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0xe3, 0xaf, 0xd5, 0x27, 0x70, 0x5b, 0x86, 0x88, 0x21, 0x74, 0x70, 0xa3, 0x0e, 0xa8,
	0x4b, 0x6d, 0xa9, 0x54, 0x62, 0x8f, 0x3a, 0x20, 0x26, 0x88, 0x10, 0x03, 0x4b, 0xd5, 0xb8, 0xae,
	0x1b, 0xb5, 0xcd, 0x8d, 0xf2, 0xa7, 0xa2, 0x6f, 0xc1, 0xc8, 0x23, 0x75, 0xec, 0x88, 0x18, 0x0a,
	0x6a, 0x1e, 0x83, 0x05, 0xc5, 0x4e, 0x44, 0x60, 0x81, 0xe9, 0xde, 0xc4, 0xf7, 0xfc, 0xae, 0x8f,
	0x7c, 0xf0, 0x90, 0xcf, 0x27, 0x7e, 0x30, 0x14, 0x81, 0x88, 0xe4, 0x86, 0xf1, 0xa1, 0x50, 0xdf,
	0x8c, 0xcf, 0xc4, 0x5a, 0xc4, 0x89, 0x1f, 0x48, 0x56, 0xd4, 0x71, 0xb2, 0x09, 0x45, 0x4c, 0xc3,
	0x08, 0x12, 0x30, 0xed, 0xaa, 0x8a, 0x96, 0x2a, 0xfa, 0xa5, 0x6a, 0x13, 0x09, 0x20, 0x97, 0x82,
	0xa9, 0x79, 0x2f, 0x9d, 0xb1, 0x69, 0x1a, 0x4d, 0x12, 0x1f, 0x02, 0x4d, 0x68, 0x9f, 0x4a, 0x90,
	0xa0, 0x5a, 0x96, 0x77, 0xfa, 0x6f, 0xd7, 0xc3, 0xcd, 0x7b, 0x0d, 0xb8, 0xcb, 0xb7, 0x99, 0x2e,
	0x6e, 0x7d, 0x5b, 0x6f, 0x21, 0xbb, 0xd6, 0x6b, 0x0c, 0xfa, 0xf4, 0xb7, 0xfd, 0xb4, 0x82, 0x71,
	0x9b, 0xeb, 0x0a, 0xb3, 0xfb, 0x81, 0x70, 0xa3, 0x72, 0x6a, 0x9a, 0xb8, 0x1e, 0x4c, 0x56, 0xc2,
	0x42, 0x36, 0xea, 0x1d, 0xbb, 0xaa, 0x37, 0xaf, 0x70, 0x6b, 0x09, 0x7c, 0x91, 0x86, 0xe3, 0x50,
	0x44, 0x3e, 0x4c, 0xad, 0x7f, 0x36, 0xea, 0x35, 0x06, 0x67, 0x54, 0xbb, 0xa2, 0xa5, 0x2b, 0x3a,
	0x2a, 0x5c, 0x39, 0x47, 0xdb, 0x7d, 0xc7, 0x78, 0x7e, 0xeb, 0x20, 0xb7, 0xa9, 0x95, 0x37, 0x4a,
	0x68, 0x5e, 0xe3, 0x93, 0xd2, 0x41, 0x81, 0xaa, 0xfd, 0x1d, 0x55, 0x9a, 0x2f, 0x58, 0x0e, 0xae,
	0xcf, 0x22, 0x21, 0xac, 0x7a, 0x7e, 0x53, 0x87, 0xe6, 0x63, 0xaf, 0xfb, 0xce, 0xb9, 0xf4, 0x93,
	0x79, 0xea, 0x51, 0x0e, 0x2b, 0xc6, 0x21, 0x5e, 0x41, 0x5c, 0x94, 0x7e, 0x3c, 0x5d, 0x30, 0xfd,
	0x6a, 0x23, 0xc1, 0x5d, 0xa5, 0x75, 0x6e, 0xb7, 0x07, 0x82, 0x76, 0x07, 0x82, 0xde, 0x0f, 0x04,
	0x3d, 0x65, 0xc4, 0xd8, 0x65, 0xc4, 0x78, 0xc9, 0x88, 0xf1, 0x70, 0x59, 0xe5, 0xfc, 0x08, 0x45,
	0x5f, 0xa7, 0x62, 0x3d, 0x60, 0x8f, 0xd5, 0x68, 0x28, 0xb8, 0xf7, 0x5f, 0x59, 0xb8, 0xf8, 0x0c,
	0x00, 0x00, 0xff, 0xff, 0x24, 0x57, 0xab, 0x46, 0x4b, 0x02, 0x00, 0x00,
}

func (m *VestingTypes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VestingTypes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VestingTypes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VestingTypes) > 0 {
		for iNdEx := len(m.VestingTypes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VestingTypes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVestingTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *VestingType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VestingType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VestingType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Free.Size()
		i -= size
		if _, err := m.Free.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintVestingTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	n1, err1 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.VestingPeriod, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.VestingPeriod):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintVestingTypes(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	n2, err2 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.LockupPeriod, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.LockupPeriod):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintVestingTypes(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintVestingTypes(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVestingTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovVestingTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VestingTypes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.VestingTypes) > 0 {
		for _, e := range m.VestingTypes {
			l = e.Size()
			n += 1 + l + sovVestingTypes(uint64(l))
		}
	}
	return n
}

func (m *VestingType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovVestingTypes(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.LockupPeriod)
	n += 1 + l + sovVestingTypes(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.VestingPeriod)
	n += 1 + l + sovVestingTypes(uint64(l))
	l = m.Free.Size()
	n += 1 + l + sovVestingTypes(uint64(l))
	return n
}

func sovVestingTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVestingTypes(x uint64) (n int) {
	return sovVestingTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VestingTypes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVestingTypes
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
			return fmt.Errorf("proto: VestingTypes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VestingTypes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingTypes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVestingTypes
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
				return ErrInvalidLengthVestingTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVestingTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingTypes = append(m.VestingTypes, &VestingType{})
			if err := m.VestingTypes[len(m.VestingTypes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVestingTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVestingTypes
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
func (m *VestingType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVestingTypes
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
			return fmt.Errorf("proto: VestingType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VestingType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVestingTypes
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
				return ErrInvalidLengthVestingTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVestingTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockupPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVestingTypes
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
				return ErrInvalidLengthVestingTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVestingTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.LockupPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVestingTypes
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
				return ErrInvalidLengthVestingTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVestingTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.VestingPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Free", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVestingTypes
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
				return ErrInvalidLengthVestingTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVestingTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Free.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVestingTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVestingTypes
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
func skipVestingTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVestingTypes
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
					return 0, ErrIntOverflowVestingTypes
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
					return 0, ErrIntOverflowVestingTypes
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
				return 0, ErrInvalidLengthVestingTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVestingTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVestingTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVestingTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVestingTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVestingTypes = fmt.Errorf("proto: unexpected end of group")
)
