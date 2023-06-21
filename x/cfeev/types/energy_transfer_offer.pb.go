// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: c4echain/cfeev/energy_transfer_offer.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type ChargerStatus int32

const (
	ChargerStatus_ACTIVE                     ChargerStatus = 0
	ChargerStatus_BUSY                       ChargerStatus = 1
	ChargerStatus_INACTIVE                   ChargerStatus = 2
	ChargerStatus_CHARGER_STATUS_UNSPECIFIED ChargerStatus = 3
)

var ChargerStatus_name = map[int32]string{
	0: "ACTIVE",
	1: "BUSY",
	2: "INACTIVE",
	3: "CHARGER_STATUS_UNSPECIFIED",
}

var ChargerStatus_value = map[string]int32{
	"ACTIVE":                     0,
	"BUSY":                       1,
	"INACTIVE":                   2,
	"CHARGER_STATUS_UNSPECIFIED": 3,
}

func (x ChargerStatus) String() string {
	return proto.EnumName(ChargerStatus_name, int32(x))
}

func (ChargerStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_60e53b5c064f8068, []int{0}
}

type PlugType int32

const (
	PlugType_Type1                 PlugType = 0
	PlugType_Type2                 PlugType = 1
	PlugType_CHAdeMO               PlugType = 2
	PlugType_CCS                   PlugType = 3
	PlugType_PLUG_TYPE_UNSPECIFIED PlugType = 4
)

var PlugType_name = map[int32]string{
	0: "Type1",
	1: "Type2",
	2: "CHAdeMO",
	3: "CCS",
	4: "PLUG_TYPE_UNSPECIFIED",
}

var PlugType_value = map[string]int32{
	"Type1":                 0,
	"Type2":                 1,
	"CHAdeMO":               2,
	"CCS":                   3,
	"PLUG_TYPE_UNSPECIFIED": 4,
}

func (x PlugType) String() string {
	return proto.EnumName(PlugType_name, int32(x))
}

func (PlugType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_60e53b5c064f8068, []int{1}
}

type EnergyTransferOffer struct {
	Id            uint64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner         string        `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	ChargerId     string        `protobuf:"bytes,3,opt,name=chargerId,proto3" json:"chargerId,omitempty"`
	ChargerStatus ChargerStatus `protobuf:"varint,4,opt,name=charger_status,json=chargerStatus,proto3,enum=chain4energy.c4echain.cfeev.ChargerStatus" json:"charger_status,omitempty"`
	Location      *Location     `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	Tariff        uint64        `protobuf:"varint,6,opt,name=tariff,proto3" json:"tariff,omitempty"`
	Name          string        `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	PlugType      PlugType      `protobuf:"varint,8,opt,name=plug_type,json=plugType,proto3,enum=chain4energy.c4echain.cfeev.PlugType" json:"plug_type,omitempty"`
}

func (m *EnergyTransferOffer) Reset()         { *m = EnergyTransferOffer{} }
func (m *EnergyTransferOffer) String() string { return proto.CompactTextString(m) }
func (*EnergyTransferOffer) ProtoMessage()    {}
func (*EnergyTransferOffer) Descriptor() ([]byte, []int) {
	return fileDescriptor_60e53b5c064f8068, []int{0}
}
func (m *EnergyTransferOffer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EnergyTransferOffer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EnergyTransferOffer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EnergyTransferOffer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnergyTransferOffer.Merge(m, src)
}
func (m *EnergyTransferOffer) XXX_Size() int {
	return m.Size()
}
func (m *EnergyTransferOffer) XXX_DiscardUnknown() {
	xxx_messageInfo_EnergyTransferOffer.DiscardUnknown(m)
}

var xxx_messageInfo_EnergyTransferOffer proto.InternalMessageInfo

func (m *EnergyTransferOffer) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EnergyTransferOffer) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *EnergyTransferOffer) GetChargerId() string {
	if m != nil {
		return m.ChargerId
	}
	return ""
}

func (m *EnergyTransferOffer) GetChargerStatus() ChargerStatus {
	if m != nil {
		return m.ChargerStatus
	}
	return ChargerStatus_ACTIVE
}

func (m *EnergyTransferOffer) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *EnergyTransferOffer) GetTariff() uint64 {
	if m != nil {
		return m.Tariff
	}
	return 0
}

func (m *EnergyTransferOffer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EnergyTransferOffer) GetPlugType() PlugType {
	if m != nil {
		return m.PlugType
	}
	return PlugType_Type1
}

type Location struct {
	Latitude  *github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=latitude,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"latitude,omitempty"`
	Longitude *github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=longitude,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"longitude,omitempty"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_60e53b5c064f8068, []int{1}
}
func (m *Location) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Location.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return m.Size()
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("chain4energy.c4echain.cfeev.ChargerStatus", ChargerStatus_name, ChargerStatus_value)
	proto.RegisterEnum("chain4energy.c4echain.cfeev.PlugType", PlugType_name, PlugType_value)
	proto.RegisterType((*EnergyTransferOffer)(nil), "chain4energy.c4echain.cfeev.EnergyTransferOffer")
	proto.RegisterType((*Location)(nil), "chain4energy.c4echain.cfeev.Location")
}

