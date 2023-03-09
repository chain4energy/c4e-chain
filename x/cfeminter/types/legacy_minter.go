package types

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	"io"
	"sort"
	"time"
)

func (params MinterConfig) Validate() error {
	if len(params.Minters) < 1 {
		return fmt.Errorf("no minters defined")
	}

	for i, minter := range params.Minters {
		if minter == nil {
			return fmt.Errorf("minter on position %d cannot be nil", i+1)
		}
	}

	sort.Sort(ByLegacySequenceId(params.Minters))
	lastPos := len(params.Minters) - 1
	id := uint32(0)
	for i, minter := range params.Minters {
		minterId, err := params.validateMinterOrderingId(minter, id)
		if err != nil {
			return err
		}
		id = minterId

		err = params.validateEndTimeExistance(minter, i, lastPos)
		if err != nil {
			return err
		}

		err = params.validateMintersEndTimeValue(minter, i, lastPos)
		if err != nil {
			return err
		}

		if err = minter.validate(); err != nil {
			return fmt.Errorf("minter with id %d validation error: %w", minter.SequenceId, err)
		}
	}
	return nil
}

type ByLegacySequenceId []*LegacyMinter

func (a ByLegacySequenceId) Len() int           { return len(a) }
func (a ByLegacySequenceId) Less(i, j int) bool { return a[i].SequenceId < a[j].SequenceId }
func (a ByLegacySequenceId) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (params MinterConfig) validateMinterOrderingId(minter *LegacyMinter, id uint32) (uint32, error) {
	if id == 0 {
		if minter.SequenceId <= id {
			return 0, fmt.Errorf("first minter sequence id must be bigger than 0, but is %d", minter.SequenceId)
		}
		id = minter.SequenceId
	} else {
		if minter.SequenceId != id+1 {
			return 0, fmt.Errorf("missing minter with sequence id %d", id+1)
		}
		id = minter.SequenceId
	}
	return id, nil
}

func (params MinterConfig) validateEndTimeExistance(minter *LegacyMinter, sequenceId int, lastPos int) error {
	if sequenceId == lastPos && minter.EndTime != nil {
		return fmt.Errorf("last minter cannot have EndTime set, but is set to %s", minter.EndTime)
	}
	if sequenceId < lastPos && minter.EndTime == nil {
		return fmt.Errorf("only last minter can have EndTime empty")
	}
	return nil
}

func (params MinterConfig) validateMintersEndTimeValue(minter *LegacyMinter, sequenceId int, lastPos int) error {
	if lastPos > 0 {
		if sequenceId == 0 {
			if minter.EndTime.Before(params.StartTime) || minter.EndTime.Equal(params.StartTime) {
				return fmt.Errorf("first minter end must be bigger than minter start")
			}
		} else if sequenceId < lastPos {
			prev := sequenceId - 1
			if minter.EndTime.Before(*params.Minters[prev].EndTime) || minter.EndTime.Equal(*params.Minters[prev].EndTime) {
				return fmt.Errorf("minter with sequence id %d mast have EndTime bigger than minter with sequence id %d", minter.SequenceId, params.Minters[prev].SequenceId)
			}
		}
	}
	return nil
}

func (params MinterConfig) ContainsMinter(sequenceId uint32) bool {
	for _, minter := range params.Minters {
		if sequenceId == minter.SequenceId {
			return true
		}
	}
	return false
}

func (m LegacyMinter) validate() error {
	switch m.Type {
	case NoMintingType:
		if m.LinearMinting != nil || m.ExponentialStepMinting != nil {
			return fmt.Errorf("for NO_MINTING type (0) LinearMinting and ExponentialStepMinting cannot be set")
		}
	case LinearMintingType:
		if m.ExponentialStepMinting != nil {
			return fmt.Errorf("for LinearMintingType type (1) ExponentialStepMinting cannot be set")
		}
		if m.EndTime == nil {
			return fmt.Errorf("for LinearMintingType type (1) EndTime must be set")
		}
		if err := m.LinearMinting.validate(); err != nil {
			return fmt.Errorf("LinearMintingType error: %w", err)
		}
	case ExponentialStepMintingType:
		if m.LinearMinting != nil {
			return fmt.Errorf("for ExponentialStepMintingType type (2) LinearMinting cannot be set")
		}
		if err := m.ExponentialStepMinting.validate(); err != nil {
			return fmt.Errorf("ExponentialStepMintingType error: %w", err)
		}
	default:
		return fmt.Errorf("unknow minting configuration type: %s", m.Type)
	}
	return nil
}

