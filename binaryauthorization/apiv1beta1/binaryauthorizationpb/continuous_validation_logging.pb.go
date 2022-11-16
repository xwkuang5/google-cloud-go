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
// source: google/cloud/binaryauthorization/v1beta1/continuous_validation_logging.proto

package binaryauthorizationpb

import (
	reflect "reflect"
	sync "sync"

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

// Audit time policy conformance verdict.
type ContinuousValidationEvent_ContinuousValidationPodEvent_PolicyConformanceVerdict int32

const (
	// We should always have a verdict. This is an error.
	ContinuousValidationEvent_ContinuousValidationPodEvent_POLICY_CONFORMANCE_VERDICT_UNSPECIFIED ContinuousValidationEvent_ContinuousValidationPodEvent_PolicyConformanceVerdict = 0
	// The pod violates the policy.
	ContinuousValidationEvent_ContinuousValidationPodEvent_VIOLATES_POLICY ContinuousValidationEvent_ContinuousValidationPodEvent_PolicyConformanceVerdict = 1
)

// Enum value maps for ContinuousValidationEvent_ContinuousValidationPodEvent_PolicyConformanceVerdict.
var (
	ContinuousValidationEvent_ContinuousValidationPodEvent_PolicyConformanceVerdict_name = map[int32]string{
		0: "POLICY_CONFORMANCE_VERDICT_UNSPECIFIED",
		1: "VIOLATES_POLICY",
	}
	ContinuousValidationEvent_ContinuousValidationPodEvent_PolicyConformanceVerdict_value = map[string]int32{
		"POLICY_CONFORMANCE_VERDICT_UNSPECIFIED": 0,
		"VIOLATES_POLICY":                        1,
	}
)

func (x ContinuousValidationEvent_ContinuousValidationPodEvent_PolicyConformanceVerdict) Enum() *ContinuousValidationEvent_ContinuousValidationPodEvent_PolicyConformanceVerdict {
	p := new(ContinuousValidationEvent_ContinuousValidationPodEvent_PolicyConformanceVerdict)
	*p = x
	return p
}

func (x ContinuousValidationEvent_ContinuousValidationPodEvent_PolicyConformanceVerdict) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContinuousValidationEvent_ContinuousValidationPodEvent_PolicyConformanceVerdict) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_enumTypes[0].Descriptor()
}

func (ContinuousValidationEvent_ContinuousValidationPodEvent_PolicyConformanceVerdict) Type() protoreflect.EnumType {
	return &file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_enumTypes[0]
}

func (x ContinuousValidationEvent_ContinuousValidationPodEvent_PolicyConformanceVerdict) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContinuousValidationEvent_ContinuousValidationPodEvent_PolicyConformanceVerdict.Descriptor instead.
func (ContinuousValidationEvent_ContinuousValidationPodEvent_PolicyConformanceVerdict) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_rawDescGZIP(), []int{0, 0, 0}
}

// Result of the audit.
type ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AuditResult int32

const (
	// Unspecified result. This is an error.
	ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AUDIT_RESULT_UNSPECIFIED ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AuditResult = 0
	// Image is allowed.
	ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_ALLOW ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AuditResult = 1
	// Image is denied.
	ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_DENY ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AuditResult = 2
)

// Enum value maps for ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AuditResult.
var (
	ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AuditResult_name = map[int32]string{
		0: "AUDIT_RESULT_UNSPECIFIED",
		1: "ALLOW",
		2: "DENY",
	}
	ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AuditResult_value = map[string]int32{
		"AUDIT_RESULT_UNSPECIFIED": 0,
		"ALLOW":                    1,
		"DENY":                     2,
	}
)

func (x ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AuditResult) Enum() *ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AuditResult {
	p := new(ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AuditResult)
	*p = x
	return p
}

func (x ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AuditResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AuditResult) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_enumTypes[1].Descriptor()
}

func (ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AuditResult) Type() protoreflect.EnumType {
	return &file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_enumTypes[1]
}

func (x ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AuditResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AuditResult.Descriptor instead.
func (ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AuditResult) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

// Represents an auditing event from Continuous Validation.
type ContinuousValidationEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of CV event.
	//
	// Types that are assignable to EventType:
	//
	//	*ContinuousValidationEvent_PodEvent
	//	*ContinuousValidationEvent_UnsupportedPolicyEvent_
	EventType isContinuousValidationEvent_EventType `protobuf_oneof:"event_type"`
}

