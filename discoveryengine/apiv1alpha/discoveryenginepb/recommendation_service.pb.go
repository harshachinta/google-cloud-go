// Copyright 2022 Google LLC
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
// source: google/cloud/discoveryengine/v1alpha/recommendation_service.proto

package discoveryenginepb

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

// Request message for Recommend method.
type RecommendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Full resource name of a
	// [ServingConfig][google.cloud.discoveryengine.v1alpha.ServingConfig]:
	// `projects/*/locations/global/collections/*/engines/*/servingConfigs/*`, or
	// `projects/*/locations/global/collections/*/dataStores/*/servingConfigs/*`
	//
	// One default serving config is created along with your recommendation engine
	// creation. The engine ID will be used as the ID of the default serving
	// config. For example, for Engine
	// `projects/*/locations/global/collections/*/engines/my-engine`, you can use
	// `projects/*/locations/global/collections/*/engines/my-engine/servingConfigs/my-engine`
	// for your
	// [RecommendationService.Recommend][google.cloud.discoveryengine.v1alpha.RecommendationService.Recommend]
	// requests.
	ServingConfig string `protobuf:"bytes,1,opt,name=serving_config,json=servingConfig,proto3" json:"serving_config,omitempty"`
	// Required. Context about the user, what they are looking at and what action
	// they took to trigger the Recommend request. Note that this user event
	// detail won't be ingested to userEvent logs. Thus, a separate userEvent
	// write request is required for event logging.
	//
	// Don't set
	// [UserEvent.user_pseudo_id][google.cloud.discoveryengine.v1alpha.UserEvent.user_pseudo_id]
	// or
	// [UserEvent.user_info.user_id][google.cloud.discoveryengine.v1alpha.UserInfo.user_id]
	// to the same fixed ID for different users. If you are trying to receive
	// non-personalized recommendations (not recommended; this can negatively
	// impact model performance), instead set
	// [UserEvent.user_pseudo_id][google.cloud.discoveryengine.v1alpha.UserEvent.user_pseudo_id]
	// to a random unique ID and leave
	// [UserEvent.user_info.user_id][google.cloud.discoveryengine.v1alpha.UserInfo.user_id]
	// unset.
	UserEvent *UserEvent `protobuf:"bytes,2,opt,name=user_event,json=userEvent,proto3" json:"user_event,omitempty"`
	// Maximum number of results to return. Set this property
	// to the number of recommendation results needed. If zero, the service will
	// choose a reasonable default. The maximum allowed value is 100. Values
	// above 100 will be coerced to 100.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Filter for restricting recommendation results with a length limit of 5,000
	// characters. Currently, only filter expressions on the `filter_tags`
	// attribute is supported.
	//
	// Examples:
	//
	//   - `(filter_tags: ANY("Red", "Blue") OR filter_tags: ANY("Hot", "Cold"))`
	//   - `(filter_tags: ANY("Red", "Blue")) AND NOT (filter_tags: ANY("Green"))`
	//
	// If `attributeFilteringSyntax` is set to true under the `params` field, then
	// attribute-based expressions are expected instead of the above described
	// tag-based syntax. Examples:
	//
	//   - (launguage: ANY("en", "es")) AND NOT (categories: ANY("Movie"))
	//   - (available: true) AND
	//     (launguage: ANY("en", "es")) OR (categories: ANY("Movie"))
	//
	// If your filter blocks all results, the API will return generic
	// (unfiltered) popular Documents. If you only want results strictly matching
	// the filters, set `strictFiltering` to True in
	// [RecommendRequest.params][google.cloud.discoveryengine.v1alpha.RecommendRequest.params]
	// to receive empty results instead.
	//
	// Note that the API will never return
	// [Document][google.cloud.discoveryengine.v1alpha.Document]s with
	// `storageStatus` of `EXPIRED` or `DELETED` regardless of filter choices.
	Filter string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	// Use validate only mode for this recommendation query. If set to true, a
	// fake model will be used that returns arbitrary Document IDs.
	// Note that the validate only mode should only be used for testing the API,
	// or if the model is not ready.
	ValidateOnly bool `protobuf:"varint,5,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	// Additional domain specific parameters for the recommendations.
	//
	// Allowed values:
	//
	//   - `returnDocument`: Boolean. If set to true, the associated Document
	//     object will be returned in
	//     [RecommendResponse.RecommendationResult.document][google.cloud.discoveryengine.v1alpha.RecommendResponse.RecommendationResult.document].
	//   - `returnScore`: Boolean. If set to true, the recommendation 'score'
	//     corresponding to each returned Document will be set in
	//     [RecommendResponse.RecommendationResult.metadata][google.cloud.discoveryengine.v1alpha.RecommendResponse.RecommendationResult.metadata].
	//     The given 'score' indicates the probability of a Document conversion
	//     given the user's context and history.
	//   - `strictFiltering`: Boolean. True by default. If set to false, the service
	//     will return generic (unfiltered) popular Documents instead of empty if
	//     your filter blocks all recommendation results.
	//   - `diversityLevel`: String. Default empty. If set to be non-empty, then
	//     it needs to be one of:
	//   - `no-diversity`
	//   - `low-diversity`
	//   - `medium-diversity`
	//   - `high-diversity`
	//   - `auto-diversity`
	//     This gives request-level control and adjusts recommendation results
	//     based on Document category.
	//   - `attributeFilteringSyntax`: Boolean. False by default. If set to true,
	//     the `filter` field is interpreted according to the new,
	//     attribute-based syntax.
	Params map[string]*structpb.Value `protobuf:"bytes,6,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The user labels applied to a resource must meet the following requirements:
	//
	//   - Each resource can have multiple labels, up to a maximum of 64.
	//   - Each label must be a key-value pair.
	//   - Keys have a minimum length of 1 character and a maximum length of 63
	//     characters and cannot be empty. Values can be empty and have a maximum
	//     length of 63 characters.
	//   - Keys and values can contain only lowercase letters, numeric characters,
	//     underscores, and dashes. All characters must use UTF-8 encoding, and
	//     international characters are allowed.
	//   - The key portion of a label must be unique. However, you can use the same
	//     key with multiple resources.
	//   - Keys must start with a lowercase letter or international character.
	//
	// See [Requirements for
	// labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements)
	// for more details.
	UserLabels map[string]string `protobuf:"bytes,8,rep,name=user_labels,json=userLabels,proto3" json:"user_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RecommendRequest) Reset() {
	*x = RecommendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendRequest) ProtoMessage() {}

func (x *RecommendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendRequest.ProtoReflect.Descriptor instead.
func (*RecommendRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_rawDescGZIP(), []int{0}
}

func (x *RecommendRequest) GetServingConfig() string {
	if x != nil {
		return x.ServingConfig
	}
	return ""
}

func (x *RecommendRequest) GetUserEvent() *UserEvent {
	if x != nil {
		return x.UserEvent
	}
	return nil
}

func (x *RecommendRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *RecommendRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *RecommendRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

func (x *RecommendRequest) GetParams() map[string]*structpb.Value {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *RecommendRequest) GetUserLabels() map[string]string {
	if x != nil {
		return x.UserLabels
	}
	return nil
}

// Response message for Recommend method.
type RecommendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of recommended Documents. The order represents the ranking (from the
	// most relevant Document to the least).
	Results []*RecommendResponse_RecommendationResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	// A unique attribution token. This should be included in the
	// [UserEvent][google.cloud.discoveryengine.v1alpha.UserEvent] logs resulting
	// from this recommendation, which enables accurate attribution of
	// recommendation model performance.
	AttributionToken string `protobuf:"bytes,2,opt,name=attribution_token,json=attributionToken,proto3" json:"attribution_token,omitempty"`
	// IDs of documents in the request that were missing from the default Branch
	// associated with the requested ServingConfig.
	MissingIds []string `protobuf:"bytes,3,rep,name=missing_ids,json=missingIds,proto3" json:"missing_ids,omitempty"`
	// True if
	// [RecommendRequest.validate_only][google.cloud.discoveryengine.v1alpha.RecommendRequest.validate_only]
	// was set.
	ValidateOnly bool `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
}