func init() {
	proto.RegisterFile("c4echain/cfeev/energy_transfer_offer.proto", fileDescriptor_60e53b5c064f8068)
}

var fileDescriptor_60e53b5c064f8068 = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x41, 0x6e, 0x9b, 0x40,
	0x14, 0xf5, 0x60, 0xc7, 0x86, 0x9f, 0xc6, 0x42, 0xd3, 0xb4, 0xa2, 0x69, 0x45, 0xac, 0x48, 0xad,
	0x2c, 0x4b, 0x01, 0xd5, 0xf5, 0x05, 0x6c, 0x42, 0x13, 0x2a, 0x37, 0x71, 0x00, 0x57, 0x4a, 0x37,
	0x88, 0xc0, 0x80, 0x51, 0x6d, 0xc6, 0x02, 0xdc, 0xd6, 0xb7, 0xe8, 0x21, 0x7a, 0x98, 0x2c, 0xb3,
	0xac, 0xba, 0x48, 0x2b, 0xfb, 0x22, 0x95, 0x07, 0x9c, 0x38, 0x1b, 0x2f, 0xba, 0x9a, 0xff, 0x87,
	0xf7, 0x1e, 0xef, 0xbf, 0xd1, 0x87, 0x96, 0xd7, 0x21, 0xde, 0xc8, 0x8d, 0x62, 0xd5, 0x0b, 0x08,
	0xf9, 0xaa, 0x92, 0x98, 0x24, 0xe1, 0xdc, 0xc9, 0x12, 0x37, 0x4e, 0x03, 0x92, 0x38, 0x34, 0x08,
	0x48, 0xa2, 0x4c, 0x13, 0x9a, 0x51, 0xfc, 0x92, 0x01, 0x3b, 0x39, 0x44, 0x59, 0x13, 0x15, 0x46,
	0x3c, 0xd8, 0x0f, 0x69, 0x48, 0x19, 0x4e, 0x5d, 0x55, 0x39, 0xe5, 0xe8, 0x0f, 0x07, 0x4f, 0x75,
	0x86, 0xb7, 0x0b, 0xc5, 0x8b, 0x95, 0x20, 0xae, 0x03, 0x17, 0xf9, 0x12, 0x6a, 0xa0, 0x66, 0xc5,
	0xe4, 0x22, 0x1f, 0xef, 0xc3, 0x0e, 0xfd, 0x16, 0x93, 0x44, 0xe2, 0x1a, 0xa8, 0x29, 0x98, 0x79,
	0x83, 0x5f, 0x81, 0xe0, 0x8d, 0xdc, 0x24, 0x24, 0x89, 0xe1, 0x4b, 0x65, 0xf6, 0xe5, 0xe1, 0x02,
	0x5f, 0x42, 0xbd, 0x68, 0x9c, 0x34, 0x73, 0xb3, 0x59, 0x2a, 0x55, 0x1a, 0xa8, 0x59, 0x6f, 0xb7,
	0x94, 0x2d, 0x3e, 0x15, 0x2d, 0xa7, 0x58, 0x8c, 0x61, 0xee, 0x79, 0x9b, 0x2d, 0xee, 0x02, 0x3f,
	0xa6, 0x9e, 0x9b, 0x45, 0x34, 0x96, 0x76, 0x1a, 0xa8, 0xb9, 0xdb, 0x7e, 0xbd, 0x55, 0xac, 0x5f,
	0x80, 0xcd, 0x7b, 0x1a, 0x7e, 0x0e, 0xd5, 0xcc, 0x4d, 0xa2, 0x20, 0x90, 0xaa, 0x6c, 0xba, 0xa2,
	0xc3, 0x18, 0x2a, 0xb1, 0x3b, 0x21, 0x52, 0x8d, 0x8d, 0xc1, 0x6a, 0xdc, 0x03, 0x61, 0x3a, 0x9e,
	0x85, 0x4e, 0x36, 0x9f, 0x12, 0x89, 0x67, 0xe6, 0xb7, 0xff, 0x6f, 0x30, 0x9e, 0x85, 0xf6, 0x7c,
	0x4a, 0x4c, 0x7e, 0x5a, 0x54, 0x47, 0x3f, 0x11, 0xf0, 0x6b, 0x1b, 0xf8, 0x03, 0xf0, 0x63, 0x37,
	0x8b, 0xb2, 0x99, 0x4f, 0x58, 0xb8, 0x42, 0x4f, 0xb9, 0xb9, 0x3b, 0x44, 0xbf, 0xef, 0x0e, 0xdf,
	0x84, 0x51, 0x36, 0x9a, 0x5d, 0x2b, 0x1e, 0x9d, 0xa8, 0x1e, 0x4d, 0x27, 0x34, 0x2d, 0x8e, 0xe3,
	0xd4, 0xff, 0xa2, 0xae, 0x0c, 0xa4, 0xca, 0x09, 0xf1, 0xcc, 0x7b, 0x3e, 0xee, 0x83, 0x30, 0xa6,
	0x71, 0x98, 0x8b, 0x71, 0xff, 0x25, 0xf6, 0x20, 0xd0, 0xb2, 0x60, 0xef, 0x51, 0xf2, 0x18, 0xa0,
	0xda, 0xd5, 0x6c, 0xe3, 0x93, 0x2e, 0x96, 0x30, 0x0f, 0x95, 0xde, 0xd0, 0xba, 0x12, 0x11, 0x7e,
	0x02, 0xbc, 0x71, 0x5e, 0xdc, 0x73, 0x58, 0x86, 0x03, 0xed, 0xac, 0x6b, 0x9e, 0xea, 0xa6, 0x63,
	0xd9, 0x5d, 0x7b, 0x68, 0x39, 0xc3, 0x73, 0x6b, 0xa0, 0x6b, 0xc6, 0x7b, 0x43, 0x3f, 0x11, 0xcb,
	0xad, 0x4b, 0xe0, 0xd7, 0x89, 0x60, 0x01, 0x76, 0x56, 0xe7, 0x5b, 0xb1, 0xb4, 0x2e, 0xdb, 0x22,
	0xc2, 0xbb, 0x50, 0xd3, 0xce, 0xba, 0x3e, 0xf9, 0x78, 0x21, 0x72, 0xb8, 0x06, 0x65, 0x4d, 0xb3,
	0xc4, 0x32, 0x7e, 0x01, 0xcf, 0x06, 0xfd, 0xe1, 0xa9, 0x63, 0x5f, 0x0d, 0xf4, 0x47, 0x92, 0x95,
	0x9e, 0x71, 0xb3, 0x90, 0xd1, 0xed, 0x42, 0x46, 0x7f, 0x17, 0x32, 0xfa, 0xb1, 0x94, 0x4b, 0xb7,
	0x4b, 0xb9, 0xf4, 0x6b, 0x29, 0x97, 0x3e, 0xab, 0x9b, 0x43, 0x6f, 0xbc, 0x91, 0xea, 0x75, 0xc8,
	0x71, 0xbe, 0x42, 0xdf, 0x8b, 0x25, 0x62, 0x09, 0x5c, 0x57, 0xd9, 0x0a, 0xbc, 0xfb, 0x17, 0x00,
	0x00, 0xff, 0xff, 0x62, 0x65, 0xc9, 0x58, 0x63, 0x03, 0x00, 0x00,
}

