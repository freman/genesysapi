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

// NewGetRecordingRecordingkeysParams creates a new GetRecordingRecordingkeysParams object
// with the default values initialized.
func NewGetRecordingRecordingkeysParams() *GetRecordingRecordingkeysParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetRecordingRecordingkeysParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRecordingRecordingkeysParamsWithTimeout creates a new GetRecordingRecordingkeysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRecordingRecordingkeysParamsWithTimeout(timeout time.Duration) *GetRecordingRecordingkeysParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetRecordingRecordingkeysParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetRecordingRecordingkeysParamsWithContext creates a new GetRecordingRecordingkeysParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRecordingRecordingkeysParamsWithContext(ctx context.Context) *GetRecordingRecordingkeysParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetRecordingRecordingkeysParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetRecordingRecordingkeysParamsWithHTTPClient creates a new GetRecordingRecordingkeysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRecordingRecordingkeysParamsWithHTTPClient(client *http.Client) *GetRecordingRecordingkeysParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetRecordingRecordingkeysParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetRecordingRecordingkeysParams contains all the parameters to send to the API endpoint
for the get recording recordingkeys operation typically these are written to a http.Request
*/
type GetRecordingRecordingkeysParams struct {

	/*PageNumber
	  Page number

	*/
	PageNumber *int32
	/*PageSize
	  Page size

	*/
	PageSize *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get recording recordingkeys params
func (o *GetRecordingRecordingkeysParams) WithTimeout(timeout time.Duration) *GetRecordingRecordingkeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get recording recordingkeys params
func (o *GetRecordingRecordingkeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get recording recordingkeys params
func (o *GetRecordingRecordingkeysParams) WithContext(ctx context.Context) *GetRecordingRecordingkeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get recording recordingkeys params
func (o *GetRecordingRecordingkeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get recording recordingkeys params
func (o *GetRecordingRecordingkeysParams) WithHTTPClient(client *http.Client) *GetRecordingRecordingkeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get recording recordingkeys params
func (o *GetRecordingRecordingkeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the get recording recordingkeys params
func (o *GetRecordingRecordingkeysParams) WithPageNumber(pageNumber *int32) *GetRecordingRecordingkeysParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get recording recordingkeys params
func (o *GetRecordingRecordingkeysParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get recording recordingkeys params
func (o *GetRecordingRecordingkeysParams) WithPageSize(pageSize *int32) *GetRecordingRecordingkeysParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get recording recordingkeys params
func (o *GetRecordingRecordingkeysParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetRecordingRecordingkeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
