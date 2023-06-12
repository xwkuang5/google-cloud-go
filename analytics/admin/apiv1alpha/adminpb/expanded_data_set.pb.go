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
// source: google/analytics/admin/v1alpha/expanded_data_set.proto

package adminpb

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

// The match type for the string filter.
type ExpandedDataSetFilter_StringFilter_MatchType int32

const (
	// Unspecified
	ExpandedDataSetFilter_StringFilter_MATCH_TYPE_UNSPECIFIED ExpandedDataSetFilter_StringFilter_MatchType = 0
	// Exact match of the string value.
	ExpandedDataSetFilter_StringFilter_EXACT ExpandedDataSetFilter_StringFilter_MatchType = 1
	// Contains the string value.
	ExpandedDataSetFilter_StringFilter_CONTAINS ExpandedDataSetFilter_StringFilter_MatchType = 2
)

// Enum value maps for ExpandedDataSetFilter_StringFilter_MatchType.
var (
	ExpandedDataSetFilter_StringFilter_MatchType_name = map[int32]string{
		0: "MATCH_TYPE_UNSPECIFIED",
		1: "EXACT",
		2: "CONTAINS",
	}
	ExpandedDataSetFilter_StringFilter_MatchType_value = map[string]int32{
		"MATCH_TYPE_UNSPECIFIED": 0,
		"EXACT":                  1,
		"CONTAINS":               2,
	}
)

func (x ExpandedDataSetFilter_StringFilter_MatchType) Enum() *ExpandedDataSetFilter_StringFilter_MatchType {
	p := new(ExpandedDataSetFilter_StringFilter_MatchType)
	*p = x
	return p
}

func (x ExpandedDataSetFilter_StringFilter_MatchType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExpandedDataSetFilter_StringFilter_MatchType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_analytics_admin_v1alpha_expanded_data_set_proto_enumTypes[0].Descriptor()
}

func (ExpandedDataSetFilter_StringFilter_MatchType) Type() protoreflect.EnumType {
	return &file_google_analytics_admin_v1alpha_expanded_data_set_proto_enumTypes[0]
}

