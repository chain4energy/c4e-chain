// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: c4echain/cfetokenization/user_devices.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
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

type UserDevices struct {
	Owner   string        `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Devices []*UserDevice `protobuf:"bytes,2,rep,name=devices,proto3" json:"devices,omitempty"`
}

func (m *UserDevices) Reset()         { *m = UserDevices{} }
func (m *UserDevices) String() string { return proto.CompactTextString(m) }
func (*UserDevices) ProtoMessage()    {}
func (*UserDevices) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8f2c6e39947bc9b, []int{0}
}
func (m *UserDevices) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserDevices) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserDevices.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserDevices) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDevices.Merge(m, src)
}
func (m *UserDevices) XXX_Size() int {
	return m.Size()
}
func (m *UserDevices) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDevices.DiscardUnknown(m)
}

var xxx_messageInfo_UserDevices proto.InternalMessageInfo

func (m *UserDevices) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *UserDevices) GetDevices() []*UserDevice {
	if m != nil {
		return m.Devices
	}
	return nil
}

type UserDevice struct {
	DeviceAddress string `protobuf:"bytes,1,opt,name=device_address,json=deviceAddress,proto3" json:"device_address,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Location      string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
}

func (m *UserDevice) Reset()         { *m = UserDevice{} }
func (m *UserDevice) String() string { return proto.CompactTextString(m) }
func (*UserDevice) ProtoMessage()    {}
func (*UserDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8f2c6e39947bc9b, []int{1}
}
func (m *UserDevice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserDevice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDevice.Merge(m, src)
}
func (m *UserDevice) XXX_Size() int {
	return m.Size()
}
func (m *UserDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDevice.DiscardUnknown(m)
}

var xxx_messageInfo_UserDevice proto.InternalMessageInfo

func (m *UserDevice) GetDeviceAddress() string {
	if m != nil {
		return m.DeviceAddress
	}
	return ""
}

func (m *UserDevice) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserDevice) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

type PendingDevice struct {
	DeviceAddress string `protobuf:"bytes,1,opt,name=device_address,json=deviceAddress,proto3" json:"device_address,omitempty"`
	UserAddress   string `protobuf:"bytes,2,opt,name=user_address,json=userAddress,proto3" json:"user_address,omitempty"`
}

func (m *PendingDevice) Reset()         { *m = PendingDevice{} }
func (m *PendingDevice) String() string { return proto.CompactTextString(m) }
func (*PendingDevice) ProtoMessage()    {}
func (*PendingDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8f2c6e39947bc9b, []int{2}
}
func (m *PendingDevice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PendingDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PendingDevice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PendingDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PendingDevice.Merge(m, src)
}
func (m *PendingDevice) XXX_Size() int {
	return m.Size()
}
func (m *PendingDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_PendingDevice.DiscardUnknown(m)
}

var xxx_messageInfo_PendingDevice proto.InternalMessageInfo

func (m *PendingDevice) GetDeviceAddress() string {
	if m != nil {
		return m.DeviceAddress
	}
	return ""
}

func (m *PendingDevice) GetUserAddress() string {
	if m != nil {
		return m.UserAddress
	}
	return ""
}

