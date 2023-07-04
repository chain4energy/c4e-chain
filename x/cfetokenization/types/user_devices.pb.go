// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: c4echain/cfetokenization/user_devices.proto

package types

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

type UserDevices struct {
	Id      uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Devices []*Device `protobuf:"bytes,2,rep,name=devices,proto3" json:"devices,omitempty"`
	Owner   string    `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
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

func (m *UserDevices) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserDevices) GetDevices() []*Device {
	if m != nil {
		return m.Devices
	}
	return nil
}

func (m *UserDevices) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type Device struct {
	DeviceId      string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DeviceAddress string `protobuf:"bytes,3,opt,name=device_address,json=deviceAddress,proto3" json:"device_address,omitempty"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8f2c6e39947bc9b, []int{1}
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

func (m *Device) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *Device) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Device) GetDeviceAddress() string {
	if m != nil {
		return m.DeviceAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*UserDevices)(nil), "chain4energy.c4echain.cfetokenization.UserDevices")
	proto.RegisterType((*Device)(nil), "chain4energy.c4echain.cfetokenization.Device")
}

func init() {
	proto.RegisterFile("c4echain/cfetokenization/user_devices.proto", fileDescriptor_a8f2c6e39947bc9b)
}

var fileDescriptor_a8f2c6e39947bc9b = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4e, 0x36, 0x49, 0x4d,
	0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x4f, 0x4e, 0x4b, 0x2d, 0xc9, 0xcf, 0x4e, 0xcd, 0xcb, 0xac, 0x4a,
	0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2f, 0x2d, 0x4e, 0x2d, 0x8a, 0x4f, 0x49, 0x2d, 0xcb, 0x4c, 0x4e,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x05, 0xab, 0x34, 0x49, 0xcd, 0x4b, 0x2d,
	0x4a, 0xaf, 0xd4, 0x83, 0xe9, 0xd4, 0x43, 0xd3, 0xa9, 0x54, 0xc3, 0xc5, 0x1d, 0x5a, 0x9c, 0x5a,
	0xe4, 0x02, 0xd1, 0x2b, 0xc4, 0xc7, 0xc5, 0x94, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x12,
	0xc4, 0x94, 0x99, 0x22, 0xe4, 0xce, 0xc5, 0x0e, 0x35, 0x56, 0x82, 0x49, 0x81, 0x59, 0x83, 0xdb,
	0x48, 0x57, 0x8f, 0x28, 0x73, 0xf5, 0x20, 0x06, 0x06, 0xc1, 0x74, 0x0b, 0x89, 0x70, 0xb1, 0xe6,
	0x97, 0xe7, 0xa5, 0x16, 0x49, 0x30, 0x2b, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x4a, 0x09, 0x5c,
	0x6c, 0x10, 0x85, 0x42, 0xd2, 0x5c, 0x9c, 0x10, 0xa5, 0xf1, 0x50, 0xfb, 0x39, 0x83, 0x38, 0x20,
	0x02, 0x9e, 0x29, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x60, 0x71, 0x30,
	0x5b, 0x48, 0x95, 0x8b, 0x0f, 0xaa, 0x21, 0x31, 0x25, 0xa5, 0x28, 0xb5, 0xb8, 0x18, 0x6a, 0x32,
	0x2f, 0x44, 0xd4, 0x11, 0x22, 0xe8, 0x14, 0x7c, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c,
	0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72,
	0x0c, 0x51, 0x96, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xc8, 0x7e,
	0xd2, 0x4f, 0x36, 0x49, 0xd5, 0x85, 0x04, 0x73, 0x05, 0x46, 0x40, 0x97, 0x54, 0x16, 0xa4, 0x16,
	0x27, 0xb1, 0x81, 0x83, 0xd8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x81, 0xc8, 0x8f, 0x50, 0x91,
	0x01, 0x00, 0x00,
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
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintUserDevices(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x1a
	}
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
	if m.Id != 0 {
		i = encodeVarintUserDevices(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
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
	if len(m.DeviceAddress) > 0 {
		i -= len(m.DeviceAddress)
		copy(dAtA[i:], m.DeviceAddress)
		i = encodeVarintUserDevices(dAtA, i, uint64(len(m.DeviceAddress)))
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
	if len(m.DeviceId) > 0 {
		i -= len(m.DeviceId)
		copy(dAtA[i:], m.DeviceId)
		i = encodeVarintUserDevices(dAtA, i, uint64(len(m.DeviceId)))
		i--
		dAtA[i] = 0xa
	}
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
	if m.Id != 0 {
		n += 1 + sovUserDevices(uint64(m.Id))
	}
	if len(m.Devices) > 0 {
		for _, e := range m.Devices {
			l = e.Size()
			n += 1 + l + sovUserDevices(uint64(l))
		}
	}
	l = len(m.Owner)
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
	l = len(m.DeviceId)
	if l > 0 {
		n += 1 + l + sovUserDevices(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUserDevices(uint64(l))
	}
	l = len(m.DeviceAddress)
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserDevices
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
			m.Devices = append(m.Devices, &Device{})
			if err := m.Devices[len(m.Devices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
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
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceId", wireType)
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
			m.DeviceId = string(dAtA[iNdEx:postIndex])
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
