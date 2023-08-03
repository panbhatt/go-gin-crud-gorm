// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth_service.proto

package pb

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

type VerifyEmailRequest struct {
	VeriificationCode    string   `protobuf:"bytes,1,opt,name=veriificationCode,proto3" json:"veriificationCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyEmailRequest) Reset()         { *m = VerifyEmailRequest{} }
func (m *VerifyEmailRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyEmailRequest) ProtoMessage()    {}
func (*VerifyEmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f39bb026ca10b68, []int{0}
}

func (m *VerifyEmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyEmailRequest.Unmarshal(m, b)
}
func (m *VerifyEmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyEmailRequest.Marshal(b, m, deterministic)
}
func (m *VerifyEmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyEmailRequest.Merge(m, src)
}
func (m *VerifyEmailRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyEmailRequest.Size(m)
}
func (m *VerifyEmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyEmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyEmailRequest proto.InternalMessageInfo

func (m *VerifyEmailRequest) GetVeriificationCode() string {
	if m != nil {
		return m.VeriificationCode
	}
	return ""
}

func init() {
	proto.RegisterType((*VerifyEmailRequest)(nil), "pb.VerifyEmailRequest")
}

func init() {
	proto.RegisterFile("auth_service.proto", fileDescriptor_0f39bb026ca10b68)
}

var fileDescriptor_0f39bb026ca10b68 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0xb1, 0x4a, 0x03, 0x41,
	0x10, 0x86, 0x3d, 0x0b, 0xc1, 0x4d, 0xe5, 0x8a, 0x22, 0x57, 0x49, 0x1a, 0x15, 0xbc, 0x3b, 0xd0,
	0xc6, 0xc2, 0xc6, 0x88, 0x88, 0xed, 0x85, 0x58, 0xd8, 0x1c, 0xbb, 0x9b, 0xc9, 0xde, 0x80, 0xb7,
	0x3b, 0xce, 0xce, 0x06, 0x7c, 0x00, 0xdf, 0x5b, 0x4c, 0x8a, 0x08, 0x47, 0xba, 0x99, 0xf9, 0xe6,
	0x87, 0xef, 0x57, 0xda, 0x64, 0xe9, 0xbb, 0x04, 0xbc, 0x46, 0x07, 0x35, 0x71, 0x94, 0xa8, 0x0f,
	0xc9, 0x96, 0x67, 0x4c, 0xae, 0x4b, 0xe8, 0x43, 0xa6, 0x2e, 0x27, 0xe0, 0x2d, 0x2a, 0xd5, 0x6e,
	0x9e, 0xce, 0x94, 0x7e, 0x07, 0xc6, 0xd5, 0xf7, 0xcb, 0x60, 0xf0, 0xb3, 0x85, 0xaf, 0x0c, 0x49,
	0xf4, 0xad, 0x3a, 0x59, 0x03, 0x23, 0xae, 0xd0, 0x19, 0xc1, 0x18, 0x9e, 0xe3, 0x12, 0x2e, 0x8a,
	0xcb, 0xe2, 0xfa, 0xb8, 0x1d, 0x83, 0xbb, 0x9f, 0x42, 0x4d, 0x9e, 0xb2, 0xf4, 0xf3, 0xad, 0x80,
	0x7e, 0x50, 0x6a, 0x8e, 0x3e, 0x2c, 0x68, 0x91, 0x80, 0xf5, 0x69, 0x4d, 0xb6, 0xde, 0xed, 0x6f,
	0x81, 0xb2, 0x94, 0x9b, 0xe3, 0x2b, 0x04, 0x60, 0x74, 0x2d, 0x24, 0x8a, 0x21, 0xc1, 0xf4, 0x40,
	0x3f, 0xaa, 0xc9, 0x3f, 0x1b, 0x7d, 0xfe, 0xf7, 0x35, 0xd6, 0xdb, 0x93, 0x9e, 0xdd, 0x7c, 0x5c,
	0x79, 0x94, 0x3e, 0xdb, 0xda, 0xc5, 0xa1, 0x21, 0x13, 0x6c, 0x6f, 0x44, 0x1a, 0x1f, 0x2b, 0x8f,
	0xa1, 0x72, 0x9c, 0x97, 0x95, 0x8f, 0x3c, 0x34, 0x64, 0xed, 0xd1, 0xa6, 0xfd, 0xfd, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xa9, 0xd2, 0xe2, 0x93, 0x3a, 0x01, 0x00, 0x00,
}