// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: portal/packet.proto

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

type PortalPacketData struct {
	// Types that are valid to be assigned to Packet:
	//	*PortalPacketData_NoData
	//	*PortalPacketData_SubscribeRatePacket
	Packet isPortalPacketData_Packet `protobuf_oneof:"packet"`
}

func (m *PortalPacketData) Reset()         { *m = PortalPacketData{} }
func (m *PortalPacketData) String() string { return proto.CompactTextString(m) }
func (*PortalPacketData) ProtoMessage()    {}
func (*PortalPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_acf0e65d77ee4aa1, []int{0}
}
func (m *PortalPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PortalPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PortalPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PortalPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortalPacketData.Merge(m, src)
}
func (m *PortalPacketData) XXX_Size() int {
	return m.Size()
}
func (m *PortalPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_PortalPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_PortalPacketData proto.InternalMessageInfo

type isPortalPacketData_Packet interface {
	isPortalPacketData_Packet()
	MarshalTo([]byte) (int, error)
	Size() int
}

type PortalPacketData_NoData struct {
	NoData *NoData `protobuf:"bytes,1,opt,name=noData,proto3,oneof" json:"noData,omitempty"`
}
type PortalPacketData_SubscribeRatePacket struct {
	SubscribeRatePacket *SubscribeRatePacketData `protobuf:"bytes,2,opt,name=subscribeRatePacket,proto3,oneof" json:"subscribeRatePacket,omitempty"`
}

func (*PortalPacketData_NoData) isPortalPacketData_Packet()              {}
func (*PortalPacketData_SubscribeRatePacket) isPortalPacketData_Packet() {}

func (m *PortalPacketData) GetPacket() isPortalPacketData_Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

func (m *PortalPacketData) GetNoData() *NoData {
	if x, ok := m.GetPacket().(*PortalPacketData_NoData); ok {
		return x.NoData
	}
	return nil
}

func (m *PortalPacketData) GetSubscribeRatePacket() *SubscribeRatePacketData {
	if x, ok := m.GetPacket().(*PortalPacketData_SubscribeRatePacket); ok {
		return x.SubscribeRatePacket
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PortalPacketData) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PortalPacketData_NoData)(nil),
		(*PortalPacketData_SubscribeRatePacket)(nil),
	}
}

type NoData struct {
}

func (m *NoData) Reset()         { *m = NoData{} }
func (m *NoData) String() string { return proto.CompactTextString(m) }
func (*NoData) ProtoMessage()    {}
func (*NoData) Descriptor() ([]byte, []int) {
	return fileDescriptor_acf0e65d77ee4aa1, []int{1}
}
func (m *NoData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NoData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NoData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NoData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoData.Merge(m, src)
}
func (m *NoData) XXX_Size() int {
	return m.Size()
}
func (m *NoData) XXX_DiscardUnknown() {
	xxx_messageInfo_NoData.DiscardUnknown(m)
}

var xxx_messageInfo_NoData proto.InternalMessageInfo

// SubscribeRatePacketData defines a struct for the packet payload
type SubscribeRatePacketData struct {
	DenomA string `protobuf:"bytes,1,opt,name=denomA,proto3" json:"denomA,omitempty"`
	DenomB string `protobuf:"bytes,2,opt,name=denomB,proto3" json:"denomB,omitempty"`
}

func (m *SubscribeRatePacketData) Reset()         { *m = SubscribeRatePacketData{} }
func (m *SubscribeRatePacketData) String() string { return proto.CompactTextString(m) }
func (*SubscribeRatePacketData) ProtoMessage()    {}
func (*SubscribeRatePacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_acf0e65d77ee4aa1, []int{2}
}
func (m *SubscribeRatePacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubscribeRatePacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubscribeRatePacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubscribeRatePacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRatePacketData.Merge(m, src)
}
func (m *SubscribeRatePacketData) XXX_Size() int {
	return m.Size()
}
func (m *SubscribeRatePacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRatePacketData.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRatePacketData proto.InternalMessageInfo

func (m *SubscribeRatePacketData) GetDenomA() string {
	if m != nil {
		return m.DenomA
	}
	return ""
}

func (m *SubscribeRatePacketData) GetDenomB() string {
	if m != nil {
		return m.DenomB
	}
	return ""
}

// SubscribeRatePacketAck defines a struct for the packet acknowledgment
type SubscribeRatePacketAck struct {
}

func (m *SubscribeRatePacketAck) Reset()         { *m = SubscribeRatePacketAck{} }
func (m *SubscribeRatePacketAck) String() string { return proto.CompactTextString(m) }
func (*SubscribeRatePacketAck) ProtoMessage()    {}
func (*SubscribeRatePacketAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_acf0e65d77ee4aa1, []int{3}
}
func (m *SubscribeRatePacketAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubscribeRatePacketAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubscribeRatePacketAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubscribeRatePacketAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRatePacketAck.Merge(m, src)
}
func (m *SubscribeRatePacketAck) XXX_Size() int {
	return m.Size()
}
func (m *SubscribeRatePacketAck) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRatePacketAck.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRatePacketAck proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PortalPacketData)(nil), "market.portal.PortalPacketData")
	proto.RegisterType((*NoData)(nil), "market.portal.NoData")
	proto.RegisterType((*SubscribeRatePacketData)(nil), "market.portal.SubscribeRatePacketData")
	proto.RegisterType((*SubscribeRatePacketAck)(nil), "market.portal.SubscribeRatePacketAck")
}

