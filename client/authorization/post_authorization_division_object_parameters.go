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
)

// NewPostAuthorizationDivisionObjectParams creates a new PostAuthorizationDivisionObjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAuthorizationDivisionObjectParams() *PostAuthorizationDivisionObjectParams {
	return &PostAuthorizationDivisionObjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAuthorizationDivisionObjectParamsWithTimeout creates a new PostAuthorizationDivisionObjectParams object
// with the ability to set a timeout on a request.
func NewPostAuthorizationDivisionObjectParamsWithTimeout(timeout time.Duration) *PostAuthorizationDivisionObjectParams {
	return &PostAuthorizationDivisionObjectParams{
		timeout: timeout,
	}
}

// NewPostAuthorizationDivisionObjectParamsWithContext creates a new PostAuthorizationDivisionObjectParams object
// with the ability to set a context for a request.
func NewPostAuthorizationDivisionObjectParamsWithContext(ctx context.Context) *PostAuthorizationDivisionObjectParams {
	return &PostAuthorizationDivisionObjectParams{
		Context: ctx,
	}
}

// NewPostAuthorizationDivisionObjectParamsWithHTTPClient creates a new PostAuthorizationDivisionObjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAuthorizationDivisionObjectParamsWithHTTPClient(client *http.Client) *PostAuthorizationDivisionObjectParams {
	return &PostAuthorizationDivisionObjectParams{
		HTTPClient: client,
	}
}

/*
PostAuthorizationDivisionObjectParams contains all the parameters to send to the API endpoint

	for the post authorization division object operation.

	Typically these are written to a http.Request.
*/
type PostAuthorizationDivisionObjectParams struct {

	/* Body.

	   Object Id List
	*/
	Body []string

	/* DivisionID.

	   Division ID
	*/
	DivisionID string

	/* ObjectType.

	   The type of the objects. Must be one of the valid object types
	*/
	ObjectType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post authorization division object params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAuthorizationDivisionObjectParams) WithDefaults() *PostAuthorizationDivisionObjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post authorization division object params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAuthorizationDivisionObjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post authorization division object params
func (o *PostAuthorizationDivisionObjectParams) WithTimeout(timeout time.Duration) *PostAuthorizationDivisionObjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post authorization division object params
func (o *PostAuthorizationDivisionObjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post authorization division object params
func (o *PostAuthorizationDivisionObjectParams) WithContext(ctx context.Context) *PostAuthorizationDivisionObjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post authorization division object params
func (o *PostAuthorizationDivisionObjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post authorization division object params
func (o *PostAuthorizationDivisionObjectParams) WithHTTPClient(client *http.Client) *PostAuthorizationDivisionObjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post authorization division object params
func (o *PostAuthorizationDivisionObjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post authorization division object params
func (o *PostAuthorizationDivisionObjectParams) WithBody(body []string) *PostAuthorizationDivisionObjectParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post authorization division object params
func (o *PostAuthorizationDivisionObjectParams) SetBody(body []string) {
	o.Body = body
}

// WithDivisionID adds the divisionID to the post authorization division object params
func (o *PostAuthorizationDivisionObjectParams) WithDivisionID(divisionID string) *PostAuthorizationDivisionObjectParams {
	o.SetDivisionID(divisionID)
	return o
}

// SetDivisionID adds the divisionId to the post authorization division object params
func (o *PostAuthorizationDivisionObjectParams) SetDivisionID(divisionID string) {
	o.DivisionID = divisionID
}

// WithObjectType adds the objectType to the post authorization division object params
func (o *PostAuthorizationDivisionObjectParams) WithObjectType(objectType string) *PostAuthorizationDivisionObjectParams {
	o.SetObjectType(objectType)
	return o
}

// SetObjectType adds the objectType to the post authorization division object params
func (o *PostAuthorizationDivisionObjectParams) SetObjectType(objectType string) {
	o.ObjectType = objectType
}

// WriteToRequest writes these params to a swagger request
func (o *PostAuthorizationDivisionObjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param divisionId
	if err := r.SetPathParam("divisionId", o.DivisionID); err != nil {
		return err
	}

	// path param objectType
	if err := r.SetPathParam("objectType", o.ObjectType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
