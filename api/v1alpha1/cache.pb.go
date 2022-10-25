// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: config/cache/v1alpha1/cache.proto

package v1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Datasource_Kind int32

const (
	Datasource_Database Datasource_Kind = 0
	Datasource_Secret   Datasource_Kind = 1
)

// Enum value maps for Datasource_Kind.
var (
	Datasource_Kind_name = map[int32]string{
		0: "Database",
		1: "Secret",
	}
	Datasource_Kind_value = map[string]int32{
		"Database": 0,
		"Secret":   1,
	}
)

func (x Datasource_Kind) Enum() *Datasource_Kind {
	p := new(Datasource_Kind)
	*p = x
	return p
}

func (x Datasource_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Datasource_Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_config_cache_v1alpha1_cache_proto_enumTypes[0].Descriptor()
}

func (Datasource_Kind) Type() protoreflect.EnumType {
	return &file_config_cache_v1alpha1_cache_proto_enumTypes[0]
}

func (x Datasource_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Datasource_Kind.Descriptor instead.
func (Datasource_Kind) EnumDescriptor() ([]byte, []int) {
	return file_config_cache_v1alpha1_cache_proto_rawDescGZIP(), []int{7, 0}
}

type CacheConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CacheSpec *CacheSpec `protobuf:"bytes,1,opt,name=cacheSpec,proto3" json:"cacheSpec,omitempty"`
}

func (x *CacheConf) Reset() {
	*x = CacheConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_cache_v1alpha1_cache_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheConf) ProtoMessage() {}

func (x *CacheConf) ProtoReflect() protoreflect.Message {
	mi := &file_config_cache_v1alpha1_cache_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheConf.ProtoReflect.Descriptor instead.
func (*CacheConf) Descriptor() ([]byte, []int) {
	return file_config_cache_v1alpha1_cache_proto_rawDescGZIP(), []int{0}
}

func (x *CacheConf) GetCacheSpec() *CacheSpec {
	if x != nil {
		return x.CacheSpec
	}
	return nil
}

// Describes the desired configuration for a Cache. Only DB Cache Service is supported atm
type CacheSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Memory resources profile required for the cache
	// mimic k8s data type
	Memory string `protobuf:"bytes,1,opt,name=memory,proto3" json:"memory,omitempty"`
	// CPU resources profile required for the cache
	Cpu string `protobuf:"bytes,2,opt,name=cpu,proto3" json:"cpu,omitempty"`
	// DatasourceRef or a ServiceBindingRef (TODO clarify)
	DataSourceRef string `protobuf:"bytes,3,opt,name=dataSourceRef,proto3" json:"dataSourceRef,omitempty"`
}

func (x *CacheSpec) Reset() {
	*x = CacheSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_cache_v1alpha1_cache_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheSpec) ProtoMessage() {}

func (x *CacheSpec) ProtoReflect() protoreflect.Message {
	mi := &file_config_cache_v1alpha1_cache_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheSpec.ProtoReflect.Descriptor instead.
func (*CacheSpec) Descriptor() ([]byte, []int) {
	return file_config_cache_v1alpha1_cache_proto_rawDescGZIP(), []int{1}
}

func (x *CacheSpec) GetMemory() string {
	if x != nil {
		return x.Memory
	}
	return ""
}

func (x *CacheSpec) GetCpu() string {
	if x != nil {
		return x.Cpu
	}
	return ""
}

func (x *CacheSpec) GetDataSourceRef() string {
	if x != nil {
		return x.DataSourceRef
	}
	return ""
}

// Describes a data source connection. Only connectionString format is supported atm.
type DataSourceSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Map?
	ConnectionProperties []string `protobuf:"bytes,1,rep,name=connectionProperties,proto3" json:"connectionProperties,omitempty"`
	ServiceBindingRef    string   `protobuf:"bytes,2,opt,name=serviceBindingRef,proto3" json:"serviceBindingRef,omitempty"`
}

func (x *DataSourceSpec) Reset() {
	*x = DataSourceSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_cache_v1alpha1_cache_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSourceSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSourceSpec) ProtoMessage() {}

func (x *DataSourceSpec) ProtoReflect() protoreflect.Message {
	mi := &file_config_cache_v1alpha1_cache_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSourceSpec.ProtoReflect.Descriptor instead.
func (*DataSourceSpec) Descriptor() ([]byte, []int) {
	return file_config_cache_v1alpha1_cache_proto_rawDescGZIP(), []int{2}
}

func (x *DataSourceSpec) GetConnectionProperties() []string {
	if x != nil {
		return x.ConnectionProperties
	}
	return nil
}

func (x *DataSourceSpec) GetServiceBindingRef() string {
	if x != nil {
		return x.ServiceBindingRef
	}
	return ""
}

// Describes a caching rule behaviours
type EagerCachingRuleSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CacheRef *NamespacedRef `protobuf:"bytes,1,opt,name=cacheRef,proto3" json:"cacheRef,omitempty"` // name+ns
}

