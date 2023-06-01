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
// 	protoc        v4.23.1
// source: google/api/serviceusage/v1/resources.proto

package serviceusagepb

import (
	reflect "reflect"
	sync "sync"

	monitoredres "google.golang.org/genproto/googleapis/api/monitoredres"
	serviceconfig "google.golang.org/genproto/googleapis/api/serviceconfig"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	apipb "google.golang.org/protobuf/types/known/apipb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Whether or not a service has been enabled for use by a consumer.
type State int32

const (
	// The default value, which indicates that the enabled state of the service
	// is unspecified or not meaningful. Currently, all consumers other than
	// projects (such as folders and organizations) are always in this state.
	State_STATE_UNSPECIFIED State = 0
	// The service cannot be used by this consumer. It has either been explicitly
	// disabled, or has never been enabled.
	State_DISABLED State = 1
	// The service has been explicitly enabled for use by this consumer.
	State_ENABLED State = 2
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "DISABLED",
		2: "ENABLED",
	}
	State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"DISABLED":          1,
		"ENABLED":           2,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_api_serviceusage_v1_resources_proto_enumTypes[0].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_google_api_serviceusage_v1_resources_proto_enumTypes[0]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_google_api_serviceusage_v1_resources_proto_rawDescGZIP(), []int{0}
}

// A service that is available for use by the consumer.
type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the consumer and service.
	//
	// A valid name would be:
	// - projects/123/services/serviceusage.googleapis.com
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The resource name of the consumer.
	//
	// A valid name would be:
	// - projects/123
	Parent string `protobuf:"bytes,5,opt,name=parent,proto3" json:"parent,omitempty"`
	// The service configuration of the available service.
	// Some fields may be filtered out of the configuration in responses to
	// the `ListServices` method. These fields are present only in responses to
	// the `GetService` method.
	Config *ServiceConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	// Whether or not the service has been enabled for use by the consumer.
	State State `protobuf:"varint,4,opt,name=state,proto3,enum=google.api.serviceusage.v1.State" json:"state,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_serviceusage_v1_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_serviceusage_v1_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_google_api_serviceusage_v1_resources_proto_rawDescGZIP(), []int{0}
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Service) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *Service) GetConfig() *ServiceConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Service) GetState() State {
	if x != nil {
		return x.State
	}
	return State_STATE_UNSPECIFIED
}

// The configuration of the service.
type ServiceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The DNS address at which this service is available.
	//
	// An example DNS address would be:
	// `calendar.googleapis.com`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The product title for this service.
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// A list of API interfaces exported by this service. Contains only the names,
	// versions, and method names of the interfaces.
	Apis []*apipb.Api `protobuf:"bytes,3,rep,name=apis,proto3" json:"apis,omitempty"`
	// Additional API documentation. Contains only the summary and the
	// documentation URL.
	Documentation *serviceconfig.Documentation `protobuf:"bytes,6,opt,name=documentation,proto3" json:"documentation,omitempty"`
	// Quota configuration.
	Quota *serviceconfig.Quota `protobuf:"bytes,10,opt,name=quota,proto3" json:"quota,omitempty"`
	// Auth configuration. Contains only the OAuth rules.
	Authentication *serviceconfig.Authentication `protobuf:"bytes,11,opt,name=authentication,proto3" json:"authentication,omitempty"`
	// Configuration controlling usage of this service.
	Usage *serviceconfig.Usage `protobuf:"bytes,15,opt,name=usage,proto3" json:"usage,omitempty"`
	// Configuration for network endpoints. Contains only the names and aliases
	// of the endpoints.
	Endpoints []*serviceconfig.Endpoint `protobuf:"bytes,18,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	// Defines the monitored resources used by this service. This is required
	// by the [Service.monitoring][google.api.Service.monitoring] and [Service.logging][google.api.Service.logging] configurations.
	MonitoredResources []*monitoredres.MonitoredResourceDescriptor `protobuf:"bytes,25,rep,name=monitored_resources,json=monitoredResources,proto3" json:"monitored_resources,omitempty"`
	// Monitoring configuration.
	// This should not include the 'producer_destinations' field.
	Monitoring *serviceconfig.Monitoring `protobuf:"bytes,28,opt,name=monitoring,proto3" json:"monitoring,omitempty"`
}

