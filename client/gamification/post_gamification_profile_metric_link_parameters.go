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

// NewPostGamificationProfileMetricLinkParams creates a new PostGamificationProfileMetricLinkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostGamificationProfileMetricLinkParams() *PostGamificationProfileMetricLinkParams {
	return &PostGamificationProfileMetricLinkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostGamificationProfileMetricLinkParamsWithTimeout creates a new PostGamificationProfileMetricLinkParams object
// with the ability to set a timeout on a request.
func NewPostGamificationProfileMetricLinkParamsWithTimeout(timeout time.Duration) *PostGamificationProfileMetricLinkParams {
	return &PostGamificationProfileMetricLinkParams{
		timeout: timeout,
	}
}

// NewPostGamificationProfileMetricLinkParamsWithContext creates a new PostGamificationProfileMetricLinkParams object
// with the ability to set a context for a request.
func NewPostGamificationProfileMetricLinkParamsWithContext(ctx context.Context) *PostGamificationProfileMetricLinkParams {
	return &PostGamificationProfileMetricLinkParams{
		Context: ctx,
	}
}

// NewPostGamificationProfileMetricLinkParamsWithHTTPClient creates a new PostGamificationProfileMetricLinkParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostGamificationProfileMetricLinkParamsWithHTTPClient(client *http.Client) *PostGamificationProfileMetricLinkParams {
	return &PostGamificationProfileMetricLinkParams{
		HTTPClient: client,
	}
}

/*
PostGamificationProfileMetricLinkParams contains all the parameters to send to the API endpoint

	for the post gamification profile metric link operation.

	Typically these are written to a http.Request.
*/
type PostGamificationProfileMetricLinkParams struct {

	/* Body.

	   linkedMetric
	*/
	Body *models.TargetPerformanceProfile

	/* SourceMetricID.

	   Source Metric Id
	*/
	SourceMetricID string

	/* SourceProfileID.

	   Source Performance Profile Id
	*/
	SourceProfileID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post gamification profile metric link params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostGamificationProfileMetricLinkParams) WithDefaults() *PostGamificationProfileMetricLinkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post gamification profile metric link params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostGamificationProfileMetricLinkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post gamification profile metric link params
func (o *PostGamificationProfileMetricLinkParams) WithTimeout(timeout time.Duration) *PostGamificationProfileMetricLinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post gamification profile metric link params
func (o *PostGamificationProfileMetricLinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post gamification profile metric link params
func (o *PostGamificationProfileMetricLinkParams) WithContext(ctx context.Context) *PostGamificationProfileMetricLinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post gamification profile metric link params
func (o *PostGamificationProfileMetricLinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post gamification profile metric link params
func (o *PostGamificationProfileMetricLinkParams) WithHTTPClient(client *http.Client) *PostGamificationProfileMetricLinkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post gamification profile metric link params
func (o *PostGamificationProfileMetricLinkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post gamification profile metric link params
func (o *PostGamificationProfileMetricLinkParams) WithBody(body *models.TargetPerformanceProfile) *PostGamificationProfileMetricLinkParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post gamification profile metric link params
func (o *PostGamificationProfileMetricLinkParams) SetBody(body *models.TargetPerformanceProfile) {
	o.Body = body
}

// WithSourceMetricID adds the sourceMetricID to the post gamification profile metric link params
func (o *PostGamificationProfileMetricLinkParams) WithSourceMetricID(sourceMetricID string) *PostGamificationProfileMetricLinkParams {
	o.SetSourceMetricID(sourceMetricID)
	return o
}

// SetSourceMetricID adds the sourceMetricId to the post gamification profile metric link params
func (o *PostGamificationProfileMetricLinkParams) SetSourceMetricID(sourceMetricID string) {
	o.SourceMetricID = sourceMetricID
}

// WithSourceProfileID adds the sourceProfileID to the post gamification profile metric link params
func (o *PostGamificationProfileMetricLinkParams) WithSourceProfileID(sourceProfileID string) *PostGamificationProfileMetricLinkParams {
	o.SetSourceProfileID(sourceProfileID)
	return o
}

// SetSourceProfileID adds the sourceProfileId to the post gamification profile metric link params
func (o *PostGamificationProfileMetricLinkParams) SetSourceProfileID(sourceProfileID string) {
	o.SourceProfileID = sourceProfileID
}

// WriteToRequest writes these params to a swagger request
func (o *PostGamificationProfileMetricLinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param sourceMetricId
	if err := r.SetPathParam("sourceMetricId", o.SourceMetricID); err != nil {
		return err
	}

	// path param sourceProfileId
	if err := r.SetPathParam("sourceProfileId", o.SourceProfileID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
