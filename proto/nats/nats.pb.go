// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/nats/nats.proto

package nats

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Request struct {
	// 手机号
	Mobile               string   `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_86369f7f3ad51865, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

type Response struct {
	Valid                bool     `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Uuid                 string   `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_86369f7f3ad51865, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Response) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "nats.Request")
	proto.RegisterType((*Response)(nil), "nats.Response")
}

func init() { proto.RegisterFile("proto/nats/nats.proto", fileDescriptor_86369f7f3ad51865) }

var fileDescriptor_86369f7f3ad51865 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x4b, 0x2c, 0x29, 0x06, 0x13, 0x7a, 0x60, 0xbe, 0x10, 0x0b, 0x88, 0xad, 0xa4,
	0xc8, 0xc5, 0x1e, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc6, 0xc5, 0x96, 0x9b, 0x9f,
	0x94, 0x99, 0x93, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe5, 0x29, 0xf9, 0x71, 0x71,
	0x04, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x89, 0x70, 0xb1, 0x96, 0x25, 0xe6, 0x64,
	0xa6, 0x80, 0x95, 0x70, 0x04, 0x41, 0x38, 0x42, 0x12, 0x5c, 0xec, 0xb9, 0xa9, 0xc5, 0xc5, 0x89,
	0xe9, 0xa9, 0x12, 0x4c, 0x60, 0xad, 0x30, 0xae, 0x90, 0x10, 0x17, 0x4b, 0x69, 0x69, 0x66, 0x8a,
	0x04, 0x33, 0x58, 0x18, 0xcc, 0x36, 0x32, 0xe7, 0x62, 0xf1, 0x4b, 0x2c, 0x29, 0x16, 0xd2, 0xe7,
	0xe2, 0xf1, 0x05, 0xdb, 0x10, 0x96, 0x5a, 0x94, 0x99, 0x56, 0x29, 0xc4, 0xab, 0x07, 0x76, 0x1d,
	0xd4, 0x39, 0x52, 0x7c, 0x30, 0x2e, 0xc4, 0x6a, 0x25, 0x86, 0x24, 0x36, 0xb0, 0xc3, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xf4, 0x15, 0x15, 0xd1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Nats service

type NatsClient interface {
	// 手机验证码
	MobileVerify(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type natsClient struct {
	c           client.Client
	serviceName string
}

func NewNatsClient(serviceName string, c client.Client) NatsClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "nats"
	}
	return &natsClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *natsClient) MobileVerify(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Nats.MobileVerify", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Nats service

type NatsHandler interface {
	// 手机验证码
	MobileVerify(context.Context, *Request, *Response) error
}

func RegisterNatsHandler(s server.Server, hdlr NatsHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Nats{hdlr}, opts...))
}

type Nats struct {
	NatsHandler
}

func (h *Nats) MobileVerify(ctx context.Context, in *Request, out *Response) error {
	return h.NatsHandler.MobileVerify(ctx, in, out)
}
