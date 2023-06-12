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
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: google/cloud/websecurityscanner/v1/scan_config_error.proto

package websecurityscannerpb

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

// Output only.
// Defines an error reason code.
// Next id: 44
type ScanConfigError_Code int32

const (
	// There is no error.
	ScanConfigError_CODE_UNSPECIFIED ScanConfigError_Code = 0
	// There is no error.
	ScanConfigError_OK ScanConfigError_Code = 0
	// Indicates an internal server error.
	// Please DO NOT USE THIS ERROR CODE unless the root cause is truly unknown.
	ScanConfigError_INTERNAL_ERROR ScanConfigError_Code = 1
	// One of the seed URLs is an App Engine URL but we cannot validate the scan
	// settings due to an App Engine API backend error.
	ScanConfigError_APPENGINE_API_BACKEND_ERROR ScanConfigError_Code = 2
	// One of the seed URLs is an App Engine URL but we cannot access the
	// App Engine API to validate scan settings.
	ScanConfigError_APPENGINE_API_NOT_ACCESSIBLE ScanConfigError_Code = 3
	// One of the seed URLs is an App Engine URL but the Default Host of the
	// App Engine is not set.
	ScanConfigError_APPENGINE_DEFAULT_HOST_MISSING ScanConfigError_Code = 4
	// Google corporate accounts can not be used for scanning.
	ScanConfigError_CANNOT_USE_GOOGLE_COM_ACCOUNT ScanConfigError_Code = 6
	// The account of the scan creator can not be used for scanning.
	ScanConfigError_CANNOT_USE_OWNER_ACCOUNT ScanConfigError_Code = 7
	// This scan targets Compute Engine, but we cannot validate scan settings
	// due to a Compute Engine API backend error.
	ScanConfigError_COMPUTE_API_BACKEND_ERROR ScanConfigError_Code = 8
	// This scan targets Compute Engine, but we cannot access the Compute Engine
	// API to validate the scan settings.
	ScanConfigError_COMPUTE_API_NOT_ACCESSIBLE ScanConfigError_Code = 9
	// The Custom Login URL does not belong to the current project.
	ScanConfigError_CUSTOM_LOGIN_URL_DOES_NOT_BELONG_TO_CURRENT_PROJECT ScanConfigError_Code = 10
	// The Custom Login URL is malformed (can not be parsed).
	ScanConfigError_CUSTOM_LOGIN_URL_MALFORMED ScanConfigError_Code = 11
	// The Custom Login URL is mapped to a non-routable IP address in DNS.
	ScanConfigError_CUSTOM_LOGIN_URL_MAPPED_TO_NON_ROUTABLE_ADDRESS ScanConfigError_Code = 12
	// The Custom Login URL is mapped to an IP address which is not reserved for
	// the current project.
	ScanConfigError_CUSTOM_LOGIN_URL_MAPPED_TO_UNRESERVED_ADDRESS ScanConfigError_Code = 13
	// The Custom Login URL has a non-routable IP address.
	ScanConfigError_CUSTOM_LOGIN_URL_HAS_NON_ROUTABLE_IP_ADDRESS ScanConfigError_Code = 14
	// The Custom Login URL has an IP address which is not reserved for the
	// current project.
	ScanConfigError_CUSTOM_LOGIN_URL_HAS_UNRESERVED_IP_ADDRESS ScanConfigError_Code = 15
	// Another scan with the same name (case-sensitive) already exists.
	ScanConfigError_DUPLICATE_SCAN_NAME ScanConfigError_Code = 16
	// A field is set to an invalid value.
	ScanConfigError_INVALID_FIELD_VALUE ScanConfigError_Code = 18
	// There was an error trying to authenticate to the scan target.
	ScanConfigError_FAILED_TO_AUTHENTICATE_TO_TARGET ScanConfigError_Code = 19
	// Finding type value is not specified in the list findings request.
	ScanConfigError_FINDING_TYPE_UNSPECIFIED ScanConfigError_Code = 20
	// Scan targets Compute Engine, yet current project was not whitelisted for
	// Google Compute Engine Scanning Alpha access.
	ScanConfigError_FORBIDDEN_TO_SCAN_COMPUTE ScanConfigError_Code = 21
	// User tries to update managed scan
	ScanConfigError_FORBIDDEN_UPDATE_TO_MANAGED_SCAN ScanConfigError_Code = 43
	// The supplied filter is malformed. For example, it can not be parsed, does
	// not have a filter type in expression, or the same filter type appears
	// more than once.
	ScanConfigError_MALFORMED_FILTER ScanConfigError_Code = 22
	// The supplied resource name is malformed (can not be parsed).
	ScanConfigError_MALFORMED_RESOURCE_NAME ScanConfigError_Code = 23
	// The current project is not in an active state.
	ScanConfigError_PROJECT_INACTIVE ScanConfigError_Code = 24
	// A required field is not set.
	ScanConfigError_REQUIRED_FIELD ScanConfigError_Code = 25
	// Project id, scanconfig id, scanrun id, or finding id are not consistent
	// with each other in resource name.
	ScanConfigError_RESOURCE_NAME_INCONSISTENT ScanConfigError_Code = 26
	// The scan being requested to start is already running.
	ScanConfigError_SCAN_ALREADY_RUNNING ScanConfigError_Code = 27
	// The scan that was requested to be stopped is not running.
	ScanConfigError_SCAN_NOT_RUNNING ScanConfigError_Code = 28
	// One of the seed URLs does not belong to the current project.
	ScanConfigError_SEED_URL_DOES_NOT_BELONG_TO_CURRENT_PROJECT ScanConfigError_Code = 29
	// One of the seed URLs is malformed (can not be parsed).
	ScanConfigError_SEED_URL_MALFORMED ScanConfigError_Code = 30
	// One of the seed URLs is mapped to a non-routable IP address in DNS.
	ScanConfigError_SEED_URL_MAPPED_TO_NON_ROUTABLE_ADDRESS ScanConfigError_Code = 31
	// One of the seed URLs is mapped to an IP address which is not reserved
	// for the current project.
	ScanConfigError_SEED_URL_MAPPED_TO_UNRESERVED_ADDRESS ScanConfigError_Code = 32
	// One of the seed URLs has on-routable IP address.
	ScanConfigError_SEED_URL_HAS_NON_ROUTABLE_IP_ADDRESS ScanConfigError_Code = 33
	// One of the seed URLs has an IP address that is not reserved
	// for the current project.
	ScanConfigError_SEED_URL_HAS_UNRESERVED_IP_ADDRESS ScanConfigError_Code = 35
	// The Web Security Scanner service account is not configured under the
	// project.
	ScanConfigError_SERVICE_ACCOUNT_NOT_CONFIGURED ScanConfigError_Code = 36
	// A project has reached the maximum number of scans.
	ScanConfigError_TOO_MANY_SCANS ScanConfigError_Code = 37
	// Resolving the details of the current project fails.
	ScanConfigError_UNABLE_TO_RESOLVE_PROJECT_INFO ScanConfigError_Code = 38
	// One or more blacklist patterns were in the wrong format.
	ScanConfigError_UNSUPPORTED_BLACKLIST_PATTERN_FORMAT ScanConfigError_Code = 39
	// The supplied filter is not supported.
	ScanConfigError_UNSUPPORTED_FILTER ScanConfigError_Code = 40
	// The supplied finding type is not supported. For example, we do not
	// provide findings of the given finding type.
	ScanConfigError_UNSUPPORTED_FINDING_TYPE ScanConfigError_Code = 41
	// The URL scheme of one or more of the supplied URLs is not supported.
	ScanConfigError_UNSUPPORTED_URL_SCHEME ScanConfigError_Code = 42
)

