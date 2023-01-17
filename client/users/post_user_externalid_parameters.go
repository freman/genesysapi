// Code generated by go-swagger; DO NOT EDIT.

package users

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

	"github.com/freman/genesysapi/models"
)

// NewPostUserExternalidParams creates a new PostUserExternalidParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostUserExternalidParams() *PostUserExternalidParams {
	return &PostUserExternalidParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostUserExternalidParamsWithTimeout creates a new PostUserExternalidParams object
// with the ability to set a timeout on a request.
func NewPostUserExternalidParamsWithTimeout(timeout time.Duration) *PostUserExternalidParams {
	return &PostUserExternalidParams{
		timeout: timeout,
	}
}

// NewPostUserExternalidParamsWithContext creates a new PostUserExternalidParams object
// with the ability to set a context for a request.
func NewPostUserExternalidParamsWithContext(ctx context.Context) *PostUserExternalidParams {
	return &PostUserExternalidParams{
		Context: ctx,
	}
}

// NewPostUserExternalidParamsWithHTTPClient creates a new PostUserExternalidParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostUserExternalidParamsWithHTTPClient(client *http.Client) *PostUserExternalidParams {
	return &PostUserExternalidParams{
		HTTPClient: client,
	}
}

/*
PostUserExternalidParams contains all the parameters to send to the API endpoint

	for the post user externalid operation.

	Typically these are written to a http.Request.
*/
type PostUserExternalidParams struct {

	// Body.
	Body *models.UserExternalIdentifier

	/* UserID.

	   User ID
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post user externalid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUserExternalidParams) WithDefaults() *PostUserExternalidParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post user externalid params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUserExternalidParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post user externalid params
func (o *PostUserExternalidParams) WithTimeout(timeout time.Duration) *PostUserExternalidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post user externalid params
func (o *PostUserExternalidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post user externalid params
func (o *PostUserExternalidParams) WithContext(ctx context.Context) *PostUserExternalidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post user externalid params
func (o *PostUserExternalidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post user externalid params
func (o *PostUserExternalidParams) WithHTTPClient(client *http.Client) *PostUserExternalidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post user externalid params
func (o *PostUserExternalidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post user externalid params
func (o *PostUserExternalidParams) WithBody(body *models.UserExternalIdentifier) *PostUserExternalidParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post user externalid params
func (o *PostUserExternalidParams) SetBody(body *models.UserExternalIdentifier) {
	o.Body = body
}

// WithUserID adds the userID to the post user externalid params
func (o *PostUserExternalidParams) WithUserID(userID string) *PostUserExternalidParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the post user externalid params
func (o *PostUserExternalidParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PostUserExternalidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}