// Code generated by go-swagger; DO NOT EDIT.

package greetings

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

// NewGetUserGreetingsDefaultsParams creates a new GetUserGreetingsDefaultsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUserGreetingsDefaultsParams() *GetUserGreetingsDefaultsParams {
	return &GetUserGreetingsDefaultsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserGreetingsDefaultsParamsWithTimeout creates a new GetUserGreetingsDefaultsParams object
// with the ability to set a timeout on a request.
func NewGetUserGreetingsDefaultsParamsWithTimeout(timeout time.Duration) *GetUserGreetingsDefaultsParams {
	return &GetUserGreetingsDefaultsParams{
		timeout: timeout,
	}
}

// NewGetUserGreetingsDefaultsParamsWithContext creates a new GetUserGreetingsDefaultsParams object
// with the ability to set a context for a request.
func NewGetUserGreetingsDefaultsParamsWithContext(ctx context.Context) *GetUserGreetingsDefaultsParams {
	return &GetUserGreetingsDefaultsParams{
		Context: ctx,
	}
}

// NewGetUserGreetingsDefaultsParamsWithHTTPClient creates a new GetUserGreetingsDefaultsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUserGreetingsDefaultsParamsWithHTTPClient(client *http.Client) *GetUserGreetingsDefaultsParams {
	return &GetUserGreetingsDefaultsParams{
		HTTPClient: client,
	}
}

/*
GetUserGreetingsDefaultsParams contains all the parameters to send to the API endpoint

	for the get user greetings defaults operation.

	Typically these are written to a http.Request.
*/
type GetUserGreetingsDefaultsParams struct {

	/* UserID.

	   User ID
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get user greetings defaults params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserGreetingsDefaultsParams) WithDefaults() *GetUserGreetingsDefaultsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get user greetings defaults params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserGreetingsDefaultsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get user greetings defaults params
func (o *GetUserGreetingsDefaultsParams) WithTimeout(timeout time.Duration) *GetUserGreetingsDefaultsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user greetings defaults params
func (o *GetUserGreetingsDefaultsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user greetings defaults params
func (o *GetUserGreetingsDefaultsParams) WithContext(ctx context.Context) *GetUserGreetingsDefaultsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user greetings defaults params
func (o *GetUserGreetingsDefaultsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user greetings defaults params
func (o *GetUserGreetingsDefaultsParams) WithHTTPClient(client *http.Client) *GetUserGreetingsDefaultsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user greetings defaults params
func (o *GetUserGreetingsDefaultsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the get user greetings defaults params
func (o *GetUserGreetingsDefaultsParams) WithUserID(userID string) *GetUserGreetingsDefaultsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get user greetings defaults params
func (o *GetUserGreetingsDefaultsParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserGreetingsDefaultsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
