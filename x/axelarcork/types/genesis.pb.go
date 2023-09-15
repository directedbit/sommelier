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
	Params                   *Params                    `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	ChainConfigurations      ChainConfigurations        `protobuf:"bytes,2,opt,name=chain_configurations,json=chainConfigurations,proto3" json:"chain_configurations"`
	CellarIds                []*CellarIDSet             `protobuf:"bytes,3,rep,name=cellar_ids,json=cellarIds,proto3" json:"cellar_ids,omitempty"`
	ScheduledCorks           *ScheduledAxelarCorks      `protobuf:"bytes,4,opt,name=scheduled_corks,json=scheduledCorks,proto3" json:"scheduled_corks,omitempty"`
	CorkResults              *AxelarCorkResults         `protobuf:"bytes,5,opt,name=cork_results,json=corkResults,proto3" json:"cork_results,omitempty"`
	AxelarContractCallNonces []*AxelarContractCallNonce `protobuf:"bytes,6,rep,name=axelar_contract_call_nonces,json=axelarContractCallNonces,proto3" json:"axelar_contract_call_nonces,omitempty"`
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

func (m *GenesisState) GetAxelarContractCallNonces() []*AxelarContractCallNonce {
	if m != nil {
		return m.AxelarContractCallNonces
	}
	return nil
}

type Params struct {
	Enabled           bool   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty" yaml:"enabled"`
	IbcChannel        string `protobuf:"bytes,2,opt,name=ibc_channel,json=ibcChannel,proto3" json:"ibc_channel,omitempty" yaml:"ibc_channel"`
	IbcPort           string `protobuf:"bytes,3,opt,name=ibc_port,json=ibcPort,proto3" json:"ibc_port,omitempty" yaml:"ibc_port"`
	GmpAccount        string `protobuf:"bytes,4,opt,name=gmp_account,json=gmpAccount,proto3" json:"gmp_account,omitempty" yaml:"gmp_account"`
	ExecutorAccount   string `protobuf:"bytes,5,opt,name=executor_account,json=executorAccount,proto3" json:"executor_account,omitempty" yaml:"executor_account"`
	TimeoutDuration   uint64 `protobuf:"varint,6,opt,name=timeout_duration,json=timeoutDuration,proto3" json:"timeout_duration,omitempty" yaml:"timeout_duration"`
	CorkTimeoutBlocks uint64 `protobuf:"varint,7,opt,name=cork_timeout_blocks,json=corkTimeoutBlocks,proto3" json:"cork_timeout_blocks,omitempty" yaml:"cork_timeout_blocks"`
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

func (m *Params) GetCorkTimeoutBlocks() uint64 {
	if m != nil {
		return m.CorkTimeoutBlocks
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "axelarcork.v1.GenesisState")
	proto.RegisterType((*Params)(nil), "axelarcork.v1.Params")
}

func init() { proto.RegisterFile("axelarcork/v1/genesis.proto", fileDescriptor_8d754a1cfeef5947) }

