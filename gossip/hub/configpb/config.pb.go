// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config.proto

package configpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	keyspb "github.com/google/trillian/crypto/keyspb"
	sigpb "github.com/google/trillian/crypto/sigpb"
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

// kind indicates
type TrackedSource_Kind int32

const (
	TrackedSource_UNKNOWN     TrackedSource_Kind = 0
	TrackedSource_RFC6962STH  TrackedSource_Kind = 1
	TrackedSource_TRILLIANSLR TrackedSource_Kind = 2
)

var TrackedSource_Kind_name = map[int32]string{
	0: "UNKNOWN",
	1: "RFC6962STH",
	2: "TRILLIANSLR",
}

var TrackedSource_Kind_value = map[string]int32{
	"UNKNOWN":     0,
	"RFC6962STH":  1,
	"TRILLIANSLR": 2,
}

func (x TrackedSource_Kind) String() string {
	return proto.EnumName(TrackedSource_Kind_name, int32(x))
}

func (TrackedSource_Kind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{3, 0}
}

// HubMultiConfig wraps up a collection of HubConfig messages with a
// corresponding map of backend specifications so that the configuration
// can be parsed as a single proto.
type HubMultiConfig struct {
	// The set of hubs that will use the backends below. All the protos in this
	// HubConfigSet must set a valid hub_backend_name for the config to be usable.
	HubConfig []*HubConfig `protobuf:"bytes,1,rep,name=hub_config,json=hubConfig,proto3" json:"hub_config,omitempty"`
	// The collection of backends that this configuration will use to send requests to.
	// The key for the map is the backend name (referenced by HubConfig.backend_name)
	// and the corresponding value defines the RPC endpoint that clients should use to
	// send requests to this hub backend. These values should be in the same format as
	// rpcBackendFlag in the hub_server main and must not be an empty string.
	HubBackends          map[string]string `protobuf:"bytes,2,rep,name=hub_backends,json=hubBackends,proto3" json:"hub_backends,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HubMultiConfig) Reset()         { *m = HubMultiConfig{} }
func (m *HubMultiConfig) String() string { return proto.CompactTextString(m) }
func (*HubMultiConfig) ProtoMessage()    {}
func (*HubMultiConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{0}
}

func (m *HubMultiConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HubMultiConfig.Unmarshal(m, b)
}
func (m *HubMultiConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HubMultiConfig.Marshal(b, m, deterministic)
}
func (m *HubMultiConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HubMultiConfig.Merge(m, src)
}
func (m *HubMultiConfig) XXX_Size() int {
	return xxx_messageInfo_HubMultiConfig.Size(m)
}
func (m *HubMultiConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HubMultiConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HubMultiConfig proto.InternalMessageInfo

func (m *HubMultiConfig) GetHubConfig() []*HubConfig {
	if m != nil {
		return m.HubConfig
	}
	return nil
}

func (m *HubMultiConfig) GetHubBackends() map[string]string {
	if m != nil {
		return m.HubBackends
	}
	return nil
}

// HubConfigSet holds a set of hub configurations as a single message to allow
// configuration of a hub using a default backend.
type HubConfigSet struct {
	Config               []*HubConfig `protobuf:"bytes,1,rep,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *HubConfigSet) Reset()         { *m = HubConfigSet{} }
func (m *HubConfigSet) String() string { return proto.CompactTextString(m) }
func (*HubConfigSet) ProtoMessage()    {}
func (*HubConfigSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{1}
}

func (m *HubConfigSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HubConfigSet.Unmarshal(m, b)
}
func (m *HubConfigSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HubConfigSet.Marshal(b, m, deterministic)
}
func (m *HubConfigSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HubConfigSet.Merge(m, src)
}
func (m *HubConfigSet) XXX_Size() int {
	return xxx_messageInfo_HubConfigSet.Size(m)
}
func (m *HubConfigSet) XXX_DiscardUnknown() {
	xxx_messageInfo_HubConfigSet.DiscardUnknown(m)
}

var xxx_messageInfo_HubConfigSet proto.InternalMessageInfo

func (m *HubConfigSet) GetConfig() []*HubConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

// HubConfig describes the configuration options for a hub instance.
type HubConfig struct {
	LogId      int64            `protobuf:"varint,1,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
	Prefix     string           `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Source     []*TrackedSource `protobuf:"bytes,3,rep,name=source,proto3" json:"source,omitempty"`
	PrivateKey *any.Any         `protobuf:"bytes,4,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// The public key is included for the convenience of test tools (and obviously
	// should match the private key above); it is not used by the hub.
	PublicKey *keyspb.PublicKey `protobuf:"bytes,5,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// backend_name (if set) indicates which backend serves this hub.  When enclosed in
	// a HubMultiConfig the name must be one of the keys for the hub_backends map.
	BackendName          string   `protobuf:"bytes,6,opt,name=backend_name,json=backendName,proto3" json:"backend_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HubConfig) Reset()         { *m = HubConfig{} }
func (m *HubConfig) String() string { return proto.CompactTextString(m) }
func (*HubConfig) ProtoMessage()    {}
func (*HubConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{2}
}

func (m *HubConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HubConfig.Unmarshal(m, b)
}
func (m *HubConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HubConfig.Marshal(b, m, deterministic)
}
func (m *HubConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HubConfig.Merge(m, src)
}
func (m *HubConfig) XXX_Size() int {
	return xxx_messageInfo_HubConfig.Size(m)
}
func (m *HubConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HubConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HubConfig proto.InternalMessageInfo

func (m *HubConfig) GetLogId() int64 {
	if m != nil {
		return m.LogId
	}
	return 0
}

func (m *HubConfig) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *HubConfig) GetSource() []*TrackedSource {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *HubConfig) GetPrivateKey() *any.Any {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *HubConfig) GetPublicKey() *keyspb.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *HubConfig) GetBackendName() string {
	if m != nil {
		return m.BackendName
	}
	return ""
}

// TrackedSource describes a source that this gossip hub will accept signed blobs from.
type TrackedSource struct {
	Name                 string                              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string                              `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	PublicKey            *keyspb.PublicKey                   `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	HashAlgorithm        sigpb.DigitallySigned_HashAlgorithm `protobuf:"varint,4,opt,name=hash_algorithm,json=hashAlgorithm,proto3,enum=sigpb.DigitallySigned_HashAlgorithm" json:"hash_algorithm,omitempty"`
	Kind                 TrackedSource_Kind                  `protobuf:"varint,5,opt,name=kind,proto3,enum=configpb.TrackedSource_Kind" json:"kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *TrackedSource) Reset()         { *m = TrackedSource{} }
func (m *TrackedSource) String() string { return proto.CompactTextString(m) }
func (*TrackedSource) ProtoMessage()    {}
func (*TrackedSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{3}
}

func (m *TrackedSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackedSource.Unmarshal(m, b)
}
func (m *TrackedSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackedSource.Marshal(b, m, deterministic)
}
func (m *TrackedSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackedSource.Merge(m, src)
}
func (m *TrackedSource) XXX_Size() int {
	return xxx_messageInfo_TrackedSource.Size(m)
}
func (m *TrackedSource) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackedSource.DiscardUnknown(m)
}

var xxx_messageInfo_TrackedSource proto.InternalMessageInfo

func (m *TrackedSource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TrackedSource) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TrackedSource) GetPublicKey() *keyspb.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *TrackedSource) GetHashAlgorithm() sigpb.DigitallySigned_HashAlgorithm {
	if m != nil {
		return m.HashAlgorithm
	}
	return sigpb.DigitallySigned_NONE
}

