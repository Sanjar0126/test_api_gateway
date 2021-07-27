// Code generated by protoc-gen-go. DO NOT EDIT.
// source: iiko_service.proto

package user_service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type IikoCredentialsId struct {
	ShipperId            string   `protobuf:"bytes,1,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IikoCredentialsId) Reset()         { *m = IikoCredentialsId{} }
func (m *IikoCredentialsId) String() string { return proto.CompactTextString(m) }
func (*IikoCredentialsId) ProtoMessage()    {}
func (*IikoCredentialsId) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c72f4b0dccd93f3, []int{0}
}

func (m *IikoCredentialsId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IikoCredentialsId.Unmarshal(m, b)
}
func (m *IikoCredentialsId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IikoCredentialsId.Marshal(b, m, deterministic)
}
func (m *IikoCredentialsId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IikoCredentialsId.Merge(m, src)
}
func (m *IikoCredentialsId) XXX_Size() int {
	return xxx_messageInfo_IikoCredentialsId.Size(m)
}
func (m *IikoCredentialsId) XXX_DiscardUnknown() {
	xxx_messageInfo_IikoCredentialsId.DiscardUnknown(m)
}

var xxx_messageInfo_IikoCredentialsId proto.InternalMessageInfo

func (m *IikoCredentialsId) GetShipperId() string {
	if m != nil {
		return m.ShipperId
	}
	return ""
}

func init() {
	proto.RegisterType((*IikoCredentialsId)(nil), "genproto.IikoCredentialsId")
}

func init() { proto.RegisterFile("iiko_service.proto", fileDescriptor_9c72f4b0dccd93f3) }

var fileDescriptor_9c72f4b0dccd93f3 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0xcc, 0xcc, 0xce,
	0x8f, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x48, 0x4f, 0xcd, 0x03, 0xb3, 0xa4, 0xb8, 0x40, 0xb2, 0x10, 0x51, 0x29, 0xe9, 0xf4, 0xfc, 0xfc,
	0xf4, 0x9c, 0x54, 0x7d, 0x30, 0x2f, 0xa9, 0x34, 0x4d, 0x3f, 0x35, 0xb7, 0xa0, 0xa4, 0x12, 0x22,
	0xa9, 0x64, 0xc4, 0x25, 0xe8, 0x99, 0x99, 0x9d, 0xef, 0x5c, 0x94, 0x9a, 0x92, 0x9a, 0x57, 0x92,
	0x99, 0x98, 0x53, 0xec, 0x99, 0x22, 0x24, 0xcb, 0xc5, 0x55, 0x9c, 0x91, 0x59, 0x50, 0x90, 0x5a,
	0x14, 0x9f, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0xc4, 0x09, 0x15, 0xf1, 0x4c, 0x31,
	0x9a, 0xcb, 0xc4, 0x25, 0x86, 0xa6, 0x29, 0x18, 0xe2, 0x0e, 0x21, 0x27, 0x2e, 0x36, 0xe7, 0xa2,
	0xd4, 0xc4, 0x92, 0x54, 0x21, 0x49, 0x3d, 0x98, 0x63, 0xf4, 0xd0, 0xd4, 0x4a, 0x49, 0xe3, 0x94,
	0xf2, 0x4c, 0x51, 0x62, 0x10, 0xb2, 0xe5, 0x62, 0x0b, 0x2d, 0x48, 0x21, 0x60, 0x86, 0x98, 0x1e,
	0xc4, 0x57, 0x7a, 0x30, 0x5f, 0xe9, 0xb9, 0x82, 0x7c, 0xa5, 0xc4, 0x20, 0x64, 0xcf, 0xc5, 0xec,
	0x9e, 0x5a, 0x22, 0x84, 0xcf, 0x12, 0x29, 0xdc, 0x06, 0x83, 0x0d, 0x60, 0x73, 0x49, 0xcd, 0x49,
	0x2d, 0x49, 0xc5, 0x6f, 0x06, 0x4e, 0x17, 0x38, 0x89, 0x47, 0x89, 0xc2, 0xf4, 0xe9, 0x97, 0x16,
	0xa7, 0x16, 0xc1, 0x62, 0x29, 0x89, 0x0d, 0x2c, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb1,
	0x3c, 0x17, 0x2a, 0xbc, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IikoCredentialsServiceClient is the client API for IikoCredentialsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IikoCredentialsServiceClient interface {
	Create(ctx context.Context, in *IikoCredentials, opts ...grpc.CallOption) (*IikoCredentialsId, error)
	Update(ctx context.Context, in *IikoCredentials, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Get(ctx context.Context, in *IikoCredentialsId, opts ...grpc.CallOption) (*IikoCredentials, error)
	Delete(ctx context.Context, in *IikoCredentialsId, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type iikoCredentialsServiceClient struct {
	cc *grpc.ClientConn
}

func NewIikoCredentialsServiceClient(cc *grpc.ClientConn) IikoCredentialsServiceClient {
	return &iikoCredentialsServiceClient{cc}
}

func (c *iikoCredentialsServiceClient) Create(ctx context.Context, in *IikoCredentials, opts ...grpc.CallOption) (*IikoCredentialsId, error) {
	out := new(IikoCredentialsId)
	err := c.cc.Invoke(ctx, "/genproto.IikoCredentialsService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iikoCredentialsServiceClient) Update(ctx context.Context, in *IikoCredentials, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/genproto.IikoCredentialsService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iikoCredentialsServiceClient) Get(ctx context.Context, in *IikoCredentialsId, opts ...grpc.CallOption) (*IikoCredentials, error) {
	out := new(IikoCredentials)
	err := c.cc.Invoke(ctx, "/genproto.IikoCredentialsService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iikoCredentialsServiceClient) Delete(ctx context.Context, in *IikoCredentialsId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/genproto.IikoCredentialsService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IikoCredentialsServiceServer is the server API for IikoCredentialsService service.
type IikoCredentialsServiceServer interface {
	Create(context.Context, *IikoCredentials) (*IikoCredentialsId, error)
	Update(context.Context, *IikoCredentials) (*emptypb.Empty, error)
	Get(context.Context, *IikoCredentialsId) (*IikoCredentials, error)
	Delete(context.Context, *IikoCredentialsId) (*emptypb.Empty, error)
}

// UnimplementedIikoCredentialsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedIikoCredentialsServiceServer struct {
}

func (*UnimplementedIikoCredentialsServiceServer) Create(ctx context.Context, req *IikoCredentials) (*IikoCredentialsId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedIikoCredentialsServiceServer) Update(ctx context.Context, req *IikoCredentials) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedIikoCredentialsServiceServer) Get(ctx context.Context, req *IikoCredentialsId) (*IikoCredentials, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedIikoCredentialsServiceServer) Delete(ctx context.Context, req *IikoCredentialsId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterIikoCredentialsServiceServer(s *grpc.Server, srv IikoCredentialsServiceServer) {
	s.RegisterService(&_IikoCredentialsService_serviceDesc, srv)
}

func _IikoCredentialsService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IikoCredentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IikoCredentialsServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.IikoCredentialsService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IikoCredentialsServiceServer).Create(ctx, req.(*IikoCredentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _IikoCredentialsService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IikoCredentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IikoCredentialsServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.IikoCredentialsService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IikoCredentialsServiceServer).Update(ctx, req.(*IikoCredentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _IikoCredentialsService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IikoCredentialsId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IikoCredentialsServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.IikoCredentialsService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IikoCredentialsServiceServer).Get(ctx, req.(*IikoCredentialsId))
	}
	return interceptor(ctx, in, info, handler)
}

func _IikoCredentialsService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IikoCredentialsId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IikoCredentialsServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.IikoCredentialsService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IikoCredentialsServiceServer).Delete(ctx, req.(*IikoCredentialsId))
	}
	return interceptor(ctx, in, info, handler)
}

var _IikoCredentialsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.IikoCredentialsService",
	HandlerType: (*IikoCredentialsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _IikoCredentialsService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _IikoCredentialsService_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _IikoCredentialsService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _IikoCredentialsService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iiko_service.proto",
}