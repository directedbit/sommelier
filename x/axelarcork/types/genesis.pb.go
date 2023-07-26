// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelarcork/v1/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState - all cork state that must be provided at genesis
type GenesisState struct {
	Params              *Params               `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	ChainConfigurations ChainConfigurations   `protobuf:"bytes,2,opt,name=chain_configurations,json=chainConfigurations,proto3" json:"chain_configurations"`
	CellarIds           []*CellarIDSet        `protobuf:"bytes,3,rep,name=cellar_ids,json=cellarIds,proto3" json:"cellar_ids,omitempty"`
	ScheduledCorks      *ScheduledAxelarCorks `protobuf:"bytes,4,opt,name=scheduled_corks,json=scheduledCorks,proto3" json:"scheduled_corks,omitempty"`
	CorkResults         *AxelarCorkResults    `protobuf:"bytes,5,opt,name=cork_results,json=corkResults,proto3" json:"cork_results,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d754a1cfeef5947, []int{0}
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

func (m *GenesisState) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *GenesisState) GetChainConfigurations() ChainConfigurations {
	if m != nil {
		return m.ChainConfigurations
	}
	return ChainConfigurations{}
}

func (m *GenesisState) GetCellarIds() []*CellarIDSet {
	if m != nil {
		return m.CellarIds
	}
	return nil
}

func (m *GenesisState) GetScheduledCorks() *ScheduledAxelarCorks {
	if m != nil {
		return m.ScheduledCorks
	}
	return nil
}

func (m *GenesisState) GetCorkResults() *AxelarCorkResults {
	if m != nil {
		return m.CorkResults
	}
	return nil
}

