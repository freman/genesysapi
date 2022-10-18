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

// NewGetRecordingsRetentionQueryParams creates a new GetRecordingsRetentionQueryParams object
// with the default values initialized.
func NewGetRecordingsRetentionQueryParams() *GetRecordingsRetentionQueryParams {
	var (
		pageSizeDefault = int32(25)
	)
	return &GetRecordingsRetentionQueryParams{
		PageSize: &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRecordingsRetentionQueryParamsWithTimeout creates a new GetRecordingsRetentionQueryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRecordingsRetentionQueryParamsWithTimeout(timeout time.Duration) *GetRecordingsRetentionQueryParams {
	var (
		pageSizeDefault = int32(25)
	)
	return &GetRecordingsRetentionQueryParams{
		PageSize: &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetRecordingsRetentionQueryParamsWithContext creates a new GetRecordingsRetentionQueryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRecordingsRetentionQueryParamsWithContext(ctx context.Context) *GetRecordingsRetentionQueryParams {
	var (
		pageSizeDefault = int32(25)
	)
	return &GetRecordingsRetentionQueryParams{
		PageSize: &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetRecordingsRetentionQueryParamsWithHTTPClient creates a new GetRecordingsRetentionQueryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRecordingsRetentionQueryParamsWithHTTPClient(client *http.Client) *GetRecordingsRetentionQueryParams {
	var (
		pageSizeDefault = int32(25)
	)
	return &GetRecordingsRetentionQueryParams{
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetRecordingsRetentionQueryParams contains all the parameters to send to the API endpoint
for the get recordings retention query operation typically these are written to a http.Request
*/
type GetRecordingsRetentionQueryParams struct {

	/*Cursor
	  Indicates where to resume query results (not required for first page)

	*/
	Cursor *string
	/*PageSize
	  Page size. Maximum is 500.

	*/
	PageSize *int32
	/*RetentionThresholdDays
	  Fetch retention data for recordings retained for more days than the provided value.

	*/
	RetentionThresholdDays int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get recordings retention query params
func (o *GetRecordingsRetentionQueryParams) WithTimeout(timeout time.Duration) *GetRecordingsRetentionQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get recordings retention query params
func (o *GetRecordingsRetentionQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get recordings retention query params
func (o *GetRecordingsRetentionQueryParams) WithContext(ctx context.Context) *GetRecordingsRetentionQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get recordings retention query params
func (o *GetRecordingsRetentionQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get recordings retention query params
func (o *GetRecordingsRetentionQueryParams) WithHTTPClient(client *http.Client) *GetRecordingsRetentionQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get recordings retention query params
func (o *GetRecordingsRetentionQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCursor adds the cursor to the get recordings retention query params
func (o *GetRecordingsRetentionQueryParams) WithCursor(cursor *string) *GetRecordingsRetentionQueryParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the get recordings retention query params
func (o *GetRecordingsRetentionQueryParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithPageSize adds the pageSize to the get recordings retention query params
func (o *GetRecordingsRetentionQueryParams) WithPageSize(pageSize *int32) *GetRecordingsRetentionQueryParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get recordings retention query params
func (o *GetRecordingsRetentionQueryParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithRetentionThresholdDays adds the retentionThresholdDays to the get recordings retention query params
func (o *GetRecordingsRetentionQueryParams) WithRetentionThresholdDays(retentionThresholdDays int32) *GetRecordingsRetentionQueryParams {
	o.SetRetentionThresholdDays(retentionThresholdDays)
	return o
}

// SetRetentionThresholdDays adds the retentionThresholdDays to the get recordings retention query params
func (o *GetRecordingsRetentionQueryParams) SetRetentionThresholdDays(retentionThresholdDays int32) {
	o.RetentionThresholdDays = retentionThresholdDays
}

// WriteToRequest writes these params to a swagger request
func (o *GetRecordingsRetentionQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param retentionThresholdDays
	qrRetentionThresholdDays := o.RetentionThresholdDays
	qRetentionThresholdDays := swag.FormatInt32(qrRetentionThresholdDays)
	if qRetentionThresholdDays != "" {
		if err := r.SetQueryParam("retentionThresholdDays", qRetentionThresholdDays); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
