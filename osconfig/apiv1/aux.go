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

package osconfig

import (
	"context"
	"time"

	"cloud.google.com/go/longrunning"
	osconfigpb "cloud.google.com/go/osconfig/apiv1/osconfigpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
)

// CreateOSPolicyAssignmentOperation manages a long-running operation from CreateOSPolicyAssignment.
type CreateOSPolicyAssignmentOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateOSPolicyAssignmentOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*osconfigpb.OSPolicyAssignment, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp osconfigpb.OSPolicyAssignment
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
func (op *CreateOSPolicyAssignmentOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*osconfigpb.OSPolicyAssignment, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp osconfigpb.OSPolicyAssignment
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
func (op *CreateOSPolicyAssignmentOperation) Metadata() (*osconfigpb.OSPolicyAssignmentOperationMetadata, error) {
	var meta osconfigpb.OSPolicyAssignmentOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateOSPolicyAssignmentOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateOSPolicyAssignmentOperation) Name() string {
	return op.lro.Name()
}

// DeleteOSPolicyAssignmentOperation manages a long-running operation from DeleteOSPolicyAssignment.
type DeleteOSPolicyAssignmentOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteOSPolicyAssignmentOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
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
func (op *DeleteOSPolicyAssignmentOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DeleteOSPolicyAssignmentOperation) Metadata() (*osconfigpb.OSPolicyAssignmentOperationMetadata, error) {
	var meta osconfigpb.OSPolicyAssignmentOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteOSPolicyAssignmentOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteOSPolicyAssignmentOperation) Name() string {
	return op.lro.Name()
}

// UpdateOSPolicyAssignmentOperation manages a long-running operation from UpdateOSPolicyAssignment.
type UpdateOSPolicyAssignmentOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *UpdateOSPolicyAssignmentOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*osconfigpb.OSPolicyAssignment, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp osconfigpb.OSPolicyAssignment
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
func (op *UpdateOSPolicyAssignmentOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*osconfigpb.OSPolicyAssignment, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp osconfigpb.OSPolicyAssignment
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
func (op *UpdateOSPolicyAssignmentOperation) Metadata() (*osconfigpb.OSPolicyAssignmentOperationMetadata, error) {
	var meta osconfigpb.OSPolicyAssignmentOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *UpdateOSPolicyAssignmentOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *UpdateOSPolicyAssignmentOperation) Name() string {
	return op.lro.Name()
}

// InventoryIterator manages a stream of *osconfigpb.Inventory.
type InventoryIterator struct {
	items    []*osconfigpb.Inventory
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
	InternalFetch func(pageSize int, pageToken string) (results []*osconfigpb.Inventory, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *InventoryIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *InventoryIterator) Next() (*osconfigpb.Inventory, error) {
	var item *osconfigpb.Inventory
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *InventoryIterator) bufLen() int {
	return len(it.items)
}

func (it *InventoryIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// OSPolicyAssignmentIterator manages a stream of *osconfigpb.OSPolicyAssignment.
type OSPolicyAssignmentIterator struct {
	items    []*osconfigpb.OSPolicyAssignment
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
	InternalFetch func(pageSize int, pageToken string) (results []*osconfigpb.OSPolicyAssignment, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *OSPolicyAssignmentIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *OSPolicyAssignmentIterator) Next() (*osconfigpb.OSPolicyAssignment, error) {
	var item *osconfigpb.OSPolicyAssignment
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *OSPolicyAssignmentIterator) bufLen() int {
	return len(it.items)
}

func (it *OSPolicyAssignmentIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// OSPolicyAssignmentReportIterator manages a stream of *osconfigpb.OSPolicyAssignmentReport.
type OSPolicyAssignmentReportIterator struct {
	items    []*osconfigpb.OSPolicyAssignmentReport
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
	InternalFetch func(pageSize int, pageToken string) (results []*osconfigpb.OSPolicyAssignmentReport, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *OSPolicyAssignmentReportIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *OSPolicyAssignmentReportIterator) Next() (*osconfigpb.OSPolicyAssignmentReport, error) {
	var item *osconfigpb.OSPolicyAssignmentReport
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *OSPolicyAssignmentReportIterator) bufLen() int {
	return len(it.items)
}

func (it *OSPolicyAssignmentReportIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// PatchDeploymentIterator manages a stream of *osconfigpb.PatchDeployment.
type PatchDeploymentIterator struct {
	items    []*osconfigpb.PatchDeployment
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
	InternalFetch func(pageSize int, pageToken string) (results []*osconfigpb.PatchDeployment, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *PatchDeploymentIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *PatchDeploymentIterator) Next() (*osconfigpb.PatchDeployment, error) {
	var item *osconfigpb.PatchDeployment
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *PatchDeploymentIterator) bufLen() int {
	return len(it.items)
}

func (it *PatchDeploymentIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// PatchJobInstanceDetailsIterator manages a stream of *osconfigpb.PatchJobInstanceDetails.
type PatchJobInstanceDetailsIterator struct {
	items    []*osconfigpb.PatchJobInstanceDetails
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
	InternalFetch func(pageSize int, pageToken string) (results []*osconfigpb.PatchJobInstanceDetails, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *PatchJobInstanceDetailsIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *PatchJobInstanceDetailsIterator) Next() (*osconfigpb.PatchJobInstanceDetails, error) {
	var item *osconfigpb.PatchJobInstanceDetails
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *PatchJobInstanceDetailsIterator) bufLen() int {
	return len(it.items)
}

func (it *PatchJobInstanceDetailsIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// PatchJobIterator manages a stream of *osconfigpb.PatchJob.
type PatchJobIterator struct {
	items    []*osconfigpb.PatchJob
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
	InternalFetch func(pageSize int, pageToken string) (results []*osconfigpb.PatchJob, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *PatchJobIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *PatchJobIterator) Next() (*osconfigpb.PatchJob, error) {
	var item *osconfigpb.PatchJob
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *PatchJobIterator) bufLen() int {
	return len(it.items)
}

func (it *PatchJobIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// VulnerabilityReportIterator manages a stream of *osconfigpb.VulnerabilityReport.
type VulnerabilityReportIterator struct {
	items    []*osconfigpb.VulnerabilityReport
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
	InternalFetch func(pageSize int, pageToken string) (results []*osconfigpb.VulnerabilityReport, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *VulnerabilityReportIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *VulnerabilityReportIterator) Next() (*osconfigpb.VulnerabilityReport, error) {
	var item *osconfigpb.VulnerabilityReport
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *VulnerabilityReportIterator) bufLen() int {
	return len(it.items)
}

func (it *VulnerabilityReportIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
