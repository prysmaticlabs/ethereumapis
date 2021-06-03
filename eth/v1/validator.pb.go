// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.8
// source: eth/v1/validator.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/protoc-gen-go/descriptor"
	github_com_prysmaticlabs_eth2_types "github.com/prysmaticlabs/eth2-types"
	_ "github.com/prysmaticlabs/ethereumapis/eth/ext"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
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

type ValidatorStatus int32

const (
	ValidatorStatus_PENDING_INITIALIZED ValidatorStatus = 0
	ValidatorStatus_PENDING_QUEUED      ValidatorStatus = 1
	ValidatorStatus_ACTIVE_ONGOING      ValidatorStatus = 2
	ValidatorStatus_ACTIVE_EXITING      ValidatorStatus = 3
	ValidatorStatus_ACTIVE_SLASHED      ValidatorStatus = 4
	ValidatorStatus_EXITED_UNSLASHED    ValidatorStatus = 5
	ValidatorStatus_EXITED_SLASHED      ValidatorStatus = 6
	ValidatorStatus_WITHDRAWAL_POSSIBLE ValidatorStatus = 7
	ValidatorStatus_WITHDRAWAL_DONE     ValidatorStatus = 8
	ValidatorStatus_ACTIVE              ValidatorStatus = 9
	ValidatorStatus_PENDING             ValidatorStatus = 10
	ValidatorStatus_EXITED              ValidatorStatus = 11
	ValidatorStatus_WITHDRAWAL          ValidatorStatus = 12
)

// Enum value maps for ValidatorStatus.
var (
	ValidatorStatus_name = map[int32]string{
		0:  "PENDING_INITIALIZED",
		1:  "PENDING_QUEUED",
		2:  "ACTIVE_ONGOING",
		3:  "ACTIVE_EXITING",
		4:  "ACTIVE_SLASHED",
		5:  "EXITED_UNSLASHED",
		6:  "EXITED_SLASHED",
		7:  "WITHDRAWAL_POSSIBLE",
		8:  "WITHDRAWAL_DONE",
		9:  "ACTIVE",
		10: "PENDING",
		11: "EXITED",
		12: "WITHDRAWAL",
	}
	ValidatorStatus_value = map[string]int32{
		"PENDING_INITIALIZED": 0,
		"PENDING_QUEUED":      1,
		"ACTIVE_ONGOING":      2,
		"ACTIVE_EXITING":      3,
		"ACTIVE_SLASHED":      4,
		"EXITED_UNSLASHED":    5,
		"EXITED_SLASHED":      6,
		"WITHDRAWAL_POSSIBLE": 7,
		"WITHDRAWAL_DONE":     8,
		"ACTIVE":              9,
		"PENDING":             10,
		"EXITED":              11,
		"WITHDRAWAL":          12,
	}
)

func (x ValidatorStatus) Enum() *ValidatorStatus {
	p := new(ValidatorStatus)
	*p = x
	return p
}

func (x ValidatorStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ValidatorStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_eth_v1_validator_proto_enumTypes[0].Descriptor()
}

func (ValidatorStatus) Type() protoreflect.EnumType {
	return &file_proto_eth_v1_validator_proto_enumTypes[0]
}

func (x ValidatorStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ValidatorStatus.Descriptor instead.
func (ValidatorStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_eth_v1_validator_proto_rawDescGZIP(), []int{0}
}

type ValidatorContainer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index     github_com_prysmaticlabs_eth2_types.ValidatorIndex `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.ValidatorIndex"`
	Balance   uint64                                             `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
	Status    ValidatorStatus                                    `protobuf:"varint,3,opt,name=status,proto3,enum=ethereum.eth.v1.ValidatorStatus" json:"status,omitempty"`
	Validator *Validator                                         `protobuf:"bytes,4,opt,name=validator,proto3" json:"validator,omitempty"`
}

func (x *ValidatorContainer) Reset() {
	*x = ValidatorContainer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eth_v1_validator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorContainer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorContainer) ProtoMessage() {}

func (x *ValidatorContainer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eth_v1_validator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatorContainer.ProtoReflect.Descriptor instead.
func (*ValidatorContainer) Descriptor() ([]byte, []int) {
	return file_proto_eth_v1_validator_proto_rawDescGZIP(), []int{0}
}

func (x *ValidatorContainer) GetIndex() github_com_prysmaticlabs_eth2_types.ValidatorIndex {
	if x != nil {
		return x.Index
	}
	return github_com_prysmaticlabs_eth2_types.ValidatorIndex(0)
}

func (x *ValidatorContainer) GetBalance() uint64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *ValidatorContainer) GetStatus() ValidatorStatus {
	if x != nil {
		return x.Status
	}
	return ValidatorStatus_PENDING_INITIALIZED
}

func (x *ValidatorContainer) GetValidator() *Validator {
	if x != nil {
		return x.Validator
	}
	return nil
}

type Validator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pubkey                     []byte                                    `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty" spec-name:"pubkey" ssz-size:"48"`
	WithdrawalCredentials      []byte                                    `protobuf:"bytes,2,opt,name=withdrawal_credentials,json=withdrawalCredentials,proto3" json:"withdrawal_credentials,omitempty" ssz-size:"32"`
	EffectiveBalance           uint64                                    `protobuf:"varint,3,opt,name=effective_balance,json=effectiveBalance,proto3" json:"effective_balance,omitempty"`
	Slashed                    bool                                      `protobuf:"varint,4,opt,name=slashed,proto3" json:"slashed,omitempty"`
	ActivationEligibilityEpoch github_com_prysmaticlabs_eth2_types.Epoch `protobuf:"varint,5,opt,name=activation_eligibility_epoch,json=activationEligibilityEpoch,proto3" json:"activation_eligibility_epoch,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.Epoch"`
	ActivationEpoch            github_com_prysmaticlabs_eth2_types.Epoch `protobuf:"varint,6,opt,name=activation_epoch,json=activationEpoch,proto3" json:"activation_epoch,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.Epoch"`
	ExitEpoch                  github_com_prysmaticlabs_eth2_types.Epoch `protobuf:"varint,7,opt,name=exit_epoch,json=exitEpoch,proto3" json:"exit_epoch,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.Epoch"`
	WithdrawableEpoch          github_com_prysmaticlabs_eth2_types.Epoch `protobuf:"varint,8,opt,name=withdrawable_epoch,json=withdrawableEpoch,proto3" json:"withdrawable_epoch,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.Epoch"`
}

