// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: c4echain/cfeclaim/claim_record.proto

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

type UserEntry struct {
	Address      string         `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	ClaimAddress string         `protobuf:"bytes,2,opt,name=claim_address,json=claimAddress,proto3" json:"claim_address,omitempty"`
	ClaimRecords []*ClaimRecord `protobuf:"bytes,3,rep,name=claim_records,json=claimRecords,proto3" json:"claim_records,omitempty"`
}

func (m *UserEntry) Reset()         { *m = UserEntry{} }
func (m *UserEntry) String() string { return proto.CompactTextString(m) }
func (*UserEntry) ProtoMessage()    {}
func (*UserEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f30000e08f2d89a, []int{0}
}
func (m *UserEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserEntry.Merge(m, src)
}
func (m *UserEntry) XXX_Size() int {
	return m.Size()
}
func (m *UserEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_UserEntry.DiscardUnknown(m)
}

var xxx_messageInfo_UserEntry proto.InternalMessageInfo

func (m *UserEntry) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *UserEntry) GetClaimAddress() string {
	if m != nil {
		return m.ClaimAddress
	}
	return ""
}

func (m *UserEntry) GetClaimRecords() []*ClaimRecord {
	if m != nil {
		return m.ClaimRecords
	}
	return nil
}

type ClaimRecord struct {
	CampaignId        uint64                                   `protobuf:"varint,1,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
	Address           string                                   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Amount            github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
	CompletedMissions []uint64                                 `protobuf:"varint,4,rep,packed,name=completedMissions,proto3" json:"completedMissions,omitempty"`
	ClaimedMissions   []uint64                                 `protobuf:"varint,5,rep,packed,name=claimedMissions,proto3" json:"claimedMissions,omitempty"`
}

func (m *ClaimRecord) Reset()         { *m = ClaimRecord{} }
func (m *ClaimRecord) String() string { return proto.CompactTextString(m) }
func (*ClaimRecord) ProtoMessage()    {}
func (*ClaimRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f30000e08f2d89a, []int{1}
}
func (m *ClaimRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimRecord.Merge(m, src)
}
func (m *ClaimRecord) XXX_Size() int {
	return m.Size()
}
func (m *ClaimRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimRecord proto.InternalMessageInfo

func (m *ClaimRecord) GetCampaignId() uint64 {
	if m != nil {
		return m.CampaignId
	}
	return 0
}

func (m *ClaimRecord) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ClaimRecord) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *ClaimRecord) GetCompletedMissions() []uint64 {
	if m != nil {
		return m.CompletedMissions
	}
	return nil
}

func (m *ClaimRecord) GetClaimedMissions() []uint64 {
	if m != nil {
		return m.ClaimedMissions
	}
	return nil
}

func init() {
	proto.RegisterType((*UserEntry)(nil), "chain4energy.c4echain.cfeclaim.UserEntry")
	proto.RegisterType((*ClaimRecord)(nil), "chain4energy.c4echain.cfeclaim.ClaimRecord")
}

func init() {
	proto.RegisterFile("c4echain/cfeclaim/claim_record.proto", fileDescriptor_7f30000e08f2d89a)
}

