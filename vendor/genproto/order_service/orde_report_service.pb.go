// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orde_report_service.proto

package order_service

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

func init() { proto.RegisterFile("orde_report_service.proto", fileDescriptor_2b5b3f2fe02ec058) }

var fileDescriptor_2b5b3f2fe02ec058 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0x2f, 0x4a, 0x49,
	0x8d, 0x2f, 0x4a, 0x2d, 0xc8, 0x2f, 0x2a, 0x89, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0x4f, 0xcd, 0x03, 0xb3, 0xa4, 0x84, 0x40, 0x8a,
	0x8a, 0xa0, 0xaa, 0x20, 0xb2, 0x46, 0x3b, 0x18, 0xb9, 0x84, 0xfc, 0x41, 0xc2, 0x41, 0x60, 0xd1,
	0x60, 0x88, 0x56, 0x21, 0x1f, 0x2e, 0x5e, 0xf7, 0xd4, 0x12, 0xdf, 0xc4, 0xcc, 0x3c, 0x88, 0xb8,
	0x90, 0xb4, 0x1e, 0xcc, 0x18, 0x3d, 0x84, 0x68, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x94,
	0x0c, 0x76, 0xc9, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x25, 0x06, 0xa1, 0x10, 0x2e, 0x01, 0xf7,
	0xd4, 0x12, 0xe7, 0xfc, 0xd2, 0xa2, 0x4c, 0x98, 0x45, 0x42, 0x72, 0x08, 0x3d, 0x28, 0x12, 0x30,
	0x33, 0x71, 0xcb, 0x17, 0x17, 0xe4, 0x83, 0x4c, 0x75, 0x92, 0x88, 0x12, 0x83, 0x29, 0xd1, 0x87,
	0xf8, 0x0c, 0xea, 0xf1, 0x24, 0x36, 0xb0, 0xa0, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x96,
	0x20, 0x6f, 0x16, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrderReportServiceClient is the client API for OrderReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderReportServiceClient interface {
	GetMainReport(ctx context.Context, in *MainReportRequest, opts ...grpc.CallOption) (*MainReportResponse, error)
	GetCourierReport(ctx context.Context, in *CourierReportRequest, opts ...grpc.CallOption) (*CourierReportRespose, error)
}

type orderReportServiceClient struct {
	cc *grpc.ClientConn
}

func NewOrderReportServiceClient(cc *grpc.ClientConn) OrderReportServiceClient {
	return &orderReportServiceClient{cc}
}

func (c *orderReportServiceClient) GetMainReport(ctx context.Context, in *MainReportRequest, opts ...grpc.CallOption) (*MainReportResponse, error) {
	out := new(MainReportResponse)
	err := c.cc.Invoke(ctx, "/genproto.OrderReportService/GetMainReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderReportServiceClient) GetCourierReport(ctx context.Context, in *CourierReportRequest, opts ...grpc.CallOption) (*CourierReportRespose, error) {
	out := new(CourierReportRespose)
	err := c.cc.Invoke(ctx, "/genproto.OrderReportService/GetCourierReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderReportServiceServer is the server API for OrderReportService service.
type OrderReportServiceServer interface {
	GetMainReport(context.Context, *MainReportRequest) (*MainReportResponse, error)
	GetCourierReport(context.Context, *CourierReportRequest) (*CourierReportRespose, error)
}

// UnimplementedOrderReportServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOrderReportServiceServer struct {
}

func (*UnimplementedOrderReportServiceServer) GetMainReport(ctx context.Context, req *MainReportRequest) (*MainReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMainReport not implemented")
}
func (*UnimplementedOrderReportServiceServer) GetCourierReport(ctx context.Context, req *CourierReportRequest) (*CourierReportRespose, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourierReport not implemented")
}

func RegisterOrderReportServiceServer(s *grpc.Server, srv OrderReportServiceServer) {
	s.RegisterService(&_OrderReportService_serviceDesc, srv)
}

func _OrderReportService_GetMainReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MainReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderReportServiceServer).GetMainReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.OrderReportService/GetMainReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderReportServiceServer).GetMainReport(ctx, req.(*MainReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderReportService_GetCourierReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourierReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderReportServiceServer).GetCourierReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.OrderReportService/GetCourierReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderReportServiceServer).GetCourierReport(ctx, req.(*CourierReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderReportService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.OrderReportService",
	HandlerType: (*OrderReportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMainReport",
			Handler:    _OrderReportService_GetMainReport_Handler,
		},
		{
			MethodName: "GetCourierReport",
			Handler:    _OrderReportService_GetCourierReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orde_report_service.proto",
}