// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: c4echain/cfeminter/event.proto

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

type EventMint struct {
	BondedRatio string `protobuf:"bytes,1,opt,name=bonded_ratio,json=bondedRatio,proto3" json:"bonded_ratio,omitempty"`
	Inflation   string `protobuf:"bytes,2,opt,name=inflation,proto3" json:"inflation,omitempty"`
	Amount      string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *EventMint) Reset()         { *m = EventMint{} }
func (m *EventMint) String() string { return proto.CompactTextString(m) }
func (*EventMint) ProtoMessage()    {}
func (*EventMint) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a184b8f00fdd2a8, []int{0}
}
func (m *EventMint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventMint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventMint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventMint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventMint.Merge(m, src)
}
func (m *EventMint) XXX_Size() int {
	return m.Size()
}
func (m *EventMint) XXX_DiscardUnknown() {
	xxx_messageInfo_EventMint.DiscardUnknown(m)
}

var xxx_messageInfo_EventMint proto.InternalMessageInfo

func (m *EventMint) GetBondedRatio() string {
	if m != nil {
		return m.BondedRatio
	}
	return ""
}

func (m *EventMint) GetInflation() string {
	if m != nil {
		return m.Inflation
	}
	return ""
}

func (m *EventMint) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func init() {
	proto.RegisterType((*EventMint)(nil), "chain4energy.c4echain.cfeminter.EventMint")
}

func init() { proto.RegisterFile("c4echain/cfeminter/event.proto", fileDescriptor_6a184b8f00fdd2a8) }

var fileDescriptor_6a184b8f00fdd2a8 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0x36, 0x49, 0x4d,
	0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x4f, 0x4e, 0x4b, 0xcd, 0xcd, 0xcc, 0x2b, 0x49, 0x2d, 0xd2, 0x4f,
	0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x07, 0x4b, 0x9a, 0xa4,
	0xe6, 0xa5, 0x16, 0xa5, 0x57, 0xea, 0xc1, 0x14, 0xeb, 0xc1, 0x15, 0x4b, 0x89, 0xa4, 0xe7, 0xa7,
	0xe7, 0x83, 0xd5, 0xea, 0x83, 0x58, 0x10, 0x6d, 0x4a, 0x29, 0x5c, 0x9c, 0xae, 0x20, 0x53, 0x7c,
	0x33, 0xf3, 0x4a, 0x84, 0x14, 0xb9, 0x78, 0x92, 0xf2, 0xf3, 0x52, 0x52, 0x53, 0xe2, 0x8b, 0x12,
	0x4b, 0x32, 0xf3, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xb8, 0x21, 0x62, 0x41, 0x20, 0x21,
	0x21, 0x19, 0x2e, 0xce, 0xcc, 0xbc, 0xb4, 0x1c, 0x10, 0x3b, 0x4f, 0x82, 0x09, 0x2c, 0x8f, 0x10,
	0x10, 0x12, 0xe3, 0x62, 0x4b, 0xcc, 0xcd, 0x2f, 0xcd, 0x2b, 0x91, 0x60, 0x06, 0x4b, 0x41, 0x79,
	0x4e, 0xbe, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84,
	0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x9c, 0x9e, 0x59,
	0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x8f, 0xec, 0x03, 0xfd, 0x64, 0x93, 0x54, 0x5d,
	0x88, 0x7f, 0x2b, 0x90, 0x7c, 0x5c, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x76, 0xbb, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xf5, 0x1b, 0xff, 0x14, 0x01, 0x00, 0x00,
}

func (m *EventMint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventMint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventMint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Inflation) > 0 {
		i -= len(m.Inflation)
		copy(dAtA[i:], m.Inflation)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Inflation)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BondedRatio) > 0 {
		i -= len(m.BondedRatio)
		copy(dAtA[i:], m.BondedRatio)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.BondedRatio)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventMint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BondedRatio)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Inflation)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventMint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventMint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventMint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BondedRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BondedRatio = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inflation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Inflation = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)
