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
// source: google/cloud/dialogflow/cx/v3/validation_message.proto

package cxpb

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

// Resource types.
type ValidationMessage_ResourceType int32

const (
	// Unspecified.
	ValidationMessage_RESOURCE_TYPE_UNSPECIFIED ValidationMessage_ResourceType = 0
	// Agent.
	ValidationMessage_AGENT ValidationMessage_ResourceType = 1
	// Intent.
	ValidationMessage_INTENT ValidationMessage_ResourceType = 2
	// Intent training phrase.
	ValidationMessage_INTENT_TRAINING_PHRASE ValidationMessage_ResourceType = 8
	// Intent parameter.
	ValidationMessage_INTENT_PARAMETER ValidationMessage_ResourceType = 9
	// Multiple intents.
	ValidationMessage_INTENTS ValidationMessage_ResourceType = 10
	// Multiple training phrases.
	ValidationMessage_INTENT_TRAINING_PHRASES ValidationMessage_ResourceType = 11
	// Entity type.
	ValidationMessage_ENTITY_TYPE ValidationMessage_ResourceType = 3
	// Multiple entity types.
	ValidationMessage_ENTITY_TYPES ValidationMessage_ResourceType = 12
	// Webhook.
	ValidationMessage_WEBHOOK ValidationMessage_ResourceType = 4
	// Flow.
	ValidationMessage_FLOW ValidationMessage_ResourceType = 5
	// Page.
	ValidationMessage_PAGE ValidationMessage_ResourceType = 6
	// Multiple pages.
	ValidationMessage_PAGES ValidationMessage_ResourceType = 13
	// Transition route group.
	ValidationMessage_TRANSITION_ROUTE_GROUP ValidationMessage_ResourceType = 7
	// Agent transition route group.
	ValidationMessage_AGENT_TRANSITION_ROUTE_GROUP ValidationMessage_ResourceType = 14
)

// Enum value maps for ValidationMessage_ResourceType.
var (
	ValidationMessage_ResourceType_name = map[int32]string{
		0:  "RESOURCE_TYPE_UNSPECIFIED",
		1:  "AGENT",
		2:  "INTENT",
		8:  "INTENT_TRAINING_PHRASE",
		9:  "INTENT_PARAMETER",
		10: "INTENTS",
		11: "INTENT_TRAINING_PHRASES",
		3:  "ENTITY_TYPE",
		12: "ENTITY_TYPES",
		4:  "WEBHOOK",
		5:  "FLOW",
		6:  "PAGE",
		13: "PAGES",
		7:  "TRANSITION_ROUTE_GROUP",
		14: "AGENT_TRANSITION_ROUTE_GROUP",
	}
	ValidationMessage_ResourceType_value = map[string]int32{
		"RESOURCE_TYPE_UNSPECIFIED":    0,
		"AGENT":                        1,
		"INTENT":                       2,
		"INTENT_TRAINING_PHRASE":       8,
		"INTENT_PARAMETER":             9,
		"INTENTS":                      10,
		"INTENT_TRAINING_PHRASES":      11,
		"ENTITY_TYPE":                  3,
		"ENTITY_TYPES":                 12,
		"WEBHOOK":                      4,
		"FLOW":                         5,
		"PAGE":                         6,
		"PAGES":                        13,
		"TRANSITION_ROUTE_GROUP":       7,
		"AGENT_TRANSITION_ROUTE_GROUP": 14,
	}
)

func (x ValidationMessage_ResourceType) Enum() *ValidationMessage_ResourceType {
	p := new(ValidationMessage_ResourceType)
	*p = x
	return p
}

func (x ValidationMessage_ResourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ValidationMessage_ResourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_dialogflow_cx_v3_validation_message_proto_enumTypes[0].Descriptor()
}

func (ValidationMessage_ResourceType) Type() protoreflect.EnumType {
	return &file_google_cloud_dialogflow_cx_v3_validation_message_proto_enumTypes[0]
}

