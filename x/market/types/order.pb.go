// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: market/order.proto

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

type Order struct {
	Uid       uint64                                   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Owner     string                                   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Status    string                                   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	OrderType string                                   `protobuf:"bytes,4,opt,name=orderType,proto3" json:"orderType,omitempty"`
	DenomAsk  string                                   `protobuf:"bytes,5,opt,name=denomAsk,proto3" json:"denomAsk,omitempty"`
	DenomBid  string                                   `protobuf:"bytes,6,opt,name=denomBid,proto3" json:"denomBid,omitempty"`
	Amount    github_com_cosmos_cosmos_sdk_types.Int   `protobuf:"bytes,7,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
	Rate      []github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,8,rep,name=rate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"rate"`
	Prev      uint64                                   `protobuf:"varint,9,opt,name=prev,proto3" json:"prev,omitempty"`
	Next      uint64                                   `protobuf:"varint,10,opt,name=next,proto3" json:"next,omitempty"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c6375df0c4a1904, []int{0}
}
func (m *Order) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Order.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return m.Size()
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

type Orders struct {
	Uids []uint64 `protobuf:"varint,1,rep,packed,name=uids,proto3" json:"uids,omitempty"`
}

func (m *Orders) Reset()         { *m = Orders{} }
func (m *Orders) String() string { return proto.CompactTextString(m) }
func (*Orders) ProtoMessage()    {}
func (*Orders) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c6375df0c4a1904, []int{1}
}
func (m *Orders) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Orders) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Orders.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Orders) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Orders.Merge(m, src)
}
func (m *Orders) XXX_Size() int {
	return m.Size()
}
func (m *Orders) XXX_DiscardUnknown() {
	xxx_messageInfo_Orders.DiscardUnknown(m)
}

var xxx_messageInfo_Orders proto.InternalMessageInfo

type OrderResponse struct {
	Uid       uint64   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Owner     string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Status    string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	OrderType string   `protobuf:"bytes,4,opt,name=orderType,proto3" json:"orderType,omitempty"`
	DenomAsk  string   `protobuf:"bytes,5,opt,name=denomAsk,proto3" json:"denomAsk,omitempty"`
	DenomBid  string   `protobuf:"bytes,6,opt,name=denomBid,proto3" json:"denomBid,omitempty"`
	Amount    string   `protobuf:"bytes,7,opt,name=amount,proto3" json:"amount,omitempty"`
	Rate      []string `protobuf:"bytes,8,rep,name=rate,proto3" json:"rate,omitempty"`
	Prev      uint64   `protobuf:"varint,9,opt,name=prev,proto3" json:"prev,omitempty"`
	Next      uint64   `protobuf:"varint,10,opt,name=next,proto3" json:"next,omitempty"`
}

func (m *OrderResponse) Reset()         { *m = OrderResponse{} }
func (m *OrderResponse) String() string { return proto.CompactTextString(m) }
func (*OrderResponse) ProtoMessage()    {}
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c6375df0c4a1904, []int{2}
}
func (m *OrderResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrderResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderResponse.Merge(m, src)
}
func (m *OrderResponse) XXX_Size() int {
	return m.Size()
}
func (m *OrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OrderResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Order)(nil), "pendulumlabs.market.market.Order")
	proto.RegisterType((*Orders)(nil), "pendulumlabs.market.market.Orders")
	proto.RegisterType((*OrderResponse)(nil), "pendulumlabs.market.market.OrderResponse")
}

func init() { proto.RegisterFile("market/order.proto", fileDescriptor_8c6375df0c4a1904) }

var fileDescriptor_8c6375df0c4a1904 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x92, 0xc1, 0x4e, 0xc2, 0x30,
	0x18, 0xc7, 0x57, 0x36, 0x26, 0x34, 0x31, 0x31, 0x0d, 0x21, 0x0d, 0x21, 0x65, 0xe1, 0x60, 0xb8,
	0xb0, 0x1d, 0x7c, 0x02, 0x77, 0x30, 0xf1, 0xa4, 0x59, 0x3c, 0x79, 0x1b, 0xb4, 0xc1, 0x05, 0xb6,
	0x2e, 0x6b, 0xab, 0xf0, 0x16, 0x3e, 0x16, 0x47, 0xbc, 0x11, 0x0f, 0x44, 0xe0, 0x29, 0xbc, 0x99,
	0x76, 0x03, 0xf1, 0xa8, 0x27, 0x4f, 0xfd, 0x7f, 0xdf, 0xaf, 0x5f, 0xd3, 0xfc, 0xbf, 0x3f, 0x44,
	0x69, 0x5c, 0x4c, 0x99, 0x0c, 0x78, 0x41, 0x59, 0xe1, 0xe7, 0x05, 0x97, 0x1c, 0x75, 0x72, 0x96,
	0x51, 0x35, 0x53, 0xe9, 0x2c, 0x1e, 0x09, 0xbf, 0xbc, 0x50, 0x1d, 0x9d, 0xd6, 0x84, 0x4f, 0xb8,
	0xb9, 0x16, 0x68, 0x55, 0x4e, 0xf4, 0xdf, 0x6a, 0xb0, 0x7e, 0xa7, 0x5f, 0x40, 0x17, 0xd0, 0x56,
	0x09, 0xc5, 0xc0, 0x03, 0x03, 0x27, 0xd2, 0x12, 0xb5, 0x60, 0x9d, 0xbf, 0x64, 0xac, 0xc0, 0x35,
	0x0f, 0x0c, 0x9a, 0x51, 0x59, 0xa0, 0x36, 0x74, 0x85, 0x8c, 0xa5, 0x12, 0xd8, 0x36, 0xed, 0xaa,
	0x42, 0x5d, 0xd8, 0x34, 0x5f, 0x79, 0x58, 0xe4, 0x0c, 0x3b, 0x06, 0x7d, 0x37, 0x50, 0x07, 0x36,
	0x28, 0xcb, 0x78, 0x7a, 0x2d, 0xa6, 0xb8, 0x6e, 0xe0, 0xb1, 0x3e, 0xb2, 0x30, 0xa1, 0xd8, 0x3d,
	0x61, 0x61, 0x42, 0xd1, 0x0d, 0x74, 0xe3, 0x94, 0xab, 0x4c, 0xe2, 0x33, 0x4d, 0x42, 0x7f, 0xb9,
	0xe9, 0x59, 0xef, 0x9b, 0xde, 0xe5, 0x24, 0x91, 0x4f, 0x6a, 0xe4, 0x8f, 0x79, 0x1a, 0x8c, 0xb9,
	0x48, 0xb9, 0xa8, 0x8e, 0xa1, 0xa0, 0xd3, 0x40, 0x2e, 0x72, 0x26, 0xfc, 0xdb, 0x4c, 0x46, 0xd5,
	0x34, 0x0a, 0xa1, 0x53, 0xc4, 0x92, 0xe1, 0x86, 0x67, 0xff, 0xe1, 0x15, 0x33, 0x8b, 0x10, 0x74,
	0xf2, 0x82, 0x3d, 0xe3, 0xa6, 0xb1, 0xc8, 0x68, 0xdd, 0xcb, 0xd8, 0x5c, 0x62, 0x58, 0xf6, 0xb4,
	0xee, 0x77, 0xa1, 0x6b, 0x2c, 0x15, 0x9a, 0xaa, 0x84, 0x0a, 0x0c, 0x3c, 0x5b, 0x53, 0xad, 0xfb,
	0x9f, 0x00, 0x9e, 0x1b, 0x1c, 0x31, 0x91, 0xf3, 0x4c, 0xb0, 0x7f, 0xea, 0x7c, 0xfb, 0xa7, 0xf3,
	0x47, 0x27, 0xd1, 0xa9, 0x93, 0xbf, 0x73, 0x26, 0xbc, 0x5f, 0x6e, 0x89, 0xb5, 0xde, 0x12, 0xb0,
	0xdc, 0x11, 0xb0, 0xda, 0x11, 0xf0, 0xb1, 0x23, 0xe0, 0x75, 0x4f, 0xac, 0xd5, 0x9e, 0x58, 0xeb,
	0x3d, 0xb1, 0x1e, 0xfd, 0x93, 0x8d, 0x1c, 0xc2, 0x3c, 0xd4, 0x69, 0x0e, 0xaa, 0xb8, 0xcf, 0x0f,
	0xc2, 0x6c, 0x67, 0xe4, 0x9a, 0x18, 0x5f, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x1b, 0xf8,
	0xee, 0x0e, 0x03, 0x00, 0x00,
}

func (m *Order) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Order) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Order) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Next != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.Next))
		i--
		dAtA[i] = 0x50
	}
	if m.Prev != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.Prev))
		i--
		dAtA[i] = 0x48
	}
	if len(m.Rate) > 0 {
		for iNdEx := len(m.Rate) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Rate[iNdEx].Size()
				i -= size
				if _, err := m.Rate[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintOrder(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.DenomBid) > 0 {
		i -= len(m.DenomBid)
		copy(dAtA[i:], m.DenomBid)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.DenomBid)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.DenomAsk) > 0 {
		i -= len(m.DenomAsk)
		copy(dAtA[i:], m.DenomAsk)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.DenomAsk)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.OrderType) > 0 {
		i -= len(m.OrderType)
		copy(dAtA[i:], m.OrderType)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.OrderType)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if m.Uid != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.Uid))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Orders) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Orders) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Orders) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Uids) > 0 {
		dAtA2 := make([]byte, len(m.Uids)*10)
		var j1 int
		for _, num := range m.Uids {
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
		i = encodeVarintOrder(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OrderResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrderResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Next != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.Next))
		i--
		dAtA[i] = 0x50
	}
	if m.Prev != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.Prev))
		i--
		dAtA[i] = 0x48
	}
	if len(m.Rate) > 0 {
		for iNdEx := len(m.Rate) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Rate[iNdEx])
			copy(dAtA[i:], m.Rate[iNdEx])
			i = encodeVarintOrder(dAtA, i, uint64(len(m.Rate[iNdEx])))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.DenomBid) > 0 {
		i -= len(m.DenomBid)
		copy(dAtA[i:], m.DenomBid)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.DenomBid)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.DenomAsk) > 0 {
		i -= len(m.DenomAsk)
		copy(dAtA[i:], m.DenomAsk)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.DenomAsk)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.OrderType) > 0 {
		i -= len(m.OrderType)
		copy(dAtA[i:], m.OrderType)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.OrderType)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if m.Uid != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.Uid))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintOrder(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Order) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Uid != 0 {
		n += 1 + sovOrder(uint64(m.Uid))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	l = len(m.OrderType)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	l = len(m.DenomAsk)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	l = len(m.DenomBid)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovOrder(uint64(l))
	if len(m.Rate) > 0 {
		for _, e := range m.Rate {
			l = e.Size()
			n += 1 + l + sovOrder(uint64(l))
		}
	}
	if m.Prev != 0 {
		n += 1 + sovOrder(uint64(m.Prev))
	}
	if m.Next != 0 {
		n += 1 + sovOrder(uint64(m.Next))
	}
	return n
}

func (m *Orders) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Uids) > 0 {
		l = 0
		for _, e := range m.Uids {
			l += sovOrder(uint64(e))
		}
		n += 1 + sovOrder(uint64(l)) + l
	}
	return n
}

func (m *OrderResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Uid != 0 {
		n += 1 + sovOrder(uint64(m.Uid))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	l = len(m.OrderType)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	l = len(m.DenomAsk)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	l = len(m.DenomBid)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	if len(m.Rate) > 0 {
		for _, s := range m.Rate {
			l = len(s)
			n += 1 + l + sovOrder(uint64(l))
		}
	}
	if m.Prev != 0 {
		n += 1 + sovOrder(uint64(m.Prev))
	}
	if m.Next != 0 {
		n += 1 + sovOrder(uint64(m.Next))
	}
	return n
}

func sovOrder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrder(x uint64) (n int) {
	return sovOrder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Order) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrder
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
			return fmt.Errorf("proto: Order: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Order: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomAsk", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomAsk = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomBid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomBid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.Rate = append(m.Rate, v)
			if err := m.Rate[len(m.Rate)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prev", wireType)
			}
			m.Prev = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Prev |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Next", wireType)
			}
			m.Next = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Next |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrder
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
func (m *Orders) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrder
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
			return fmt.Errorf("proto: Orders: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Orders: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowOrder
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
				m.Uids = append(m.Uids, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowOrder
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
					return ErrInvalidLengthOrder
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthOrder
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
				if elementCount != 0 && len(m.Uids) == 0 {
					m.Uids = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowOrder
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
					m.Uids = append(m.Uids, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Uids", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrder
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
func (m *OrderResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrder
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
			return fmt.Errorf("proto: OrderResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrderResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomAsk", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomAsk = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomBid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomBid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rate = append(m.Rate, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prev", wireType)
			}
			m.Prev = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Prev |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Next", wireType)
			}
			m.Next = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Next |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrder
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
func skipOrder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrder
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
					return 0, ErrIntOverflowOrder
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
					return 0, ErrIntOverflowOrder
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
				return 0, ErrInvalidLengthOrder
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrder
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrder
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrder        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrder          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrder = fmt.Errorf("proto: unexpected end of group")
)
