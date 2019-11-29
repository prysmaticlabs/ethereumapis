// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eth/v1alpha1/node.proto

package eth

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PeerDirection int32

const (
	PeerDirection_UNKNOWN  PeerDirection = 0
	PeerDirection_INBOUND  PeerDirection = 1
	PeerDirection_OUTBOUND PeerDirection = 2
)

var PeerDirection_name = map[int32]string{
	0: "UNKNOWN",
	1: "INBOUND",
	2: "OUTBOUND",
}

var PeerDirection_value = map[string]int32{
	"UNKNOWN":  0,
	"INBOUND":  1,
	"OUTBOUND": 2,
}

func (x PeerDirection) String() string {
	return proto.EnumName(PeerDirection_name, int32(x))
}

func (PeerDirection) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4d5c81e0a630f5ab, []int{0}
}

type SyncStatus struct {
	Syncing              bool     `protobuf:"varint,1,opt,name=syncing,proto3" json:"syncing,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncStatus) Reset()         { *m = SyncStatus{} }
func (m *SyncStatus) String() string { return proto.CompactTextString(m) }
func (*SyncStatus) ProtoMessage()    {}
func (*SyncStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d5c81e0a630f5ab, []int{0}
}

func (m *SyncStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncStatus.Unmarshal(m, b)
}
func (m *SyncStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncStatus.Marshal(b, m, deterministic)
}
func (m *SyncStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncStatus.Merge(m, src)
}
func (m *SyncStatus) XXX_Size() int {
	return xxx_messageInfo_SyncStatus.Size(m)
}
func (m *SyncStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncStatus.DiscardUnknown(m)
}

var xxx_messageInfo_SyncStatus proto.InternalMessageInfo

func (m *SyncStatus) GetSyncing() bool {
	if m != nil {
		return m.Syncing
	}
	return false
}

type Genesis struct {
	GenesisTime            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=genesis_time,json=genesisTime,proto3" json:"genesis_time,omitempty"`
	DepositContractAddress []byte               `protobuf:"bytes,2,opt,name=deposit_contract_address,json=depositContractAddress,proto3" json:"deposit_contract_address,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}             `json:"-"`
	XXX_unrecognized       []byte               `json:"-"`
	XXX_sizecache          int32                `json:"-"`
}

func (m *Genesis) Reset()         { *m = Genesis{} }
func (m *Genesis) String() string { return proto.CompactTextString(m) }
func (*Genesis) ProtoMessage()    {}
func (*Genesis) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d5c81e0a630f5ab, []int{1}
}

func (m *Genesis) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Genesis.Unmarshal(m, b)
}
func (m *Genesis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Genesis.Marshal(b, m, deterministic)
}
func (m *Genesis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Genesis.Merge(m, src)
}
func (m *Genesis) XXX_Size() int {
	return xxx_messageInfo_Genesis.Size(m)
}
func (m *Genesis) XXX_DiscardUnknown() {
	xxx_messageInfo_Genesis.DiscardUnknown(m)
}

var xxx_messageInfo_Genesis proto.InternalMessageInfo

func (m *Genesis) GetGenesisTime() *timestamp.Timestamp {
	if m != nil {
		return m.GenesisTime
	}
	return nil
}

func (m *Genesis) GetDepositContractAddress() []byte {
	if m != nil {
		return m.DepositContractAddress
	}
	return nil
}

type Version struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Metadata             string   `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d5c81e0a630f5ab, []int{2}
}

func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Version) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

type ImplementedServices struct {
	Services             []string `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImplementedServices) Reset()         { *m = ImplementedServices{} }
func (m *ImplementedServices) String() string { return proto.CompactTextString(m) }
func (*ImplementedServices) ProtoMessage()    {}
func (*ImplementedServices) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d5c81e0a630f5ab, []int{3}
}

func (m *ImplementedServices) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImplementedServices.Unmarshal(m, b)
}
func (m *ImplementedServices) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImplementedServices.Marshal(b, m, deterministic)
}
func (m *ImplementedServices) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImplementedServices.Merge(m, src)
}
func (m *ImplementedServices) XXX_Size() int {
	return xxx_messageInfo_ImplementedServices.Size(m)
}
func (m *ImplementedServices) XXX_DiscardUnknown() {
	xxx_messageInfo_ImplementedServices.DiscardUnknown(m)
}

var xxx_messageInfo_ImplementedServices proto.InternalMessageInfo

func (m *ImplementedServices) GetServices() []string {
	if m != nil {
		return m.Services
	}
	return nil
}

type Peers struct {
	Peers                []*Peer  `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Peers) Reset()         { *m = Peers{} }
