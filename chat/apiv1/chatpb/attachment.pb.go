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
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: google/chat/v1/attachment.proto

package chatpb

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

// The source of the attachment.
type Attachment_Source int32

const (
	// Reserved.
	Attachment_SOURCE_UNSPECIFIED Attachment_Source = 0
	// The file is a Google Drive file.
	Attachment_DRIVE_FILE Attachment_Source = 1
	// The file is uploaded to Chat.
	Attachment_UPLOADED_CONTENT Attachment_Source = 2
)

// Enum value maps for Attachment_Source.
var (
	Attachment_Source_name = map[int32]string{
		0: "SOURCE_UNSPECIFIED",
		1: "DRIVE_FILE",
		2: "UPLOADED_CONTENT",
	}
	Attachment_Source_value = map[string]int32{
		"SOURCE_UNSPECIFIED": 0,
		"DRIVE_FILE":         1,
		"UPLOADED_CONTENT":   2,
	}
)

func (x Attachment_Source) Enum() *Attachment_Source {
	p := new(Attachment_Source)
	*p = x
	return p
}

func (x Attachment_Source) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Attachment_Source) Descriptor() protoreflect.EnumDescriptor {
	return file_google_chat_v1_attachment_proto_enumTypes[0].Descriptor()
}

func (Attachment_Source) Type() protoreflect.EnumType {
	return &file_google_chat_v1_attachment_proto_enumTypes[0]
}

func (x Attachment_Source) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Attachment_Source.Descriptor instead.
func (Attachment_Source) EnumDescriptor() ([]byte, []int) {
	return file_google_chat_v1_attachment_proto_rawDescGZIP(), []int{0, 0}
}

