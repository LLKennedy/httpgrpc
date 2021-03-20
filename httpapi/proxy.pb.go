// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.10.1
// source: proxy.proto

package httpapi

import (
	proto "github.com/golang/protobuf/proto"
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

// All valid HTTP method names
type Method int32

const (
	Method_UNKNOWN Method = 0
	Method_GET     Method = 1
	Method_HEAD    Method = 2
	Method_POST    Method = 3
	Method_PUT     Method = 4
	Method_DELETE  Method = 5
	Method_CONNECT Method = 6
	Method_OPTIONS Method = 7
	Method_TRACE   Method = 8
	Method_PATCH   Method = 9
)

// Enum value maps for Method.
var (
	Method_name = map[int32]string{
		0: "UNKNOWN",
		1: "GET",
		2: "HEAD",
		3: "POST",
		4: "PUT",
		5: "DELETE",
		6: "CONNECT",
		7: "OPTIONS",
		8: "TRACE",
		9: "PATCH",
	}
	Method_value = map[string]int32{
		"UNKNOWN": 0,
		"GET":     1,
		"HEAD":    2,
		"POST":    3,
		"PUT":     4,
		"DELETE":  5,
		"CONNECT": 6,
		"OPTIONS": 7,
		"TRACE":   8,
		"PATCH":   9,
	}
)

func (x Method) Enum() *Method {
	p := new(Method)
	*p = x
	return p
}

func (x Method) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Method) Descriptor() protoreflect.EnumDescriptor {
	return file_proxy_proto_enumTypes[0].Descriptor()
}

func (Method) Type() protoreflect.EnumType {
	return &file_proxy_proto_enumTypes[0]
}

func (x Method) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Method.Descriptor instead.
func (Method) EnumDescriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{0}
}

type StreamedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to MessageType:
	//	*StreamedRequest_Init
	//	*StreamedRequest_Request
	MessageType isStreamedRequest_MessageType `protobuf_oneof:"message_type"`
}

func (x *StreamedRequest) Reset() {
	*x = StreamedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamedRequest) ProtoMessage() {}

func (x *StreamedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamedRequest.ProtoReflect.Descriptor instead.
func (*StreamedRequest) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{0}
}

func (m *StreamedRequest) GetMessageType() isStreamedRequest_MessageType {
	if m != nil {
		return m.MessageType
	}
	return nil
}

func (x *StreamedRequest) GetInit() *RoutingInformation {
	if x, ok := x.GetMessageType().(*StreamedRequest_Init); ok {
		return x.Init
	}
	return nil
}

func (x *StreamedRequest) GetRequest() []byte {
	if x, ok := x.GetMessageType().(*StreamedRequest_Request); ok {
		return x.Request
	}
	return nil
}

type isStreamedRequest_MessageType interface {
	isStreamedRequest_MessageType()
}

type StreamedRequest_Init struct {
	// Init will only be sent once on initialisation of the websocket that this stream is linked to, it contains the routing information normally transmitted by the Request message
	Init *RoutingInformation `protobuf:"bytes,1,opt,name=init,proto3,oneof"`
}

type StreamedRequest_Request struct {
	// Request will be sent for every message but the very first, and is the JSON extracted from the websocket message
	Request []byte `protobuf:"bytes,2,opt,name=request,proto3,oneof"`
}

func (*StreamedRequest_Init) isStreamedRequest_MessageType() {}

func (*StreamedRequest_Request) isStreamedRequest_MessageType() {}

type RoutingInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// HTTP method from requestor
	Method Method `protobuf:"varint,1,opt,name=method,proto3,enum=httpapi.Method" json:"method,omitempty"`
	// Desired procedure name
	Procedure string `protobuf:"bytes,2,opt,name=procedure,proto3" json:"procedure,omitempty"`
	// Headers won't be automatically translated into the GRPC requests, but they should be delivered in case manual conversion is needed
	Headers map[string]*MultiVal `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RoutingInformation) Reset() {
	*x = RoutingInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoutingInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutingInformation) ProtoMessage() {}

func (x *RoutingInformation) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutingInformation.ProtoReflect.Descriptor instead.
func (*RoutingInformation) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{1}
}

func (x *RoutingInformation) GetMethod() Method {
	if x != nil {
		return x.Method
	}
	return Method_UNKNOWN
}

func (x *RoutingInformation) GetProcedure() string {
	if x != nil {
		return x.Procedure
	}
	return ""
}

func (x *RoutingInformation) GetHeaders() map[string]*MultiVal {
	if x != nil {
		return x.Headers
	}
	return nil
}

// TODO: work out what we can do here
type StreamedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response []byte `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *StreamedResponse) Reset() {
	*x = StreamedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamedResponse) ProtoMessage() {}

