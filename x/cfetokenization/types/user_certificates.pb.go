// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: c4echain/cfetokenization/user_certificates.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

type CertificateStatus int32

const (
	CertificateStatus_UNKNOWN_CERTIFICATE_STATUS CertificateStatus = 0
	CertificateStatus_VALID                      CertificateStatus = 1
	CertificateStatus_INVALID                    CertificateStatus = 2
	CertificateStatus_BURNED                     CertificateStatus = 3
)

var CertificateStatus_name = map[int32]string{
	0: "UNKNOWN_CERTIFICATE_STATUS",
	1: "VALID",
	2: "INVALID",
	3: "BURNED",
}

var CertificateStatus_value = map[string]int32{
	"UNKNOWN_CERTIFICATE_STATUS": 0,
	"VALID":                      1,
	"INVALID":                    2,
	"BURNED":                     3,
}

func (x CertificateStatus) String() string {
	return proto.EnumName(CertificateStatus_name, int32(x))
}

func (CertificateStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a17fd18775f25ee5, []int{0}
}

type UserCertificates struct {
	Owner        string         `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Certificates []*Certificate `protobuf:"bytes,2,rep,name=certificates,proto3" json:"certificates,omitempty"`
}

func (m *UserCertificates) Reset()         { *m = UserCertificates{} }
func (m *UserCertificates) String() string { return proto.CompactTextString(m) }
func (*UserCertificates) ProtoMessage()    {}
func (*UserCertificates) Descriptor() ([]byte, []int) {
	return fileDescriptor_a17fd18775f25ee5, []int{0}
}
func (m *UserCertificates) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserCertificates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserCertificates.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserCertificates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCertificates.Merge(m, src)
}
func (m *UserCertificates) XXX_Size() int {
	return m.Size()
}
func (m *UserCertificates) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCertificates.DiscardUnknown(m)
}

var xxx_messageInfo_UserCertificates proto.InternalMessageInfo

func (m *UserCertificates) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *UserCertificates) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

type Certificate struct {
	Id                 uint64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CertyficateTypeId  uint64            `protobuf:"varint,2,opt,name=certyficate_type_id,json=certyficateTypeId,proto3" json:"certyficate_type_id,omitempty"`
	Power              uint64            `protobuf:"varint,3,opt,name=power,proto3" json:"power,omitempty"`
	DeviceAddress      string            `protobuf:"bytes,4,opt,name=device_address,json=deviceAddress,proto3" json:"device_address,omitempty"`
	AllowedAuthorities []string          `protobuf:"bytes,5,rep,name=allowed_authorities,json=allowedAuthorities,proto3" json:"allowed_authorities,omitempty"`
	Authority          string            `protobuf:"bytes,6,opt,name=authority,proto3" json:"authority,omitempty"`
	CertificateStatus  CertificateStatus `protobuf:"varint,7,opt,name=certificate_status,json=certificateStatus,proto3,enum=chain4energy.c4echain.cfetokenization.CertificateStatus" json:"certificate_status,omitempty"`
}

func (m *Certificate) Reset()         { *m = Certificate{} }
func (m *Certificate) String() string { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()    {}
func (*Certificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_a17fd18775f25ee5, []int{1}
}
func (m *Certificate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Certificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Certificate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Certificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificate.Merge(m, src)
}
func (m *Certificate) XXX_Size() int {
	return m.Size()
}
func (m *Certificate) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificate.DiscardUnknown(m)
}

var xxx_messageInfo_Certificate proto.InternalMessageInfo

func (m *Certificate) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Certificate) GetCertyficateTypeId() uint64 {
	if m != nil {
		return m.CertyficateTypeId
	}
	return 0
}

func (m *Certificate) GetPower() uint64 {
	if m != nil {
		return m.Power
	}
	return 0
}

func (m *Certificate) GetDeviceAddress() string {
	if m != nil {
		return m.DeviceAddress
	}
	return ""
}

func (m *Certificate) GetAllowedAuthorities() []string {
	if m != nil {
		return m.AllowedAuthorities
	}
	return nil
}

func (m *Certificate) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *Certificate) GetCertificateStatus() CertificateStatus {
	if m != nil {
		return m.CertificateStatus
	}
	return CertificateStatus_UNKNOWN_CERTIFICATE_STATUS
}

type CertificateOffer struct {
	Id            uint64                                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CertificateId uint64                                   `protobuf:"varint,2,opt,name=certificate_id,json=certificateId,proto3" json:"certificate_id,omitempty"`
	Owner         string                                   `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Buyer         string                                   `protobuf:"bytes,4,opt,name=buyer,proto3" json:"buyer,omitempty"`
	Price         github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,5,rep,name=price,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"price"`
}

func (m *CertificateOffer) Reset()         { *m = CertificateOffer{} }
func (m *CertificateOffer) String() string { return proto.CompactTextString(m) }
func (*CertificateOffer) ProtoMessage()    {}
func (*CertificateOffer) Descriptor() ([]byte, []int) {
	return fileDescriptor_a17fd18775f25ee5, []int{2}
}
func (m *CertificateOffer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CertificateOffer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CertificateOffer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CertificateOffer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateOffer.Merge(m, src)
}
func (m *CertificateOffer) XXX_Size() int {
	return m.Size()
}
func (m *CertificateOffer) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateOffer.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateOffer proto.InternalMessageInfo

func (m *CertificateOffer) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CertificateOffer) GetCertificateId() uint64 {
	if m != nil {
		return m.CertificateId
	}
	return 0
}

func (m *CertificateOffer) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *CertificateOffer) GetBuyer() string {
	if m != nil {
		return m.Buyer
	}
	return ""
}

func (m *CertificateOffer) GetPrice() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Price
	}
	return nil
}

func init() {
	proto.RegisterEnum("chain4energy.c4echain.cfetokenization.CertificateStatus", CertificateStatus_name, CertificateStatus_value)
	proto.RegisterType((*UserCertificates)(nil), "chain4energy.c4echain.cfetokenization.UserCertificates")
	proto.RegisterType((*Certificate)(nil), "chain4energy.c4echain.cfetokenization.Certificate")
	proto.RegisterType((*CertificateOffer)(nil), "chain4energy.c4echain.cfetokenization.CertificateOffer")
}

func init() {
	proto.RegisterFile("c4echain/cfetokenization/user_certificates.proto", fileDescriptor_a17fd18775f25ee5)
}

var fileDescriptor_a17fd18775f25ee5 = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xc1, 0x6e, 0xda, 0x4c,
	0x10, 0xc6, 0x26, 0x24, 0x62, 0xf9, 0x83, 0xc8, 0x86, 0x83, 0x7f, 0x54, 0x39, 0x08, 0x09, 0x09,
	0x55, 0x8a, 0x9d, 0x50, 0x0e, 0xed, 0x11, 0x08, 0x95, 0x50, 0x2b, 0x22, 0x19, 0x48, 0xa4, 0x5e,
	0x2c, 0x63, 0x0f, 0xb0, 0x4a, 0xe2, 0x45, 0xbb, 0x4b, 0xa8, 0x7b, 0xea, 0x23, 0xf4, 0x39, 0xfa,
	0x18, 0x3d, 0xe5, 0x98, 0x63, 0x4e, 0x6d, 0x05, 0x2f, 0x52, 0xb1, 0x4b, 0x9a, 0x4d, 0xb9, 0xb4,
	0x27, 0x7b, 0xbe, 0xf9, 0x66, 0xf6, 0xdb, 0x6f, 0x66, 0xd1, 0x49, 0xd8, 0x80, 0x70, 0x1a, 0x90,
	0xd8, 0x0d, 0xc7, 0x20, 0xe8, 0x15, 0xc4, 0xe4, 0x53, 0x20, 0x08, 0x8d, 0xdd, 0x39, 0x07, 0xe6,
	0x87, 0xc0, 0x04, 0x19, 0x93, 0x30, 0x10, 0xc0, 0x9d, 0x19, 0xa3, 0x82, 0xe2, 0xaa, 0xa4, 0x37,
	0x20, 0x06, 0x36, 0x49, 0x9c, 0xc7, 0x72, 0xe7, 0x8f, 0xf2, 0x92, 0x1d, 0x52, 0x7e, 0x43, 0xb9,
	0x3b, 0x0a, 0x38, 0xb8, 0xb7, 0xa7, 0x23, 0x10, 0xc1, 0xa9, 0x1b, 0x52, 0x12, 0xab, 0x36, 0xa5,
	0xe2, 0x84, 0x4e, 0xa8, 0xfc, 0x75, 0xd7, 0x7f, 0x0a, 0xad, 0x7c, 0x36, 0x50, 0x61, 0xc8, 0x81,
	0xb5, 0xb5, 0x73, 0x71, 0x11, 0x65, 0xe8, 0x22, 0x06, 0x66, 0x19, 0x65, 0xa3, 0x96, 0xf5, 0x54,
	0x80, 0x2f, 0xd0, 0x7f, 0xba, 0x3a, 0xcb, 0x2c, 0xa7, 0x6b, 0xb9, 0x7a, 0xdd, 0xf9, 0x2b, 0x79,
	0x8e, 0x76, 0x80, 0xf7, 0xac, 0x4f, 0xe5, 0x9b, 0x89, 0x72, 0x5a, 0x16, 0xe7, 0x91, 0x49, 0x22,
	0x79, 0xf4, 0x8e, 0x67, 0x92, 0x08, 0x3b, 0xe8, 0x70, 0xcd, 0x4f, 0x54, 0xda, 0x17, 0xc9, 0x0c,
	0x7c, 0x12, 0x59, 0xa6, 0x24, 0x1c, 0x68, 0xa9, 0x41, 0x32, 0x83, 0x6e, 0xb4, 0x56, 0x3f, 0xa3,
	0x0b, 0x60, 0x56, 0x5a, 0x32, 0x54, 0x80, 0xab, 0x28, 0x1f, 0xc1, 0x2d, 0x09, 0xc1, 0x0f, 0xa2,
	0x88, 0x01, 0xe7, 0xd6, 0x8e, 0xbc, 0xdc, 0xbe, 0x42, 0x9b, 0x0a, 0xc4, 0x2e, 0x3a, 0x0c, 0xae,
	0xaf, 0xe9, 0x02, 0x22, 0x3f, 0x98, 0x8b, 0x29, 0x65, 0x44, 0x10, 0xe0, 0x56, 0xa6, 0x9c, 0xae,
	0x65, 0x3d, 0xbc, 0x49, 0x35, 0x9f, 0x32, 0xf8, 0x05, 0xca, 0x3e, 0x12, 0x13, 0x6b, 0x57, 0xb6,
	0x7c, 0x02, 0xf0, 0x04, 0x61, 0xed, 0xae, 0x3e, 0x17, 0x81, 0x98, 0x73, 0x6b, 0xaf, 0x6c, 0xd4,
	0xf2, 0xf5, 0xd7, 0xff, 0xee, 0x5c, 0x5f, 0xd6, 0xab, 0x4b, 0x3f, 0x83, 0x2a, 0x0f, 0x06, 0x2a,
	0x68, 0xc4, 0xf3, 0xf1, 0x18, 0xd8, 0x96, 0x93, 0x55, 0x94, 0xd7, 0xd5, 0xfc, 0x36, 0x71, 0x5f,
	0x43, 0x95, 0x81, 0x6a, 0xfc, 0x69, 0x7d, 0xfc, 0x45, 0x94, 0x19, 0xcd, 0x13, 0x60, 0x1b, 0xdf,
	0x54, 0x80, 0x03, 0x94, 0x99, 0x31, 0x12, 0x82, 0x74, 0x28, 0x57, 0xff, 0xdf, 0x51, 0x5b, 0xe8,
	0xac, 0xb7, 0xd0, 0xd9, 0x6c, 0xa1, 0xd3, 0xa6, 0x24, 0x6e, 0x9d, 0xdc, 0x7d, 0x3f, 0x4a, 0x7d,
	0xfd, 0x71, 0x54, 0x9b, 0x10, 0x31, 0x9d, 0x8f, 0x9c, 0x90, 0xde, 0xb8, 0x9b, 0x95, 0x55, 0x9f,
	0x63, 0x1e, 0x5d, 0xb9, 0xeb, 0xe9, 0x72, 0x59, 0xc0, 0x3d, 0xd5, 0xf9, 0xe5, 0x25, 0x3a, 0xd8,
	0xb2, 0x00, 0xdb, 0xa8, 0x34, 0xec, 0xbd, 0xeb, 0x9d, 0x5f, 0xf6, 0xfc, 0x76, 0xc7, 0x1b, 0x74,
	0xdf, 0x76, 0xdb, 0xcd, 0x41, 0xc7, 0xef, 0x0f, 0x9a, 0x83, 0x61, 0xbf, 0x90, 0xc2, 0x59, 0x94,
	0xb9, 0x68, 0xbe, 0xef, 0x9e, 0x15, 0x0c, 0x9c, 0x43, 0x7b, 0xdd, 0x9e, 0x0a, 0x4c, 0x8c, 0xd0,
	0x6e, 0x6b, 0xe8, 0xf5, 0x3a, 0x67, 0x85, 0x74, 0xab, 0x7f, 0xb7, 0xb4, 0x8d, 0xfb, 0xa5, 0x6d,
	0xfc, 0x5c, 0xda, 0xc6, 0x97, 0x95, 0x9d, 0xba, 0x5f, 0xd9, 0xa9, 0x87, 0x95, 0x9d, 0xfa, 0xf0,
	0x46, 0xd7, 0xa8, 0x0d, 0xc9, 0x0d, 0x1b, 0x70, 0xac, 0x5e, 0xef, 0xc7, 0xad, 0xf7, 0x2b, 0xa5,
	0x8f, 0x76, 0xe5, 0xbb, 0x7a, 0xf5, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x08, 0x5b, 0xc5, 0xe8,
	0x03, 0x00, 0x00,
}

func (m *UserCertificates) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserCertificates) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserCertificates) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
				i = encodeVarintUserCertificates(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintUserCertificates(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Certificate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Certificate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Certificate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CertificateStatus != 0 {
		i = encodeVarintUserCertificates(dAtA, i, uint64(m.CertificateStatus))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintUserCertificates(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.AllowedAuthorities) > 0 {
		for iNdEx := len(m.AllowedAuthorities) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllowedAuthorities[iNdEx])
			copy(dAtA[i:], m.AllowedAuthorities[iNdEx])
			i = encodeVarintUserCertificates(dAtA, i, uint64(len(m.AllowedAuthorities[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.DeviceAddress) > 0 {
		i -= len(m.DeviceAddress)
		copy(dAtA[i:], m.DeviceAddress)
		i = encodeVarintUserCertificates(dAtA, i, uint64(len(m.DeviceAddress)))
		i--
		dAtA[i] = 0x22
	}
	if m.Power != 0 {
		i = encodeVarintUserCertificates(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x18
	}
	if m.CertyficateTypeId != 0 {
		i = encodeVarintUserCertificates(dAtA, i, uint64(m.CertyficateTypeId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintUserCertificates(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CertificateOffer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CertificateOffer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CertificateOffer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Price) > 0 {
		for iNdEx := len(m.Price) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Price[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUserCertificates(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Buyer) > 0 {
		i -= len(m.Buyer)
		copy(dAtA[i:], m.Buyer)
		i = encodeVarintUserCertificates(dAtA, i, uint64(len(m.Buyer)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintUserCertificates(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x1a
	}
	if m.CertificateId != 0 {
		i = encodeVarintUserCertificates(dAtA, i, uint64(m.CertificateId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintUserCertificates(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintUserCertificates(dAtA []byte, offset int, v uint64) int {
	offset -= sovUserCertificates(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UserCertificates) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovUserCertificates(uint64(l))
	}
	if len(m.Certificates) > 0 {
		for _, e := range m.Certificates {
			l = e.Size()
			n += 1 + l + sovUserCertificates(uint64(l))
		}
	}
	return n
}

func (m *Certificate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovUserCertificates(uint64(m.Id))
	}
	if m.CertyficateTypeId != 0 {
		n += 1 + sovUserCertificates(uint64(m.CertyficateTypeId))
	}
	if m.Power != 0 {
		n += 1 + sovUserCertificates(uint64(m.Power))
	}
	l = len(m.DeviceAddress)
	if l > 0 {
		n += 1 + l + sovUserCertificates(uint64(l))
	}
	if len(m.AllowedAuthorities) > 0 {
		for _, s := range m.AllowedAuthorities {
			l = len(s)
			n += 1 + l + sovUserCertificates(uint64(l))
		}
	}
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovUserCertificates(uint64(l))
	}
	if m.CertificateStatus != 0 {
		n += 1 + sovUserCertificates(uint64(m.CertificateStatus))
	}
	return n
}

func (m *CertificateOffer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovUserCertificates(uint64(m.Id))
	}
	if m.CertificateId != 0 {
		n += 1 + sovUserCertificates(uint64(m.CertificateId))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovUserCertificates(uint64(l))
	}
	l = len(m.Buyer)
	if l > 0 {
		n += 1 + l + sovUserCertificates(uint64(l))
	}
	if len(m.Price) > 0 {
		for _, e := range m.Price {
			l = e.Size()
			n += 1 + l + sovUserCertificates(uint64(l))
		}
	}
	return n
}

func sovUserCertificates(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUserCertificates(x uint64) (n int) {
	return sovUserCertificates(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserCertificates) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUserCertificates
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
			return fmt.Errorf("proto: UserCertificates: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserCertificates: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserCertificates
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
				return ErrInvalidLengthUserCertificates
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserCertificates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserCertificates
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
				return ErrInvalidLengthUserCertificates
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUserCertificates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Certificates = append(m.Certificates, &Certificate{})
			if err := m.Certificates[len(m.Certificates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUserCertificates(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUserCertificates
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
func (m *Certificate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUserCertificates
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
			return fmt.Errorf("proto: Certificate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Certificate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserCertificates
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertyficateTypeId", wireType)
			}
			m.CertyficateTypeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserCertificates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CertyficateTypeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserCertificates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserCertificates
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
				return ErrInvalidLengthUserCertificates
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserCertificates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedAuthorities", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserCertificates
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
				return ErrInvalidLengthUserCertificates
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserCertificates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowedAuthorities = append(m.AllowedAuthorities, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserCertificates
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
				return ErrInvalidLengthUserCertificates
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserCertificates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertificateStatus", wireType)
			}
			m.CertificateStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserCertificates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CertificateStatus |= CertificateStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUserCertificates(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUserCertificates
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
func (m *CertificateOffer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUserCertificates
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
			return fmt.Errorf("proto: CertificateOffer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CertificateOffer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserCertificates
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertificateId", wireType)
			}
			m.CertificateId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserCertificates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CertificateId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserCertificates
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
				return ErrInvalidLengthUserCertificates
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserCertificates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Buyer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserCertificates
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
				return ErrInvalidLengthUserCertificates
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserCertificates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Buyer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserCertificates
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
				return ErrInvalidLengthUserCertificates
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUserCertificates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = append(m.Price, types.Coin{})
			if err := m.Price[len(m.Price)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUserCertificates(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUserCertificates
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
func skipUserCertificates(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUserCertificates
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
					return 0, ErrIntOverflowUserCertificates
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
					return 0, ErrIntOverflowUserCertificates
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
				return 0, ErrInvalidLengthUserCertificates
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUserCertificates
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUserCertificates
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUserCertificates        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUserCertificates          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUserCertificates = fmt.Errorf("proto: unexpected end of group")
)
