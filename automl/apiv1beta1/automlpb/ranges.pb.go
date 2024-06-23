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
// source: google/cloud/automl/v1beta1/ranges.proto

package automlpb

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

// A range between two double numbers.
type DoubleRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Start of the range, inclusive.
	Start float64 `protobuf:"fixed64,1,opt,name=start,proto3" json:"start,omitempty"`
	// End of the range, exclusive.
	End float64 `protobuf:"fixed64,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *DoubleRange) Reset() {
	*x = DoubleRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_ranges_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoubleRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoubleRange) ProtoMessage() {}

func (x *DoubleRange) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_ranges_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoubleRange.ProtoReflect.Descriptor instead.
func (*DoubleRange) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_ranges_proto_rawDescGZIP(), []int{0}
}

func (x *DoubleRange) GetStart() float64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *DoubleRange) GetEnd() float64 {
	if x != nil {
		return x.End
	}
	return 0
}

var File_google_cloud_automl_v1beta1_ranges_proto protoreflect.FileDescriptor

var file_google_cloud_automl_v1beta1_ranges_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x22, 0x35, 0x0a, 0x0b, 0x44, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x42, 0xa8,
	0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x42, 0x0b, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x61, 0x70,
	0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x70,
	0x62, 0x3b, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x70, 0x62, 0xca, 0x02, 0x1b, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x6c,
	0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x4c,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_cloud_automl_v1beta1_ranges_proto_rawDescOnce sync.Once
	file_google_cloud_automl_v1beta1_ranges_proto_rawDescData = file_google_cloud_automl_v1beta1_ranges_proto_rawDesc
)

func file_google_cloud_automl_v1beta1_ranges_proto_rawDescGZIP() []byte {
	file_google_cloud_automl_v1beta1_ranges_proto_rawDescOnce.Do(func() {
		file_google_cloud_automl_v1beta1_ranges_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_automl_v1beta1_ranges_proto_rawDescData)
	})
	return file_google_cloud_automl_v1beta1_ranges_proto_rawDescData
}

var file_google_cloud_automl_v1beta1_ranges_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_automl_v1beta1_ranges_proto_goTypes = []any{
	(*DoubleRange)(nil), // 0: google.cloud.automl.v1beta1.DoubleRange
}
var file_google_cloud_automl_v1beta1_ranges_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_automl_v1beta1_ranges_proto_init() }
func file_google_cloud_automl_v1beta1_ranges_proto_init() {
	if File_google_cloud_automl_v1beta1_ranges_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_automl_v1beta1_ranges_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DoubleRange); i {
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
			RawDescriptor: file_google_cloud_automl_v1beta1_ranges_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_automl_v1beta1_ranges_proto_goTypes,
		DependencyIndexes: file_google_cloud_automl_v1beta1_ranges_proto_depIdxs,
		MessageInfos:      file_google_cloud_automl_v1beta1_ranges_proto_msgTypes,
	}.Build()
	File_google_cloud_automl_v1beta1_ranges_proto = out.File
	file_google_cloud_automl_v1beta1_ranges_proto_rawDesc = nil
	file_google_cloud_automl_v1beta1_ranges_proto_goTypes = nil
	file_google_cloud_automl_v1beta1_ranges_proto_depIdxs = nil
}