func (x *Validator) Reset() {
	*x = Validator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eth_v1_validator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Validator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Validator) ProtoMessage() {}

func (x *Validator) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eth_v1_validator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Validator.ProtoReflect.Descriptor instead.
func (*Validator) Descriptor() ([]byte, []int) {
	return file_proto_eth_v1_validator_proto_rawDescGZIP(), []int{1}
}

func (x *Validator) GetPubkey() []byte {
	if x != nil {
		return x.Pubkey
	}
	return nil
}

func (x *Validator) GetWithdrawalCredentials() []byte {
	if x != nil {
		return x.WithdrawalCredentials
	}
	return nil
}

func (x *Validator) GetEffectiveBalance() uint64 {
	if x != nil {
		return x.EffectiveBalance
	}
	return 0
}

func (x *Validator) GetSlashed() bool {
	if x != nil {
		return x.Slashed
	}
	return false
}

func (x *Validator) GetActivationEligibilityEpoch() github_com_prysmaticlabs_eth2_types.Epoch {
	if x != nil {
		return x.ActivationEligibilityEpoch
	}
	return github_com_prysmaticlabs_eth2_types.Epoch(0)
}

func (x *Validator) GetActivationEpoch() github_com_prysmaticlabs_eth2_types.Epoch {
	if x != nil {
		return x.ActivationEpoch
	}
	return github_com_prysmaticlabs_eth2_types.Epoch(0)
}

func (x *Validator) GetExitEpoch() github_com_prysmaticlabs_eth2_types.Epoch {
	if x != nil {
		return x.ExitEpoch
	}
	return github_com_prysmaticlabs_eth2_types.Epoch(0)
}

func (x *Validator) GetWithdrawableEpoch() github_com_prysmaticlabs_eth2_types.Epoch {
	if x != nil {
		return x.WithdrawableEpoch
	}
	return github_com_prysmaticlabs_eth2_types.Epoch(0)
}

var File_proto_eth_v1_validator_proto protoreflect.FileDescriptor

