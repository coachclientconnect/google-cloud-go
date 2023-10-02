// Copyright 2021 Google LLC
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
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: google/cloud/retail/v2alpha/merchant_center_account_link_service.proto

package retailpb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request for
// [MerchantCenterAccountLinkService.ListMerchantCenterAccountLinks][google.cloud.retail.v2alpha.MerchantCenterAccountLinkService.ListMerchantCenterAccountLinks]
// method.
type ListMerchantCenterAccountLinksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The parent Catalog of the resource.
	// It must match this format:
	// projects/{PROJECT_NUMBER}/locations/global/catalogs/{CATALOG_ID}
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
}

func (x *ListMerchantCenterAccountLinksRequest) Reset() {
	*x = ListMerchantCenterAccountLinksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMerchantCenterAccountLinksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMerchantCenterAccountLinksRequest) ProtoMessage() {}

func (x *ListMerchantCenterAccountLinksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMerchantCenterAccountLinksRequest.ProtoReflect.Descriptor instead.
func (*ListMerchantCenterAccountLinksRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListMerchantCenterAccountLinksRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

// Response for
// [MerchantCenterAccountLinkService.ListMerchantCenterAccountLinks][google.cloud.retail.v2alpha.MerchantCenterAccountLinkService.ListMerchantCenterAccountLinks]
// method.
type ListMerchantCenterAccountLinksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The links.
	MerchantCenterAccountLinks []*MerchantCenterAccountLink `protobuf:"bytes,1,rep,name=merchant_center_account_links,json=merchantCenterAccountLinks,proto3" json:"merchant_center_account_links,omitempty"`
}

func (x *ListMerchantCenterAccountLinksResponse) Reset() {
	*x = ListMerchantCenterAccountLinksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMerchantCenterAccountLinksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMerchantCenterAccountLinksResponse) ProtoMessage() {}

func (x *ListMerchantCenterAccountLinksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMerchantCenterAccountLinksResponse.ProtoReflect.Descriptor instead.
func (*ListMerchantCenterAccountLinksResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListMerchantCenterAccountLinksResponse) GetMerchantCenterAccountLinks() []*MerchantCenterAccountLink {
	if x != nil {
		return x.MerchantCenterAccountLinks
	}
	return nil
}

// Request for
// [MerchantCenterAccountLinkService.CreateMerchantCenterAccountLink][google.cloud.retail.v2alpha.MerchantCenterAccountLinkService.CreateMerchantCenterAccountLink]
// method.
type CreateMerchantCenterAccountLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The branch resource where this MerchantCenterAccountLink will be
	// created. Format:
	// projects/{PROJECT_NUMBER}/locations/global/catalogs/{CATALOG_ID}}
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The
	// [MerchantCenterAccountLink][google.cloud.retail.v2alpha.MerchantCenterAccountLink]
	// to create.
	//
	// If the caller does not have permission to create the
	// [MerchantCenterAccountLink][google.cloud.retail.v2alpha.MerchantCenterAccountLink],
	// regardless of whether or not it exists, a PERMISSION_DENIED error is
	// returned.
	MerchantCenterAccountLink *MerchantCenterAccountLink `protobuf:"bytes,2,opt,name=merchant_center_account_link,json=merchantCenterAccountLink,proto3" json:"merchant_center_account_link,omitempty"`
}

func (x *CreateMerchantCenterAccountLinkRequest) Reset() {
	*x = CreateMerchantCenterAccountLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMerchantCenterAccountLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMerchantCenterAccountLinkRequest) ProtoMessage() {}

func (x *CreateMerchantCenterAccountLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMerchantCenterAccountLinkRequest.ProtoReflect.Descriptor instead.
func (*CreateMerchantCenterAccountLinkRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateMerchantCenterAccountLinkRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateMerchantCenterAccountLinkRequest) GetMerchantCenterAccountLink() *MerchantCenterAccountLink {
	if x != nil {
		return x.MerchantCenterAccountLink
	}
	return nil
}

// Request for
// [MerchantCenterAccountLinkService.DeleteMerchantCenterAccountLink][google.cloud.retail.v2alpha.MerchantCenterAccountLinkService.DeleteMerchantCenterAccountLink]
// method.
type DeleteMerchantCenterAccountLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Full resource name. Format:
	// projects/{project_number}/locations/{location_id}/catalogs/{catalog_id}/merchantCenterAccountLinks/{merchant_center_account_link_id}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteMerchantCenterAccountLinkRequest) Reset() {
	*x = DeleteMerchantCenterAccountLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMerchantCenterAccountLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMerchantCenterAccountLinkRequest) ProtoMessage() {}

func (x *DeleteMerchantCenterAccountLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMerchantCenterAccountLinkRequest.ProtoReflect.Descriptor instead.
func (*DeleteMerchantCenterAccountLinkRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteMerchantCenterAccountLinkRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto protoreflect.FileDescriptor

var file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_rawDesc = []byte{
	0x0a, 0x46, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6c, 0x69,
	0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x25, 0x4c, 0x69,
	0x73, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x25, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x1f, 0x0a, 0x1d, 0x72, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x22, 0xa3, 0x01, 0x0a, 0x26, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x79, 0x0a,
	0x1d, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x1a, 0x6d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x22, 0xe5, 0x01, 0x0a, 0x26, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x25, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x1f, 0x0a, 0x1d, 0x72, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x12, 0x7c, 0x0a, 0x1c, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6c, 0x69,
	0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x19, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b,
	0x22, 0x75, 0x0a, 0x26, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c,
	0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x37, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x31,
	0x0a, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xfe, 0x07, 0x0a, 0x20, 0x4d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8a, 0x02, 0x0a,
	0x1e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12,
	0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5f, 0xda, 0x41, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x50, 0x12, 0x4e, 0x2f, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x2a, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x6d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0xa1, 0x03, 0x0a, 0x1f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x43, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67,
	0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x99, 0x02, 0xca, 0x41, 0x7c, 0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x43,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xda, 0x41, 0x23, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2c, 0x6d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x6e, 0x3a,
	0x1c, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x4e, 0x2f,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x2f,
	0x2a, 0x7d, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0xdd, 0x01,
	0x0a, 0x1f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x12, 0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x5d,
	0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x50, 0x2a, 0x4e, 0x2f,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x2a, 0x2f,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x2f, 0x2a, 0x7d, 0x1a, 0x49, 0xca,
	0x41, 0x15, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0xe9, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x25, 0x4d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x70, 0x62, 0x3b, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x70, 0x62, 0xa2, 0x02,
	0x06, 0x52, 0x45, 0x54, 0x41, 0x49, 0x4c, 0xaa, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x56, 0x32,
	0x41, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5c, 0x56, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0xea, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x3a, 0x3a, 0x56, 0x32, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_rawDescOnce sync.Once
	file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_rawDescData = file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_rawDesc
)

func file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_rawDescGZIP() []byte {
	file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_rawDescData)
	})
	return file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_rawDescData
}

