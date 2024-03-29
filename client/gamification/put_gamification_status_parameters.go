// Code generated by go-swagger; DO NOT EDIT.

package gamification

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

// NewPutGamificationStatusParams creates a new PutGamificationStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutGamificationStatusParams() *PutGamificationStatusParams {
	return &PutGamificationStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutGamificationStatusParamsWithTimeout creates a new PutGamificationStatusParams object
// with the ability to set a timeout on a request.
func NewPutGamificationStatusParamsWithTimeout(timeout time.Duration) *PutGamificationStatusParams {
	return &PutGamificationStatusParams{
		timeout: timeout,
	}
}

// NewPutGamificationStatusParamsWithContext creates a new PutGamificationStatusParams object
// with the ability to set a context for a request.
func NewPutGamificationStatusParamsWithContext(ctx context.Context) *PutGamificationStatusParams {
	return &PutGamificationStatusParams{
		Context: ctx,
	}
}

// NewPutGamificationStatusParamsWithHTTPClient creates a new PutGamificationStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutGamificationStatusParamsWithHTTPClient(client *http.Client) *PutGamificationStatusParams {
	return &PutGamificationStatusParams{
		HTTPClient: client,
	}
}

/*
PutGamificationStatusParams contains all the parameters to send to the API endpoint

	for the put gamification status operation.

	Typically these are written to a http.Request.
*/
type PutGamificationStatusParams struct {

	/* Status.

	   Gamification status
	*/
	Status *models.GamificationStatus

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put gamification status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutGamificationStatusParams) WithDefaults() *PutGamificationStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put gamification status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutGamificationStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put gamification status params
func (o *PutGamificationStatusParams) WithTimeout(timeout time.Duration) *PutGamificationStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put gamification status params
func (o *PutGamificationStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put gamification status params
func (o *PutGamificationStatusParams) WithContext(ctx context.Context) *PutGamificationStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put gamification status params
func (o *PutGamificationStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put gamification status params
func (o *PutGamificationStatusParams) WithHTTPClient(client *http.Client) *PutGamificationStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put gamification status params
func (o *PutGamificationStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStatus adds the status to the put gamification status params
func (o *PutGamificationStatusParams) WithStatus(status *models.GamificationStatus) *PutGamificationStatusParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the put gamification status params
func (o *PutGamificationStatusParams) SetStatus(status *models.GamificationStatus) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *PutGamificationStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Status != nil {
		if err := r.SetBodyParam(o.Status); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
