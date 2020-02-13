// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hi.proto

package msg

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

type HiReq struct {
	Route                string   `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"`
	Words                string   `protobuf:"bytes,2,opt,name=words,proto3" json:"words,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HiReq) Reset()         { *m = HiReq{} }
func (m *HiReq) String() string { return proto.CompactTextString(m) }
func (*HiReq) ProtoMessage()    {}
func (*HiReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d092a8920edeec73, []int{0}
}

func (m *HiReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HiReq.Unmarshal(m, b)
}
func (m *HiReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HiReq.Marshal(b, m, deterministic)
}
func (m *HiReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HiReq.Merge(m, src)
}
func (m *HiReq) XXX_Size() int {
	return xxx_messageInfo_HiReq.Size(m)
}
func (m *HiReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HiReq.DiscardUnknown(m)
}

var xxx_messageInfo_HiReq proto.InternalMessageInfo

func (m *HiReq) GetRoute() string {
	if m != nil {
		return m.Route
	}
	return ""
}

func (m *HiReq) GetWords() string {
	if m != nil {
		return m.Words
	}
	return ""
}

type HiResp struct {
	Words                string   `protobuf:"bytes,1,opt,name=words,proto3" json:"words,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HiResp) Reset()         { *m = HiResp{} }
func (m *HiResp) String() string { return proto.CompactTextString(m) }
func (*HiResp) ProtoMessage()    {}
func (*HiResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d092a8920edeec73, []int{1}
}

func (m *HiResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HiResp.Unmarshal(m, b)
}
func (m *HiResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HiResp.Marshal(b, m, deterministic)
}
func (m *HiResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HiResp.Merge(m, src)
}
func (m *HiResp) XXX_Size() int {
	return xxx_messageInfo_HiResp.Size(m)
}
func (m *HiResp) XXX_DiscardUnknown() {
	xxx_messageInfo_HiResp.DiscardUnknown(m)
}

var xxx_messageInfo_HiResp proto.InternalMessageInfo

func (m *HiResp) GetWords() string {
	if m != nil {
		return m.Words
	}
	return ""
}

func init() {
	proto.RegisterType((*HiReq)(nil), "msg.HiReq")
	proto.RegisterType((*HiResp)(nil), "msg.HiResp")
}

func init() { proto.RegisterFile("hi.proto", fileDescriptor_d092a8920edeec73) }

var fileDescriptor_d092a8920edeec73 = []byte{
	// 96 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc8, 0xc8, 0xd4, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2d, 0x4e, 0x57, 0x32, 0xe6, 0x62, 0xf5, 0xc8, 0x0c,
	0x4a, 0x2d, 0x14, 0x12, 0xe1, 0x62, 0x2d, 0xca, 0x2f, 0x2d, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0x82, 0x70, 0x40, 0xa2, 0xe5, 0xf9, 0x45, 0x29, 0xc5, 0x12, 0x4c, 0x10, 0x51, 0x30,
	0x47, 0x49, 0x8e, 0x8b, 0x0d, 0xa4, 0xa9, 0xb8, 0x00, 0x21, 0xcf, 0x88, 0x24, 0x9f, 0xc4, 0x06,
	0xb6, 0xc0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x58, 0x72, 0x63, 0x44, 0x6c, 0x00, 0x00, 0x00,
}