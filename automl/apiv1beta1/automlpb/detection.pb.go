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
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: google/cloud/automl/v1beta1/detection.proto

package automlpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Annotation details for image object detection.
type ImageObjectDetectionAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The rectangle representing the object location.
	BoundingBox *BoundingPoly `protobuf:"bytes,1,opt,name=bounding_box,json=boundingBox,proto3" json:"bounding_box,omitempty"`
	// Output only. The confidence that this annotation is positive for the parent example,
	// value in [0, 1], higher means higher positivity confidence.
	Score float32 `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *ImageObjectDetectionAnnotation) Reset() {
	*x = ImageObjectDetectionAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_detection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageObjectDetectionAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageObjectDetectionAnnotation) ProtoMessage() {}

func (x *ImageObjectDetectionAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_detection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageObjectDetectionAnnotation.ProtoReflect.Descriptor instead.
func (*ImageObjectDetectionAnnotation) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_detection_proto_rawDescGZIP(), []int{0}
}

func (x *ImageObjectDetectionAnnotation) GetBoundingBox() *BoundingPoly {
	if x != nil {
		return x.BoundingBox
	}
	return nil
}

func (x *ImageObjectDetectionAnnotation) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

// Annotation details for video object tracking.
type VideoObjectTrackingAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The instance of the object, expressed as a positive integer. Used to tell
	// apart objects of the same type (i.e. AnnotationSpec) when multiple are
	// present on a single example.
	// NOTE: Instance ID prediction quality is not a part of model evaluation and
	// is done as best effort. Especially in cases when an entity goes
	// off-screen for a longer time (minutes), when it comes back it may be given
	// a new instance ID.
	InstanceId string `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	// Required. A time (frame) of a video to which this annotation pertains.
	// Represented as the duration since the video's start.
	TimeOffset *durationpb.Duration `protobuf:"bytes,2,opt,name=time_offset,json=timeOffset,proto3" json:"time_offset,omitempty"`
	// Required. The rectangle representing the object location on the frame (i.e.
	// at the time_offset of the video).
	BoundingBox *BoundingPoly `protobuf:"bytes,3,opt,name=bounding_box,json=boundingBox,proto3" json:"bounding_box,omitempty"`
	// Output only. The confidence that this annotation is positive for the video at
	// the time_offset, value in [0, 1], higher means higher positivity
	// confidence. For annotations created by the user the score is 1. When
	// user approves an annotation, the original float score is kept (and not
	// changed to 1).
	Score float32 `protobuf:"fixed32,4,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *VideoObjectTrackingAnnotation) Reset() {
	*x = VideoObjectTrackingAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_detection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoObjectTrackingAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoObjectTrackingAnnotation) ProtoMessage() {}

func (x *VideoObjectTrackingAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_detection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoObjectTrackingAnnotation.ProtoReflect.Descriptor instead.
func (*VideoObjectTrackingAnnotation) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_detection_proto_rawDescGZIP(), []int{1}
}

func (x *VideoObjectTrackingAnnotation) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *VideoObjectTrackingAnnotation) GetTimeOffset() *durationpb.Duration {
	if x != nil {
		return x.TimeOffset
	}
	return nil
}

func (x *VideoObjectTrackingAnnotation) GetBoundingBox() *BoundingPoly {
	if x != nil {
		return x.BoundingBox
	}
	return nil
}

func (x *VideoObjectTrackingAnnotation) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

// Bounding box matching model metrics for a single intersection-over-union
// threshold and multiple label match confidence thresholds.
type BoundingBoxMetricsEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The intersection-over-union threshold value used to compute
	// this metrics entry.
	IouThreshold float32 `protobuf:"fixed32,1,opt,name=iou_threshold,json=iouThreshold,proto3" json:"iou_threshold,omitempty"`
	// Output only. The mean average precision, most often close to au_prc.
	MeanAveragePrecision float32 `protobuf:"fixed32,2,opt,name=mean_average_precision,json=meanAveragePrecision,proto3" json:"mean_average_precision,omitempty"`
	// Output only. Metrics for each label-match confidence_threshold from
	// 0.05,0.10,...,0.95,0.96,0.97,0.98,0.99. Precision-recall curve is
	// derived from them.
	ConfidenceMetricsEntries []*BoundingBoxMetricsEntry_ConfidenceMetricsEntry `protobuf:"bytes,3,rep,name=confidence_metrics_entries,json=confidenceMetricsEntries,proto3" json:"confidence_metrics_entries,omitempty"`
}

func (x *BoundingBoxMetricsEntry) Reset() {
	*x = BoundingBoxMetricsEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_detection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoundingBoxMetricsEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoundingBoxMetricsEntry) ProtoMessage() {}

func (x *BoundingBoxMetricsEntry) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_detection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoundingBoxMetricsEntry.ProtoReflect.Descriptor instead.
func (*BoundingBoxMetricsEntry) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_detection_proto_rawDescGZIP(), []int{2}
}

func (x *BoundingBoxMetricsEntry) GetIouThreshold() float32 {
	if x != nil {
		return x.IouThreshold
	}
	return 0
}

func (x *BoundingBoxMetricsEntry) GetMeanAveragePrecision() float32 {
	if x != nil {
		return x.MeanAveragePrecision
	}
	return 0
}

func (x *BoundingBoxMetricsEntry) GetConfidenceMetricsEntries() []*BoundingBoxMetricsEntry_ConfidenceMetricsEntry {
	if x != nil {
		return x.ConfidenceMetricsEntries
	}
	return nil
}

// Model evaluation metrics for image object detection problems.
// Evaluates prediction quality of labeled bounding boxes.
type ImageObjectDetectionEvaluationMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The total number of bounding boxes (i.e. summed over all
	// images) the ground truth used to create this evaluation had.
	EvaluatedBoundingBoxCount int32 `protobuf:"varint,1,opt,name=evaluated_bounding_box_count,json=evaluatedBoundingBoxCount,proto3" json:"evaluated_bounding_box_count,omitempty"`
	// Output only. The bounding boxes match metrics for each
	// Intersection-over-union threshold 0.05,0.10,...,0.95,0.96,0.97,0.98,0.99
	// and each label confidence threshold 0.05,0.10,...,0.95,0.96,0.97,0.98,0.99
	// pair.
	BoundingBoxMetricsEntries []*BoundingBoxMetricsEntry `protobuf:"bytes,2,rep,name=bounding_box_metrics_entries,json=boundingBoxMetricsEntries,proto3" json:"bounding_box_metrics_entries,omitempty"`
	// Output only. The single metric for bounding boxes evaluation:
	// the mean_average_precision averaged over all bounding_box_metrics_entries.
	BoundingBoxMeanAveragePrecision float32 `protobuf:"fixed32,3,opt,name=bounding_box_mean_average_precision,json=boundingBoxMeanAveragePrecision,proto3" json:"bounding_box_mean_average_precision,omitempty"`
}

func (x *ImageObjectDetectionEvaluationMetrics) Reset() {
	*x = ImageObjectDetectionEvaluationMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_detection_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageObjectDetectionEvaluationMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageObjectDetectionEvaluationMetrics) ProtoMessage() {}

func (x *ImageObjectDetectionEvaluationMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_detection_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageObjectDetectionEvaluationMetrics.ProtoReflect.Descriptor instead.
func (*ImageObjectDetectionEvaluationMetrics) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_detection_proto_rawDescGZIP(), []int{3}
}

func (x *ImageObjectDetectionEvaluationMetrics) GetEvaluatedBoundingBoxCount() int32 {
	if x != nil {
		return x.EvaluatedBoundingBoxCount
	}
	return 0
}

func (x *ImageObjectDetectionEvaluationMetrics) GetBoundingBoxMetricsEntries() []*BoundingBoxMetricsEntry {
	if x != nil {
		return x.BoundingBoxMetricsEntries
	}
	return nil
}

func (x *ImageObjectDetectionEvaluationMetrics) GetBoundingBoxMeanAveragePrecision() float32 {
	if x != nil {
		return x.BoundingBoxMeanAveragePrecision
	}
	return 0
}

// Model evaluation metrics for video object tracking problems.
// Evaluates prediction quality of both labeled bounding boxes and labeled
// tracks (i.e. series of bounding boxes sharing same label and instance ID).
type VideoObjectTrackingEvaluationMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The number of video frames used to create this evaluation.
	EvaluatedFrameCount int32 `protobuf:"varint,1,opt,name=evaluated_frame_count,json=evaluatedFrameCount,proto3" json:"evaluated_frame_count,omitempty"`
	// Output only. The total number of bounding boxes (i.e. summed over all
	// frames) the ground truth used to create this evaluation had.
	EvaluatedBoundingBoxCount int32 `protobuf:"varint,2,opt,name=evaluated_bounding_box_count,json=evaluatedBoundingBoxCount,proto3" json:"evaluated_bounding_box_count,omitempty"`
	// Output only. The bounding boxes match metrics for each
	// Intersection-over-union threshold 0.05,0.10,...,0.95,0.96,0.97,0.98,0.99
	// and each label confidence threshold 0.05,0.10,...,0.95,0.96,0.97,0.98,0.99
	// pair.
	BoundingBoxMetricsEntries []*BoundingBoxMetricsEntry `protobuf:"bytes,4,rep,name=bounding_box_metrics_entries,json=boundingBoxMetricsEntries,proto3" json:"bounding_box_metrics_entries,omitempty"`
	// Output only. The single metric for bounding boxes evaluation:
	// the mean_average_precision averaged over all bounding_box_metrics_entries.
	BoundingBoxMeanAveragePrecision float32 `protobuf:"fixed32,6,opt,name=bounding_box_mean_average_precision,json=boundingBoxMeanAveragePrecision,proto3" json:"bounding_box_mean_average_precision,omitempty"`
}

func (x *VideoObjectTrackingEvaluationMetrics) Reset() {
	*x = VideoObjectTrackingEvaluationMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_detection_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoObjectTrackingEvaluationMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoObjectTrackingEvaluationMetrics) ProtoMessage() {}

func (x *VideoObjectTrackingEvaluationMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_detection_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoObjectTrackingEvaluationMetrics.ProtoReflect.Descriptor instead.
func (*VideoObjectTrackingEvaluationMetrics) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_detection_proto_rawDescGZIP(), []int{4}
}

func (x *VideoObjectTrackingEvaluationMetrics) GetEvaluatedFrameCount() int32 {
	if x != nil {
		return x.EvaluatedFrameCount
	}
	return 0
}

func (x *VideoObjectTrackingEvaluationMetrics) GetEvaluatedBoundingBoxCount() int32 {
	if x != nil {
		return x.EvaluatedBoundingBoxCount
	}
	return 0
}

func (x *VideoObjectTrackingEvaluationMetrics) GetBoundingBoxMetricsEntries() []*BoundingBoxMetricsEntry {
	if x != nil {
		return x.BoundingBoxMetricsEntries
	}
	return nil
}

func (x *VideoObjectTrackingEvaluationMetrics) GetBoundingBoxMeanAveragePrecision() float32 {
	if x != nil {
		return x.BoundingBoxMeanAveragePrecision
	}
	return 0
}

// Metrics for a single confidence threshold.
type BoundingBoxMetricsEntry_ConfidenceMetricsEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The confidence threshold value used to compute the metrics.
	ConfidenceThreshold float32 `protobuf:"fixed32,1,opt,name=confidence_threshold,json=confidenceThreshold,proto3" json:"confidence_threshold,omitempty"`
	// Output only. Recall under the given confidence threshold.
	Recall float32 `protobuf:"fixed32,2,opt,name=recall,proto3" json:"recall,omitempty"`
	// Output only. Precision under the given confidence threshold.
	Precision float32 `protobuf:"fixed32,3,opt,name=precision,proto3" json:"precision,omitempty"`
	// Output only. The harmonic mean of recall and precision.
	F1Score float32 `protobuf:"fixed32,4,opt,name=f1_score,json=f1Score,proto3" json:"f1_score,omitempty"`
}

func (x *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) Reset() {
	*x = BoundingBoxMetricsEntry_ConfidenceMetricsEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_detection_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoundingBoxMetricsEntry_ConfidenceMetricsEntry) ProtoMessage() {}

func (x *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_detection_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoundingBoxMetricsEntry_ConfidenceMetricsEntry.ProtoReflect.Descriptor instead.
func (*BoundingBoxMetricsEntry_ConfidenceMetricsEntry) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_detection_proto_rawDescGZIP(), []int{2, 0}
}

func (x *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) GetConfidenceThreshold() float32 {
	if x != nil {
		return x.ConfidenceThreshold
	}
	return 0
}

func (x *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) GetRecall() float32 {
	if x != nil {
		return x.Recall
	}
	return 0
}

func (x *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) GetPrecision() float32 {
	if x != nil {
		return x.Precision
	}
	return 0
}

func (x *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) GetF1Score() float32 {
	if x != nil {
		return x.F1Score
	}
	return 0
}

var File_google_cloud_automl_v1beta1_detection_proto protoreflect.FileDescriptor

var file_google_cloud_automl_v1beta1_detection_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x64, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x2a, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x1e, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x0a, 0x0c, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x6f, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x42, 0x6f,
	0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x79, 0x52, 0x0b, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0xe0, 0x01,
	0x0a, 0x1d, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x3a, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x4c, 0x0a, 0x0c,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x6f, 0x78, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x79, 0x52, 0x0b, 0x62,
	0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x22, 0x9f, 0x03, 0x0a, 0x17, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x0d,
	0x69, 0x6f, 0x75, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0c, 0x69, 0x6f, 0x75, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x12, 0x34, 0x0a, 0x16, 0x6d, 0x65, 0x61, 0x6e, 0x5f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x5f, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x14, 0x6d, 0x65, 0x61, 0x6e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x50, 0x72,
	0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x89, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x65,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x42, 0x6f, 0x75, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x18, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x1a, 0x9c, 0x01, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e,
	0x63, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x31,
	0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x68, 0x72,
	0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x13, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x61, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x06, 0x72, 0x65, 0x63, 0x61, 0x6c, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x65,
	0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x70, 0x72,
	0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x31, 0x5f, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x66, 0x31, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x22, 0xad, 0x02, 0x0a, 0x25, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x3f, 0x0a, 0x1c,
	0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x5f, 0x62, 0x6f, 0x78, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x19, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x64, 0x42, 0x6f, 0x75,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x75, 0x0a,
	0x1c, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x6f, 0x78, 0x5f, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x19, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x4c, 0x0a, 0x23, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x5f, 0x62, 0x6f, 0x78, 0x5f, 0x6d, 0x65, 0x61, 0x6e, 0x5f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x5f, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x1f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78, 0x4d, 0x65,
	0x61, 0x6e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x50, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0xe0, 0x02, 0x0a, 0x24, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x65,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x65, 0x76, 0x61, 0x6c,
	0x75, 0x61, 0x74, 0x65, 0x64, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x3f, 0x0a, 0x1c, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x6f, 0x78, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x19, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x75, 0x0a, 0x1c, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x6f, 0x78,
	0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x19, 0x62, 0x6f,
	0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x4c, 0x0a, 0x23, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x5f, 0x62, 0x6f, 0x78, 0x5f, 0x6d, 0x65, 0x61, 0x6e, 0x5f, 0x61, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x1f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f,
	0x78, 0x4d, 0x65, 0x61, 0x6e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x50, 0x72, 0x65, 0x63,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x9b, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x50, 0x01, 0x5a, 0x37, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x70, 0x62, 0x3b, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x6c, 0x70, 0x62, 0xca, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x6c, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xea, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x4c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_automl_v1beta1_detection_proto_rawDescOnce sync.Once
	file_google_cloud_automl_v1beta1_detection_proto_rawDescData = file_google_cloud_automl_v1beta1_detection_proto_rawDesc
)

func file_google_cloud_automl_v1beta1_detection_proto_rawDescGZIP() []byte {
	file_google_cloud_automl_v1beta1_detection_proto_rawDescOnce.Do(func() {
		file_google_cloud_automl_v1beta1_detection_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_automl_v1beta1_detection_proto_rawDescData)
	})
	return file_google_cloud_automl_v1beta1_detection_proto_rawDescData
}

var file_google_cloud_automl_v1beta1_detection_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_automl_v1beta1_detection_proto_goTypes = []interface{}{
	(*ImageObjectDetectionAnnotation)(nil),                 // 0: google.cloud.automl.v1beta1.ImageObjectDetectionAnnotation
	(*VideoObjectTrackingAnnotation)(nil),                  // 1: google.cloud.automl.v1beta1.VideoObjectTrackingAnnotation
	(*BoundingBoxMetricsEntry)(nil),                        // 2: google.cloud.automl.v1beta1.BoundingBoxMetricsEntry
	(*ImageObjectDetectionEvaluationMetrics)(nil),          // 3: google.cloud.automl.v1beta1.ImageObjectDetectionEvaluationMetrics
	(*VideoObjectTrackingEvaluationMetrics)(nil),           // 4: google.cloud.automl.v1beta1.VideoObjectTrackingEvaluationMetrics
	(*BoundingBoxMetricsEntry_ConfidenceMetricsEntry)(nil), // 5: google.cloud.automl.v1beta1.BoundingBoxMetricsEntry.ConfidenceMetricsEntry
	(*BoundingPoly)(nil),                                   // 6: google.cloud.automl.v1beta1.BoundingPoly
	(*durationpb.Duration)(nil),                            // 7: google.protobuf.Duration
}
var file_google_cloud_automl_v1beta1_detection_proto_depIdxs = []int32{
	6, // 0: google.cloud.automl.v1beta1.ImageObjectDetectionAnnotation.bounding_box:type_name -> google.cloud.automl.v1beta1.BoundingPoly
	7, // 1: google.cloud.automl.v1beta1.VideoObjectTrackingAnnotation.time_offset:type_name -> google.protobuf.Duration
	6, // 2: google.cloud.automl.v1beta1.VideoObjectTrackingAnnotation.bounding_box:type_name -> google.cloud.automl.v1beta1.BoundingPoly
	5, // 3: google.cloud.automl.v1beta1.BoundingBoxMetricsEntry.confidence_metrics_entries:type_name -> google.cloud.automl.v1beta1.BoundingBoxMetricsEntry.ConfidenceMetricsEntry
	2, // 4: google.cloud.automl.v1beta1.ImageObjectDetectionEvaluationMetrics.bounding_box_metrics_entries:type_name -> google.cloud.automl.v1beta1.BoundingBoxMetricsEntry
	2, // 5: google.cloud.automl.v1beta1.VideoObjectTrackingEvaluationMetrics.bounding_box_metrics_entries:type_name -> google.cloud.automl.v1beta1.BoundingBoxMetricsEntry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_cloud_automl_v1beta1_detection_proto_init() }
func file_google_cloud_automl_v1beta1_detection_proto_init() {
	if File_google_cloud_automl_v1beta1_detection_proto != nil {
		return
	}
	file_google_cloud_automl_v1beta1_geometry_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_automl_v1beta1_detection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageObjectDetectionAnnotation); i {
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
		file_google_cloud_automl_v1beta1_detection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoObjectTrackingAnnotation); i {
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
		file_google_cloud_automl_v1beta1_detection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoundingBoxMetricsEntry); i {
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
		file_google_cloud_automl_v1beta1_detection_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageObjectDetectionEvaluationMetrics); i {
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
		file_google_cloud_automl_v1beta1_detection_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoObjectTrackingEvaluationMetrics); i {
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
		file_google_cloud_automl_v1beta1_detection_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoundingBoxMetricsEntry_ConfidenceMetricsEntry); i {
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
			RawDescriptor: file_google_cloud_automl_v1beta1_detection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_automl_v1beta1_detection_proto_goTypes,
		DependencyIndexes: file_google_cloud_automl_v1beta1_detection_proto_depIdxs,
		MessageInfos:      file_google_cloud_automl_v1beta1_detection_proto_msgTypes,
	}.Build()
	File_google_cloud_automl_v1beta1_detection_proto = out.File
	file_google_cloud_automl_v1beta1_detection_proto_rawDesc = nil
	file_google_cloud_automl_v1beta1_detection_proto_goTypes = nil
	file_google_cloud_automl_v1beta1_detection_proto_depIdxs = nil
}