func (x *ServiceConfig) Reset() {
	*x = ServiceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_serviceusage_v1_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceConfig) ProtoMessage() {}

func (x *ServiceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_serviceusage_v1_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceConfig.ProtoReflect.Descriptor instead.
func (*ServiceConfig) Descriptor() ([]byte, []int) {
	return file_google_api_serviceusage_v1_resources_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceConfig) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ServiceConfig) GetApis() []*apipb.Api {
	if x != nil {
		return x.Apis
	}
	return nil
}

func (x *ServiceConfig) GetDocumentation() *serviceconfig.Documentation {
	if x != nil {
		return x.Documentation
	}
	return nil
}

func (x *ServiceConfig) GetQuota() *serviceconfig.Quota {
	if x != nil {
		return x.Quota
	}
	return nil
}

func (x *ServiceConfig) GetAuthentication() *serviceconfig.Authentication {
	if x != nil {
		return x.Authentication
	}
	return nil
}

func (x *ServiceConfig) GetUsage() *serviceconfig.Usage {
	if x != nil {
		return x.Usage
	}
	return nil
}

func (x *ServiceConfig) GetEndpoints() []*serviceconfig.Endpoint {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *ServiceConfig) GetMonitoredResources() []*monitoredres.MonitoredResourceDescriptor {
	if x != nil {
		return x.MonitoredResources
	}
	return nil
}

func (x *ServiceConfig) GetMonitoring() *serviceconfig.Monitoring {
	if x != nil {
		return x.Monitoring
	}
	return nil
}

// The operation metadata returned for the batchend services operation.
type OperationMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The full name of the resources that this operation is directly
	// associated with.
	ResourceNames []string `protobuf:"bytes,2,rep,name=resource_names,json=resourceNames,proto3" json:"resource_names,omitempty"`
}

func (x *OperationMetadata) Reset() {
	*x = OperationMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_serviceusage_v1_resources_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationMetadata) ProtoMessage() {}

func (x *OperationMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_serviceusage_v1_resources_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationMetadata.ProtoReflect.Descriptor instead.
func (*OperationMetadata) Descriptor() ([]byte, []int) {
	return file_google_api_serviceusage_v1_resources_proto_rawDescGZIP(), []int{2}
}

func (x *OperationMetadata) GetResourceNames() []string {
	if x != nil {
		return x.ResourceNames
	}
	return nil
}

var File_google_api_serviceusage_v1_resources_proto protoreflect.FileDescriptor

var file_google_api_serviceusage_v1_resources_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64,
	0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12,
	0x41, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x37, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x80, 0x04, 0x0a, 0x0d,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x61, 0x70, 0x69, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x04, 0x61, 0x70, 0x69,
	0x73, 0x12, 0x3f, 0x0a, 0x0d, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x51,
	0x75, 0x6f, 0x74, 0x61, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x42, 0x0a, 0x0e, 0x61,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x27, 0x0a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x58, 0x0a, 0x13,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x18, 0x19, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x52, 0x12, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x0a, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x52, 0x0a, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x3a,
	0x0a, 0x11, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x2a, 0x39, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49,
	0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x41, 0x42,
	0x4c, 0x45, 0x44, 0x10, 0x02, 0x42, 0xd8, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x75, 0x73, 0x61, 0x67, 0x65, 0x70,
	0x62, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x75, 0x73, 0x61, 0x67, 0x65, 0x70, 0x62,
	0xaa, 0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x5c, 0x56, 0x31, 0xea, 0x02,
	0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_api_serviceusage_v1_resources_proto_rawDescOnce sync.Once
	file_google_api_serviceusage_v1_resources_proto_rawDescData = file_google_api_serviceusage_v1_resources_proto_rawDesc
)