var file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_goTypes = []interface{}{
	(*ListMerchantCenterAccountLinksRequest)(nil),  // 0: google.cloud.retail.v2alpha.ListMerchantCenterAccountLinksRequest
	(*ListMerchantCenterAccountLinksResponse)(nil), // 1: google.cloud.retail.v2alpha.ListMerchantCenterAccountLinksResponse
	(*CreateMerchantCenterAccountLinkRequest)(nil), // 2: google.cloud.retail.v2alpha.CreateMerchantCenterAccountLinkRequest
	(*DeleteMerchantCenterAccountLinkRequest)(nil), // 3: google.cloud.retail.v2alpha.DeleteMerchantCenterAccountLinkRequest
	(*MerchantCenterAccountLink)(nil),              // 4: google.cloud.retail.v2alpha.MerchantCenterAccountLink
	(*longrunningpb.Operation)(nil),                // 5: google.longrunning.Operation
	(*emptypb.Empty)(nil),                          // 6: google.protobuf.Empty
}
var file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_depIdxs = []int32{
	4, // 0: google.cloud.retail.v2alpha.ListMerchantCenterAccountLinksResponse.merchant_center_account_links:type_name -> google.cloud.retail.v2alpha.MerchantCenterAccountLink
	4, // 1: google.cloud.retail.v2alpha.CreateMerchantCenterAccountLinkRequest.merchant_center_account_link:type_name -> google.cloud.retail.v2alpha.MerchantCenterAccountLink
	0, // 2: google.cloud.retail.v2alpha.MerchantCenterAccountLinkService.ListMerchantCenterAccountLinks:input_type -> google.cloud.retail.v2alpha.ListMerchantCenterAccountLinksRequest
	2, // 3: google.cloud.retail.v2alpha.MerchantCenterAccountLinkService.CreateMerchantCenterAccountLink:input_type -> google.cloud.retail.v2alpha.CreateMerchantCenterAccountLinkRequest
	3, // 4: google.cloud.retail.v2alpha.MerchantCenterAccountLinkService.DeleteMerchantCenterAccountLink:input_type -> google.cloud.retail.v2alpha.DeleteMerchantCenterAccountLinkRequest
	1, // 5: google.cloud.retail.v2alpha.MerchantCenterAccountLinkService.ListMerchantCenterAccountLinks:output_type -> google.cloud.retail.v2alpha.ListMerchantCenterAccountLinksResponse
	5, // 6: google.cloud.retail.v2alpha.MerchantCenterAccountLinkService.CreateMerchantCenterAccountLink:output_type -> google.longrunning.Operation
	6, // 7: google.cloud.retail.v2alpha.MerchantCenterAccountLinkService.DeleteMerchantCenterAccountLink:output_type -> google.protobuf.Empty
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_init() }
func file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_init() {
	if File_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto != nil {
		return
	}
	file_google_cloud_retail_v2alpha_merchant_center_account_link_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMerchantCenterAccountLinksRequest); i {
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
		file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMerchantCenterAccountLinksResponse); i {
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
		file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMerchantCenterAccountLinkRequest); i {
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
		file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMerchantCenterAccountLinkRequest); i {
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
			RawDescriptor: file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_depIdxs,
		MessageInfos:      file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_msgTypes,
	}.Build()
	File_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto = out.File
	file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_rawDesc = nil
	file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_goTypes = nil
	file_google_cloud_retail_v2alpha_merchant_center_account_link_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MerchantCenterAccountLinkServiceClient is the client API for MerchantCenterAccountLinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MerchantCenterAccountLinkServiceClient interface {
	// Lists all
	// [MerchantCenterAccountLink][google.cloud.retail.v2alpha.MerchantCenterAccountLink]s
	// under the specified parent [Catalog][google.cloud.retail.v2alpha.Catalog].
	ListMerchantCenterAccountLinks(ctx context.Context, in *ListMerchantCenterAccountLinksRequest, opts ...grpc.CallOption) (*ListMerchantCenterAccountLinksResponse, error)
	// Creates a
	// [MerchantCenterAccountLink][google.cloud.retail.v2alpha.MerchantCenterAccountLink].
	CreateMerchantCenterAccountLink(ctx context.Context, in *CreateMerchantCenterAccountLinkRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Deletes a
	// [MerchantCenterAccountLink][google.cloud.retail.v2alpha.MerchantCenterAccountLink].
	// If the
	// [MerchantCenterAccountLink][google.cloud.retail.v2alpha.MerchantCenterAccountLink]
	// to delete does not exist, a NOT_FOUND error is returned.
	DeleteMerchantCenterAccountLink(ctx context.Context, in *DeleteMerchantCenterAccountLinkRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type merchantCenterAccountLinkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMerchantCenterAccountLinkServiceClient(cc grpc.ClientConnInterface) MerchantCenterAccountLinkServiceClient {
	return &merchantCenterAccountLinkServiceClient{cc}
}

func (c *merchantCenterAccountLinkServiceClient) ListMerchantCenterAccountLinks(ctx context.Context, in *ListMerchantCenterAccountLinksRequest, opts ...grpc.CallOption) (*ListMerchantCenterAccountLinksResponse, error) {
	out := new(ListMerchantCenterAccountLinksResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.retail.v2alpha.MerchantCenterAccountLinkService/ListMerchantCenterAccountLinks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantCenterAccountLinkServiceClient) CreateMerchantCenterAccountLink(ctx context.Context, in *CreateMerchantCenterAccountLinkRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.retail.v2alpha.MerchantCenterAccountLinkService/CreateMerchantCenterAccountLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantCenterAccountLinkServiceClient) DeleteMerchantCenterAccountLink(ctx context.Context, in *DeleteMerchantCenterAccountLinkRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/google.cloud.retail.v2alpha.MerchantCenterAccountLinkService/DeleteMerchantCenterAccountLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MerchantCenterAccountLinkServiceServer is the server API for MerchantCenterAccountLinkService service.
type MerchantCenterAccountLinkServiceServer interface {
	// Lists all
	// [MerchantCenterAccountLink][google.cloud.retail.v2alpha.MerchantCenterAccountLink]s
	// under the specified parent [Catalog][google.cloud.retail.v2alpha.Catalog].
	ListMerchantCenterAccountLinks(context.Context, *ListMerchantCenterAccountLinksRequest) (*ListMerchantCenterAccountLinksResponse, error)
	// Creates a
	// [MerchantCenterAccountLink][google.cloud.retail.v2alpha.MerchantCenterAccountLink].
	CreateMerchantCenterAccountLink(context.Context, *CreateMerchantCenterAccountLinkRequest) (*longrunningpb.Operation, error)
	// Deletes a
	// [MerchantCenterAccountLink][google.cloud.retail.v2alpha.MerchantCenterAccountLink].
	// If the
	// [MerchantCenterAccountLink][google.cloud.retail.v2alpha.MerchantCenterAccountLink]
	// to delete does not exist, a NOT_FOUND error is returned.
	DeleteMerchantCenterAccountLink(context.Context, *DeleteMerchantCenterAccountLinkRequest) (*emptypb.Empty, error)
}

// UnimplementedMerchantCenterAccountLinkServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMerchantCenterAccountLinkServiceServer struct {
}

func (*UnimplementedMerchantCenterAccountLinkServiceServer) ListMerchantCenterAccountLinks(context.Context, *ListMerchantCenterAccountLinksRequest) (*ListMerchantCenterAccountLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMerchantCenterAccountLinks not implemented")
}
func (*UnimplementedMerchantCenterAccountLinkServiceServer) CreateMerchantCenterAccountLink(context.Context, *CreateMerchantCenterAccountLinkRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMerchantCenterAccountLink not implemented")
}
func (*UnimplementedMerchantCenterAccountLinkServiceServer) DeleteMerchantCenterAccountLink(context.Context, *DeleteMerchantCenterAccountLinkRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMerchantCenterAccountLink not implemented")
}

func RegisterMerchantCenterAccountLinkServiceServer(s *grpc.Server, srv MerchantCenterAccountLinkServiceServer) {
	s.RegisterService(&_MerchantCenterAccountLinkService_serviceDesc, srv)
}

func _MerchantCenterAccountLinkService_ListMerchantCenterAccountLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMerchantCenterAccountLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantCenterAccountLinkServiceServer).ListMerchantCenterAccountLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.retail.v2alpha.MerchantCenterAccountLinkService/ListMerchantCenterAccountLinks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantCenterAccountLinkServiceServer).ListMerchantCenterAccountLinks(ctx, req.(*ListMerchantCenterAccountLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantCenterAccountLinkService_CreateMerchantCenterAccountLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMerchantCenterAccountLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantCenterAccountLinkServiceServer).CreateMerchantCenterAccountLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.retail.v2alpha.MerchantCenterAccountLinkService/CreateMerchantCenterAccountLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantCenterAccountLinkServiceServer).CreateMerchantCenterAccountLink(ctx, req.(*CreateMerchantCenterAccountLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantCenterAccountLinkService_DeleteMerchantCenterAccountLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMerchantCenterAccountLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantCenterAccountLinkServiceServer).DeleteMerchantCenterAccountLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.retail.v2alpha.MerchantCenterAccountLinkService/DeleteMerchantCenterAccountLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantCenterAccountLinkServiceServer).DeleteMerchantCenterAccountLink(ctx, req.(*DeleteMerchantCenterAccountLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MerchantCenterAccountLinkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.retail.v2alpha.MerchantCenterAccountLinkService",
	HandlerType: (*MerchantCenterAccountLinkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMerchantCenterAccountLinks",
			Handler:    _MerchantCenterAccountLinkService_ListMerchantCenterAccountLinks_Handler,
		},
		{
			MethodName: "CreateMerchantCenterAccountLink",
			Handler:    _MerchantCenterAccountLinkService_CreateMerchantCenterAccountLink_Handler,
		},
		{
			MethodName: "DeleteMerchantCenterAccountLink",
			Handler:    _MerchantCenterAccountLinkService_DeleteMerchantCenterAccountLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/retail/v2alpha/merchant_center_account_link_service.proto",
}
