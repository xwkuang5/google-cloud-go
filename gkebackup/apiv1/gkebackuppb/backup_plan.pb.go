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
// source: google/cloud/gkebackup/v1/backup_plan.proto

package gkebackuppb

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

// Defines the configuration and scheduling for a "line" of Backups.
type BackupPlan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The full name of the BackupPlan resource.
	// Format: `projects/*/locations/*/backupPlans/*`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Server generated global unique identifier of
	// [UUID](https://en.wikipedia.org/wiki/Universally_unique_identifier) format.
	Uid string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	// Output only. The timestamp when this BackupPlan resource was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The timestamp when this BackupPlan resource was last
	// updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// User specified descriptive string for this BackupPlan.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Required. Immutable. The source cluster from which Backups will be created
	// via this BackupPlan. Valid formats:
	//
	// - `projects/*/locations/*/clusters/*`
	// - `projects/*/zones/*/clusters/*`
	Cluster string `protobuf:"bytes,6,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// RetentionPolicy governs lifecycle of Backups created under this plan.
	RetentionPolicy *BackupPlan_RetentionPolicy `protobuf:"bytes,7,opt,name=retention_policy,json=retentionPolicy,proto3" json:"retention_policy,omitempty"`
	// A set of custom labels supplied by user.
	Labels map[string]string `protobuf:"bytes,8,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Defines a schedule for automatic Backup creation via this BackupPlan.
	BackupSchedule *BackupPlan_Schedule `protobuf:"bytes,9,opt,name=backup_schedule,json=backupSchedule,proto3" json:"backup_schedule,omitempty"`
	// Output only. `etag` is used for optimistic concurrency control as a way to
	// help prevent simultaneous updates of a backup plan from overwriting each
	// other. It is strongly suggested that systems make use of the 'etag' in the
	// read-modify-write cycle to perform BackupPlan updates in order to avoid
	// race conditions: An `etag` is returned in the response to `GetBackupPlan`,
	// and systems are expected to put that etag in the request to
	// `UpdateBackupPlan` or `DeleteBackupPlan` to ensure that their change
	// will be applied to the same version of the resource.
	Etag string `protobuf:"bytes,10,opt,name=etag,proto3" json:"etag,omitempty"`
	// This flag indicates whether this BackupPlan has been deactivated.
	// Setting this field to True locks the BackupPlan such that no further
	// updates will be allowed (except deletes), including the deactivated field
	// itself. It also prevents any new Backups from being created via this
	// BackupPlan (including scheduled Backups).
	//
	// Default: False
	Deactivated bool `protobuf:"varint,11,opt,name=deactivated,proto3" json:"deactivated,omitempty"`
	// Defines the configuration of Backups created via this BackupPlan.
	BackupConfig *BackupPlan_BackupConfig `protobuf:"bytes,12,opt,name=backup_config,json=backupConfig,proto3" json:"backup_config,omitempty"`
	// Output only. The number of Kubernetes Pods backed up in the
	// last successful Backup created via this BackupPlan.
	ProtectedPodCount int32 `protobuf:"varint,13,opt,name=protected_pod_count,json=protectedPodCount,proto3" json:"protected_pod_count,omitempty"`
}

func (x *BackupPlan) Reset() {
	*x = BackupPlan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_gkebackup_v1_backup_plan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupPlan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupPlan) ProtoMessage() {}

func (x *BackupPlan) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gkebackup_v1_backup_plan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupPlan.ProtoReflect.Descriptor instead.
func (*BackupPlan) Descriptor() ([]byte, []int) {
	return file_google_cloud_gkebackup_v1_backup_plan_proto_rawDescGZIP(), []int{0}
}

func (x *BackupPlan) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BackupPlan) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *BackupPlan) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *BackupPlan) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *BackupPlan) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BackupPlan) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *BackupPlan) GetRetentionPolicy() *BackupPlan_RetentionPolicy {
	if x != nil {
		return x.RetentionPolicy
	}
	return nil
}

func (x *BackupPlan) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *BackupPlan) GetBackupSchedule() *BackupPlan_Schedule {
	if x != nil {
		return x.BackupSchedule
	}
	return nil
}

func (x *BackupPlan) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *BackupPlan) GetDeactivated() bool {
	if x != nil {
		return x.Deactivated
	}
	return false
}

func (x *BackupPlan) GetBackupConfig() *BackupPlan_BackupConfig {
	if x != nil {
		return x.BackupConfig
	}
	return nil
}

