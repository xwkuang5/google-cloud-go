// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: google/cloud/aiplatform/v1/index.proto

package aiplatformpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The update method of an Index.
type Index_IndexUpdateMethod int32

const (
	// Should not be used.
	Index_INDEX_UPDATE_METHOD_UNSPECIFIED Index_IndexUpdateMethod = 0
	// BatchUpdate: user can call UpdateIndex with files on Cloud Storage of
	// datapoints to update.
	Index_BATCH_UPDATE Index_IndexUpdateMethod = 1
	// StreamUpdate: user can call UpsertDatapoints/DeleteDatapoints to update
	// the Index and the updates will be applied in corresponding
	// DeployedIndexes in nearly real-time.
	Index_STREAM_UPDATE Index_IndexUpdateMethod = 2
)

// Enum value maps for Index_IndexUpdateMethod.
var (
	Index_IndexUpdateMethod_name = map[int32]string{
		0: "INDEX_UPDATE_METHOD_UNSPECIFIED",
		1: "BATCH_UPDATE",
		2: "STREAM_UPDATE",
	}
	Index_IndexUpdateMethod_value = map[string]int32{
		"INDEX_UPDATE_METHOD_UNSPECIFIED": 0,
		"BATCH_UPDATE":                    1,
		"STREAM_UPDATE":                   2,
	}
)

func (x Index_IndexUpdateMethod) Enum() *Index_IndexUpdateMethod {
	p := new(Index_IndexUpdateMethod)
	*p = x
	return p
}

func (x Index_IndexUpdateMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Index_IndexUpdateMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_aiplatform_v1_index_proto_enumTypes[0].Descriptor()
}

func (Index_IndexUpdateMethod) Type() protoreflect.EnumType {
	return &file_google_cloud_aiplatform_v1_index_proto_enumTypes[0]
}

func (x Index_IndexUpdateMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Index_IndexUpdateMethod.Descriptor instead.
func (Index_IndexUpdateMethod) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_index_proto_rawDescGZIP(), []int{0, 0}
}

// A representation of a collection of database items organized in a way that
// allows for approximate nearest neighbor (a.k.a ANN) algorithms search.
type Index struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the Index.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The display name of the Index.
	// The name can be up to 128 characters long and can be consist of any UTF-8
	// characters.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The description of the Index.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Immutable. Points to a YAML file stored on Google Cloud Storage describing additional
	// information about the Index, that is specific to it. Unset if the Index
	// does not have any additional information.
	// The schema is defined as an OpenAPI 3.0.2 [Schema
	// Object](https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.2.md#schemaObject).
	// Note: The URI given on output will be immutable and probably different,
	// including the URI scheme, than the one given on input. The output URI will
	// point to a location where the user only has a read access.
	MetadataSchemaUri string `protobuf:"bytes,4,opt,name=metadata_schema_uri,json=metadataSchemaUri,proto3" json:"metadata_schema_uri,omitempty"`
	// An additional information about the Index; the schema of the metadata can
	// be found in [metadata_schema][google.cloud.aiplatform.v1.Index.metadata_schema_uri].
	Metadata *structpb.Value `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Output only. The pointers to DeployedIndexes created from this Index.
	// An Index can be only deleted if all its DeployedIndexes had been undeployed
	// first.
	DeployedIndexes []*DeployedIndexRef `protobuf:"bytes,7,rep,name=deployed_indexes,json=deployedIndexes,proto3" json:"deployed_indexes,omitempty"`
	// Used to perform consistent read-modify-write updates. If not set, a blind
	// "overwrite" update happens.
	Etag string `protobuf:"bytes,8,opt,name=etag,proto3" json:"etag,omitempty"`
	// The labels with user-defined metadata to organize your Indexes.
	//
	// Label keys and values can be no longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	//
	// See https://goo.gl/xmQnxf for more information and examples of labels.
	Labels map[string]string `protobuf:"bytes,9,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Output only. Timestamp when this Index was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Timestamp when this Index was most recently updated.
	// This also includes any update to the contents of the Index.
	// Note that Operations working on this Index may have their
	// [Operations.metadata.generic_metadata.update_time]
	// [google.cloud.aiplatform.v1.GenericOperationMetadata.update_time] a little after the value of this
	// timestamp, yet that does not mean their results are not already reflected
	// in the Index. Result of any successfully completed Operation on the Index
	// is reflected in it.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. Stats of the index resource.
	IndexStats *IndexStats `protobuf:"bytes,14,opt,name=index_stats,json=indexStats,proto3" json:"index_stats,omitempty"`
	// Immutable. The update method to use with this Index. If not set, BATCH_UPDATE will be
	// used by default.
	IndexUpdateMethod Index_IndexUpdateMethod `protobuf:"varint,16,opt,name=index_update_method,json=indexUpdateMethod,proto3,enum=google.cloud.aiplatform.v1.Index_IndexUpdateMethod" json:"index_update_method,omitempty"`
}

