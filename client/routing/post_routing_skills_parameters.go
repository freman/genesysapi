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

// NewPostRoutingSkillsParams creates a new PostRoutingSkillsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostRoutingSkillsParams() *PostRoutingSkillsParams {
	return &PostRoutingSkillsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostRoutingSkillsParamsWithTimeout creates a new PostRoutingSkillsParams object
// with the ability to set a timeout on a request.
func NewPostRoutingSkillsParamsWithTimeout(timeout time.Duration) *PostRoutingSkillsParams {
	return &PostRoutingSkillsParams{
		timeout: timeout,
	}
}

// NewPostRoutingSkillsParamsWithContext creates a new PostRoutingSkillsParams object
// with the ability to set a context for a request.
func NewPostRoutingSkillsParamsWithContext(ctx context.Context) *PostRoutingSkillsParams {
	return &PostRoutingSkillsParams{
		Context: ctx,
	}
}

// NewPostRoutingSkillsParamsWithHTTPClient creates a new PostRoutingSkillsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostRoutingSkillsParamsWithHTTPClient(client *http.Client) *PostRoutingSkillsParams {
	return &PostRoutingSkillsParams{
		HTTPClient: client,
	}
}

/*
PostRoutingSkillsParams contains all the parameters to send to the API endpoint

	for the post routing skills operation.

	Typically these are written to a http.Request.
*/
type PostRoutingSkillsParams struct {

	/* Body.

	   Skill
	*/
	Body *models.RoutingSkill

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post routing skills params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRoutingSkillsParams) WithDefaults() *PostRoutingSkillsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post routing skills params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRoutingSkillsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post routing skills params
func (o *PostRoutingSkillsParams) WithTimeout(timeout time.Duration) *PostRoutingSkillsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post routing skills params
func (o *PostRoutingSkillsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post routing skills params
func (o *PostRoutingSkillsParams) WithContext(ctx context.Context) *PostRoutingSkillsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post routing skills params
func (o *PostRoutingSkillsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post routing skills params
func (o *PostRoutingSkillsParams) WithHTTPClient(client *http.Client) *PostRoutingSkillsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post routing skills params
func (o *PostRoutingSkillsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post routing skills params
func (o *PostRoutingSkillsParams) WithBody(body *models.RoutingSkill) *PostRoutingSkillsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post routing skills params
func (o *PostRoutingSkillsParams) SetBody(body *models.RoutingSkill) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRoutingSkillsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
