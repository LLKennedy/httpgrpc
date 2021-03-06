// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package httpapi

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ExposedServiceClient is the client API for ExposedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExposedServiceClient interface {
	// ProxyUnary sends a message to the GRPC API expecting an immediate single response
	ProxyUnary(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	// ProxyStream handles streamed requests and responses from websockets
	ProxyStream(ctx context.Context, opts ...grpc.CallOption) (ExposedService_ProxyStreamClient, error)
}

type exposedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExposedServiceClient(cc grpc.ClientConnInterface) ExposedServiceClient {
	return &exposedServiceClient{cc}
}

func (c *exposedServiceClient) ProxyUnary(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/httpapi.ExposedService/ProxyUnary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exposedServiceClient) ProxyStream(ctx context.Context, opts ...grpc.CallOption) (ExposedService_ProxyStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ExposedService_serviceDesc.Streams[0], "/httpapi.ExposedService/ProxyStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &exposedServiceProxyStreamClient{stream}
	return x, nil
}

type ExposedService_ProxyStreamClient interface {
	Send(*StreamedRequest) error
	Recv() (*StreamedResponse, error)
	grpc.ClientStream
}

type exposedServiceProxyStreamClient struct {
	grpc.ClientStream
}

func (x *exposedServiceProxyStreamClient) Send(m *StreamedRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *exposedServiceProxyStreamClient) Recv() (*StreamedResponse, error) {
	m := new(StreamedResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExposedServiceServer is the server API for ExposedService service.
// All implementations must embed UnimplementedExposedServiceServer
// for forward compatibility
type ExposedServiceServer interface {
	// ProxyUnary sends a message to the GRPC API expecting an immediate single response
	ProxyUnary(context.Context, *Request) (*Response, error)
	// ProxyStream handles streamed requests and responses from websockets
	ProxyStream(ExposedService_ProxyStreamServer) error
	mustEmbedUnimplementedExposedServiceServer()
}

// UnimplementedExposedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExposedServiceServer struct {
}

func (UnimplementedExposedServiceServer) ProxyUnary(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProxyUnary not implemented")
}
func (UnimplementedExposedServiceServer) ProxyStream(ExposedService_ProxyStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ProxyStream not implemented")
}
func (UnimplementedExposedServiceServer) mustEmbedUnimplementedExposedServiceServer() {}

// UnsafeExposedServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExposedServiceServer will
// result in compilation errors.
type UnsafeExposedServiceServer interface {
	mustEmbedUnimplementedExposedServiceServer()
}

func RegisterExposedServiceServer(s *grpc.Server, srv ExposedServiceServer) {
	s.RegisterService(&_ExposedService_serviceDesc, srv)
}

func _ExposedService_ProxyUnary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExposedServiceServer).ProxyUnary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/httpapi.ExposedService/ProxyUnary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExposedServiceServer).ProxyUnary(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExposedService_ProxyStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExposedServiceServer).ProxyStream(&exposedServiceProxyStreamServer{stream})
}

type ExposedService_ProxyStreamServer interface {
	Send(*StreamedResponse) error
	Recv() (*StreamedRequest, error)
	grpc.ServerStream
}

type exposedServiceProxyStreamServer struct {
	grpc.ServerStream
}

func (x *exposedServiceProxyStreamServer) Send(m *StreamedResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *exposedServiceProxyStreamServer) Recv() (*StreamedRequest, error) {
	m := new(StreamedRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ExposedService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "httpapi.ExposedService",
	HandlerType: (*ExposedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProxyUnary",
			Handler:    _ExposedService_ProxyUnary_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ProxyStream",
			Handler:       _ExposedService_ProxyStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proxy.proto",
}