func (x *Index) Reset() {
	*x = Index{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_index_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Index) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Index) ProtoMessage() {}

func (x *Index) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_index_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Index.ProtoReflect.Descriptor instead.
func (*Index) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_index_proto_rawDescGZIP(), []int{0}
}

func (x *Index) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Index) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Index) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Index) GetMetadataSchemaUri() string {
	if x != nil {
		return x.MetadataSchemaUri
	}
	return ""
}

func (x *Index) GetMetadata() *structpb.Value {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Index) GetDeployedIndexes() []*DeployedIndexRef {
	if x != nil {
		return x.DeployedIndexes
	}
	return nil
}

func (x *Index) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *Index) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Index) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Index) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Index) GetIndexStats() *IndexStats {
	if x != nil {
		return x.IndexStats
	}
	return nil
}

func (x *Index) GetIndexUpdateMethod() Index_IndexUpdateMethod {
	if x != nil {
		return x.IndexUpdateMethod
	}
	return Index_INDEX_UPDATE_METHOD_UNSPECIFIED
}

// A datapoint of Index.
type IndexDatapoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Unique identifier of the datapoint.
	DatapointId string `protobuf:"bytes,1,opt,name=datapoint_id,json=datapointId,proto3" json:"datapoint_id,omitempty"`
	// Required. Feature embedding vector. An array of numbers with the length of
	// [NearestNeighborSearchConfig.dimensions].
	FeatureVector []float32 `protobuf:"fixed32,2,rep,packed,name=feature_vector,json=featureVector,proto3" json:"feature_vector,omitempty"`
	// Optional. List of Restrict of the datapoint, used to perform "restricted searches"
	// where boolean rule are used to filter the subset of the database eligible
	// for matching.
	// See: https://cloud.google.com/vertex-ai/docs/matching-engine/filtering
	Restricts []*IndexDatapoint_Restriction `protobuf:"bytes,4,rep,name=restricts,proto3" json:"restricts,omitempty"`
	// Optional. CrowdingTag of the datapoint, the number of neighbors to return in each
	// crowding can be configured during query.
	CrowdingTag *IndexDatapoint_CrowdingTag `protobuf:"bytes,5,opt,name=crowding_tag,json=crowdingTag,proto3" json:"crowding_tag,omitempty"`
}

func (x *IndexDatapoint) Reset() {
	*x = IndexDatapoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_index_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexDatapoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexDatapoint) ProtoMessage() {}

func (x *IndexDatapoint) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_index_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexDatapoint.ProtoReflect.Descriptor instead.
func (*IndexDatapoint) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_index_proto_rawDescGZIP(), []int{1}
}

func (x *IndexDatapoint) GetDatapointId() string {
	if x != nil {
		return x.DatapointId
	}
	return ""
}

func (x *IndexDatapoint) GetFeatureVector() []float32 {
	if x != nil {
		return x.FeatureVector
	}
	return nil
}

func (x *IndexDatapoint) GetRestricts() []*IndexDatapoint_Restriction {
	if x != nil {
		return x.Restricts
	}
	return nil
}

func (x *IndexDatapoint) GetCrowdingTag() *IndexDatapoint_CrowdingTag {
	if x != nil {
		return x.CrowdingTag
	}
	return nil
}

// Stats of the Index.
type IndexStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The number of vectors in the Index.
	VectorsCount int64 `protobuf:"varint,1,opt,name=vectors_count,json=vectorsCount,proto3" json:"vectors_count,omitempty"`
	// Output only. The number of shards in the Index.
	ShardsCount int32 `protobuf:"varint,2,opt,name=shards_count,json=shardsCount,proto3" json:"shards_count,omitempty"`
}

func (x *IndexStats) Reset() {
	*x = IndexStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_index_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexStats) ProtoMessage() {}

func (x *IndexStats) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_index_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexStats.ProtoReflect.Descriptor instead.
func (*IndexStats) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_index_proto_rawDescGZIP(), []int{2}
}

