// Code generated by protoc-gen-go. DO NOT EDIT.
// source: add.proto

package proto

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

type AddRequest struct {
	A                    uint64   `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    uint64   `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_077cd88a1973142f, []int{0}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetA() uint64 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *AddRequest) GetB() uint64 {
	if m != nil {
		return m.B
	}
	return 0
}

type AddResponse struct {
	Result               uint64   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()    {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_077cd88a1973142f, []int{1}
}

func (m *AddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddResponse.Unmarshal(m, b)
}
func (m *AddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddResponse.Marshal(b, m, deterministic)
}
func (m *AddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResponse.Merge(m, src)
}
func (m *AddResponse) XXX_Size() int {
	return xxx_messageInfo_AddResponse.Size(m)
}
func (m *AddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddResponse proto.InternalMessageInfo

func (m *AddResponse) GetResult() uint64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*AddRequest)(nil), "proto.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "proto.AddResponse")
}

func init() { proto.RegisterFile("add.proto", fileDescriptor_077cd88a1973142f) }

var fileDescriptor_077cd88a1973142f = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x4c, 0x49, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x1a, 0x5c, 0x5c, 0x8e, 0x29, 0x29,
	0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x3c, 0x5c, 0x8c, 0x89, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x2c, 0x41, 0x8c, 0x89, 0x20, 0x5e, 0x92, 0x04, 0x13, 0x84, 0x97, 0xa4, 0xa4, 0xca, 0xc5,
	0x0d, 0x56, 0x59, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc6, 0xc5, 0x56, 0x94, 0x5a, 0x5c,
	0x9a, 0x53, 0x02, 0x55, 0x0f, 0xe5, 0x19, 0x39, 0x80, 0x0d, 0x0c, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c,
	0x4e, 0x15, 0x32, 0xe2, 0x62, 0x77, 0xce, 0xcf, 0x2d, 0x28, 0x2d, 0x49, 0x15, 0x12, 0x84, 0x58,
	0xac, 0x87, 0xb0, 0x4e, 0x4a, 0x08, 0x59, 0x08, 0x62, 0xae, 0x12, 0x43, 0x12, 0x1b, 0x58, 0xd0,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x60, 0x35, 0x3a, 0xad, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AddServiceClient is the client API for AddService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddServiceClient interface {
	Compute(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
}

type addServiceClient struct {
	cc *grpc.ClientConn
}

func NewAddServiceClient(cc *grpc.ClientConn) AddServiceClient {
	return &addServiceClient{cc}
}

func (c *addServiceClient) Compute(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/proto.AddService/Compute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddServiceServer is the server API for AddService service.
type AddServiceServer interface {
	Compute(context.Context, *AddRequest) (*AddResponse, error)
}

// UnimplementedAddServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAddServiceServer struct {
}

func (*UnimplementedAddServiceServer) Compute(ctx context.Context, req *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Compute not implemented")
}

func RegisterAddServiceServer(s *grpc.Server, srv AddServiceServer) {
	s.RegisterService(&_AddService_serviceDesc, srv)
}

func _AddService_Compute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).Compute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddService/Compute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).Compute(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AddService",
	HandlerType: (*AddServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Compute",
			Handler:    _AddService_Compute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "add.proto",
}
