// Code generated by go-swagger; DO NOT EDIT.

package learning

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

// NewPostLearningAssessmentsScoringParams creates a new PostLearningAssessmentsScoringParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostLearningAssessmentsScoringParams() *PostLearningAssessmentsScoringParams {
	return &PostLearningAssessmentsScoringParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostLearningAssessmentsScoringParamsWithTimeout creates a new PostLearningAssessmentsScoringParams object
// with the ability to set a timeout on a request.
func NewPostLearningAssessmentsScoringParamsWithTimeout(timeout time.Duration) *PostLearningAssessmentsScoringParams {
	return &PostLearningAssessmentsScoringParams{
		timeout: timeout,
	}
}

// NewPostLearningAssessmentsScoringParamsWithContext creates a new PostLearningAssessmentsScoringParams object
// with the ability to set a context for a request.
func NewPostLearningAssessmentsScoringParamsWithContext(ctx context.Context) *PostLearningAssessmentsScoringParams {
	return &PostLearningAssessmentsScoringParams{
		Context: ctx,
	}
}

// NewPostLearningAssessmentsScoringParamsWithHTTPClient creates a new PostLearningAssessmentsScoringParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostLearningAssessmentsScoringParamsWithHTTPClient(client *http.Client) *PostLearningAssessmentsScoringParams {
	return &PostLearningAssessmentsScoringParams{
		HTTPClient: client,
	}
}

/*
PostLearningAssessmentsScoringParams contains all the parameters to send to the API endpoint

	for the post learning assessments scoring operation.

	Typically these are written to a http.Request.
*/
type PostLearningAssessmentsScoringParams struct {

	/* Body.

	   Assessment form and answers to score
	*/
	Body *models.LearningAssessmentScoringRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post learning assessments scoring params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLearningAssessmentsScoringParams) WithDefaults() *PostLearningAssessmentsScoringParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post learning assessments scoring params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostLearningAssessmentsScoringParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post learning assessments scoring params
func (o *PostLearningAssessmentsScoringParams) WithTimeout(timeout time.Duration) *PostLearningAssessmentsScoringParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post learning assessments scoring params
func (o *PostLearningAssessmentsScoringParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post learning assessments scoring params
func (o *PostLearningAssessmentsScoringParams) WithContext(ctx context.Context) *PostLearningAssessmentsScoringParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post learning assessments scoring params
func (o *PostLearningAssessmentsScoringParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post learning assessments scoring params
func (o *PostLearningAssessmentsScoringParams) WithHTTPClient(client *http.Client) *PostLearningAssessmentsScoringParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post learning assessments scoring params
func (o *PostLearningAssessmentsScoringParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post learning assessments scoring params
func (o *PostLearningAssessmentsScoringParams) WithBody(body *models.LearningAssessmentScoringRequest) *PostLearningAssessmentsScoringParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post learning assessments scoring params
func (o *PostLearningAssessmentsScoringParams) SetBody(body *models.LearningAssessmentScoringRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostLearningAssessmentsScoringParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
