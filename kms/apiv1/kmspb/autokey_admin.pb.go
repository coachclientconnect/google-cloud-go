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
// source: google/cloud/kms/v1/autokey_admin.proto

package kmspb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for
// [UpdateAutokeyConfig][google.cloud.kms.v1.AutokeyAdmin.UpdateAutokeyConfig].
type UpdateAutokeyConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. [AutokeyConfig][google.cloud.kms.v1.AutokeyConfig] with values to
	// update.
	AutokeyConfig *AutokeyConfig `protobuf:"bytes,1,opt,name=autokey_config,json=autokeyConfig,proto3" json:"autokey_config,omitempty"`
	// Required. Masks which fields of the
	// [AutokeyConfig][google.cloud.kms.v1.AutokeyConfig] to update, e.g.
	// `keyProject`.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateAutokeyConfigRequest) Reset() {
	*x = UpdateAutokeyConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_kms_v1_autokey_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAutokeyConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAutokeyConfigRequest) ProtoMessage() {}

func (x *UpdateAutokeyConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_kms_v1_autokey_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAutokeyConfigRequest.ProtoReflect.Descriptor instead.
func (*UpdateAutokeyConfigRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_kms_v1_autokey_admin_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateAutokeyConfigRequest) GetAutokeyConfig() *AutokeyConfig {
	if x != nil {
		return x.AutokeyConfig
	}
	return nil
}

func (x *UpdateAutokeyConfigRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// Request message for
// [GetAutokeyConfig][google.cloud.kms.v1.AutokeyAdmin.GetAutokeyConfig].
type GetAutokeyConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Name of the [AutokeyConfig][google.cloud.kms.v1.AutokeyConfig]
	// resource, e.g. `folders/{FOLDER_NUMBER}/autokeyConfig`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetAutokeyConfigRequest) Reset() {
	*x = GetAutokeyConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_kms_v1_autokey_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAutokeyConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAutokeyConfigRequest) ProtoMessage() {}

func (x *GetAutokeyConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_kms_v1_autokey_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAutokeyConfigRequest.ProtoReflect.Descriptor instead.
func (*GetAutokeyConfigRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_kms_v1_autokey_admin_proto_rawDescGZIP(), []int{1}
}

func (x *GetAutokeyConfigRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Cloud KMS Autokey configuration for a folder.
type AutokeyConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier. Name of the [AutokeyConfig][google.cloud.kms.v1.AutokeyConfig]
	// resource, e.g. `folders/{FOLDER_NUMBER}/autokeyConfig`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. Name of the key project, e.g. `projects/{PROJECT_ID}` or
	// `projects/{PROJECT_NUMBER}`, where Cloud KMS Autokey will provision a new
	// [CryptoKey][google.cloud.kms.v1.CryptoKey] when a
	// [KeyHandle][google.cloud.kms.v1.KeyHandle] is created. On
	// [UpdateAutokeyConfig][google.cloud.kms.v1.AutokeyAdmin.UpdateAutokeyConfig],
	// the caller will require `cloudkms.cryptoKeys.setIamPolicy` permission on
	// this key project. Once configured, for Cloud KMS Autokey to function
	// properly, this key project must have the Cloud KMS API activated and the
	// Cloud KMS Service Agent for this key project must be granted the
	// `cloudkms.admin` role (or pertinent permissions). A request with an empty
	// key project field will clear the configuration.
	KeyProject string `protobuf:"bytes,2,opt,name=key_project,json=keyProject,proto3" json:"key_project,omitempty"`
}

func (x *AutokeyConfig) Reset() {
	*x = AutokeyConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_kms_v1_autokey_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutokeyConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutokeyConfig) ProtoMessage() {}

func (x *AutokeyConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_kms_v1_autokey_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutokeyConfig.ProtoReflect.Descriptor instead.
func (*AutokeyConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_kms_v1_autokey_admin_proto_rawDescGZIP(), []int{2}
}

func (x *AutokeyConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AutokeyConfig) GetKeyProject() string {
	if x != nil {
		return x.KeyProject
	}
	return ""
}

// Request message for
// [ShowEffectiveAutokeyConfig][google.cloud.kms.v1.AutokeyAdmin.ShowEffectiveAutokeyConfig].
type ShowEffectiveAutokeyConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Name of the resource project to the show effective Cloud KMS
	// Autokey configuration for. This may be helpful for interrogating the effect
	// of nested folder configurations on a given resource project.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
}

func (x *ShowEffectiveAutokeyConfigRequest) Reset() {
	*x = ShowEffectiveAutokeyConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_kms_v1_autokey_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowEffectiveAutokeyConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowEffectiveAutokeyConfigRequest) ProtoMessage() {}

func (x *ShowEffectiveAutokeyConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_kms_v1_autokey_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowEffectiveAutokeyConfigRequest.ProtoReflect.Descriptor instead.
func (*ShowEffectiveAutokeyConfigRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_kms_v1_autokey_admin_proto_rawDescGZIP(), []int{3}
}

func (x *ShowEffectiveAutokeyConfigRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

// Response message for
// [ShowEffectiveAutokeyConfig][google.cloud.kms.v1.AutokeyAdmin.ShowEffectiveAutokeyConfig].
type ShowEffectiveAutokeyConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the key project configured in the resource project's folder
	// ancestry.
	KeyProject string `protobuf:"bytes,1,opt,name=key_project,json=keyProject,proto3" json:"key_project,omitempty"`
}

func (x *ShowEffectiveAutokeyConfigResponse) Reset() {
	*x = ShowEffectiveAutokeyConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_kms_v1_autokey_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowEffectiveAutokeyConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowEffectiveAutokeyConfigResponse) ProtoMessage() {}

func (x *ShowEffectiveAutokeyConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_kms_v1_autokey_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowEffectiveAutokeyConfigResponse.ProtoReflect.Descriptor instead.
func (*ShowEffectiveAutokeyConfigResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_kms_v1_autokey_admin_proto_rawDescGZIP(), []int{4}
}

func (x *ShowEffectiveAutokeyConfigResponse) GetKeyProject() string {
	if x != nil {
		return x.KeyProject
	}
	return ""
}

var File_google_cloud_kms_v1_autokey_admin_proto protoreflect.FileDescriptor

var file_google_cloud_kms_v1_autokey_admin_proto_rawDesc = []byte{
	0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6b,
	0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x65, 0x79, 0x5f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75,
	0x74, 0x6f, 0x6b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x4e, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x65, 0x79, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x73, 0x6b, 0x22, 0x5c, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x6f, 0x6b,
	0x65, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x41, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2d, 0xe0,
	0x41, 0x02, 0xfa, 0x41, 0x27, 0x0a, 0x25, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6b, 0x6d, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41,
	0x75, 0x74, 0x6f, 0x6b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0xb9, 0x01, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x65, 0x79, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a,
	0x0b, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0a, 0x6b, 0x65, 0x79, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x3a, 0x69, 0xea, 0x41, 0x66, 0x0a, 0x25, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6b,
	0x6d, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x1e, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x7d, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2a,
	0x0e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x32,
	0x0d, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x70,
	0x0a, 0x21, 0x53, 0x68, 0x6f, 0x77, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x41,
	0x75, 0x74, 0x6f, 0x6b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x33, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2d, 0x0a, 0x2b, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x22, 0x45, 0x0a, 0x22, 0x53, 0x68, 0x6f, 0x77, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6b, 0x65, 0x79,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x32, 0xc8, 0x05, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x6f,
	0x6b, 0x65, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0xd2, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74,
	0x6f, 0x6b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x65, 0x79, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x66, 0xda, 0x41, 0x1a, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x65,
	0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x43, 0x3a, 0x0e, 0x61, 0x75, 0x74, 0x6f,
	0x6b, 0x65, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x32, 0x31, 0x2f, 0x76, 0x31, 0x2f,
	0x7b, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x65, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x7d, 0x12, 0x97, 0x01,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x6f,
	0x6b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x65, 0x79, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x22, 0x31, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x24, 0x12, 0x22, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x65, 0x79,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x7d, 0x12, 0xd2, 0x01, 0x0a, 0x1a, 0x53, 0x68, 0x6f, 0x77,
	0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x65, 0x79,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x6f,
	0x77, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x65,
	0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x6d,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x6f, 0x77, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x43, 0xda, 0x41, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x34, 0x12, 0x32, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a,
	0x7d, 0x3a, 0x73, 0x68, 0x6f, 0x77, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x41,
	0x75, 0x74, 0x6f, 0x6b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x74, 0xca, 0x41,
	0x17, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6b, 0x6d, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x57, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6b,
	0x6d, 0x73, 0x42, 0x59, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x41,
	0x75, 0x74, 0x6f, 0x6b, 0x65, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x29, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6b, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x76,
	0x31, 0x2f, 0x6b, 0x6d, 0x73, 0x70, 0x62, 0x3b, 0x6b, 0x6d, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_kms_v1_autokey_admin_proto_rawDescOnce sync.Once
	file_google_cloud_kms_v1_autokey_admin_proto_rawDescData = file_google_cloud_kms_v1_autokey_admin_proto_rawDesc
)

func file_google_cloud_kms_v1_autokey_admin_proto_rawDescGZIP() []byte {
	file_google_cloud_kms_v1_autokey_admin_proto_rawDescOnce.Do(func() {
		file_google_cloud_kms_v1_autokey_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_kms_v1_autokey_admin_proto_rawDescData)
	})
	return file_google_cloud_kms_v1_autokey_admin_proto_rawDescData
}

var file_google_cloud_kms_v1_autokey_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_kms_v1_autokey_admin_proto_goTypes = []interface{}{
	(*UpdateAutokeyConfigRequest)(nil),         // 0: google.cloud.kms.v1.UpdateAutokeyConfigRequest
	(*GetAutokeyConfigRequest)(nil),            // 1: google.cloud.kms.v1.GetAutokeyConfigRequest
	(*AutokeyConfig)(nil),                      // 2: google.cloud.kms.v1.AutokeyConfig
	(*ShowEffectiveAutokeyConfigRequest)(nil),  // 3: google.cloud.kms.v1.ShowEffectiveAutokeyConfigRequest
	(*ShowEffectiveAutokeyConfigResponse)(nil), // 4: google.cloud.kms.v1.ShowEffectiveAutokeyConfigResponse
	(*fieldmaskpb.FieldMask)(nil),              // 5: google.protobuf.FieldMask
}
var file_google_cloud_kms_v1_autokey_admin_proto_depIdxs = []int32{
	2, // 0: google.cloud.kms.v1.UpdateAutokeyConfigRequest.autokey_config:type_name -> google.cloud.kms.v1.AutokeyConfig
	5, // 1: google.cloud.kms.v1.UpdateAutokeyConfigRequest.update_mask:type_name -> google.protobuf.FieldMask
	0, // 2: google.cloud.kms.v1.AutokeyAdmin.UpdateAutokeyConfig:input_type -> google.cloud.kms.v1.UpdateAutokeyConfigRequest
	1, // 3: google.cloud.kms.v1.AutokeyAdmin.GetAutokeyConfig:input_type -> google.cloud.kms.v1.GetAutokeyConfigRequest
	3, // 4: google.cloud.kms.v1.AutokeyAdmin.ShowEffectiveAutokeyConfig:input_type -> google.cloud.kms.v1.ShowEffectiveAutokeyConfigRequest
	2, // 5: google.cloud.kms.v1.AutokeyAdmin.UpdateAutokeyConfig:output_type -> google.cloud.kms.v1.AutokeyConfig
	2, // 6: google.cloud.kms.v1.AutokeyAdmin.GetAutokeyConfig:output_type -> google.cloud.kms.v1.AutokeyConfig
	4, // 7: google.cloud.kms.v1.AutokeyAdmin.ShowEffectiveAutokeyConfig:output_type -> google.cloud.kms.v1.ShowEffectiveAutokeyConfigResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_kms_v1_autokey_admin_proto_init() }
func file_google_cloud_kms_v1_autokey_admin_proto_init() {
	if File_google_cloud_kms_v1_autokey_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_kms_v1_autokey_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAutokeyConfigRequest); i {
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
		file_google_cloud_kms_v1_autokey_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAutokeyConfigRequest); i {
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
		file_google_cloud_kms_v1_autokey_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutokeyConfig); i {
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
		file_google_cloud_kms_v1_autokey_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowEffectiveAutokeyConfigRequest); i {
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
		file_google_cloud_kms_v1_autokey_admin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowEffectiveAutokeyConfigResponse); i {
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
			RawDescriptor: file_google_cloud_kms_v1_autokey_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_kms_v1_autokey_admin_proto_goTypes,
		DependencyIndexes: file_google_cloud_kms_v1_autokey_admin_proto_depIdxs,
		MessageInfos:      file_google_cloud_kms_v1_autokey_admin_proto_msgTypes,
	}.Build()
	File_google_cloud_kms_v1_autokey_admin_proto = out.File
	file_google_cloud_kms_v1_autokey_admin_proto_rawDesc = nil
	file_google_cloud_kms_v1_autokey_admin_proto_goTypes = nil
	file_google_cloud_kms_v1_autokey_admin_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AutokeyAdminClient is the client API for AutokeyAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AutokeyAdminClient interface {
	// Updates the [AutokeyConfig][google.cloud.kms.v1.AutokeyConfig] for a
	// folder. The caller must have both `cloudkms.autokeyConfigs.update`
	// permission on the parent folder and `cloudkms.cryptoKeys.setIamPolicy`
	// permission on the provided key project. A
	// [KeyHandle][google.cloud.kms.v1.KeyHandle] creation in the folder's
	// descendant projects will use this configuration to determine where to
	// create the resulting [CryptoKey][google.cloud.kms.v1.CryptoKey].
	UpdateAutokeyConfig(ctx context.Context, in *UpdateAutokeyConfigRequest, opts ...grpc.CallOption) (*AutokeyConfig, error)
	// Returns the [AutokeyConfig][google.cloud.kms.v1.AutokeyConfig] for a
	// folder.
	GetAutokeyConfig(ctx context.Context, in *GetAutokeyConfigRequest, opts ...grpc.CallOption) (*AutokeyConfig, error)
	// Returns the effective Cloud KMS Autokey configuration for a given project.
	ShowEffectiveAutokeyConfig(ctx context.Context, in *ShowEffectiveAutokeyConfigRequest, opts ...grpc.CallOption) (*ShowEffectiveAutokeyConfigResponse, error)
}

type autokeyAdminClient struct {
	cc grpc.ClientConnInterface
}

func NewAutokeyAdminClient(cc grpc.ClientConnInterface) AutokeyAdminClient {
	return &autokeyAdminClient{cc}
}

func (c *autokeyAdminClient) UpdateAutokeyConfig(ctx context.Context, in *UpdateAutokeyConfigRequest, opts ...grpc.CallOption) (*AutokeyConfig, error) {
	out := new(AutokeyConfig)
	err := c.cc.Invoke(ctx, "/google.cloud.kms.v1.AutokeyAdmin/UpdateAutokeyConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autokeyAdminClient) GetAutokeyConfig(ctx context.Context, in *GetAutokeyConfigRequest, opts ...grpc.CallOption) (*AutokeyConfig, error) {
	out := new(AutokeyConfig)
	err := c.cc.Invoke(ctx, "/google.cloud.kms.v1.AutokeyAdmin/GetAutokeyConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autokeyAdminClient) ShowEffectiveAutokeyConfig(ctx context.Context, in *ShowEffectiveAutokeyConfigRequest, opts ...grpc.CallOption) (*ShowEffectiveAutokeyConfigResponse, error) {
	out := new(ShowEffectiveAutokeyConfigResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.kms.v1.AutokeyAdmin/ShowEffectiveAutokeyConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AutokeyAdminServer is the server API for AutokeyAdmin service.
type AutokeyAdminServer interface {
	// Updates the [AutokeyConfig][google.cloud.kms.v1.AutokeyConfig] for a
	// folder. The caller must have both `cloudkms.autokeyConfigs.update`
	// permission on the parent folder and `cloudkms.cryptoKeys.setIamPolicy`
	// permission on the provided key project. A
	// [KeyHandle][google.cloud.kms.v1.KeyHandle] creation in the folder's
	// descendant projects will use this configuration to determine where to
	// create the resulting [CryptoKey][google.cloud.kms.v1.CryptoKey].
	UpdateAutokeyConfig(context.Context, *UpdateAutokeyConfigRequest) (*AutokeyConfig, error)
	// Returns the [AutokeyConfig][google.cloud.kms.v1.AutokeyConfig] for a
	// folder.
	GetAutokeyConfig(context.Context, *GetAutokeyConfigRequest) (*AutokeyConfig, error)
	// Returns the effective Cloud KMS Autokey configuration for a given project.
	ShowEffectiveAutokeyConfig(context.Context, *ShowEffectiveAutokeyConfigRequest) (*ShowEffectiveAutokeyConfigResponse, error)
}

// UnimplementedAutokeyAdminServer can be embedded to have forward compatible implementations.
type UnimplementedAutokeyAdminServer struct {
}

func (*UnimplementedAutokeyAdminServer) UpdateAutokeyConfig(context.Context, *UpdateAutokeyConfigRequest) (*AutokeyConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAutokeyConfig not implemented")
}
func (*UnimplementedAutokeyAdminServer) GetAutokeyConfig(context.Context, *GetAutokeyConfigRequest) (*AutokeyConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAutokeyConfig not implemented")
}
func (*UnimplementedAutokeyAdminServer) ShowEffectiveAutokeyConfig(context.Context, *ShowEffectiveAutokeyConfigRequest) (*ShowEffectiveAutokeyConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowEffectiveAutokeyConfig not implemented")
}

func RegisterAutokeyAdminServer(s *grpc.Server, srv AutokeyAdminServer) {
	s.RegisterService(&_AutokeyAdmin_serviceDesc, srv)
}

func _AutokeyAdmin_UpdateAutokeyConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAutokeyConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutokeyAdminServer).UpdateAutokeyConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.kms.v1.AutokeyAdmin/UpdateAutokeyConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutokeyAdminServer).UpdateAutokeyConfig(ctx, req.(*UpdateAutokeyConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutokeyAdmin_GetAutokeyConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAutokeyConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutokeyAdminServer).GetAutokeyConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.kms.v1.AutokeyAdmin/GetAutokeyConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutokeyAdminServer).GetAutokeyConfig(ctx, req.(*GetAutokeyConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutokeyAdmin_ShowEffectiveAutokeyConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowEffectiveAutokeyConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutokeyAdminServer).ShowEffectiveAutokeyConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.kms.v1.AutokeyAdmin/ShowEffectiveAutokeyConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutokeyAdminServer).ShowEffectiveAutokeyConfig(ctx, req.(*ShowEffectiveAutokeyConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AutokeyAdmin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.kms.v1.AutokeyAdmin",
	HandlerType: (*AutokeyAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateAutokeyConfig",
			Handler:    _AutokeyAdmin_UpdateAutokeyConfig_Handler,
		},
		{
			MethodName: "GetAutokeyConfig",
			Handler:    _AutokeyAdmin_GetAutokeyConfig_Handler,
		},
		{
			MethodName: "ShowEffectiveAutokeyConfig",
			Handler:    _AutokeyAdmin_ShowEffectiveAutokeyConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/kms/v1/autokey_admin.proto",
}
