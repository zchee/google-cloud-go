// Copyright 2024 Google LLC
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
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: google/cloud/aiplatform/v1beta1/feature_view_sync.proto

package aiplatformpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	interval "google.golang.org/genproto/googleapis/type/interval"
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

// FeatureViewSync is a representation of sync operation which copies data from
// data source to Feature View in Online Store.
type FeatureViewSync struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier. Name of the FeatureViewSync. Format:
	// `projects/{project}/locations/{location}/featureOnlineStores/{feature_online_store}/featureViews/{feature_view}/featureViewSyncs/{feature_view_sync}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Time when this FeatureViewSync is created. Creation of a
	// FeatureViewSync means that the job is pending / waiting for sufficient
	// resources but may not have started the actual data transfer yet.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Time when this FeatureViewSync is finished.
	RunTime *interval.Interval `protobuf:"bytes,5,opt,name=run_time,json=runTime,proto3" json:"run_time,omitempty"`
	// Output only. Final status of the FeatureViewSync.
	FinalStatus *status.Status `protobuf:"bytes,4,opt,name=final_status,json=finalStatus,proto3" json:"final_status,omitempty"`
	// Output only. Summary of the sync job.
	SyncSummary *FeatureViewSync_SyncSummary `protobuf:"bytes,6,opt,name=sync_summary,json=syncSummary,proto3" json:"sync_summary,omitempty"`
	// Output only. Reserved for future use.
	SatisfiesPzs bool `protobuf:"varint,7,opt,name=satisfies_pzs,json=satisfiesPzs,proto3" json:"satisfies_pzs,omitempty"`
	// Output only. Reserved for future use.
	SatisfiesPzi bool `protobuf:"varint,8,opt,name=satisfies_pzi,json=satisfiesPzi,proto3" json:"satisfies_pzi,omitempty"`
}

func (x *FeatureViewSync) Reset() {
	*x = FeatureViewSync{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureViewSync) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureViewSync) ProtoMessage() {}

func (x *FeatureViewSync) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureViewSync.ProtoReflect.Descriptor instead.
func (*FeatureViewSync) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_rawDescGZIP(), []int{0}
}

func (x *FeatureViewSync) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FeatureViewSync) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *FeatureViewSync) GetRunTime() *interval.Interval {
	if x != nil {
		return x.RunTime
	}
	return nil
}

func (x *FeatureViewSync) GetFinalStatus() *status.Status {
	if x != nil {
		return x.FinalStatus
	}
	return nil
}

func (x *FeatureViewSync) GetSyncSummary() *FeatureViewSync_SyncSummary {
	if x != nil {
		return x.SyncSummary
	}
	return nil
}

func (x *FeatureViewSync) GetSatisfiesPzs() bool {
	if x != nil {
		return x.SatisfiesPzs
	}
	return false
}

func (x *FeatureViewSync) GetSatisfiesPzi() bool {
	if x != nil {
		return x.SatisfiesPzi
	}
	return false
}

// Summary from the Sync job. For continuous syncs, the summary is updated
// periodically. For batch syncs, it gets updated on completion of the sync.
type FeatureViewSync_SyncSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Total number of rows synced.
	RowSynced int64 `protobuf:"varint,1,opt,name=row_synced,json=rowSynced,proto3" json:"row_synced,omitempty"`
	// Output only. BigQuery slot milliseconds consumed for the sync job.
	TotalSlot int64 `protobuf:"varint,2,opt,name=total_slot,json=totalSlot,proto3" json:"total_slot,omitempty"`
}

func (x *FeatureViewSync_SyncSummary) Reset() {
	*x = FeatureViewSync_SyncSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureViewSync_SyncSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureViewSync_SyncSummary) ProtoMessage() {}

func (x *FeatureViewSync_SyncSummary) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureViewSync_SyncSummary.ProtoReflect.Descriptor instead.
func (*FeatureViewSync_SyncSummary) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_rawDescGZIP(), []int{0, 0}
}

