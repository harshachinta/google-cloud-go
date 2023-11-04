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

package billing

import (
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	billingpb "cloud.google.com/go/billing/apiv1/billingpb"
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

var newCloudCatalogClientHook clientHook

// CloudCatalogCallOptions contains the retry settings for each method of CloudCatalogClient.
type CloudCatalogCallOptions struct {
	ListServices []gax.CallOption
	ListSkus     []gax.CallOption
}

func defaultCloudCatalogGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("cloudbilling.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("cloudbilling.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://cloudbilling.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCloudCatalogCallOptions() *CloudCatalogCallOptions {
	return &CloudCatalogCallOptions{
		ListServices: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		ListSkus: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

func defaultCloudCatalogRESTCallOptions() *CloudCatalogCallOptions {
	return &CloudCatalogCallOptions{
		ListServices: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		ListSkus: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

// internalCloudCatalogClient is an interface that defines the methods available from Cloud Billing API.
type internalCloudCatalogClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListServices(context.Context, *billingpb.ListServicesRequest, ...gax.CallOption) *ServiceIterator
	ListSkus(context.Context, *billingpb.ListSkusRequest, ...gax.CallOption) *SkuIterator
}

// CloudCatalogClient is a client for interacting with Cloud Billing API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A catalog of Google Cloud Platform services and SKUs.
// Provides pricing information and metadata on Google Cloud Platform services
// and SKUs.
type CloudCatalogClient struct {
	// The internal transport-dependent client.
	internalClient internalCloudCatalogClient

	// The call options for this service.
	CallOptions *CloudCatalogCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CloudCatalogClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CloudCatalogClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *CloudCatalogClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListServices lists all public cloud services.
func (c *CloudCatalogClient) ListServices(ctx context.Context, req *billingpb.ListServicesRequest, opts ...gax.CallOption) *ServiceIterator {
	return c.internalClient.ListServices(ctx, req, opts...)
}

// ListSkus lists all publicly available SKUs for a given cloud service.
func (c *CloudCatalogClient) ListSkus(ctx context.Context, req *billingpb.ListSkusRequest, opts ...gax.CallOption) *SkuIterator {
	return c.internalClient.ListSkus(ctx, req, opts...)
}

// cloudCatalogGRPCClient is a client for interacting with Cloud Billing API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type cloudCatalogGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing CloudCatalogClient
	CallOptions **CloudCatalogCallOptions

	// The gRPC API client.
	cloudCatalogClient billingpb.CloudCatalogClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewCloudCatalogClient creates a new cloud catalog client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A catalog of Google Cloud Platform services and SKUs.
// Provides pricing information and metadata on Google Cloud Platform services
// and SKUs.
func NewCloudCatalogClient(ctx context.Context, opts ...option.ClientOption) (*CloudCatalogClient, error) {
	clientOpts := defaultCloudCatalogGRPCClientOptions()
	if newCloudCatalogClientHook != nil {
		hookOpts, err := newCloudCatalogClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := CloudCatalogClient{CallOptions: defaultCloudCatalogCallOptions()}

	c := &cloudCatalogGRPCClient{
		connPool:           connPool,
		cloudCatalogClient: billingpb.NewCloudCatalogClient(connPool),
		CallOptions:        &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *cloudCatalogGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *cloudCatalogGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *cloudCatalogGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type cloudCatalogRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing CloudCatalogClient
	CallOptions **CloudCatalogCallOptions
}

// NewCloudCatalogRESTClient creates a new cloud catalog rest client.
//
// A catalog of Google Cloud Platform services and SKUs.
// Provides pricing information and metadata on Google Cloud Platform services
// and SKUs.
func NewCloudCatalogRESTClient(ctx context.Context, opts ...option.ClientOption) (*CloudCatalogClient, error) {
	clientOpts := append(defaultCloudCatalogRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultCloudCatalogRESTCallOptions()
	c := &cloudCatalogRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &CloudCatalogClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultCloudCatalogRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://cloudbilling.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://cloudbilling.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://cloudbilling.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *cloudCatalogRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *cloudCatalogRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *cloudCatalogRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *cloudCatalogGRPCClient) ListServices(ctx context.Context, req *billingpb.ListServicesRequest, opts ...gax.CallOption) *ServiceIterator {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).ListServices[0:len((*c.CallOptions).ListServices):len((*c.CallOptions).ListServices)], opts...)
	it := &ServiceIterator{}
	req = proto.Clone(req).(*billingpb.ListServicesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*billingpb.Service, string, error) {
		resp := &billingpb.ListServicesResponse{}
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
			resp, err = c.cloudCatalogClient.ListServices(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetServices(), resp.GetNextPageToken(), nil
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

func (c *cloudCatalogGRPCClient) ListSkus(ctx context.Context, req *billingpb.ListSkusRequest, opts ...gax.CallOption) *SkuIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListSkus[0:len((*c.CallOptions).ListSkus):len((*c.CallOptions).ListSkus)], opts...)
	it := &SkuIterator{}
	req = proto.Clone(req).(*billingpb.ListSkusRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*billingpb.Sku, string, error) {
		resp := &billingpb.ListSkusResponse{}
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
			resp, err = c.cloudCatalogClient.ListSkus(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetSkus(), resp.GetNextPageToken(), nil
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

// ListServices lists all public cloud services.
func (c *cloudCatalogRESTClient) ListServices(ctx context.Context, req *billingpb.ListServicesRequest, opts ...gax.CallOption) *ServiceIterator {
	it := &ServiceIterator{}
	req = proto.Clone(req).(*billingpb.ListServicesRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*billingpb.Service, string, error) {
		resp := &billingpb.ListServicesResponse{}
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
		baseUrl.Path += fmt.Sprintf("/v1/services")

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
		return resp.GetServices(), resp.GetNextPageToken(), nil
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

// ListSkus lists all publicly available SKUs for a given cloud service.
func (c *cloudCatalogRESTClient) ListSkus(ctx context.Context, req *billingpb.ListSkusRequest, opts ...gax.CallOption) *SkuIterator {
	it := &SkuIterator{}
	req = proto.Clone(req).(*billingpb.ListSkusRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*billingpb.Sku, string, error) {
		resp := &billingpb.ListSkusResponse{}
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
		baseUrl.Path += fmt.Sprintf("/v1/%v/skus", req.GetParent())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetCurrencyCode() != "" {
			params.Add("currencyCode", fmt.Sprintf("%v", req.GetCurrencyCode()))
		}
		if req.GetEndTime() != nil {
			endTime, err := protojson.Marshal(req.GetEndTime())
			if err != nil {
				return nil, "", err
			}
			params.Add("endTime", string(endTime[1:len(endTime)-1]))
		}
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req.GetStartTime() != nil {
			startTime, err := protojson.Marshal(req.GetStartTime())
			if err != nil {
				return nil, "", err
			}
			params.Add("startTime", string(startTime[1:len(startTime)-1]))
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
		return resp.GetSkus(), resp.GetNextPageToken(), nil
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
