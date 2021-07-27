// Code generated by protoc-gen-go. DO NOT EDIT.
// source: system_user_service.proto

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

type SystemUserId struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SystemUserId) Reset()         { *m = SystemUserId{} }
func (m *SystemUserId) String() string { return proto.CompactTextString(m) }
func (*SystemUserId) ProtoMessage()    {}
func (*SystemUserId) Descriptor() ([]byte, []int) {
	return fileDescriptor_6120f12d5c23b9ee, []int{0}
}

func (m *SystemUserId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemUserId.Unmarshal(m, b)
}
func (m *SystemUserId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemUserId.Marshal(b, m, deterministic)
}
func (m *SystemUserId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemUserId.Merge(m, src)
}
func (m *SystemUserId) XXX_Size() int {
	return xxx_messageInfo_SystemUserId.Size(m)
}
func (m *SystemUserId) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemUserId.DiscardUnknown(m)
}

var xxx_messageInfo_SystemUserId proto.InternalMessageInfo

func (m *SystemUserId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetSystemUserByUsernameRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSystemUserByUsernameRequest) Reset()         { *m = GetSystemUserByUsernameRequest{} }
func (m *GetSystemUserByUsernameRequest) String() string { return proto.CompactTextString(m) }
func (*GetSystemUserByUsernameRequest) ProtoMessage()    {}
func (*GetSystemUserByUsernameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6120f12d5c23b9ee, []int{1}
}

