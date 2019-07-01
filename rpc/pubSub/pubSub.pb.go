// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/pubSub.proto

package pubSub

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type String struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *String) Reset()         { *m = String{} }
func (m *String) String() string { return proto.CompactTextString(m) }
func (*String) ProtoMessage()    {}
func (*String) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f998d012dca9cc5, []int{0}
}

func (m *String) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_String.Unmarshal(m, b)
}
func (m *String) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_String.Marshal(b, m, deterministic)
}
func (m *String) XXX_Merge(src proto.Message) {
	xxx_messageInfo_String.Merge(m, src)
}
func (m *String) XXX_Size() int {
	return xxx_messageInfo_String.Size(m)
}
func (m *String) XXX_DiscardUnknown() {
	xxx_messageInfo_String.DiscardUnknown(m)
}

var xxx_messageInfo_String proto.InternalMessageInfo

func (m *String) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Token struct {
	Value                uint64   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f998d012dca9cc5, []int{1}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*String)(nil), "rpc.String")
	proto.RegisterType((*Token)(nil), "rpc.Token")
}

func init() { proto.RegisterFile("rpc/pubSub.proto", fileDescriptor_4f998d012dca9cc5) }

var fileDescriptor_4f998d012dca9cc5 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x2a, 0x48, 0xd6,
	0x2f, 0x28, 0x4d, 0x0a, 0x2e, 0x4d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a,
	0x48, 0x56, 0x92, 0xe3, 0x62, 0x0b, 0x2e, 0x29, 0xca, 0xcc, 0x4b, 0x17, 0x12, 0xe1, 0x62, 0x2d,
	0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x94, 0x64, 0xb9,
	0x58, 0x43, 0xf2, 0xb3, 0x53, 0xf3, 0x50, 0xa5, 0x59, 0xa0, 0xd2, 0x46, 0x8d, 0x8c, 0x5c, 0xbc,
	0x01, 0x60, 0x43, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0x94, 0xb9, 0xd8, 0x03, 0x4a,
	0x93, 0x72, 0x32, 0x8b, 0x33, 0x84, 0xb8, 0xf5, 0x8a, 0x0a, 0x92, 0xf5, 0x20, 0xc6, 0x4b, 0x21,
	0x73, 0x84, 0xd4, 0xb8, 0x38, 0xdd, 0x53, 0x4b, 0x82, 0x4b, 0x8a, 0x52, 0x13, 0x73, 0x85, 0xb8,
	0xc0, 0x32, 0x60, 0x5b, 0x50, 0x54, 0x19, 0x30, 0x0a, 0x29, 0x73, 0x71, 0x04, 0x97, 0x26, 0x15,
	0x27, 0x17, 0x65, 0x26, 0xa1, 0x9a, 0x86, 0xa4, 0x27, 0x89, 0x0d, 0xec, 0x1d, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x65, 0x2d, 0x71, 0xc3, 0xe2, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PubSubServiceClient is the client API for PubSubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PubSubServiceClient interface {
	Publish(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error)
	GetStream(ctx context.Context, in *Token, opts ...grpc.CallOption) (PubSubService_GetStreamClient, error)
	Subscrib(ctx context.Context, in *String, opts ...grpc.CallOption) (*Token, error)
}

type pubSubServiceClient struct {
	cc *grpc.ClientConn
}

func NewPubSubServiceClient(cc *grpc.ClientConn) PubSubServiceClient {
	return &pubSubServiceClient{cc}
}

func (c *pubSubServiceClient) Publish(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/rpc.PubSubService/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubSubServiceClient) GetStream(ctx context.Context, in *Token, opts ...grpc.CallOption) (PubSubService_GetStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PubSubService_serviceDesc.Streams[0], "/rpc.PubSubService/GetStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &pubSubServiceGetStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PubSubService_GetStreamClient interface {
	Recv() (*String, error)
	grpc.ClientStream
}

type pubSubServiceGetStreamClient struct {
	grpc.ClientStream
}

func (x *pubSubServiceGetStreamClient) Recv() (*String, error) {
	m := new(String)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pubSubServiceClient) Subscrib(ctx context.Context, in *String, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/rpc.PubSubService/Subscrib", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PubSubServiceServer is the server API for PubSubService service.
type PubSubServiceServer interface {
	Publish(context.Context, *String) (*String, error)
	GetStream(*Token, PubSubService_GetStreamServer) error
	Subscrib(context.Context, *String) (*Token, error)
}

// UnimplementedPubSubServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPubSubServiceServer struct {
}

func (*UnimplementedPubSubServiceServer) Publish(ctx context.Context, req *String) (*String, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (*UnimplementedPubSubServiceServer) GetStream(req *Token, srv PubSubService_GetStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStream not implemented")
}
func (*UnimplementedPubSubServiceServer) Subscrib(ctx context.Context, req *String) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscrib not implemented")
}

func RegisterPubSubServiceServer(s *grpc.Server, srv PubSubServiceServer) {
	s.RegisterService(&_PubSubService_serviceDesc, srv)
}

func _PubSubService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSubServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.PubSubService/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSubServiceServer).Publish(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubSubService_GetStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Token)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PubSubServiceServer).GetStream(m, &pubSubServiceGetStreamServer{stream})
}

type PubSubService_GetStreamServer interface {
	Send(*String) error
	grpc.ServerStream
}

type pubSubServiceGetStreamServer struct {
	grpc.ServerStream
}

func (x *pubSubServiceGetStreamServer) Send(m *String) error {
	return x.ServerStream.SendMsg(m)
}

func _PubSubService_Subscrib_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSubServiceServer).Subscrib(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.PubSubService/Subscrib",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSubServiceServer).Subscrib(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

var _PubSubService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.PubSubService",
	HandlerType: (*PubSubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _PubSubService_Publish_Handler,
		},
		{
			MethodName: "Subscrib",
			Handler:    _PubSubService_Subscrib_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStream",
			Handler:       _PubSubService_GetStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpc/pubSub.proto",
}