func (x ExpandedDataSetFilter_StringFilter_MatchType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExpandedDataSetFilter_StringFilter_MatchType.Descriptor instead.
func (ExpandedDataSetFilter_StringFilter_MatchType) EnumDescriptor() ([]byte, []int) {
	return file_google_analytics_admin_v1alpha_expanded_data_set_proto_rawDescGZIP(), []int{0, 0, 0}
}

// A specific filter for a single dimension
type ExpandedDataSetFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// One of the above filters.
	//
	// Types that are assignable to OneFilter:
	//	*ExpandedDataSetFilter_StringFilter_
	//	*ExpandedDataSetFilter_InListFilter_
	OneFilter isExpandedDataSetFilter_OneFilter `protobuf_oneof:"one_filter"`
	// Required. The dimension name to filter.
	FieldName string `protobuf:"bytes,1,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
}

func (x *ExpandedDataSetFilter) Reset() {
	*x = ExpandedDataSetFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_analytics_admin_v1alpha_expanded_data_set_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpandedDataSetFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandedDataSetFilter) ProtoMessage() {}

func (x *ExpandedDataSetFilter) ProtoReflect() protoreflect.Message {
	mi := &file_google_analytics_admin_v1alpha_expanded_data_set_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandedDataSetFilter.ProtoReflect.Descriptor instead.
func (*ExpandedDataSetFilter) Descriptor() ([]byte, []int) {
	return file_google_analytics_admin_v1alpha_expanded_data_set_proto_rawDescGZIP(), []int{0}
}

func (m *ExpandedDataSetFilter) GetOneFilter() isExpandedDataSetFilter_OneFilter {
	if m != nil {
		return m.OneFilter
	}
	return nil
}

func (x *ExpandedDataSetFilter) GetStringFilter() *ExpandedDataSetFilter_StringFilter {
	if x, ok := x.GetOneFilter().(*ExpandedDataSetFilter_StringFilter_); ok {
		return x.StringFilter
	}
	return nil
}

func (x *ExpandedDataSetFilter) GetInListFilter() *ExpandedDataSetFilter_InListFilter {
	if x, ok := x.GetOneFilter().(*ExpandedDataSetFilter_InListFilter_); ok {
		return x.InListFilter
	}
	return nil
}

func (x *ExpandedDataSetFilter) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

type isExpandedDataSetFilter_OneFilter interface {
	isExpandedDataSetFilter_OneFilter()
}

type ExpandedDataSetFilter_StringFilter_ struct {
	// A filter for a string-type dimension that matches a particular pattern.
	StringFilter *ExpandedDataSetFilter_StringFilter `protobuf:"bytes,2,opt,name=string_filter,json=stringFilter,proto3,oneof"`
}

type ExpandedDataSetFilter_InListFilter_ struct {
	// A filter for a string dimension that matches a particular list of
	// options.
	InListFilter *ExpandedDataSetFilter_InListFilter `protobuf:"bytes,3,opt,name=in_list_filter,json=inListFilter,proto3,oneof"`
}

func (*ExpandedDataSetFilter_StringFilter_) isExpandedDataSetFilter_OneFilter() {}

func (*ExpandedDataSetFilter_InListFilter_) isExpandedDataSetFilter_OneFilter() {}

// A logical expression of EnhancedDataSet dimension filters.
type ExpandedDataSetFilterExpression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The expression applied to a filter.
	//
	// Types that are assignable to Expr:
	//	*ExpandedDataSetFilterExpression_AndGroup
	//	*ExpandedDataSetFilterExpression_NotExpression
	//	*ExpandedDataSetFilterExpression_Filter
	Expr isExpandedDataSetFilterExpression_Expr `protobuf_oneof:"expr"`
}

func (x *ExpandedDataSetFilterExpression) Reset() {
	*x = ExpandedDataSetFilterExpression{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_analytics_admin_v1alpha_expanded_data_set_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpandedDataSetFilterExpression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandedDataSetFilterExpression) ProtoMessage() {}

func (x *ExpandedDataSetFilterExpression) ProtoReflect() protoreflect.Message {
	mi := &file_google_analytics_admin_v1alpha_expanded_data_set_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandedDataSetFilterExpression.ProtoReflect.Descriptor instead.
func (*ExpandedDataSetFilterExpression) Descriptor() ([]byte, []int) {
	return file_google_analytics_admin_v1alpha_expanded_data_set_proto_rawDescGZIP(), []int{1}
}

func (m *ExpandedDataSetFilterExpression) GetExpr() isExpandedDataSetFilterExpression_Expr {
	if m != nil {
		return m.Expr
	}
	return nil
}

func (x *ExpandedDataSetFilterExpression) GetAndGroup() *ExpandedDataSetFilterExpressionList {
	if x, ok := x.GetExpr().(*ExpandedDataSetFilterExpression_AndGroup); ok {
		return x.AndGroup
	}
	return nil
}

func (x *ExpandedDataSetFilterExpression) GetNotExpression() *ExpandedDataSetFilterExpression {
	if x, ok := x.GetExpr().(*ExpandedDataSetFilterExpression_NotExpression); ok {
		return x.NotExpression
	}
	return nil
}

func (x *ExpandedDataSetFilterExpression) GetFilter() *ExpandedDataSetFilter {
	if x, ok := x.GetExpr().(*ExpandedDataSetFilterExpression_Filter); ok {
		return x.Filter
	}
	return nil
}

type isExpandedDataSetFilterExpression_Expr interface {
	isExpandedDataSetFilterExpression_Expr()
}

type ExpandedDataSetFilterExpression_AndGroup struct {
	// A list of expressions to be AND’ed together. It must contain a
	// ExpandedDataSetFilterExpression with either not_expression or
	// dimension_filter. This must be set for the top level
	// ExpandedDataSetFilterExpression.
	AndGroup *ExpandedDataSetFilterExpressionList `protobuf:"bytes,1,opt,name=and_group,json=andGroup,proto3,oneof"`
}

type ExpandedDataSetFilterExpression_NotExpression struct {
	// A filter expression to be NOT'ed (that is, inverted, complemented). It
	// must include a dimension_filter. This cannot be set on the
	// top level ExpandedDataSetFilterExpression.
	NotExpression *ExpandedDataSetFilterExpression `protobuf:"bytes,2,opt,name=not_expression,json=notExpression,proto3,oneof"`
}

type ExpandedDataSetFilterExpression_Filter struct {
	// A filter on a single dimension. This cannot be set on the top
	// level ExpandedDataSetFilterExpression.
	Filter *ExpandedDataSetFilter `protobuf:"bytes,3,opt,name=filter,proto3,oneof"`
}

func (*ExpandedDataSetFilterExpression_AndGroup) isExpandedDataSetFilterExpression_Expr() {}

func (*ExpandedDataSetFilterExpression_NotExpression) isExpandedDataSetFilterExpression_Expr() {}

func (*ExpandedDataSetFilterExpression_Filter) isExpandedDataSetFilterExpression_Expr() {}

// A list of ExpandedDataSet filter expressions.
type ExpandedDataSetFilterExpressionList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of ExpandedDataSet filter expressions.
	FilterExpressions []*ExpandedDataSetFilterExpression `protobuf:"bytes,1,rep,name=filter_expressions,json=filterExpressions,proto3" json:"filter_expressions,omitempty"`
}

func (x *ExpandedDataSetFilterExpressionList) Reset() {
	*x = ExpandedDataSetFilterExpressionList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_analytics_admin_v1alpha_expanded_data_set_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpandedDataSetFilterExpressionList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandedDataSetFilterExpressionList) ProtoMessage() {}

func (x *ExpandedDataSetFilterExpressionList) ProtoReflect() protoreflect.Message {
	mi := &file_google_analytics_admin_v1alpha_expanded_data_set_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandedDataSetFilterExpressionList.ProtoReflect.Descriptor instead.
func (*ExpandedDataSetFilterExpressionList) Descriptor() ([]byte, []int) {
	return file_google_analytics_admin_v1alpha_expanded_data_set_proto_rawDescGZIP(), []int{2}
}

func (x *ExpandedDataSetFilterExpressionList) GetFilterExpressions() []*ExpandedDataSetFilterExpression {
	if x != nil {
		return x.FilterExpressions
	}
	return nil
}

// A resource message representing a GA4 ExpandedDataSet.
type ExpandedDataSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name for this ExpandedDataSet resource.
	// Format: properties/{property_id}/expandedDataSets/{expanded_data_set}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The display name of the ExpandedDataSet.
	// Max 200 chars.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Optional. The description of the ExpandedDataSet.
	// Max 50 chars.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Immutable. The list of dimensions included in the ExpandedDataSet.
	// See the [API
	// Dimensions](https://developers.google.com/analytics/devguides/reporting/data/v1/api-schema#dimensions)
	// for the list of dimension names.
	DimensionNames []string `protobuf:"bytes,4,rep,name=dimension_names,json=dimensionNames,proto3" json:"dimension_names,omitempty"`
	// Immutable. The list of metrics included in the ExpandedDataSet.
	// See the [API
	// Metrics](https://developers.google.com/analytics/devguides/reporting/data/v1/api-schema#metrics)
	// for the list of dimension names.
	MetricNames []string `protobuf:"bytes,5,rep,name=metric_names,json=metricNames,proto3" json:"metric_names,omitempty"`
	// Immutable. A logical expression of ExpandedDataSet filters applied to
	// dimension included in the ExpandedDataSet. This filter is used to reduce
	// the number of rows and thus the chance of encountering `other` row.
	DimensionFilterExpression *ExpandedDataSetFilterExpression `protobuf:"bytes,6,opt,name=dimension_filter_expression,json=dimensionFilterExpression,proto3" json:"dimension_filter_expression,omitempty"`
	// Output only. Time when expanded data set began (or will begin) collecing
	// data.
	DataCollectionStartTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=data_collection_start_time,json=dataCollectionStartTime,proto3" json:"data_collection_start_time,omitempty"`
}

