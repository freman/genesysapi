// Code generated by go-swagger; DO NOT EDIT.

package alerting

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

// NewDeleteAlertingInteractionstatsAlertParams creates a new DeleteAlertingInteractionstatsAlertParams object
// with the default values initialized.
func NewDeleteAlertingInteractionstatsAlertParams() *DeleteAlertingInteractionstatsAlertParams {
	var ()
	return &DeleteAlertingInteractionstatsAlertParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAlertingInteractionstatsAlertParamsWithTimeout creates a new DeleteAlertingInteractionstatsAlertParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAlertingInteractionstatsAlertParamsWithTimeout(timeout time.Duration) *DeleteAlertingInteractionstatsAlertParams {
	var ()
	return &DeleteAlertingInteractionstatsAlertParams{

		timeout: timeout,
	}
}

// NewDeleteAlertingInteractionstatsAlertParamsWithContext creates a new DeleteAlertingInteractionstatsAlertParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAlertingInteractionstatsAlertParamsWithContext(ctx context.Context) *DeleteAlertingInteractionstatsAlertParams {
	var ()
	return &DeleteAlertingInteractionstatsAlertParams{

		Context: ctx,
	}
}

// NewDeleteAlertingInteractionstatsAlertParamsWithHTTPClient creates a new DeleteAlertingInteractionstatsAlertParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAlertingInteractionstatsAlertParamsWithHTTPClient(client *http.Client) *DeleteAlertingInteractionstatsAlertParams {
	var ()
	return &DeleteAlertingInteractionstatsAlertParams{
		HTTPClient: client,
	}
}

/*DeleteAlertingInteractionstatsAlertParams contains all the parameters to send to the API endpoint
for the delete alerting interactionstats alert operation typically these are written to a http.Request
*/
type DeleteAlertingInteractionstatsAlertParams struct {

	/*AlertID
	  Alert ID

	*/
	AlertID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete alerting interactionstats alert params
func (o *DeleteAlertingInteractionstatsAlertParams) WithTimeout(timeout time.Duration) *DeleteAlertingInteractionstatsAlertParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete alerting interactionstats alert params
func (o *DeleteAlertingInteractionstatsAlertParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete alerting interactionstats alert params
func (o *DeleteAlertingInteractionstatsAlertParams) WithContext(ctx context.Context) *DeleteAlertingInteractionstatsAlertParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete alerting interactionstats alert params
func (o *DeleteAlertingInteractionstatsAlertParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete alerting interactionstats alert params
func (o *DeleteAlertingInteractionstatsAlertParams) WithHTTPClient(client *http.Client) *DeleteAlertingInteractionstatsAlertParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete alerting interactionstats alert params
func (o *DeleteAlertingInteractionstatsAlertParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlertID adds the alertID to the delete alerting interactionstats alert params
func (o *DeleteAlertingInteractionstatsAlertParams) WithAlertID(alertID string) *DeleteAlertingInteractionstatsAlertParams {
	o.SetAlertID(alertID)
	return o
}

// SetAlertID adds the alertId to the delete alerting interactionstats alert params
func (o *DeleteAlertingInteractionstatsAlertParams) SetAlertID(alertID string) {
	o.AlertID = alertID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAlertingInteractionstatsAlertParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param alertId
	if err := r.SetPathParam("alertId", o.AlertID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
