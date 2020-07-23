// Code generated by go-swagger; DO NOT EDIT.

package conversations

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

// NewGetAnalyticsConversationsDetailsJobResultsParams creates a new GetAnalyticsConversationsDetailsJobResultsParams object
// with the default values initialized.
func NewGetAnalyticsConversationsDetailsJobResultsParams() *GetAnalyticsConversationsDetailsJobResultsParams {
	var ()
	return &GetAnalyticsConversationsDetailsJobResultsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAnalyticsConversationsDetailsJobResultsParamsWithTimeout creates a new GetAnalyticsConversationsDetailsJobResultsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAnalyticsConversationsDetailsJobResultsParamsWithTimeout(timeout time.Duration) *GetAnalyticsConversationsDetailsJobResultsParams {
	var ()
	return &GetAnalyticsConversationsDetailsJobResultsParams{

		timeout: timeout,
	}
}

// NewGetAnalyticsConversationsDetailsJobResultsParamsWithContext creates a new GetAnalyticsConversationsDetailsJobResultsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAnalyticsConversationsDetailsJobResultsParamsWithContext(ctx context.Context) *GetAnalyticsConversationsDetailsJobResultsParams {
	var ()
	return &GetAnalyticsConversationsDetailsJobResultsParams{

		Context: ctx,
	}
}

// NewGetAnalyticsConversationsDetailsJobResultsParamsWithHTTPClient creates a new GetAnalyticsConversationsDetailsJobResultsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAnalyticsConversationsDetailsJobResultsParamsWithHTTPClient(client *http.Client) *GetAnalyticsConversationsDetailsJobResultsParams {
	var ()
	return &GetAnalyticsConversationsDetailsJobResultsParams{
		HTTPClient: client,
	}
}

/*GetAnalyticsConversationsDetailsJobResultsParams contains all the parameters to send to the API endpoint
for the get analytics conversations details job results operation typically these are written to a http.Request
*/
type GetAnalyticsConversationsDetailsJobResultsParams struct {

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

// WithTimeout adds the timeout to the get analytics conversations details job results params
func (o *GetAnalyticsConversationsDetailsJobResultsParams) WithTimeout(timeout time.Duration) *GetAnalyticsConversationsDetailsJobResultsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get analytics conversations details job results params
func (o *GetAnalyticsConversationsDetailsJobResultsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get analytics conversations details job results params
func (o *GetAnalyticsConversationsDetailsJobResultsParams) WithContext(ctx context.Context) *GetAnalyticsConversationsDetailsJobResultsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get analytics conversations details job results params
func (o *GetAnalyticsConversationsDetailsJobResultsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get analytics conversations details job results params
func (o *GetAnalyticsConversationsDetailsJobResultsParams) WithHTTPClient(client *http.Client) *GetAnalyticsConversationsDetailsJobResultsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get analytics conversations details job results params
func (o *GetAnalyticsConversationsDetailsJobResultsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCursor adds the cursor to the get analytics conversations details job results params
func (o *GetAnalyticsConversationsDetailsJobResultsParams) WithCursor(cursor *string) *GetAnalyticsConversationsDetailsJobResultsParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the get analytics conversations details job results params
func (o *GetAnalyticsConversationsDetailsJobResultsParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithJobID adds the jobID to the get analytics conversations details job results params
func (o *GetAnalyticsConversationsDetailsJobResultsParams) WithJobID(jobID string) *GetAnalyticsConversationsDetailsJobResultsParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the get analytics conversations details job results params
func (o *GetAnalyticsConversationsDetailsJobResultsParams) SetJobID(jobID string) {
	o.JobID = jobID
}

// WithPageSize adds the pageSize to the get analytics conversations details job results params
func (o *GetAnalyticsConversationsDetailsJobResultsParams) WithPageSize(pageSize *int32) *GetAnalyticsConversationsDetailsJobResultsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get analytics conversations details job results params
func (o *GetAnalyticsConversationsDetailsJobResultsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetAnalyticsConversationsDetailsJobResultsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
