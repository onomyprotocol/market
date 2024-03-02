// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: market/proposal.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/x/bank/types"
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

// FundTreasuryProposal details a dao fund treasury proposal.
type DenomMetadataProposal struct {
	Sender      string          `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Title       string          `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string          `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Metadata    *types.Metadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *DenomMetadataProposal) Reset()         { *m = DenomMetadataProposal{} }
func (m *DenomMetadataProposal) String() string { return proto.CompactTextString(m) }
func (*DenomMetadataProposal) ProtoMessage()    {}
func (*DenomMetadataProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_b890bb392e4d4e20, []int{0}
}
func (m *DenomMetadataProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DenomMetadataProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DenomMetadataProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DenomMetadataProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DenomMetadataProposal.Merge(m, src)
}
func (m *DenomMetadataProposal) XXX_Size() int {
	return m.Size()
}
func (m *DenomMetadataProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_DenomMetadataProposal.DiscardUnknown(m)
}

var xxx_messageInfo_DenomMetadataProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DenomMetadataProposal)(nil), "pendulumlabs.market.market.DenomMetadataProposal")
}

func init() { proto.RegisterFile("market/proposal.proto", fileDescriptor_b890bb392e4d4e20) }

var fileDescriptor_b890bb392e4d4e20 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x63, 0x7e, 0xaa, 0xe2, 0x6e, 0x51, 0x8b, 0xa2, 0x48, 0xb8, 0x11, 0x53, 0x17, 0x6c,
	0x15, 0x26, 0x18, 0x11, 0x23, 0x48, 0xa8, 0x23, 0x9b, 0x93, 0x58, 0x21, 0x6a, 0x9c, 0x6b, 0xc5,
	0x0e, 0x82, 0x37, 0x60, 0xe4, 0x11, 0x3a, 0x32, 0xf0, 0x20, 0x8c, 0x1d, 0x19, 0x51, 0xb2, 0xf0,
	0x18, 0xa8, 0xb6, 0x83, 0x98, 0x7c, 0xcf, 0xbd, 0xc7, 0x3a, 0x47, 0x1f, 0x9e, 0x49, 0xde, 0xac,
	0x85, 0x61, 0xaa, 0x01, 0x05, 0x9a, 0x57, 0x54, 0x35, 0x60, 0x20, 0x8c, 0x95, 0xa8, 0xf3, 0xb6,
	0x6a, 0x65, 0xc5, 0x53, 0x4d, 0x9d, 0xc7, 0x3f, 0xf1, 0xb4, 0x80, 0x02, 0xac, 0x8d, 0xed, 0x26,
	0xf7, 0x23, 0x26, 0x19, 0x68, 0x09, 0x9a, 0xa5, 0xbc, 0x5e, 0xb3, 0xa7, 0x65, 0x2a, 0x0c, 0x5f,
	0x5a, 0xe1, 0xee, 0xa7, 0x1f, 0x08, 0xcf, 0x6e, 0x44, 0x0d, 0xf2, 0x4e, 0x18, 0x9e, 0x73, 0xc3,
	0xef, 0x7d, 0x62, 0x78, 0x8c, 0x47, 0x5a, 0xd4, 0xb9, 0x68, 0x22, 0x94, 0xa0, 0xc5, 0xd1, 0xca,
	0xab, 0x70, 0x8a, 0x0f, 0x4d, 0x69, 0x2a, 0x11, 0xed, 0xd9, 0xb5, 0x13, 0x61, 0x82, 0x27, 0xb9,
	0xd0, 0x59, 0x53, 0x2a, 0x53, 0x42, 0x1d, 0xed, 0xdb, 0xdb, 0xff, 0x55, 0x78, 0x89, 0xc7, 0xd2,
	0x67, 0x44, 0x07, 0x09, 0x5a, 0x4c, 0xce, 0x4f, 0xa8, 0x2b, 0x47, 0x6d, 0x1f, 0x5f, 0x8e, 0x0e,
	0x45, 0x56, 0x7f, 0xf6, 0xab, 0xf1, 0xeb, 0x66, 0x1e, 0xfc, 0x6c, 0xe6, 0xc1, 0xf5, 0xed, 0x7b,
	0x47, 0xd0, 0x67, 0x47, 0xd0, 0xb6, 0x23, 0xe8, 0xbb, 0x23, 0xe8, 0xad, 0x27, 0xc1, 0xb6, 0x27,
	0xc1, 0x57, 0x4f, 0x82, 0x07, 0x5a, 0x94, 0xe6, 0xb1, 0x4d, 0x69, 0x06, 0x92, 0x0d, 0xa4, 0xce,
	0x76, 0xa8, 0x98, 0xc7, 0xf9, 0x3c, 0x0c, 0xe6, 0x45, 0x09, 0x9d, 0x8e, 0x2c, 0x83, 0x8b, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xac, 0xe8, 0x18, 0x5f, 0x6e, 0x01, 0x00, 0x00,
}

func (m *DenomMetadataProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DenomMetadataProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DenomMetadataProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintProposal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DenomMetadataProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DenomMetadataProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: DenomMetadataProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DenomMetadataProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &types.Metadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
