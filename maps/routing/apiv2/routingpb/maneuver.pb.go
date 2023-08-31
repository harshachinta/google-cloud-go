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
// source: google/maps/routing/v2/maneuver.proto

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

// A set of values that specify the navigation action to take for the current
// step (e.g., turn left, merge, straight, etc.).
type Maneuver int32

const (
	// Not used.
	Maneuver_MANEUVER_UNSPECIFIED Maneuver = 0
	// Turn slightly to the left.
	Maneuver_TURN_SLIGHT_LEFT Maneuver = 1
	// Turn sharply to the left.
	Maneuver_TURN_SHARP_LEFT Maneuver = 2
	// Make a left u-turn.
	Maneuver_UTURN_LEFT Maneuver = 3
	// Turn left.
	Maneuver_TURN_LEFT Maneuver = 4
	// Turn slightly to the right.
	Maneuver_TURN_SLIGHT_RIGHT Maneuver = 5
	// Turn sharply to the right.
	Maneuver_TURN_SHARP_RIGHT Maneuver = 6
	// Make a right u-turn.
	Maneuver_UTURN_RIGHT Maneuver = 7
	// Turn right.
	Maneuver_TURN_RIGHT Maneuver = 8
	// Go straight.
	Maneuver_STRAIGHT Maneuver = 9
	// Take the left ramp.
	Maneuver_RAMP_LEFT Maneuver = 10
	// Take the right ramp.
	Maneuver_RAMP_RIGHT Maneuver = 11
	// Merge into traffic.
	Maneuver_MERGE Maneuver = 12
	// Take the left fork.
	Maneuver_FORK_LEFT Maneuver = 13
	// Take the right fork.
	Maneuver_FORK_RIGHT Maneuver = 14
	// Take the ferry.
	Maneuver_FERRY Maneuver = 15
	// Take the train leading onto the ferry.
	Maneuver_FERRY_TRAIN Maneuver = 16
	// Turn left at the roundabout.
	Maneuver_ROUNDABOUT_LEFT Maneuver = 17
	// Turn right at the roundabout.
	Maneuver_ROUNDABOUT_RIGHT Maneuver = 18
	// Initial maneuver.
	Maneuver_DEPART Maneuver = 19
	// Used to indicate a street name change.
	Maneuver_NAME_CHANGE Maneuver = 20
)

// Enum value maps for Maneuver.
var (
	Maneuver_name = map[int32]string{
		0:  "MANEUVER_UNSPECIFIED",
		1:  "TURN_SLIGHT_LEFT",
		2:  "TURN_SHARP_LEFT",
		3:  "UTURN_LEFT",
		4:  "TURN_LEFT",
		5:  "TURN_SLIGHT_RIGHT",
		6:  "TURN_SHARP_RIGHT",
		7:  "UTURN_RIGHT",
		8:  "TURN_RIGHT",
		9:  "STRAIGHT",
		10: "RAMP_LEFT",
		11: "RAMP_RIGHT",
		12: "MERGE",
		13: "FORK_LEFT",
		14: "FORK_RIGHT",
		15: "FERRY",
		16: "FERRY_TRAIN",
		17: "ROUNDABOUT_LEFT",
		18: "ROUNDABOUT_RIGHT",
		19: "DEPART",
		20: "NAME_CHANGE",
	}
	Maneuver_value = map[string]int32{
		"MANEUVER_UNSPECIFIED": 0,
		"TURN_SLIGHT_LEFT":     1,
		"TURN_SHARP_LEFT":      2,
		"UTURN_LEFT":           3,
		"TURN_LEFT":            4,
		"TURN_SLIGHT_RIGHT":    5,
		"TURN_SHARP_RIGHT":     6,
		"UTURN_RIGHT":          7,
		"TURN_RIGHT":           8,
		"STRAIGHT":             9,
		"RAMP_LEFT":            10,
		"RAMP_RIGHT":           11,
		"MERGE":                12,
		"FORK_LEFT":            13,
		"FORK_RIGHT":           14,
		"FERRY":                15,
		"FERRY_TRAIN":          16,
		"ROUNDABOUT_LEFT":      17,
		"ROUNDABOUT_RIGHT":     18,
		"DEPART":               19,
		"NAME_CHANGE":          20,
	}
)

func (x Maneuver) Enum() *Maneuver {
	p := new(Maneuver)
	*p = x
	return p
}

func (x Maneuver) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Maneuver) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_routing_v2_maneuver_proto_enumTypes[0].Descriptor()
}

func (Maneuver) Type() protoreflect.EnumType {
	return &file_google_maps_routing_v2_maneuver_proto_enumTypes[0]
}

func (x Maneuver) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Maneuver.Descriptor instead.
func (Maneuver) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_routing_v2_maneuver_proto_rawDescGZIP(), []int{0}
}

var File_google_maps_routing_v2_maneuver_proto protoreflect.FileDescriptor

