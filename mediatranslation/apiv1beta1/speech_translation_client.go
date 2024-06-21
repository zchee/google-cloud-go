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

package mediatranslation

import (
	"context"
	"math"

	mediatranslationpb "cloud.google.com/go/mediatranslation/apiv1beta1/mediatranslationpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
)

var newSpeechTranslationClientHook clientHook

// SpeechTranslationCallOptions contains the retry settings for each method of SpeechTranslationClient.
type SpeechTranslationCallOptions struct {
	StreamingTranslateSpeech []gax.CallOption
}

func defaultSpeechTranslationGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("mediatranslation.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("mediatranslation.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("mediatranslation.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://mediatranslation.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultSpeechTranslationCallOptions() *SpeechTranslationCallOptions {
	return &SpeechTranslationCallOptions{
		StreamingTranslateSpeech: []gax.CallOption{},
	}
}

// internalSpeechTranslationClient is an interface that defines the methods available from Media Translation API.
type internalSpeechTranslationClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	StreamingTranslateSpeech(context.Context, ...gax.CallOption) (mediatranslationpb.SpeechTranslationService_StreamingTranslateSpeechClient, error)
}

// SpeechTranslationClient is a client for interacting with Media Translation API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Provides translation from/to media types.
type SpeechTranslationClient struct {
	// The internal transport-dependent client.
	internalClient internalSpeechTranslationClient

	// The call options for this service.
	CallOptions *SpeechTranslationCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *SpeechTranslationClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *SpeechTranslationClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *SpeechTranslationClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// StreamingTranslateSpeech performs bidirectional streaming speech translation: receive results while
// sending audio. This method is only available via the gRPC API (not REST).
func (c *SpeechTranslationClient) StreamingTranslateSpeech(ctx context.Context, opts ...gax.CallOption) (mediatranslationpb.SpeechTranslationService_StreamingTranslateSpeechClient, error) {
	return c.internalClient.StreamingTranslateSpeech(ctx, opts...)
}

// speechTranslationGRPCClient is a client for interacting with Media Translation API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type speechTranslationGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing SpeechTranslationClient
	CallOptions **SpeechTranslationCallOptions

	// The gRPC API client.
	speechTranslationClient mediatranslationpb.SpeechTranslationServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewSpeechTranslationClient creates a new speech translation service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Provides translation from/to media types.
func NewSpeechTranslationClient(ctx context.Context, opts ...option.ClientOption) (*SpeechTranslationClient, error) {
	clientOpts := defaultSpeechTranslationGRPCClientOptions()
	if newSpeechTranslationClientHook != nil {
		hookOpts, err := newSpeechTranslationClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := SpeechTranslationClient{CallOptions: defaultSpeechTranslationCallOptions()}

	c := &speechTranslationGRPCClient{
		connPool:                connPool,
		speechTranslationClient: mediatranslationpb.NewSpeechTranslationServiceClient(connPool),
		CallOptions:             &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *speechTranslationGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *speechTranslationGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *speechTranslationGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *speechTranslationGRPCClient) StreamingTranslateSpeech(ctx context.Context, opts ...gax.CallOption) (mediatranslationpb.SpeechTranslationService_StreamingTranslateSpeechClient, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	var resp mediatranslationpb.SpeechTranslationService_StreamingTranslateSpeechClient
	opts = append((*c.CallOptions).StreamingTranslateSpeech[0:len((*c.CallOptions).StreamingTranslateSpeech):len((*c.CallOptions).StreamingTranslateSpeech)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.speechTranslationClient.StreamingTranslateSpeech(ctx, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
