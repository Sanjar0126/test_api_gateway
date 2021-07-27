// Code generated by protoc-gen-go. DO NOT EDIT.
// source: branch.proto

package user_service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

type Branch struct {
	Id                   string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ShipperId            string                  `protobuf:"bytes,2,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
	Name                 string                  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Image                *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Phone                string                  `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	IsActive             bool                    `protobuf:"varint,6,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Address              string                  `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	Location             *Location               `protobuf:"bytes,8,opt,name=location,proto3" json:"location,omitempty"`
	CreatedAt            string                  `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string                  `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Destination          *wrapperspb.StringValue `protobuf:"bytes,11,opt,name=destination,proto3" json:"destination,omitempty"`
	WorkHourStart        string                  `protobuf:"bytes,12,opt,name=work_hour_start,json=workHourStart,proto3" json:"work_hour_start,omitempty"`
	WorkHourEnd          string                  `protobuf:"bytes,13,opt,name=work_hour_end,json=workHourEnd,proto3" json:"work_hour_end,omitempty"`
	JowiId               *wrapperspb.StringValue `protobuf:"bytes,14,opt,name=jowi_id,json=jowiId,proto3" json:"jowi_id,omitempty"`
	IikoId               *wrapperspb.StringValue `protobuf:"bytes,15,opt,name=iiko_id,json=iikoId,proto3" json:"iiko_id,omitempty"`
	IikoTerminalId       *wrapperspb.StringValue `protobuf:"bytes,16,opt,name=iiko_terminal_id,json=iikoTerminalId,proto3" json:"iiko_terminal_id,omitempty"`
	FareId               *wrapperspb.StringValue `protobuf:"bytes,17,opt,name=fare_id,json=fareId,proto3" json:"fare_id,omitempty"`
	TgChatId             *wrapperspb.StringValue `protobuf:"bytes,18,opt,name=tg_chat_id,json=tgChatId,proto3" json:"tg_chat_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Branch) Reset()         { *m = Branch{} }
func (m *Branch) String() string { return proto.CompactTextString(m) }
func (*Branch) ProtoMessage()    {}
func (*Branch) Descriptor() ([]byte, []int) {
	return fileDescriptor_20ce1f5884b7e047, []int{0}
}

func (m *Branch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Branch.Unmarshal(m, b)
}
func (m *Branch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Branch.Marshal(b, m, deterministic)
}
func (m *Branch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Branch.Merge(m, src)
}
func (m *Branch) XXX_Size() int {
	return xxx_messageInfo_Branch.Size(m)
}
func (m *Branch) XXX_DiscardUnknown() {
	xxx_messageInfo_Branch.DiscardUnknown(m)
}

var xxx_messageInfo_Branch proto.InternalMessageInfo

func (m *Branch) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Branch) GetShipperId() string {
	if m != nil {
		return m.ShipperId
	}
	return ""
}

func (m *Branch) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Branch) GetImage() *wrapperspb.StringValue {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *Branch) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Branch) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *Branch) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Branch) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Branch) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Branch) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Branch) GetDestination() *wrapperspb.StringValue {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *Branch) GetWorkHourStart() string {
	if m != nil {
		return m.WorkHourStart
	}
	return ""
}

func (m *Branch) GetWorkHourEnd() string {
	if m != nil {
		return m.WorkHourEnd
	}
	return ""
}

func (m *Branch) GetJowiId() *wrapperspb.StringValue {
	if m != nil {
		return m.JowiId
	}
	return nil
}

func (m *Branch) GetIikoId() *wrapperspb.StringValue {
	if m != nil {
		return m.IikoId
	}
	return nil
}

func (m *Branch) GetIikoTerminalId() *wrapperspb.StringValue {
	if m != nil {
		return m.IikoTerminalId
	}
	return nil
}

func (m *Branch) GetFareId() *wrapperspb.StringValue {
	if m != nil {
		return m.FareId
	}
	return nil
}

func (m *Branch) GetTgChatId() *wrapperspb.StringValue {
	if m != nil {
		return m.TgChatId
	}
	return nil
}

type Location struct {
	Long                 float64  `protobuf:"fixed64,1,opt,name=long,proto3" json:"long,omitempty"`
	Lat                  float64  `protobuf:"fixed64,2,opt,name=lat,proto3" json:"lat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_20ce1f5884b7e047, []int{1}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetLong() float64 {
	if m != nil {
		return m.Long
	}
	return 0
}

func (m *Location) GetLat() float64 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func init() {
	proto.RegisterType((*Branch)(nil), "genproto.Branch")
	proto.RegisterType((*Location)(nil), "genproto.Location")
}

func init() { proto.RegisterFile("branch.proto", fileDescriptor_20ce1f5884b7e047) }

