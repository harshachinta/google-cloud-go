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
// source: google/cloud/dialogflow/cx/v3beta1/deployment.proto

package cxpb

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
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The state of the deployment.
type Deployment_State int32

const (
	// State unspecified.
	Deployment_STATE_UNSPECIFIED Deployment_State = 0
	// The deployment is running.
	Deployment_RUNNING Deployment_State = 1
	// The deployment succeeded.
	Deployment_SUCCEEDED Deployment_State = 2
	// The deployment failed.
	Deployment_FAILED Deployment_State = 3
)

// Enum value maps for Deployment_State.
var (
	Deployment_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "RUNNING",
		2: "SUCCEEDED",
		3: "FAILED",
	}
	Deployment_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"RUNNING":           1,
		"SUCCEEDED":         2,
		"FAILED":            3,
	}
)

func (x Deployment_State) Enum() *Deployment_State {
	p := new(Deployment_State)
	*p = x
	return p
}

func (x Deployment_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Deployment_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_enumTypes[0].Descriptor()
}

func (Deployment_State) Type() protoreflect.EnumType {
	return &file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_enumTypes[0]
}

func (x Deployment_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Deployment_State.Descriptor instead.
func (Deployment_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_rawDescGZIP(), []int{0, 0}
}

// Represents a deployment in an environment. A deployment happens when a flow
// version configured to be active in the environment. You can configure running
// pre-deployment steps, e.g. running validation test cases, experiment
// auto-rollout, etc.
type Deployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the deployment.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent
	// ID>/environments/<Environment ID>/deployments/<Deployment ID>.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The name of the flow version for this deployment.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent
	// ID>/flows/<Flow ID>/versions/<Verion ID>.
	FlowVersion string `protobuf:"bytes,2,opt,name=flow_version,json=flowVersion,proto3" json:"flow_version,omitempty"`
	// The current state of the deployment.
	State Deployment_State `protobuf:"varint,3,opt,name=state,proto3,enum=google.cloud.dialogflow.cx.v3beta1.Deployment_State" json:"state,omitempty"`
	// Result of the deployment.
	Result *Deployment_Result `protobuf:"bytes,4,opt,name=result,proto3" json:"result,omitempty"`
	// Start time of this deployment.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// End time of this deployment.
	EndTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *Deployment) Reset() {
	*x = Deployment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deployment) ProtoMessage() {}

func (x *Deployment) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deployment.ProtoReflect.Descriptor instead.
func (*Deployment) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_rawDescGZIP(), []int{0}
}

func (x *Deployment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Deployment) GetFlowVersion() string {
	if x != nil {
		return x.FlowVersion
	}
	return ""
}

func (x *Deployment) GetState() Deployment_State {
	if x != nil {
		return x.State
	}
	return Deployment_STATE_UNSPECIFIED
}

func (x *Deployment) GetResult() *Deployment_Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *Deployment) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Deployment) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

// The request message for
// [Deployments.ListDeployments][google.cloud.dialogflow.cx.v3beta1.Deployments.ListDeployments].
type ListDeploymentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The [Environment][google.cloud.dialogflow.cx.v3beta1.Environment]
	// to list all environments for. Format: `projects/<Project
	// ID>/locations/<Location ID>/agents/<Agent ID>/environments/<Environment
	// ID>`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of items to return in a single page. By default 20 and
	// at most 100.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token value returned from a previous list request.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListDeploymentsRequest) Reset() {
	*x = ListDeploymentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDeploymentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeploymentsRequest) ProtoMessage() {}

