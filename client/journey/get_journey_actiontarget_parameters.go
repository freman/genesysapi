// Code generated by go-swagger; DO NOT EDIT.

package journey

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

// NewGetJourneyActiontargetParams creates a new GetJourneyActiontargetParams object
// with the default values initialized.
func NewGetJourneyActiontargetParams() *GetJourneyActiontargetParams {
	var ()
	return &GetJourneyActiontargetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetJourneyActiontargetParamsWithTimeout creates a new GetJourneyActiontargetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetJourneyActiontargetParamsWithTimeout(timeout time.Duration) *GetJourneyActiontargetParams {
	var ()
	return &GetJourneyActiontargetParams{

		timeout: timeout,
	}
}

// NewGetJourneyActiontargetParamsWithContext creates a new GetJourneyActiontargetParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetJourneyActiontargetParamsWithContext(ctx context.Context) *GetJourneyActiontargetParams {
	var ()
	return &GetJourneyActiontargetParams{

		Context: ctx,
	}
}

// NewGetJourneyActiontargetParamsWithHTTPClient creates a new GetJourneyActiontargetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetJourneyActiontargetParamsWithHTTPClient(client *http.Client) *GetJourneyActiontargetParams {
	var ()
	return &GetJourneyActiontargetParams{
		HTTPClient: client,
	}
}

/*GetJourneyActiontargetParams contains all the parameters to send to the API endpoint
for the get journey actiontarget operation typically these are written to a http.Request
*/
type GetJourneyActiontargetParams struct {

	/*ActionTargetID
	  ID of the action target.

	*/
	ActionTargetID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get journey actiontarget params
func (o *GetJourneyActiontargetParams) WithTimeout(timeout time.Duration) *GetJourneyActiontargetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get journey actiontarget params
func (o *GetJourneyActiontargetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get journey actiontarget params
func (o *GetJourneyActiontargetParams) WithContext(ctx context.Context) *GetJourneyActiontargetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get journey actiontarget params
func (o *GetJourneyActiontargetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get journey actiontarget params
func (o *GetJourneyActiontargetParams) WithHTTPClient(client *http.Client) *GetJourneyActiontargetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get journey actiontarget params
func (o *GetJourneyActiontargetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionTargetID adds the actionTargetID to the get journey actiontarget params
func (o *GetJourneyActiontargetParams) WithActionTargetID(actionTargetID string) *GetJourneyActiontargetParams {
	o.SetActionTargetID(actionTargetID)
	return o
}

// SetActionTargetID adds the actionTargetId to the get journey actiontarget params
func (o *GetJourneyActiontargetParams) SetActionTargetID(actionTargetID string) {
	o.ActionTargetID = actionTargetID
}

// WriteToRequest writes these params to a swagger request
func (o *GetJourneyActiontargetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param actionTargetId
	if err := r.SetPathParam("actionTargetId", o.ActionTargetID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}