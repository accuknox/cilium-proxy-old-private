// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/accesslog/v2/accesslog.proto

package v2

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	route "github.com/cilium/proxy/go/envoy/api/v2/route"
	_type "github.com/cilium/proxy/go/envoy/type"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type ComparisonFilter_Op int32

const (
	// =
	ComparisonFilter_EQ ComparisonFilter_Op = 0
	// >=
	ComparisonFilter_GE ComparisonFilter_Op = 1
	// <=
	ComparisonFilter_LE ComparisonFilter_Op = 2
)

var ComparisonFilter_Op_name = map[int32]string{
	0: "EQ",
	1: "GE",
	2: "LE",
}

var ComparisonFilter_Op_value = map[string]int32{
	"EQ": 0,
	"GE": 1,
	"LE": 2,
}

func (x ComparisonFilter_Op) String() string {
	return proto.EnumName(ComparisonFilter_Op_name, int32(x))
}

func (ComparisonFilter_Op) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_67bfd82f8b509e9f, []int{2, 0}
}

type AccessLog struct {
	// The name of the access log implementation to instantiate. The name must
	// match a statically registered access log. Current built-in loggers include:
	//
	// #. "envoy.file_access_log"
	// #. "envoy.http_grpc_access_log"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Filter which is used to determine if the access log needs to be written.
	Filter *AccessLogFilter `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	// Custom configuration that depends on the access log being instantiated. Built-in
	// configurations include:
	//
	// #. "envoy.file_access_log": :ref:`FileAccessLog
	//    <envoy_api_msg_config.accesslog.v2.FileAccessLog>`
	// #. "envoy.http_grpc_access_log": :ref:`HttpGrpcAccessLogConfig
	//    <envoy_api_msg_config.accesslog.v2.HttpGrpcAccessLogConfig>`
	//
	// Types that are valid to be assigned to ConfigType:
	//	*AccessLog_Config
	//	*AccessLog_TypedConfig
	ConfigType           isAccessLog_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AccessLog) Reset()         { *m = AccessLog{} }
func (m *AccessLog) String() string { return proto.CompactTextString(m) }
func (*AccessLog) ProtoMessage()    {}
func (*AccessLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_67bfd82f8b509e9f, []int{0}
}

func (m *AccessLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessLog.Unmarshal(m, b)
}
func (m *AccessLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessLog.Marshal(b, m, deterministic)
}
func (m *AccessLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessLog.Merge(m, src)
}
func (m *AccessLog) XXX_Size() int {
	return xxx_messageInfo_AccessLog.Size(m)
}
func (m *AccessLog) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessLog.DiscardUnknown(m)
}

var xxx_messageInfo_AccessLog proto.InternalMessageInfo

func (m *AccessLog) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AccessLog) GetFilter() *AccessLogFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type isAccessLog_ConfigType interface {
	isAccessLog_ConfigType()
}

type AccessLog_Config struct {
	Config *_struct.Struct `protobuf:"bytes,3,opt,name=config,proto3,oneof"`
}

type AccessLog_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,4,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*AccessLog_Config) isAccessLog_ConfigType() {}

func (*AccessLog_TypedConfig) isAccessLog_ConfigType() {}

func (m *AccessLog) GetConfigType() isAccessLog_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

// Deprecated: Do not use.
func (m *AccessLog) GetConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*AccessLog_Config); ok {
		return x.Config
	}
	return nil
}

func (m *AccessLog) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*AccessLog_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AccessLog) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AccessLog_Config)(nil),
		(*AccessLog_TypedConfig)(nil),
	}
}

