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
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: google/cloud/aiplatform/v1/feature_view.proto

package aiplatformpb

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

// FeatureView is representation of values that the FeatureOnlineStore will
// serve based on its syncConfig.
type FeatureView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Source:
	//
	//	*FeatureView_BigQuerySource_
	//	*FeatureView_FeatureRegistrySource_
	Source isFeatureView_Source `protobuf_oneof:"source"`
	// Identifier. Name of the FeatureView. Format:
	// `projects/{project}/locations/{location}/featureOnlineStores/{feature_online_store}/featureViews/{feature_view}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Timestamp when this FeatureView was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Timestamp when this FeatureView was last updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Optional. Used to perform consistent read-modify-write updates. If not set,
	// a blind "overwrite" update happens.
	Etag string `protobuf:"bytes,4,opt,name=etag,proto3" json:"etag,omitempty"`
	// Optional. The labels with user-defined metadata to organize your
	// FeatureViews.
	//
	// Label keys and values can be no longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	//
	// See https://goo.gl/xmQnxf for more information on and examples of labels.
	// No more than 64 user labels can be associated with one
	// FeatureOnlineStore(System labels are excluded)." System reserved label keys
	// are prefixed with "aiplatform.googleapis.com/" and are immutable.
	Labels map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Configures when data is to be synced/updated for this FeatureView. At the
	// end of the sync the latest featureValues for each entityId of this
	// FeatureView are made ready for online serving.
	SyncConfig *FeatureView_SyncConfig `protobuf:"bytes,7,opt,name=sync_config,json=syncConfig,proto3" json:"sync_config,omitempty"`
}

func (x *FeatureView) Reset() {
	*x = FeatureView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_feature_view_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureView) ProtoMessage() {}

func (x *FeatureView) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_feature_view_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureView.ProtoReflect.Descriptor instead.
func (*FeatureView) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_feature_view_proto_rawDescGZIP(), []int{0}
}

func (m *FeatureView) GetSource() isFeatureView_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *FeatureView) GetBigQuerySource() *FeatureView_BigQuerySource {
	if x, ok := x.GetSource().(*FeatureView_BigQuerySource_); ok {
		return x.BigQuerySource
	}
	return nil
}

func (x *FeatureView) GetFeatureRegistrySource() *FeatureView_FeatureRegistrySource {
	if x, ok := x.GetSource().(*FeatureView_FeatureRegistrySource_); ok {
		return x.FeatureRegistrySource
	}
	return nil
}

func (x *FeatureView) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FeatureView) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *FeatureView) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *FeatureView) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *FeatureView) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *FeatureView) GetSyncConfig() *FeatureView_SyncConfig {
	if x != nil {
		return x.SyncConfig
	}
	return nil
}

type isFeatureView_Source interface {
	isFeatureView_Source()
}

type FeatureView_BigQuerySource_ struct {
	// Optional. Configures how data is supposed to be extracted from a BigQuery
	// source to be loaded onto the FeatureOnlineStore.
	BigQuerySource *FeatureView_BigQuerySource `protobuf:"bytes,6,opt,name=big_query_source,json=bigQuerySource,proto3,oneof"`
}

type FeatureView_FeatureRegistrySource_ struct {
	// Optional. Configures the features from a Feature Registry source that
	// need to be loaded onto the FeatureOnlineStore.
	FeatureRegistrySource *FeatureView_FeatureRegistrySource `protobuf:"bytes,9,opt,name=feature_registry_source,json=featureRegistrySource,proto3,oneof"`
}

func (*FeatureView_BigQuerySource_) isFeatureView_Source() {}

func (*FeatureView_FeatureRegistrySource_) isFeatureView_Source() {}

type FeatureView_BigQuerySource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The BigQuery view URI that will be materialized on each sync
	// trigger based on FeatureView.SyncConfig.
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// Required. Columns to construct entity_id / row keys.
	EntityIdColumns []string `protobuf:"bytes,2,rep,name=entity_id_columns,json=entityIdColumns,proto3" json:"entity_id_columns,omitempty"`
}

func (x *FeatureView_BigQuerySource) Reset() {
	*x = FeatureView_BigQuerySource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_feature_view_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureView_BigQuerySource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureView_BigQuerySource) ProtoMessage() {}

func (x *FeatureView_BigQuerySource) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_feature_view_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureView_BigQuerySource.ProtoReflect.Descriptor instead.
func (*FeatureView_BigQuerySource) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_feature_view_proto_rawDescGZIP(), []int{0, 0}
}

func (x *FeatureView_BigQuerySource) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *FeatureView_BigQuerySource) GetEntityIdColumns() []string {
	if x != nil {
		return x.EntityIdColumns
	}
	return nil
}