var fileDescriptor_20ce1f5884b7e047 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x8f, 0xd3, 0x3c,
	0x10, 0xc6, 0x95, 0xee, 0xb6, 0x4d, 0xa7, 0xdb, 0x6e, 0xdf, 0xd1, 0x8b, 0xb0, 0xf8, 0xa7, 0xaa,
	0x07, 0xd4, 0x53, 0x8a, 0x16, 0x71, 0xe1, 0x80, 0xd4, 0x45, 0x20, 0x2a, 0x71, 0xca, 0x22, 0x0e,
	0x5c, 0x22, 0x37, 0xf6, 0x26, 0x66, 0x53, 0x3b, 0xb2, 0x9d, 0xed, 0xc7, 0xe6, 0x2b, 0x20, 0x4f,
	0x12, 0xe0, 0x98, 0xdb, 0xcc, 0x33, 0xcf, 0xaf, 0x8f, 0xa7, 0x19, 0xb8, 0x3a, 0x5a, 0xae, 0xf3,
	0x32, 0xa9, 0xad, 0xf1, 0x06, 0xe3, 0x42, 0x6a, 0xaa, 0x9e, 0xbd, 0x2a, 0x8c, 0x29, 0x2a, 0xb9,
	0xa3, 0xee, 0xd8, 0xdc, 0xef, 0xce, 0x96, 0xd7, 0xb5, 0xb4, 0xae, 0x75, 0x6e, 0x7e, 0x8d, 0x61,
	0x72, 0x4b, 0x28, 0x2e, 0x61, 0xa4, 0x04, 0x8b, 0xd6, 0xd1, 0x76, 0x96, 0x8e, 0x94, 0xc0, 0x97,
	0x00, 0xae, 0x54, 0xc1, 0x9c, 0x29, 0xc1, 0x46, 0xa4, 0xcf, 0x3a, 0xe5, 0x20, 0x10, 0xe1, 0x52,
	0xf3, 0x93, 0x64, 0x17, 0x34, 0xa0, 0x1a, 0x6f, 0x60, 0xac, 0x4e, 0xbc, 0x90, 0xec, 0x72, 0x1d,
	0x6d, 0xe7, 0x37, 0x2f, 0x92, 0x36, 0x3d, 0xe9, 0xd3, 0x93, 0x3b, 0x6f, 0x95, 0x2e, 0xbe, 0xf3,
	0xaa, 0x91, 0x69, 0x6b, 0xc5, 0xff, 0x61, 0x5c, 0x97, 0x46, 0x4b, 0x36, 0xa6, 0x1f, 0x6a, 0x1b,
	0x7c, 0x0e, 0x33, 0xe5, 0x32, 0x9e, 0x7b, 0xf5, 0x28, 0xd9, 0x64, 0x1d, 0x6d, 0xe3, 0x34, 0x56,
	0x6e, 0x4f, 0x3d, 0x32, 0x98, 0x72, 0x21, 0xac, 0x74, 0x8e, 0x4d, 0x09, 0xea, 0x5b, 0x4c, 0x20,
	0xae, 0x4c, 0xce, 0xbd, 0x32, 0x9a, 0xc5, 0xf4, 0x06, 0x4c, 0xfa, 0xff, 0x22, 0xf9, 0xda, 0x4d,
	0xd2, 0x3f, 0x9e, 0xb0, 0x63, 0x6e, 0x25, 0xf7, 0x52, 0x64, 0xdc, 0xb3, 0x59, 0xbb, 0x63, 0xa7,
	0xec, 0x7d, 0x18, 0x37, 0xb5, 0xe8, 0xc7, 0xd0, 0x8e, 0x3b, 0x65, 0xef, 0xf1, 0x03, 0xcc, 0x85,
	0x74, 0x5e, 0xe9, 0x36, 0x70, 0x3e, 0x60, 0xe9, 0x7f, 0x01, 0x7c, 0x0d, 0xd7, 0x67, 0x63, 0x1f,
	0xb2, 0xd2, 0x34, 0x36, 0x73, 0x9e, 0x5b, 0xcf, 0xae, 0x28, 0x63, 0x11, 0xe4, 0x2f, 0xa6, 0xb1,
	0x77, 0x41, 0xc4, 0x0d, 0x2c, 0xfe, 0xfa, 0xa4, 0x16, 0x6c, 0x41, 0xae, 0x79, 0xef, 0xfa, 0xa4,
	0x05, 0xbe, 0x83, 0xe9, 0x4f, 0x73, 0x56, 0xe1, 0x53, 0x2d, 0x07, 0xbc, 0x63, 0x12, 0xcc, 0x07,
	0xc2, 0x94, 0x7a, 0x30, 0x01, 0xbb, 0x1e, 0x82, 0x05, 0xf3, 0x41, 0xe0, 0x67, 0x58, 0x11, 0xe6,
	0xa5, 0x3d, 0x29, 0xcd, 0xab, 0xc0, 0xaf, 0x06, 0xf0, 0xcb, 0x40, 0x7d, 0xeb, 0xa0, 0x36, 0xfe,
	0x9e, 0x5b, 0x19, 0xf0, 0xff, 0x86, 0xc4, 0x07, 0xf3, 0x41, 0xe0, 0x7b, 0x00, 0x5f, 0x64, 0x79,
	0xc9, 0x7d, 0x20, 0x71, 0x00, 0x19, 0xfb, 0xe2, 0x63, 0xc9, 0xfd, 0x41, 0x6c, 0xde, 0x40, 0xdc,
	0x1f, 0x42, 0xb8, 0xe1, 0xca, 0xe8, 0x82, 0x8e, 0x3e, 0x4a, 0xa9, 0xc6, 0x15, 0x5c, 0x54, 0xdc,
	0xd3, 0xbd, 0x47, 0x69, 0x28, 0x6f, 0x9f, 0xfe, 0x78, 0xd2, 0xdf, 0xd0, 0xae, 0x71, 0xd2, 0x66,
	0x4e, 0xda, 0x47, 0x95, 0xcb, 0xe3, 0x84, 0xb4, 0xb7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x03,
	0x46, 0x66, 0xde, 0x7d, 0x03, 0x00, 0x00,
}