type AccessLogFilter struct {
	// Types that are valid to be assigned to FilterSpecifier:
	//	*AccessLogFilter_StatusCodeFilter
	//	*AccessLogFilter_DurationFilter
	//	*AccessLogFilter_NotHealthCheckFilter
	//	*AccessLogFilter_TraceableFilter
	//	*AccessLogFilter_RuntimeFilter
	//	*AccessLogFilter_AndFilter
	//	*AccessLogFilter_OrFilter
	//	*AccessLogFilter_HeaderFilter
	//	*AccessLogFilter_ResponseFlagFilter
	FilterSpecifier      isAccessLogFilter_FilterSpecifier `protobuf_oneof:"filter_specifier"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *AccessLogFilter) Reset()         { *m = AccessLogFilter{} }
func (m *AccessLogFilter) String() string { return proto.CompactTextString(m) }
func (*AccessLogFilter) ProtoMessage()    {}
func (*AccessLogFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_67bfd82f8b509e9f, []int{1}
}

func (m *AccessLogFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessLogFilter.Unmarshal(m, b)
}
func (m *AccessLogFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessLogFilter.Marshal(b, m, deterministic)
}
func (m *AccessLogFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessLogFilter.Merge(m, src)
}
func (m *AccessLogFilter) XXX_Size() int {
	return xxx_messageInfo_AccessLogFilter.Size(m)
}
func (m *AccessLogFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessLogFilter.DiscardUnknown(m)
}

var xxx_messageInfo_AccessLogFilter proto.InternalMessageInfo

type isAccessLogFilter_FilterSpecifier interface {
	isAccessLogFilter_FilterSpecifier()
}

type AccessLogFilter_StatusCodeFilter struct {
	StatusCodeFilter *StatusCodeFilter `protobuf:"bytes,1,opt,name=status_code_filter,json=statusCodeFilter,proto3,oneof"`
}

type AccessLogFilter_DurationFilter struct {
	DurationFilter *DurationFilter `protobuf:"bytes,2,opt,name=duration_filter,json=durationFilter,proto3,oneof"`
}

type AccessLogFilter_NotHealthCheckFilter struct {
	NotHealthCheckFilter *NotHealthCheckFilter `protobuf:"bytes,3,opt,name=not_health_check_filter,json=notHealthCheckFilter,proto3,oneof"`
}

type AccessLogFilter_TraceableFilter struct {
	TraceableFilter *TraceableFilter `protobuf:"bytes,4,opt,name=traceable_filter,json=traceableFilter,proto3,oneof"`
}

type AccessLogFilter_RuntimeFilter struct {
	RuntimeFilter *RuntimeFilter `protobuf:"bytes,5,opt,name=runtime_filter,json=runtimeFilter,proto3,oneof"`
}

type AccessLogFilter_AndFilter struct {
	AndFilter *AndFilter `protobuf:"bytes,6,opt,name=and_filter,json=andFilter,proto3,oneof"`
}

type AccessLogFilter_OrFilter struct {
	OrFilter *OrFilter `protobuf:"bytes,7,opt,name=or_filter,json=orFilter,proto3,oneof"`
}

type AccessLogFilter_HeaderFilter struct {
	HeaderFilter *HeaderFilter `protobuf:"bytes,8,opt,name=header_filter,json=headerFilter,proto3,oneof"`
}

type AccessLogFilter_ResponseFlagFilter struct {
	ResponseFlagFilter *ResponseFlagFilter `protobuf:"bytes,9,opt,name=response_flag_filter,json=responseFlagFilter,proto3,oneof"`
}

func (*AccessLogFilter_StatusCodeFilter) isAccessLogFilter_FilterSpecifier() {}

func (*AccessLogFilter_DurationFilter) isAccessLogFilter_FilterSpecifier() {}

func (*AccessLogFilter_NotHealthCheckFilter) isAccessLogFilter_FilterSpecifier() {}

func (*AccessLogFilter_TraceableFilter) isAccessLogFilter_FilterSpecifier() {}

func (*AccessLogFilter_RuntimeFilter) isAccessLogFilter_FilterSpecifier() {}

func (*AccessLogFilter_AndFilter) isAccessLogFilter_FilterSpecifier() {}

func (*AccessLogFilter_OrFilter) isAccessLogFilter_FilterSpecifier() {}

func (*AccessLogFilter_HeaderFilter) isAccessLogFilter_FilterSpecifier() {}

func (*AccessLogFilter_ResponseFlagFilter) isAccessLogFilter_FilterSpecifier() {}

func (m *AccessLogFilter) GetFilterSpecifier() isAccessLogFilter_FilterSpecifier {
	if m != nil {
		return m.FilterSpecifier
	}
	return nil
}

func (m *AccessLogFilter) GetStatusCodeFilter() *StatusCodeFilter {
	if x, ok := m.GetFilterSpecifier().(*AccessLogFilter_StatusCodeFilter); ok {
		return x.StatusCodeFilter
	}
	return nil
}

func (m *AccessLogFilter) GetDurationFilter() *DurationFilter {
	if x, ok := m.GetFilterSpecifier().(*AccessLogFilter_DurationFilter); ok {
		return x.DurationFilter
	}
	return nil
}

func (m *AccessLogFilter) GetNotHealthCheckFilter() *NotHealthCheckFilter {
	if x, ok := m.GetFilterSpecifier().(*AccessLogFilter_NotHealthCheckFilter); ok {
		return x.NotHealthCheckFilter
	}
	return nil
}

func (m *AccessLogFilter) GetTraceableFilter() *TraceableFilter {
	if x, ok := m.GetFilterSpecifier().(*AccessLogFilter_TraceableFilter); ok {
		return x.TraceableFilter
	}
	return nil
}

func (m *AccessLogFilter) GetRuntimeFilter() *RuntimeFilter {
	if x, ok := m.GetFilterSpecifier().(*AccessLogFilter_RuntimeFilter); ok {
		return x.RuntimeFilter
	}
	return nil
}

func (m *AccessLogFilter) GetAndFilter() *AndFilter {
	if x, ok := m.GetFilterSpecifier().(*AccessLogFilter_AndFilter); ok {
		return x.AndFilter
	}
	return nil
}

func (m *AccessLogFilter) GetOrFilter() *OrFilter {
	if x, ok := m.GetFilterSpecifier().(*AccessLogFilter_OrFilter); ok {
		return x.OrFilter
	}
	return nil
}

func (m *AccessLogFilter) GetHeaderFilter() *HeaderFilter {
	if x, ok := m.GetFilterSpecifier().(*AccessLogFilter_HeaderFilter); ok {
		return x.HeaderFilter
	}
	return nil
}

func (m *AccessLogFilter) GetResponseFlagFilter() *ResponseFlagFilter {
	if x, ok := m.GetFilterSpecifier().(*AccessLogFilter_ResponseFlagFilter); ok {
		return x.ResponseFlagFilter
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AccessLogFilter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AccessLogFilter_StatusCodeFilter)(nil),
		(*AccessLogFilter_DurationFilter)(nil),
		(*AccessLogFilter_NotHealthCheckFilter)(nil),
		(*AccessLogFilter_TraceableFilter)(nil),
		(*AccessLogFilter_RuntimeFilter)(nil),
		(*AccessLogFilter_AndFilter)(nil),
		(*AccessLogFilter_OrFilter)(nil),
		(*AccessLogFilter_HeaderFilter)(nil),
		(*AccessLogFilter_ResponseFlagFilter)(nil),
	}
}

// Filter on an integer comparison.
type ComparisonFilter struct {
	// Comparison operator.
	Op ComparisonFilter_Op `protobuf:"varint,1,opt,name=op,proto3,enum=envoy.config.filter.accesslog.v2.ComparisonFilter_Op" json:"op,omitempty"`
	// Value to compare against.
	Value                *core.RuntimeUInt32 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ComparisonFilter) Reset()         { *m = ComparisonFilter{} }
func (m *ComparisonFilter) String() string { return proto.CompactTextString(m) }
func (*ComparisonFilter) ProtoMessage()    {}
func (*ComparisonFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_67bfd82f8b509e9f, []int{2}
}

func (m *ComparisonFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComparisonFilter.Unmarshal(m, b)
}
func (m *ComparisonFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComparisonFilter.Marshal(b, m, deterministic)
}
func (m *ComparisonFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComparisonFilter.Merge(m, src)
}
func (m *ComparisonFilter) XXX_Size() int {
	return xxx_messageInfo_ComparisonFilter.Size(m)
}
func (m *ComparisonFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ComparisonFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ComparisonFilter proto.InternalMessageInfo

func (m *ComparisonFilter) GetOp() ComparisonFilter_Op {
	if m != nil {
		return m.Op
	}
	return ComparisonFilter_EQ
}

func (m *ComparisonFilter) GetValue() *core.RuntimeUInt32 {
	if m != nil {
		return m.Value
	}
	return nil
}

// Filters on HTTP response/status code.
type StatusCodeFilter struct {
	// Comparison.
	Comparison           *ComparisonFilter `protobuf:"bytes,1,opt,name=comparison,proto3" json:"comparison,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StatusCodeFilter) Reset()         { *m = StatusCodeFilter{} }
func (m *StatusCodeFilter) String() string { return proto.CompactTextString(m) }
func (*StatusCodeFilter) ProtoMessage()    {}
func (*StatusCodeFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_67bfd82f8b509e9f, []int{3}
}

func (m *StatusCodeFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusCodeFilter.Unmarshal(m, b)
}
func (m *StatusCodeFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusCodeFilter.Marshal(b, m, deterministic)
}
func (m *StatusCodeFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusCodeFilter.Merge(m, src)
}
func (m *StatusCodeFilter) XXX_Size() int {
	return xxx_messageInfo_StatusCodeFilter.Size(m)
}
func (m *StatusCodeFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusCodeFilter.DiscardUnknown(m)
}

var xxx_messageInfo_StatusCodeFilter proto.InternalMessageInfo

func (m *StatusCodeFilter) GetComparison() *ComparisonFilter {
	if m != nil {
		return m.Comparison
	}
	return nil
}

// Filters on total request duration in milliseconds.
type DurationFilter struct {
	// Comparison.
	Comparison           *ComparisonFilter `protobuf:"bytes,1,opt,name=comparison,proto3" json:"comparison,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DurationFilter) Reset()         { *m = DurationFilter{} }
func (m *DurationFilter) String() string { return proto.CompactTextString(m) }
func (*DurationFilter) ProtoMessage()    {}
func (*DurationFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_67bfd82f8b509e9f, []int{4}
}

func (m *DurationFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DurationFilter.Unmarshal(m, b)
}
func (m *DurationFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DurationFilter.Marshal(b, m, deterministic)
}
func (m *DurationFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DurationFilter.Merge(m, src)
}
func (m *DurationFilter) XXX_Size() int {
	return xxx_messageInfo_DurationFilter.Size(m)
}
func (m *DurationFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_DurationFilter.DiscardUnknown(m)
}

var xxx_messageInfo_DurationFilter proto.InternalMessageInfo

func (m *DurationFilter) GetComparison() *ComparisonFilter {
	if m != nil {
		return m.Comparison
	}
	return nil
}

// Filters for requests that are not health check requests. A health check
// request is marked by the health check filter.
type NotHealthCheckFilter struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotHealthCheckFilter) Reset()         { *m = NotHealthCheckFilter{} }
func (m *NotHealthCheckFilter) String() string { return proto.CompactTextString(m) }
func (*NotHealthCheckFilter) ProtoMessage()    {}
func (*NotHealthCheckFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_67bfd82f8b509e9f, []int{5}
}

func (m *NotHealthCheckFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotHealthCheckFilter.Unmarshal(m, b)
}
func (m *NotHealthCheckFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotHealthCheckFilter.Marshal(b, m, deterministic)
}
func (m *NotHealthCheckFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotHealthCheckFilter.Merge(m, src)
}
func (m *NotHealthCheckFilter) XXX_Size() int {
	return xxx_messageInfo_NotHealthCheckFilter.Size(m)
}
func (m *NotHealthCheckFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_NotHealthCheckFilter.DiscardUnknown(m)
}

var xxx_messageInfo_NotHealthCheckFilter proto.InternalMessageInfo

// Filters for requests that are traceable. See the tracing overview for more
// information on how a request becomes traceable.
type TraceableFilter struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TraceableFilter) Reset()         { *m = TraceableFilter{} }
func (m *TraceableFilter) String() string { return proto.CompactTextString(m) }
func (*TraceableFilter) ProtoMessage()    {}
func (*TraceableFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_67bfd82f8b509e9f, []int{6}
}

func (m *TraceableFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceableFilter.Unmarshal(m, b)
}
func (m *TraceableFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceableFilter.Marshal(b, m, deterministic)
}
func (m *TraceableFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceableFilter.Merge(m, src)
}
func (m *TraceableFilter) XXX_Size() int {
	return xxx_messageInfo_TraceableFilter.Size(m)
}
func (m *TraceableFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceableFilter.DiscardUnknown(m)
}

var xxx_messageInfo_TraceableFilter proto.InternalMessageInfo

// Filters for random sampling of requests.
type RuntimeFilter struct {
	// Runtime key to get an optional overridden numerator for use in the *percent_sampled* field.
	// If found in runtime, this value will replace the default numerator.
	RuntimeKey string `protobuf:"bytes,1,opt,name=runtime_key,json=runtimeKey,proto3" json:"runtime_key,omitempty"`
	// The default sampling percentage. If not specified, defaults to 0% with denominator of 100.
	PercentSampled *_type.FractionalPercent `protobuf:"bytes,2,opt,name=percent_sampled,json=percentSampled,proto3" json:"percent_sampled,omitempty"`
	// By default, sampling pivots on the header
	// :ref:`x-request-id<config_http_conn_man_headers_x-request-id>` being present. If
	// :ref:`x-request-id<config_http_conn_man_headers_x-request-id>` is present, the filter will
	// consistently sample across multiple hosts based on the runtime key value and the value
	// extracted from :ref:`x-request-id<config_http_conn_man_headers_x-request-id>`. If it is
	// missing, or *use_independent_randomness* is set to true, the filter will randomly sample based
	// on the runtime key value alone. *use_independent_randomness* can be used for logging kill
	// switches within complex nested :ref:`AndFilter
	// <envoy_api_msg_config.filter.accesslog.v2.AndFilter>` and :ref:`OrFilter
	// <envoy_api_msg_config.filter.accesslog.v2.OrFilter>` blocks that are easier to reason about
	// from a probability perspective (i.e., setting to true will cause the filter to behave like
	// an independent random variable when composed within logical operator filters).
	UseIndependentRandomness bool     `protobuf:"varint,3,opt,name=use_independent_randomness,json=useIndependentRandomness,proto3" json:"use_independent_randomness,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *RuntimeFilter) Reset()         { *m = RuntimeFilter{} }
func (m *RuntimeFilter) String() string { return proto.CompactTextString(m) }
func (*RuntimeFilter) ProtoMessage()    {}
func (*RuntimeFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_67bfd82f8b509e9f, []int{7}
}

func (m *RuntimeFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuntimeFilter.Unmarshal(m, b)
}
func (m *RuntimeFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuntimeFilter.Marshal(b, m, deterministic)
}
func (m *RuntimeFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuntimeFilter.Merge(m, src)
}
func (m *RuntimeFilter) XXX_Size() int {
	return xxx_messageInfo_RuntimeFilter.Size(m)
}
func (m *RuntimeFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_RuntimeFilter.DiscardUnknown(m)
}

var xxx_messageInfo_RuntimeFilter proto.InternalMessageInfo

func (m *RuntimeFilter) GetRuntimeKey() string {
	if m != nil {
		return m.RuntimeKey
	}
	return ""
}

func (m *RuntimeFilter) GetPercentSampled() *_type.FractionalPercent {
	if m != nil {
		return m.PercentSampled
	}
	return nil
}

func (m *RuntimeFilter) GetUseIndependentRandomness() bool {
	if m != nil {
		return m.UseIndependentRandomness
	}
	return false
}

// Performs a logical “and” operation on the result of each filter in filters.
// Filters are evaluated sequentially and if one of them returns false, the
// filter returns false immediately.
type AndFilter struct {
	Filters              []*AccessLogFilter `protobuf:"bytes,1,rep,name=filters,proto3" json:"filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AndFilter) Reset()         { *m = AndFilter{} }
func (m *AndFilter) String() string { return proto.CompactTextString(m) }
func (*AndFilter) ProtoMessage()    {}
func (*AndFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_67bfd82f8b509e9f, []int{8}
}

func (m *AndFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AndFilter.Unmarshal(m, b)
}
func (m *AndFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AndFilter.Marshal(b, m, deterministic)
}
func (m *AndFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AndFilter.Merge(m, src)
}
func (m *AndFilter) XXX_Size() int {
	return xxx_messageInfo_AndFilter.Size(m)
}
func (m *AndFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_AndFilter.DiscardUnknown(m)
}

var xxx_messageInfo_AndFilter proto.InternalMessageInfo

func (m *AndFilter) GetFilters() []*AccessLogFilter {
	if m != nil {
		return m.Filters
	}
	return nil
}

// Performs a logical “or” operation on the result of each individual filter.
// Filters are evaluated sequentially and if one of them returns true, the
// filter returns true immediately.
type OrFilter struct {
	Filters              []*AccessLogFilter `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *OrFilter) Reset()         { *m = OrFilter{} }
func (m *OrFilter) String() string { return proto.CompactTextString(m) }
func (*OrFilter) ProtoMessage()    {}
func (*OrFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_67bfd82f8b509e9f, []int{9}
}

func (m *OrFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrFilter.Unmarshal(m, b)
}
func (m *OrFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrFilter.Marshal(b, m, deterministic)
}
func (m *OrFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrFilter.Merge(m, src)
}
func (m *OrFilter) XXX_Size() int {
	return xxx_messageInfo_OrFilter.Size(m)
}
func (m *OrFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_OrFilter.DiscardUnknown(m)
}

var xxx_messageInfo_OrFilter proto.InternalMessageInfo

func (m *OrFilter) GetFilters() []*AccessLogFilter {
	if m != nil {
		return m.Filters
	}
	return nil
}

// Filters requests based on the presence or value of a request header.
type HeaderFilter struct {
	// Only requests with a header which matches the specified HeaderMatcher will pass the filter
	// check.
	Header               *route.HeaderMatcher `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *HeaderFilter) Reset()         { *m = HeaderFilter{} }
func (m *HeaderFilter) String() string { return proto.CompactTextString(m) }
func (*HeaderFilter) ProtoMessage()    {}
func (*HeaderFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_67bfd82f8b509e9f, []int{10}
}

func (m *HeaderFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderFilter.Unmarshal(m, b)
}
func (m *HeaderFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderFilter.Marshal(b, m, deterministic)
}
func (m *HeaderFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderFilter.Merge(m, src)
}
func (m *HeaderFilter) XXX_Size() int {
	return xxx_messageInfo_HeaderFilter.Size(m)
}
func (m *HeaderFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderFilter.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderFilter proto.InternalMessageInfo

func (m *HeaderFilter) GetHeader() *route.HeaderMatcher {
	if m != nil {
		return m.Header
	}
	return nil
}

// Filters requests that received responses with an Envoy response flag set.
// A list of the response flags can be found
// in the access log formatter :ref:`documentation<config_access_log_format_response_flags>`.
type ResponseFlagFilter struct {
	// Only responses with the any of the flags listed in this field will be logged.
	// This field is optional. If it is not specified, then any response flag will pass
	// the filter check.
	Flags                []string `protobuf:"bytes,1,rep,name=flags,proto3" json:"flags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseFlagFilter) Reset()         { *m = ResponseFlagFilter{} }
func (m *ResponseFlagFilter) String() string { return proto.CompactTextString(m) }
func (*ResponseFlagFilter) ProtoMessage()    {}
func (*ResponseFlagFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_67bfd82f8b509e9f, []int{11}
}

func (m *ResponseFlagFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseFlagFilter.Unmarshal(m, b)
}
func (m *ResponseFlagFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseFlagFilter.Marshal(b, m, deterministic)
}
func (m *ResponseFlagFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseFlagFilter.Merge(m, src)
}
func (m *ResponseFlagFilter) XXX_Size() int {
	return xxx_messageInfo_ResponseFlagFilter.Size(m)
}
func (m *ResponseFlagFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseFlagFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseFlagFilter proto.InternalMessageInfo

func (m *ResponseFlagFilter) GetFlags() []string {
	if m != nil {
		return m.Flags
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.config.filter.accesslog.v2.ComparisonFilter_Op", ComparisonFilter_Op_name, ComparisonFilter_Op_value)
	proto.RegisterType((*AccessLog)(nil), "envoy.config.filter.accesslog.v2.AccessLog")
	proto.RegisterType((*AccessLogFilter)(nil), "envoy.config.filter.accesslog.v2.AccessLogFilter")
	proto.RegisterType((*ComparisonFilter)(nil), "envoy.config.filter.accesslog.v2.ComparisonFilter")
	proto.RegisterType((*StatusCodeFilter)(nil), "envoy.config.filter.accesslog.v2.StatusCodeFilter")
	proto.RegisterType((*DurationFilter)(nil), "envoy.config.filter.accesslog.v2.DurationFilter")
	proto.RegisterType((*NotHealthCheckFilter)(nil), "envoy.config.filter.accesslog.v2.NotHealthCheckFilter")
	proto.RegisterType((*TraceableFilter)(nil), "envoy.config.filter.accesslog.v2.TraceableFilter")
	proto.RegisterType((*RuntimeFilter)(nil), "envoy.config.filter.accesslog.v2.RuntimeFilter")
	proto.RegisterType((*AndFilter)(nil), "envoy.config.filter.accesslog.v2.AndFilter")
	proto.RegisterType((*OrFilter)(nil), "envoy.config.filter.accesslog.v2.OrFilter")
	proto.RegisterType((*HeaderFilter)(nil), "envoy.config.filter.accesslog.v2.HeaderFilter")
	proto.RegisterType((*ResponseFlagFilter)(nil), "envoy.config.filter.accesslog.v2.ResponseFlagFilter")
}

func init() {
	proto.RegisterFile("envoy/config/filter/accesslog/v2/accesslog.proto", fileDescriptor_67bfd82f8b509e9f)
}

var fileDescriptor_67bfd82f8b509e9f = []byte{
	// 985 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x96, 0x41, 0x6f, 0xdb, 0x36,
	0x14, 0xc7, 0x23, 0xc6, 0x71, 0xed, 0x97, 0xc4, 0xf6, 0x88, 0x60, 0x71, 0x83, 0x6c, 0xf0, 0x74,
	0x2a, 0x3a, 0x40, 0xea, 0xdc, 0xb5, 0xc0, 0x80, 0x5d, 0x62, 0xc7, 0xae, 0x8d, 0x79, 0x4d, 0xc2,
	0xd4, 0x58, 0xb0, 0x01, 0x15, 0x68, 0x89, 0xb6, 0x85, 0xca, 0xa2, 0x40, 0xc9, 0xc6, 0x7c, 0xd8,
	0x65, 0xc7, 0x1d, 0xfb, 0x69, 0x86, 0xed, 0xd2, 0x7d, 0x9c, 0x62, 0xb7, 0x7d, 0x82, 0x41, 0x24,
	0xa5, 0xd8, 0x4e, 0x01, 0x07, 0xc3, 0x7a, 0x90, 0x25, 0xf2, 0xf1, 0xff, 0x7b, 0x7c, 0xcf, 0x7c,
	0x4f, 0x82, 0x27, 0x2c, 0x5c, 0xf0, 0xa5, 0xed, 0xf2, 0x70, 0xec, 0x4f, 0xec, 0xb1, 0x1f, 0x24,
	0x4c, 0xd8, 0xd4, 0x75, 0x59, 0x1c, 0x07, 0x7c, 0x62, 0x2f, 0x9a, 0xb7, 0x03, 0x2b, 0x12, 0x3c,
	0xe1, 0xb8, 0x21, 0x15, 0x96, 0x52, 0x58, 0x4a, 0x61, 0xdd, 0x2e, 0x5a, 0x34, 0x4f, 0x4e, 0x15,
	0x93, 0x46, 0x7e, 0xaa, 0x77, 0xb9, 0x60, 0xf6, 0x88, 0xc6, 0x4c, 0xe9, 0x4f, 0x3e, 0x5f, 0xb3,
	0x0a, 0x3e, 0x4f, 0x98, 0xfa, 0xd5, 0xf6, 0xba, 0xb2, 0x27, 0xcb, 0x88, 0xd9, 0x11, 0x13, 0x2e,
	0x0b, 0x13, 0x6d, 0x79, 0x38, 0xe1, 0x7c, 0x12, 0x30, 0x5b, 0x8e, 0x46, 0xf3, 0xb1, 0x4d, 0xc3,
	0xa5, 0x36, 0x9d, 0x6e, 0x9a, 0xe2, 0x44, 0xcc, 0xdd, 0x4c, 0x78, 0xbc, 0xa0, 0x81, 0xef, 0xd1,
	0x84, 0xd9, 0xd9, 0x83, 0x32, 0x98, 0x7f, 0x1b, 0x50, 0x3e, 0x93, 0x5b, 0x1f, 0xf0, 0x09, 0xc6,
	0x50, 0x08, 0xe9, 0x8c, 0xd5, 0x8d, 0x86, 0xf1, 0xa8, 0x4c, 0xe4, 0x33, 0xee, 0x43, 0x51, 0x85,
	0x58, 0x47, 0x0d, 0xe3, 0xd1, 0x7e, 0xf3, 0x2b, 0x6b, 0x5b, 0xf8, 0x56, 0x0e, 0xec, 0x4a, 0x23,
	0xd1, 0x00, 0xfc, 0x0c, 0x8a, 0x4a, 0x55, 0xdf, 0x95, 0xa8, 0x63, 0x4b, 0x6d, 0xda, 0xca, 0x36,
	0x6d, 0x5d, 0xcb, 0x4d, 0xb7, 0x50, 0xdd, 0xe8, 0xed, 0x10, 0xbd, 0x18, 0x7f, 0x03, 0x07, 0x69,
	0x2e, 0x3c, 0x47, 0x8b, 0x0b, 0x52, 0x7c, 0x74, 0x47, 0x7c, 0x16, 0x2e, 0x7b, 0x3b, 0x64, 0x5f,
	0xae, 0x6d, 0xcb, 0xa5, 0xad, 0x43, 0xd8, 0x57, 0x22, 0x27, 0x9d, 0x35, 0xff, 0x29, 0x42, 0x75,
	0x63, 0x73, 0x78, 0x04, 0x38, 0x4e, 0x68, 0x32, 0x8f, 0x1d, 0x97, 0x7b, 0xcc, 0xd1, 0xb1, 0x1a,
	0xd2, 0x47, 0x73, 0x7b, 0xac, 0xd7, 0x52, 0xdb, 0xe6, 0x1e, 0x53, 0xbc, 0xde, 0x0e, 0xa9, 0xc5,
	0x1b, 0x73, 0xf8, 0x27, 0xa8, 0x7a, 0x73, 0x41, 0x13, 0x9f, 0x87, 0xce, 0x5a, 0x32, 0x9f, 0x6c,
	0x77, 0x70, 0xae, 0x85, 0x39, 0xbe, 0xe2, 0xad, 0xcd, 0x60, 0x0e, 0xc7, 0x21, 0x4f, 0x9c, 0x29,
	0xa3, 0x41, 0x32, 0x75, 0xdc, 0x29, 0x73, 0xdf, 0x64, 0x4e, 0x54, 0x9a, 0x9f, 0x6f, 0x77, 0xf2,
	0x92, 0x27, 0x3d, 0xa9, 0x6f, 0xa7, 0xf2, 0xdc, 0xd5, 0x51, 0xf8, 0x81, 0x79, 0xfc, 0x1a, 0x6a,
	0x89, 0xa0, 0x2e, 0xa3, 0xa3, 0x20, 0xcf, 0x57, 0xe1, 0xbe, 0x67, 0xe3, 0x55, 0xa6, 0xcc, 0x9d,
	0x54, 0x93, 0xf5, 0x29, 0x7c, 0x03, 0x15, 0x31, 0x0f, 0x13, 0x7f, 0x96, 0xd3, 0xf7, 0x24, 0xdd,
	0xde, 0x4e, 0x27, 0x4a, 0x97, 0xb3, 0x0f, 0xc5, 0xea, 0x04, 0x1e, 0x00, 0xd0, 0xd0, 0xcb, 0xa8,
	0x45, 0x49, 0xfd, 0xf2, 0x1e, 0xe7, 0x39, 0xf4, 0x72, 0x62, 0x99, 0x66, 0x03, 0xdc, 0x87, 0x32,
	0x17, 0x19, 0xec, 0x81, 0x84, 0x3d, 0xde, 0x0e, 0xbb, 0x10, 0x39, 0xab, 0xc4, 0xf5, 0x33, 0x1e,
	0xc2, 0xe1, 0x94, 0x51, 0x8f, 0xe5, 0xb8, 0x92, 0xc4, 0x59, 0xdb, 0x71, 0x3d, 0x29, 0xcb, 0x91,
	0x07, 0xd3, 0x95, 0x31, 0x9e, 0xc2, 0x91, 0x60, 0x71, 0xc4, 0xc3, 0x98, 0x39, 0xe3, 0x80, 0x4e,
	0x32, 0x7a, 0x59, 0xd2, 0xbf, 0xbe, 0x47, 0x3e, 0xb5, 0xba, 0x1b, 0xd0, 0x49, 0xee, 0x03, 0x8b,
	0x3b, 0xb3, 0xad, 0x87, 0x50, 0x53, 0x7a, 0x27, 0x8e, 0x98, 0xeb, 0x8f, 0x7d, 0x26, 0xf0, 0xde,
	0xef, 0xef, 0xdf, 0xed, 0x1a, 0xe6, 0x9f, 0x06, 0xd4, 0xda, 0x7c, 0x16, 0x51, 0xe1, 0xc7, 0xf9,
	0xa1, 0xbd, 0x02, 0xc4, 0x23, 0x59, 0x65, 0x95, 0xe6, 0xb3, 0xed, 0xfb, 0xd8, 0xd4, 0x5b, 0x17,
	0x51, 0x0b, 0xfe, 0x78, 0xff, 0x6e, 0x77, 0xef, 0x57, 0x03, 0xd5, 0x0c, 0x82, 0x78, 0x84, 0x9f,
	0xc3, 0xde, 0x82, 0x06, 0x73, 0xa6, 0x4b, 0xab, 0xa1, 0xa9, 0x34, 0xf2, 0x53, 0x42, 0xda, 0x84,
	0xb3, 0xe3, 0x31, 0xec, 0x87, 0xc9, 0xd3, 0x26, 0x51, 0xcb, 0xcd, 0x53, 0x40, 0x17, 0x11, 0x2e,
	0x02, 0xea, 0x5c, 0xd5, 0x76, 0xd2, 0xfb, 0x8b, 0x4e, 0xcd, 0x48, 0xef, 0x83, 0x4e, 0x0d, 0x99,
	0x02, 0x6a, 0x9b, 0x25, 0x8e, 0x5f, 0x03, 0xb8, 0xf9, 0x86, 0xee, 0xdf, 0x2a, 0x36, 0x83, 0xd0,
	0x11, 0xfc, 0x26, 0x23, 0x58, 0x21, 0x9a, 0x11, 0x54, 0xd6, 0xab, 0xfe, 0xa3, 0x7b, 0xfc, 0x14,
	0x8e, 0x3e, 0xd4, 0x02, 0xcc, 0x4f, 0xa0, 0xba, 0x51, 0xb0, 0xe6, 0x5f, 0x06, 0x1c, 0xae, 0x95,
	0x19, 0x7e, 0x0c, 0xfb, 0x59, 0xbd, 0xbe, 0x61, 0x4b, 0xf5, 0xf2, 0x68, 0x95, 0x53, 0x4f, 0x05,
	0x81, 0x1a, 0x06, 0x01, 0x6d, 0xfd, 0x8e, 0x2d, 0x71, 0x17, 0xaa, 0xfa, 0x95, 0xe6, 0xc4, 0x74,
	0x16, 0x05, 0xcc, 0xd3, 0x7f, 0xd7, 0x67, 0x3a, 0x9a, 0xb4, 0x4f, 0x5b, 0x5d, 0x41, 0xdd, 0x34,
	0x7a, 0x1a, 0x5c, 0xaa, 0xc5, 0xa4, 0xa2, 0x55, 0xd7, 0x4a, 0x84, 0xbf, 0x85, 0x93, 0x79, 0xcc,
	0x1c, 0x3f, 0xf4, 0x58, 0xc4, 0x42, 0x2f, 0xe5, 0x09, 0x1a, 0x7a, 0x7c, 0x16, 0xb2, 0x38, 0x96,
	0x7d, 0xaf, 0x44, 0xea, 0xf3, 0x98, 0xf5, 0x6f, 0x17, 0x90, 0xdc, 0x6e, 0x7a, 0x50, 0xce, 0x6b,
	0x1a, 0xff, 0x00, 0x0f, 0x54, 0xee, 0xe2, 0xba, 0xd1, 0xd8, 0xfd, 0x4f, 0x6f, 0x38, 0x9d, 0xd7,
	0xb7, 0x06, 0x2a, 0x21, 0x92, 0xd1, 0x4c, 0x17, 0x4a, 0x59, 0xb1, 0xaf, 0x3a, 0x41, 0xff, 0xab,
	0x93, 0x21, 0x1c, 0xac, 0xb6, 0x00, 0xdc, 0x81, 0xa2, 0x6a, 0x01, 0xfa, 0x94, 0x7c, 0xb1, 0x5e,
	0x06, 0xea, 0x3b, 0x43, 0x29, 0xbe, 0xa7, 0x89, 0x3b, 0xdd, 0x38, 0x14, 0x5a, 0x6c, 0xfe, 0x02,
	0xf8, 0x6e, 0xed, 0xe3, 0x09, 0xec, 0xa5, 0x6d, 0x44, 0x25, 0xaa, 0xdc, 0xba, 0x4a, 0x85, 0x83,
	0xb7, 0x46, 0xdf, 0x7c, 0x21, 0x3a, 0x04, 0x0d, 0x7a, 0x04, 0x0d, 0xd3, 0xeb, 0x15, 0x41, 0x03,
	0x42, 0xd0, 0x30, 0xbd, 0xba, 0x04, 0x0d, 0xdb, 0x04, 0x0d, 0x2f, 0x08, 0x7a, 0x49, 0x08, 0x3a,
	0xef, 0x13, 0xd4, 0xed, 0x13, 0x44, 0x06, 0xa4, 0x30, 0x3c, 0xeb, 0xdc, 0x90, 0x02, 0x19, 0x5c,
	0x77, 0x08, 0x3a, 0x6f, 0x93, 0xdd, 0x21, 0xb9, 0x21, 0x8a, 0xdf, 0xea, 0x82, 0xe5, 0x73, 0xb5,
	0xf3, 0x48, 0xf0, 0x9f, 0x97, 0x5b, 0x93, 0xd5, 0xaa, 0x9c, 0x65, 0xa3, 0xcb, 0xf4, 0x7b, 0xe0,
	0xd2, 0xf8, 0x11, 0x2d, 0x9a, 0xa3, 0xa2, 0xfc, 0x38, 0x78, 0xfa, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x72, 0x3d, 0x8a, 0x58, 0xe5, 0x09, 0x00, 0x00,
}
