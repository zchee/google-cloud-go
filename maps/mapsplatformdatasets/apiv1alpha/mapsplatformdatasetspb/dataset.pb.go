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
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: google/maps/mapsplatformdatasets/v1alpha/dataset.proto

package mapsplatformdatasetspb

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

// Usage specifies where the data is intended to be used to inform how to
// process the data.
type Usage int32

const (
	// The usage of this dataset is not set.
	Usage_USAGE_UNSPECIFIED Usage = 0
	// This dataset will be used for data driven styling.
	Usage_USAGE_DATA_DRIVEN_STYLING Usage = 1
	// This dataset will be used for area affordances in routing.
	Usage_USAGE_AREA_AFFORDANCES Usage = 2
	// This dataset will be used for assisted driving in routing.
	Usage_USAGE_ASSISTED_DRIVING Usage = 3
)

// Enum value maps for Usage.
var (
	Usage_name = map[int32]string{
		0: "USAGE_UNSPECIFIED",
		1: "USAGE_DATA_DRIVEN_STYLING",
		2: "USAGE_AREA_AFFORDANCES",
		3: "USAGE_ASSISTED_DRIVING",
	}
	Usage_value = map[string]int32{
		"USAGE_UNSPECIFIED":         0,
		"USAGE_DATA_DRIVEN_STYLING": 1,
		"USAGE_AREA_AFFORDANCES":    2,
		"USAGE_ASSISTED_DRIVING":    3,
	}
)

func (x Usage) Enum() *Usage {
	p := new(Usage)
	*p = x
	return p
}

func (x Usage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Usage) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_enumTypes[0].Descriptor()
}

func (Usage) Type() protoreflect.EnumType {
	return &file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_enumTypes[0]
}

func (x Usage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Usage.Descriptor instead.
func (Usage) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_rawDescGZIP(), []int{0}
}

// State specifies the status of the import of the latest dataset version.
type State int32

const (
	// The state of this dataset is not set.
	State_STATE_UNSPECIFIED State = 0
	// The dataset version is getting imported.
	State_STATE_IMPORTING State = 1
	// The dataset version succeeded in getting imported.
	State_STATE_IMPORT_SUCCEEDED State = 2
	// The dataset version failed to get imported.
	State_STATE_IMPORT_FAILED State = 3
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "STATE_IMPORTING",
		2: "STATE_IMPORT_SUCCEEDED",
		3: "STATE_IMPORT_FAILED",
	}
	State_value = map[string]int32{
		"STATE_UNSPECIFIED":      0,
		"STATE_IMPORTING":        1,
		"STATE_IMPORT_SUCCEEDED": 2,
		"STATE_IMPORT_FAILED":    3,
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
	return file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_enumTypes[1].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_enumTypes[1]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_rawDescGZIP(), []int{1}
}

// A representation of a maps platform dataset.
type Dataset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name,
	// projects/{project}/datasets/{dataset_id}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Human readable name, shown in the console UI. Set by customer.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// A description of this dataset; set by the customer.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// The version of the dataset.
	VersionId string `protobuf:"bytes,4,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	// Specified use case(s) for this dataset.
	Usage []Usage `protobuf:"varint,5,rep,packed,name=usage,proto3,enum=google.maps.mapsplatformdatasets.v1alpha.Usage" json:"usage,omitempty"`
	// Details about the source of the data for the dataset.
	//
	// Types that are assignable to DataSource:
	//
	//	*Dataset_LocalFileSource
	//	*Dataset_GcsSource
	DataSource isDataset_DataSource `protobuf_oneof:"data_source"`
	// The status of the import of the latest dataset version.
	Status State `protobuf:"varint,12,opt,name=status,proto3,enum=google.maps.mapsplatformdatasets.v1alpha.State" json:"status,omitempty"`
	// Output only. Time when the dataset was first created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Time when the dataset metadata was last updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. Time when this version of dataset was created. (It happened when importing
	// data to the dataset)
	VersionCreateTime *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=version_create_time,json=versionCreateTime,proto3" json:"version_create_time,omitempty"`
	// Output only. The description for this version of dataset. It is provided when importing
	// data to the dataset.
	VersionDescription string `protobuf:"bytes,11,opt,name=version_description,json=versionDescription,proto3" json:"version_description,omitempty"`
}

func (x *Dataset) Reset() {
	*x = Dataset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dataset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dataset) ProtoMessage() {}

func (x *Dataset) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dataset.ProtoReflect.Descriptor instead.
func (*Dataset) Descriptor() ([]byte, []int) {
	return file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_rawDescGZIP(), []int{0}
}

func (x *Dataset) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Dataset) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Dataset) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Dataset) GetVersionId() string {
	if x != nil {
		return x.VersionId
	}
	return ""
}

func (x *Dataset) GetUsage() []Usage {
	if x != nil {
		return x.Usage
	}
	return nil
}

func (m *Dataset) GetDataSource() isDataset_DataSource {
	if m != nil {
		return m.DataSource
	}
	return nil
}

func (x *Dataset) GetLocalFileSource() *LocalFileSource {
	if x, ok := x.GetDataSource().(*Dataset_LocalFileSource); ok {
		return x.LocalFileSource
	}
	return nil
}

func (x *Dataset) GetGcsSource() *GcsSource {
	if x, ok := x.GetDataSource().(*Dataset_GcsSource); ok {
		return x.GcsSource
	}
	return nil
}

func (x *Dataset) GetStatus() State {
	if x != nil {
		return x.Status
	}
	return State_STATE_UNSPECIFIED
}

func (x *Dataset) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Dataset) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Dataset) GetVersionCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.VersionCreateTime
	}
	return nil
}