// Enum value maps for ScanConfigError_Code.
var (
	ScanConfigError_Code_name = map[int32]string{
		0: "CODE_UNSPECIFIED",
		// Duplicate value: 0: "OK",
		1:  "INTERNAL_ERROR",
		2:  "APPENGINE_API_BACKEND_ERROR",
		3:  "APPENGINE_API_NOT_ACCESSIBLE",
		4:  "APPENGINE_DEFAULT_HOST_MISSING",
		6:  "CANNOT_USE_GOOGLE_COM_ACCOUNT",
		7:  "CANNOT_USE_OWNER_ACCOUNT",
		8:  "COMPUTE_API_BACKEND_ERROR",
		9:  "COMPUTE_API_NOT_ACCESSIBLE",
		10: "CUSTOM_LOGIN_URL_DOES_NOT_BELONG_TO_CURRENT_PROJECT",
		11: "CUSTOM_LOGIN_URL_MALFORMED",
		12: "CUSTOM_LOGIN_URL_MAPPED_TO_NON_ROUTABLE_ADDRESS",
		13: "CUSTOM_LOGIN_URL_MAPPED_TO_UNRESERVED_ADDRESS",
		14: "CUSTOM_LOGIN_URL_HAS_NON_ROUTABLE_IP_ADDRESS",
		15: "CUSTOM_LOGIN_URL_HAS_UNRESERVED_IP_ADDRESS",
		16: "DUPLICATE_SCAN_NAME",
		18: "INVALID_FIELD_VALUE",
		19: "FAILED_TO_AUTHENTICATE_TO_TARGET",
		20: "FINDING_TYPE_UNSPECIFIED",
		21: "FORBIDDEN_TO_SCAN_COMPUTE",
		43: "FORBIDDEN_UPDATE_TO_MANAGED_SCAN",
		22: "MALFORMED_FILTER",
		23: "MALFORMED_RESOURCE_NAME",
		24: "PROJECT_INACTIVE",
		25: "REQUIRED_FIELD",
		26: "RESOURCE_NAME_INCONSISTENT",
		27: "SCAN_ALREADY_RUNNING",
		28: "SCAN_NOT_RUNNING",
		29: "SEED_URL_DOES_NOT_BELONG_TO_CURRENT_PROJECT",
		30: "SEED_URL_MALFORMED",
		31: "SEED_URL_MAPPED_TO_NON_ROUTABLE_ADDRESS",
		32: "SEED_URL_MAPPED_TO_UNRESERVED_ADDRESS",
		33: "SEED_URL_HAS_NON_ROUTABLE_IP_ADDRESS",
		35: "SEED_URL_HAS_UNRESERVED_IP_ADDRESS",
		36: "SERVICE_ACCOUNT_NOT_CONFIGURED",
		37: "TOO_MANY_SCANS",
		38: "UNABLE_TO_RESOLVE_PROJECT_INFO",
		39: "UNSUPPORTED_BLACKLIST_PATTERN_FORMAT",
		40: "UNSUPPORTED_FILTER",
		41: "UNSUPPORTED_FINDING_TYPE",
		42: "UNSUPPORTED_URL_SCHEME",
	}
	ScanConfigError_Code_value = map[string]int32{
		"CODE_UNSPECIFIED":                                    0,
		"OK":                                                  0,
		"INTERNAL_ERROR":                                      1,
		"APPENGINE_API_BACKEND_ERROR":                         2,
		"APPENGINE_API_NOT_ACCESSIBLE":                        3,
		"APPENGINE_DEFAULT_HOST_MISSING":                      4,
		"CANNOT_USE_GOOGLE_COM_ACCOUNT":                       6,
		"CANNOT_USE_OWNER_ACCOUNT":                            7,
		"COMPUTE_API_BACKEND_ERROR":                           8,
		"COMPUTE_API_NOT_ACCESSIBLE":                          9,
		"CUSTOM_LOGIN_URL_DOES_NOT_BELONG_TO_CURRENT_PROJECT": 10,
		"CUSTOM_LOGIN_URL_MALFORMED":                          11,
		"CUSTOM_LOGIN_URL_MAPPED_TO_NON_ROUTABLE_ADDRESS":     12,
		"CUSTOM_LOGIN_URL_MAPPED_TO_UNRESERVED_ADDRESS":       13,
		"CUSTOM_LOGIN_URL_HAS_NON_ROUTABLE_IP_ADDRESS":        14,
		"CUSTOM_LOGIN_URL_HAS_UNRESERVED_IP_ADDRESS":          15,
		"DUPLICATE_SCAN_NAME":                                 16,
		"INVALID_FIELD_VALUE":                                 18,
		"FAILED_TO_AUTHENTICATE_TO_TARGET":                    19,
		"FINDING_TYPE_UNSPECIFIED":                            20,
		"FORBIDDEN_TO_SCAN_COMPUTE":                           21,
		"FORBIDDEN_UPDATE_TO_MANAGED_SCAN":                    43,
		"MALFORMED_FILTER":                                    22,
		"MALFORMED_RESOURCE_NAME":                             23,
		"PROJECT_INACTIVE":                                    24,
		"REQUIRED_FIELD":                                      25,
		"RESOURCE_NAME_INCONSISTENT":                          26,
		"SCAN_ALREADY_RUNNING":                                27,
		"SCAN_NOT_RUNNING":                                    28,
		"SEED_URL_DOES_NOT_BELONG_TO_CURRENT_PROJECT":         29,
		"SEED_URL_MALFORMED":                                  30,
		"SEED_URL_MAPPED_TO_NON_ROUTABLE_ADDRESS":             31,
		"SEED_URL_MAPPED_TO_UNRESERVED_ADDRESS":               32,
		"SEED_URL_HAS_NON_ROUTABLE_IP_ADDRESS":                33,
		"SEED_URL_HAS_UNRESERVED_IP_ADDRESS":                  35,
		"SERVICE_ACCOUNT_NOT_CONFIGURED":                      36,
		"TOO_MANY_SCANS":                                      37,
		"UNABLE_TO_RESOLVE_PROJECT_INFO":                      38,
		"UNSUPPORTED_BLACKLIST_PATTERN_FORMAT":                39,
		"UNSUPPORTED_FILTER":                                  40,
		"UNSUPPORTED_FINDING_TYPE":                            41,
		"UNSUPPORTED_URL_SCHEME":                              42,
	}
)

