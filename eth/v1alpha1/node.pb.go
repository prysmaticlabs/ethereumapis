// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eth/v1alpha1/node.proto

package eth

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type SyncState int32

const (
	SyncState_SYNC_UNKNOWN  SyncState = 0
	SyncState_SYNC_INACTIVE SyncState = 1
	SyncState_SYNC_CATCHUP  SyncState = 2
	SyncState_SYNC_FULL     SyncState = 3
)

var SyncState_name = map[int32]string{
	0: "SYNC_UNKNOWN",
	1: "SYNC_INACTIVE",
	2: "SYNC_CATCHUP",
	3: "SYNC_FULL",
}

var SyncState_value = map[string]int32{
	"SYNC_UNKNOWN":  0,
	"SYNC_INACTIVE": 1,
	"SYNC_CATCHUP":  2,
	"SYNC_FULL":     3,
}

func (x SyncState) String() string {
	return proto.EnumName(SyncState_name, int32(x))
}

func (SyncState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4d5c81e0a630f5ab, []int{0}
}

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
	return fileDescriptor_4d5c81e0a630f5ab, []int{1}
}

type NodeInfo struct {
	NodeId                     string    `protobuf:"bytes,1001,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Version                    string    `protobuf:"bytes,1002,opt,name=version,proto3" json:"version,omitempty"`
	Addresses                  []string  `protobuf:"bytes,2001,rep,name=addresses,proto3" json:"addresses,omitempty"`
	Peers                      []*Peer   `protobuf:"bytes,2002,rep,name=peers,proto3" json:"peers,omitempty"`
	SyncState                  SyncState `protobuf:"varint,3001,opt,name=sync_state,json=syncState,proto3,enum=ethereum.eth.v1alpha1.SyncState" json:"sync_state,omitempty"`
	CurrentEpoch               uint64    `protobuf:"varint,3002,opt,name=current_epoch,json=currentEpoch,proto3" json:"current_epoch,omitempty"`
	CurrentSlot                uint64    `protobuf:"varint,3003,opt,name=current_slot,json=currentSlot,proto3" json:"current_slot,omitempty"`
	CurrentBlockRoot           uint64    `protobuf:"varint,3004,opt,name=current_block_root,json=currentBlockRoot,proto3" json:"current_block_root,omitempty"`
	FinalizedEpoch             uint64    `protobuf:"varint,3005,opt,name=finalized_epoch,json=finalizedEpoch,proto3" json:"finalized_epoch,omitempty"`
	FinalizedSlot              uint64    `protobuf:"varint,3006,opt,name=finalized_slot,json=finalizedSlot,proto3" json:"finalized_slot,omitempty"`
	FinalizedBlockRoot         uint64    `protobuf:"varint,3007,opt,name=finalized_block_root,json=finalizedBlockRoot,proto3" json:"finalized_block_root,omitempty"`
	JustifiedEpoch             uint64    `protobuf:"varint,3008,opt,name=justified_epoch,json=justifiedEpoch,proto3" json:"justified_epoch,omitempty"`
	JustifiedSlot              uint64    `protobuf:"varint,3009,opt,name=justified_slot,json=justifiedSlot,proto3" json:"justified_slot,omitempty"`
	JustifiedBlockRoot         uint64    `protobuf:"varint,3010,opt,name=justified_block_root,json=justifiedBlockRoot,proto3" json:"justified_block_root,omitempty"`
	PreviousJustifiedEpoch     uint64    `protobuf:"varint,3011,opt,name=previous_justified_epoch,json=previousJustifiedEpoch,proto3" json:"previous_justified_epoch,omitempty"`
	PreviousJustifiedSlot      uint64    `protobuf:"varint,3012,opt,name=previous_justified_slot,json=previousJustifiedSlot,proto3" json:"previous_justified_slot,omitempty"`
	PreviousJustifiedBlockRoot uint64    `protobuf:"varint,3013,opt,name=previous_justified_block_root,json=previousJustifiedBlockRoot,proto3" json:"previous_justified_block_root,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}  `json:"-"`
	XXX_unrecognized           []byte    `json:"-"`
	XXX_sizecache              int32     `json:"-"`
}

