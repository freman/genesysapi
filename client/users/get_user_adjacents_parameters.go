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
	"github.com/go-openapi/swag"
)

// NewGetUserAdjacentsParams creates a new GetUserAdjacentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUserAdjacentsParams() *GetUserAdjacentsParams {
	return &GetUserAdjacentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserAdjacentsParamsWithTimeout creates a new GetUserAdjacentsParams object
// with the ability to set a timeout on a request.
func NewGetUserAdjacentsParamsWithTimeout(timeout time.Duration) *GetUserAdjacentsParams {
	return &GetUserAdjacentsParams{
		timeout: timeout,
	}
}

// NewGetUserAdjacentsParamsWithContext creates a new GetUserAdjacentsParams object
// with the ability to set a context for a request.
func NewGetUserAdjacentsParamsWithContext(ctx context.Context) *GetUserAdjacentsParams {
	return &GetUserAdjacentsParams{
		Context: ctx,
	}
}

// NewGetUserAdjacentsParamsWithHTTPClient creates a new GetUserAdjacentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUserAdjacentsParamsWithHTTPClient(client *http.Client) *GetUserAdjacentsParams {
	return &GetUserAdjacentsParams{
		HTTPClient: client,
	}
}

/*
GetUserAdjacentsParams contains all the parameters to send to the API endpoint

	for the get user adjacents operation.

	Typically these are written to a http.Request.
*/
type GetUserAdjacentsParams struct {

	/* Expand.

	   Which fields, if any, to expand
	*/
	Expand []string

	/* UserID.

	   User ID
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get user adjacents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserAdjacentsParams) WithDefaults() *GetUserAdjacentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get user adjacents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserAdjacentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get user adjacents params
func (o *GetUserAdjacentsParams) WithTimeout(timeout time.Duration) *GetUserAdjacentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user adjacents params
func (o *GetUserAdjacentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user adjacents params
func (o *GetUserAdjacentsParams) WithContext(ctx context.Context) *GetUserAdjacentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user adjacents params
func (o *GetUserAdjacentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user adjacents params
func (o *GetUserAdjacentsParams) WithHTTPClient(client *http.Client) *GetUserAdjacentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user adjacents params
func (o *GetUserAdjacentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get user adjacents params
func (o *GetUserAdjacentsParams) WithExpand(expand []string) *GetUserAdjacentsParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get user adjacents params
func (o *GetUserAdjacentsParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithUserID adds the userID to the get user adjacents params
func (o *GetUserAdjacentsParams) WithUserID(userID string) *GetUserAdjacentsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get user adjacents params
func (o *GetUserAdjacentsParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserAdjacentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Expand != nil {

		// binding items for expand
		joinedExpand := o.bindParamExpand(reg)

		// query array param expand
		if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
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

// bindParamGetUserAdjacents binds the parameter expand
func (o *GetUserAdjacentsParams) bindParamExpand(formats strfmt.Registry) []string {
	expandIR := o.Expand

	var expandIC []string
	for _, expandIIR := range expandIR { // explode []string

		expandIIV := expandIIR // string as string
		expandIC = append(expandIC, expandIIV)
	}

	// items.CollectionFormat: "multi"
	expandIS := swag.JoinByFormat(expandIC, "multi")

	return expandIS
}