var fileDescriptor_7f30000e08f2d89a = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x41, 0x4e, 0xe3, 0x30,
	0x14, 0x4d, 0xda, 0x4e, 0x47, 0x75, 0x67, 0x34, 0x9a, 0x68, 0x16, 0x99, 0x2e, 0xdc, 0xaa, 0xb0,
	0x88, 0x04, 0xb5, 0x69, 0xe9, 0x05, 0x68, 0xc5, 0x02, 0x09, 0x24, 0x14, 0x89, 0x0d, 0x9b, 0xca,
	0x71, 0x4c, 0x6a, 0xd1, 0xd8, 0x51, 0x9c, 0x22, 0x7a, 0x00, 0xf6, 0x1c, 0x80, 0x13, 0x70, 0x92,
	0x2e, 0xbb, 0x64, 0x05, 0xa8, 0xbd, 0x08, 0xaa, 0x9d, 0x88, 0x00, 0x12, 0x1b, 0x27, 0xfe, 0xef,
	0xbd, 0xaf, 0xf7, 0x9f, 0x3f, 0xd8, 0xa5, 0x43, 0x46, 0xa7, 0x84, 0x0b, 0x4c, 0xaf, 0x18, 0x9d,
	0x11, 0x1e, 0x63, 0x7d, 0x4e, 0x52, 0x46, 0x65, 0x1a, 0xa2, 0x24, 0x95, 0x99, 0x74, 0xa0, 0xa6,
	0x0c, 0x99, 0x60, 0x69, 0xb4, 0x40, 0x85, 0x04, 0x15, 0x92, 0xd6, 0xbf, 0x48, 0x46, 0x52, 0x53,
	0xf1, 0xf6, 0xcf, 0xa8, 0x5a, 0x90, 0x4a, 0x15, 0x4b, 0x85, 0x03, 0xa2, 0x18, 0xbe, 0xe9, 0x07,
	0x2c, 0x23, 0x7d, 0x4c, 0x25, 0x17, 0x06, 0xef, 0x3e, 0xd8, 0xa0, 0x71, 0xa1, 0x58, 0x7a, 0x2c,
	0xb2, 0x74, 0xe1, 0xb8, 0xe0, 0x27, 0x09, 0xc3, 0x94, 0x29, 0xe5, 0xda, 0x1d, 0xdb, 0x6b, 0xf8,
	0xc5, 0xd5, 0xd9, 0x01, 0xbf, 0x8d, 0xa7, 0x02, 0xaf, 0x68, 0xfc, 0x97, 0x2e, 0x1e, 0xe5, 0xa4,
	0xf3, 0x82, 0x64, 0x8c, 0x2b, 0xb7, 0xda, 0xa9, 0x7a, 0xcd, 0xc1, 0x1e, 0xfa, 0xde, 0x3a, 0x1a,
	0x6f, 0x4f, 0x5f, 0x6b, 0xf2, 0x8e, 0xe6, 0xa2, 0xba, 0x77, 0x15, 0xd0, 0x2c, 0xa1, 0x4e, 0x1b,
	0x34, 0x29, 0x89, 0x13, 0xc2, 0x23, 0x31, 0xe1, 0xa1, 0x36, 0x59, 0xf3, 0x41, 0x51, 0x3a, 0x09,
	0xcb, 0x13, 0x54, 0x3e, 0x4e, 0x40, 0x41, 0x9d, 0xc4, 0x72, 0x2e, 0xb2, 0xdc, 0xd5, 0x7f, 0x64,
	0xa2, 0x41, 0xdb, 0x68, 0x50, 0x1e, 0x0d, 0x1a, 0x4b, 0x2e, 0x46, 0x07, 0xcb, 0xe7, 0xb6, 0xf5,
	0xf8, 0xd2, 0xf6, 0x22, 0x9e, 0x4d, 0xe7, 0x01, 0xa2, 0x32, 0xc6, 0x79, 0x8e, 0xe6, 0xd3, 0x53,
	0xe1, 0x35, 0xce, 0x16, 0x09, 0x53, 0x5a, 0xa0, 0xfc, 0xbc, 0xb5, 0xb3, 0x0f, 0xfe, 0x52, 0x19,
	0x27, 0x33, 0x96, 0xb1, 0xf0, 0x8c, 0x2b, 0xc5, 0xa5, 0x50, 0x6e, 0xad, 0x53, 0xf5, 0x6a, 0xfe,
	0x57, 0xc0, 0xf1, 0xc0, 0x1f, 0x3d, 0x6d, 0x89, 0xfb, 0x43, 0x73, 0x3f, 0x97, 0x47, 0xa7, 0xcb,
	0x35, 0xb4, 0x57, 0x6b, 0x68, 0xbf, 0xae, 0xa1, 0x7d, 0xbf, 0x81, 0xd6, 0x6a, 0x03, 0xad, 0xa7,
	0x0d, 0xb4, 0x2e, 0x07, 0x65, 0x8f, 0xa5, 0x98, 0x31, 0x1d, 0xb2, 0x9e, 0xd9, 0xaa, 0xdb, 0xf7,
	0xbd, 0xd2, 0x9e, 0x83, 0xba, 0x7e, 0xfb, 0xc3, 0xb7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x80,
	0x45, 0x30, 0x79, 0x02, 0x00, 0x00,
}

