// Code generated by go-swagger; DO NOT EDIT.

package routing

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

// NewGetRoutingMessageRecipientsParams creates a new GetRoutingMessageRecipientsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRoutingMessageRecipientsParams() *GetRoutingMessageRecipientsParams {
	return &GetRoutingMessageRecipientsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoutingMessageRecipientsParamsWithTimeout creates a new GetRoutingMessageRecipientsParams object
// with the ability to set a timeout on a request.
func NewGetRoutingMessageRecipientsParamsWithTimeout(timeout time.Duration) *GetRoutingMessageRecipientsParams {
	return &GetRoutingMessageRecipientsParams{
		timeout: timeout,
	}
}

// NewGetRoutingMessageRecipientsParamsWithContext creates a new GetRoutingMessageRecipientsParams object
// with the ability to set a context for a request.
func NewGetRoutingMessageRecipientsParamsWithContext(ctx context.Context) *GetRoutingMessageRecipientsParams {
	return &GetRoutingMessageRecipientsParams{
		Context: ctx,
	}
}

// NewGetRoutingMessageRecipientsParamsWithHTTPClient creates a new GetRoutingMessageRecipientsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRoutingMessageRecipientsParamsWithHTTPClient(client *http.Client) *GetRoutingMessageRecipientsParams {
	return &GetRoutingMessageRecipientsParams{
		HTTPClient: client,
	}
}

/*
GetRoutingMessageRecipientsParams contains all the parameters to send to the API endpoint

	for the get routing message recipients operation.

	Typically these are written to a http.Request.
*/
type GetRoutingMessageRecipientsParams struct {

	/* MessengerType.

	   Messenger Type
	*/
	MessengerType *string

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

// WithDefaults hydrates default values in the get routing message recipients params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoutingMessageRecipientsParams) WithDefaults() *GetRoutingMessageRecipientsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get routing message recipients params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoutingMessageRecipientsParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)
	)

	val := GetRoutingMessageRecipientsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get routing message recipients params
func (o *GetRoutingMessageRecipientsParams) WithTimeout(timeout time.Duration) *GetRoutingMessageRecipientsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get routing message recipients params
func (o *GetRoutingMessageRecipientsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get routing message recipients params
func (o *GetRoutingMessageRecipientsParams) WithContext(ctx context.Context) *GetRoutingMessageRecipientsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get routing message recipients params
func (o *GetRoutingMessageRecipientsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get routing message recipients params
func (o *GetRoutingMessageRecipientsParams) WithHTTPClient(client *http.Client) *GetRoutingMessageRecipientsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get routing message recipients params
func (o *GetRoutingMessageRecipientsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMessengerType adds the messengerType to the get routing message recipients params
func (o *GetRoutingMessageRecipientsParams) WithMessengerType(messengerType *string) *GetRoutingMessageRecipientsParams {
	o.SetMessengerType(messengerType)
	return o
}

// SetMessengerType adds the messengerType to the get routing message recipients params
func (o *GetRoutingMessageRecipientsParams) SetMessengerType(messengerType *string) {
	o.MessengerType = messengerType
}

// WithPageNumber adds the pageNumber to the get routing message recipients params
func (o *GetRoutingMessageRecipientsParams) WithPageNumber(pageNumber *int32) *GetRoutingMessageRecipientsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get routing message recipients params
func (o *GetRoutingMessageRecipientsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get routing message recipients params
func (o *GetRoutingMessageRecipientsParams) WithPageSize(pageSize *int32) *GetRoutingMessageRecipientsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get routing message recipients params
func (o *GetRoutingMessageRecipientsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoutingMessageRecipientsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.MessengerType != nil {

		// query param messengerType
		var qrMessengerType string

		if o.MessengerType != nil {
			qrMessengerType = *o.MessengerType
		}
		qMessengerType := qrMessengerType
		if qMessengerType != "" {

			if err := r.SetQueryParam("messengerType", qMessengerType); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
