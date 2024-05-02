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
// source: google/maps/places/v1/fuel_options.proto

package placespb

import (
	reflect "reflect"
	sync "sync"

	money "google.golang.org/genproto/googleapis/type/money"
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

// Types of fuel.
type FuelOptions_FuelPrice_FuelType int32

const (
	// Unspecified fuel type.
	FuelOptions_FuelPrice_FUEL_TYPE_UNSPECIFIED FuelOptions_FuelPrice_FuelType = 0
	// Diesel fuel.
	FuelOptions_FuelPrice_DIESEL FuelOptions_FuelPrice_FuelType = 1
	// Regular unleaded.
	FuelOptions_FuelPrice_REGULAR_UNLEADED FuelOptions_FuelPrice_FuelType = 2
	// Midgrade.
	FuelOptions_FuelPrice_MIDGRADE FuelOptions_FuelPrice_FuelType = 3
	// Premium.
	FuelOptions_FuelPrice_PREMIUM FuelOptions_FuelPrice_FuelType = 4
	// SP 91.
	FuelOptions_FuelPrice_SP91 FuelOptions_FuelPrice_FuelType = 5
	// SP 91 E10.
	FuelOptions_FuelPrice_SP91_E10 FuelOptions_FuelPrice_FuelType = 6
	// SP 92.
	FuelOptions_FuelPrice_SP92 FuelOptions_FuelPrice_FuelType = 7
	// SP 95.
	FuelOptions_FuelPrice_SP95 FuelOptions_FuelPrice_FuelType = 8
	// SP95 E10.
	FuelOptions_FuelPrice_SP95_E10 FuelOptions_FuelPrice_FuelType = 9
	// SP 98.
	FuelOptions_FuelPrice_SP98 FuelOptions_FuelPrice_FuelType = 10
	// SP 99.
	FuelOptions_FuelPrice_SP99 FuelOptions_FuelPrice_FuelType = 11
	// SP 100.
	FuelOptions_FuelPrice_SP100 FuelOptions_FuelPrice_FuelType = 12
	// LPG.
	FuelOptions_FuelPrice_LPG FuelOptions_FuelPrice_FuelType = 13
	// E 80.
	FuelOptions_FuelPrice_E80 FuelOptions_FuelPrice_FuelType = 14
	// E 85.
	FuelOptions_FuelPrice_E85 FuelOptions_FuelPrice_FuelType = 15
	// Methane.
	FuelOptions_FuelPrice_METHANE FuelOptions_FuelPrice_FuelType = 16
	// Bio-diesel.
	FuelOptions_FuelPrice_BIO_DIESEL FuelOptions_FuelPrice_FuelType = 17
	// Truck diesel.
	FuelOptions_FuelPrice_TRUCK_DIESEL FuelOptions_FuelPrice_FuelType = 18
)

// Enum value maps for FuelOptions_FuelPrice_FuelType.
var (
	FuelOptions_FuelPrice_FuelType_name = map[int32]string{
		0:  "FUEL_TYPE_UNSPECIFIED",
		1:  "DIESEL",
		2:  "REGULAR_UNLEADED",
		3:  "MIDGRADE",
		4:  "PREMIUM",
		5:  "SP91",
		6:  "SP91_E10",
		7:  "SP92",
		8:  "SP95",
		9:  "SP95_E10",
		10: "SP98",
		11: "SP99",
		12: "SP100",
		13: "LPG",
		14: "E80",
		15: "E85",
		16: "METHANE",
		17: "BIO_DIESEL",
		18: "TRUCK_DIESEL",
	}
	FuelOptions_FuelPrice_FuelType_value = map[string]int32{
		"FUEL_TYPE_UNSPECIFIED": 0,
		"DIESEL":                1,
		"REGULAR_UNLEADED":      2,
		"MIDGRADE":              3,
		"PREMIUM":               4,
		"SP91":                  5,
		"SP91_E10":              6,
		"SP92":                  7,
		"SP95":                  8,
		"SP95_E10":              9,
		"SP98":                  10,
		"SP99":                  11,
		"SP100":                 12,
		"LPG":                   13,
		"E80":                   14,
		"E85":                   15,
		"METHANE":               16,
		"BIO_DIESEL":            17,
		"TRUCK_DIESEL":          18,
	}
)

func (x FuelOptions_FuelPrice_FuelType) Enum() *FuelOptions_FuelPrice_FuelType {
	p := new(FuelOptions_FuelPrice_FuelType)
	*p = x
	return p
}

func (x FuelOptions_FuelPrice_FuelType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FuelOptions_FuelPrice_FuelType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_places_v1_fuel_options_proto_enumTypes[0].Descriptor()
}

func (FuelOptions_FuelPrice_FuelType) Type() protoreflect.EnumType {
	return &file_google_maps_places_v1_fuel_options_proto_enumTypes[0]
}

func (x FuelOptions_FuelPrice_FuelType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FuelOptions_FuelPrice_FuelType.Descriptor instead.
func (FuelOptions_FuelPrice_FuelType) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_places_v1_fuel_options_proto_rawDescGZIP(), []int{0, 0, 0}
}

// The most recent information about fuel options in a gas station. This
// information is updated regularly.
type FuelOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The last known fuel price for each type of fuel this station has. There is
	// one entry per fuel type this station has. Order is not important.
	FuelPrices []*FuelOptions_FuelPrice `protobuf:"bytes,1,rep,name=fuel_prices,json=fuelPrices,proto3" json:"fuel_prices,omitempty"`
}

