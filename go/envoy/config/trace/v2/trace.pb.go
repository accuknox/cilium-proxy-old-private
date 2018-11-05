// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/trace/v2/trace.proto

package v2

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The tracing configuration specifies global
// settings for the HTTP tracer used by Envoy. The configuration is defined by
// the :ref:`Bootstrap <envoy_api_msg_config.bootstrap.v2.Bootstrap>` :ref:`tracing
// <envoy_api_field_config.bootstrap.v2.Bootstrap.tracing>` field. Envoy may support other tracers
// in the future, but right now the HTTP tracer is the only one supported.
type Tracing struct {
	// Provides configuration for the HTTP tracer.
	Http                 *Tracing_Http `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Tracing) Reset()         { *m = Tracing{} }
func (m *Tracing) String() string { return proto.CompactTextString(m) }
func (*Tracing) ProtoMessage()    {}
func (*Tracing) Descriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{0}
}

func (m *Tracing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tracing.Unmarshal(m, b)
}
func (m *Tracing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tracing.Marshal(b, m, deterministic)
}
func (m *Tracing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tracing.Merge(m, src)
}
func (m *Tracing) XXX_Size() int {
	return xxx_messageInfo_Tracing.Size(m)
}
func (m *Tracing) XXX_DiscardUnknown() {
	xxx_messageInfo_Tracing.DiscardUnknown(m)
}

var xxx_messageInfo_Tracing proto.InternalMessageInfo

func (m *Tracing) GetHttp() *Tracing_Http {
	if m != nil {
		return m.Http
	}
	return nil
}

type Tracing_Http struct {
	// The name of the HTTP trace driver to instantiate. The name must match a
	// supported HTTP trace driver. *envoy.lightstep*, *envoy.zipkin*, and
	// *envoy.dynamic.ot* are built-in trace drivers.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Trace driver specific configuration which depends on the driver being
	// instantiated. See the :ref:`LightstepConfig
	// <envoy_api_msg_config.trace.v2.LightstepConfig>`, :ref:`ZipkinConfig
	// <envoy_api_msg_config.trace.v2.ZipkinConfig>`, and :ref:`DynamicOtConfig
	// <envoy_api_msg_config.trace.v2.DynamicOtConfig>` trace drivers for examples.
	Config               *_struct.Struct `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Tracing_Http) Reset()         { *m = Tracing_Http{} }
func (m *Tracing_Http) String() string { return proto.CompactTextString(m) }
func (*Tracing_Http) ProtoMessage()    {}
func (*Tracing_Http) Descriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{0, 0}
}

func (m *Tracing_Http) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tracing_Http.Unmarshal(m, b)
}
func (m *Tracing_Http) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tracing_Http.Marshal(b, m, deterministic)
}
func (m *Tracing_Http) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tracing_Http.Merge(m, src)
}
func (m *Tracing_Http) XXX_Size() int {
	return xxx_messageInfo_Tracing_Http.Size(m)
}
func (m *Tracing_Http) XXX_DiscardUnknown() {
	xxx_messageInfo_Tracing_Http.DiscardUnknown(m)
}

var xxx_messageInfo_Tracing_Http proto.InternalMessageInfo

func (m *Tracing_Http) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tracing_Http) GetConfig() *_struct.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

// Configuration for the LightStep tracer.
type LightstepConfig struct {
	// The cluster manager cluster that hosts the LightStep collectors.
	CollectorCluster string `protobuf:"bytes,1,opt,name=collector_cluster,json=collectorCluster,proto3" json:"collector_cluster,omitempty"`
	// File containing the access token to the `LightStep
	// <http://lightstep.com/>`_ API.
	AccessTokenFile      string   `protobuf:"bytes,2,opt,name=access_token_file,json=accessTokenFile,proto3" json:"access_token_file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LightstepConfig) Reset()         { *m = LightstepConfig{} }
func (m *LightstepConfig) String() string { return proto.CompactTextString(m) }
func (*LightstepConfig) ProtoMessage()    {}
func (*LightstepConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{1}
}

func (m *LightstepConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LightstepConfig.Unmarshal(m, b)
}
func (m *LightstepConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LightstepConfig.Marshal(b, m, deterministic)
}
func (m *LightstepConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LightstepConfig.Merge(m, src)
}
func (m *LightstepConfig) XXX_Size() int {
	return xxx_messageInfo_LightstepConfig.Size(m)
}
func (m *LightstepConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LightstepConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LightstepConfig proto.InternalMessageInfo

func (m *LightstepConfig) GetCollectorCluster() string {
	if m != nil {
		return m.CollectorCluster
	}
	return ""
}

func (m *LightstepConfig) GetAccessTokenFile() string {
	if m != nil {
		return m.AccessTokenFile
	}
	return ""
}

type ZipkinConfig struct {
	// The cluster manager cluster that hosts the Zipkin collectors. Note that the
	// Zipkin cluster must be defined in the :ref:`Bootstrap static cluster
	// resources <envoy_api_field_config.bootstrap.v2.Bootstrap.StaticResources.clusters>`.
	CollectorCluster string `protobuf:"bytes,1,opt,name=collector_cluster,json=collectorCluster,proto3" json:"collector_cluster,omitempty"`
	// The API endpoint of the Zipkin service where the spans will be sent. When
	// using a standard Zipkin installation, the API endpoint is typically
	// /api/v1/spans, which is the default value.
	CollectorEndpoint string `protobuf:"bytes,2,opt,name=collector_endpoint,json=collectorEndpoint,proto3" json:"collector_endpoint,omitempty"`
	// Determines whether a 128bit trace id will be used when creating a new
	// trace instance. The default value is false, which will result in a 64 bit trace id being used.
	TraceId_128Bit bool `protobuf:"varint,3,opt,name=trace_id_128bit,json=traceId128bit,proto3" json:"trace_id_128bit,omitempty"`
	// Determines whether client and server spans will shared the same span id.
	// The default value is true.
	SharedSpanContext    *wrappers.BoolValue `protobuf:"bytes,4,opt,name=shared_span_context,json=sharedSpanContext,proto3" json:"shared_span_context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ZipkinConfig) Reset()         { *m = ZipkinConfig{} }
func (m *ZipkinConfig) String() string { return proto.CompactTextString(m) }
func (*ZipkinConfig) ProtoMessage()    {}
func (*ZipkinConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{2}
}

func (m *ZipkinConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZipkinConfig.Unmarshal(m, b)
}
func (m *ZipkinConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZipkinConfig.Marshal(b, m, deterministic)
}
func (m *ZipkinConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZipkinConfig.Merge(m, src)
}
func (m *ZipkinConfig) XXX_Size() int {
	return xxx_messageInfo_ZipkinConfig.Size(m)
}
func (m *ZipkinConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ZipkinConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ZipkinConfig proto.InternalMessageInfo

func (m *ZipkinConfig) GetCollectorCluster() string {
	if m != nil {
		return m.CollectorCluster
	}
	return ""
}

func (m *ZipkinConfig) GetCollectorEndpoint() string {
	if m != nil {
		return m.CollectorEndpoint
	}
	return ""
}

func (m *ZipkinConfig) GetTraceId_128Bit() bool {
	if m != nil {
		return m.TraceId_128Bit
	}
	return false
}

func (m *ZipkinConfig) GetSharedSpanContext() *wrappers.BoolValue {
	if m != nil {
		return m.SharedSpanContext
	}
	return nil
}

// DynamicOtConfig is used to dynamically load a tracer from a shared library
// that implements the `OpenTracing dynamic loading API
// <https://github.com/opentracing/opentracing-cpp>`_.
type DynamicOtConfig struct {
	// Dynamic library implementing the `OpenTracing API
	// <https://github.com/opentracing/opentracing-cpp>`_.
	Library string `protobuf:"bytes,1,opt,name=library,proto3" json:"library,omitempty"`
	// The configuration to use when creating a tracer from the given dynamic
	// library.
	Config               *_struct.Struct `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DynamicOtConfig) Reset()         { *m = DynamicOtConfig{} }
func (m *DynamicOtConfig) String() string { return proto.CompactTextString(m) }
func (*DynamicOtConfig) ProtoMessage()    {}
func (*DynamicOtConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{3}
}

func (m *DynamicOtConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynamicOtConfig.Unmarshal(m, b)
}
func (m *DynamicOtConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynamicOtConfig.Marshal(b, m, deterministic)
}
func (m *DynamicOtConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynamicOtConfig.Merge(m, src)
}
func (m *DynamicOtConfig) XXX_Size() int {
	return xxx_messageInfo_DynamicOtConfig.Size(m)
}
func (m *DynamicOtConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DynamicOtConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DynamicOtConfig proto.InternalMessageInfo

func (m *DynamicOtConfig) GetLibrary() string {
	if m != nil {
		return m.Library
	}
	return ""
}

func (m *DynamicOtConfig) GetConfig() *_struct.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

// Configuration structure.
type TraceServiceConfig struct {
	// The upstream gRPC cluster that hosts the metrics service.
	GrpcService          *core.GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TraceServiceConfig) Reset()         { *m = TraceServiceConfig{} }
func (m *TraceServiceConfig) String() string { return proto.CompactTextString(m) }
func (*TraceServiceConfig) ProtoMessage()    {}
func (*TraceServiceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{4}
}

func (m *TraceServiceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceServiceConfig.Unmarshal(m, b)
}
func (m *TraceServiceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceServiceConfig.Marshal(b, m, deterministic)
}
func (m *TraceServiceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceServiceConfig.Merge(m, src)
}
func (m *TraceServiceConfig) XXX_Size() int {
	return xxx_messageInfo_TraceServiceConfig.Size(m)
}
func (m *TraceServiceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceServiceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TraceServiceConfig proto.InternalMessageInfo

func (m *TraceServiceConfig) GetGrpcService() *core.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func init() {
	proto.RegisterType((*Tracing)(nil), "envoy.config.trace.v2.Tracing")
	proto.RegisterType((*Tracing_Http)(nil), "envoy.config.trace.v2.Tracing.Http")
	proto.RegisterType((*LightstepConfig)(nil), "envoy.config.trace.v2.LightstepConfig")
	proto.RegisterType((*ZipkinConfig)(nil), "envoy.config.trace.v2.ZipkinConfig")
	proto.RegisterType((*DynamicOtConfig)(nil), "envoy.config.trace.v2.DynamicOtConfig")
	proto.RegisterType((*TraceServiceConfig)(nil), "envoy.config.trace.v2.TraceServiceConfig")
}

func init() { proto.RegisterFile("envoy/config/trace/v2/trace.proto", fileDescriptor_0785d24fc8ab55c7) }

var fileDescriptor_0785d24fc8ab55c7 = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcf, 0x6e, 0x13, 0x3f,
	0x10, 0xd6, 0xa6, 0xfb, 0x6b, 0x7f, 0x75, 0x8b, 0xd2, 0x18, 0xa1, 0x46, 0x11, 0x54, 0x21, 0x45,
	0xa8, 0x27, 0xaf, 0x08, 0x02, 0x7a, 0x4e, 0xf9, 0x2f, 0x10, 0xd2, 0xa6, 0xea, 0xa1, 0x97, 0x95,
	0xe3, 0x4c, 0x36, 0x56, 0x5d, 0xdb, 0xf2, 0x4e, 0x16, 0x72, 0xe3, 0xcc, 0x23, 0xf0, 0x28, 0x9c,
	0x78, 0x1d, 0xee, 0x3c, 0x00, 0x5a, 0xdb, 0xa1, 0x85, 0x70, 0x41, 0xdc, 0xec, 0xf9, 0xbe, 0x6f,
	0xe6, 0x9b, 0xf1, 0x98, 0xdc, 0x05, 0x5d, 0x9b, 0x65, 0x26, 0x8c, 0x9e, 0xc9, 0x32, 0x43, 0xc7,
	0x05, 0x64, 0xf5, 0x30, 0x1c, 0x98, 0x75, 0x06, 0x0d, 0xbd, 0xe5, 0x29, 0x2c, 0x50, 0x58, 0x40,
	0xea, 0x61, 0xef, 0x5e, 0x50, 0x72, 0x2b, 0x1b, 0x81, 0x30, 0x0e, 0xb2, 0xd2, 0x59, 0x51, 0x54,
	0xe0, 0x6a, 0xb9, 0x12, 0xf7, 0x6e, 0x97, 0xc6, 0x94, 0x0a, 0x32, 0x7f, 0x9b, 0x2c, 0x66, 0x59,
	0x85, 0x6e, 0x21, 0x30, 0xa2, 0x07, 0xbf, 0xa3, 0xef, 0x1d, 0xb7, 0x16, 0x5c, 0x15, 0xf1, 0xfd,
	0x9a, 0x2b, 0x39, 0xe5, 0x08, 0xd9, 0xea, 0x10, 0x80, 0xc1, 0xe7, 0x84, 0x6c, 0x9d, 0x3a, 0x2e,
	0xa4, 0x2e, 0xe9, 0x13, 0x92, 0xce, 0x11, 0x6d, 0x37, 0xe9, 0x27, 0x47, 0x3b, 0xc3, 0x43, 0xf6,
	0x47, 0xbb, 0x2c, 0xb2, 0xd9, 0x4b, 0x44, 0x9b, 0x7b, 0x41, 0xef, 0x8c, 0xa4, 0xcd, 0x8d, 0xde,
	0x21, 0xa9, 0xe6, 0x97, 0xe0, 0x13, 0x6c, 0x8f, 0xb6, 0xbf, 0x7c, 0xfb, 0xba, 0x91, 0xba, 0x56,
	0x3f, 0xc9, 0x7d, 0x98, 0x66, 0x64, 0x33, 0x24, 0xeb, 0xb6, 0x7c, 0x85, 0x7d, 0x16, 0x5c, 0xb3,
	0x95, 0x6b, 0x36, 0xf6, 0x3d, 0xe5, 0x91, 0x36, 0xf8, 0x98, 0x90, 0xf6, 0x1b, 0x59, 0xce, 0xb1,
	0x42, 0xb0, 0x27, 0x3e, 0x46, 0x1f, 0x93, 0x8e, 0x30, 0x4a, 0x81, 0x40, 0xe3, 0x0a, 0xa1, 0x16,
	0x15, 0x82, 0x5b, 0x2f, 0xb8, 0xf7, 0x93, 0x73, 0x12, 0x28, 0xf4, 0x11, 0xe9, 0x70, 0x21, 0xa0,
	0xaa, 0x0a, 0x34, 0x17, 0xa0, 0x8b, 0x99, 0x54, 0xe0, 0x7d, 0xfc, 0xa2, 0x6b, 0x07, 0xce, 0x69,
	0x43, 0x79, 0x2e, 0x15, 0x0c, 0xbe, 0x27, 0x64, 0xf7, 0x5c, 0xda, 0x0b, 0xa9, 0xff, 0xb1, 0xfe,
	0x31, 0xa1, 0x57, 0x3a, 0xd0, 0x53, 0x6b, 0xa4, 0xc6, 0x75, 0x03, 0x57, 0xc9, 0x9f, 0x45, 0x0e,
	0xbd, 0x4f, 0xda, 0x7e, 0xf8, 0x85, 0x9c, 0x16, 0x0f, 0x86, 0xc7, 0x13, 0x89, 0xdd, 0x8d, 0x7e,
	0x72, 0xf4, 0x7f, 0x7e, 0xc3, 0x87, 0x5f, 0x4d, 0x43, 0x90, 0xbe, 0x26, 0x37, 0xab, 0x39, 0x77,
	0x30, 0x2d, 0x2a, 0xcb, 0x75, 0x21, 0x8c, 0x46, 0xf8, 0x80, 0xdd, 0xd4, 0xcf, 0xba, 0xb7, 0x36,
	0xeb, 0x91, 0x31, 0xea, 0x8c, 0xab, 0x05, 0xe4, 0x9d, 0x20, 0x1b, 0x5b, 0xde, 0x34, 0xd9, 0x88,
	0x06, 0x25, 0x69, 0x3f, 0x5d, 0x6a, 0x7e, 0x29, 0xc5, 0x3b, 0x8c, 0x8d, 0x1f, 0x92, 0x2d, 0x25,
	0x27, 0x8e, 0xbb, 0xe5, 0x7a, 0xbb, 0x2b, 0xe4, 0xef, 0x9f, 0x58, 0x10, 0xda, 0x2c, 0x14, 0x8c,
	0xc3, 0xb2, 0xc7, 0x5a, 0x6f, 0xc9, 0xee, 0xf5, 0x2f, 0x10, 0x37, 0xf2, 0x20, 0x6e, 0x24, 0xb7,
	0xb2, 0x59, 0xc4, 0xe6, 0xa7, 0xb0, 0x17, 0xce, 0x8a, 0xa8, 0x1d, 0x91, 0xc6, 0xd0, 0x7f, 0x9f,
	0x92, 0xd6, 0x5e, 0x92, 0xef, 0x94, 0xd7, 0x80, 0xf4, 0xbc, 0x55, 0x0f, 0x27, 0x9b, 0xde, 0xc3,
	0xc3, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x79, 0x45, 0x14, 0x07, 0xaa, 0x03, 0x00, 0x00,
}