// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/paymentApi/paymentApi.proto

/*
Package go_micro_api_paymentApi is a generated protocol buffer package.

It is generated from these files:
	proto/paymentApi/paymentApi.proto

It has these top-level messages:
	Pair
	Request
	Response
*/
package go_micro_api_paymentApi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Pair struct {
	Key    string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Values []string `protobuf:"bytes,2,rep,name=values" json:"values,omitempty"`
}

func (m *Pair) Reset()                    { *m = Pair{} }
func (m *Pair) String() string            { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()               {}
func (*Pair) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Pair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Pair) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type Request struct {
	Method string           `protobuf:"bytes,1,opt,name=method" json:"method,omitempty"`
	Path   string           `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Header map[string]*Pair `protobuf:"bytes,3,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Get    map[string]*Pair `protobuf:"bytes,4,rep,name=get" json:"get,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Post   map[string]*Pair `protobuf:"bytes,5,rep,name=post" json:"post,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body   string           `protobuf:"bytes,6,opt,name=body" json:"body,omitempty"`
	Url    string           `protobuf:"bytes,7,opt,name=url" json:"url,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Request) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Request) GetGet() map[string]*Pair {
	if m != nil {
		return m.Get
	}
	return nil
}

func (m *Request) GetPost() map[string]*Pair {
	if m != nil {
		return m.Post
	}
	return nil
}

func (m *Request) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Request) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type Response struct {
	StatusCode int32            `protobuf:"varint,1,opt,name=statusCode" json:"statusCode,omitempty"`
	Header     map[string]*Pair `protobuf:"bytes,2,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body       string           `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *Response) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Response) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*Pair)(nil), "go.micro.api.paymentApi.Pair")
	proto.RegisterType((*Request)(nil), "go.micro.api.paymentApi.Request")
	proto.RegisterType((*Response)(nil), "go.micro.api.paymentApi.Response")
}

func init() { proto.RegisterFile("proto/paymentApi/paymentApi.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0xb5, 0x1f, 0x14, 0x18, 0x3c, 0x98, 0x3d, 0xe8, 0x86, 0x44, 0x53, 0x7a, 0x42, 0xa3, 0xd5,
	0xc0, 0xc5, 0x68, 0x62, 0x62, 0x94, 0xe8, 0xb1, 0xd9, 0x04, 0xe3, 0x75, 0xa1, 0x2b, 0x10, 0x4b,
	0xb7, 0x76, 0xb7, 0x26, 0xfd, 0x45, 0xfe, 0x2c, 0xff, 0x8a, 0xd9, 0x6d, 0xa9, 0x1c, 0xc0, 0x5e,
	0xf0, 0x36, 0x3b, 0x33, 0xef, 0x65, 0xe6, 0xbd, 0x59, 0xe8, 0x25, 0x29, 0x97, 0xfc, 0x32, 0xa1,
	0xf9, 0x92, 0xc5, 0xf2, 0x3e, 0x59, 0xac, 0x85, 0xbe, 0xae, 0xa1, 0xa3, 0x19, 0xf7, 0x97, 0x8b,
	0x69, 0xca, 0x7d, 0xaa, 0x72, 0x55, 0xd9, 0xbb, 0x02, 0x3b, 0xa0, 0x8b, 0x14, 0x1d, 0x80, 0xf5,
	0xce, 0x72, 0x6c, 0xb8, 0x46, 0xbf, 0x4d, 0x54, 0x88, 0x0e, 0xc1, 0xf9, 0xa4, 0x51, 0xc6, 0x04,
	0x36, 0x5d, 0xab, 0xdf, 0x26, 0xe5, 0xcb, 0xfb, 0xb2, 0xa1, 0x49, 0xd8, 0x47, 0xc6, 0x84, 0x54,
	0x3d, 0x4b, 0x26, 0xe7, 0x3c, 0x2c, 0x81, 0xe5, 0x0b, 0x21, 0xb0, 0x13, 0x2a, 0xe7, 0xd8, 0xd4,
	0x59, 0x1d, 0xa3, 0x47, 0x70, 0xe6, 0x8c, 0x86, 0x2c, 0xc5, 0x96, 0x6b, 0xf5, 0x3b, 0x83, 0x73,
	0x7f, 0xcb, 0x4c, 0x7e, 0xc9, 0xee, 0x3f, 0xeb, 0xf6, 0x51, 0x2c, 0xd3, 0x9c, 0x94, 0x58, 0x74,
	0x0b, 0xd6, 0x8c, 0x49, 0x6c, 0x6b, 0x8a, 0xd3, 0x5a, 0x8a, 0x27, 0x26, 0x0b, 0xbc, 0x42, 0xa1,
	0x3b, 0xb0, 0x13, 0x2e, 0x24, 0x6e, 0x68, 0xf4, 0x59, 0x2d, 0x3a, 0xe0, 0xa2, 0x84, 0x6b, 0x9c,
	0x5a, 0x6b, 0xc2, 0xc3, 0x1c, 0x3b, 0xc5, 0x5a, 0x2a, 0x56, 0xc2, 0x65, 0x69, 0x84, 0x9b, 0x85,
	0x70, 0x59, 0x1a, 0x75, 0x5f, 0xa1, 0xb3, 0x36, 0xf9, 0x06, 0x65, 0x87, 0xd0, 0xd0, 0x5a, 0x6a,
	0x79, 0x3a, 0x83, 0xe3, 0xad, 0x73, 0x28, 0x67, 0x48, 0xd1, 0x7b, 0x63, 0x5e, 0x1b, 0xdd, 0x31,
	0xb4, 0x56, 0x0b, 0xed, 0x92, 0xf6, 0x05, 0xda, 0xd5, 0xa6, 0x3b, 0xe4, 0xf5, 0xbe, 0x0d, 0x68,
	0x11, 0x26, 0x12, 0x1e, 0x0b, 0x86, 0x4e, 0x00, 0x84, 0xa4, 0x32, 0x13, 0x0f, 0x3c, 0x64, 0x9a,
	0xbe, 0x41, 0xd6, 0x32, 0x68, 0x54, 0x9d, 0x87, 0xa9, 0xdd, 0xb9, 0xf8, 0xc3, 0x9d, 0x82, 0x72,
	0xe3, 0x7d, 0xac, 0x2c, 0xb2, 0x7e, 0x2d, 0xfa, 0x3f, 0x43, 0x06, 0x53, 0x80, 0xa0, 0xaa, 0xa2,
	0x31, 0xec, 0x07, 0x34, 0x0f, 0x68, 0x44, 0xd8, 0x5b, 0x16, 0x87, 0xc8, 0xad, 0x3b, 0xb0, 0x6e,
	0xaf, 0x76, 0x49, 0x6f, 0x6f, 0xe2, 0xe8, 0x2f, 0x3c, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x27,
	0x4f, 0x8c, 0xe1, 0xe7, 0x03, 0x00, 0x00,
}