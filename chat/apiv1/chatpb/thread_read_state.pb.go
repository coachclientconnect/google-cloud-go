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
// source: google/chat/v1/thread_read_state.proto

package chatpb

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

// A user's read state within a thread, used to identify read and unread
// messages.
type ThreadReadState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name of the thread read state.
	//
	// Format: `users/{user}/spaces/{space}/threads/{thread}/threadReadState`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The time when the user's thread read state was updated. Usually this
	// corresponds with the timestamp of the last read message in a thread.
	LastReadTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_read_time,json=lastReadTime,proto3" json:"last_read_time,omitempty"`
}

func (x *ThreadReadState) Reset() {
	*x = ThreadReadState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_chat_v1_thread_read_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThreadReadState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThreadReadState) ProtoMessage() {}

func (x *ThreadReadState) ProtoReflect() protoreflect.Message {
	mi := &file_google_chat_v1_thread_read_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThreadReadState.ProtoReflect.Descriptor instead.
func (*ThreadReadState) Descriptor() ([]byte, []int) {
	return file_google_chat_v1_thread_read_state_proto_rawDescGZIP(), []int{0}
}

func (x *ThreadReadState) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ThreadReadState) GetLastReadTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastReadTime
	}
	return nil
}

// Request message for GetThreadReadStateRequest API.
type GetThreadReadStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Resource name of the thread read state to retrieve.
	//
	// Only supports getting read state for the calling user.
	//
	// To refer to the calling user, set one of the following:
	//
	// - The `me` alias. For example,
	// `users/me/spaces/{space}/threads/{thread}/threadReadState`.
	//
	// - Their Workspace email address. For example,
	// `users/user@example.com/spaces/{space}/threads/{thread}/threadReadState`.
	//
	// - Their user id. For example,
	// `users/123456789/spaces/{space}/threads/{thread}/threadReadState`.
	//
	// Format: users/{user}/spaces/{space}/threads/{thread}/threadReadState
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetThreadReadStateRequest) Reset() {
	*x = GetThreadReadStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_chat_v1_thread_read_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetThreadReadStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetThreadReadStateRequest) ProtoMessage() {}

func (x *GetThreadReadStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_chat_v1_thread_read_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetThreadReadStateRequest.ProtoReflect.Descriptor instead.
func (*GetThreadReadStateRequest) Descriptor() ([]byte, []int) {
	return file_google_chat_v1_thread_read_state_proto_rawDescGZIP(), []int{1}
}

func (x *GetThreadReadStateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_google_chat_v1_thread_read_state_proto protoreflect.FileDescriptor

var file_google_chat_v1_thread_read_state_proto_rawDesc = []byte{
	0x0a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x01, 0x0a, 0x0f, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a,
	0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x3a,
	0x77, 0xea, 0x41, 0x74, 0x0a, 0x23, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3c, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x7d, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x7b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x7d, 0x2f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x7b,
	0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x7d, 0x2f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x52, 0x65,
	0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x32, 0x0f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x52,
	0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x5c, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x54,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x2b, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x25, 0x0a, 0x23, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0xad, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x14, 0x54,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f,
	0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x3b, 0x63, 0x68, 0x61,
	0x74, 0x70, 0x62, 0xa2, 0x02, 0x0b, 0x44, 0x59, 0x4e, 0x41, 0x50, 0x49, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x13, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x73, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x13, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x41, 0x70, 0x70, 0x73, 0x5c, 0x43, 0x68, 0x61, 0x74, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x16,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x70, 0x70, 0x73, 0x3a, 0x3a, 0x43, 0x68,
	0x61, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_chat_v1_thread_read_state_proto_rawDescOnce sync.Once
	file_google_chat_v1_thread_read_state_proto_rawDescData = file_google_chat_v1_thread_read_state_proto_rawDesc
)

func file_google_chat_v1_thread_read_state_proto_rawDescGZIP() []byte {
	file_google_chat_v1_thread_read_state_proto_rawDescOnce.Do(func() {
		file_google_chat_v1_thread_read_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_chat_v1_thread_read_state_proto_rawDescData)
	})
	return file_google_chat_v1_thread_read_state_proto_rawDescData
}

var file_google_chat_v1_thread_read_state_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_chat_v1_thread_read_state_proto_goTypes = []any{
	(*ThreadReadState)(nil),           // 0: google.chat.v1.ThreadReadState
	(*GetThreadReadStateRequest)(nil), // 1: google.chat.v1.GetThreadReadStateRequest
	(*timestamppb.Timestamp)(nil),     // 2: google.protobuf.Timestamp
}
var file_google_chat_v1_thread_read_state_proto_depIdxs = []int32{
	2, // 0: google.chat.v1.ThreadReadState.last_read_time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_chat_v1_thread_read_state_proto_init() }
func file_google_chat_v1_thread_read_state_proto_init() {
	if File_google_chat_v1_thread_read_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_chat_v1_thread_read_state_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ThreadReadState); i {
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
		file_google_chat_v1_thread_read_state_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetThreadReadStateRequest); i {
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
			RawDescriptor: file_google_chat_v1_thread_read_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_chat_v1_thread_read_state_proto_goTypes,
		DependencyIndexes: file_google_chat_v1_thread_read_state_proto_depIdxs,
		MessageInfos:      file_google_chat_v1_thread_read_state_proto_msgTypes,
	}.Build()
	File_google_chat_v1_thread_read_state_proto = out.File
	file_google_chat_v1_thread_read_state_proto_rawDesc = nil
	file_google_chat_v1_thread_read_state_proto_goTypes = nil
	file_google_chat_v1_thread_read_state_proto_depIdxs = nil
}