func (x ValidationMessage_ResourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ValidationMessage_ResourceType.Descriptor instead.
func (ValidationMessage_ResourceType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3_validation_message_proto_rawDescGZIP(), []int{0, 0}
}

// Severity level.
type ValidationMessage_Severity int32

const (
	// Unspecified.
	ValidationMessage_SEVERITY_UNSPECIFIED ValidationMessage_Severity = 0
	// The agent doesn't follow Dialogflow best practices.
	ValidationMessage_INFO ValidationMessage_Severity = 1
	// The agent may not behave as expected.
	ValidationMessage_WARNING ValidationMessage_Severity = 2
	// The agent may experience failures.
	ValidationMessage_ERROR ValidationMessage_Severity = 3
)

// Enum value maps for ValidationMessage_Severity.
var (
	ValidationMessage_Severity_name = map[int32]string{
		0: "SEVERITY_UNSPECIFIED",
		1: "INFO",
		2: "WARNING",
		3: "ERROR",
	}
	ValidationMessage_Severity_value = map[string]int32{
		"SEVERITY_UNSPECIFIED": 0,
		"INFO":                 1,
		"WARNING":              2,
		"ERROR":                3,
	}
)

func (x ValidationMessage_Severity) Enum() *ValidationMessage_Severity {
	p := new(ValidationMessage_Severity)
	*p = x
	return p
}

func (x ValidationMessage_Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ValidationMessage_Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_dialogflow_cx_v3_validation_message_proto_enumTypes[1].Descriptor()
}

func (ValidationMessage_Severity) Type() protoreflect.EnumType {
	return &file_google_cloud_dialogflow_cx_v3_validation_message_proto_enumTypes[1]
}

func (x ValidationMessage_Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ValidationMessage_Severity.Descriptor instead.
func (ValidationMessage_Severity) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3_validation_message_proto_rawDescGZIP(), []int{0, 1}
}

// Agent/flow validation message.
type ValidationMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of the resources where the message is found.
	ResourceType ValidationMessage_ResourceType `protobuf:"varint,1,opt,name=resource_type,json=resourceType,proto3,enum=google.cloud.dialogflow.cx.v3.ValidationMessage_ResourceType" json:"resource_type,omitempty"`
	// The names of the resources where the message is found.
	//
	// Deprecated: Marked as deprecated in google/cloud/dialogflow/cx/v3/validation_message.proto.
	Resources []string `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
	// The resource names of the resources where the message is found.
	ResourceNames []*ResourceName `protobuf:"bytes,6,rep,name=resource_names,json=resourceNames,proto3" json:"resource_names,omitempty"`
	// Indicates the severity of the message.
	Severity ValidationMessage_Severity `protobuf:"varint,3,opt,name=severity,proto3,enum=google.cloud.dialogflow.cx.v3.ValidationMessage_Severity" json:"severity,omitempty"`
	// The message detail.
	Detail string `protobuf:"bytes,4,opt,name=detail,proto3" json:"detail,omitempty"`
}

func (x *ValidationMessage) Reset() {
	*x = ValidationMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_cx_v3_validation_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationMessage) ProtoMessage() {}

func (x *ValidationMessage) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3_validation_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationMessage.ProtoReflect.Descriptor instead.
func (*ValidationMessage) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3_validation_message_proto_rawDescGZIP(), []int{0}
}

func (x *ValidationMessage) GetResourceType() ValidationMessage_ResourceType {
	if x != nil {
		return x.ResourceType
	}
	return ValidationMessage_RESOURCE_TYPE_UNSPECIFIED
}

// Deprecated: Marked as deprecated in google/cloud/dialogflow/cx/v3/validation_message.proto.
func (x *ValidationMessage) GetResources() []string {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *ValidationMessage) GetResourceNames() []*ResourceName {
	if x != nil {
		return x.ResourceNames
	}
	return nil
}

func (x *ValidationMessage) GetSeverity() ValidationMessage_Severity {
	if x != nil {
		return x.Severity
	}
	return ValidationMessage_SEVERITY_UNSPECIFIED
}

func (x *ValidationMessage) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

// Resource name and display name.
type ResourceName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Display name.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *ResourceName) Reset() {
	*x = ResourceName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_cx_v3_validation_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceName) ProtoMessage() {}

func (x *ResourceName) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3_validation_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceName.ProtoReflect.Descriptor instead.
func (*ResourceName) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3_validation_message_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceName) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

var File_google_cloud_dialogflow_cx_v3_validation_message_proto protoreflect.FileDescriptor

var file_google_cloud_dialogflow_cx_v3_validation_message_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x78, 0x2f, 0x76, 0x33, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x22, 0xd4, 0x05, 0x0a, 0x11, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x62, 0x0a,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63,
	0x78, 0x2e, 0x76, 0x33, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x20, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x12, 0x52, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f,
	0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x55, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x76, 0x65,
	0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0xad, 0x02, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x45, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x1a, 0x0a,
	0x16, 0x49, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47,
	0x5f, 0x50, 0x48, 0x52, 0x41, 0x53, 0x45, 0x10, 0x08, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x4e, 0x54,
	0x45, 0x4e, 0x54, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x10, 0x09, 0x12,
	0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x0a, 0x12, 0x1b, 0x0a, 0x17,
	0x49, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f,
	0x50, 0x48, 0x52, 0x41, 0x53, 0x45, 0x53, 0x10, 0x0b, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x4e, 0x54,
	0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x4e,
	0x54, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x53, 0x10, 0x0c, 0x12, 0x0b, 0x0a, 0x07,
	0x57, 0x45, 0x42, 0x48, 0x4f, 0x4f, 0x4b, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x4c, 0x4f,
	0x57, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x41, 0x47, 0x45, 0x10, 0x06, 0x12, 0x09, 0x0a,
	0x05, 0x50, 0x41, 0x47, 0x45, 0x53, 0x10, 0x0d, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x52, 0x41, 0x4e,
	0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x5f, 0x47, 0x52, 0x4f,
	0x55, 0x50, 0x10, 0x07, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x5f, 0x47,
	0x52, 0x4f, 0x55, 0x50, 0x10, 0x0e, 0x22, 0x46, 0x0a, 0x08, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x49, 0x4e, 0x46, 0x4f, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e,
	0x47, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x22, 0x45,
	0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0xbc, 0x01, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f,
	0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x42, 0x16, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x69, 0x61, 0x6c, 0x6f,
	0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x33, 0x2f, 0x63,
	0x78, 0x70, 0x62, 0x3b, 0x63, 0x78, 0x70, 0x62, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x02, 0x44, 0x46,
	0xaa, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x78, 0x2e, 0x56, 0x33,
	0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x3a, 0x3a, 0x43, 0x58,
	0x3a, 0x3a, 0x56, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_dialogflow_cx_v3_validation_message_proto_rawDescOnce sync.Once
	file_google_cloud_dialogflow_cx_v3_validation_message_proto_rawDescData = file_google_cloud_dialogflow_cx_v3_validation_message_proto_rawDesc
)

func file_google_cloud_dialogflow_cx_v3_validation_message_proto_rawDescGZIP() []byte {
	file_google_cloud_dialogflow_cx_v3_validation_message_proto_rawDescOnce.Do(func() {
		file_google_cloud_dialogflow_cx_v3_validation_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_dialogflow_cx_v3_validation_message_proto_rawDescData)
	})
	return file_google_cloud_dialogflow_cx_v3_validation_message_proto_rawDescData
}

var file_google_cloud_dialogflow_cx_v3_validation_message_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_dialogflow_cx_v3_validation_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_dialogflow_cx_v3_validation_message_proto_goTypes = []interface{}{
	(ValidationMessage_ResourceType)(0), // 0: google.cloud.dialogflow.cx.v3.ValidationMessage.ResourceType
	(ValidationMessage_Severity)(0),     // 1: google.cloud.dialogflow.cx.v3.ValidationMessage.Severity
	(*ValidationMessage)(nil),           // 2: google.cloud.dialogflow.cx.v3.ValidationMessage
	(*ResourceName)(nil),                // 3: google.cloud.dialogflow.cx.v3.ResourceName
}
var file_google_cloud_dialogflow_cx_v3_validation_message_proto_depIdxs = []int32{
	0, // 0: google.cloud.dialogflow.cx.v3.ValidationMessage.resource_type:type_name -> google.cloud.dialogflow.cx.v3.ValidationMessage.ResourceType
	3, // 1: google.cloud.dialogflow.cx.v3.ValidationMessage.resource_names:type_name -> google.cloud.dialogflow.cx.v3.ResourceName
	1, // 2: google.cloud.dialogflow.cx.v3.ValidationMessage.severity:type_name -> google.cloud.dialogflow.cx.v3.ValidationMessage.Severity
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_dialogflow_cx_v3_validation_message_proto_init() }
func file_google_cloud_dialogflow_cx_v3_validation_message_proto_init() {
	if File_google_cloud_dialogflow_cx_v3_validation_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_dialogflow_cx_v3_validation_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidationMessage); i {
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
		file_google_cloud_dialogflow_cx_v3_validation_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceName); i {
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
			RawDescriptor: file_google_cloud_dialogflow_cx_v3_validation_message_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_dialogflow_cx_v3_validation_message_proto_goTypes,
		DependencyIndexes: file_google_cloud_dialogflow_cx_v3_validation_message_proto_depIdxs,
		EnumInfos:         file_google_cloud_dialogflow_cx_v3_validation_message_proto_enumTypes,
		MessageInfos:      file_google_cloud_dialogflow_cx_v3_validation_message_proto_msgTypes,
	}.Build()
	File_google_cloud_dialogflow_cx_v3_validation_message_proto = out.File
	file_google_cloud_dialogflow_cx_v3_validation_message_proto_rawDesc = nil
	file_google_cloud_dialogflow_cx_v3_validation_message_proto_goTypes = nil
	file_google_cloud_dialogflow_cx_v3_validation_message_proto_depIdxs = nil
}
