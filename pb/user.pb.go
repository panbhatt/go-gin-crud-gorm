// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type UserRole int32

const (
	User_user  UserRole = 0
	User_admin UserRole = 1
)

var UserRole_name = map[int32]string{
	0: "user",
	1: "admin",
}

var UserRole_value = map[string]int32{
	"user":  0,
	"admin": 1,
}

func (x UserRole) String() string {
	return proto.EnumName(UserRole_name, int32(x))
}

func (UserRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0, 0}
}

type User struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string               `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type UserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type GenericResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenericResponse) Reset()         { *m = GenericResponse{} }
func (m *GenericResponse) String() string { return proto.CompactTextString(m) }
func (*GenericResponse) ProtoMessage()    {}
func (*GenericResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *GenericResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenericResponse.Unmarshal(m, b)
}
func (m *GenericResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenericResponse.Marshal(b, m, deterministic)
}
func (m *GenericResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericResponse.Merge(m, src)
}
func (m *GenericResponse) XXX_Size() int {
	return xxx_messageInfo_GenericResponse.Size(m)
}
func (m *GenericResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenericResponse proto.InternalMessageInfo

func (m *GenericResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *GenericResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("pb.UserRole", UserRole_name, UserRole_value)
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*UserResponse)(nil), "pb.UserResponse")
	proto.RegisterType((*GenericResponse)(nil), "pb.GenericResponse")
}

func init() {
	proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf)
}

var fileDescriptor_116e343673f7ffaf = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x18, 0xc5, 0x6d, 0xed, 0xe6, 0xf6, 0x4d, 0x74, 0x04, 0x91, 0x32, 0x05, 0x47, 0x2f, 0x4e, 0x70,
	0x29, 0xcc, 0x93, 0xc7, 0xe9, 0xc1, 0xfb, 0xd0, 0x8b, 0x17, 0x49, 0xda, 0xcf, 0x2c, 0xb0, 0x34,
	0x21, 0xf9, 0xfa, 0x97, 0xfa, 0x0f, 0xc9, 0xd2, 0xd6, 0xab, 0xb7, 0xbc, 0x97, 0xf7, 0x78, 0xfc,
	0x3e, 0x80, 0x36, 0xa0, 0xe7, 0xce, 0x5b, 0xb2, 0x2c, 0x75, 0x72, 0x71, 0xa7, 0xac, 0x55, 0x07,
	0x2c, 0xa3, 0x23, 0xdb, 0xef, 0x92, 0xb4, 0xc1, 0x40, 0xc2, 0xb8, 0x2e, 0x54, 0xfc, 0x24, 0x90,
	0x7d, 0x04, 0xf4, 0xec, 0x02, 0x52, 0x5d, 0xe7, 0xc9, 0x32, 0x59, 0x4d, 0x77, 0xa9, 0xae, 0x19,
	0x83, 0xac, 0x11, 0x06, 0xf3, 0x34, 0x3a, 0xf1, 0xcd, 0xae, 0x60, 0x84, 0x46, 0xe8, 0x43, 0x7e,
	0x1a, 0xcd, 0x4e, 0xb0, 0x67, 0x80, 0xca, 0xa3, 0x20, 0xac, 0xbf, 0x04, 0xe5, 0xa3, 0x65, 0xb2,
	0x9a, 0x6d, 0x16, 0xbc, 0x1b, 0xe6, 0xc3, 0x30, 0x7f, 0x1f, 0x86, 0x77, 0xd3, 0x3e, 0xbd, 0xa5,
	0x63, 0xb5, 0x75, 0xf5, 0x50, 0x1d, 0xff, 0x5f, 0xed, 0xd3, 0x5b, 0x2a, 0x6e, 0x20, 0xf3, 0xf6,
	0x80, 0x6c, 0x02, 0xd9, 0x91, 0x79, 0x7e, 0xc2, 0xa6, 0x30, 0x12, 0xb5, 0xd1, 0xcd, 0x3c, 0x29,
	0x1e, 0xe1, 0xfc, 0x08, 0xb5, 0xc3, 0xe0, 0x6c, 0x13, 0x90, 0xdd, 0x76, 0xa1, 0x88, 0x37, 0xdb,
	0x4c, 0xb8, 0x93, 0x3c, 0xfe, 0x47, 0xb7, 0x78, 0x85, 0xcb, 0x37, 0x6c, 0xd0, 0xeb, 0xea, 0xaf,
	0x70, 0x0d, 0xe3, 0x40, 0x82, 0xda, 0xd0, 0x5f, 0xa4, 0x57, 0x2c, 0x87, 0x33, 0x83, 0x21, 0x08,
	0x35, 0x1c, 0x66, 0x90, 0x2f, 0x0f, 0x9f, 0xf7, 0x4a, 0xd3, 0xbe, 0x95, 0xbc, 0xb2, 0xa6, 0x74,
	0xa2, 0x91, 0x7b, 0x41, 0x54, 0x2a, 0xbb, 0x56, 0xba, 0x59, 0x57, 0xbe, 0xad, 0xd7, 0xca, 0x7a,
	0x53, 0x3a, 0x29, 0xc7, 0x91, 0xec, 0xe9, 0x37, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x87, 0x50, 0x5a,
	0xad, 0x01, 0x00, 0x00,
}