func (x *ExpandedDataSet) Reset() {
	*x = ExpandedDataSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_analytics_admin_v1alpha_expanded_data_set_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpandedDataSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandedDataSet) ProtoMessage() {}

func (x *ExpandedDataSet) ProtoReflect() protoreflect.Message {
	mi := &file_google_analytics_admin_v1alpha_expanded_data_set_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandedDataSet.ProtoReflect.Descriptor instead.
func (*ExpandedDataSet) Descriptor() ([]byte, []int) {
	return file_google_analytics_admin_v1alpha_expanded_data_set_proto_rawDescGZIP(), []int{3}
}

func (x *ExpandedDataSet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExpandedDataSet) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *ExpandedDataSet) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ExpandedDataSet) GetDimensionNames() []string {
	if x != nil {
		return x.DimensionNames
	}
	return nil
}

func (x *ExpandedDataSet) GetMetricNames() []string {
	if x != nil {
		return x.MetricNames
	}
	return nil
}

func (x *ExpandedDataSet) GetDimensionFilterExpression() *ExpandedDataSetFilterExpression {
	if x != nil {
		return x.DimensionFilterExpression
	}
	return nil
}

func (x *ExpandedDataSet) GetDataCollectionStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DataCollectionStartTime
	}
	return nil
}