func (m *EnergyTransferOffer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnergyTransferOffer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EnergyTransferOffer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PlugType != 0 {
		i = encodeVarintEnergyTransferOffer(dAtA, i, uint64(m.PlugType))
		i--
		dAtA[i] = 0x40
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintEnergyTransferOffer(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Tariff != 0 {
		i = encodeVarintEnergyTransferOffer(dAtA, i, uint64(m.Tariff))
		i--
		dAtA[i] = 0x30
	}
	if m.Location != nil {
		{
			size, err := m.Location.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEnergyTransferOffer(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.ChargerStatus != 0 {
		i = encodeVarintEnergyTransferOffer(dAtA, i, uint64(m.ChargerStatus))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ChargerId) > 0 {
		i -= len(m.ChargerId)
		copy(dAtA[i:], m.ChargerId)
		i = encodeVarintEnergyTransferOffer(dAtA, i, uint64(len(m.ChargerId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintEnergyTransferOffer(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintEnergyTransferOffer(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Location) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Location) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Location) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Longitude != nil {
		{
			size := m.Longitude.Size()
			i -= size
			if _, err := m.Longitude.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEnergyTransferOffer(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Latitude != nil {
		{
			size := m.Latitude.Size()
			i -= size
			if _, err := m.Latitude.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEnergyTransferOffer(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEnergyTransferOffer(dAtA []byte, offset int, v uint64) int {
	offset -= sovEnergyTransferOffer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EnergyTransferOffer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovEnergyTransferOffer(uint64(m.Id))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovEnergyTransferOffer(uint64(l))
	}
	l = len(m.ChargerId)
	if l > 0 {
		n += 1 + l + sovEnergyTransferOffer(uint64(l))
	}
	if m.ChargerStatus != 0 {
		n += 1 + sovEnergyTransferOffer(uint64(m.ChargerStatus))
	}
	if m.Location != nil {
		l = m.Location.Size()
		n += 1 + l + sovEnergyTransferOffer(uint64(l))
	}
	if m.Tariff != 0 {
		n += 1 + sovEnergyTransferOffer(uint64(m.Tariff))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovEnergyTransferOffer(uint64(l))
	}
	if m.PlugType != 0 {
		n += 1 + sovEnergyTransferOffer(uint64(m.PlugType))
	}
	return n
}

func (m *Location) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Latitude != nil {
		l = m.Latitude.Size()
		n += 1 + l + sovEnergyTransferOffer(uint64(l))
	}
	if m.Longitude != nil {
		l = m.Longitude.Size()
		n += 1 + l + sovEnergyTransferOffer(uint64(l))
	}
	return n
}

func sovEnergyTransferOffer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEnergyTransferOffer(x uint64) (n int) {
	return sovEnergyTransferOffer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EnergyTransferOffer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnergyTransferOffer
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
			return fmt.Errorf("proto: EnergyTransferOffer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EnergyTransferOffer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnergyTransferOffer
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
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnergyTransferOffer
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
				return ErrInvalidLengthEnergyTransferOffer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEnergyTransferOffer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChargerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnergyTransferOffer
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
				return ErrInvalidLengthEnergyTransferOffer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEnergyTransferOffer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChargerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChargerStatus", wireType)
			}
			m.ChargerStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnergyTransferOffer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChargerStatus |= ChargerStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Location", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnergyTransferOffer
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
				return ErrInvalidLengthEnergyTransferOffer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEnergyTransferOffer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Location == nil {
				m.Location = &Location{}
			}
			if err := m.Location.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tariff", wireType)
			}
			m.Tariff = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnergyTransferOffer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Tariff |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnergyTransferOffer
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
				return ErrInvalidLengthEnergyTransferOffer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEnergyTransferOffer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlugType", wireType)
			}
			m.PlugType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnergyTransferOffer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PlugType |= PlugType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEnergyTransferOffer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEnergyTransferOffer
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
func (m *Location) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnergyTransferOffer
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
			return fmt.Errorf("proto: Location: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Location: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Latitude", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnergyTransferOffer
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
				return ErrInvalidLengthEnergyTransferOffer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEnergyTransferOffer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.Latitude = &v
			if err := m.Latitude.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Longitude", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnergyTransferOffer
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
				return ErrInvalidLengthEnergyTransferOffer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEnergyTransferOffer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.Longitude = &v
			if err := m.Longitude.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEnergyTransferOffer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEnergyTransferOffer
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
func skipEnergyTransferOffer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEnergyTransferOffer
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
					return 0, ErrIntOverflowEnergyTransferOffer
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
					return 0, ErrIntOverflowEnergyTransferOffer
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
				return 0, ErrInvalidLengthEnergyTransferOffer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEnergyTransferOffer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEnergyTransferOffer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEnergyTransferOffer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEnergyTransferOffer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEnergyTransferOffer = fmt.Errorf("proto: unexpected end of group")
)
