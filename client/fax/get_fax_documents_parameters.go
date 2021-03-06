// Code generated by go-swagger; DO NOT EDIT.

package fax

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

// NewGetFaxDocumentsParams creates a new GetFaxDocumentsParams object
// with the default values initialized.
func NewGetFaxDocumentsParams() *GetFaxDocumentsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetFaxDocumentsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFaxDocumentsParamsWithTimeout creates a new GetFaxDocumentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFaxDocumentsParamsWithTimeout(timeout time.Duration) *GetFaxDocumentsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetFaxDocumentsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetFaxDocumentsParamsWithContext creates a new GetFaxDocumentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFaxDocumentsParamsWithContext(ctx context.Context) *GetFaxDocumentsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetFaxDocumentsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetFaxDocumentsParamsWithHTTPClient creates a new GetFaxDocumentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFaxDocumentsParamsWithHTTPClient(client *http.Client) *GetFaxDocumentsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetFaxDocumentsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetFaxDocumentsParams contains all the parameters to send to the API endpoint
for the get fax documents operation typically these are written to a http.Request
*/
type GetFaxDocumentsParams struct {

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

// WithTimeout adds the timeout to the get fax documents params
func (o *GetFaxDocumentsParams) WithTimeout(timeout time.Duration) *GetFaxDocumentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fax documents params
func (o *GetFaxDocumentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fax documents params
func (o *GetFaxDocumentsParams) WithContext(ctx context.Context) *GetFaxDocumentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fax documents params
func (o *GetFaxDocumentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fax documents params
func (o *GetFaxDocumentsParams) WithHTTPClient(client *http.Client) *GetFaxDocumentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fax documents params
func (o *GetFaxDocumentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the get fax documents params
func (o *GetFaxDocumentsParams) WithPageNumber(pageNumber *int32) *GetFaxDocumentsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get fax documents params
func (o *GetFaxDocumentsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get fax documents params
func (o *GetFaxDocumentsParams) WithPageSize(pageSize *int32) *GetFaxDocumentsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get fax documents params
func (o *GetFaxDocumentsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetFaxDocumentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
