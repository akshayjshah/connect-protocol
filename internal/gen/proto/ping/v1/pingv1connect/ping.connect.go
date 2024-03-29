// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/ping/v1/ping.proto

package pingv1connect

import (
	v1 "cncf/protocol/internal/gen/proto/ping/v1"
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// PingServiceName is the fully-qualified name of the PingService service.
	PingServiceName = "ping.v1.PingService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PingServicePingProcedure is the fully-qualified name of the PingService's Ping RPC.
	PingServicePingProcedure = "/ping.v1.PingService/Ping"
	// PingServicePingsProcedure is the fully-qualified name of the PingService's Pings RPC.
	PingServicePingsProcedure = "/ping.v1.PingService/Pings"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	pingServiceServiceDescriptor     = v1.File_proto_ping_v1_ping_proto.Services().ByName("PingService")
	pingServicePingMethodDescriptor  = pingServiceServiceDescriptor.Methods().ByName("Ping")
	pingServicePingsMethodDescriptor = pingServiceServiceDescriptor.Methods().ByName("Pings")
)

// PingServiceClient is a client for the ping.v1.PingService service.
type PingServiceClient interface {
	Ping(context.Context, *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error)
	Pings(context.Context, *connect.Request[v1.PingsRequest]) (*connect.ServerStreamForClient[v1.PingsResponse], error)
}

// NewPingServiceClient constructs a client for the ping.v1.PingService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPingServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) PingServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &pingServiceClient{
		ping: connect.NewClient[v1.PingRequest, v1.PingResponse](
			httpClient,
			baseURL+PingServicePingProcedure,
			connect.WithSchema(pingServicePingMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		pings: connect.NewClient[v1.PingsRequest, v1.PingsResponse](
			httpClient,
			baseURL+PingServicePingsProcedure,
			connect.WithSchema(pingServicePingsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// pingServiceClient implements PingServiceClient.
type pingServiceClient struct {
	ping  *connect.Client[v1.PingRequest, v1.PingResponse]
	pings *connect.Client[v1.PingsRequest, v1.PingsResponse]
}

// Ping calls ping.v1.PingService.Ping.
func (c *pingServiceClient) Ping(ctx context.Context, req *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error) {
	return c.ping.CallUnary(ctx, req)
}

// Pings calls ping.v1.PingService.Pings.
func (c *pingServiceClient) Pings(ctx context.Context, req *connect.Request[v1.PingsRequest]) (*connect.ServerStreamForClient[v1.PingsResponse], error) {
	return c.pings.CallServerStream(ctx, req)
}

// PingServiceHandler is an implementation of the ping.v1.PingService service.
type PingServiceHandler interface {
	Ping(context.Context, *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error)
	Pings(context.Context, *connect.Request[v1.PingsRequest], *connect.ServerStream[v1.PingsResponse]) error
}

// NewPingServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPingServiceHandler(svc PingServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	pingServicePingHandler := connect.NewUnaryHandler(
		PingServicePingProcedure,
		svc.Ping,
		connect.WithSchema(pingServicePingMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	pingServicePingsHandler := connect.NewServerStreamHandler(
		PingServicePingsProcedure,
		svc.Pings,
		connect.WithSchema(pingServicePingsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/ping.v1.PingService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PingServicePingProcedure:
			pingServicePingHandler.ServeHTTP(w, r)
		case PingServicePingsProcedure:
			pingServicePingsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPingServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPingServiceHandler struct{}

func (UnimplementedPingServiceHandler) Ping(context.Context, *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ping.v1.PingService.Ping is not implemented"))
}

func (UnimplementedPingServiceHandler) Pings(context.Context, *connect.Request[v1.PingsRequest], *connect.ServerStream[v1.PingsResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("ping.v1.PingService.Pings is not implemented"))
}