func (x *FuelOptions) Reset() {
	*x = FuelOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_places_v1_fuel_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FuelOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FuelOptions) ProtoMessage() {}

func (x *FuelOptions) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_places_v1_fuel_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FuelOptions.ProtoReflect.Descriptor instead.
func (*FuelOptions) Descriptor() ([]byte, []int) {
	return file_google_maps_places_v1_fuel_options_proto_rawDescGZIP(), []int{0}
}

func (x *FuelOptions) GetFuelPrices() []*FuelOptions_FuelPrice {
	if x != nil {
		return x.FuelPrices
	}
	return nil
}

// Fuel price information for a given type.
type FuelOptions_FuelPrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of fuel.
	Type FuelOptions_FuelPrice_FuelType `protobuf:"varint,1,opt,name=type,proto3,enum=google.maps.places.v1.FuelOptions_FuelPrice_FuelType" json:"type,omitempty"`
	// The price of the fuel.
	Price *money.Money `protobuf:"bytes,2,opt,name=price,proto3" json:"price,omitempty"`
	// The time the fuel price was last updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *FuelOptions_FuelPrice) Reset() {
	*x = FuelOptions_FuelPrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_places_v1_fuel_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FuelOptions_FuelPrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FuelOptions_FuelPrice) ProtoMessage() {}

func (x *FuelOptions_FuelPrice) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_places_v1_fuel_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FuelOptions_FuelPrice.ProtoReflect.Descriptor instead.
func (*FuelOptions_FuelPrice) Descriptor() ([]byte, []int) {
	return file_google_maps_places_v1_fuel_options_proto_rawDescGZIP(), []int{0, 0}
}

func (x *FuelOptions_FuelPrice) GetType() FuelOptions_FuelPrice_FuelType {
	if x != nil {
		return x.Type
	}
	return FuelOptions_FuelPrice_FUEL_TYPE_UNSPECIFIED
}

func (x *FuelOptions_FuelPrice) GetPrice() *money.Money {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *FuelOptions_FuelPrice) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

var File_google_maps_places_v1_fuel_options_proto protoreflect.FileDescriptor

