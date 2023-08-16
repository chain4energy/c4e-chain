// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: c4echain/cfedistributor/params.proto

package v2

import (
	fmt "fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	"gopkg.in/yaml.v2"
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

// Params defines the parameters for the module.
type Params struct {
	SubDistributors []SubDistributor `protobuf:"bytes,1,rep,name=sub_distributors,json=subDistributors,proto3" json:"sub_distributors"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd9302be7fe27078, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

var (
	ValidatorsRewardsCollector = "validators_rewards_collector"
	ModuleAccount              = "MODULE_ACCOUNT"
	Main                       = "MAIN"
)
var (
	DefaultSubDistributors = []SubDistributor{
		{
			Name: "default_distributor",
			Destinations: Destinations{
				PrimaryShare: Account{
					Id:   ValidatorsRewardsCollector,
					Type: ModuleAccount,
				},
				BurnShare: sdk.ZeroDec(),
			},
			Sources: []*Account{
				{
					Id:   "",
					Type: Main,
				},
			},
		},
	}
)

// NewParams creates a new Params instance
func NewParams(subDistributors []SubDistributor) Params {
	return Params{SubDistributors: subDistributors}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultSubDistributors)
}

func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetSubDistributors() []SubDistributor {
	if m != nil {
		return m.SubDistributors
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "chain4energy.c4echain.cfedistributor.v2.Params")
}

func init() {
	proto.RegisterFile("c4echain/cfedistributor/params.proto", fileDescriptor_fd9302be7fe27078)
}

var fileDescriptor_fd9302be7fe27078 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0x36, 0x49, 0x4d,
	0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x4f, 0x4e, 0x4b, 0x4d, 0xc9, 0x2c, 0x2e, 0x29, 0xca, 0x4c, 0x2a,
	0x2d, 0xc9, 0x2f, 0xd2, 0x2f, 0x48, 0x2c, 0x4a, 0xcc, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x52, 0x01, 0x2b, 0x31, 0x49, 0xcd, 0x4b, 0x2d, 0x4a, 0xaf, 0xd4, 0x83, 0x69, 0xd1, 0x43,
	0xd5, 0x22, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0xd6, 0xa0, 0x0f, 0x62, 0x41, 0xf4, 0x4a, 0xe9,
	0xe2, 0xb2, 0xa1, 0xb8, 0x34, 0x29, 0x1e, 0x89, 0x0f, 0x51, 0xae, 0x54, 0xca, 0xc5, 0x16, 0x00,
	0xb6, 0x5a, 0x28, 0x95, 0x4b, 0x00, 0x4d, 0x49, 0xb1, 0x04, 0xa3, 0x02, 0xb3, 0x06, 0xb7, 0x91,
	0x89, 0x1e, 0x31, 0xee, 0xd1, 0x0b, 0x2e, 0x4d, 0x72, 0x41, 0x70, 0x9d, 0x58, 0x4e, 0xdc, 0x93,
	0x67, 0x08, 0xe2, 0x2f, 0x46, 0x11, 0x2d, 0xb6, 0x62, 0x99, 0xb1, 0x40, 0x9e, 0xc1, 0x29, 0xe8,
	0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e,
	0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x2c, 0xd2, 0x33, 0x4b, 0x32, 0x4a,
	0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x91, 0xad, 0xd5, 0x4f, 0x36, 0x49, 0xd5, 0x85, 0x78, 0xac,
	0x02, 0xdd, 0x6b, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x1f, 0x19, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x6a, 0x80, 0x74, 0x71, 0x64, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SubDistributors) > 0 {
		for iNdEx := len(m.SubDistributors) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SubDistributors[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SubDistributors) > 0 {
		for _, e := range m.SubDistributors {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubDistributors", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubDistributors = append(m.SubDistributors, SubDistributor{})
			if err := m.SubDistributors[len(m.SubDistributors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
