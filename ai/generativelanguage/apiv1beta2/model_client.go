// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package generativelanguage

import (
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"

	generativelanguagepb "cloud.google.com/go/ai/generativelanguage/apiv1beta2/generativelanguagepb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newModelClientHook clientHook

// ModelCallOptions contains the retry settings for each method of ModelClient.
type ModelCallOptions struct {
	GetModel   []gax.CallOption
	ListModels []gax.CallOption
}

func defaultModelGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("generativelanguage.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("generativelanguage.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://generativelanguage.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultModelCallOptions() *ModelCallOptions {
	return &ModelCallOptions{
		GetModel:   []gax.CallOption{},
		ListModels: []gax.CallOption{},
	}
}

func defaultModelRESTCallOptions() *ModelCallOptions {
	return &ModelCallOptions{
		GetModel:   []gax.CallOption{},
		ListModels: []gax.CallOption{},
	}
}

// internalModelClient is an interface that defines the methods available from Generative Language API.
type internalModelClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetModel(context.Context, *generativelanguagepb.GetModelRequest, ...gax.CallOption) (*generativelanguagepb.Model, error)
	ListModels(context.Context, *generativelanguagepb.ListModelsRequest, ...gax.CallOption) *ModelIterator
}

// ModelClient is a client for interacting with Generative Language API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Provides methods for getting metadata information about Generative Models.
type ModelClient struct {
	// The internal transport-dependent client.
	internalClient internalModelClient

	// The call options for this service.
	CallOptions *ModelCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ModelClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ModelClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *ModelClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetModel gets information about a specific Model.
func (c *ModelClient) GetModel(ctx context.Context, req *generativelanguagepb.GetModelRequest, opts ...gax.CallOption) (*generativelanguagepb.Model, error) {
	return c.internalClient.GetModel(ctx, req, opts...)
}

// ListModels lists models available through the API.
func (c *ModelClient) ListModels(ctx context.Context, req *generativelanguagepb.ListModelsRequest, opts ...gax.CallOption) *ModelIterator {
	return c.internalClient.ListModels(ctx, req, opts...)
}

// modelGRPCClient is a client for interacting with Generative Language API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type modelGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing ModelClient
	CallOptions **ModelCallOptions

	// The gRPC API client.
	modelClient generativelanguagepb.ModelServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewModelClient creates a new model service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Provides methods for getting metadata information about Generative Models.
func NewModelClient(ctx context.Context, opts ...option.ClientOption) (*ModelClient, error) {
	clientOpts := defaultModelGRPCClientOptions()
	if newModelClientHook != nil {
		hookOpts, err := newModelClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ModelClient{CallOptions: defaultModelCallOptions()}

	c := &modelGRPCClient{
		connPool:    connPool,
		modelClient: generativelanguagepb.NewModelServiceClient(connPool),
		CallOptions: &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *modelGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *modelGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *modelGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type modelRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing ModelClient
	CallOptions **ModelCallOptions
}

// NewModelRESTClient creates a new model service rest client.
//
// Provides methods for getting metadata information about Generative Models.
func NewModelRESTClient(ctx context.Context, opts ...option.ClientOption) (*ModelClient, error) {
	clientOpts := append(defaultModelRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultModelRESTCallOptions()
	c := &modelRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &ModelClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultModelRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://generativelanguage.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://generativelanguage.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://generativelanguage.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *modelRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *modelRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *modelRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *modelGRPCClient) GetModel(ctx context.Context, req *generativelanguagepb.GetModelRequest, opts ...gax.CallOption) (*generativelanguagepb.Model, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetModel[0:len((*c.CallOptions).GetModel):len((*c.CallOptions).GetModel)], opts...)
	var resp *generativelanguagepb.Model
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.modelClient.GetModel(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *modelGRPCClient) ListModels(ctx context.Context, req *generativelanguagepb.ListModelsRequest, opts ...gax.CallOption) *ModelIterator {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).ListModels[0:len((*c.CallOptions).ListModels):len((*c.CallOptions).ListModels)], opts...)
	it := &ModelIterator{}
	req = proto.Clone(req).(*generativelanguagepb.ListModelsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*generativelanguagepb.Model, string, error) {
		resp := &generativelanguagepb.ListModelsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.modelClient.ListModels(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetModels(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// GetModel gets information about a specific Model.
func (c *modelRESTClient) GetModel(ctx context.Context, req *generativelanguagepb.GetModelRequest, opts ...gax.CallOption) (*generativelanguagepb.Model, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta2/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetModel[0:len((*c.CallOptions).GetModel):len((*c.CallOptions).GetModel)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &generativelanguagepb.Model{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// ListModels lists models available through the API.
func (c *modelRESTClient) ListModels(ctx context.Context, req *generativelanguagepb.ListModelsRequest, opts ...gax.CallOption) *ModelIterator {
	it := &ModelIterator{}
	req = proto.Clone(req).(*generativelanguagepb.ListModelsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*generativelanguagepb.Model, string, error) {
		resp := &generativelanguagepb.ListModelsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/v1beta2/models")

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		hds := append(c.xGoogHeaders, "Content-Type", "application/json")
		headers := gax.BuildHeaders(ctx, hds...)
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			httpRsp, err := c.httpClient.Do(httpReq)
			if err != nil {
				return err
			}
			defer httpRsp.Body.Close()

			if err = googleapi.CheckResponse(httpRsp); err != nil {
				return err
			}

			buf, err := io.ReadAll(httpRsp.Body)
			if err != nil {
				return err
			}

			if err := unm.Unmarshal(buf, resp); err != nil {
				return err
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetModels(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// ModelIterator manages a stream of *generativelanguagepb.Model.
type ModelIterator struct {
	items    []*generativelanguagepb.Model
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*generativelanguagepb.Model, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ModelIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ModelIterator) Next() (*generativelanguagepb.Model, error) {
	var item *generativelanguagepb.Model
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ModelIterator) bufLen() int {
	return len(it.items)
}

func (it *ModelIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
