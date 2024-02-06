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

package places

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"

	placespb "cloud.google.com/go/maps/places/apiv1/placespb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	SearchNearby  []gax.CallOption
	SearchText    []gax.CallOption
	GetPhotoMedia []gax.CallOption
	GetPlace      []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("places.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("places.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("places.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://places.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		SearchNearby:  []gax.CallOption{},
		SearchText:    []gax.CallOption{},
		GetPhotoMedia: []gax.CallOption{},
		GetPlace:      []gax.CallOption{},
	}
}

func defaultRESTCallOptions() *CallOptions {
	return &CallOptions{
		SearchNearby:  []gax.CallOption{},
		SearchText:    []gax.CallOption{},
		GetPhotoMedia: []gax.CallOption{},
		GetPlace:      []gax.CallOption{},
	}
}

// internalClient is an interface that defines the methods available from Places API (New).
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	SearchNearby(context.Context, *placespb.SearchNearbyRequest, ...gax.CallOption) (*placespb.SearchNearbyResponse, error)
	SearchText(context.Context, *placespb.SearchTextRequest, ...gax.CallOption) (*placespb.SearchTextResponse, error)
	GetPhotoMedia(context.Context, *placespb.GetPhotoMediaRequest, ...gax.CallOption) (*placespb.PhotoMedia, error)
	GetPlace(context.Context, *placespb.GetPlaceRequest, ...gax.CallOption) (*placespb.Place, error)
}

