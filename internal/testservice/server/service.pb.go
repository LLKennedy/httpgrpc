// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package server

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

type FibonacciRequest struct {
	// The 1-indexed point in the Fibonacci sequence
	N                    uint64   `protobuf:"varint,1,opt,name=n,proto3" json:"n,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibonacciRequest) Reset()         { *m = FibonacciRequest{} }
func (m *FibonacciRequest) String() string { return proto.CompactTextString(m) }
func (*FibonacciRequest) ProtoMessage()    {}
func (*FibonacciRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *FibonacciRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibonacciRequest.Unmarshal(m, b)
}
func (m *FibonacciRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibonacciRequest.Marshal(b, m, deterministic)
}
func (m *FibonacciRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibonacciRequest.Merge(m, src)
}
func (m *FibonacciRequest) XXX_Size() int {
	return xxx_messageInfo_FibonacciRequest.Size(m)
}
func (m *FibonacciRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FibonacciRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FibonacciRequest proto.InternalMessageInfo

func (m *FibonacciRequest) GetN() uint64 {
	if m != nil {
		return m.N
	}
	return 0
}

type FibonacciResponse struct {
	// The number found in the nth place of the sequence
	Number               uint64   `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibonacciResponse) Reset()         { *m = FibonacciResponse{} }
