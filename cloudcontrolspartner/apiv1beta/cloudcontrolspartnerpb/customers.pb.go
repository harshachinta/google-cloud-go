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
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: google/cloud/cloudcontrolspartner/v1beta/customers.proto

package cloudcontrolspartnerpb

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

// Enum for possible onboarding steps
type CustomerOnboardingStep_Step int32

const (
	// Unspecified step
	CustomerOnboardingStep_STEP_UNSPECIFIED CustomerOnboardingStep_Step = 0
	// KAJ Enrollment
	CustomerOnboardingStep_KAJ_ENROLLMENT CustomerOnboardingStep_Step = 1
	// Customer Environment
	CustomerOnboardingStep_CUSTOMER_ENVIRONMENT CustomerOnboardingStep_Step = 2
)

// Enum value maps for CustomerOnboardingStep_Step.
var (
	CustomerOnboardingStep_Step_name = map[int32]string{
		0: "STEP_UNSPECIFIED",
		1: "KAJ_ENROLLMENT",
		2: "CUSTOMER_ENVIRONMENT",
	}
	CustomerOnboardingStep_Step_value = map[string]int32{
		"STEP_UNSPECIFIED":     0,
		"KAJ_ENROLLMENT":       1,
		"CUSTOMER_ENVIRONMENT": 2,
	}
)

func (x CustomerOnboardingStep_Step) Enum() *CustomerOnboardingStep_Step {
	p := new(CustomerOnboardingStep_Step)
	*p = x
	return p
}

func (x CustomerOnboardingStep_Step) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CustomerOnboardingStep_Step) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_enumTypes[0].Descriptor()
}

func (CustomerOnboardingStep_Step) Type() protoreflect.EnumType {
	return &file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_enumTypes[0]
}

func (x CustomerOnboardingStep_Step) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CustomerOnboardingStep_Step.Descriptor instead.
func (CustomerOnboardingStep_Step) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_rawDescGZIP(), []int{5, 0}
}

// Contains metadata around a Cloud Controls Partner Customer
type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier. Format:
	// organizations/{organization}/locations/{location}/customers/{customer}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The customer organization's display name. E.g. "google.com".
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Container for customer onboarding steps
	CustomerOnboardingState *CustomerOnboardingState `protobuf:"bytes,3,opt,name=customer_onboarding_state,json=customerOnboardingState,proto3" json:"customer_onboarding_state,omitempty"`
	// Indicates whether a customer is fully onboarded
	IsOnboarded bool `protobuf:"varint,4,opt,name=is_onboarded,json=isOnboarded,proto3" json:"is_onboarded,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_rawDescGZIP(), []int{0}
}

func (x *Customer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Customer) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Customer) GetCustomerOnboardingState() *CustomerOnboardingState {
	if x != nil {
		return x.CustomerOnboardingState
	}
	return nil
}

func (x *Customer) GetIsOnboarded() bool {
	if x != nil {
		return x.IsOnboarded
	}
	return false
}

// Request to list customers
type ListCustomersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Parent resource
	// Format: organizations/{organization}/locations/{location}
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of Customers to return. The service may return fewer
	// than this value. If unspecified, at most 500 Customers will be returned.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A page token, received from a previous `ListCustomers` call.
	// Provide this to retrieve the subsequent page.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Optional. Filtering results
	Filter string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	// Optional. Hint for how to order the results
	OrderBy string `protobuf:"bytes,5,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
}

func (x *ListCustomersRequest) Reset() {
	*x = ListCustomersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCustomersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCustomersRequest) ProtoMessage() {}

