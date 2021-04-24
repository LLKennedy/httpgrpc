package proxy

import (
	"context"
	"reflect"

	"github.com/LLKennedy/mercury/httpapi"
	"google.golang.org/grpc"
)

type apiMethodPattern int

const (
	apiMethodPatternUnknown apiMethodPattern = iota
	apiMethodPatternStructStruct
	apiMethodPatternStreamStruct
	apiMethodPatternStructStream
	apiMethodPatternStreamStream
)

// We use defaultServer in the case that s is nil
var defaultServer = &Server{}

// ExceptionHandler is an exception handler function
type ExceptionHandler func(ctx context.Context, req *httpapi.Request) (handled bool, res *httpapi.Response, err error)

// Server is an HTTP to GRPC proxy server
type Server struct {
	rawAPI           interface{} // unparsed API
	grpcServer       *grpc.Server
	api              map[string]map[string]apiMethod // the api of innerServer
	innerServer      interface{}                     // the actual protobuf endpoints we want to use
	exceptionHandler ExceptionHandler
	httpapi.UnimplementedExposedServiceServer
	skipForwardingMetadata bool
}

type apiMethod struct {
	pattern    apiMethodPattern
	reflection reflect.Method
	value      reflect.Value
}

func (s *Server) getGrpcServer() *grpc.Server {
	if s == nil {
		return defaultServer.grpcServer
	}
	return s.grpcServer
}

func (s *Server) setGrpcServer(in *grpc.Server) {
	if s == nil {
		defaultServer.grpcServer = in
		return
	}
	s.grpcServer = in
}

func (s *Server) getAPI() map[string]map[string]apiMethod {
	if s == nil {
		return defaultServer.api
	}
	return s.api
}

func (s *Server) setAPI(in map[string]map[string]apiMethod) {
	if s == nil {
		defaultServer.api = in
		return
	}
	s.api = in
}

func (s *Server) getInnerServer() interface{} {
	if s == nil {
		return defaultServer.innerServer
	}
	return s.innerServer
}

func (s *Server) setInnerServer(in interface{}) {
	if s == nil {
		defaultServer.innerServer = in
		return
	}
	s.innerServer = in
}
func (s *Server) getSkipForwardingMetadata() bool {
	if s == nil {
		return defaultServer.skipForwardingMetadata
	}
	return s.skipForwardingMetadata
}

func (s *Server) setSkipForwardingMetadata(in bool) {
	if s == nil {
		defaultServer.skipForwardingMetadata = in
	}
	s.skipForwardingMetadata = in
}

func (s *Server) handleExceptions(ctx context.Context, req *httpapi.Request) (handled bool, res *httpapi.Response, err error) {
	if s == nil || s.exceptionHandler == nil {
		handled = false
		return
	}
	return s.exceptionHandler(ctx, req)
}
