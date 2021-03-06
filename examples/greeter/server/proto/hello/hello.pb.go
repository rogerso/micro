// Code generated by protoc-gen-go.
// source: examples/greeter/server/proto/hello/hello.proto
// DO NOT EDIT!

/*
Package go_micro_srv_greeter is a generated protocol buffer package.

It is generated from these files:
	examples/greeter/server/proto/hello/hello.proto

It has these top-level messages:
	Request
	Response
*/
package go_micro_srv_greeter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Response struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.greeter.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.greeter.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Say service

type SayClient interface {
	Hello(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type sayClient struct {
	c           client.Client
	serviceName string
}

func NewSayClient(serviceName string, c client.Client) SayClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.greeter"
	}
	return &sayClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *sayClient) Hello(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Say.Hello", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Say service

type SayHandler interface {
	Hello(context.Context, *Request, *Response) error
}

func RegisterSayHandler(s server.Server, hdlr SayHandler) {
	s.Handle(s.NewHandler(&Say{hdlr}))
}

type Say struct {
	SayHandler
}

func (h *Say) Hello(ctx context.Context, in *Request, out *Response) error {
	return h.SayHandler.Hello(ctx, in, out)
}

var fileDescriptor0 = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xd2, 0x4f, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x4f, 0x2f, 0x4a, 0x4d, 0x2d, 0x49, 0x2d, 0xd2, 0x2f, 0x4e, 0x2d,
	0x2a, 0x03, 0x52, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0xfa, 0x19, 0xa9, 0x39, 0x39, 0x50, 0x52, 0x0f,
	0x2c, 0x22, 0x24, 0x92, 0x9e, 0xaf, 0x97, 0x9b, 0x99, 0x5c, 0x94, 0xaf, 0x57, 0x5c, 0x54, 0xa6,
	0x07, 0xd5, 0xa4, 0x24, 0xce, 0xc5, 0x1e, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0xc4, 0xc3,
	0xc5, 0x92, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x09, 0x94, 0xe0, 0x08, 0x4a,
	0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0xe2, 0xe6, 0x62, 0xce, 0x2d, 0x4e, 0x87, 0x48, 0x18,
	0xf9, 0x73, 0x31, 0x07, 0x27, 0x56, 0x0a, 0x79, 0x70, 0xb1, 0x7a, 0x80, 0x4c, 0x17, 0x92, 0xd5,
	0xc3, 0x66, 0xb0, 0x1e, 0xd4, 0x54, 0x29, 0x39, 0x5c, 0xd2, 0x10, 0xb3, 0x95, 0x18, 0x92, 0xd8,
	0xc0, 0xee, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x03, 0x56, 0x54, 0x80, 0xd2, 0x00, 0x00,
	0x00,
}
