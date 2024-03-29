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

// NewGetConversationsMessagingSupportedcontentParams creates a new GetConversationsMessagingSupportedcontentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConversationsMessagingSupportedcontentParams() *GetConversationsMessagingSupportedcontentParams {
	return &GetConversationsMessagingSupportedcontentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConversationsMessagingSupportedcontentParamsWithTimeout creates a new GetConversationsMessagingSupportedcontentParams object
// with the ability to set a timeout on a request.
func NewGetConversationsMessagingSupportedcontentParamsWithTimeout(timeout time.Duration) *GetConversationsMessagingSupportedcontentParams {
	return &GetConversationsMessagingSupportedcontentParams{
		timeout: timeout,
	}
}

// NewGetConversationsMessagingSupportedcontentParamsWithContext creates a new GetConversationsMessagingSupportedcontentParams object
// with the ability to set a context for a request.
func NewGetConversationsMessagingSupportedcontentParamsWithContext(ctx context.Context) *GetConversationsMessagingSupportedcontentParams {
	return &GetConversationsMessagingSupportedcontentParams{
		Context: ctx,
	}
}

// NewGetConversationsMessagingSupportedcontentParamsWithHTTPClient creates a new GetConversationsMessagingSupportedcontentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetConversationsMessagingSupportedcontentParamsWithHTTPClient(client *http.Client) *GetConversationsMessagingSupportedcontentParams {
	return &GetConversationsMessagingSupportedcontentParams{
		HTTPClient: client,
	}
}

/*
GetConversationsMessagingSupportedcontentParams contains all the parameters to send to the API endpoint

	for the get conversations messaging supportedcontent operation.

	Typically these are written to a http.Request.
*/
type GetConversationsMessagingSupportedcontentParams struct {

	/* PageNumber.

	   Page number

	   Format: int32
	   Default: 1
	*/
	PageNumber *int32

	/* PageSize.

	   Page size

	   Format: int32
	   Default: 25
	*/
	PageSize *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get conversations messaging supportedcontent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConversationsMessagingSupportedcontentParams) WithDefaults() *GetConversationsMessagingSupportedcontentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get conversations messaging supportedcontent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConversationsMessagingSupportedcontentParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)
	)

	val := GetConversationsMessagingSupportedcontentParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get conversations messaging supportedcontent params
func (o *GetConversationsMessagingSupportedcontentParams) WithTimeout(timeout time.Duration) *GetConversationsMessagingSupportedcontentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get conversations messaging supportedcontent params
func (o *GetConversationsMessagingSupportedcontentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get conversations messaging supportedcontent params
func (o *GetConversationsMessagingSupportedcontentParams) WithContext(ctx context.Context) *GetConversationsMessagingSupportedcontentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get conversations messaging supportedcontent params
func (o *GetConversationsMessagingSupportedcontentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get conversations messaging supportedcontent params
func (o *GetConversationsMessagingSupportedcontentParams) WithHTTPClient(client *http.Client) *GetConversationsMessagingSupportedcontentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get conversations messaging supportedcontent params
func (o *GetConversationsMessagingSupportedcontentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the get conversations messaging supportedcontent params
func (o *GetConversationsMessagingSupportedcontentParams) WithPageNumber(pageNumber *int32) *GetConversationsMessagingSupportedcontentParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get conversations messaging supportedcontent params
func (o *GetConversationsMessagingSupportedcontentParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get conversations messaging supportedcontent params
func (o *GetConversationsMessagingSupportedcontentParams) WithPageSize(pageSize *int32) *GetConversationsMessagingSupportedcontentParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get conversations messaging supportedcontent params
func (o *GetConversationsMessagingSupportedcontentParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetConversationsMessagingSupportedcontentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
