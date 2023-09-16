// Copyright 2023 Google LLC
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
// source: google/maps/routing/v2/traffic_model.proto

package routingpb

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

// Specifies the assumptions to use when calculating time in traffic. This
// setting affects the value returned in the `duration` field in the
// response, which contains the predicted time in traffic based on historical
// averages.
type TrafficModel int32

const (
	// Unused. If specified, will default to `BEST_GUESS`.
	TrafficModel_TRAFFIC_MODEL_UNSPECIFIED TrafficModel = 0
	// Indicates that the returned `duration` should be the best
	// estimate of travel time given what is known about both historical traffic
	// conditions and live traffic. Live traffic becomes more important the closer
	// the `departure_time` is to now.
	TrafficModel_BEST_GUESS TrafficModel = 1
	// Indicates that the returned duration should be longer than the
	// actual travel time on most days, though occasional days with particularly
	// bad traffic conditions may exceed this value.
	TrafficModel_PESSIMISTIC TrafficModel = 2
	// Indicates that the returned duration should be shorter than the actual
	// travel time on most days, though occasional days with particularly good
	// traffic conditions may be faster than this value.
	TrafficModel_OPTIMISTIC TrafficModel = 3
)

// Enum value maps for TrafficModel.
var (
	TrafficModel_name = map[int32]string{
		0: "TRAFFIC_MODEL_UNSPECIFIED",
		1: "BEST_GUESS",
		2: "PESSIMISTIC",
		3: "OPTIMISTIC",
	}
	TrafficModel_value = map[string]int32{
		"TRAFFIC_MODEL_UNSPECIFIED": 0,
		"BEST_GUESS":                1,
		"PESSIMISTIC":               2,
		"OPTIMISTIC":                3,
	}
)

func (x TrafficModel) Enum() *TrafficModel {
	p := new(TrafficModel)
	*p = x
	return p
}

func (x TrafficModel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TrafficModel) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_routing_v2_traffic_model_proto_enumTypes[0].Descriptor()
}

func (TrafficModel) Type() protoreflect.EnumType {
	return &file_google_maps_routing_v2_traffic_model_proto_enumTypes[0]
}

func (x TrafficModel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TrafficModel.Descriptor instead.
func (TrafficModel) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_routing_v2_traffic_model_proto_rawDescGZIP(), []int{0}
}

var File_google_maps_routing_v2_traffic_model_proto protoreflect.FileDescriptor

var file_google_maps_routing_v2_traffic_model_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x32, 0x2a, 0x5e, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x19, 0x54, 0x52, 0x41, 0x46, 0x46, 0x49, 0x43, 0x5f,
	0x4d, 0x4f, 0x44, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x42, 0x45, 0x53, 0x54, 0x5f, 0x47, 0x55, 0x45, 0x53,
	0x53, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x45, 0x53, 0x53, 0x49, 0x4d, 0x49, 0x53, 0x54,
	0x49, 0x43, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x50, 0x54, 0x49, 0x4d, 0x49, 0x53, 0x54,
	0x49, 0x43, 0x10, 0x03, 0x42, 0xc6, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x32, 0x42, 0x11, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x61,
	0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x3b, 0x72, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x70, 0x62, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x05, 0x47, 0x4d, 0x52, 0x56, 0x32, 0xaa,
	0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x52, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x4d, 0x61, 0x70, 0x73, 0x5c, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5c, 0x56,
	0x32, 0xea, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x4d, 0x61, 0x70, 0x73,
	0x3a, 0x3a, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_routing_v2_traffic_model_proto_rawDescOnce sync.Once
	file_google_maps_routing_v2_traffic_model_proto_rawDescData = file_google_maps_routing_v2_traffic_model_proto_rawDesc
)

func file_google_maps_routing_v2_traffic_model_proto_rawDescGZIP() []byte {
	file_google_maps_routing_v2_traffic_model_proto_rawDescOnce.Do(func() {
		file_google_maps_routing_v2_traffic_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_routing_v2_traffic_model_proto_rawDescData)
	})
	return file_google_maps_routing_v2_traffic_model_proto_rawDescData
}

var file_google_maps_routing_v2_traffic_model_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_maps_routing_v2_traffic_model_proto_goTypes = []interface{}{
	(TrafficModel)(0), // 0: google.maps.routing.v2.TrafficModel
}
var file_google_maps_routing_v2_traffic_model_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_maps_routing_v2_traffic_model_proto_init() }
func file_google_maps_routing_v2_traffic_model_proto_init() {
	if File_google_maps_routing_v2_traffic_model_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_maps_routing_v2_traffic_model_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_routing_v2_traffic_model_proto_goTypes,
		DependencyIndexes: file_google_maps_routing_v2_traffic_model_proto_depIdxs,
		EnumInfos:         file_google_maps_routing_v2_traffic_model_proto_enumTypes,
	}.Build()
	File_google_maps_routing_v2_traffic_model_proto = out.File
	file_google_maps_routing_v2_traffic_model_proto_rawDesc = nil
	file_google_maps_routing_v2_traffic_model_proto_goTypes = nil
	file_google_maps_routing_v2_traffic_model_proto_depIdxs = nil
}
