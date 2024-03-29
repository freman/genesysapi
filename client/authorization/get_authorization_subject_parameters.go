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

// NewGetAuthorizationSubjectParams creates a new GetAuthorizationSubjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAuthorizationSubjectParams() *GetAuthorizationSubjectParams {
	return &GetAuthorizationSubjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthorizationSubjectParamsWithTimeout creates a new GetAuthorizationSubjectParams object
// with the ability to set a timeout on a request.
func NewGetAuthorizationSubjectParamsWithTimeout(timeout time.Duration) *GetAuthorizationSubjectParams {
	return &GetAuthorizationSubjectParams{
		timeout: timeout,
	}
}

// NewGetAuthorizationSubjectParamsWithContext creates a new GetAuthorizationSubjectParams object
// with the ability to set a context for a request.
func NewGetAuthorizationSubjectParamsWithContext(ctx context.Context) *GetAuthorizationSubjectParams {
	return &GetAuthorizationSubjectParams{
		Context: ctx,
	}
}

// NewGetAuthorizationSubjectParamsWithHTTPClient creates a new GetAuthorizationSubjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAuthorizationSubjectParamsWithHTTPClient(client *http.Client) *GetAuthorizationSubjectParams {
	return &GetAuthorizationSubjectParams{
		HTTPClient: client,
	}
}

/*
GetAuthorizationSubjectParams contains all the parameters to send to the API endpoint

	for the get authorization subject operation.

	Typically these are written to a http.Request.
*/
type GetAuthorizationSubjectParams struct {

	/* SubjectID.

	   Subject ID (user or group)
	*/
	SubjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get authorization subject params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAuthorizationSubjectParams) WithDefaults() *GetAuthorizationSubjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get authorization subject params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAuthorizationSubjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get authorization subject params
func (o *GetAuthorizationSubjectParams) WithTimeout(timeout time.Duration) *GetAuthorizationSubjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get authorization subject params
func (o *GetAuthorizationSubjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get authorization subject params
func (o *GetAuthorizationSubjectParams) WithContext(ctx context.Context) *GetAuthorizationSubjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get authorization subject params
func (o *GetAuthorizationSubjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get authorization subject params
func (o *GetAuthorizationSubjectParams) WithHTTPClient(client *http.Client) *GetAuthorizationSubjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get authorization subject params
func (o *GetAuthorizationSubjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSubjectID adds the subjectID to the get authorization subject params
func (o *GetAuthorizationSubjectParams) WithSubjectID(subjectID string) *GetAuthorizationSubjectParams {
	o.SetSubjectID(subjectID)
	return o
}

// SetSubjectID adds the subjectId to the get authorization subject params
func (o *GetAuthorizationSubjectParams) SetSubjectID(subjectID string) {
	o.SubjectID = subjectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthorizationSubjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
