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
// source: google/cloud/texttospeech/v1/cloud_tts_lrs.proto

package texttospeechpb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// The top-level message sent by the client for the
// `SynthesizeLongAudio` method.
type SynthesizeLongAudioRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource states of the request in the form of
	// `projects/*/locations/*`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The Synthesizer requires either plain text or SSML as input.
	// While Long Audio is in preview, SSML is temporarily unsupported.
	Input *SynthesisInput `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
	// Required. The configuration of the synthesized audio.
	AudioConfig *AudioConfig `protobuf:"bytes,3,opt,name=audio_config,json=audioConfig,proto3" json:"audio_config,omitempty"`
	// Required. Specifies a Cloud Storage URI for the synthesis results. Must be
	// specified in the format: `gs://bucket_name/object_name`, and the bucket
	// must already exist.
	OutputGcsUri string `protobuf:"bytes,4,opt,name=output_gcs_uri,json=outputGcsUri,proto3" json:"output_gcs_uri,omitempty"`
	// Required. The desired voice of the synthesized audio.
	Voice *VoiceSelectionParams `protobuf:"bytes,5,opt,name=voice,proto3" json:"voice,omitempty"`
}

func (x *SynthesizeLongAudioRequest) Reset() {
	*x = SynthesizeLongAudioRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SynthesizeLongAudioRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SynthesizeLongAudioRequest) ProtoMessage() {}

func (x *SynthesizeLongAudioRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SynthesizeLongAudioRequest.ProtoReflect.Descriptor instead.
func (*SynthesizeLongAudioRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_rawDescGZIP(), []int{0}
}

func (x *SynthesizeLongAudioRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *SynthesizeLongAudioRequest) GetInput() *SynthesisInput {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *SynthesizeLongAudioRequest) GetAudioConfig() *AudioConfig {
	if x != nil {
		return x.AudioConfig
	}
	return nil
}

func (x *SynthesizeLongAudioRequest) GetOutputGcsUri() string {
	if x != nil {
		return x.OutputGcsUri
	}
	return ""
}

func (x *SynthesizeLongAudioRequest) GetVoice() *VoiceSelectionParams {
	if x != nil {
		return x.Voice
	}
	return nil
}

// The message returned to the client by the `SynthesizeLongAudio` method.
type SynthesizeLongAudioResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SynthesizeLongAudioResponse) Reset() {
	*x = SynthesizeLongAudioResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SynthesizeLongAudioResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SynthesizeLongAudioResponse) ProtoMessage() {}

func (x *SynthesizeLongAudioResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SynthesizeLongAudioResponse.ProtoReflect.Descriptor instead.
func (*SynthesizeLongAudioResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_rawDescGZIP(), []int{1}
}

// Metadata for response returned by the `SynthesizeLongAudio` method.
type SynthesizeLongAudioMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Time when the request was received.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Time of the most recent processing update.
	LastUpdateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_update_time,json=lastUpdateTime,proto3" json:"last_update_time,omitempty"`
	// The progress of the most recent processing update in percentage, ie. 70.0%.
	ProgressPercentage float64 `protobuf:"fixed64,3,opt,name=progress_percentage,json=progressPercentage,proto3" json:"progress_percentage,omitempty"`
}

func (x *SynthesizeLongAudioMetadata) Reset() {
	*x = SynthesizeLongAudioMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SynthesizeLongAudioMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SynthesizeLongAudioMetadata) ProtoMessage() {}

func (x *SynthesizeLongAudioMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SynthesizeLongAudioMetadata.ProtoReflect.Descriptor instead.
func (*SynthesizeLongAudioMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_rawDescGZIP(), []int{2}
}

func (x *SynthesizeLongAudioMetadata) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *SynthesizeLongAudioMetadata) GetLastUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdateTime
	}
	return nil
}

func (x *SynthesizeLongAudioMetadata) GetProgressPercentage() float64 {
	if x != nil {
		return x.ProgressPercentage
	}
	return 0
}

var File_google_cloud_texttospeech_v1_cloud_tts_lrs_proto protoreflect.FileDescriptor

