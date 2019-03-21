// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/buffer/v2/buffer.proto

package v2

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/lyft/protoc-gen-validate/validate"
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

type Buffer struct {
	// The maximum request size that the filter will buffer before the connection
	// manager will stop buffering and return a 413 response.
	MaxRequestBytes      *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=max_request_bytes,json=maxRequestBytes,proto3" json:"max_request_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Buffer) Reset()         { *m = Buffer{} }
func (m *Buffer) String() string { return proto.CompactTextString(m) }
func (*Buffer) ProtoMessage()    {}
func (*Buffer) Descriptor() ([]byte, []int) {
	return fileDescriptor_e782fc75ce4c789f, []int{0}
}

func (m *Buffer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Buffer.Unmarshal(m, b)
}
func (m *Buffer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Buffer.Marshal(b, m, deterministic)
}
func (m *Buffer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Buffer.Merge(m, src)
}
func (m *Buffer) XXX_Size() int {
	return xxx_messageInfo_Buffer.Size(m)
}
func (m *Buffer) XXX_DiscardUnknown() {
	xxx_messageInfo_Buffer.DiscardUnknown(m)
}

var xxx_messageInfo_Buffer proto.InternalMessageInfo

func (m *Buffer) GetMaxRequestBytes() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxRequestBytes
	}
	return nil
}

type BufferPerRoute struct {
	// Types that are valid to be assigned to Override:
	//	*BufferPerRoute_Disabled
	//	*BufferPerRoute_Buffer
	Override             isBufferPerRoute_Override `protobuf_oneof:"override"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *BufferPerRoute) Reset()         { *m = BufferPerRoute{} }
func (m *BufferPerRoute) String() string { return proto.CompactTextString(m) }
func (*BufferPerRoute) ProtoMessage()    {}
func (*BufferPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_e782fc75ce4c789f, []int{1}
}

func (m *BufferPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BufferPerRoute.Unmarshal(m, b)
}
func (m *BufferPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BufferPerRoute.Marshal(b, m, deterministic)
}
func (m *BufferPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BufferPerRoute.Merge(m, src)
}
func (m *BufferPerRoute) XXX_Size() int {
	return xxx_messageInfo_BufferPerRoute.Size(m)
}
func (m *BufferPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_BufferPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_BufferPerRoute proto.InternalMessageInfo

type isBufferPerRoute_Override interface {
	isBufferPerRoute_Override()
}

type BufferPerRoute_Disabled struct {
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3,oneof"`
}

type BufferPerRoute_Buffer struct {
	Buffer *Buffer `protobuf:"bytes,2,opt,name=buffer,proto3,oneof"`
}

func (*BufferPerRoute_Disabled) isBufferPerRoute_Override() {}

func (*BufferPerRoute_Buffer) isBufferPerRoute_Override() {}

func (m *BufferPerRoute) GetOverride() isBufferPerRoute_Override {
	if m != nil {
		return m.Override
	}
	return nil
}

func (m *BufferPerRoute) GetDisabled() bool {
	if x, ok := m.GetOverride().(*BufferPerRoute_Disabled); ok {
		return x.Disabled
	}
	return false
}

func (m *BufferPerRoute) GetBuffer() *Buffer {
	if x, ok := m.GetOverride().(*BufferPerRoute_Buffer); ok {
		return x.Buffer
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BufferPerRoute) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BufferPerRoute_Disabled)(nil),
		(*BufferPerRoute_Buffer)(nil),
	}
}

func init() {
	proto.RegisterType((*Buffer)(nil), "envoy.config.filter.http.buffer.v2.Buffer")
	proto.RegisterType((*BufferPerRoute)(nil), "envoy.config.filter.http.buffer.v2.BufferPerRoute")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/buffer/v2/buffer.proto", fileDescriptor_e782fc75ce4c789f)
}