func (m *LinearMinting) validate() error {
	if m == nil {
		return fmt.Errorf("for LinearMintingType type (1) LinearMinting must be set")
	}
	if m.Amount.IsNil() {
		return fmt.Errorf("amount cannot be nil")
	}
	if m.Amount.IsNegative() {
		return fmt.Errorf("amount cannot be less than 0")
	}
	return nil
}

func (m *ExponentialStepMinting) validate() error {
	if m == nil {
		return fmt.Errorf("for ExponentialStepMintingType type (2) ExponentialStepMinting must be set")
	}
	if m.Amount.IsNil() {
		return fmt.Errorf("amount cannot be nil")
	}
	if m.Amount.IsNegative() {
		return fmt.Errorf("amount cannot be less than 0")
	}
	if m.AmountMultiplier.IsNil() {
		return fmt.Errorf("amountMultiplier cannot be nil")
	}
	if m.AmountMultiplier.IsNegative() {
		return fmt.Errorf("amountMultiplier cannot be less than 0")
	}
	if m.StepDuration <= 0 {
		return fmt.Errorf("stepDuration must be bigger than 0")
	}
	return nil
}

func (m LegacyMinterState) Validate() error {
	if m.AmountMinted.IsNil() {
		return fmt.Errorf("minter state validation error: amountMinted cannot be nil")
	}
	if m.AmountMinted.IsNegative() {
		return fmt.Errorf("minter state validation error: amountMinted cannot be less than 0")
	}
	if m.RemainderFromPreviousPeriod.IsNil() {
		return fmt.Errorf("minter state validation error: remainderFromPreviousPeriod cannot be nil")
	}
	if m.RemainderFromPreviousPeriod.IsNegative() {
		return fmt.Errorf("minter state validation error: remainderFromPreviousPeriod cannot be less than 0")
	}
	if m.RemainderToMint.IsNil() {
		return fmt.Errorf("minter state validation error: remainderToMint cannot be nil")
	}
	if m.RemainderToMint.IsNegative() {
		return fmt.Errorf("minter state validation error: remainderToMint cannot be less than 0")
	}
	return nil
}

func (m *LegacyMinterState) Reset()         { *m = LegacyMinterState{} }
func (m *LegacyMinterState) String() string { return proto.CompactTextString(m) }
func (*LegacyMinterState) ProtoMessage()    {}
func (*LegacyMinterState) Descriptor() ([]byte, []int) {
	return fileDescriptor_1112145d4942e936, []int{4}
}
func (m *LegacyMinterState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LegacyMinterState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MinterState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LegacyMinterState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MinterState.Merge(m, src)
}
func (m *LegacyMinterState) XXX_Size() int {
	return m.Size()
}
func (m *LegacyMinterState) XXX_DiscardUnknown() {
	xxx_messageInfo_LegacyMinterState.DiscardUnknown(m)
}

var xxx_messageInfo_LegacyMinterState proto.InternalMessageInfo

func (m *LegacyMinterState) GetSequenceId() uint32 {
	if m != nil {
		return m.SequenceId
	}
	return 0
}

func (m *LegacyMinterState) GetLastMintBlockTime() time.Time {
	if m != nil {
		return m.LastMintBlockTime
	}
	return time.Time{}
}