// A filter for a string-type dimension that matches a particular pattern.
type ExpandedDataSetFilter_StringFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The match type for the string filter.
	MatchType ExpandedDataSetFilter_StringFilter_MatchType `protobuf:"varint,1,opt,name=match_type,json=matchType,proto3,enum=google.analytics.admin.v1alpha.ExpandedDataSetFilter_StringFilter_MatchType" json:"match_type,omitempty"`
	// Required. The string value to be matched against.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// Optional. If true, the match is case-sensitive. If false, the match is
	// case-insensitive.
	// Must be true when match_type is EXACT.
	// Must be false when match_type is CONTAINS.
	CaseSensitive bool `protobuf:"varint,3,opt,name=case_sensitive,json=caseSensitive,proto3" json:"case_sensitive,omitempty"`
}

func (x *ExpandedDataSetFilter_StringFilter) Reset() {
	*x = ExpandedDataSetFilter_StringFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_analytics_admin_v1alpha_expanded_data_set_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpandedDataSetFilter_StringFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandedDataSetFilter_StringFilter) ProtoMessage() {}

func (x *ExpandedDataSetFilter_StringFilter) ProtoReflect() protoreflect.Message {
	mi := &file_google_analytics_admin_v1alpha_expanded_data_set_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandedDataSetFilter_StringFilter.ProtoReflect.Descriptor instead.
func (*ExpandedDataSetFilter_StringFilter) Descriptor() ([]byte, []int) {
	return file_google_analytics_admin_v1alpha_expanded_data_set_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ExpandedDataSetFilter_StringFilter) GetMatchType() ExpandedDataSetFilter_StringFilter_MatchType {
	if x != nil {
		return x.MatchType
	}
	return ExpandedDataSetFilter_StringFilter_MATCH_TYPE_UNSPECIFIED
}

func (x *ExpandedDataSetFilter_StringFilter) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ExpandedDataSetFilter_StringFilter) GetCaseSensitive() bool {
	if x != nil {
		return x.CaseSensitive
	}
	return false
}

// A filter for a string dimension that matches a particular list of options.
type ExpandedDataSetFilter_InListFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The list of possible string values to match against. Must be
	// non-empty.
	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	// Optional. If true, the match is case-sensitive. If false, the match is
	// case-insensitive.
	// Must be true.
	CaseSensitive bool `protobuf:"varint,2,opt,name=case_sensitive,json=caseSensitive,proto3" json:"case_sensitive,omitempty"`
}

func (x *ExpandedDataSetFilter_InListFilter) Reset() {
	*x = ExpandedDataSetFilter_InListFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_analytics_admin_v1alpha_expanded_data_set_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpandedDataSetFilter_InListFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandedDataSetFilter_InListFilter) ProtoMessage() {}

func (x *ExpandedDataSetFilter_InListFilter) ProtoReflect() protoreflect.Message {
	mi := &file_google_analytics_admin_v1alpha_expanded_data_set_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandedDataSetFilter_InListFilter.ProtoReflect.Descriptor instead.
func (*ExpandedDataSetFilter_InListFilter) Descriptor() ([]byte, []int) {
	return file_google_analytics_admin_v1alpha_expanded_data_set_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ExpandedDataSetFilter_InListFilter) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *ExpandedDataSetFilter_InListFilter) GetCaseSensitive() bool {
	if x != nil {
		return x.CaseSensitive
	}
	return false
}

var File_google_analytics_admin_v1alpha_expanded_data_set_proto protoreflect.FileDescriptor

