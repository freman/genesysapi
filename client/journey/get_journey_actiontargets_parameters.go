// Code generated by go-swagger; DO NOT EDIT.

package journey

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

// NewGetJourneyActiontargetsParams creates a new GetJourneyActiontargetsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetJourneyActiontargetsParams() *GetJourneyActiontargetsParams {
	return &GetJourneyActiontargetsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetJourneyActiontargetsParamsWithTimeout creates a new GetJourneyActiontargetsParams object
// with the ability to set a timeout on a request.
func NewGetJourneyActiontargetsParamsWithTimeout(timeout time.Duration) *GetJourneyActiontargetsParams {
	return &GetJourneyActiontargetsParams{
		timeout: timeout,
	}
}

// NewGetJourneyActiontargetsParamsWithContext creates a new GetJourneyActiontargetsParams object
// with the ability to set a context for a request.
func NewGetJourneyActiontargetsParamsWithContext(ctx context.Context) *GetJourneyActiontargetsParams {
	return &GetJourneyActiontargetsParams{
		Context: ctx,
	}
}

// NewGetJourneyActiontargetsParamsWithHTTPClient creates a new GetJourneyActiontargetsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetJourneyActiontargetsParamsWithHTTPClient(client *http.Client) *GetJourneyActiontargetsParams {
	return &GetJourneyActiontargetsParams{
		HTTPClient: client,
	}
}

/*
GetJourneyActiontargetsParams contains all the parameters to send to the API endpoint

	for the get journey actiontargets operation.

	Typically these are written to a http.Request.
*/
type GetJourneyActiontargetsParams struct {

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

// WithDefaults hydrates default values in the get journey actiontargets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetJourneyActiontargetsParams) WithDefaults() *GetJourneyActiontargetsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get journey actiontargets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetJourneyActiontargetsParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)
	)

	val := GetJourneyActiontargetsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get journey actiontargets params
func (o *GetJourneyActiontargetsParams) WithTimeout(timeout time.Duration) *GetJourneyActiontargetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get journey actiontargets params
func (o *GetJourneyActiontargetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get journey actiontargets params
func (o *GetJourneyActiontargetsParams) WithContext(ctx context.Context) *GetJourneyActiontargetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get journey actiontargets params
func (o *GetJourneyActiontargetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get journey actiontargets params
func (o *GetJourneyActiontargetsParams) WithHTTPClient(client *http.Client) *GetJourneyActiontargetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get journey actiontargets params
func (o *GetJourneyActiontargetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the get journey actiontargets params
func (o *GetJourneyActiontargetsParams) WithPageNumber(pageNumber *int32) *GetJourneyActiontargetsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get journey actiontargets params
func (o *GetJourneyActiontargetsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get journey actiontargets params
func (o *GetJourneyActiontargetsParams) WithPageSize(pageSize *int32) *GetJourneyActiontargetsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get journey actiontargets params
func (o *GetJourneyActiontargetsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetJourneyActiontargetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
