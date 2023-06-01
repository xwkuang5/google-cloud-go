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
// 	protoc        v4.23.1
// source: google/cloud/support/v2/case.proto

package supportpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The status of a support case.
type Case_State int32

const (
	// Case is in an unknown state.
	Case_STATE_UNSPECIFIED Case_State = 0
	// The case has been created but no one is assigned to work on it yet.
	Case_NEW Case_State = 1
	// The case is currently being handled by Google support.
	Case_IN_PROGRESS_GOOGLE_SUPPORT Case_State = 2
	// Google is waiting for a response.
	Case_ACTION_REQUIRED Case_State = 3
	// A solution has been offered for the case, but it isn't yet closed.
	Case_SOLUTION_PROVIDED Case_State = 4
	// The case has been resolved.
	Case_CLOSED Case_State = 5
)

// Enum value maps for Case_State.
var (
	Case_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "NEW",
		2: "IN_PROGRESS_GOOGLE_SUPPORT",
		3: "ACTION_REQUIRED",
		4: "SOLUTION_PROVIDED",
		5: "CLOSED",
	}
	Case_State_value = map[string]int32{
		"STATE_UNSPECIFIED":          0,
		"NEW":                        1,
		"IN_PROGRESS_GOOGLE_SUPPORT": 2,
		"ACTION_REQUIRED":            3,
		"SOLUTION_PROVIDED":          4,
		"CLOSED":                     5,
	}
)

func (x Case_State) Enum() *Case_State {
	p := new(Case_State)
	*p = x
	return p
}

func (x Case_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Case_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_support_v2_case_proto_enumTypes[0].Descriptor()
}

func (Case_State) Type() protoreflect.EnumType {
	return &file_google_cloud_support_v2_case_proto_enumTypes[0]
}

func (x Case_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Case_State.Descriptor instead.
func (Case_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_support_v2_case_proto_rawDescGZIP(), []int{0, 0}
}

// The case Priority. P0 is most urgent and P4 the least.
type Case_Priority int32

const (
	// Priority is undefined or has not been set yet.
	Case_PRIORITY_UNSPECIFIED Case_Priority = 0
	// Extreme impact on a production service. Service is hard down.
	Case_P0 Case_Priority = 1
	// Critical impact on a production service. Service is currently unusable.
	Case_P1 Case_Priority = 2
	// Severe impact on a production service. Service is usable but greatly
	// impaired.
	Case_P2 Case_Priority = 3
	// Medium impact on a production service.  Service is available, but
	// moderately impaired.
	Case_P3 Case_Priority = 4
	// General questions or minor issues.  Production service is fully
	// available.
	Case_P4 Case_Priority = 5
)

// Enum value maps for Case_Priority.
var (
	Case_Priority_name = map[int32]string{
		0: "PRIORITY_UNSPECIFIED",
		1: "P0",
		2: "P1",
		3: "P2",
		4: "P3",
		5: "P4",
	}
	Case_Priority_value = map[string]int32{
		"PRIORITY_UNSPECIFIED": 0,
		"P0":                   1,
		"P1":                   2,
		"P2":                   3,
		"P3":                   4,
		"P4":                   5,
	}
)

func (x Case_Priority) Enum() *Case_Priority {
	p := new(Case_Priority)
	*p = x
	return p
}

func (x Case_Priority) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Case_Priority) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_support_v2_case_proto_enumTypes[1].Descriptor()
}

func (Case_Priority) Type() protoreflect.EnumType {
	return &file_google_cloud_support_v2_case_proto_enumTypes[1]
}

func (x Case_Priority) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Case_Priority.Descriptor instead.
func (Case_Priority) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_support_v2_case_proto_rawDescGZIP(), []int{0, 1}
}

