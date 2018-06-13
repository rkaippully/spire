// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dummy.proto

package dummy

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/spiffe/spire/proto/common"
import plugin "github.com/spiffe/spire/proto/common/plugin"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Empty from public import github.com/spiffe/spire/proto/common/common.proto
type Empty = common.Empty

// AttestationData from public import github.com/spiffe/spire/proto/common/common.proto
type AttestationData = common.AttestationData

// Selector from public import github.com/spiffe/spire/proto/common/common.proto
type Selector = common.Selector

// Selectors from public import github.com/spiffe/spire/proto/common/common.proto
type Selectors = common.Selectors

// RegistrationEntry from public import github.com/spiffe/spire/proto/common/common.proto
type RegistrationEntry = common.RegistrationEntry

// RegistrationEntries from public import github.com/spiffe/spire/proto/common/common.proto
type RegistrationEntries = common.RegistrationEntries

// ConfigureRequest from public import github.com/spiffe/spire/proto/common/plugin/plugin.proto
type ConfigureRequest = plugin.ConfigureRequest

// ConfigureResponse from public import github.com/spiffe/spire/proto/common/plugin/plugin.proto
type ConfigureResponse = plugin.ConfigureResponse

// GetPluginInfoRequest from public import github.com/spiffe/spire/proto/common/plugin/plugin.proto
type GetPluginInfoRequest = plugin.GetPluginInfoRequest

// GetPluginInfoResponse from public import github.com/spiffe/spire/proto/common/plugin/plugin.proto
type GetPluginInfoResponse = plugin.GetPluginInfoResponse

type NoStreamRequest struct {
	Value                int64    `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NoStreamRequest) Reset()         { *m = NoStreamRequest{} }
func (m *NoStreamRequest) String() string { return proto.CompactTextString(m) }
func (*NoStreamRequest) ProtoMessage()    {}
func (*NoStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dummy_50641bf2c7f7ff0f, []int{0}
}
func (m *NoStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NoStreamRequest.Unmarshal(m, b)
}
func (m *NoStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NoStreamRequest.Marshal(b, m, deterministic)
}
func (dst *NoStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoStreamRequest.Merge(dst, src)
}
func (m *NoStreamRequest) XXX_Size() int {
	return xxx_messageInfo_NoStreamRequest.Size(m)
}
func (m *NoStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NoStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NoStreamRequest proto.InternalMessageInfo

func (m *NoStreamRequest) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type NoStreamResponse struct {
	Value                int64    `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NoStreamResponse) Reset()         { *m = NoStreamResponse{} }