// Configuration for Sync. Only one option is set.
type FeatureView_SyncConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cron schedule (https://en.wikipedia.org/wiki/Cron) to launch scheduled
	// runs. To explicitly set a timezone to the cron tab, apply a prefix in
	// the cron tab: "CRON_TZ=${IANA_TIME_ZONE}" or "TZ=${IANA_TIME_ZONE}".
	// The ${IANA_TIME_ZONE} may only be a valid string from IANA time zone
	// database. For example, "CRON_TZ=America/New_York 1 * * * *", or
	// "TZ=America/New_York 1 * * * *".
	Cron string `protobuf:"bytes,1,opt,name=cron,proto3" json:"cron,omitempty"`
}

func (x *FeatureView_SyncConfig) Reset() {
	*x = FeatureView_SyncConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_feature_view_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureView_SyncConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureView_SyncConfig) ProtoMessage() {}

func (x *FeatureView_SyncConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_feature_view_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureView_SyncConfig.ProtoReflect.Descriptor instead.
func (*FeatureView_SyncConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_feature_view_proto_rawDescGZIP(), []int{0, 1}
}

func (x *FeatureView_SyncConfig) GetCron() string {
	if x != nil {
		return x.Cron
	}
	return ""
}

// A Feature Registry source for features that need to be synced to Online
// Store.
type FeatureView_FeatureRegistrySource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. List of features that need to be synced to Online Store.
	FeatureGroups []*FeatureView_FeatureRegistrySource_FeatureGroup `protobuf:"bytes,1,rep,name=feature_groups,json=featureGroups,proto3" json:"feature_groups,omitempty"`
	// Optional. The project number of the parent project of the Feature Groups.
	ProjectNumber *int64 `protobuf:"varint,2,opt,name=project_number,json=projectNumber,proto3,oneof" json:"project_number,omitempty"`
}

func (x *FeatureView_FeatureRegistrySource) Reset() {
	*x = FeatureView_FeatureRegistrySource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_feature_view_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureView_FeatureRegistrySource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureView_FeatureRegistrySource) ProtoMessage() {}

func (x *FeatureView_FeatureRegistrySource) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_feature_view_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureView_FeatureRegistrySource.ProtoReflect.Descriptor instead.
func (*FeatureView_FeatureRegistrySource) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_feature_view_proto_rawDescGZIP(), []int{0, 2}
}

func (x *FeatureView_FeatureRegistrySource) GetFeatureGroups() []*FeatureView_FeatureRegistrySource_FeatureGroup {
	if x != nil {
		return x.FeatureGroups
	}
	return nil
}

func (x *FeatureView_FeatureRegistrySource) GetProjectNumber() int64 {
	if x != nil && x.ProjectNumber != nil {
		return *x.ProjectNumber
	}
	return 0
}

// Features belonging to a single feature group that will be
// synced to Online Store.
type FeatureView_FeatureRegistrySource_FeatureGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Identifier of the feature group.
	FeatureGroupId string `protobuf:"bytes,1,opt,name=feature_group_id,json=featureGroupId,proto3" json:"feature_group_id,omitempty"`
	// Required. Identifiers of features under the feature group.
	FeatureIds []string `protobuf:"bytes,2,rep,name=feature_ids,json=featureIds,proto3" json:"feature_ids,omitempty"`
}

func (x *FeatureView_FeatureRegistrySource_FeatureGroup) Reset() {
	*x = FeatureView_FeatureRegistrySource_FeatureGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_feature_view_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureView_FeatureRegistrySource_FeatureGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureView_FeatureRegistrySource_FeatureGroup) ProtoMessage() {}

func (x *FeatureView_FeatureRegistrySource_FeatureGroup) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_feature_view_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureView_FeatureRegistrySource_FeatureGroup.ProtoReflect.Descriptor instead.
func (*FeatureView_FeatureRegistrySource_FeatureGroup) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_feature_view_proto_rawDescGZIP(), []int{0, 2, 0}
}

func (x *FeatureView_FeatureRegistrySource_FeatureGroup) GetFeatureGroupId() string {
	if x != nil {
		return x.FeatureGroupId
	}
	return ""
}

func (x *FeatureView_FeatureRegistrySource_FeatureGroup) GetFeatureIds() []string {
	if x != nil {
		return x.FeatureIds
	}
	return nil
}

