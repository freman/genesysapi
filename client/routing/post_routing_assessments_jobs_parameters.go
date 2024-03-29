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

// NewPostRoutingAssessmentsJobsParams creates a new PostRoutingAssessmentsJobsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostRoutingAssessmentsJobsParams() *PostRoutingAssessmentsJobsParams {
	return &PostRoutingAssessmentsJobsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostRoutingAssessmentsJobsParamsWithTimeout creates a new PostRoutingAssessmentsJobsParams object
// with the ability to set a timeout on a request.
func NewPostRoutingAssessmentsJobsParamsWithTimeout(timeout time.Duration) *PostRoutingAssessmentsJobsParams {
	return &PostRoutingAssessmentsJobsParams{
		timeout: timeout,
	}
}

// NewPostRoutingAssessmentsJobsParamsWithContext creates a new PostRoutingAssessmentsJobsParams object
// with the ability to set a context for a request.
func NewPostRoutingAssessmentsJobsParamsWithContext(ctx context.Context) *PostRoutingAssessmentsJobsParams {
	return &PostRoutingAssessmentsJobsParams{
		Context: ctx,
	}
}

// NewPostRoutingAssessmentsJobsParamsWithHTTPClient creates a new PostRoutingAssessmentsJobsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostRoutingAssessmentsJobsParamsWithHTTPClient(client *http.Client) *PostRoutingAssessmentsJobsParams {
	return &PostRoutingAssessmentsJobsParams{
		HTTPClient: client,
	}
}

/*
PostRoutingAssessmentsJobsParams contains all the parameters to send to the API endpoint

	for the post routing assessments jobs operation.

	Typically these are written to a http.Request.
*/
type PostRoutingAssessmentsJobsParams struct {

	// Body.
	Body *models.CreateBenefitAssessmentJobRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post routing assessments jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRoutingAssessmentsJobsParams) WithDefaults() *PostRoutingAssessmentsJobsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post routing assessments jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRoutingAssessmentsJobsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post routing assessments jobs params
func (o *PostRoutingAssessmentsJobsParams) WithTimeout(timeout time.Duration) *PostRoutingAssessmentsJobsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post routing assessments jobs params
func (o *PostRoutingAssessmentsJobsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post routing assessments jobs params
func (o *PostRoutingAssessmentsJobsParams) WithContext(ctx context.Context) *PostRoutingAssessmentsJobsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post routing assessments jobs params
func (o *PostRoutingAssessmentsJobsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post routing assessments jobs params
func (o *PostRoutingAssessmentsJobsParams) WithHTTPClient(client *http.Client) *PostRoutingAssessmentsJobsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post routing assessments jobs params
func (o *PostRoutingAssessmentsJobsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post routing assessments jobs params
func (o *PostRoutingAssessmentsJobsParams) WithBody(body *models.CreateBenefitAssessmentJobRequest) *PostRoutingAssessmentsJobsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post routing assessments jobs params
func (o *PostRoutingAssessmentsJobsParams) SetBody(body *models.CreateBenefitAssessmentJobRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRoutingAssessmentsJobsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