func (x *ListDeploymentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeploymentsRequest.ProtoReflect.Descriptor instead.
func (*ListDeploymentsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_rawDescGZIP(), []int{1}
}

func (x *ListDeploymentsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListDeploymentsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListDeploymentsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// The response message for
// [Deployments.ListDeployments][google.cloud.dialogflow.cx.v3beta1.Deployments.ListDeployments].
type ListDeploymentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of deployments. There will be a maximum number of items
	// returned based on the page_size field in the request. The list may in some
	// cases be empty or contain fewer entries than page_size even if this isn't
	// the last page.
	Deployments []*Deployment `protobuf:"bytes,1,rep,name=deployments,proto3" json:"deployments,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no more
	// results in the list.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListDeploymentsResponse) Reset() {
	*x = ListDeploymentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDeploymentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeploymentsResponse) ProtoMessage() {}

func (x *ListDeploymentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeploymentsResponse.ProtoReflect.Descriptor instead.
func (*ListDeploymentsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_rawDescGZIP(), []int{2}
}

func (x *ListDeploymentsResponse) GetDeployments() []*Deployment {
	if x != nil {
		return x.Deployments
	}
	return nil
}

func (x *ListDeploymentsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// The request message for
// [Deployments.GetDeployment][google.cloud.dialogflow.cx.v3beta1.Deployments.GetDeployment].
type GetDeploymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the
	// [Deployment][google.cloud.dialogflow.cx.v3beta1.Deployment]. Format:
	// `projects/<Project ID>/locations/<Location ID>/agents/<Agent
	// ID>/environments/<Environment ID>/deployments/<Deployment ID>`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetDeploymentRequest) Reset() {
	*x = GetDeploymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeploymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeploymentRequest) ProtoMessage() {}

func (x *GetDeploymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeploymentRequest.ProtoReflect.Descriptor instead.
func (*GetDeploymentRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_rawDescGZIP(), []int{3}
}

func (x *GetDeploymentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Result of the deployment.
type Deployment_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Results of test cases running before the deployment.
	// Format: `projects/<Project ID>/locations/<Location ID>/agents/<Agent
	// ID>/testCases/<TestCase ID>/results/<TestCaseResult ID>`.
	DeploymentTestResults []string `protobuf:"bytes,1,rep,name=deployment_test_results,json=deploymentTestResults,proto3" json:"deployment_test_results,omitempty"`
	// The name of the experiment triggered by this deployment.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent
	// ID>/environments/<Environment ID>/experiments/<Experiment ID>.
	Experiment string `protobuf:"bytes,2,opt,name=experiment,proto3" json:"experiment,omitempty"`
}

func (x *Deployment_Result) Reset() {
	*x = Deployment_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deployment_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deployment_Result) ProtoMessage() {}

func (x *Deployment_Result) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deployment_Result.ProtoReflect.Descriptor instead.
func (*Deployment_Result) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Deployment_Result) GetDeploymentTestResults() []string {
	if x != nil {
		return x.DeploymentTestResults
	}
	return nil
}

func (x *Deployment_Result) GetExperiment() string {
	if x != nil {
		return x.Experiment
	}
	return ""
}

var File_google_cloud_dialogflow_cx_v3beta1_deployment_proto protoreflect.FileDescriptor

var file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x78, 0x2f, 0x76, 0x33, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63,
	0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x06,
	0x0a, 0x0a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x49, 0x0a, 0x0c, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x64, 0x69, 0x61,
	0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b,
	0x66, 0x6c, 0x6f, 0x77, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x4d, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0xba, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x65, 0x0a, 0x17, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x42, 0x2d, 0xfa, 0x41, 0x2a, 0x0a, 0x28, 0x64, 0x69, 0x61, 0x6c, 0x6f,
	0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x15, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x49, 0x0a, 0x0a, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29,
	0xfa, 0x41, 0x26, 0x0a, 0x24, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x46, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15,
	0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10,
	0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x3a, 0x96, 0x01,
	0xea, 0x41, 0x92, 0x01, 0x0a, 0x24, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x6a, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x7d, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x7d, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f,
	0x7b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x7d, 0x2f, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x7d, 0x22, 0x9a, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x44, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x2c, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x26, 0x12, 0x24, 0x64, 0x69, 0x61, 0x6c, 0x6f,
	0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x93, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x50, 0x0a, 0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63,
	0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x58, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x40, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x2c, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x26, 0x0a, 0x24, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x32, 0xcc, 0x04, 0x0a, 0x0b, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0xe9, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78,
	0x2e, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x5d, 0xda, 0x41, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x4e, 0x12, 0x4c, 0x2f, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x2a, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x2a, 0x7d, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0xd6, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76,
	0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f,
	0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x5b, 0xda, 0x41, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4e, 0x12, 0x4c, 0x2f, 0x76, 0x33, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x2a, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x1a, 0x78, 0xca, 0x41, 0x19, 0x64, 0x69, 0x61,
	0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x59, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f,
	0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c,
	0x6f, 0x77, 0x42, 0xc9, 0x01, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0f, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x36, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77,
	0x2f, 0x63, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63,
	0x78, 0x70, 0x62, 0x3b, 0x63, 0x78, 0x70, 0x62, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x02, 0x44, 0x46,
	0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x78, 0x2e, 0x56, 0x33,
	0x42, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f,
	0x77, 0x3a, 0x3a, 0x43, 0x58, 0x3a, 0x3a, 0x56, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_rawDescOnce sync.Once
	file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_rawDescData = file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_rawDesc
)

