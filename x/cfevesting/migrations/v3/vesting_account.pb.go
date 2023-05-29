// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: c4echain/cfevesting/vesting_account.proto

package v3

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type VestingAccountTrace struct {
	Id                 uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address            string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Genesis            bool   `protobuf:"varint,3,opt,name=genesis,proto3" json:"genesis,omitempty"`
	FromGenesisPool    bool   `protobuf:"varint,4,opt,name=from_genesis_pool,json=fromGenesisPool,proto3" json:"from_genesis_pool,omitempty"`
	FromGenesisAccount bool   `protobuf:"varint,5,opt,name=from_genesis_account,json=fromGenesisAccount,proto3" json:"from_genesis_account,omitempty"`
}

func (m *VestingAccountTrace) Reset()         { *m = VestingAccountTrace{} }
func (m *VestingAccountTrace) String() string { return proto.CompactTextString(m) }
func (*VestingAccountTrace) ProtoMessage()    {}
func (*VestingAccountTrace) Descriptor() ([]byte, []int) {
	return fileDescriptor_785587b07f9ac0c9, []int{0}
}
func (m *VestingAccountTrace) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VestingAccountTrace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VestingAccountTrace.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VestingAccountTrace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VestingAccountTrace.Merge(m, src)
}
func (m *VestingAccountTrace) XXX_Size() int {
	return m.Size()
}
func (m *VestingAccountTrace) XXX_DiscardUnknown() {
	xxx_messageInfo_VestingAccountTrace.DiscardUnknown(m)
}

var xxx_messageInfo_VestingAccountTrace proto.InternalMessageInfo

func (m *VestingAccountTrace) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *VestingAccountTrace) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *VestingAccountTrace) GetGenesis() bool {
	if m != nil {
		return m.Genesis
	}
	return false
}

func (m *VestingAccountTrace) GetFromGenesisPool() bool {
	if m != nil {
		return m.FromGenesisPool
	}
	return false
}

func (m *VestingAccountTrace) GetFromGenesisAccount() bool {
	if m != nil {
		return m.FromGenesisAccount
	}
	return false
}

func init() {
	proto.RegisterType((*VestingAccountTrace)(nil), "chain4energy.c4echain.cfevesting.VestingAccountTrace")
}

func init() {
	proto.RegisterFile("c4echain/cfevesting/vesting_account.proto", fileDescriptor_785587b07f9ac0c9)
}

var fileDescriptor_785587b07f9ac0c9 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4c, 0x36, 0x49, 0x4d,
	0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x4f, 0x4e, 0x4b, 0x2d, 0x4b, 0x2d, 0x2e, 0xc9, 0xcc, 0x4b, 0xd7,
	0x87, 0xd2, 0xf1, 0x89, 0xc9, 0xc9, 0xf9, 0xa5, 0x79, 0x25, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x42, 0x0a, 0x60, 0x75, 0x26, 0xa9, 0x79, 0xa9, 0x45, 0xe9, 0x95, 0x7a, 0x30, 0x7d, 0x7a, 0x08,
	0x7d, 0x4a, 0xdb, 0x19, 0xb9, 0x84, 0xc3, 0x20, 0x6c, 0x47, 0x88, 0xd6, 0x90, 0xa2, 0xc4, 0xe4,
	0x54, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x96, 0x20, 0xa6, 0xcc,
	0x14, 0x21, 0x09, 0x2e, 0xf6, 0xc4, 0x94, 0x94, 0xa2, 0xd4, 0xe2, 0x62, 0x09, 0x26, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x18, 0x17, 0x24, 0x93, 0x9e, 0x9a, 0x97, 0x5a, 0x9c, 0x59, 0x2c, 0xc1, 0xac,
	0xc0, 0xa8, 0xc1, 0x11, 0x04, 0xe3, 0x0a, 0x69, 0x71, 0x09, 0xa6, 0x15, 0xe5, 0xe7, 0xc6, 0x43,
	0xf9, 0xf1, 0x05, 0xf9, 0xf9, 0x39, 0x12, 0x2c, 0x60, 0x35, 0xfc, 0x20, 0x09, 0x77, 0x88, 0x78,
	0x40, 0x7e, 0x7e, 0x8e, 0x90, 0x01, 0x97, 0x08, 0x8a, 0x5a, 0xa8, 0x3f, 0x24, 0x58, 0xc1, 0xca,
	0x85, 0x90, 0x94, 0x43, 0x9d, 0xe9, 0xe4, 0x77, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c,
	0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72,
	0x0c, 0x51, 0x26, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xc8, 0x01,
	0xa0, 0x9f, 0x6c, 0x92, 0xaa, 0x0b, 0x09, 0xb9, 0x0a, 0xe4, 0xb0, 0x2b, 0xa9, 0x2c, 0x48, 0x2d,
	0x4e, 0x62, 0x03, 0x07, 0x99, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xae, 0xe5, 0xff, 0xda, 0x5f,
	0x01, 0x00, 0x00,
}

func (m *VestingAccountTrace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VestingAccountTrace) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VestingAccountTrace) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FromGenesisAccount {
		i--
		if m.FromGenesisAccount {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.FromGenesisPool {
		i--
		if m.FromGenesisPool {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Genesis {
		i--
		if m.Genesis {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintVestingAccount(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintVestingAccount(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintVestingAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovVestingAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VestingAccountTrace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovVestingAccount(uint64(m.Id))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovVestingAccount(uint64(l))
	}
	if m.Genesis {
		n += 2
	}
	if m.FromGenesisPool {
		n += 2
	}
	if m.FromGenesisAccount {
		n += 2
	}
	return n
}

func sovVestingAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVestingAccount(x uint64) (n int) {
	return sovVestingAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VestingAccountTrace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVestingAccount
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
			return fmt.Errorf("proto: VestingAccountTrace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VestingAccountTrace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVestingAccount
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
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVestingAccount
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
				return ErrInvalidLengthVestingAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVestingAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Genesis", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVestingAccount
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
			m.Genesis = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromGenesisPool", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVestingAccount
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
			m.FromGenesisPool = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromGenesisAccount", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVestingAccount
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
			m.FromGenesisAccount = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipVestingAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVestingAccount
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
func skipVestingAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVestingAccount
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
					return 0, ErrIntOverflowVestingAccount
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
					return 0, ErrIntOverflowVestingAccount
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
				return 0, ErrInvalidLengthVestingAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVestingAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVestingAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVestingAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVestingAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVestingAccount = fmt.Errorf("proto: unexpected end of group")
)
