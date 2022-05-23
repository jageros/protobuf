// Code generated by protoc-gen-go. DO NOT EDIT.
// source: multi/multi1.proto

package multitest

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

type Multi1 struct {
	Multi2               *Multi2         `protobuf:"bytes,1,req,name=multi2" json:"multi2,omitempty"`
	Color                *Multi2_Color   `protobuf:"varint,2,opt,name=color,enum=multitest.Multi2_Color" json:"color,omitempty"`
	HatType              *Multi3_HatType `protobuf:"varint,3,opt,name=hat_type,json=hatType,enum=multitest.Multi3_HatType" json:"hat_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Multi1) Reset()         { *m = Multi1{} }
func (m *Multi1) String() string { return proto.CompactTextString(m) }
func (*Multi1) ProtoMessage()    {}
func (*Multi1) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0bffc140cd1b1d9, []int{0}
}

func (m *Multi1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Multi1.Unmarshal(m, b)
}
func (m *Multi1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Multi1.Marshal(b, m, deterministic)
}
func (m *Multi1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Multi1.Merge(m, src)
}
func (m *Multi1) XXX_Size() int {
	return xxx_messageInfo_Multi1.Size(m)
}
func (m *Multi1) XXX_DiscardUnknown() {
	xxx_messageInfo_Multi1.DiscardUnknown(m)
}

var xxx_messageInfo_Multi1 proto.InternalMessageInfo

func (m *Multi1) GetMulti2() *Multi2 {
	if m != nil {
		return m.Multi2
	}
	return nil
}

func (m *Multi1) GetColor() Multi2_Color {
	if m != nil && m.Color != nil {
		return *m.Color
	}
	return Multi2_BLUE
}

func (m *Multi1) GetHatType() Multi3_HatType {
	if m != nil && m.HatType != nil {
		return *m.HatType
	}
	return Multi3_FEDORA
}

func init() {
	proto.RegisterType((*Multi1)(nil), "multitest.Multi1")
}

func init() { proto.RegisterFile("multi/multi1.proto", fileDescriptor_e0bffc140cd1b1d9) }

var fileDescriptor_e0bffc140cd1b1d9 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0x2d, 0xcd, 0x29,
	0xc9, 0xd4, 0x07, 0x93, 0x86, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x9c, 0x60, 0x5e, 0x49,
	0x6a, 0x71, 0x89, 0x14, 0xb2, 0xb4, 0x11, 0x44, 0x1a, 0x45, 0xcc, 0x18, 0x22, 0xa6, 0x34, 0x83,
	0x91, 0x8b, 0xcd, 0x17, 0x6c, 0x86, 0x90, 0x26, 0x17, 0x1b, 0x44, 0xb9, 0x04, 0xa3, 0x02, 0x93,
	0x06, 0xb7, 0x91, 0xa0, 0x1e, 0xdc, 0x38, 0x3d, 0xb0, 0x12, 0xa3, 0x20, 0xa8, 0x02, 0x21, 0x5d,
	0x2e, 0xd6, 0xe4, 0xfc, 0x9c, 0xfc, 0x22, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x3e, 0x23, 0x71, 0x0c,
	0x95, 0x7a, 0xce, 0x20, 0xe9, 0x20, 0x88, 0x2a, 0x21, 0x13, 0x2e, 0x8e, 0x8c, 0xc4, 0x92, 0xf8,
	0x92, 0xca, 0x82, 0x54, 0x09, 0x66, 0xb0, 0x0e, 0x49, 0x74, 0x1d, 0xc6, 0x7a, 0x1e, 0x89, 0x25,
	0x21, 0x95, 0x05, 0xa9, 0x41, 0xec, 0x19, 0x10, 0x86, 0x93, 0x73, 0x94, 0x63, 0x7a, 0x66, 0x49,
	0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x7a, 0x7e, 0x4e, 0x62, 0x5e, 0xba, 0x3e, 0xd8,
	0xd5, 0x49, 0xa5, 0x69, 0x10, 0x46, 0xb2, 0x6e, 0x7a, 0x6a, 0x9e, 0x6e, 0x7a, 0xbe, 0x3e, 0xc8,
	0xa0, 0x94, 0xc4, 0x92, 0x44, 0x88, 0xe7, 0xac, 0xe1, 0x86, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x60, 0x7d, 0xfc, 0x9f, 0x27, 0x01, 0x00, 0x00,
}
