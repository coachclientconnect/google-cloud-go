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
// source: google/maps/places/v1/contextual_content.proto

package placespb

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

// Experimental: See
// https://developers.google.com/maps/documentation/places/web-service/experimental/places-generative
// for more details.
//
// Content that is contextual to the place query.
type ContextualContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of reviews about this place, contexual to the place query.
	Reviews []*Review `protobuf:"bytes,1,rep,name=reviews,proto3" json:"reviews,omitempty"`
	// Information (including references) about photos of this place, contexual to
	// the place query.
	Photos []*Photo `protobuf:"bytes,2,rep,name=photos,proto3" json:"photos,omitempty"`
	// Experimental: See
	// https://developers.google.com/maps/documentation/places/web-service/experimental/places-generative
	// for more details.
	//
	// Justifications for the place.
	Justifications []*ContextualContent_Justification `protobuf:"bytes,3,rep,name=justifications,proto3" json:"justifications,omitempty"`
}

func (x *ContextualContent) Reset() {
	*x = ContextualContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_places_v1_contextual_content_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextualContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextualContent) ProtoMessage() {}

func (x *ContextualContent) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_places_v1_contextual_content_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextualContent.ProtoReflect.Descriptor instead.
func (*ContextualContent) Descriptor() ([]byte, []int) {
	return file_google_maps_places_v1_contextual_content_proto_rawDescGZIP(), []int{0}
}

func (x *ContextualContent) GetReviews() []*Review {
	if x != nil {
		return x.Reviews
	}
	return nil
}

func (x *ContextualContent) GetPhotos() []*Photo {
	if x != nil {
		return x.Photos
	}
	return nil
}

func (x *ContextualContent) GetJustifications() []*ContextualContent_Justification {
	if x != nil {
		return x.Justifications
	}
	return nil
}

// Experimental: See
// https://developers.google.com/maps/documentation/places/web-service/experimental/places-generative
// for more details.
//
// Justifications for the place. Justifications answers the question of why a
// place could interest an end user.
type ContextualContent_Justification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Justification:
	//
	//	*ContextualContent_Justification_ReviewJustification_
	//	*ContextualContent_Justification_BusinessAvailabilityAttributesJustification_
	Justification isContextualContent_Justification_Justification `protobuf_oneof:"justification"`
}

func (x *ContextualContent_Justification) Reset() {
	*x = ContextualContent_Justification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_places_v1_contextual_content_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextualContent_Justification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextualContent_Justification) ProtoMessage() {}

func (x *ContextualContent_Justification) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_places_v1_contextual_content_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextualContent_Justification.ProtoReflect.Descriptor instead.
func (*ContextualContent_Justification) Descriptor() ([]byte, []int) {
	return file_google_maps_places_v1_contextual_content_proto_rawDescGZIP(), []int{0, 0}
}

func (m *ContextualContent_Justification) GetJustification() isContextualContent_Justification_Justification {
	if m != nil {
		return m.Justification
	}
	return nil
}

func (x *ContextualContent_Justification) GetReviewJustification() *ContextualContent_Justification_ReviewJustification {
	if x, ok := x.GetJustification().(*ContextualContent_Justification_ReviewJustification_); ok {
		return x.ReviewJustification
	}
	return nil
}

func (x *ContextualContent_Justification) GetBusinessAvailabilityAttributesJustification() *ContextualContent_Justification_BusinessAvailabilityAttributesJustification {
	if x, ok := x.GetJustification().(*ContextualContent_Justification_BusinessAvailabilityAttributesJustification_); ok {
		return x.BusinessAvailabilityAttributesJustification
	}
	return nil
}

type isContextualContent_Justification_Justification interface {
	isContextualContent_Justification_Justification()
}

type ContextualContent_Justification_ReviewJustification_ struct {
	// Experimental: See
	// https://developers.google.com/maps/documentation/places/web-service/experimental/places-generative
	// for more details.
	ReviewJustification *ContextualContent_Justification_ReviewJustification `protobuf:"bytes,1,opt,name=review_justification,json=reviewJustification,proto3,oneof"`
}

