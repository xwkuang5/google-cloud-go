// Copyright 2023 Google LLC
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
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: google/cloud/datacatalog/v1/dataplex_spec.proto

package datacatalogpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Common Dataplex fields.
type DataplexSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Fully qualified resource name of an asset in Dataplex, to which the
	// underlying data source (Cloud Storage bucket or BigQuery dataset) of the
	// entity is attached.
	Asset string `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	// Format of the data.
	DataFormat *PhysicalSchema `protobuf:"bytes,2,opt,name=data_format,json=dataFormat,proto3" json:"data_format,omitempty"`
	// Compression format of the data, e.g., zip, gzip etc.
	CompressionFormat string `protobuf:"bytes,3,opt,name=compression_format,json=compressionFormat,proto3" json:"compression_format,omitempty"`
	// Project ID of the underlying Cloud Storage or BigQuery data. Note that
	// this may not be the same project as the correspondingly Dataplex lake /
	// zone / asset.
	ProjectId string `protobuf:"bytes,4,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *DataplexSpec) Reset() {
	*x = DataplexSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datacatalog_v1_dataplex_spec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataplexSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataplexSpec) ProtoMessage() {}

func (x *DataplexSpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datacatalog_v1_dataplex_spec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataplexSpec.ProtoReflect.Descriptor instead.
func (*DataplexSpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_datacatalog_v1_dataplex_spec_proto_rawDescGZIP(), []int{0}
}

func (x *DataplexSpec) GetAsset() string {
	if x != nil {
		return x.Asset
	}
	return ""
}

func (x *DataplexSpec) GetDataFormat() *PhysicalSchema {
	if x != nil {
		return x.DataFormat
	}
	return nil
}

func (x *DataplexSpec) GetCompressionFormat() string {
	if x != nil {
		return x.CompressionFormat
	}
	return ""
}

func (x *DataplexSpec) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

// Entry specyfication for a Dataplex fileset.
type DataplexFilesetSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Common Dataplex fields.
	DataplexSpec *DataplexSpec `protobuf:"bytes,1,opt,name=dataplex_spec,json=dataplexSpec,proto3" json:"dataplex_spec,omitempty"`
}

func (x *DataplexFilesetSpec) Reset() {
	*x = DataplexFilesetSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datacatalog_v1_dataplex_spec_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataplexFilesetSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataplexFilesetSpec) ProtoMessage() {}

func (x *DataplexFilesetSpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datacatalog_v1_dataplex_spec_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataplexFilesetSpec.ProtoReflect.Descriptor instead.
func (*DataplexFilesetSpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_datacatalog_v1_dataplex_spec_proto_rawDescGZIP(), []int{1}
}

func (x *DataplexFilesetSpec) GetDataplexSpec() *DataplexSpec {
	if x != nil {
		return x.DataplexSpec
	}
	return nil
}

// Entry specification for a Dataplex table.
type DataplexTableSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of external tables registered by Dataplex in other systems based on
	// the same underlying data.
	//
	// External tables allow to query this data in those systems.
	ExternalTables []*DataplexExternalTable `protobuf:"bytes,1,rep,name=external_tables,json=externalTables,proto3" json:"external_tables,omitempty"`
	// Common Dataplex fields.
	DataplexSpec *DataplexSpec `protobuf:"bytes,2,opt,name=dataplex_spec,json=dataplexSpec,proto3" json:"dataplex_spec,omitempty"`
	// Indicates if the table schema is managed by the user or not.
	UserManaged bool `protobuf:"varint,3,opt,name=user_managed,json=userManaged,proto3" json:"user_managed,omitempty"`
}

func (x *DataplexTableSpec) Reset() {
	*x = DataplexTableSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datacatalog_v1_dataplex_spec_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataplexTableSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataplexTableSpec) ProtoMessage() {}

