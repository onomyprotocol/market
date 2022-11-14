// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: market/drop.proto

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

type Drop struct {
	Uid    uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Owner  string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Pair   string `protobuf:"bytes,3,opt,name=pair,proto3" json:"pair,omitempty"`
	Drops  string `protobuf:"bytes,4,opt,name=drops,proto3" json:"drops,omitempty"`
	Sum    string `protobuf:"bytes,5,opt,name=sum,proto3" json:"sum,omitempty"`
	Active bool   `protobuf:"varint,6,opt,name=active,proto3" json:"active,omitempty"`
}

func (m *Drop) Reset()         { *m = Drop{} }
func (m *Drop) String() string { return proto.CompactTextString(m) }
func (*Drop) ProtoMessage()    {}
func (*Drop) Descriptor() ([]byte, []int) {
	return fileDescriptor_3961bee11a1276cb, []int{0}
}
func (m *Drop) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Drop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Drop.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Drop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Drop.Merge(m, src)
}
func (m *Drop) XXX_Size() int {
	return m.Size()
}
func (m *Drop) XXX_DiscardUnknown() {
	xxx_messageInfo_Drop.DiscardUnknown(m)
}

var xxx_messageInfo_Drop proto.InternalMessageInfo

func (m *Drop) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *Drop) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Drop) GetPair() string {
	if m != nil {
		return m.Pair
	}
	return ""
}

func (m *Drop) GetDrops() string {
	if m != nil {
		return m.Drops
	}
	return ""
}

func (m *Drop) GetSum() string {
	if m != nil {
		return m.Sum
	}
	return ""
}

func (m *Drop) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func init() {
	proto.RegisterType((*Drop)(nil), "onomyprotocol.market.market.Drop")
}

func init() { proto.RegisterFile("market/drop.proto", fileDescriptor_3961bee11a1276cb) }

var fileDescriptor_3961bee11a1276cb = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0x4d, 0x2c, 0xca,
	0x4e, 0x2d, 0xd1, 0x4f, 0x29, 0xca, 0x2f, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0xce,
	0xcf, 0xcb, 0xcf, 0xad, 0x04, 0xb3, 0x93, 0xf3, 0x73, 0xf4, 0x20, 0x0a, 0xa0, 0x94, 0x52, 0x13,
	0x23, 0x17, 0x8b, 0x4b, 0x51, 0x7e, 0x81, 0x90, 0x00, 0x17, 0x73, 0x69, 0x66, 0x8a, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x4b, 0x10, 0x88, 0x29, 0x24, 0xc2, 0xc5, 0x9a, 0x5f, 0x9e, 0x97, 0x5a, 0x24,
	0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe1, 0x08, 0x09, 0x71, 0xb1, 0x14, 0x24, 0x66, 0x16,
	0x49, 0x30, 0x83, 0x05, 0xc1, 0x6c, 0x90, 0x4a, 0x90, 0x7d, 0xc5, 0x12, 0x2c, 0x10, 0x95, 0x60,
	0x0e, 0xc8, 0xc4, 0xe2, 0xd2, 0x5c, 0x09, 0x56, 0xb0, 0x18, 0x88, 0x29, 0x24, 0xc6, 0xc5, 0x96,
	0x98, 0x5c, 0x92, 0x59, 0x96, 0x2a, 0xc1, 0xa6, 0xc0, 0xa8, 0xc1, 0x11, 0x04, 0xe5, 0x39, 0x79,
	0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb,
	0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x5e, 0x7a, 0x66, 0x49, 0x46,
	0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0x8a, 0x37, 0xf4, 0xa1, 0xfe, 0xac, 0x80, 0x31, 0x4a,
	0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0xf2, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9c,
	0x32, 0xce, 0x66, 0x07, 0x01, 0x00, 0x00,
}

func (m *Drop) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Drop) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Drop) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Active {
		i--
		if m.Active {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.Sum) > 0 {
		i -= len(m.Sum)
		copy(dAtA[i:], m.Sum)
		i = encodeVarintDrop(dAtA, i, uint64(len(m.Sum)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Drops) > 0 {
		i -= len(m.Drops)
		copy(dAtA[i:], m.Drops)
		i = encodeVarintDrop(dAtA, i, uint64(len(m.Drops)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Pair) > 0 {
		i -= len(m.Pair)
		copy(dAtA[i:], m.Pair)
		i = encodeVarintDrop(dAtA, i, uint64(len(m.Pair)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintDrop(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if m.Uid != 0 {
		i = encodeVarintDrop(dAtA, i, uint64(m.Uid))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDrop(dAtA []byte, offset int, v uint64) int {
	offset -= sovDrop(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Drop) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Uid != 0 {
		n += 1 + sovDrop(uint64(m.Uid))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovDrop(uint64(l))
	}
	l = len(m.Pair)
	if l > 0 {
		n += 1 + l + sovDrop(uint64(l))
	}
	l = len(m.Drops)
	if l > 0 {
		n += 1 + l + sovDrop(uint64(l))
	}
	l = len(m.Sum)
	if l > 0 {
		n += 1 + l + sovDrop(uint64(l))
	}
	if m.Active {
		n += 2
	}
	return n
}

func sovDrop(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDrop(x uint64) (n int) {
	return sovDrop(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Drop) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDrop
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
			return fmt.Errorf("proto: Drop: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Drop: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrop
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= uint64(b&0x7F) << shift
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
					return ErrIntOverflowDrop
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
				return ErrInvalidLengthDrop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDrop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pair", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrop
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
				return ErrInvalidLengthDrop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDrop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pair = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Drops", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrop
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
				return ErrInvalidLengthDrop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDrop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Drops = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrop
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
				return ErrInvalidLengthDrop
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDrop
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Active", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrop
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
			m.Active = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipDrop(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDrop
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
func skipDrop(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDrop
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
					return 0, ErrIntOverflowDrop
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
					return 0, ErrIntOverflowDrop
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
				return 0, ErrInvalidLengthDrop
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDrop
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDrop
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDrop        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDrop          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDrop = fmt.Errorf("proto: unexpected end of group")
)
