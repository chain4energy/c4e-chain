// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: c4echain/cfetokenization/genesis.proto

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

// GenesisState defines the cfetokenization module's genesis state.
type GenesisState struct {
	Params               Params             `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	CertificateTypeList  []CertificateType  `protobuf:"bytes,2,rep,name=certificateTypeList,proto3" json:"certificateTypeList"`
	CertificateTypeCount uint64             `protobuf:"varint,3,opt,name=certificateTypeCount,proto3" json:"certificateTypeCount,omitempty"`
	UserDevicesList      []UserDevices      `protobuf:"bytes,4,rep,name=userDevicesList,proto3" json:"userDevicesList"`
	UserCertificatesList []UserCertificates `protobuf:"bytes,5,rep,name=userCertificatesList,proto3" json:"userCertificatesList"`
	DevicesList          []Device           `protobuf:"bytes,6,rep,name=devicesList,proto3" json:"devicesList"`
	Certificates         []CertificateOffer `protobuf:"bytes,7,rep,name=certificates,proto3" json:"certificates"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ae4ca1053feafab, []int{0}
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

func (m *GenesisState) GetCertificateTypeList() []CertificateType {
	if m != nil {
		return m.CertificateTypeList
	}
	return nil
}

func (m *GenesisState) GetCertificateTypeCount() uint64 {
	if m != nil {
		return m.CertificateTypeCount
	}
	return 0
}

func (m *GenesisState) GetUserDevicesList() []UserDevices {
	if m != nil {
		return m.UserDevicesList
	}
	return nil
}

func (m *GenesisState) GetUserCertificatesList() []UserCertificates {
	if m != nil {
		return m.UserCertificatesList
	}
	return nil
}

func (m *GenesisState) GetDevicesList() []Device {
	if m != nil {
		return m.DevicesList
	}
	return nil
}

func (m *GenesisState) GetCertificates() []CertificateOffer {
	if m != nil {
		return m.Certificates
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "chain4energy.c4echain.cfetokenization.GenesisState")
}

func init() {
	proto.RegisterFile("c4echain/cfetokenization/genesis.proto", fileDescriptor_6ae4ca1053feafab)
}

var fileDescriptor_6ae4ca1053feafab = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x6a, 0xf2, 0x40,
	0x14, 0xc5, 0x93, 0xcf, 0x3f, 0x1f, 0x8c, 0x42, 0x61, 0xea, 0x22, 0xb8, 0x48, 0xa5, 0x60, 0x11,
	0x8a, 0x49, 0x49, 0xa5, 0xa5, 0x5b, 0x2d, 0x74, 0xd1, 0x42, 0x8b, 0xd6, 0x4d, 0x37, 0x12, 0xe3,
	0x4d, 0x1c, 0x8a, 0x99, 0x34, 0x19, 0x4b, 0xed, 0x53, 0xf8, 0x58, 0x2e, 0x5d, 0x76, 0x55, 0x8a,
	0xbe, 0x48, 0xc9, 0x64, 0xa4, 0xa3, 0xd5, 0x12, 0x77, 0x61, 0x72, 0xcf, 0xf9, 0x9d, 0x7b, 0xe7,
	0x0e, 0x3a, 0x71, 0x1a, 0xe0, 0x0c, 0x6d, 0xe2, 0x9b, 0x8e, 0x0b, 0x8c, 0x3e, 0x83, 0x4f, 0xde,
	0x6d, 0x46, 0xa8, 0x6f, 0x7a, 0xe0, 0x43, 0x44, 0x22, 0x23, 0x08, 0x29, 0xa3, 0xb8, 0xca, 0x8b,
	0x1a, 0xe0, 0x43, 0xe8, 0x4d, 0x8c, 0x95, 0xc8, 0xd8, 0x10, 0x95, 0x4b, 0x1e, 0xf5, 0x28, 0x57,
	0x98, 0xf1, 0x57, 0x22, 0x2e, 0x57, 0x77, 0x42, 0x02, 0x3b, 0xb4, 0x47, 0x82, 0x51, 0x36, 0x77,
	0x96, 0x39, 0x10, 0x32, 0xe2, 0x12, 0xc7, 0x66, 0xd0, 0x63, 0x93, 0x00, 0x84, 0xe0, 0x74, 0xa7,
	0x60, 0x1c, 0x41, 0xd8, 0x1b, 0xc0, 0x2b, 0x71, 0x60, 0xe5, 0x7e, 0xf6, 0x77, 0xb1, 0x84, 0x10,
	0x8a, 0xe3, 0x69, 0x0e, 0x15, 0x6f, 0x92, 0x29, 0x74, 0x98, 0xcd, 0x00, 0xdf, 0xa2, 0x7c, 0x12,
	0x58, 0x53, 0x2b, 0x6a, 0xad, 0x60, 0xd5, 0x8d, 0x54, 0x53, 0x31, 0x1e, 0xb8, 0xa8, 0x99, 0x9d,
	0x7d, 0x1e, 0x29, 0x6d, 0x61, 0x81, 0x7d, 0x74, 0x28, 0x31, 0x1f, 0x27, 0x01, 0xdc, 0x91, 0x88,
	0x69, 0xff, 0x2a, 0x99, 0x5a, 0xc1, 0xba, 0x48, 0xe9, 0xdc, 0x5a, 0x77, 0x10, 0x88, 0x6d, 0xc6,
	0xd8, 0x42, 0xa5, 0x8d, 0xe3, 0x16, 0x1d, 0xfb, 0x4c, 0xcb, 0x54, 0xd4, 0x5a, 0xb6, 0xbd, 0xf5,
	0x1f, 0xee, 0xa3, 0x83, 0x78, 0x38, 0xd7, 0xc9, 0x20, 0x79, 0xbe, 0x2c, 0xcf, 0x67, 0xa5, 0xcc,
	0xd7, 0xfd, 0x51, 0x8b, 0x6c, 0x9b, 0x86, 0xf8, 0x05, 0x95, 0xe2, 0x23, 0xa9, 0x93, 0x04, 0x94,
	0xe3, 0xa0, 0xcb, 0x3d, 0x40, 0xb2, 0x85, 0xa0, 0x6d, 0xb5, 0xc6, 0x5d, 0x54, 0x18, 0x48, 0x2d,
	0xe5, 0x39, 0x29, 0xed, 0x65, 0x26, 0xd9, 0x85, 0xbf, 0xec, 0x83, 0x6d, 0x54, 0x94, 0xb7, 0x48,
	0xfb, 0xbf, 0x57, 0x07, 0x52, 0xca, 0x7b, 0xd7, 0x85, 0x50, 0x10, 0xd6, 0x2c, 0x9b, 0x9d, 0xd9,
	0x42, 0x57, 0xe7, 0x0b, 0x5d, 0xfd, 0x5a, 0xe8, 0xea, 0x74, 0xa9, 0x2b, 0xf3, 0xa5, 0xae, 0x7c,
	0x2c, 0x75, 0xe5, 0xe9, 0xca, 0x23, 0x6c, 0x38, 0xee, 0x1b, 0x0e, 0x1d, 0x99, 0x32, 0x30, 0x7e,
	0x54, 0xf5, 0x64, 0xef, 0xdf, 0x7e, 0x6d, 0x7e, 0xfc, 0x96, 0xa2, 0x7e, 0x9e, 0xaf, 0xfb, 0xf9,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x62, 0x73, 0x09, 0x0c, 0x04, 0x00, 0x00,
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
	if len(m.Certificates) > 0 {
		for iNdEx := len(m.Certificates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Certificates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.DevicesList) > 0 {
		for iNdEx := len(m.DevicesList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DevicesList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.UserCertificatesList) > 0 {
		for iNdEx := len(m.UserCertificatesList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UserCertificatesList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.UserDevicesList) > 0 {
		for iNdEx := len(m.UserDevicesList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UserDevicesList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if m.CertificateTypeCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.CertificateTypeCount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.CertificateTypeList) > 0 {
		for iNdEx := len(m.CertificateTypeList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CertificateTypeList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.CertificateTypeList) > 0 {
		for _, e := range m.CertificateTypeList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.CertificateTypeCount != 0 {
		n += 1 + sovGenesis(uint64(m.CertificateTypeCount))
	}
	if len(m.UserDevicesList) > 0 {
		for _, e := range m.UserDevicesList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UserCertificatesList) > 0 {
		for _, e := range m.UserCertificatesList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DevicesList) > 0 {
		for _, e := range m.DevicesList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Certificates) > 0 {
		for _, e := range m.Certificates {
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
				return fmt.Errorf("proto: wrong wireType = %d for field CertificateTypeList", wireType)
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
			m.CertificateTypeList = append(m.CertificateTypeList, CertificateType{})
			if err := m.CertificateTypeList[len(m.CertificateTypeList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertificateTypeCount", wireType)
			}
			m.CertificateTypeCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CertificateTypeCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserDevicesList", wireType)
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
			m.UserDevicesList = append(m.UserDevicesList, UserDevices{})
			if err := m.UserDevicesList[len(m.UserDevicesList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserCertificatesList", wireType)
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
			m.UserCertificatesList = append(m.UserCertificatesList, UserCertificates{})
			if err := m.UserCertificatesList[len(m.UserCertificatesList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevicesList", wireType)
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
			m.DevicesList = append(m.DevicesList, Device{})
			if err := m.DevicesList[len(m.DevicesList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificates", wireType)
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
			m.Certificates = append(m.Certificates, CertificateOffer{})
			if err := m.Certificates[len(m.Certificates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
