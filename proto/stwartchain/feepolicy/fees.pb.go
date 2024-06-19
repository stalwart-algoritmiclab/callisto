// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stwartchain/feepolicy/fees.proto

package feepolicy

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type Fees struct {
	AmountFrom  string `protobuf:"bytes,1,opt,name=amountFrom,proto3" json:"amountFrom,omitempty"`
	Fee         string `protobuf:"bytes,2,opt,name=fee,proto3" json:"fee,omitempty"`
	RefReward   string `protobuf:"bytes,3,opt,name=refReward,proto3" json:"refReward,omitempty"`
	StakeReward string `protobuf:"bytes,4,opt,name=stakeReward,proto3" json:"stakeReward,omitempty"`
	MinAmount   uint64 `protobuf:"varint,5,opt,name=minAmount,proto3" json:"minAmount,omitempty"`
	NoRefReward bool   `protobuf:"varint,6,opt,name=noRefReward,proto3" json:"noRefReward,omitempty"`
	Creator     string `protobuf:"bytes,7,opt,name=creator,proto3" json:"creator,omitempty"`
	Id          uint64 `protobuf:"varint,8,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *Fees) Reset()         { *m = Fees{} }
func (m *Fees) String() string { return proto.CompactTextString(m) }
func (*Fees) ProtoMessage()    {}
func (*Fees) Descriptor() ([]byte, []int) {
	return fileDescriptor_6002665507933a6d, []int{0}
}
func (m *Fees) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Fees) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Fees.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Fees) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fees.Merge(m, src)
}
func (m *Fees) XXX_Size() int {
	return m.Size()
}
func (m *Fees) XXX_DiscardUnknown() {
	xxx_messageInfo_Fees.DiscardUnknown(m)
}

var xxx_messageInfo_Fees proto.InternalMessageInfo

func (m *Fees) GetAmountFrom() string {
	if m != nil {
		return m.AmountFrom
	}
	return ""
}

func (m *Fees) GetFee() string {
	if m != nil {
		return m.Fee
	}
	return ""
}

func (m *Fees) GetRefReward() string {
	if m != nil {
		return m.RefReward
	}
	return ""
}

func (m *Fees) GetStakeReward() string {
	if m != nil {
		return m.StakeReward
	}
	return ""
}

func (m *Fees) GetMinAmount() uint64 {
	if m != nil {
		return m.MinAmount
	}
	return 0
}

func (m *Fees) GetNoRefReward() bool {
	if m != nil {
		return m.NoRefReward
	}
	return false
}

func (m *Fees) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Fees) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Fees)(nil), "stwartchain.feepolicy.Fees")
}

func init() { proto.RegisterFile("stwartchain/feepolicy/fees.proto", fileDescriptor_6002665507933a6d) }

var fileDescriptor_6002665507933a6d = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x86, 0xe3, 0x34, 0x5f, 0x7f, 0xfc, 0x49, 0x08, 0x59, 0x42, 0xf2, 0x80, 0xac, 0x88, 0x29,
	0x0b, 0xc9, 0xc0, 0x15, 0x80, 0x44, 0x2f, 0x20, 0x23, 0x0b, 0x72, 0x92, 0x13, 0x6a, 0x9a, 0xc4,
	0x91, 0x7d, 0x50, 0xe9, 0x5d, 0x70, 0x59, 0x8c, 0x1d, 0x19, 0x51, 0x72, 0x23, 0xc8, 0x2e, 0xb4,
	0xd9, 0x8e, 0x9f, 0xf7, 0x9c, 0xd7, 0xd2, 0x43, 0x63, 0x8b, 0x3b, 0x69, 0xb0, 0xdc, 0x48, 0xd5,
	0x65, 0x35, 0x40, 0xaf, 0x1b, 0x55, 0xee, 0xdd, 0x64, 0xd3, 0xde, 0x68, 0xd4, 0xec, 0x6a, 0xb2,
	0x91, 0x9e, 0x36, 0x6e, 0x06, 0x42, 0xa3, 0x35, 0x80, 0x65, 0x82, 0x52, 0xd9, 0xea, 0xb7, 0x0e,
	0xd7, 0x46, 0xb7, 0x9c, 0xc4, 0x24, 0x59, 0xe5, 0x13, 0xc2, 0x2e, 0xe9, 0xac, 0x06, 0xe0, 0xa1,
	0x0f, 0xdc, 0xc8, 0xae, 0xe9, 0xca, 0x40, 0x9d, 0xc3, 0x4e, 0x9a, 0x8a, 0xcf, 0x3c, 0x3f, 0x03,
	0x16, 0xd3, 0xff, 0x16, 0xe5, 0x16, 0x7e, 0xf3, 0xc8, 0xe7, 0x53, 0xe4, 0xee, 0x5b, 0xd5, 0xdd,
	0xfb, 0x2f, 0xf8, 0xbf, 0x98, 0x24, 0x51, 0x7e, 0x06, 0xee, 0xbe, 0xd3, 0xf9, 0xa9, 0x7f, 0x1e,
	0x93, 0x64, 0x99, 0x4f, 0x11, 0xe3, 0x74, 0x51, 0x1a, 0x90, 0xa8, 0x0d, 0x5f, 0xf8, 0xf6, 0xbf,
	0x27, 0xbb, 0xa0, 0xa1, 0xaa, 0xf8, 0xd2, 0x57, 0x86, 0xaa, 0x7a, 0x78, 0xfe, 0x1c, 0x04, 0x39,
	0x0c, 0x82, 0x7c, 0x0f, 0x82, 0x7c, 0x8c, 0x22, 0x38, 0x8c, 0x22, 0xf8, 0x1a, 0x45, 0xf0, 0xf4,
	0xf8, 0xa2, 0xb0, 0x91, 0x45, 0x6a, 0x51, 0x36, 0x4e, 0x4f, 0x8a, 0x50, 0x6e, 0x32, 0xf5, 0xaa,
	0x74, 0xd6, 0x3a, 0x95, 0x85, 0x2c, 0xb7, 0xd0, 0x55, 0xd9, 0x51, 0xde, 0xed, 0xd1, 0xef, 0xfb,
	0xc4, 0x30, 0xee, 0x7b, 0xb0, 0xc5, 0xdc, 0x3b, 0xbe, 0xfb, 0x09, 0x00, 0x00, 0xff, 0xff, 0x4e,
	0x9c, 0x92, 0xe5, 0x87, 0x01, 0x00, 0x00,
}

func (m *Fees) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Fees) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Fees) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintFees(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x40
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintFees(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x3a
	}
	if m.NoRefReward {
		i--
		if m.NoRefReward {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.MinAmount != 0 {
		i = encodeVarintFees(dAtA, i, uint64(m.MinAmount))
		i--
		dAtA[i] = 0x28
	}
	if len(m.StakeReward) > 0 {
		i -= len(m.StakeReward)
		copy(dAtA[i:], m.StakeReward)
		i = encodeVarintFees(dAtA, i, uint64(len(m.StakeReward)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.RefReward) > 0 {
		i -= len(m.RefReward)
		copy(dAtA[i:], m.RefReward)
		i = encodeVarintFees(dAtA, i, uint64(len(m.RefReward)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Fee) > 0 {
		i -= len(m.Fee)
		copy(dAtA[i:], m.Fee)
		i = encodeVarintFees(dAtA, i, uint64(len(m.Fee)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AmountFrom) > 0 {
		i -= len(m.AmountFrom)
		copy(dAtA[i:], m.AmountFrom)
		i = encodeVarintFees(dAtA, i, uint64(len(m.AmountFrom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFees(dAtA []byte, offset int, v uint64) int {
	offset -= sovFees(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Fees) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AmountFrom)
	if l > 0 {
		n += 1 + l + sovFees(uint64(l))
	}
	l = len(m.Fee)
	if l > 0 {
		n += 1 + l + sovFees(uint64(l))
	}
	l = len(m.RefReward)
	if l > 0 {
		n += 1 + l + sovFees(uint64(l))
	}
	l = len(m.StakeReward)
	if l > 0 {
		n += 1 + l + sovFees(uint64(l))
	}
	if m.MinAmount != 0 {
		n += 1 + sovFees(uint64(m.MinAmount))
	}
	if m.NoRefReward {
		n += 2
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovFees(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovFees(uint64(m.Id))
	}
	return n
}

func sovFees(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFees(x uint64) (n int) {
	return sovFees(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Fees) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFees
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
			return fmt.Errorf("proto: Fees: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fees: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountFrom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFees
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
				return ErrInvalidLengthFees
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFees
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AmountFrom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFees
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
				return ErrInvalidLengthFees
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFees
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefReward", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFees
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
				return ErrInvalidLengthFees
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFees
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RefReward = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakeReward", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFees
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
				return ErrInvalidLengthFees
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFees
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakeReward = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinAmount", wireType)
			}
			m.MinAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFees
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoRefReward", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFees
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.NoRefReward = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFees
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
				return ErrInvalidLengthFees
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFees
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFees
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
		default:
			iNdEx = preIndex
			skippy, err := skipFees(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFees
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
func skipFees(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFees
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
					return 0, ErrIntOverflowFees
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
					return 0, ErrIntOverflowFees
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
				return 0, ErrInvalidLengthFees
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFees
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFees
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFees        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFees          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFees = fmt.Errorf("proto: unexpected end of group")
)
