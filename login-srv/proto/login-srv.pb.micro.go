// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/login-srv.proto

package loginsrv

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for LoginSrv service

func NewLoginSrvEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for LoginSrv service

type LoginSrvService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (LoginSrv_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (LoginSrv_PingPongService, error)
}

type loginSrvService struct {
	c    client.Client
	name string
}

func NewLoginSrvService(name string, c client.Client) LoginSrvService {
	return &loginSrvService{
		c:    c,
		name: name,
	}
}

func (c *loginSrvService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "LoginSrv.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginSrvService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (LoginSrv_StreamService, error) {
	req := c.c.NewRequest(c.name, "LoginSrv.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &loginSrvServiceStream{stream}, nil
}

type LoginSrv_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type loginSrvServiceStream struct {
	stream client.Stream
}

func (x *loginSrvServiceStream) Close() error {
	return x.stream.Close()
}

func (x *loginSrvServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *loginSrvServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *loginSrvServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *loginSrvServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *loginSrvService) PingPong(ctx context.Context, opts ...client.CallOption) (LoginSrv_PingPongService, error) {
	req := c.c.NewRequest(c.name, "LoginSrv.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &loginSrvServicePingPong{stream}, nil
}

type LoginSrv_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type loginSrvServicePingPong struct {
	stream client.Stream
}

func (x *loginSrvServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *loginSrvServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *loginSrvServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *loginSrvServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *loginSrvServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *loginSrvServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for LoginSrv service

type LoginSrvHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, LoginSrv_StreamStream) error
	PingPong(context.Context, LoginSrv_PingPongStream) error
}

func RegisterLoginSrvHandler(s server.Server, hdlr LoginSrvHandler, opts ...server.HandlerOption) error {
	type loginSrv interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type LoginSrv struct {
		loginSrv
	}
	h := &loginSrvHandler{hdlr}
	return s.Handle(s.NewHandler(&LoginSrv{h}, opts...))
}

type loginSrvHandler struct {
	LoginSrvHandler
}

func (h *loginSrvHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.LoginSrvHandler.Call(ctx, in, out)
}

func (h *loginSrvHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.LoginSrvHandler.Stream(ctx, m, &loginSrvStreamStream{stream})
}

type LoginSrv_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type loginSrvStreamStream struct {
	stream server.Stream
}

func (x *loginSrvStreamStream) Close() error {
	return x.stream.Close()
}

func (x *loginSrvStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *loginSrvStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *loginSrvStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *loginSrvStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *loginSrvHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.LoginSrvHandler.PingPong(ctx, &loginSrvPingPongStream{stream})
}

type LoginSrv_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type loginSrvPingPongStream struct {
	stream server.Stream
}

func (x *loginSrvPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *loginSrvPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *loginSrvPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *loginSrvPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *loginSrvPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *loginSrvPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}