func (m *NodeInfo) Reset()         { *m = NodeInfo{} }
func (m *NodeInfo) String() string { return proto.CompactTextString(m) }
func (*NodeInfo) ProtoMessage()    {}
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d5c81e0a630f5ab, []int{0}
}

func (m *NodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeInfo.Unmarshal(m, b)
}
func (m *NodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeInfo.Marshal(b, m, deterministic)
}
func (m *NodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeInfo.Merge(m, src)
}
func (m *NodeInfo) XXX_Size() int {
	return xxx_messageInfo_NodeInfo.Size(m)
}
func (m *NodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NodeInfo proto.InternalMessageInfo

func (m *NodeInfo) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *NodeInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *NodeInfo) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *NodeInfo) GetPeers() []*Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

func (m *NodeInfo) GetSyncState() SyncState {
	if m != nil {
		return m.SyncState
	}
	return SyncState_SYNC_UNKNOWN
}

func (m *NodeInfo) GetCurrentEpoch() uint64 {
	if m != nil {
		return m.CurrentEpoch
	}
	return 0
}

func (m *NodeInfo) GetCurrentSlot() uint64 {
	if m != nil {
		return m.CurrentSlot
	}
	return 0
}

func (m *NodeInfo) GetCurrentBlockRoot() uint64 {
	if m != nil {
		return m.CurrentBlockRoot
	}
	return 0
}

func (m *NodeInfo) GetFinalizedEpoch() uint64 {
	if m != nil {
		return m.FinalizedEpoch
	}
	return 0
}

func (m *NodeInfo) GetFinalizedSlot() uint64 {
	if m != nil {
		return m.FinalizedSlot
	}
	return 0
}

func (m *NodeInfo) GetFinalizedBlockRoot() uint64 {
	if m != nil {
		return m.FinalizedBlockRoot
	}
	return 0
}

func (m *NodeInfo) GetJustifiedEpoch() uint64 {
	if m != nil {
		return m.JustifiedEpoch
	}
	return 0
}

func (m *NodeInfo) GetJustifiedSlot() uint64 {
	if m != nil {
		return m.JustifiedSlot
	}
	return 0
}

func (m *NodeInfo) GetJustifiedBlockRoot() uint64 {
	if m != nil {
		return m.JustifiedBlockRoot
	}
	return 0
}

func (m *NodeInfo) GetPreviousJustifiedEpoch() uint64 {
	if m != nil {
		return m.PreviousJustifiedEpoch
	}
	return 0
}

func (m *NodeInfo) GetPreviousJustifiedSlot() uint64 {
	if m != nil {
		return m.PreviousJustifiedSlot
	}
	return 0
}

func (m *NodeInfo) GetPreviousJustifiedBlockRoot() uint64 {
	if m != nil {
		return m.PreviousJustifiedBlockRoot
	}
	return 0
}