var file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74,
	0x65, 0x78, 0x74, 0x74, 0x6f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x74, 0x74, 0x73, 0x5f, 0x6c, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x74, 0x65, 0x78, 0x74, 0x74, 0x6f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x74, 0x6f, 0x73, 0x70, 0x65,
	0x65, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x74, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6c,
	0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x02, 0x0a,
	0x1a, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x4c, 0x6f, 0x6e, 0x67, 0x41,
	0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x47, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x74, 0x6f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x51, 0x0a, 0x0c,
	0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x74, 0x6f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x29, 0x0a, 0x0e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x67, 0x63, 0x73, 0x5f, 0x75, 0x72,
	0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0c, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x47, 0x63, 0x73, 0x55, 0x72, 0x69, 0x12, 0x4d, 0x0a, 0x05, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x74, 0x6f, 0x73,
	0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x05, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x22, 0x1d, 0x0a, 0x1b, 0x53, 0x79, 0x6e,
	0x74, 0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x4c, 0x6f, 0x6e, 0x67, 0x41, 0x75, 0x64, 0x69, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xcf, 0x01, 0x0a, 0x1b, 0x53, 0x79, 0x6e,
	0x74, 0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x4c, 0x6f, 0x6e, 0x67, 0x41, 0x75, 0x64, 0x69, 0x6f,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x12, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x32, 0xee, 0x02, 0x0a, 0x1f, 0x54,
	0x65, 0x78, 0x74, 0x54, 0x6f, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x4c, 0x6f, 0x6e, 0x67, 0x41,
	0x75, 0x64, 0x69, 0x6f, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x12, 0xf9,
	0x01, 0x0a, 0x13, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x4c, 0x6f, 0x6e,
	0x67, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x74, 0x6f, 0x73, 0x70, 0x65, 0x65,
	0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x7a, 0x65,
	0x4c, 0x6f, 0x6e, 0x67, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75,
	0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x88, 0x01, 0xca, 0x41, 0x3a, 0x0a, 0x1b, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x7a,
	0x65, 0x4c, 0x6f, 0x6e, 0x67, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1b, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x4c, 0x6f,
	0x6e, 0x67, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x45, 0x3a, 0x01, 0x2a, 0x22, 0x40, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x7a,
	0x65, 0x4c, 0x6f, 0x6e, 0x67, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x1a, 0x4f, 0xca, 0x41, 0x1b, 0x74,
	0x65, 0x78, 0x74, 0x74, 0x6f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0xf2, 0x01, 0x0a, 0x20,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x74, 0x65, 0x78, 0x74, 0x74, 0x6f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x76, 0x31,
	0x42, 0x23, 0x54, 0x65, 0x78, 0x74, 0x54, 0x6f, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x4c, 0x6f,
	0x6e, 0x67, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x78,
	0x74, 0x74, 0x6f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f,
	0x74, 0x65, 0x78, 0x74, 0x74, 0x6f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x70, 0x62, 0x3b, 0x74,
	0x65, 0x78, 0x74, 0x74, 0x6f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x70, 0x62, 0xf8, 0x01, 0x01,
	0xaa, 0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x54, 0x65, 0x78, 0x74, 0x54, 0x6f, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x54,
	0x65, 0x78, 0x74, 0x54, 0x6f, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x5c, 0x56, 0x31, 0xea, 0x02,
	0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x54, 0x65, 0x78, 0x74, 0x54, 0x6f, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_rawDescOnce sync.Once
	file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_rawDescData = file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_rawDesc
)

func file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_rawDescGZIP() []byte {
	file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_rawDescOnce.Do(func() {
		file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_rawDescData)
	})
	return file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_rawDescData
}

