// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/api.proto

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

type Test_State int32

const (
	Test_UNKNOWN  Test_State = 0
	Test_PASSED   Test_State = 1
	Test_FAILED   Test_State = 2
	Test_SKIPPED  Test_State = 3
	Test_PENDING  Test_State = 4
	Test_PANICKED Test_State = 5
	Test_TIMEDOUT Test_State = 6
	Test_INVALID  Test_State = 7
)

var Test_State_name = map[int32]string{
	0: "UNKNOWN",
	1: "PASSED",
	2: "FAILED",
	3: "SKIPPED",
	4: "PENDING",
	5: "PANICKED",
	6: "TIMEDOUT",
	7: "INVALID",
}
var Test_State_value = map[string]int32{
	"UNKNOWN":  0,
	"PASSED":   1,
	"FAILED":   2,
	"SKIPPED":  3,
	"PENDING":  4,
	"PANICKED": 5,
	"TIMEDOUT": 6,
	"INVALID":  7,
}

func (x Test_State) String() string {
	return proto.EnumName(Test_State_name, int32(x))
}
func (Test_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_api_c11a3ff407495f5c, []int{1, 0}
}

type SuiteSummary struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Codebase             string   `protobuf:"bytes,2,opt,name=codebase" json:"codebase,omitempty"`
	Commit               string   `protobuf:"bytes,3,opt,name=commit" json:"commit,omitempty"`
	Tests                []*Test  `protobuf:"bytes,4,rep,name=tests" json:"tests,omitempty"`
	StartedAt            int64    `protobuf:"varint,5,opt,name=started_at,json=startedAt" json:"started_at,omitempty"`
	FinishedAt           int64    `protobuf:"varint,6,opt,name=finished_at,json=finishedAt" json:"finished_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SuiteSummary) Reset()         { *m = SuiteSummary{} }
func (m *SuiteSummary) String() string { return proto.CompactTextString(m) }
func (*SuiteSummary) ProtoMessage()    {}
func (*SuiteSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_c11a3ff407495f5c, []int{0}
}
func (m *SuiteSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SuiteSummary.Unmarshal(m, b)
}
func (m *SuiteSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SuiteSummary.Marshal(b, m, deterministic)
}
func (dst *SuiteSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuiteSummary.Merge(dst, src)
}
func (m *SuiteSummary) XXX_Size() int {
	return xxx_messageInfo_SuiteSummary.Size(m)
}
func (m *SuiteSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_SuiteSummary.DiscardUnknown(m)
}

var xxx_messageInfo_SuiteSummary proto.InternalMessageInfo

func (m *SuiteSummary) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SuiteSummary) GetCodebase() string {
	if m != nil {
		return m.Codebase
	}
	return ""
}

func (m *SuiteSummary) GetCommit() string {
	if m != nil {
		return m.Commit
	}
	return ""
}

func (m *SuiteSummary) GetTests() []*Test {
	if m != nil {
		return m.Tests
	}
	return nil
}

func (m *SuiteSummary) GetStartedAt() int64 {
	if m != nil {
		return m.StartedAt
	}
	return 0
}

func (m *SuiteSummary) GetFinishedAt() int64 {
	if m != nil {
		return m.FinishedAt
	}
	return 0
}

type Test struct {
	Name                 string     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	State                Test_State `protobuf:"varint,2,opt,name=state,enum=api.Test_State" json:"state,omitempty"`
	Failure              *Failure   `protobuf:"bytes,3,opt,name=failure" json:"failure,omitempty"`
	Location             string     `protobuf:"bytes,4,opt,name=location" json:"location,omitempty"`
	StartedAt            int64      `protobuf:"varint,5,opt,name=started_at,json=startedAt" json:"started_at,omitempty"`
	FinishedAt           int64      `protobuf:"varint,6,opt,name=finished_at,json=finishedAt" json:"finished_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Test) Reset()         { *m = Test{} }
