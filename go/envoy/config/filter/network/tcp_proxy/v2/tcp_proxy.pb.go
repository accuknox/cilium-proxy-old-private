// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/tcp_proxy/v2/tcp_proxy.proto

package v2

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	v2 "github.com/cilium/proxy/go/envoy/config/filter/accesslog/v2"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type TcpProxy struct {
	// The prefix to use when emitting :ref:`statistics
	// <config_network_filters_tcp_proxy_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// Types that are valid to be assigned to ClusterSpecifier:
	//	*TcpProxy_Cluster
	//	*TcpProxy_WeightedClusters
	ClusterSpecifier isTcpProxy_ClusterSpecifier `protobuf_oneof:"cluster_specifier"`
	// Optional endpoint metadata match criteria. Only endpoints in the upstream
	// cluster with metadata matching that set in metadata_match will be
	// considered. The filter name should be specified as *envoy.lb*.
	MetadataMatch *core.Metadata `protobuf:"bytes,9,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	// The idle timeout for connections managed by the TCP proxy filter. The idle timeout
	// is defined as the period in which there are no bytes sent or received on either
	// the upstream or downstream connection. If not set, connections will never be closed
	// by the TCP proxy due to being idle.
	IdleTimeout *duration.Duration `protobuf:"bytes,8,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	// [#not-implemented-hide:] The idle timeout for connections managed by the TCP proxy
	// filter. The idle timeout is defined as the period in which there is no
	// active traffic. If not set, there is no idle timeout. When the idle timeout
	// is reached the connection will be closed. The distinction between
	// downstream_idle_timeout/upstream_idle_timeout provides a means to set
	// timeout based on the last byte sent on the downstream/upstream connection.
	DownstreamIdleTimeout *duration.Duration `protobuf:"bytes,3,opt,name=downstream_idle_timeout,json=downstreamIdleTimeout,proto3" json:"downstream_idle_timeout,omitempty"`
	// [#not-implemented-hide:]
	UpstreamIdleTimeout *duration.Duration `protobuf:"bytes,4,opt,name=upstream_idle_timeout,json=upstreamIdleTimeout,proto3" json:"upstream_idle_timeout,omitempty"`
	// Configuration for :ref:`access logs <arch_overview_access_logs>`
	// emitted by the this tcp_proxy.
	AccessLog []*v2.AccessLog `protobuf:"bytes,5,rep,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	// TCP Proxy filter configuration using deprecated V1 format. This is required for complex
	// routing until filter chain matching in the listener is implemented.
	//
	// Example:
	//
	// .. code-block:: yaml
	//
	//     - name: "envoy.tcp_proxy"
	//       config:
	//         deprecated_v1: true
	//         value:
	//           stat_prefix: "prefix"
	//           access_log:
	//             - ...
	//           route_config:
	//             routes:
	//               - cluster: "cluster"
	//                 destination_ip_list:
	//                   - "10.1.0.0/8"
	//                 destination_ports: "8080"
	//                 source_ip_list:
	//                   - "10.1.0.0/16"
	//                   - "2001:db8::/32"
	//                 source_ports: "8000,9000-9999"
	//
	// .. attention::
	//
	//   Using the deprecated V1 configuration excludes the use of any V2 configuration options. Only
	//   the V1 configuration is used. All available fields are shown in the example, although the
	//   access log configuration is omitted for simplicity. The access log configuration uses the
	//   :repo:`deprecated V1 access log configuration<source/common/json/config_schemas.cc>`.
	//
	// .. attention::
	//
	//   In the deprecated V1 configuration, source and destination CIDR ranges are specified as a
	//   list of strings with each string in CIDR notation. Source and destination ports are
	//   specified as single strings containing a comma-separated list of ports and/or port ranges.
	//
	DeprecatedV1 *TcpProxy_DeprecatedV1 `protobuf:"bytes,6,opt,name=deprecated_v1,json=deprecatedV1,proto3" json:"deprecated_v1,omitempty"` // Deprecated: Do not use.
	// The maximum number of unsuccessful connection attempts that will be made before
	// giving up. If the parameter is not specified, 1 connection attempt will be made.
	MaxConnectAttempts   *wrappers.UInt32Value `protobuf:"bytes,7,opt,name=max_connect_attempts,json=maxConnectAttempts,proto3" json:"max_connect_attempts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TcpProxy) Reset()         { *m = TcpProxy{} }
func (m *TcpProxy) String() string { return proto.CompactTextString(m) }
func (*TcpProxy) ProtoMessage()    {}
func (*TcpProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f6b35dbcbad27ba, []int{0}
}

func (m *TcpProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy.Unmarshal(m, b)
}
func (m *TcpProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy.Marshal(b, m, deterministic)
}
func (m *TcpProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy.Merge(m, src)
}
func (m *TcpProxy) XXX_Size() int {
	return xxx_messageInfo_TcpProxy.Size(m)
}
func (m *TcpProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy proto.InternalMessageInfo

func (m *TcpProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

type isTcpProxy_ClusterSpecifier interface {
	isTcpProxy_ClusterSpecifier()
}

type TcpProxy_Cluster struct {
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3,oneof"`
}

type TcpProxy_WeightedClusters struct {
	WeightedClusters *TcpProxy_WeightedCluster `protobuf:"bytes,10,opt,name=weighted_clusters,json=weightedClusters,proto3,oneof"`
}

func (*TcpProxy_Cluster) isTcpProxy_ClusterSpecifier() {}

func (*TcpProxy_WeightedClusters) isTcpProxy_ClusterSpecifier() {}

func (m *TcpProxy) GetClusterSpecifier() isTcpProxy_ClusterSpecifier {
	if m != nil {
		return m.ClusterSpecifier
	}
	return nil
}

func (m *TcpProxy) GetCluster() string {
	if x, ok := m.GetClusterSpecifier().(*TcpProxy_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *TcpProxy) GetWeightedClusters() *TcpProxy_WeightedCluster {
	if x, ok := m.GetClusterSpecifier().(*TcpProxy_WeightedClusters); ok {
		return x.WeightedClusters
	}
	return nil
}

func (m *TcpProxy) GetMetadataMatch() *core.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

func (m *TcpProxy) GetIdleTimeout() *duration.Duration {
	if m != nil {
		return m.IdleTimeout
	}
	return nil
}

func (m *TcpProxy) GetDownstreamIdleTimeout() *duration.Duration {
	if m != nil {
		return m.DownstreamIdleTimeout
	}
	return nil
}

func (m *TcpProxy) GetUpstreamIdleTimeout() *duration.Duration {
	if m != nil {
		return m.UpstreamIdleTimeout
	}
	return nil
}

func (m *TcpProxy) GetAccessLog() []*v2.AccessLog {
	if m != nil {
		return m.AccessLog
	}
	return nil
}

// Deprecated: Do not use.
func (m *TcpProxy) GetDeprecatedV1() *TcpProxy_DeprecatedV1 {
	if m != nil {
		return m.DeprecatedV1
	}
	return nil
}

func (m *TcpProxy) GetMaxConnectAttempts() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConnectAttempts
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TcpProxy) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TcpProxy_Cluster)(nil),
		(*TcpProxy_WeightedClusters)(nil),
	}
}