type ContextualContent_Justification_BusinessAvailabilityAttributesJustification_ struct {
	// Experimental: See
	// https://developers.google.com/maps/documentation/places/web-service/experimental/places-generative
	// for more details.
	BusinessAvailabilityAttributesJustification *ContextualContent_Justification_BusinessAvailabilityAttributesJustification `protobuf:"bytes,2,opt,name=business_availability_attributes_justification,json=businessAvailabilityAttributesJustification,proto3,oneof"`
}

func (*ContextualContent_Justification_ReviewJustification_) isContextualContent_Justification_Justification() {
}

func (*ContextualContent_Justification_BusinessAvailabilityAttributesJustification_) isContextualContent_Justification_Justification() {
}

// Experimental: See
// https://developers.google.com/maps/documentation/places/web-service/experimental/places-generative
// for more details.
//
// User review justifications. This highlights a section of the user review
// that would interest an end user. For instance, if the search query is
// "firewood pizza", the review justification highlights the text relevant
// to the search query.
type ContextualContent_Justification_ReviewJustification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HighlightedText *ContextualContent_Justification_ReviewJustification_HighlightedText `protobuf:"bytes,1,opt,name=highlighted_text,json=highlightedText,proto3" json:"highlighted_text,omitempty"`
	// The review that the highlighted text is generated from.
	Review *Review `protobuf:"bytes,2,opt,name=review,proto3" json:"review,omitempty"`
}

func (x *ContextualContent_Justification_ReviewJustification) Reset() {
	*x = ContextualContent_Justification_ReviewJustification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_places_v1_contextual_content_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextualContent_Justification_ReviewJustification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextualContent_Justification_ReviewJustification) ProtoMessage() {}

func (x *ContextualContent_Justification_ReviewJustification) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_places_v1_contextual_content_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextualContent_Justification_ReviewJustification.ProtoReflect.Descriptor instead.
func (*ContextualContent_Justification_ReviewJustification) Descriptor() ([]byte, []int) {
	return file_google_maps_places_v1_contextual_content_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *ContextualContent_Justification_ReviewJustification) GetHighlightedText() *ContextualContent_Justification_ReviewJustification_HighlightedText {
	if x != nil {
		return x.HighlightedText
	}
	return nil
}

func (x *ContextualContent_Justification_ReviewJustification) GetReview() *Review {
	if x != nil {
		return x.Review
	}
	return nil
}

// Experimental: See
// https://developers.google.com/maps/documentation/places/web-service/experimental/places-generative
// for more details.
// BusinessAvailabilityAttributes justifications. This shows some attributes
// a business has that could interest an end user.
type ContextualContent_Justification_BusinessAvailabilityAttributesJustification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If a place provides takeout.
	Takeout bool `protobuf:"varint,1,opt,name=takeout,proto3" json:"takeout,omitempty"`
	// If a place provides delivery.
	Delivery bool `protobuf:"varint,2,opt,name=delivery,proto3" json:"delivery,omitempty"`
	// If a place provides dine-in.
	DineIn bool `protobuf:"varint,3,opt,name=dine_in,json=dineIn,proto3" json:"dine_in,omitempty"`
}

func (x *ContextualContent_Justification_BusinessAvailabilityAttributesJustification) Reset() {
	*x = ContextualContent_Justification_BusinessAvailabilityAttributesJustification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_places_v1_contextual_content_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextualContent_Justification_BusinessAvailabilityAttributesJustification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextualContent_Justification_BusinessAvailabilityAttributesJustification) ProtoMessage() {}

func (x *ContextualContent_Justification_BusinessAvailabilityAttributesJustification) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_places_v1_contextual_content_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextualContent_Justification_BusinessAvailabilityAttributesJustification.ProtoReflect.Descriptor instead.
func (*ContextualContent_Justification_BusinessAvailabilityAttributesJustification) Descriptor() ([]byte, []int) {
	return file_google_maps_places_v1_contextual_content_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *ContextualContent_Justification_BusinessAvailabilityAttributesJustification) GetTakeout() bool {
	if x != nil {
		return x.Takeout
	}
	return false
}

