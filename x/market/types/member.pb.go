// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: market/member.proto

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

type Member struct {
	Pair     string                                 `protobuf:"bytes,1,opt,name=pair,proto3" json:"pair,omitempty"`
	DenomA   string                                 `protobuf:"bytes,2,opt,name=denomA,proto3" json:"denomA,omitempty"`
	DenomB   string                                 `protobuf:"bytes,3,opt,name=denomB,proto3" json:"denomB,omitempty"`
	Balance  github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=balance,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"balance"`
	Previous github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=previous,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"previous"`
	Limit    uint64                                 `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	Stop     uint64                                 `protobuf:"varint,7,opt,name=stop,proto3" json:"stop,omitempty"`
}

func (m *Member) Reset()         { *m = Member{} }
func (m *Member) String() string { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()    {}
func (*Member) Descriptor() ([]byte, []int) {
	return fileDescriptor_f075dd79c09a0a2c, []int{0}
}
func (m *Member) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Member) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Member.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Member) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Member.Merge(m, src)
}
func (m *Member) XXX_Size() int {
	return m.Size()
}
func (m *Member) XXX_DiscardUnknown() {
	xxx_messageInfo_Member.DiscardUnknown(m)
}

var xxx_messageInfo_Member proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Member)(nil), "pendulumlabs.market.market.Member")
}

func init() { proto.RegisterFile("market/member.proto", fileDescriptor_f075dd79c09a0a2c) }

var fileDescriptor_f075dd79c09a0a2c = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x51, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x8d, 0xfb, 0xa5, 0xe9, 0x87, 0x47, 0x53, 0x21, 0xab, 0x83, 0x5b, 0x31, 0xa0, 0x2e, 0x75,
	0x06, 0x9e, 0x80, 0x4c, 0x80, 0x84, 0x84, 0x32, 0xb2, 0xe5, 0xc7, 0x0a, 0x51, 0xe3, 0xd8, 0xb2,
	0x1d, 0x04, 0xef, 0xc0, 0xc0, 0x63, 0x65, 0xec, 0x58, 0x31, 0x54, 0x34, 0x79, 0x11, 0x14, 0x27,
	0xa9, 0x3a, 0x33, 0xdd, 0x73, 0x8f, 0xef, 0x39, 0xb6, 0xcf, 0x85, 0x97, 0x3c, 0x52, 0x5b, 0x66,
	0x7c, 0xce, 0x78, 0xcc, 0x14, 0x95, 0x4a, 0x18, 0x81, 0x16, 0x92, 0x95, 0x69, 0x55, 0x54, 0xbc,
	0x88, 0x62, 0x4d, 0xfb, 0x89, 0xa1, 0x2c, 0xe6, 0x99, 0xc8, 0x84, 0x1d, 0xf3, 0x3b, 0xd4, 0x2b,
	0xae, 0x3f, 0x27, 0xd0, 0x7b, 0xb2, 0x16, 0x08, 0x41, 0x57, 0x46, 0xb9, 0xc2, 0x60, 0x05, 0xd6,
	0x17, 0xa1, 0xc5, 0xe8, 0x0a, 0x7a, 0x29, 0x2b, 0x05, 0xbf, 0xc3, 0x13, 0xcb, 0x0e, 0xdd, 0x89,
	0x0f, 0xf0, 0xbf, 0x33, 0x3e, 0x40, 0xf7, 0x70, 0x16, 0x47, 0x45, 0x54, 0x26, 0x0c, 0xbb, 0xdd,
	0x41, 0x40, 0xeb, 0xc3, 0xd2, 0xf9, 0x3e, 0x2c, 0x6f, 0xb2, 0xdc, 0xbc, 0x56, 0x31, 0x4d, 0x04,
	0xf7, 0x13, 0xa1, 0xb9, 0xd0, 0x43, 0xd9, 0xe8, 0x74, 0xeb, 0x9b, 0x0f, 0xc9, 0x34, 0x7d, 0x28,
	0x4d, 0x38, 0xca, 0xd1, 0x23, 0xfc, 0x2f, 0x15, 0x7b, 0xcb, 0x45, 0xa5, 0xf1, 0xf4, 0x4f, 0x56,
	0x27, 0x3d, 0x9a, 0xc3, 0x69, 0x91, 0xf3, 0xdc, 0x60, 0x6f, 0x05, 0xd6, 0x6e, 0xd8, 0x37, 0xdd,
	0x7f, 0xb5, 0x11, 0x12, 0xcf, 0x2c, 0x69, 0x71, 0xf0, 0x5c, 0x1f, 0x89, 0xb3, 0x3f, 0x12, 0x50,
	0x37, 0x04, 0xec, 0x1a, 0x02, 0x7e, 0x1a, 0x02, 0xbe, 0x5a, 0xe2, 0xec, 0x5a, 0xe2, 0xec, 0x5b,
	0xe2, 0xbc, 0xd0, 0xb3, 0xdb, 0xc7, 0xb4, 0x37, 0x5d, 0xdc, 0xfe, 0xb0, 0x90, 0xf7, 0x11, 0xd8,
	0x97, 0xc4, 0x9e, 0xcd, 0xf9, 0xf6, 0x37, 0x00, 0x00, 0xff, 0xff, 0xc7, 0xdc, 0x58, 0x8e, 0xb0,
	0x01, 0x00, 0x00,
}

func (m *Member) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Member) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Member) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Stop != 0 {
		i = encodeVarintMember(dAtA, i, uint64(m.Stop))
		i--
		dAtA[i] = 0x38
	}
	if m.Limit != 0 {
		i = encodeVarintMember(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.Previous.Size()
		i -= size
		if _, err := m.Previous.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMember(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.Balance.Size()
		i -= size
		if _, err := m.Balance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMember(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.DenomB) > 0 {
		i -= len(m.DenomB)
		copy(dAtA[i:], m.DenomB)
		i = encodeVarintMember(dAtA, i, uint64(len(m.DenomB)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DenomA) > 0 {
		i -= len(m.DenomA)
		copy(dAtA[i:], m.DenomA)
		i = encodeVarintMember(dAtA, i, uint64(len(m.DenomA)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Pair) > 0 {
		i -= len(m.Pair)
		copy(dAtA[i:], m.Pair)
		i = encodeVarintMember(dAtA, i, uint64(len(m.Pair)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMember(dAtA []byte, offset int, v uint64) int {
	offset -= sovMember(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Member) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Pair)
	if l > 0 {
		n += 1 + l + sovMember(uint64(l))
	}
	l = len(m.DenomA)
	if l > 0 {
		n += 1 + l + sovMember(uint64(l))
	}
	l = len(m.DenomB)
	if l > 0 {
		n += 1 + l + sovMember(uint64(l))
	}
	l = m.Balance.Size()
	n += 1 + l + sovMember(uint64(l))
	l = m.Previous.Size()
	n += 1 + l + sovMember(uint64(l))
	if m.Limit != 0 {
		n += 1 + sovMember(uint64(m.Limit))
	}
	if m.Stop != 0 {
		n += 1 + sovMember(uint64(m.Stop))
	}
	return n
}

func sovMember(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMember(x uint64) (n int) {
	return sovMember(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Member) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMember
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
			return fmt.Errorf("proto: Member: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Member: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pair", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
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
				return ErrInvalidLengthMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMember
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pair = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomA", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
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
				return ErrInvalidLengthMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMember
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomA = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomB", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
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
				return ErrInvalidLengthMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMember
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomB = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
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
				return ErrInvalidLengthMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMember
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Previous", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
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
				return ErrInvalidLengthMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMember
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Previous.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stop", wireType)
			}
			m.Stop = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Stop |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMember(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMember
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
func skipMember(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMember
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
					return 0, ErrIntOverflowMember
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
					return 0, ErrIntOverflowMember
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
				return 0, ErrInvalidLengthMember
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMember
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMember
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMember        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMember          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMember = fmt.Errorf("proto: unexpected end of group")
)
