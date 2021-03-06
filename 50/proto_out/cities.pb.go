// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cities.proto

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

type Cities struct {
	City                 []string `protobuf:"bytes,1,rep,name=city" json:"city,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cities) Reset()         { *m = Cities{} }
func (m *Cities) String() string { return proto.CompactTextString(m) }
func (*Cities) ProtoMessage()    {}
func (*Cities) Descriptor() ([]byte, []int) {
	return fileDescriptor_403b24a7741af34a, []int{0}
}

func (m *Cities) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cities.Unmarshal(m, b)
}
func (m *Cities) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cities.Marshal(b, m, deterministic)
}
func (m *Cities) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cities.Merge(m, src)
}
func (m *Cities) XXX_Size() int {
	return xxx_messageInfo_Cities.Size(m)
}
func (m *Cities) XXX_DiscardUnknown() {
	xxx_messageInfo_Cities.DiscardUnknown(m)
}

var xxx_messageInfo_Cities proto.InternalMessageInfo

func (m *Cities) GetCity() []string {
	if m != nil {
		return m.City
	}
	return nil
}

func init() {
	proto.RegisterType((*Cities)(nil), "Cities")
}

func init() {
	proto.RegisterFile("cities.proto", fileDescriptor_403b24a7741af34a)
}

var fileDescriptor_403b24a7741af34a = []byte{
	// 62 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0x2c, 0xc9,
	0x4c, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe1, 0x62, 0x73, 0x06, 0xf3, 0x85,
	0x84, 0xb8, 0x58, 0x92, 0x33, 0x4b, 0x2a, 0x25, 0x18, 0x15, 0x98, 0x35, 0x38, 0x83, 0xc0, 0x6c,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x24, 0x5d, 0x2c, 0x32, 0x2c, 0x00, 0x00, 0x00,
}