func (x *IndexStats) GetVectorsCount() int64 {
	if x != nil {
		return x.VectorsCount
	}
	return 0
}

func (x *IndexStats) GetShardsCount() int32 {
	if x != nil {
		return x.ShardsCount
	}
	return 0
}

// Restriction of a datapoint which describe its attributes(tokens) from each
// of several attribute categories(namespaces).
type IndexDatapoint_Restriction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The namespace of this restriction. eg: color.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// The attributes to allow in this namespace. eg: 'red'
	AllowList []string `protobuf:"bytes,2,rep,name=allow_list,json=allowList,proto3" json:"allow_list,omitempty"`
	// The attributes to deny in this namespace. eg: 'blue'
	DenyList []string `protobuf:"bytes,3,rep,name=deny_list,json=denyList,proto3" json:"deny_list,omitempty"`
}

func (x *IndexDatapoint_Restriction) Reset() {
	*x = IndexDatapoint_Restriction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_index_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexDatapoint_Restriction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexDatapoint_Restriction) ProtoMessage() {}

func (x *IndexDatapoint_Restriction) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_index_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexDatapoint_Restriction.ProtoReflect.Descriptor instead.
func (*IndexDatapoint_Restriction) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_index_proto_rawDescGZIP(), []int{1, 0}
}

func (x *IndexDatapoint_Restriction) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *IndexDatapoint_Restriction) GetAllowList() []string {
	if x != nil {
		return x.AllowList
	}
	return nil
}

func (x *IndexDatapoint_Restriction) GetDenyList() []string {
	if x != nil {
		return x.DenyList
	}
	return nil
}

// Crowding tag is a constraint on a neighbor list produced by nearest
// neighbor search requiring that no more than some value k' of the k
// neighbors returned have the same value of crowding_attribute.
type IndexDatapoint_CrowdingTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The attribute value used for crowding.  The maximum number of neighbors
	// to return per crowding attribute value
	// (per_crowding_attribute_num_neighbors) is configured per-query. This
	// field is ignored if per_crowding_attribute_num_neighbors is larger than
	// the total number of neighbors to return for a given query.
	CrowdingAttribute string `protobuf:"bytes,1,opt,name=crowding_attribute,json=crowdingAttribute,proto3" json:"crowding_attribute,omitempty"`
}

func (x *IndexDatapoint_CrowdingTag) Reset() {
	*x = IndexDatapoint_CrowdingTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_index_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexDatapoint_CrowdingTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexDatapoint_CrowdingTag) ProtoMessage() {}

func (x *IndexDatapoint_CrowdingTag) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_index_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexDatapoint_CrowdingTag.ProtoReflect.Descriptor instead.
func (*IndexDatapoint_CrowdingTag) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_index_proto_rawDescGZIP(), []int{1, 1}
}

func (x *IndexDatapoint_CrowdingTag) GetCrowdingAttribute() string {
	if x != nil {
		return x.CrowdingAttribute
	}
	return ""
}

var File_google_cloud_aiplatform_v1_index_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1_index_proto_rawDesc = []byte{
	0x0a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x72, 0x65, 0x66, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x07, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x17,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x33, 0x0a, 0x13, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x05, 0x52, 0x11, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x55, 0x72, 0x69, 0x12, 0x32, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x5c, 0x0a, 0x10, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52,
	0x65, 0x66, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x12, 0x45, 0x0a, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x68, 0x0a, 0x13, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x11, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x1a,
	0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5d, 0x0a, 0x11, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x23, 0x0a, 0x1f, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f,
	0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x42, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x55, 0x50,
	0x44, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d,
	0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x3a, 0x5d, 0xea, 0x41, 0x5a, 0x0a, 0x1f,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x37, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73,
	0x2f, 0x7b, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x7d, 0x22, 0xc6, 0x03, 0x0a, 0x0e, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0c, 0x64,
	0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x0e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x76,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x02, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x0d, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12,
	0x59, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x52,
	0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x09, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x12, 0x5e, 0x0a, 0x0c, 0x63, 0x72,
	0x6f, 0x77, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x6f,
	0x77, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x67, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b, 0x63,
	0x72, 0x6f, 0x77, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x67, 0x1a, 0x67, 0x0a, 0x0b, 0x52, 0x65,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x6e, 0x79, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x6e, 0x79, 0x4c,
	0x69, 0x73, 0x74, 0x1a, 0x3c, 0x0a, 0x0b, 0x43, 0x72, 0x6f, 0x77, 0x64, 0x69, 0x6e, 0x67, 0x54,
	0x61, 0x67, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x72, 0x6f, 0x77, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x63, 0x72, 0x6f, 0x77, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x22, 0x5e, 0x0a, 0x0a, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12,
	0x28, 0x0a, 0x0d, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0c, 0x76, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0c, 0x73, 0x68, 0x61,
	0x72, 0x64, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x42, 0xce, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x44, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c,
	0x56, 0x31, 0xea, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1_index_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1_index_proto_rawDescData = file_google_cloud_aiplatform_v1_index_proto_rawDesc
)