func (m *LegacyMinterState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LegacyMinterState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LegacyMinterState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.RemainderFromPreviousPeriod.Size()
		i -= size
		if _, err := m.RemainderFromPreviousPeriod.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMinter(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	n6, err6 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LastMintBlockTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LastMintBlockTime):])
	if err6 != nil {
		return 0, err6
	}
	i -= n6
	i = encodeVarintMinter(dAtA, i, uint64(n6))
	i--
	dAtA[i] = 0x22
	{
		size := m.RemainderToMint.Size()
		i -= size
		if _, err := m.RemainderToMint.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMinter(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.AmountMinted.Size()
		i -= size
		if _, err := m.AmountMinted.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMinter(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.SequenceId != 0 {
		i = encodeVarintMinter(dAtA, i, uint64(m.SequenceId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LegacyMinterState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SequenceId != 0 {
		n += 1 + sovMinter(uint64(m.SequenceId))
	}
	l = m.AmountMinted.Size()
	n += 1 + l + sovMinter(uint64(l))
	l = m.RemainderToMint.Size()
	n += 1 + l + sovMinter(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastMintBlockTime)
	n += 1 + l + sovMinter(uint64(l))
	l = m.RemainderFromPreviousPeriod.Size()
	n += 1 + l + sovMinter(uint64(l))
	return n
}

func (m *LegacyMinterState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMinter
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
			return fmt.Errorf("proto: MinterState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MinterState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequenceId", wireType)
			}
			m.SequenceId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SequenceId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountMinted", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
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
				return ErrInvalidLengthMinter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountMinted.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemainderToMint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
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
				return ErrInvalidLengthMinter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RemainderToMint.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastMintBlockTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
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
				return ErrInvalidLengthMinter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LastMintBlockTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemainderFromPreviousPeriod", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
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
				return ErrInvalidLengthMinter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RemainderFromPreviousPeriod.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMinter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMinter
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

var fileDescriptor_1112145d4942e936 = []byte{
	// 638 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x4f, 0x6f, 0xd3, 0x3e,
	0x18, 0xc7, 0x9b, 0xb6, 0xbf, 0xfd, 0x71, 0xd7, 0x1f, 0xcc, 0x9a, 0xa6, 0x32, 0xa4, 0xb4, 0xea,
	0x01, 0x7a, 0x59, 0x22, 0x6d, 0x93, 0x38, 0x70, 0xa2, 0x1b, 0x13, 0x93, 0xa8, 0x34, 0x65, 0x9b,
	0x90, 0xc6, 0x21, 0x4a, 0x93, 0xa7, 0x99, 0xb5, 0xc4, 0xce, 0x6c, 0x07, 0x6d, 0xaf, 0x80, 0xeb,
	0x4e, 0x88, 0xf7, 0xc0, 0xdb, 0xe0, 0xb0, 0xe3, 0x8e, 0x88, 0xc3, 0x40, 0xdb, 0x1b, 0x41, 0xb6,
	0x9b, 0xb6, 0xc0, 0xa4, 0x52, 0x38, 0x25, 0x76, 0x9e, 0xef, 0xe7, 0xfb, 0xb5, 0xfd, 0x38, 0x68,
	0x35, 0x1c, 0x40, 0x4a, 0xa8, 0x04, 0xee, 0x9a, 0x87, 0x93, 0x71, 0x26, 0x19, 0x6e, 0x86, 0x27,
	0x01, 0xa1, 0x5b, 0x40, 0x81, 0xc7, 0x17, 0x4e, 0xb8, 0x05, 0x7a, 0xec, 0x8c, 0xaa, 0xd7, 0x56,
	0x62, 0x16, 0x33, 0x5d, 0xeb, 0xaa, 0x37, 0x23, 0x5b, 0x6b, 0xc6, 0x8c, 0xc5, 0x09, 0xb8, 0x7a,
	0xd4, 0xcf, 0x07, 0xae, 0x24, 0x29, 0x08, 0x19, 0xa4, 0xd9, 0xb0, 0xc0, 0xfe, 0xb5, 0x20, 0xca,
	0x79, 0x20, 0x09, 0xa3, 0xe6, 0x7b, 0xfb, 0x83, 0x85, 0x96, 0x7a, 0xda, 0x61, 0x9b, 0xd1, 0x01,
	0x89, 0xf1, 0x36, 0x42, 0x42, 0x06, 0x5c, 0xfa, 0x8a, 0xd4, 0x28, 0xb7, 0xac, 0x4e, 0x6d, 0x63,
	0xcd, 0x31, 0x14, 0xa7, 0xa0, 0x38, 0x87, 0x85, 0x4d, 0x77, 0xe1, 0xea, 0xa6, 0x59, 0xba, 0xfc,
	0xd6, 0xb4, 0xbc, 0x45, 0xad, 0x53, 0x5f, 0xf0, 0x0b, 0x34, 0x6f, 0x62, 0x8b, 0x46, 0xa5, 0x55,
	0xe9, 0xd4, 0x36, 0x9e, 0x3a, 0x53, 0xd6, 0xe7, 0x98, 0x10, 0x5e, 0xa1, 0x6b, 0x7f, 0x2e, 0xa3,
	0x39, 0x33, 0x87, 0x9b, 0xa8, 0x26, 0xe0, 0x2c, 0x07, 0x1a, 0x82, 0x4f, 0xa2, 0x86, 0xd5, 0xb2,
	0x3a, 0x75, 0x0f, 0x15, 0x53, 0x7b, 0x11, 0x7e, 0x8e, 0x16, 0x80, 0x46, 0x7f, 0x9a, 0xb8, 0xaa,
	0xd3, 0xce, 0x03, 0x8d, 0x74, 0x56, 0x8c, 0xaa, 0xf2, 0x22, 0x83, 0x46, 0xa5, 0x65, 0x75, 0x16,
	0x3d, 0xfd, 0x8e, 0x8f, 0xd0, 0xff, 0x09, 0xa1, 0x10, 0x70, 0x5f, 0xc5, 0x21, 0x34, 0x6e, 0x54,
	0x35, 0xd6, 0x99, 0xba, 0x8c, 0xd7, 0x5a, 0xd6, 0x33, 0x2a, 0xaf, 0x9e, 0x4c, 0x0e, 0xf1, 0x19,
	0x6a, 0xc0, 0x79, 0xc6, 0x28, 0x50, 0x49, 0x82, 0xc4, 0x17, 0x12, 0xb2, 0x91, 0xc1, 0x7f, 0xda,
	0xe0, 0xd9, 0x54, 0x83, 0x97, 0x63, 0xc0, 0x81, 0x84, 0xac, 0x70, 0x5a, 0x85, 0x7b, 0xe7, 0xdb,
	0x6f, 0x50, 0xfd, 0xa7, 0x48, 0x78, 0x17, 0xcd, 0x05, 0x29, 0xcb, 0xa9, 0xd4, 0xfb, 0xb8, 0xd8,
	0x75, 0xd4, 0xf9, 0x7d, 0xbd, 0x69, 0x3e, 0x89, 0x89, 0x3c, 0xc9, 0xfb, 0x4e, 0xc8, 0x52, 0x37,
	0x64, 0x22, 0x65, 0x62, 0xf8, 0x58, 0x17, 0xd1, 0xa9, 0xab, 0x76, 0x45, 0x38, 0x7b, 0x54, 0x7a,
	0x43, 0x75, 0xfb, 0x7d, 0x19, 0xad, 0xde, 0x9f, 0x65, 0xc2, 0xa2, 0xfc, 0x2f, 0x16, 0xf8, 0x15,
	0xaa, 0xeb, 0x2d, 0x2a, 0x5a, 0x56, 0x27, 0xae, 0x6d, 0x3c, 0xfa, 0xed, 0x6c, 0x77, 0x86, 0x05,
	0xa6, 0x19, 0x3f, 0xaa, 0xe3, 0x5d, 0x52, 0xca, 0x62, 0x1e, 0xbf, 0x45, 0xcb, 0x86, 0xe9, 0xa7,
	0x79, 0x22, 0x49, 0x96, 0x10, 0xe0, 0xfa, 0x48, 0x67, 0x0b, 0xb7, 0x03, 0xa1, 0xf7, 0xd0, 0x80,
	0x7a, 0x23, 0x4e, 0xfb, 0x53, 0x05, 0xd5, 0x4c, 0xa7, 0x1e, 0xc8, 0x40, 0xc2, 0xf4, 0x76, 0x3d,
	0x40, 0xf5, 0x22, 0x8d, 0x92, 0x45, 0x7f, 0xb9, 0x4d, 0x4b, 0xc3, 0x24, 0x9a, 0x81, 0x8f, 0xd1,
	0x32, 0x87, 0x34, 0x20, 0x34, 0x02, 0xee, 0x4b, 0xa6, 0xd1, 0xa6, 0xa7, 0x67, 0x5e, 0xe2, 0x83,
	0x11, 0xe8, 0x90, 0x29, 0x3a, 0x3e, 0x42, 0x2b, 0x49, 0x20, 0x4c, 0x5c, 0xbf, 0x9f, 0xb0, 0xf0,
	0xd4, 0xdc, 0xb5, 0xea, 0x0c, 0x7f, 0x87, 0x65, 0x45, 0x50, 0xb4, 0xae, 0xd2, 0xeb, 0x9b, 0x27,
	0x90, 0x3d, 0x8e, 0x3c, 0xe0, 0x2c, 0xf5, 0x33, 0x0e, 0xef, 0x08, 0xcb, 0x85, 0x9f, 0x01, 0x27,
	0x2c, 0xd2, 0x97, 0x62, 0xf6, 0xfc, 0x8f, 0x47, 0xd4, 0x5d, 0xce, 0xd2, 0xfd, 0x21, 0x73, 0x5f,
	0x23, 0xbb, 0xbd, 0xab, 0x5b, 0xdb, 0xba, 0xbe, 0xb5, 0xad, 0xef, 0xb7, 0xb6, 0x75, 0x79, 0x67,
	0x97, 0xae, 0xef, 0xec, 0xd2, 0x97, 0x3b, 0xbb, 0x74, 0xbc, 0x39, 0x89, 0x9f, 0xb8, 0x85, 0x6e,
	0xb8, 0x05, 0xeb, 0x7a, 0xc2, 0x3d, 0x77, 0xc7, 0xbf, 0x6f, 0xed, 0xd7, 0x9f, 0xd3, 0x8b, 0xde,
	0xfc, 0x11, 0x00, 0x00, 0xff, 0xff, 0x64, 0x1c, 0xad, 0xfd, 0xd8, 0x05, 0x00, 0x00,
}
