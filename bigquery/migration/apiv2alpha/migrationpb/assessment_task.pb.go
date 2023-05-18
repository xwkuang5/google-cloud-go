// Copyright 2021 Google LLC
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
// 	protoc        v3.21.12
// source: google/cloud/bigquery/migration/v2alpha/assessment_task.proto

package migrationpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Assessment task config.
type AssessmentTaskDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The Cloud Storage path for assessment input files.
	InputPath string `protobuf:"bytes,1,opt,name=input_path,json=inputPath,proto3" json:"input_path,omitempty"`
	// Required. The BigQuery dataset for output.
	OutputDataset string `protobuf:"bytes,2,opt,name=output_dataset,json=outputDataset,proto3" json:"output_dataset,omitempty"`
	// Optional. An optional Cloud Storage path to write the query logs (which is
	// then used as an input path on the translation task)
	QuerylogsPath string `protobuf:"bytes,3,opt,name=querylogs_path,json=querylogsPath,proto3" json:"querylogs_path,omitempty"`
	// Required. The data source or data warehouse type (eg: TERADATA/REDSHIFT)
	// from which the input data is extracted.
	DataSource string `protobuf:"bytes,4,opt,name=data_source,json=dataSource,proto3" json:"data_source,omitempty"`
}

func (x *AssessmentTaskDetails) Reset() {
	*x = AssessmentTaskDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssessmentTaskDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssessmentTaskDetails) ProtoMessage() {}

func (x *AssessmentTaskDetails) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssessmentTaskDetails.ProtoReflect.Descriptor instead.
func (*AssessmentTaskDetails) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_rawDescGZIP(), []int{0}
}

func (x *AssessmentTaskDetails) GetInputPath() string {
	if x != nil {
		return x.InputPath
	}
	return ""
}

func (x *AssessmentTaskDetails) GetOutputDataset() string {
	if x != nil {
		return x.OutputDataset
	}
	return ""
}

func (x *AssessmentTaskDetails) GetQuerylogsPath() string {
	if x != nil {
		return x.QuerylogsPath
	}
	return ""
}

func (x *AssessmentTaskDetails) GetDataSource() string {
	if x != nil {
		return x.DataSource
	}
	return ""
}

// Details for an assessment task orchestration result.
type AssessmentOrchestrationResultDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The version used for the output table schemas.
	OutputTablesSchemaVersion string `protobuf:"bytes,1,opt,name=output_tables_schema_version,json=outputTablesSchemaVersion,proto3" json:"output_tables_schema_version,omitempty"`
}

func (x *AssessmentOrchestrationResultDetails) Reset() {
	*x = AssessmentOrchestrationResultDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssessmentOrchestrationResultDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssessmentOrchestrationResultDetails) ProtoMessage() {}

func (x *AssessmentOrchestrationResultDetails) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssessmentOrchestrationResultDetails.ProtoReflect.Descriptor instead.
func (*AssessmentOrchestrationResultDetails) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_rawDescGZIP(), []int{1}
}

func (x *AssessmentOrchestrationResultDetails) GetOutputTablesSchemaVersion() string {
	if x != nil {
		return x.OutputTablesSchemaVersion
	}
	return ""
}

var File_google_cloud_bigquery_migration_v2alpha_assessment_task_proto protoreflect.FileDescriptor

var file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69,
	0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x15, 0x41, 0x73,
	0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x22, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x2a, 0x0a, 0x0e, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x12, 0x2a, 0x0a, 0x0e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x6c, 0x6f, 0x67, 0x73,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01,
	0x52, 0x0d, 0x71, 0x75, 0x65, 0x72, 0x79, 0x6c, 0x6f, 0x67, 0x73, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x24, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x6c, 0x0a, 0x24, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d,
	0x65, 0x6e, 0x74, 0x4f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x44, 0x0a,
	0x1c, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x5f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x19, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x42, 0xe3, 0x01, 0x0a, 0x2b, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x42, 0x13, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f,
	0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x3b, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x70, 0x62, 0xaa, 0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x32, 0x41, 0x6c, 0x70, 0x68, 0x61, 0xca,
	0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x42,
	0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x5c, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5c, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_rawDescOnce sync.Once
	file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_rawDescData = file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_rawDesc
)

func file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_rawDescGZIP() []byte {
	file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_rawDescOnce.Do(func() {
		file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_rawDescData)
	})
	return file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_rawDescData
}

var file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_goTypes = []interface{}{
	(*AssessmentTaskDetails)(nil),                // 0: google.cloud.bigquery.migration.v2alpha.AssessmentTaskDetails
	(*AssessmentOrchestrationResultDetails)(nil), // 1: google.cloud.bigquery.migration.v2alpha.AssessmentOrchestrationResultDetails
}
var file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_init() }
func file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_init() {
	if File_google_cloud_bigquery_migration_v2alpha_assessment_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssessmentTaskDetails); i {
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
		file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssessmentOrchestrationResultDetails); i {
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
			RawDescriptor: file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_goTypes,
		DependencyIndexes: file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_depIdxs,
		MessageInfos:      file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_msgTypes,
	}.Build()
	File_google_cloud_bigquery_migration_v2alpha_assessment_task_proto = out.File
	file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_rawDesc = nil
	file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_goTypes = nil
	file_google_cloud_bigquery_migration_v2alpha_assessment_task_proto_depIdxs = nil
}
