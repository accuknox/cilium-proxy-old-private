// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/core/config_source.proto

package core

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

// APIs may be fetched via either REST or gRPC.
type ApiConfigSource_ApiType int32

const (
	// REST-JSON legacy corresponds to the v1 API.
	ApiConfigSource_REST_LEGACY ApiConfigSource_ApiType = 0
	// REST-JSON v2 API. The `canonical JSON encoding
	// <https://developers.google.com/protocol-buffers/docs/proto3#json>`_ for
	// the v2 protos is used.
	ApiConfigSource_REST ApiConfigSource_ApiType = 1
	// gRPC v2 API.
	ApiConfigSource_GRPC ApiConfigSource_ApiType = 2
)

var ApiConfigSource_ApiType_name = map[int32]string{
	0: "REST_LEGACY",
	1: "REST",
	2: "GRPC",
}

var ApiConfigSource_ApiType_value = map[string]int32{
	"REST_LEGACY": 0,
	"REST":        1,
	"GRPC":        2,
}

func (x ApiConfigSource_ApiType) String() string {
	return proto.EnumName(ApiConfigSource_ApiType_name, int32(x))
}

func (ApiConfigSource_ApiType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1ffcc55cf4c30535, []int{0, 0}
}

// API configuration source. This identifies the API type and cluster that Envoy
// will use to fetch an xDS API.
type ApiConfigSource struct {
	ApiType ApiConfigSource_ApiType `protobuf:"varint,1,opt,name=api_type,json=apiType,proto3,enum=envoy.api.v2.core.ApiConfigSource_ApiType" json:"api_type,omitempty"`
	// Cluster names should be used only with REST_LEGACY/REST. If > 1
	// cluster is defined, clusters will be cycled through if any kind of failure
	// occurs.
	//
	// .. note::
	//
	//  The cluster with name ``cluster_name`` must be statically defined and its
	//  type must not be ``EDS``.
	ClusterNames []string `protobuf:"bytes,2,rep,name=cluster_names,json=clusterNames,proto3" json:"cluster_names,omitempty"`
	// Multiple gRPC services be provided for GRPC. If > 1 cluster is defined,
	// services will be cycled through if any kind of failure occurs.
	GrpcServices []*GrpcService `protobuf:"bytes,4,rep,name=grpc_services,json=grpcServices,proto3" json:"grpc_services,omitempty"`
	// For REST APIs, the delay between successive polls.
	RefreshDelay *duration.Duration `protobuf:"bytes,3,opt,name=refresh_delay,json=refreshDelay,proto3" json:"refresh_delay,omitempty"`
	// For REST APIs, the request timeout. If not set, a default value of 1s will be used.
	RequestTimeout       *duration.Duration `protobuf:"bytes,5,opt,name=request_timeout,json=requestTimeout,proto3" json:"request_timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ApiConfigSource) Reset()         { *m = ApiConfigSource{} }
func (m *ApiConfigSource) String() string { return proto.CompactTextString(m) }
func (*ApiConfigSource) ProtoMessage()    {}
func (*ApiConfigSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ffcc55cf4c30535, []int{0}
}

func (m *ApiConfigSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiConfigSource.Unmarshal(m, b)
}
func (m *ApiConfigSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiConfigSource.Marshal(b, m, deterministic)
}
func (m *ApiConfigSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiConfigSource.Merge(m, src)
}
func (m *ApiConfigSource) XXX_Size() int {
	return xxx_messageInfo_ApiConfigSource.Size(m)
}
func (m *ApiConfigSource) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiConfigSource.DiscardUnknown(m)
}

var xxx_messageInfo_ApiConfigSource proto.InternalMessageInfo

func (m *ApiConfigSource) GetApiType() ApiConfigSource_ApiType {
	if m != nil {
		return m.ApiType
	}
	return ApiConfigSource_REST_LEGACY
}

func (m *ApiConfigSource) GetClusterNames() []string {
	if m != nil {
		return m.ClusterNames
	}
	return nil
}

func (m *ApiConfigSource) GetGrpcServices() []*GrpcService {
	if m != nil {
		return m.GrpcServices
	}
	return nil
}

func (m *ApiConfigSource) GetRefreshDelay() *duration.Duration {
	if m != nil {
		return m.RefreshDelay
	}
	return nil
}

func (m *ApiConfigSource) GetRequestTimeout() *duration.Duration {
	if m != nil {
		return m.RequestTimeout
	}
	return nil
}

// Aggregated Discovery Service (ADS) options. This is currently empty, but when
// set in :ref:`ConfigSource <envoy_api_msg_core.ConfigSource>` can be used to
// specify that ADS is to be used.
type AggregatedConfigSource struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AggregatedConfigSource) Reset()         { *m = AggregatedConfigSource{} }
func (m *AggregatedConfigSource) String() string { return proto.CompactTextString(m) }
func (*AggregatedConfigSource) ProtoMessage()    {}
func (*AggregatedConfigSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ffcc55cf4c30535, []int{1}
}

func (m *AggregatedConfigSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggregatedConfigSource.Unmarshal(m, b)
}
func (m *AggregatedConfigSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggregatedConfigSource.Marshal(b, m, deterministic)
}
func (m *AggregatedConfigSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregatedConfigSource.Merge(m, src)
}
func (m *AggregatedConfigSource) XXX_Size() int {
	return xxx_messageInfo_AggregatedConfigSource.Size(m)
}
func (m *AggregatedConfigSource) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregatedConfigSource.DiscardUnknown(m)
}

var xxx_messageInfo_AggregatedConfigSource proto.InternalMessageInfo

// Configuration for :ref:`listeners <config_listeners>`, :ref:`clusters
// <config_cluster_manager>`, :ref:`routes
// <envoy_api_msg_RouteConfiguration>`, :ref:`endpoints
// <arch_overview_service_discovery>` etc. may either be sourced from the
// filesystem or from an xDS API source. Filesystem configs are watched with
// inotify for updates.
type ConfigSource struct {
	// Types that are valid to be assigned to ConfigSourceSpecifier:
	//	*ConfigSource_Path
	//	*ConfigSource_ApiConfigSource
	//	*ConfigSource_Ads
	ConfigSourceSpecifier isConfigSource_ConfigSourceSpecifier `protobuf_oneof:"config_source_specifier"`
	XXX_NoUnkeyedLiteral  struct{}                             `json:"-"`
	XXX_unrecognized      []byte                               `json:"-"`
	XXX_sizecache         int32                                `json:"-"`
}

func (m *ConfigSource) Reset()         { *m = ConfigSource{} }
func (m *ConfigSource) String() string { return proto.CompactTextString(m) }
func (*ConfigSource) ProtoMessage()    {}
func (*ConfigSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ffcc55cf4c30535, []int{2}
}

func (m *ConfigSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigSource.Unmarshal(m, b)
}
func (m *ConfigSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigSource.Marshal(b, m, deterministic)
}
func (m *ConfigSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigSource.Merge(m, src)
}
func (m *ConfigSource) XXX_Size() int {
	return xxx_messageInfo_ConfigSource.Size(m)
}
func (m *ConfigSource) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigSource.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigSource proto.InternalMessageInfo

type isConfigSource_ConfigSourceSpecifier interface {
	isConfigSource_ConfigSourceSpecifier()
}

type ConfigSource_Path struct {
	Path string `protobuf:"bytes,1,opt,name=path,proto3,oneof"`
}

type ConfigSource_ApiConfigSource struct {
	ApiConfigSource *ApiConfigSource `protobuf:"bytes,2,opt,name=api_config_source,json=apiConfigSource,proto3,oneof"`
}

type ConfigSource_Ads struct {
	Ads *AggregatedConfigSource `protobuf:"bytes,3,opt,name=ads,proto3,oneof"`
}

func (*ConfigSource_Path) isConfigSource_ConfigSourceSpecifier() {}

func (*ConfigSource_ApiConfigSource) isConfigSource_ConfigSourceSpecifier() {}

func (*ConfigSource_Ads) isConfigSource_ConfigSourceSpecifier() {}

func (m *ConfigSource) GetConfigSourceSpecifier() isConfigSource_ConfigSourceSpecifier {
	if m != nil {
		return m.ConfigSourceSpecifier
	}
	return nil
}

func (m *ConfigSource) GetPath() string {
	if x, ok := m.GetConfigSourceSpecifier().(*ConfigSource_Path); ok {
		return x.Path
	}
	return ""
}

func (m *ConfigSource) GetApiConfigSource() *ApiConfigSource {
	if x, ok := m.GetConfigSourceSpecifier().(*ConfigSource_ApiConfigSource); ok {
		return x.ApiConfigSource
	}
	return nil
}

func (m *ConfigSource) GetAds() *AggregatedConfigSource {
	if x, ok := m.GetConfigSourceSpecifier().(*ConfigSource_Ads); ok {
		return x.Ads
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ConfigSource) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ConfigSource_OneofMarshaler, _ConfigSource_OneofUnmarshaler, _ConfigSource_OneofSizer, []interface{}{
		(*ConfigSource_Path)(nil),
		(*ConfigSource_ApiConfigSource)(nil),
		(*ConfigSource_Ads)(nil),
	}
}

func _ConfigSource_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ConfigSource)
	// config_source_specifier
	switch x := m.ConfigSourceSpecifier.(type) {
	case *ConfigSource_Path:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Path)
	case *ConfigSource_ApiConfigSource:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ApiConfigSource); err != nil {
			return err
		}
	case *ConfigSource_Ads:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ads); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ConfigSource.ConfigSourceSpecifier has unexpected type %T", x)
	}
	return nil
}