// An attachment in Google Chat.
type Attachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name of the attachment, in the form
	// `spaces/*/messages/*/attachments/*`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The original file name for the content, not the full path.
	ContentName string `protobuf:"bytes,2,opt,name=content_name,json=contentName,proto3" json:"content_name,omitempty"`
	// Output only. The content type (MIME type) of the file.
	ContentType string `protobuf:"bytes,3,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// The data reference to the attachment.
	//
	// Types that are assignable to DataRef:
	//
	//	*Attachment_AttachmentDataRef
	//	*Attachment_DriveDataRef
	DataRef isAttachment_DataRef `protobuf_oneof:"data_ref"`
	// Output only. The thumbnail URL which should be used to preview the
	// attachment to a human user. Chat apps shouldn't use this URL to download
	// attachment content.
	ThumbnailUri string `protobuf:"bytes,5,opt,name=thumbnail_uri,json=thumbnailUri,proto3" json:"thumbnail_uri,omitempty"`
	// Output only. The download URL which should be used to allow a human user to
	// download the attachment. Chat apps shouldn't use this URL to download
	// attachment content.
	DownloadUri string `protobuf:"bytes,6,opt,name=download_uri,json=downloadUri,proto3" json:"download_uri,omitempty"`
	// Output only. The source of the attachment.
	Source Attachment_Source `protobuf:"varint,9,opt,name=source,proto3,enum=google.chat.v1.Attachment_Source" json:"source,omitempty"`
}

func (x *Attachment) Reset() {
	*x = Attachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_chat_v1_attachment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attachment) ProtoMessage() {}

func (x *Attachment) ProtoReflect() protoreflect.Message {
	mi := &file_google_chat_v1_attachment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attachment.ProtoReflect.Descriptor instead.
func (*Attachment) Descriptor() ([]byte, []int) {
	return file_google_chat_v1_attachment_proto_rawDescGZIP(), []int{0}
}

func (x *Attachment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Attachment) GetContentName() string {
	if x != nil {
		return x.ContentName
	}
	return ""
}

func (x *Attachment) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (m *Attachment) GetDataRef() isAttachment_DataRef {
	if m != nil {
		return m.DataRef
	}
	return nil
}

func (x *Attachment) GetAttachmentDataRef() *AttachmentDataRef {
	if x, ok := x.GetDataRef().(*Attachment_AttachmentDataRef); ok {
		return x.AttachmentDataRef
	}
	return nil
}

func (x *Attachment) GetDriveDataRef() *DriveDataRef {
	if x, ok := x.GetDataRef().(*Attachment_DriveDataRef); ok {
		return x.DriveDataRef
	}
	return nil
}

func (x *Attachment) GetThumbnailUri() string {
	if x != nil {
		return x.ThumbnailUri
	}
	return ""
}

func (x *Attachment) GetDownloadUri() string {
	if x != nil {
		return x.DownloadUri
	}
	return ""
}

func (x *Attachment) GetSource() Attachment_Source {
	if x != nil {
		return x.Source
	}
	return Attachment_SOURCE_UNSPECIFIED
}

type isAttachment_DataRef interface {
	isAttachment_DataRef()
}

type Attachment_AttachmentDataRef struct {
	// A reference to the attachment data. This field is used with the media API
	// to download the attachment data.
	AttachmentDataRef *AttachmentDataRef `protobuf:"bytes,4,opt,name=attachment_data_ref,json=attachmentDataRef,proto3,oneof"`
}

type Attachment_DriveDataRef struct {
	// Output only. A reference to the Google Drive attachment. This field is
	// used with the Google Drive API.
	DriveDataRef *DriveDataRef `protobuf:"bytes,7,opt,name=drive_data_ref,json=driveDataRef,proto3,oneof"`
}

func (*Attachment_AttachmentDataRef) isAttachment_DataRef() {}

func (*Attachment_DriveDataRef) isAttachment_DataRef() {}

// A reference to the data of a drive attachment.
type DriveDataRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID for the drive file. Use with the Drive API.
	DriveFileId string `protobuf:"bytes,2,opt,name=drive_file_id,json=driveFileId,proto3" json:"drive_file_id,omitempty"`
}

func (x *DriveDataRef) Reset() {
	*x = DriveDataRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_chat_v1_attachment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriveDataRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriveDataRef) ProtoMessage() {}

func (x *DriveDataRef) ProtoReflect() protoreflect.Message {
	mi := &file_google_chat_v1_attachment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriveDataRef.ProtoReflect.Descriptor instead.
func (*DriveDataRef) Descriptor() ([]byte, []int) {
	return file_google_chat_v1_attachment_proto_rawDescGZIP(), []int{1}
}

func (x *DriveDataRef) GetDriveFileId() string {
	if x != nil {
		return x.DriveFileId
	}
	return ""
}

// A reference to the attachment data.
type AttachmentDataRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the attachment data. This field is used with the media
	// API to download the attachment data.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Opaque token containing a reference to an uploaded attachment. Treated by
	// clients as an opaque string and used to create or update Chat messages with
	// attachments.
	AttachmentUploadToken string `protobuf:"bytes,2,opt,name=attachment_upload_token,json=attachmentUploadToken,proto3" json:"attachment_upload_token,omitempty"`
}

func (x *AttachmentDataRef) Reset() {
	*x = AttachmentDataRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_chat_v1_attachment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachmentDataRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachmentDataRef) ProtoMessage() {}

func (x *AttachmentDataRef) ProtoReflect() protoreflect.Message {
	mi := &file_google_chat_v1_attachment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachmentDataRef.ProtoReflect.Descriptor instead.
func (*AttachmentDataRef) Descriptor() ([]byte, []int) {
	return file_google_chat_v1_attachment_proto_rawDescGZIP(), []int{2}
}

func (x *AttachmentDataRef) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AttachmentDataRef) GetAttachmentUploadToken() string {
	if x != nil {
		return x.AttachmentUploadToken
	}
	return ""
}

// Request to get an attachment.
type GetAttachmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Resource name of the attachment, in the form
	// `spaces/*/messages/*/attachments/*`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetAttachmentRequest) Reset() {
	*x = GetAttachmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_chat_v1_attachment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAttachmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAttachmentRequest) ProtoMessage() {}

func (x *GetAttachmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_chat_v1_attachment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAttachmentRequest.ProtoReflect.Descriptor instead.
func (*GetAttachmentRequest) Descriptor() ([]byte, []int) {
	return file_google_chat_v1_attachment_proto_rawDescGZIP(), []int{3}
}

func (x *GetAttachmentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request to upload an attachment.
type UploadAttachmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Resource name of the Chat space in which the attachment is
	// uploaded. Format "spaces/{space}".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The filename of the attachment, including the file extension.
	Filename string `protobuf:"bytes,4,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *UploadAttachmentRequest) Reset() {
	*x = UploadAttachmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_chat_v1_attachment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadAttachmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadAttachmentRequest) ProtoMessage() {}

func (x *UploadAttachmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_chat_v1_attachment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadAttachmentRequest.ProtoReflect.Descriptor instead.
func (*UploadAttachmentRequest) Descriptor() ([]byte, []int) {
	return file_google_chat_v1_attachment_proto_rawDescGZIP(), []int{4}
}

func (x *UploadAttachmentRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *UploadAttachmentRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

// Response of uploading an attachment.
type UploadAttachmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reference to the uploaded attachment.
	AttachmentDataRef *AttachmentDataRef `protobuf:"bytes,1,opt,name=attachment_data_ref,json=attachmentDataRef,proto3" json:"attachment_data_ref,omitempty"`
}

func (x *UploadAttachmentResponse) Reset() {
	*x = UploadAttachmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_chat_v1_attachment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadAttachmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadAttachmentResponse) ProtoMessage() {}

func (x *UploadAttachmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_chat_v1_attachment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadAttachmentResponse.ProtoReflect.Descriptor instead.
func (*UploadAttachmentResponse) Descriptor() ([]byte, []int) {
	return file_google_chat_v1_attachment_proto_rawDescGZIP(), []int{5}
}

func (x *UploadAttachmentResponse) GetAttachmentDataRef() *AttachmentDataRef {
	if x != nil {
		return x.AttachmentDataRef
	}
	return nil
}

var File_google_chat_v1_attachment_proto protoreflect.FileDescriptor

var file_google_chat_v1_attachment_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x04,
	0x0a, 0x0a, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x26, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x53, 0x0a, 0x13, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x66,
	0x48, 0x00, 0x52, 0x11, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x66, 0x12, 0x49, 0x0a, 0x0e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x66, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x48, 0x00, 0x52, 0x0c, 0x64, 0x72, 0x69, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x66,
	0x12, 0x28, 0x0a, 0x0d, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x75, 0x72,
	0x69, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0c, 0x74, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x72, 0x69, 0x12, 0x26, 0x0a, 0x0c, 0x64, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55,
	0x72, 0x69, 0x12, 0x3e, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x22, 0x46, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x12,
	0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x52, 0x49, 0x56, 0x45, 0x5f, 0x46, 0x49,
	0x4c, 0x45, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x45, 0x44,
	0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x3a, 0x5f, 0xea, 0x41, 0x5c, 0x0a,
	0x1e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x3a, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x7d, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x7d, 0x2f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x7d, 0x42, 0x0a, 0x0a, 0x08, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x72, 0x65, 0x66, 0x22, 0x32, 0x0a, 0x0c, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x66, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x70, 0x0a, 0x11, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x66,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x17, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x52, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x26, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x20, 0x0a, 0x1e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x77, 0x0a, 0x17, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0xe0, 0x41,
	0x02, 0xfa, 0x41, 0x1d, 0x12, 0x1b, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x6d, 0x0a, 0x18, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x13, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x66, 0x52, 0x11, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x66, 0x42, 0xa8, 0x01, 0x0a, 0x12, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31,
	0x42, 0x0f, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x61, 0x70,
	0x69, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x3b, 0x63, 0x68, 0x61, 0x74, 0x70,
	0x62, 0xa2, 0x02, 0x0b, 0x44, 0x59, 0x4e, 0x41, 0x50, 0x49, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xaa,
	0x02, 0x13, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x73, 0x2e, 0x43, 0x68,
	0x61, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x13, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41,
	0x70, 0x70, 0x73, 0x5c, 0x43, 0x68, 0x61, 0x74, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x16, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x70, 0x70, 0x73, 0x3a, 0x3a, 0x43, 0x68, 0x61, 0x74,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_chat_v1_attachment_proto_rawDescOnce sync.Once
	file_google_chat_v1_attachment_proto_rawDescData = file_google_chat_v1_attachment_proto_rawDesc
)

func file_google_chat_v1_attachment_proto_rawDescGZIP() []byte {
	file_google_chat_v1_attachment_proto_rawDescOnce.Do(func() {
		file_google_chat_v1_attachment_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_chat_v1_attachment_proto_rawDescData)
	})
	return file_google_chat_v1_attachment_proto_rawDescData
}

var file_google_chat_v1_attachment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_chat_v1_attachment_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_chat_v1_attachment_proto_goTypes = []interface{}{
	(Attachment_Source)(0),           // 0: google.chat.v1.Attachment.Source
	(*Attachment)(nil),               // 1: google.chat.v1.Attachment
	(*DriveDataRef)(nil),             // 2: google.chat.v1.DriveDataRef
	(*AttachmentDataRef)(nil),        // 3: google.chat.v1.AttachmentDataRef
	(*GetAttachmentRequest)(nil),     // 4: google.chat.v1.GetAttachmentRequest
	(*UploadAttachmentRequest)(nil),  // 5: google.chat.v1.UploadAttachmentRequest
	(*UploadAttachmentResponse)(nil), // 6: google.chat.v1.UploadAttachmentResponse
}
var file_google_chat_v1_attachment_proto_depIdxs = []int32{
	3, // 0: google.chat.v1.Attachment.attachment_data_ref:type_name -> google.chat.v1.AttachmentDataRef
	2, // 1: google.chat.v1.Attachment.drive_data_ref:type_name -> google.chat.v1.DriveDataRef
	0, // 2: google.chat.v1.Attachment.source:type_name -> google.chat.v1.Attachment.Source
	3, // 3: google.chat.v1.UploadAttachmentResponse.attachment_data_ref:type_name -> google.chat.v1.AttachmentDataRef
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_chat_v1_attachment_proto_init() }
func file_google_chat_v1_attachment_proto_init() {
	if File_google_chat_v1_attachment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_chat_v1_attachment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attachment); i {
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
		file_google_chat_v1_attachment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriveDataRef); i {
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
		file_google_chat_v1_attachment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachmentDataRef); i {
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
		file_google_chat_v1_attachment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAttachmentRequest); i {
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
		file_google_chat_v1_attachment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadAttachmentRequest); i {
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
		file_google_chat_v1_attachment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadAttachmentResponse); i {
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
	file_google_chat_v1_attachment_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Attachment_AttachmentDataRef)(nil),
		(*Attachment_DriveDataRef)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_chat_v1_attachment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_chat_v1_attachment_proto_goTypes,
		DependencyIndexes: file_google_chat_v1_attachment_proto_depIdxs,
		EnumInfos:         file_google_chat_v1_attachment_proto_enumTypes,
		MessageInfos:      file_google_chat_v1_attachment_proto_msgTypes,
	}.Build()
	File_google_chat_v1_attachment_proto = out.File
	file_google_chat_v1_attachment_proto_rawDesc = nil
	file_google_chat_v1_attachment_proto_goTypes = nil
	file_google_chat_v1_attachment_proto_depIdxs = nil
}
