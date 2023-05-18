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
// 	protoc        v3.21.12
// source: google/appengine/v1/domain.proto

package appenginepb

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

// A domain that a user has been authorized to administer. To authorize use
// of a domain, verify ownership via
// [Search Console](https://search.google.com/search-console/welcome).
type AuthorizedDomain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Full path to the `AuthorizedDomain` resource in the API. Example:
	// `apps/myapp/authorizedDomains/example.com`.
	//
	// @OutputOnly
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Fully qualified domain name of the domain authorized for use. Example:
	// `example.com`.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AuthorizedDomain) Reset() {
	*x = AuthorizedDomain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_v1_domain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizedDomain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizedDomain) ProtoMessage() {}

func (x *AuthorizedDomain) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_v1_domain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizedDomain.ProtoReflect.Descriptor instead.
func (*AuthorizedDomain) Descriptor() ([]byte, []int) {
	return file_google_appengine_v1_domain_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorizedDomain) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AuthorizedDomain) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_google_appengine_v1_domain_proto protoreflect.FileDescriptor

var file_google_appengine_v1_domain_proto_rawDesc = []byte{
	0x0a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x36, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x64, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42,
	0xbc, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f,
	0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f,
	0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0x3b, 0x61, 0x70, 0x70, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0xaa, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0xea,
	0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_appengine_v1_domain_proto_rawDescOnce sync.Once
	file_google_appengine_v1_domain_proto_rawDescData = file_google_appengine_v1_domain_proto_rawDesc
)

func file_google_appengine_v1_domain_proto_rawDescGZIP() []byte {
	file_google_appengine_v1_domain_proto_rawDescOnce.Do(func() {
		file_google_appengine_v1_domain_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_appengine_v1_domain_proto_rawDescData)
	})
	return file_google_appengine_v1_domain_proto_rawDescData
}

var file_google_appengine_v1_domain_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_appengine_v1_domain_proto_goTypes = []interface{}{
	(*AuthorizedDomain)(nil), // 0: google.appengine.v1.AuthorizedDomain
}
var file_google_appengine_v1_domain_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_appengine_v1_domain_proto_init() }
func file_google_appengine_v1_domain_proto_init() {
	if File_google_appengine_v1_domain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_appengine_v1_domain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizedDomain); i {
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
			RawDescriptor: file_google_appengine_v1_domain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_appengine_v1_domain_proto_goTypes,
		DependencyIndexes: file_google_appengine_v1_domain_proto_depIdxs,
		MessageInfos:      file_google_appengine_v1_domain_proto_msgTypes,
	}.Build()
	File_google_appengine_v1_domain_proto = out.File
	file_google_appengine_v1_domain_proto_rawDesc = nil
	file_google_appengine_v1_domain_proto_goTypes = nil
	file_google_appengine_v1_domain_proto_depIdxs = nil
}