// TCP Proxy filter configuration using V1 format, until Envoy gets the
// ability to match source/destination at the listener level (called
// :ref:`filter chain match <envoy_api_msg_listener.FilterChainMatch>`).
type TcpProxy_DeprecatedV1 struct {
	// The route table for the filter. All filter instances must have a route
	// table, even if it is empty.
	Routes               []*TcpProxy_DeprecatedV1_TCPRoute `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *TcpProxy_DeprecatedV1) Reset()         { *m = TcpProxy_DeprecatedV1{} }
func (m *TcpProxy_DeprecatedV1) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_DeprecatedV1) ProtoMessage()    {}
func (*TcpProxy_DeprecatedV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f6b35dbcbad27ba, []int{0, 0}
}

func (m *TcpProxy_DeprecatedV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_DeprecatedV1.Unmarshal(m, b)
}
func (m *TcpProxy_DeprecatedV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_DeprecatedV1.Marshal(b, m, deterministic)
}
func (m *TcpProxy_DeprecatedV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_DeprecatedV1.Merge(m, src)
}
func (m *TcpProxy_DeprecatedV1) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_DeprecatedV1.Size(m)
}
func (m *TcpProxy_DeprecatedV1) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_DeprecatedV1.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_DeprecatedV1 proto.InternalMessageInfo

func (m *TcpProxy_DeprecatedV1) GetRoutes() []*TcpProxy_DeprecatedV1_TCPRoute {
	if m != nil {
		return m.Routes
	}
	return nil
}

// A TCP proxy route consists of a set of optional L4 criteria and the
// name of a cluster. If a downstream connection matches all the
// specified criteria, the cluster in the route is used for the
// corresponding upstream connection. Routes are tried in the order
// specified until a match is found. If no match is found, the connection
// is closed. A route with no criteria is valid and always produces a
// match.
type TcpProxy_DeprecatedV1_TCPRoute struct {
	// The cluster to connect to when a the downstream network connection
	// matches the specified criteria.
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// An optional list of IP address subnets in the form
	// “ip_address/xx”. The criteria is satisfied if the destination IP
	// address of the downstream connection is contained in at least one of
	// the specified subnets. If the parameter is not specified or the list
	// is empty, the destination IP address is ignored. The destination IP
	// address of the downstream connection might be different from the
	// addresses on which the proxy is listening if the connection has been
	// redirected.
	DestinationIpList []*core.CidrRange `protobuf:"bytes,2,rep,name=destination_ip_list,json=destinationIpList,proto3" json:"destination_ip_list,omitempty"`
	// An optional string containing a comma-separated list of port numbers
	// or ranges. The criteria is satisfied if the destination port of the
	// downstream connection is contained in at least one of the specified
	// ranges. If the parameter is not specified, the destination port is
	// ignored. The destination port address of the downstream connection
	// might be different from the port on which the proxy is listening if
	// the connection has been redirected.
	DestinationPorts string `protobuf:"bytes,3,opt,name=destination_ports,json=destinationPorts,proto3" json:"destination_ports,omitempty"`
	// An optional list of IP address subnets in the form
	// “ip_address/xx”. The criteria is satisfied if the source IP address
	// of the downstream connection is contained in at least one of the
	// specified subnets. If the parameter is not specified or the list is
	// empty, the source IP address is ignored.
	SourceIpList []*core.CidrRange `protobuf:"bytes,4,rep,name=source_ip_list,json=sourceIpList,proto3" json:"source_ip_list,omitempty"`
	// An optional string containing a comma-separated list of port numbers
	// or ranges. The criteria is satisfied if the source port of the
	// downstream connection is contained in at least one of the specified
	// ranges. If the parameter is not specified, the source port is
	// ignored.
	SourcePorts          string   `protobuf:"bytes,5,opt,name=source_ports,json=sourcePorts,proto3" json:"source_ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) Reset()         { *m = TcpProxy_DeprecatedV1_TCPRoute{} }
func (m *TcpProxy_DeprecatedV1_TCPRoute) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_DeprecatedV1_TCPRoute) ProtoMessage()    {}
func (*TcpProxy_DeprecatedV1_TCPRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f6b35dbcbad27ba, []int{0, 0, 0}
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.Unmarshal(m, b)
}
func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.Marshal(b, m, deterministic)
}
func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.Merge(m, src)
}
func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.Size(m)
}
func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute proto.InternalMessageInfo

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetDestinationIpList() []*core.CidrRange {
	if m != nil {
		return m.DestinationIpList
	}
	return nil
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetDestinationPorts() string {
	if m != nil {
		return m.DestinationPorts
	}
	return ""
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetSourceIpList() []*core.CidrRange {
	if m != nil {
		return m.SourceIpList
	}
	return nil
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetSourcePorts() string {
	if m != nil {
		return m.SourcePorts
	}
	return ""
}

// Allows for specification of multiple upstream clusters along with weights
// that indicate the percentage of traffic to be forwarded to each cluster.
// The router selects an upstream cluster based on these weights.
type TcpProxy_WeightedCluster struct {
	// Specifies one or more upstream clusters associated with the route.
	Clusters             []*TcpProxy_WeightedCluster_ClusterWeight `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *TcpProxy_WeightedCluster) Reset()         { *m = TcpProxy_WeightedCluster{} }
func (m *TcpProxy_WeightedCluster) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_WeightedCluster) ProtoMessage()    {}
func (*TcpProxy_WeightedCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f6b35dbcbad27ba, []int{0, 1}
}

func (m *TcpProxy_WeightedCluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_WeightedCluster.Unmarshal(m, b)
}
func (m *TcpProxy_WeightedCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_WeightedCluster.Marshal(b, m, deterministic)
}
func (m *TcpProxy_WeightedCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_WeightedCluster.Merge(m, src)
}
func (m *TcpProxy_WeightedCluster) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_WeightedCluster.Size(m)
}
func (m *TcpProxy_WeightedCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_WeightedCluster.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_WeightedCluster proto.InternalMessageInfo

func (m *TcpProxy_WeightedCluster) GetClusters() []*TcpProxy_WeightedCluster_ClusterWeight {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type TcpProxy_WeightedCluster_ClusterWeight struct {
	// Name of the upstream cluster.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// When a request matches the route, the choice of an upstream cluster is
	// determined by its weight. The sum of weights across all entries in the
	// clusters array determines the total weight.
	Weight               uint32   `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TcpProxy_WeightedCluster_ClusterWeight) Reset() {
	*m = TcpProxy_WeightedCluster_ClusterWeight{}
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_WeightedCluster_ClusterWeight) ProtoMessage()    {}
func (*TcpProxy_WeightedCluster_ClusterWeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f6b35dbcbad27ba, []int{0, 1, 0}
}

func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Unmarshal(m, b)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Marshal(b, m, deterministic)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Merge(m, src)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Size(m)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight proto.InternalMessageInfo

func (m *TcpProxy_WeightedCluster_ClusterWeight) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TcpProxy_WeightedCluster_ClusterWeight) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func init() {
	proto.RegisterType((*TcpProxy)(nil), "envoy.config.filter.network.tcp_proxy.v2.TcpProxy")
	proto.RegisterType((*TcpProxy_DeprecatedV1)(nil), "envoy.config.filter.network.tcp_proxy.v2.TcpProxy.DeprecatedV1")
	proto.RegisterType((*TcpProxy_DeprecatedV1_TCPRoute)(nil), "envoy.config.filter.network.tcp_proxy.v2.TcpProxy.DeprecatedV1.TCPRoute")
	proto.RegisterType((*TcpProxy_WeightedCluster)(nil), "envoy.config.filter.network.tcp_proxy.v2.TcpProxy.WeightedCluster")
	proto.RegisterType((*TcpProxy_WeightedCluster_ClusterWeight)(nil), "envoy.config.filter.network.tcp_proxy.v2.TcpProxy.WeightedCluster.ClusterWeight")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/tcp_proxy/v2/tcp_proxy.proto", fileDescriptor_1f6b35dbcbad27ba)
}

var fileDescriptor_1f6b35dbcbad27ba = []byte{
	// 799 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x6e, 0xdc, 0x44,
	0x14, 0x8e, 0xbd, 0x49, 0xba, 0x3b, 0x9b, 0x2d, 0xcd, 0xb4, 0x55, 0x8d, 0x09, 0x34, 0x85, 0x9b,
	0x55, 0x2a, 0xd9, 0x74, 0x2b, 0x21, 0xee, 0x50, 0x9d, 0x5e, 0x24, 0x28, 0x2b, 0x6d, 0xad, 0x50,
	0x04, 0x37, 0xd6, 0xc4, 0x9e, 0x75, 0x46, 0xd8, 0x9e, 0x61, 0x66, 0xbc, 0xbb, 0x7d, 0x0b, 0xe0,
	0x0a, 0xf1, 0x08, 0x3c, 0x01, 0xe2, 0xaa, 0x4f, 0x02, 0xd7, 0x7d, 0x0b, 0x34, 0x3f, 0xde, 0xf5,
	0xb2, 0x8b, 0x1a, 0x91, 0x2b, 0x7b, 0xce, 0x39, 0xdf, 0xf7, 0xcd, 0xf9, 0x1b, 0xf0, 0x25, 0xae,
	0x66, 0xf4, 0x4d, 0x98, 0xd2, 0x6a, 0x4a, 0xf2, 0x70, 0x4a, 0x0a, 0x89, 0x79, 0x58, 0x61, 0x39,
	0xa7, 0xfc, 0x87, 0x50, 0xa6, 0x2c, 0x61, 0x9c, 0x2e, 0xde, 0x84, 0xb3, 0xd1, 0xea, 0x10, 0x30,
	0x4e, 0x25, 0x85, 0x43, 0x8d, 0x0c, 0x0c, 0x32, 0x30, 0xc8, 0xc0, 0x22, 0x83, 0x55, 0xf0, 0x6c,
	0xe4, 0x7f, 0xbe, 0x4d, 0x03, 0xa5, 0x29, 0x16, 0xa2, 0xa0, 0xb9, 0xe2, 0x5e, 0x1e, 0x0c, 0xb7,
	0xff, 0xd8, 0x20, 0x10, 0x23, 0xca, 0x9b, 0x52, 0x8e, 0x43, 0x94, 0x65, 0x1c, 0x0b, 0x61, 0x03,
	0x8e, 0x36, 0x03, 0xae, 0x90, 0xc0, 0xd6, 0xfb, 0x49, 0x4e, 0x69, 0x5e, 0xe0, 0x50, 0x9f, 0xae,
	0xea, 0x69, 0x98, 0xd5, 0x1c, 0x49, 0x42, 0xab, 0xff, 0xf2, 0xcf, 0x39, 0x62, 0x0c, 0xf3, 0x86,
	0xfd, 0xd1, 0x0c, 0x15, 0x24, 0x43, 0x12, 0x87, 0xcd, 0x8f, 0x75, 0x3c, 0xc8, 0x69, 0x4e, 0xf5,
	0x6f, 0xa8, 0xfe, 0x8c, 0xf5, 0xd3, 0xdf, 0xfa, 0xa0, 0x7b, 0x99, 0xb2, 0x89, 0xca, 0x17, 0x9e,
	0x80, 0xbe, 0x90, 0x48, 0x26, 0x8c, 0xe3, 0x29, 0x59, 0x78, 0xce, 0xb1, 0x33, 0xec, 0x45, 0xbd,
	0x3f, 0xdf, 0xbd, 0xed, 0xec, 0x72, 0xf7, 0xd8, 0x89, 0x81, 0xf2, 0x4e, 0xb4, 0x13, 0xfa, 0xe0,
	0x4e, 0x5a, 0xd4, 0x42, 0x62, 0xee, 0xb9, 0x2a, 0xee, 0x6c, 0x27, 0x6e, 0x0c, 0xf0, 0x47, 0x70,
	0x38, 0xc7, 0x24, 0xbf, 0x96, 0x38, 0x4b, 0xac, 0x4d, 0x78, 0xe0, 0xd8, 0x19, 0xf6, 0x47, 0x51,
	0x70, 0xd3, 0xd2, 0x07, 0xcd, 0xb5, 0x82, 0x6f, 0x2d, 0xd7, 0xa9, 0xa1, 0x3a, 0xdb, 0x89, 0xef,
	0xcd, 0xd7, 0x4d, 0x02, 0x46, 0xe0, 0x6e, 0x89, 0x25, 0xca, 0x90, 0x44, 0x49, 0x89, 0x64, 0x7a,
	0xed, 0xf5, 0xb4, 0xde, 0x47, 0x56, 0x0f, 0x31, 0xa2, 0x38, 0x55, 0xb5, 0x83, 0xb1, 0x0d, 0x8c,
	0x07, 0x0d, 0x64, 0xac, 0x10, 0xf0, 0x02, 0x1c, 0x90, 0xac, 0xc0, 0x89, 0x24, 0x25, 0xa6, 0xb5,
	0xf4, 0xba, 0x9a, 0xe1, 0xc3, 0xc0, 0x54, 0x3c, 0x68, 0x2a, 0x1e, 0xbc, 0xb4, 0x1d, 0x89, 0xee,
	0xfe, 0xfa, 0xf7, 0x63, 0x47, 0x95, 0x67, 0xef, 0x77, 0xc7, 0x3d, 0xd9, 0x89, 0xfb, 0x0a, 0x7e,
	0x69, 0xd0, 0xf0, 0x15, 0x78, 0x94, 0xd1, 0x79, 0x25, 0x24, 0xc7, 0xa8, 0x4c, 0xd6, 0x88, 0x3b,
	0xef, 0x21, 0x8e, 0x1f, 0xae, 0x90, 0xe7, 0x2d, 0xca, 0x31, 0x78, 0x58, 0xb3, 0x6d, 0x84, 0xbb,
	0xef, 0x23, 0xbc, 0xdf, 0xe0, 0xda, 0x74, 0x5f, 0x03, 0x60, 0x86, 0x37, 0x29, 0x68, 0xee, 0xed,
	0x1d, 0x77, 0x86, 0xfd, 0xd1, 0xd3, 0xad, 0xfd, 0x59, 0xcd, 0xf8, 0x6c, 0x14, 0xbc, 0xd0, 0x87,
	0x0b, 0x9a, 0xc7, 0x3d, 0xd4, 0xfc, 0xc2, 0x6b, 0x30, 0xc8, 0x30, 0xe3, 0x38, 0x45, 0xaa, 0xe9,
	0xb3, 0x67, 0xde, 0xbe, 0xbe, 0xd2, 0x57, 0xff, 0xa3, 0xdd, 0x2f, 0x97, 0x3c, 0xaf, 0x9f, 0x45,
	0xae, 0xe7, 0xc4, 0x07, 0x59, 0xcb, 0x02, 0xbf, 0x03, 0x0f, 0x4a, 0xb4, 0x48, 0x52, 0x5a, 0x55,
	0x38, 0x95, 0x09, 0x92, 0x12, 0x97, 0x4c, 0x0a, 0xef, 0x8e, 0x16, 0x3c, 0xda, 0xa8, 0xc1, 0x37,
	0xe7, 0x95, 0x7c, 0x3e, 0x7a, 0x8d, 0x8a, 0x1a, 0xdb, 0x59, 0x3e, 0x71, 0x87, 0x4e, 0x0c, 0x4b,
	0xb4, 0x38, 0x35, 0x1c, 0x2f, 0x2c, 0x85, 0xff, 0x53, 0x07, 0x1c, 0xb4, 0xd5, 0x61, 0x01, 0xf6,
	0x39, 0xad, 0x25, 0x16, 0x9e, 0xa3, 0xab, 0x73, 0x76, 0xcb, 0x74, 0x82, 0xcb, 0xd3, 0x49, 0xac,
	0x08, 0x23, 0xa0, 0xc7, 0xe6, 0x17, 0xc7, 0xed, 0x3a, 0xb1, 0xd5, 0xf0, 0x7f, 0x76, 0x41, 0xb7,
	0x09, 0x80, 0x9f, 0xad, 0xf6, 0x6b, 0x63, 0x0f, 0x97, 0x8b, 0x76, 0x01, 0xee, 0x67, 0x58, 0x48,
	0x52, 0xe9, 0x2e, 0x27, 0x84, 0x25, 0x05, 0x11, 0xd2, 0x73, 0xf5, 0x65, 0x8f, 0xb6, 0x8c, 0xfe,
	0x29, 0xc9, 0x78, 0x8c, 0xaa, 0x1c, 0xc7, 0x87, 0x2d, 0xe0, 0x39, 0xbb, 0x20, 0x42, 0xc2, 0xa7,
	0xa0, 0x6d, 0x4c, 0x18, 0xe5, 0x52, 0xe8, 0x59, 0xed, 0xc5, 0xf7, 0x5a, 0x8e, 0x89, 0xb2, 0xab,
	0x85, 0x13, 0xb4, 0xe6, 0x29, 0x5e, 0xaa, 0xee, 0xde, 0x40, 0xf5, 0xc0, 0x60, 0xac, 0xe0, 0x13,
	0x60, 0xcf, 0x56, 0x6b, 0x4f, 0x6b, 0xf5, 0x8d, 0x4d, 0xcb, 0xf8, 0x7f, 0x39, 0xe0, 0x83, 0x7f,
	0xed, 0x3f, 0x5c, 0x80, 0xee, 0xf2, 0x55, 0x31, 0x7d, 0x99, 0xdc, 0xfe, 0x55, 0x09, 0xec, 0xd7,
	0x98, 0xd7, 0xfa, 0xb3, 0x54, 0xf3, 0x5f, 0x81, 0xc1, 0x5a, 0x18, 0xfc, 0x18, 0xec, 0x56, 0xa8,
	0xc4, 0x9b, 0x2d, 0xd2, 0x66, 0xf8, 0x04, 0xec, 0x9b, 0x97, 0x4a, 0xbf, 0x91, 0x83, 0xf6, 0xfc,
	0x59, 0x47, 0xe4, 0x83, 0x43, 0x4b, 0x9f, 0x08, 0x86, 0x53, 0x32, 0x25, 0x98, 0xc3, 0xbd, 0x3f,
	0xde, 0xbd, 0xed, 0x38, 0xd1, 0x18, 0x7c, 0x41, 0xa8, 0x49, 0xcd, 0xdc, 0xff, 0xa6, 0x59, 0x46,
	0x83, 0x26, 0xcd, 0x89, 0x5a, 0x83, 0x89, 0xf3, 0xbd, 0x3b, 0x1b, 0x5d, 0xed, 0xeb, 0x9d, 0x78,
	0xfe, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x74, 0xf0, 0xef, 0x38, 0x07, 0x00, 0x00,
}
