// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AppClient is the client API for App service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppClient interface {
	// Fibonacci returns the nth number in the Fibonacci sequence. It does not start with an HTTP method and is therefore not exposed
	Fibonacci(ctx context.Context, in *FibonacciRequest, opts ...grpc.CallOption) (*FibonacciResponse, error)
	// Random returns a random integer in the desired range. It may be accessed via a Get request to the proxy at, for example, /api/Service/Random
	Random(ctx context.Context, in *RandomRequest, opts ...grpc.CallOption) (*RandomResponse, error)
	// UploadPhoto allows the upload of a photo to some persistence store. It may be accessed via  Post request to the proxy at, for example, /api/Service/UploadPhoto
	UploadPhoto(ctx context.Context, in *UploadPhotoRequest, opts ...grpc.CallOption) (*UploadPhotoResponse, error)
	// Feed sends streamed inputs
	Feed(ctx context.Context, opts ...grpc.CallOption) (App_FeedClient, error)
	// Broadcast asks the App to broadcast data in a stream
	Broadcast(ctx context.Context, in *BroadcastRequest, opts ...grpc.CallOption) (App_BroadcastClient, error)
	// ConvertToString streams conversions of the input stream to strings
	ConvertToString(ctx context.Context, opts ...grpc.CallOption) (App_ConvertToStringClient, error)
}

type appClient struct {
	cc grpc.ClientConnInterface
}

func NewAppClient(cc grpc.ClientConnInterface) AppClient {
	return &appClient{cc}
}