func file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_rawDescGZIP() []byte {
	file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_rawDescOnce.Do(func() {
		file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_rawDescData)
	})
	return file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_rawDescData
}

var file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_goTypes = []interface{}{
	(Deployment_State)(0),           // 0: google.cloud.dialogflow.cx.v3beta1.Deployment.State
	(*Deployment)(nil),              // 1: google.cloud.dialogflow.cx.v3beta1.Deployment
	(*ListDeploymentsRequest)(nil),  // 2: google.cloud.dialogflow.cx.v3beta1.ListDeploymentsRequest
	(*ListDeploymentsResponse)(nil), // 3: google.cloud.dialogflow.cx.v3beta1.ListDeploymentsResponse
	(*GetDeploymentRequest)(nil),    // 4: google.cloud.dialogflow.cx.v3beta1.GetDeploymentRequest
	(*Deployment_Result)(nil),       // 5: google.cloud.dialogflow.cx.v3beta1.Deployment.Result
	(*timestamppb.Timestamp)(nil),   // 6: google.protobuf.Timestamp
}
var file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_depIdxs = []int32{
	0, // 0: google.cloud.dialogflow.cx.v3beta1.Deployment.state:type_name -> google.cloud.dialogflow.cx.v3beta1.Deployment.State
	5, // 1: google.cloud.dialogflow.cx.v3beta1.Deployment.result:type_name -> google.cloud.dialogflow.cx.v3beta1.Deployment.Result
	6, // 2: google.cloud.dialogflow.cx.v3beta1.Deployment.start_time:type_name -> google.protobuf.Timestamp
	6, // 3: google.cloud.dialogflow.cx.v3beta1.Deployment.end_time:type_name -> google.protobuf.Timestamp
	1, // 4: google.cloud.dialogflow.cx.v3beta1.ListDeploymentsResponse.deployments:type_name -> google.cloud.dialogflow.cx.v3beta1.Deployment
	2, // 5: google.cloud.dialogflow.cx.v3beta1.Deployments.ListDeployments:input_type -> google.cloud.dialogflow.cx.v3beta1.ListDeploymentsRequest
	4, // 6: google.cloud.dialogflow.cx.v3beta1.Deployments.GetDeployment:input_type -> google.cloud.dialogflow.cx.v3beta1.GetDeploymentRequest
	3, // 7: google.cloud.dialogflow.cx.v3beta1.Deployments.ListDeployments:output_type -> google.cloud.dialogflow.cx.v3beta1.ListDeploymentsResponse
	1, // 8: google.cloud.dialogflow.cx.v3beta1.Deployments.GetDeployment:output_type -> google.cloud.dialogflow.cx.v3beta1.Deployment
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_init() }
func file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_init() {
	if File_google_cloud_dialogflow_cx_v3beta1_deployment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deployment); i {
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
		file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDeploymentsRequest); i {
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
		file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDeploymentsResponse); i {
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
		file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeploymentRequest); i {
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
		file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deployment_Result); i {
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
			RawDescriptor: file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_goTypes,
		DependencyIndexes: file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_depIdxs,
		EnumInfos:         file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_enumTypes,
		MessageInfos:      file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_msgTypes,
	}.Build()
	File_google_cloud_dialogflow_cx_v3beta1_deployment_proto = out.File
	file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_rawDesc = nil
	file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_goTypes = nil
	file_google_cloud_dialogflow_cx_v3beta1_deployment_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DeploymentsClient is the client API for Deployments service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeploymentsClient interface {
	// Returns the list of all deployments in the specified
	// [Environment][google.cloud.dialogflow.cx.v3beta1.Environment].
	ListDeployments(ctx context.Context, in *ListDeploymentsRequest, opts ...grpc.CallOption) (*ListDeploymentsResponse, error)
	// Retrieves the specified
	// [Deployment][google.cloud.dialogflow.cx.v3beta1.Deployment].
	GetDeployment(ctx context.Context, in *GetDeploymentRequest, opts ...grpc.CallOption) (*Deployment, error)
}