func (m *FibonacciResponse) String() string { return proto.CompactTextString(m) }
func (*FibonacciResponse) ProtoMessage()    {}
func (*FibonacciResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *FibonacciResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibonacciResponse.Unmarshal(m, b)
}
func (m *FibonacciResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibonacciResponse.Marshal(b, m, deterministic)
}
func (m *FibonacciResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibonacciResponse.Merge(m, src)
}
func (m *FibonacciResponse) XXX_Size() int {
	return xxx_messageInfo_FibonacciResponse.Size(m)
}
func (m *FibonacciResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FibonacciResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FibonacciResponse proto.InternalMessageInfo

func (m *FibonacciResponse) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type RandomRequest struct {
	// The lowest inclusive integer for the resulting random number
	LowerBound int64 `protobuf:"varint,1,opt,name=lower_bound,json=lowerBound,proto3" json:"lower_bound,omitempty"`
	// The highest inclusive integer for the resulting random number
	UpperBound           int64    `protobuf:"varint,2,opt,name=upper_bound,json=upperBound,proto3" json:"upper_bound,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RandomRequest) Reset()         { *m = RandomRequest{} }
func (m *RandomRequest) String() string { return proto.CompactTextString(m) }
func (*RandomRequest) ProtoMessage()    {}
func (*RandomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *RandomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RandomRequest.Unmarshal(m, b)
}
func (m *RandomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RandomRequest.Marshal(b, m, deterministic)
}
func (m *RandomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RandomRequest.Merge(m, src)
}
func (m *RandomRequest) XXX_Size() int {
	return xxx_messageInfo_RandomRequest.Size(m)
}
func (m *RandomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RandomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RandomRequest proto.InternalMessageInfo

func (m *RandomRequest) GetLowerBound() int64 {
	if m != nil {
		return m.LowerBound
	}
	return 0
}

func (m *RandomRequest) GetUpperBound() int64 {
	if m != nil {
		return m.UpperBound
	}
	return 0
}

type RandomResponse struct {
	// The generated number
	Number               int64    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RandomResponse) Reset()         { *m = RandomResponse{} }
func (m *RandomResponse) String() string { return proto.CompactTextString(m) }
func (*RandomResponse) ProtoMessage()    {}
func (*RandomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *RandomResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RandomResponse.Unmarshal(m, b)
}
func (m *RandomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RandomResponse.Marshal(b, m, deterministic)
}
func (m *RandomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RandomResponse.Merge(m, src)
}
func (m *RandomResponse) XXX_Size() int {
	return xxx_messageInfo_RandomResponse.Size(m)
}
func (m *RandomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RandomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RandomResponse proto.InternalMessageInfo

func (m *RandomResponse) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type UploadPhotoRequest struct {
	// The raw bytes of the photo
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadPhotoRequest) Reset()         { *m = UploadPhotoRequest{} }
func (m *UploadPhotoRequest) String() string { return proto.CompactTextString(m) }
func (*UploadPhotoRequest) ProtoMessage()    {}
func (*UploadPhotoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *UploadPhotoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadPhotoRequest.Unmarshal(m, b)
}
func (m *UploadPhotoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadPhotoRequest.Marshal(b, m, deterministic)
}
func (m *UploadPhotoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadPhotoRequest.Merge(m, src)
}
func (m *UploadPhotoRequest) XXX_Size() int {
	return xxx_messageInfo_UploadPhotoRequest.Size(m)
}
func (m *UploadPhotoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadPhotoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadPhotoRequest proto.InternalMessageInfo

func (m *UploadPhotoRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type UploadPhotoResponse struct {
	// The uuid generated to identify and retreive this photo
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadPhotoResponse) Reset()         { *m = UploadPhotoResponse{} }
func (m *UploadPhotoResponse) String() string { return proto.CompactTextString(m) }
func (*UploadPhotoResponse) ProtoMessage()    {}
func (*UploadPhotoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *UploadPhotoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadPhotoResponse.Unmarshal(m, b)
}
func (m *UploadPhotoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadPhotoResponse.Marshal(b, m, deterministic)
}
func (m *UploadPhotoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadPhotoResponse.Merge(m, src)
}
func (m *UploadPhotoResponse) XXX_Size() int {
	return xxx_messageInfo_UploadPhotoResponse.Size(m)
}
func (m *UploadPhotoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadPhotoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadPhotoResponse proto.InternalMessageInfo

func (m *UploadPhotoResponse) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func init() {
	proto.RegisterType((*FibonacciRequest)(nil), "server.FibonacciRequest")
	proto.RegisterType((*FibonacciResponse)(nil), "server.FibonacciResponse")
	proto.RegisterType((*RandomRequest)(nil), "server.RandomRequest")
	proto.RegisterType((*RandomResponse)(nil), "server.RandomResponse")
	proto.RegisterType((*UploadPhotoRequest)(nil), "server.UploadPhotoRequest")
	proto.RegisterType((*UploadPhotoResponse)(nil), "server.UploadPhotoResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x09, 0xad, 0x82, 0x72, 0xb4, 0xfc, 0x39, 0x44, 0x55, 0xc2, 0x40, 0xe5, 0x29, 0x08,
	0x29, 0x03, 0xac, 0x4c, 0x1d, 0x60, 0x61, 0x28, 0x46, 0xcc, 0x28, 0x7f, 0x2c, 0x11, 0xa9, 0xf5,
	0x85, 0xd8, 0x86, 0xcf, 0xca, 0xb7, 0x41, 0xb1, 0xe3, 0xa8, 0x2d, 0xd9, 0x7c, 0xef, 0x7e, 0x79,
	0x77, 0xf7, 0x14, 0x98, 0x2a, 0xd1, 0x7c, 0x57, 0x85, 0x48, 0xeb, 0x86, 0x34, 0x61, 0xd8, 0x96,
	0xa2, 0x61, 0x0b, 0x38, 0x7b, 0xaa, 0x72, 0x92, 0x59, 0x51, 0x54, 0x5c, 0x7c, 0x19, 0xa1, 0x34,
	0x4e, 0x20, 0x90, 0xf3, 0x60, 0x11, 0x24, 0x63, 0x1e, 0x48, 0x76, 0x07, 0xe7, 0x5b, 0x84, 0xaa,
	0x49, 0x2a, 0x81, 0x33, 0x08, 0xa5, 0xd9, 0xe4, 0xa2, 0xe9, 0xb8, 0xae, 0x62, 0xaf, 0x30, 0xe5,
	0x99, 0x2c, 0x69, 0xe3, 0xbd, 0x6e, 0xe0, 0x78, 0x4d, 0x3f, 0xa2, 0xf9, 0xc8, 0xc9, 0xc8, 0xd2,
	0xd2, 0x23, 0x0e, 0x56, 0x5a, 0xb6, 0x4a, 0x0b, 0x98, 0xba, 0xee, 0x81, 0x43, 0x07, 0x58, 0xc9,
	0x02, 0x2c, 0x81, 0x13, 0x6f, 0x39, 0x38, 0x7c, 0xd4, 0x0f, 0x4f, 0x00, 0xdf, 0xeb, 0x35, 0x65,
	0xe5, 0xea, 0x93, 0x34, 0xf9, 0x0d, 0x10, 0xc6, 0x65, 0xa6, 0x33, 0xcb, 0x4e, 0xb8, 0x7d, 0xb3,
	0x5b, 0xb8, 0xd8, 0x21, 0x3b, 0x63, 0x84, 0xb1, 0x31, 0x95, 0xdb, 0x32, 0xe2, 0xf6, 0x7d, 0xff,
	0x1b, 0xc0, 0xd1, 0x9b, 0x8b, 0x0e, 0x97, 0x10, 0xf5, 0x51, 0xe0, 0x3c, 0x75, 0x11, 0xa6, 0xfb,
	0xf9, 0xc5, 0x57, 0x03, 0x1d, 0x37, 0x81, 0x1d, 0xe0, 0x23, 0x44, 0xcf, 0x42, 0xbb, 0x8b, 0xf0,
	0xd2, 0x93, 0x3b, 0xa1, 0xc5, 0xb3, 0x7d, 0xb9, 0xff, 0xfa, 0x05, 0x4e, 0x57, 0xa4, 0xf4, 0xd6,
	0xf2, 0x18, 0x7b, 0xf8, 0xff, 0xed, 0xf1, 0xf5, 0x60, 0xcf, 0xbb, 0xe5, 0xa1, 0xfd, 0x17, 0x1e,
	0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa9, 0xa5, 0x0f, 0x69, 0x1c, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	// Fibonacci returns the nth number in the Fibonacci sequence. It does not start with an HTTP method and is therefore not exposed
	Fibonacci(ctx context.Context, in *FibonacciRequest, opts ...grpc.CallOption) (*FibonacciResponse, error)
	// GetRandom returns a random integer in the desired range. It may be accessed via a Get request to the proxy at, for example, /api/Service/Random
	GetRandom(ctx context.Context, in *RandomRequest, opts ...grpc.CallOption) (*RandomResponse, error)
	// PostUploadPhoto allows the upload of a photo to some persistence store. It may be accessed via  Post request to the proxy at, for example, /api/Service/UploadPhoto
	PostUploadPhoto(ctx context.Context, in *UploadPhotoRequest, opts ...grpc.CallOption) (*UploadPhotoResponse, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Fibonacci(ctx context.Context, in *FibonacciRequest, opts ...grpc.CallOption) (*FibonacciResponse, error) {
	out := new(FibonacciResponse)
	err := c.cc.Invoke(ctx, "/server.Service/Fibonacci", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetRandom(ctx context.Context, in *RandomRequest, opts ...grpc.CallOption) (*RandomResponse, error) {
	out := new(RandomResponse)
	err := c.cc.Invoke(ctx, "/server.Service/GetRandom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) PostUploadPhoto(ctx context.Context, in *UploadPhotoRequest, opts ...grpc.CallOption) (*UploadPhotoResponse, error) {
	out := new(UploadPhotoResponse)
	err := c.cc.Invoke(ctx, "/server.Service/PostUploadPhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	// Fibonacci returns the nth number in the Fibonacci sequence. It does not start with an HTTP method and is therefore not exposed
	Fibonacci(context.Context, *FibonacciRequest) (*FibonacciResponse, error)
	// GetRandom returns a random integer in the desired range. It may be accessed via a Get request to the proxy at, for example, /api/Service/Random
	GetRandom(context.Context, *RandomRequest) (*RandomResponse, error)
	// PostUploadPhoto allows the upload of a photo to some persistence store. It may be accessed via  Post request to the proxy at, for example, /api/Service/UploadPhoto
	PostUploadPhoto(context.Context, *UploadPhotoRequest) (*UploadPhotoResponse, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) Fibonacci(ctx context.Context, req *FibonacciRequest) (*FibonacciResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fibonacci not implemented")
}
func (*UnimplementedServiceServer) GetRandom(ctx context.Context, req *RandomRequest) (*RandomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRandom not implemented")
}
func (*UnimplementedServiceServer) PostUploadPhoto(ctx context.Context, req *UploadPhotoRequest) (*UploadPhotoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostUploadPhoto not implemented")
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Fibonacci_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FibonacciRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Fibonacci(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Service/Fibonacci",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Fibonacci(ctx, req.(*FibonacciRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetRandom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RandomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetRandom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Service/GetRandom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetRandom(ctx, req.(*RandomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_PostUploadPhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadPhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).PostUploadPhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Service/PostUploadPhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).PostUploadPhoto(ctx, req.(*UploadPhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fibonacci",
			Handler:    _Service_Fibonacci_Handler,
		},
		{
			MethodName: "GetRandom",
			Handler:    _Service_GetRandom_Handler,
		},
		{
			MethodName: "PostUploadPhoto",
			Handler:    _Service_PostUploadPhoto_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
