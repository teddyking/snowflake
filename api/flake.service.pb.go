// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flake.service.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FlakeListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlakeListRequest) Reset()         { *m = FlakeListRequest{} }
func (m *FlakeListRequest) String() string { return proto.CompactTextString(m) }
func (*FlakeListRequest) ProtoMessage()    {}
func (*FlakeListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_flake_service_1a2b33afea705e78, []int{0}
}
func (m *FlakeListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlakeListRequest.Unmarshal(m, b)
}
func (m *FlakeListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlakeListRequest.Marshal(b, m, deterministic)
}
func (dst *FlakeListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlakeListRequest.Merge(dst, src)
}
func (m *FlakeListRequest) XXX_Size() int {
	return xxx_messageInfo_FlakeListRequest.Size(m)
}
func (m *FlakeListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FlakeListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FlakeListRequest proto.InternalMessageInfo

type FlakeListResponse struct {
	Tests                []*Test  `protobuf:"bytes,1,rep,name=Tests" json:"Tests,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlakeListResponse) Reset()         { *m = FlakeListResponse{} }
func (m *FlakeListResponse) String() string { return proto.CompactTextString(m) }
func (*FlakeListResponse) ProtoMessage()    {}
func (*FlakeListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_flake_service_1a2b33afea705e78, []int{1}
}
func (m *FlakeListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlakeListResponse.Unmarshal(m, b)
}
func (m *FlakeListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlakeListResponse.Marshal(b, m, deterministic)
}
func (dst *FlakeListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlakeListResponse.Merge(dst, src)
}
func (m *FlakeListResponse) XXX_Size() int {
	return xxx_messageInfo_FlakeListResponse.Size(m)
}
func (m *FlakeListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FlakeListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FlakeListResponse proto.InternalMessageInfo

func (m *FlakeListResponse) GetTests() []*Test {
	if m != nil {
		return m.Tests
	}
	return nil
}

func init() {
	proto.RegisterType((*FlakeListRequest)(nil), "api.FlakeListRequest")
	proto.RegisterType((*FlakeListResponse)(nil), "api.FlakeListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Flake service

type FlakeClient interface {
	List(ctx context.Context, in *FlakeListRequest, opts ...grpc.CallOption) (*FlakeListResponse, error)
}

type flakeClient struct {
	cc *grpc.ClientConn
}

func NewFlakeClient(cc *grpc.ClientConn) FlakeClient {
	return &flakeClient{cc}
}

func (c *flakeClient) List(ctx context.Context, in *FlakeListRequest, opts ...grpc.CallOption) (*FlakeListResponse, error) {
	out := new(FlakeListResponse)
	err := grpc.Invoke(ctx, "/api.Flake/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Flake service

type FlakeServer interface {
	List(context.Context, *FlakeListRequest) (*FlakeListResponse, error)
}

func RegisterFlakeServer(s *grpc.Server, srv FlakeServer) {
	s.RegisterService(&_Flake_serviceDesc, srv)
}

func _Flake_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlakeListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlakeServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Flake/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlakeServer).List(ctx, req.(*FlakeListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Flake_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Flake",
	HandlerType: (*FlakeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Flake_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flake.service.proto",
}

func init() { proto.RegisterFile("flake.service.proto", fileDescriptor_flake_service_1a2b33afea705e78) }

var fileDescriptor_flake_service_1a2b33afea705e78 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xcb, 0x49, 0xcc,
	0x4e, 0xd5, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x4e, 0x2c, 0xc8, 0x94, 0xe2, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x86, 0x88, 0x28, 0x09, 0x71,
	0x09, 0xb8, 0x81, 0x14, 0xfa, 0x64, 0x16, 0x97, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x28,
	0x99, 0x70, 0x09, 0x22, 0x89, 0x15, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0xc9, 0x73, 0xb1, 0x86,
	0xa4, 0x16, 0x97, 0x14, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0x71, 0xea, 0x25, 0x16, 0x64,
	0xea, 0x81, 0x44, 0x82, 0x20, 0xe2, 0x46, 0x0e, 0x5c, 0xac, 0x60, 0x5d, 0x42, 0xe6, 0x5c, 0x2c,
	0x20, 0x9d, 0x42, 0xa2, 0x60, 0x25, 0xe8, 0xa6, 0x4b, 0x89, 0xa1, 0x0b, 0x43, 0x2c, 0x50, 0x62,
	0x48, 0x62, 0x03, 0x3b, 0xc9, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x37, 0x59, 0x0d, 0xbb,
	0x00, 0x00, 0x00,
}
