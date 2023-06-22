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

package metricsscope

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	metricsscopepb "cloud.google.com/go/monitoring/metricsscope/apiv1/metricsscopepb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newMetricsScopesClientHook clientHook

// MetricsScopesCallOptions contains the retry settings for each method of MetricsScopesClient.
type MetricsScopesCallOptions struct {
	GetMetricsScope                     []gax.CallOption
	ListMetricsScopesByMonitoredProject []gax.CallOption
	CreateMonitoredProject              []gax.CallOption
	DeleteMonitoredProject              []gax.CallOption
}

func defaultMetricsScopesGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("monitoring.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("monitoring.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://monitoring.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultMetricsScopesCallOptions() *MetricsScopesCallOptions {
	return &MetricsScopesCallOptions{
		GetMetricsScope:                     []gax.CallOption{},
		ListMetricsScopesByMonitoredProject: []gax.CallOption{},
		CreateMonitoredProject:              []gax.CallOption{},
		DeleteMonitoredProject:              []gax.CallOption{},
	}
}

// internalMetricsScopesClient is an interface that defines the methods available from Cloud Monitoring API.
type internalMetricsScopesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetMetricsScope(context.Context, *metricsscopepb.GetMetricsScopeRequest, ...gax.CallOption) (*metricsscopepb.MetricsScope, error)
	ListMetricsScopesByMonitoredProject(context.Context, *metricsscopepb.ListMetricsScopesByMonitoredProjectRequest, ...gax.CallOption) (*metricsscopepb.ListMetricsScopesByMonitoredProjectResponse, error)
	CreateMonitoredProject(context.Context, *metricsscopepb.CreateMonitoredProjectRequest, ...gax.CallOption) (*CreateMonitoredProjectOperation, error)
	CreateMonitoredProjectOperation(name string) *CreateMonitoredProjectOperation
	DeleteMonitoredProject(context.Context, *metricsscopepb.DeleteMonitoredProjectRequest, ...gax.CallOption) (*DeleteMonitoredProjectOperation, error)
	DeleteMonitoredProjectOperation(name string) *DeleteMonitoredProjectOperation
}