func (m *Peers) String() string { return proto.CompactTextString(m) }
func (*Peers) ProtoMessage()    {}
func (*Peers) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d5c81e0a630f5ab, []int{4}
}

func (m *Peers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Peers.Unmarshal(m, b)
}
func (m *Peers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Peers.Marshal(b, m, deterministic)
}
func (m *Peers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peers.Merge(m, src)
}
func (m *Peers) XXX_Size() int {
	return xxx_messageInfo_Peers.Size(m)
}
func (m *Peers) XXX_DiscardUnknown() {
	xxx_messageInfo_Peers.DiscardUnknown(m)
}

var xxx_messageInfo_Peers proto.InternalMessageInfo

func (m *Peers) GetPeers() []*Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

type Peer struct {
	Address              string        `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Direction            PeerDirection `protobuf:"varint,2,opt,name=direction,proto3,enum=ethereum.eth.v1alpha1.PeerDirection" json:"direction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Peer) Reset()         { *m = Peer{} }
func (m *Peer) String() string { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()    {}
func (*Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d5c81e0a630f5ab, []int{5}
}

func (m *Peer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Peer.Unmarshal(m, b)
}
func (m *Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Peer.Marshal(b, m, deterministic)
}
func (m *Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peer.Merge(m, src)
}
func (m *Peer) XXX_Size() int {
	return xxx_messageInfo_Peer.Size(m)
}
func (m *Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_Peer proto.InternalMessageInfo

func (m *Peer) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Peer) GetDirection() PeerDirection {
	if m != nil {
		return m.Direction
	}
	return PeerDirection_UNKNOWN
}

func init() {
	proto.RegisterEnum("ethereum.eth.v1alpha1.PeerDirection", PeerDirection_name, PeerDirection_value)
	proto.RegisterType((*SyncStatus)(nil), "ethereum.eth.v1alpha1.SyncStatus")
	proto.RegisterType((*Genesis)(nil), "ethereum.eth.v1alpha1.Genesis")
	proto.RegisterType((*Version)(nil), "ethereum.eth.v1alpha1.Version")
	proto.RegisterType((*ImplementedServices)(nil), "ethereum.eth.v1alpha1.ImplementedServices")
	proto.RegisterType((*Peers)(nil), "ethereum.eth.v1alpha1.Peers")
	proto.RegisterType((*Peer)(nil), "ethereum.eth.v1alpha1.Peer")
}

func init() { proto.RegisterFile("eth/v1alpha1/node.proto", fileDescriptor_4d5c81e0a630f5ab) }

var fileDescriptor_4d5c81e0a630f5ab = []byte{
	// 606 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xd1, 0x6a, 0x13, 0x41,
	0x14, 0x86, 0xdd, 0xd8, 0x9a, 0x66, 0xd2, 0x4a, 0x19, 0x69, 0x1b, 0xb7, 0x55, 0xe3, 0x2a, 0x52,
	0x7a, 0xb1, 0x4b, 0x22, 0xa8, 0x28, 0x22, 0xc6, 0x86, 0x52, 0x94, 0xb4, 0x6c, 0x5b, 0x05, 0x29,
	0x94, 0xc9, 0xee, 0x31, 0x3b, 0x90, 0x9d, 0x59, 0x76, 0x4e, 0x0a, 0xb9, 0xed, 0x13, 0x08, 0xbe,
	0x85, 0xd7, 0x3e, 0x85, 0xb7, 0xbe, 0x82, 0x0f, 0x22, 0xb3, 0x3b, 0x63, 0x5a, 0x9b, 0x04, 0xbc,
	0x9b, 0x7f, 0xce, 0xf9, 0xcf, 0x77, 0x76, 0xe7, 0x27, 0x1b, 0x80, 0x49, 0x70, 0xde, 0x62, 0xc3,
	0x2c, 0x61, 0xad, 0x40, 0xc8, 0x18, 0xfc, 0x2c, 0x97, 0x28, 0xe9, 0x1a, 0x60, 0x02, 0x39, 0x8c,
	0x52, 0x1f, 0x30, 0xf1, 0x6d, 0x87, 0xbb, 0x35, 0x90, 0x72, 0x30, 0x84, 0x80, 0x65, 0x3c, 0x60,
	0x42, 0x48, 0x64, 0xc8, 0xa5, 0x50, 0xa5, 0xc9, 0xdd, 0x34, 0xd5, 0x42, 0xf5, 0x47, 0x5f, 0x02,
	0x48, 0x33, 0x1c, 0x9b, 0xe2, 0x83, 0x7f, 0x8b, 0xc8, 0x53, 0x50, 0xc8, 0xd2, 0xac, 0x6c, 0xf0,
	0x9e, 0x10, 0x72, 0x34, 0x16, 0xd1, 0x11, 0x32, 0x1c, 0x29, 0xda, 0x20, 0x55, 0x35, 0x16, 0x11,
	0x17, 0x83, 0x86, 0xd3, 0x74, 0xb6, 0x97, 0x42, 0x2b, 0xbd, 0x0b, 0x87, 0x54, 0xf7, 0x40, 0x80,
	0xe2, 0x8a, 0xbe, 0x26, 0xcb, 0x83, 0xf2, 0x78, 0xa6, 0xc7, 0x15, 0xad, 0xf5, 0xb6, 0xeb, 0x97,
	0x2c, 0xdf, 0xb2, 0xfc, 0x63, 0xcb, 0x0a, 0xeb, 0xa6, 0x5f, 0xdf, 0xd0, 0x17, 0xa4, 0x11, 0x43,
	0x26, 0x15, 0xc7, 0xb3, 0x48, 0x0a, 0xcc, 0x59, 0x84, 0x67, 0x2c, 0x8e, 0x73, 0x50, 0xaa, 0x51,
	0x69, 0x3a, 0xdb, 0xcb, 0xe1, 0xba, 0xa9, 0xbf, 0x33, 0xe5, 0xb7, 0x65, 0xd5, 0x7b, 0x43, 0xaa,
	0x1f, 0x21, 0x57, 0x5c, 0x0a, 0xbd, 0xe9, 0x79, 0x79, 0x2c, 0xf0, 0xb5, 0xd0, 0x4a, 0xea, 0x92,
	0xa5, 0x14, 0x90, 0xc5, 0x0c, 0x59, 0x31, 0xae, 0x16, 0xfe, 0xd5, 0x5e, 0x8b, 0xdc, 0xd9, 0x4f,
	0xb3, 0x21, 0xa4, 0x20, 0x10, 0xe2, 0x23, 0xc8, 0xcf, 0x79, 0x04, 0x4a, 0x5b, 0x94, 0x39, 0x37,
	0x9c, 0xe6, 0x4d, 0x6d, 0xb1, 0xda, 0x7b, 0x49, 0x16, 0x0f, 0x01, 0x72, 0x45, 0x5b, 0x64, 0x31,
	0xd3, 0x87, 0xa2, 0xa3, 0xde, 0xde, 0xf4, 0xa7, 0x3e, 0x96, 0xaf, 0x9b, 0xc3, 0xb2, 0xd3, 0x8b,
	0xc9, 0x82, 0x96, 0x7a, 0x59, 0xfb, 0x81, 0x66, 0x59, 0x23, 0x69, 0x87, 0xd4, 0x62, 0x9e, 0x43,
	0xa4, 0x1f, 0xb4, 0xd8, 0xf6, 0x76, 0xfb, 0xf1, 0x9c, 0xc1, 0xbb, 0xb6, 0x37, 0x9c, 0xd8, 0x76,
	0x9e, 0x93, 0x95, 0x2b, 0x35, 0x5a, 0x27, 0xd5, 0x93, 0xde, 0xfb, 0xde, 0xc1, 0xa7, 0xde, 0xea,
	0x0d, 0x2d, 0xf6, 0x7b, 0x9d, 0x83, 0x93, 0xde, 0xee, 0xaa, 0x43, 0x97, 0xc9, 0xd2, 0xc1, 0xc9,
	0x71, 0xa9, 0x2a, 0xed, 0x1f, 0x0b, 0x64, 0xa1, 0x27, 0x63, 0xa0, 0x82, 0xac, 0xec, 0x01, 0x5e,
	0xca, 0xc1, 0xfa, 0xb5, 0xb7, 0xec, 0xea, 0x50, 0xb9, 0x0f, 0x67, 0xec, 0x36, 0xb1, 0x7a, 0xde,
	0xc5, 0xaf, 0xdf, 0xdf, 0x2a, 0x5b, 0xd4, 0x0d, 0xae, 0xa5, 0x3c, 0x30, 0x61, 0xa2, 0x09, 0x21,
	0x7b, 0x80, 0x36, 0x4e, 0xb3, 0x60, 0xf7, 0x67, 0xc0, 0x8c, 0x6f, 0x2e, 0xc9, 0xe4, 0xcd, 0x90,
	0x6c, 0x68, 0xfe, 0x97, 0x64, 0x7c, 0x73, 0x49, 0x36, 0x76, 0x17, 0x0e, 0xd9, 0xf8, 0xc0, 0x15,
	0x4e, 0xcb, 0xd7, 0x2c, 0xee, 0xce, 0x0c, 0xee, 0x94, 0x19, 0xde, 0xa3, 0x62, 0x87, 0x7b, 0x74,
	0x73, 0xda, 0x7f, 0xb5, 0xa0, 0x88, 0xd4, 0xf4, 0x0e, 0x65, 0x60, 0x67, 0x51, 0xb7, 0xe6, 0x04,
	0x4c, 0x79, 0xcd, 0x82, 0xe3, 0xd2, 0xc6, 0x14, 0x4e, 0x91, 0xea, 0xce, 0x57, 0x87, 0xdc, 0x95,
	0xf9, 0x60, 0xfa, 0x94, 0x4e, 0x4d, 0x27, 0xea, 0x50, 0x13, 0x0f, 0x9d, 0xcf, 0xcf, 0x06, 0x1c,
	0x93, 0x51, 0xdf, 0x8f, 0x64, 0x1a, 0x64, 0xf9, 0x58, 0xa5, 0x0c, 0x79, 0x34, 0x64, 0x7d, 0x15,
	0x58, 0x33, 0xcb, 0xb8, 0xba, 0x42, 0x7a, 0x05, 0x98, 0x7c, 0xaf, 0xac, 0x75, 0xed, 0xf0, 0xee,
	0xa5, 0xe1, 0x3f, 0x27, 0xf7, 0xa7, 0x5d, 0x4c, 0x4e, 0xed, 0x7d, 0xff, 0x56, 0xf1, 0x89, 0x4f,
	0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc2, 0xc4, 0xf7, 0x46, 0x5a, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeClient interface {
	GetSyncStatus(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*SyncStatus, error)
	GetGenesis(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Genesis, error)
	GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Version, error)
	ListImplementedServices(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ImplementedServices, error)
	ListPeers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Peers, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) GetSyncStatus(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*SyncStatus, error) {
	out := new(SyncStatus)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1alpha1.Node/GetSyncStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) GetGenesis(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Genesis, error) {
	out := new(Genesis)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1alpha1.Node/GetGenesis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Version, error) {
	out := new(Version)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1alpha1.Node/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) ListImplementedServices(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ImplementedServices, error) {
	out := new(ImplementedServices)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1alpha1.Node/ListImplementedServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) ListPeers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Peers, error) {
	out := new(Peers)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1alpha1.Node/ListPeers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServer is the server API for Node service.
type NodeServer interface {
	GetSyncStatus(context.Context, *empty.Empty) (*SyncStatus, error)
	GetGenesis(context.Context, *empty.Empty) (*Genesis, error)
	GetVersion(context.Context, *empty.Empty) (*Version, error)
	ListImplementedServices(context.Context, *empty.Empty) (*ImplementedServices, error)
	ListPeers(context.Context, *empty.Empty) (*Peers, error)
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_GetSyncStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).GetSyncStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1alpha1.Node/GetSyncStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).GetSyncStatus(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_GetGenesis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).GetGenesis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1alpha1.Node/GetGenesis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).GetGenesis(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1alpha1.Node/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).GetVersion(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_ListImplementedServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).ListImplementedServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1alpha1.Node/ListImplementedServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).ListImplementedServices(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_ListPeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).ListPeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1alpha1.Node/ListPeers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).ListPeers(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ethereum.eth.v1alpha1.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSyncStatus",
			Handler:    _Node_GetSyncStatus_Handler,
		},
		{
			MethodName: "GetGenesis",
			Handler:    _Node_GetGenesis_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _Node_GetVersion_Handler,
		},
		{
			MethodName: "ListImplementedServices",
			Handler:    _Node_ListImplementedServices_Handler,
		},
		{
			MethodName: "ListPeers",
			Handler:    _Node_ListPeers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eth/v1alpha1/node.proto",
}