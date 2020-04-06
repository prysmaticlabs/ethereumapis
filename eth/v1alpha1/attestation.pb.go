// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eth/v1alpha1/attestation.proto

package eth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Attestation struct {
	AggregationBits      []byte           `protobuf:"bytes,1,opt,name=aggregation_bits,json=aggregationBits,proto3" json:"aggregation_bits,omitempty"`
	Data                 *AttestationData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Signature            []byte           `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Attestation) Reset()         { *m = Attestation{} }
func (m *Attestation) String() string { return proto.CompactTextString(m) }
func (*Attestation) ProtoMessage()    {}
func (*Attestation) Descriptor() ([]byte, []int) {
	return fileDescriptor_7894c78397fc93a1, []int{0}
}

func (m *Attestation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attestation.Unmarshal(m, b)
}
func (m *Attestation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attestation.Marshal(b, m, deterministic)
}
func (m *Attestation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attestation.Merge(m, src)
}
func (m *Attestation) XXX_Size() int {
	return xxx_messageInfo_Attestation.Size(m)
}
func (m *Attestation) XXX_DiscardUnknown() {
	xxx_messageInfo_Attestation.DiscardUnknown(m)
}

var xxx_messageInfo_Attestation proto.InternalMessageInfo

func (m *Attestation) GetAggregationBits() []byte {
	if m != nil {
		return m.AggregationBits
	}
	return nil
}

func (m *Attestation) GetData() *AttestationData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Attestation) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type AggregateAttestationAndProof struct {
	AggregatorIndex      uint64       `protobuf:"varint,1,opt,name=aggregator_index,json=aggregatorIndex,proto3" json:"aggregator_index,omitempty"`
	Aggregate            *Attestation `protobuf:"bytes,3,opt,name=aggregate,proto3" json:"aggregate,omitempty"`
	SelectionProof       []byte       `protobuf:"bytes,2,opt,name=selection_proof,json=selectionProof,proto3" json:"selection_proof,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AggregateAttestationAndProof) Reset()         { *m = AggregateAttestationAndProof{} }
func (m *AggregateAttestationAndProof) String() string { return proto.CompactTextString(m) }
func (*AggregateAttestationAndProof) ProtoMessage()    {}
func (*AggregateAttestationAndProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_7894c78397fc93a1, []int{1}
}

func (m *AggregateAttestationAndProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggregateAttestationAndProof.Unmarshal(m, b)
}
func (m *AggregateAttestationAndProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggregateAttestationAndProof.Marshal(b, m, deterministic)
}
func (m *AggregateAttestationAndProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregateAttestationAndProof.Merge(m, src)
}
func (m *AggregateAttestationAndProof) XXX_Size() int {
	return xxx_messageInfo_AggregateAttestationAndProof.Size(m)
}
func (m *AggregateAttestationAndProof) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregateAttestationAndProof.DiscardUnknown(m)
}

var xxx_messageInfo_AggregateAttestationAndProof proto.InternalMessageInfo

func (m *AggregateAttestationAndProof) GetAggregatorIndex() uint64 {
	if m != nil {
		return m.AggregatorIndex
	}
	return 0
}

func (m *AggregateAttestationAndProof) GetAggregate() *Attestation {
	if m != nil {
		return m.Aggregate
	}
	return nil
}

func (m *AggregateAttestationAndProof) GetSelectionProof() []byte {
	if m != nil {
		return m.SelectionProof
	}
	return nil
}

type SignedAggregateAttestationAndProof struct {
	Message              *AggregateAttestationAndProof `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Signature            []byte                        `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SignedAggregateAttestationAndProof) Reset()         { *m = SignedAggregateAttestationAndProof{} }
func (m *SignedAggregateAttestationAndProof) String() string { return proto.CompactTextString(m) }
func (*SignedAggregateAttestationAndProof) ProtoMessage()    {}
func (*SignedAggregateAttestationAndProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_7894c78397fc93a1, []int{2}
}

func (m *SignedAggregateAttestationAndProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedAggregateAttestationAndProof.Unmarshal(m, b)
}
func (m *SignedAggregateAttestationAndProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedAggregateAttestationAndProof.Marshal(b, m, deterministic)
}
func (m *SignedAggregateAttestationAndProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedAggregateAttestationAndProof.Merge(m, src)
}
func (m *SignedAggregateAttestationAndProof) XXX_Size() int {
	return xxx_messageInfo_SignedAggregateAttestationAndProof.Size(m)
}
func (m *SignedAggregateAttestationAndProof) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedAggregateAttestationAndProof.DiscardUnknown(m)
}