// A support case.
type Case struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name for the case.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The short summary of the issue reported in this case.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// A broad description of the issue.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// The issue classification applicable to this case.
	Classification *CaseClassification `protobuf:"bytes,4,opt,name=classification,proto3" json:"classification,omitempty"`
	// The timezone of the user who created the support case.
	// It should be in a format IANA recognizes: https://www.iana.org/time-zones.
	// There is no additional validation done by the API.
	TimeZone string `protobuf:"bytes,8,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
	// The email addresses to receive updates on this case.
	SubscriberEmailAddresses []string `protobuf:"bytes,9,rep,name=subscriber_email_addresses,json=subscriberEmailAddresses,proto3" json:"subscriber_email_addresses,omitempty"`
	// Output only. The current status of the support case.
	State Case_State `protobuf:"varint,12,opt,name=state,proto3,enum=google.cloud.support.v2.Case_State" json:"state,omitempty"`
	// Output only. The time this case was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The time this case was last updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The user who created the case.
	//
	// Note: The name and email will be obfuscated if the case was created by
	// Google Support.
	Creator *Actor `protobuf:"bytes,15,opt,name=creator,proto3" json:"creator,omitempty"`
	// A user-supplied email address to send case update notifications for. This
	// should only be used in BYOID flows, where we cannot infer the user's email
	// address directly from their EUCs.
	ContactEmail string `protobuf:"bytes,35,opt,name=contact_email,json=contactEmail,proto3" json:"contact_email,omitempty"`
	// Whether the case is currently escalated.
	Escalated bool `protobuf:"varint,17,opt,name=escalated,proto3" json:"escalated,omitempty"`
	// Whether this case was created for internal API testing and should not be
	// acted on by the support team.
	TestCase bool `protobuf:"varint,19,opt,name=test_case,json=testCase,proto3" json:"test_case,omitempty"`
	// The language the user has requested to receive support in. This should be a
	// BCP 47 language code (e.g., `"en"`, `"zh-CN"`, `"zh-TW"`, `"ja"`, `"ko"`).
	// If no language or an unsupported language is specified, this field defaults
	// to English (en).
	//
	// Language selection during case creation may affect your available support
	// options. For a list of supported languages and their support working hours,
	// see: https://cloud.google.com/support/docs/language-working-hours
	LanguageCode string `protobuf:"bytes,23,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// The priority of this case.
	Priority Case_Priority `protobuf:"varint,32,opt,name=priority,proto3,enum=google.cloud.support.v2.Case_Priority" json:"priority,omitempty"`
}

func (x *Case) Reset() {
	*x = Case{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_support_v2_case_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Case) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Case) ProtoMessage() {}

func (x *Case) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_support_v2_case_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Case.ProtoReflect.Descriptor instead.
func (*Case) Descriptor() ([]byte, []int) {
	return file_google_cloud_support_v2_case_proto_rawDescGZIP(), []int{0}
}

func (x *Case) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Case) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Case) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Case) GetClassification() *CaseClassification {
	if x != nil {
		return x.Classification
	}
	return nil
}

func (x *Case) GetTimeZone() string {
	if x != nil {
		return x.TimeZone
	}
	return ""
}

func (x *Case) GetSubscriberEmailAddresses() []string {
	if x != nil {
		return x.SubscriberEmailAddresses
	}
	return nil
}

func (x *Case) GetState() Case_State {
	if x != nil {
		return x.State
	}
	return Case_STATE_UNSPECIFIED
}

func (x *Case) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Case) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Case) GetCreator() *Actor {
	if x != nil {
		return x.Creator
	}
	return nil
}

func (x *Case) GetContactEmail() string {
	if x != nil {
		return x.ContactEmail
	}
	return ""
}

func (x *Case) GetEscalated() bool {
	if x != nil {
		return x.Escalated
	}
	return false
}

func (x *Case) GetTestCase() bool {
	if x != nil {
		return x.TestCase
	}
	return false
}

func (x *Case) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

func (x *Case) GetPriority() Case_Priority {
	if x != nil {
		return x.Priority
	}
	return Case_PRIORITY_UNSPECIFIED
}