var file_proto_eth_v1_validator_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x65, 0x78, 0x74,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0,
	0x01, 0x0a, 0x12, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x4c, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x36, 0x82, 0xb5, 0x18, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61,
	0x62, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x32, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x38, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e,
	0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x22, 0xb2, 0x04, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x28, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42,
	0x10, 0x8a, 0xb5, 0x18, 0x02, 0x34, 0x38, 0x9a, 0xb5, 0x18, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65,
	0x79, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x12, 0x3d, 0x0a, 0x16, 0x77, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33,
	0x32, 0x52, 0x15, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x65, 0x66, 0x66, 0x65,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x10, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x65, 0x64, 0x12,
	0x6f, 0x0a, 0x1c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6c,
	0x69, 0x67, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x04, 0x42, 0x2d, 0x82, 0xb5, 0x18, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c,
	0x61, 0x62, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x32, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45,
	0x70, 0x6f, 0x63, 0x68, 0x52, 0x1a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x6c, 0x69, 0x67, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x45, 0x70, 0x6f, 0x63, 0x68,
	0x12, 0x58, 0x0a, 0x10, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65,
	0x70, 0x6f, 0x63, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x42, 0x2d, 0x82, 0xb5, 0x18, 0x29,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d,
	0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x32, 0x2d, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x52, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x4c, 0x0a, 0x0a, 0x65, 0x78,
	0x69, 0x74, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x42, 0x2d,
	0x82, 0xb5, 0x18, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x65, 0x74, 0x68,
	0x32, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x52, 0x09, 0x65,
	0x78, 0x69, 0x74, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x5c, 0x0a, 0x12, 0x77, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x2d, 0x82, 0xb5, 0x18, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61,
	0x62, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x32, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x70,
	0x6f, 0x63, 0x68, 0x52, 0x11, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x62, 0x6c,
	0x65, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x2a, 0x87, 0x02, 0x0a, 0x0f, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x45,
	0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x49, 0x5a, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x51,
	0x55, 0x45, 0x55, 0x45, 0x44, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x43, 0x54, 0x49, 0x56,
	0x45, 0x5f, 0x4f, 0x4e, 0x47, 0x4f, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x41,
	0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x45, 0x58, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12,
	0x12, 0x0a, 0x0e, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x53, 0x4c, 0x41, 0x53, 0x48, 0x45,
	0x44, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x58, 0x49, 0x54, 0x45, 0x44, 0x5f, 0x55, 0x4e,
	0x53, 0x4c, 0x41, 0x53, 0x48, 0x45, 0x44, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x58, 0x49,
	0x54, 0x45, 0x44, 0x5f, 0x53, 0x4c, 0x41, 0x53, 0x48, 0x45, 0x44, 0x10, 0x06, 0x12, 0x17, 0x0a,
	0x13, 0x57, 0x49, 0x54, 0x48, 0x44, 0x52, 0x41, 0x57, 0x41, 0x4c, 0x5f, 0x50, 0x4f, 0x53, 0x53,
	0x49, 0x42, 0x4c, 0x45, 0x10, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x57, 0x49, 0x54, 0x48, 0x44, 0x52,
	0x41, 0x57, 0x41, 0x4c, 0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x41,
	0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x09, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49,
	0x4e, 0x47, 0x10, 0x0a, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x58, 0x49, 0x54, 0x45, 0x44, 0x10, 0x0b,
	0x12, 0x0e, 0x0a, 0x0a, 0x57, 0x49, 0x54, 0x48, 0x44, 0x52, 0x41, 0x57, 0x41, 0x4c, 0x10, 0x0c,
	0x42, 0x79, 0x0a, 0x13, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c,
	0x61, 0x62, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0xaa, 0x02, 0x0f, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x2e, 0x45, 0x74, 0x68, 0x2e, 0x76, 0x31, 0xca, 0x02, 0x0f, 0x45, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x5c, 0x45, 0x74, 0x68, 0x5c, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_eth_v1_validator_proto_rawDescOnce sync.Once
	file_proto_eth_v1_validator_proto_rawDescData = file_proto_eth_v1_validator_proto_rawDesc
)

func file_proto_eth_v1_validator_proto_rawDescGZIP() []byte {
	file_proto_eth_v1_validator_proto_rawDescOnce.Do(func() {
		file_proto_eth_v1_validator_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_eth_v1_validator_proto_rawDescData)
	})
	return file_proto_eth_v1_validator_proto_rawDescData
}

var file_proto_eth_v1_validator_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_eth_v1_validator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_eth_v1_validator_proto_goTypes = []interface{}{
	(ValidatorStatus)(0),       // 0: ethereum.eth.v1.ValidatorStatus
	(*ValidatorContainer)(nil), // 1: ethereum.eth.v1.ValidatorContainer
	(*Validator)(nil),          // 2: ethereum.eth.v1.Validator
}
var file_proto_eth_v1_validator_proto_depIdxs = []int32{
	0, // 0: ethereum.eth.v1.ValidatorContainer.status:type_name -> ethereum.eth.v1.ValidatorStatus
	2, // 1: ethereum.eth.v1.ValidatorContainer.validator:type_name -> ethereum.eth.v1.Validator
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_eth_v1_validator_proto_init() }
func file_proto_eth_v1_validator_proto_init() {
	if File_proto_eth_v1_validator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_eth_v1_validator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorContainer); i {
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
		file_proto_eth_v1_validator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Validator); i {
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
			RawDescriptor: file_proto_eth_v1_validator_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_eth_v1_validator_proto_goTypes,
		DependencyIndexes: file_proto_eth_v1_validator_proto_depIdxs,
		EnumInfos:         file_proto_eth_v1_validator_proto_enumTypes,
		MessageInfos:      file_proto_eth_v1_validator_proto_msgTypes,
	}.Build()
	File_proto_eth_v1_validator_proto = out.File
	file_proto_eth_v1_validator_proto_rawDesc = nil
	file_proto_eth_v1_validator_proto_goTypes = nil
	file_proto_eth_v1_validator_proto_depIdxs = nil
}
