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

package publicca

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	publiccapb "cloud.google.com/go/security/publicca/apiv1beta1/publiccapb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

var newPublicCertificateAuthorityClientHook clientHook

// PublicCertificateAuthorityCallOptions contains the retry settings for each method of PublicCertificateAuthorityClient.
type PublicCertificateAuthorityCallOptions struct {
	CreateExternalAccountKey []gax.CallOption
}

func defaultPublicCertificateAuthorityGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("publicca.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("publicca.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://publicca.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultPublicCertificateAuthorityCallOptions() *PublicCertificateAuthorityCallOptions {
	return &PublicCertificateAuthorityCallOptions{
		CreateExternalAccountKey: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

func defaultPublicCertificateAuthorityRESTCallOptions() *PublicCertificateAuthorityCallOptions {
	return &PublicCertificateAuthorityCallOptions{
		CreateExternalAccountKey: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
	}
}

// internalPublicCertificateAuthorityClient is an interface that defines the methods available from Public Certificate Authority API.
type internalPublicCertificateAuthorityClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateExternalAccountKey(context.Context, *publiccapb.CreateExternalAccountKeyRequest, ...gax.CallOption) (*publiccapb.ExternalAccountKey, error)
}

// PublicCertificateAuthorityClient is a client for interacting with Public Certificate Authority API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Manages the resources required for ACME external account
// binding (at https://tools.ietf.org/html/rfc8555#section-7.3.4) for
// the public certificate authority service.
type PublicCertificateAuthorityClient struct {
	// The internal transport-dependent client.
	internalClient internalPublicCertificateAuthorityClient

	// The call options for this service.
	CallOptions *PublicCertificateAuthorityCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *PublicCertificateAuthorityClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *PublicCertificateAuthorityClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *PublicCertificateAuthorityClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateExternalAccountKey creates a new ExternalAccountKey bound to the project.
func (c *PublicCertificateAuthorityClient) CreateExternalAccountKey(ctx context.Context, req *publiccapb.CreateExternalAccountKeyRequest, opts ...gax.CallOption) (*publiccapb.ExternalAccountKey, error) {
	return c.internalClient.CreateExternalAccountKey(ctx, req, opts...)
}

// publicCertificateAuthorityGRPCClient is a client for interacting with Public Certificate Authority API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type publicCertificateAuthorityGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing PublicCertificateAuthorityClient
	CallOptions **PublicCertificateAuthorityCallOptions

	// The gRPC API client.
	publicCertificateAuthorityClient publiccapb.PublicCertificateAuthorityServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewPublicCertificateAuthorityClient creates a new public certificate authority service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Manages the resources required for ACME external account
// binding (at https://tools.ietf.org/html/rfc8555#section-7.3.4) for
// the public certificate authority service.
func NewPublicCertificateAuthorityClient(ctx context.Context, opts ...option.ClientOption) (*PublicCertificateAuthorityClient, error) {
	clientOpts := defaultPublicCertificateAuthorityGRPCClientOptions()
	if newPublicCertificateAuthorityClientHook != nil {
		hookOpts, err := newPublicCertificateAuthorityClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := PublicCertificateAuthorityClient{CallOptions: defaultPublicCertificateAuthorityCallOptions()}

	c := &publicCertificateAuthorityGRPCClient{
		connPool:                         connPool,
		publicCertificateAuthorityClient: publiccapb.NewPublicCertificateAuthorityServiceClient(connPool),
		CallOptions:                      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *publicCertificateAuthorityGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *publicCertificateAuthorityGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *publicCertificateAuthorityGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type publicCertificateAuthorityRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing PublicCertificateAuthorityClient
	CallOptions **PublicCertificateAuthorityCallOptions
}

// NewPublicCertificateAuthorityRESTClient creates a new public certificate authority service rest client.
//
// Manages the resources required for ACME external account
// binding (at https://tools.ietf.org/html/rfc8555#section-7.3.4) for
// the public certificate authority service.
func NewPublicCertificateAuthorityRESTClient(ctx context.Context, opts ...option.ClientOption) (*PublicCertificateAuthorityClient, error) {
	clientOpts := append(defaultPublicCertificateAuthorityRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultPublicCertificateAuthorityRESTCallOptions()
	c := &publicCertificateAuthorityRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &PublicCertificateAuthorityClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultPublicCertificateAuthorityRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://publicca.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://publicca.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://publicca.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *publicCertificateAuthorityRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *publicCertificateAuthorityRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *publicCertificateAuthorityRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *publicCertificateAuthorityGRPCClient) CreateExternalAccountKey(ctx context.Context, req *publiccapb.CreateExternalAccountKeyRequest, opts ...gax.CallOption) (*publiccapb.ExternalAccountKey, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateExternalAccountKey[0:len((*c.CallOptions).CreateExternalAccountKey):len((*c.CallOptions).CreateExternalAccountKey)], opts...)
	var resp *publiccapb.ExternalAccountKey
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.publicCertificateAuthorityClient.CreateExternalAccountKey(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateExternalAccountKey creates a new ExternalAccountKey bound to the project.
func (c *publicCertificateAuthorityRESTClient) CreateExternalAccountKey(ctx context.Context, req *publiccapb.CreateExternalAccountKeyRequest, opts ...gax.CallOption) (*publiccapb.ExternalAccountKey, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetExternalAccountKey()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v/externalAccountKeys", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).CreateExternalAccountKey[0:len((*c.CallOptions).CreateExternalAccountKey):len((*c.CallOptions).CreateExternalAccountKey)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &publiccapb.ExternalAccountKey{}
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