var file_google_maps_routing_v2_maneuver_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x61, 0x6e, 0x65, 0x75, 0x76, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2a,
	0xf7, 0x02, 0x0a, 0x08, 0x4d, 0x61, 0x6e, 0x65, 0x75, 0x76, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x14,
	0x4d, 0x41, 0x4e, 0x45, 0x55, 0x56, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x55, 0x52, 0x4e, 0x5f, 0x53,
	0x4c, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x4c, 0x45, 0x46, 0x54, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f,
	0x54, 0x55, 0x52, 0x4e, 0x5f, 0x53, 0x48, 0x41, 0x52, 0x50, 0x5f, 0x4c, 0x45, 0x46, 0x54, 0x10,
	0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x54, 0x55, 0x52, 0x4e, 0x5f, 0x4c, 0x45, 0x46, 0x54, 0x10,
	0x03, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x55, 0x52, 0x4e, 0x5f, 0x4c, 0x45, 0x46, 0x54, 0x10, 0x04,
	0x12, 0x15, 0x0a, 0x11, 0x54, 0x55, 0x52, 0x4e, 0x5f, 0x53, 0x4c, 0x49, 0x47, 0x48, 0x54, 0x5f,
	0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x55, 0x52, 0x4e, 0x5f,
	0x53, 0x48, 0x41, 0x52, 0x50, 0x5f, 0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x06, 0x12, 0x0f, 0x0a,
	0x0b, 0x55, 0x54, 0x55, 0x52, 0x4e, 0x5f, 0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x07, 0x12, 0x0e,
	0x0a, 0x0a, 0x54, 0x55, 0x52, 0x4e, 0x5f, 0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x08, 0x12, 0x0c,
	0x0a, 0x08, 0x53, 0x54, 0x52, 0x41, 0x49, 0x47, 0x48, 0x54, 0x10, 0x09, 0x12, 0x0d, 0x0a, 0x09,
	0x52, 0x41, 0x4d, 0x50, 0x5f, 0x4c, 0x45, 0x46, 0x54, 0x10, 0x0a, 0x12, 0x0e, 0x0a, 0x0a, 0x52,
	0x41, 0x4d, 0x50, 0x5f, 0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x0b, 0x12, 0x09, 0x0a, 0x05, 0x4d,
	0x45, 0x52, 0x47, 0x45, 0x10, 0x0c, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x4f, 0x52, 0x4b, 0x5f, 0x4c,
	0x45, 0x46, 0x54, 0x10, 0x0d, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x4f, 0x52, 0x4b, 0x5f, 0x52, 0x49,
	0x47, 0x48, 0x54, 0x10, 0x0e, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x45, 0x52, 0x52, 0x59, 0x10, 0x0f,
	0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x45, 0x52, 0x52, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x10,
	0x10, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x41, 0x42, 0x4f, 0x55, 0x54, 0x5f,
	0x4c, 0x45, 0x46, 0x54, 0x10, 0x11, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x41,
	0x42, 0x4f, 0x55, 0x54, 0x5f, 0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x12, 0x12, 0x0a, 0x0a, 0x06,
	0x44, 0x45, 0x50, 0x41, 0x52, 0x54, 0x10, 0x13, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x41, 0x4d, 0x45,
	0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x14, 0x42, 0xc2, 0x01, 0x0a, 0x1a, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x42, 0x0d, 0x4d, 0x61, 0x6e, 0x65, 0x75, 0x76,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6d,
	0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x3b, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x70, 0x62, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x05, 0x47, 0x4d, 0x52, 0x56, 0x32,
	0xaa, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x52,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x4d, 0x61, 0x70, 0x73, 0x5c, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5c,
	0x56, 0x32, 0xea, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x4d, 0x61, 0x70,
	0x73, 0x3a, 0x3a, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_routing_v2_maneuver_proto_rawDescOnce sync.Once
	file_google_maps_routing_v2_maneuver_proto_rawDescData = file_google_maps_routing_v2_maneuver_proto_rawDesc
)

func file_google_maps_routing_v2_maneuver_proto_rawDescGZIP() []byte {
	file_google_maps_routing_v2_maneuver_proto_rawDescOnce.Do(func() {
		file_google_maps_routing_v2_maneuver_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_routing_v2_maneuver_proto_rawDescData)
	})
	return file_google_maps_routing_v2_maneuver_proto_rawDescData
}

var file_google_maps_routing_v2_maneuver_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_maps_routing_v2_maneuver_proto_goTypes = []interface{}{
	(Maneuver)(0), // 0: google.maps.routing.v2.Maneuver
}
var file_google_maps_routing_v2_maneuver_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_maps_routing_v2_maneuver_proto_init() }
func file_google_maps_routing_v2_maneuver_proto_init() {
	if File_google_maps_routing_v2_maneuver_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_maps_routing_v2_maneuver_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_routing_v2_maneuver_proto_goTypes,
		DependencyIndexes: file_google_maps_routing_v2_maneuver_proto_depIdxs,
		EnumInfos:         file_google_maps_routing_v2_maneuver_proto_enumTypes,
	}.Build()
	File_google_maps_routing_v2_maneuver_proto = out.File
	file_google_maps_routing_v2_maneuver_proto_rawDesc = nil
	file_google_maps_routing_v2_maneuver_proto_goTypes = nil
	file_google_maps_routing_v2_maneuver_proto_depIdxs = nil
}
