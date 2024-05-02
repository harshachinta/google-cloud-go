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
// source: google/maps/places/v1/ev_charging.proto

package placespb

import (
	reflect "reflect"
	sync "sync"

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

// See http://ieeexplore.ieee.org/stamp/stamp.jsp?arnumber=6872107 for
// additional information/context on EV charging connector types.
type EVConnectorType int32

const (
	// Unspecified connector.
	EVConnectorType_EV_CONNECTOR_TYPE_UNSPECIFIED EVConnectorType = 0
	// Other connector types.
	EVConnectorType_EV_CONNECTOR_TYPE_OTHER EVConnectorType = 1
	// J1772 type 1 connector.
	EVConnectorType_EV_CONNECTOR_TYPE_J1772 EVConnectorType = 2
	// IEC 62196 type 2 connector. Often referred to as MENNEKES.
	EVConnectorType_EV_CONNECTOR_TYPE_TYPE_2 EVConnectorType = 3
	// CHAdeMO type connector.
	EVConnectorType_EV_CONNECTOR_TYPE_CHADEMO EVConnectorType = 4
	// Combined Charging System (AC and DC). Based on SAE.
	//
	//	Type-1 J-1772 connector
	EVConnectorType_EV_CONNECTOR_TYPE_CCS_COMBO_1 EVConnectorType = 5
	// Combined Charging System (AC and DC). Based on Type-2
	// Mennekes connector
	EVConnectorType_EV_CONNECTOR_TYPE_CCS_COMBO_2 EVConnectorType = 6
	// The generic TESLA connector. This is NACS in the North America but can be
	// non-NACS in other parts of the world (e.g. CCS Combo 2 (CCS2) or GB/T).
	// This value is less representative of an actual connector type, and more
	// represents the ability to charge a Tesla brand vehicle at a Tesla owned
	// charging station.
	EVConnectorType_EV_CONNECTOR_TYPE_TESLA EVConnectorType = 7
	// GB/T type corresponds to the GB/T standard in China. This type covers all
	// GB_T types.
	EVConnectorType_EV_CONNECTOR_TYPE_UNSPECIFIED_GB_T EVConnectorType = 8
	// Unspecified wall outlet.
	EVConnectorType_EV_CONNECTOR_TYPE_UNSPECIFIED_WALL_OUTLET EVConnectorType = 9
)

// Enum value maps for EVConnectorType.
var (
	EVConnectorType_name = map[int32]string{
		0: "EV_CONNECTOR_TYPE_UNSPECIFIED",
		1: "EV_CONNECTOR_TYPE_OTHER",
		2: "EV_CONNECTOR_TYPE_J1772",
		3: "EV_CONNECTOR_TYPE_TYPE_2",
		4: "EV_CONNECTOR_TYPE_CHADEMO",
		5: "EV_CONNECTOR_TYPE_CCS_COMBO_1",
		6: "EV_CONNECTOR_TYPE_CCS_COMBO_2",
		7: "EV_CONNECTOR_TYPE_TESLA",
		8: "EV_CONNECTOR_TYPE_UNSPECIFIED_GB_T",
		9: "EV_CONNECTOR_TYPE_UNSPECIFIED_WALL_OUTLET",
	}
	EVConnectorType_value = map[string]int32{
		"EV_CONNECTOR_TYPE_UNSPECIFIED":             0,
		"EV_CONNECTOR_TYPE_OTHER":                   1,
		"EV_CONNECTOR_TYPE_J1772":                   2,
		"EV_CONNECTOR_TYPE_TYPE_2":                  3,
		"EV_CONNECTOR_TYPE_CHADEMO":                 4,
		"EV_CONNECTOR_TYPE_CCS_COMBO_1":             5,
		"EV_CONNECTOR_TYPE_CCS_COMBO_2":             6,
		"EV_CONNECTOR_TYPE_TESLA":                   7,
		"EV_CONNECTOR_TYPE_UNSPECIFIED_GB_T":        8,
		"EV_CONNECTOR_TYPE_UNSPECIFIED_WALL_OUTLET": 9,
	}
)

func (x EVConnectorType) Enum() *EVConnectorType {
	p := new(EVConnectorType)
	*p = x
	return p
}

func (x EVConnectorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EVConnectorType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_places_v1_ev_charging_proto_enumTypes[0].Descriptor()
}

func (EVConnectorType) Type() protoreflect.EnumType {
	return &file_google_maps_places_v1_ev_charging_proto_enumTypes[0]
}

func (x EVConnectorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EVConnectorType.Descriptor instead.
func (EVConnectorType) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_places_v1_ev_charging_proto_rawDescGZIP(), []int{0}
}

// Information about the EV Charge Station hosted in Place.
// Terminology follows
// https://afdc.energy.gov/fuels/electricity_infrastructure.html One port
// could charge one car at a time. One port has one or more connectors. One
// station has one or more ports.
type EVChargeOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of connectors at this station. However, because some ports can have
	// multiple connectors but only be able to charge one car at a time (e.g.) the
	// number of connectors may be greater than the total number of cars which can
	// charge simultaneously.
	ConnectorCount int32 `protobuf:"varint,1,opt,name=connector_count,json=connectorCount,proto3" json:"connector_count,omitempty"`
	// A list of EV charging connector aggregations that contain connectors of the
	// same type and same charge rate.
	ConnectorAggregation []*EVChargeOptions_ConnectorAggregation `protobuf:"bytes,2,rep,name=connector_aggregation,json=connectorAggregation,proto3" json:"connector_aggregation,omitempty"`
}