func file_google_api_serviceusage_v1_resources_proto_rawDescGZIP() []byte {
	file_google_api_serviceusage_v1_resources_proto_rawDescOnce.Do(func() {
		file_google_api_serviceusage_v1_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_api_serviceusage_v1_resources_proto_rawDescData)
	})
	return file_google_api_serviceusage_v1_resources_proto_rawDescData
}

var file_google_api_serviceusage_v1_resources_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_api_serviceusage_v1_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_api_serviceusage_v1_resources_proto_goTypes = []interface{}{
	(State)(0),                                       // 0: google.api.serviceusage.v1.State
	(*Service)(nil),                                  // 1: google.api.serviceusage.v1.Service
	(*ServiceConfig)(nil),                            // 2: google.api.serviceusage.v1.ServiceConfig
	(*OperationMetadata)(nil),                        // 3: google.api.serviceusage.v1.OperationMetadata
	(*apipb.Api)(nil),                                // 4: google.protobuf.Api
	(*serviceconfig.Documentation)(nil),              // 5: google.api.Documentation
	(*serviceconfig.Quota)(nil),                      // 6: google.api.Quota
	(*serviceconfig.Authentication)(nil),             // 7: google.api.Authentication
	(*serviceconfig.Usage)(nil),                      // 8: google.api.Usage
	(*serviceconfig.Endpoint)(nil),                   // 9: google.api.Endpoint
	(*monitoredres.MonitoredResourceDescriptor)(nil), // 10: google.api.MonitoredResourceDescriptor
	(*serviceconfig.Monitoring)(nil),                 // 11: google.api.Monitoring
}
var file_google_api_serviceusage_v1_resources_proto_depIdxs = []int32{
	2,  // 0: google.api.serviceusage.v1.Service.config:type_name -> google.api.serviceusage.v1.ServiceConfig
	0,  // 1: google.api.serviceusage.v1.Service.state:type_name -> google.api.serviceusage.v1.State
	4,  // 2: google.api.serviceusage.v1.ServiceConfig.apis:type_name -> google.protobuf.Api
	5,  // 3: google.api.serviceusage.v1.ServiceConfig.documentation:type_name -> google.api.Documentation
	6,  // 4: google.api.serviceusage.v1.ServiceConfig.quota:type_name -> google.api.Quota
	7,  // 5: google.api.serviceusage.v1.ServiceConfig.authentication:type_name -> google.api.Authentication
	8,  // 6: google.api.serviceusage.v1.ServiceConfig.usage:type_name -> google.api.Usage
	9,  // 7: google.api.serviceusage.v1.ServiceConfig.endpoints:type_name -> google.api.Endpoint
	10, // 8: google.api.serviceusage.v1.ServiceConfig.monitored_resources:type_name -> google.api.MonitoredResourceDescriptor
	11, // 9: google.api.serviceusage.v1.ServiceConfig.monitoring:type_name -> google.api.Monitoring
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_google_api_serviceusage_v1_resources_proto_init() }
func file_google_api_serviceusage_v1_resources_proto_init() {
	if File_google_api_serviceusage_v1_resources_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_api_serviceusage_v1_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_google_api_serviceusage_v1_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceConfig); i {
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
		file_google_api_serviceusage_v1_resources_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationMetadata); i {
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
			RawDescriptor: file_google_api_serviceusage_v1_resources_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_api_serviceusage_v1_resources_proto_goTypes,
		DependencyIndexes: file_google_api_serviceusage_v1_resources_proto_depIdxs,
		EnumInfos:         file_google_api_serviceusage_v1_resources_proto_enumTypes,
		MessageInfos:      file_google_api_serviceusage_v1_resources_proto_msgTypes,
	}.Build()
	File_google_api_serviceusage_v1_resources_proto = out.File
	file_google_api_serviceusage_v1_resources_proto_rawDesc = nil
	file_google_api_serviceusage_v1_resources_proto_goTypes = nil
	file_google_api_serviceusage_v1_resources_proto_depIdxs = nil
}