func (x *Dataset) GetVersionDescription() string {
	if x != nil {
		return x.VersionDescription
	}
	return ""
}

type isDataset_DataSource interface {
	isDataset_DataSource()
}

type Dataset_LocalFileSource struct {
	// A local file source for the dataset for a single upload.
	LocalFileSource *LocalFileSource `protobuf:"bytes,6,opt,name=local_file_source,json=localFileSource,proto3,oneof"`
}

type Dataset_GcsSource struct {
	// A Google Cloud Storage file source for the dataset for a single upload.
	GcsSource *GcsSource `protobuf:"bytes,7,opt,name=gcs_source,json=gcsSource,proto3,oneof"`
}

func (*Dataset_LocalFileSource) isDataset_DataSource() {}

func (*Dataset_GcsSource) isDataset_DataSource() {}

var File_google_maps_mapsplatformdatasets_v1alpha_dataset_proto protoreflect.FileDescriptor

var file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x6d, 0x61,
	0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x6d, 0x61, 0x70, 0x73,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x06, 0x0a, 0x07,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x45, 0x0a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x2f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70,
	0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x12, 0x67, 0x0a, 0x11, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e,
	0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0f,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x54, 0x0a, 0x0a, 0x67, 0x63, 0x73, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70,
	0x73, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47,
	0x63, 0x73, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x09, 0x67, 0x63, 0x73, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d,
	0x61, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x40,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x4f, 0x0a, 0x13, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x11, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x13, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x12, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x57, 0xea, 0x41, 0x54, 0x0a, 0x2b,
	0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x25, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x7d, 0x42, 0x0d, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2a, 0x75, 0x0a, 0x05, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x53,
	0x41, 0x47, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x1d, 0x0a, 0x19, 0x55, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f,
	0x44, 0x52, 0x49, 0x56, 0x45, 0x4e, 0x5f, 0x53, 0x54, 0x59, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x01,
	0x12, 0x1a, 0x0a, 0x16, 0x55, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x41, 0x52, 0x45, 0x41, 0x5f, 0x41,
	0x46, 0x46, 0x4f, 0x52, 0x44, 0x41, 0x4e, 0x43, 0x45, 0x53, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16,
	0x55, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x41, 0x53, 0x53, 0x49, 0x53, 0x54, 0x45, 0x44, 0x5f, 0x44,
	0x52, 0x49, 0x56, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x2a, 0x68, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x1a, 0x0a,
	0x16, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44,
	0x10, 0x03, 0x42, 0xfc, 0x01, 0x0a, 0x2c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x42, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x66, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x6d, 0x61,
	0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d, 0x61,
	0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x73, 0x70, 0x62, 0x3b, 0x6d, 0x61, 0x70, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x70, 0x62, 0xaa, 0x02, 0x28, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x56,
	0x31, 0x41, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x28, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x4d, 0x61, 0x70, 0x73, 0x5c, 0x4d, 0x61, 0x70, 0x73, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_rawDescOnce sync.Once
	file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_rawDescData = file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_rawDesc
)

func file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_rawDescGZIP() []byte {
	file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_rawDescOnce.Do(func() {
		file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_rawDescData)
	})
	return file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_rawDescData
}

var file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_goTypes = []interface{}{
	(Usage)(0),                    // 0: google.maps.mapsplatformdatasets.v1alpha.Usage
	(State)(0),                    // 1: google.maps.mapsplatformdatasets.v1alpha.State
	(*Dataset)(nil),               // 2: google.maps.mapsplatformdatasets.v1alpha.Dataset
	(*LocalFileSource)(nil),       // 3: google.maps.mapsplatformdatasets.v1alpha.LocalFileSource
	(*GcsSource)(nil),             // 4: google.maps.mapsplatformdatasets.v1alpha.GcsSource
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_depIdxs = []int32{
	0, // 0: google.maps.mapsplatformdatasets.v1alpha.Dataset.usage:type_name -> google.maps.mapsplatformdatasets.v1alpha.Usage
	3, // 1: google.maps.mapsplatformdatasets.v1alpha.Dataset.local_file_source:type_name -> google.maps.mapsplatformdatasets.v1alpha.LocalFileSource
	4, // 2: google.maps.mapsplatformdatasets.v1alpha.Dataset.gcs_source:type_name -> google.maps.mapsplatformdatasets.v1alpha.GcsSource
	1, // 3: google.maps.mapsplatformdatasets.v1alpha.Dataset.status:type_name -> google.maps.mapsplatformdatasets.v1alpha.State
	5, // 4: google.maps.mapsplatformdatasets.v1alpha.Dataset.create_time:type_name -> google.protobuf.Timestamp
	5, // 5: google.maps.mapsplatformdatasets.v1alpha.Dataset.update_time:type_name -> google.protobuf.Timestamp
	5, // 6: google.maps.mapsplatformdatasets.v1alpha.Dataset.version_create_time:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_init() }
func file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_init() {
	if File_google_maps_mapsplatformdatasets_v1alpha_dataset_proto != nil {
		return
	}
	file_google_maps_mapsplatformdatasets_v1alpha_data_source_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dataset); i {
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
	file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Dataset_LocalFileSource)(nil),
		(*Dataset_GcsSource)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_goTypes,
		DependencyIndexes: file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_depIdxs,
		EnumInfos:         file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_enumTypes,
		MessageInfos:      file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_msgTypes,
	}.Build()
	File_google_maps_mapsplatformdatasets_v1alpha_dataset_proto = out.File
	file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_rawDesc = nil
	file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_goTypes = nil
	file_google_maps_mapsplatformdatasets_v1alpha_dataset_proto_depIdxs = nil
}
