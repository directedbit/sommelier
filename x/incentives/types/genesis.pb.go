// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: incentives/v1/genesis.proto

package types

import (
	fmt "fmt"
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

type GenesisState struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_179cfb82d3e2b395, []int{0}
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

// Params incentives parameters
type Params struct {
	// DistributionPerBlock defines the coin to be sent to the distribution module from the community pool every block
	DistributionPerBlock types.Coin `protobuf:"bytes,1,opt,name=distribution_per_block,json=distributionPerBlock,proto3" json:"distribution_per_block"`
	// IncentivesCutoffHeight defines the block height after which the incentives module will stop sending coins to the distribution module from
	// the community pool
	IncentivesCutoffHeight uint64 `protobuf:"varint,2,opt,name=incentives_cutoff_height,json=incentivesCutoffHeight,proto3" json:"incentives_cutoff_height,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_179cfb82d3e2b395, []int{1}
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

func (m *Params) GetDistributionPerBlock() types.Coin {
	if m != nil {
		return m.DistributionPerBlock
	}
	return types.Coin{}
}

func (m *Params) GetIncentivesCutoffHeight() uint64 {
	if m != nil {
		return m.IncentivesCutoffHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "incentives.v1.GenesisState")
	proto.RegisterType((*Params)(nil), "incentives.v1.Params")
}

func init() { proto.RegisterFile("incentives/v1/genesis.proto", fileDescriptor_179cfb82d3e2b395) }

var fileDescriptor_179cfb82d3e2b395 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0xd0, 0xb1, 0x4f, 0x3a, 0x31,
	0x14, 0x07, 0xf0, 0xeb, 0x2f, 0x84, 0xe1, 0x7e, 0xba, 0x5c, 0x90, 0x20, 0x26, 0x95, 0x30, 0x31,
	0xb5, 0x39, 0x70, 0x70, 0x86, 0x41, 0x07, 0x07, 0x82, 0x71, 0x71, 0xb9, 0xdc, 0xd5, 0x47, 0xa9,
	0x72, 0x7d, 0x97, 0xb6, 0x34, 0xf2, 0x5f, 0xb8, 0xfa, 0x1f, 0x31, 0x32, 0x3a, 0x19, 0x03, 0xff,
	0x88, 0xe1, 0xee, 0x12, 0x70, 0x6b, 0xf2, 0x79, 0xef, 0xdb, 0x6f, 0x5e, 0x78, 0xa5, 0xb4, 0x00,
	0xed, 0x94, 0x07, 0xcb, 0x7d, 0xcc, 0x25, 0x68, 0xb0, 0xca, 0xb2, 0xc2, 0xa0, 0xc3, 0xe8, 0xfc,
	0x88, 0xcc, 0xc7, 0xdd, 0x96, 0x44, 0x89, 0xa5, 0xf0, 0xc3, 0xab, 0x1a, 0xea, 0x52, 0x81, 0x36,
	0x47, 0xcb, 0xb3, 0xd4, 0x02, 0xf7, 0x71, 0x06, 0x2e, 0x8d, 0xb9, 0x40, 0xa5, 0x2b, 0xef, 0x4f,
	0xc2, 0xb3, 0xbb, 0x2a, 0xf5, 0xd1, 0xa5, 0x0e, 0xa2, 0x51, 0xd8, 0x2c, 0x52, 0x93, 0xe6, 0xb6,
	0x43, 0x7a, 0x64, 0xf0, 0x7f, 0x78, 0xc1, 0xfe, 0xfc, 0xc2, 0xa6, 0x25, 0x8e, 0x1b, 0x9b, 0xef,
	0xeb, 0x60, 0x56, 0x8f, 0xf6, 0x3f, 0x49, 0xd8, 0xac, 0x20, 0x7a, 0x0a, 0xdb, 0x2f, 0xca, 0x3a,
	0xa3, 0xb2, 0x95, 0x53, 0xa8, 0x93, 0x02, 0x4c, 0x92, 0x2d, 0x51, 0xbc, 0xd5, 0x79, 0x97, 0xac,
	0x2a, 0xc4, 0x0e, 0x85, 0x58, 0x5d, 0x88, 0x4d, 0x50, 0xe9, 0x3a, 0xb3, 0x75, 0xba, 0x3e, 0x05,
	0x33, 0x3e, 0x2c, 0x47, 0xb7, 0x61, 0xe7, 0xd8, 0x23, 0x11, 0x2b, 0x87, 0xf3, 0x79, 0xb2, 0x00,
	0x25, 0x17, 0xae, 0xf3, 0xaf, 0x47, 0x06, 0x8d, 0x59, 0xfb, 0xe8, 0x93, 0x92, 0xef, 0x4b, 0x1d,
	0x3f, 0x6c, 0x76, 0x94, 0x6c, 0x77, 0x94, 0xfc, 0xec, 0x28, 0xf9, 0xd8, 0xd3, 0x60, 0xbb, 0xa7,
	0xc1, 0xd7, 0x9e, 0x06, 0xcf, 0x43, 0xa9, 0xdc, 0x62, 0x95, 0x31, 0x81, 0x39, 0x2f, 0x40, 0xca,
	0xf5, 0xab, 0xe7, 0x16, 0xf3, 0x1c, 0x96, 0x0a, 0x0c, 0xf7, 0x37, 0xfc, 0x9d, 0x9f, 0x9c, 0xdf,
	0xad, 0x0b, 0xb0, 0x59, 0xb3, 0xbc, 0xda, 0xe8, 0x37, 0x00, 0x00, 0xff, 0xff, 0x86, 0x6f, 0x5d,
	0x07, 0x99, 0x01, 0x00, 0x00,
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
	if m.IncentivesCutoffHeight != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.IncentivesCutoffHeight))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.DistributionPerBlock.MarshalToSizedBuffer(dAtA[:i])
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
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.DistributionPerBlock.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.IncentivesCutoffHeight != 0 {
		n += 1 + sovGenesis(uint64(m.IncentivesCutoffHeight))
	}
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionPerBlock", wireType)
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
			if err := m.DistributionPerBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncentivesCutoffHeight", wireType)
			}
			m.IncentivesCutoffHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IncentivesCutoffHeight |= uint64(b&0x7F) << shift
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