var fileDescriptor_8d754a1cfeef5947 = []byte{
	// 592 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xd1, 0x6e, 0xd3, 0x3c,
	0x1c, 0xc5, 0xdb, 0xaf, 0x5b, 0xb7, 0xb9, 0xfb, 0x56, 0xf0, 0x06, 0x44, 0x9d, 0x94, 0x56, 0x41,
	0x42, 0xbb, 0x80, 0x44, 0x1b, 0x12, 0x08, 0xee, 0xd6, 0x4c, 0xa0, 0x49, 0xd3, 0x34, 0x79, 0x5c,
	0xc1, 0x45, 0xe4, 0x38, 0x26, 0x0d, 0x73, 0xe2, 0xc8, 0x76, 0xaa, 0xed, 0x2d, 0x78, 0x1c, 0x1e,
	0x61, 0x77, 0xec, 0x92, 0xab, 0x08, 0x6d, 0x6f, 0xd0, 0x27, 0x40, 0x71, 0xd2, 0x35, 0x0d, 0xe3,
	0x2e, 0xff, 0x73, 0x7e, 0xe7, 0xb8, 0xe9, 0x3f, 0x06, 0xbb, 0xf8, 0x92, 0x32, 0x2c, 0x08, 0x17,
	0x17, 0xce, 0x74, 0xdf, 0x09, 0x69, 0x42, 0x65, 0x24, 0xed, 0x54, 0x70, 0xc5, 0xe1, 0xff, 0x0b,
	0xd3, 0x9e, 0xee, 0x0f, 0xcc, 0x65, 0xb6, 0x66, 0x6a, 0x7c, 0xb0, 0x13, 0xf2, 0x90, 0xeb, 0x47,
	0xa7, 0x78, 0x2a, 0x55, 0xeb, 0x67, 0x07, 0x6c, 0x7e, 0x2c, 0x6b, 0xcf, 0x15, 0x56, 0x14, 0xbe,
	0x02, 0xdd, 0x14, 0x0b, 0x1c, 0x4b, 0xa3, 0x3d, 0x6a, 0xef, 0xf5, 0x0e, 0x9e, 0xd8, 0x4b, 0xc7,
	0xd8, 0x67, 0xda, 0x44, 0x15, 0x04, 0xbf, 0x80, 0x1d, 0x32, 0xc1, 0x51, 0xe2, 0x11, 0x9e, 0x7c,
	0x8d, 0xc2, 0x4c, 0x60, 0x15, 0xf1, 0x44, 0x1a, 0xff, 0xe9, 0xb0, 0xd5, 0x08, 0xbb, 0x05, 0xea,
	0x2e, 0x91, 0xe3, 0x95, 0xeb, 0x7c, 0xd8, 0x42, 0xdb, 0xe4, 0x6f, 0x0b, 0xbe, 0x03, 0x80, 0x50,
	0xc6, 0xb0, 0xf0, 0xa2, 0x40, 0x1a, 0x9d, 0x51, 0x67, 0xaf, 0x77, 0x30, 0x68, 0x56, 0x6a, 0xe0,
	0xf8, 0xe8, 0x9c, 0x2a, 0xb4, 0x51, 0xd2, 0xc7, 0x81, 0x84, 0x27, 0xa0, 0x2f, 0xc9, 0x84, 0x06,
	0x19, 0xa3, 0x81, 0x57, 0xb0, 0xd2, 0x58, 0xd1, 0x3f, 0xe9, 0x79, 0x23, 0x7f, 0x3e, 0xa7, 0x0e,
	0xb5, 0xec, 0x16, 0x28, 0xda, 0xba, 0xcf, 0xea, 0x19, 0xba, 0x60, 0xb3, 0xe0, 0x3d, 0x41, 0x65,
	0xc6, 0x94, 0x34, 0x56, 0x75, 0xd5, 0xa8, 0x51, 0xb5, 0x68, 0x40, 0x25, 0x87, 0x7a, 0x64, 0x31,
	0x40, 0x3a, 0x5f, 0x67, 0xf1, 0x5f, 0x29, 0x81, 0x89, 0xf2, 0x08, 0x66, 0xcc, 0x4b, 0x78, 0x42,
	0xa8, 0x34, 0xba, 0xfa, 0xf5, 0x5e, 0xfc, 0xa3, 0xb3, 0x0c, 0xb8, 0x98, 0xb1, 0xd3, 0x02, 0x47,
	0x06, 0x7e, 0xd8, 0x90, 0xd6, 0x8f, 0x0e, 0xe8, 0x96, 0x4b, 0x82, 0x2f, 0xc1, 0x1a, 0x4d, 0xb0,
	0xcf, 0x68, 0xa0, 0x97, 0xb9, 0x3e, 0x86, 0xb3, 0x7c, 0xb8, 0x75, 0x85, 0x63, 0xf6, 0xde, 0xaa,
	0x0c, 0x0b, 0xcd, 0x11, 0xf8, 0x16, 0xf4, 0x22, 0x9f, 0x78, 0x64, 0x82, 0x93, 0x84, 0x32, 0xbd,
	0xc1, 0x8d, 0xf1, 0xd3, 0x59, 0x3e, 0x84, 0x65, 0xa2, 0x66, 0x5a, 0x08, 0x44, 0x3e, 0x71, 0xcb,
	0x01, 0xda, 0x60, 0xbd, 0xf0, 0x52, 0x2e, 0x94, 0xd1, 0xd1, 0xa9, 0xed, 0x59, 0x3e, 0xec, 0x2f,
	0x52, 0x85, 0x63, 0xa1, 0xb5, 0xc8, 0x27, 0x67, 0x5c, 0xa8, 0xe2, 0xa0, 0x30, 0x4e, 0x3d, 0x4c,
	0x08, 0xcf, 0x12, 0xa5, 0xf7, 0xb2, 0x74, 0x50, 0xcd, 0xb4, 0x10, 0x08, 0xe3, 0xf4, 0xb0, 0x1c,
	0xe0, 0x07, 0xf0, 0x88, 0x5e, 0x52, 0x92, 0x29, 0x2e, 0xee, 0xd3, 0xab, 0x3a, 0xbd, 0x3b, 0xcb,
	0x87, 0xcf, 0xaa, 0x17, 0x6b, 0x10, 0x16, 0xea, 0xcf, 0xa5, 0x5a, 0x8f, 0x8a, 0x62, 0xca, 0x33,
	0xe5, 0x05, 0xd5, 0xc7, 0x66, 0x74, 0x47, 0xed, 0xbd, 0x95, 0x7a, 0x4f, 0x93, 0xb0, 0x50, 0xbf,
	0x92, 0x8e, 0x2a, 0x05, 0x9e, 0x82, 0x6d, 0xfd, 0x59, 0xcc, 0x51, 0x9f, 0x71, 0x72, 0x21, 0x8d,
	0x35, 0x5d, 0x65, 0xce, 0xf2, 0xe1, 0xa0, 0xac, 0x7a, 0x00, 0xb2, 0xd0, 0xe3, 0x42, 0xfd, 0x54,
	0x8a, 0x63, 0xad, 0x8d, 0x4f, 0xae, 0x6f, 0xcd, 0xf6, 0xcd, 0xad, 0xd9, 0xfe, 0x7d, 0x6b, 0xb6,
	0xbf, 0xdf, 0x99, 0xad, 0x9b, 0x3b, 0xb3, 0xf5, 0xeb, 0xce, 0x6c, 0x7d, 0x3e, 0x08, 0x23, 0x35,
	0xc9, 0x7c, 0x9b, 0xf0, 0xd8, 0x49, 0x69, 0x18, 0x5e, 0x7d, 0x9b, 0x3a, 0x92, 0xc7, 0x31, 0x65,
	0x11, 0x15, 0xce, 0xf4, 0x8d, 0x73, 0x59, 0xbb, 0xf0, 0x8e, 0xba, 0x4a, 0xa9, 0xf4, 0xbb, 0xfa,
	0x86, 0xbf, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x16, 0xb3, 0xbe, 0x45, 0x04, 0x00, 0x00,
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
	if len(m.AxelarContractCallNonces) > 0 {
		for iNdEx := len(m.AxelarContractCallNonces) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AxelarContractCallNonces[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
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
	if m.CorkTimeoutBlocks != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.CorkTimeoutBlocks))
		i--
		dAtA[i] = 0x38
	}
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
	if len(m.AxelarContractCallNonces) > 0 {
		for _, e := range m.AxelarContractCallNonces {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
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
	if m.CorkTimeoutBlocks != 0 {
		n += 1 + sovGenesis(uint64(m.CorkTimeoutBlocks))
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
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AxelarContractCallNonces", wireType)
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
			m.AxelarContractCallNonces = append(m.AxelarContractCallNonces, &AxelarContractCallNonce{})
			if err := m.AxelarContractCallNonces[len(m.AxelarContractCallNonces)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CorkTimeoutBlocks", wireType)
			}
			m.CorkTimeoutBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CorkTimeoutBlocks |= uint64(b&0x7F) << shift
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
