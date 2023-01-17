// Code generated by go-swagger; DO NOT EDIT.

package authorization

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

// NewGetAuthorizationDivisionParams creates a new GetAuthorizationDivisionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAuthorizationDivisionParams() *GetAuthorizationDivisionParams {
	return &GetAuthorizationDivisionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthorizationDivisionParamsWithTimeout creates a new GetAuthorizationDivisionParams object
// with the ability to set a timeout on a request.
func NewGetAuthorizationDivisionParamsWithTimeout(timeout time.Duration) *GetAuthorizationDivisionParams {
	return &GetAuthorizationDivisionParams{
		timeout: timeout,
	}
}

// NewGetAuthorizationDivisionParamsWithContext creates a new GetAuthorizationDivisionParams object
// with the ability to set a context for a request.
func NewGetAuthorizationDivisionParamsWithContext(ctx context.Context) *GetAuthorizationDivisionParams {
	return &GetAuthorizationDivisionParams{
		Context: ctx,
	}
}

// NewGetAuthorizationDivisionParamsWithHTTPClient creates a new GetAuthorizationDivisionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAuthorizationDivisionParamsWithHTTPClient(client *http.Client) *GetAuthorizationDivisionParams {
	return &GetAuthorizationDivisionParams{
		HTTPClient: client,
	}
}

/*
GetAuthorizationDivisionParams contains all the parameters to send to the API endpoint

	for the get authorization division operation.

	Typically these are written to a http.Request.
*/
type GetAuthorizationDivisionParams struct {

	/* DivisionID.

	   Division ID
	*/
	DivisionID string

	/* ObjectCount.

	   Get count of objects in this division, grouped by type
	*/
	ObjectCount *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get authorization division params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAuthorizationDivisionParams) WithDefaults() *GetAuthorizationDivisionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get authorization division params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAuthorizationDivisionParams) SetDefaults() {
	var (
		objectCountDefault = bool(false)
	)

	val := GetAuthorizationDivisionParams{
		ObjectCount: &objectCountDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get authorization division params
func (o *GetAuthorizationDivisionParams) WithTimeout(timeout time.Duration) *GetAuthorizationDivisionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get authorization division params
func (o *GetAuthorizationDivisionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get authorization division params
func (o *GetAuthorizationDivisionParams) WithContext(ctx context.Context) *GetAuthorizationDivisionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get authorization division params
func (o *GetAuthorizationDivisionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get authorization division params
func (o *GetAuthorizationDivisionParams) WithHTTPClient(client *http.Client) *GetAuthorizationDivisionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get authorization division params
func (o *GetAuthorizationDivisionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDivisionID adds the divisionID to the get authorization division params
func (o *GetAuthorizationDivisionParams) WithDivisionID(divisionID string) *GetAuthorizationDivisionParams {
	o.SetDivisionID(divisionID)
	return o
}

// SetDivisionID adds the divisionId to the get authorization division params
func (o *GetAuthorizationDivisionParams) SetDivisionID(divisionID string) {
	o.DivisionID = divisionID
}

// WithObjectCount adds the objectCount to the get authorization division params
func (o *GetAuthorizationDivisionParams) WithObjectCount(objectCount *bool) *GetAuthorizationDivisionParams {
	o.SetObjectCount(objectCount)
	return o
}

// SetObjectCount adds the objectCount to the get authorization division params
func (o *GetAuthorizationDivisionParams) SetObjectCount(objectCount *bool) {
	o.ObjectCount = objectCount
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthorizationDivisionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param divisionId
	if err := r.SetPathParam("divisionId", o.DivisionID); err != nil {
		return err
	}

	if o.ObjectCount != nil {

		// query param objectCount
		var qrObjectCount bool

		if o.ObjectCount != nil {
			qrObjectCount = *o.ObjectCount
		}
		qObjectCount := swag.FormatBool(qrObjectCount)
		if qObjectCount != "" {

			if err := r.SetQueryParam("objectCount", qObjectCount); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