var fileDescriptor_e782fc75ce4c789f = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xdd, 0x4a, 0xc3, 0x30,
	0x1c, 0xc5, 0xd7, 0xee, 0x83, 0x9a, 0x81, 0x6e, 0x45, 0x70, 0x88, 0xc8, 0xd8, 0x8d, 0xb2, 0x8b,
	0x44, 0xba, 0x37, 0x08, 0x08, 0xea, 0xd5, 0xa8, 0xe8, 0x85, 0x37, 0x23, 0x5d, 0xff, 0xad, 0x91,
	0xae, 0xa9, 0x69, 0x5a, 0xb7, 0x57, 0xf0, 0x21, 0x7c, 0x0e, 0xf1, 0x6a, 0xaf, 0xb3, 0xb7, 0x90,
	0x26, 0x99, 0x37, 0x5e, 0xe8, 0xdd, 0x21, 0x27, 0xe7, 0xfc, 0x92, 0x83, 0x08, 0xe4, 0xb5, 0xd8,
	0x90, 0xa5, 0xc8, 0x13, 0x9e, 0x92, 0x84, 0x67, 0x0a, 0x24, 0x79, 0x56, 0xaa, 0x20, 0x51, 0x95,
	0x24, 0x20, 0x49, 0x1d, 0x58, 0x85, 0x0b, 0x29, 0x94, 0xf0, 0x27, 0x3a, 0x80, 0x4d, 0x00, 0x9b,
	0x00, 0x6e, 0x02, 0xd8, 0x5e, 0xab, 0x83, 0xd3, 0xf3, 0x54, 0x88, 0x34, 0x03, 0xa2, 0x13, 0x51,
	0x95, 0x90, 0xb8, 0x92, 0x4c, 0x71, 0x91, 0x9b, 0x8e, 0xdf, 0xfe, 0x9b, 0x64, 0x45, 0x01, 0xb2,
	0xb4, 0xfe, 0x49, 0xcd, 0x32, 0x1e, 0x33, 0x05, 0x64, 0x2f, 0xac, 0x71, 0x9c, 0x8a, 0x54, 0x68,
	0x49, 0x1a, 0x65, 0x4e, 0x27, 0x4b, 0xd4, 0xa3, 0x9a, 0xed, 0xdf, 0xa3, 0xe1, 0x8a, 0xad, 0x17,
	0x12, 0x5e, 0x2b, 0x28, 0xd5, 0x22, 0xda, 0x28, 0x28, 0x47, 0xce, 0xd8, 0xb9, 0xec, 0x07, 0x67,
	0xd8, 0x40, 0xf1, 0x1e, 0x8a, 0x1f, 0x6e, 0x73, 0x35, 0x0b, 0x1e, 0x59, 0x56, 0x01, 0x3d, 0xf8,
	0xda, 0x6d, 0xdb, 0x9d, 0xa9, 0x3b, 0x6e, 0x85, 0x47, 0x2b, 0xb6, 0x0e, 0x4d, 0x01, 0x6d, 0xf2,
	0x77, 0x1d, 0xcf, 0x1d, 0xb4, 0x27, 0x1f, 0x0e, 0x3a, 0x34, 0x94, 0x39, 0xc8, 0x50, 0x54, 0x0a,
	0xfc, 0x0b, 0xe4, 0xc5, 0xbc, 0x64, 0x51, 0x06, 0xb1, 0x86, 0x78, 0xb6, 0xe6, 0xc5, 0xf5, 0x9c,
	0x9b, 0x56, 0xf8, 0x63, 0xfa, 0x73, 0xd4, 0x33, 0xe3, 0x8c, 0x5c, 0xfd, 0x96, 0x29, 0xfe, 0x7b,
	0x44, 0x6c, 0x60, 0x14, 0x35, 0x95, 0xdd, 0x77, 0xc7, 0x1d, 0x34, 0x9d, 0xb6, 0x87, 0x0e, 0x91,
	0x27, 0x6a, 0x90, 0x92, 0xc7, 0xe0, 0x77, 0x3f, 0x77, 0xdb, 0xb6, 0x43, 0xaf, 0xd1, 0x15, 0x17,
	0xa6, 0xb8, 0x90, 0x62, 0xbd, 0xf9, 0x07, 0x83, 0xf6, 0xed, 0x8f, 0x9a, 0x49, 0xe6, 0xce, 0x93,
	0x5b, 0x07, 0x51, 0x4f, 0xef, 0x33, 0xfb, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x88, 0x6e, 0xa6, 0x5d,
	0x19, 0x02, 0x00, 0x00,
}
