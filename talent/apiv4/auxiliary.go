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

package talent

import (
	"context"
	"time"

	"cloud.google.com/go/longrunning"
	talentpb "cloud.google.com/go/talent/apiv4/talentpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
)

// BatchCreateJobsOperation manages a long-running operation from BatchCreateJobs.
type BatchCreateJobsOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *BatchCreateJobsOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*talentpb.BatchCreateJobsResponse, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp talentpb.BatchCreateJobsResponse
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
func (op *BatchCreateJobsOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*talentpb.BatchCreateJobsResponse, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp talentpb.BatchCreateJobsResponse
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
func (op *BatchCreateJobsOperation) Metadata() (*talentpb.BatchOperationMetadata, error) {
	var meta talentpb.BatchOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *BatchCreateJobsOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *BatchCreateJobsOperation) Name() string {
	return op.lro.Name()
}

// BatchDeleteJobsOperation manages a long-running operation from BatchDeleteJobs.
type BatchDeleteJobsOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *BatchDeleteJobsOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*talentpb.BatchDeleteJobsResponse, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp talentpb.BatchDeleteJobsResponse
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
func (op *BatchDeleteJobsOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*talentpb.BatchDeleteJobsResponse, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp talentpb.BatchDeleteJobsResponse
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
func (op *BatchDeleteJobsOperation) Metadata() (*talentpb.BatchOperationMetadata, error) {
	var meta talentpb.BatchOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *BatchDeleteJobsOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *BatchDeleteJobsOperation) Name() string {
	return op.lro.Name()
}

// BatchUpdateJobsOperation manages a long-running operation from BatchUpdateJobs.
type BatchUpdateJobsOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *BatchUpdateJobsOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*talentpb.BatchUpdateJobsResponse, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp talentpb.BatchUpdateJobsResponse
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
func (op *BatchUpdateJobsOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*talentpb.BatchUpdateJobsResponse, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp talentpb.BatchUpdateJobsResponse
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
func (op *BatchUpdateJobsOperation) Metadata() (*talentpb.BatchOperationMetadata, error) {
	var meta talentpb.BatchOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *BatchUpdateJobsOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *BatchUpdateJobsOperation) Name() string {
	return op.lro.Name()
}

// CompanyIterator manages a stream of *talentpb.Company.
type CompanyIterator struct {
	items    []*talentpb.Company
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
	InternalFetch func(pageSize int, pageToken string) (results []*talentpb.Company, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *CompanyIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *CompanyIterator) Next() (*talentpb.Company, error) {
	var item *talentpb.Company
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *CompanyIterator) bufLen() int {
	return len(it.items)
}

func (it *CompanyIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// JobIterator manages a stream of *talentpb.Job.
type JobIterator struct {
	items    []*talentpb.Job
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
	InternalFetch func(pageSize int, pageToken string) (results []*talentpb.Job, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *JobIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *JobIterator) Next() (*talentpb.Job, error) {
	var item *talentpb.Job
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *JobIterator) bufLen() int {
	return len(it.items)
}

func (it *JobIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// TenantIterator manages a stream of *talentpb.Tenant.
type TenantIterator struct {
	items    []*talentpb.Tenant
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
	InternalFetch func(pageSize int, pageToken string) (results []*talentpb.Tenant, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *TenantIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TenantIterator) Next() (*talentpb.Tenant, error) {
	var item *talentpb.Tenant
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TenantIterator) bufLen() int {
	return len(it.items)
}

func (it *TenantIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