func (x *FeatureViewSync_SyncSummary) GetRowSynced() int64 {
	if x != nil {
		return x.RowSynced
	}
	return 0
}

func (x *FeatureViewSync_SyncSummary) GetTotalSlot() int64 {
	if x != nil {
		return x.TotalSlot
	}
	return 0
}

var File_google_cloud_aiplatform_v1beta1_feature_view_sync_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x73,
	0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x05, 0x0a,
	0x0f, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x53, 0x79, 0x6e, 0x63,
	0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x72,
	0x75, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x0b, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x64,
	0x0a, 0x0c, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69,
	0x65, 0x77, 0x53, 0x79, 0x6e, 0x63, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x73, 0x79, 0x6e, 0x63, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x12, 0x28, 0x0a, 0x0d, 0x73, 0x61, 0x74, 0x69, 0x73, 0x66, 0x69, 0x65,
	0x73, 0x5f, 0x70, 0x7a, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0c, 0x73, 0x61, 0x74, 0x69, 0x73, 0x66, 0x69, 0x65, 0x73, 0x50, 0x7a, 0x73, 0x12, 0x28,
	0x0a, 0x0d, 0x73, 0x61, 0x74, 0x69, 0x73, 0x66, 0x69, 0x65, 0x73, 0x5f, 0x70, 0x7a, 0x69, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0c, 0x73, 0x61, 0x74, 0x69,
	0x73, 0x66, 0x69, 0x65, 0x73, 0x50, 0x7a, 0x69, 0x1a, 0x55, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x22, 0x0a, 0x0a, 0x72, 0x6f, 0x77, 0x5f, 0x73,
	0x79, 0x6e, 0x63, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x09, 0x72, 0x6f, 0x77, 0x53, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x6c, 0x6f, 0x74, 0x3a,
	0xc3, 0x01, 0xea, 0x41, 0xbf, 0x01, 0x0a, 0x29, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x53, 0x79, 0x6e,
	0x63, 0x12, 0x91, 0x01, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f,
	0x7b, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x7d, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69,
	0x65, 0x77, 0x73, 0x2f, 0x7b, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x76, 0x69, 0x65,
	0x77, 0x7d, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x53, 0x79,
	0x6e, 0x63, 0x73, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x76, 0x69, 0x65, 0x77,
	0x5f, 0x73, 0x79, 0x6e, 0x63, 0x42, 0xeb, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x14, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0x3b, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x1f, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02,
	0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x41, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_rawDescData = file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_rawDesc
)

func file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_rawDescData
}

var file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_goTypes = []any{
	(*FeatureViewSync)(nil),             // 0: google.cloud.aiplatform.v1beta1.FeatureViewSync
	(*FeatureViewSync_SyncSummary)(nil), // 1: google.cloud.aiplatform.v1beta1.FeatureViewSync.SyncSummary
	(*timestamppb.Timestamp)(nil),       // 2: google.protobuf.Timestamp
	(*interval.Interval)(nil),           // 3: google.type.Interval
	(*status.Status)(nil),               // 4: google.rpc.Status
}
var file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_depIdxs = []int32{
	2, // 0: google.cloud.aiplatform.v1beta1.FeatureViewSync.create_time:type_name -> google.protobuf.Timestamp
	3, // 1: google.cloud.aiplatform.v1beta1.FeatureViewSync.run_time:type_name -> google.type.Interval
	4, // 2: google.cloud.aiplatform.v1beta1.FeatureViewSync.final_status:type_name -> google.rpc.Status
	1, // 3: google.cloud.aiplatform.v1beta1.FeatureViewSync.sync_summary:type_name -> google.cloud.aiplatform.v1beta1.FeatureViewSync.SyncSummary
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_init() }
func file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_init() {
	if File_google_cloud_aiplatform_v1beta1_feature_view_sync_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*FeatureViewSync); i {
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
		file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*FeatureViewSync_SyncSummary); i {
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
			RawDescriptor: file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1beta1_feature_view_sync_proto = out.File
	file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_goTypes = nil
	file_google_cloud_aiplatform_v1beta1_feature_view_sync_proto_depIdxs = nil
}
