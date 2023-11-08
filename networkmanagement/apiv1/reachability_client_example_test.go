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

package networkmanagement_test

import (
	"context"

	networkmanagement "cloud.google.com/go/networkmanagement/apiv1"
	networkmanagementpb "cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb"
	"google.golang.org/api/iterator"
)

func ExampleNewReachabilityClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := networkmanagement.NewReachabilityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleNewReachabilityRESTClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := networkmanagement.NewReachabilityRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleReachabilityClient_CreateConnectivityTest() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := networkmanagement.NewReachabilityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &networkmanagementpb.CreateConnectivityTestRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb#CreateConnectivityTestRequest.
	}
	op, err := c.CreateConnectivityTest(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleReachabilityClient_DeleteConnectivityTest() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := networkmanagement.NewReachabilityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &networkmanagementpb.DeleteConnectivityTestRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb#DeleteConnectivityTestRequest.
	}
	op, err := c.DeleteConnectivityTest(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleReachabilityClient_GetConnectivityTest() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := networkmanagement.NewReachabilityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &networkmanagementpb.GetConnectivityTestRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb#GetConnectivityTestRequest.
	}
	resp, err := c.GetConnectivityTest(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleReachabilityClient_ListConnectivityTests() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := networkmanagement.NewReachabilityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &networkmanagementpb.ListConnectivityTestsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb#ListConnectivityTestsRequest.
	}
	it := c.ListConnectivityTests(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleReachabilityClient_RerunConnectivityTest() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := networkmanagement.NewReachabilityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &networkmanagementpb.RerunConnectivityTestRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb#RerunConnectivityTestRequest.
	}
	op, err := c.RerunConnectivityTest(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleReachabilityClient_UpdateConnectivityTest() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := networkmanagement.NewReachabilityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &networkmanagementpb.UpdateConnectivityTestRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb#UpdateConnectivityTestRequest.
	}
	op, err := c.UpdateConnectivityTest(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
