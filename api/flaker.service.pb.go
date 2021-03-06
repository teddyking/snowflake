// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flaker.service.proto

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

type FlakerListReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlakerListReq) Reset()         { *m = FlakerListReq{} }
func (m *FlakerListReq) String() string { return proto.CompactTextString(m) }
func (*FlakerListReq) ProtoMessage()    {}
func (*FlakerListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_flaker_service_dc8cad22e71b50ae, []int{0}
}
func (m *FlakerListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlakerListReq.Unmarshal(m, b)
}
func (m *FlakerListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlakerListReq.Marshal(b, m, deterministic)
}
func (dst *FlakerListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlakerListReq.Merge(dst, src)
}
func (m *FlakerListReq) XXX_Size() int {
	return xxx_messageInfo_FlakerListReq.Size(m)
}
func (m *FlakerListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FlakerListReq.DiscardUnknown(m)
}

var xxx_messageInfo_FlakerListReq proto.InternalMessageInfo

type FlakerListRes struct {
	Flakes               []*Flake `protobuf:"bytes,1,rep,name=flakes" json:"flakes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlakerListRes) Reset()         { *m = FlakerListRes{} }
func (m *FlakerListRes) String() string { return proto.CompactTextString(m) }
func (*FlakerListRes) ProtoMessage()    {}
func (*FlakerListRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_flaker_service_dc8cad22e71b50ae, []int{1}
}
func (m *FlakerListRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlakerListRes.Unmarshal(m, b)
}
func (m *FlakerListRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlakerListRes.Marshal(b, m, deterministic)
}
func (dst *FlakerListRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlakerListRes.Merge(dst, src)
}
func (m *FlakerListRes) XXX_Size() int {
	return xxx_messageInfo_FlakerListRes.Size(m)
}
func (m *FlakerListRes) XXX_DiscardUnknown() {
	xxx_messageInfo_FlakerListRes.DiscardUnknown(m)
}

var xxx_messageInfo_FlakerListRes proto.InternalMessageInfo

func (m *FlakerListRes) GetFlakes() []*Flake {
	if m != nil {
		return m.Flakes
	}
	return nil
}

func init() {
	proto.RegisterType((*FlakerListReq)(nil), "api.FlakerListReq")
	proto.RegisterType((*FlakerListRes)(nil), "api.FlakerListRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Flaker service

type FlakerClient interface {
	List(ctx context.Context, in *FlakerListReq, opts ...grpc.CallOption) (*FlakerListRes, error)
}

type flakerClient struct {
	cc *grpc.ClientConn
}

func NewFlakerClient(cc *grpc.ClientConn) FlakerClient {
	return &flakerClient{cc}
}

func (c *flakerClient) List(ctx context.Context, in *FlakerListReq, opts ...grpc.CallOption) (*FlakerListRes, error) {
	out := new(FlakerListRes)
	err := grpc.Invoke(ctx, "/api.Flaker/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Flaker service

type FlakerServer interface {
	List(context.Context, *FlakerListReq) (*FlakerListRes, error)
}

func RegisterFlakerServer(s *grpc.Server, srv FlakerServer) {
	s.RegisterService(&_Flaker_serviceDesc, srv)
}

func _Flaker_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlakerListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlakerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Flaker/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlakerServer).List(ctx, req.(*FlakerListReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Flaker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Flaker",
	HandlerType: (*FlakerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Flaker_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flaker.service.proto",
}

func init() {
	proto.RegisterFile("flaker.service.proto", fileDescriptor_flaker_service_dc8cad22e71b50ae)
}

var fileDescriptor_flaker_service_dc8cad22e71b50ae = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0xcb, 0x49, 0xcc,
	0x4e, 0x2d, 0xd2, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x94, 0xe2, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x86, 0x88, 0x28, 0xf1,
	0x73, 0xf1, 0xba, 0x81, 0x55, 0xfa, 0x64, 0x16, 0x97, 0x04, 0xa5, 0x16, 0x2a, 0x19, 0xa3, 0x0a,
	0x14, 0x0b, 0x29, 0x71, 0xb1, 0x81, 0xcd, 0x2a, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0xe2,
	0xd2, 0x4b, 0x2c, 0xc8, 0xd4, 0x03, 0xab, 0x09, 0x82, 0xca, 0x18, 0x59, 0x71, 0xb1, 0x41, 0x34,
	0x09, 0x19, 0x70, 0xb1, 0x80, 0x34, 0x0a, 0x09, 0x21, 0x54, 0xc1, 0x8c, 0x96, 0xc2, 0x14, 0x2b,
	0x56, 0x62, 0x48, 0x62, 0x03, 0x3b, 0xc4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x4d, 0xbf,
	0x6c, 0xb2, 0x00, 0x00, 0x00,
}
