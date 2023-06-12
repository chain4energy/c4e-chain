// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: c4echain/cfeclaim/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// GenesisState defines the cfeclaim module's genesis state.
type GenesisState struct {
	Campaigns    []Campaign  `protobuf:"bytes,2,rep,name=campaigns,proto3" json:"campaigns"`
	UsersEntries []UserEntry `protobuf:"bytes,3,rep,name=users_entries,json=usersEntries,proto3" json:"users_entries"`
	Missions     []Mission   `protobuf:"bytes,5,rep,name=missions,proto3" json:"missions"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_873bc3e9e2d08cf5, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetCampaigns() []Campaign {
	if m != nil {
		return m.Campaigns
	}
	return nil
}

func (m *GenesisState) GetUsersEntries() []UserEntry {
	if m != nil {
		return m.UsersEntries
	}
	return nil
}

func (m *GenesisState) GetMissions() []Mission {
	if m != nil {
		return m.Missions
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "chain4energy.c4echain.cfeclaim.GenesisState")
}

func init() { proto.RegisterFile("c4echain/cfeclaim/genesis.proto", fileDescriptor_873bc3e9e2d08cf5) }

var fileDescriptor_873bc3e9e2d08cf5 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x36, 0x49, 0x4d,
	0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x4f, 0x4e, 0x4b, 0x4d, 0xce, 0x49, 0xcc, 0xcc, 0xd5, 0x4f, 0x4f,
	0xcd, 0x4b, 0x2d, 0xce, 0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x03, 0xcb, 0x9a,
	0xa4, 0xe6, 0xa5, 0x16, 0xa5, 0x57, 0xea, 0xc1, 0x54, 0xeb, 0xc1, 0x54, 0x4b, 0x89, 0xa4, 0xe7,
	0xa7, 0xe7, 0x83, 0x95, 0xea, 0x83, 0x58, 0x10, 0x5d, 0x52, 0x0a, 0x98, 0xc6, 0x26, 0x27, 0xe6,
	0x16, 0x24, 0x66, 0xa6, 0xe7, 0x41, 0x55, 0x60, 0xb1, 0x38, 0x37, 0xb3, 0xb8, 0x38, 0x33, 0x1f,
	0xa6, 0x40, 0x05, 0x8b, 0x11, 0x20, 0x32, 0xbe, 0x28, 0x35, 0x39, 0xbf, 0x28, 0x05, 0xa2, 0x4a,
	0xe9, 0x2f, 0x23, 0x17, 0x8f, 0x3b, 0xc4, 0xc1, 0xc1, 0x25, 0x89, 0x25, 0xa9, 0x42, 0x3e, 0x5c,
	0x9c, 0x30, 0x9b, 0x8a, 0x25, 0x98, 0x14, 0x98, 0x35, 0xb8, 0x8d, 0x34, 0xf4, 0xf0, 0xfb, 0x41,
	0xcf, 0x19, 0xaa, 0xc1, 0x89, 0xe5, 0xc4, 0x3d, 0x79, 0x86, 0x20, 0x84, 0x01, 0x42, 0x21, 0x5c,
	0xbc, 0xa5, 0xc5, 0xa9, 0x45, 0xc5, 0xf1, 0xa9, 0x79, 0x25, 0x45, 0x99, 0xa9, 0xc5, 0x12, 0xcc,
	0x60, 0x13, 0x35, 0x09, 0x99, 0x18, 0x5a, 0x9c, 0x5a, 0xe4, 0x9a, 0x57, 0x52, 0x54, 0x09, 0x35,
	0x92, 0x07, 0x6c, 0x8a, 0x2b, 0xc4, 0x10, 0x21, 0x4f, 0x2e, 0x0e, 0xa8, 0x5f, 0x8b, 0x25, 0x58,
	0xc1, 0x06, 0xaa, 0x13, 0x32, 0xd0, 0x17, 0xa2, 0x1e, 0x6a, 0x1c, 0x5c, 0xbb, 0x93, 0xcf, 0x89,
	0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3,
	0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x19, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26,
	0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x23, 0x1b, 0xae, 0x9f, 0x6c, 0x92, 0xaa, 0x0b, 0x09, 0xd8, 0x0a,
	0x44, 0xd0, 0x96, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x03, 0xd5, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0xd0, 0xff, 0x9b, 0xb7, 0x16, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Missions) > 0 {
		for iNdEx := len(m.Missions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Missions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.UsersEntries) > 0 {
		for iNdEx := len(m.UsersEntries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UsersEntries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Campaigns) > 0 {
		for iNdEx := len(m.Campaigns) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Campaigns[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Campaigns) > 0 {
		for _, e := range m.Campaigns {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UsersEntries) > 0 {
		for _, e := range m.UsersEntries {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Missions) > 0 {
		for _, e := range m.Missions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Campaigns", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Campaigns = append(m.Campaigns, Campaign{})
			if err := m.Campaigns[len(m.Campaigns)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UsersEntries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UsersEntries = append(m.UsersEntries, UserEntry{})
			if err := m.UsersEntries[len(m.UsersEntries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Missions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Missions = append(m.Missions, Mission{})
			if err := m.Missions[len(m.Missions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