func (c *appClient) Fibonacci(ctx context.Context, in *FibonacciRequest, opts ...grpc.CallOption) (*FibonacciResponse, error) {
	out := new(FibonacciResponse)
	err := c.cc.Invoke(ctx, "/service.App/Fibonacci", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) Random(ctx context.Context, in *RandomRequest, opts ...grpc.CallOption) (*RandomResponse, error) {
	out := new(RandomResponse)
	err := c.cc.Invoke(ctx, "/service.App/Random", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) UploadPhoto(ctx context.Context, in *UploadPhotoRequest, opts ...grpc.CallOption) (*UploadPhotoResponse, error) {
	out := new(UploadPhotoResponse)
	err := c.cc.Invoke(ctx, "/service.App/UploadPhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) Feed(ctx context.Context, opts ...grpc.CallOption) (App_FeedClient, error) {
	stream, err := c.cc.NewStream(ctx, &_App_serviceDesc.Streams[0], "/service.App/Feed", opts...)
	if err != nil {
		return nil, err
	}
	x := &appFeedClient{stream}
	return x, nil
}

type App_FeedClient interface {
	Send(*FeedData) error
	CloseAndRecv() (*FeedResponse, error)
	grpc.ClientStream
}

type appFeedClient struct {
	grpc.ClientStream
}

func (x *appFeedClient) Send(m *FeedData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *appFeedClient) CloseAndRecv() (*FeedResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FeedResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *appClient) Broadcast(ctx context.Context, in *BroadcastRequest, opts ...grpc.CallOption) (App_BroadcastClient, error) {
	stream, err := c.cc.NewStream(ctx, &_App_serviceDesc.Streams[1], "/service.App/Broadcast", opts...)
	if err != nil {
		return nil, err
	}
	x := &appBroadcastClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type App_BroadcastClient interface {
	Recv() (*BroadcastData, error)
	grpc.ClientStream
}

type appBroadcastClient struct {
	grpc.ClientStream
}

func (x *appBroadcastClient) Recv() (*BroadcastData, error) {
	m := new(BroadcastData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *appClient) ConvertToString(ctx context.Context, opts ...grpc.CallOption) (App_ConvertToStringClient, error) {
	stream, err := c.cc.NewStream(ctx, &_App_serviceDesc.Streams[2], "/service.App/ConvertToString", opts...)
	if err != nil {
		return nil, err
	}
	x := &appConvertToStringClient{stream}
	return x, nil
}

type App_ConvertToStringClient interface {
	Send(*ConvertInput) error
	Recv() (*ConvertOutput, error)
	grpc.ClientStream
}

type appConvertToStringClient struct {
	grpc.ClientStream
}

func (x *appConvertToStringClient) Send(m *ConvertInput) error {
	return x.ClientStream.SendMsg(m)
}

func (x *appConvertToStringClient) Recv() (*ConvertOutput, error) {
	m := new(ConvertOutput)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AppServer is the server API for App service.
// All implementations must embed UnimplementedAppServer
// for forward compatibility
type AppServer interface {
	// Fibonacci returns the nth number in the Fibonacci sequence. It does not start with an HTTP method and is therefore not exposed
	Fibonacci(context.Context, *FibonacciRequest) (*FibonacciResponse, error)
	// Random returns a random integer in the desired range. It may be accessed via a Get request to the proxy at, for example, /api/Service/Random
	Random(context.Context, *RandomRequest) (*RandomResponse, error)
	// UploadPhoto allows the upload of a photo to some persistence store. It may be accessed via  Post request to the proxy at, for example, /api/Service/UploadPhoto
	UploadPhoto(context.Context, *UploadPhotoRequest) (*UploadPhotoResponse, error)
	// Feed sends streamed inputs
	Feed(App_FeedServer) error
	// Broadcast asks the App to broadcast data in a stream
	Broadcast(*BroadcastRequest, App_BroadcastServer) error
	// ConvertToString streams conversions of the input stream to strings
	ConvertToString(App_ConvertToStringServer) error
	mustEmbedUnimplementedAppServer()
}

// UnimplementedAppServer must be embedded to have forward compatible implementations.
type UnimplementedAppServer struct {
}

func (UnimplementedAppServer) Fibonacci(context.Context, *FibonacciRequest) (*FibonacciResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fibonacci not implemented")
}
func (UnimplementedAppServer) Random(context.Context, *RandomRequest) (*RandomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Random not implemented")
}
func (UnimplementedAppServer) UploadPhoto(context.Context, *UploadPhotoRequest) (*UploadPhotoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadPhoto not implemented")
}
func (UnimplementedAppServer) Feed(App_FeedServer) error {
	return status.Errorf(codes.Unimplemented, "method Feed not implemented")
}
func (UnimplementedAppServer) Broadcast(*BroadcastRequest, App_BroadcastServer) error {
	return status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedAppServer) ConvertToString(App_ConvertToStringServer) error {
	return status.Errorf(codes.Unimplemented, "method ConvertToString not implemented")
}
func (UnimplementedAppServer) mustEmbedUnimplementedAppServer() {}

// UnsafeAppServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppServer will
// result in compilation errors.
type UnsafeAppServer interface {
	mustEmbedUnimplementedAppServer()
}

func RegisterAppServer(s *grpc.Server, srv AppServer) {
	s.RegisterService(&_App_serviceDesc, srv)
}

func _App_Fibonacci_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FibonacciRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Fibonacci(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.App/Fibonacci",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Fibonacci(ctx, req.(*FibonacciRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_Random_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RandomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Random(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.App/Random",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Random(ctx, req.(*RandomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_UploadPhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadPhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).UploadPhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.App/UploadPhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).UploadPhoto(ctx, req.(*UploadPhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_Feed_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AppServer).Feed(&appFeedServer{stream})
}

type App_FeedServer interface {
	SendAndClose(*FeedResponse) error
	Recv() (*FeedData, error)
	grpc.ServerStream
}

type appFeedServer struct {
	grpc.ServerStream
}

func (x *appFeedServer) SendAndClose(m *FeedResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *appFeedServer) Recv() (*FeedData, error) {
	m := new(FeedData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _App_Broadcast_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BroadcastRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AppServer).Broadcast(m, &appBroadcastServer{stream})
}

type App_BroadcastServer interface {
	Send(*BroadcastData) error
	grpc.ServerStream
}

type appBroadcastServer struct {
	grpc.ServerStream
}

func (x *appBroadcastServer) Send(m *BroadcastData) error {
	return x.ServerStream.SendMsg(m)
}

func _App_ConvertToString_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AppServer).ConvertToString(&appConvertToStringServer{stream})
}

type App_ConvertToStringServer interface {
	Send(*ConvertOutput) error
	Recv() (*ConvertInput, error)
	grpc.ServerStream
}

type appConvertToStringServer struct {
	grpc.ServerStream
}

func (x *appConvertToStringServer) Send(m *ConvertOutput) error {
	return x.ServerStream.SendMsg(m)
}

func (x *appConvertToStringServer) Recv() (*ConvertInput, error) {
	m := new(ConvertInput)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _App_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.App",
	HandlerType: (*AppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fibonacci",
			Handler:    _App_Fibonacci_Handler,
		},
		{
			MethodName: "Random",
			Handler:    _App_Random_Handler,
		},
		{
			MethodName: "UploadPhoto",
			Handler:    _App_UploadPhoto_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Feed",
			Handler:       _App_Feed_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Broadcast",
			Handler:       _App_Broadcast_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ConvertToString",
			Handler:       _App_ConvertToString_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}

// ExposedAppClient is the client API for ExposedApp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExposedAppClient interface {
	GetRandom(ctx context.Context, in *RandomRequest, opts ...grpc.CallOption) (*RandomResponse, error)
	PostUploadPhoto(ctx context.Context, in *UploadPhotoRequest, opts ...grpc.CallOption) (*UploadPhotoResponse, error)
	PostFeed(ctx context.Context, opts ...grpc.CallOption) (ExposedApp_PostFeedClient, error)
	GetBroadcast(ctx context.Context, in *BroadcastRequest, opts ...grpc.CallOption) (ExposedApp_GetBroadcastClient, error)
	PostConvertToString(ctx context.Context, opts ...grpc.CallOption) (ExposedApp_PostConvertToStringClient, error)
}

type exposedAppClient struct {
	cc grpc.ClientConnInterface
}

func NewExposedAppClient(cc grpc.ClientConnInterface) ExposedAppClient {
	return &exposedAppClient{cc}
}

func (c *exposedAppClient) GetRandom(ctx context.Context, in *RandomRequest, opts ...grpc.CallOption) (*RandomResponse, error) {
	out := new(RandomResponse)
	err := c.cc.Invoke(ctx, "/service.ExposedApp/GetRandom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exposedAppClient) PostUploadPhoto(ctx context.Context, in *UploadPhotoRequest, opts ...grpc.CallOption) (*UploadPhotoResponse, error) {
	out := new(UploadPhotoResponse)
	err := c.cc.Invoke(ctx, "/service.ExposedApp/PostUploadPhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exposedAppClient) PostFeed(ctx context.Context, opts ...grpc.CallOption) (ExposedApp_PostFeedClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ExposedApp_serviceDesc.Streams[0], "/service.ExposedApp/PostFeed", opts...)
	if err != nil {
		return nil, err
	}
	x := &exposedAppPostFeedClient{stream}
	return x, nil
}

type ExposedApp_PostFeedClient interface {
	Send(*FeedData) error
	CloseAndRecv() (*FeedResponse, error)
	grpc.ClientStream
}

type exposedAppPostFeedClient struct {
	grpc.ClientStream
}

func (x *exposedAppPostFeedClient) Send(m *FeedData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *exposedAppPostFeedClient) CloseAndRecv() (*FeedResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FeedResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *exposedAppClient) GetBroadcast(ctx context.Context, in *BroadcastRequest, opts ...grpc.CallOption) (ExposedApp_GetBroadcastClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ExposedApp_serviceDesc.Streams[1], "/service.ExposedApp/GetBroadcast", opts...)
	if err != nil {
		return nil, err
	}
	x := &exposedAppGetBroadcastClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExposedApp_GetBroadcastClient interface {
	Recv() (*BroadcastData, error)
	grpc.ClientStream
}

type exposedAppGetBroadcastClient struct {
	grpc.ClientStream
}

func (x *exposedAppGetBroadcastClient) Recv() (*BroadcastData, error) {
	m := new(BroadcastData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *exposedAppClient) PostConvertToString(ctx context.Context, opts ...grpc.CallOption) (ExposedApp_PostConvertToStringClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ExposedApp_serviceDesc.Streams[2], "/service.ExposedApp/PostConvertToString", opts...)
	if err != nil {
		return nil, err
	}
	x := &exposedAppPostConvertToStringClient{stream}
	return x, nil
}

type ExposedApp_PostConvertToStringClient interface {
	Send(*ConvertInput) error
	Recv() (*ConvertOutput, error)
	grpc.ClientStream
}

type exposedAppPostConvertToStringClient struct {
	grpc.ClientStream
}

func (x *exposedAppPostConvertToStringClient) Send(m *ConvertInput) error {
	return x.ClientStream.SendMsg(m)
}

func (x *exposedAppPostConvertToStringClient) Recv() (*ConvertOutput, error) {
	m := new(ConvertOutput)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExposedAppServer is the server API for ExposedApp service.
// All implementations must embed UnimplementedExposedAppServer
// for forward compatibility
type ExposedAppServer interface {
	GetRandom(context.Context, *RandomRequest) (*RandomResponse, error)
	PostUploadPhoto(context.Context, *UploadPhotoRequest) (*UploadPhotoResponse, error)
	PostFeed(ExposedApp_PostFeedServer) error
	GetBroadcast(*BroadcastRequest, ExposedApp_GetBroadcastServer) error
	PostConvertToString(ExposedApp_PostConvertToStringServer) error
	mustEmbedUnimplementedExposedAppServer()
}

// UnimplementedExposedAppServer must be embedded to have forward compatible implementations.
type UnimplementedExposedAppServer struct {
}

func (UnimplementedExposedAppServer) GetRandom(context.Context, *RandomRequest) (*RandomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRandom not implemented")
}
func (UnimplementedExposedAppServer) PostUploadPhoto(context.Context, *UploadPhotoRequest) (*UploadPhotoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostUploadPhoto not implemented")
}
func (UnimplementedExposedAppServer) PostFeed(ExposedApp_PostFeedServer) error {
	return status.Errorf(codes.Unimplemented, "method PostFeed not implemented")
}
func (UnimplementedExposedAppServer) GetBroadcast(*BroadcastRequest, ExposedApp_GetBroadcastServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBroadcast not implemented")
}
func (UnimplementedExposedAppServer) PostConvertToString(ExposedApp_PostConvertToStringServer) error {
	return status.Errorf(codes.Unimplemented, "method PostConvertToString not implemented")
}
func (UnimplementedExposedAppServer) mustEmbedUnimplementedExposedAppServer() {}

// UnsafeExposedAppServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExposedAppServer will
// result in compilation errors.
type UnsafeExposedAppServer interface {
	mustEmbedUnimplementedExposedAppServer()
}

func RegisterExposedAppServer(s *grpc.Server, srv ExposedAppServer) {
	s.RegisterService(&_ExposedApp_serviceDesc, srv)
}

func _ExposedApp_GetRandom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RandomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExposedAppServer).GetRandom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ExposedApp/GetRandom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExposedAppServer).GetRandom(ctx, req.(*RandomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExposedApp_PostUploadPhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadPhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExposedAppServer).PostUploadPhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ExposedApp/PostUploadPhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExposedAppServer).PostUploadPhoto(ctx, req.(*UploadPhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExposedApp_PostFeed_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExposedAppServer).PostFeed(&exposedAppPostFeedServer{stream})
}

type ExposedApp_PostFeedServer interface {
	SendAndClose(*FeedResponse) error
	Recv() (*FeedData, error)
	grpc.ServerStream
}

type exposedAppPostFeedServer struct {
	grpc.ServerStream
}

func (x *exposedAppPostFeedServer) SendAndClose(m *FeedResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *exposedAppPostFeedServer) Recv() (*FeedData, error) {
	m := new(FeedData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ExposedApp_GetBroadcast_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BroadcastRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExposedAppServer).GetBroadcast(m, &exposedAppGetBroadcastServer{stream})
}

type ExposedApp_GetBroadcastServer interface {
	Send(*BroadcastData) error
	grpc.ServerStream
}

type exposedAppGetBroadcastServer struct {
	grpc.ServerStream
}

func (x *exposedAppGetBroadcastServer) Send(m *BroadcastData) error {
	return x.ServerStream.SendMsg(m)
}

func _ExposedApp_PostConvertToString_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExposedAppServer).PostConvertToString(&exposedAppPostConvertToStringServer{stream})
}

type ExposedApp_PostConvertToStringServer interface {
	Send(*ConvertOutput) error
	Recv() (*ConvertInput, error)
	grpc.ServerStream
}

type exposedAppPostConvertToStringServer struct {
	grpc.ServerStream
}

func (x *exposedAppPostConvertToStringServer) Send(m *ConvertOutput) error {
	return x.ServerStream.SendMsg(m)
}

func (x *exposedAppPostConvertToStringServer) Recv() (*ConvertInput, error) {
	m := new(ConvertInput)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ExposedApp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.ExposedApp",
	HandlerType: (*ExposedAppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRandom",
			Handler:    _ExposedApp_GetRandom_Handler,
		},
		{
			MethodName: "PostUploadPhoto",
			Handler:    _ExposedApp_PostUploadPhoto_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PostFeed",
			Handler:       _ExposedApp_PostFeed_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetBroadcast",
			Handler:       _ExposedApp_GetBroadcast_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PostConvertToString",
			Handler:       _ExposedApp_PostConvertToString_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}