func (x *ListCustomersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCustomersRequest.ProtoReflect.Descriptor instead.
func (*ListCustomersRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_rawDescGZIP(), []int{1}
}

func (x *ListCustomersRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListCustomersRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListCustomersRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListCustomersRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *ListCustomersRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

// Response message for list customer Customers requests
type ListCustomersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of customers
	Customers []*Customer `protobuf:"bytes,1,rep,name=customers,proto3" json:"customers,omitempty"`
	// A token that can be sent as `page_token` to retrieve the next page.
	// If this field is omitted, there are no subsequent pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// Locations that could not be reached.
	Unreachable []string `protobuf:"bytes,3,rep,name=unreachable,proto3" json:"unreachable,omitempty"`
}

func (x *ListCustomersResponse) Reset() {
	*x = ListCustomersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCustomersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCustomersResponse) ProtoMessage() {}

func (x *ListCustomersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCustomersResponse.ProtoReflect.Descriptor instead.
func (*ListCustomersResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_rawDescGZIP(), []int{2}
}

func (x *ListCustomersResponse) GetCustomers() []*Customer {
	if x != nil {
		return x.Customers
	}
	return nil
}

func (x *ListCustomersResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListCustomersResponse) GetUnreachable() []string {
	if x != nil {
		return x.Unreachable
	}
	return nil
}

// Message for getting a customer
type GetCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Format:
	// organizations/{organization}/locations/{location}/customers/{customer}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetCustomerRequest) Reset() {
	*x = GetCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerRequest) ProtoMessage() {}

func (x *GetCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_rawDescGZIP(), []int{3}
}

func (x *GetCustomerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Container for customer onboarding steps
type CustomerOnboardingState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of customer onboarding steps
	OnboardingSteps []*CustomerOnboardingStep `protobuf:"bytes,1,rep,name=onboarding_steps,json=onboardingSteps,proto3" json:"onboarding_steps,omitempty"`
}

func (x *CustomerOnboardingState) Reset() {
	*x = CustomerOnboardingState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerOnboardingState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerOnboardingState) ProtoMessage() {}

func (x *CustomerOnboardingState) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerOnboardingState.ProtoReflect.Descriptor instead.
func (*CustomerOnboardingState) Descriptor() ([]byte, []int) {
	return file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_rawDescGZIP(), []int{4}
}

func (x *CustomerOnboardingState) GetOnboardingSteps() []*CustomerOnboardingStep {
	if x != nil {
		return x.OnboardingSteps
	}
	return nil
}

// Container for customer onboarding information
type CustomerOnboardingStep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The onboarding step
	Step CustomerOnboardingStep_Step `protobuf:"varint,1,opt,name=step,proto3,enum=google.cloud.cloudcontrolspartner.v1beta.CustomerOnboardingStep_Step" json:"step,omitempty"`
	// The starting time of the onboarding step
	StartTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// The completion time of the onboarding step
	CompletionTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=completion_time,json=completionTime,proto3" json:"completion_time,omitempty"`
	// Output only. Current state of the step
	CompletionState CompletionState `protobuf:"varint,4,opt,name=completion_state,json=completionState,proto3,enum=google.cloud.cloudcontrolspartner.v1beta.CompletionState" json:"completion_state,omitempty"`
}

func (x *CustomerOnboardingStep) Reset() {
	*x = CustomerOnboardingStep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerOnboardingStep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerOnboardingStep) ProtoMessage() {}

func (x *CustomerOnboardingStep) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerOnboardingStep.ProtoReflect.Descriptor instead.
func (*CustomerOnboardingStep) Descriptor() ([]byte, []int) {
	return file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_rawDescGZIP(), []int{5}
}

func (x *CustomerOnboardingStep) GetStep() CustomerOnboardingStep_Step {
	if x != nil {
		return x.Step
	}
	return CustomerOnboardingStep_STEP_UNSPECIFIED
}

func (x *CustomerOnboardingStep) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *CustomerOnboardingStep) GetCompletionTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CompletionTime
	}
	return nil
}

func (x *CustomerOnboardingStep) GetCompletionState() CompletionState {
	if x != nil {
		return x.CompletionState
	}
	return CompletionState_COMPLETION_STATE_UNSPECIFIED
}

var File_google_cloud_cloudcontrolspartner_v1beta_customers_proto protoreflect.FileDescriptor

var file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xfa, 0x02, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12,
	0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x7d, 0x0a, 0x19, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x17, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4f, 0x6e, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73,
	0x5f, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x69, 0x73, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x65, 0x64, 0x3a, 0x8f, 0x01,
	0xea, 0x41, 0x8b, 0x01, 0x0a, 0x2c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x12, 0x46, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f,
	0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x7d, 0x2a, 0x09, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x32, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22,
	0xdd, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x34, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2e,
	0x12, 0x2c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x1e, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x22,
	0xb3, 0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x09, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x52, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x63, 0x68, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x63,
	0x68, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x5e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x34, 0xe0, 0x41, 0x02, 0xfa, 0x41,
	0x2e, 0x0a, 0x2c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x17, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x6b, 0x0a, 0x10, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f,
	0x73, 0x74, 0x65, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4f,
	0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x65, 0x70, 0x52, 0x0f, 0x6f,
	0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x65, 0x70, 0x73, 0x22, 0xaa,
	0x03, 0x0a, 0x16, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4f, 0x6e, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x65, 0x70, 0x12, 0x59, 0x0a, 0x04, 0x73, 0x74, 0x65,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x45, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4f, 0x6e, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x65, 0x70, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x52, 0x04,
	0x73, 0x74, 0x65, 0x70, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x43, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x69, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0f,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22,
	0x4a, 0x0a, 0x04, 0x53, 0x74, 0x65, 0x70, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x45, 0x50, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a,
	0x0e, 0x4b, 0x41, 0x4a, 0x5f, 0x45, 0x4e, 0x52, 0x4f, 0x4c, 0x4c, 0x4d, 0x45, 0x4e, 0x54, 0x10,
	0x01, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x45, 0x4e,
	0x56, 0x49, 0x52, 0x4f, 0x4e, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x42, 0xa6, 0x02, 0x0a, 0x2c,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x42, 0x0e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x60,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x70, 0x62,
	0xaa, 0x02, 0x28, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x50, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0xca, 0x02, 0x28, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5c,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0xea, 0x02, 0x2b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x73, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_rawDescOnce sync.Once
	file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_rawDescData = file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_rawDesc
)

