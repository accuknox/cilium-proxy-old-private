// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/ratelimit/v2/rls.proto

package v2

import (
	context "context"
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	ratelimit "github.com/cilium/proxy/go/envoy/api/v2/ratelimit"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/lyft/protoc-gen-validate/validate"
	grpc "google.golang.org/grpc"
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

type RateLimitResponse_Code int32

const (
	// The response code is not known.
	RateLimitResponse_UNKNOWN RateLimitResponse_Code = 0
	// The response code to notify that the number of requests are under limit.
	RateLimitResponse_OK RateLimitResponse_Code = 1
	// The response code to notify that the number of requests are over limit.
	RateLimitResponse_OVER_LIMIT RateLimitResponse_Code = 2
)

var RateLimitResponse_Code_name = map[int32]string{
	0: "UNKNOWN",
	1: "OK",
	2: "OVER_LIMIT",
}

var RateLimitResponse_Code_value = map[string]int32{
	"UNKNOWN":    0,
	"OK":         1,
	"OVER_LIMIT": 2,
}

func (x RateLimitResponse_Code) String() string {
	return proto.EnumName(RateLimitResponse_Code_name, int32(x))
}

func (RateLimitResponse_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1de95711edb19ee8, []int{1, 0}
}

type RateLimitResponse_RateLimit_Unit int32

const (
	// The time unit is not known.
	RateLimitResponse_RateLimit_UNKNOWN RateLimitResponse_RateLimit_Unit = 0
	// The time unit representing a second.
	RateLimitResponse_RateLimit_SECOND RateLimitResponse_RateLimit_Unit = 1
	// The time unit representing a minute.
	RateLimitResponse_RateLimit_MINUTE RateLimitResponse_RateLimit_Unit = 2
	// The time unit representing an hour.
	RateLimitResponse_RateLimit_HOUR RateLimitResponse_RateLimit_Unit = 3
	// The time unit representing a day.
	RateLimitResponse_RateLimit_DAY RateLimitResponse_RateLimit_Unit = 4
)

var RateLimitResponse_RateLimit_Unit_name = map[int32]string{
	0: "UNKNOWN",
	1: "SECOND",
	2: "MINUTE",
	3: "HOUR",
	4: "DAY",
}

var RateLimitResponse_RateLimit_Unit_value = map[string]int32{
	"UNKNOWN": 0,
	"SECOND":  1,
	"MINUTE":  2,
	"HOUR":    3,
	"DAY":     4,
}

func (x RateLimitResponse_RateLimit_Unit) String() string {
	return proto.EnumName(RateLimitResponse_RateLimit_Unit_name, int32(x))
}

func (RateLimitResponse_RateLimit_Unit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1de95711edb19ee8, []int{1, 0, 0}
}

// Main message for a rate limit request. The rate limit service is designed to be fully generic
// in the sense that it can operate on arbitrary hierarchical key/value pairs. The loaded
// configuration will parse the request and find the most specific limit to apply. In addition,
// a RateLimitRequest can contain multiple "descriptors" to limit on. When multiple descriptors
// are provided, the server will limit on *ALL* of them and return an OVER_LIMIT response if any
// of them are over limit. This enables more complex application level rate limiting scenarios
// if desired.
type RateLimitRequest struct {
	// All rate limit requests must specify a domain. This enables the configuration to be per
	// application without fear of overlap. E.g., "envoy".
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// All rate limit requests must specify at least one RateLimitDescriptor. Each descriptor is
	// processed by the service (see below). If any of the descriptors are over limit, the entire
	// request is considered to be over limit.
	Descriptors []*ratelimit.RateLimitDescriptor `protobuf:"bytes,2,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	// Rate limit requests can optionally specify the number of hits a request adds to the matched
	// limit. If the value is not set in the message, a request increases the matched limit by 1.
	HitsAddend           uint32   `protobuf:"varint,3,opt,name=hits_addend,json=hitsAddend,proto3" json:"hits_addend,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimitRequest) Reset()         { *m = RateLimitRequest{} }
func (m *RateLimitRequest) String() string { return proto.CompactTextString(m) }
func (*RateLimitRequest) ProtoMessage()    {}
func (*RateLimitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1de95711edb19ee8, []int{0}
}

func (m *RateLimitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitRequest.Unmarshal(m, b)
}
func (m *RateLimitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitRequest.Marshal(b, m, deterministic)
}
func (m *RateLimitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitRequest.Merge(m, src)
}
func (m *RateLimitRequest) XXX_Size() int {
	return xxx_messageInfo_RateLimitRequest.Size(m)
}
func (m *RateLimitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitRequest proto.InternalMessageInfo

func (m *RateLimitRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimitRequest) GetDescriptors() []*ratelimit.RateLimitDescriptor {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

func (m *RateLimitRequest) GetHitsAddend() uint32 {
	if m != nil {
		return m.HitsAddend
	}
	return 0
}

// A response from a ShouldRateLimit call.
type RateLimitResponse struct {
	// The overall response code which takes into account all of the descriptors that were passed
	// in the RateLimitRequest message.
	OverallCode RateLimitResponse_Code `protobuf:"varint,1,opt,name=overall_code,json=overallCode,proto3,enum=envoy.service.ratelimit.v2.RateLimitResponse_Code" json:"overall_code,omitempty"`
	// A list of DescriptorStatus messages which matches the length of the descriptor list passed
	// in the RateLimitRequest. This can be used by the caller to determine which individual
	// descriptors failed and/or what the currently configured limits are for all of them.
	Statuses []*RateLimitResponse_DescriptorStatus `protobuf:"bytes,2,rep,name=statuses,proto3" json:"statuses,omitempty"`
	// A list of headers to add to the response
	Headers              []*core.HeaderValue `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RateLimitResponse) Reset()         { *m = RateLimitResponse{} }
func (m *RateLimitResponse) String() string { return proto.CompactTextString(m) }
func (*RateLimitResponse) ProtoMessage()    {}
func (*RateLimitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1de95711edb19ee8, []int{1}
}

func (m *RateLimitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitResponse.Unmarshal(m, b)
}
func (m *RateLimitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitResponse.Marshal(b, m, deterministic)
}
func (m *RateLimitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitResponse.Merge(m, src)
}
func (m *RateLimitResponse) XXX_Size() int {
	return xxx_messageInfo_RateLimitResponse.Size(m)
}
func (m *RateLimitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitResponse proto.InternalMessageInfo

func (m *RateLimitResponse) GetOverallCode() RateLimitResponse_Code {
	if m != nil {
		return m.OverallCode
	}
	return RateLimitResponse_UNKNOWN
}

func (m *RateLimitResponse) GetStatuses() []*RateLimitResponse_DescriptorStatus {
	if m != nil {
		return m.Statuses
	}
	return nil
}

func (m *RateLimitResponse) GetHeaders() []*core.HeaderValue {
	if m != nil {
		return m.Headers
	}
	return nil
}

// Defines an actual rate limit in terms of requests per unit of time and the unit itself.
type RateLimitResponse_RateLimit struct {
	// The number of requests per unit of time.
	RequestsPerUnit uint32 `protobuf:"varint,1,opt,name=requests_per_unit,json=requestsPerUnit,proto3" json:"requests_per_unit,omitempty"`
	// The unit of time.
	Unit                 RateLimitResponse_RateLimit_Unit `protobuf:"varint,2,opt,name=unit,proto3,enum=envoy.service.ratelimit.v2.RateLimitResponse_RateLimit_Unit" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *RateLimitResponse_RateLimit) Reset()         { *m = RateLimitResponse_RateLimit{} }
func (m *RateLimitResponse_RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimitResponse_RateLimit) ProtoMessage()    {}
func (*RateLimitResponse_RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_1de95711edb19ee8, []int{1, 0}
}

func (m *RateLimitResponse_RateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitResponse_RateLimit.Unmarshal(m, b)
}
func (m *RateLimitResponse_RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitResponse_RateLimit.Marshal(b, m, deterministic)
}
func (m *RateLimitResponse_RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitResponse_RateLimit.Merge(m, src)
}
func (m *RateLimitResponse_RateLimit) XXX_Size() int {
	return xxx_messageInfo_RateLimitResponse_RateLimit.Size(m)
}
func (m *RateLimitResponse_RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitResponse_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitResponse_RateLimit proto.InternalMessageInfo

func (m *RateLimitResponse_RateLimit) GetRequestsPerUnit() uint32 {
	if m != nil {
		return m.RequestsPerUnit
	}
	return 0
}

func (m *RateLimitResponse_RateLimit) GetUnit() RateLimitResponse_RateLimit_Unit {
	if m != nil {
		return m.Unit
	}
	return RateLimitResponse_RateLimit_UNKNOWN
}

type RateLimitResponse_DescriptorStatus struct {
	// The response code for an individual descriptor.
	Code RateLimitResponse_Code `protobuf:"varint,1,opt,name=code,proto3,enum=envoy.service.ratelimit.v2.RateLimitResponse_Code" json:"code,omitempty"`
	// The current limit as configured by the server. Useful for debugging, etc.
	CurrentLimit *RateLimitResponse_RateLimit `protobuf:"bytes,2,opt,name=current_limit,json=currentLimit,proto3" json:"current_limit,omitempty"`
	// The limit remaining in the current time unit.
	LimitRemaining       uint32   `protobuf:"varint,3,opt,name=limit_remaining,json=limitRemaining,proto3" json:"limit_remaining,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimitResponse_DescriptorStatus) Reset()         { *m = RateLimitResponse_DescriptorStatus{} }
func (m *RateLimitResponse_DescriptorStatus) String() string { return proto.CompactTextString(m) }
func (*RateLimitResponse_DescriptorStatus) ProtoMessage()    {}
func (*RateLimitResponse_DescriptorStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_1de95711edb19ee8, []int{1, 1}
}

func (m *RateLimitResponse_DescriptorStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitResponse_DescriptorStatus.Unmarshal(m, b)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitResponse_DescriptorStatus.Marshal(b, m, deterministic)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitResponse_DescriptorStatus.Merge(m, src)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_Size() int {
	return xxx_messageInfo_RateLimitResponse_DescriptorStatus.Size(m)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitResponse_DescriptorStatus.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitResponse_DescriptorStatus proto.InternalMessageInfo

func (m *RateLimitResponse_DescriptorStatus) GetCode() RateLimitResponse_Code {
	if m != nil {
		return m.Code
	}
	return RateLimitResponse_UNKNOWN
}

func (m *RateLimitResponse_DescriptorStatus) GetCurrentLimit() *RateLimitResponse_RateLimit {
	if m != nil {
		return m.CurrentLimit
	}
	return nil
}

func (m *RateLimitResponse_DescriptorStatus) GetLimitRemaining() uint32 {
	if m != nil {
		return m.LimitRemaining
	}
	return 0
}

func init() {
	proto.RegisterEnum("envoy.service.ratelimit.v2.RateLimitResponse_Code", RateLimitResponse_Code_name, RateLimitResponse_Code_value)
	proto.RegisterEnum("envoy.service.ratelimit.v2.RateLimitResponse_RateLimit_Unit", RateLimitResponse_RateLimit_Unit_name, RateLimitResponse_RateLimit_Unit_value)
	proto.RegisterType((*RateLimitRequest)(nil), "envoy.service.ratelimit.v2.RateLimitRequest")
	proto.RegisterType((*RateLimitResponse)(nil), "envoy.service.ratelimit.v2.RateLimitResponse")
	proto.RegisterType((*RateLimitResponse_RateLimit)(nil), "envoy.service.ratelimit.v2.RateLimitResponse.RateLimit")
	proto.RegisterType((*RateLimitResponse_DescriptorStatus)(nil), "envoy.service.ratelimit.v2.RateLimitResponse.DescriptorStatus")
}

func init() {
	proto.RegisterFile("envoy/service/ratelimit/v2/rls.proto", fileDescriptor_1de95711edb19ee8)
}

var fileDescriptor_1de95711edb19ee8 = []byte{
	// 582 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x86, 0xeb, 0x24, 0x5f, 0xda, 0x4e, 0xfa, 0xe3, 0xee, 0xc1, 0x47, 0x64, 0x21, 0xa8, 0x22,
	0x04, 0x11, 0x05, 0x5b, 0x32, 0x07, 0x20, 0x84, 0x2a, 0xf5, 0x0f, 0xb5, 0x6a, 0x9b, 0x44, 0x9b,
	0xa6, 0x88, 0x0a, 0xc9, 0xda, 0xc6, 0x23, 0xba, 0x92, 0xeb, 0x35, 0xbb, 0x6b, 0x8b, 0x9e, 0x73,
	0x15, 0xdc, 0x11, 0x17, 0xc1, 0x01, 0x77, 0x82, 0xbc, 0x76, 0x9c, 0x14, 0x04, 0x22, 0x70, 0x66,
	0xcf, 0xcc, 0xfb, 0x78, 0xde, 0x99, 0xf5, 0xc2, 0x03, 0x8c, 0x33, 0x71, 0xe3, 0x29, 0x94, 0x19,
	0x1f, 0xa3, 0x27, 0x99, 0xc6, 0x88, 0x5f, 0x73, 0xed, 0x65, 0xbe, 0x27, 0x23, 0xe5, 0x26, 0x52,
	0x68, 0x41, 0x1c, 0x53, 0xe5, 0x96, 0x55, 0x6e, 0x55, 0xe5, 0x66, 0xbe, 0x73, 0xb7, 0x20, 0xb0,
	0x84, 0xe7, 0x9a, 0xb1, 0x90, 0xe8, 0x5d, 0x32, 0x85, 0x85, 0xd2, 0x79, 0x78, 0x2b, 0x3b, 0xc5,
	0x4f, 0x11, 0x45, 0xdd, 0x9d, 0x8c, 0x45, 0x3c, 0x64, 0x1a, 0xbd, 0xc9, 0x43, 0x91, 0xe8, 0x7c,
	0xb6, 0xc0, 0xa6, 0x4c, 0xe3, 0x49, 0x5e, 0x4c, 0xf1, 0x43, 0x8a, 0x4a, 0x93, 0xff, 0xa1, 0x19,
	0x8a, 0x6b, 0xc6, 0xe3, 0xb6, 0xb5, 0x69, 0x75, 0x97, 0x69, 0xf9, 0x46, 0x4e, 0xa1, 0x15, 0xa2,
	0x1a, 0x4b, 0x9e, 0x68, 0x21, 0x55, 0xbb, 0xb6, 0x59, 0xef, 0xb6, 0xfc, 0x2d, 0xb7, 0xe8, 0x9e,
	0x25, 0xdc, 0xcd, 0xfc, 0x99, 0xe6, 0x2b, 0xec, 0x7e, 0xa5, 0xa1, 0xb3, 0x7a, 0x72, 0x1f, 0x5a,
	0x57, 0x5c, 0xab, 0x80, 0x85, 0x21, 0xc6, 0x61, 0xbb, 0xbe, 0x69, 0x75, 0x57, 0x29, 0xe4, 0xa1,
	0x1d, 0x13, 0xe9, 0x7c, 0xfd, 0x0f, 0x36, 0x66, 0x9a, 0x53, 0x89, 0x88, 0x15, 0x92, 0x11, 0xac,
	0x88, 0x0c, 0x25, 0x8b, 0xa2, 0x60, 0x2c, 0x42, 0x34, 0x3d, 0xae, 0xf9, 0xbe, 0xfb, 0xeb, 0x21,
	0xba, 0x3f, 0x41, 0xdc, 0x3d, 0x11, 0x22, 0x6d, 0x95, 0x9c, 0xfc, 0x85, 0x5c, 0xc0, 0x92, 0xd2,
	0x4c, 0xa7, 0x0a, 0x27, 0xce, 0xb6, 0xe7, 0x43, 0x4e, 0x6d, 0x0e, 0x0d, 0x87, 0x56, 0x3c, 0xf2,
	0x02, 0x16, 0xaf, 0x90, 0x85, 0x28, 0x55, 0xbb, 0x6e, 0xd0, 0xf7, 0x6e, 0x0f, 0x2d, 0x5f, 0xab,
	0x7b, 0x68, 0x2a, 0xce, 0x59, 0x94, 0x22, 0x9d, 0x94, 0x3b, 0x5f, 0x2c, 0x58, 0xae, 0x3e, 0x45,
	0x1e, 0xc3, 0x86, 0x2c, 0x76, 0xa4, 0x82, 0x04, 0x65, 0x90, 0xc6, 0x5c, 0x1b, 0xff, 0xab, 0x74,
	0x7d, 0x92, 0x18, 0xa0, 0x1c, 0xc5, 0x5c, 0x93, 0x01, 0x34, 0x4c, 0xba, 0x66, 0xc6, 0xf3, 0x6a,
	0x3e, 0x2f, 0x55, 0xc4, 0xcd, 0x59, 0xd4, 0x90, 0x3a, 0xdb, 0xd0, 0x30, 0xe4, 0x16, 0x2c, 0x8e,
	0x7a, 0xc7, 0xbd, 0xfe, 0x9b, 0x9e, 0xbd, 0x40, 0x00, 0x9a, 0xc3, 0x83, 0xbd, 0x7e, 0x6f, 0xdf,
	0xb6, 0xf2, 0xe7, 0xd3, 0xa3, 0xde, 0xe8, 0xec, 0xc0, 0xae, 0x91, 0x25, 0x68, 0x1c, 0xf6, 0x47,
	0xd4, 0xae, 0x93, 0x45, 0xa8, 0xef, 0xef, 0xbc, 0xb5, 0x1b, 0xce, 0x37, 0x0b, 0xec, 0x1f, 0x87,
	0x44, 0x5e, 0x43, 0xe3, 0x1f, 0xb7, 0x68, 0xf4, 0xe4, 0x1d, 0xac, 0x8e, 0x53, 0x29, 0x31, 0xd6,
	0x81, 0x11, 0x18, 0xdf, 0x2d, 0xff, 0xf9, 0x5f, 0xfa, 0xa6, 0x2b, 0x25, 0xad, 0x18, 0xfc, 0x23,
	0x58, 0x37, 0xaa, 0x40, 0x62, 0xfe, 0x27, 0xf0, 0xf8, 0x7d, 0x79, 0x5c, 0xd7, 0xa2, 0x42, 0x5f,
	0x46, 0x3b, 0x5b, 0xd0, 0x30, 0xa7, 0xe9, 0xd6, 0x8c, 0x9a, 0x50, 0xeb, 0x1f, 0xdb, 0x16, 0x59,
	0x03, 0xe8, 0x9f, 0x1f, 0xd0, 0xe0, 0xe4, 0xe8, 0xf4, 0xe8, 0xcc, 0xae, 0xf9, 0x9f, 0x66, 0x7f,
	0xbe, 0x61, 0xd1, 0x21, 0x49, 0x60, 0x7d, 0x78, 0x25, 0xd2, 0x28, 0x9c, 0xae, 0xfd, 0xc9, 0x1f,
	0x9a, 0x30, 0x07, 0xc0, 0x79, 0x3a, 0x97, 0xe5, 0xce, 0xc2, 0xee, 0x4b, 0xe8, 0x72, 0x51, 0x88,
	0x12, 0x29, 0x3e, 0xde, 0xfc, 0x46, 0xbf, 0xbb, 0x44, 0x23, 0x35, 0xc8, 0x6f, 0x8e, 0x81, 0x75,
	0x51, 0xcb, 0xfc, 0xcb, 0xa6, 0xb9, 0x46, 0x9e, 0x7d, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x84, 0xf4,
	0x57, 0x38, 0xe9, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RateLimitServiceClient is the client API for RateLimitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RateLimitServiceClient interface {
	// Determine whether rate limiting should take place.
	ShouldRateLimit(ctx context.Context, in *RateLimitRequest, opts ...grpc.CallOption) (*RateLimitResponse, error)
}

type rateLimitServiceClient struct {
	cc *grpc.ClientConn
}

func NewRateLimitServiceClient(cc *grpc.ClientConn) RateLimitServiceClient {
	return &rateLimitServiceClient{cc}
}

func (c *rateLimitServiceClient) ShouldRateLimit(ctx context.Context, in *RateLimitRequest, opts ...grpc.CallOption) (*RateLimitResponse, error) {
	out := new(RateLimitResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.ratelimit.v2.RateLimitService/ShouldRateLimit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateLimitServiceServer is the server API for RateLimitService service.
type RateLimitServiceServer interface {
	// Determine whether rate limiting should take place.
	ShouldRateLimit(context.Context, *RateLimitRequest) (*RateLimitResponse, error)
}

func RegisterRateLimitServiceServer(s *grpc.Server, srv RateLimitServiceServer) {
	s.RegisterService(&_RateLimitService_serviceDesc, srv)
}

func _RateLimitService_ShouldRateLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimitServiceServer).ShouldRateLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.ratelimit.v2.RateLimitService/ShouldRateLimit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimitServiceServer).ShouldRateLimit(ctx, req.(*RateLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RateLimitService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.ratelimit.v2.RateLimitService",
	HandlerType: (*RateLimitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShouldRateLimit",
			Handler:    _RateLimitService_ShouldRateLimit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "envoy/service/ratelimit/v2/rls.proto",
}