func (x *ContinuousValidationEvent) Reset() {
	*x = ContinuousValidationEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContinuousValidationEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContinuousValidationEvent) ProtoMessage() {}

func (x *ContinuousValidationEvent) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContinuousValidationEvent.ProtoReflect.Descriptor instead.
func (*ContinuousValidationEvent) Descriptor() ([]byte, []int) {
	return file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_rawDescGZIP(), []int{0}
}

func (m *ContinuousValidationEvent) GetEventType() isContinuousValidationEvent_EventType {
	if m != nil {
		return m.EventType
	}
	return nil
}

func (x *ContinuousValidationEvent) GetPodEvent() *ContinuousValidationEvent_ContinuousValidationPodEvent {
	if x, ok := x.GetEventType().(*ContinuousValidationEvent_PodEvent); ok {
		return x.PodEvent
	}
	return nil
}

func (x *ContinuousValidationEvent) GetUnsupportedPolicyEvent() *ContinuousValidationEvent_UnsupportedPolicyEvent {
	if x, ok := x.GetEventType().(*ContinuousValidationEvent_UnsupportedPolicyEvent_); ok {
		return x.UnsupportedPolicyEvent
	}
	return nil
}

type isContinuousValidationEvent_EventType interface {
	isContinuousValidationEvent_EventType()
}

type ContinuousValidationEvent_PodEvent struct {
	// Pod event.
	PodEvent *ContinuousValidationEvent_ContinuousValidationPodEvent `protobuf:"bytes,1,opt,name=pod_event,json=podEvent,proto3,oneof"`
}

type ContinuousValidationEvent_UnsupportedPolicyEvent_ struct {
	// Unsupported policy event.
	UnsupportedPolicyEvent *ContinuousValidationEvent_UnsupportedPolicyEvent `protobuf:"bytes,2,opt,name=unsupported_policy_event,json=unsupportedPolicyEvent,proto3,oneof"`
}

func (*ContinuousValidationEvent_PodEvent) isContinuousValidationEvent_EventType() {}

func (*ContinuousValidationEvent_UnsupportedPolicyEvent_) isContinuousValidationEvent_EventType() {}