func (m *NoStreamResponse) String() string { return proto.CompactTextString(m) }
func (*NoStreamResponse) ProtoMessage()    {}
func (*NoStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dummy_50641bf2c7f7ff0f, []int{1}
}
func (m *NoStreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NoStreamResponse.Unmarshal(m, b)
}
func (m *NoStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NoStreamResponse.Marshal(b, m, deterministic)
}
func (dst *NoStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoStreamResponse.Merge(dst, src)
}
func (m *NoStreamResponse) XXX_Size() int {
	return xxx_messageInfo_NoStreamResponse.Size(m)
}
func (m *NoStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NoStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NoStreamResponse proto.InternalMessageInfo

func (m *NoStreamResponse) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type ClientStreamRequest struct {
	Value                int64    `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientStreamRequest) Reset()         { *m = ClientStreamRequest{} }
func (m *ClientStreamRequest) String() string { return proto.CompactTextString(m) }
func (*ClientStreamRequest) ProtoMessage()    {}
func (*ClientStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dummy_50641bf2c7f7ff0f, []int{2}
}
func (m *ClientStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientStreamRequest.Unmarshal(m, b)
}
func (m *ClientStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientStreamRequest.Marshal(b, m, deterministic)
}
func (dst *ClientStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientStreamRequest.Merge(dst, src)
}
func (m *ClientStreamRequest) XXX_Size() int {
	return xxx_messageInfo_ClientStreamRequest.Size(m)
}
func (m *ClientStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientStreamRequest proto.InternalMessageInfo

func (m *ClientStreamRequest) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type ClientStreamResponse struct {
	Value                int64    `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientStreamResponse) Reset()         { *m = ClientStreamResponse{} }
func (m *ClientStreamResponse) String() string { return proto.CompactTextString(m) }
func (*ClientStreamResponse) ProtoMessage()    {}
func (*ClientStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dummy_50641bf2c7f7ff0f, []int{3}
}
func (m *ClientStreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientStreamResponse.Unmarshal(m, b)
}
func (m *ClientStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientStreamResponse.Marshal(b, m, deterministic)
}
func (dst *ClientStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientStreamResponse.Merge(dst, src)
}
func (m *ClientStreamResponse) XXX_Size() int {
	return xxx_messageInfo_ClientStreamResponse.Size(m)
}
func (m *ClientStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClientStreamResponse proto.InternalMessageInfo

func (m *ClientStreamResponse) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type ServerStreamRequest struct {
	Value                int64    `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerStreamRequest) Reset()         { *m = ServerStreamRequest{} }
func (m *ServerStreamRequest) String() string { return proto.CompactTextString(m) }
func (*ServerStreamRequest) ProtoMessage()    {}
func (*ServerStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dummy_50641bf2c7f7ff0f, []int{4}
}
func (m *ServerStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerStreamRequest.Unmarshal(m, b)
}
func (m *ServerStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerStreamRequest.Marshal(b, m, deterministic)
}
func (dst *ServerStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerStreamRequest.Merge(dst, src)
}
func (m *ServerStreamRequest) XXX_Size() int {
	return xxx_messageInfo_ServerStreamRequest.Size(m)
}
func (m *ServerStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServerStreamRequest proto.InternalMessageInfo

func (m *ServerStreamRequest) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type ServerStreamResponse struct {
	Value                int64    `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerStreamResponse) Reset()         { *m = ServerStreamResponse{} }
func (m *ServerStreamResponse) String() string { return proto.CompactTextString(m) }
func (*ServerStreamResponse) ProtoMessage()    {}
func (*ServerStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dummy_50641bf2c7f7ff0f, []int{5}
}
func (m *ServerStreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerStreamResponse.Unmarshal(m, b)
}
func (m *ServerStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerStreamResponse.Marshal(b, m, deterministic)
}
func (dst *ServerStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerStreamResponse.Merge(dst, src)
}
func (m *ServerStreamResponse) XXX_Size() int {
	return xxx_messageInfo_ServerStreamResponse.Size(m)
}
func (m *ServerStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServerStreamResponse proto.InternalMessageInfo

func (m *ServerStreamResponse) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type BothStreamRequest struct {
	Value                int64    `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BothStreamRequest) Reset()         { *m = BothStreamRequest{} }
func (m *BothStreamRequest) String() string { return proto.CompactTextString(m) }
func (*BothStreamRequest) ProtoMessage()    {}
func (*BothStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dummy_50641bf2c7f7ff0f, []int{6}
}
func (m *BothStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BothStreamRequest.Unmarshal(m, b)
}
func (m *BothStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BothStreamRequest.Marshal(b, m, deterministic)
}
func (dst *BothStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BothStreamRequest.Merge(dst, src)
}
func (m *BothStreamRequest) XXX_Size() int {
	return xxx_messageInfo_BothStreamRequest.Size(m)
}
func (m *BothStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BothStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BothStreamRequest proto.InternalMessageInfo

func (m *BothStreamRequest) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type BothStreamResponse struct {
	Value                int64    `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BothStreamResponse) Reset()         { *m = BothStreamResponse{} }
func (m *BothStreamResponse) String() string { return proto.CompactTextString(m) }
func (*BothStreamResponse) ProtoMessage()    {}
func (*BothStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dummy_50641bf2c7f7ff0f, []int{7}
}
func (m *BothStreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BothStreamResponse.Unmarshal(m, b)
}
func (m *BothStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BothStreamResponse.Marshal(b, m, deterministic)
}
func (dst *BothStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BothStreamResponse.Merge(dst, src)
}
func (m *BothStreamResponse) XXX_Size() int {
	return xxx_messageInfo_BothStreamResponse.Size(m)
}
func (m *BothStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BothStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BothStreamResponse proto.InternalMessageInfo

func (m *BothStreamResponse) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*NoStreamRequest)(nil), "spire.dummy.NoStreamRequest")
	proto.RegisterType((*NoStreamResponse)(nil), "spire.dummy.NoStreamResponse")
	proto.RegisterType((*ClientStreamRequest)(nil), "spire.dummy.ClientStreamRequest")
	proto.RegisterType((*ClientStreamResponse)(nil), "spire.dummy.ClientStreamResponse")
	proto.RegisterType((*ServerStreamRequest)(nil), "spire.dummy.ServerStreamRequest")
	proto.RegisterType((*ServerStreamResponse)(nil), "spire.dummy.ServerStreamResponse")
	proto.RegisterType((*BothStreamRequest)(nil), "spire.dummy.BothStreamRequest")
	proto.RegisterType((*BothStreamResponse)(nil), "spire.dummy.BothStreamResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Dummy service

type DummyClient interface {
	// No streaming
	NoStream(ctx context.Context, in *NoStreamRequest, opts ...grpc.CallOption) (*NoStreamResponse, error)
	// Stream things to
	ClientStream(ctx context.Context, opts ...grpc.CallOption) (Dummy_ClientStreamClient, error)
	// Stream things back
	ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...grpc.CallOption) (Dummy_ServerStreamClient, error)
	BothStream(ctx context.Context, opts ...grpc.CallOption) (Dummy_BothStreamClient, error)
	// Plugin Methods
	Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error)
	GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error)
}