func file_google_cloud_aiplatform_v1_index_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1_index_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1_index_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1_index_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1_index_proto_rawDescData
}

var file_google_cloud_aiplatform_v1_index_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_aiplatform_v1_index_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_aiplatform_v1_index_proto_goTypes = []interface{}{
	(Index_IndexUpdateMethod)(0),       // 0: google.cloud.aiplatform.v1.Index.IndexUpdateMethod
	(*Index)(nil),                      // 1: google.cloud.aiplatform.v1.Index
	(*IndexDatapoint)(nil),             // 2: google.cloud.aiplatform.v1.IndexDatapoint
	(*IndexStats)(nil),                 // 3: google.cloud.aiplatform.v1.IndexStats
	nil,                                // 4: google.cloud.aiplatform.v1.Index.LabelsEntry
	(*IndexDatapoint_Restriction)(nil), // 5: google.cloud.aiplatform.v1.IndexDatapoint.Restriction
	(*IndexDatapoint_CrowdingTag)(nil), // 6: google.cloud.aiplatform.v1.IndexDatapoint.CrowdingTag
	(*structpb.Value)(nil),             // 7: google.protobuf.Value
	(*DeployedIndexRef)(nil),           // 8: google.cloud.aiplatform.v1.DeployedIndexRef
	(*timestamppb.Timestamp)(nil),      // 9: google.protobuf.Timestamp
}
var file_google_cloud_aiplatform_v1_index_proto_depIdxs = []int32{
	7, // 0: google.cloud.aiplatform.v1.Index.metadata:type_name -> google.protobuf.Value
	8, // 1: google.cloud.aiplatform.v1.Index.deployed_indexes:type_name -> google.cloud.aiplatform.v1.DeployedIndexRef
	4, // 2: google.cloud.aiplatform.v1.Index.labels:type_name -> google.cloud.aiplatform.v1.Index.LabelsEntry
	9, // 3: google.cloud.aiplatform.v1.Index.create_time:type_name -> google.protobuf.Timestamp
	9, // 4: google.cloud.aiplatform.v1.Index.update_time:type_name -> google.protobuf.Timestamp
	3, // 5: google.cloud.aiplatform.v1.Index.index_stats:type_name -> google.cloud.aiplatform.v1.IndexStats
	0, // 6: google.cloud.aiplatform.v1.Index.index_update_method:type_name -> google.cloud.aiplatform.v1.Index.IndexUpdateMethod
	5, // 7: google.cloud.aiplatform.v1.IndexDatapoint.restricts:type_name -> google.cloud.aiplatform.v1.IndexDatapoint.Restriction
	6, // 8: google.cloud.aiplatform.v1.IndexDatapoint.crowding_tag:type_name -> google.cloud.aiplatform.v1.IndexDatapoint.CrowdingTag
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1_index_proto_init() }
func file_google_cloud_aiplatform_v1_index_proto_init() {
	if File_google_cloud_aiplatform_v1_index_proto != nil {
		return
	}
	file_google_cloud_aiplatform_v1_deployed_index_ref_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1_index_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Index); i {
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
		file_google_cloud_aiplatform_v1_index_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexDatapoint); i {
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
		file_google_cloud_aiplatform_v1_index_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexStats); i {
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
		file_google_cloud_aiplatform_v1_index_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexDatapoint_Restriction); i {
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
		file_google_cloud_aiplatform_v1_index_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexDatapoint_CrowdingTag); i {
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
			RawDescriptor: file_google_cloud_aiplatform_v1_index_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1_index_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1_index_proto_depIdxs,
		EnumInfos:         file_google_cloud_aiplatform_v1_index_proto_enumTypes,
		MessageInfos:      file_google_cloud_aiplatform_v1_index_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1_index_proto = out.File
	file_google_cloud_aiplatform_v1_index_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1_index_proto_goTypes = nil
	file_google_cloud_aiplatform_v1_index_proto_depIdxs = nil
}