func _ConfigSource_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ConfigSource)
	switch tag {
	case 1: // config_source_specifier.path
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.ConfigSourceSpecifier = &ConfigSource_Path{x}
		return true, err
	case 2: // config_source_specifier.api_config_source
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ApiConfigSource)
		err := b.DecodeMessage(msg)
		m.ConfigSourceSpecifier = &ConfigSource_ApiConfigSource{msg}
		return true, err
	case 3: // config_source_specifier.ads
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AggregatedConfigSource)
		err := b.DecodeMessage(msg)
		m.ConfigSourceSpecifier = &ConfigSource_Ads{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ConfigSource_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ConfigSource)
	// config_source_specifier
	switch x := m.ConfigSourceSpecifier.(type) {
	case *ConfigSource_Path:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Path)))
		n += len(x.Path)
	case *ConfigSource_ApiConfigSource:
		s := proto.Size(x.ApiConfigSource)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ConfigSource_Ads:
		s := proto.Size(x.Ads)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterEnum("envoy.api.v2.core.ApiConfigSource_ApiType", ApiConfigSource_ApiType_name, ApiConfigSource_ApiType_value)
	proto.RegisterType((*ApiConfigSource)(nil), "envoy.api.v2.core.ApiConfigSource")
	proto.RegisterType((*AggregatedConfigSource)(nil), "envoy.api.v2.core.AggregatedConfigSource")
	proto.RegisterType((*ConfigSource)(nil), "envoy.api.v2.core.ConfigSource")
}

