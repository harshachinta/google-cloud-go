// Copyright 2024 Google LLC
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

package inventory

import (
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	kmspb "cloud.google.com/go/kms/apiv1/kmspb"
	inventorypb "cloud.google.com/go/kms/inventory/apiv1/inventorypb"
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

var newKeyDashboardClientHook clientHook

// KeyDashboardCallOptions contains the retry settings for each method of KeyDashboardClient.
type KeyDashboardCallOptions struct {
	ListCryptoKeys []gax.CallOption
}

func defaultKeyDashboardGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("kmsinventory.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("kmsinventory.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("kmsinventory.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://kmsinventory.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultKeyDashboardCallOptions() *KeyDashboardCallOptions {
	return &KeyDashboardCallOptions{
		ListCryptoKeys: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

func defaultKeyDashboardRESTCallOptions() *KeyDashboardCallOptions {
	return &KeyDashboardCallOptions{
		ListCryptoKeys: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

// internalKeyDashboardClient is an interface that defines the methods available from KMS Inventory API.
type internalKeyDashboardClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListCryptoKeys(context.Context, *inventorypb.ListCryptoKeysRequest, ...gax.CallOption) *CryptoKeyIterator
}

// KeyDashboardClient is a client for interacting with KMS Inventory API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Provides a cross-region view of all Cloud KMS keys in a given Cloud project.
type KeyDashboardClient struct {
	// The internal transport-dependent client.
	internalClient internalKeyDashboardClient

	// The call options for this service.
	CallOptions *KeyDashboardCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *KeyDashboardClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *KeyDashboardClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *KeyDashboardClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListCryptoKeys returns cryptographic keys managed by Cloud KMS in a given Cloud project.
// Note that this data is sourced from snapshots, meaning it may not
// completely reflect the actual state of key metadata at call time.
func (c *KeyDashboardClient) ListCryptoKeys(ctx context.Context, req *inventorypb.ListCryptoKeysRequest, opts ...gax.CallOption) *CryptoKeyIterator {
	return c.internalClient.ListCryptoKeys(ctx, req, opts...)
}

// keyDashboardGRPCClient is a client for interacting with KMS Inventory API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type keyDashboardGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing KeyDashboardClient
	CallOptions **KeyDashboardCallOptions

	// The gRPC API client.
	keyDashboardClient inventorypb.KeyDashboardServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewKeyDashboardClient creates a new key dashboard service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Provides a cross-region view of all Cloud KMS keys in a given Cloud project.
func NewKeyDashboardClient(ctx context.Context, opts ...option.ClientOption) (*KeyDashboardClient, error) {
	clientOpts := defaultKeyDashboardGRPCClientOptions()
	if newKeyDashboardClientHook != nil {
		hookOpts, err := newKeyDashboardClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := KeyDashboardClient{CallOptions: defaultKeyDashboardCallOptions()}

	c := &keyDashboardGRPCClient{
		connPool:           connPool,
		keyDashboardClient: inventorypb.NewKeyDashboardServiceClient(connPool),
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
func (c *keyDashboardGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *keyDashboardGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *keyDashboardGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type keyDashboardRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing KeyDashboardClient
	CallOptions **KeyDashboardCallOptions
}

// NewKeyDashboardRESTClient creates a new key dashboard service rest client.
//
// Provides a cross-region view of all Cloud KMS keys in a given Cloud project.
func NewKeyDashboardRESTClient(ctx context.Context, opts ...option.ClientOption) (*KeyDashboardClient, error) {
	clientOpts := append(defaultKeyDashboardRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultKeyDashboardRESTCallOptions()
	c := &keyDashboardRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &KeyDashboardClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultKeyDashboardRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://kmsinventory.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://kmsinventory.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://kmsinventory.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://kmsinventory.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *keyDashboardRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *keyDashboardRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *keyDashboardRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *keyDashboardGRPCClient) ListCryptoKeys(ctx context.Context, req *inventorypb.ListCryptoKeysRequest, opts ...gax.CallOption) *CryptoKeyIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListCryptoKeys[0:len((*c.CallOptions).ListCryptoKeys):len((*c.CallOptions).ListCryptoKeys)], opts...)
	it := &CryptoKeyIterator{}
	req = proto.Clone(req).(*inventorypb.ListCryptoKeysRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*kmspb.CryptoKey, string, error) {
		resp := &inventorypb.ListCryptoKeysResponse{}
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
			resp, err = c.keyDashboardClient.ListCryptoKeys(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetCryptoKeys(), resp.GetNextPageToken(), nil
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

// ListCryptoKeys returns cryptographic keys managed by Cloud KMS in a given Cloud project.
// Note that this data is sourced from snapshots, meaning it may not
// completely reflect the actual state of key metadata at call time.
func (c *keyDashboardRESTClient) ListCryptoKeys(ctx context.Context, req *inventorypb.ListCryptoKeysRequest, opts ...gax.CallOption) *CryptoKeyIterator {
	it := &CryptoKeyIterator{}
	req = proto.Clone(req).(*inventorypb.ListCryptoKeysRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*kmspb.CryptoKey, string, error) {
		resp := &inventorypb.ListCryptoKeysResponse{}
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
		baseUrl.Path += fmt.Sprintf("/v1/%v/cryptoKeys", req.GetParent())

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
		return resp.GetCryptoKeys(), resp.GetNextPageToken(), nil
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