var file_google_analytics_admin_v1alpha_expanded_data_set_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73,
	0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x05, 0x0a, 0x15, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64,
	0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x69, 0x0a, 0x0d, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x6a, 0x0a, 0x0e, 0x69, 0x6e,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0c, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x89, 0x02, 0x0a, 0x0c, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x70, 0x0a, 0x0a, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x4c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2a, 0x0a, 0x0e, 0x63, 0x61, 0x73, 0x65,
	0x5f, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0d, 0x63, 0x61, 0x73, 0x65, 0x53, 0x65, 0x6e, 0x73, 0x69,
	0x74, 0x69, 0x76, 0x65, 0x22, 0x40, 0x0a, 0x09, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x45, 0x58, 0x41, 0x43, 0x54, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4e, 0x54,
	0x41, 0x49, 0x4e, 0x53, 0x10, 0x02, 0x1a, 0x57, 0x0a, 0x0c, 0x49, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x0e, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x6e, 0x73,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x01,
	0x52, 0x0d, 0x63, 0x61, 0x73, 0x65, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x42,
	0x0c, 0x0a, 0x0a, 0x6f, 0x6e, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0xc8, 0x02,
	0x0a, 0x1f, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65,
	0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x62, 0x0a, 0x09, 0x61, 0x6e, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x08, 0x61, 0x6e, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x68, 0x0a, 0x0e, 0x6e, 0x6f, 0x74, 0x5f, 0x65, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x45,
	0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x0d, 0x6e, 0x6f, 0x74, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x4f, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x42, 0x06, 0x0a, 0x04, 0x65, 0x78, 0x70, 0x72, 0x22, 0x95, 0x01, 0x0a, 0x23, 0x45, 0x78, 0x70,
	0x61, 0x6e, 0x64, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x6e, 0x0a, 0x12, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x45, 0x78,
	0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0xa4, 0x04, 0x0a, 0x0f, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x65, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x0f,
	0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x0e, 0x64, 0x69, 0x6d, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0c, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x12, 0x84, 0x01, 0x0a, 0x1b, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64,
	0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x45,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x19,
	0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x45,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x5c, 0x0a, 0x1a, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x17,
	0x64, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x3a, 0x6e, 0xea, 0x41, 0x6b, 0x0a, 0x2d, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x78, 0x70, 0x61,
	0x6e, 0x64, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x12, 0x3a, 0x70, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x7d, 0x2f, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x65, 0x74, 0x73, 0x2f, 0x7b, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x73, 0x65, 0x74, 0x7d, 0x42, 0x7c, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x14, 0x45,
	0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0x3b, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_analytics_admin_v1alpha_expanded_data_set_proto_rawDescOnce sync.Once
	file_google_analytics_admin_v1alpha_expanded_data_set_proto_rawDescData = file_google_analytics_admin_v1alpha_expanded_data_set_proto_rawDesc
)

func file_google_analytics_admin_v1alpha_expanded_data_set_proto_rawDescGZIP() []byte {
	file_google_analytics_admin_v1alpha_expanded_data_set_proto_rawDescOnce.Do(func() {
		file_google_analytics_admin_v1alpha_expanded_data_set_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_analytics_admin_v1alpha_expanded_data_set_proto_rawDescData)
	})
	return file_google_analytics_admin_v1alpha_expanded_data_set_proto_rawDescData
}

