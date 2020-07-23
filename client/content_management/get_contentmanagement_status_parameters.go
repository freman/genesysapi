// Code generated by go-swagger; DO NOT EDIT.

package content_management

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

// NewGetContentmanagementStatusParams creates a new GetContentmanagementStatusParams object
// with the default values initialized.
func NewGetContentmanagementStatusParams() *GetContentmanagementStatusParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetContentmanagementStatusParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetContentmanagementStatusParamsWithTimeout creates a new GetContentmanagementStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetContentmanagementStatusParamsWithTimeout(timeout time.Duration) *GetContentmanagementStatusParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetContentmanagementStatusParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetContentmanagementStatusParamsWithContext creates a new GetContentmanagementStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetContentmanagementStatusParamsWithContext(ctx context.Context) *GetContentmanagementStatusParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetContentmanagementStatusParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetContentmanagementStatusParamsWithHTTPClient creates a new GetContentmanagementStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetContentmanagementStatusParamsWithHTTPClient(client *http.Client) *GetContentmanagementStatusParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetContentmanagementStatusParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetContentmanagementStatusParams contains all the parameters to send to the API endpoint
for the get contentmanagement status operation typically these are written to a http.Request
*/
type GetContentmanagementStatusParams struct {

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

// WithTimeout adds the timeout to the get contentmanagement status params
func (o *GetContentmanagementStatusParams) WithTimeout(timeout time.Duration) *GetContentmanagementStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get contentmanagement status params
func (o *GetContentmanagementStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get contentmanagement status params
func (o *GetContentmanagementStatusParams) WithContext(ctx context.Context) *GetContentmanagementStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get contentmanagement status params
func (o *GetContentmanagementStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get contentmanagement status params
func (o *GetContentmanagementStatusParams) WithHTTPClient(client *http.Client) *GetContentmanagementStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get contentmanagement status params
func (o *GetContentmanagementStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the get contentmanagement status params
func (o *GetContentmanagementStatusParams) WithPageNumber(pageNumber *int32) *GetContentmanagementStatusParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get contentmanagement status params
func (o *GetContentmanagementStatusParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get contentmanagement status params
func (o *GetContentmanagementStatusParams) WithPageSize(pageSize *int32) *GetContentmanagementStatusParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get contentmanagement status params
func (o *GetContentmanagementStatusParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetContentmanagementStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