type Device struct {
	DeviceAddress   string         `protobuf:"bytes,1,opt,name=device_address,json=deviceAddress,proto3" json:"device_address,omitempty"`
	Measurements    []*Measurement `protobuf:"bytes,2,rep,name=measurements,proto3" json:"measurements,omitempty"`
	ActivePowerSum  uint64         `protobuf:"varint,3,opt,name=active_power_sum,json=activePowerSum,proto3" json:"active_power_sum,omitempty"`
	ReversePowerSum uint64         `protobuf:"varint,4,opt,name=reverse_power_sum,json=reversePowerSum,proto3" json:"reverse_power_sum,omitempty"`
	UsedActivePower uint64         `protobuf:"varint,5,opt,name=used_active_power,json=usedActivePower,proto3" json:"used_active_power,omitempty"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8f2c6e39947bc9b, []int{3}
}
func (m *Device) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Device.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(m, src)
}
func (m *Device) XXX_Size() int {
	return m.Size()
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

func (m *Device) GetDeviceAddress() string {
	if m != nil {
		return m.DeviceAddress
	}
	return ""
}

func (m *Device) GetMeasurements() []*Measurement {
	if m != nil {
		return m.Measurements
	}
	return nil
}

func (m *Device) GetActivePowerSum() uint64 {
	if m != nil {
		return m.ActivePowerSum
	}
	return 0
}

func (m *Device) GetReversePowerSum() uint64 {
	if m != nil {
		return m.ReversePowerSum
	}
	return 0
}

func (m *Device) GetUsedActivePower() uint64 {
	if m != nil {
		return m.UsedActivePower
	}
	return 0
}

type Measurement struct {
	Timestamp    time.Time `protobuf:"bytes,1,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	ActivePower  uint64    `protobuf:"varint,2,opt,name=active_power,json=activePower,proto3" json:"active_power,omitempty"`
	ReversePower uint64    `protobuf:"varint,3,opt,name=reverse_power,json=reversePower,proto3" json:"reverse_power,omitempty"`
	Metadata     string    `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *Measurement) Reset()         { *m = Measurement{} }
func (m *Measurement) String() string { return proto.CompactTextString(m) }
func (*Measurement) ProtoMessage()    {}
func (*Measurement) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8f2c6e39947bc9b, []int{4}
}
func (m *Measurement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Measurement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Measurement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Measurement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Measurement.Merge(m, src)
}
func (m *Measurement) XXX_Size() int {
	return m.Size()
}
func (m *Measurement) XXX_DiscardUnknown() {
	xxx_messageInfo_Measurement.DiscardUnknown(m)
}

var xxx_messageInfo_Measurement proto.InternalMessageInfo

func (m *Measurement) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *Measurement) GetActivePower() uint64 {
	if m != nil {
		return m.ActivePower
	}
	return 0
}

func (m *Measurement) GetReversePower() uint64 {
	if m != nil {
		return m.ReversePower
	}
	return 0
}

func (m *Measurement) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

func init() {
	proto.RegisterType((*UserDevices)(nil), "chain4energy.c4echain.cfetokenization.UserDevices")
	proto.RegisterType((*UserDevice)(nil), "chain4energy.c4echain.cfetokenization.UserDevice")
	proto.RegisterType((*PendingDevice)(nil), "chain4energy.c4echain.cfetokenization.PendingDevice")
	proto.RegisterType((*Device)(nil), "chain4energy.c4echain.cfetokenization.Device")
	proto.RegisterType((*Measurement)(nil), "chain4energy.c4echain.cfetokenization.Measurement")
}

func init() {
	proto.RegisterFile("c4echain/cfetokenization/user_devices.proto", fileDescriptor_a8f2c6e39947bc9b)
}

var fileDescriptor_a8f2c6e39947bc9b = []byte{
	// 487 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x8d, 0xd3, 0x07, 0xcd, 0x75, 0x52, 0x60, 0xd4, 0x85, 0x95, 0x85, 0x53, 0x8c, 0x2a, 0x45,
	0x20, 0x6c, 0x11, 0xb2, 0x61, 0xd9, 0x88, 0x1d, 0x42, 0xaa, 0x5c, 0x40, 0x82, 0x8d, 0x35, 0xb1,
	0x6f, 0x5d, 0x8b, 0xda, 0x63, 0xcd, 0x8c, 0x53, 0xca, 0x1f, 0xb0, 0xeb, 0xcf, 0xf0, 0x0f, 0x5d,
	0x76, 0xc9, 0x0a, 0x50, 0xf2, 0x23, 0xc8, 0x33, 0x76, 0xec, 0xc0, 0xa6, 0xdd, 0xcd, 0x9c, 0x7b,
	0xee, 0x39, 0x73, 0x1f, 0x03, 0xcf, 0xc3, 0x29, 0x86, 0xe7, 0x34, 0xc9, 0xbc, 0xf0, 0x0c, 0x25,
	0xfb, 0x82, 0x59, 0xf2, 0x8d, 0xca, 0x84, 0x65, 0x5e, 0x21, 0x90, 0x07, 0x11, 0x2e, 0x92, 0x10,
	0x85, 0x9b, 0x73, 0x26, 0x19, 0x39, 0x52, 0xcc, 0x29, 0x66, 0xc8, 0xe3, 0x2b, 0xb7, 0xce, 0x74,
	0xff, 0xc9, 0x1c, 0x8e, 0x62, 0xc6, 0xe2, 0x0b, 0xf4, 0x54, 0xd2, 0xbc, 0x38, 0xf3, 0x64, 0x92,
	0xa2, 0x90, 0x34, 0xcd, 0xb5, 0xce, 0xf0, 0x20, 0x66, 0x31, 0x53, 0x47, 0xaf, 0x3c, 0x69, 0xd4,
	0xc9, 0xc1, 0xfc, 0x20, 0x90, 0xbf, 0xd1, 0x96, 0xe4, 0x00, 0x76, 0xd8, 0x65, 0x86, 0xdc, 0x32,
	0x0e, 0x8d, 0x71, 0xcf, 0xd7, 0x17, 0xf2, 0x16, 0x1e, 0x54, 0x6f, 0xb2, 0xba, 0x87, 0x5b, 0x63,
	0x73, 0xf2, 0xd2, 0xbd, 0xd3, 0xa3, 0xdc, 0x46, 0xda, 0xaf, 0x15, 0x9c, 0x10, 0xa0, 0x81, 0xc9,
	0x11, 0xec, 0xeb, 0x40, 0x40, 0xa3, 0x88, 0xa3, 0x10, 0x95, 0xf3, 0x40, 0xa3, 0xc7, 0x1a, 0x24,
	0x04, 0xb6, 0x33, 0x9a, 0xa2, 0xd5, 0x55, 0x41, 0x75, 0x26, 0x43, 0xd8, 0xbb, 0x60, 0xa1, 0x32,
	0xb2, 0xb6, 0x14, 0xbe, 0xbe, 0x3b, 0x9f, 0x60, 0x70, 0x82, 0x59, 0x94, 0x64, 0xf1, 0xfd, 0x7c,
	0x9e, 0x40, 0x5f, 0x8d, 0xa0, 0x26, 0x69, 0x3f, 0xb3, 0xc4, 0x2a, 0x8a, 0xf3, 0xbd, 0x0b, 0xbb,
	0xf7, 0x13, 0xfd, 0x08, 0xfd, 0x14, 0xa9, 0x28, 0x38, 0xa6, 0x98, 0xc9, 0xba, 0x87, 0x93, 0x3b,
	0xf6, 0xf0, 0x5d, 0x93, 0xea, 0x6f, 0xe8, 0x90, 0x31, 0x3c, 0xa2, 0xa1, 0x4c, 0x16, 0x18, 0xe4,
	0xec, 0x12, 0x79, 0x20, 0x8a, 0x54, 0x35, 0x62, 0xdb, 0xdf, 0xd7, 0xf8, 0x49, 0x09, 0x9f, 0x16,
	0x29, 0x79, 0x06, 0x8f, 0x39, 0x2e, 0x90, 0x8b, 0x36, 0x75, 0x5b, 0x51, 0x1f, 0x56, 0x81, 0x36,
	0xb7, 0x10, 0x18, 0x05, 0x6d, 0x69, 0x6b, 0x47, 0x73, 0xcb, 0xc0, 0x71, 0x23, 0xed, 0xfc, 0x30,
	0xc0, 0x6c, 0xbd, 0x8f, 0xcc, 0xa0, 0xb7, 0x5e, 0x3b, 0xd5, 0x0b, 0x73, 0x32, 0x74, 0xf5, 0x62,
	0xba, 0xf5, 0x62, 0xba, 0xef, 0x6b, 0xc6, 0x6c, 0xef, 0xe6, 0xd7, 0xa8, 0x73, 0xfd, 0x7b, 0x64,
	0xf8, 0x4d, 0x5a, 0x39, 0x82, 0x0d, 0xeb, 0xae, 0xb2, 0x36, 0x5b, 0x15, 0x91, 0xa7, 0x30, 0xd8,
	0x28, 0xa7, 0xaa, 0xba, 0xdf, 0x2e, 0xa5, 0x5c, 0x8f, 0x14, 0x25, 0x8d, 0xa8, 0xa4, 0xaa, 0xd4,
	0x9e, 0xbf, 0xbe, 0xcf, 0x4e, 0x6f, 0x96, 0xb6, 0x71, 0xbb, 0xb4, 0x8d, 0x3f, 0x4b, 0xdb, 0xb8,
	0x5e, 0xd9, 0x9d, 0xdb, 0x95, 0xdd, 0xf9, 0xb9, 0xb2, 0x3b, 0x9f, 0x5f, 0xc7, 0x89, 0x3c, 0x2f,
	0xe6, 0x6e, 0xc8, 0x52, 0xaf, 0x3d, 0x1f, 0x2f, 0x9c, 0xe2, 0x0b, 0xfd, 0x67, 0xbf, 0xfe, 0xf7,
	0x6b, 0xe5, 0x55, 0x8e, 0x62, 0xbe, 0xab, 0x2a, 0x7c, 0xf5, 0x37, 0x00, 0x00, 0xff, 0xff, 0x12,
	0x8a, 0xc5, 0x75, 0xde, 0x03, 0x00, 0x00,
}

func (m *UserDevices) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserDevices) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserDevices) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Devices) > 0 {
		for iNdEx := len(m.Devices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Devices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUserDevices(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintUserDevices(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UserDevice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserDevice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserDevice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Location) > 0 {
		i -= len(m.Location)
		copy(dAtA[i:], m.Location)
		i = encodeVarintUserDevices(dAtA, i, uint64(len(m.Location)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintUserDevices(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DeviceAddress) > 0 {
		i -= len(m.DeviceAddress)
		copy(dAtA[i:], m.DeviceAddress)
		i = encodeVarintUserDevices(dAtA, i, uint64(len(m.DeviceAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PendingDevice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PendingDevice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PendingDevice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.UserAddress) > 0 {
		i -= len(m.UserAddress)
		copy(dAtA[i:], m.UserAddress)
		i = encodeVarintUserDevices(dAtA, i, uint64(len(m.UserAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DeviceAddress) > 0 {
		i -= len(m.DeviceAddress)
		copy(dAtA[i:], m.DeviceAddress)
		i = encodeVarintUserDevices(dAtA, i, uint64(len(m.DeviceAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Device) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Device) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Device) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UsedActivePower != 0 {
		i = encodeVarintUserDevices(dAtA, i, uint64(m.UsedActivePower))
		i--
		dAtA[i] = 0x28
	}
	if m.ReversePowerSum != 0 {
		i = encodeVarintUserDevices(dAtA, i, uint64(m.ReversePowerSum))
		i--
		dAtA[i] = 0x20
	}
	if m.ActivePowerSum != 0 {
		i = encodeVarintUserDevices(dAtA, i, uint64(m.ActivePowerSum))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Measurements) > 0 {
		for iNdEx := len(m.Measurements) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Measurements[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUserDevices(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.DeviceAddress) > 0 {
		i -= len(m.DeviceAddress)
		copy(dAtA[i:], m.DeviceAddress)
		i = encodeVarintUserDevices(dAtA, i, uint64(len(m.DeviceAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Measurement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Measurement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Measurement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Metadata) > 0 {
		i -= len(m.Metadata)
		copy(dAtA[i:], m.Metadata)
		i = encodeVarintUserDevices(dAtA, i, uint64(len(m.Metadata)))
		i--
		dAtA[i] = 0x22
	}
	if m.ReversePower != 0 {
		i = encodeVarintUserDevices(dAtA, i, uint64(m.ReversePower))
		i--
		dAtA[i] = 0x18
	}
	if m.ActivePower != 0 {
		i = encodeVarintUserDevices(dAtA, i, uint64(m.ActivePower))
		i--
		dAtA[i] = 0x10
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintUserDevices(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintUserDevices(dAtA []byte, offset int, v uint64) int {
	offset -= sovUserDevices(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UserDevices) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovUserDevices(uint64(l))
	}
	if len(m.Devices) > 0 {
		for _, e := range m.Devices {
			l = e.Size()
			n += 1 + l + sovUserDevices(uint64(l))
		}
	}
	return n
}

func (m *UserDevice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DeviceAddress)
	if l > 0 {
		n += 1 + l + sovUserDevices(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUserDevices(uint64(l))
	}
	l = len(m.Location)
	if l > 0 {
		n += 1 + l + sovUserDevices(uint64(l))
	}
	return n
}

func (m *PendingDevice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DeviceAddress)
	if l > 0 {
		n += 1 + l + sovUserDevices(uint64(l))
	}
	l = len(m.UserAddress)
	if l > 0 {
		n += 1 + l + sovUserDevices(uint64(l))
	}
	return n
}

func (m *Device) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DeviceAddress)
	if l > 0 {
		n += 1 + l + sovUserDevices(uint64(l))
	}
	if len(m.Measurements) > 0 {
		for _, e := range m.Measurements {
			l = e.Size()
			n += 1 + l + sovUserDevices(uint64(l))
		}
	}
	if m.ActivePowerSum != 0 {
		n += 1 + sovUserDevices(uint64(m.ActivePowerSum))
	}
	if m.ReversePowerSum != 0 {
		n += 1 + sovUserDevices(uint64(m.ReversePowerSum))
	}
	if m.UsedActivePower != 0 {
		n += 1 + sovUserDevices(uint64(m.UsedActivePower))
	}
	return n
}

func (m *Measurement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovUserDevices(uint64(l))
	if m.ActivePower != 0 {
		n += 1 + sovUserDevices(uint64(m.ActivePower))
	}
	if m.ReversePower != 0 {
		n += 1 + sovUserDevices(uint64(m.ReversePower))
	}
	l = len(m.Metadata)
	if l > 0 {
		n += 1 + l + sovUserDevices(uint64(l))
	}
	return n
}

func sovUserDevices(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUserDevices(x uint64) (n int) {
	return sovUserDevices(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserDevices) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUserDevices
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
			return fmt.Errorf("proto: UserDevices: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserDevices: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserDevices
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
				return ErrInvalidLengthUserDevices
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserDevices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Devices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserDevices
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
				return ErrInvalidLengthUserDevices
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUserDevices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Devices = append(m.Devices, &UserDevice{})
			if err := m.Devices[len(m.Devices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUserDevices(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUserDevices
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
func (m *UserDevice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUserDevices
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
			return fmt.Errorf("proto: UserDevice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserDevice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserDevices
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
				return ErrInvalidLengthUserDevices
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserDevices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserDevices
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
				return ErrInvalidLengthUserDevices
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserDevices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Location", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserDevices
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
				return ErrInvalidLengthUserDevices
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserDevices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Location = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUserDevices(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUserDevices
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
func (m *PendingDevice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUserDevices
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
			return fmt.Errorf("proto: PendingDevice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PendingDevice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserDevices
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
				return ErrInvalidLengthUserDevices
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserDevices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserDevices
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
				return ErrInvalidLengthUserDevices
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserDevices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUserDevices(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUserDevices
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
func (m *Device) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUserDevices
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
			return fmt.Errorf("proto: Device: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Device: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserDevices
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
				return ErrInvalidLengthUserDevices
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserDevices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Measurements", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserDevices
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
				return ErrInvalidLengthUserDevices
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUserDevices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Measurements = append(m.Measurements, &Measurement{})
			if err := m.Measurements[len(m.Measurements)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActivePowerSum", wireType)
			}
			m.ActivePowerSum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserDevices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ActivePowerSum |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReversePowerSum", wireType)
			}
			m.ReversePowerSum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserDevices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReversePowerSum |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UsedActivePower", wireType)
			}
			m.UsedActivePower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserDevices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UsedActivePower |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUserDevices(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUserDevices
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
func (m *Measurement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUserDevices
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
			return fmt.Errorf("proto: Measurement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Measurement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserDevices
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
				return ErrInvalidLengthUserDevices
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUserDevices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActivePower", wireType)
			}
			m.ActivePower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserDevices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ActivePower |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReversePower", wireType)
			}
			m.ReversePower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserDevices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReversePower |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserDevices
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
				return ErrInvalidLengthUserDevices
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserDevices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUserDevices(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUserDevices
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
func skipUserDevices(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUserDevices
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
					return 0, ErrIntOverflowUserDevices
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
					return 0, ErrIntOverflowUserDevices
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
				return 0, ErrInvalidLengthUserDevices
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUserDevices
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUserDevices
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUserDevices        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUserDevices          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUserDevices = fmt.Errorf("proto: unexpected end of group")
)