type deploymentsClient struct {
	cc grpc.ClientConnInterface
}

func NewDeploymentsClient(cc grpc.ClientConnInterface) DeploymentsClient {
	return &deploymentsClient{cc}
}

func (c *deploymentsClient) ListDeployments(ctx context.Context, in *ListDeploymentsRequest, opts ...grpc.CallOption) (*ListDeploymentsResponse, error) {
	out := new(ListDeploymentsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.cx.v3beta1.Deployments/ListDeployments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentsClient) GetDeployment(ctx context.Context, in *GetDeploymentRequest, opts ...grpc.CallOption) (*Deployment, error) {
	out := new(Deployment)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.cx.v3beta1.Deployments/GetDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeploymentsServer is the server API for Deployments service.
type DeploymentsServer interface {
	// Returns the list of all deployments in the specified
	// [Environment][google.cloud.dialogflow.cx.v3beta1.Environment].
	ListDeployments(context.Context, *ListDeploymentsRequest) (*ListDeploymentsResponse, error)
	// Retrieves the specified
	// [Deployment][google.cloud.dialogflow.cx.v3beta1.Deployment].
	GetDeployment(context.Context, *GetDeploymentRequest) (*Deployment, error)
}

// UnimplementedDeploymentsServer can be embedded to have forward compatible implementations.
type UnimplementedDeploymentsServer struct {
}

func (*UnimplementedDeploymentsServer) ListDeployments(context.Context, *ListDeploymentsRequest) (*ListDeploymentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDeployments not implemented")
}
func (*UnimplementedDeploymentsServer) GetDeployment(context.Context, *GetDeploymentRequest) (*Deployment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeployment not implemented")
}

func RegisterDeploymentsServer(s *grpc.Server, srv DeploymentsServer) {
	s.RegisterService(&_Deployments_serviceDesc, srv)
}

func _Deployments_ListDeployments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeploymentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentsServer).ListDeployments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.cx.v3beta1.Deployments/ListDeployments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentsServer).ListDeployments(ctx, req.(*ListDeploymentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deployments_GetDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentsServer).GetDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.cx.v3beta1.Deployments/GetDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentsServer).GetDeployment(ctx, req.(*GetDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Deployments_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.dialogflow.cx.v3beta1.Deployments",
	HandlerType: (*DeploymentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDeployments",
			Handler:    _Deployments_ListDeployments_Handler,
		},
		{
			MethodName: "GetDeployment",
			Handler:    _Deployments_GetDeployment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/dialogflow/cx/v3beta1/deployment.proto",
}