func (x *EVChargeOptions) Reset() {
	*x = EVChargeOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_places_v1_ev_charging_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EVChargeOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EVChargeOptions) ProtoMessage() {}

func (x *EVChargeOptions) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_places_v1_ev_charging_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EVChargeOptions.ProtoReflect.Descriptor instead.
func (*EVChargeOptions) Descriptor() ([]byte, []int) {
	return file_google_maps_places_v1_ev_charging_proto_rawDescGZIP(), []int{0}
}

func (x *EVChargeOptions) GetConnectorCount() int32 {
	if x != nil {
		return x.ConnectorCount
	}
	return 0
}

func (x *EVChargeOptions) GetConnectorAggregation() []*EVChargeOptions_ConnectorAggregation {
	if x != nil {
		return x.ConnectorAggregation
	}
	return nil
}

// EV charging information grouped by [type, max_charge_rate_kw].
// Shows EV charge aggregation of connectors that have the same type and max
// charge rate in kw.
type EVChargeOptions_ConnectorAggregation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The connector type of this aggregation.
	Type EVConnectorType `protobuf:"varint,1,opt,name=type,proto3,enum=google.maps.places.v1.EVConnectorType" json:"type,omitempty"`
	// The static max charging rate in kw of each connector in the aggregation.
	MaxChargeRateKw float64 `protobuf:"fixed64,2,opt,name=max_charge_rate_kw,json=maxChargeRateKw,proto3" json:"max_charge_rate_kw,omitempty"`
	// Number of connectors in this aggregation.
	Count int32 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	// Number of connectors in this aggregation that are currently available.
	AvailableCount *int32 `protobuf:"varint,4,opt,name=available_count,json=availableCount,proto3,oneof" json:"available_count,omitempty"`
	// Number of connectors in this aggregation that are currently out of
	// service.
	OutOfServiceCount *int32 `protobuf:"varint,5,opt,name=out_of_service_count,json=outOfServiceCount,proto3,oneof" json:"out_of_service_count,omitempty"`
	// The timestamp when the connector availability information in this
	// aggregation was last updated.
	AvailabilityLastUpdateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=availability_last_update_time,json=availabilityLastUpdateTime,proto3" json:"availability_last_update_time,omitempty"`
}

func (x *EVChargeOptions_ConnectorAggregation) Reset() {
	*x = EVChargeOptions_ConnectorAggregation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_places_v1_ev_charging_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EVChargeOptions_ConnectorAggregation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EVChargeOptions_ConnectorAggregation) ProtoMessage() {}

func (x *EVChargeOptions_ConnectorAggregation) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_places_v1_ev_charging_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EVChargeOptions_ConnectorAggregation.ProtoReflect.Descriptor instead.
func (*EVChargeOptions_ConnectorAggregation) Descriptor() ([]byte, []int) {
	return file_google_maps_places_v1_ev_charging_proto_rawDescGZIP(), []int{0, 0}
}

func (x *EVChargeOptions_ConnectorAggregation) GetType() EVConnectorType {
	if x != nil {
		return x.Type
	}
	return EVConnectorType_EV_CONNECTOR_TYPE_UNSPECIFIED
}

func (x *EVChargeOptions_ConnectorAggregation) GetMaxChargeRateKw() float64 {
	if x != nil {
		return x.MaxChargeRateKw
	}
	return 0
}

func (x *EVChargeOptions_ConnectorAggregation) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *EVChargeOptions_ConnectorAggregation) GetAvailableCount() int32 {
	if x != nil && x.AvailableCount != nil {
		return *x.AvailableCount
	}
	return 0
}

func (x *EVChargeOptions_ConnectorAggregation) GetOutOfServiceCount() int32 {
	if x != nil && x.OutOfServiceCount != nil {
		return *x.OutOfServiceCount
	}
	return 0
}

func (x *EVChargeOptions_ConnectorAggregation) GetAvailabilityLastUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AvailabilityLastUpdateTime
	}
	return nil
}

var File_google_maps_places_v1_ev_charging_proto protoreflect.FileDescriptor