func (x *StreamedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamedResponse.ProtoReflect.Descriptor instead.
func (*StreamedResponse) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{2}
}

func (x *StreamedResponse) GetResponse() []byte {
	if x != nil {
		return x.Response
	}
	return nil
}

// The request data from the HTTP request
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// HTTP method from requestor
	Method Method `protobuf:"varint,1,opt,name=method,proto3,enum=httpapi.Method" json:"method,omitempty"`
	// Desired procedure name
	Procedure string `protobuf:"bytes,2,opt,name=procedure,proto3" json:"procedure,omitempty"`
	// JSON payload of the request, if present
	Payload []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	// Query params of the request, if present
	Params map[string]*MultiVal `protobuf:"bytes,4,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Headers won't be automatically translated into the GRPC requests, but they should be delivered in case manual conversion is needed
	Headers map[string]*MultiVal `protobuf:"bytes,5,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{3}
}

func (x *Request) GetMethod() Method {
	if x != nil {
		return x.Method
	}
	return Method_UNKNOWN
}

func (x *Request) GetProcedure() string {
	if x != nil {
		return x.Procedure
	}
	return ""
}

func (x *Request) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Request) GetParams() map[string]*MultiVal {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *Request) GetHeaders() map[string]*MultiVal {
	if x != nil {
		return x.Headers
	}
	return nil
}

