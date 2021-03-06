// Code generated by protoc-gen-go. DO NOT EDIT.
// source: person.proto

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

type Person struct {
	Name                 *string  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Age                  *int32   `protobuf:"varint,2,req,name=age" json:"age,omitempty"`
	ActivityTs           *uint32  `protobuf:"fixed32,3,opt,name=activity_ts,json=activityTs" json:"activity_ts,omitempty"`
	Address              *Point   `protobuf:"bytes,4,opt,name=address" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Person) GetAge() int32 {
	if m != nil && m.Age != nil {
		return *m.Age
	}
	return 0
}

func (m *Person) GetActivityTs() uint32 {
	if m != nil && m.ActivityTs != nil {
		return *m.ActivityTs
	}
	return 0
}

func (m *Person) GetAddress() *Point {
	if m != nil {
		return m.Address
	}
	return nil
}

func init() {
	proto.RegisterType((*Person)(nil), "Person")
}

func init() {
	proto.RegisterFile("person.proto", fileDescriptor_4c9e10cf24b1156d)
}

var fileDescriptor_4c9e10cf24b1156d = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x48, 0x2d, 0x2a,
	0xce, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0xe2, 0x29, 0xc8, 0xcf, 0xcc, 0x2b, 0x29,
	0x86, 0xf0, 0x94, 0x0a, 0xb9, 0xd8, 0x02, 0xc0, 0xb2, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9,
	0xa9, 0x12, 0x8c, 0x0a, 0x4c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x00, 0x17, 0x73, 0x62, 0x7a,
	0xaa, 0x04, 0x93, 0x02, 0x93, 0x06, 0x6b, 0x10, 0x88, 0x29, 0x24, 0xcf, 0xc5, 0x9d, 0x98, 0x5c,
	0x92, 0x59, 0x96, 0x59, 0x52, 0x19, 0x5f, 0x52, 0x2c, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x1e, 0xc4,
	0x05, 0x13, 0x0a, 0x29, 0x16, 0x52, 0xe0, 0x62, 0x4f, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0x96,
	0x60, 0x51, 0x60, 0xd4, 0xe0, 0x36, 0x62, 0xd3, 0x0b, 0x00, 0x59, 0x18, 0x04, 0x13, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xaf, 0x63, 0x9b, 0xf8, 0x8f, 0x00, 0x00, 0x00,
}