func (x *DataplexTableSpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datacatalog_v1_dataplex_spec_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataplexTableSpec.ProtoReflect.Descriptor instead.
func (*DataplexTableSpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_datacatalog_v1_dataplex_spec_proto_rawDescGZIP(), []int{2}
}

func (x *DataplexTableSpec) GetExternalTables() []*DataplexExternalTable {
	if x != nil {
		return x.ExternalTables
	}
	return nil
}

func (x *DataplexTableSpec) GetDataplexSpec() *DataplexSpec {
	if x != nil {
		return x.DataplexSpec
	}
	return nil
}

func (x *DataplexTableSpec) GetUserManaged() bool {
	if x != nil {
		return x.UserManaged
	}
	return false
}

// External table registered by Dataplex.
// Dataplex publishes data discovered from an asset into multiple other systems
// (BigQuery, DPMS) in form of tables. We call them "external tables". External
// tables are also synced into the Data Catalog.
// This message contains pointers to
// those external tables (fully qualified name, resource name et cetera) within
// the Data Catalog.
type DataplexExternalTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Service in which the external table is registered.
	System IntegratedSystem `protobuf:"varint,1,opt,name=system,proto3,enum=google.cloud.datacatalog.v1.IntegratedSystem" json:"system,omitempty"`
	// Fully qualified name (FQN) of the external table.
	FullyQualifiedName string `protobuf:"bytes,28,opt,name=fully_qualified_name,json=fullyQualifiedName,proto3" json:"fully_qualified_name,omitempty"`
	// Google Cloud resource name of the external table.
	GoogleCloudResource string `protobuf:"bytes,3,opt,name=google_cloud_resource,json=googleCloudResource,proto3" json:"google_cloud_resource,omitempty"`
	// Name of the Data Catalog entry representing the external table.
	DataCatalogEntry string `protobuf:"bytes,4,opt,name=data_catalog_entry,json=dataCatalogEntry,proto3" json:"data_catalog_entry,omitempty"`
}

func (x *DataplexExternalTable) Reset() {
	*x = DataplexExternalTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datacatalog_v1_dataplex_spec_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataplexExternalTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataplexExternalTable) ProtoMessage() {}

func (x *DataplexExternalTable) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datacatalog_v1_dataplex_spec_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataplexExternalTable.ProtoReflect.Descriptor instead.
func (*DataplexExternalTable) Descriptor() ([]byte, []int) {
	return file_google_cloud_datacatalog_v1_dataplex_spec_proto_rawDescGZIP(), []int{3}
}

func (x *DataplexExternalTable) GetSystem() IntegratedSystem {
	if x != nil {
		return x.System
	}
	return IntegratedSystem_INTEGRATED_SYSTEM_UNSPECIFIED
}

func (x *DataplexExternalTable) GetFullyQualifiedName() string {
	if x != nil {
		return x.FullyQualifiedName
	}
	return ""
}

func (x *DataplexExternalTable) GetGoogleCloudResource() string {
	if x != nil {
		return x.GoogleCloudResource
	}
	return ""
}

func (x *DataplexExternalTable) GetDataCatalogEntry() string {
	if x != nil {
		return x.DataCatalogEntry
	}
	return ""
}

var File_google_cloud_datacatalog_v1_dataplex_spec_proto protoreflect.FileDescriptor

var file_google_cloud_datacatalog_v1_dataplex_spec_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x28,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01, 0x0a, 0x0c,
	0x44, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x78, 0x53, 0x70, 0x65, 0x63, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x12, 0x4c, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x12, 0x2d, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x65,
	0x0a, 0x13, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x78, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x65,
	0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x4e, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x65,
	0x78, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x70,
	0x6c, 0x65, 0x78, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x65,
	0x78, 0x53, 0x70, 0x65, 0x63, 0x22, 0xe3, 0x01, 0x0a, 0x11, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6c,
	0x65, 0x78, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x5b, 0x0a, 0x0f, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x78, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x4e, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61,
	0x70, 0x6c, 0x65, 0x78, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x70, 0x6c, 0x65, 0x78, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61,
	0x70, 0x6c, 0x65, 0x78, 0x53, 0x70, 0x65, 0x63, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x75, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x22, 0xf2, 0x01, 0x0a, 0x15,
	0x44, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x78, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x65, 0x64, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x30, 0x0a, 0x14,
	0x66, 0x75, 0x6c, 0x6c, 0x79, 0x5f, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x66, 0x75, 0x6c, 0x6c,
	0x79, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x32,
	0x0a, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x64, 0x61, 0x74, 0x61, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x42, 0xd9, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x78, 0x53, 0x70,
	0x65, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x70, 0x62, 0x3b, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x70, 0x62, 0xf8, 0x01, 0x01, 0xaa,
	0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1b,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x61, 0x74,
	0x61, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x61, 0x74,
	0x61, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_datacatalog_v1_dataplex_spec_proto_rawDescOnce sync.Once
	file_google_cloud_datacatalog_v1_dataplex_spec_proto_rawDescData = file_google_cloud_datacatalog_v1_dataplex_spec_proto_rawDesc
)