func (x *BackupPlan) GetProtectedPodCount() int32 {
	if x != nil {
		return x.ProtectedPodCount
	}
	return 0
}

// RetentionPolicy defines a Backup retention policy for a BackupPlan.
type BackupPlan_RetentionPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Minimum age for Backups created via this BackupPlan (in days).
	// This field MUST be an integer value between 0-90 (inclusive).
	// A Backup created under this BackupPlan will NOT be deletable until it
	// reaches Backup's (create_time + backup_delete_lock_days).
	// Updating this field of a BackupPlan does NOT affect existing Backups
	// under it. Backups created AFTER a successful update will inherit
	// the new value.
	//
	// Default: 0 (no delete blocking)
	BackupDeleteLockDays int32 `protobuf:"varint,1,opt,name=backup_delete_lock_days,json=backupDeleteLockDays,proto3" json:"backup_delete_lock_days,omitempty"`
	// The default maximum age of a Backup created via this BackupPlan.
	// This field MUST be an integer value >= 0 and <= 365.
	// If specified, a Backup created under this BackupPlan will be
	// automatically deleted after its age reaches (create_time +
	// backup_retain_days).
	// If not specified, Backups created under this BackupPlan will NOT be
	// subject to automatic deletion.
	// Updating this field does NOT affect existing Backups under it. Backups
	// created AFTER a successful update will automatically pick up the new
	// value.
	// NOTE: backup_retain_days must be >=
	// [backup_delete_lock_days][google.cloud.gkebackup.v1.BackupPlan.RetentionPolicy.backup_delete_lock_days].
	// If
	// [cron_schedule][google.cloud.gkebackup.v1.BackupPlan.Schedule.cron_schedule]
	// is defined, then this must be
	// <= 360 * the creation interval.
	//
	// Default: 0 (no automatic deletion)
	BackupRetainDays int32 `protobuf:"varint,2,opt,name=backup_retain_days,json=backupRetainDays,proto3" json:"backup_retain_days,omitempty"`
	// This flag denotes whether the retention policy of this BackupPlan is
	// locked.  If set to True, no further update is allowed on this policy,
	// including the `locked` field itself.
	//
	// Default: False
	Locked bool `protobuf:"varint,3,opt,name=locked,proto3" json:"locked,omitempty"`
}

func (x *BackupPlan_RetentionPolicy) Reset() {
	*x = BackupPlan_RetentionPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_gkebackup_v1_backup_plan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupPlan_RetentionPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupPlan_RetentionPolicy) ProtoMessage() {}

