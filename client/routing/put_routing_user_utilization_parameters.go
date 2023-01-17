// Code generated by go-swagger; DO NOT EDIT.

package routing

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

// NewPutRoutingUserUtilizationParams creates a new PutRoutingUserUtilizationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutRoutingUserUtilizationParams() *PutRoutingUserUtilizationParams {
	return &PutRoutingUserUtilizationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutRoutingUserUtilizationParamsWithTimeout creates a new PutRoutingUserUtilizationParams object
// with the ability to set a timeout on a request.
func NewPutRoutingUserUtilizationParamsWithTimeout(timeout time.Duration) *PutRoutingUserUtilizationParams {
	return &PutRoutingUserUtilizationParams{
		timeout: timeout,
	}
}

// NewPutRoutingUserUtilizationParamsWithContext creates a new PutRoutingUserUtilizationParams object
// with the ability to set a context for a request.
func NewPutRoutingUserUtilizationParamsWithContext(ctx context.Context) *PutRoutingUserUtilizationParams {
	return &PutRoutingUserUtilizationParams{
		Context: ctx,
	}
}

// NewPutRoutingUserUtilizationParamsWithHTTPClient creates a new PutRoutingUserUtilizationParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutRoutingUserUtilizationParamsWithHTTPClient(client *http.Client) *PutRoutingUserUtilizationParams {
	return &PutRoutingUserUtilizationParams{
		HTTPClient: client,
	}
}

/*
PutRoutingUserUtilizationParams contains all the parameters to send to the API endpoint

	for the put routing user utilization operation.

	Typically these are written to a http.Request.
*/
type PutRoutingUserUtilizationParams struct {

	/* Body.

	   utilization
	*/
	Body *models.Utilization

	/* UserID.

	   User ID
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put routing user utilization params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRoutingUserUtilizationParams) WithDefaults() *PutRoutingUserUtilizationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put routing user utilization params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutRoutingUserUtilizationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put routing user utilization params
func (o *PutRoutingUserUtilizationParams) WithTimeout(timeout time.Duration) *PutRoutingUserUtilizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put routing user utilization params
func (o *PutRoutingUserUtilizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put routing user utilization params
func (o *PutRoutingUserUtilizationParams) WithContext(ctx context.Context) *PutRoutingUserUtilizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put routing user utilization params
func (o *PutRoutingUserUtilizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put routing user utilization params
func (o *PutRoutingUserUtilizationParams) WithHTTPClient(client *http.Client) *PutRoutingUserUtilizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put routing user utilization params
func (o *PutRoutingUserUtilizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put routing user utilization params
func (o *PutRoutingUserUtilizationParams) WithBody(body *models.Utilization) *PutRoutingUserUtilizationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put routing user utilization params
func (o *PutRoutingUserUtilizationParams) SetBody(body *models.Utilization) {
	o.Body = body
}

// WithUserID adds the userID to the put routing user utilization params
func (o *PutRoutingUserUtilizationParams) WithUserID(userID string) *PutRoutingUserUtilizationParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the put routing user utilization params
func (o *PutRoutingUserUtilizationParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PutRoutingUserUtilizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
