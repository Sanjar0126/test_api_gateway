// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shipper_service.proto

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

type GetShipperByNameRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShipperByNameRequest) Reset()         { *m = GetShipperByNameRequest{} }
func (m *GetShipperByNameRequest) String() string { return proto.CompactTextString(m) }
func (*GetShipperByNameRequest) ProtoMessage()    {}
func (*GetShipperByNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e270737089a546a, []int{0}
}

func (m *GetShipperByNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShipperByNameRequest.Unmarshal(m, b)
}
func (m *GetShipperByNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShipperByNameRequest.Marshal(b, m, deterministic)
}
func (m *GetShipperByNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShipperByNameRequest.Merge(m, src)
}
func (m *GetShipperByNameRequest) XXX_Size() int {
	return xxx_messageInfo_GetShipperByNameRequest.Size(m)
}
func (m *GetShipperByNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShipperByNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetShipperByNameRequest proto.InternalMessageInfo

func (m *GetShipperByNameRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ShipperId struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShipperId) Reset()         { *m = ShipperId{} }
func (m *ShipperId) String() string { return proto.CompactTextString(m) }
func (*ShipperId) ProtoMessage()    {}
func (*ShipperId) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e270737089a546a, []int{1}
}

func (m *ShipperId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShipperId.Unmarshal(m, b)
}
func (m *ShipperId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShipperId.Marshal(b, m, deterministic)
}
func (m *ShipperId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShipperId.Merge(m, src)
}
func (m *ShipperId) XXX_Size() int {
	return xxx_messageInfo_ShipperId.Size(m)
}
func (m *ShipperId) XXX_DiscardUnknown() {
	xxx_messageInfo_ShipperId.DiscardUnknown(m)
}

var xxx_messageInfo_ShipperId proto.InternalMessageInfo