type dummyClient struct {
	cc *grpc.ClientConn
}

func NewDummyClient(cc *grpc.ClientConn) DummyClient {
	return &dummyClient{cc}
}

func (c *dummyClient) NoStream(ctx context.Context, in *NoStreamRequest, opts ...grpc.CallOption) (*NoStreamResponse, error) {
	out := new(NoStreamResponse)
	err := grpc.Invoke(ctx, "/spire.dummy.Dummy/NoStream", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dummyClient) ClientStream(ctx context.Context, opts ...grpc.CallOption) (Dummy_ClientStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Dummy_serviceDesc.Streams[0], c.cc, "/spire.dummy.Dummy/ClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &dummyClientStreamClient{stream}
	return x, nil
}

type Dummy_ClientStreamClient interface {
	Send(*ClientStreamRequest) error
	CloseAndRecv() (*ClientStreamResponse, error)
	grpc.ClientStream
}

type dummyClientStreamClient struct {
	grpc.ClientStream
}

func (x *dummyClientStreamClient) Send(m *ClientStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dummyClientStreamClient) CloseAndRecv() (*ClientStreamResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ClientStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dummyClient) ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...grpc.CallOption) (Dummy_ServerStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Dummy_serviceDesc.Streams[1], c.cc, "/spire.dummy.Dummy/ServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &dummyServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Dummy_ServerStreamClient interface {
	Recv() (*ServerStreamResponse, error)
	grpc.ClientStream
}

type dummyServerStreamClient struct {
	grpc.ClientStream
}