func (m *TrackedSource) GetKind() TrackedSource_Kind {
	if m != nil {
		return m.Kind
	}
	return TrackedSource_UNKNOWN
}

func init() {
	proto.RegisterEnum("configpb.TrackedSource_Kind", TrackedSource_Kind_name, TrackedSource_Kind_value)
	proto.RegisterType((*HubMultiConfig)(nil), "configpb.HubMultiConfig")
	proto.RegisterMapType((map[string]string)(nil), "configpb.HubMultiConfig.HubBackendsEntry")
	proto.RegisterType((*HubConfigSet)(nil), "configpb.HubConfigSet")
	proto.RegisterType((*HubConfig)(nil), "configpb.HubConfig")
	proto.RegisterType((*TrackedSource)(nil), "configpb.TrackedSource")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor_3eaf2c85e69e9ea4) }

var fileDescriptor_3eaf2c85e69e9ea4 = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xfd, 0xec, 0xa4, 0xfe, 0xc8, 0x38, 0x35, 0x61, 0x29, 0x60, 0x2a, 0x2e, 0x42, 0xc4, 0x45,
	0x10, 0x92, 0x5d, 0xb9, 0xb4, 0xe2, 0x47, 0x42, 0x0a, 0x05, 0x94, 0x28, 0x21, 0xa0, 0x4d, 0x10,
	0x97, 0xd1, 0xda, 0xde, 0xd8, 0xab, 0x6c, 0xbc, 0x96, 0x7f, 0x2a, 0xfc, 0x2a, 0x3c, 0x16, 0x2f,
	0xc3, 0x2d, 0xf2, 0x7a, 0xd3, 0x3f, 0x09, 0x95, 0x1b, 0x7b, 0xce, 0xec, 0x39, 0x3b, 0x67, 0x66,
	0x16, 0xba, 0x81, 0x48, 0xd6, 0x2c, 0x72, 0xd2, 0x4c, 0x14, 0x02, 0xdd, 0x69, 0x50, 0xea, 0x1f,
	0x9e, 0x44, 0xac, 0x88, 0x4b, 0xdf, 0x09, 0xc4, 0xd6, 0x8d, 0x84, 0x88, 0x38, 0x75, 0x8b, 0x8c,
	0x71, 0xce, 0x48, 0xe2, 0x06, 0x59, 0x95, 0x16, 0xc2, 0xdd, 0xd0, 0x2a, 0x4f, 0x7d, 0xf5, 0x6b,
	0x2e, 0x38, 0x3c, 0xbe, 0x5d, 0x96, 0xd7, 0xf7, 0x37, 0x5f, 0x25, 0x7a, 0xac, 0x98, 0x12, 0xf9,
	0xe5, 0xda, 0x25, 0x49, 0xd5, 0x1c, 0x0d, 0x7e, 0x69, 0x60, 0x8d, 0x4b, 0xff, 0x73, 0xc9, 0x0b,
	0x76, 0x26, 0xbd, 0x21, 0x0f, 0x20, 0x2e, 0xfd, 0x55, 0xe3, 0xd4, 0xd6, 0xfa, 0xad, 0xa1, 0xe9,
	0xdd, 0x77, 0x76, 0xc6, 0x9d, 0x71, 0xe9, 0x37, 0x44, 0xdc, 0x89, 0x77, 0x21, 0x9a, 0x41, 0xb7,
	0xd6, 0xf8, 0x24, 0xd8, 0xd0, 0x24, 0xcc, 0x6d, 0x5d, 0xaa, 0x9e, 0x5f, 0x53, 0x5d, 0xa9, 0x51,
	0xc3, 0xf7, 0x8a, 0xfb, 0x31, 0x29, 0xb2, 0x0a, 0x9b, 0xf1, 0x65, 0xe6, 0xf0, 0x1d, 0xf4, 0x6e,
	0x12, 0x50, 0x0f, 0x5a, 0x1b, 0x5a, 0xd9, 0x5a, 0x5f, 0x1b, 0x76, 0x70, 0x1d, 0xa2, 0x03, 0xd8,
	0x3b, 0x27, 0xbc, 0xa4, 0xb6, 0x2e, 0x73, 0x0d, 0x78, 0xa3, 0xbf, 0xd2, 0x06, 0x6f, 0xa1, 0x7b,
	0xe1, 0x72, 0x41, 0x0b, 0xf4, 0x02, 0x8c, 0xdb, 0xbb, 0x51, 0x94, 0xc1, 0x6f, 0x0d, 0x3a, 0x17,
	0x59, 0xf4, 0x00, 0x0c, 0x2e, 0xa2, 0x15, 0x0b, 0x65, 0xe5, 0x16, 0xde, 0xe3, 0x22, 0x9a, 0x84,
	0xe8, 0x21, 0x18, 0x69, 0x46, 0xd7, 0xec, 0x87, 0x2a, 0xae, 0x10, 0x72, 0xc1, 0xc8, 0x45, 0x99,
	0x05, 0xd4, 0x6e, 0xc9, 0x4a, 0x8f, 0x2e, 0x2b, 0x2d, 0xb3, 0xba, 0x9f, 0x70, 0x21, 0x8f, 0xb1,
	0xa2, 0xa1, 0x13, 0x30, 0xd3, 0x8c, 0x9d, 0x93, 0x82, 0xae, 0xea, 0xf6, 0xda, 0x7d, 0x6d, 0x68,
	0x7a, 0x07, 0x4e, 0xb3, 0x30, 0x67, 0xb7, 0x30, 0x67, 0x94, 0x54, 0x18, 0x14, 0x71, 0x4a, 0x2b,
	0x74, 0x04, 0x90, 0x96, 0x3e, 0x67, 0x81, 0x54, 0xed, 0x49, 0xd5, 0x3d, 0x47, 0xbd, 0x94, 0xaf,
	0xf2, 0x64, 0x4a, 0x2b, 0xdc, 0x49, 0x77, 0x21, 0x7a, 0x0a, 0x5d, 0xb5, 0x9d, 0x55, 0x42, 0xb6,
	0xd4, 0x36, 0xa4, 0x6f, 0x53, 0xe5, 0xe6, 0x64, 0x4b, 0x07, 0x3f, 0x75, 0xd8, 0xbf, 0xe6, 0x12,
	0x21, 0x68, 0x4b, 0x72, 0x33, 0x75, 0x19, 0x23, 0x0b, 0x74, 0x16, 0xaa, 0xb6, 0x75, 0x16, 0xde,
	0xb0, 0xd2, 0xfa, 0x07, 0x2b, 0x53, 0xb0, 0x62, 0x92, 0xc7, 0x2b, 0xc2, 0x23, 0x91, 0xb1, 0x22,
	0xde, 0xca, 0xb6, 0x2d, 0xef, 0x99, 0xd3, 0x3c, 0xda, 0x0f, 0x2c, 0x62, 0x05, 0xe1, 0xbc, 0x5a,
	0xb0, 0x28, 0xa1, 0xa1, 0x33, 0x26, 0x79, 0x3c, 0xda, 0x71, 0xf1, 0x7e, 0x7c, 0x15, 0xa2, 0x23,
	0x68, 0x6f, 0x58, 0x12, 0xca, 0x19, 0x58, 0xde, 0x93, 0xbf, 0xcc, 0xdb, 0x99, 0xb2, 0x24, 0xc4,
	0x92, 0x39, 0x78, 0x09, 0xed, 0x1a, 0x21, 0x13, 0xfe, 0xff, 0x36, 0x9f, 0xce, 0xbf, 0x7c, 0x9f,
	0xf7, 0xfe, 0x43, 0x16, 0x00, 0xfe, 0x74, 0x76, 0xfa, 0xfa, 0xd4, 0x5b, 0x2c, 0xc7, 0x3d, 0x0d,
	0xdd, 0x05, 0x73, 0x89, 0x27, 0xb3, 0xd9, 0x64, 0x34, 0x5f, 0xcc, 0x70, 0x4f, 0xf7, 0x0d, 0xb9,
	0x8b, 0xe3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x12, 0x2f, 0x77, 0xd0, 0x03, 0x00, 0x00,
}
