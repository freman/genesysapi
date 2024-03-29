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

// NewGetUserRolesParams creates a new GetUserRolesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUserRolesParams() *GetUserRolesParams {
	return &GetUserRolesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserRolesParamsWithTimeout creates a new GetUserRolesParams object
// with the ability to set a timeout on a request.
func NewGetUserRolesParamsWithTimeout(timeout time.Duration) *GetUserRolesParams {
	return &GetUserRolesParams{
		timeout: timeout,
	}
}

// NewGetUserRolesParamsWithContext creates a new GetUserRolesParams object
// with the ability to set a context for a request.
func NewGetUserRolesParamsWithContext(ctx context.Context) *GetUserRolesParams {
	return &GetUserRolesParams{
		Context: ctx,
	}
}

// NewGetUserRolesParamsWithHTTPClient creates a new GetUserRolesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUserRolesParamsWithHTTPClient(client *http.Client) *GetUserRolesParams {
	return &GetUserRolesParams{
		HTTPClient: client,
	}
}

/*
GetUserRolesParams contains all the parameters to send to the API endpoint

	for the get user roles operation.

	Typically these are written to a http.Request.
*/
type GetUserRolesParams struct {

	/* SubjectID.

	   User ID
	*/
	SubjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get user roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserRolesParams) WithDefaults() *GetUserRolesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get user roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserRolesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get user roles params
func (o *GetUserRolesParams) WithTimeout(timeout time.Duration) *GetUserRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user roles params
func (o *GetUserRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user roles params
func (o *GetUserRolesParams) WithContext(ctx context.Context) *GetUserRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user roles params
func (o *GetUserRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user roles params
func (o *GetUserRolesParams) WithHTTPClient(client *http.Client) *GetUserRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user roles params
func (o *GetUserRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSubjectID adds the subjectID to the get user roles params
func (o *GetUserRolesParams) WithSubjectID(subjectID string) *GetUserRolesParams {
	o.SetSubjectID(subjectID)
	return o
}

// SetSubjectID adds the subjectId to the get user roles params
func (o *GetUserRolesParams) SetSubjectID(subjectID string) {
	o.SubjectID = subjectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param subjectId
	if err := r.SetPathParam("subjectId", o.SubjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