var xxx_messageInfo_SignedAggregateAttestationAndProof proto.InternalMessageInfo

func (m *SignedAggregateAttestationAndProof) GetMessage() *AggregateAttestationAndProof {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *SignedAggregateAttestationAndProof) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type AttestationData struct {
	Slot                 uint64      `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	CommitteeIndex       uint64      `protobuf:"varint,2,opt,name=committee_index,json=committeeIndex,proto3" json:"committee_index,omitempty"`
	BeaconBlockRoot      []byte      `protobuf:"bytes,3,opt,name=beacon_block_root,json=beaconBlockRoot,proto3" json:"beacon_block_root,omitempty"`
	Source               *Checkpoint `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Target               *Checkpoint `protobuf:"bytes,5,opt,name=target,proto3" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AttestationData) Reset()         { *m = AttestationData{} }
func (m *AttestationData) String() string { return proto.CompactTextString(m) }
func (*AttestationData) ProtoMessage()    {}
func (*AttestationData) Descriptor() ([]byte, []int) {
	return fileDescriptor_7894c78397fc93a1, []int{3}
}

func (m *AttestationData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttestationData.Unmarshal(m, b)
}
func (m *AttestationData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttestationData.Marshal(b, m, deterministic)
}
func (m *AttestationData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttestationData.Merge(m, src)
}
func (m *AttestationData) XXX_Size() int {
	return xxx_messageInfo_AttestationData.Size(m)
}
func (m *AttestationData) XXX_DiscardUnknown() {
	xxx_messageInfo_AttestationData.DiscardUnknown(m)
}

var xxx_messageInfo_AttestationData proto.InternalMessageInfo

func (m *AttestationData) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *AttestationData) GetCommitteeIndex() uint64 {
	if m != nil {
		return m.CommitteeIndex
	}
	return 0
}

func (m *AttestationData) GetBeaconBlockRoot() []byte {
	if m != nil {
		return m.BeaconBlockRoot
	}
	return nil
}

func (m *AttestationData) GetSource() *Checkpoint {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *AttestationData) GetTarget() *Checkpoint {
	if m != nil {
		return m.Target
	}
	return nil
}

type Crosslink struct {
	Shard                uint64   `protobuf:"varint,1,opt,name=shard,proto3" json:"shard,omitempty"`
	ParentRoot           []byte   `protobuf:"bytes,2,opt,name=parent_root,json=parentRoot,proto3" json:"parent_root,omitempty"`
	StartEpoch           uint64   `protobuf:"varint,3,opt,name=start_epoch,json=startEpoch,proto3" json:"start_epoch,omitempty"`
	EndEpoch             uint64   `protobuf:"varint,4,opt,name=end_epoch,json=endEpoch,proto3" json:"end_epoch,omitempty"`
	DataRoot             []byte   `protobuf:"bytes,5,opt,name=data_root,json=dataRoot,proto3" json:"data_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Crosslink) Reset()         { *m = Crosslink{} }
func (m *Crosslink) String() string { return proto.CompactTextString(m) }
func (*Crosslink) ProtoMessage()    {}
func (*Crosslink) Descriptor() ([]byte, []int) {
	return fileDescriptor_7894c78397fc93a1, []int{4}
}

func (m *Crosslink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Crosslink.Unmarshal(m, b)
}
func (m *Crosslink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Crosslink.Marshal(b, m, deterministic)
}
func (m *Crosslink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Crosslink.Merge(m, src)
}
func (m *Crosslink) XXX_Size() int {
	return xxx_messageInfo_Crosslink.Size(m)
}
func (m *Crosslink) XXX_DiscardUnknown() {
	xxx_messageInfo_Crosslink.DiscardUnknown(m)
}

var xxx_messageInfo_Crosslink proto.InternalMessageInfo

func (m *Crosslink) GetShard() uint64 {
	if m != nil {
		return m.Shard
	}
	return 0
}

func (m *Crosslink) GetParentRoot() []byte {
	if m != nil {
		return m.ParentRoot
	}
	return nil
}

func (m *Crosslink) GetStartEpoch() uint64 {
	if m != nil {
		return m.StartEpoch
	}
	return 0
}

func (m *Crosslink) GetEndEpoch() uint64 {
	if m != nil {
		return m.EndEpoch
	}
	return 0
}

func (m *Crosslink) GetDataRoot() []byte {
	if m != nil {
		return m.DataRoot
	}
	return nil
}

type Checkpoint struct {
	Epoch                uint64   `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	Root                 []byte   `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Checkpoint) Reset()         { *m = Checkpoint{} }
func (m *Checkpoint) String() string { return proto.CompactTextString(m) }
func (*Checkpoint) ProtoMessage()    {}
func (*Checkpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_7894c78397fc93a1, []int{5}
}

func (m *Checkpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Checkpoint.Unmarshal(m, b)
}
func (m *Checkpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Checkpoint.Marshal(b, m, deterministic)
}
func (m *Checkpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Checkpoint.Merge(m, src)
}
func (m *Checkpoint) XXX_Size() int {
	return xxx_messageInfo_Checkpoint.Size(m)
}
func (m *Checkpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Checkpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Checkpoint proto.InternalMessageInfo

func (m *Checkpoint) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *Checkpoint) GetRoot() []byte {
	if m != nil {
		return m.Root
	}
	return nil
}

func init() {
	proto.RegisterType((*Attestation)(nil), "ethereum.eth.v1alpha1.Attestation")
	proto.RegisterType((*AggregateAttestationAndProof)(nil), "ethereum.eth.v1alpha1.AggregateAttestationAndProof")
	proto.RegisterType((*SignedAggregateAttestationAndProof)(nil), "ethereum.eth.v1alpha1.SignedAggregateAttestationAndProof")
	proto.RegisterType((*AttestationData)(nil), "ethereum.eth.v1alpha1.AttestationData")
	proto.RegisterType((*Crosslink)(nil), "ethereum.eth.v1alpha1.Crosslink")
	proto.RegisterType((*Checkpoint)(nil), "ethereum.eth.v1alpha1.Checkpoint")
}

func init() { proto.RegisterFile("eth/v1alpha1/attestation.proto", fileDescriptor_7894c78397fc93a1) }

var fileDescriptor_7894c78397fc93a1 = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0x96, 0xf3, 0xbb, 0xfd, 0x9b, 0x71, 0xd5, 0x94, 0x15, 0x95, 0x82, 0xa8, 0x68, 0xf1, 0x01,
	0x0a, 0x07, 0x47, 0x4d, 0xa5, 0x4a, 0xc0, 0x85, 0xa4, 0xe4, 0xc0, 0x01, 0x29, 0x32, 0x37, 0x54,
	0x29, 0xda, 0xd8, 0x83, 0xbd, 0x8a, 0xed, 0xb5, 0x76, 0x27, 0x08, 0x1e, 0x03, 0x2e, 0x9c, 0x78,
	0x01, 0xce, 0x3c, 0x05, 0x4f, 0xc3, 0x23, 0xa0, 0x5d, 0xdb, 0xb1, 0xa9, 0xda, 0xc2, 0xcd, 0xf3,
	0xcd, 0xce, 0x37, 0xdf, 0x7c, 0xb3, 0x5e, 0x78, 0x80, 0x94, 0x8e, 0x3e, 0x9c, 0xf2, 0xac, 0x4c,
	0xf9, 0xe9, 0x88, 0x13, 0xa1, 0x26, 0x4e, 0x42, 0x16, 0x41, 0xa9, 0x24, 0x49, 0x76, 0x80, 0x94,
	0xa2, 0xc2, 0x75, 0x1e, 0x20, 0xa5, 0x41, 0x73, 0xd0, 0xff, 0xe2, 0x80, 0x37, 0x69, 0x0f, 0xb3,
	0x27, 0xb0, 0xcf, 0x93, 0x44, 0x61, 0x62, 0xc3, 0xc5, 0x52, 0x90, 0x1e, 0x3a, 0xc7, 0xce, 0xc9,
	0x6e, 0x38, 0xe8, 0xe0, 0x53, 0x41, 0x9a, 0x3d, 0x07, 0x37, 0xe6, 0xc4, 0x87, 0xbd, 0x63, 0xe7,
	0xc4, 0x1b, 0x3f, 0x0a, 0xae, 0x6d, 0x10, 0x74, 0xc8, 0x5f, 0x71, 0xe2, 0xa1, 0xad, 0x61, 0x87,
	0xd0, 0xd7, 0x22, 0x29, 0x38, 0xad, 0x15, 0x0e, 0xff, 0xb3, 0xfc, 0x2d, 0xe0, 0xff, 0x70, 0xe0,
	0x70, 0x52, 0x77, 0xc3, 0x0e, 0xc1, 0xa4, 0x88, 0xe7, 0x4a, 0xca, 0xf7, 0x5d, 0x95, 0x52, 0x2d,
	0x44, 0x11, 0xe3, 0x47, 0xab, 0xd2, 0x6d, 0x55, 0x4a, 0xf5, 0xda, 0xc0, 0xec, 0x25, 0xf4, 0x1b,
	0xa8, 0xea, 0xe4, 0x8d, 0xfd, 0xbf, 0x4b, 0x0d, 0xdb, 0x22, 0xf6, 0x18, 0x06, 0x1a, 0x33, 0x8c,
	0xac, 0x21, 0xa5, 0xe9, 0x6f, 0x47, 0xde, 0x0d, 0xf7, 0x36, 0xb0, 0x55, 0xe5, 0x7f, 0x76, 0xc0,
	0x7f, 0x2b, 0x92, 0x02, 0xe3, 0x5b, 0xc5, 0xbf, 0x81, 0xff, 0x73, 0xd4, 0x9a, 0x27, 0x68, 0x35,
	0x7b, 0xe3, 0xb3, 0x9b, 0xf4, 0xdc, 0xc2, 0x12, 0x36, 0x1c, 0x7f, 0x5a, 0xd9, 0xbb, 0x6a, 0xe5,
	0x2f, 0x07, 0x06, 0x57, 0x56, 0xc0, 0x18, 0xb8, 0x3a, 0x93, 0x54, 0x3b, 0x66, 0xbf, 0xcd, 0x90,
	0x91, 0xcc, 0x73, 0x41, 0x84, 0x58, 0x1b, 0xda, 0xb3, 0xe9, 0xbd, 0x0d, 0x5c, 0xf9, 0xf9, 0x14,
	0xee, 0x2c, 0x91, 0x47, 0xe6, 0x6e, 0x64, 0x32, 0x5a, 0x2d, 0x94, 0x94, 0x54, 0x6f, 0x70, 0x50,
	0x25, 0xa6, 0x06, 0x0f, 0xa5, 0x24, 0xf6, 0x0c, 0xb6, 0xb5, 0x5c, 0xab, 0x08, 0x87, 0xae, 0x1d,
	0xf4, 0xe1, 0x0d, 0x83, 0x5e, 0xa4, 0x18, 0xad, 0x4a, 0x29, 0x0a, 0x0a, 0xeb, 0x02, 0x53, 0x4a,
	0x5c, 0x25, 0x48, 0xc3, 0xad, 0x7f, 0x2e, 0xad, 0x0a, 0xfc, 0x6f, 0x0e, 0xf4, 0x2f, 0x94, 0xd4,
	0x3a, 0x13, 0xc5, 0x8a, 0xdd, 0x85, 0x2d, 0x9d, 0x72, 0x15, 0xd7, 0xd3, 0x56, 0x01, 0x3b, 0x02,
	0xaf, 0xe4, 0x0a, 0x0b, 0xaa, 0xf4, 0x57, 0xb6, 0x41, 0x05, 0x59, 0xe9, 0x47, 0xe0, 0x69, 0xe2,
	0x8a, 0x16, 0x58, 0xca, 0x28, 0xb5, 0x03, 0xba, 0x21, 0x58, 0x68, 0x66, 0x10, 0x76, 0x1f, 0xfa,
	0x58, 0xc4, 0x75, 0xda, 0xb5, 0xe9, 0x1d, 0x2c, 0xe2, 0x4d, 0xd2, 0x5c, 0xf3, 0x8a, 0x7c, 0xcb,
	0x92, 0xef, 0x18, 0xc0, 0x50, 0xfb, 0xe7, 0x00, 0xad, 0x6a, 0xa3, 0xaf, 0xe2, 0xa8, 0xf5, 0xd9,
	0xc0, 0xac, 0xa8, 0x23, 0xcc, 0x7e, 0x4f, 0xbf, 0x3a, 0x70, 0x4f, 0xaa, 0xe4, 0x7a, 0x23, 0xa6,
	0xfb, 0x9d, 0x2d, 0xcf, 0xcd, 0x1f, 0x3f, 0x77, 0xde, 0x9d, 0x27, 0x82, 0xd2, 0xf5, 0x32, 0x88,
	0x64, 0x3e, 0x2a, 0xd5, 0x27, 0x9d, 0x73, 0x12, 0x51, 0xc6, 0x97, 0x7a, 0xd4, 0x70, 0xf0, 0x52,
	0xd8, 0x60, 0xf3, 0x72, 0xbc, 0x40, 0x4a, 0xbf, 0xf7, 0x0e, 0x66, 0x4d, 0x8f, 0x59, 0xa7, 0xc7,
	0xcf, 0x16, 0xbf, 0x9c, 0x51, 0x7a, 0xd9, 0xe0, 0xcb, 0x6d, 0xfb, 0xc4, 0x9c, 0xfd, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0xf5, 0xbd, 0xf5, 0xfd, 0x84, 0x04, 0x00, 0x00,
}