var file_google_analytics_admin_v1alpha_expanded_data_set_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_analytics_admin_v1alpha_expanded_data_set_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_analytics_admin_v1alpha_expanded_data_set_proto_goTypes = []interface{}{
	(ExpandedDataSetFilter_StringFilter_MatchType)(0), // 0: google.analytics.admin.v1alpha.ExpandedDataSetFilter.StringFilter.MatchType
	(*ExpandedDataSetFilter)(nil),                     // 1: google.analytics.admin.v1alpha.ExpandedDataSetFilter
	(*ExpandedDataSetFilterExpression)(nil),           // 2: google.analytics.admin.v1alpha.ExpandedDataSetFilterExpression
	(*ExpandedDataSetFilterExpressionList)(nil),       // 3: google.analytics.admin.v1alpha.ExpandedDataSetFilterExpressionList
	(*ExpandedDataSet)(nil),                           // 4: google.analytics.admin.v1alpha.ExpandedDataSet
	(*ExpandedDataSetFilter_StringFilter)(nil),        // 5: google.analytics.admin.v1alpha.ExpandedDataSetFilter.StringFilter
	(*ExpandedDataSetFilter_InListFilter)(nil),        // 6: google.analytics.admin.v1alpha.ExpandedDataSetFilter.InListFilter
	(*timestamppb.Timestamp)(nil),                     // 7: google.protobuf.Timestamp
}
var file_google_analytics_admin_v1alpha_expanded_data_set_proto_depIdxs = []int32{
	5, // 0: google.analytics.admin.v1alpha.ExpandedDataSetFilter.string_filter:type_name -> google.analytics.admin.v1alpha.ExpandedDataSetFilter.StringFilter
	6, // 1: google.analytics.admin.v1alpha.ExpandedDataSetFilter.in_list_filter:type_name -> google.analytics.admin.v1alpha.ExpandedDataSetFilter.InListFilter
	3, // 2: google.analytics.admin.v1alpha.ExpandedDataSetFilterExpression.and_group:type_name -> google.analytics.admin.v1alpha.ExpandedDataSetFilterExpressionList
	2, // 3: google.analytics.admin.v1alpha.ExpandedDataSetFilterExpression.not_expression:type_name -> google.analytics.admin.v1alpha.ExpandedDataSetFilterExpression
	1, // 4: google.analytics.admin.v1alpha.ExpandedDataSetFilterExpression.filter:type_name -> google.analytics.admin.v1alpha.ExpandedDataSetFilter
	2, // 5: google.analytics.admin.v1alpha.ExpandedDataSetFilterExpressionList.filter_expressions:type_name -> google.analytics.admin.v1alpha.ExpandedDataSetFilterExpression
	2, // 6: google.analytics.admin.v1alpha.ExpandedDataSet.dimension_filter_expression:type_name -> google.analytics.admin.v1alpha.ExpandedDataSetFilterExpression
	7, // 7: google.analytics.admin.v1alpha.ExpandedDataSet.data_collection_start_time:type_name -> google.protobuf.Timestamp
	0, // 8: google.analytics.admin.v1alpha.ExpandedDataSetFilter.StringFilter.match_type:type_name -> google.analytics.admin.v1alpha.ExpandedDataSetFilter.StringFilter.MatchType
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_google_analytics_admin_v1alpha_expanded_data_set_proto_init() }
func file_google_analytics_admin_v1alpha_expanded_data_set_proto_init() {
	if File_google_analytics_admin_v1alpha_expanded_data_set_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_analytics_admin_v1alpha_expanded_data_set_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpandedDataSetFilter); i {
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
		file_google_analytics_admin_v1alpha_expanded_data_set_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpandedDataSetFilterExpression); i {
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
		file_google_analytics_admin_v1alpha_expanded_data_set_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpandedDataSetFilterExpressionList); i {
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
		file_google_analytics_admin_v1alpha_expanded_data_set_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpandedDataSet); i {
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
		file_google_analytics_admin_v1alpha_expanded_data_set_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpandedDataSetFilter_StringFilter); i {
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
		file_google_analytics_admin_v1alpha_expanded_data_set_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpandedDataSetFilter_InListFilter); i {
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
	file_google_analytics_admin_v1alpha_expanded_data_set_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ExpandedDataSetFilter_StringFilter_)(nil),
		(*ExpandedDataSetFilter_InListFilter_)(nil),
	}
	file_google_analytics_admin_v1alpha_expanded_data_set_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ExpandedDataSetFilterExpression_AndGroup)(nil),
		(*ExpandedDataSetFilterExpression_NotExpression)(nil),
		(*ExpandedDataSetFilterExpression_Filter)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_analytics_admin_v1alpha_expanded_data_set_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_analytics_admin_v1alpha_expanded_data_set_proto_goTypes,
		DependencyIndexes: file_google_analytics_admin_v1alpha_expanded_data_set_proto_depIdxs,
		EnumInfos:         file_google_analytics_admin_v1alpha_expanded_data_set_proto_enumTypes,
		MessageInfos:      file_google_analytics_admin_v1alpha_expanded_data_set_proto_msgTypes,
	}.Build()
	File_google_analytics_admin_v1alpha_expanded_data_set_proto = out.File
	file_google_analytics_admin_v1alpha_expanded_data_set_proto_rawDesc = nil
	file_google_analytics_admin_v1alpha_expanded_data_set_proto_goTypes = nil
	file_google_analytics_admin_v1alpha_expanded_data_set_proto_depIdxs = nil
}
