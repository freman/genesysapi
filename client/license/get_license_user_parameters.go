// Code generated by go-swagger; DO NOT EDIT.

package license

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

// NewGetLicenseUserParams creates a new GetLicenseUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLicenseUserParams() *GetLicenseUserParams {
	return &GetLicenseUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLicenseUserParamsWithTimeout creates a new GetLicenseUserParams object
// with the ability to set a timeout on a request.
func NewGetLicenseUserParamsWithTimeout(timeout time.Duration) *GetLicenseUserParams {
	return &GetLicenseUserParams{
		timeout: timeout,
	}
}

// NewGetLicenseUserParamsWithContext creates a new GetLicenseUserParams object
// with the ability to set a context for a request.
func NewGetLicenseUserParamsWithContext(ctx context.Context) *GetLicenseUserParams {
	return &GetLicenseUserParams{
		Context: ctx,
	}
}

// NewGetLicenseUserParamsWithHTTPClient creates a new GetLicenseUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLicenseUserParamsWithHTTPClient(client *http.Client) *GetLicenseUserParams {
	return &GetLicenseUserParams{
		HTTPClient: client,
	}
}

/*
GetLicenseUserParams contains all the parameters to send to the API endpoint

	for the get license user operation.

	Typically these are written to a http.Request.
*/
type GetLicenseUserParams struct {

	/* UserID.

	   ID
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get license user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLicenseUserParams) WithDefaults() *GetLicenseUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get license user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLicenseUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get license user params
func (o *GetLicenseUserParams) WithTimeout(timeout time.Duration) *GetLicenseUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get license user params
func (o *GetLicenseUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get license user params
func (o *GetLicenseUserParams) WithContext(ctx context.Context) *GetLicenseUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get license user params
func (o *GetLicenseUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get license user params
func (o *GetLicenseUserParams) WithHTTPClient(client *http.Client) *GetLicenseUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get license user params
func (o *GetLicenseUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the get license user params
func (o *GetLicenseUserParams) WithUserID(userID string) *GetLicenseUserParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get license user params
func (o *GetLicenseUserParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLicenseUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