func file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_rawDescGZIP() []byte {
	file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_rawDescOnce.Do(func() {
		file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_rawDescData)
	})
	return file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_rawDescData
}

var file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_goTypes = []interface{}{
	(CustomerOnboardingStep_Step)(0), // 0: google.cloud.cloudcontrolspartner.v1beta.CustomerOnboardingStep.Step
	(*Customer)(nil),                 // 1: google.cloud.cloudcontrolspartner.v1beta.Customer
	(*ListCustomersRequest)(nil),     // 2: google.cloud.cloudcontrolspartner.v1beta.ListCustomersRequest
	(*ListCustomersResponse)(nil),    // 3: google.cloud.cloudcontrolspartner.v1beta.ListCustomersResponse
	(*GetCustomerRequest)(nil),       // 4: google.cloud.cloudcontrolspartner.v1beta.GetCustomerRequest
	(*CustomerOnboardingState)(nil),  // 5: google.cloud.cloudcontrolspartner.v1beta.CustomerOnboardingState
	(*CustomerOnboardingStep)(nil),   // 6: google.cloud.cloudcontrolspartner.v1beta.CustomerOnboardingStep
	(*timestamppb.Timestamp)(nil),    // 7: google.protobuf.Timestamp
	(CompletionState)(0),             // 8: google.cloud.cloudcontrolspartner.v1beta.CompletionState
}
var file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_depIdxs = []int32{
	5, // 0: google.cloud.cloudcontrolspartner.v1beta.Customer.customer_onboarding_state:type_name -> google.cloud.cloudcontrolspartner.v1beta.CustomerOnboardingState
	1, // 1: google.cloud.cloudcontrolspartner.v1beta.ListCustomersResponse.customers:type_name -> google.cloud.cloudcontrolspartner.v1beta.Customer
	6, // 2: google.cloud.cloudcontrolspartner.v1beta.CustomerOnboardingState.onboarding_steps:type_name -> google.cloud.cloudcontrolspartner.v1beta.CustomerOnboardingStep
	0, // 3: google.cloud.cloudcontrolspartner.v1beta.CustomerOnboardingStep.step:type_name -> google.cloud.cloudcontrolspartner.v1beta.CustomerOnboardingStep.Step
	7, // 4: google.cloud.cloudcontrolspartner.v1beta.CustomerOnboardingStep.start_time:type_name -> google.protobuf.Timestamp
	7, // 5: google.cloud.cloudcontrolspartner.v1beta.CustomerOnboardingStep.completion_time:type_name -> google.protobuf.Timestamp
	8, // 6: google.cloud.cloudcontrolspartner.v1beta.CustomerOnboardingStep.completion_state:type_name -> google.cloud.cloudcontrolspartner.v1beta.CompletionState
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_init() }
func file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_init() {
	if File_google_cloud_cloudcontrolspartner_v1beta_customers_proto != nil {
		return
	}
	file_google_cloud_cloudcontrolspartner_v1beta_completion_state_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customer); i {
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
		file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCustomersRequest); i {
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
		file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCustomersResponse); i {
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
		file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerRequest); i {
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
		file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerOnboardingState); i {
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
		file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerOnboardingStep); i {
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
			RawDescriptor: file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_goTypes,
		DependencyIndexes: file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_depIdxs,
		EnumInfos:         file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_enumTypes,
		MessageInfos:      file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_msgTypes,
	}.Build()
	File_google_cloud_cloudcontrolspartner_v1beta_customers_proto = out.File
	file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_rawDesc = nil
	file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_goTypes = nil
	file_google_cloud_cloudcontrolspartner_v1beta_customers_proto_depIdxs = nil
}