// An auditing event for one Pod.
type ContinuousValidationEvent_ContinuousValidationPodEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The k8s namespace of the Pod.
	PodNamespace string `protobuf:"bytes,7,opt,name=pod_namespace,json=podNamespace,proto3" json:"pod_namespace,omitempty"`
	// The name of the Pod.
	Pod string `protobuf:"bytes,1,opt,name=pod,proto3" json:"pod,omitempty"`
	// Deploy time of the Pod from k8s.
	DeployTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=deploy_time,json=deployTime,proto3" json:"deploy_time,omitempty"`
	// Termination time of the Pod from k8s, or nothing if still running.
	EndTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Auditing verdict for this Pod.
	Verdict ContinuousValidationEvent_ContinuousValidationPodEvent_PolicyConformanceVerdict `protobuf:"varint,4,opt,name=verdict,proto3,enum=google.cloud.binaryauthorization.v1beta1.ContinuousValidationEvent_ContinuousValidationPodEvent_PolicyConformanceVerdict" json:"verdict,omitempty"`
	// List of images with auditing details.
	Images []*ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails `protobuf:"bytes,5,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *ContinuousValidationEvent_ContinuousValidationPodEvent) Reset() {
	*x = ContinuousValidationEvent_ContinuousValidationPodEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContinuousValidationEvent_ContinuousValidationPodEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContinuousValidationEvent_ContinuousValidationPodEvent) ProtoMessage() {}

func (x *ContinuousValidationEvent_ContinuousValidationPodEvent) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContinuousValidationEvent_ContinuousValidationPodEvent.ProtoReflect.Descriptor instead.
func (*ContinuousValidationEvent_ContinuousValidationPodEvent) Descriptor() ([]byte, []int) {
	return file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ContinuousValidationEvent_ContinuousValidationPodEvent) GetPodNamespace() string {
	if x != nil {
		return x.PodNamespace
	}
	return ""
}

func (x *ContinuousValidationEvent_ContinuousValidationPodEvent) GetPod() string {
	if x != nil {
		return x.Pod
	}
	return ""
}

func (x *ContinuousValidationEvent_ContinuousValidationPodEvent) GetDeployTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeployTime
	}
	return nil
}

func (x *ContinuousValidationEvent_ContinuousValidationPodEvent) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *ContinuousValidationEvent_ContinuousValidationPodEvent) GetVerdict() ContinuousValidationEvent_ContinuousValidationPodEvent_PolicyConformanceVerdict {
	if x != nil {
		return x.Verdict
	}
	return ContinuousValidationEvent_ContinuousValidationPodEvent_POLICY_CONFORMANCE_VERDICT_UNSPECIFIED
}

func (x *ContinuousValidationEvent_ContinuousValidationPodEvent) GetImages() []*ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails {
	if x != nil {
		return x.Images
	}
	return nil
}

// An event describing that the project policy is unsupported by CV.
type ContinuousValidationEvent_UnsupportedPolicyEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A description of the unsupported policy.
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ContinuousValidationEvent_UnsupportedPolicyEvent) Reset() {
	*x = ContinuousValidationEvent_UnsupportedPolicyEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContinuousValidationEvent_UnsupportedPolicyEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContinuousValidationEvent_UnsupportedPolicyEvent) ProtoMessage() {}

func (x *ContinuousValidationEvent_UnsupportedPolicyEvent) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContinuousValidationEvent_UnsupportedPolicyEvent.ProtoReflect.Descriptor instead.
func (*ContinuousValidationEvent_UnsupportedPolicyEvent) Descriptor() ([]byte, []int) {
	return file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ContinuousValidationEvent_UnsupportedPolicyEvent) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// Container image with auditing details.
type ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the image.
	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	// The result of the audit for this image.
	Result ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AuditResult `protobuf:"varint,2,opt,name=result,proto3,enum=google.cloud.binaryauthorization.v1beta1.ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AuditResult" json:"result,omitempty"`
	// Description of the above result.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails) Reset() {
	*x = ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails) ProtoMessage() {}

func (x *ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails.ProtoReflect.Descriptor instead.
func (*ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails) Descriptor() ([]byte, []int) {
	return file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails) GetResult() ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AuditResult {
	if x != nil {
		return x.Result
	}
	return ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AUDIT_RESULT_UNSPECIFIED
}

func (x *ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto protoreflect.FileDescriptor

var file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x69,
	0x6e, 0x75, 0x6f, 0x75, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6e,
	0x61, 0x72, 0x79, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x09, 0x0a, 0x19, 0x43, 0x6f,
	0x6e, 0x74, 0x69, 0x6e, 0x75, 0x6f, 0x75, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x7f, 0x0a, 0x09, 0x70, 0x6f, 0x64, 0x5f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x60, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x6f, 0x75, 0x73,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x6f, 0x75, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x08,
	0x70, 0x6f, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x96, 0x01, 0x0a, 0x18, 0x75, 0x6e, 0x73,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x5a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x6f, 0x75,
	0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x55, 0x6e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x16, 0x75, 0x6e, 0x73, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x1a, 0xe3, 0x06, 0x0a, 0x1c, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x6f, 0x75, 0x73,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x64, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6f, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x6f, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x93, 0x01,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x64, 0x69, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x69,
	0x6e, 0x75, 0x6f, 0x75, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x6f, 0x75, 0x73, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x64, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x63, 0x65, 0x56, 0x65, 0x72, 0x64, 0x69, 0x63, 0x74, 0x52, 0x07, 0x76, 0x65, 0x72, 0x64,
	0x69, 0x63, 0x74, 0x12, 0x85, 0x01, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x6f, 0x75, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e,
	0x75, 0x6f, 0x75, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f,
	0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x9c, 0x02, 0x0a, 0x0c,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x91, 0x01, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x6f, 0x75, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75,
	0x6f, 0x75, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x64,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x40, 0x0a, 0x0b, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c, 0x0a, 0x18, 0x41, 0x55, 0x44, 0x49, 0x54,
	0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x4e, 0x59, 0x10, 0x02, 0x22, 0x5b, 0x0a, 0x18, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x56,
	0x65, 0x72, 0x64, 0x69, 0x63, 0x74, 0x12, 0x2a, 0x0a, 0x26, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59,
	0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x56, 0x45, 0x52,
	0x44, 0x49, 0x43, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x56, 0x49, 0x4f, 0x4c, 0x41, 0x54, 0x45, 0x53, 0x5f, 0x50,
	0x4f, 0x4c, 0x49, 0x43, 0x59, 0x10, 0x01, 0x1a, 0x3a, 0x0a, 0x16, 0x55, 0x6e, 0x73, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x0c, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x42, 0xb6, 0x02, 0x0a, 0x2c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x42, 0x20, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x6f, 0x75, 0x73, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b,
	0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0xf8, 0x01, 0x01, 0xaa, 0x02, 0x28, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61,
	0x31, 0xca, 0x02, 0x28, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x2b, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x42, 0x69,
	0x6e, 0x61, 0x72, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_rawDescOnce sync.Once
	file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_rawDescData = file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_rawDesc
)

func file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_rawDescGZIP() []byte {
	file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_rawDescOnce.Do(func() {
		file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_rawDescData)
	})
	return file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_rawDescData
}

var file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_goTypes = []interface{}{
	(ContinuousValidationEvent_ContinuousValidationPodEvent_PolicyConformanceVerdict)(0), // 0: google.cloud.binaryauthorization.v1beta1.ContinuousValidationEvent.ContinuousValidationPodEvent.PolicyConformanceVerdict
	(ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AuditResult)(0), // 1: google.cloud.binaryauthorization.v1beta1.ContinuousValidationEvent.ContinuousValidationPodEvent.ImageDetails.AuditResult
	(*ContinuousValidationEvent)(nil),                                           // 2: google.cloud.binaryauthorization.v1beta1.ContinuousValidationEvent
	(*ContinuousValidationEvent_ContinuousValidationPodEvent)(nil),              // 3: google.cloud.binaryauthorization.v1beta1.ContinuousValidationEvent.ContinuousValidationPodEvent
	(*ContinuousValidationEvent_UnsupportedPolicyEvent)(nil),                    // 4: google.cloud.binaryauthorization.v1beta1.ContinuousValidationEvent.UnsupportedPolicyEvent
	(*ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails)(nil), // 5: google.cloud.binaryauthorization.v1beta1.ContinuousValidationEvent.ContinuousValidationPodEvent.ImageDetails
	(*timestamppb.Timestamp)(nil),                                               // 6: google.protobuf.Timestamp
}
var file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_depIdxs = []int32{
	3, // 0: google.cloud.binaryauthorization.v1beta1.ContinuousValidationEvent.pod_event:type_name -> google.cloud.binaryauthorization.v1beta1.ContinuousValidationEvent.ContinuousValidationPodEvent
	4, // 1: google.cloud.binaryauthorization.v1beta1.ContinuousValidationEvent.unsupported_policy_event:type_name -> google.cloud.binaryauthorization.v1beta1.ContinuousValidationEvent.UnsupportedPolicyEvent
	6, // 2: google.cloud.binaryauthorization.v1beta1.ContinuousValidationEvent.ContinuousValidationPodEvent.deploy_time:type_name -> google.protobuf.Timestamp
	6, // 3: google.cloud.binaryauthorization.v1beta1.ContinuousValidationEvent.ContinuousValidationPodEvent.end_time:type_name -> google.protobuf.Timestamp
	0, // 4: google.cloud.binaryauthorization.v1beta1.ContinuousValidationEvent.ContinuousValidationPodEvent.verdict:type_name -> google.cloud.binaryauthorization.v1beta1.ContinuousValidationEvent.ContinuousValidationPodEvent.PolicyConformanceVerdict
	5, // 5: google.cloud.binaryauthorization.v1beta1.ContinuousValidationEvent.ContinuousValidationPodEvent.images:type_name -> google.cloud.binaryauthorization.v1beta1.ContinuousValidationEvent.ContinuousValidationPodEvent.ImageDetails
	1, // 6: google.cloud.binaryauthorization.v1beta1.ContinuousValidationEvent.ContinuousValidationPodEvent.ImageDetails.result:type_name -> google.cloud.binaryauthorization.v1beta1.ContinuousValidationEvent.ContinuousValidationPodEvent.ImageDetails.AuditResult
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_init() }
func file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_init() {
	if File_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContinuousValidationEvent); i {
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
		file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContinuousValidationEvent_ContinuousValidationPodEvent); i {
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
		file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContinuousValidationEvent_UnsupportedPolicyEvent); i {
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
		file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails); i {
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
	file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ContinuousValidationEvent_PodEvent)(nil),
		(*ContinuousValidationEvent_UnsupportedPolicyEvent_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_goTypes,
		DependencyIndexes: file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_depIdxs,
		EnumInfos:         file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_enumTypes,
		MessageInfos:      file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_msgTypes,
	}.Build()
	File_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto = out.File
	file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_rawDesc = nil
	file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_goTypes = nil
	file_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto_depIdxs = nil
}
