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

package networkconnectivity

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	iampb "cloud.google.com/go/iam/apiv1/iampb"
	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	networkconnectivitypb "cloud.google.com/go/networkconnectivity/apiv1/networkconnectivitypb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

var newPolicyBasedRoutingClientHook clientHook

// PolicyBasedRoutingCallOptions contains the retry settings for each method of PolicyBasedRoutingClient.
type PolicyBasedRoutingCallOptions struct {
	ListPolicyBasedRoutes  []gax.CallOption
	GetPolicyBasedRoute    []gax.CallOption
	CreatePolicyBasedRoute []gax.CallOption
	DeletePolicyBasedRoute []gax.CallOption
	GetLocation            []gax.CallOption
	ListLocations          []gax.CallOption
	GetIamPolicy           []gax.CallOption
	SetIamPolicy           []gax.CallOption
	TestIamPermissions     []gax.CallOption
	CancelOperation        []gax.CallOption
	DeleteOperation        []gax.CallOption
	GetOperation           []gax.CallOption
	ListOperations         []gax.CallOption
}

func defaultPolicyBasedRoutingGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("networkconnectivity.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("networkconnectivity.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("networkconnectivity.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://networkconnectivity.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultPolicyBasedRoutingCallOptions() *PolicyBasedRoutingCallOptions {
	return &PolicyBasedRoutingCallOptions{
		ListPolicyBasedRoutes: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetPolicyBasedRoute: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreatePolicyBasedRoute: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		DeletePolicyBasedRoute: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetLocation:        []gax.CallOption{},
		ListLocations:      []gax.CallOption{},
		GetIamPolicy:       []gax.CallOption{},
		SetIamPolicy:       []gax.CallOption{},
		TestIamPermissions: []gax.CallOption{},
		CancelOperation:    []gax.CallOption{},
		DeleteOperation:    []gax.CallOption{},
		GetOperation:       []gax.CallOption{},
		ListOperations:     []gax.CallOption{},
	}
}

// internalPolicyBasedRoutingClient is an interface that defines the methods available from Network Connectivity API.
type internalPolicyBasedRoutingClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListPolicyBasedRoutes(context.Context, *networkconnectivitypb.ListPolicyBasedRoutesRequest, ...gax.CallOption) *PolicyBasedRouteIterator
	GetPolicyBasedRoute(context.Context, *networkconnectivitypb.GetPolicyBasedRouteRequest, ...gax.CallOption) (*networkconnectivitypb.PolicyBasedRoute, error)
	CreatePolicyBasedRoute(context.Context, *networkconnectivitypb.CreatePolicyBasedRouteRequest, ...gax.CallOption) (*CreatePolicyBasedRouteOperation, error)
	CreatePolicyBasedRouteOperation(name string) *CreatePolicyBasedRouteOperation
	DeletePolicyBasedRoute(context.Context, *networkconnectivitypb.DeletePolicyBasedRouteRequest, ...gax.CallOption) (*DeletePolicyBasedRouteOperation, error)
	DeletePolicyBasedRouteOperation(name string) *DeletePolicyBasedRouteOperation
	GetLocation(context.Context, *locationpb.GetLocationRequest, ...gax.CallOption) (*locationpb.Location, error)
	ListLocations(context.Context, *locationpb.ListLocationsRequest, ...gax.CallOption) *LocationIterator
	GetIamPolicy(context.Context, *iampb.GetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	SetIamPolicy(context.Context, *iampb.SetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	TestIamPermissions(context.Context, *iampb.TestIamPermissionsRequest, ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error)
	CancelOperation(context.Context, *longrunningpb.CancelOperationRequest, ...gax.CallOption) error
	DeleteOperation(context.Context, *longrunningpb.DeleteOperationRequest, ...gax.CallOption) error
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
	ListOperations(context.Context, *longrunningpb.ListOperationsRequest, ...gax.CallOption) *OperationIterator
}

// PolicyBasedRoutingClient is a client for interacting with Network Connectivity API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Policy-Based Routing allows GCP customers to specify flexibile routing
// policies for Layer 4 traffic traversing through the connected service.
type PolicyBasedRoutingClient struct {
	// The internal transport-dependent client.
	internalClient internalPolicyBasedRoutingClient

	// The call options for this service.
	CallOptions *PolicyBasedRoutingCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *PolicyBasedRoutingClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *PolicyBasedRoutingClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *PolicyBasedRoutingClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListPolicyBasedRoutes lists PolicyBasedRoutes in a given project and location.
func (c *PolicyBasedRoutingClient) ListPolicyBasedRoutes(ctx context.Context, req *networkconnectivitypb.ListPolicyBasedRoutesRequest, opts ...gax.CallOption) *PolicyBasedRouteIterator {
	return c.internalClient.ListPolicyBasedRoutes(ctx, req, opts...)
}

// GetPolicyBasedRoute gets details of a single PolicyBasedRoute.
func (c *PolicyBasedRoutingClient) GetPolicyBasedRoute(ctx context.Context, req *networkconnectivitypb.GetPolicyBasedRouteRequest, opts ...gax.CallOption) (*networkconnectivitypb.PolicyBasedRoute, error) {
	return c.internalClient.GetPolicyBasedRoute(ctx, req, opts...)
}

// CreatePolicyBasedRoute creates a new PolicyBasedRoute in a given project and location.
func (c *PolicyBasedRoutingClient) CreatePolicyBasedRoute(ctx context.Context, req *networkconnectivitypb.CreatePolicyBasedRouteRequest, opts ...gax.CallOption) (*CreatePolicyBasedRouteOperation, error) {
	return c.internalClient.CreatePolicyBasedRoute(ctx, req, opts...)
}

// CreatePolicyBasedRouteOperation returns a new CreatePolicyBasedRouteOperation from a given name.
// The name must be that of a previously created CreatePolicyBasedRouteOperation, possibly from a different process.
func (c *PolicyBasedRoutingClient) CreatePolicyBasedRouteOperation(name string) *CreatePolicyBasedRouteOperation {
	return c.internalClient.CreatePolicyBasedRouteOperation(name)
}

// DeletePolicyBasedRoute deletes a single PolicyBasedRoute.
func (c *PolicyBasedRoutingClient) DeletePolicyBasedRoute(ctx context.Context, req *networkconnectivitypb.DeletePolicyBasedRouteRequest, opts ...gax.CallOption) (*DeletePolicyBasedRouteOperation, error) {
	return c.internalClient.DeletePolicyBasedRoute(ctx, req, opts...)
}

// DeletePolicyBasedRouteOperation returns a new DeletePolicyBasedRouteOperation from a given name.
// The name must be that of a previously created DeletePolicyBasedRouteOperation, possibly from a different process.
func (c *PolicyBasedRoutingClient) DeletePolicyBasedRouteOperation(name string) *DeletePolicyBasedRouteOperation {
	return c.internalClient.DeletePolicyBasedRouteOperation(name)
}

// GetLocation gets information about a location.
func (c *PolicyBasedRoutingClient) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
	return c.internalClient.GetLocation(ctx, req, opts...)
}

// ListLocations lists information about the supported locations for this service.
func (c *PolicyBasedRoutingClient) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
	return c.internalClient.ListLocations(ctx, req, opts...)
}

// GetIamPolicy gets the access control policy for a resource. Returns an empty policy
// if the resource exists and does not have a policy set.
func (c *PolicyBasedRoutingClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// SetIamPolicy sets the access control policy on the specified resource. Replaces
// any existing policy.
//
// Can return NOT_FOUND, INVALID_ARGUMENT, and PERMISSION_DENIED
// errors.
func (c *PolicyBasedRoutingClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// TestIamPermissions returns permissions that a caller has on the specified resource. If the
// resource does not exist, this will return an empty set of
// permissions, not a NOT_FOUND error.
//
// Note: This operation is designed to be used for building
// permission-aware UIs and command-line tools, not for authorization
// checking. This operation may “fail open” without warning.
func (c *PolicyBasedRoutingClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// CancelOperation is a utility method from google.longrunning.Operations.
func (c *PolicyBasedRoutingClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.CancelOperation(ctx, req, opts...)
}

// DeleteOperation is a utility method from google.longrunning.Operations.
func (c *PolicyBasedRoutingClient) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteOperation(ctx, req, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *PolicyBasedRoutingClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *PolicyBasedRoutingClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.ListOperations(ctx, req, opts...)
}

// policyBasedRoutingGRPCClient is a client for interacting with Network Connectivity API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type policyBasedRoutingGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing PolicyBasedRoutingClient
	CallOptions **PolicyBasedRoutingCallOptions

	// The gRPC API client.
	policyBasedRoutingClient networkconnectivitypb.PolicyBasedRoutingServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	operationsClient longrunningpb.OperationsClient

	iamPolicyClient iampb.IAMPolicyClient

	locationsClient locationpb.LocationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewPolicyBasedRoutingClient creates a new policy based routing service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Policy-Based Routing allows GCP customers to specify flexibile routing
// policies for Layer 4 traffic traversing through the connected service.
func NewPolicyBasedRoutingClient(ctx context.Context, opts ...option.ClientOption) (*PolicyBasedRoutingClient, error) {
	clientOpts := defaultPolicyBasedRoutingGRPCClientOptions()
	if newPolicyBasedRoutingClientHook != nil {
		hookOpts, err := newPolicyBasedRoutingClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := PolicyBasedRoutingClient{CallOptions: defaultPolicyBasedRoutingCallOptions()}

	c := &policyBasedRoutingGRPCClient{
		connPool:                 connPool,
		policyBasedRoutingClient: networkconnectivitypb.NewPolicyBasedRoutingServiceClient(connPool),
		CallOptions:              &client.CallOptions,
		operationsClient:         longrunningpb.NewOperationsClient(connPool),
		iamPolicyClient:          iampb.NewIAMPolicyClient(connPool),
		locationsClient:          locationpb.NewLocationsClient(connPool),
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
func (c *policyBasedRoutingGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *policyBasedRoutingGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *policyBasedRoutingGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *policyBasedRoutingGRPCClient) ListPolicyBasedRoutes(ctx context.Context, req *networkconnectivitypb.ListPolicyBasedRoutesRequest, opts ...gax.CallOption) *PolicyBasedRouteIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListPolicyBasedRoutes[0:len((*c.CallOptions).ListPolicyBasedRoutes):len((*c.CallOptions).ListPolicyBasedRoutes)], opts...)
	it := &PolicyBasedRouteIterator{}
	req = proto.Clone(req).(*networkconnectivitypb.ListPolicyBasedRoutesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*networkconnectivitypb.PolicyBasedRoute, string, error) {
		resp := &networkconnectivitypb.ListPolicyBasedRoutesResponse{}
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
			resp, err = c.policyBasedRoutingClient.ListPolicyBasedRoutes(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetPolicyBasedRoutes(), resp.GetNextPageToken(), nil
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

func (c *policyBasedRoutingGRPCClient) GetPolicyBasedRoute(ctx context.Context, req *networkconnectivitypb.GetPolicyBasedRouteRequest, opts ...gax.CallOption) (*networkconnectivitypb.PolicyBasedRoute, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetPolicyBasedRoute[0:len((*c.CallOptions).GetPolicyBasedRoute):len((*c.CallOptions).GetPolicyBasedRoute)], opts...)
	var resp *networkconnectivitypb.PolicyBasedRoute
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.policyBasedRoutingClient.GetPolicyBasedRoute(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *policyBasedRoutingGRPCClient) CreatePolicyBasedRoute(ctx context.Context, req *networkconnectivitypb.CreatePolicyBasedRouteRequest, opts ...gax.CallOption) (*CreatePolicyBasedRouteOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreatePolicyBasedRoute[0:len((*c.CallOptions).CreatePolicyBasedRoute):len((*c.CallOptions).CreatePolicyBasedRoute)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.policyBasedRoutingClient.CreatePolicyBasedRoute(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreatePolicyBasedRouteOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *policyBasedRoutingGRPCClient) DeletePolicyBasedRoute(ctx context.Context, req *networkconnectivitypb.DeletePolicyBasedRouteRequest, opts ...gax.CallOption) (*DeletePolicyBasedRouteOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeletePolicyBasedRoute[0:len((*c.CallOptions).DeletePolicyBasedRoute):len((*c.CallOptions).DeletePolicyBasedRoute)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.policyBasedRoutingClient.DeletePolicyBasedRoute(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeletePolicyBasedRouteOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *policyBasedRoutingGRPCClient) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetLocation[0:len((*c.CallOptions).GetLocation):len((*c.CallOptions).GetLocation)], opts...)
	var resp *locationpb.Location
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.locationsClient.GetLocation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *policyBasedRoutingGRPCClient) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListLocations[0:len((*c.CallOptions).ListLocations):len((*c.CallOptions).ListLocations)], opts...)
	it := &LocationIterator{}
	req = proto.Clone(req).(*locationpb.ListLocationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*locationpb.Location, string, error) {
		resp := &locationpb.ListLocationsResponse{}
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
			resp, err = c.locationsClient.ListLocations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetLocations(), resp.GetNextPageToken(), nil
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

func (c *policyBasedRoutingGRPCClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.iamPolicyClient.GetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *policyBasedRoutingGRPCClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.iamPolicyClient.SetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *policyBasedRoutingGRPCClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
	var resp *iampb.TestIamPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.iamPolicyClient.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *policyBasedRoutingGRPCClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CancelOperation[0:len((*c.CallOptions).CancelOperation):len((*c.CallOptions).CancelOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.operationsClient.CancelOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *policyBasedRoutingGRPCClient) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteOperation[0:len((*c.CallOptions).DeleteOperation):len((*c.CallOptions).DeleteOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.operationsClient.DeleteOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *policyBasedRoutingGRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
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

func (c *policyBasedRoutingGRPCClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListOperations[0:len((*c.CallOptions).ListOperations):len((*c.CallOptions).ListOperations)], opts...)
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		resp := &longrunningpb.ListOperationsResponse{}
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
			resp, err = c.operationsClient.ListOperations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetOperations(), resp.GetNextPageToken(), nil
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

// CreatePolicyBasedRouteOperation returns a new CreatePolicyBasedRouteOperation from a given name.
// The name must be that of a previously created CreatePolicyBasedRouteOperation, possibly from a different process.
func (c *policyBasedRoutingGRPCClient) CreatePolicyBasedRouteOperation(name string) *CreatePolicyBasedRouteOperation {
	return &CreatePolicyBasedRouteOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// DeletePolicyBasedRouteOperation returns a new DeletePolicyBasedRouteOperation from a given name.
// The name must be that of a previously created DeletePolicyBasedRouteOperation, possibly from a different process.
func (c *policyBasedRoutingGRPCClient) DeletePolicyBasedRouteOperation(name string) *DeletePolicyBasedRouteOperation {
	return &DeletePolicyBasedRouteOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}