var file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_goTypes = []interface{}{
	(*SynthesizeLongAudioRequest)(nil),  // 0: google.cloud.texttospeech.v1.SynthesizeLongAudioRequest
	(*SynthesizeLongAudioResponse)(nil), // 1: google.cloud.texttospeech.v1.SynthesizeLongAudioResponse
	(*SynthesizeLongAudioMetadata)(nil), // 2: google.cloud.texttospeech.v1.SynthesizeLongAudioMetadata
	(*SynthesisInput)(nil),              // 3: google.cloud.texttospeech.v1.SynthesisInput
	(*AudioConfig)(nil),                 // 4: google.cloud.texttospeech.v1.AudioConfig
	(*VoiceSelectionParams)(nil),        // 5: google.cloud.texttospeech.v1.VoiceSelectionParams
	(*timestamppb.Timestamp)(nil),       // 6: google.protobuf.Timestamp
	(*longrunningpb.Operation)(nil),     // 7: google.longrunning.Operation
}
var file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_depIdxs = []int32{
	3, // 0: google.cloud.texttospeech.v1.SynthesizeLongAudioRequest.input:type_name -> google.cloud.texttospeech.v1.SynthesisInput
	4, // 1: google.cloud.texttospeech.v1.SynthesizeLongAudioRequest.audio_config:type_name -> google.cloud.texttospeech.v1.AudioConfig
	5, // 2: google.cloud.texttospeech.v1.SynthesizeLongAudioRequest.voice:type_name -> google.cloud.texttospeech.v1.VoiceSelectionParams
	6, // 3: google.cloud.texttospeech.v1.SynthesizeLongAudioMetadata.start_time:type_name -> google.protobuf.Timestamp
	6, // 4: google.cloud.texttospeech.v1.SynthesizeLongAudioMetadata.last_update_time:type_name -> google.protobuf.Timestamp
	0, // 5: google.cloud.texttospeech.v1.TextToSpeechLongAudioSynthesize.SynthesizeLongAudio:input_type -> google.cloud.texttospeech.v1.SynthesizeLongAudioRequest
	7, // 6: google.cloud.texttospeech.v1.TextToSpeechLongAudioSynthesize.SynthesizeLongAudio:output_type -> google.longrunning.Operation
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_init() }
func file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_init() {
	if File_google_cloud_texttospeech_v1_cloud_tts_lrs_proto != nil {
		return
	}
	file_google_cloud_texttospeech_v1_cloud_tts_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SynthesizeLongAudioRequest); i {
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
		file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SynthesizeLongAudioResponse); i {
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
		file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SynthesizeLongAudioMetadata); i {
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
			RawDescriptor: file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_goTypes,
		DependencyIndexes: file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_depIdxs,
		MessageInfos:      file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_msgTypes,
	}.Build()
	File_google_cloud_texttospeech_v1_cloud_tts_lrs_proto = out.File
	file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_rawDesc = nil
	file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_goTypes = nil
	file_google_cloud_texttospeech_v1_cloud_tts_lrs_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TextToSpeechLongAudioSynthesizeClient is the client API for TextToSpeechLongAudioSynthesize service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TextToSpeechLongAudioSynthesizeClient interface {
	// Synthesizes long form text asynchronously.
	SynthesizeLongAudio(ctx context.Context, in *SynthesizeLongAudioRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
}

type textToSpeechLongAudioSynthesizeClient struct {
	cc grpc.ClientConnInterface
}

func NewTextToSpeechLongAudioSynthesizeClient(cc grpc.ClientConnInterface) TextToSpeechLongAudioSynthesizeClient {
	return &textToSpeechLongAudioSynthesizeClient{cc}
}

func (c *textToSpeechLongAudioSynthesizeClient) SynthesizeLongAudio(ctx context.Context, in *SynthesizeLongAudioRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.texttospeech.v1.TextToSpeechLongAudioSynthesize/SynthesizeLongAudio", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TextToSpeechLongAudioSynthesizeServer is the server API for TextToSpeechLongAudioSynthesize service.
type TextToSpeechLongAudioSynthesizeServer interface {
	// Synthesizes long form text asynchronously.
	SynthesizeLongAudio(context.Context, *SynthesizeLongAudioRequest) (*longrunningpb.Operation, error)
}

// UnimplementedTextToSpeechLongAudioSynthesizeServer can be embedded to have forward compatible implementations.
type UnimplementedTextToSpeechLongAudioSynthesizeServer struct {
}

func (*UnimplementedTextToSpeechLongAudioSynthesizeServer) SynthesizeLongAudio(context.Context, *SynthesizeLongAudioRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SynthesizeLongAudio not implemented")
}

func RegisterTextToSpeechLongAudioSynthesizeServer(s *grpc.Server, srv TextToSpeechLongAudioSynthesizeServer) {
	s.RegisterService(&_TextToSpeechLongAudioSynthesize_serviceDesc, srv)
}

func _TextToSpeechLongAudioSynthesize_SynthesizeLongAudio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SynthesizeLongAudioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextToSpeechLongAudioSynthesizeServer).SynthesizeLongAudio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.texttospeech.v1.TextToSpeechLongAudioSynthesize/SynthesizeLongAudio",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextToSpeechLongAudioSynthesizeServer).SynthesizeLongAudio(ctx, req.(*SynthesizeLongAudioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TextToSpeechLongAudioSynthesize_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.texttospeech.v1.TextToSpeechLongAudioSynthesize",
	HandlerType: (*TextToSpeechLongAudioSynthesizeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SynthesizeLongAudio",
			Handler:    _TextToSpeechLongAudioSynthesize_SynthesizeLongAudio_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/texttospeech/v1/cloud_tts_lrs.proto",
}
