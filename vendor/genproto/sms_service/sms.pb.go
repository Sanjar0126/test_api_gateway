// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sms.proto

package sms_service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Sms struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ShipperId            string   `protobuf:"bytes,2,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
	CreatedAt            string   `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Text                 string   `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	Recipients           []string `protobuf:"bytes,5,rep,name=recipients,proto3" json:"recipients,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sms) Reset()         { *m = Sms{} }
func (m *Sms) String() string { return proto.CompactTextString(m) }
func (*Sms) ProtoMessage()    {}
func (*Sms) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8d8bdc537111860, []int{0}
}

func (m *Sms) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sms.Unmarshal(m, b)
}
func (m *Sms) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sms.Marshal(b, m, deterministic)
}
func (m *Sms) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sms.Merge(m, src)
}
func (m *Sms) XXX_Size() int {
	return xxx_messageInfo_Sms.Size(m)
}
func (m *Sms) XXX_DiscardUnknown() {
	xxx_messageInfo_Sms.DiscardUnknown(m)
}

var xxx_messageInfo_Sms proto.InternalMessageInfo

func (m *Sms) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Sms) GetShipperId() string {
	if m != nil {
		return m.ShipperId
	}
	return ""
}

func (m *Sms) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Sms) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Sms) GetRecipients() []string {
	if m != nil {
		return m.Recipients
	}
	return nil
}

func init() {
	proto.RegisterType((*Sms)(nil), "genproto.Sms")
}

func init() { proto.RegisterFile("sms.proto", fileDescriptor_c8d8bdc537111860) }

var fileDescriptor_c8d8bdc537111860 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0xce, 0x2d, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0x4f, 0xcd, 0x03, 0xb3, 0x94, 0xda, 0x19, 0xb9,
	0x98, 0x83, 0x73, 0x8b, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x83, 0x98, 0x32, 0x53, 0x84, 0x64, 0xb9, 0xb8, 0x8a, 0x33, 0x32, 0x0b, 0x0a, 0x52, 0x8b, 0xe2,
	0x33, 0x53, 0x24, 0x98, 0xc0, 0xe2, 0x9c, 0x50, 0x11, 0x4f, 0xb0, 0x74, 0x72, 0x51, 0x6a, 0x62,
	0x49, 0x6a, 0x4a, 0x7c, 0x62, 0x89, 0x04, 0x33, 0x44, 0x1a, 0x2a, 0xe2, 0x58, 0x22, 0x24, 0xc4,
	0xc5, 0x52, 0x92, 0x5a, 0x51, 0x22, 0xc1, 0x02, 0x96, 0x00, 0xb3, 0x85, 0xe4, 0xb8, 0xb8, 0x8a,
	0x52, 0x93, 0x33, 0x0b, 0x32, 0x53, 0xf3, 0x4a, 0x8a, 0x25, 0x58, 0x15, 0x98, 0x35, 0x38, 0x83,
	0x90, 0x44, 0x9c, 0xc4, 0xa2, 0x44, 0x60, 0xae, 0xd2, 0x2f, 0xce, 0x2d, 0x8e, 0x2f, 0x4e, 0x2d,
	0x2a, 0xcb, 0x4c, 0x4e, 0x4d, 0x62, 0x03, 0x0b, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x24,
	0x1f, 0x9f, 0x1f, 0xbf, 0x00, 0x00, 0x00,
}
