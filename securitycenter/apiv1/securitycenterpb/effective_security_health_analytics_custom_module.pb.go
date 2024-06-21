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
// 	protoc-gen-go v1.34.1
// 	protoc        v4.25.3
// source: google/cloud/securitycenter/v1/effective_security_health_analytics_custom_module.proto

package securitycenterpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The enablement state of the module.
type EffectiveSecurityHealthAnalyticsCustomModule_EnablementState int32

const (
	// Unspecified enablement state.
	EffectiveSecurityHealthAnalyticsCustomModule_ENABLEMENT_STATE_UNSPECIFIED EffectiveSecurityHealthAnalyticsCustomModule_EnablementState = 0
	// The module is enabled at the given level.
	EffectiveSecurityHealthAnalyticsCustomModule_ENABLED EffectiveSecurityHealthAnalyticsCustomModule_EnablementState = 1
	// The module is disabled at the given level.
	EffectiveSecurityHealthAnalyticsCustomModule_DISABLED EffectiveSecurityHealthAnalyticsCustomModule_EnablementState = 2
)

// Enum value maps for EffectiveSecurityHealthAnalyticsCustomModule_EnablementState.
var (
	EffectiveSecurityHealthAnalyticsCustomModule_EnablementState_name = map[int32]string{
		0: "ENABLEMENT_STATE_UNSPECIFIED",
		1: "ENABLED",
		2: "DISABLED",
	}
	EffectiveSecurityHealthAnalyticsCustomModule_EnablementState_value = map[string]int32{
		"ENABLEMENT_STATE_UNSPECIFIED": 0,
		"ENABLED":                      1,
		"DISABLED":                     2,
	}
)

func (x EffectiveSecurityHealthAnalyticsCustomModule_EnablementState) Enum() *EffectiveSecurityHealthAnalyticsCustomModule_EnablementState {
	p := new(EffectiveSecurityHealthAnalyticsCustomModule_EnablementState)
	*p = x
	return p
}

func (x EffectiveSecurityHealthAnalyticsCustomModule_EnablementState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EffectiveSecurityHealthAnalyticsCustomModule_EnablementState) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_enumTypes[0].Descriptor()
}

func (EffectiveSecurityHealthAnalyticsCustomModule_EnablementState) Type() protoreflect.EnumType {
	return &file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_enumTypes[0]
}

func (x EffectiveSecurityHealthAnalyticsCustomModule_EnablementState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EffectiveSecurityHealthAnalyticsCustomModule_EnablementState.Descriptor instead.
func (EffectiveSecurityHealthAnalyticsCustomModule_EnablementState) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_rawDescGZIP(), []int{0, 0}
}