// The response to send to the requestor
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Manually set status code rather than rely on error-handling at the proxy
	StatusCode uint32 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	// JSON data to return to the requestor
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// Headers won't be automatically returned from GRPC requests, but they should be accepted in case manual conversion is needed
	WriteHeaders map[string]*MultiVal `protobuf:"bytes,6,rep,name=write_headers,json=writeHeaders,proto3" json:"write_headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{4}
}

func (x *Response) GetStatusCode() uint32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *Response) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Response) GetWriteHeaders() map[string]*MultiVal {
	if x != nil {
		return x.WriteHeaders
	}
	return nil
}

type MultiVal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *MultiVal) Reset() {
	*x = MultiVal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiVal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiVal) ProtoMessage() {}

func (x *MultiVal) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiVal.ProtoReflect.Descriptor instead.
func (*MultiVal) Descriptor() ([]byte, []int) {
	return file_proxy_proto_rawDescGZIP(), []int{5}
}

func (x *MultiVal) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_proxy_proto protoreflect.FileDescriptor

var file_proxy_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x68,
	0x74, 0x74, 0x70, 0x61, 0x70, 0x69, 0x22, 0x70, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x04, 0x69, 0x6e, 0x69,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x61, 0x70,
	0x69, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0xee, 0x01, 0x0a, 0x12, 0x52, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x27, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0f, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x64, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x12, 0x42, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x61, 0x70,
	0x69, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x4d, 0x0a, 0x0c, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x56, 0x61, 0x6c, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2e, 0x0a, 0x10, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xf6, 0x02, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x61, 0x70, 0x69, 0x2e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x34, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x61, 0x70, 0x69,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x37, 0x0a, 0x07,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x4c, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x61, 0x70, 0x69, 0x2e,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x56, 0x61, 0x6c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x4d, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x61, 0x70, 0x69, 0x2e, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x56, 0x61, 0x6c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xe3, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x48, 0x0a, 0x0d, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x1a, 0x52, 0x0a, 0x11, 0x57, 0x72, 0x69, 0x74, 0x65, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x56, 0x61, 0x6c, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x22, 0x0a, 0x08, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x56, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2a, 0x77, 0x0a, 0x06,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x48, 0x45, 0x41, 0x44, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4f, 0x53, 0x54, 0x10, 0x03,
	0x12, 0x07, 0x0a, 0x03, 0x50, 0x55, 0x54, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c,
	0x45, 0x54, 0x45, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54,
	0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x07, 0x12,
	0x09, 0x0a, 0x05, 0x54, 0x52, 0x41, 0x43, 0x45, 0x10, 0x08, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x41,
	0x54, 0x43, 0x48, 0x10, 0x09, 0x32, 0x8f, 0x01, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x65,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x10, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x61, 0x70, 0x69,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a,
	0x0b, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x18, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x61, 0x70, 0x69,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x4c, 0x4b, 0x65, 0x6e, 0x6e, 0x65, 0x64, 0x79, 0x2f,
	0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proxy_proto_rawDescOnce sync.Once
	file_proxy_proto_rawDescData = file_proxy_proto_rawDesc
)

func file_proxy_proto_rawDescGZIP() []byte {
	file_proxy_proto_rawDescOnce.Do(func() {
		file_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_proxy_proto_rawDescData)
	})
	return file_proxy_proto_rawDescData
}

var file_proxy_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proxy_proto_goTypes = []interface{}{
	(Method)(0),                // 0: httpapi.Method
	(*StreamedRequest)(nil),    // 1: httpapi.StreamedRequest
	(*RoutingInformation)(nil), // 2: httpapi.RoutingInformation
	(*StreamedResponse)(nil),   // 3: httpapi.StreamedResponse
	(*Request)(nil),            // 4: httpapi.Request
	(*Response)(nil),           // 5: httpapi.Response
	(*MultiVal)(nil),           // 6: httpapi.MultiVal
	nil,                        // 7: httpapi.RoutingInformation.HeadersEntry
	nil,                        // 8: httpapi.Request.ParamsEntry
	nil,                        // 9: httpapi.Request.HeadersEntry
	nil,                        // 10: httpapi.Response.WriteHeadersEntry
}
var file_proxy_proto_depIdxs = []int32{
	2,  // 0: httpapi.StreamedRequest.init:type_name -> httpapi.RoutingInformation
	0,  // 1: httpapi.RoutingInformation.method:type_name -> httpapi.Method
	7,  // 2: httpapi.RoutingInformation.headers:type_name -> httpapi.RoutingInformation.HeadersEntry
	0,  // 3: httpapi.Request.method:type_name -> httpapi.Method
	8,  // 4: httpapi.Request.params:type_name -> httpapi.Request.ParamsEntry
	9,  // 5: httpapi.Request.headers:type_name -> httpapi.Request.HeadersEntry
	10, // 6: httpapi.Response.write_headers:type_name -> httpapi.Response.WriteHeadersEntry
	6,  // 7: httpapi.RoutingInformation.HeadersEntry.value:type_name -> httpapi.MultiVal
	6,  // 8: httpapi.Request.ParamsEntry.value:type_name -> httpapi.MultiVal
	6,  // 9: httpapi.Request.HeadersEntry.value:type_name -> httpapi.MultiVal
	6,  // 10: httpapi.Response.WriteHeadersEntry.value:type_name -> httpapi.MultiVal
	4,  // 11: httpapi.ExposedService.ProxyUnary:input_type -> httpapi.Request
	1,  // 12: httpapi.ExposedService.ProxyStream:input_type -> httpapi.StreamedRequest
	5,  // 13: httpapi.ExposedService.ProxyUnary:output_type -> httpapi.Response
	3,  // 14: httpapi.ExposedService.ProxyStream:output_type -> httpapi.StreamedResponse
	13, // [13:15] is the sub-list for method output_type
	11, // [11:13] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_proxy_proto_init() }
func file_proxy_proto_init() {
	if File_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamedRequest); i {
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
		file_proxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoutingInformation); i {
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
		file_proxy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamedResponse); i {
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
		file_proxy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proxy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proxy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiVal); i {
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
	file_proxy_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*StreamedRequest_Init)(nil),
		(*StreamedRequest_Request)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proxy_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proxy_proto_goTypes,
		DependencyIndexes: file_proxy_proto_depIdxs,
		EnumInfos:         file_proxy_proto_enumTypes,
		MessageInfos:      file_proxy_proto_msgTypes,
	}.Build()
	File_proxy_proto = out.File
	file_proxy_proto_rawDesc = nil
	file_proxy_proto_goTypes = nil
	file_proxy_proto_depIdxs = nil
}