func (x *ContextualContent_Justification_BusinessAvailabilityAttributesJustification) GetDelivery() bool {
	if x != nil {
		return x.Delivery
	}
	return false
}

func (x *ContextualContent_Justification_BusinessAvailabilityAttributesJustification) GetDineIn() bool {
	if x != nil {
		return x.DineIn
	}
	return false
}

// The text highlighted by the justification. This is a subset of the
// review itself. The exact word to highlight is marked by the
// HighlightedTextRange. There could be several words in the text being
// highlighted.
type ContextualContent_Justification_ReviewJustification_HighlightedText struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// The list of the ranges of the highlighted text.
	HighlightedTextRanges []*ContextualContent_Justification_ReviewJustification_HighlightedText_HighlightedTextRange `protobuf:"bytes,2,rep,name=highlighted_text_ranges,json=highlightedTextRanges,proto3" json:"highlighted_text_ranges,omitempty"`
}

func (x *ContextualContent_Justification_ReviewJustification_HighlightedText) Reset() {
	*x = ContextualContent_Justification_ReviewJustification_HighlightedText{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_places_v1_contextual_content_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextualContent_Justification_ReviewJustification_HighlightedText) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextualContent_Justification_ReviewJustification_HighlightedText) ProtoMessage() {}

func (x *ContextualContent_Justification_ReviewJustification_HighlightedText) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_places_v1_contextual_content_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextualContent_Justification_ReviewJustification_HighlightedText.ProtoReflect.Descriptor instead.
func (*ContextualContent_Justification_ReviewJustification_HighlightedText) Descriptor() ([]byte, []int) {
	return file_google_maps_places_v1_contextual_content_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

func (x *ContextualContent_Justification_ReviewJustification_HighlightedText) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ContextualContent_Justification_ReviewJustification_HighlightedText) GetHighlightedTextRanges() []*ContextualContent_Justification_ReviewJustification_HighlightedText_HighlightedTextRange {
	if x != nil {
		return x.HighlightedTextRanges
	}
	return nil
}

// The range of highlighted text.
type ContextualContent_Justification_ReviewJustification_HighlightedText_HighlightedTextRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartIndex int32 `protobuf:"varint,1,opt,name=start_index,json=startIndex,proto3" json:"start_index,omitempty"`
	EndIndex   int32 `protobuf:"varint,2,opt,name=end_index,json=endIndex,proto3" json:"end_index,omitempty"`
}

func (x *ContextualContent_Justification_ReviewJustification_HighlightedText_HighlightedTextRange) Reset() {
	*x = ContextualContent_Justification_ReviewJustification_HighlightedText_HighlightedTextRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_places_v1_contextual_content_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextualContent_Justification_ReviewJustification_HighlightedText_HighlightedTextRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextualContent_Justification_ReviewJustification_HighlightedText_HighlightedTextRange) ProtoMessage() {
}

func (x *ContextualContent_Justification_ReviewJustification_HighlightedText_HighlightedTextRange) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_places_v1_contextual_content_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextualContent_Justification_ReviewJustification_HighlightedText_HighlightedTextRange.ProtoReflect.Descriptor instead.
func (*ContextualContent_Justification_ReviewJustification_HighlightedText_HighlightedTextRange) Descriptor() ([]byte, []int) {
	return file_google_maps_places_v1_contextual_content_proto_rawDescGZIP(), []int{0, 0, 0, 0, 0}
}

func (x *ContextualContent_Justification_ReviewJustification_HighlightedText_HighlightedTextRange) GetStartIndex() int32 {
	if x != nil {
		return x.StartIndex
	}
	return 0
}

func (x *ContextualContent_Justification_ReviewJustification_HighlightedText_HighlightedTextRange) GetEndIndex() int32 {
	if x != nil {
		return x.EndIndex
	}
	return 0
}

var File_google_maps_places_v1_contextual_content_proto protoreflect.FileDescriptor

