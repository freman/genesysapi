// Code generated by go-swagger; DO NOT EDIT.

package uploads

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

// NewPostLanguageunderstandingMinerUploadsParams creates a new PostLanguageunderstandingMinerUploadsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostLanguageunderstandingMinerUploadsParams() *PostLanguageunderstandingMinerUploadsParams {
	return &PostLanguageunderstandingMinerUploadsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostLanguageunderstandingMinerUploadsParamsWithTimeout creates a new PostLanguageunderstandingMinerUploadsParams object
// with the ability to set a timeout on a request.
func NewPostLanguageunderstandingMinerUploadsParamsWithTimeout(timeout time.Duration) *PostLanguageunderstandingMinerUploadsParams {
	return &PostLanguageunderstandingMinerUploadsParams{
		timeout: timeout,
	}
}

// NewPostLanguageunderstandingMinerUploadsParamsWithContext creates a new PostLanguageunderstandingMinerUploadsParams object
// with the ability to set a context for a request.
func NewPostLanguageunderstandingMinerUploadsParamsWithContext(ctx context.Context) *PostLanguageunderstandingMinerUploadsParams {
	return &PostLanguageunderstandingMinerUploadsParams{
		Context: ctx,
	}
}

// NewPostLanguageunderstandingMinerUploadsParamsWithHTTPClient creates a new PostLanguageunderstandingMinerUploadsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostLanguageunderstandingMinerUploadsParamsWithHTTPClient(client *http.Client) *PostLanguageunderstandingMinerUploadsParams {
	return &PostLanguageunderstandingMinerUploadsParams{
		HTTPClient: client,
	}
}

/*
PostLanguageunderstandingMinerUploadsParams contains all the parameters to send to the API endpoint

	for the post languageunderstanding miner uploads operation.

	Typically these are written to a http.Request.
*/
type PostLanguageunderstandingMinerUploadsParams struct {

	/* Body.

	   query
	*/
	Body models.Empty

	/* MinerID.

	   Miner ID
	*/
	MinerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post languageunderstanding miner uploads params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLanguageunderstandingMinerUploadsParams) WithDefaults() *PostLanguageunderstandingMinerUploadsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post languageunderstanding miner uploads params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLanguageunderstandingMinerUploadsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post languageunderstanding miner uploads params
func (o *PostLanguageunderstandingMinerUploadsParams) WithTimeout(timeout time.Duration) *PostLanguageunderstandingMinerUploadsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post languageunderstanding miner uploads params
func (o *PostLanguageunderstandingMinerUploadsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post languageunderstanding miner uploads params
func (o *PostLanguageunderstandingMinerUploadsParams) WithContext(ctx context.Context) *PostLanguageunderstandingMinerUploadsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post languageunderstanding miner uploads params
func (o *PostLanguageunderstandingMinerUploadsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post languageunderstanding miner uploads params
func (o *PostLanguageunderstandingMinerUploadsParams) WithHTTPClient(client *http.Client) *PostLanguageunderstandingMinerUploadsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post languageunderstanding miner uploads params
func (o *PostLanguageunderstandingMinerUploadsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post languageunderstanding miner uploads params
func (o *PostLanguageunderstandingMinerUploadsParams) WithBody(body models.Empty) *PostLanguageunderstandingMinerUploadsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post languageunderstanding miner uploads params
func (o *PostLanguageunderstandingMinerUploadsParams) SetBody(body models.Empty) {
	o.Body = body
}

// WithMinerID adds the minerID to the post languageunderstanding miner uploads params
func (o *PostLanguageunderstandingMinerUploadsParams) WithMinerID(minerID string) *PostLanguageunderstandingMinerUploadsParams {
	o.SetMinerID(minerID)
	return o
}

// SetMinerID adds the minerId to the post languageunderstanding miner uploads params
func (o *PostLanguageunderstandingMinerUploadsParams) SetMinerID(minerID string) {
	o.MinerID = minerID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLanguageunderstandingMinerUploadsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param minerId
	if err := r.SetPathParam("minerId", o.MinerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
