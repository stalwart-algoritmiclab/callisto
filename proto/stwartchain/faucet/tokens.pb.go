// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stwartchain/faucet/tokens.proto

package faucet

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/cosmos/gogoproto/proto"
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

type Tokens struct {
	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Denom   string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	Amount  string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Creator string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Tokens) Reset()         { *m = Tokens{} }
func (m *Tokens) String() string { return proto.CompactTextString(m) }
func (*Tokens) ProtoMessage()    {}
func (*Tokens) Descriptor() ([]byte, []int) {
	return fileDescriptor_cee1ad4d97921188, []int{0}
}
func (m *Tokens) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Tokens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Tokens.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Tokens) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tokens.Merge(m, src)
}
func (m *Tokens) XXX_Size() int {
	return m.Size()
}
func (m *Tokens) XXX_DiscardUnknown() {
	xxx_messageInfo_Tokens.DiscardUnknown(m)
}

var xxx_messageInfo_Tokens proto.InternalMessageInfo

func (m *Tokens) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Tokens) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *Tokens) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *Tokens) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Tokens)(nil), "stwartchain.faucet.Tokens")
}

func init() { proto.RegisterFile("stwartchain/faucet/tokens.proto", fileDescriptor_cee1ad4d97921188) }

var fileDescriptor_cee1ad4d97921188 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x2e, 0x29, 0x4f,
	0x2c, 0x2a, 0x49, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x4f, 0x4b, 0x2c, 0x4d, 0x4e, 0x2d, 0xd1, 0x2f,
	0xc9, 0xcf, 0x4e, 0xcd, 0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x42, 0x52, 0xa0,
	0x07, 0x51, 0xa0, 0x94, 0xc0, 0xc5, 0x16, 0x02, 0x56, 0x23, 0xc4, 0xc7, 0xc5, 0x94, 0x99, 0x22,
	0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x12, 0xc4, 0x94, 0x99, 0x22, 0x24, 0xc2, 0xc5, 0x9a, 0x92, 0x9a,
	0x97, 0x9f, 0x2b, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe1, 0x08, 0x89, 0x71, 0xb1, 0x25,
	0xe6, 0xe6, 0x97, 0xe6, 0x95, 0x48, 0x30, 0x83, 0x85, 0xa1, 0x3c, 0x21, 0x09, 0x2e, 0xf6, 0xe4,
	0xa2, 0xd4, 0xc4, 0x92, 0xfc, 0x22, 0x09, 0x16, 0xb0, 0x04, 0x8c, 0xeb, 0x14, 0x73, 0xe2, 0x91,
	0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1,
	0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x4e, 0xe9, 0x99, 0x25, 0x39, 0x89, 0x49, 0x7a,
	0xc5, 0x25, 0x89, 0x39, 0x20, 0x87, 0xe9, 0x95, 0xa4, 0x26, 0x67, 0xe8, 0x67, 0x66, 0x65, 0xe6,
	0xeb, 0xe7, 0x82, 0xfc, 0x90, 0x94, 0x98, 0x9c, 0x9d, 0x9a, 0x97, 0xa2, 0x0f, 0x71, 0xb6, 0x2e,
	0xc4, 0x63, 0x15, 0x70, 0xaf, 0x55, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0xbd, 0x66, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x71, 0xc9, 0x65, 0x8c, 0xfd, 0x00, 0x00, 0x00,
}

func (m *Tokens) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Tokens) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Tokens) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTokens(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintTokens(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTokens(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintTokens(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTokens(dAtA []byte, offset int, v uint64) int {
	offset -= sovTokens(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Tokens) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTokens(uint64(m.Id))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTokens(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovTokens(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTokens(uint64(l))
	}
	return n
}

func sovTokens(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTokens(x uint64) (n int) {
	return sovTokens(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Tokens) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTokens
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
			return fmt.Errorf("proto: Tokens: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Tokens: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokens
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
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokens
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
				return ErrInvalidLengthTokens
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTokens
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokens
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
				return ErrInvalidLengthTokens
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTokens
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokens
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
				return ErrInvalidLengthTokens
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTokens
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTokens(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTokens
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
func skipTokens(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTokens
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
					return 0, ErrIntOverflowTokens
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
					return 0, ErrIntOverflowTokens
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
				return 0, ErrInvalidLengthTokens
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTokens
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTokens
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTokens        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTokens          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTokens = fmt.Errorf("proto: unexpected end of group")
)