var file_google_maps_places_v1_ev_charging_proto_rawDesc = []byte{
	0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb4, 0x04, 0x0a, 0x0f, 0x45, 0x56, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x70,
	0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x56, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x41,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x14, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x85, 0x03, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x41, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x56, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x68, 0x61,
	0x72, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65,
	0x4b, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x0f, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x00, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x14, 0x6f, 0x75, 0x74, 0x5f, 0x6f, 0x66,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x11, 0x6f, 0x75, 0x74, 0x4f, 0x66, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x5d, 0x0a, 0x1d,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x1a, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x61, 0x73,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x12, 0x0a, 0x10, 0x5f,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x17, 0x0a, 0x15, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x6f, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2a, 0xe5, 0x02, 0x0a, 0x0f, 0x45, 0x56, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x1d,
	0x45, 0x56, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x1b, 0x0a, 0x17, 0x45, 0x56, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17,
	0x45, 0x56, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4a, 0x31, 0x37, 0x37, 0x32, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x56, 0x5f,
	0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x32, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x45, 0x56, 0x5f, 0x43, 0x4f,
	0x4e, 0x4e, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x48, 0x41,
	0x44, 0x45, 0x4d, 0x4f, 0x10, 0x04, 0x12, 0x21, 0x0a, 0x1d, 0x45, 0x56, 0x5f, 0x43, 0x4f, 0x4e,
	0x4e, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x43, 0x53, 0x5f,
	0x43, 0x4f, 0x4d, 0x42, 0x4f, 0x5f, 0x31, 0x10, 0x05, 0x12, 0x21, 0x0a, 0x1d, 0x45, 0x56, 0x5f,
	0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43,
	0x43, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x42, 0x4f, 0x5f, 0x32, 0x10, 0x06, 0x12, 0x1b, 0x0a, 0x17,
	0x45, 0x56, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x54, 0x45, 0x53, 0x4c, 0x41, 0x10, 0x07, 0x12, 0x26, 0x0a, 0x22, 0x45, 0x56, 0x5f,
	0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x5f, 0x47, 0x42, 0x5f, 0x54, 0x10,
	0x08, 0x12, 0x2d, 0x0a, 0x29, 0x45, 0x56, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x4f,
	0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x5f, 0x57, 0x41, 0x4c, 0x4c, 0x5f, 0x4f, 0x55, 0x54, 0x4c, 0x45, 0x54, 0x10, 0x09,
	0x42, 0xa3, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x6d, 0x61, 0x70, 0x73, 0x2e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0f,
	0x45, 0x76, 0x43, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x70,
	0x62, 0x3b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x70, 0x62, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x06,
	0x47, 0x4d, 0x50, 0x53, 0x56, 0x31, 0xaa, 0x02, 0x15, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x4d, 0x61, 0x70, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x15, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x4d, 0x61, 0x70, 0x73, 0x5c, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_places_v1_ev_charging_proto_rawDescOnce sync.Once
	file_google_maps_places_v1_ev_charging_proto_rawDescData = file_google_maps_places_v1_ev_charging_proto_rawDesc
)

func file_google_maps_places_v1_ev_charging_proto_rawDescGZIP() []byte {
	file_google_maps_places_v1_ev_charging_proto_rawDescOnce.Do(func() {
		file_google_maps_places_v1_ev_charging_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_places_v1_ev_charging_proto_rawDescData)
	})
	return file_google_maps_places_v1_ev_charging_proto_rawDescData
}

var file_google_maps_places_v1_ev_charging_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_maps_places_v1_ev_charging_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_maps_places_v1_ev_charging_proto_goTypes = []interface{}{
	(EVConnectorType)(0),                         // 0: google.maps.places.v1.EVConnectorType
	(*EVChargeOptions)(nil),                      // 1: google.maps.places.v1.EVChargeOptions
	(*EVChargeOptions_ConnectorAggregation)(nil), // 2: google.maps.places.v1.EVChargeOptions.ConnectorAggregation
	(*timestamppb.Timestamp)(nil),                // 3: google.protobuf.Timestamp
}
var file_google_maps_places_v1_ev_charging_proto_depIdxs = []int32{
	2, // 0: google.maps.places.v1.EVChargeOptions.connector_aggregation:type_name -> google.maps.places.v1.EVChargeOptions.ConnectorAggregation
	0, // 1: google.maps.places.v1.EVChargeOptions.ConnectorAggregation.type:type_name -> google.maps.places.v1.EVConnectorType
	3, // 2: google.maps.places.v1.EVChargeOptions.ConnectorAggregation.availability_last_update_time:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_maps_places_v1_ev_charging_proto_init() }
func file_google_maps_places_v1_ev_charging_proto_init() {
	if File_google_maps_places_v1_ev_charging_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_maps_places_v1_ev_charging_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EVChargeOptions); i {
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
		file_google_maps_places_v1_ev_charging_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EVChargeOptions_ConnectorAggregation); i {
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
	file_google_maps_places_v1_ev_charging_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_maps_places_v1_ev_charging_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_places_v1_ev_charging_proto_goTypes,
		DependencyIndexes: file_google_maps_places_v1_ev_charging_proto_depIdxs,
		EnumInfos:         file_google_maps_places_v1_ev_charging_proto_enumTypes,
		MessageInfos:      file_google_maps_places_v1_ev_charging_proto_msgTypes,
	}.Build()
	File_google_maps_places_v1_ev_charging_proto = out.File
	file_google_maps_places_v1_ev_charging_proto_rawDesc = nil
	file_google_maps_places_v1_ev_charging_proto_goTypes = nil
	file_google_maps_places_v1_ev_charging_proto_depIdxs = nil
}
