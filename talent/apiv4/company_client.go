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

package talent

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	talentpb "cloud.google.com/go/talent/apiv4/talentpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newCompanyClientHook clientHook

// CompanyCallOptions contains the retry settings for each method of CompanyClient.
type CompanyCallOptions struct {
	CreateCompany []gax.CallOption
	GetCompany    []gax.CallOption
	UpdateCompany []gax.CallOption
	DeleteCompany []gax.CallOption
	ListCompanies []gax.CallOption
	GetOperation  []gax.CallOption
}

func defaultCompanyGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("jobs.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("jobs.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://jobs.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCompanyCallOptions() *CompanyCallOptions {
	return &CompanyCallOptions{
		CreateCompany: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
		},
		GetCompany: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateCompany: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
		},
		DeleteCompany: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListCompanies: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetOperation: []gax.CallOption{},
	}
}

func defaultCompanyRESTCallOptions() *CompanyCallOptions {
	return &CompanyCallOptions{
		CreateCompany: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
		},
		GetCompany: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable)
			}),
		},
		UpdateCompany: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
		},
		DeleteCompany: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable)
			}),
		},
		ListCompanies: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable)
			}),
		},
		GetOperation: []gax.CallOption{},
	}
}

// internalCompanyClient is an interface that defines the methods available from Cloud Talent Solution API.
type internalCompanyClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateCompany(context.Context, *talentpb.CreateCompanyRequest, ...gax.CallOption) (*talentpb.Company, error)
	GetCompany(context.Context, *talentpb.GetCompanyRequest, ...gax.CallOption) (*talentpb.Company, error)
	UpdateCompany(context.Context, *talentpb.UpdateCompanyRequest, ...gax.CallOption) (*talentpb.Company, error)
	DeleteCompany(context.Context, *talentpb.DeleteCompanyRequest, ...gax.CallOption) error
	ListCompanies(context.Context, *talentpb.ListCompaniesRequest, ...gax.CallOption) *CompanyIterator
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
}

