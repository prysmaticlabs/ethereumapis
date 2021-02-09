// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: eth/v1alpha1/attestation.proto

package eth

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Attestation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AggregationBits []byte           `protobuf:"bytes,1,opt,name=aggregation_bits,json=aggregationBits,proto3" json:"aggregation_bits,omitempty"`
	Data            *AttestationData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Signature       []byte           `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Attestation) Reset() {
	*x = Attestation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1alpha1_attestation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attestation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attestation) ProtoMessage() {}

func (x *Attestation) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1alpha1_attestation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attestation.ProtoReflect.Descriptor instead.
func (*Attestation) Descriptor() ([]byte, []int) {
	return file_eth_v1alpha1_attestation_proto_rawDescGZIP(), []int{0}
}

func (x *Attestation) GetAggregationBits() []byte {
	if x != nil {
		return x.AggregationBits
	}
	return nil
}

func (x *Attestation) GetData() *AttestationData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Attestation) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type AggregateAttestationAndProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AggregatorIndex uint64       `protobuf:"varint,1,opt,name=aggregator_index,json=aggregatorIndex,proto3" json:"aggregator_index,omitempty"`
	Aggregate       *Attestation `protobuf:"bytes,3,opt,name=aggregate,proto3" json:"aggregate,omitempty"`
	SelectionProof  []byte       `protobuf:"bytes,2,opt,name=selection_proof,json=selectionProof,proto3" json:"selection_proof,omitempty"`
}

func (x *AggregateAttestationAndProof) Reset() {
	*x = AggregateAttestationAndProof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1alpha1_attestation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregateAttestationAndProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregateAttestationAndProof) ProtoMessage() {}

func (x *AggregateAttestationAndProof) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1alpha1_attestation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregateAttestationAndProof.ProtoReflect.Descriptor instead.
func (*AggregateAttestationAndProof) Descriptor() ([]byte, []int) {
	return file_eth_v1alpha1_attestation_proto_rawDescGZIP(), []int{1}
}

func (x *AggregateAttestationAndProof) GetAggregatorIndex() uint64 {
	if x != nil {
		return x.AggregatorIndex
	}
	return 0
}

func (x *AggregateAttestationAndProof) GetAggregate() *Attestation {
	if x != nil {
		return x.Aggregate
	}
	return nil
}

func (x *AggregateAttestationAndProof) GetSelectionProof() []byte {
	if x != nil {
		return x.SelectionProof
	}
	return nil
}

type SignedAggregateAttestationAndProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   *AggregateAttestationAndProof `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Signature []byte                        `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SignedAggregateAttestationAndProof) Reset() {
	*x = SignedAggregateAttestationAndProof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1alpha1_attestation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedAggregateAttestationAndProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedAggregateAttestationAndProof) ProtoMessage() {}

func (x *SignedAggregateAttestationAndProof) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1alpha1_attestation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedAggregateAttestationAndProof.ProtoReflect.Descriptor instead.
func (*SignedAggregateAttestationAndProof) Descriptor() ([]byte, []int) {
	return file_eth_v1alpha1_attestation_proto_rawDescGZIP(), []int{2}
}

func (x *SignedAggregateAttestationAndProof) GetMessage() *AggregateAttestationAndProof {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *SignedAggregateAttestationAndProof) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type AttestationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot            uint64      `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	CommitteeIndex  uint64      `protobuf:"varint,2,opt,name=committee_index,json=committeeIndex,proto3" json:"committee_index,omitempty"`
	BeaconBlockRoot []byte      `protobuf:"bytes,3,opt,name=beacon_block_root,json=beaconBlockRoot,proto3" json:"beacon_block_root,omitempty"`
	Source          *Checkpoint `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Target          *Checkpoint `protobuf:"bytes,5,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *AttestationData) Reset() {
	*x = AttestationData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1alpha1_attestation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttestationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttestationData) ProtoMessage() {}

