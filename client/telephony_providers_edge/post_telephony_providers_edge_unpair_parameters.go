// Code generated by go-swagger; DO NOT EDIT.

package telephony_providers_edge

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

// NewPostTelephonyProvidersEdgeUnpairParams creates a new PostTelephonyProvidersEdgeUnpairParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostTelephonyProvidersEdgeUnpairParams() *PostTelephonyProvidersEdgeUnpairParams {
	return &PostTelephonyProvidersEdgeUnpairParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostTelephonyProvidersEdgeUnpairParamsWithTimeout creates a new PostTelephonyProvidersEdgeUnpairParams object
// with the ability to set a timeout on a request.
func NewPostTelephonyProvidersEdgeUnpairParamsWithTimeout(timeout time.Duration) *PostTelephonyProvidersEdgeUnpairParams {
	return &PostTelephonyProvidersEdgeUnpairParams{
		timeout: timeout,
	}
}

// NewPostTelephonyProvidersEdgeUnpairParamsWithContext creates a new PostTelephonyProvidersEdgeUnpairParams object
// with the ability to set a context for a request.
func NewPostTelephonyProvidersEdgeUnpairParamsWithContext(ctx context.Context) *PostTelephonyProvidersEdgeUnpairParams {
	return &PostTelephonyProvidersEdgeUnpairParams{
		Context: ctx,
	}
}

// NewPostTelephonyProvidersEdgeUnpairParamsWithHTTPClient creates a new PostTelephonyProvidersEdgeUnpairParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostTelephonyProvidersEdgeUnpairParamsWithHTTPClient(client *http.Client) *PostTelephonyProvidersEdgeUnpairParams {
	return &PostTelephonyProvidersEdgeUnpairParams{
		HTTPClient: client,
	}
}

/*
PostTelephonyProvidersEdgeUnpairParams contains all the parameters to send to the API endpoint

	for the post telephony providers edge unpair operation.

	Typically these are written to a http.Request.
*/
type PostTelephonyProvidersEdgeUnpairParams struct {

	/* EdgeID.

	   Edge Id
	*/
	EdgeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post telephony providers edge unpair params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostTelephonyProvidersEdgeUnpairParams) WithDefaults() *PostTelephonyProvidersEdgeUnpairParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post telephony providers edge unpair params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostTelephonyProvidersEdgeUnpairParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post telephony providers edge unpair params
func (o *PostTelephonyProvidersEdgeUnpairParams) WithTimeout(timeout time.Duration) *PostTelephonyProvidersEdgeUnpairParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post telephony providers edge unpair params
func (o *PostTelephonyProvidersEdgeUnpairParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post telephony providers edge unpair params
func (o *PostTelephonyProvidersEdgeUnpairParams) WithContext(ctx context.Context) *PostTelephonyProvidersEdgeUnpairParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post telephony providers edge unpair params
func (o *PostTelephonyProvidersEdgeUnpairParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post telephony providers edge unpair params
func (o *PostTelephonyProvidersEdgeUnpairParams) WithHTTPClient(client *http.Client) *PostTelephonyProvidersEdgeUnpairParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post telephony providers edge unpair params
func (o *PostTelephonyProvidersEdgeUnpairParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEdgeID adds the edgeID to the post telephony providers edge unpair params
func (o *PostTelephonyProvidersEdgeUnpairParams) WithEdgeID(edgeID string) *PostTelephonyProvidersEdgeUnpairParams {
	o.SetEdgeID(edgeID)
	return o
}

// SetEdgeID adds the edgeId to the post telephony providers edge unpair params
func (o *PostTelephonyProvidersEdgeUnpairParams) SetEdgeID(edgeID string) {
	o.EdgeID = edgeID
}

// WriteToRequest writes these params to a swagger request
func (o *PostTelephonyProvidersEdgeUnpairParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param edgeId
	if err := r.SetPathParam("edgeId", o.EdgeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
