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

package automl

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	automlpb "cloud.google.com/go/automl/apiv1beta1/automlpb"
	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var newPredictionClientHook clientHook

// PredictionCallOptions contains the retry settings for each method of PredictionClient.
type PredictionCallOptions struct {
	Predict      []gax.CallOption
	BatchPredict []gax.CallOption
}

func defaultPredictionGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("automl.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("automl.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://automl.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultPredictionCallOptions() *PredictionCallOptions {
	return &PredictionCallOptions{
		Predict: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		BatchPredict: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

func defaultPredictionRESTCallOptions() *PredictionCallOptions {
	return &PredictionCallOptions{
		Predict: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		BatchPredict: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

// internalPredictionClient is an interface that defines the methods available from Cloud AutoML API.
type internalPredictionClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Predict(context.Context, *automlpb.PredictRequest, ...gax.CallOption) (*automlpb.PredictResponse, error)
	BatchPredict(context.Context, *automlpb.BatchPredictRequest, ...gax.CallOption) (*BatchPredictOperation, error)
	BatchPredictOperation(name string) *BatchPredictOperation
}

// PredictionClient is a client for interacting with Cloud AutoML API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// AutoML Prediction API.
//
// On any input that is documented to expect a string parameter in
// snake_case or kebab-case, either of those cases is accepted.
type PredictionClient struct {
	// The internal transport-dependent client.
	internalClient internalPredictionClient

	// The call options for this service.
	CallOptions *PredictionCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *PredictionClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *PredictionClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *PredictionClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// Predict perform an online prediction. The prediction result will be directly
// returned in the response.
// Available for following ML problems, and their expected request payloads:
//
//	Image Classification - Image in .JPEG, .GIF or .PNG format, image_bytes
//	up to 30MB.
//
//	Image Object Detection - Image in .JPEG, .GIF or .PNG format, image_bytes
//	up to 30MB.
//
//	Text Classification - TextSnippet, content up to 60,000 characters,
//	UTF-8 encoded.
//
//	Text Extraction - TextSnippet, content up to 30,000 characters,
//	UTF-8 NFC encoded.
//
//	Translation - TextSnippet, content up to 25,000 characters, UTF-8
//	encoded.
//
//	Tables - Row, with column values matching the columns of the model,
//	up to 5MB. Not available for FORECASTING
//
// prediction_type.
//
//	Text Sentiment - TextSnippet, content up 500 characters, UTF-8
//	encoded.
func (c *PredictionClient) Predict(ctx context.Context, req *automlpb.PredictRequest, opts ...gax.CallOption) (*automlpb.PredictResponse, error) {
	return c.internalClient.Predict(ctx, req, opts...)
}

// BatchPredict perform a batch prediction. Unlike the online Predict, batch
// prediction result won’t be immediately available in the response. Instead,
// a long running operation object is returned. User can poll the operation
// result via GetOperation
// method. Once the operation is done, BatchPredictResult is returned in
// the response field.
// Available for following ML problems:
//
//	Image Classification
//
//	Image Object Detection
//
//	Video Classification
//
//	Video Object Tracking * Text Extraction
//
//	Tables
func (c *PredictionClient) BatchPredict(ctx context.Context, req *automlpb.BatchPredictRequest, opts ...gax.CallOption) (*BatchPredictOperation, error) {
	return c.internalClient.BatchPredict(ctx, req, opts...)
}

// BatchPredictOperation returns a new BatchPredictOperation from a given name.
// The name must be that of a previously created BatchPredictOperation, possibly from a different process.
func (c *PredictionClient) BatchPredictOperation(name string) *BatchPredictOperation {
	return c.internalClient.BatchPredictOperation(name)
}

// predictionGRPCClient is a client for interacting with Cloud AutoML API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type predictionGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing PredictionClient
	CallOptions **PredictionCallOptions

	// The gRPC API client.
	predictionClient automlpb.PredictionServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewPredictionClient creates a new prediction service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// AutoML Prediction API.
//
// On any input that is documented to expect a string parameter in
// snake_case or kebab-case, either of those cases is accepted.
func NewPredictionClient(ctx context.Context, opts ...option.ClientOption) (*PredictionClient, error) {
	clientOpts := defaultPredictionGRPCClientOptions()
	if newPredictionClientHook != nil {
		hookOpts, err := newPredictionClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := PredictionClient{CallOptions: defaultPredictionCallOptions()}

	c := &predictionGRPCClient{
		connPool:         connPool,
		predictionClient: automlpb.NewPredictionServiceClient(connPool),
		CallOptions:      &client.CallOptions,
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
func (c *predictionGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *predictionGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *predictionGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type predictionRESTClient struct {
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

	// Points back to the CallOptions field of the containing PredictionClient
	CallOptions **PredictionCallOptions
}

// NewPredictionRESTClient creates a new prediction service rest client.
//
// AutoML Prediction API.
//
// On any input that is documented to expect a string parameter in
// snake_case or kebab-case, either of those cases is accepted.
func NewPredictionRESTClient(ctx context.Context, opts ...option.ClientOption) (*PredictionClient, error) {
	clientOpts := append(defaultPredictionRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultPredictionRESTCallOptions()
	c := &predictionRESTClient{
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

	return &PredictionClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultPredictionRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://automl.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://automl.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://automl.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *predictionRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *predictionRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *predictionRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *predictionGRPCClient) Predict(ctx context.Context, req *automlpb.PredictRequest, opts ...gax.CallOption) (*automlpb.PredictResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Predict[0:len((*c.CallOptions).Predict):len((*c.CallOptions).Predict)], opts...)
	var resp *automlpb.PredictResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.predictionClient.Predict(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *predictionGRPCClient) BatchPredict(ctx context.Context, req *automlpb.BatchPredictRequest, opts ...gax.CallOption) (*BatchPredictOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).BatchPredict[0:len((*c.CallOptions).BatchPredict):len((*c.CallOptions).BatchPredict)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.predictionClient.BatchPredict(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &BatchPredictOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// Predict perform an online prediction. The prediction result will be directly
// returned in the response.
// Available for following ML problems, and their expected request payloads:
//
//	Image Classification - Image in .JPEG, .GIF or .PNG format, image_bytes
//	up to 30MB.
//
//	Image Object Detection - Image in .JPEG, .GIF or .PNG format, image_bytes
//	up to 30MB.
//
//	Text Classification - TextSnippet, content up to 60,000 characters,
//	UTF-8 encoded.
//
//	Text Extraction - TextSnippet, content up to 30,000 characters,
//	UTF-8 NFC encoded.
//
//	Translation - TextSnippet, content up to 25,000 characters, UTF-8
//	encoded.
//
//	Tables - Row, with column values matching the columns of the model,
//	up to 5MB. Not available for FORECASTING
//
// prediction_type.
//
//	Text Sentiment - TextSnippet, content up 500 characters, UTF-8
//	encoded.
func (c *predictionRESTClient) Predict(ctx context.Context, req *automlpb.PredictRequest, opts ...gax.CallOption) (*automlpb.PredictResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v:predict", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).Predict[0:len((*c.CallOptions).Predict):len((*c.CallOptions).Predict)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &automlpb.PredictResponse{}
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

// BatchPredict perform a batch prediction. Unlike the online Predict, batch
// prediction result won’t be immediately available in the response. Instead,
// a long running operation object is returned. User can poll the operation
// result via GetOperation
// method. Once the operation is done, BatchPredictResult is returned in
// the response field.
// Available for following ML problems:
//
//	Image Classification
//
//	Image Object Detection
//
//	Video Classification
//
//	Video Object Tracking * Text Extraction
//
//	Tables
func (c *predictionRESTClient) BatchPredict(ctx context.Context, req *automlpb.BatchPredictRequest, opts ...gax.CallOption) (*BatchPredictOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v:batchPredict", req.GetName())

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

	override := fmt.Sprintf("/v1beta1/%s", resp.GetName())
	return &BatchPredictOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// BatchPredictOperation manages a long-running operation from BatchPredict.
type BatchPredictOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// BatchPredictOperation returns a new BatchPredictOperation from a given name.
// The name must be that of a previously created BatchPredictOperation, possibly from a different process.
func (c *predictionGRPCClient) BatchPredictOperation(name string) *BatchPredictOperation {
	return &BatchPredictOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// BatchPredictOperation returns a new BatchPredictOperation from a given name.
// The name must be that of a previously created BatchPredictOperation, possibly from a different process.
func (c *predictionRESTClient) BatchPredictOperation(name string) *BatchPredictOperation {
	override := fmt.Sprintf("/v1beta1/%s", name)
	return &BatchPredictOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *BatchPredictOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*automlpb.BatchPredictResult, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp automlpb.BatchPredictResult
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
func (op *BatchPredictOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*automlpb.BatchPredictResult, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp automlpb.BatchPredictResult
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
func (op *BatchPredictOperation) Metadata() (*automlpb.OperationMetadata, error) {
	var meta automlpb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *BatchPredictOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *BatchPredictOperation) Name() string {
	return op.lro.Name()
}