func init() {
	proto.RegisterFile("envoy/api/v2/core/config_source.proto", fileDescriptor_1ffcc55cf4c30535)
}

var fileDescriptor_1ffcc55cf4c30535 = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x6e, 0xda, 0x40,
	0x10, 0x66, 0xb1, 0x29, 0x64, 0x81, 0x40, 0xac, 0xa8, 0x71, 0x73, 0xa0, 0x88, 0xb6, 0x12, 0xcd,
	0xc1, 0x96, 0xe8, 0xb9, 0x07, 0x7e, 0x22, 0x38, 0x54, 0x55, 0xba, 0x70, 0xe9, 0xc9, 0xda, 0x98,
	0x61, 0xb3, 0x12, 0x61, 0xb7, 0xbb, 0x6b, 0x24, 0xae, 0x7d, 0x8a, 0xaa, 0x4f, 0x50, 0xf5, 0x09,
	0xaa, 0x9e, 0xf2, 0x16, 0x3d, 0x56, 0xea, 0x2d, 0x6f, 0x51, 0xad, 0x6d, 0x94, 0x90, 0x20, 0xe5,
	0x36, 0xf3, 0x79, 0xbe, 0xcf, 0xdf, 0x7c, 0xb3, 0xf8, 0x0d, 0xac, 0xd6, 0x62, 0x13, 0x52, 0xc9,
	0xc3, 0x75, 0x2f, 0x8c, 0x85, 0x82, 0x30, 0x16, 0xab, 0x05, 0x67, 0x91, 0x16, 0x89, 0x8a, 0x21,
	0x90, 0x4a, 0x18, 0xe1, 0x1d, 0xa5, 0x63, 0x01, 0x95, 0x3c, 0x58, 0xf7, 0x02, 0x3b, 0x76, 0xfa,
	0xfa, 0x31, 0x93, 0x29, 0x19, 0x47, 0x1a, 0xd4, 0x9a, 0x6f, 0x89, 0xa7, 0x2d, 0x26, 0x04, 0x5b,
	0x42, 0x98, 0x76, 0x97, 0xc9, 0x22, 0x9c, 0x27, 0x8a, 0x1a, 0x2e, 0x56, 0xf9, 0xf7, 0x93, 0x35,
	0x5d, 0xf2, 0x39, 0x35, 0x10, 0x6e, 0x8b, 0xfc, 0xc3, 0x31, 0x13, 0x4c, 0xa4, 0x65, 0x68, 0xab,
	0x0c, 0xed, 0x7c, 0x77, 0x70, 0xa3, 0x2f, 0xf9, 0x30, 0xb5, 0x38, 0x4d, 0x1d, 0x7a, 0x9f, 0x70,
	0x85, 0x4a, 0x1e, 0x99, 0x8d, 0x04, 0x1f, 0xb5, 0x51, 0xf7, 0xb0, 0x77, 0x16, 0x3c, 0xb2, 0x1b,
	0x3c, 0x60, 0xd9, 0x7e, 0xb6, 0x91, 0x30, 0xc0, 0xbf, 0x6f, 0x6f, 0x9c, 0xd2, 0x57, 0x54, 0x6c,
	0x22, 0x52, 0xa6, 0x19, 0xe8, 0xbd, 0xc2, 0xf5, 0x78, 0x99, 0x68, 0x03, 0x2a, 0x5a, 0xd1, 0x6b,
	0xd0, 0x7e, 0xb1, 0xed, 0x74, 0x0f, 0x48, 0x2d, 0x07, 0x3f, 0x5a, 0xcc, 0x1b, 0xe2, 0xfa, 0xfd,
	0x85, 0xb5, 0xef, 0xb6, 0x9d, 0x6e, 0xb5, 0xd7, 0xda, 0xf3, 0xf3, 0xb1, 0x92, 0xf1, 0x34, 0x1b,
	0x23, 0x35, 0x76, 0xd7, 0x68, 0x6f, 0x84, 0xeb, 0x0a, 0x16, 0x0a, 0xf4, 0x55, 0x34, 0x87, 0x25,
	0xdd, 0xf8, 0x4e, 0x1b, 0x75, 0xab, 0xbd, 0x17, 0x41, 0x96, 0x5b, 0xb0, 0xcd, 0x2d, 0x18, 0xe5,
	0xb9, 0x0d, 0xdc, 0x6f, 0x7f, 0x5f, 0x22, 0x52, 0xcb, 0x59, 0x23, 0x4b, 0xf2, 0x66, 0xb8, 0xa1,
	0xe0, 0x4b, 0x02, 0xda, 0x44, 0x86, 0x5f, 0x83, 0x48, 0x8c, 0x5f, 0x7a, 0x4a, 0xa7, 0x69, 0x75,
	0xec, 0xf2, 0xe5, 0x9f, 0xc8, 0x3d, 0x2b, 0x56, 0x0a, 0xe4, 0x30, 0xd7, 0x98, 0x65, 0x12, 0x9d,
	0x00, 0x97, 0xf3, 0x94, 0xbc, 0x06, 0xae, 0x92, 0xf3, 0xe9, 0x2c, 0xfa, 0x70, 0x3e, 0xee, 0x0f,
	0x3f, 0x37, 0x0b, 0x5e, 0x05, 0xbb, 0x16, 0x68, 0x22, 0x5b, 0x8d, 0xc9, 0xc5, 0xb0, 0x59, 0xec,
	0xf8, 0xf8, 0x79, 0x9f, 0x31, 0x05, 0x8c, 0x1a, 0x98, 0xdf, 0x0f, 0xbb, 0xf3, 0x07, 0xe1, 0xda,
	0xce, 0xcd, 0x8e, 0xb1, 0x2b, 0xa9, 0xb9, 0x4a, 0xef, 0x75, 0x30, 0x29, 0x90, 0xb4, 0xf3, 0x2e,
	0xf0, 0x91, 0xbd, 0xe4, 0xce, 0x03, 0xf4, 0x8b, 0xe9, 0x22, 0x9d, 0xa7, 0x4f, 0x3a, 0x29, 0x90,
	0x06, 0x7d, 0xf0, 0x36, 0xde, 0x63, 0x87, 0xce, 0x75, 0x1e, 0xea, 0xdb, 0x7d, 0x1a, 0x7b, 0x0d,
	0x4f, 0x0a, 0xc4, 0xf2, 0x06, 0x6d, 0x7c, 0xb2, 0x63, 0x26, 0xd2, 0x12, 0x62, 0xbe, 0xe0, 0xa0,
	0xbc, 0xd2, 0xaf, 0xdb, 0x1b, 0x07, 0x0d, 0xdc, 0x1f, 0xff, 0x5a, 0xe8, 0xf2, 0x59, 0x1a, 0xef,
	0xbb, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xfb, 0xbb, 0xef, 0x4e, 0x03, 0x00, 0x00,
}