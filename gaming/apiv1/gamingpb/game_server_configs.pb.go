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
// source: google/cloud/gaming/v1/game_server_configs.proto

package gamingpb

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

// Request message for GameServerConfigsService.ListGameServerConfigs.
type ListGameServerConfigsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The parent resource name, in the following form:
	// `projects/{project}/locations/{location}/gameServerDeployments/{deployment}/configs/*`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. The maximum number of items to return.  If unspecified, server
	// will pick an appropriate default. Server may return fewer items than
	// requested. A caller should only rely on response's
	// [next_page_token][google.cloud.gaming.v1.ListGameServerConfigsResponse.next_page_token] to
	// determine if there are more GameServerConfigs left to be queried.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. The next_page_token value returned from a previous list request, if any.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Optional. The filter to apply to list results.
	Filter string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	// Optional. Specifies the ordering of results following syntax at
	// https://cloud.google.com/apis/design/design_patterns#sorting_order.
	OrderBy string `protobuf:"bytes,5,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
}

func (x *ListGameServerConfigsRequest) Reset() {
	*x = ListGameServerConfigsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_gaming_v1_game_server_configs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGameServerConfigsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGameServerConfigsRequest) ProtoMessage() {}

func (x *ListGameServerConfigsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gaming_v1_game_server_configs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGameServerConfigsRequest.ProtoReflect.Descriptor instead.
func (*ListGameServerConfigsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_gaming_v1_game_server_configs_proto_rawDescGZIP(), []int{0}
}

func (x *ListGameServerConfigsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListGameServerConfigsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListGameServerConfigsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListGameServerConfigsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *ListGameServerConfigsRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

// Response message for GameServerConfigsService.ListGameServerConfigs.
type ListGameServerConfigsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of game server configs.
	GameServerConfigs []*GameServerConfig `protobuf:"bytes,1,rep,name=game_server_configs,json=gameServerConfigs,proto3" json:"game_server_configs,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no more
	// results in the list.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// List of locations that could not be reached.
	Unreachable []string `protobuf:"bytes,4,rep,name=unreachable,proto3" json:"unreachable,omitempty"`
}

func (x *ListGameServerConfigsResponse) Reset() {
	*x = ListGameServerConfigsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_gaming_v1_game_server_configs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGameServerConfigsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGameServerConfigsResponse) ProtoMessage() {}

func (x *ListGameServerConfigsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gaming_v1_game_server_configs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGameServerConfigsResponse.ProtoReflect.Descriptor instead.
func (*ListGameServerConfigsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_gaming_v1_game_server_configs_proto_rawDescGZIP(), []int{1}
}

func (x *ListGameServerConfigsResponse) GetGameServerConfigs() []*GameServerConfig {
	if x != nil {
		return x.GameServerConfigs
	}
	return nil
}

func (x *ListGameServerConfigsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListGameServerConfigsResponse) GetUnreachable() []string {
	if x != nil {
		return x.Unreachable
	}
	return nil
}

// Request message for GameServerConfigsService.GetGameServerConfig.
type GetGameServerConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the game server config to retrieve, in the following form:
	// `projects/{project}/locations/{location}/gameServerDeployments/{deployment}/configs/{config}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetGameServerConfigRequest) Reset() {
	*x = GetGameServerConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_gaming_v1_game_server_configs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGameServerConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGameServerConfigRequest) ProtoMessage() {}