var File_google_cloud_aiplatform_v1_feature_view_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1_feature_view_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x09, 0x0a, 0x0b, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x12, 0x67, 0x0a, 0x10, 0x62, 0x69, 0x67, 0x5f,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x2e, 0x42, 0x69, 0x67, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48,
	0x00, 0x52, 0x0e, 0x62, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x7c, 0x0a, 0x17, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x2e, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x00, 0x52, 0x15, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x04,
	0x65, 0x74, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x04, 0x65, 0x74, 0x61, 0x67, 0x12, 0x50, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x2e, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x53, 0x0a, 0x0b, 0x73, 0x79, 0x6e, 0x63, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x56, 0x69, 0x65, 0x77, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x0a, 0x73, 0x79, 0x6e, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x58, 0x0a, 0x0e,
	0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x15,
	0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x2f, 0x0a, 0x11, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f,
	0x69, 0x64, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x43,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x1a, 0x20, 0x0a, 0x0a, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x72, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x72, 0x6f, 0x6e, 0x1a, 0xb8, 0x02, 0x0a, 0x15, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x76, 0x0a, 0x0e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56,
	0x69, 0x65, 0x77, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0d, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x2f, 0x0a, 0x0e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x00, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x1a, 0x63, 0x0a, 0x0c, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2d, 0x0a, 0x10, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0e, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0b, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x64, 0x73,
	0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x9b,
	0x01, 0xea, 0x41, 0x97, 0x01, 0x0a, 0x25, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x12, 0x6e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x7b, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7d,
	0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x73, 0x2f, 0x7b, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x7d, 0x42, 0x08, 0x0a, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0xce, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70,
	0x69, 0x76, 0x31, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62,
	0x3b, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0xaa, 0x02, 0x1a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1_feature_view_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1_feature_view_proto_rawDescData = file_google_cloud_aiplatform_v1_feature_view_proto_rawDesc
)

func file_google_cloud_aiplatform_v1_feature_view_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1_feature_view_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1_feature_view_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1_feature_view_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1_feature_view_proto_rawDescData
}

var file_google_cloud_aiplatform_v1_feature_view_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_aiplatform_v1_feature_view_proto_goTypes = []interface{}{
	(*FeatureView)(nil),                       // 0: google.cloud.aiplatform.v1.FeatureView
	(*FeatureView_BigQuerySource)(nil),        // 1: google.cloud.aiplatform.v1.FeatureView.BigQuerySource
	(*FeatureView_SyncConfig)(nil),            // 2: google.cloud.aiplatform.v1.FeatureView.SyncConfig
	(*FeatureView_FeatureRegistrySource)(nil), // 3: google.cloud.aiplatform.v1.FeatureView.FeatureRegistrySource
	nil, // 4: google.cloud.aiplatform.v1.FeatureView.LabelsEntry
	(*FeatureView_FeatureRegistrySource_FeatureGroup)(nil), // 5: google.cloud.aiplatform.v1.FeatureView.FeatureRegistrySource.FeatureGroup
	(*timestamppb.Timestamp)(nil),                          // 6: google.protobuf.Timestamp
}
var file_google_cloud_aiplatform_v1_feature_view_proto_depIdxs = []int32{
	1, // 0: google.cloud.aiplatform.v1.FeatureView.big_query_source:type_name -> google.cloud.aiplatform.v1.FeatureView.BigQuerySource
	3, // 1: google.cloud.aiplatform.v1.FeatureView.feature_registry_source:type_name -> google.cloud.aiplatform.v1.FeatureView.FeatureRegistrySource
	6, // 2: google.cloud.aiplatform.v1.FeatureView.create_time:type_name -> google.protobuf.Timestamp
	6, // 3: google.cloud.aiplatform.v1.FeatureView.update_time:type_name -> google.protobuf.Timestamp
	4, // 4: google.cloud.aiplatform.v1.FeatureView.labels:type_name -> google.cloud.aiplatform.v1.FeatureView.LabelsEntry
	2, // 5: google.cloud.aiplatform.v1.FeatureView.sync_config:type_name -> google.cloud.aiplatform.v1.FeatureView.SyncConfig
	5, // 6: google.cloud.aiplatform.v1.FeatureView.FeatureRegistrySource.feature_groups:type_name -> google.cloud.aiplatform.v1.FeatureView.FeatureRegistrySource.FeatureGroup
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1_feature_view_proto_init() }
func file_google_cloud_aiplatform_v1_feature_view_proto_init() {
	if File_google_cloud_aiplatform_v1_feature_view_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1_feature_view_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureView); i {
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
		file_google_cloud_aiplatform_v1_feature_view_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureView_BigQuerySource); i {
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
		file_google_cloud_aiplatform_v1_feature_view_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureView_SyncConfig); i {
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
		file_google_cloud_aiplatform_v1_feature_view_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureView_FeatureRegistrySource); i {
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
		file_google_cloud_aiplatform_v1_feature_view_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureView_FeatureRegistrySource_FeatureGroup); i {
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
	file_google_cloud_aiplatform_v1_feature_view_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*FeatureView_BigQuerySource_)(nil),
		(*FeatureView_FeatureRegistrySource_)(nil),
	}
	file_google_cloud_aiplatform_v1_feature_view_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_aiplatform_v1_feature_view_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1_feature_view_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1_feature_view_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1_feature_view_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1_feature_view_proto = out.File
	file_google_cloud_aiplatform_v1_feature_view_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1_feature_view_proto_goTypes = nil
	file_google_cloud_aiplatform_v1_feature_view_proto_depIdxs = nil
}