func (x *AttestationData) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1alpha1_attestation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttestationData.ProtoReflect.Descriptor instead.
func (*AttestationData) Descriptor() ([]byte, []int) {
	return file_eth_v1alpha1_attestation_proto_rawDescGZIP(), []int{3}
}

func (x *AttestationData) GetSlot() uint64 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *AttestationData) GetCommitteeIndex() uint64 {
	if x != nil {
		return x.CommitteeIndex
	}
	return 0
}

func (x *AttestationData) GetBeaconBlockRoot() []byte {
	if x != nil {
		return x.BeaconBlockRoot
	}
	return nil
}

func (x *AttestationData) GetSource() *Checkpoint {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *AttestationData) GetTarget() *Checkpoint {
	if x != nil {
		return x.Target
	}
	return nil
}

type Checkpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Epoch uint64 `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	Root  []byte `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
}

func (x *Checkpoint) Reset() {
	*x = Checkpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1alpha1_attestation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Checkpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Checkpoint) ProtoMessage() {}

func (x *Checkpoint) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1alpha1_attestation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Checkpoint.ProtoReflect.Descriptor instead.
func (*Checkpoint) Descriptor() ([]byte, []int) {
	return file_eth_v1alpha1_attestation_proto_rawDescGZIP(), []int{4}
}

func (x *Checkpoint) GetEpoch() uint64 {
	if x != nil {
		return x.Epoch
	}
	return 0
}

func (x *Checkpoint) GetRoot() []byte {
	if x != nil {
		return x.Root
	}
	return nil
}

var File_eth_v1alpha1_attestation_proto protoreflect.FileDescriptor

var file_eth_v1alpha1_attestation_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x15, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x01, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x65, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6d, 0x0a, 0x10, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x42, 0x42, 0xf2, 0xde, 0x1f, 0x0e, 0x73, 0x73, 0x7a, 0x2d, 0x6d, 0x61, 0x78, 0x3a, 0x22, 0x32,
	0x30, 0x34, 0x38, 0x22, 0xfa, 0xde, 0x1f, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73,
	0x2f, 0x67, 0x6f, 0x2d, 0x62, 0x69, 0x74, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x42, 0x69, 0x74,
	0x6c, 0x69, 0x73, 0x74, 0x52, 0x0f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x69, 0x74, 0x73, 0x12, 0x3a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x65,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x2f, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x42, 0x11, 0xf2, 0xde, 0x1f, 0x0d, 0x73, 0x73, 0x7a, 0x2d, 0x73, 0x69,
	0x7a, 0x65, 0x3a, 0x22, 0x39, 0x36, 0x22, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x22, 0xc7, 0x01, 0x0a, 0x1c, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x64, 0x50, 0x72,
	0x6f, 0x6f, 0x66, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x40,
	0x0a, 0x09, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x12, 0x3a, 0x0a, 0x0f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72,
	0x6f, 0x6f, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x11, 0xf2, 0xde, 0x1f, 0x0d, 0x73,
	0x73, 0x7a, 0x2d, 0x73, 0x69, 0x7a, 0x65, 0x3a, 0x22, 0x39, 0x36, 0x22, 0x52, 0x0e, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0xa4, 0x01, 0x0a,
	0x22, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x64, 0x50, 0x72,
	0x6f, 0x6f, 0x66, 0x12, 0x4d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e,
	0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x11, 0xf2, 0xde, 0x1f, 0x0d, 0x73, 0x73, 0x7a, 0x2d, 0x73,
	0x69, 0x7a, 0x65, 0x3a, 0x22, 0x39, 0x36, 0x22, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x22, 0x83, 0x02, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x3d, 0x0a, 0x11, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x5f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42,
	0x11, 0xf2, 0xde, 0x1f, 0x0d, 0x73, 0x73, 0x7a, 0x2d, 0x73, 0x69, 0x7a, 0x65, 0x3a, 0x22, 0x33,
	0x32, 0x22, 0x52, 0x0f, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x6f, 0x6f, 0x74, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x39,
	0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x78, 0x0a, 0x0a, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x2d, 0xfa, 0xde, 0x1f, 0x29, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63,
	0x6c, 0x61, 0x62, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x32, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x45, 0x70, 0x6f, 0x63, 0x68, 0x52, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x25, 0x0a, 0x04,
	0x72, 0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x11, 0xf2, 0xde, 0x1f, 0x0d,
	0x73, 0x73, 0x7a, 0x2d, 0x73, 0x69, 0x7a, 0x65, 0x3a, 0x22, 0x33, 0x32, 0x22, 0x52, 0x04, 0x72,
	0x6f, 0x6f, 0x74, 0x42, 0x97, 0x01, 0x0a, 0x19, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x42, 0x10, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f,
	0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x65, 0x74, 0x68,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x65, 0x74, 0x68, 0xaa, 0x02, 0x15,
	0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x45, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x15, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x5c, 0x45, 0x74, 0x68, 0x5c, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eth_v1alpha1_attestation_proto_rawDescOnce sync.Once
	file_eth_v1alpha1_attestation_proto_rawDescData = file_eth_v1alpha1_attestation_proto_rawDesc
)

