// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eth/v1/beacon_debug.proto

package v1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ForkChoiceHeadsResponse struct {
	Data                 []*ForkChoiceHead `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ForkChoiceHeadsResponse) Reset()         { *m = ForkChoiceHeadsResponse{} }
func (m *ForkChoiceHeadsResponse) String() string { return proto.CompactTextString(m) }
func (*ForkChoiceHeadsResponse) ProtoMessage()    {}
func (*ForkChoiceHeadsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2b2a594eaf2d4f, []int{0}
}
func (m *ForkChoiceHeadsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ForkChoiceHeadsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ForkChoiceHeadsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ForkChoiceHeadsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForkChoiceHeadsResponse.Merge(m, src)
}
func (m *ForkChoiceHeadsResponse) XXX_Size() int {
	return m.Size()
}
func (m *ForkChoiceHeadsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ForkChoiceHeadsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ForkChoiceHeadsResponse proto.InternalMessageInfo

func (m *ForkChoiceHeadsResponse) GetData() []*ForkChoiceHead {
	if m != nil {
		return m.Data
	}
	return nil
}

type ForkChoiceHead struct {
	Root                 []byte   `protobuf:"bytes,1,opt,name=root,proto3" json:"root,omitempty" ssz-size:"32"`
	Slot                 uint64   `protobuf:"varint,2,opt,name=slot,proto3" json:"slot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForkChoiceHead) Reset()         { *m = ForkChoiceHead{} }
func (m *ForkChoiceHead) String() string { return proto.CompactTextString(m) }
func (*ForkChoiceHead) ProtoMessage()    {}
func (*ForkChoiceHead) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2b2a594eaf2d4f, []int{1}
}
func (m *ForkChoiceHead) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ForkChoiceHead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ForkChoiceHead.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ForkChoiceHead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForkChoiceHead.Merge(m, src)
}
func (m *ForkChoiceHead) XXX_Size() int {
	return m.Size()
}
func (m *ForkChoiceHead) XXX_DiscardUnknown() {
	xxx_messageInfo_ForkChoiceHead.DiscardUnknown(m)
}

var xxx_messageInfo_ForkChoiceHead proto.InternalMessageInfo

func (m *ForkChoiceHead) GetRoot() []byte {
	if m != nil {
		return m.Root
	}
	return nil
}

func (m *ForkChoiceHead) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