var file_google_maps_places_v1_contextual_content_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x75,
	0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x6d, 0x61, 0x70, 0x73, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x68, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf,
	0x09, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x75, 0x61, 0x6c, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d,
	0x61, 0x70, 0x73, 0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x52, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x34, 0x0a,
	0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x06, 0x70, 0x68, 0x6f,
	0x74, 0x6f, 0x73, 0x12, 0x5e, 0x0a, 0x0e, 0x6a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x75, 0x61, 0x6c, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x4a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x6a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x1a, 0xea, 0x07, 0x0a, 0x0d, 0x4a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x7f, 0x0a, 0x14, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f,
	0x6a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x4a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70,
	0x73, 0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x75, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x4a, 0x75,
	0x73, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x4a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x00, 0x52, 0x13, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x4a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xc9, 0x01, 0x0a, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x5f, 0x6a, 0x75, 0x73, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x62, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x75,
	0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x4a, 0x75, 0x73, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x4a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x2b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x4a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0xfc, 0x03, 0x0a, 0x13, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x4a, 0x75, 0x73,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x85, 0x01, 0x0a, 0x10, 0x68,
	0x69, 0x67, 0x68, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x5a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d,
	0x61, 0x70, 0x73, 0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x75, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e,
	0x4a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x4a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x48, 0x69, 0x67, 0x68, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x54, 0x65, 0x78,
	0x74, 0x52, 0x0f, 0x68, 0x69, 0x67, 0x68, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x54, 0x65,
	0x78, 0x74, 0x12, 0x35, 0x0a, 0x06, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73,
	0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x52, 0x06, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x1a, 0xa5, 0x02, 0x0a, 0x0f, 0x48, 0x69,
	0x67, 0x68, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x12, 0xa7, 0x01, 0x0a, 0x17, 0x68, 0x69, 0x67, 0x68, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x65,
	0x64, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x6f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70,
	0x73, 0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x75, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x4a, 0x75,
	0x73, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x4a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x48, 0x69, 0x67, 0x68, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x2e,
	0x48, 0x69, 0x67, 0x68, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x15, 0x68, 0x69, 0x67, 0x68, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x65,
	0x64, 0x54, 0x65, 0x78, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x1a, 0x54, 0x0a, 0x14, 0x48,
	0x69, 0x67, 0x68, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x1a, 0x7c, 0x0a, 0x2b, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x4a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x6b, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x74, 0x61, 0x6b, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x69, 0x6e, 0x65, 0x5f, 0x69,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x64, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x42,
	0x0f, 0x0a, 0x0d, 0x6a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0xaa, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x6d, 0x61, 0x70, 0x73, 0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x16,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x75, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x61,
	0x70, 0x73, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x70, 0x62, 0x3b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x70,
	0x62, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x06, 0x47, 0x4d, 0x50, 0x53, 0x56, 0x31, 0xaa, 0x02, 0x15,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x15, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x4d,
	0x61, 0x70, 0x73, 0x5c, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_places_v1_contextual_content_proto_rawDescOnce sync.Once
	file_google_maps_places_v1_contextual_content_proto_rawDescData = file_google_maps_places_v1_contextual_content_proto_rawDesc
)

func file_google_maps_places_v1_contextual_content_proto_rawDescGZIP() []byte {
	file_google_maps_places_v1_contextual_content_proto_rawDescOnce.Do(func() {
		file_google_maps_places_v1_contextual_content_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_places_v1_contextual_content_proto_rawDescData)
	})
	return file_google_maps_places_v1_contextual_content_proto_rawDescData
}

