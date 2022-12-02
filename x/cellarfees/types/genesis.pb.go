// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cellarfees/v1/genesis.proto

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

// GenesisState defines the cellarfees module's genesis state.
type GenesisState struct {
	Params               Params                                 `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	FeeAccrualCounters   FeeAccrualCounters                     `protobuf:"bytes,2,opt,name=fee_accrual_counters,json=feeAccrualCounters,proto3" json:"fee_accrual_counters"`
	LastRewardSupplyPeak github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=last_reward_supply_peak,json=lastRewardSupplyPeak,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"last_reward_supply_peak"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_856aa03b4cb6eca9, []int{0}
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

func (m *GenesisState) GetFeeAccrualCounters() FeeAccrualCounters {
	if m != nil {
		return m.FeeAccrualCounters
	}
	return FeeAccrualCounters{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "cellarfees.v1.GenesisState")
}

func init() { proto.RegisterFile("cellarfees/v1/genesis.proto", fileDescriptor_856aa03b4cb6eca9) }

var fileDescriptor_856aa03b4cb6eca9 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0xd0, 0xc1, 0x4e, 0x2a, 0x31,
	0x14, 0x06, 0xe0, 0x19, 0xee, 0x0d, 0x89, 0xa3, 0x6e, 0x26, 0x18, 0xc9, 0x98, 0x14, 0x74, 0x61,
	0xd8, 0xd8, 0x06, 0xf0, 0x05, 0xc4, 0x44, 0x63, 0xe2, 0x82, 0xc0, 0x4a, 0x37, 0x93, 0x52, 0x0e,
	0x15, 0xe9, 0xd0, 0xa6, 0xed, 0x8c, 0xf2, 0x00, 0xee, 0x7d, 0x2c, 0x96, 0x2c, 0x8d, 0x0b, 0x62,
	0xe0, 0x45, 0x0c, 0x9d, 0x49, 0x04, 0x5d, 0xb5, 0xc9, 0xff, 0xf7, 0x3b, 0xcd, 0x09, 0x4e, 0x18,
	0x08, 0x41, 0xf5, 0x08, 0xc0, 0x90, 0xac, 0x49, 0x38, 0x4c, 0xc1, 0x8c, 0x0d, 0x56, 0x5a, 0x5a,
	0x19, 0x1e, 0xfe, 0x84, 0x38, 0x6b, 0x46, 0x15, 0x2e, 0xb9, 0x74, 0x09, 0xd9, 0xdc, 0xf2, 0x52,
	0x14, 0xed, 0x0a, 0x8a, 0x6a, 0x9a, 0x14, 0x40, 0x84, 0x76, 0xb3, 0x2d, 0xce, 0xe5, 0x67, 0x6f,
	0xa5, 0xe0, 0xe0, 0x36, 0x1f, 0xd9, 0xb7, 0xd4, 0x42, 0xd8, 0x0e, 0xca, 0x39, 0x50, 0xf5, 0xeb,
	0x7e, 0x63, 0xbf, 0x75, 0x84, 0x77, 0xbe, 0x80, 0xbb, 0x2e, 0xec, 0xfc, 0x9f, 0x2f, 0x6b, 0x5e,
	0xaf, 0xa8, 0x86, 0x0f, 0x41, 0x65, 0x04, 0x10, 0x53, 0xc6, 0x74, 0x4a, 0x45, 0xcc, 0x64, 0x3a,
	0xb5, 0xa0, 0x4d, 0xb5, 0xe4, 0x88, 0xd3, 0x5f, 0xc4, 0x0d, 0xc0, 0x55, 0xde, 0xbc, 0x2e, 0x8a,
	0x05, 0x17, 0x8e, 0xfe, 0x24, 0x21, 0x04, 0xc7, 0x82, 0x1a, 0x1b, 0x6b, 0x78, 0xa1, 0x7a, 0x18,
	0x9b, 0x54, 0x29, 0x31, 0x8b, 0x15, 0xd0, 0x49, 0xf5, 0x5f, 0xdd, 0x6f, 0xec, 0x75, 0xf0, 0xe6,
	0xe9, 0xe7, 0xb2, 0x76, 0xce, 0xc7, 0xf6, 0x29, 0x1d, 0x60, 0x26, 0x13, 0xc2, 0xa4, 0x49, 0xa4,
	0x29, 0x8e, 0x0b, 0x33, 0x9c, 0x10, 0x3b, 0x53, 0x60, 0xf0, 0xdd, 0xd4, 0xf6, 0x2a, 0x1b, 0xae,
	0xe7, 0xb4, 0xbe, 0xc3, 0xba, 0x40, 0x27, 0x9d, 0xfb, 0xf9, 0x0a, 0xf9, 0x8b, 0x15, 0xf2, 0xbf,
	0x56, 0xc8, 0x7f, 0x5f, 0x23, 0x6f, 0xb1, 0x46, 0xde, 0xc7, 0x1a, 0x79, 0x8f, 0xad, 0x2d, 0x57,
	0x01, 0xe7, 0xb3, 0xe7, 0x8c, 0x18, 0x99, 0x24, 0x20, 0xc6, 0xa0, 0x49, 0x76, 0x49, 0x5e, 0xb7,
	0xb6, 0x9a, 0xcf, 0x19, 0x94, 0xdd, 0x72, 0xdb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x94,
	0xd2, 0xf4, 0xdc, 0x01, 0x00, 0x00,
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
		size := m.LastRewardSupplyPeak.Size()
		i -= size
		if _, err := m.LastRewardSupplyPeak.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.FeeAccrualCounters.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
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
	l = m.FeeAccrualCounters.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.LastRewardSupplyPeak.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field FeeAccrualCounters", wireType)
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
			if err := m.FeeAccrualCounters.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastRewardSupplyPeak", wireType)
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
			if err := m.LastRewardSupplyPeak.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
