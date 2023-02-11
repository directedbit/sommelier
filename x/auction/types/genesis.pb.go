// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auction/v1/genesis.proto

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

type GenesisState struct {
	Params        Params        `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Auctions      []*Auction    `protobuf:"bytes,2,rep,name=auctions,proto3" json:"auctions,omitempty"`
	Bids          []*Bid        `protobuf:"bytes,3,rep,name=bids,proto3" json:"bids,omitempty"`
	TokenPrices   []*TokenPrice `protobuf:"bytes,4,rep,name=token_prices,json=tokenPrices,proto3" json:"token_prices,omitempty"`
	LastAuctionId uint32        `protobuf:"varint,5,opt,name=last_auction_id,json=lastAuctionId,proto3" json:"last_auction_id,omitempty"`
	LastBidId     uint64        `protobuf:"varint,6,opt,name=last_bid_id,json=lastBidId,proto3" json:"last_bid_id,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a762e9d6ba7af420, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetAuctions() []*Auction {
	if m != nil {
		return m.Auctions
	}
	return nil
}

func (m *GenesisState) GetBids() []*Bid {
	if m != nil {
		return m.Bids
	}
	return nil
}

func (m *GenesisState) GetTokenPrices() []*TokenPrice {
	if m != nil {
		return m.TokenPrices
	}
	return nil
}

func (m *GenesisState) GetLastAuctionId() uint32 {
	if m != nil {
		return m.LastAuctionId
	}
	return 0
}

func (m *GenesisState) GetLastBidId() uint64 {
	if m != nil {
		return m.LastBidId
	}
	return 0
}

