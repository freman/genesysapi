// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewGetAnalyticsUsersDetailsJobResultsParams creates a new GetAnalyticsUsersDetailsJobResultsParams object
// with the default values initialized.
func NewGetAnalyticsUsersDetailsJobResultsParams() *GetAnalyticsUsersDetailsJobResultsParams {
	var ()
	return &GetAnalyticsUsersDetailsJobResultsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAnalyticsUsersDetailsJobResultsParamsWithTimeout creates a new GetAnalyticsUsersDetailsJobResultsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAnalyticsUsersDetailsJobResultsParamsWithTimeout(timeout time.Duration) *GetAnalyticsUsersDetailsJobResultsParams {
	var ()
	return &GetAnalyticsUsersDetailsJobResultsParams{

		timeout: timeout,
	}
}

// NewGetAnalyticsUsersDetailsJobResultsParamsWithContext creates a new GetAnalyticsUsersDetailsJobResultsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAnalyticsUsersDetailsJobResultsParamsWithContext(ctx context.Context) *GetAnalyticsUsersDetailsJobResultsParams {
	var ()
	return &GetAnalyticsUsersDetailsJobResultsParams{

		Context: ctx,
	}
}

// NewGetAnalyticsUsersDetailsJobResultsParamsWithHTTPClient creates a new GetAnalyticsUsersDetailsJobResultsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAnalyticsUsersDetailsJobResultsParamsWithHTTPClient(client *http.Client) *GetAnalyticsUsersDetailsJobResultsParams {
	var ()
	return &GetAnalyticsUsersDetailsJobResultsParams{
		HTTPClient: client,
	}
}

/*GetAnalyticsUsersDetailsJobResultsParams contains all the parameters to send to the API endpoint
for the get analytics users details job results operation typically these are written to a http.Request
*/
type GetAnalyticsUsersDetailsJobResultsParams struct {

	/*Cursor
	  Indicates where to resume query results (not required for first page)

	*/
	Cursor *string
	/*JobID
	  jobId

	*/
	JobID string
	/*PageSize
	  The desired maximum number of results

	*/
	PageSize *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get analytics users details job results params
func (o *GetAnalyticsUsersDetailsJobResultsParams) WithTimeout(timeout time.Duration) *GetAnalyticsUsersDetailsJobResultsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get analytics users details job results params
func (o *GetAnalyticsUsersDetailsJobResultsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get analytics users details job results params
func (o *GetAnalyticsUsersDetailsJobResultsParams) WithContext(ctx context.Context) *GetAnalyticsUsersDetailsJobResultsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get analytics users details job results params
func (o *GetAnalyticsUsersDetailsJobResultsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get analytics users details job results params
func (o *GetAnalyticsUsersDetailsJobResultsParams) WithHTTPClient(client *http.Client) *GetAnalyticsUsersDetailsJobResultsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get analytics users details job results params
func (o *GetAnalyticsUsersDetailsJobResultsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCursor adds the cursor to the get analytics users details job results params
func (o *GetAnalyticsUsersDetailsJobResultsParams) WithCursor(cursor *string) *GetAnalyticsUsersDetailsJobResultsParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the get analytics users details job results params
func (o *GetAnalyticsUsersDetailsJobResultsParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithJobID adds the jobID to the get analytics users details job results params
func (o *GetAnalyticsUsersDetailsJobResultsParams) WithJobID(jobID string) *GetAnalyticsUsersDetailsJobResultsParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the get analytics users details job results params
func (o *GetAnalyticsUsersDetailsJobResultsParams) SetJobID(jobID string) {
	o.JobID = jobID
}

// WithPageSize adds the pageSize to the get analytics users details job results params
func (o *GetAnalyticsUsersDetailsJobResultsParams) WithPageSize(pageSize *int32) *GetAnalyticsUsersDetailsJobResultsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get analytics users details job results params
func (o *GetAnalyticsUsersDetailsJobResultsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetAnalyticsUsersDetailsJobResultsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cursor != nil {

		// query param cursor
		var qrCursor string
		if o.Cursor != nil {
			qrCursor = *o.Cursor
		}
		qCursor := qrCursor
		if qCursor != "" {
			if err := r.SetQueryParam("cursor", qCursor); err != nil {
				return err
			}
		}

	}

	// path param jobId
	if err := r.SetPathParam("jobId", o.JobID); err != nil {
		return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
