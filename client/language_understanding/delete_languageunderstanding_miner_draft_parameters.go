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

// NewDeleteLanguageunderstandingMinerDraftParams creates a new DeleteLanguageunderstandingMinerDraftParams object
// with the default values initialized.
func NewDeleteLanguageunderstandingMinerDraftParams() *DeleteLanguageunderstandingMinerDraftParams {
	var ()
	return &DeleteLanguageunderstandingMinerDraftParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLanguageunderstandingMinerDraftParamsWithTimeout creates a new DeleteLanguageunderstandingMinerDraftParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLanguageunderstandingMinerDraftParamsWithTimeout(timeout time.Duration) *DeleteLanguageunderstandingMinerDraftParams {
	var ()
	return &DeleteLanguageunderstandingMinerDraftParams{

		timeout: timeout,
	}
}

// NewDeleteLanguageunderstandingMinerDraftParamsWithContext creates a new DeleteLanguageunderstandingMinerDraftParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLanguageunderstandingMinerDraftParamsWithContext(ctx context.Context) *DeleteLanguageunderstandingMinerDraftParams {
	var ()
	return &DeleteLanguageunderstandingMinerDraftParams{

		Context: ctx,
	}
}

// NewDeleteLanguageunderstandingMinerDraftParamsWithHTTPClient creates a new DeleteLanguageunderstandingMinerDraftParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLanguageunderstandingMinerDraftParamsWithHTTPClient(client *http.Client) *DeleteLanguageunderstandingMinerDraftParams {
	var ()
	return &DeleteLanguageunderstandingMinerDraftParams{
		HTTPClient: client,
	}
}

/*DeleteLanguageunderstandingMinerDraftParams contains all the parameters to send to the API endpoint
for the delete languageunderstanding miner draft operation typically these are written to a http.Request
*/
type DeleteLanguageunderstandingMinerDraftParams struct {

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

// WithTimeout adds the timeout to the delete languageunderstanding miner draft params
func (o *DeleteLanguageunderstandingMinerDraftParams) WithTimeout(timeout time.Duration) *DeleteLanguageunderstandingMinerDraftParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete languageunderstanding miner draft params
func (o *DeleteLanguageunderstandingMinerDraftParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete languageunderstanding miner draft params
func (o *DeleteLanguageunderstandingMinerDraftParams) WithContext(ctx context.Context) *DeleteLanguageunderstandingMinerDraftParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete languageunderstanding miner draft params
func (o *DeleteLanguageunderstandingMinerDraftParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete languageunderstanding miner draft params
func (o *DeleteLanguageunderstandingMinerDraftParams) WithHTTPClient(client *http.Client) *DeleteLanguageunderstandingMinerDraftParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete languageunderstanding miner draft params
func (o *DeleteLanguageunderstandingMinerDraftParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDraftID adds the draftID to the delete languageunderstanding miner draft params
func (o *DeleteLanguageunderstandingMinerDraftParams) WithDraftID(draftID string) *DeleteLanguageunderstandingMinerDraftParams {
	o.SetDraftID(draftID)
	return o
}

// SetDraftID adds the draftId to the delete languageunderstanding miner draft params
func (o *DeleteLanguageunderstandingMinerDraftParams) SetDraftID(draftID string) {
	o.DraftID = draftID
}

// WithMinerID adds the minerID to the delete languageunderstanding miner draft params
func (o *DeleteLanguageunderstandingMinerDraftParams) WithMinerID(minerID string) *DeleteLanguageunderstandingMinerDraftParams {
	o.SetMinerID(minerID)
	return o
}

// SetMinerID adds the minerId to the delete languageunderstanding miner draft params
func (o *DeleteLanguageunderstandingMinerDraftParams) SetMinerID(minerID string) {
	o.MinerID = minerID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLanguageunderstandingMinerDraftParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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