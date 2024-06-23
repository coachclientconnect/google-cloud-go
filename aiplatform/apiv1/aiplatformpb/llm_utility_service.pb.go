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
// source: google/cloud/aiplatform/v1/llm_utility_service.proto

package aiplatformpb

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
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for ComputeTokens RPC call.
type ComputeTokensRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the Endpoint requested to get lists of tokens and
	// token ids.
	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// Required. The instances that are the input to token computing API call.
	// Schema is identical to the prediction schema of the text model, even for
	// the non-text models, like chat models, or Codey models.
	Instances []*structpb.Value `protobuf:"bytes,2,rep,name=instances,proto3" json:"instances,omitempty"`
}

func (x *ComputeTokensRequest) Reset() {
	*x = ComputeTokensRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_llm_utility_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeTokensRequest) ProtoMessage() {}

func (x *ComputeTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_llm_utility_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeTokensRequest.ProtoReflect.Descriptor instead.
func (*ComputeTokensRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_llm_utility_service_proto_rawDescGZIP(), []int{0}
}

func (x *ComputeTokensRequest) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *ComputeTokensRequest) GetInstances() []*structpb.Value {
	if x != nil {
		return x.Instances
	}
	return nil
}

// Tokens info with a list of tokens and the corresponding list of token ids.
type TokensInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of tokens from the input.
	Tokens [][]byte `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
	// A list of token ids from the input.
	TokenIds []int64 `protobuf:"varint,2,rep,packed,name=token_ids,json=tokenIds,proto3" json:"token_ids,omitempty"`
}

func (x *TokensInfo) Reset() {
	*x = TokensInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_llm_utility_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokensInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokensInfo) ProtoMessage() {}

func (x *TokensInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_llm_utility_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokensInfo.ProtoReflect.Descriptor instead.
func (*TokensInfo) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_llm_utility_service_proto_rawDescGZIP(), []int{1}
}

func (x *TokensInfo) GetTokens() [][]byte {
	if x != nil {
		return x.Tokens
	}
	return nil
}

func (x *TokensInfo) GetTokenIds() []int64 {
	if x != nil {
		return x.TokenIds
	}
	return nil
}

// Response message for ComputeTokens RPC call.
type ComputeTokensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Lists of tokens info from the input. A ComputeTokensRequest could have
	// multiple instances with a prompt in each instance. We also need to return
	// lists of tokens info for the request with multiple instances.
	TokensInfo []*TokensInfo `protobuf:"bytes,1,rep,name=tokens_info,json=tokensInfo,proto3" json:"tokens_info,omitempty"`
}

func (x *ComputeTokensResponse) Reset() {
	*x = ComputeTokensResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_llm_utility_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeTokensResponse) ProtoMessage() {}

func (x *ComputeTokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_llm_utility_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeTokensResponse.ProtoReflect.Descriptor instead.
func (*ComputeTokensResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_llm_utility_service_proto_rawDescGZIP(), []int{2}
}

func (x *ComputeTokensResponse) GetTokensInfo() []*TokensInfo {
	if x != nil {
		return x.TokensInfo
	}
	return nil
}

var File_google_cloud_aiplatform_v1_llm_utility_service_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1_llm_utility_service_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6c, 0x6d,
	0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x46, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x2a, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x24, 0x0a, 0x22, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x09, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x73, 0x22, 0x41, 0x0a, 0x0a, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x73, 0x22, 0x60, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x47, 0x0a, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xac, 0x05, 0x0a, 0x11, 0x4c, 0x6c,
	0x6d, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x9d, 0x02, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12,
	0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0xac, 0x01, 0xda, 0x41, 0x12, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2c, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x90, 0x01, 0x3a,
	0x01, 0x2a, 0x5a, 0x4c, 0x3a, 0x01, 0x2a, 0x22, 0x47, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x22, 0x3d, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x3d,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x2f, 0x2a, 0x7d, 0x3a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12,
	0xa7, 0x02, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x12, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xb0, 0x01, 0xda, 0x41, 0x12, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x94, 0x01, 0x3a, 0x01, 0x2a, 0x5a, 0x4e, 0x3a, 0x01, 0x2a, 0x22, 0x49, 0x2f,
	0x76, 0x31, 0x2f, 0x7b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x2a, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x73, 0x2f, 0x2a,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0x3f, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x63, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x1a, 0x4d, 0xca, 0x41, 0x19, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0xd4, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x16, 0x4c, 0x6c, 0x6d,
	0x55, 0x74, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0x3b, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x70, 0x62, 0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0xea,
	0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1_llm_utility_service_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1_llm_utility_service_proto_rawDescData = file_google_cloud_aiplatform_v1_llm_utility_service_proto_rawDesc
)

func file_google_cloud_aiplatform_v1_llm_utility_service_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1_llm_utility_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1_llm_utility_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1_llm_utility_service_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1_llm_utility_service_proto_rawDescData
}

var file_google_cloud_aiplatform_v1_llm_utility_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_aiplatform_v1_llm_utility_service_proto_goTypes = []any{
	(*ComputeTokensRequest)(nil),  // 0: google.cloud.aiplatform.v1.ComputeTokensRequest
	(*TokensInfo)(nil),            // 1: google.cloud.aiplatform.v1.TokensInfo
	(*ComputeTokensResponse)(nil), // 2: google.cloud.aiplatform.v1.ComputeTokensResponse
	(*structpb.Value)(nil),        // 3: google.protobuf.Value
	(*CountTokensRequest)(nil),    // 4: google.cloud.aiplatform.v1.CountTokensRequest
	(*CountTokensResponse)(nil),   // 5: google.cloud.aiplatform.v1.CountTokensResponse
}
var file_google_cloud_aiplatform_v1_llm_utility_service_proto_depIdxs = []int32{
	3, // 0: google.cloud.aiplatform.v1.ComputeTokensRequest.instances:type_name -> google.protobuf.Value
	1, // 1: google.cloud.aiplatform.v1.ComputeTokensResponse.tokens_info:type_name -> google.cloud.aiplatform.v1.TokensInfo
	4, // 2: google.cloud.aiplatform.v1.LlmUtilityService.CountTokens:input_type -> google.cloud.aiplatform.v1.CountTokensRequest
	0, // 3: google.cloud.aiplatform.v1.LlmUtilityService.ComputeTokens:input_type -> google.cloud.aiplatform.v1.ComputeTokensRequest
	5, // 4: google.cloud.aiplatform.v1.LlmUtilityService.CountTokens:output_type -> google.cloud.aiplatform.v1.CountTokensResponse
	2, // 5: google.cloud.aiplatform.v1.LlmUtilityService.ComputeTokens:output_type -> google.cloud.aiplatform.v1.ComputeTokensResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1_llm_utility_service_proto_init() }
func file_google_cloud_aiplatform_v1_llm_utility_service_proto_init() {
	if File_google_cloud_aiplatform_v1_llm_utility_service_proto != nil {
		return
	}
	file_google_cloud_aiplatform_v1_prediction_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1_llm_utility_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ComputeTokensRequest); i {
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
		file_google_cloud_aiplatform_v1_llm_utility_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TokensInfo); i {
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
		file_google_cloud_aiplatform_v1_llm_utility_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ComputeTokensResponse); i {
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
			RawDescriptor: file_google_cloud_aiplatform_v1_llm_utility_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_aiplatform_v1_llm_utility_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1_llm_utility_service_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1_llm_utility_service_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1_llm_utility_service_proto = out.File
	file_google_cloud_aiplatform_v1_llm_utility_service_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1_llm_utility_service_proto_goTypes = nil
	file_google_cloud_aiplatform_v1_llm_utility_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LlmUtilityServiceClient is the client API for LlmUtilityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LlmUtilityServiceClient interface {
	// Perform a token counting.
	CountTokens(ctx context.Context, in *CountTokensRequest, opts ...grpc.CallOption) (*CountTokensResponse, error)
	// Return a list of tokens based on the input text.
	ComputeTokens(ctx context.Context, in *ComputeTokensRequest, opts ...grpc.CallOption) (*ComputeTokensResponse, error)
}

type llmUtilityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLlmUtilityServiceClient(cc grpc.ClientConnInterface) LlmUtilityServiceClient {
	return &llmUtilityServiceClient{cc}
}

func (c *llmUtilityServiceClient) CountTokens(ctx context.Context, in *CountTokensRequest, opts ...grpc.CallOption) (*CountTokensResponse, error) {
	out := new(CountTokensResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.aiplatform.v1.LlmUtilityService/CountTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *llmUtilityServiceClient) ComputeTokens(ctx context.Context, in *ComputeTokensRequest, opts ...grpc.CallOption) (*ComputeTokensResponse, error) {
	out := new(ComputeTokensResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.aiplatform.v1.LlmUtilityService/ComputeTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LlmUtilityServiceServer is the server API for LlmUtilityService service.
type LlmUtilityServiceServer interface {
	// Perform a token counting.
	CountTokens(context.Context, *CountTokensRequest) (*CountTokensResponse, error)
	// Return a list of tokens based on the input text.
	ComputeTokens(context.Context, *ComputeTokensRequest) (*ComputeTokensResponse, error)
}

// UnimplementedLlmUtilityServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLlmUtilityServiceServer struct {
}

func (*UnimplementedLlmUtilityServiceServer) CountTokens(context.Context, *CountTokensRequest) (*CountTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountTokens not implemented")
}
func (*UnimplementedLlmUtilityServiceServer) ComputeTokens(context.Context, *ComputeTokensRequest) (*ComputeTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComputeTokens not implemented")
}

func RegisterLlmUtilityServiceServer(s *grpc.Server, srv LlmUtilityServiceServer) {
	s.RegisterService(&_LlmUtilityService_serviceDesc, srv)
}

func _LlmUtilityService_CountTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LlmUtilityServiceServer).CountTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.aiplatform.v1.LlmUtilityService/CountTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LlmUtilityServiceServer).CountTokens(ctx, req.(*CountTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LlmUtilityService_ComputeTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComputeTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LlmUtilityServiceServer).ComputeTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.aiplatform.v1.LlmUtilityService/ComputeTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LlmUtilityServiceServer).ComputeTokens(ctx, req.(*ComputeTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LlmUtilityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.aiplatform.v1.LlmUtilityService",
	HandlerType: (*LlmUtilityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CountTokens",
			Handler:    _LlmUtilityService_CountTokens_Handler,
		},
		{
			MethodName: "ComputeTokens",
			Handler:    _LlmUtilityService_ComputeTokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/aiplatform/v1/llm_utility_service.proto",
}