// Code generated by go-swagger; DO NOT EDIT.

package web_messaging

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

// NewGetWebmessagingMessagesParams creates a new GetWebmessagingMessagesParams object
// with the default values initialized.
func NewGetWebmessagingMessagesParams() *GetWebmessagingMessagesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetWebmessagingMessagesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWebmessagingMessagesParamsWithTimeout creates a new GetWebmessagingMessagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWebmessagingMessagesParamsWithTimeout(timeout time.Duration) *GetWebmessagingMessagesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetWebmessagingMessagesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetWebmessagingMessagesParamsWithContext creates a new GetWebmessagingMessagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWebmessagingMessagesParamsWithContext(ctx context.Context) *GetWebmessagingMessagesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetWebmessagingMessagesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetWebmessagingMessagesParamsWithHTTPClient creates a new GetWebmessagingMessagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWebmessagingMessagesParamsWithHTTPClient(client *http.Client) *GetWebmessagingMessagesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetWebmessagingMessagesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetWebmessagingMessagesParams contains all the parameters to send to the API endpoint
for the get webmessaging messages operation typically these are written to a http.Request
*/
type GetWebmessagingMessagesParams struct {

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

// WithTimeout adds the timeout to the get webmessaging messages params
func (o *GetWebmessagingMessagesParams) WithTimeout(timeout time.Duration) *GetWebmessagingMessagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get webmessaging messages params
func (o *GetWebmessagingMessagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get webmessaging messages params
func (o *GetWebmessagingMessagesParams) WithContext(ctx context.Context) *GetWebmessagingMessagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get webmessaging messages params
func (o *GetWebmessagingMessagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get webmessaging messages params
func (o *GetWebmessagingMessagesParams) WithHTTPClient(client *http.Client) *GetWebmessagingMessagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get webmessaging messages params
func (o *GetWebmessagingMessagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the get webmessaging messages params
func (o *GetWebmessagingMessagesParams) WithPageNumber(pageNumber *int32) *GetWebmessagingMessagesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get webmessaging messages params
func (o *GetWebmessagingMessagesParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get webmessaging messages params
func (o *GetWebmessagingMessagesParams) WithPageSize(pageSize *int32) *GetWebmessagingMessagesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get webmessaging messages params
func (o *GetWebmessagingMessagesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetWebmessagingMessagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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