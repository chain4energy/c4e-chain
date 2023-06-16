// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: c4echain/cfeev/genesis.proto

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

// GenesisState defines the cfeev module's genesis state.
type GenesisState struct {
	Params                   Params                `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	EnergyTransferOffers     []EnergyTransferOffer `protobuf:"bytes,2,rep,name=energy_transfer_offers,json=energyTransferOffers,proto3" json:"energy_transfer_offers"`
	EnergyTransferOfferCount uint64                `protobuf:"varint,3,opt,name=energy_transfer_offer_count,json=energyTransferOfferCount,proto3" json:"energy_transfer_offer_count,omitempty"`
	EnergyTransfers          []EnergyTransfer      `protobuf:"bytes,4,rep,name=energy_transfers,json=energyTransfers,proto3" json:"energy_transfers"`
	EnergyTransferCount      uint64                `protobuf:"varint,5,opt,name=energy_transfer_count,json=energyTransferCount,proto3" json:"energy_transfer_count,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_36ea69ce9f43b4b6, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetEnergyTransferOffers() []EnergyTransferOffer {
	if m != nil {
		return m.EnergyTransferOffers
	}
	return nil
}

func (m *GenesisState) GetEnergyTransferOfferCount() uint64 {
	if m != nil {
		return m.EnergyTransferOfferCount
	}
	return 0
}

func (m *GenesisState) GetEnergyTransfers() []EnergyTransfer {
	if m != nil {
		return m.EnergyTransfers
	}
	return nil
}

func (m *GenesisState) GetEnergyTransferCount() uint64 {
	if m != nil {
		return m.EnergyTransferCount
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "chain4energy.c4echain.cfeev.GenesisState")
}

func init() { proto.RegisterFile("c4echain/cfeev/genesis.proto", fileDescriptor_36ea69ce9f43b4b6) }

var fileDescriptor_36ea69ce9f43b4b6 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0x36, 0x49, 0x4d,
	0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x4f, 0x4e, 0x4b, 0x4d, 0x2d, 0xd3, 0x4f, 0x4f, 0xcd, 0x4b, 0x2d,
	0xce, 0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x06, 0x4b, 0x99, 0xa4, 0xe6, 0xa5,
	0x16, 0xa5, 0x57, 0xea, 0xc1, 0x94, 0xea, 0x81, 0x95, 0x4a, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83,
	0xd5, 0xe9, 0x83, 0x58, 0x10, 0x2d, 0x52, 0xd2, 0x68, 0x06, 0x16, 0x24, 0x16, 0x25, 0xe6, 0x42,
	0xcd, 0x93, 0xd2, 0x42, 0x93, 0x84, 0x18, 0x1c, 0x5f, 0x52, 0x94, 0x98, 0x57, 0x9c, 0x96, 0x5a,
	0x14, 0x9f, 0x9f, 0x96, 0x96, 0x5a, 0x04, 0x55, 0xab, 0x82, 0x5f, 0x2d, 0x44, 0x95, 0xd2, 0x04,
	0x66, 0x2e, 0x1e, 0x77, 0x88, 0x9b, 0x83, 0x4b, 0x12, 0x4b, 0x52, 0x85, 0x1c, 0xb9, 0xd8, 0x20,
	0x56, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x29, 0xeb, 0xe1, 0xf1, 0x83, 0x5e, 0x00, 0x58,
	0xa9, 0x13, 0xcb, 0x89, 0x7b, 0xf2, 0x0c, 0x41, 0x50, 0x8d, 0x42, 0x39, 0x5c, 0x62, 0x58, 0x1d,
	0x56, 0x2c, 0xc1, 0xa4, 0xc0, 0xac, 0xc1, 0x6d, 0x64, 0x80, 0xd7, 0x48, 0x57, 0xb0, 0x68, 0x08,
	0x54, 0xa7, 0x3f, 0x48, 0x23, 0xd4, 0x7c, 0x91, 0x54, 0x4c, 0xa9, 0x62, 0x21, 0x5b, 0x2e, 0x69,
	0xac, 0xb6, 0xc5, 0x27, 0xe7, 0x97, 0xe6, 0x95, 0x48, 0x30, 0x2b, 0x30, 0x6a, 0xb0, 0x04, 0x49,
	0x60, 0xd1, 0xea, 0x0c, 0x92, 0x17, 0x8a, 0xe1, 0x12, 0x40, 0xd3, 0x5e, 0x2c, 0xc1, 0x02, 0x76,
	0xa6, 0x36, 0x09, 0xce, 0x84, 0xba, 0x90, 0x1f, 0xd5, 0x9a, 0x62, 0x21, 0x23, 0x2e, 0x51, 0x74,
	0xc7, 0x41, 0x9c, 0xc5, 0x0a, 0x76, 0x96, 0x30, 0xaa, 0x7a, 0xb0, 0x8b, 0x9c, 0x3c, 0x4f, 0x3c,
	0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e,
	0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x3f, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49,
	0x2f, 0x39, 0x3f, 0x57, 0x1f, 0xd9, 0x6d, 0xfa, 0xc9, 0x26, 0xa9, 0xba, 0x90, 0xb8, 0xae, 0x80,
	0xc6, 0x76, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38, 0x92, 0x8d, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xce, 0x19, 0x82, 0x4e, 0xa6, 0x02, 0x00, 0x00,
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
	if m.EnergyTransferCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.EnergyTransferCount))
		i--
		dAtA[i] = 0x28
	}
	if len(m.EnergyTransfers) > 0 {
		for iNdEx := len(m.EnergyTransfers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EnergyTransfers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.EnergyTransferOfferCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.EnergyTransferOfferCount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.EnergyTransferOffers) > 0 {
		for iNdEx := len(m.EnergyTransferOffers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EnergyTransferOffers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.EnergyTransferOffers) > 0 {
		for _, e := range m.EnergyTransferOffers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.EnergyTransferOfferCount != 0 {
		n += 1 + sovGenesis(uint64(m.EnergyTransferOfferCount))
	}
	if len(m.EnergyTransfers) > 0 {
		for _, e := range m.EnergyTransfers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.EnergyTransferCount != 0 {
		n += 1 + sovGenesis(uint64(m.EnergyTransferCount))
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
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnergyTransferOffers", wireType)
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
			m.EnergyTransferOffers = append(m.EnergyTransferOffers, EnergyTransferOffer{})
			if err := m.EnergyTransferOffers[len(m.EnergyTransferOffers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnergyTransferOfferCount", wireType)
			}
			m.EnergyTransferOfferCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EnergyTransferOfferCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnergyTransfers", wireType)
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
			m.EnergyTransfers = append(m.EnergyTransfers, EnergyTransfer{})
			if err := m.EnergyTransfers[len(m.EnergyTransfers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnergyTransferCount", wireType)
			}
			m.EnergyTransferCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EnergyTransferCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