type Params struct {
	PriceMaxBlockAge                     uint64                                 `protobuf:"varint,1,opt,name=price_max_block_age,json=priceMaxBlockAge,proto3" json:"price_max_block_age,omitempty"`
	MinimumBidInUsomm                    uint64                                 `protobuf:"varint,2,opt,name=minimum_bid_in_usomm,json=minimumBidInUsomm,proto3" json:"minimum_bid_in_usomm,omitempty"`
	MinimumSaleTokensUsdValue            github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=minimum_sale_tokens_usd_value,json=minimumSaleTokensUsdValue,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"minimum_sale_tokens_usd_value"`
	AuctionMaxBlockAge                   uint64                                 `protobuf:"varint,4,opt,name=auction_max_block_age,json=auctionMaxBlockAge,proto3" json:"auction_max_block_age,omitempty"`
	AuctionPriceDecreaseAccelerationRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=auction_price_decrease_acceleration_rate,json=auctionPriceDecreaseAccelerationRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"auction_price_decrease_acceleration_rate"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_a762e9d6ba7af420, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetPriceMaxBlockAge() uint64 {
	if m != nil {
		return m.PriceMaxBlockAge
	}
	return 0
}

func (m *Params) GetMinimumBidInUsomm() uint64 {
	if m != nil {
		return m.MinimumBidInUsomm
	}
	return 0
}

func (m *Params) GetAuctionMaxBlockAge() uint64 {
	if m != nil {
		return m.AuctionMaxBlockAge
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "auction.v1.GenesisState")
	proto.RegisterType((*Params)(nil), "auction.v1.Params")
}

func init() { proto.RegisterFile("auction/v1/genesis.proto", fileDescriptor_a762e9d6ba7af420) }

var fileDescriptor_a762e9d6ba7af420 = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0xe3, 0xc6, 0x44, 0xf4, 0xd2, 0xaa, 0x70, 0x2d, 0xc8, 0x54, 0xc2, 0x8d, 0x0a, 0xaa,
	0xbc, 0xd4, 0x26, 0x85, 0x85, 0x31, 0x56, 0x25, 0x94, 0x01, 0xa9, 0x72, 0x29, 0x03, 0xcb, 0xe9,
	0xe2, 0x7b, 0x65, 0x8e, 0xd8, 0x3e, 0xcb, 0x77, 0x8e, 0xd2, 0x2f, 0xc0, 0xcc, 0x84, 0xc4, 0x37,
	0xea, 0xd8, 0x11, 0x31, 0x54, 0x28, 0xf9, 0x22, 0xe8, 0xce, 0x97, 0x62, 0x46, 0x26, 0xbf, 0x7a,
	0x9f, 0xdf, 0xfb, 0xef, 0xb1, 0x0e, 0x79, 0xb4, 0x49, 0x15, 0x17, 0x65, 0xb4, 0x18, 0x47, 0x19,
	0x94, 0x20, 0xb9, 0x0c, 0xab, 0x5a, 0x28, 0x81, 0x91, 0x55, 0xc2, 0xc5, 0xf8, 0x70, 0xbf, 0x43,
	0xa9, 0x65, 0x0b, 0x1c, 0x76, 0x4b, 0x37, 0x6c, 0xab, 0x1c, 0x64, 0x22, 0x13, 0x26, 0x8c, 0x74,
	0xd4, 0x66, 0x8f, 0x7f, 0x6c, 0xa1, 0x9d, 0x77, 0xed, 0x88, 0x4b, 0x45, 0x15, 0xe0, 0x57, 0x68,
	0x50, 0xd1, 0x9a, 0x16, 0xd2, 0x73, 0x46, 0x4e, 0x30, 0x3c, 0xc3, 0xe1, 0xdf, 0x91, 0xe1, 0x85,
	0x51, 0x62, 0xf7, 0xe6, 0xee, 0xa8, 0x97, 0x58, 0x0e, 0x47, 0xe8, 0xa1, 0x45, 0xa4, 0xb7, 0x35,
	0xea, 0x07, 0xc3, 0xb3, 0xfd, 0x6e, 0xcd, 0xa4, 0x0d, 0x93, 0x7b, 0x08, 0xbf, 0x40, 0xee, 0x8c,
	0x33, 0xe9, 0xf5, 0x0d, 0xbc, 0xd7, 0x85, 0x63, 0xce, 0x12, 0x23, 0xe2, 0xb7, 0x68, 0x47, 0x89,
	0x39, 0x94, 0xa4, 0xaa, 0x79, 0x0a, 0xd2, 0x73, 0x0d, 0xfc, 0xb4, 0x0b, 0x7f, 0xd0, 0xfa, 0x85,
	0x96, 0x93, 0xa1, 0xba, 0x8f, 0x25, 0x3e, 0x41, 0x7b, 0x39, 0x95, 0x8a, 0x58, 0x94, 0x70, 0xe6,
	0x3d, 0x18, 0x39, 0xc1, 0x6e, 0xb2, 0xab, 0xd3, 0x76, 0x9f, 0x29, 0xc3, 0x3e, 0x1a, 0x1a, 0x6e,
	0xc6, 0x99, 0x66, 0x06, 0x23, 0x27, 0x70, 0x93, 0x6d, 0x9d, 0x8a, 0x39, 0x9b, 0xb2, 0xe3, 0xef,
	0x7d, 0x34, 0x68, 0x2f, 0xc6, 0xa7, 0x68, 0xdf, 0xec, 0x41, 0x0a, 0xba, 0x24, 0xb3, 0x5c, 0xa4,
	0x73, 0x42, 0x33, 0x30, 0x16, 0xb9, 0xc9, 0x23, 0x23, 0xbd, 0xa7, 0xcb, 0x58, 0x0b, 0x93, 0x0c,
	0x70, 0x84, 0x0e, 0x0a, 0x5e, 0xf2, 0xa2, 0x29, 0xda, 0xe6, 0x25, 0x69, 0xa4, 0x28, 0x0a, 0x6f,
	0xcb, 0xf0, 0x8f, 0xad, 0xa6, 0xa7, 0x94, 0x57, 0x5a, 0xc0, 0x15, 0x7a, 0xbe, 0x29, 0x90, 0x34,
	0x07, 0x62, 0xce, 0x91, 0xa4, 0x91, 0x8c, 0x2c, 0x68, 0xde, 0x80, 0xd7, 0x1f, 0x39, 0xc1, 0x76,
	0x1c, 0x6a, 0xe3, 0x7f, 0xdd, 0x1d, 0x9d, 0x64, 0x5c, 0x7d, 0x6e, 0x66, 0x61, 0x2a, 0x8a, 0x28,
	0x15, 0xb2, 0x10, 0xd2, 0x7e, 0x4e, 0x25, 0x9b, 0x47, 0xea, 0xba, 0x02, 0x19, 0x9e, 0x43, 0x9a,
	0x3c, 0xb3, 0x4d, 0x2f, 0x69, 0x0e, 0xc6, 0x2d, 0x79, 0x25, 0xd9, 0x47, 0xdd, 0x10, 0x8f, 0xd1,
	0x93, 0x8d, 0x3f, 0xff, 0xde, 0xe4, 0x9a, 0x1d, 0xb1, 0x15, 0xbb, 0x57, 0x7d, 0x75, 0x50, 0xb0,
	0xa9, 0x69, 0xdd, 0x60, 0x90, 0xd6, 0x40, 0x25, 0x10, 0x9a, 0xa6, 0x90, 0x43, 0x4d, 0x8d, 0x56,
	0x53, 0x05, 0xc6, 0xf1, 0xff, 0x5f, 0xf8, 0xa5, 0xed, 0x6f, 0xfe, 0xe4, 0xb9, 0xed, 0x3e, 0xe9,
	0x34, 0x4f, 0xa8, 0x82, 0x78, 0x7a, 0xb3, 0xf2, 0x9d, 0xdb, 0x95, 0xef, 0xfc, 0x5e, 0xf9, 0xce,
	0xb7, 0xb5, 0xdf, 0xbb, 0x5d, 0xfb, 0xbd, 0x9f, 0x6b, 0xbf, 0xf7, 0x29, 0xea, 0xcc, 0xa9, 0x20,
	0xcb, 0xae, 0xbf, 0x2c, 0x22, 0x6d, 0x30, 0xe4, 0x1c, 0xea, 0x68, 0xf1, 0x26, 0x5a, 0x6e, 0x5e,
	0x45, 0x3b, 0x74, 0x36, 0x30, 0xcf, 0xe0, 0xf5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x63,
	0xd6, 0x24, 0x73, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastBidId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastBidId))
		i--
		dAtA[i] = 0x30
	}
	if m.LastAuctionId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastAuctionId))
		i--
		dAtA[i] = 0x28
	}
	if len(m.TokenPrices) > 0 {
		for iNdEx := len(m.TokenPrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokenPrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Bids) > 0 {
		for iNdEx := len(m.Bids) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Bids[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Auctions) > 0 {
		for iNdEx := len(m.Auctions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Auctions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.AuctionPriceDecreaseAccelerationRate.Size()
		i -= size
		if _, err := m.AuctionPriceDecreaseAccelerationRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.AuctionMaxBlockAge != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.AuctionMaxBlockAge))
		i--
		dAtA[i] = 0x20
	}
	{
		size := m.MinimumSaleTokensUsdValue.Size()
		i -= size
		if _, err := m.MinimumSaleTokensUsdValue.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.MinimumBidInUsomm != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MinimumBidInUsomm))
		i--
		dAtA[i] = 0x10
	}
	if m.PriceMaxBlockAge != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.PriceMaxBlockAge))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Auctions) > 0 {
		for _, e := range m.Auctions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Bids) > 0 {
		for _, e := range m.Bids {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TokenPrices) > 0 {
		for _, e := range m.TokenPrices {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.LastAuctionId != 0 {
		n += 1 + sovGenesis(uint64(m.LastAuctionId))
	}
	if m.LastBidId != 0 {
		n += 1 + sovGenesis(uint64(m.LastBidId))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PriceMaxBlockAge != 0 {
		n += 1 + sovGenesis(uint64(m.PriceMaxBlockAge))
	}
	if m.MinimumBidInUsomm != 0 {
		n += 1 + sovGenesis(uint64(m.MinimumBidInUsomm))
	}
	l = m.MinimumSaleTokensUsdValue.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.AuctionMaxBlockAge != 0 {
		n += 1 + sovGenesis(uint64(m.AuctionMaxBlockAge))
	}
	l = m.AuctionPriceDecreaseAccelerationRate.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Auctions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Auctions = append(m.Auctions, &Auction{})
			if err := m.Auctions[len(m.Auctions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bids", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bids = append(m.Bids, &Bid{})
			if err := m.Bids[len(m.Bids)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenPrices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenPrices = append(m.TokenPrices, &TokenPrice{})
			if err := m.TokenPrices[len(m.TokenPrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastAuctionId", wireType)
			}
			m.LastAuctionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastAuctionId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBidId", wireType)
			}
			m.LastBidId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastBidId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceMaxBlockAge", wireType)
			}
			m.PriceMaxBlockAge = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PriceMaxBlockAge |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinimumBidInUsomm", wireType)
			}
			m.MinimumBidInUsomm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinimumBidInUsomm |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinimumSaleTokensUsdValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinimumSaleTokensUsdValue.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionMaxBlockAge", wireType)
			}
			m.AuctionMaxBlockAge = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionMaxBlockAge |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionPriceDecreaseAccelerationRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AuctionPriceDecreaseAccelerationRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