func (m *UserEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClaimRecords) > 0 {
		for iNdEx := len(m.ClaimRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClaimRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintClaimRecord(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ClaimAddress) > 0 {
		i -= len(m.ClaimAddress)
		copy(dAtA[i:], m.ClaimAddress)
		i = encodeVarintClaimRecord(dAtA, i, uint64(len(m.ClaimAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintClaimRecord(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClaimRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClaimedMissions) > 0 {
		dAtA2 := make([]byte, len(m.ClaimedMissions)*10)
		var j1 int
		for _, num := range m.ClaimedMissions {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintClaimRecord(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.CompletedMissions) > 0 {
		dAtA4 := make([]byte, len(m.CompletedMissions)*10)
		var j3 int
		for _, num := range m.CompletedMissions {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintClaimRecord(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintClaimRecord(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintClaimRecord(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.CampaignId != 0 {
		i = encodeVarintClaimRecord(dAtA, i, uint64(m.CampaignId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintClaimRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovClaimRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UserEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovClaimRecord(uint64(l))
	}
	l = len(m.ClaimAddress)
	if l > 0 {
		n += 1 + l + sovClaimRecord(uint64(l))
	}
	if len(m.ClaimRecords) > 0 {
		for _, e := range m.ClaimRecords {
			l = e.Size()
			n += 1 + l + sovClaimRecord(uint64(l))
		}
	}
	return n
}

func (m *ClaimRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CampaignId != 0 {
		n += 1 + sovClaimRecord(uint64(m.CampaignId))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovClaimRecord(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovClaimRecord(uint64(l))
		}
	}
	if len(m.CompletedMissions) > 0 {
		l = 0
		for _, e := range m.CompletedMissions {
			l += sovClaimRecord(uint64(e))
		}
		n += 1 + sovClaimRecord(uint64(l)) + l
	}
	if len(m.ClaimedMissions) > 0 {
		l = 0
		for _, e := range m.ClaimedMissions {
			l += sovClaimRecord(uint64(e))
		}
		n += 1 + sovClaimRecord(uint64(l)) + l
	}
	return n
}

func sovClaimRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClaimRecord(x uint64) (n int) {
	return sovClaimRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaimRecord
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
			return fmt.Errorf("proto: UserEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimRecord
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
				return ErrInvalidLengthClaimRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaimRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimRecord
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
				return ErrInvalidLengthClaimRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaimRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimRecord
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
				return ErrInvalidLengthClaimRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaimRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimRecords = append(m.ClaimRecords, &ClaimRecord{})
			if err := m.ClaimRecords[len(m.ClaimRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClaimRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaimRecord
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
func (m *ClaimRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaimRecord
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
			return fmt.Errorf("proto: ClaimRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CampaignId", wireType)
			}
			m.CampaignId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CampaignId |= uint64(b&0x7F) << shift
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
					return ErrIntOverflowClaimRecord
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
				return ErrInvalidLengthClaimRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaimRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimRecord
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
				return ErrInvalidLengthClaimRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaimRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaimRecord
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.CompletedMissions = append(m.CompletedMissions, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaimRecord
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClaimRecord
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthClaimRecord
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.CompletedMissions) == 0 {
					m.CompletedMissions = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaimRecord
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.CompletedMissions = append(m.CompletedMissions, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field CompletedMissions", wireType)
			}
		case 5:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaimRecord
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ClaimedMissions = append(m.ClaimedMissions, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaimRecord
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClaimRecord
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthClaimRecord
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.ClaimedMissions) == 0 {
					m.ClaimedMissions = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaimRecord
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ClaimedMissions = append(m.ClaimedMissions, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimedMissions", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClaimRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaimRecord
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
func skipClaimRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClaimRecord
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
					return 0, ErrIntOverflowClaimRecord
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
					return 0, ErrIntOverflowClaimRecord
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
				return 0, ErrInvalidLengthClaimRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClaimRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClaimRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClaimRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClaimRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClaimRecord = fmt.Errorf("proto: unexpected end of group")
)