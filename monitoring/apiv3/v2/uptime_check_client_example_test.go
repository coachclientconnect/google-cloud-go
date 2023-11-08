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

package monitoring_test

import (
	"context"

	monitoring "cloud.google.com/go/monitoring/apiv3/v2"
	monitoringpb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	"google.golang.org/api/iterator"
)

func ExampleNewUptimeCheckClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := monitoring.NewUptimeCheckClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleUptimeCheckClient_CreateUptimeCheckConfig() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := monitoring.NewUptimeCheckClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &monitoringpb.CreateUptimeCheckConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/monitoring/apiv3/v2/monitoringpb#CreateUptimeCheckConfigRequest.
	}
	resp, err := c.CreateUptimeCheckConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleUptimeCheckClient_DeleteUptimeCheckConfig() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := monitoring.NewUptimeCheckClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &monitoringpb.DeleteUptimeCheckConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/monitoring/apiv3/v2/monitoringpb#DeleteUptimeCheckConfigRequest.
	}
	err = c.DeleteUptimeCheckConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleUptimeCheckClient_GetUptimeCheckConfig() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := monitoring.NewUptimeCheckClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &monitoringpb.GetUptimeCheckConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/monitoring/apiv3/v2/monitoringpb#GetUptimeCheckConfigRequest.
	}
	resp, err := c.GetUptimeCheckConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleUptimeCheckClient_ListUptimeCheckConfigs() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := monitoring.NewUptimeCheckClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &monitoringpb.ListUptimeCheckConfigsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/monitoring/apiv3/v2/monitoringpb#ListUptimeCheckConfigsRequest.
	}
	it := c.ListUptimeCheckConfigs(ctx, req)
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

func ExampleUptimeCheckClient_ListUptimeCheckIps() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := monitoring.NewUptimeCheckClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &monitoringpb.ListUptimeCheckIpsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/monitoring/apiv3/v2/monitoringpb#ListUptimeCheckIpsRequest.
	}
	it := c.ListUptimeCheckIps(ctx, req)
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

func ExampleUptimeCheckClient_UpdateUptimeCheckConfig() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := monitoring.NewUptimeCheckClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &monitoringpb.UpdateUptimeCheckConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/monitoring/apiv3/v2/monitoringpb#UpdateUptimeCheckConfigRequest.
	}
	resp, err := c.UpdateUptimeCheckConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