func file_google_cloud_datacatalog_v1_dataplex_spec_proto_rawDescGZIP() []byte {
	file_google_cloud_datacatalog_v1_dataplex_spec_proto_rawDescOnce.Do(func() {
		file_google_cloud_datacatalog_v1_dataplex_spec_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_datacatalog_v1_dataplex_spec_proto_rawDescData)
	})
	return file_google_cloud_datacatalog_v1_dataplex_spec_proto_rawDescData
}

var file_google_cloud_datacatalog_v1_dataplex_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_datacatalog_v1_dataplex_spec_proto_goTypes = []interface{}{
	(*DataplexSpec)(nil),          // 0: google.cloud.datacatalog.v1.DataplexSpec
	(*DataplexFilesetSpec)(nil),   // 1: google.cloud.datacatalog.v1.DataplexFilesetSpec
	(*DataplexTableSpec)(nil),     // 2: google.cloud.datacatalog.v1.DataplexTableSpec
	(*DataplexExternalTable)(nil), // 3: google.cloud.datacatalog.v1.DataplexExternalTable
	(*PhysicalSchema)(nil),        // 4: google.cloud.datacatalog.v1.PhysicalSchema
	(IntegratedSystem)(0),         // 5: google.cloud.datacatalog.v1.IntegratedSystem
}
var file_google_cloud_datacatalog_v1_dataplex_spec_proto_depIdxs = []int32{
	4, // 0: google.cloud.datacatalog.v1.DataplexSpec.data_format:type_name -> google.cloud.datacatalog.v1.PhysicalSchema
	0, // 1: google.cloud.datacatalog.v1.DataplexFilesetSpec.dataplex_spec:type_name -> google.cloud.datacatalog.v1.DataplexSpec
	3, // 2: google.cloud.datacatalog.v1.DataplexTableSpec.external_tables:type_name -> google.cloud.datacatalog.v1.DataplexExternalTable
	0, // 3: google.cloud.datacatalog.v1.DataplexTableSpec.dataplex_spec:type_name -> google.cloud.datacatalog.v1.DataplexSpec
	5, // 4: google.cloud.datacatalog.v1.DataplexExternalTable.system:type_name -> google.cloud.datacatalog.v1.IntegratedSystem
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_datacatalog_v1_dataplex_spec_proto_init() }
func file_google_cloud_datacatalog_v1_dataplex_spec_proto_init() {
	if File_google_cloud_datacatalog_v1_dataplex_spec_proto != nil {
		return
	}
	file_google_cloud_datacatalog_v1_common_proto_init()
	file_google_cloud_datacatalog_v1_physical_schema_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_datacatalog_v1_dataplex_spec_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataplexSpec); i {
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
		file_google_cloud_datacatalog_v1_dataplex_spec_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataplexFilesetSpec); i {
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
		file_google_cloud_datacatalog_v1_dataplex_spec_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataplexTableSpec); i {
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
		file_google_cloud_datacatalog_v1_dataplex_spec_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataplexExternalTable); i {
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
			RawDescriptor: file_google_cloud_datacatalog_v1_dataplex_spec_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_datacatalog_v1_dataplex_spec_proto_goTypes,
		DependencyIndexes: file_google_cloud_datacatalog_v1_dataplex_spec_proto_depIdxs,
		MessageInfos:      file_google_cloud_datacatalog_v1_dataplex_spec_proto_msgTypes,
	}.Build()
	File_google_cloud_datacatalog_v1_dataplex_spec_proto = out.File
	file_google_cloud_datacatalog_v1_dataplex_spec_proto_rawDesc = nil
	file_google_cloud_datacatalog_v1_dataplex_spec_proto_goTypes = nil
	file_google_cloud_datacatalog_v1_dataplex_spec_proto_depIdxs = nil
}