func (x *GetGameServerConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gaming_v1_game_server_configs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGameServerConfigRequest.ProtoReflect.Descriptor instead.
func (*GetGameServerConfigRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_gaming_v1_game_server_configs_proto_rawDescGZIP(), []int{2}
}

func (x *GetGameServerConfigRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request message for GameServerConfigsService.CreateGameServerConfig.
type CreateGameServerConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The parent resource name, in the following form:
	// `projects/{project}/locations/{location}/gameServerDeployments/{deployment}/`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The ID of the game server config resource to be created.
	ConfigId string `protobuf:"bytes,2,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	// Required. The game server config resource to be created.
	GameServerConfig *GameServerConfig `protobuf:"bytes,3,opt,name=game_server_config,json=gameServerConfig,proto3" json:"game_server_config,omitempty"`
}

func (x *CreateGameServerConfigRequest) Reset() {
	*x = CreateGameServerConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_gaming_v1_game_server_configs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGameServerConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGameServerConfigRequest) ProtoMessage() {}

func (x *CreateGameServerConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gaming_v1_game_server_configs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGameServerConfigRequest.ProtoReflect.Descriptor instead.
func (*CreateGameServerConfigRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_gaming_v1_game_server_configs_proto_rawDescGZIP(), []int{3}
}

func (x *CreateGameServerConfigRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateGameServerConfigRequest) GetConfigId() string {
	if x != nil {
		return x.ConfigId
	}
	return ""
}

func (x *CreateGameServerConfigRequest) GetGameServerConfig() *GameServerConfig {
	if x != nil {
		return x.GameServerConfig
	}
	return nil
}

// Request message for GameServerConfigsService.DeleteGameServerConfig.
type DeleteGameServerConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the game server config to delete, in the following form:
	// `projects/{project}/locations/{location}/gameServerDeployments/{deployment}/configs/{config}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteGameServerConfigRequest) Reset() {
	*x = DeleteGameServerConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_gaming_v1_game_server_configs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGameServerConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGameServerConfigRequest) ProtoMessage() {}

func (x *DeleteGameServerConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gaming_v1_game_server_configs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGameServerConfigRequest.ProtoReflect.Descriptor instead.
func (*DeleteGameServerConfigRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_gaming_v1_game_server_configs_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteGameServerConfigRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Autoscaling config for an Agones fleet.
type ScalingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the Scaling Config
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. Agones fleet autoscaler spec. Example spec:
	// https://agones.dev/site/docs/reference/fleetautoscaler/
	FleetAutoscalerSpec string `protobuf:"bytes,2,opt,name=fleet_autoscaler_spec,json=fleetAutoscalerSpec,proto3" json:"fleet_autoscaler_spec,omitempty"`
	// Labels used to identify the game server clusters to which this Agones
	// scaling config applies. A game server cluster is subject to this Agones
	// scaling config if its labels match any of the selector entries.
	Selectors []*LabelSelector `protobuf:"bytes,4,rep,name=selectors,proto3" json:"selectors,omitempty"`
	// The schedules to which this Scaling Config applies.
	Schedules []*Schedule `protobuf:"bytes,5,rep,name=schedules,proto3" json:"schedules,omitempty"`
}

func (x *ScalingConfig) Reset() {
	*x = ScalingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_gaming_v1_game_server_configs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScalingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScalingConfig) ProtoMessage() {}

func (x *ScalingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gaming_v1_game_server_configs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScalingConfig.ProtoReflect.Descriptor instead.
func (*ScalingConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_gaming_v1_game_server_configs_proto_rawDescGZIP(), []int{5}
}

func (x *ScalingConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ScalingConfig) GetFleetAutoscalerSpec() string {
	if x != nil {
		return x.FleetAutoscalerSpec
	}
	return ""
}

func (x *ScalingConfig) GetSelectors() []*LabelSelector {
	if x != nil {
		return x.Selectors
	}
	return nil
}

func (x *ScalingConfig) GetSchedules() []*Schedule {
	if x != nil {
		return x.Schedules
	}
	return nil
}

// Fleet configs for Agones.
type FleetConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Agones fleet spec. Example spec:
	// `https://agones.dev/site/docs/reference/fleet/`.
	FleetSpec string `protobuf:"bytes,1,opt,name=fleet_spec,json=fleetSpec,proto3" json:"fleet_spec,omitempty"`
	// The name of the FleetConfig.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FleetConfig) Reset() {
	*x = FleetConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_gaming_v1_game_server_configs_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FleetConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FleetConfig) ProtoMessage() {}

func (x *FleetConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gaming_v1_game_server_configs_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FleetConfig.ProtoReflect.Descriptor instead.
func (*FleetConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_gaming_v1_game_server_configs_proto_rawDescGZIP(), []int{6}
}

func (x *FleetConfig) GetFleetSpec() string {
	if x != nil {
		return x.FleetSpec
	}
	return ""
}

func (x *FleetConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// A game server config resource.
type GameServerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the game server config, in the following form:
	// `projects/{project}/locations/{location}/gameServerDeployments/{deployment}/configs/{config}`.
	// For example,
	// `projects/my-project/locations/global/gameServerDeployments/my-game/configs/my-config`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The creation time.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The last-modified time.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The labels associated with this game server config. Each label is a
	// key-value pair.
	Labels map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// FleetConfig contains a list of Agones fleet specs. Only one FleetConfig
	// is allowed.
	FleetConfigs []*FleetConfig `protobuf:"bytes,5,rep,name=fleet_configs,json=fleetConfigs,proto3" json:"fleet_configs,omitempty"`
	// The autoscaling settings.
	ScalingConfigs []*ScalingConfig `protobuf:"bytes,6,rep,name=scaling_configs,json=scalingConfigs,proto3" json:"scaling_configs,omitempty"`
	// The description of the game server config.
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *GameServerConfig) Reset() {
	*x = GameServerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_gaming_v1_game_server_configs_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameServerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameServerConfig) ProtoMessage() {}

func (x *GameServerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gaming_v1_game_server_configs_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameServerConfig.ProtoReflect.Descriptor instead.
func (*GameServerConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_gaming_v1_game_server_configs_proto_rawDescGZIP(), []int{7}
}

func (x *GameServerConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GameServerConfig) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *GameServerConfig) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *GameServerConfig) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *GameServerConfig) GetFleetConfigs() []*FleetConfig {
	if x != nil {
		return x.FleetConfigs
	}
	return nil
}

func (x *GameServerConfig) GetScalingConfigs() []*ScalingConfig {
	if x != nil {
		return x.ScalingConfigs
	}
	return nil
}

func (x *GameServerConfig) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_google_cloud_gaming_v1_game_server_configs_proto protoreflect.FileDescriptor

var file_google_cloud_gaming_v1_game_server_configs_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x01, 0x0a,
	0x1c, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a,
	0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x34, 0xe0,
	0x41, 0x02, 0xfa, 0x41, 0x2e, 0x12, 0x2c, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03,
	0xe0, 0x41, 0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x1b, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1e,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x22, 0xc3,
	0x01, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x58, 0x0a, 0x13, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x11, 0x67, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x63, 0x68, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x63, 0x68,
	0x61, 0x62, 0x6c, 0x65, 0x22, 0x66, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x48, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x34, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2e, 0x0a, 0x2c, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xec, 0x01, 0x0a,
	0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4c,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x34,
	0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2e, 0x12, 0x2c, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x09,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x5b,
	0x0a, 0x12, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x10, 0x67, 0x61, 0x6d, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x69, 0x0a, 0x1d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x34, 0xe0, 0x41, 0x02, 0xfa,
	0x41, 0x2e, 0x0a, 0x2c, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xe6, 0x01, 0x0a, 0x0d, 0x53, 0x63, 0x61, 0x6c, 0x69,
	0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x37, 0x0a, 0x15, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x72, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x13, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x41, 0x75, 0x74, 0x6f,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x43, 0x0a, 0x09, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12,
	0x3e, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x52, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x22,
	0x40, 0x0a, 0x0b, 0x46, 0x6c, 0x65, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1d,
	0x0a, 0x0a, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x81, 0x05, 0x0a, 0x10, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4c,
	0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x48, 0x0a, 0x0d,
	0x66, 0x6c, 0x65, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6c, 0x65,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0c, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x4e, 0x0a, 0x0f, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e,
	0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0e, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x3a, 0x8f, 0x01, 0xea, 0x41, 0x8b, 0x01, 0x0a, 0x2c, 0x67, 0x61, 0x6d, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x5b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x7d, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x7d, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f, 0x7b, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x7d, 0x42, 0x52, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x70, 0x62,
	0x3b, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_cloud_gaming_v1_game_server_configs_proto_rawDescOnce sync.Once
	file_google_cloud_gaming_v1_game_server_configs_proto_rawDescData = file_google_cloud_gaming_v1_game_server_configs_proto_rawDesc
)

