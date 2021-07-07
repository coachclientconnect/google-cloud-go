// Copyright 2021 Google LLC
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

package compute_test

import (
	"context"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
)

func ExampleNewRegionInstanceGroupsRESTClient() {
	ctx := context.Background()
	c, err := compute.NewRegionInstanceGroupsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleRegionInstanceGroupsClient_Get() {
	ctx := context.Background()
	c, err := compute.NewRegionInstanceGroupsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.GetRegionInstanceGroupRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.Get(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegionInstanceGroupsClient_List() {
	ctx := context.Background()
	c, err := compute.NewRegionInstanceGroupsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.ListRegionInstanceGroupsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.List(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegionInstanceGroupsClient_ListInstances() {
	ctx := context.Background()
	c, err := compute.NewRegionInstanceGroupsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.ListInstancesRegionInstanceGroupsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ListInstances(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegionInstanceGroupsClient_SetNamedPorts() {
	ctx := context.Background()
	c, err := compute.NewRegionInstanceGroupsRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.SetNamedPortsRegionInstanceGroupRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.SetNamedPorts(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}