func (x *BackupPlan_RetentionPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gkebackup_v1_backup_plan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupPlan_RetentionPolicy.ProtoReflect.Descriptor instead.
func (*BackupPlan_RetentionPolicy) Descriptor() ([]byte, []int) {
	return file_google_cloud_gkebackup_v1_backup_plan_proto_rawDescGZIP(), []int{0, 0}
}

func (x *BackupPlan_RetentionPolicy) GetBackupDeleteLockDays() int32 {
	if x != nil {
		return x.BackupDeleteLockDays
	}
	return 0
}

func (x *BackupPlan_RetentionPolicy) GetBackupRetainDays() int32 {
	if x != nil {
		return x.BackupRetainDays
	}
	return 0
}

func (x *BackupPlan_RetentionPolicy) GetLocked() bool {
	if x != nil {
		return x.Locked
	}
	return false
}

// Schedule defines scheduling parameters for automatically creating Backups
// via this BackupPlan.
type BackupPlan_Schedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A standard [cron](https://wikipedia.com/wiki/cron) string that defines a
	// repeating schedule for creating Backups via this BackupPlan. If this is
	// defined, then
	// [backup_retain_days][google.cloud.gkebackup.v1.BackupPlan.RetentionPolicy.backup_retain_days]
	// must also be defined.
	//
	// Default (empty): no automatic backup creation will occur.
	CronSchedule string `protobuf:"bytes,1,opt,name=cron_schedule,json=cronSchedule,proto3" json:"cron_schedule,omitempty"`
	// This flag denotes whether automatic Backup creation is paused for this
	// BackupPlan.
	//
	// Default: False
	Paused bool `protobuf:"varint,2,opt,name=paused,proto3" json:"paused,omitempty"`
}

func (x *BackupPlan_Schedule) Reset() {
	*x = BackupPlan_Schedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_gkebackup_v1_backup_plan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupPlan_Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupPlan_Schedule) ProtoMessage() {}

func (x *BackupPlan_Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gkebackup_v1_backup_plan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupPlan_Schedule.ProtoReflect.Descriptor instead.
func (*BackupPlan_Schedule) Descriptor() ([]byte, []int) {
	return file_google_cloud_gkebackup_v1_backup_plan_proto_rawDescGZIP(), []int{0, 1}
}

func (x *BackupPlan_Schedule) GetCronSchedule() string {
	if x != nil {
		return x.CronSchedule
	}
	return ""
}

func (x *BackupPlan_Schedule) GetPaused() bool {
	if x != nil {
		return x.Paused
	}
	return false
}

// BackupConfig defines the configuration of Backups created via this
// BackupPlan.
type BackupPlan_BackupConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This defines the "scope" of the Backup - which namespaced
	// resources in the cluster will be included in a Backup.
	// Exactly one of the fields of backup_scope MUST be specified.
	//
	// Types that are assignable to BackupScope:
	//	*BackupPlan_BackupConfig_AllNamespaces
	//	*BackupPlan_BackupConfig_SelectedNamespaces
	//	*BackupPlan_BackupConfig_SelectedApplications
	BackupScope isBackupPlan_BackupConfig_BackupScope `protobuf_oneof:"backup_scope"`
	// This flag specifies whether volume data should be backed up when
	// PVCs are included in the scope of a Backup.
	//
	// Default: False
	IncludeVolumeData bool `protobuf:"varint,4,opt,name=include_volume_data,json=includeVolumeData,proto3" json:"include_volume_data,omitempty"`
	// This flag specifies whether Kubernetes Secret resources should be
	// included when they fall into the scope of Backups.
	//
	// Default: False
	IncludeSecrets bool `protobuf:"varint,5,opt,name=include_secrets,json=includeSecrets,proto3" json:"include_secrets,omitempty"`
	// This defines a customer managed encryption key that will be used to
	// encrypt the "config" portion (the Kubernetes resources) of Backups
	// created via this plan.
	//
	// Default (empty): Config backup artifacts will not be encrypted.
	EncryptionKey *EncryptionKey `protobuf:"bytes,6,opt,name=encryption_key,json=encryptionKey,proto3" json:"encryption_key,omitempty"`
}

func (x *BackupPlan_BackupConfig) Reset() {
	*x = BackupPlan_BackupConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_gkebackup_v1_backup_plan_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupPlan_BackupConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupPlan_BackupConfig) ProtoMessage() {}

func (x *BackupPlan_BackupConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_gkebackup_v1_backup_plan_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupPlan_BackupConfig.ProtoReflect.Descriptor instead.
func (*BackupPlan_BackupConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_gkebackup_v1_backup_plan_proto_rawDescGZIP(), []int{0, 2}
}

func (m *BackupPlan_BackupConfig) GetBackupScope() isBackupPlan_BackupConfig_BackupScope {
	if m != nil {
		return m.BackupScope
	}
	return nil
}

func (x *BackupPlan_BackupConfig) GetAllNamespaces() bool {
	if x, ok := x.GetBackupScope().(*BackupPlan_BackupConfig_AllNamespaces); ok {
		return x.AllNamespaces
	}
	return false
}

func (x *BackupPlan_BackupConfig) GetSelectedNamespaces() *Namespaces {
	if x, ok := x.GetBackupScope().(*BackupPlan_BackupConfig_SelectedNamespaces); ok {
		return x.SelectedNamespaces
	}
	return nil
}

func (x *BackupPlan_BackupConfig) GetSelectedApplications() *NamespacedNames {
	if x, ok := x.GetBackupScope().(*BackupPlan_BackupConfig_SelectedApplications); ok {
		return x.SelectedApplications
	}
	return nil
}

func (x *BackupPlan_BackupConfig) GetIncludeVolumeData() bool {
	if x != nil {
		return x.IncludeVolumeData
	}
	return false
}

func (x *BackupPlan_BackupConfig) GetIncludeSecrets() bool {
	if x != nil {
		return x.IncludeSecrets
	}
	return false
}

func (x *BackupPlan_BackupConfig) GetEncryptionKey() *EncryptionKey {
	if x != nil {
		return x.EncryptionKey
	}
	return nil
}

type isBackupPlan_BackupConfig_BackupScope interface {
	isBackupPlan_BackupConfig_BackupScope()
}

type BackupPlan_BackupConfig_AllNamespaces struct {
	// If True, include all namespaced resources
	AllNamespaces bool `protobuf:"varint,1,opt,name=all_namespaces,json=allNamespaces,proto3,oneof"`
}

type BackupPlan_BackupConfig_SelectedNamespaces struct {
	// If set, include just the resources in the listed namespaces.
	SelectedNamespaces *Namespaces `protobuf:"bytes,2,opt,name=selected_namespaces,json=selectedNamespaces,proto3,oneof"`
}

type BackupPlan_BackupConfig_SelectedApplications struct {
	// If set, include just the resources referenced by the listed
	// ProtectedApplications.
	SelectedApplications *NamespacedNames `protobuf:"bytes,3,opt,name=selected_applications,json=selectedApplications,proto3,oneof"`
}

func (*BackupPlan_BackupConfig_AllNamespaces) isBackupPlan_BackupConfig_BackupScope() {}

func (*BackupPlan_BackupConfig_SelectedNamespaces) isBackupPlan_BackupConfig_BackupScope() {}

func (*BackupPlan_BackupConfig_SelectedApplications) isBackupPlan_BackupConfig_BackupScope() {}

var File_google_cloud_gkebackup_v1_backup_plan_proto protoreflect.FileDescriptor

var file_google_cloud_gkebackup_v1_backup_plan_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67,
	0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6b, 0x65, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x67, 0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x0c,
	0x0a, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x17, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x40, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x2b, 0xe0, 0x41, 0x05, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x22, 0x0a, 0x20,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x60, 0x0a, 0x10, 0x72, 0x65, 0x74,
	0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x50, 0x6c, 0x61, 0x6e, 0x2e, 0x52, 0x65, 0x74, 0x65, 0x6e,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0f, 0x72, 0x65, 0x74, 0x65,
	0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x49, 0x0a, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6b, 0x65, 0x62, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x50, 0x6c,
	0x61, 0x6e, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x57, 0x0a, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67,
	0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x50, 0x6c, 0x61, 0x6e, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52,
	0x0e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12,
	0x17, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x64,
	0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x12, 0x57, 0x0a, 0x0d, 0x62, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x67, 0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x50, 0x6c, 0x61, 0x6e, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0c, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x33, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x5f, 0x70, 0x6f, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x50, 0x6f, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x8e, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x74,
	0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x35, 0x0a, 0x17,
	0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x44,
	0x61, 0x79, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x72, 0x65,
	0x74, 0x61, 0x69, 0x6e, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x10, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6e, 0x44, 0x61, 0x79,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x1a, 0x47, 0x0a, 0x08, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x72, 0x6f, 0x6e, 0x5f, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x72,
	0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61,
	0x75, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x61, 0x75, 0x73,
	0x65, 0x64, 0x1a, 0xae, 0x03, 0x0a, 0x0c, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x27, 0x0a, 0x0e, 0x61, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0d, 0x61,
	0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x58, 0x0a, 0x13,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x48, 0x00, 0x52, 0x12, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x61, 0x0a, 0x15, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x48, 0x00, 0x52, 0x14, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x69, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0e, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x12, 0x4f, 0x0a, 0x0e, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6b, 0x65, 0x62, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x4b, 0x65, 0x79, 0x52, 0x0d, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x4b, 0x65, 0x79, 0x42, 0x0e, 0x0a, 0x0c, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x6b,
	0xea, 0x41, 0x68, 0x0a, 0x23, 0x67, 0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x41, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x7d, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x2f, 0x7b, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x7d, 0x42, 0xc6, 0x01, 0x0a, 0x1d,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x67, 0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x42,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x50, 0x6c, 0x61, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2f,
	0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x67, 0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x70,
	0x62, 0x3b, 0x67, 0x6b, 0x65, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x70, 0x62, 0xaa, 0x02, 0x19,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x47, 0x6b, 0x65,
	0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x47, 0x6b, 0x65, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x47, 0x6b, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_gkebackup_v1_backup_plan_proto_rawDescOnce sync.Once
	file_google_cloud_gkebackup_v1_backup_plan_proto_rawDescData = file_google_cloud_gkebackup_v1_backup_plan_proto_rawDesc
)

func file_google_cloud_gkebackup_v1_backup_plan_proto_rawDescGZIP() []byte {
	file_google_cloud_gkebackup_v1_backup_plan_proto_rawDescOnce.Do(func() {
		file_google_cloud_gkebackup_v1_backup_plan_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_gkebackup_v1_backup_plan_proto_rawDescData)
	})
	return file_google_cloud_gkebackup_v1_backup_plan_proto_rawDescData
}