func (x *EagerCachingRuleSpec) Reset() {
	*x = EagerCachingRuleSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_cache_v1alpha1_cache_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EagerCachingRuleSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EagerCachingRuleSpec) ProtoMessage() {}

func (x *EagerCachingRuleSpec) ProtoReflect() protoreflect.Message {
	mi := &file_config_cache_v1alpha1_cache_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EagerCachingRuleSpec.ProtoReflect.Descriptor instead.
func (*EagerCachingRuleSpec) Descriptor() ([]byte, []int) {
	return file_config_cache_v1alpha1_cache_proto_rawDescGZIP(), []int{3}
}

func (x *EagerCachingRuleSpec) GetCacheRef() *NamespacedRef {
	if x != nil {
		return x.CacheRef
	}
	return nil
}

// Describes a caching rule behaviours
type LazyCachingRuleSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CacheRef *NamespacedRef `protobuf:"bytes,1,opt,name=cacheRef,proto3" json:"cacheRef,omitempty"` // name+ns
}

func (x *LazyCachingRuleSpec) Reset() {
	*x = LazyCachingRuleSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_cache_v1alpha1_cache_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LazyCachingRuleSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LazyCachingRuleSpec) ProtoMessage() {}

func (x *LazyCachingRuleSpec) ProtoReflect() protoreflect.Message {
	mi := &file_config_cache_v1alpha1_cache_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LazyCachingRuleSpec.ProtoReflect.Descriptor instead.
func (*LazyCachingRuleSpec) Descriptor() ([]byte, []int) {
	return file_config_cache_v1alpha1_cache_proto_rawDescGZIP(), []int{4}
}

func (x *LazyCachingRuleSpec) GetCacheRef() *NamespacedRef {
	if x != nil {
		return x.CacheRef
	}
	return nil
}

type NamespacedRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *NamespacedRef) Reset() {
	*x = NamespacedRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_cache_v1alpha1_cache_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespacedRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespacedRef) ProtoMessage() {}

func (x *NamespacedRef) ProtoReflect() protoreflect.Message {
	mi := &file_config_cache_v1alpha1_cache_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespacedRef.ProtoReflect.Descriptor instead.
func (*NamespacedRef) Descriptor() ([]byte, []int) {
	return file_config_cache_v1alpha1_cache_proto_rawDescGZIP(), []int{5}
}

func (x *NamespacedRef) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NamespacedRef) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

// Describes how the cache entry (key,value) are produced by the cache. Only Infinispan
// is supported atm. The existence of a proto definition shared between application and cache is
// implied.
type DataProducer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DataProducer) Reset() {
	*x = DataProducer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_cache_v1alpha1_cache_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataProducer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataProducer) ProtoMessage() {}

func (x *DataProducer) ProtoReflect() protoreflect.Message {
	mi := &file_config_cache_v1alpha1_cache_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataProducer.ProtoReflect.Descriptor instead.
func (*DataProducer) Descriptor() ([]byte, []int) {
	return file_config_cache_v1alpha1_cache_proto_rawDescGZIP(), []int{6}
}

func (x *DataProducer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Datasource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Group string          `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	Kind  Datasource_Kind `protobuf:"varint,3,opt,name=kind,proto3,enum=gingersnap.config.cache.v1alpha1.Datasource_Kind" json:"kind,omitempty"`
	Id    string          `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Datasource) Reset() {
	*x = Datasource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_cache_v1alpha1_cache_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Datasource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Datasource) ProtoMessage() {}

func (x *Datasource) ProtoReflect() protoreflect.Message {
	mi := &file_config_cache_v1alpha1_cache_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Datasource.ProtoReflect.Descriptor instead.
func (*Datasource) Descriptor() ([]byte, []int) {
	return file_config_cache_v1alpha1_cache_proto_rawDescGZIP(), []int{7}
}

func (x *Datasource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Datasource) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Datasource) GetKind() Datasource_Kind {
	if x != nil {
		return x.Kind
	}
	return Datasource_Database
}

func (x *Datasource) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_config_cache_v1alpha1_cache_proto protoreflect.FileDescriptor

var file_config_cache_v1alpha1_cache_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x20, 0x67, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x73, 0x6e, 0x61, 0x70, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x56, 0x0a, 0x09, 0x43, 0x61, 0x63, 0x68, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x12, 0x49, 0x0a, 0x09, 0x63, 0x61, 0x63, 0x68, 0x65, 0x53, 0x70, 0x65, 0x63, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x73, 0x6e,
	0x61, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x70,
	0x65, 0x63, 0x52, 0x09, 0x63, 0x61, 0x63, 0x68, 0x65, 0x53, 0x70, 0x65, 0x63, 0x22, 0x5b, 0x0a,
	0x09, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x63, 0x70, 0x75, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x65, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x61, 0x74,
	0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x22, 0x72, 0x0a, 0x0e, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x32, 0x0a, 0x14,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x14, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x12, 0x2c, 0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x66, 0x22, 0x63,
	0x0a, 0x14, 0x45, 0x61, 0x67, 0x65, 0x72, 0x43, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x75,
	0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x4b, 0x0a, 0x08, 0x63, 0x61, 0x63, 0x68, 0x65, 0x52,
	0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x69, 0x6e, 0x67, 0x65,
	0x72, 0x73, 0x6e, 0x61, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x64, 0x52, 0x65, 0x66, 0x52, 0x08, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x52, 0x65, 0x66, 0x22, 0x62, 0x0a, 0x13, 0x4c, 0x61, 0x7a, 0x79, 0x43, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x4b, 0x0a, 0x08, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x52, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67,
	0x69, 0x6e, 0x67, 0x65, 0x72, 0x73, 0x6e, 0x61, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x64, 0x52, 0x65, 0x66, 0x52, 0x08, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x66, 0x22, 0x41, 0x0a, 0x0d, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x64, 0x52, 0x65, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x22, 0x0a, 0x0c, 0x44, 0x61,
	0x74, 0x61, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xaf,
	0x01, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x45, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x67, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x73, 0x6e,
	0x61, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20,
	0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x10, 0x01,
	0x42, 0x10, 0x50, 0x01, 0x5a, 0x0c, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_cache_v1alpha1_cache_proto_rawDescOnce sync.Once
	file_config_cache_v1alpha1_cache_proto_rawDescData = file_config_cache_v1alpha1_cache_proto_rawDesc
)