// MetricsScopesClient is a client for interacting with Cloud Monitoring API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Manages Cloud Monitoring Metrics Scopes, and the monitoring of Google Cloud
// projects and AWS accounts.
type MetricsScopesClient struct {
	// The internal transport-dependent client.
	internalClient internalMetricsScopesClient

	// The call options for this service.
	CallOptions *MetricsScopesCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *MetricsScopesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *MetricsScopesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *MetricsScopesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetMetricsScope returns a specific Metrics Scope.
func (c *MetricsScopesClient) GetMetricsScope(ctx context.Context, req *metricsscopepb.GetMetricsScopeRequest, opts ...gax.CallOption) (*metricsscopepb.MetricsScope, error) {
	return c.internalClient.GetMetricsScope(ctx, req, opts...)
}

// ListMetricsScopesByMonitoredProject returns a list of every Metrics Scope that a specific MonitoredProject
// has been added to. The metrics scope representing the specified monitored
// project will always be the first entry in the response.
func (c *MetricsScopesClient) ListMetricsScopesByMonitoredProject(ctx context.Context, req *metricsscopepb.ListMetricsScopesByMonitoredProjectRequest, opts ...gax.CallOption) (*metricsscopepb.ListMetricsScopesByMonitoredProjectResponse, error) {
	return c.internalClient.ListMetricsScopesByMonitoredProject(ctx, req, opts...)
}

// CreateMonitoredProject adds a MonitoredProject with the given project ID
// to the specified Metrics Scope.
func (c *MetricsScopesClient) CreateMonitoredProject(ctx context.Context, req *metricsscopepb.CreateMonitoredProjectRequest, opts ...gax.CallOption) (*CreateMonitoredProjectOperation, error) {
	return c.internalClient.CreateMonitoredProject(ctx, req, opts...)
}

// CreateMonitoredProjectOperation returns a new CreateMonitoredProjectOperation from a given name.
// The name must be that of a previously created CreateMonitoredProjectOperation, possibly from a different process.
func (c *MetricsScopesClient) CreateMonitoredProjectOperation(name string) *CreateMonitoredProjectOperation {
	return c.internalClient.CreateMonitoredProjectOperation(name)
}

// DeleteMonitoredProject deletes a MonitoredProject from the specified Metrics Scope.
func (c *MetricsScopesClient) DeleteMonitoredProject(ctx context.Context, req *metricsscopepb.DeleteMonitoredProjectRequest, opts ...gax.CallOption) (*DeleteMonitoredProjectOperation, error) {
	return c.internalClient.DeleteMonitoredProject(ctx, req, opts...)
}

// DeleteMonitoredProjectOperation returns a new DeleteMonitoredProjectOperation from a given name.
// The name must be that of a previously created DeleteMonitoredProjectOperation, possibly from a different process.
func (c *MetricsScopesClient) DeleteMonitoredProjectOperation(name string) *DeleteMonitoredProjectOperation {
	return c.internalClient.DeleteMonitoredProjectOperation(name)
}

// metricsScopesGRPCClient is a client for interacting with Cloud Monitoring API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type metricsScopesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing MetricsScopesClient
	CallOptions **MetricsScopesCallOptions

	// The gRPC API client.
	metricsScopesClient metricsscopepb.MetricsScopesClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewMetricsScopesClient creates a new metrics scopes client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Manages Cloud Monitoring Metrics Scopes, and the monitoring of Google Cloud
// projects and AWS accounts.
func NewMetricsScopesClient(ctx context.Context, opts ...option.ClientOption) (*MetricsScopesClient, error) {
	clientOpts := defaultMetricsScopesGRPCClientOptions()
	if newMetricsScopesClientHook != nil {
		hookOpts, err := newMetricsScopesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := MetricsScopesClient{CallOptions: defaultMetricsScopesCallOptions()}

	c := &metricsScopesGRPCClient{
		connPool:            connPool,
		metricsScopesClient: metricsscopepb.NewMetricsScopesClient(connPool),
		CallOptions:         &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *metricsScopesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *metricsScopesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *metricsScopesGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *metricsScopesGRPCClient) GetMetricsScope(ctx context.Context, req *metricsscopepb.GetMetricsScopeRequest, opts ...gax.CallOption) (*metricsscopepb.MetricsScope, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetMetricsScope[0:len((*c.CallOptions).GetMetricsScope):len((*c.CallOptions).GetMetricsScope)], opts...)
	var resp *metricsscopepb.MetricsScope
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.metricsScopesClient.GetMetricsScope(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *metricsScopesGRPCClient) ListMetricsScopesByMonitoredProject(ctx context.Context, req *metricsscopepb.ListMetricsScopesByMonitoredProjectRequest, opts ...gax.CallOption) (*metricsscopepb.ListMetricsScopesByMonitoredProjectResponse, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ListMetricsScopesByMonitoredProject[0:len((*c.CallOptions).ListMetricsScopesByMonitoredProject):len((*c.CallOptions).ListMetricsScopesByMonitoredProject)], opts...)
	var resp *metricsscopepb.ListMetricsScopesByMonitoredProjectResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.metricsScopesClient.ListMetricsScopesByMonitoredProject(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *metricsScopesGRPCClient) CreateMonitoredProject(ctx context.Context, req *metricsscopepb.CreateMonitoredProjectRequest, opts ...gax.CallOption) (*CreateMonitoredProjectOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateMonitoredProject[0:len((*c.CallOptions).CreateMonitoredProject):len((*c.CallOptions).CreateMonitoredProject)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.metricsScopesClient.CreateMonitoredProject(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateMonitoredProjectOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *metricsScopesGRPCClient) DeleteMonitoredProject(ctx context.Context, req *metricsscopepb.DeleteMonitoredProjectRequest, opts ...gax.CallOption) (*DeleteMonitoredProjectOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteMonitoredProject[0:len((*c.CallOptions).DeleteMonitoredProject):len((*c.CallOptions).DeleteMonitoredProject)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.metricsScopesClient.DeleteMonitoredProject(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteMonitoredProjectOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// CreateMonitoredProjectOperation manages a long-running operation from CreateMonitoredProject.
type CreateMonitoredProjectOperation struct {
	lro *longrunning.Operation
}

// CreateMonitoredProjectOperation returns a new CreateMonitoredProjectOperation from a given name.
// The name must be that of a previously created CreateMonitoredProjectOperation, possibly from a different process.
func (c *metricsScopesGRPCClient) CreateMonitoredProjectOperation(name string) *CreateMonitoredProjectOperation {
	return &CreateMonitoredProjectOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateMonitoredProjectOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*metricsscopepb.MonitoredProject, error) {
	var resp metricsscopepb.MonitoredProject
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *CreateMonitoredProjectOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*metricsscopepb.MonitoredProject, error) {
	var resp metricsscopepb.MonitoredProject
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *CreateMonitoredProjectOperation) Metadata() (*metricsscopepb.OperationMetadata, error) {
	var meta metricsscopepb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateMonitoredProjectOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateMonitoredProjectOperation) Name() string {
	return op.lro.Name()
}

// DeleteMonitoredProjectOperation manages a long-running operation from DeleteMonitoredProject.
type DeleteMonitoredProjectOperation struct {
	lro *longrunning.Operation
}

// DeleteMonitoredProjectOperation returns a new DeleteMonitoredProjectOperation from a given name.
// The name must be that of a previously created DeleteMonitoredProjectOperation, possibly from a different process.
func (c *metricsScopesGRPCClient) DeleteMonitoredProjectOperation(name string) *DeleteMonitoredProjectOperation {
	return &DeleteMonitoredProjectOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteMonitoredProjectOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.WaitWithInterval(ctx, nil, time.Minute, opts...)
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *DeleteMonitoredProjectOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DeleteMonitoredProjectOperation) Metadata() (*metricsscopepb.OperationMetadata, error) {
	var meta metricsscopepb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteMonitoredProjectOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteMonitoredProjectOperation) Name() string {
	return op.lro.Name()
}