func file_eth_v1alpha1_attestation_proto_rawDescGZIP() []byte {
	file_eth_v1alpha1_attestation_proto_rawDescOnce.Do(func() {
		file_eth_v1alpha1_attestation_proto_rawDescData = protoimpl.X.CompressGZIP(file_eth_v1alpha1_attestation_proto_rawDescData)
	})
	return file_eth_v1alpha1_attestation_proto_rawDescData
}

var file_eth_v1alpha1_attestation_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_eth_v1alpha1_attestation_proto_goTypes = []interface{}{
	(*Attestation)(nil),                        // 0: ethereum.eth.v1alpha1.Attestation
	(*AggregateAttestationAndProof)(nil),       // 1: ethereum.eth.v1alpha1.AggregateAttestationAndProof
	(*SignedAggregateAttestationAndProof)(nil), // 2: ethereum.eth.v1alpha1.SignedAggregateAttestationAndProof
	(*AttestationData)(nil),                    // 3: ethereum.eth.v1alpha1.AttestationData
	(*Checkpoint)(nil),                         // 4: ethereum.eth.v1alpha1.Checkpoint
}
var file_eth_v1alpha1_attestation_proto_depIdxs = []int32{
	3, // 0: ethereum.eth.v1alpha1.Attestation.data:type_name -> ethereum.eth.v1alpha1.AttestationData
	0, // 1: ethereum.eth.v1alpha1.AggregateAttestationAndProof.aggregate:type_name -> ethereum.eth.v1alpha1.Attestation
	1, // 2: ethereum.eth.v1alpha1.SignedAggregateAttestationAndProof.message:type_name -> ethereum.eth.v1alpha1.AggregateAttestationAndProof
	4, // 3: ethereum.eth.v1alpha1.AttestationData.source:type_name -> ethereum.eth.v1alpha1.Checkpoint
	4, // 4: ethereum.eth.v1alpha1.AttestationData.target:type_name -> ethereum.eth.v1alpha1.Checkpoint
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_eth_v1alpha1_attestation_proto_init() }
func file_eth_v1alpha1_attestation_proto_init() {
	if File_eth_v1alpha1_attestation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eth_v1alpha1_attestation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attestation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eth_v1alpha1_attestation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AggregateAttestationAndProof); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eth_v1alpha1_attestation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedAggregateAttestationAndProof); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eth_v1alpha1_attestation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestationData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eth_v1alpha1_attestation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Checkpoint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eth_v1alpha1_attestation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eth_v1alpha1_attestation_proto_goTypes,
		DependencyIndexes: file_eth_v1alpha1_attestation_proto_depIdxs,
		MessageInfos:      file_eth_v1alpha1_attestation_proto_msgTypes,
	}.Build()
	File_eth_v1alpha1_attestation_proto = out.File
	file_eth_v1alpha1_attestation_proto_rawDesc = nil
	file_eth_v1alpha1_attestation_proto_goTypes = nil
	file_eth_v1alpha1_attestation_proto_depIdxs = nil
}