var file_google_cloud_gkebackup_v1_backup_plan_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_gkebackup_v1_backup_plan_proto_goTypes = []interface{}{
	(*BackupPlan)(nil),                 // 0: google.cloud.gkebackup.v1.BackupPlan
	(*BackupPlan_RetentionPolicy)(nil), // 1: google.cloud.gkebackup.v1.BackupPlan.RetentionPolicy
	(*BackupPlan_Schedule)(nil),        // 2: google.cloud.gkebackup.v1.BackupPlan.Schedule
	(*BackupPlan_BackupConfig)(nil),    // 3: google.cloud.gkebackup.v1.BackupPlan.BackupConfig
	nil,                                // 4: google.cloud.gkebackup.v1.BackupPlan.LabelsEntry
	(*timestamppb.Timestamp)(nil),      // 5: google.protobuf.Timestamp
	(*Namespaces)(nil),                 // 6: google.cloud.gkebackup.v1.Namespaces
	(*NamespacedNames)(nil),            // 7: google.cloud.gkebackup.v1.NamespacedNames
	(*EncryptionKey)(nil),              // 8: google.cloud.gkebackup.v1.EncryptionKey
}
var file_google_cloud_gkebackup_v1_backup_plan_proto_depIdxs = []int32{
	5, // 0: google.cloud.gkebackup.v1.BackupPlan.create_time:type_name -> google.protobuf.Timestamp
	5, // 1: google.cloud.gkebackup.v1.BackupPlan.update_time:type_name -> google.protobuf.Timestamp
	1, // 2: google.cloud.gkebackup.v1.BackupPlan.retention_policy:type_name -> google.cloud.gkebackup.v1.BackupPlan.RetentionPolicy
	4, // 3: google.cloud.gkebackup.v1.BackupPlan.labels:type_name -> google.cloud.gkebackup.v1.BackupPlan.LabelsEntry
	2, // 4: google.cloud.gkebackup.v1.BackupPlan.backup_schedule:type_name -> google.cloud.gkebackup.v1.BackupPlan.Schedule
	3, // 5: google.cloud.gkebackup.v1.BackupPlan.backup_config:type_name -> google.cloud.gkebackup.v1.BackupPlan.BackupConfig
	6, // 6: google.cloud.gkebackup.v1.BackupPlan.BackupConfig.selected_namespaces:type_name -> google.cloud.gkebackup.v1.Namespaces
	7, // 7: google.cloud.gkebackup.v1.BackupPlan.BackupConfig.selected_applications:type_name -> google.cloud.gkebackup.v1.NamespacedNames
	8, // 8: google.cloud.gkebackup.v1.BackupPlan.BackupConfig.encryption_key:type_name -> google.cloud.gkebackup.v1.EncryptionKey
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_google_cloud_gkebackup_v1_backup_plan_proto_init() }
func file_google_cloud_gkebackup_v1_backup_plan_proto_init() {
	if File_google_cloud_gkebackup_v1_backup_plan_proto != nil {
		return
	}
	file_google_cloud_gkebackup_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_gkebackup_v1_backup_plan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupPlan); i {
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
		file_google_cloud_gkebackup_v1_backup_plan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupPlan_RetentionPolicy); i {
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
		file_google_cloud_gkebackup_v1_backup_plan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupPlan_Schedule); i {
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
		file_google_cloud_gkebackup_v1_backup_plan_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupPlan_BackupConfig); i {
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
	file_google_cloud_gkebackup_v1_backup_plan_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*BackupPlan_BackupConfig_AllNamespaces)(nil),
		(*BackupPlan_BackupConfig_SelectedNamespaces)(nil),
		(*BackupPlan_BackupConfig_SelectedApplications)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_gkebackup_v1_backup_plan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_gkebackup_v1_backup_plan_proto_goTypes,
		DependencyIndexes: file_google_cloud_gkebackup_v1_backup_plan_proto_depIdxs,
		MessageInfos:      file_google_cloud_gkebackup_v1_backup_plan_proto_msgTypes,
	}.Build()
	File_google_cloud_gkebackup_v1_backup_plan_proto = out.File
	file_google_cloud_gkebackup_v1_backup_plan_proto_rawDesc = nil
	file_google_cloud_gkebackup_v1_backup_plan_proto_goTypes = nil
	file_google_cloud_gkebackup_v1_backup_plan_proto_depIdxs = nil
}