func file_config_cache_v1alpha1_cache_proto_rawDescGZIP() []byte {
	file_config_cache_v1alpha1_cache_proto_rawDescOnce.Do(func() {
		file_config_cache_v1alpha1_cache_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_cache_v1alpha1_cache_proto_rawDescData)
	})
	return file_config_cache_v1alpha1_cache_proto_rawDescData
}

var file_config_cache_v1alpha1_cache_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_config_cache_v1alpha1_cache_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_config_cache_v1alpha1_cache_proto_goTypes = []interface{}{
	(Datasource_Kind)(0),         // 0: gingersnap.config.cache.v1alpha1.Datasource.Kind
	(*CacheConf)(nil),            // 1: gingersnap.config.cache.v1alpha1.CacheConf
	(*CacheSpec)(nil),            // 2: gingersnap.config.cache.v1alpha1.CacheSpec
	(*DataSourceSpec)(nil),       // 3: gingersnap.config.cache.v1alpha1.DataSourceSpec
	(*EagerCachingRuleSpec)(nil), // 4: gingersnap.config.cache.v1alpha1.EagerCachingRuleSpec
	(*LazyCachingRuleSpec)(nil),  // 5: gingersnap.config.cache.v1alpha1.LazyCachingRuleSpec
	(*NamespacedRef)(nil),        // 6: gingersnap.config.cache.v1alpha1.NamespacedRef
	(*DataProducer)(nil),         // 7: gingersnap.config.cache.v1alpha1.DataProducer
	(*Datasource)(nil),           // 8: gingersnap.config.cache.v1alpha1.Datasource
}
var file_config_cache_v1alpha1_cache_proto_depIdxs = []int32{
	2, // 0: gingersnap.config.cache.v1alpha1.CacheConf.cacheSpec:type_name -> gingersnap.config.cache.v1alpha1.CacheSpec
	6, // 1: gingersnap.config.cache.v1alpha1.EagerCachingRuleSpec.cacheRef:type_name -> gingersnap.config.cache.v1alpha1.NamespacedRef
	6, // 2: gingersnap.config.cache.v1alpha1.LazyCachingRuleSpec.cacheRef:type_name -> gingersnap.config.cache.v1alpha1.NamespacedRef
	0, // 3: gingersnap.config.cache.v1alpha1.Datasource.kind:type_name -> gingersnap.config.cache.v1alpha1.Datasource.Kind
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_config_cache_v1alpha1_cache_proto_init() }
func file_config_cache_v1alpha1_cache_proto_init() {
	if File_config_cache_v1alpha1_cache_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_cache_v1alpha1_cache_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CacheConf); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_cache_v1alpha1_cache_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CacheSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_cache_v1alpha1_cache_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSourceSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_cache_v1alpha1_cache_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EagerCachingRuleSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_cache_v1alpha1_cache_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LazyCachingRuleSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_cache_v1alpha1_cache_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespacedRef); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_cache_v1alpha1_cache_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataProducer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_cache_v1alpha1_cache_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Datasource); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_cache_v1alpha1_cache_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_cache_v1alpha1_cache_proto_goTypes,
		DependencyIndexes: file_config_cache_v1alpha1_cache_proto_depIdxs,
		EnumInfos:         file_config_cache_v1alpha1_cache_proto_enumTypes,
		MessageInfos:      file_config_cache_v1alpha1_cache_proto_msgTypes,
	}.Build()
	File_config_cache_v1alpha1_cache_proto = out.File
	file_config_cache_v1alpha1_cache_proto_rawDesc = nil
	file_config_cache_v1alpha1_cache_proto_goTypes = nil
	file_config_cache_v1alpha1_cache_proto_depIdxs = nil
}