type Params struct {
	Enabled         bool   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty" yaml:"enabled"`
	IbcChannel      string `protobuf:"bytes,2,opt,name=ibc_channel,json=ibcChannel,proto3" json:"ibc_channel,omitempty" yaml:"ibc_channel"`
	IbcPort         string `protobuf:"bytes,3,opt,name=ibc_port,json=ibcPort,proto3" json:"ibc_port,omitempty" yaml:"ibc_port"`
	GmpAccount      string `protobuf:"bytes,4,opt,name=gmp_account,json=gmpAccount,proto3" json:"gmp_account,omitempty" yaml:"gmp_account"`
	ExecutorAccount string `protobuf:"bytes,5,opt,name=executor_account,json=executorAccount,proto3" json:"executor_account,omitempty" yaml:"executor_account"`
	TimeoutDuration uint64 `protobuf:"varint,6,opt,name=timeout_duration,json=timeoutDuration,proto3" json:"timeout_duration,omitempty" yaml:"timeout_duration"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d754a1cfeef5947, []int{1}
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

func (m *Params) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Params) GetIbcChannel() string {
	if m != nil {
		return m.IbcChannel
	}
	return ""
}

func (m *Params) GetIbcPort() string {
	if m != nil {
		return m.IbcPort
	}
	return ""
}

func (m *Params) GetGmpAccount() string {
	if m != nil {
		return m.GmpAccount
	}
	return ""
}

func (m *Params) GetExecutorAccount() string {
	if m != nil {
		return m.ExecutorAccount
	}
	return ""
}

func (m *Params) GetTimeoutDuration() uint64 {
	if m != nil {
		return m.TimeoutDuration
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "axelarcork.v1.GenesisState")
	proto.RegisterType((*Params)(nil), "axelarcork.v1.Params")
}

func init() { proto.RegisterFile("axelarcork/v1/genesis.proto", fileDescriptor_8d754a1cfeef5947) }

var fileDescriptor_8d754a1cfeef5947 = []byte{
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x9b, 0xfd, 0xe9, 0x36, 0x77, 0xac, 0xc8, 0x1b, 0x10, 0x75, 0x52, 0x5a, 0x85, 0x4b,
	0x0f, 0x90, 0x68, 0x45, 0x02, 0xc1, 0x6d, 0xed, 0x04, 0x9a, 0xb4, 0xc3, 0xe4, 0xde, 0xe0, 0x10,
	0x39, 0x8e, 0x49, 0x0d, 0x49, 0x1c, 0xd9, 0x4e, 0xd5, 0x7e, 0x0b, 0x3e, 0xd6, 0x8e, 0x3b, 0x72,
	0xaa, 0x50, 0x7b, 0xe1, 0xdc, 0x4f, 0x80, 0xe2, 0xa4, 0x6b, 0x1a, 0x6e, 0x7e, 0xdf, 0xe7, 0xf7,
	0x3c, 0x6f, 0xf2, 0xca, 0x06, 0x97, 0x78, 0x46, 0x23, 0x2c, 0x08, 0x17, 0x3f, 0xdd, 0xe9, 0x95,
	0x1b, 0xd2, 0x84, 0x4a, 0x26, 0x9d, 0x54, 0x70, 0xc5, 0xe1, 0xb3, 0xad, 0xe8, 0x4c, 0xaf, 0x3a,
	0xd6, 0x2e, 0x5b, 0x11, 0x35, 0xde, 0xb9, 0x08, 0x79, 0xc8, 0xf5, 0xd1, 0xcd, 0x4f, 0x45, 0xd7,
	0xfe, 0xbb, 0x07, 0x4e, 0xbf, 0x14, 0xb1, 0x63, 0x85, 0x15, 0x85, 0x6f, 0x41, 0x33, 0xc5, 0x02,
	0xc7, 0xd2, 0x34, 0x7a, 0x46, 0xbf, 0x35, 0x78, 0xe1, 0xec, 0x8c, 0x71, 0xee, 0xb5, 0x88, 0x4a,
	0x08, 0x7e, 0x03, 0x17, 0x64, 0x82, 0x59, 0xe2, 0x11, 0x9e, 0x7c, 0x67, 0x61, 0x26, 0xb0, 0x62,
	0x3c, 0x91, 0xe6, 0x9e, 0x36, 0xdb, 0x35, 0xf3, 0x28, 0x47, 0x47, 0x3b, 0xe4, 0xf0, 0xe0, 0x61,
	0xd1, 0x6d, 0xa0, 0x73, 0xf2, 0xbf, 0x04, 0x3f, 0x02, 0x40, 0x68, 0x14, 0x61, 0xe1, 0xb1, 0x40,
	0x9a, 0xfb, 0xbd, 0xfd, 0x7e, 0x6b, 0xd0, 0xa9, 0x47, 0x6a, 0xe0, 0xf6, 0x66, 0x4c, 0x15, 0x3a,
	0x29, 0xe8, 0xdb, 0x40, 0xc2, 0x3b, 0xd0, 0x96, 0x64, 0x42, 0x83, 0x2c, 0xa2, 0x81, 0x97, 0xb3,
	0xd2, 0x3c, 0xd0, 0x9f, 0xf4, 0xba, 0xe6, 0x1f, 0x6f, 0xa8, 0x6b, 0xdd, 0x1e, 0xe5, 0x28, 0x3a,
	0x7b, 0xf2, 0xea, 0x1a, 0x8e, 0xc0, 0x69, 0xce, 0x7b, 0x82, 0xca, 0x2c, 0x52, 0xd2, 0x3c, 0xd4,
	0x51, 0xbd, 0x5a, 0xd4, 0x36, 0x01, 0x15, 0x1c, 0x6a, 0x91, 0x6d, 0x91, 0xaf, 0xba, 0x59, 0x6c,
	0x0f, 0xbe, 0x01, 0x47, 0x34, 0xc1, 0x7e, 0x44, 0x03, 0xbd, 0xe5, 0xe3, 0x21, 0x5c, 0x2f, 0xba,
	0x67, 0x73, 0x1c, 0x47, 0x9f, 0xec, 0x52, 0xb0, 0xd1, 0x06, 0x81, 0x1f, 0x40, 0x8b, 0xf9, 0xc4,
	0x23, 0x13, 0x9c, 0x24, 0x34, 0xd2, 0xab, 0x3d, 0x19, 0xbe, 0x5c, 0x2f, 0xba, 0xb0, 0x70, 0x54,
	0x44, 0x1b, 0x01, 0xe6, 0x93, 0x51, 0x51, 0x40, 0x07, 0x1c, 0xe7, 0x5a, 0xca, 0x85, 0x32, 0xf7,
	0xb5, 0xeb, 0x7c, 0xbd, 0xe8, 0xb6, 0xb7, 0xae, 0x5c, 0xb1, 0xd1, 0x11, 0xf3, 0xc9, 0x3d, 0x17,
	0x2a, 0x1f, 0x14, 0xc6, 0xa9, 0x87, 0x09, 0xe1, 0x59, 0xa2, 0xf4, 0xc2, 0x76, 0x06, 0x55, 0x44,
	0x1b, 0x81, 0x30, 0x4e, 0xaf, 0x8b, 0x02, 0x7e, 0x06, 0xcf, 0xe9, 0x8c, 0x92, 0x4c, 0x71, 0xf1,
	0xe4, 0x3e, 0xd4, 0xee, 0xcb, 0xf5, 0xa2, 0xfb, 0xaa, 0xfc, 0xb1, 0x1a, 0x61, 0xa3, 0xf6, 0xa6,
	0x55, 0xc9, 0x51, 0x2c, 0xa6, 0x3c, 0x53, 0x5e, 0x50, 0xde, 0x02, 0xb3, 0xd9, 0x33, 0xfa, 0x07,
	0xd5, 0x9c, 0x3a, 0x61, 0xa3, 0x76, 0xd9, 0xba, 0x29, 0x3b, 0xc3, 0xbb, 0x87, 0xa5, 0x65, 0x3c,
	0x2e, 0x2d, 0xe3, 0xcf, 0xd2, 0x32, 0x7e, 0xad, 0xac, 0xc6, 0xe3, 0xca, 0x6a, 0xfc, 0x5e, 0x59,
	0x8d, 0xaf, 0x83, 0x90, 0xa9, 0x49, 0xe6, 0x3b, 0x84, 0xc7, 0x6e, 0x4a, 0xc3, 0x70, 0xfe, 0x63,
	0xea, 0x4a, 0x1e, 0xc7, 0x34, 0x62, 0x54, 0xb8, 0xd3, 0xf7, 0xee, 0xac, 0xf2, 0x72, 0x5c, 0x35,
	0x4f, 0xa9, 0xf4, 0x9b, 0xfa, 0xa9, 0xbc, 0xfb, 0x17, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x72, 0x81,
	0x59, 0x8e, 0x03, 0x00, 0x00,
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
	if m.CorkResults != nil {
		{
			size, err := m.CorkResults.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.ScheduledCorks != nil {
		{
			size, err := m.ScheduledCorks.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.CellarIds) > 0 {
		for iNdEx := len(m.CellarIds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CellarIds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	{
		size, err := m.ChainConfigurations.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Params != nil {
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
	}
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
	if m.TimeoutDuration != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.TimeoutDuration))
		i--
		dAtA[i] = 0x30
	}
	if len(m.ExecutorAccount) > 0 {
		i -= len(m.ExecutorAccount)
		copy(dAtA[i:], m.ExecutorAccount)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ExecutorAccount)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.GmpAccount) > 0 {
		i -= len(m.GmpAccount)
		copy(dAtA[i:], m.GmpAccount)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.GmpAccount)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.IbcPort) > 0 {
		i -= len(m.IbcPort)
		copy(dAtA[i:], m.IbcPort)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.IbcPort)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.IbcChannel) > 0 {
		i -= len(m.IbcChannel)
		copy(dAtA[i:], m.IbcChannel)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.IbcChannel)))
		i--
		dAtA[i] = 0x12
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
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
	if m.Params != nil {
		l = m.Params.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.ChainConfigurations.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.CellarIds) > 0 {
		for _, e := range m.CellarIds {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.ScheduledCorks != nil {
		l = m.ScheduledCorks.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.CorkResults != nil {
		l = m.CorkResults.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Enabled {
		n += 2
	}
	l = len(m.IbcChannel)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.IbcPort)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.GmpAccount)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ExecutorAccount)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.TimeoutDuration != 0 {
		n += 1 + sovGenesis(uint64(m.TimeoutDuration))
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
			if m.Params == nil {
				m.Params = &Params{}
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainConfigurations", wireType)
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
			if err := m.ChainConfigurations.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CellarIds", wireType)
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
			m.CellarIds = append(m.CellarIds, &CellarIDSet{})
			if err := m.CellarIds[len(m.CellarIds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScheduledCorks", wireType)
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
			if m.ScheduledCorks == nil {
				m.ScheduledCorks = &ScheduledAxelarCorks{}
			}
			if err := m.ScheduledCorks.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CorkResults", wireType)
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
			if m.CorkResults == nil {
				m.CorkResults = &AxelarCorkResults{}
			}
			if err := m.CorkResults.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
			m.Enabled = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcChannel", wireType)
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
			m.IbcChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcPort", wireType)
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
			m.IbcPort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GmpAccount", wireType)
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
			m.GmpAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutorAccount", wireType)
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
			m.ExecutorAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutDuration", wireType)
			}
			m.TimeoutDuration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutDuration |= uint64(b&0x7F) << shift
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