func file_google_cloud_gaming_v1_game_server_configs_proto_rawDescGZIP() []byte {
	file_google_cloud_gaming_v1_game_server_configs_proto_rawDescOnce.Do(func() {
		file_google_cloud_gaming_v1_game_server_configs_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_gaming_v1_game_server_configs_proto_rawDescData)
	})
	return file_google_cloud_gaming_v1_game_server_configs_proto_rawDescData
}

var file_google_cloud_gaming_v1_game_server_configs_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_google_cloud_gaming_v1_game_server_configs_proto_goTypes = []interface{}{
	(*ListGameServerConfigsRequest)(nil),  // 0: google.cloud.gaming.v1.ListGameServerConfigsRequest
	(*ListGameServerConfigsResponse)(nil), // 1: google.cloud.gaming.v1.ListGameServerConfigsResponse
	(*GetGameServerConfigRequest)(nil),    // 2: google.cloud.gaming.v1.GetGameServerConfigRequest
	(*CreateGameServerConfigRequest)(nil), // 3: google.cloud.gaming.v1.CreateGameServerConfigRequest
	(*DeleteGameServerConfigRequest)(nil), // 4: google.cloud.gaming.v1.DeleteGameServerConfigRequest
	(*ScalingConfig)(nil),                 // 5: google.cloud.gaming.v1.ScalingConfig
	(*FleetConfig)(nil),                   // 6: google.cloud.gaming.v1.FleetConfig
	(*GameServerConfig)(nil),              // 7: google.cloud.gaming.v1.GameServerConfig
	nil,                                   // 8: google.cloud.gaming.v1.GameServerConfig.LabelsEntry
	(*LabelSelector)(nil),                 // 9: google.cloud.gaming.v1.LabelSelector
	(*Schedule)(nil),                      // 10: google.cloud.gaming.v1.Schedule
	(*timestamppb.Timestamp)(nil),         // 11: google.protobuf.Timestamp
}
var file_google_cloud_gaming_v1_game_server_configs_proto_depIdxs = []int32{
	7,  // 0: google.cloud.gaming.v1.ListGameServerConfigsResponse.game_server_configs:type_name -> google.cloud.gaming.v1.GameServerConfig
	7,  // 1: google.cloud.gaming.v1.CreateGameServerConfigRequest.game_server_config:type_name -> google.cloud.gaming.v1.GameServerConfig
	9,  // 2: google.cloud.gaming.v1.ScalingConfig.selectors:type_name -> google.cloud.gaming.v1.LabelSelector
	10, // 3: google.cloud.gaming.v1.ScalingConfig.schedules:type_name -> google.cloud.gaming.v1.Schedule
	11, // 4: google.cloud.gaming.v1.GameServerConfig.create_time:type_name -> google.protobuf.Timestamp
	11, // 5: google.cloud.gaming.v1.GameServerConfig.update_time:type_name -> google.protobuf.Timestamp
	8,  // 6: google.cloud.gaming.v1.GameServerConfig.labels:type_name -> google.cloud.gaming.v1.GameServerConfig.LabelsEntry
	6,  // 7: google.cloud.gaming.v1.GameServerConfig.fleet_configs:type_name -> google.cloud.gaming.v1.FleetConfig
	5,  // 8: google.cloud.gaming.v1.GameServerConfig.scaling_configs:type_name -> google.cloud.gaming.v1.ScalingConfig
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_google_cloud_gaming_v1_game_server_configs_proto_init() }
func file_google_cloud_gaming_v1_game_server_configs_proto_init() {
	if File_google_cloud_gaming_v1_game_server_configs_proto != nil {
		return
	}
	file_google_cloud_gaming_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_gaming_v1_game_server_configs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGameServerConfigsRequest); i {
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
		file_google_cloud_gaming_v1_game_server_configs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGameServerConfigsResponse); i {
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
		file_google_cloud_gaming_v1_game_server_configs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGameServerConfigRequest); i {
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
		file_google_cloud_gaming_v1_game_server_configs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGameServerConfigRequest); i {
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
		file_google_cloud_gaming_v1_game_server_configs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGameServerConfigRequest); i {
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
		file_google_cloud_gaming_v1_game_server_configs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScalingConfig); i {
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
		file_google_cloud_gaming_v1_game_server_configs_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FleetConfig); i {
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
		file_google_cloud_gaming_v1_game_server_configs_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameServerConfig); i {
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
			RawDescriptor: file_google_cloud_gaming_v1_game_server_configs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_gaming_v1_game_server_configs_proto_goTypes,
		DependencyIndexes: file_google_cloud_gaming_v1_game_server_configs_proto_depIdxs,
		MessageInfos:      file_google_cloud_gaming_v1_game_server_configs_proto_msgTypes,
	}.Build()
	File_google_cloud_gaming_v1_game_server_configs_proto = out.File
	file_google_cloud_gaming_v1_game_server_configs_proto_rawDesc = nil
	file_google_cloud_gaming_v1_game_server_configs_proto_goTypes = nil
	file_google_cloud_gaming_v1_game_server_configs_proto_depIdxs = nil
}