func init() { proto.RegisterFile("portal/packet.proto", fileDescriptor_acf0e65d77ee4aa1) }

var fileDescriptor_acf0e65d77ee4aa1 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0xc8, 0x2f, 0x2a,
	0x49, 0xcc, 0xd1, 0x2f, 0x48, 0x4c, 0xce, 0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0xcd, 0x4d, 0x2c, 0x02, 0xf3, 0xc0, 0x72, 0x4a, 0x2b, 0x19, 0xb9, 0x04, 0x02, 0xc0, 0xcc,
	0x00, 0xb0, 0x2a, 0x97, 0xc4, 0x92, 0x44, 0x21, 0x7d, 0x2e, 0xb6, 0xbc, 0x7c, 0x10, 0x4b, 0x82,
	0x51, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x54, 0x0f, 0x45, 0x93, 0x9e, 0x1f, 0x58, 0xd2, 0x83, 0x21,
	0x08, 0xaa, 0x4c, 0x28, 0x8a, 0x4b, 0xb8, 0xb8, 0x34, 0xa9, 0x38, 0xb9, 0x28, 0x33, 0x29, 0x35,
	0x28, 0xb1, 0x24, 0x15, 0x62, 0x96, 0x04, 0x13, 0x58, 0xb7, 0x1a, 0x9a, 0xee, 0x60, 0x4c, 0x95,
	0x50, 0xe3, 0xb0, 0x19, 0xe2, 0xc4, 0xc1, 0xc5, 0x06, 0xf1, 0x80, 0x12, 0x07, 0x17, 0x1b, 0xc4,
	0x66, 0x25, 0x4f, 0x2e, 0x71, 0x1c, 0xa6, 0x08, 0x89, 0x71, 0xb1, 0xa5, 0xa4, 0xe6, 0xe5, 0xe7,
	0x3a, 0x82, 0xdd, 0xce, 0x19, 0x04, 0xe5, 0xc1, 0xc5, 0x9d, 0xc0, 0xae, 0x82, 0x89, 0x3b, 0x29,
	0x49, 0x70, 0x89, 0x61, 0x31, 0xca, 0x31, 0x39, 0xdb, 0x49, 0xff, 0xc4, 0x23, 0x39, 0xc6, 0x0b,
	0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86,
	0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x44, 0x21, 0x1e, 0xd2, 0xaf, 0xd0, 0x87, 0x86, 0x70, 0x49, 0x65,
	0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38, 0x84, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x98, 0xab,
	0x05, 0x61, 0x78, 0x01, 0x00, 0x00,
}

func (m *PortalPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PortalPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PortalPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Packet != nil {
		{
			size := m.Packet.Size()
			i -= size
			if _, err := m.Packet.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *PortalPacketData_NoData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PortalPacketData_NoData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NoData != nil {
		{
			size, err := m.NoData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPacket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *PortalPacketData_SubscribeRatePacket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PortalPacketData_SubscribeRatePacket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SubscribeRatePacket != nil {
		{
			size, err := m.SubscribeRatePacket.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPacket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *NoData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NoData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NoData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *SubscribeRatePacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubscribeRatePacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubscribeRatePacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DenomB) > 0 {
		i -= len(m.DenomB)
		copy(dAtA[i:], m.DenomB)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.DenomB)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DenomA) > 0 {
		i -= len(m.DenomA)
		copy(dAtA[i:], m.DenomA)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.DenomA)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SubscribeRatePacketAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubscribeRatePacketAck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubscribeRatePacketAck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PortalPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Packet != nil {
		n += m.Packet.Size()
	}
	return n
}

func (m *PortalPacketData_NoData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NoData != nil {
		l = m.NoData.Size()
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}
func (m *PortalPacketData_SubscribeRatePacket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SubscribeRatePacket != nil {
		l = m.SubscribeRatePacket.Size()
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}
func (m *NoData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *SubscribeRatePacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DenomA)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.DenomB)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func (m *SubscribeRatePacketAck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPacket(x uint64) (n int) {
	return sovPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PortalPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: PortalPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PortalPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &NoData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &PortalPacketData_NoData{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubscribeRatePacket", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &SubscribeRatePacketData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &PortalPacketData_SubscribeRatePacket{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *NoData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: NoData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NoData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *SubscribeRatePacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: SubscribeRatePacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubscribeRatePacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomA", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomA = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomB", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomB = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *SubscribeRatePacketAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: SubscribeRatePacketAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubscribeRatePacketAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func skipPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
				return 0, ErrInvalidLengthPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPacket = fmt.Errorf("proto: unexpected end of group")
)