func (m *Test) String() string { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()    {}
func (*Test) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_c11a3ff407495f5c, []int{1}
}
func (m *Test) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test.Unmarshal(m, b)
}
func (m *Test) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test.Marshal(b, m, deterministic)
}
func (dst *Test) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test.Merge(dst, src)
}
func (m *Test) XXX_Size() int {
	return xxx_messageInfo_Test.Size(m)
}
func (m *Test) XXX_DiscardUnknown() {
	xxx_messageInfo_Test.DiscardUnknown(m)
}

var xxx_messageInfo_Test proto.InternalMessageInfo

func (m *Test) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Test) GetState() Test_State {
	if m != nil {
		return m.State
	}
	return Test_UNKNOWN
}

func (m *Test) GetFailure() *Failure {
	if m != nil {
		return m.Failure
	}
	return nil
}

func (m *Test) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Test) GetStartedAt() int64 {
	if m != nil {
		return m.StartedAt
	}
	return 0
}

func (m *Test) GetFinishedAt() int64 {
	if m != nil {
		return m.FinishedAt
	}
	return 0
}

type Failure struct {
	Message              string   `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Failure) Reset()         { *m = Failure{} }
func (m *Failure) String() string { return proto.CompactTextString(m) }
func (*Failure) ProtoMessage()    {}
func (*Failure) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_c11a3ff407495f5c, []int{2}
}
func (m *Failure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Failure.Unmarshal(m, b)
}
func (m *Failure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Failure.Marshal(b, m, deterministic)
}
func (dst *Failure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Failure.Merge(dst, src)
}
func (m *Failure) XXX_Size() int {
	return xxx_messageInfo_Failure.Size(m)
}
func (m *Failure) XXX_DiscardUnknown() {
	xxx_messageInfo_Failure.DiscardUnknown(m)
}

var xxx_messageInfo_Failure proto.InternalMessageInfo

func (m *Failure) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CreateRequest struct {
	Summary              *SuiteSummary `protobuf:"bytes,1,opt,name=summary" json:"summary,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_c11a3ff407495f5c, []int{3}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(dst, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetSummary() *SuiteSummary {
	if m != nil {
		return m.Summary
	}
	return nil
}

type CreateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_c11a3ff407495f5c, []int{4}
}
func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (dst *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(dst, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_c11a3ff407495f5c, []int{5}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (dst *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(dst, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

type ListResponse struct {
	SuiteSummaries       []*SuiteSummary `protobuf:"bytes,1,rep,name=SuiteSummaries" json:"SuiteSummaries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_c11a3ff407495f5c, []int{6}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (dst *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(dst, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetSuiteSummaries() []*SuiteSummary {
	if m != nil {
		return m.SuiteSummaries
	}
	return nil
}

type GetRequest struct {
	Codebase             string   `protobuf:"bytes,1,opt,name=codebase" json:"codebase,omitempty"`
	Commit               string   `protobuf:"bytes,2,opt,name=commit" json:"commit,omitempty"`
	Location             string   `protobuf:"bytes,3,opt,name=location" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_c11a3ff407495f5c, []int{7}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetCodebase() string {
	if m != nil {
		return m.Codebase
	}
	return ""
}

func (m *GetRequest) GetCommit() string {
	if m != nil {
		return m.Commit
	}
	return ""
}

func (m *GetRequest) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

type GetResponse struct {
	Test                 *Test    `protobuf:"bytes,1,opt,name=test" json:"test,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_c11a3ff407495f5c, []int{8}
}
func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (dst *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(dst, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetTest() *Test {
	if m != nil {
		return m.Test
	}
	return nil
}

func init() {
	proto.RegisterType((*SuiteSummary)(nil), "api.SuiteSummary")
	proto.RegisterType((*Test)(nil), "api.Test")
	proto.RegisterType((*Failure)(nil), "api.Failure")
	proto.RegisterType((*CreateRequest)(nil), "api.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "api.CreateResponse")
	proto.RegisterType((*ListRequest)(nil), "api.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "api.ListResponse")
	proto.RegisterType((*GetRequest)(nil), "api.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "api.GetResponse")
	proto.RegisterEnum("api.Test_State", Test_State_name, Test_State_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Suite service

type SuiteClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type suiteClient struct {
	cc *grpc.ClientConn
}

func NewSuiteClient(cc *grpc.ClientConn) SuiteClient {
	return &suiteClient{cc}
}

func (c *suiteClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc.Invoke(ctx, "/api.Suite/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *suiteClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/api.Suite/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *suiteClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/api.Suite/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Suite service

type SuiteServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
}

func RegisterSuiteServer(s *grpc.Server, srv SuiteServer) {
	s.RegisterService(&_Suite_serviceDesc, srv)
}

func _Suite_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuiteServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Suite/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuiteServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Suite_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuiteServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Suite/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuiteServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Suite_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuiteServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Suite/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuiteServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Suite_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Suite",
	HandlerType: (*SuiteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Suite_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Suite_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Suite_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}

func init() { proto.RegisterFile("api/api.proto", fileDescriptor_api_c11a3ff407495f5c) }

var fileDescriptor_api_c11a3ff407495f5c = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0x5e, 0x9a, 0xaf, 0xf5, 0x4d, 0x5b, 0x3c, 0x23, 0xa1, 0xa8, 0xd2, 0x44, 0x15, 0x04, 0xaa,
	0xf8, 0x28, 0x52, 0x77, 0x42, 0xe2, 0x12, 0x2d, 0x59, 0x15, 0xb5, 0x64, 0x55, 0xd2, 0xc1, 0x05,
	0x09, 0x79, 0x9d, 0x07, 0x96, 0x96, 0xa6, 0xd4, 0xee, 0x81, 0x5f, 0xc2, 0x3f, 0x81, 0xbf, 0x87,
	0x6c, 0x27, 0x5d, 0x82, 0xc6, 0x89, 0x5b, 0x9e, 0x8f, 0xbc, 0x7e, 0xf3, 0xf8, 0x09, 0xf4, 0xc9,
	0x96, 0xbd, 0x25, 0x5b, 0x36, 0xd9, 0xee, 0x4a, 0x51, 0x62, 0x93, 0x6c, 0x59, 0xf0, 0xdb, 0x80,
	0x5e, 0xbe, 0x67, 0x82, 0xe6, 0xfb, 0xa2, 0x20, 0xbb, 0x1f, 0x18, 0x83, 0xb5, 0x21, 0x05, 0xf5,
	0x8d, 0x91, 0x31, 0xee, 0x66, 0xea, 0x19, 0x0f, 0xe1, 0x78, 0x5d, 0xde, 0xd0, 0x6b, 0xc2, 0xa9,
	0xdf, 0x51, 0xfc, 0x01, 0xe3, 0x27, 0xe0, 0xac, 0xcb, 0xa2, 0x60, 0xc2, 0x37, 0x95, 0x52, 0x21,
	0xfc, 0x14, 0x6c, 0x41, 0xb9, 0xe0, 0xbe, 0x35, 0x32, 0xc7, 0xde, 0xb4, 0x3b, 0x91, 0x07, 0xaf,
	0x28, 0x17, 0x99, 0xe6, 0xf1, 0x29, 0x00, 0x17, 0x64, 0x27, 0xe8, 0xcd, 0x17, 0x22, 0x7c, 0x7b,
	0x64, 0x8c, 0xcd, 0xac, 0x5b, 0x31, 0xa1, 0x7c, 0xdf, 0xbb, 0x65, 0x1b, 0xc6, 0xbf, 0x69, 0xdd,
	0x51, 0x3a, 0xd4, 0x54, 0x28, 0x82, 0x5f, 0x1d, 0xb0, 0xe4, 0xbc, 0x07, 0x37, 0x7e, 0x0e, 0x36,
	0x17, 0x44, 0xe8, 0x75, 0x07, 0xd3, 0x47, 0x87, 0xd3, 0x27, 0xb9, 0xa4, 0x33, 0xad, 0xe2, 0x17,
	0xe0, 0xde, 0x12, 0x76, 0xb7, 0xdf, 0x51, 0xb5, 0xbd, 0x37, 0xed, 0x29, 0xe3, 0x85, 0xe6, 0xb2,
	0x5a, 0x94, 0x01, 0xdc, 0x95, 0x6b, 0x22, 0x58, 0xb9, 0xf1, 0x2d, 0x1d, 0x40, 0x8d, 0xff, 0xfb,
	0x3b, 0x4a, 0xb0, 0xd5, 0x4e, 0xd8, 0x03, 0xf7, 0x2a, 0x9d, 0xa7, 0x97, 0x9f, 0x52, 0x74, 0x84,
	0x01, 0x9c, 0x65, 0x98, 0xe7, 0x71, 0x84, 0x0c, 0xf9, 0x7c, 0x11, 0x26, 0x8b, 0x38, 0x42, 0x1d,
	0x69, 0xca, 0xe7, 0xc9, 0x72, 0x19, 0x47, 0xc8, 0x94, 0x60, 0x19, 0xa7, 0x51, 0x92, 0xce, 0x90,
	0x85, 0x7b, 0x70, 0xbc, 0x0c, 0xd3, 0xe4, 0x7c, 0x1e, 0x47, 0xc8, 0x96, 0x68, 0x95, 0x7c, 0x88,
	0xa3, 0xcb, 0xab, 0x15, 0x72, 0xa4, 0x31, 0x49, 0x3f, 0x86, 0x8b, 0x24, 0x42, 0x6e, 0xf0, 0x0c,
	0xdc, 0xea, 0x03, 0xb1, 0x0f, 0x6e, 0x41, 0x39, 0x27, 0x5f, 0xeb, 0xf4, 0x6a, 0x18, 0xbc, 0x87,
	0xfe, 0xf9, 0x8e, 0xca, 0xa8, 0xe8, 0xf7, 0xbd, 0x4c, 0xf9, 0x15, 0xb8, 0x5c, 0x57, 0x44, 0x59,
	0xbd, 0xe9, 0x89, 0x8a, 0xaa, 0xd9, 0x9d, 0xac, 0x76, 0x04, 0x08, 0x06, 0xf5, 0xdb, 0x7c, 0x5b,
	0x6e, 0x38, 0x0d, 0xfa, 0xe0, 0x2d, 0x18, 0x17, 0xd5, 0xb4, 0x20, 0x81, 0x9e, 0x86, 0x5a, 0xc6,
	0xef, 0x60, 0xd0, 0x98, 0xc4, 0x28, 0xf7, 0x0d, 0x55, 0x9b, 0x07, 0x0e, 0xf9, 0xcb, 0x18, 0x7c,
	0x06, 0x98, 0xd1, 0x7a, 0x70, 0xab, 0xaa, 0xc6, 0x3f, 0xab, 0xda, 0x69, 0x55, 0xb5, 0x79, 0xbb,
	0x66, 0xfb, 0x76, 0x83, 0xd7, 0xe0, 0xa9, 0xe9, 0xd5, 0x9e, 0xa7, 0x60, 0xc9, 0xf6, 0x56, 0x11,
	0x34, 0x4a, 0xad, 0xe8, 0xe9, 0x4f, 0x03, 0x6c, 0xb5, 0x1e, 0x3e, 0x03, 0x47, 0x27, 0x80, 0xb1,
	0x32, 0xb5, 0xc2, 0x1c, 0x3e, 0x6e, 0x71, 0x55, 0x44, 0x47, 0xf8, 0x0d, 0x58, 0x32, 0x15, 0x8c,
	0x94, 0xdc, 0xc8, 0x6b, 0x78, 0xd2, 0x60, 0x0e, 0xf6, 0x97, 0x60, 0xce, 0xa8, 0xc0, 0xba, 0xdc,
	0xf7, 0x19, 0x0c, 0xd1, 0x3d, 0x51, 0x7b, 0xaf, 0x1d, 0xf5, 0xcf, 0x9f, 0xfd, 0x09, 0x00, 0x00,
	0xff, 0xff, 0x85, 0x58, 0x5f, 0xf9, 0x04, 0x04, 0x00, 0x00,
}
