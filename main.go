package main

import (
	"context"
	"errors"
	"net/http"

	pingv1 "cncf/protocol/internal/gen/proto/ping/v1"
	"cncf/protocol/internal/gen/proto/ping/v1/pingv1connect"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type pinger struct{}

func (p *pinger) Ping(
	ctx context.Context,
	req *connect.Request[pingv1.PingRequest],
) (*connect.Response[pingv1.PingResponse], error) {
	if req.Msg.Text == "explode" {
		err := connect.NewError(connect.CodeInvalidArgument, errors.New("boom"))
		if d, e := connect.NewErrorDetail(timestamppb.Now()); e == nil {
			err.AddDetail(d)
		}
		return nil, err
	}
	return connect.NewResponse(&pingv1.PingResponse{
		Text: req.Msg.Text,
	}), nil
}

func (p *pinger) Pings(
	ctx context.Context,
	req *connect.Request[pingv1.PingsRequest],
	stream *connect.ServerStream[pingv1.PingsResponse],
) error {
	if err := stream.Send(&pingv1.PingsResponse{Text: req.Msg.Text}); err != nil {
		return err
	}
	err := connect.NewError(connect.CodeInvalidArgument, errors.New("boom"))
	if d, e := connect.NewErrorDetail(timestamppb.Now()); e == nil {
		err.AddDetail(d)
	}
	return err
}

func main() {
	reflector := grpcreflect.NewStaticReflector(
		pingv1connect.PingServiceName,
	)
	mux := http.NewServeMux()
	mux.Handle(pingv1connect.NewPingServiceHandler(&pinger{}))
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
