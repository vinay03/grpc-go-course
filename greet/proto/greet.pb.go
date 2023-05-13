// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greet/proto/greet.proto

package proto

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

type GreetRequest struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetRequest) Reset()         { *m = GreetRequest{} }
func (m *GreetRequest) String() string { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()    {}
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cc3401bebadabd3, []int{0}
}

func (m *GreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetRequest.Unmarshal(m, b)
}
func (m *GreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetRequest.Marshal(b, m, deterministic)
}
func (m *GreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetRequest.Merge(m, src)
}
func (m *GreetRequest) XXX_Size() int {
	return xxx_messageInfo_GreetRequest.Size(m)
}
func (m *GreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetRequest proto.InternalMessageInfo

func (m *GreetRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

type GreetResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetResponse) Reset()         { *m = GreetResponse{} }
func (m *GreetResponse) String() string { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()    {}
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cc3401bebadabd3, []int{1}
}

func (m *GreetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetResponse.Unmarshal(m, b)
}
func (m *GreetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetResponse.Marshal(b, m, deterministic)
}
func (m *GreetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetResponse.Merge(m, src)
}
func (m *GreetResponse) XXX_Size() int {
	return xxx_messageInfo_GreetResponse.Size(m)
}
func (m *GreetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetResponse proto.InternalMessageInfo

func (m *GreetResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*GreetRequest)(nil), "greet.GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "greet.GreetResponse")
}

func init() { proto.RegisterFile("greet/proto/greet.proto", fileDescriptor_4cc3401bebadabd3) }

var fileDescriptor_4cc3401bebadabd3 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x59, 0xa1, 0x85, 0x1d, 0x54, 0x30, 0x8a, 0x8a, 0x20, 0x48, 0x2f, 0xf6, 0xb2, 0x9b,
	0xa5, 0xbd, 0x78, 0xd1, 0x83, 0x28, 0x5e, 0xd4, 0x43, 0x15, 0x04, 0x2f, 0x92, 0xae, 0x63, 0x1a,
	0xe8, 0x26, 0xeb, 0x24, 0xbb, 0xb0, 0xbf, 0xcd, 0x3f, 0x27, 0x1d, 0x73, 0xa8, 0xc7, 0x9c, 0xf2,
	0xde, 0xc7, 0xbc, 0x37, 0x09, 0x81, 0x13, 0x4d, 0x88, 0x41, 0xb6, 0xe4, 0x82, 0x93, 0xac, 0x4b,
	0xd6, 0x62, 0xc4, 0x66, 0x52, 0xc0, 0xee, 0xc3, 0x46, 0x2c, 0xf0, 0xbb, 0x43, 0x1f, 0xc4, 0x39,
	0xc0, 0x97, 0x21, 0x1f, 0x3e, 0xac, 0x6a, 0xf0, 0x34, 0xbb, 0xc8, 0xa6, 0xf9, 0x22, 0x67, 0xf2,
	0xac, 0x1a, 0x9c, 0x5c, 0xc2, 0x5e, 0x1c, 0xf7, 0xad, 0xb3, 0x1e, 0xc5, 0x31, 0x8c, 0x09, 0x7d,
	0xb7, 0x0e, 0x71, 0x36, 0xba, 0xd9, 0xcf, 0x4e, 0x2c, 0x7e, 0x41, 0xea, 0x4d, 0x8d, 0x62, 0x06,
	0x23, 0xf6, 0xe2, 0xb0, 0xfc, 0xbb, 0xc6, 0xf6, 0xda, 0xb3, 0xa3, 0xff, 0x30, 0x96, 0x5f, 0xc3,
	0x3e, 0x83, 0x27, 0x65, 0x87, 0x57, 0xd3, 0xa0, 0x4f, 0x08, 0x57, 0x99, 0xb8, 0x82, 0xfc, 0xd1,
	0x59, 0x9d, 0xba, 0x76, 0x9a, 0x89, 0x9b, 0xf8, 0xcc, 0xfb, 0x1e, 0x69, 0x70, 0x16, 0x93, 0xd2,
	0xd5, 0x26, 0x7f, 0xc0, 0xf0, 0xcd, 0x84, 0xd5, 0x1d, 0xaa, 0xcf, 0xb5, 0x49, 0xea, 0xb8, 0x95,
	0xef, 0x85, 0x36, 0x61, 0xd5, 0x2d, 0xcb, 0xda, 0x35, 0xb2, 0x37, 0x56, 0x0d, 0xd5, 0x5c, 0x6a,
	0x6a, 0xeb, 0x42, 0xbb, 0xa2, 0x76, 0x1d, 0x79, 0x94, 0x5b, 0x3f, 0xbb, 0x1c, 0xf3, 0x31, 0xff,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0x89, 0xd2, 0x1a, 0xe3, 0xef, 0x01, 0x00, 0x00,
}