func (x *RecommendResponse) Reset() {
	*x = RecommendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendResponse) ProtoMessage() {}

func (x *RecommendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendResponse.ProtoReflect.Descriptor instead.
func (*RecommendResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_rawDescGZIP(), []int{1}
}

func (x *RecommendResponse) GetResults() []*RecommendResponse_RecommendationResult {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *RecommendResponse) GetAttributionToken() string {
	if x != nil {
		return x.AttributionToken
	}
	return ""
}

func (x *RecommendResponse) GetMissingIds() []string {
	if x != nil {
		return x.MissingIds
	}
	return nil
}

func (x *RecommendResponse) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

// RecommendationResult represents a generic recommendation result with
// associated metadata.
type RecommendResponse_RecommendationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource ID of the recommended Document.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Set if `returnDocument` is set to true in
	// [RecommendRequest.params][google.cloud.discoveryengine.v1alpha.RecommendRequest.params].
	Document *Document `protobuf:"bytes,2,opt,name=document,proto3" json:"document,omitempty"`
	// Additional Document metadata / annotations.
	//
	// Possible values:
	//
	//   - `score`: Recommendation score in double value. Is set if
	//     `returnScore` is set to true in
	//     [RecommendRequest.params][google.cloud.discoveryengine.v1alpha.RecommendRequest.params].
	Metadata map[string]*structpb.Value `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RecommendResponse_RecommendationResult) Reset() {
	*x = RecommendResponse_RecommendationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommendResponse_RecommendationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendResponse_RecommendationResult) ProtoMessage() {}

func (x *RecommendResponse_RecommendationResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendResponse_RecommendationResult.ProtoReflect.Descriptor instead.
func (*RecommendResponse_RecommendationResult) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *RecommendResponse_RecommendationResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RecommendResponse_RecommendationResult) GetDocument() *Document {
	if x != nil {
		return x.Document
	}
	return nil
}

func (x *RecommendResponse_RecommendationResult) GetMetadata() map[string]*structpb.Value {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_google_cloud_discoveryengine_v1alpha_recommendation_service_proto protoreflect.FileDescriptor

var file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_rawDesc = []byte{
	0x0a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x04, 0x0a, 0x10, 0x52, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5b, 0x0a, 0x0e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x34, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2e, 0x0a, 0x2c, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x53, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f,
	0x6e, 0x6c, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x5a, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x67, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x51, 0x0a, 0x0b,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x3d, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb0,
	0x04, 0x0a, 0x11, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x2b, 0x0a, 0x11,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x1a,
	0xbf, 0x02, 0x0a, 0x14, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x4a, 0x0a, 0x08, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x76, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x5a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x53, 0x0a, 0x0d,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x32, 0xa8, 0x04, 0x0a, 0x15, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xba, 0x03, 0x0a, 0x09,
	0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x12, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xbb, 0x02, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0xb4, 0x02, 0x3a, 0x01, 0x2a, 0x5a, 0x6b, 0x3a, 0x01, 0x2a, 0x22, 0x66, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x72, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x5a, 0x68, 0x3a, 0x01, 0x2a, 0x22, 0x63, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x73,
	0x2f, 0x2a, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x22, 0x58,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x72,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x1a, 0x52, 0xca, 0x41, 0x1e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0xa6, 0x02, 0x0a,
	0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x1a, 0x52, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x52, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0x3b, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0xa2, 0x02, 0x0f, 0x44, 0x49,
	0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x59, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0xaa, 0x02, 0x24,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x31, 0x41,
	0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xea, 0x02, 0x27, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_rawDescOnce sync.Once
	file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_rawDescData = file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_rawDesc
)

func file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_rawDescGZIP() []byte {
	file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_rawDescData)
	})
	return file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_rawDescData
}

var file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_goTypes = []interface{}{
	(*RecommendRequest)(nil),  // 0: google.cloud.discoveryengine.v1alpha.RecommendRequest
	(*RecommendResponse)(nil), // 1: google.cloud.discoveryengine.v1alpha.RecommendResponse
	nil,                       // 2: google.cloud.discoveryengine.v1alpha.RecommendRequest.ParamsEntry
	nil,                       // 3: google.cloud.discoveryengine.v1alpha.RecommendRequest.UserLabelsEntry
	(*RecommendResponse_RecommendationResult)(nil), // 4: google.cloud.discoveryengine.v1alpha.RecommendResponse.RecommendationResult
	nil,                    // 5: google.cloud.discoveryengine.v1alpha.RecommendResponse.RecommendationResult.MetadataEntry
	(*UserEvent)(nil),      // 6: google.cloud.discoveryengine.v1alpha.UserEvent
	(*structpb.Value)(nil), // 7: google.protobuf.Value
	(*Document)(nil),       // 8: google.cloud.discoveryengine.v1alpha.Document
}
var file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_depIdxs = []int32{
	6, // 0: google.cloud.discoveryengine.v1alpha.RecommendRequest.user_event:type_name -> google.cloud.discoveryengine.v1alpha.UserEvent
	2, // 1: google.cloud.discoveryengine.v1alpha.RecommendRequest.params:type_name -> google.cloud.discoveryengine.v1alpha.RecommendRequest.ParamsEntry
	3, // 2: google.cloud.discoveryengine.v1alpha.RecommendRequest.user_labels:type_name -> google.cloud.discoveryengine.v1alpha.RecommendRequest.UserLabelsEntry
	4, // 3: google.cloud.discoveryengine.v1alpha.RecommendResponse.results:type_name -> google.cloud.discoveryengine.v1alpha.RecommendResponse.RecommendationResult
	7, // 4: google.cloud.discoveryengine.v1alpha.RecommendRequest.ParamsEntry.value:type_name -> google.protobuf.Value
	8, // 5: google.cloud.discoveryengine.v1alpha.RecommendResponse.RecommendationResult.document:type_name -> google.cloud.discoveryengine.v1alpha.Document
	5, // 6: google.cloud.discoveryengine.v1alpha.RecommendResponse.RecommendationResult.metadata:type_name -> google.cloud.discoveryengine.v1alpha.RecommendResponse.RecommendationResult.MetadataEntry
	7, // 7: google.cloud.discoveryengine.v1alpha.RecommendResponse.RecommendationResult.MetadataEntry.value:type_name -> google.protobuf.Value
	0, // 8: google.cloud.discoveryengine.v1alpha.RecommendationService.Recommend:input_type -> google.cloud.discoveryengine.v1alpha.RecommendRequest
	1, // 9: google.cloud.discoveryengine.v1alpha.RecommendationService.Recommend:output_type -> google.cloud.discoveryengine.v1alpha.RecommendResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_init() }
func file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_init() {
	if File_google_cloud_discoveryengine_v1alpha_recommendation_service_proto != nil {
		return
	}
	file_google_cloud_discoveryengine_v1alpha_document_proto_init()
	file_google_cloud_discoveryengine_v1alpha_user_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommendRequest); i {
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
		file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommendResponse); i {
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
		file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommendResponse_RecommendationResult); i {
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
			RawDescriptor: file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_depIdxs,
		MessageInfos:      file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_msgTypes,
	}.Build()
	File_google_cloud_discoveryengine_v1alpha_recommendation_service_proto = out.File
	file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_rawDesc = nil
	file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_goTypes = nil
	file_google_cloud_discoveryengine_v1alpha_recommendation_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RecommendationServiceClient is the client API for RecommendationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecommendationServiceClient interface {
	// Makes a recommendation, which requires a contextual user event.
	Recommend(ctx context.Context, in *RecommendRequest, opts ...grpc.CallOption) (*RecommendResponse, error)
}

type recommendationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecommendationServiceClient(cc grpc.ClientConnInterface) RecommendationServiceClient {
	return &recommendationServiceClient{cc}
}

func (c *recommendationServiceClient) Recommend(ctx context.Context, in *RecommendRequest, opts ...grpc.CallOption) (*RecommendResponse, error) {
	out := new(RecommendResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.discoveryengine.v1alpha.RecommendationService/Recommend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecommendationServiceServer is the server API for RecommendationService service.
type RecommendationServiceServer interface {
	// Makes a recommendation, which requires a contextual user event.
	Recommend(context.Context, *RecommendRequest) (*RecommendResponse, error)
}

// UnimplementedRecommendationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRecommendationServiceServer struct {
}

func (*UnimplementedRecommendationServiceServer) Recommend(context.Context, *RecommendRequest) (*RecommendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recommend not implemented")
}

func RegisterRecommendationServiceServer(s *grpc.Server, srv RecommendationServiceServer) {
	s.RegisterService(&_RecommendationService_serviceDesc, srv)
}

func _RecommendationService_Recommend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecommendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServiceServer).Recommend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.discoveryengine.v1alpha.RecommendationService/Recommend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServiceServer).Recommend(ctx, req.(*RecommendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RecommendationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.discoveryengine.v1alpha.RecommendationService",
	HandlerType: (*RecommendationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Recommend",
			Handler:    _RecommendationService_Recommend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/discoveryengine/v1alpha/recommendation_service.proto",
}
