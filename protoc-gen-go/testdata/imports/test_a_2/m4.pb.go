// Code generated by protoc-gen-go. DO NOT EDIT.
// source: imports/test_a_2/m4.proto

package test_a_2

import (
	fmt "fmt"
	proto "github.com/jageros/protobuf/proto"
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

type M4 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M4) Reset()         { *m = M4{} }
func (m *M4) String() string { return proto.CompactTextString(m) }
func (*M4) ProtoMessage()    {}
func (*M4) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdd24f82f6c5a786, []int{0}
}

func (m *M4) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M4.Unmarshal(m, b)
}
func (m *M4) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M4.Marshal(b, m, deterministic)
}
func (m *M4) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M4.Merge(m, src)
}
func (m *M4) XXX_Size() int {
	return xxx_messageInfo_M4.Size(m)
}
func (m *M4) XXX_DiscardUnknown() {
	xxx_messageInfo_M4.DiscardUnknown(m)
}

var xxx_messageInfo_M4 proto.InternalMessageInfo

func init() {
	proto.RegisterType((*M4)(nil), "test.a.M4")
}

func init() { proto.RegisterFile("imports/test_a_2/m4.proto", fileDescriptor_fdd24f82f6c5a786) }

var fileDescriptor_fdd24f82f6c5a786 = []byte{
	// 114 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0xcc, 0x2d, 0xc8,
	0x2f, 0x2a, 0x29, 0xd6, 0x2f, 0x49, 0x2d, 0x2e, 0x89, 0x4f, 0x8c, 0x37, 0xd2, 0xcf, 0x35, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x09, 0xe9, 0x25, 0x2a, 0xb1, 0x70, 0x31, 0xf9,
	0x9a, 0x38, 0xb9, 0x44, 0x39, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea,
	0xa7, 0xe7, 0xe7, 0x24, 0xe6, 0xa5, 0xeb, 0x83, 0x15, 0x26, 0x95, 0xa6, 0x41, 0x18, 0xc9, 0xba,
	0xe9, 0xa9, 0x79, 0xba, 0xe9, 0xf9, 0x60, 0xb3, 0x52, 0x12, 0x4b, 0x12, 0xf5, 0xd1, 0x0d, 0x4f,
	0x62, 0x03, 0x2b, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x58, 0xcb, 0x10, 0xc8, 0x77, 0x00,
	0x00, 0x00,
}