type Peer struct {
	Enr                  string        `protobuf:"bytes,1,opt,name=enr,proto3" json:"enr,omitempty"`
	Address              string        `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Direction            PeerDirection `protobuf:"varint,3,opt,name=direction,proto3,enum=ethereum.eth.v1alpha1.PeerDirection" json:"direction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Peer) Reset()         { *m = Peer{} }
func (m *Peer) String() string { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()    {}
func (*Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d5c81e0a630f5ab, []int{1}
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

func (m *Peer) GetEnr() string {
	if m != nil {
		return m.Enr
	}
	return ""
}

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
	return fileDescriptor_4d5c81e0a630f5ab, []int{2}
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

func init() {
	proto.RegisterEnum("ethereum.eth.v1alpha1.SyncState", SyncState_name, SyncState_value)
	proto.RegisterEnum("ethereum.eth.v1alpha1.PeerDirection", PeerDirection_name, PeerDirection_value)
	proto.RegisterType((*NodeInfo)(nil), "ethereum.eth.v1alpha1.NodeInfo")
	proto.RegisterType((*Peer)(nil), "ethereum.eth.v1alpha1.Peer")
	proto.RegisterType((*ImplementedServices)(nil), "ethereum.eth.v1alpha1.ImplementedServices")
}

func init() { proto.RegisterFile("eth/v1alpha1/node.proto", fileDescriptor_4d5c81e0a630f5ab) }

var fileDescriptor_4d5c81e0a630f5ab = []byte{
	// 781 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xcd, 0x8e, 0xe3, 0x44,
	0x10, 0xc6, 0xc9, 0x30, 0x19, 0x57, 0x26, 0xbb, 0xa6, 0x61, 0x36, 0x9e, 0x0c, 0xa3, 0x8d, 0xc2,
	0x0a, 0x45, 0x23, 0x61, 0x2b, 0x41, 0x62, 0x85, 0xb8, 0xb0, 0xc9, 0x06, 0x08, 0x8c, 0x3c, 0x21,
	0x99, 0x80, 0x40, 0x2b, 0x45, 0x8e, 0xdd, 0x89, 0x1b, 0x1c, 0xb7, 0xe5, 0xee, 0x44, 0x0a, 0xc7,
	0x7d, 0x01, 0x90, 0x78, 0x0b, 0xde, 0x80, 0xff, 0xff, 0x0b, 0x37, 0xe0, 0x15, 0xb8, 0xc0, 0x99,
	0x07, 0x40, 0xed, 0xf6, 0x4f, 0xc2, 0x24, 0x7b, 0xeb, 0xfa, 0xea, 0xfb, 0xaa, 0xbe, 0xb6, 0xab,
	0x0b, 0xaa, 0x98, 0x7b, 0xe6, 0xaa, 0x65, 0xfb, 0xa1, 0x67, 0xb7, 0xcc, 0x80, 0xba, 0xd8, 0x08,
	0x23, 0xca, 0x29, 0x3a, 0xc1, 0xdc, 0xc3, 0x11, 0x5e, 0x2e, 0x0c, 0xcc, 0x3d, 0x23, 0x65, 0xd4,
	0x9e, 0x9f, 0x53, 0x3a, 0xf7, 0xb1, 0x69, 0x87, 0xc4, 0xb4, 0x83, 0x80, 0x72, 0x9b, 0x13, 0x1a,
	0x30, 0x29, 0xaa, 0x9d, 0x25, 0xd9, 0x38, 0x9a, 0x2e, 0x67, 0x26, 0x5e, 0x84, 0x7c, 0x2d, 0x93,
	0x8d, 0x4f, 0x0f, 0xe1, 0xc8, 0xa2, 0x2e, 0xee, 0x07, 0x33, 0x8a, 0x74, 0x28, 0x89, 0x66, 0x13,
	0xe2, 0xea, 0x7f, 0x97, 0xea, 0x4a, 0x53, 0x1d, 0x1e, 0x8a, 0xb8, 0xef, 0xa2, 0x53, 0x28, 0xad,
	0x70, 0xc4, 0x08, 0x0d, 0xf4, 0x7f, 0x64, 0x26, 0x8d, 0xd1, 0x39, 0xa8, 0xb6, 0xeb, 0x46, 0x98,
	0x31, 0xcc, 0xf4, 0xdf, 0x6f, 0xd7, 0x8b, 0x4d, 0x75, 0x98, 0x23, 0xa8, 0x0d, 0x4f, 0x87, 0x18,
	0x47, 0x4c, 0xff, 0x43, 0xa4, 0xca, 0xed, 0x33, 0x63, 0xe7, 0x1d, 0x8c, 0x01, 0xc6, 0xd1, 0x50,
	0x52, 0xd1, 0xeb, 0x00, 0x6c, 0x1d, 0x38, 0x13, 0xc6, 0x6d, 0x8e, 0xf5, 0x2f, 0xab, 0x75, 0xa5,
	0x79, 0xab, 0x5d, 0xdf, 0x23, 0x1c, 0xad, 0x03, 0x67, 0x24, 0x88, 0x43, 0x95, 0xa5, 0x47, 0x74,
	0x0f, 0x2a, 0xce, 0x32, 0x8a, 0x70, 0xc0, 0x27, 0x38, 0xa4, 0x8e, 0xa7, 0x7f, 0x25, 0x8a, 0x1c,
	0x0c, 0x8f, 0x13, 0xb4, 0x27, 0x40, 0xd4, 0x80, 0x34, 0x9e, 0x30, 0x9f, 0x72, 0xfd, 0x6b, 0x49,
	0x2a, 0x27, 0xe0, 0xc8, 0xa7, 0x1c, 0xbd, 0x04, 0x28, 0xe5, 0x4c, 0x7d, 0xea, 0x7c, 0x3c, 0x89,
	0x28, 0xe5, 0xfa, 0x37, 0x92, 0xa9, 0x25, 0xa9, 0x8e, 0xc8, 0x0c, 0x29, 0xe5, 0xa8, 0x09, 0xb7,
	0x67, 0x24, 0xb0, 0x7d, 0xf2, 0x09, 0x76, 0x93, 0xd6, 0xdf, 0x4a, 0xee, 0xad, 0x0c, 0x97, 0xcd,
	0x5f, 0x84, 0x1c, 0x91, 0xed, 0xbf, 0x93, 0xc4, 0x4a, 0x06, 0xc7, 0x06, 0x5a, 0xf0, 0x5c, 0xce,
	0xdb, 0xb0, 0xf0, 0xbd, 0x64, 0xa3, 0x2c, 0xb9, 0x65, 0xe2, 0xa3, 0x25, 0xe3, 0x64, 0x46, 0x32,
	0x13, 0x3f, 0x24, 0x26, 0x32, 0x3c, 0x33, 0x91, 0x33, 0x63, 0x13, 0x3f, 0x26, 0x26, 0x32, 0x38,
	0x35, 0x91, 0xf3, 0x36, 0x4c, 0xfc, 0x94, 0x98, 0xc8, 0x92, 0xb9, 0x89, 0x57, 0x41, 0x0f, 0x23,
	0xbc, 0x22, 0x74, 0xc9, 0x26, 0xff, 0x77, 0xf3, 0xb3, 0x94, 0xdd, 0x49, 0x09, 0x6f, 0x6f, 0xbb,
	0xba, 0x0f, 0xd5, 0x1d, 0xd2, 0xd8, 0xde, 0x2f, 0x52, 0x79, 0x72, 0x43, 0x19, 0xdb, 0xec, 0xc0,
	0xf9, 0x0e, 0xe1, 0x86, 0xdf, 0x5f, 0xa5, 0xbc, 0x76, 0x43, 0x9e, 0xf9, 0x6e, 0xac, 0xe0, 0x40,
	0xcc, 0x22, 0xd2, 0xa0, 0x88, 0x83, 0x48, 0x57, 0xe2, 0x69, 0x17, 0x47, 0xf1, 0x3c, 0x92, 0xb9,
	0xd6, 0x0b, 0xf2, 0x0d, 0x24, 0x21, 0xea, 0x80, 0xea, 0x92, 0x08, 0x3b, 0xe2, 0xd9, 0xe9, 0xc5,
	0x78, 0x5c, 0xef, 0x3d, 0x61, 0xce, 0x1f, 0xa6, 0xdc, 0x61, 0x2e, 0x6b, 0xb4, 0xe0, 0xd9, 0xfe,
	0x22, 0xf4, 0xf1, 0x02, 0x07, 0x1c, 0xbb, 0x23, 0x1c, 0xad, 0x88, 0x83, 0x19, 0xaa, 0xc1, 0x11,
	0x4b, 0xce, 0xba, 0x12, 0x3f, 0xae, 0x2c, 0xbe, 0x78, 0x17, 0xd4, 0x6c, 0xfa, 0x91, 0x06, 0xc7,
	0xa3, 0x0f, 0xac, 0xee, 0x64, 0x6c, 0xbd, 0x63, 0x5d, 0xbd, 0x6f, 0x69, 0x4f, 0xa1, 0x67, 0xa0,
	0x12, 0x23, 0x7d, 0xeb, 0x41, 0xf7, 0xba, 0xff, 0x5e, 0x4f, 0x53, 0x32, 0x52, 0xf7, 0xc1, 0x75,
	0xf7, 0xad, 0xf1, 0x40, 0x2b, 0xa0, 0x0a, 0xa8, 0x31, 0xf2, 0xc6, 0xf8, 0xf2, 0x52, 0x2b, 0x5e,
	0xdc, 0x87, 0xca, 0x96, 0x43, 0x54, 0x86, 0x52, 0x5e, 0xb1, 0x0c, 0xa5, 0xbe, 0xd5, 0xb9, 0x1a,
	0x5b, 0x0f, 0x35, 0x05, 0x1d, 0xc3, 0xd1, 0xd5, 0xf8, 0x5a, 0x46, 0x85, 0xf6, 0xbf, 0x0a, 0x1c,
	0x88, 0x45, 0x82, 0xe6, 0x50, 0x7e, 0x13, 0xf3, 0x6c, 0xa7, 0xdc, 0x31, 0xe4, 0xfa, 0x31, 0xd2,
	0xf5, 0x63, 0xf4, 0xc4, 0xfa, 0xa9, 0xdd, 0xdd, 0xf3, 0x7d, 0x52, 0x61, 0xe3, 0xee, 0xe3, 0x3f,
	0xff, 0xfa, 0xbc, 0x70, 0x8a, 0xaa, 0xe6, 0x8d, 0x6d, 0x68, 0x12, 0x51, 0xf9, 0xb1, 0x02, 0xd5,
	0x4b, 0xc2, 0xf8, 0xae, 0xaf, 0xb6, 0xaf, 0xeb, 0xc5, 0x9e, 0xae, 0x3b, 0x6a, 0x34, 0x5e, 0x88,
	0x0d, 0x9c, 0xa3, 0xb3, 0x1d, 0x06, 0xd2, 0x5f, 0xd0, 0xf9, 0x4c, 0x81, 0x53, 0x1a, 0xcd, 0x77,
	0x97, 0xed, 0xa8, 0xe2, 0x36, 0x03, 0x61, 0x61, 0xa0, 0x7c, 0xf8, 0xca, 0x9c, 0x70, 0x6f, 0x39,
	0x35, 0x1c, 0xba, 0x30, 0xc3, 0x68, 0xcd, 0x16, 0x36, 0x27, 0x8e, 0x6f, 0x4f, 0x99, 0x99, 0x8a,
	0xed, 0x90, 0xb0, 0xad, 0x66, 0xaf, 0x61, 0xee, 0x7d, 0x51, 0x38, 0xe9, 0xa5, 0xc5, 0x7b, 0x1b,
	0xc5, 0x7f, 0xcb, 0xf1, 0x47, 0x3d, 0xee, 0x3d, 0x4a, 0xf1, 0xe9, 0x61, 0x7c, 0xe7, 0x97, 0xff,
	0x0b, 0x00, 0x00, 0xff, 0xff, 0x80, 0x7c, 0xd3, 0x5a, 0x46, 0x06, 0x00, 0x00,
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
	GetNodeInfo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*NodeInfo, error)
	ListImplementedServices(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ImplementedServices, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) GetNodeInfo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*NodeInfo, error) {
	out := new(NodeInfo)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1alpha1.Node/GetNodeInfo", in, out, opts...)
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

// NodeServer is the server API for Node service.
type NodeServer interface {
	GetNodeInfo(context.Context, *empty.Empty) (*NodeInfo, error)
	ListImplementedServices(context.Context, *empty.Empty) (*ImplementedServices, error)
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_GetNodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).GetNodeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1alpha1.Node/GetNodeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).GetNodeInfo(ctx, req.(*empty.Empty))
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

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ethereum.eth.v1alpha1.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNodeInfo",
			Handler:    _Node_GetNodeInfo_Handler,
		},
		{
			MethodName: "ListImplementedServices",
			Handler:    _Node_ListImplementedServices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eth/v1alpha1/node.proto",
}