func (m *GetSystemUserByUsernameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSystemUserByUsernameRequest.Unmarshal(m, b)
}
func (m *GetSystemUserByUsernameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSystemUserByUsernameRequest.Marshal(b, m, deterministic)
}
func (m *GetSystemUserByUsernameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSystemUserByUsernameRequest.Merge(m, src)
}
func (m *GetSystemUserByUsernameRequest) XXX_Size() int {
	return xxx_messageInfo_GetSystemUserByUsernameRequest.Size(m)
}
func (m *GetSystemUserByUsernameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSystemUserByUsernameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSystemUserByUsernameRequest proto.InternalMessageInfo

func (m *GetSystemUserByUsernameRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type GetSystemUserByCredentialsRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSystemUserByCredentialsRequest) Reset()         { *m = GetSystemUserByCredentialsRequest{} }
func (m *GetSystemUserByCredentialsRequest) String() string { return proto.CompactTextString(m) }
func (*GetSystemUserByCredentialsRequest) ProtoMessage()    {}
func (*GetSystemUserByCredentialsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6120f12d5c23b9ee, []int{2}
}

func (m *GetSystemUserByCredentialsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSystemUserByCredentialsRequest.Unmarshal(m, b)
}
func (m *GetSystemUserByCredentialsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSystemUserByCredentialsRequest.Marshal(b, m, deterministic)
}
func (m *GetSystemUserByCredentialsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSystemUserByCredentialsRequest.Merge(m, src)
}
func (m *GetSystemUserByCredentialsRequest) XXX_Size() int {
	return xxx_messageInfo_GetSystemUserByCredentialsRequest.Size(m)
}
func (m *GetSystemUserByCredentialsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSystemUserByCredentialsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSystemUserByCredentialsRequest proto.InternalMessageInfo

func (m *GetSystemUserByCredentialsRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetSystemUserByCredentialsRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type ChangeSystemUserPasswordRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeSystemUserPasswordRequest) Reset()         { *m = ChangeSystemUserPasswordRequest{} }
func (m *ChangeSystemUserPasswordRequest) String() string { return proto.CompactTextString(m) }
func (*ChangeSystemUserPasswordRequest) ProtoMessage()    {}
func (*ChangeSystemUserPasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6120f12d5c23b9ee, []int{3}
}

func (m *ChangeSystemUserPasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeSystemUserPasswordRequest.Unmarshal(m, b)
}
func (m *ChangeSystemUserPasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeSystemUserPasswordRequest.Marshal(b, m, deterministic)
}
func (m *ChangeSystemUserPasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeSystemUserPasswordRequest.Merge(m, src)
}
func (m *ChangeSystemUserPasswordRequest) XXX_Size() int {
	return xxx_messageInfo_ChangeSystemUserPasswordRequest.Size(m)
}
func (m *ChangeSystemUserPasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeSystemUserPasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeSystemUserPasswordRequest proto.InternalMessageInfo

func (m *ChangeSystemUserPasswordRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ChangeSystemUserPasswordRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type GetAllSystemUsersRequest struct {
	Page                 uint64   `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit                uint64   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search               string   `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllSystemUsersRequest) Reset()         { *m = GetAllSystemUsersRequest{} }
func (m *GetAllSystemUsersRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllSystemUsersRequest) ProtoMessage()    {}
func (*GetAllSystemUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6120f12d5c23b9ee, []int{4}
}

func (m *GetAllSystemUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllSystemUsersRequest.Unmarshal(m, b)
}
func (m *GetAllSystemUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllSystemUsersRequest.Marshal(b, m, deterministic)
}
func (m *GetAllSystemUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllSystemUsersRequest.Merge(m, src)
}
func (m *GetAllSystemUsersRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllSystemUsersRequest.Size(m)
}
func (m *GetAllSystemUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllSystemUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllSystemUsersRequest proto.InternalMessageInfo

func (m *GetAllSystemUsersRequest) GetPage() uint64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetAllSystemUsersRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetAllSystemUsersRequest) GetSearch() string {
	if m != nil {
		return m.Search
	}
	return ""
}

type GetAllSystemUsersResponse struct {
	SystemUsers          []*SystemUser `protobuf:"bytes,1,rep,name=system_users,json=systemUsers,proto3" json:"system_users,omitempty"`
	Count                uint64        `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetAllSystemUsersResponse) Reset()         { *m = GetAllSystemUsersResponse{} }
func (m *GetAllSystemUsersResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllSystemUsersResponse) ProtoMessage()    {}
func (*GetAllSystemUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6120f12d5c23b9ee, []int{5}
}

func (m *GetAllSystemUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllSystemUsersResponse.Unmarshal(m, b)
}
func (m *GetAllSystemUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllSystemUsersResponse.Marshal(b, m, deterministic)
}
func (m *GetAllSystemUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllSystemUsersResponse.Merge(m, src)
}
func (m *GetAllSystemUsersResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllSystemUsersResponse.Size(m)
}
func (m *GetAllSystemUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllSystemUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllSystemUsersResponse proto.InternalMessageInfo

func (m *GetAllSystemUsersResponse) GetSystemUsers() []*SystemUser {
	if m != nil {
		return m.SystemUsers
	}
	return nil
}

func (m *GetAllSystemUsersResponse) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type DeleteSystemUserRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ShipperId            string   `protobuf:"bytes,2,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteSystemUserRequest) Reset()         { *m = DeleteSystemUserRequest{} }
func (m *DeleteSystemUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteSystemUserRequest) ProtoMessage()    {}
func (*DeleteSystemUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6120f12d5c23b9ee, []int{6}
}

func (m *DeleteSystemUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSystemUserRequest.Unmarshal(m, b)
}
func (m *DeleteSystemUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSystemUserRequest.Marshal(b, m, deterministic)
}
func (m *DeleteSystemUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSystemUserRequest.Merge(m, src)
}
func (m *DeleteSystemUserRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteSystemUserRequest.Size(m)
}
func (m *DeleteSystemUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSystemUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSystemUserRequest proto.InternalMessageInfo

func (m *DeleteSystemUserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeleteSystemUserRequest) GetShipperId() string {
	if m != nil {
		return m.ShipperId
	}
	return ""
}

func init() {
	proto.RegisterType((*SystemUserId)(nil), "genproto.SystemUserId")
	proto.RegisterType((*GetSystemUserByUsernameRequest)(nil), "genproto.GetSystemUserByUsernameRequest")
	proto.RegisterType((*GetSystemUserByCredentialsRequest)(nil), "genproto.GetSystemUserByCredentialsRequest")
	proto.RegisterType((*ChangeSystemUserPasswordRequest)(nil), "genproto.ChangeSystemUserPasswordRequest")
	proto.RegisterType((*GetAllSystemUsersRequest)(nil), "genproto.GetAllSystemUsersRequest")
	proto.RegisterType((*GetAllSystemUsersResponse)(nil), "genproto.GetAllSystemUsersResponse")
	proto.RegisterType((*DeleteSystemUserRequest)(nil), "genproto.DeleteSystemUserRequest")
}

func init() { proto.RegisterFile("system_user_service.proto", fileDescriptor_6120f12d5c23b9ee) }

var fileDescriptor_6120f12d5c23b9ee = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x86, 0xfb, 0x45, 0xb5, 0x9d, 0x8d, 0x89, 0x1d, 0x8d, 0x2d, 0x0b, 0x62, 0x6c, 0xe6, 0xa6,
	0x08, 0x29, 0x95, 0x86, 0x10, 0x5c, 0x70, 0xc3, 0xca, 0x54, 0x7a, 0x81, 0x04, 0xa9, 0x2a, 0x24,
	0x40, 0xaa, 0xb2, 0xe6, 0x90, 0x06, 0xa5, 0x49, 0xb0, 0x1d, 0x50, 0x7f, 0x2c, 0xff, 0x05, 0xc5,
	0x4e, 0xea, 0x30, 0x9a, 0xc0, 0x4d, 0xe5, 0xf3, 0xf5, 0xe8, 0xf8, 0x75, 0xdf, 0xc0, 0xa9, 0x58,
	0x0b, 0x49, 0xab, 0x79, 0x26, 0x88, 0xcf, 0x05, 0xf1, 0x1f, 0xe1, 0x82, 0x9c, 0x94, 0x27, 0x32,
	0xc1, 0x9d, 0x80, 0x62, 0x75, 0xb2, 0x0f, 0x2b, 0x4d, 0xba, 0x68, 0x3f, 0x08, 0x92, 0x24, 0x88,
	0x68, 0xa8, 0xa2, 0x9b, 0xec, 0xeb, 0x90, 0x56, 0xa9, 0x5c, 0xeb, 0x22, 0x3b, 0x83, 0xfd, 0xa9,
	0x9a, 0x98, 0x09, 0xe2, 0x13, 0x1f, 0x0f, 0xa0, 0x13, 0xfa, 0x56, 0xfb, 0xbc, 0x3d, 0xd8, 0x75,
	0x3b, 0xa1, 0xcf, 0x5e, 0xc1, 0xd9, 0x98, 0xa4, 0x69, 0xb9, 0x5a, 0xe7, 0xbf, 0xb1, 0xb7, 0x22,
	0x97, 0xbe, 0x67, 0x24, 0x24, 0xda, 0xb0, 0x93, 0x15, 0xa9, 0x62, 0x6e, 0x13, 0xb3, 0xcf, 0x70,
	0x71, 0x6b, 0x7a, 0xc4, 0xc9, 0xa7, 0x58, 0x86, 0x5e, 0x24, 0xfe, 0x03, 0x90, 0xd7, 0x52, 0x4f,
	0x88, 0x9f, 0x09, 0xf7, 0xad, 0x8e, 0xae, 0x95, 0x31, 0x7b, 0x07, 0x8f, 0x46, 0x4b, 0x2f, 0x0e,
	0xc8, 0xf0, 0xdf, 0x17, 0xb5, 0x12, 0x7d, 0xeb, 0x36, 0x8d, 0xb8, 0x2f, 0x60, 0x8d, 0x49, 0xbe,
	0x8e, 0x22, 0x83, 0xdb, 0xac, 0x88, 0xd0, 0x4b, 0xbd, 0x40, 0xaf, 0xd7, 0x73, 0xd5, 0x19, 0x8f,
	0xe0, 0x4e, 0x14, 0xae, 0x42, 0xa9, 0x40, 0x3d, 0x57, 0x07, 0x78, 0x0c, 0x7d, 0x41, 0x1e, 0x5f,
	0x2c, 0xad, 0xae, 0xe2, 0x17, 0x11, 0xfb, 0x06, 0xa7, 0x5b, 0xe8, 0x22, 0x4d, 0x62, 0x41, 0xf8,
	0x02, 0xf6, 0x2b, 0xcf, 0x26, 0xac, 0xf6, 0x79, 0x77, 0xb0, 0x77, 0x79, 0xe4, 0x94, 0xaf, 0xea,
	0x98, 0x21, 0x77, 0x4f, 0x18, 0x40, 0xbe, 0xc3, 0x22, 0xc9, 0xe2, 0xcd, 0x0e, 0x2a, 0x60, 0x6f,
	0xe1, 0xe4, 0x0d, 0x45, 0x24, 0x2b, 0xc2, 0xd4, 0x09, 0xf2, 0x10, 0x40, 0x2c, 0xc3, 0x34, 0x25,
	0x3e, 0x0f, 0x4b, 0x49, 0x76, 0x8b, 0xcc, 0xc4, 0xbf, 0xfc, 0xd5, 0x83, 0x43, 0x03, 0x99, 0xea,
	0xff, 0x1c, 0xbe, 0x84, 0xfe, 0x88, 0x93, 0x27, 0x09, 0xb7, 0xae, 0x68, 0x1f, 0x6f, 0xcb, 0x4e,
	0x7c, 0xd6, 0xc2, 0xe7, 0xd0, 0x1d, 0x93, 0xc4, 0x9a, 0x06, 0x7b, 0x2b, 0x8e, 0xb5, 0x70, 0x0a,
	0x7d, 0x2d, 0x1e, 0x32, 0xd3, 0x51, 0xf7, 0x58, 0xf6, 0xe3, 0xc6, 0x1e, 0x2d, 0x39, 0x6b, 0xe5,
	0xb7, 0x98, 0xa5, 0x7e, 0xe3, 0x2d, 0x94, 0x6f, 0x9c, 0xd2, 0x37, 0xce, 0x75, 0xee, 0x1b, 0xd6,
	0xc2, 0x6b, 0xe8, 0x6b, 0x7d, 0xf1, 0xc2, 0x4c, 0xd6, 0x28, 0xde, 0x80, 0xf9, 0x00, 0x77, 0xc7,
	0x24, 0x8d, 0xa1, 0x70, 0xf0, 0xc7, 0xe2, 0x0d, 0x9e, 0xab, 0x15, 0xea, 0x23, 0xdc, 0x53, 0xc8,
	0x8a, 0xcb, 0xf0, 0x69, 0x2d, 0xf5, 0x6f, 0x2f, 0xd6, 0x82, 0x67, 0x70, 0xa0, 0xbd, 0x56, 0x3a,
	0x0c, 0x9f, 0x98, 0xce, 0x7f, 0xb8, 0xb0, 0x5e, 0x82, 0xab, 0x93, 0x4f, 0xf7, 0x4b, 0xca, 0xb0,
	0xfa, 0x59, 0xbb, 0xe9, 0xab, 0xdc, 0xb3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x12, 0x6c, 0xe9,
	0x81, 0xf4, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SystemUserServiceClient is the client API for SystemUserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SystemUserServiceClient interface {
	Create(ctx context.Context, in *SystemUser, opts ...grpc.CallOption) (*SystemUserId, error)
	Get(ctx context.Context, in *SystemUserId, opts ...grpc.CallOption) (*SystemUser, error)
	GetAll(ctx context.Context, in *GetAllSystemUsersRequest, opts ...grpc.CallOption) (*GetAllSystemUsersResponse, error)
	Update(ctx context.Context, in *SystemUser, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *DeleteSystemUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetByUsername(ctx context.Context, in *GetSystemUserByUsernameRequest, opts ...grpc.CallOption) (*SystemUser, error)
	GetByCredentials(ctx context.Context, in *GetSystemUserByCredentialsRequest, opts ...grpc.CallOption) (*SystemUser, error)
	ChangePassword(ctx context.Context, in *ChangeSystemUserPasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type systemUserServiceClient struct {
	cc *grpc.ClientConn
}

func NewSystemUserServiceClient(cc *grpc.ClientConn) SystemUserServiceClient {
	return &systemUserServiceClient{cc}
}

func (c *systemUserServiceClient) Create(ctx context.Context, in *SystemUser, opts ...grpc.CallOption) (*SystemUserId, error) {
	out := new(SystemUserId)
	err := c.cc.Invoke(ctx, "/genproto.SystemUserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) Get(ctx context.Context, in *SystemUserId, opts ...grpc.CallOption) (*SystemUser, error) {
	out := new(SystemUser)
	err := c.cc.Invoke(ctx, "/genproto.SystemUserService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) GetAll(ctx context.Context, in *GetAllSystemUsersRequest, opts ...grpc.CallOption) (*GetAllSystemUsersResponse, error) {
	out := new(GetAllSystemUsersResponse)
	err := c.cc.Invoke(ctx, "/genproto.SystemUserService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) Update(ctx context.Context, in *SystemUser, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/genproto.SystemUserService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) Delete(ctx context.Context, in *DeleteSystemUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/genproto.SystemUserService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) GetByUsername(ctx context.Context, in *GetSystemUserByUsernameRequest, opts ...grpc.CallOption) (*SystemUser, error) {
	out := new(SystemUser)
	err := c.cc.Invoke(ctx, "/genproto.SystemUserService/GetByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) GetByCredentials(ctx context.Context, in *GetSystemUserByCredentialsRequest, opts ...grpc.CallOption) (*SystemUser, error) {
	out := new(SystemUser)
	err := c.cc.Invoke(ctx, "/genproto.SystemUserService/GetByCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) ChangePassword(ctx context.Context, in *ChangeSystemUserPasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/genproto.SystemUserService/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemUserServiceServer is the server API for SystemUserService service.
type SystemUserServiceServer interface {
	Create(context.Context, *SystemUser) (*SystemUserId, error)
	Get(context.Context, *SystemUserId) (*SystemUser, error)
	GetAll(context.Context, *GetAllSystemUsersRequest) (*GetAllSystemUsersResponse, error)
	Update(context.Context, *SystemUser) (*emptypb.Empty, error)
	Delete(context.Context, *DeleteSystemUserRequest) (*emptypb.Empty, error)
	GetByUsername(context.Context, *GetSystemUserByUsernameRequest) (*SystemUser, error)
	GetByCredentials(context.Context, *GetSystemUserByCredentialsRequest) (*SystemUser, error)
	ChangePassword(context.Context, *ChangeSystemUserPasswordRequest) (*emptypb.Empty, error)
}

// UnimplementedSystemUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSystemUserServiceServer struct {
}

func (*UnimplementedSystemUserServiceServer) Create(ctx context.Context, req *SystemUser) (*SystemUserId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedSystemUserServiceServer) Get(ctx context.Context, req *SystemUserId) (*SystemUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedSystemUserServiceServer) GetAll(ctx context.Context, req *GetAllSystemUsersRequest) (*GetAllSystemUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedSystemUserServiceServer) Update(ctx context.Context, req *SystemUser) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedSystemUserServiceServer) Delete(ctx context.Context, req *DeleteSystemUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedSystemUserServiceServer) GetByUsername(ctx context.Context, req *GetSystemUserByUsernameRequest) (*SystemUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByUsername not implemented")
}
func (*UnimplementedSystemUserServiceServer) GetByCredentials(ctx context.Context, req *GetSystemUserByCredentialsRequest) (*SystemUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByCredentials not implemented")
}
func (*UnimplementedSystemUserServiceServer) ChangePassword(ctx context.Context, req *ChangeSystemUserPasswordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}

func RegisterSystemUserServiceServer(s *grpc.Server, srv SystemUserServiceServer) {
	s.RegisterService(&_SystemUserService_serviceDesc, srv)
}

func _SystemUserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.SystemUserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).Create(ctx, req.(*SystemUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.SystemUserService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).Get(ctx, req.(*SystemUserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllSystemUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.SystemUserService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).GetAll(ctx, req.(*GetAllSystemUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.SystemUserService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).Update(ctx, req.(*SystemUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSystemUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.SystemUserService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).Delete(ctx, req.(*DeleteSystemUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_GetByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSystemUserByUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).GetByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.SystemUserService/GetByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).GetByUsername(ctx, req.(*GetSystemUserByUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_GetByCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSystemUserByCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).GetByCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.SystemUserService/GetByCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).GetByCredentials(ctx, req.(*GetSystemUserByCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeSystemUserPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.SystemUserService/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).ChangePassword(ctx, req.(*ChangeSystemUserPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SystemUserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.SystemUserService",
	HandlerType: (*SystemUserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SystemUserService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _SystemUserService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _SystemUserService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SystemUserService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SystemUserService_Delete_Handler,
		},
		{
			MethodName: "GetByUsername",
			Handler:    _SystemUserService_GetByUsername_Handler,
		},
		{
			MethodName: "GetByCredentials",
			Handler:    _SystemUserService_GetByCredentials_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _SystemUserService_ChangePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "system_user_service.proto",
}
