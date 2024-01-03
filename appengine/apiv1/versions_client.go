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

package appengine

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	appenginepb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
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

var newVersionsClientHook clientHook

// VersionsCallOptions contains the retry settings for each method of VersionsClient.
type VersionsCallOptions struct {
	ListVersions  []gax.CallOption
	GetVersion    []gax.CallOption
	CreateVersion []gax.CallOption
	UpdateVersion []gax.CallOption
	DeleteVersion []gax.CallOption
}

func defaultVersionsGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("appengine.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("appengine.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://appengine.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultVersionsCallOptions() *VersionsCallOptions {
	return &VersionsCallOptions{
		ListVersions: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetVersion: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		CreateVersion: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		UpdateVersion: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		DeleteVersion: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

func defaultVersionsRESTCallOptions() *VersionsCallOptions {
	return &VersionsCallOptions{
		ListVersions: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetVersion: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		CreateVersion: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		UpdateVersion: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		DeleteVersion: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

// internalVersionsClient is an interface that defines the methods available from App Engine Admin API.
type internalVersionsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListVersions(context.Context, *appenginepb.ListVersionsRequest, ...gax.CallOption) *VersionIterator
	GetVersion(context.Context, *appenginepb.GetVersionRequest, ...gax.CallOption) (*appenginepb.Version, error)
	CreateVersion(context.Context, *appenginepb.CreateVersionRequest, ...gax.CallOption) (*CreateVersionOperation, error)
	CreateVersionOperation(name string) *CreateVersionOperation
	UpdateVersion(context.Context, *appenginepb.UpdateVersionRequest, ...gax.CallOption) (*UpdateVersionOperation, error)
	UpdateVersionOperation(name string) *UpdateVersionOperation
	DeleteVersion(context.Context, *appenginepb.DeleteVersionRequest, ...gax.CallOption) (*DeleteVersionOperation, error)
	DeleteVersionOperation(name string) *DeleteVersionOperation
}

// VersionsClient is a client for interacting with App Engine Admin API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Manages versions of a service.
type VersionsClient struct {
	// The internal transport-dependent client.
	internalClient internalVersionsClient

	// The call options for this service.
	CallOptions *VersionsCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *VersionsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *VersionsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *VersionsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListVersions lists the versions of a service.
func (c *VersionsClient) ListVersions(ctx context.Context, req *appenginepb.ListVersionsRequest, opts ...gax.CallOption) *VersionIterator {
	return c.internalClient.ListVersions(ctx, req, opts...)
}

// GetVersion gets the specified Version resource.
// By default, only a BASIC_VIEW will be returned.
// Specify the FULL_VIEW parameter to get the full resource.
func (c *VersionsClient) GetVersion(ctx context.Context, req *appenginepb.GetVersionRequest, opts ...gax.CallOption) (*appenginepb.Version, error) {
	return c.internalClient.GetVersion(ctx, req, opts...)
}

// CreateVersion deploys code and resource files to a new version.
func (c *VersionsClient) CreateVersion(ctx context.Context, req *appenginepb.CreateVersionRequest, opts ...gax.CallOption) (*CreateVersionOperation, error) {
	return c.internalClient.CreateVersion(ctx, req, opts...)
}

// CreateVersionOperation returns a new CreateVersionOperation from a given name.
// The name must be that of a previously created CreateVersionOperation, possibly from a different process.
func (c *VersionsClient) CreateVersionOperation(name string) *CreateVersionOperation {
	return c.internalClient.CreateVersionOperation(name)
}

// UpdateVersion updates the specified Version resource.
// You can specify the following fields depending on the App Engine
// environment and type of scaling that the version resource uses:
//
// Standard environment
//
//	instance_class (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#Version.FIELDS.instance_class)
//
// automatic scaling in the standard environment:
//
//	automatic_scaling.min_idle_instances (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#Version.FIELDS.automatic_scaling)
//
//	automatic_scaling.max_idle_instances (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#Version.FIELDS.automatic_scaling)
//
//	automaticScaling.standard_scheduler_settings.max_instances (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#StandardSchedulerSettings)
//
//	automaticScaling.standard_scheduler_settings.min_instances (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#StandardSchedulerSettings)
//
//	automaticScaling.standard_scheduler_settings.target_cpu_utilization (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#StandardSchedulerSettings)
//
//	automaticScaling.standard_scheduler_settings.target_throughput_utilization (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#StandardSchedulerSettings)
//
// basic scaling or manual scaling in the standard environment:
//
//	serving_status (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#Version.FIELDS.serving_status)
//
//	manual_scaling.instances (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#manualscaling)
//
// Flexible environment
//
//	serving_status (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#Version.FIELDS.serving_status)
//
// automatic scaling in the flexible environment:
//
//	automatic_scaling.min_total_instances (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#Version.FIELDS.automatic_scaling)
//
//	automatic_scaling.max_total_instances (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#Version.FIELDS.automatic_scaling)
//
//	automatic_scaling.cool_down_period_sec (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#Version.FIELDS.automatic_scaling)
//
//	automatic_scaling.cpu_utilization.target_utilization (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#Version.FIELDS.automatic_scaling)
//
// manual scaling in the flexible environment:
//
//	manual_scaling.instances (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#manualscaling)
func (c *VersionsClient) UpdateVersion(ctx context.Context, req *appenginepb.UpdateVersionRequest, opts ...gax.CallOption) (*UpdateVersionOperation, error) {
	return c.internalClient.UpdateVersion(ctx, req, opts...)
}

// UpdateVersionOperation returns a new UpdateVersionOperation from a given name.
// The name must be that of a previously created UpdateVersionOperation, possibly from a different process.
func (c *VersionsClient) UpdateVersionOperation(name string) *UpdateVersionOperation {
	return c.internalClient.UpdateVersionOperation(name)
}

// DeleteVersion deletes an existing Version resource.
func (c *VersionsClient) DeleteVersion(ctx context.Context, req *appenginepb.DeleteVersionRequest, opts ...gax.CallOption) (*DeleteVersionOperation, error) {
	return c.internalClient.DeleteVersion(ctx, req, opts...)
}

// DeleteVersionOperation returns a new DeleteVersionOperation from a given name.
// The name must be that of a previously created DeleteVersionOperation, possibly from a different process.
func (c *VersionsClient) DeleteVersionOperation(name string) *DeleteVersionOperation {
	return c.internalClient.DeleteVersionOperation(name)
}

// versionsGRPCClient is a client for interacting with App Engine Admin API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type versionsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing VersionsClient
	CallOptions **VersionsCallOptions

	// The gRPC API client.
	versionsClient appenginepb.VersionsClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewVersionsClient creates a new versions client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Manages versions of a service.
func NewVersionsClient(ctx context.Context, opts ...option.ClientOption) (*VersionsClient, error) {
	clientOpts := defaultVersionsGRPCClientOptions()
	if newVersionsClientHook != nil {
		hookOpts, err := newVersionsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := VersionsClient{CallOptions: defaultVersionsCallOptions()}

	c := &versionsGRPCClient{
		connPool:       connPool,
		versionsClient: appenginepb.NewVersionsClient(connPool),
		CallOptions:    &client.CallOptions,
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
func (c *versionsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *versionsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *versionsGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type versionsRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing VersionsClient
	CallOptions **VersionsCallOptions
}

// NewVersionsRESTClient creates a new versions rest client.
//
// Manages versions of a service.
func NewVersionsRESTClient(ctx context.Context, opts ...option.ClientOption) (*VersionsClient, error) {
	clientOpts := append(defaultVersionsRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultVersionsRESTCallOptions()
	c := &versionsRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	lroOpts := []option.ClientOption{
		option.WithHTTPClient(httpClient),
		option.WithEndpoint(endpoint),
	}
	opClient, err := lroauto.NewOperationsRESTClient(ctx, lroOpts...)
	if err != nil {
		return nil, err
	}
	c.LROClient = &opClient

	return &VersionsClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultVersionsRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://appengine.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://appengine.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://appengine.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *versionsRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *versionsRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *versionsRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *versionsGRPCClient) ListVersions(ctx context.Context, req *appenginepb.ListVersionsRequest, opts ...gax.CallOption) *VersionIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListVersions[0:len((*c.CallOptions).ListVersions):len((*c.CallOptions).ListVersions)], opts...)
	it := &VersionIterator{}
	req = proto.Clone(req).(*appenginepb.ListVersionsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*appenginepb.Version, string, error) {
		resp := &appenginepb.ListVersionsResponse{}
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
			resp, err = c.versionsClient.ListVersions(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetVersions(), resp.GetNextPageToken(), nil
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

func (c *versionsGRPCClient) GetVersion(ctx context.Context, req *appenginepb.GetVersionRequest, opts ...gax.CallOption) (*appenginepb.Version, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetVersion[0:len((*c.CallOptions).GetVersion):len((*c.CallOptions).GetVersion)], opts...)
	var resp *appenginepb.Version
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.versionsClient.GetVersion(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *versionsGRPCClient) CreateVersion(ctx context.Context, req *appenginepb.CreateVersionRequest, opts ...gax.CallOption) (*CreateVersionOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateVersion[0:len((*c.CallOptions).CreateVersion):len((*c.CallOptions).CreateVersion)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.versionsClient.CreateVersion(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateVersionOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *versionsGRPCClient) UpdateVersion(ctx context.Context, req *appenginepb.UpdateVersionRequest, opts ...gax.CallOption) (*UpdateVersionOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateVersion[0:len((*c.CallOptions).UpdateVersion):len((*c.CallOptions).UpdateVersion)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.versionsClient.UpdateVersion(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &UpdateVersionOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *versionsGRPCClient) DeleteVersion(ctx context.Context, req *appenginepb.DeleteVersionRequest, opts ...gax.CallOption) (*DeleteVersionOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteVersion[0:len((*c.CallOptions).DeleteVersion):len((*c.CallOptions).DeleteVersion)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.versionsClient.DeleteVersion(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteVersionOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// ListVersions lists the versions of a service.
func (c *versionsRESTClient) ListVersions(ctx context.Context, req *appenginepb.ListVersionsRequest, opts ...gax.CallOption) *VersionIterator {
	it := &VersionIterator{}
	req = proto.Clone(req).(*appenginepb.ListVersionsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*appenginepb.Version, string, error) {
		resp := &appenginepb.ListVersionsResponse{}
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
		baseUrl.Path += fmt.Sprintf("/v1/%v/versions", req.GetParent())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req.GetView() != 0 {
			params.Add("view", fmt.Sprintf("%v", req.GetView()))
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
		return resp.GetVersions(), resp.GetNextPageToken(), nil
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

// GetVersion gets the specified Version resource.
// By default, only a BASIC_VIEW will be returned.
// Specify the FULL_VIEW parameter to get the full resource.
func (c *versionsRESTClient) GetVersion(ctx context.Context, req *appenginepb.GetVersionRequest, opts ...gax.CallOption) (*appenginepb.Version, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetView() != 0 {
		params.Add("view", fmt.Sprintf("%v", req.GetView()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetVersion[0:len((*c.CallOptions).GetVersion):len((*c.CallOptions).GetVersion)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &appenginepb.Version{}
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

// CreateVersion deploys code and resource files to a new version.
func (c *versionsRESTClient) CreateVersion(ctx context.Context, req *appenginepb.CreateVersionRequest, opts ...gax.CallOption) (*CreateVersionOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetVersion()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v/versions", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
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

	override := fmt.Sprintf("/v1/%s", resp.GetName())
	return &CreateVersionOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// UpdateVersion updates the specified Version resource.
// You can specify the following fields depending on the App Engine
// environment and type of scaling that the version resource uses:
//
// Standard environment
//
//	instance_class (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#Version.FIELDS.instance_class)
//
// automatic scaling in the standard environment:
//
//	automatic_scaling.min_idle_instances (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#Version.FIELDS.automatic_scaling)
//
//	automatic_scaling.max_idle_instances (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#Version.FIELDS.automatic_scaling)
//
//	automaticScaling.standard_scheduler_settings.max_instances (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#StandardSchedulerSettings)
//
//	automaticScaling.standard_scheduler_settings.min_instances (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#StandardSchedulerSettings)
//
//	automaticScaling.standard_scheduler_settings.target_cpu_utilization (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#StandardSchedulerSettings)
//
//	automaticScaling.standard_scheduler_settings.target_throughput_utilization (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#StandardSchedulerSettings)
//
// basic scaling or manual scaling in the standard environment:
//
//	serving_status (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#Version.FIELDS.serving_status)
//
//	manual_scaling.instances (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#manualscaling)
//
// Flexible environment
//
//	serving_status (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#Version.FIELDS.serving_status)
//
// automatic scaling in the flexible environment:
//
//	automatic_scaling.min_total_instances (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#Version.FIELDS.automatic_scaling)
//
//	automatic_scaling.max_total_instances (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#Version.FIELDS.automatic_scaling)
//
//	automatic_scaling.cool_down_period_sec (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#Version.FIELDS.automatic_scaling)
//
//	automatic_scaling.cpu_utilization.target_utilization (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#Version.FIELDS.automatic_scaling)
//
// manual scaling in the flexible environment:
//
//	manual_scaling.instances (at https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#manualscaling)
func (c *versionsRESTClient) UpdateVersion(ctx context.Context, req *appenginepb.UpdateVersionRequest, opts ...gax.CallOption) (*UpdateVersionOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetVersion()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

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
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
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

	override := fmt.Sprintf("/v1/%s", resp.GetName())
	return &UpdateVersionOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// DeleteVersion deletes an existing Version resource.
func (c *versionsRESTClient) DeleteVersion(ctx context.Context, req *appenginepb.DeleteVersionRequest, opts ...gax.CallOption) (*DeleteVersionOperation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
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

	override := fmt.Sprintf("/v1/%s", resp.GetName())
	return &DeleteVersionOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// CreateVersionOperation returns a new CreateVersionOperation from a given name.
// The name must be that of a previously created CreateVersionOperation, possibly from a different process.
func (c *versionsGRPCClient) CreateVersionOperation(name string) *CreateVersionOperation {
	return &CreateVersionOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// CreateVersionOperation returns a new CreateVersionOperation from a given name.
// The name must be that of a previously created CreateVersionOperation, possibly from a different process.
func (c *versionsRESTClient) CreateVersionOperation(name string) *CreateVersionOperation {
	override := fmt.Sprintf("/v1/%s", name)
	return &CreateVersionOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}

// DeleteVersionOperation returns a new DeleteVersionOperation from a given name.
// The name must be that of a previously created DeleteVersionOperation, possibly from a different process.
func (c *versionsGRPCClient) DeleteVersionOperation(name string) *DeleteVersionOperation {
	return &DeleteVersionOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// DeleteVersionOperation returns a new DeleteVersionOperation from a given name.
// The name must be that of a previously created DeleteVersionOperation, possibly from a different process.
func (c *versionsRESTClient) DeleteVersionOperation(name string) *DeleteVersionOperation {
	override := fmt.Sprintf("/v1/%s", name)
	return &DeleteVersionOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}

// UpdateVersionOperation returns a new UpdateVersionOperation from a given name.
// The name must be that of a previously created UpdateVersionOperation, possibly from a different process.
func (c *versionsGRPCClient) UpdateVersionOperation(name string) *UpdateVersionOperation {
	return &UpdateVersionOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// UpdateVersionOperation returns a new UpdateVersionOperation from a given name.
// The name must be that of a previously created UpdateVersionOperation, possibly from a different process.
func (c *versionsRESTClient) UpdateVersionOperation(name string) *UpdateVersionOperation {
	override := fmt.Sprintf("/v1/%s", name)
	return &UpdateVersionOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}