var file_google_maps_places_v1_contextual_content_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_maps_places_v1_contextual_content_proto_goTypes = []any{
	(*ContextualContent)(nil),                                                                        // 0: google.maps.places.v1.ContextualContent
	(*ContextualContent_Justification)(nil),                                                          // 1: google.maps.places.v1.ContextualContent.Justification
	(*ContextualContent_Justification_ReviewJustification)(nil),                                      // 2: google.maps.places.v1.ContextualContent.Justification.ReviewJustification
	(*ContextualContent_Justification_BusinessAvailabilityAttributesJustification)(nil),              // 3: google.maps.places.v1.ContextualContent.Justification.BusinessAvailabilityAttributesJustification
	(*ContextualContent_Justification_ReviewJustification_HighlightedText)(nil),                      // 4: google.maps.places.v1.ContextualContent.Justification.ReviewJustification.HighlightedText
	(*ContextualContent_Justification_ReviewJustification_HighlightedText_HighlightedTextRange)(nil), // 5: google.maps.places.v1.ContextualContent.Justification.ReviewJustification.HighlightedText.HighlightedTextRange
	(*Review)(nil), // 6: google.maps.places.v1.Review
	(*Photo)(nil),  // 7: google.maps.places.v1.Photo
}
var file_google_maps_places_v1_contextual_content_proto_depIdxs = []int32{
	6, // 0: google.maps.places.v1.ContextualContent.reviews:type_name -> google.maps.places.v1.Review
	7, // 1: google.maps.places.v1.ContextualContent.photos:type_name -> google.maps.places.v1.Photo
	1, // 2: google.maps.places.v1.ContextualContent.justifications:type_name -> google.maps.places.v1.ContextualContent.Justification
	2, // 3: google.maps.places.v1.ContextualContent.Justification.review_justification:type_name -> google.maps.places.v1.ContextualContent.Justification.ReviewJustification
	3, // 4: google.maps.places.v1.ContextualContent.Justification.business_availability_attributes_justification:type_name -> google.maps.places.v1.ContextualContent.Justification.BusinessAvailabilityAttributesJustification
	4, // 5: google.maps.places.v1.ContextualContent.Justification.ReviewJustification.highlighted_text:type_name -> google.maps.places.v1.ContextualContent.Justification.ReviewJustification.HighlightedText
	6, // 6: google.maps.places.v1.ContextualContent.Justification.ReviewJustification.review:type_name -> google.maps.places.v1.Review
	5, // 7: google.maps.places.v1.ContextualContent.Justification.ReviewJustification.HighlightedText.highlighted_text_ranges:type_name -> google.maps.places.v1.ContextualContent.Justification.ReviewJustification.HighlightedText.HighlightedTextRange
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_google_maps_places_v1_contextual_content_proto_init() }
func file_google_maps_places_v1_contextual_content_proto_init() {
	if File_google_maps_places_v1_contextual_content_proto != nil {
		return
	}
	file_google_maps_places_v1_photo_proto_init()
	file_google_maps_places_v1_review_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_maps_places_v1_contextual_content_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ContextualContent); i {
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
		file_google_maps_places_v1_contextual_content_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ContextualContent_Justification); i {
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
		file_google_maps_places_v1_contextual_content_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ContextualContent_Justification_ReviewJustification); i {
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
		file_google_maps_places_v1_contextual_content_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ContextualContent_Justification_BusinessAvailabilityAttributesJustification); i {
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
		file_google_maps_places_v1_contextual_content_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ContextualContent_Justification_ReviewJustification_HighlightedText); i {
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
		file_google_maps_places_v1_contextual_content_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ContextualContent_Justification_ReviewJustification_HighlightedText_HighlightedTextRange); i {
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
	file_google_maps_places_v1_contextual_content_proto_msgTypes[1].OneofWrappers = []any{
		(*ContextualContent_Justification_ReviewJustification_)(nil),
		(*ContextualContent_Justification_BusinessAvailabilityAttributesJustification_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_maps_places_v1_contextual_content_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_places_v1_contextual_content_proto_goTypes,
		DependencyIndexes: file_google_maps_places_v1_contextual_content_proto_depIdxs,
		MessageInfos:      file_google_maps_places_v1_contextual_content_proto_msgTypes,
	}.Build()
	File_google_maps_places_v1_contextual_content_proto = out.File
	file_google_maps_places_v1_contextual_content_proto_rawDesc = nil
	file_google_maps_places_v1_contextual_content_proto_goTypes = nil
	file_google_maps_places_v1_contextual_content_proto_depIdxs = nil
}
