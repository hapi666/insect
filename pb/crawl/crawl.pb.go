// Code generated by protoc-gen-go. DO NOT EDIT.
// source: crawl.proto

package crawl

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

type CrawlRequest struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Timeout              int64    `protobuf:"varint,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Headers              []*KV    `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CrawlRequest) Reset()         { *m = CrawlRequest{} }
func (m *CrawlRequest) String() string { return proto.CompactTextString(m) }
func (*CrawlRequest) ProtoMessage()    {}
func (*CrawlRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5877e6cf3ac6969d, []int{0}
}

func (m *CrawlRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CrawlRequest.Unmarshal(m, b)
}
func (m *CrawlRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CrawlRequest.Marshal(b, m, deterministic)
}
func (m *CrawlRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrawlRequest.Merge(m, src)
}
func (m *CrawlRequest) XXX_Size() int {
	return xxx_messageInfo_CrawlRequest.Size(m)
}
func (m *CrawlRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CrawlRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CrawlRequest proto.InternalMessageInfo

func (m *CrawlRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *CrawlRequest) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *CrawlRequest) GetHeaders() []*KV {
	if m != nil {
		return m.Headers
	}
	return nil
}

type KV struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KV) Reset()         { *m = KV{} }
func (m *KV) String() string { return proto.CompactTextString(m) }
func (*KV) ProtoMessage()    {}
func (*KV) Descriptor() ([]byte, []int) {
	return fileDescriptor_5877e6cf3ac6969d, []int{1}
}

func (m *KV) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KV.Unmarshal(m, b)
}
func (m *KV) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KV.Marshal(b, m, deterministic)
}
func (m *KV) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KV.Merge(m, src)
}
func (m *KV) XXX_Size() int {
	return xxx_messageInfo_KV.Size(m)
}
func (m *KV) XXX_DiscardUnknown() {
	xxx_messageInfo_KV.DiscardUnknown(m)
}

var xxx_messageInfo_KV proto.InternalMessageInfo

func (m *KV) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KV) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type CrawlResponse struct {
	Signal               string   `protobuf:"bytes,1,opt,name=signal,proto3" json:"signal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CrawlResponse) Reset()         { *m = CrawlResponse{} }
func (m *CrawlResponse) String() string { return proto.CompactTextString(m) }
func (*CrawlResponse) ProtoMessage()    {}
func (*CrawlResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5877e6cf3ac6969d, []int{2}
}

func (m *CrawlResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CrawlResponse.Unmarshal(m, b)
}
func (m *CrawlResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CrawlResponse.Marshal(b, m, deterministic)
}
func (m *CrawlResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrawlResponse.Merge(m, src)
}
func (m *CrawlResponse) XXX_Size() int {
	return xxx_messageInfo_CrawlResponse.Size(m)
}
func (m *CrawlResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CrawlResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CrawlResponse proto.InternalMessageInfo

func (m *CrawlResponse) GetSignal() string {
	if m != nil {
		return m.Signal
	}
	return ""
}

func init() {
	proto.RegisterType((*CrawlRequest)(nil), "crawl.CrawlRequest")
	proto.RegisterType((*KV)(nil), "crawl.KV")
	proto.RegisterType((*CrawlResponse)(nil), "crawl.CrawlResponse")
}

func init() { proto.RegisterFile("crawl.proto", fileDescriptor_5877e6cf3ac6969d) }

var fileDescriptor_5877e6cf3ac6969d = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2e, 0x4a, 0x2c,
	0xcf, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0xe2, 0xb9, 0x78, 0x9c,
	0x41, 0x8c, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x01, 0x2e, 0xe6, 0xd2, 0xa2, 0x1c,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x10, 0x53, 0x48, 0x82, 0x8b, 0xbd, 0x24, 0x33, 0x37,
	0x35, 0xbf, 0xb4, 0x44, 0x82, 0x49, 0x81, 0x51, 0x83, 0x39, 0x08, 0xc6, 0x15, 0x52, 0xe6, 0x62,
	0xcf, 0x48, 0x4d, 0x4c, 0x49, 0x2d, 0x2a, 0x96, 0x60, 0x56, 0x60, 0xd6, 0xe0, 0x36, 0xe2, 0xd4,
	0x83, 0xd8, 0xe0, 0x1d, 0x16, 0x04, 0x93, 0x51, 0xd2, 0xe1, 0x62, 0xf2, 0x0e, 0x03, 0x19, 0x9b,
	0x9d, 0x5a, 0x09, 0x33, 0x36, 0x3b, 0xb5, 0x52, 0x48, 0x84, 0x8b, 0xb5, 0x2c, 0x31, 0xa7, 0x34,
	0x15, 0x6c, 0x28, 0x67, 0x10, 0x84, 0xa3, 0xa4, 0xce, 0xc5, 0x0b, 0x75, 0x4e, 0x71, 0x41, 0x7e,
	0x5e, 0x71, 0xaa, 0x90, 0x18, 0x17, 0x5b, 0x71, 0x66, 0x7a, 0x5e, 0x22, 0xcc, 0x49, 0x50, 0x9e,
	0x91, 0x2d, 0x17, 0x2b, 0x58, 0xa1, 0x90, 0x09, 0x8c, 0x21, 0x0c, 0xb5, 0x1c, 0xd9, 0x3b, 0x52,
	0x22, 0xa8, 0x82, 0x10, 0x43, 0x95, 0x18, 0x92, 0xd8, 0xc0, 0x81, 0x60, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x2d, 0x2b, 0x97, 0x19, 0x13, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CrawlClient is the client API for Crawl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CrawlClient interface {
	Crawl(ctx context.Context, in *CrawlRequest, opts ...grpc.CallOption) (*CrawlResponse, error)
}

type crawlClient struct {
	cc *grpc.ClientConn
}

func NewCrawlClient(cc *grpc.ClientConn) CrawlClient {
	return &crawlClient{cc}
}

func (c *crawlClient) Crawl(ctx context.Context, in *CrawlRequest, opts ...grpc.CallOption) (*CrawlResponse, error) {
	out := new(CrawlResponse)
	err := c.cc.Invoke(ctx, "/crawl.Crawl/Crawl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrawlServer is the server API for Crawl service.
type CrawlServer interface {
	Crawl(context.Context, *CrawlRequest) (*CrawlResponse, error)
}

// UnimplementedCrawlServer can be embedded to have forward compatible implementations.
type UnimplementedCrawlServer struct {
}

func (*UnimplementedCrawlServer) Crawl(ctx context.Context, req *CrawlRequest) (*CrawlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Crawl not implemented")
}

func RegisterCrawlServer(s *grpc.Server, srv CrawlServer) {
	s.RegisterService(&_Crawl_serviceDesc, srv)
}

func _Crawl_Crawl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CrawlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrawlServer).Crawl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crawl.Crawl/Crawl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrawlServer).Crawl(ctx, req.(*CrawlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Crawl_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crawl.Crawl",
	HandlerType: (*CrawlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Crawl",
			Handler:    _Crawl_Crawl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crawl.proto",
}
