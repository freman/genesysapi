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

	"github.com/freman/genesysapi/models"
)

// NewPostJourneyActionmapsEstimatesJobsParams creates a new PostJourneyActionmapsEstimatesJobsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostJourneyActionmapsEstimatesJobsParams() *PostJourneyActionmapsEstimatesJobsParams {
	return &PostJourneyActionmapsEstimatesJobsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostJourneyActionmapsEstimatesJobsParamsWithTimeout creates a new PostJourneyActionmapsEstimatesJobsParams object
// with the ability to set a timeout on a request.
func NewPostJourneyActionmapsEstimatesJobsParamsWithTimeout(timeout time.Duration) *PostJourneyActionmapsEstimatesJobsParams {
	return &PostJourneyActionmapsEstimatesJobsParams{
		timeout: timeout,
	}
}

// NewPostJourneyActionmapsEstimatesJobsParamsWithContext creates a new PostJourneyActionmapsEstimatesJobsParams object
// with the ability to set a context for a request.
func NewPostJourneyActionmapsEstimatesJobsParamsWithContext(ctx context.Context) *PostJourneyActionmapsEstimatesJobsParams {
	return &PostJourneyActionmapsEstimatesJobsParams{
		Context: ctx,
	}
}

// NewPostJourneyActionmapsEstimatesJobsParamsWithHTTPClient creates a new PostJourneyActionmapsEstimatesJobsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostJourneyActionmapsEstimatesJobsParamsWithHTTPClient(client *http.Client) *PostJourneyActionmapsEstimatesJobsParams {
	return &PostJourneyActionmapsEstimatesJobsParams{
		HTTPClient: client,
	}
}

/*
PostJourneyActionmapsEstimatesJobsParams contains all the parameters to send to the API endpoint

	for the post journey actionmaps estimates jobs operation.

	Typically these are written to a http.Request.
*/
type PostJourneyActionmapsEstimatesJobsParams struct {

	/* Body.

	   audience estimator request
	*/
	Body *models.ActionMapEstimateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post journey actionmaps estimates jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostJourneyActionmapsEstimatesJobsParams) WithDefaults() *PostJourneyActionmapsEstimatesJobsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post journey actionmaps estimates jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostJourneyActionmapsEstimatesJobsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post journey actionmaps estimates jobs params
func (o *PostJourneyActionmapsEstimatesJobsParams) WithTimeout(timeout time.Duration) *PostJourneyActionmapsEstimatesJobsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post journey actionmaps estimates jobs params
func (o *PostJourneyActionmapsEstimatesJobsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post journey actionmaps estimates jobs params
func (o *PostJourneyActionmapsEstimatesJobsParams) WithContext(ctx context.Context) *PostJourneyActionmapsEstimatesJobsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post journey actionmaps estimates jobs params
func (o *PostJourneyActionmapsEstimatesJobsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post journey actionmaps estimates jobs params
func (o *PostJourneyActionmapsEstimatesJobsParams) WithHTTPClient(client *http.Client) *PostJourneyActionmapsEstimatesJobsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post journey actionmaps estimates jobs params
func (o *PostJourneyActionmapsEstimatesJobsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post journey actionmaps estimates jobs params
func (o *PostJourneyActionmapsEstimatesJobsParams) WithBody(body *models.ActionMapEstimateRequest) *PostJourneyActionmapsEstimatesJobsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post journey actionmaps estimates jobs params
func (o *PostJourneyActionmapsEstimatesJobsParams) SetBody(body *models.ActionMapEstimateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostJourneyActionmapsEstimatesJobsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