type BeaconStateResponse struct {
	Data                 *BeaconState `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BeaconStateResponse) Reset()         { *m = BeaconStateResponse{} }
func (m *BeaconStateResponse) String() string { return proto.CompactTextString(m) }
func (*BeaconStateResponse) ProtoMessage()    {}
func (*BeaconStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2b2a594eaf2d4f, []int{2}
}
func (m *BeaconStateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BeaconStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BeaconStateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BeaconStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeaconStateResponse.Merge(m, src)
}
func (m *BeaconStateResponse) XXX_Size() int {
	return m.Size()
}
func (m *BeaconStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BeaconStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BeaconStateResponse proto.InternalMessageInfo

func (m *BeaconStateResponse) GetData() *BeaconState {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ForkChoiceHeadsResponse)(nil), "ethereum.eth.v1.ForkChoiceHeadsResponse")
	proto.RegisterType((*ForkChoiceHead)(nil), "ethereum.eth.v1.ForkChoiceHead")
	proto.RegisterType((*BeaconStateResponse)(nil), "ethereum.eth.v1.BeaconStateResponse")
}

func init() { proto.RegisterFile("eth/v1/beacon_debug.proto", fileDescriptor_7b2b2a594eaf2d4f) }

var fileDescriptor_7b2b2a594eaf2d4f = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xcf, 0x8a, 0xd3, 0x40,
	0x18, 0x67, 0x6a, 0xf1, 0x30, 0xab, 0xbb, 0x38, 0x85, 0xb5, 0xc6, 0xda, 0xad, 0x41, 0x25, 0x07,
	0x77, 0xc6, 0xb6, 0x37, 0x8f, 0xd5, 0xba, 0x82, 0x22, 0x4b, 0xbd, 0xc9, 0x42, 0x99, 0x24, 0x63,
	0x32, 0x6c, 0x9a, 0x89, 0x99, 0x2f, 0x81, 0xdd, 0xc5, 0x83, 0x82, 0x2f, 0xa0, 0x6f, 0xe1, 0x93,
	0x88, 0x27, 0xc1, 0xbb, 0x48, 0xf1, 0x09, 0x7c, 0x02, 0xc9, 0x64, 0x5a, 0xdb, 0xee, 0xa2, 0xa7,
	0xcc, 0x7c, 0xbf, 0xdf, 0xef, 0x9b, 0xdf, 0xf7, 0x27, 0xf8, 0x86, 0x80, 0x98, 0x95, 0x7d, 0xe6,
	0x0b, 0x1e, 0xa8, 0x74, 0x1a, 0x0a, 0xbf, 0x88, 0x68, 0x96, 0x2b, 0x50, 0x64, 0x47, 0x40, 0x2c,
	0x72, 0x51, 0xcc, 0xa8, 0x80, 0x98, 0x96, 0x7d, 0x67, 0x3f, 0x92, 0x10, 0x17, 0x3e, 0x0d, 0xd4,
	0x8c, 0x45, 0x2a, 0x52, 0xcc, 0xf0, 0xfc, 0xe2, 0xb5, 0xb9, 0x99, 0x8b, 0x39, 0xd5, 0x7a, 0xa7,
	0x13, 0x29, 0x15, 0x25, 0x82, 0xf1, 0x4c, 0x32, 0x9e, 0xa6, 0x0a, 0x38, 0x48, 0x95, 0x6a, 0x8b,
	0xde, 0xb4, 0xe8, 0x32, 0x87, 0x98, 0x65, 0x70, 0x62, 0xc1, 0xb6, 0x75, 0xc5, 0x01, 0x84, 0xae,
	0x75, 0x16, 0xd9, 0xf0, 0xeb, 0x27, 0x2a, 0x38, 0xbe, 0x18, 0xaa, 0x74, 0xc2, 0x42, 0xb7, 0xd7,
	0xa1, 0x20, 0xe6, 0x32, 0x9d, 0x6a, 0x91, 0x97, 0x32, 0x58, 0x50, 0x76, 0x2d, 0xa5, 0xe4, 0x89,
	0x0c, 0x39, 0xa8, 0xbc, 0x8e, 0xbb, 0x2f, 0xf0, 0xf5, 0x27, 0x2a, 0x3f, 0x7e, 0x14, 0x2b, 0x19,
	0x88, 0xa7, 0x82, 0x87, 0x7a, 0x22, 0x74, 0xa6, 0x52, 0x2d, 0xc8, 0x10, 0x37, 0x43, 0x0e, 0xbc,
	0x8d, 0x7a, 0x97, 0xbc, 0xad, 0xc1, 0x1e, 0xdd, 0xe8, 0x17, 0x5d, 0xd7, 0x4d, 0x0c, 0xd9, 0x7d,
	0x86, 0xb7, 0xd7, 0xe3, 0xe4, 0x2e, 0x6e, 0xe6, 0x4a, 0x41, 0x1b, 0xf5, 0x90, 0x77, 0x65, 0x74,
	0xed, 0xf7, 0x8f, 0xbd, 0xab, 0x5a, 0x9f, 0xee, 0x6b, 0x79, 0x2a, 0x1e, 0xba, 0xc3, 0x81, 0x3b,
	0x31, 0x30, 0x21, 0xb8, 0xa9, 0x13, 0x05, 0xed, 0x46, 0x0f, 0x79, 0xcd, 0x89, 0x39, 0xbb, 0x07,
	0xb8, 0x35, 0x32, 0x25, 0xbd, 0xac, 0x8a, 0x5d, 0x1a, 0x7b, 0xb0, 0x34, 0x86, 0xbc, 0xad, 0x41,
	0xe7, 0x9c, 0xb1, 0x55, 0x8d, 0x61, 0x0e, 0x3e, 0x36, 0xf0, 0x56, 0x1d, 0x7d, 0x5c, 0x6d, 0x00,
	0xf9, 0x80, 0xf0, 0xf6, 0x81, 0x80, 0x15, 0x22, 0xb9, 0x75, 0x2e, 0x8d, 0x7d, 0xf4, 0x4d, 0x21,
	0x34, 0x38, 0x77, 0xfe, 0xf9, 0x8a, 0x75, 0xe6, 0xd2, 0xf7, 0xdf, 0x7f, 0x7d, 0x6a, 0x78, 0xe4,
	0x1e, 0xb3, 0xed, 0x36, 0x0b, 0x67, 0xe7, 0xc2, 0xcc, 0xc8, 0x34, 0x3b, 0x33, 0xdf, 0xa9, 0x0c,
	0xdf, 0x92, 0x33, 0xdc, 0x7a, 0x2e, 0x35, 0x6c, 0x4c, 0x80, 0xec, 0xd2, 0x7a, 0x7b, 0xe8, 0x62,
	0x7b, 0xe8, 0xb8, 0xda, 0x1e, 0xc7, 0xfb, 0xcf, 0x0c, 0x96, 0xb3, 0x73, 0x5d, 0x63, 0xa4, 0x43,
	0x9c, 0x0b, 0x8d, 0xc4, 0x15, 0x77, 0xf4, 0x0e, 0x7d, 0x99, 0x77, 0xd1, 0xb7, 0x79, 0x17, 0xfd,
	0x9c, 0x77, 0x11, 0x6e, 0xa9, 0x3c, 0xda, 0xcc, 0x3f, 0xc2, 0xa6, 0x5f, 0x87, 0x95, 0x8d, 0x43,
	0xf4, 0xea, 0xfe, 0xca, 0x1f, 0x92, 0xe5, 0x27, 0x7a, 0xc6, 0x41, 0x06, 0x09, 0xf7, 0x35, 0x5b,
	0xe8, 0x78, 0x26, 0xb5, 0x7d, 0xf2, 0x73, 0x63, 0x67, 0xbc, 0xc8, 0x36, 0x36, 0xd9, 0xbe, 0xfe,
	0x8d, 0x1c, 0x8d, 0x21, 0x3e, 0x2a, 0xfb, 0xfe, 0x65, 0x53, 0xe1, 0xf0, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x7d, 0x2e, 0xb8, 0x94, 0xa8, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BeaconDebugClient is the client API for BeaconDebug service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BeaconDebugClient interface {
	GetBeaconState(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (*BeaconStateResponse, error)
	ListForkChoiceHeads(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*ForkChoiceHeadsResponse, error)
}

type beaconDebugClient struct {
	cc *grpc.ClientConn
}

func NewBeaconDebugClient(cc *grpc.ClientConn) BeaconDebugClient {
	return &beaconDebugClient{cc}
}

func (c *beaconDebugClient) GetBeaconState(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (*BeaconStateResponse, error) {
	out := new(BeaconStateResponse)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1.BeaconDebug/GetBeaconState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beaconDebugClient) ListForkChoiceHeads(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*ForkChoiceHeadsResponse, error) {
	out := new(ForkChoiceHeadsResponse)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1.BeaconDebug/ListForkChoiceHeads", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BeaconDebugServer is the server API for BeaconDebug service.
type BeaconDebugServer interface {
	GetBeaconState(context.Context, *StateRequest) (*BeaconStateResponse, error)
	ListForkChoiceHeads(context.Context, *types.Empty) (*ForkChoiceHeadsResponse, error)
}

// UnimplementedBeaconDebugServer can be embedded to have forward compatible implementations.
type UnimplementedBeaconDebugServer struct {
}

func (*UnimplementedBeaconDebugServer) GetBeaconState(ctx context.Context, req *StateRequest) (*BeaconStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBeaconState not implemented")
}
func (*UnimplementedBeaconDebugServer) ListForkChoiceHeads(ctx context.Context, req *types.Empty) (*ForkChoiceHeadsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListForkChoiceHeads not implemented")
}

func RegisterBeaconDebugServer(s *grpc.Server, srv BeaconDebugServer) {
	s.RegisterService(&_BeaconDebug_serviceDesc, srv)
}

func _BeaconDebug_GetBeaconState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconDebugServer).GetBeaconState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1.BeaconDebug/GetBeaconState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconDebugServer).GetBeaconState(ctx, req.(*StateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BeaconDebug_ListForkChoiceHeads_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconDebugServer).ListForkChoiceHeads(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1.BeaconDebug/ListForkChoiceHeads",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconDebugServer).ListForkChoiceHeads(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _BeaconDebug_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ethereum.eth.v1.BeaconDebug",
	HandlerType: (*BeaconDebugServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBeaconState",
			Handler:    _BeaconDebug_GetBeaconState_Handler,
		},
		{
			MethodName: "ListForkChoiceHeads",
			Handler:    _BeaconDebug_ListForkChoiceHeads_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eth/v1/beacon_debug.proto",
}

func (m *ForkChoiceHeadsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ForkChoiceHeadsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ForkChoiceHeadsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Data) > 0 {
		for iNdEx := len(m.Data) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Data[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBeaconDebug(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ForkChoiceHead) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ForkChoiceHead) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ForkChoiceHead) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Slot != 0 {
		i = encodeVarintBeaconDebug(dAtA, i, uint64(m.Slot))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Root) > 0 {
		i -= len(m.Root)
		copy(dAtA[i:], m.Root)
		i = encodeVarintBeaconDebug(dAtA, i, uint64(len(m.Root)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BeaconStateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BeaconStateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BeaconStateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBeaconDebug(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBeaconDebug(dAtA []byte, offset int, v uint64) int {
	offset -= sovBeaconDebug(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ForkChoiceHeadsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Data) > 0 {
		for _, e := range m.Data {
			l = e.Size()
			n += 1 + l + sovBeaconDebug(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ForkChoiceHead) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Root)
	if l > 0 {
		n += 1 + l + sovBeaconDebug(uint64(l))
	}
	if m.Slot != 0 {
		n += 1 + sovBeaconDebug(uint64(m.Slot))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BeaconStateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovBeaconDebug(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovBeaconDebug(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBeaconDebug(x uint64) (n int) {
	return sovBeaconDebug(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ForkChoiceHeadsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBeaconDebug
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
			return fmt.Errorf("proto: ForkChoiceHeadsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ForkChoiceHeadsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBeaconDebug
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
				return ErrInvalidLengthBeaconDebug
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBeaconDebug
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, &ForkChoiceHead{})
			if err := m.Data[len(m.Data)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBeaconDebug(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBeaconDebug
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBeaconDebug
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ForkChoiceHead) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBeaconDebug
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
			return fmt.Errorf("proto: ForkChoiceHead: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ForkChoiceHead: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Root", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBeaconDebug
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBeaconDebug
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBeaconDebug
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Root = append(m.Root[:0], dAtA[iNdEx:postIndex]...)
			if m.Root == nil {
				m.Root = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slot", wireType)
			}
			m.Slot = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBeaconDebug
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Slot |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBeaconDebug(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBeaconDebug
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBeaconDebug
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BeaconStateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBeaconDebug
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
			return fmt.Errorf("proto: BeaconStateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BeaconStateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBeaconDebug
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
				return ErrInvalidLengthBeaconDebug
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBeaconDebug
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &BeaconState{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBeaconDebug(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBeaconDebug
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBeaconDebug
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBeaconDebug(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBeaconDebug
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
					return 0, ErrIntOverflowBeaconDebug
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
					return 0, ErrIntOverflowBeaconDebug
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
				return 0, ErrInvalidLengthBeaconDebug
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBeaconDebug
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBeaconDebug
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBeaconDebug        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBeaconDebug          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBeaconDebug = fmt.Errorf("proto: unexpected end of group")
)