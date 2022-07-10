// Code generated by protoc-gen-go. DO NOT EDIT.
// source: names.proto

package proto_out

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

type Names struct {
	Name                 []string `protobuf:"bytes,1,rep,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Names) Reset()         { *m = Names{} }
func (m *Names) String() string { return proto.CompactTextString(m) }
func (*Names) ProtoMessage()    {}
func (*Names) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4268625867c617c, []int{0}
}

func (m *Names) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Names.Unmarshal(m, b)
}
func (m *Names) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Names.Marshal(b, m, deterministic)
}
func (m *Names) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Names.Merge(m, src)
}
func (m *Names) XXX_Size() int {
	return xxx_messageInfo_Names.Size(m)
}
func (m *Names) XXX_DiscardUnknown() {
	xxx_messageInfo_Names.DiscardUnknown(m)
}

var xxx_messageInfo_Names proto.InternalMessageInfo

func (m *Names) GetName() []string {
	if m != nil {
		return m.Name
	}
	return nil
}

func init() {
	proto.RegisterType((*Names)(nil), "Names")
}

func init() {
	proto.RegisterFile("names.proto", fileDescriptor_f4268625867c617c)
}

var fileDescriptor_f4268625867c617c = []byte{
	// 58 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x4b, 0xcc, 0x4d,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe6, 0x62, 0xf5, 0x03, 0x71, 0x85, 0x84,
	0xb8, 0x58, 0x40, 0xe2, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0x9c, 0x41, 0x60, 0x36, 0x20, 0x00, 0x00,
	0xff, 0xff, 0xa4, 0x0f, 0x98, 0xd2, 0x2a, 0x00, 0x00, 0x00,
}