func (x ScanConfigError_Code) Enum() *ScanConfigError_Code {
	p := new(ScanConfigError_Code)
	*p = x
	return p
}

func (x ScanConfigError_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScanConfigError_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_websecurityscanner_v1_scan_config_error_proto_enumTypes[0].Descriptor()
}

func (ScanConfigError_Code) Type() protoreflect.EnumType {
	return &file_google_cloud_websecurityscanner_v1_scan_config_error_proto_enumTypes[0]
}

func (x ScanConfigError_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ScanConfigError_Code.Descriptor instead.
func (ScanConfigError_Code) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_websecurityscanner_v1_scan_config_error_proto_rawDescGZIP(), []int{0, 0}
}

// Defines a custom error message used by CreateScanConfig and UpdateScanConfig
// APIs when scan configuration validation fails. It is also reported as part of
// a ScanRunErrorTrace message if scan validation fails due to a scan
// configuration error.
type ScanConfigError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Indicates the reason code for a configuration failure.
	Code ScanConfigError_Code `protobuf:"varint,1,opt,name=code,proto3,enum=google.cloud.websecurityscanner.v1.ScanConfigError_Code" json:"code,omitempty"`
	// Output only. Indicates the full name of the ScanConfig field that triggers this error,
	// for example "scan_config.max_qps". This field is provided for
	// troubleshooting purposes only and its actual value can change in the
	// future.
	FieldName string `protobuf:"bytes,2,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
}

func (x *ScanConfigError) Reset() {
	*x = ScanConfigError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_websecurityscanner_v1_scan_config_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanConfigError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanConfigError) ProtoMessage() {}

func (x *ScanConfigError) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_websecurityscanner_v1_scan_config_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanConfigError.ProtoReflect.Descriptor instead.
func (*ScanConfigError) Descriptor() ([]byte, []int) {
	return file_google_cloud_websecurityscanner_v1_scan_config_error_proto_rawDescGZIP(), []int{0}
}

func (x *ScanConfigError) GetCode() ScanConfigError_Code {
	if x != nil {
		return x.Code
	}
	return ScanConfigError_CODE_UNSPECIFIED
}

func (x *ScanConfigError) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

var File_google_cloud_websecurityscanner_v1_scan_config_error_proto protoreflect.FileDescriptor

var file_google_cloud_websecurityscanner_v1_scan_config_error_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x77,
	0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x22, 0xfa, 0x0b, 0x0a, 0x0f, 0x53, 0x63, 0x61, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x4c, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0xf9, 0x0a, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f,
	0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4e, 0x54, 0x45,
	0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b,
	0x41, 0x50, 0x50, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0x5f, 0x41, 0x50, 0x49, 0x5f, 0x42, 0x41,
	0x43, 0x4b, 0x45, 0x4e, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x20, 0x0a,
	0x1c, 0x41, 0x50, 0x50, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0x5f, 0x41, 0x50, 0x49, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x49, 0x42, 0x4c, 0x45, 0x10, 0x03, 0x12,
	0x22, 0x0a, 0x1e, 0x41, 0x50, 0x50, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0x5f, 0x44, 0x45, 0x46,
	0x41, 0x55, 0x4c, 0x54, 0x5f, 0x48, 0x4f, 0x53, 0x54, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e,
	0x47, 0x10, 0x04, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x55, 0x53,
	0x45, 0x5f, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x5f, 0x41, 0x43, 0x43,
	0x4f, 0x55, 0x4e, 0x54, 0x10, 0x06, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54,
	0x5f, 0x55, 0x53, 0x45, 0x5f, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55,
	0x4e, 0x54, 0x10, 0x07, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x4f, 0x4d, 0x50, 0x55, 0x54, 0x45, 0x5f,
	0x41, 0x50, 0x49, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x45, 0x4e, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x08, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x4f, 0x4d, 0x50, 0x55, 0x54, 0x45, 0x5f, 0x41,
	0x50, 0x49, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x49, 0x42, 0x4c,
	0x45, 0x10, 0x09, 0x12, 0x37, 0x0a, 0x33, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x4c, 0x4f,
	0x47, 0x49, 0x4e, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x44, 0x4f, 0x45, 0x53, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x42, 0x45, 0x4c, 0x4f, 0x4e, 0x47, 0x5f, 0x54, 0x4f, 0x5f, 0x43, 0x55, 0x52, 0x52, 0x45,
	0x4e, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x0a, 0x12, 0x1e, 0x0a, 0x1a,
	0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x55, 0x52, 0x4c,
	0x5f, 0x4d, 0x41, 0x4c, 0x46, 0x4f, 0x52, 0x4d, 0x45, 0x44, 0x10, 0x0b, 0x12, 0x33, 0x0a, 0x2f,
	0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x55, 0x52, 0x4c,
	0x5f, 0x4d, 0x41, 0x50, 0x50, 0x45, 0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x4e, 0x4f, 0x4e, 0x5f, 0x52,
	0x4f, 0x55, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x10,
	0x0c, 0x12, 0x31, 0x0a, 0x2d, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x4c, 0x4f, 0x47, 0x49,
	0x4e, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x4d, 0x41, 0x50, 0x50, 0x45, 0x44, 0x5f, 0x54, 0x4f, 0x5f,
	0x55, 0x4e, 0x52, 0x45, 0x53, 0x45, 0x52, 0x56, 0x45, 0x44, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45,
	0x53, 0x53, 0x10, 0x0d, 0x12, 0x30, 0x0a, 0x2c, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x4c,
	0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x48, 0x41, 0x53, 0x5f, 0x4e, 0x4f, 0x4e,
	0x5f, 0x52, 0x4f, 0x55, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x49, 0x50, 0x5f, 0x41, 0x44, 0x44,
	0x52, 0x45, 0x53, 0x53, 0x10, 0x0e, 0x12, 0x2e, 0x0a, 0x2a, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d,
	0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x48, 0x41, 0x53, 0x5f, 0x55,
	0x4e, 0x52, 0x45, 0x53, 0x45, 0x52, 0x56, 0x45, 0x44, 0x5f, 0x49, 0x50, 0x5f, 0x41, 0x44, 0x44,
	0x52, 0x45, 0x53, 0x53, 0x10, 0x0f, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43,
	0x41, 0x54, 0x45, 0x5f, 0x53, 0x43, 0x41, 0x4e, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x10, 0x12,
	0x17, 0x0a, 0x13, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44,
	0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x12, 0x12, 0x24, 0x0a, 0x20, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41,
	0x54, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x10, 0x13, 0x12, 0x1c,
	0x0a, 0x18, 0x46, 0x49, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x14, 0x12, 0x1d, 0x0a, 0x19,
	0x46, 0x4f, 0x52, 0x42, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x5f, 0x54, 0x4f, 0x5f, 0x53, 0x43, 0x41,
	0x4e, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x55, 0x54, 0x45, 0x10, 0x15, 0x12, 0x24, 0x0a, 0x20, 0x46,
	0x4f, 0x52, 0x42, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f,
	0x54, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x44, 0x5f, 0x53, 0x43, 0x41, 0x4e, 0x10,
	0x2b, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x41, 0x4c, 0x46, 0x4f, 0x52, 0x4d, 0x45, 0x44, 0x5f, 0x46,
	0x49, 0x4c, 0x54, 0x45, 0x52, 0x10, 0x16, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x41, 0x4c, 0x46, 0x4f,
	0x52, 0x4d, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x4e, 0x41,
	0x4d, 0x45, 0x10, 0x17, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f,
	0x49, 0x4e, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x18, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45,
	0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x10, 0x19, 0x12, 0x1e,
	0x0a, 0x1a, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f,
	0x49, 0x4e, 0x43, 0x4f, 0x4e, 0x53, 0x49, 0x53, 0x54, 0x45, 0x4e, 0x54, 0x10, 0x1a, 0x12, 0x18,
	0x0a, 0x14, 0x53, 0x43, 0x41, 0x4e, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x52,
	0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x1b, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x43, 0x41, 0x4e,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x1c, 0x12, 0x2f,
	0x0a, 0x2b, 0x53, 0x45, 0x45, 0x44, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x44, 0x4f, 0x45, 0x53, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x42, 0x45, 0x4c, 0x4f, 0x4e, 0x47, 0x5f, 0x54, 0x4f, 0x5f, 0x43, 0x55,
	0x52, 0x52, 0x45, 0x4e, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x1d, 0x12,
	0x16, 0x0a, 0x12, 0x53, 0x45, 0x45, 0x44, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x4d, 0x41, 0x4c, 0x46,
	0x4f, 0x52, 0x4d, 0x45, 0x44, 0x10, 0x1e, 0x12, 0x2b, 0x0a, 0x27, 0x53, 0x45, 0x45, 0x44, 0x5f,
	0x55, 0x52, 0x4c, 0x5f, 0x4d, 0x41, 0x50, 0x50, 0x45, 0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x4e, 0x4f,
	0x4e, 0x5f, 0x52, 0x4f, 0x55, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45,
	0x53, 0x53, 0x10, 0x1f, 0x12, 0x29, 0x0a, 0x25, 0x53, 0x45, 0x45, 0x44, 0x5f, 0x55, 0x52, 0x4c,
	0x5f, 0x4d, 0x41, 0x50, 0x50, 0x45, 0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x55, 0x4e, 0x52, 0x45, 0x53,
	0x45, 0x52, 0x56, 0x45, 0x44, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x10, 0x20, 0x12,
	0x28, 0x0a, 0x24, 0x53, 0x45, 0x45, 0x44, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x48, 0x41, 0x53, 0x5f,
	0x4e, 0x4f, 0x4e, 0x5f, 0x52, 0x4f, 0x55, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x49, 0x50, 0x5f,
	0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x10, 0x21, 0x12, 0x26, 0x0a, 0x22, 0x53, 0x45, 0x45,
	0x44, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x48, 0x41, 0x53, 0x5f, 0x55, 0x4e, 0x52, 0x45, 0x53, 0x45,
	0x52, 0x56, 0x45, 0x44, 0x5f, 0x49, 0x50, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x10,
	0x23, 0x12, 0x22, 0x0a, 0x1e, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x41, 0x43, 0x43,
	0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x55,
	0x52, 0x45, 0x44, 0x10, 0x24, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e,
	0x59, 0x5f, 0x53, 0x43, 0x41, 0x4e, 0x53, 0x10, 0x25, 0x12, 0x22, 0x0a, 0x1e, 0x55, 0x4e, 0x41,
	0x42, 0x4c, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x5f, 0x50,
	0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x26, 0x12, 0x28, 0x0a,
	0x24, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x42, 0x4c, 0x41,
	0x43, 0x4b, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x50, 0x41, 0x54, 0x54, 0x45, 0x52, 0x4e, 0x5f, 0x46,
	0x4f, 0x52, 0x4d, 0x41, 0x54, 0x10, 0x27, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e, 0x53, 0x55, 0x50,
	0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x10, 0x28, 0x12,
	0x1c, 0x0a, 0x18, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x46,
	0x49, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x29, 0x12, 0x1a, 0x0a,
	0x16, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x55, 0x52, 0x4c,
	0x5f, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x45, 0x10, 0x2a, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x8a, 0x02,
	0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x14, 0x53, 0x63, 0x61, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x56, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x77,
	0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x70, 0x62, 0x3b, 0x77, 0x65, 0x62, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x73,
	0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x70, 0x62, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x57, 0x65, 0x62, 0x53, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x22,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x57, 0x65, 0x62,
	0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5c,
	0x56, 0x31, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x57, 0x65, 0x62, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x53,
	0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_cloud_websecurityscanner_v1_scan_config_error_proto_rawDescOnce sync.Once
	file_google_cloud_websecurityscanner_v1_scan_config_error_proto_rawDescData = file_google_cloud_websecurityscanner_v1_scan_config_error_proto_rawDesc
)

func file_google_cloud_websecurityscanner_v1_scan_config_error_proto_rawDescGZIP() []byte {
	file_google_cloud_websecurityscanner_v1_scan_config_error_proto_rawDescOnce.Do(func() {
		file_google_cloud_websecurityscanner_v1_scan_config_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_websecurityscanner_v1_scan_config_error_proto_rawDescData)
	})
	return file_google_cloud_websecurityscanner_v1_scan_config_error_proto_rawDescData
}

var file_google_cloud_websecurityscanner_v1_scan_config_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_websecurityscanner_v1_scan_config_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_websecurityscanner_v1_scan_config_error_proto_goTypes = []interface{}{
	(ScanConfigError_Code)(0), // 0: google.cloud.websecurityscanner.v1.ScanConfigError.Code
	(*ScanConfigError)(nil),   // 1: google.cloud.websecurityscanner.v1.ScanConfigError
}
var file_google_cloud_websecurityscanner_v1_scan_config_error_proto_depIdxs = []int32{
	0, // 0: google.cloud.websecurityscanner.v1.ScanConfigError.code:type_name -> google.cloud.websecurityscanner.v1.ScanConfigError.Code
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_websecurityscanner_v1_scan_config_error_proto_init() }
func file_google_cloud_websecurityscanner_v1_scan_config_error_proto_init() {
	if File_google_cloud_websecurityscanner_v1_scan_config_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_websecurityscanner_v1_scan_config_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanConfigError); i {
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
			RawDescriptor: file_google_cloud_websecurityscanner_v1_scan_config_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_websecurityscanner_v1_scan_config_error_proto_goTypes,
		DependencyIndexes: file_google_cloud_websecurityscanner_v1_scan_config_error_proto_depIdxs,
		EnumInfos:         file_google_cloud_websecurityscanner_v1_scan_config_error_proto_enumTypes,
		MessageInfos:      file_google_cloud_websecurityscanner_v1_scan_config_error_proto_msgTypes,
	}.Build()
	File_google_cloud_websecurityscanner_v1_scan_config_error_proto = out.File
	file_google_cloud_websecurityscanner_v1_scan_config_error_proto_rawDesc = nil
	file_google_cloud_websecurityscanner_v1_scan_config_error_proto_goTypes = nil
	file_google_cloud_websecurityscanner_v1_scan_config_error_proto_depIdxs = nil
}
