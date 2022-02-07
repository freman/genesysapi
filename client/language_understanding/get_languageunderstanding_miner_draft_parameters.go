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
)

// NewGetLanguageunderstandingMinerDraftParams creates a new GetLanguageunderstandingMinerDraftParams object
// with the default values initialized.
func NewGetLanguageunderstandingMinerDraftParams() *GetLanguageunderstandingMinerDraftParams {
	var ()
	return &GetLanguageunderstandingMinerDraftParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLanguageunderstandingMinerDraftParamsWithTimeout creates a new GetLanguageunderstandingMinerDraftParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLanguageunderstandingMinerDraftParamsWithTimeout(timeout time.Duration) *GetLanguageunderstandingMinerDraftParams {
	var ()
	return &GetLanguageunderstandingMinerDraftParams{

		timeout: timeout,
	}
}

// NewGetLanguageunderstandingMinerDraftParamsWithContext creates a new GetLanguageunderstandingMinerDraftParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLanguageunderstandingMinerDraftParamsWithContext(ctx context.Context) *GetLanguageunderstandingMinerDraftParams {
	var ()
	return &GetLanguageunderstandingMinerDraftParams{

		Context: ctx,
	}
}

// NewGetLanguageunderstandingMinerDraftParamsWithHTTPClient creates a new GetLanguageunderstandingMinerDraftParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLanguageunderstandingMinerDraftParamsWithHTTPClient(client *http.Client) *GetLanguageunderstandingMinerDraftParams {
	var ()
	return &GetLanguageunderstandingMinerDraftParams{
		HTTPClient: client,
	}
}

/*GetLanguageunderstandingMinerDraftParams contains all the parameters to send to the API endpoint
for the get languageunderstanding miner draft operation typically these are written to a http.Request
*/
type GetLanguageunderstandingMinerDraftParams struct {

	/*DraftID
	  Draft ID

	*/
	DraftID string
	/*MinerID
	  Miner ID

	*/
	MinerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get languageunderstanding miner draft params
func (o *GetLanguageunderstandingMinerDraftParams) WithTimeout(timeout time.Duration) *GetLanguageunderstandingMinerDraftParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get languageunderstanding miner draft params
func (o *GetLanguageunderstandingMinerDraftParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get languageunderstanding miner draft params
func (o *GetLanguageunderstandingMinerDraftParams) WithContext(ctx context.Context) *GetLanguageunderstandingMinerDraftParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get languageunderstanding miner draft params
func (o *GetLanguageunderstandingMinerDraftParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get languageunderstanding miner draft params
func (o *GetLanguageunderstandingMinerDraftParams) WithHTTPClient(client *http.Client) *GetLanguageunderstandingMinerDraftParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get languageunderstanding miner draft params
func (o *GetLanguageunderstandingMinerDraftParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDraftID adds the draftID to the get languageunderstanding miner draft params
func (o *GetLanguageunderstandingMinerDraftParams) WithDraftID(draftID string) *GetLanguageunderstandingMinerDraftParams {
	o.SetDraftID(draftID)
	return o
}

// SetDraftID adds the draftId to the get languageunderstanding miner draft params
func (o *GetLanguageunderstandingMinerDraftParams) SetDraftID(draftID string) {
	o.DraftID = draftID
}

// WithMinerID adds the minerID to the get languageunderstanding miner draft params
func (o *GetLanguageunderstandingMinerDraftParams) WithMinerID(minerID string) *GetLanguageunderstandingMinerDraftParams {
	o.SetMinerID(minerID)
	return o
}

// SetMinerID adds the minerId to the get languageunderstanding miner draft params
func (o *GetLanguageunderstandingMinerDraftParams) SetMinerID(minerID string) {
	o.MinerID = minerID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLanguageunderstandingMinerDraftParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param draftId
	if err := r.SetPathParam("draftId", o.DraftID); err != nil {
		return err
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