var file_google_maps_places_v1_fuel_options_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x75, 0x65, 0x6c, 0x5f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x04, 0x0a, 0x0b,
	0x46, 0x75, 0x65, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4d, 0x0a, 0x0b, 0x66,
	0x75, 0x65, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x75, 0x65, 0x6c, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x46, 0x75, 0x65, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x0a,
	0x66, 0x75, 0x65, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x1a, 0xc5, 0x03, 0x0a, 0x09, 0x46,
	0x75, 0x65, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x6d, 0x61, 0x70, 0x73, 0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x75, 0x65, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x46, 0x75, 0x65, 0x6c, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x75, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x85, 0x02, 0x0a, 0x08, 0x46,
	0x75, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x46, 0x55, 0x45, 0x4c, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x49, 0x45, 0x53, 0x45, 0x4c, 0x10, 0x01, 0x12, 0x14,
	0x0a, 0x10, 0x52, 0x45, 0x47, 0x55, 0x4c, 0x41, 0x52, 0x5f, 0x55, 0x4e, 0x4c, 0x45, 0x41, 0x44,
	0x45, 0x44, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x49, 0x44, 0x47, 0x52, 0x41, 0x44, 0x45,
	0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x45, 0x4d, 0x49, 0x55, 0x4d, 0x10, 0x04, 0x12,
	0x08, 0x0a, 0x04, 0x53, 0x50, 0x39, 0x31, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x50, 0x39,
	0x31, 0x5f, 0x45, 0x31, 0x30, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x50, 0x39, 0x32, 0x10,
	0x07, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x50, 0x39, 0x35, 0x10, 0x08, 0x12, 0x0c, 0x0a, 0x08, 0x53,
	0x50, 0x39, 0x35, 0x5f, 0x45, 0x31, 0x30, 0x10, 0x09, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x50, 0x39,
	0x38, 0x10, 0x0a, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x50, 0x39, 0x39, 0x10, 0x0b, 0x12, 0x09, 0x0a,
	0x05, 0x53, 0x50, 0x31, 0x30, 0x30, 0x10, 0x0c, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x50, 0x47, 0x10,
	0x0d, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x38, 0x30, 0x10, 0x0e, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x38,
	0x35, 0x10, 0x0f, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x45, 0x54, 0x48, 0x41, 0x4e, 0x45, 0x10, 0x10,
	0x12, 0x0e, 0x0a, 0x0a, 0x42, 0x49, 0x4f, 0x5f, 0x44, 0x49, 0x45, 0x53, 0x45, 0x4c, 0x10, 0x11,
	0x12, 0x10, 0x0a, 0x0c, 0x54, 0x52, 0x55, 0x43, 0x4b, 0x5f, 0x44, 0x49, 0x45, 0x53, 0x45, 0x4c,
	0x10, 0x12, 0x42, 0xa4, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x42, 0x10, 0x46, 0x75, 0x65, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x73, 0x70, 0x62, 0x3b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x70, 0x62, 0xf8, 0x01, 0x01,
	0xa2, 0x02, 0x06, 0x47, 0x4d, 0x50, 0x53, 0x56, 0x31, 0xaa, 0x02, 0x15, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x15, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x4d, 0x61, 0x70, 0x73, 0x5c,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_maps_places_v1_fuel_options_proto_rawDescOnce sync.Once
	file_google_maps_places_v1_fuel_options_proto_rawDescData = file_google_maps_places_v1_fuel_options_proto_rawDesc
)

func file_google_maps_places_v1_fuel_options_proto_rawDescGZIP() []byte {
	file_google_maps_places_v1_fuel_options_proto_rawDescOnce.Do(func() {
		file_google_maps_places_v1_fuel_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_places_v1_fuel_options_proto_rawDescData)
	})
	return file_google_maps_places_v1_fuel_options_proto_rawDescData
}

var file_google_maps_places_v1_fuel_options_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_maps_places_v1_fuel_options_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_maps_places_v1_fuel_options_proto_goTypes = []interface{}{
	(FuelOptions_FuelPrice_FuelType)(0), // 0: google.maps.places.v1.FuelOptions.FuelPrice.FuelType
	(*FuelOptions)(nil),                 // 1: google.maps.places.v1.FuelOptions
	(*FuelOptions_FuelPrice)(nil),       // 2: google.maps.places.v1.FuelOptions.FuelPrice
	(*money.Money)(nil),                 // 3: google.type.Money
	(*timestamppb.Timestamp)(nil),       // 4: google.protobuf.Timestamp
}
var file_google_maps_places_v1_fuel_options_proto_depIdxs = []int32{
	2, // 0: google.maps.places.v1.FuelOptions.fuel_prices:type_name -> google.maps.places.v1.FuelOptions.FuelPrice
	0, // 1: google.maps.places.v1.FuelOptions.FuelPrice.type:type_name -> google.maps.places.v1.FuelOptions.FuelPrice.FuelType
	3, // 2: google.maps.places.v1.FuelOptions.FuelPrice.price:type_name -> google.type.Money
	4, // 3: google.maps.places.v1.FuelOptions.FuelPrice.update_time:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_maps_places_v1_fuel_options_proto_init() }
func file_google_maps_places_v1_fuel_options_proto_init() {
	if File_google_maps_places_v1_fuel_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_maps_places_v1_fuel_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FuelOptions); i {
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
		file_google_maps_places_v1_fuel_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FuelOptions_FuelPrice); i {
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
			RawDescriptor: file_google_maps_places_v1_fuel_options_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_places_v1_fuel_options_proto_goTypes,
		DependencyIndexes: file_google_maps_places_v1_fuel_options_proto_depIdxs,
		EnumInfos:         file_google_maps_places_v1_fuel_options_proto_enumTypes,
		MessageInfos:      file_google_maps_places_v1_fuel_options_proto_msgTypes,
	}.Build()
	File_google_maps_places_v1_fuel_options_proto = out.File
	file_google_maps_places_v1_fuel_options_proto_rawDesc = nil
	file_google_maps_places_v1_fuel_options_proto_goTypes = nil
	file_google_maps_places_v1_fuel_options_proto_depIdxs = nil
}
