// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package go_micro_srv_user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_58f123e63f2dcdb3, []int{0}
}
func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (dst *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(dst, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UserResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_58f123e63f2dcdb3, []int{1}
}
func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (dst *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(dst, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "go.micro.srv.user.UserRequest")
	proto.RegisterType((*UserResponse)(nil), "go.micro.srv.user.UserResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_58f123e63f2dcdb3) }

var fileDescriptor_user_58f123e63f2dcdb3 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4c, 0xcf, 0xd7, 0xcb, 0xcd, 0x4c, 0x2e, 0xca,
	0xd7, 0x2b, 0x2e, 0x2a, 0xd3, 0x03, 0x49, 0x28, 0xb9, 0x72, 0x71, 0x87, 0x16, 0xa7, 0x16, 0x05,
	0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x49, 0x71, 0x71, 0x80, 0x84, 0xf3, 0x12, 0x73, 0x53,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x90, 0x5c, 0x41, 0x62, 0x71, 0x71, 0x79,
	0x7e, 0x51, 0x8a, 0x04, 0x13, 0x44, 0x0e, 0xc6, 0x57, 0x92, 0xe3, 0xe2, 0x81, 0x18, 0x53, 0x5c,
	0x90, 0x9f, 0x57, 0x9c, 0x2a, 0xc4, 0xc7, 0xc5, 0x94, 0x99, 0x02, 0x35, 0x81, 0x29, 0x33, 0xc5,
	0x28, 0x96, 0x8b, 0xcb, 0xb9, 0x28, 0x35, 0xb1, 0x24, 0x15, 0xa4, 0x4a, 0xc8, 0x1f, 0x85, 0x27,
	0xa7, 0x87, 0xe1, 0x2c, 0x3d, 0x24, 0x37, 0x49, 0xc9, 0xe3, 0x94, 0x87, 0x58, 0xa6, 0xc4, 0x90,
	0xc4, 0x06, 0xf6, 0x9f, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x98, 0x5b, 0xc2, 0xbc, 0xed, 0x00,
	0x00, 0x00,
}