func (x *dummyServerStreamClient) Recv() (*ServerStreamResponse, error) {
	m := new(ServerStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dummyClient) BothStream(ctx context.Context, opts ...grpc.CallOption) (Dummy_BothStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Dummy_serviceDesc.Streams[2], c.cc, "/spire.dummy.Dummy/BothStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &dummyBothStreamClient{stream}
	return x, nil
}

type Dummy_BothStreamClient interface {
	Send(*BothStreamRequest) error
	Recv() (*BothStreamResponse, error)
	grpc.ClientStream
}

type dummyBothStreamClient struct {
	grpc.ClientStream
}

func (x *dummyBothStreamClient) Send(m *BothStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dummyBothStreamClient) Recv() (*BothStreamResponse, error) {
	m := new(BothStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dummyClient) Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error) {
	out := new(plugin.ConfigureResponse)
	err := grpc.Invoke(ctx, "/spire.dummy.Dummy/Configure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dummyClient) GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error) {
	out := new(plugin.GetPluginInfoResponse)
	err := grpc.Invoke(ctx, "/spire.dummy.Dummy/GetPluginInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Dummy service

type DummyServer interface {
	// No streaming
	NoStream(context.Context, *NoStreamRequest) (*NoStreamResponse, error)
	// Stream things to
	ClientStream(Dummy_ClientStreamServer) error
	// Stream things back
	ServerStream(*ServerStreamRequest, Dummy_ServerStreamServer) error
	BothStream(Dummy_BothStreamServer) error
	// Plugin Methods
	Configure(context.Context, *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error)
	GetPluginInfo(context.Context, *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error)
}

func RegisterDummyServer(s *grpc.Server, srv DummyServer) {
	s.RegisterService(&_Dummy_serviceDesc, srv)
}

func _Dummy_NoStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DummyServer).NoStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.dummy.Dummy/NoStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DummyServer).NoStream(ctx, req.(*NoStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dummy_ClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DummyServer).ClientStream(&dummyClientStreamServer{stream})
}

type Dummy_ClientStreamServer interface {
	SendAndClose(*ClientStreamResponse) error
	Recv() (*ClientStreamRequest, error)
	grpc.ServerStream
}

type dummyClientStreamServer struct {
	grpc.ServerStream
}

func (x *dummyClientStreamServer) SendAndClose(m *ClientStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dummyClientStreamServer) Recv() (*ClientStreamRequest, error) {
	m := new(ClientStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Dummy_ServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ServerStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DummyServer).ServerStream(m, &dummyServerStreamServer{stream})
}

type Dummy_ServerStreamServer interface {
	Send(*ServerStreamResponse) error
	grpc.ServerStream
}

type dummyServerStreamServer struct {
	grpc.ServerStream
}

func (x *dummyServerStreamServer) Send(m *ServerStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Dummy_BothStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DummyServer).BothStream(&dummyBothStreamServer{stream})
}

type Dummy_BothStreamServer interface {
	Send(*BothStreamResponse) error
	Recv() (*BothStreamRequest, error)
	grpc.ServerStream
}

type dummyBothStreamServer struct {
	grpc.ServerStream
}

func (x *dummyBothStreamServer) Send(m *BothStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dummyBothStreamServer) Recv() (*BothStreamRequest, error) {
	m := new(BothStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Dummy_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DummyServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.dummy.Dummy/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DummyServer).Configure(ctx, req.(*plugin.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dummy_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DummyServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.dummy.Dummy/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DummyServer).GetPluginInfo(ctx, req.(*plugin.GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dummy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.dummy.Dummy",
	HandlerType: (*DummyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NoStream",
			Handler:    _Dummy_NoStream_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _Dummy_Configure_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _Dummy_GetPluginInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ClientStream",
			Handler:       _Dummy_ClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ServerStream",
			Handler:       _Dummy_ServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BothStream",
			Handler:       _Dummy_BothStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "dummy.proto",
}

func init() { proto.RegisterFile("dummy.proto", fileDescriptor_dummy_50641bf2c7f7ff0f) }

var fileDescriptor_dummy_50641bf2c7f7ff0f = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4d, 0x4f, 0xc2, 0x30,
	0x18, 0xc7, 0x5d, 0x14, 0x5f, 0x1e, 0x34, 0x6a, 0xe5, 0x60, 0x16, 0x5f, 0x90, 0x44, 0x05, 0x34,
	0x1b, 0xea, 0xc5, 0x33, 0x98, 0x10, 0x2f, 0x06, 0x21, 0x5e, 0xb8, 0x01, 0x3e, 0x83, 0x25, 0x6c,
	0x9d, 0x5d, 0x4b, 0xe2, 0x07, 0xf1, 0xfb, 0x1a, 0xda, 0x55, 0x36, 0x1c, 0x63, 0xa7, 0x06, 0x9e,
	0xdf, 0xff, 0xd7, 0xed, 0xdf, 0x15, 0x8a, 0x9f, 0xc2, 0xf3, 0xbe, 0xad, 0x80, 0x51, 0x4e, 0x49,
	0x31, 0x0c, 0x5c, 0x86, 0x96, 0xfc, 0xcb, 0x7c, 0x18, 0xbb, 0x7c, 0x22, 0x86, 0xd6, 0x88, 0x7a,
	0x76, 0x18, 0xb8, 0x8e, 0x83, 0xb6, 0x1c, 0xdb, 0x92, 0xb5, 0x47, 0xd4, 0xf3, 0xa8, 0x1f, 0x2d,
	0x2a, 0x6f, 0x3e, 0xe7, 0x8a, 0x04, 0x53, 0x31, 0x76, 0xf5, 0xa2, 0x92, 0x95, 0x5b, 0x38, 0x7c,
	0xa3, 0x3d, 0xce, 0x70, 0xe0, 0x75, 0xf1, 0x4b, 0x60, 0xc8, 0x49, 0x09, 0x0a, 0xb3, 0xc1, 0x54,
	0xe0, 0xa9, 0x51, 0x36, 0xaa, 0x9b, 0x5d, 0xf5, 0xa3, 0x52, 0x85, 0xa3, 0x05, 0x18, 0x06, 0xd4,
	0x0f, 0x71, 0x05, 0x79, 0x07, 0x27, 0xad, 0xa9, 0x8b, 0x3e, 0xcf, 0xa3, 0xbd, 0x87, 0x52, 0x12,
	0x5e, 0xa7, 0xee, 0x21, 0x9b, 0x21, 0xcb, 0xa9, 0x4e, 0xc2, 0x99, 0xea, 0x1a, 0x1c, 0x37, 0x29,
	0x9f, 0xe4, 0x11, 0xd7, 0x81, 0xc4, 0xd1, 0x2c, 0xed, 0xe3, 0xcf, 0x16, 0x14, 0x5e, 0xe6, 0xc7,
	0x4a, 0xda, 0xb0, 0xab, 0x0b, 0x24, 0x67, 0x56, 0xec, 0xc0, 0xad, 0xa5, 0x03, 0x30, 0xcf, 0x57,
	0x4c, 0xa3, 0x8d, 0x3e, 0x60, 0x3f, 0x5e, 0x19, 0x29, 0x27, 0xf0, 0x94, 0xea, 0xcd, 0xab, 0x0c,
	0x42, 0x49, 0xab, 0xc6, 0x5c, 0x1b, 0xaf, 0x6b, 0x49, 0x9b, 0x52, 0xfb, 0x92, 0x36, 0xad, 0xeb,
	0x86, 0x41, 0xde, 0x01, 0x16, 0x65, 0x91, 0x8b, 0x44, 0xe4, 0x5f, 0xe1, 0xe6, 0xe5, 0xca, 0xb9,
	0x7e, 0xce, 0x86, 0x41, 0xfa, 0xb0, 0xd7, 0xa2, 0xbe, 0xe3, 0x8e, 0x05, 0x43, 0x72, 0x1d, 0x25,
	0xf4, 0x7d, 0x50, 0x1f, 0xf7, 0xdf, 0x5c, 0x8b, 0x6f, 0xd6, 0x61, 0x51, 0xb9, 0x0e, 0x1c, 0xb4,
	0x91, 0x77, 0xe4, 0xf8, 0xd5, 0x77, 0x28, 0xa9, 0xa5, 0x06, 0x13, 0x8c, 0xde, 0xa3, 0x9e, 0x07,
	0x55, 0xfb, 0x34, 0x77, 0xfa, 0x05, 0xf9, 0x8e, 0x9d, 0x8d, 0x8e, 0x31, 0xdc, 0x96, 0x77, 0xf1,
	0xe9, 0x37, 0x00, 0x00, 0xff, 0xff, 0x1f, 0xa1, 0x11, 0x36, 0x14, 0x04, 0x00, 0x00,
}
