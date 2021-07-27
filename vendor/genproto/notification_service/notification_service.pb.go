// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notification_service.proto

package notification_service

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

type ChangeStatusRequest struct {
	NotificationId       string   `protobuf:"bytes,1,opt,name=notification_id,json=notificationId,proto3" json:"notification_id,omitempty"`
	ReceiverId           string   `protobuf:"bytes,2,opt,name=receiver_id,json=receiverId,proto3" json:"receiver_id,omitempty"`
	StatusId             string   `protobuf:"bytes,3,opt,name=status_id,json=statusId,proto3" json:"status_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeStatusRequest) Reset()         { *m = ChangeStatusRequest{} }
func (m *ChangeStatusRequest) String() string { return proto.CompactTextString(m) }
func (*ChangeStatusRequest) ProtoMessage()    {}
func (*ChangeStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c524632753ba8b80, []int{0}
}

func (m *ChangeStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeStatusRequest.Unmarshal(m, b)
}
func (m *ChangeStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeStatusRequest.Marshal(b, m, deterministic)
}
func (m *ChangeStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeStatusRequest.Merge(m, src)
}
func (m *ChangeStatusRequest) XXX_Size() int {
	return xxx_messageInfo_ChangeStatusRequest.Size(m)
}
func (m *ChangeStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeStatusRequest proto.InternalMessageInfo

func (m *ChangeStatusRequest) GetNotificationId() string {
	if m != nil {
		return m.NotificationId
	}
	return ""
}

func (m *ChangeStatusRequest) GetReceiverId() string {
	if m != nil {
		return m.ReceiverId
	}
	return ""
}

func (m *ChangeStatusRequest) GetStatusId() string {
	if m != nil {
		return m.StatusId
	}
	return ""
}

type GetAllUserNotificationsRequest struct {
	ReceiverId           string   `protobuf:"bytes,1,opt,name=receiver_id,json=receiverId,proto3" json:"receiver_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllUserNotificationsRequest) Reset()         { *m = GetAllUserNotificationsRequest{} }
func (m *GetAllUserNotificationsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllUserNotificationsRequest) ProtoMessage()    {}
func (*GetAllUserNotificationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c524632753ba8b80, []int{1}
}

func (m *GetAllUserNotificationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllUserNotificationsRequest.Unmarshal(m, b)
}
func (m *GetAllUserNotificationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllUserNotificationsRequest.Marshal(b, m, deterministic)
}
func (m *GetAllUserNotificationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllUserNotificationsRequest.Merge(m, src)
}
func (m *GetAllUserNotificationsRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllUserNotificationsRequest.Size(m)
}
func (m *GetAllUserNotificationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllUserNotificationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllUserNotificationsRequest proto.InternalMessageInfo

func (m *GetAllUserNotificationsRequest) GetReceiverId() string {
	if m != nil {
		return m.ReceiverId
	}
	return ""
}

type GetAllUserNotificationsResponse struct {
	Notifications        []*Notification `protobuf:"bytes,1,rep,name=notifications,proto3" json:"notifications,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetAllUserNotificationsResponse) Reset()         { *m = GetAllUserNotificationsResponse{} }
func (m *GetAllUserNotificationsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllUserNotificationsResponse) ProtoMessage()    {}
func (*GetAllUserNotificationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c524632753ba8b80, []int{2}
}

func (m *GetAllUserNotificationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllUserNotificationsResponse.Unmarshal(m, b)
}
func (m *GetAllUserNotificationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllUserNotificationsResponse.Marshal(b, m, deterministic)
}
func (m *GetAllUserNotificationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllUserNotificationsResponse.Merge(m, src)
}
func (m *GetAllUserNotificationsResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllUserNotificationsResponse.Size(m)
}
func (m *GetAllUserNotificationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllUserNotificationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllUserNotificationsResponse proto.InternalMessageInfo

func (m *GetAllUserNotificationsResponse) GetNotifications() []*Notification {
	if m != nil {
		return m.Notifications
	}
	return nil
}

func init() {
	proto.RegisterType((*ChangeStatusRequest)(nil), "genproto.ChangeStatusRequest")
	proto.RegisterType((*GetAllUserNotificationsRequest)(nil), "genproto.GetAllUserNotificationsRequest")
	proto.RegisterType((*GetAllUserNotificationsResponse)(nil), "genproto.GetAllUserNotificationsResponse")
}

func init() { proto.RegisterFile("notification_service.proto", fileDescriptor_c524632753ba8b80) }

var fileDescriptor_c524632753ba8b80 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6b, 0xf2, 0x40,
	0x18, 0x84, 0x8d, 0x82, 0xe8, 0xeb, 0xf7, 0xb5, 0xb0, 0x82, 0x95, 0x88, 0x55, 0x72, 0xa9, 0xbd,
	0xac, 0x60, 0x6f, 0xa5, 0x17, 0x2b, 0x45, 0xbc, 0xf4, 0xa0, 0xf4, 0xd2, 0x8b, 0xc4, 0xe4, 0x35,
	0x5d, 0x48, 0xb3, 0xe9, 0xee, 0x46, 0x28, 0xf4, 0x27, 0xf4, 0x47, 0x97, 0xec, 0x92, 0xba, 0x8a,
	0x91, 0xde, 0x92, 0x99, 0xc9, 0x93, 0xcd, 0x4c, 0xc0, 0x4d, 0xb8, 0x62, 0x5b, 0x16, 0xf8, 0x8a,
	0xf1, 0x64, 0x2d, 0x51, 0xec, 0x58, 0x80, 0x34, 0x15, 0x5c, 0x71, 0xd2, 0x88, 0x30, 0xd1, 0x57,
	0x6e, 0x2f, 0xe2, 0x3c, 0x8a, 0x71, 0xac, 0xef, 0x36, 0xd9, 0x76, 0x8c, 0xef, 0xa9, 0xfa, 0x34,
	0x31, 0x97, 0xd8, 0x08, 0xa3, 0x79, 0x5f, 0xd0, 0x9e, 0xbd, 0xf9, 0x49, 0x84, 0x2b, 0xe5, 0xab,
	0x4c, 0x2e, 0xf1, 0x23, 0x43, 0xa9, 0xc8, 0x0d, 0x5c, 0x1e, 0xbc, 0x8f, 0x85, 0x5d, 0x67, 0xe8,
	0x8c, 0x9a, 0xcb, 0x0b, 0x5b, 0x5e, 0x84, 0x64, 0x00, 0x2d, 0x81, 0x01, 0xb2, 0x1d, 0x8a, 0x3c,
	0x54, 0xd5, 0x21, 0x28, 0xa4, 0x45, 0x48, 0x7a, 0xd0, 0x94, 0x1a, 0x9d, 0xdb, 0x35, 0x6d, 0x37,
	0x8c, 0xb0, 0x08, 0xbd, 0x29, 0x5c, 0xcf, 0x51, 0x4d, 0xe3, 0xf8, 0x45, 0xa2, 0x78, 0xb6, 0xc8,
	0xbf, 0x07, 0x39, 0xe2, 0x3b, 0xc7, 0x7c, 0x6f, 0x0d, 0x83, 0x52, 0x84, 0x4c, 0x79, 0x22, 0x91,
	0x3c, 0xc0, 0x7f, 0xfb, 0xd4, 0xb2, 0xeb, 0x0c, 0x6b, 0xa3, 0xd6, 0xa4, 0x43, 0x8b, 0xda, 0xa8,
	0xfd, 0xdc, 0xf2, 0x30, 0x3c, 0xf9, 0xae, 0x42, 0xdb, 0xf6, 0x57, 0xa6, 0x7a, 0x72, 0x0f, 0xf5,
	0x99, 0x40, 0x5f, 0x21, 0x29, 0x01, 0xb9, 0x1d, 0x6a, 0xd6, 0xa0, 0xc5, 0x1a, 0xf4, 0x29, 0x5f,
	0xc3, 0xab, 0x90, 0x39, 0xfc, 0xb3, 0x5b, 0x27, 0xfd, 0x3d, 0xe1, 0xc4, 0x1a, 0x67, 0x40, 0x29,
	0x5c, 0x95, 0x7c, 0x3d, 0x19, 0xed, 0x99, 0xe7, 0x3b, 0x76, 0x6f, 0xff, 0x90, 0x34, 0x55, 0x7a,
	0x95, 0xc7, 0xc1, 0x6b, 0xbf, 0x48, 0x8f, 0x4f, 0xfd, 0x92, 0x9b, 0xba, 0xf6, 0xee, 0x7e, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x10, 0xd4, 0xbe, 0x1a, 0xb1, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationServiceClient interface {
	Create(ctx context.Context, in *Notification, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ChangeStatus(ctx context.Context, in *ChangeStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAllUserNotifications(ctx context.Context, in *GetAllUserNotificationsRequest, opts ...grpc.CallOption) (*GetAllUserNotificationsResponse, error)
}

type notificationServiceClient struct {
	cc *grpc.ClientConn
}

func NewNotificationServiceClient(cc *grpc.ClientConn) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) Create(ctx context.Context, in *Notification, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/genproto.NotificationService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) ChangeStatus(ctx context.Context, in *ChangeStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/genproto.NotificationService/ChangeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) GetAllUserNotifications(ctx context.Context, in *GetAllUserNotificationsRequest, opts ...grpc.CallOption) (*GetAllUserNotificationsResponse, error) {
	out := new(GetAllUserNotificationsResponse)
	err := c.cc.Invoke(ctx, "/genproto.NotificationService/GetAllUserNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
type NotificationServiceServer interface {
	Create(context.Context, *Notification) (*emptypb.Empty, error)
	ChangeStatus(context.Context, *ChangeStatusRequest) (*emptypb.Empty, error)
	GetAllUserNotifications(context.Context, *GetAllUserNotificationsRequest) (*GetAllUserNotificationsResponse, error)
}

// UnimplementedNotificationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (*UnimplementedNotificationServiceServer) Create(ctx context.Context, req *Notification) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedNotificationServiceServer) ChangeStatus(ctx context.Context, req *ChangeStatusRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeStatus not implemented")
}
func (*UnimplementedNotificationServiceServer) GetAllUserNotifications(ctx context.Context, req *GetAllUserNotificationsRequest) (*GetAllUserNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUserNotifications not implemented")
}

func RegisterNotificationServiceServer(s *grpc.Server, srv NotificationServiceServer) {
	s.RegisterService(&_NotificationService_serviceDesc, srv)
}

func _NotificationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Notification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.NotificationService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).Create(ctx, req.(*Notification))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_ChangeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).ChangeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.NotificationService/ChangeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).ChangeStatus(ctx, req.(*ChangeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_GetAllUserNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUserNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).GetAllUserNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.NotificationService/GetAllUserNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).GetAllUserNotifications(ctx, req.(*GetAllUserNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NotificationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _NotificationService_Create_Handler,
		},
		{
			MethodName: "ChangeStatus",
			Handler:    _NotificationService_ChangeStatus_Handler,
		},
		{
			MethodName: "GetAllUserNotifications",
			Handler:    _NotificationService_GetAllUserNotifications_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notification_service.proto",
}