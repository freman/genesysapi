// Code generated by go-swagger; DO NOT EDIT.

package language_understanding

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

// NewPostLanguageunderstandingMinerExecuteParams creates a new PostLanguageunderstandingMinerExecuteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostLanguageunderstandingMinerExecuteParams() *PostLanguageunderstandingMinerExecuteParams {
	return &PostLanguageunderstandingMinerExecuteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostLanguageunderstandingMinerExecuteParamsWithTimeout creates a new PostLanguageunderstandingMinerExecuteParams object
// with the ability to set a timeout on a request.
func NewPostLanguageunderstandingMinerExecuteParamsWithTimeout(timeout time.Duration) *PostLanguageunderstandingMinerExecuteParams {
	return &PostLanguageunderstandingMinerExecuteParams{
		timeout: timeout,
	}
}

// NewPostLanguageunderstandingMinerExecuteParamsWithContext creates a new PostLanguageunderstandingMinerExecuteParams object
// with the ability to set a context for a request.
func NewPostLanguageunderstandingMinerExecuteParamsWithContext(ctx context.Context) *PostLanguageunderstandingMinerExecuteParams {
	return &PostLanguageunderstandingMinerExecuteParams{
		Context: ctx,
	}
}

// NewPostLanguageunderstandingMinerExecuteParamsWithHTTPClient creates a new PostLanguageunderstandingMinerExecuteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostLanguageunderstandingMinerExecuteParamsWithHTTPClient(client *http.Client) *PostLanguageunderstandingMinerExecuteParams {
	return &PostLanguageunderstandingMinerExecuteParams{
		HTTPClient: client,
	}
}

/*
PostLanguageunderstandingMinerExecuteParams contains all the parameters to send to the API endpoint

	for the post languageunderstanding miner execute operation.

	Typically these are written to a http.Request.
*/
type PostLanguageunderstandingMinerExecuteParams struct {

	// Body.
	Body *models.MinerExecuteRequest

	/* MinerID.

	   Miner ID
	*/
	MinerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post languageunderstanding miner execute params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLanguageunderstandingMinerExecuteParams) WithDefaults() *PostLanguageunderstandingMinerExecuteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post languageunderstanding miner execute params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLanguageunderstandingMinerExecuteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post languageunderstanding miner execute params
func (o *PostLanguageunderstandingMinerExecuteParams) WithTimeout(timeout time.Duration) *PostLanguageunderstandingMinerExecuteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post languageunderstanding miner execute params
func (o *PostLanguageunderstandingMinerExecuteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post languageunderstanding miner execute params
func (o *PostLanguageunderstandingMinerExecuteParams) WithContext(ctx context.Context) *PostLanguageunderstandingMinerExecuteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post languageunderstanding miner execute params
func (o *PostLanguageunderstandingMinerExecuteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post languageunderstanding miner execute params
func (o *PostLanguageunderstandingMinerExecuteParams) WithHTTPClient(client *http.Client) *PostLanguageunderstandingMinerExecuteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post languageunderstanding miner execute params
func (o *PostLanguageunderstandingMinerExecuteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post languageunderstanding miner execute params
func (o *PostLanguageunderstandingMinerExecuteParams) WithBody(body *models.MinerExecuteRequest) *PostLanguageunderstandingMinerExecuteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post languageunderstanding miner execute params
func (o *PostLanguageunderstandingMinerExecuteParams) SetBody(body *models.MinerExecuteRequest) {
	o.Body = body
}

// WithMinerID adds the minerID to the post languageunderstanding miner execute params
func (o *PostLanguageunderstandingMinerExecuteParams) WithMinerID(minerID string) *PostLanguageunderstandingMinerExecuteParams {
	o.SetMinerID(minerID)
	return o
}

// SetMinerID adds the minerId to the post languageunderstanding miner execute params
func (o *PostLanguageunderstandingMinerExecuteParams) SetMinerID(minerID string) {
	o.MinerID = minerID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLanguageunderstandingMinerExecuteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
