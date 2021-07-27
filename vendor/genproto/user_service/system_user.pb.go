// Code generated by protoc-gen-go. DO NOT EDIT.
// source: system_user.proto

package user_service

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

type SystemUser struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Phone                string   `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	IsBlocked            bool     `protobuf:"varint,6,opt,name=is_blocked,json=isBlocked,proto3" json:"is_blocked,omitempty"`
	CreatedAt            string   `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SystemUser) Reset()         { *m = SystemUser{} }
func (m *SystemUser) String() string { return proto.CompactTextString(m) }
func (*SystemUser) ProtoMessage()    {}
func (*SystemUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_366d5dd5987cd68b, []int{0}
}

func (m *SystemUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemUser.Unmarshal(m, b)
}
func (m *SystemUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemUser.Marshal(b, m, deterministic)
}
func (m *SystemUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemUser.Merge(m, src)
}
func (m *SystemUser) XXX_Size() int {
	return xxx_messageInfo_SystemUser.Size(m)
}
func (m *SystemUser) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemUser.DiscardUnknown(m)
}

var xxx_messageInfo_SystemUser proto.InternalMessageInfo

func (m *SystemUser) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SystemUser) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SystemUser) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SystemUser) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *SystemUser) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *SystemUser) GetIsBlocked() bool {
	if m != nil {
		return m.IsBlocked
	}
	return false
}

func (m *SystemUser) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *SystemUser) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func init() {
	proto.RegisterType((*SystemUser)(nil), "genproto.SystemUser")
}

func init() { proto.RegisterFile("system_user.proto", fileDescriptor_366d5dd5987cd68b) }

var fileDescriptor_366d5dd5987cd68b = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x49, 0xdd, 0x5d, 0xd3, 0x39, 0x08, 0x06, 0xc5, 0x20, 0x08, 0x8b, 0xa7, 0x3d, 0xe9,
	0xc1, 0x27, 0xd8, 0x7d, 0x84, 0x8a, 0x17, 0x2f, 0x21, 0x6d, 0x06, 0x0d, 0xda, 0x26, 0x64, 0x52,
	0xc5, 0x67, 0xf5, 0x65, 0x24, 0x93, 0xd6, 0xdb, 0xff, 0x7f, 0x1f, 0x33, 0xcc, 0xc0, 0x25, 0xfd,
	0x50, 0xc6, 0xd1, 0xcc, 0x84, 0xe9, 0x21, 0xa6, 0x90, 0x83, 0x92, 0x6f, 0x38, 0x71, 0xba, 0xff,
	0x15, 0x00, 0xcf, 0xec, 0x5f, 0x08, 0x93, 0xba, 0x80, 0xc6, 0x3b, 0x2d, 0xf6, 0xe2, 0xd0, 0x76,
	0x8d, 0x77, 0x4a, 0xc1, 0x66, 0xb2, 0x23, 0xea, 0x86, 0x09, 0x67, 0x75, 0x0b, 0xb2, 0xac, 0x62,
	0x7e, 0xc6, 0xfc, 0xbf, 0x17, 0x17, 0x2d, 0xd1, 0x77, 0x48, 0x4e, 0x6f, 0xaa, 0x5b, 0xbb, 0xba,
	0x82, 0x6d, 0x7c, 0x0f, 0x13, 0xea, 0x2d, 0x8b, 0x5a, 0xd4, 0x1d, 0x80, 0x27, 0xd3, 0x7f, 0x86,
	0xe1, 0x03, 0x9d, 0xde, 0xed, 0xc5, 0x41, 0x76, 0xad, 0xa7, 0x53, 0x05, 0x45, 0x0f, 0x09, 0x6d,
	0x46, 0x67, 0x6c, 0xd6, 0xe7, 0x3c, 0xd9, 0x2e, 0xe4, 0x98, 0x8b, 0x9e, 0xa3, 0x5b, 0xb5, 0xac,
	0x7a, 0x21, 0xc7, 0x7c, 0xba, 0x79, 0xbd, 0x5e, 0x3f, 0x7d, 0x2c, 0x37, 0x1a, 0xc2, 0xf4, 0xe5,
	0x07, 0xec, 0x77, 0xcc, 0x9e, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x6e, 0x18, 0xb3, 0x1c,
	0x01, 0x00, 0x00,
}