// CompanyClient is a client for interacting with Cloud Talent Solution API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service that handles company management, including CRUD and enumeration.
type CompanyClient struct {
	// The internal transport-dependent client.
	internalClient internalCompanyClient

	// The call options for this service.
	CallOptions *CompanyCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CompanyClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CompanyClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *CompanyClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateCompany creates a new company entity.
func (c *CompanyClient) CreateCompany(ctx context.Context, req *talentpb.CreateCompanyRequest, opts ...gax.CallOption) (*talentpb.Company, error) {
	return c.internalClient.CreateCompany(ctx, req, opts...)
}

// GetCompany retrieves specified company.
func (c *CompanyClient) GetCompany(ctx context.Context, req *talentpb.GetCompanyRequest, opts ...gax.CallOption) (*talentpb.Company, error) {
	return c.internalClient.GetCompany(ctx, req, opts...)
}

// UpdateCompany updates specified company.
func (c *CompanyClient) UpdateCompany(ctx context.Context, req *talentpb.UpdateCompanyRequest, opts ...gax.CallOption) (*talentpb.Company, error) {
	return c.internalClient.UpdateCompany(ctx, req, opts...)
}

// DeleteCompany deletes specified company.
// Prerequisite: The company has no jobs associated with it.
func (c *CompanyClient) DeleteCompany(ctx context.Context, req *talentpb.DeleteCompanyRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteCompany(ctx, req, opts...)
}

// ListCompanies lists all companies associated with the project.
func (c *CompanyClient) ListCompanies(ctx context.Context, req *talentpb.ListCompaniesRequest, opts ...gax.CallOption) *CompanyIterator {
	return c.internalClient.ListCompanies(ctx, req, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *CompanyClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// companyGRPCClient is a client for interacting with Cloud Talent Solution API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type companyGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing CompanyClient
	CallOptions **CompanyCallOptions

	// The gRPC API client.
	companyClient talentpb.CompanyServiceClient

	operationsClient longrunningpb.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewCompanyClient creates a new company service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service that handles company management, including CRUD and enumeration.
func NewCompanyClient(ctx context.Context, opts ...option.ClientOption) (*CompanyClient, error) {
	clientOpts := defaultCompanyGRPCClientOptions()
	if newCompanyClientHook != nil {
		hookOpts, err := newCompanyClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := CompanyClient{CallOptions: defaultCompanyCallOptions()}

	c := &companyGRPCClient{
		connPool:         connPool,
		companyClient:    talentpb.NewCompanyServiceClient(connPool),
		CallOptions:      &client.CallOptions,
		operationsClient: longrunningpb.NewOperationsClient(connPool),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *companyGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *companyGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *companyGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type companyRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing CompanyClient
	CallOptions **CompanyCallOptions
}

// NewCompanyRESTClient creates a new company service rest client.
//
// A service that handles company management, including CRUD and enumeration.
func NewCompanyRESTClient(ctx context.Context, opts ...option.ClientOption) (*CompanyClient, error) {
	clientOpts := append(defaultCompanyRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultCompanyRESTCallOptions()
	c := &companyRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &CompanyClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultCompanyRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://jobs.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://jobs.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://jobs.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *companyRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *companyRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *companyRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *companyGRPCClient) CreateCompany(ctx context.Context, req *talentpb.CreateCompanyRequest, opts ...gax.CallOption) (*talentpb.Company, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateCompany[0:len((*c.CallOptions).CreateCompany):len((*c.CallOptions).CreateCompany)], opts...)
	var resp *talentpb.Company
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.companyClient.CreateCompany(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *companyGRPCClient) GetCompany(ctx context.Context, req *talentpb.GetCompanyRequest, opts ...gax.CallOption) (*talentpb.Company, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetCompany[0:len((*c.CallOptions).GetCompany):len((*c.CallOptions).GetCompany)], opts...)
	var resp *talentpb.Company
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.companyClient.GetCompany(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *companyGRPCClient) UpdateCompany(ctx context.Context, req *talentpb.UpdateCompanyRequest, opts ...gax.CallOption) (*talentpb.Company, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "company.name", url.QueryEscape(req.GetCompany().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateCompany[0:len((*c.CallOptions).UpdateCompany):len((*c.CallOptions).UpdateCompany)], opts...)
	var resp *talentpb.Company
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.companyClient.UpdateCompany(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *companyGRPCClient) DeleteCompany(ctx context.Context, req *talentpb.DeleteCompanyRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteCompany[0:len((*c.CallOptions).DeleteCompany):len((*c.CallOptions).DeleteCompany)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.companyClient.DeleteCompany(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *companyGRPCClient) ListCompanies(ctx context.Context, req *talentpb.ListCompaniesRequest, opts ...gax.CallOption) *CompanyIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListCompanies[0:len((*c.CallOptions).ListCompanies):len((*c.CallOptions).ListCompanies)], opts...)
	it := &CompanyIterator{}
	req = proto.Clone(req).(*talentpb.ListCompaniesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*talentpb.Company, string, error) {
		resp := &talentpb.ListCompaniesResponse{}
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
			resp, err = c.companyClient.ListCompanies(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetCompanies(), resp.GetNextPageToken(), nil
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

func (c *companyGRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.operationsClient.GetOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateCompany creates a new company entity.
func (c *companyRESTClient) CreateCompany(ctx context.Context, req *talentpb.CreateCompanyRequest, opts ...gax.CallOption) (*talentpb.Company, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetCompany()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v4/%v/companies", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).CreateCompany[0:len((*c.CallOptions).CreateCompany):len((*c.CallOptions).CreateCompany)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &talentpb.Company{}
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

// GetCompany retrieves specified company.
func (c *companyRESTClient) GetCompany(ctx context.Context, req *talentpb.GetCompanyRequest, opts ...gax.CallOption) (*talentpb.Company, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v4/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).GetCompany[0:len((*c.CallOptions).GetCompany):len((*c.CallOptions).GetCompany)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &talentpb.Company{}
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

// UpdateCompany updates specified company.
func (c *companyRESTClient) UpdateCompany(ctx context.Context, req *talentpb.UpdateCompanyRequest, opts ...gax.CallOption) (*talentpb.Company, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetCompany()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v4/%v", req.GetCompany().GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetUpdateMask() != nil {
		updateMask, err := protojson.Marshal(req.GetUpdateMask())
		if err != nil {
			return nil, err
		}
		params.Add("updateMask", string(updateMask[1:len(updateMask)-1]))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "company.name", url.QueryEscape(req.GetCompany().GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).UpdateCompany[0:len((*c.CallOptions).UpdateCompany):len((*c.CallOptions).UpdateCompany)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &talentpb.Company{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
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

// DeleteCompany deletes specified company.
// Prerequisite: The company has no jobs associated with it.
func (c *companyRESTClient) DeleteCompany(ctx context.Context, req *talentpb.DeleteCompanyRequest, opts ...gax.CallOption) error {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return err
	}
	baseUrl.Path += fmt.Sprintf("/v4/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	return gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
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

		// Returns nil if there is no error, otherwise wraps
		// the response code and body into a non-nil error
		return googleapi.CheckResponse(httpRsp)
	}, opts...)
}

// ListCompanies lists all companies associated with the project.
func (c *companyRESTClient) ListCompanies(ctx context.Context, req *talentpb.ListCompaniesRequest, opts ...gax.CallOption) *CompanyIterator {
	it := &CompanyIterator{}
	req = proto.Clone(req).(*talentpb.ListCompaniesRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*talentpb.Company, string, error) {
		resp := &talentpb.ListCompaniesResponse{}
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
		baseUrl.Path += fmt.Sprintf("/v4/%v/companies", req.GetParent())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req.GetRequireOpenJobs() {
			params.Add("requireOpenJobs", fmt.Sprintf("%v", req.GetRequireOpenJobs()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
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
		return resp.GetCompanies(), resp.GetNextPageToken(), nil
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

// GetOperation is a utility method from google.longrunning.Operations.
func (c *companyRESTClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v4/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
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

// CompanyIterator manages a stream of *talentpb.Company.
type CompanyIterator struct {
	items    []*talentpb.Company
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
	InternalFetch func(pageSize int, pageToken string) (results []*talentpb.Company, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *CompanyIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *CompanyIterator) Next() (*talentpb.Company, error) {
	var item *talentpb.Company
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *CompanyIterator) bufLen() int {
	return len(it.items)
}

func (it *CompanyIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