// Client is a client for interacting with Places API (New).
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service definition for the Places API.
// Note: every request actually requires a field mask set outside of
// the request proto (all/’*’, is not assumed).  That can be set via either a
// side channel (SystemParameterContext) over RPC, or a header
// (X-Goog-FieldMask) over HTTP. See:
// https://cloud.google.com/apis/docs/system-parameters (at https://cloud.google.com/apis/docs/system-parameters)
type Client struct {
	// The internal transport-dependent client.
	internalClient internalClient

	// The call options for this service.
	CallOptions *CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// SearchNearby search for places near locations.
func (c *Client) SearchNearby(ctx context.Context, req *placespb.SearchNearbyRequest, opts ...gax.CallOption) (*placespb.SearchNearbyResponse, error) {
	return c.internalClient.SearchNearby(ctx, req, opts...)
}

// SearchText text query based place search.
func (c *Client) SearchText(ctx context.Context, req *placespb.SearchTextRequest, opts ...gax.CallOption) (*placespb.SearchTextResponse, error) {
	return c.internalClient.SearchText(ctx, req, opts...)
}

// GetPhotoMedia get a photo media with a photo reference string.
func (c *Client) GetPhotoMedia(ctx context.Context, req *placespb.GetPhotoMediaRequest, opts ...gax.CallOption) (*placespb.PhotoMedia, error) {
	return c.internalClient.GetPhotoMedia(ctx, req, opts...)
}

// GetPlace get place details with a place id (in a name) string.
func (c *Client) GetPlace(ctx context.Context, req *placespb.GetPlaceRequest, opts ...gax.CallOption) (*placespb.Place, error) {
	return c.internalClient.GetPlace(ctx, req, opts...)
}

// gRPCClient is a client for interacting with Places API (New) over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	// The gRPC API client.
	client placespb.PlacesClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewClient creates a new places client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service definition for the Places API.
// Note: every request actually requires a field mask set outside of
// the request proto (all/’*’, is not assumed).  That can be set via either a
// side channel (SystemParameterContext) over RPC, or a header
// (X-Goog-FieldMask) over HTTP. See:
// https://cloud.google.com/apis/docs/system-parameters (at https://cloud.google.com/apis/docs/system-parameters)
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := defaultGRPCClientOptions()
	if newClientHook != nil {
		hookOpts, err := newClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := Client{CallOptions: defaultCallOptions()}

	c := &gRPCClient{
		connPool:    connPool,
		client:      placespb.NewPlacesClient(connPool),
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
func (c *gRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type restClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions
}

// NewRESTClient creates a new places rest client.
//
// Service definition for the Places API.
// Note: every request actually requires a field mask set outside of
// the request proto (all/’*’, is not assumed).  That can be set via either a
// side channel (SystemParameterContext) over RPC, or a header
// (X-Goog-FieldMask) over HTTP. See:
// https://cloud.google.com/apis/docs/system-parameters (at https://cloud.google.com/apis/docs/system-parameters)
func NewRESTClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := append(defaultRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultRESTCallOptions()
	c := &restClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &Client{internalClient: c, CallOptions: callOpts}, nil
}

func defaultRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://places.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://places.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://places.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://places.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *restClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *restClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *restClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *gRPCClient) SearchNearby(ctx context.Context, req *placespb.SearchNearbyRequest, opts ...gax.CallOption) (*placespb.SearchNearbyResponse, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).SearchNearby[0:len((*c.CallOptions).SearchNearby):len((*c.CallOptions).SearchNearby)], opts...)
	var resp *placespb.SearchNearbyResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.SearchNearby(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) SearchText(ctx context.Context, req *placespb.SearchTextRequest, opts ...gax.CallOption) (*placespb.SearchTextResponse, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).SearchText[0:len((*c.CallOptions).SearchText):len((*c.CallOptions).SearchText)], opts...)
	var resp *placespb.SearchTextResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.SearchText(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) GetPhotoMedia(ctx context.Context, req *placespb.GetPhotoMediaRequest, opts ...gax.CallOption) (*placespb.PhotoMedia, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetPhotoMedia[0:len((*c.CallOptions).GetPhotoMedia):len((*c.CallOptions).GetPhotoMedia)], opts...)
	var resp *placespb.PhotoMedia
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetPhotoMedia(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) GetPlace(ctx context.Context, req *placespb.GetPlaceRequest, opts ...gax.CallOption) (*placespb.Place, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetPlace[0:len((*c.CallOptions).GetPlace):len((*c.CallOptions).GetPlace)], opts...)
	var resp *placespb.Place
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetPlace(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SearchNearby search for places near locations.
func (c *restClient) SearchNearby(ctx context.Context, req *placespb.SearchNearbyRequest, opts ...gax.CallOption) (*placespb.SearchNearbyResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/places:searchNearby")

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := append(c.xGoogHeaders, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).SearchNearby[0:len((*c.CallOptions).SearchNearby):len((*c.CallOptions).SearchNearby)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &placespb.SearchNearbyResponse{}
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

// SearchText text query based place search.
func (c *restClient) SearchText(ctx context.Context, req *placespb.SearchTextRequest, opts ...gax.CallOption) (*placespb.SearchTextResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/places:searchText")

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := append(c.xGoogHeaders, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).SearchText[0:len((*c.CallOptions).SearchText):len((*c.CallOptions).SearchText)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &placespb.SearchTextResponse{}
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

// GetPhotoMedia get a photo media with a photo reference string.
func (c *restClient) GetPhotoMedia(ctx context.Context, req *placespb.GetPhotoMediaRequest, opts ...gax.CallOption) (*placespb.PhotoMedia, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetMaxHeightPx() != 0 {
		params.Add("maxHeightPx", fmt.Sprintf("%v", req.GetMaxHeightPx()))
	}
	if req.GetMaxWidthPx() != 0 {
		params.Add("maxWidthPx", fmt.Sprintf("%v", req.GetMaxWidthPx()))
	}
	if req.GetSkipHttpRedirect() {
		params.Add("skipHttpRedirect", fmt.Sprintf("%v", req.GetSkipHttpRedirect()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetPhotoMedia[0:len((*c.CallOptions).GetPhotoMedia):len((*c.CallOptions).GetPhotoMedia)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &placespb.PhotoMedia{}
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

// GetPlace get place details with a place id (in a name) string.
func (c *restClient) GetPlace(ctx context.Context, req *placespb.GetPlaceRequest, opts ...gax.CallOption) (*placespb.Place, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetLanguageCode() != "" {
		params.Add("languageCode", fmt.Sprintf("%v", req.GetLanguageCode()))
	}
	if req.GetRegionCode() != "" {
		params.Add("regionCode", fmt.Sprintf("%v", req.GetRegionCode()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetPlace[0:len((*c.CallOptions).GetPlace):len((*c.CallOptions).GetPlace)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &placespb.Place{}
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