func (m *ShipperId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetAllShippersRequest struct {
	Page                 uint64   `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit                uint64   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	HasIiko              bool     `protobuf:"varint,3,opt,name=has_iiko,json=hasIiko,proto3" json:"has_iiko,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllShippersRequest) Reset()         { *m = GetAllShippersRequest{} }
func (m *GetAllShippersRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllShippersRequest) ProtoMessage()    {}
func (*GetAllShippersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e270737089a546a, []int{2}
}

func (m *GetAllShippersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllShippersRequest.Unmarshal(m, b)
}
func (m *GetAllShippersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllShippersRequest.Marshal(b, m, deterministic)
}
func (m *GetAllShippersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllShippersRequest.Merge(m, src)
}
func (m *GetAllShippersRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllShippersRequest.Size(m)
}
func (m *GetAllShippersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllShippersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllShippersRequest proto.InternalMessageInfo

func (m *GetAllShippersRequest) GetPage() uint64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetAllShippersRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetAllShippersRequest) GetHasIiko() bool {
	if m != nil {
		return m.HasIiko
	}
	return false
}

type GetAllShippersResponse struct {
	Shippers             []*Shipper `protobuf:"bytes,1,rep,name=shippers,proto3" json:"shippers,omitempty"`
	Count                uint64     `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetAllShippersResponse) Reset()         { *m = GetAllShippersResponse{} }
func (m *GetAllShippersResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllShippersResponse) ProtoMessage()    {}
func (*GetAllShippersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e270737089a546a, []int{3}
}

func (m *GetAllShippersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllShippersResponse.Unmarshal(m, b)
}
func (m *GetAllShippersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllShippersResponse.Marshal(b, m, deterministic)
}
func (m *GetAllShippersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllShippersResponse.Merge(m, src)
}
func (m *GetAllShippersResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllShippersResponse.Size(m)
}
func (m *GetAllShippersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllShippersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllShippersResponse proto.InternalMessageInfo

func (m *GetAllShippersResponse) GetShippers() []*Shipper {
	if m != nil {
		return m.Shippers
	}
	return nil
}

func (m *GetAllShippersResponse) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*GetShipperByNameRequest)(nil), "genproto.GetShipperByNameRequest")
	proto.RegisterType((*ShipperId)(nil), "genproto.ShipperId")
	proto.RegisterType((*GetAllShippersRequest)(nil), "genproto.GetAllShippersRequest")
	proto.RegisterType((*GetAllShippersResponse)(nil), "genproto.GetAllShippersResponse")
}

func init() { proto.RegisterFile("shipper_service.proto", fileDescriptor_5e270737089a546a) }

var fileDescriptor_5e270737089a546a = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x41, 0xaf, 0x93, 0x40,
	0x10, 0x06, 0xfa, 0x44, 0x3a, 0xc6, 0x97, 0x38, 0xda, 0xf7, 0x90, 0x77, 0x10, 0xf7, 0xd4, 0x4b,
	0x21, 0xa9, 0x31, 0x9e, 0x6d, 0x35, 0xa4, 0x07, 0x3d, 0xd0, 0x78, 0x31, 0x9a, 0x86, 0x96, 0x91,
	0x6e, 0x0a, 0x2c, 0xb2, 0x8b, 0x49, 0x7f, 0xba, 0x37, 0x53, 0x96, 0xb6, 0x6a, 0xe9, 0x6d, 0x66,
	0xbe, 0x6f, 0xbf, 0x9d, 0xf9, 0xf2, 0xc1, 0x48, 0x6e, 0x79, 0x55, 0x51, 0xbd, 0x92, 0x54, 0xff,
	0xe2, 0x1b, 0x0a, 0xaa, 0x5a, 0x28, 0x81, 0x4e, 0x46, 0x65, 0x5b, 0x79, 0x4f, 0x3b, 0x82, 0x06,
	0xbc, 0x87, 0x4c, 0x88, 0x2c, 0xa7, 0xb0, 0xed, 0xd6, 0xcd, 0x8f, 0x90, 0x8a, 0x4a, 0xed, 0x35,
	0xc8, 0x26, 0x70, 0x1f, 0x91, 0x5a, 0xea, 0x07, 0xb3, 0xfd, 0xe7, 0xa4, 0xa0, 0x98, 0x7e, 0x36,
	0x24, 0x15, 0x22, 0xdc, 0x94, 0x49, 0x41, 0xae, 0xe9, 0x9b, 0xe3, 0x61, 0xdc, 0xd6, 0xec, 0x01,
	0x86, 0x1d, 0x77, 0x91, 0xe2, 0x2d, 0x58, 0x3c, 0xed, 0x60, 0x8b, 0xa7, 0xec, 0x1b, 0x8c, 0x22,
	0x52, 0xef, 0xf3, 0xbc, 0xa3, 0xc8, 0xbf, 0x94, 0xaa, 0x24, 0xd3, 0x4a, 0x37, 0x71, 0x5b, 0xe3,
	0x0b, 0x78, 0x94, 0xf3, 0x82, 0x2b, 0xd7, 0x6a, 0x87, 0xba, 0xc1, 0x97, 0xe0, 0x6c, 0x13, 0xb9,
	0xe2, 0x7c, 0x27, 0xdc, 0x81, 0x6f, 0x8e, 0x9d, 0xf8, 0xf1, 0x36, 0x91, 0x0b, 0xbe, 0x13, 0xec,
	0x3b, 0xdc, 0xfd, 0xaf, 0x2e, 0x2b, 0x51, 0x4a, 0xc2, 0x09, 0x38, 0xdd, 0xc5, 0xd2, 0x35, 0xfd,
	0xc1, 0xf8, 0xc9, 0xf4, 0x59, 0x70, 0x34, 0x23, 0xe8, 0xd8, 0xf1, 0x89, 0x72, 0xf8, 0x79, 0x23,
	0x9a, 0xf2, 0xf4, 0x73, 0xdb, 0x4c, 0x7f, 0x5b, 0x70, 0xdb, 0x71, 0x97, 0xda, 0x57, 0x9c, 0x82,
	0x3d, 0xaf, 0x29, 0x51, 0x84, 0x97, 0x7a, 0xde, 0xf3, 0x8b, 0xd1, 0x22, 0x65, 0x06, 0x86, 0x30,
	0x88, 0x48, 0x61, 0x1f, 0xea, 0x5d, 0xaa, 0x30, 0x03, 0x3f, 0x81, 0xad, 0xcf, 0xc2, 0x57, 0x67,
	0xb8, 0xd7, 0x46, 0xcf, 0xbf, 0x4e, 0xd0, 0x4e, 0x30, 0x03, 0xdf, 0x82, 0xfd, 0xa5, 0x4a, 0xaf,
	0xec, 0x7c, 0x17, 0xe8, 0x28, 0x04, 0xc7, 0x28, 0x04, 0x1f, 0x0f, 0x51, 0x60, 0x06, 0xbe, 0x03,
	0xfb, 0x03, 0xe5, 0xa4, 0xa8, 0x7f, 0xf3, 0xeb, 0x0f, 0xe7, 0x30, 0x8c, 0x48, 0xe9, 0xe0, 0xe0,
	0xeb, 0x7f, 0x16, 0xec, 0x0b, 0x55, 0xaf, 0x07, 0xb3, 0xfb, 0xaf, 0xa3, 0xe3, 0x34, 0x6c, 0xe4,
	0x39, 0xd9, 0x6b, 0xbb, 0x9d, 0xbd, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x23, 0x16, 0x4b,
	0xf3, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShipperServiceClient is the client API for ShipperService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShipperServiceClient interface {
	Create(ctx context.Context, in *Shipper, opts ...grpc.CallOption) (*ShipperId, error)
	Get(ctx context.Context, in *ShipperId, opts ...grpc.CallOption) (*Shipper, error)
	GetAll(ctx context.Context, in *GetAllShippersRequest, opts ...grpc.CallOption) (*GetAllShippersResponse, error)
	Update(ctx context.Context, in *Shipper, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *ShipperId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetByName(ctx context.Context, in *GetShipperByNameRequest, opts ...grpc.CallOption) (*Shipper, error)
}

type shipperServiceClient struct {
	cc *grpc.ClientConn
}

func NewShipperServiceClient(cc *grpc.ClientConn) ShipperServiceClient {
	return &shipperServiceClient{cc}
}

func (c *shipperServiceClient) Create(ctx context.Context, in *Shipper, opts ...grpc.CallOption) (*ShipperId, error) {
	out := new(ShipperId)
	err := c.cc.Invoke(ctx, "/genproto.ShipperService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipperServiceClient) Get(ctx context.Context, in *ShipperId, opts ...grpc.CallOption) (*Shipper, error) {
	out := new(Shipper)
	err := c.cc.Invoke(ctx, "/genproto.ShipperService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipperServiceClient) GetAll(ctx context.Context, in *GetAllShippersRequest, opts ...grpc.CallOption) (*GetAllShippersResponse, error) {
	out := new(GetAllShippersResponse)
	err := c.cc.Invoke(ctx, "/genproto.ShipperService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipperServiceClient) Update(ctx context.Context, in *Shipper, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/genproto.ShipperService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipperServiceClient) Delete(ctx context.Context, in *ShipperId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/genproto.ShipperService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipperServiceClient) GetByName(ctx context.Context, in *GetShipperByNameRequest, opts ...grpc.CallOption) (*Shipper, error) {
	out := new(Shipper)
	err := c.cc.Invoke(ctx, "/genproto.ShipperService/GetByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShipperServiceServer is the server API for ShipperService service.
type ShipperServiceServer interface {
	Create(context.Context, *Shipper) (*ShipperId, error)
	Get(context.Context, *ShipperId) (*Shipper, error)
	GetAll(context.Context, *GetAllShippersRequest) (*GetAllShippersResponse, error)
	Update(context.Context, *Shipper) (*emptypb.Empty, error)
	Delete(context.Context, *ShipperId) (*emptypb.Empty, error)
	GetByName(context.Context, *GetShipperByNameRequest) (*Shipper, error)
}

// UnimplementedShipperServiceServer can be embedded to have forward compatible implementations.
type UnimplementedShipperServiceServer struct {
}

func (*UnimplementedShipperServiceServer) Create(ctx context.Context, req *Shipper) (*ShipperId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedShipperServiceServer) Get(ctx context.Context, req *ShipperId) (*Shipper, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedShipperServiceServer) GetAll(ctx context.Context, req *GetAllShippersRequest) (*GetAllShippersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedShipperServiceServer) Update(ctx context.Context, req *Shipper) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedShipperServiceServer) Delete(ctx context.Context, req *ShipperId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedShipperServiceServer) GetByName(ctx context.Context, req *GetShipperByNameRequest) (*Shipper, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByName not implemented")
}

func RegisterShipperServiceServer(s *grpc.Server, srv ShipperServiceServer) {
	s.RegisterService(&_ShipperService_serviceDesc, srv)
}

func _ShipperService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Shipper)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipperServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ShipperService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipperServiceServer).Create(ctx, req.(*Shipper))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipperService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShipperId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipperServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ShipperService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipperServiceServer).Get(ctx, req.(*ShipperId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipperService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllShippersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipperServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ShipperService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipperServiceServer).GetAll(ctx, req.(*GetAllShippersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipperService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Shipper)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipperServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ShipperService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipperServiceServer).Update(ctx, req.(*Shipper))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipperService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShipperId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipperServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ShipperService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipperServiceServer).Delete(ctx, req.(*ShipperId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipperService_GetByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShipperByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipperServiceServer).GetByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ShipperService/GetByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipperServiceServer).GetByName(ctx, req.(*GetShipperByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShipperService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.ShipperService",
	HandlerType: (*ShipperServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ShipperService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ShipperService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ShipperService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ShipperService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ShipperService_Delete_Handler,
		},
		{
			MethodName: "GetByName",
			Handler:    _ShipperService_GetByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shipper_service.proto",
}
