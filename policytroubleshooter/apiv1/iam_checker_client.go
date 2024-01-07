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

package policytroubleshooter

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	policytroubleshooterpb "cloud.google.com/go/policytroubleshooter/apiv1/policytroubleshooterpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var newIamCheckerClientHook clientHook

// IamCheckerCallOptions contains the retry settings for each method of IamCheckerClient.
type IamCheckerCallOptions struct {
	TroubleshootIamPolicy []gax.CallOption
}

func defaultIamCheckerGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("policytroubleshooter.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("policytroubleshooter.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://policytroubleshooter.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultIamCheckerCallOptions() *IamCheckerCallOptions {
	return &IamCheckerCallOptions{
		TroubleshootIamPolicy: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

func defaultIamCheckerRESTCallOptions() *IamCheckerCallOptions {
	return &IamCheckerCallOptions{
		TroubleshootIamPolicy: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

// internalIamCheckerClient is an interface that defines the methods available from Policy Troubleshooter API.
type internalIamCheckerClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	TroubleshootIamPolicy(context.Context, *policytroubleshooterpb.TroubleshootIamPolicyRequest, ...gax.CallOption) (*policytroubleshooterpb.TroubleshootIamPolicyResponse, error)
}

// IamCheckerClient is a client for interacting with Policy Troubleshooter API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// IAM Policy Troubleshooter service.
//
// This service helps you troubleshoot access issues for Google Cloud resources.
type IamCheckerClient struct {
	// The internal transport-dependent client.
	internalClient internalIamCheckerClient

	// The call options for this service.
	CallOptions *IamCheckerCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *IamCheckerClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *IamCheckerClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *IamCheckerClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// TroubleshootIamPolicy checks whether a principal has a specific permission for a specific
// resource, and explains why the principal does or does not have that
// permission.
func (c *IamCheckerClient) TroubleshootIamPolicy(ctx context.Context, req *policytroubleshooterpb.TroubleshootIamPolicyRequest, opts ...gax.CallOption) (*policytroubleshooterpb.TroubleshootIamPolicyResponse, error) {
	return c.internalClient.TroubleshootIamPolicy(ctx, req, opts...)
}

// iamCheckerGRPCClient is a client for interacting with Policy Troubleshooter API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type iamCheckerGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing IamCheckerClient
	CallOptions **IamCheckerCallOptions

	// The gRPC API client.
	iamCheckerClient policytroubleshooterpb.IamCheckerClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewIamCheckerClient creates a new iam checker client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// IAM Policy Troubleshooter service.
//
// This service helps you troubleshoot access issues for Google Cloud resources.
func NewIamCheckerClient(ctx context.Context, opts ...option.ClientOption) (*IamCheckerClient, error) {
	clientOpts := defaultIamCheckerGRPCClientOptions()
	if newIamCheckerClientHook != nil {
		hookOpts, err := newIamCheckerClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := IamCheckerClient{CallOptions: defaultIamCheckerCallOptions()}

	c := &iamCheckerGRPCClient{
		connPool:         connPool,
		iamCheckerClient: policytroubleshooterpb.NewIamCheckerClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *iamCheckerGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *iamCheckerGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *iamCheckerGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type iamCheckerRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing IamCheckerClient
	CallOptions **IamCheckerCallOptions
}

// NewIamCheckerRESTClient creates a new iam checker rest client.
//
// IAM Policy Troubleshooter service.
//
// This service helps you troubleshoot access issues for Google Cloud resources.
func NewIamCheckerRESTClient(ctx context.Context, opts ...option.ClientOption) (*IamCheckerClient, error) {
	clientOpts := append(defaultIamCheckerRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultIamCheckerRESTCallOptions()
	c := &iamCheckerRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &IamCheckerClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultIamCheckerRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://policytroubleshooter.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://policytroubleshooter.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://policytroubleshooter.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *iamCheckerRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *iamCheckerRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *iamCheckerRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *iamCheckerGRPCClient) TroubleshootIamPolicy(ctx context.Context, req *policytroubleshooterpb.TroubleshootIamPolicyRequest, opts ...gax.CallOption) (*policytroubleshooterpb.TroubleshootIamPolicyResponse, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).TroubleshootIamPolicy[0:len((*c.CallOptions).TroubleshootIamPolicy):len((*c.CallOptions).TroubleshootIamPolicy)], opts...)
	var resp *policytroubleshooterpb.TroubleshootIamPolicyResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.iamCheckerClient.TroubleshootIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// TroubleshootIamPolicy checks whether a principal has a specific permission for a specific
// resource, and explains why the principal does or does not have that
// permission.
func (c *iamCheckerRESTClient) TroubleshootIamPolicy(ctx context.Context, req *policytroubleshooterpb.TroubleshootIamPolicyRequest, opts ...gax.CallOption) (*policytroubleshooterpb.TroubleshootIamPolicyResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/iam:troubleshoot")

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := append(c.xGoogHeaders, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).TroubleshootIamPolicy[0:len((*c.CallOptions).TroubleshootIamPolicy):len((*c.CallOptions).TroubleshootIamPolicy)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &policytroubleshooterpb.TroubleshootIamPolicyResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
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