// A classification object with a product type and value.
type CaseClassification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique ID for a classification. Must be specified for case creation.
	//
	// To retrieve valid classification IDs for case creation, use
	// `caseClassifications.search`.
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The display name of the classification.
	DisplayName string `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *CaseClassification) Reset() {
	*x = CaseClassification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_support_v2_case_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CaseClassification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaseClassification) ProtoMessage() {}

func (x *CaseClassification) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_support_v2_case_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaseClassification.ProtoReflect.Descriptor instead.
func (*CaseClassification) Descriptor() ([]byte, []int) {
	return file_google_cloud_support_v2_case_proto_rawDescGZIP(), []int{1}
}

func (x *CaseClassification) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CaseClassification) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

var File_google_cloud_support_v2_case_proto protoreflect.FileDescriptor

var file_google_cloud_support_v2_case_proto_rawDesc = []byte{
	0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2f,
	0x76, 0x32, 0x2f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x98, 0x08, 0x0a, 0x04, 0x43, 0x61, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x53, 0x0a, 0x0e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x76, 0x32, 0x2e, 0x43, 0x61, 0x73, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x7a,
	0x6f, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x5a,
	0x6f, 0x6e, 0x65, 0x12, 0x3c, 0x0a, 0x1a, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x18, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x12, 0x3e, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x61, 0x73, 0x65, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x32,
	0x2e, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x23, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x74, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x42, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x20, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x32,
	0x2e, 0x43, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x7f, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x45, 0x57, 0x10,
	0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53,
	0x5f, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x10,
	0x02, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55,
	0x49, 0x52, 0x45, 0x44, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x4f, 0x4c, 0x55, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0a, 0x0a,
	0x06, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x05, 0x22, 0x4c, 0x0a, 0x08, 0x50, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54,
	0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x06, 0x0a, 0x02, 0x50, 0x30, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x50, 0x31, 0x10, 0x02, 0x12,
	0x06, 0x0a, 0x02, 0x50, 0x32, 0x10, 0x03, 0x12, 0x06, 0x0a, 0x02, 0x50, 0x33, 0x10, 0x04, 0x12,
	0x06, 0x0a, 0x02, 0x50, 0x34, 0x10, 0x05, 0x3a, 0x71, 0xea, 0x41, 0x6e, 0x0a, 0x20, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x73, 0x65, 0x12, 0x29,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x61, 0x73,
	0x65, 0x73, 0x2f, 0x7b, 0x63, 0x61, 0x73, 0x65, 0x7d, 0x12, 0x1f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x63, 0x61,
	0x73, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x61, 0x73, 0x65, 0x7d, 0x22, 0x47, 0x0a, 0x12, 0x43, 0x61,
	0x73, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x42, 0xb2, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x76, 0x32, 0x42, 0x09, 0x43, 0x61, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x35, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x70, 0x62, 0x3b, 0x73, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x70, 0x62, 0xaa, 0x02, 0x17, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x56,
	0x32, 0xca, 0x02, 0x17, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x5c, 0x56, 0x32, 0xea, 0x02, 0x1a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_support_v2_case_proto_rawDescOnce sync.Once
	file_google_cloud_support_v2_case_proto_rawDescData = file_google_cloud_support_v2_case_proto_rawDesc
)

func file_google_cloud_support_v2_case_proto_rawDescGZIP() []byte {
	file_google_cloud_support_v2_case_proto_rawDescOnce.Do(func() {
		file_google_cloud_support_v2_case_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_support_v2_case_proto_rawDescData)
	})
	return file_google_cloud_support_v2_case_proto_rawDescData
}

var file_google_cloud_support_v2_case_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_support_v2_case_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_support_v2_case_proto_goTypes = []interface{}{
	(Case_State)(0),               // 0: google.cloud.support.v2.Case.State
	(Case_Priority)(0),            // 1: google.cloud.support.v2.Case.Priority
	(*Case)(nil),                  // 2: google.cloud.support.v2.Case
	(*CaseClassification)(nil),    // 3: google.cloud.support.v2.CaseClassification
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*Actor)(nil),                 // 5: google.cloud.support.v2.Actor
}
var file_google_cloud_support_v2_case_proto_depIdxs = []int32{
	3, // 0: google.cloud.support.v2.Case.classification:type_name -> google.cloud.support.v2.CaseClassification
	0, // 1: google.cloud.support.v2.Case.state:type_name -> google.cloud.support.v2.Case.State
	4, // 2: google.cloud.support.v2.Case.create_time:type_name -> google.protobuf.Timestamp
	4, // 3: google.cloud.support.v2.Case.update_time:type_name -> google.protobuf.Timestamp
	5, // 4: google.cloud.support.v2.Case.creator:type_name -> google.cloud.support.v2.Actor
	1, // 5: google.cloud.support.v2.Case.priority:type_name -> google.cloud.support.v2.Case.Priority
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_cloud_support_v2_case_proto_init() }
func file_google_cloud_support_v2_case_proto_init() {
	if File_google_cloud_support_v2_case_proto != nil {
		return
	}
	file_google_cloud_support_v2_actor_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_support_v2_case_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Case); i {
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
		file_google_cloud_support_v2_case_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CaseClassification); i {
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
			RawDescriptor: file_google_cloud_support_v2_case_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_support_v2_case_proto_goTypes,
		DependencyIndexes: file_google_cloud_support_v2_case_proto_depIdxs,
		EnumInfos:         file_google_cloud_support_v2_case_proto_enumTypes,
		MessageInfos:      file_google_cloud_support_v2_case_proto_msgTypes,
	}.Build()
	File_google_cloud_support_v2_case_proto = out.File
	file_google_cloud_support_v2_case_proto_rawDesc = nil
	file_google_cloud_support_v2_case_proto_goTypes = nil
	file_google_cloud_support_v2_case_proto_depIdxs = nil
}
