// Code generated by go-swagger; DO NOT EDIT.

package recording

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetRecordingJobsParams creates a new GetRecordingJobsParams object
// with the default values initialized.
func NewGetRecordingJobsParams() *GetRecordingJobsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("userId")
	)
	return &GetRecordingJobsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRecordingJobsParamsWithTimeout creates a new GetRecordingJobsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRecordingJobsParamsWithTimeout(timeout time.Duration) *GetRecordingJobsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("userId")
	)
	return &GetRecordingJobsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,

		timeout: timeout,
	}
}

// NewGetRecordingJobsParamsWithContext creates a new GetRecordingJobsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRecordingJobsParamsWithContext(ctx context.Context) *GetRecordingJobsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("userId")
	)
	return &GetRecordingJobsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,

		Context: ctx,
	}
}

// NewGetRecordingJobsParamsWithHTTPClient creates a new GetRecordingJobsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRecordingJobsParamsWithHTTPClient(client *http.Client) *GetRecordingJobsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("userId")
	)
	return &GetRecordingJobsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		HTTPClient: client,
	}
}

/*GetRecordingJobsParams contains all the parameters to send to the API endpoint
for the get recording jobs operation typically these are written to a http.Request
*/
type GetRecordingJobsParams struct {

	/*JobType
	  Job Type (Can be left empty for both)

	*/
	JobType *string
	/*PageNumber
	  Page number

	*/
	PageNumber *int32
	/*PageSize
	  Page size

	*/
	PageSize *int32
	/*ShowOnlyMyJobs
	  Show only my jobs

	*/
	ShowOnlyMyJobs *bool
	/*SortBy
	  Sort by

	*/
	SortBy *string
	/*State
	  Filter by state

	*/
	State *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get recording jobs params
func (o *GetRecordingJobsParams) WithTimeout(timeout time.Duration) *GetRecordingJobsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get recording jobs params
func (o *GetRecordingJobsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get recording jobs params
func (o *GetRecordingJobsParams) WithContext(ctx context.Context) *GetRecordingJobsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get recording jobs params
func (o *GetRecordingJobsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get recording jobs params
func (o *GetRecordingJobsParams) WithHTTPClient(client *http.Client) *GetRecordingJobsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get recording jobs params
func (o *GetRecordingJobsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobType adds the jobType to the get recording jobs params
func (o *GetRecordingJobsParams) WithJobType(jobType *string) *GetRecordingJobsParams {
	o.SetJobType(jobType)
	return o
}

// SetJobType adds the jobType to the get recording jobs params
func (o *GetRecordingJobsParams) SetJobType(jobType *string) {
	o.JobType = jobType
}

// WithPageNumber adds the pageNumber to the get recording jobs params
func (o *GetRecordingJobsParams) WithPageNumber(pageNumber *int32) *GetRecordingJobsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get recording jobs params
func (o *GetRecordingJobsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get recording jobs params
func (o *GetRecordingJobsParams) WithPageSize(pageSize *int32) *GetRecordingJobsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get recording jobs params
func (o *GetRecordingJobsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithShowOnlyMyJobs adds the showOnlyMyJobs to the get recording jobs params
func (o *GetRecordingJobsParams) WithShowOnlyMyJobs(showOnlyMyJobs *bool) *GetRecordingJobsParams {
	o.SetShowOnlyMyJobs(showOnlyMyJobs)
	return o
}

// SetShowOnlyMyJobs adds the showOnlyMyJobs to the get recording jobs params
func (o *GetRecordingJobsParams) SetShowOnlyMyJobs(showOnlyMyJobs *bool) {
	o.ShowOnlyMyJobs = showOnlyMyJobs
}

// WithSortBy adds the sortBy to the get recording jobs params
func (o *GetRecordingJobsParams) WithSortBy(sortBy *string) *GetRecordingJobsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get recording jobs params
func (o *GetRecordingJobsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithState adds the state to the get recording jobs params
func (o *GetRecordingJobsParams) WithState(state *string) *GetRecordingJobsParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the get recording jobs params
func (o *GetRecordingJobsParams) SetState(state *string) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *GetRecordingJobsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.JobType != nil {

		// query param jobType
		var qrJobType string
		if o.JobType != nil {
			qrJobType = *o.JobType
		}
		qJobType := qrJobType
		if qJobType != "" {
			if err := r.SetQueryParam("jobType", qJobType); err != nil {
				return err
			}
		}

	}

	if o.PageNumber != nil {

		// query param pageNumber
		var qrPageNumber int32
		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := swag.FormatInt32(qrPageNumber)
		if qPageNumber != "" {
			if err := r.SetQueryParam("pageNumber", qPageNumber); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int32
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}

	}

	if o.ShowOnlyMyJobs != nil {

		// query param showOnlyMyJobs
		var qrShowOnlyMyJobs bool
		if o.ShowOnlyMyJobs != nil {
			qrShowOnlyMyJobs = *o.ShowOnlyMyJobs
		}
		qShowOnlyMyJobs := swag.FormatBool(qrShowOnlyMyJobs)
		if qShowOnlyMyJobs != "" {
			if err := r.SetQueryParam("showOnlyMyJobs", qShowOnlyMyJobs); err != nil {
				return err
			}
		}

	}

	if o.SortBy != nil {

		// query param sortBy
		var qrSortBy string
		if o.SortBy != nil {
			qrSortBy = *o.SortBy
		}
		qSortBy := qrSortBy
		if qSortBy != "" {
			if err := r.SetQueryParam("sortBy", qSortBy); err != nil {
				return err
			}
		}

	}

	if o.State != nil {

		// query param state
		var qrState string
		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {
			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}