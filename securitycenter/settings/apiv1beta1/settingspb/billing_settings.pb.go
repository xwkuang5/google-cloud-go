// Copyright 2020 Google LLC
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
// source: google/cloud/securitycenter/settings/v1beta1/billing_settings.proto

package settingspb

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

// Billing tier options
type BillingTier int32

const (
	// Default value. This value is unused.
	BillingTier_BILLING_TIER_UNSPECIFIED BillingTier = 0
	// The standard billing tier.
	BillingTier_STANDARD BillingTier = 1
	// The premium billing tier.
	BillingTier_PREMIUM BillingTier = 2
)

// Enum value maps for BillingTier.
var (
	BillingTier_name = map[int32]string{
		0: "BILLING_TIER_UNSPECIFIED",
		1: "STANDARD",
		2: "PREMIUM",
	}
	BillingTier_value = map[string]int32{
		"BILLING_TIER_UNSPECIFIED": 0,
		"STANDARD":                 1,
		"PREMIUM":                  2,
	}
)

func (x BillingTier) Enum() *BillingTier {
	p := new(BillingTier)
	*p = x
	return p
}

func (x BillingTier) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BillingTier) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_enumTypes[0].Descriptor()
}

func (BillingTier) Type() protoreflect.EnumType {
	return &file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_enumTypes[0]
}

func (x BillingTier) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BillingTier.Descriptor instead.
func (BillingTier) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_rawDescGZIP(), []int{0}
}

// Billing type
type BillingType int32

const (
	// Default billing type
	BillingType_BILLING_TYPE_UNSPECIFIED BillingType = 0
	// Subscription for Premium billing tier
	BillingType_SUBSCRIPTION BillingType = 1
	// Trial subscription for Premium billing tier
	BillingType_TRIAL_SUBSCRIPTION BillingType = 2
	// Alpha customer for Premium billing tier
	BillingType_ALPHA BillingType = 3
)

// Enum value maps for BillingType.
var (
	BillingType_name = map[int32]string{
		0: "BILLING_TYPE_UNSPECIFIED",
		1: "SUBSCRIPTION",
		2: "TRIAL_SUBSCRIPTION",
		3: "ALPHA",
	}
	BillingType_value = map[string]int32{
		"BILLING_TYPE_UNSPECIFIED": 0,
		"SUBSCRIPTION":             1,
		"TRIAL_SUBSCRIPTION":       2,
		"ALPHA":                    3,
	}
)

func (x BillingType) Enum() *BillingType {
	p := new(BillingType)
	*p = x
	return p
}

func (x BillingType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BillingType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_enumTypes[1].Descriptor()
}

func (BillingType) Type() protoreflect.EnumType {
	return &file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_enumTypes[1]
}

func (x BillingType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BillingType.Descriptor instead.
func (BillingType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_rawDescGZIP(), []int{1}
}

// Billing settings
type BillingSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Billing tier selected by customer
	BillingTier BillingTier `protobuf:"varint,1,opt,name=billing_tier,json=billingTier,proto3,enum=google.cloud.securitycenter.settings.v1beta1.BillingTier" json:"billing_tier,omitempty"`
	// Output only. Type of billing method
	BillingType BillingType `protobuf:"varint,2,opt,name=billing_type,json=billingType,proto3,enum=google.cloud.securitycenter.settings.v1beta1.BillingType" json:"billing_type,omitempty"`
	// Output only. The absolute point in time when the subscription became effective.
	// Can be compared to expire_time value to determine full contract duration
	StartTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Output only. The absolute point in time when the subscription expires.
	//
	// If this field is populated and billing_tier is STANDARD, this is
	// indication of a point in the _past_ when a PREMIUM access ended.
	ExpireTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
}

func (x *BillingSettings) Reset() {
	*x = BillingSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingSettings) ProtoMessage() {}

func (x *BillingSettings) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingSettings.ProtoReflect.Descriptor instead.
func (*BillingSettings) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_rawDescGZIP(), []int{0}
}

func (x *BillingSettings) GetBillingTier() BillingTier {
	if x != nil {
		return x.BillingTier
	}
	return BillingTier_BILLING_TIER_UNSPECIFIED
}

func (x *BillingSettings) GetBillingType() BillingType {
	if x != nil {
		return x.BillingType
	}
	return BillingType_BILLING_TYPE_UNSPECIFIED
}

func (x *BillingSettings) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *BillingSettings) GetExpireTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpireTime
	}
	return nil
}

var File_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_rawDesc = []byte{
	0x0a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x02, 0x0a, 0x0f, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x61, 0x0a, 0x0c, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x42,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x65, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0b, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x65, 0x72, 0x12, 0x61, 0x0a, 0x0c,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x0b, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x3e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x40, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x2a, 0x46, 0x0a, 0x0b, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x65, 0x72,
	0x12, 0x1c, 0x0a, 0x18, 0x42, 0x49, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x49, 0x45, 0x52,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c,
	0x0a, 0x08, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x50, 0x52, 0x45, 0x4d, 0x49, 0x55, 0x4d, 0x10, 0x02, 0x2a, 0x60, 0x0a, 0x0b, 0x42, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x42, 0x49, 0x4c, 0x4c,
	0x49, 0x4e, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52,
	0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x52, 0x49, 0x41,
	0x4c, 0x5f, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02,
	0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c, 0x50, 0x48, 0x41, 0x10, 0x03, 0x42, 0xac, 0x02, 0x0a, 0x30,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x42, 0x14, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x70, 0x62, 0x3b, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x70, 0x62, 0xf8, 0x01, 0x01, 0xaa, 0x02, 0x2c, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x2c, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x5c,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x30, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_rawDescData = file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_rawDesc
)

func file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_rawDescData
}

var file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_goTypes = []interface{}{
	(BillingTier)(0),              // 0: google.cloud.securitycenter.settings.v1beta1.BillingTier
	(BillingType)(0),              // 1: google.cloud.securitycenter.settings.v1beta1.BillingType
	(*BillingSettings)(nil),       // 2: google.cloud.securitycenter.settings.v1beta1.BillingSettings
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_depIdxs = []int32{
	0, // 0: google.cloud.securitycenter.settings.v1beta1.BillingSettings.billing_tier:type_name -> google.cloud.securitycenter.settings.v1beta1.BillingTier
	1, // 1: google.cloud.securitycenter.settings.v1beta1.BillingSettings.billing_type:type_name -> google.cloud.securitycenter.settings.v1beta1.BillingType
	3, // 2: google.cloud.securitycenter.settings.v1beta1.BillingSettings.start_time:type_name -> google.protobuf.Timestamp
	3, // 3: google.cloud.securitycenter.settings.v1beta1.BillingSettings.expire_time:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_init() }
func file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_init() {
	if File_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillingSettings); i {
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
			RawDescriptor: file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_depIdxs,
		EnumInfos:         file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_enumTypes,
		MessageInfos:      file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto = out.File
	file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_rawDesc = nil
	file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_goTypes = nil
	file_google_cloud_securitycenter_settings_v1beta1_billing_settings_proto_depIdxs = nil
}