// An EffectiveSecurityHealthAnalyticsCustomModule is the representation of
// a Security Health Analytics custom module at a specified level of the
// resource hierarchy: organization, folder, or project. If a custom module is
// inherited from a parent organization or folder, the value of the
// `enablementState` property in EffectiveSecurityHealthAnalyticsCustomModule is
// set to the value that is effective in the parent, instead of  `INHERITED`.
// For example, if the module is enabled in a parent organization or folder, the
// effective enablement_state for the module in all child folders or projects is
// also `enabled`. EffectiveSecurityHealthAnalyticsCustomModule is read-only.
type EffectiveSecurityHealthAnalyticsCustomModule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the custom module.
	// Its format is
	// "organizations/{organization}/securityHealthAnalyticsSettings/effectiveCustomModules/{customModule}",
	// or
	// "folders/{folder}/securityHealthAnalyticsSettings/effectiveCustomModules/{customModule}",
	// or
	// "projects/{project}/securityHealthAnalyticsSettings/effectiveCustomModules/{customModule}"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The user-specified configuration for the module.
	CustomConfig *CustomConfig `protobuf:"bytes,2,opt,name=custom_config,json=customConfig,proto3" json:"custom_config,omitempty"`
	// Output only. The effective state of enablement for the module at the given
	// level of the hierarchy.
	EnablementState EffectiveSecurityHealthAnalyticsCustomModule_EnablementState `protobuf:"varint,3,opt,name=enablement_state,json=enablementState,proto3,enum=google.cloud.securitycenter.v1.EffectiveSecurityHealthAnalyticsCustomModule_EnablementState" json:"enablement_state,omitempty"`
	// Output only. The display name for the custom module. The name must be
	// between 1 and 128 characters, start with a lowercase letter, and contain
	// alphanumeric characters or underscores only.
	DisplayName string `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *EffectiveSecurityHealthAnalyticsCustomModule) Reset() {
	*x = EffectiveSecurityHealthAnalyticsCustomModule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EffectiveSecurityHealthAnalyticsCustomModule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EffectiveSecurityHealthAnalyticsCustomModule) ProtoMessage() {}

func (x *EffectiveSecurityHealthAnalyticsCustomModule) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EffectiveSecurityHealthAnalyticsCustomModule.ProtoReflect.Descriptor instead.
func (*EffectiveSecurityHealthAnalyticsCustomModule) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_rawDescGZIP(), []int{0}
}

func (x *EffectiveSecurityHealthAnalyticsCustomModule) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EffectiveSecurityHealthAnalyticsCustomModule) GetCustomConfig() *CustomConfig {
	if x != nil {
		return x.CustomConfig
	}
	return nil
}

func (x *EffectiveSecurityHealthAnalyticsCustomModule) GetEnablementState() EffectiveSecurityHealthAnalyticsCustomModule_EnablementState {
	if x != nil {
		return x.EnablementState
	}
	return EffectiveSecurityHealthAnalyticsCustomModule_ENABLEMENT_STATE_UNSPECIFIED
}

func (x *EffectiveSecurityHealthAnalyticsCustomModule) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

var File_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_rawDesc = []byte{
	0x0a, 0x56, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb0, 0x06, 0x0a, 0x2c, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x41, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x56, 0x0a, 0x0d,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x8c, 0x01, 0x0a, 0x10, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x5c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x0f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4e, 0x0a, 0x0f, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x20,
	0x0a, 0x1c, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a,
	0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x3a, 0x87, 0x03, 0xea, 0x41,
	0x83, 0x03, 0x0a, 0x4a, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x6d,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x41, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x65, 0x66, 0x66,
	0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x7d, 0x12, 0x61, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x7d, 0x2f,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x41, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f,
	0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x7d,
	0x12, 0x63, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x7d, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x2f, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x65, 0x66, 0x66,
	0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x7d, 0x42, 0x8b, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x31, 0x45, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0xaa, 0x02, 0x1e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xea,
	0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_rawDescData = file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_rawDesc
)

func file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_rawDescData
}

var file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_goTypes = []interface{}{
	(EffectiveSecurityHealthAnalyticsCustomModule_EnablementState)(0), // 0: google.cloud.securitycenter.v1.EffectiveSecurityHealthAnalyticsCustomModule.EnablementState
	(*EffectiveSecurityHealthAnalyticsCustomModule)(nil),              // 1: google.cloud.securitycenter.v1.EffectiveSecurityHealthAnalyticsCustomModule
	(*CustomConfig)(nil), // 2: google.cloud.securitycenter.v1.CustomConfig
}
var file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_depIdxs = []int32{
	2, // 0: google.cloud.securitycenter.v1.EffectiveSecurityHealthAnalyticsCustomModule.custom_config:type_name -> google.cloud.securitycenter.v1.CustomConfig
	0, // 1: google.cloud.securitycenter.v1.EffectiveSecurityHealthAnalyticsCustomModule.enablement_state:type_name -> google.cloud.securitycenter.v1.EffectiveSecurityHealthAnalyticsCustomModule.EnablementState
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_init()
}
func file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_init() {
	if File_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto != nil {
		return
	}
	file_google_cloud_securitycenter_v1_security_health_analytics_custom_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EffectiveSecurityHealthAnalyticsCustomModule); i {
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
			RawDescriptor: file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_depIdxs,
		EnumInfos:         file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_enumTypes,
		MessageInfos:      file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto = out.File
	file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_rawDesc = nil
	file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_goTypes = nil
	file_google_cloud_securitycenter_v1_effective_security_health_analytics_custom_module_proto_depIdxs = nil
}
