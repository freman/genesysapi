// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewPatchUserRoutingskillsBulkParams creates a new PatchUserRoutingskillsBulkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchUserRoutingskillsBulkParams() *PatchUserRoutingskillsBulkParams {
	return &PatchUserRoutingskillsBulkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchUserRoutingskillsBulkParamsWithTimeout creates a new PatchUserRoutingskillsBulkParams object
// with the ability to set a timeout on a request.
func NewPatchUserRoutingskillsBulkParamsWithTimeout(timeout time.Duration) *PatchUserRoutingskillsBulkParams {
	return &PatchUserRoutingskillsBulkParams{
		timeout: timeout,
	}
}

// NewPatchUserRoutingskillsBulkParamsWithContext creates a new PatchUserRoutingskillsBulkParams object
// with the ability to set a context for a request.
func NewPatchUserRoutingskillsBulkParamsWithContext(ctx context.Context) *PatchUserRoutingskillsBulkParams {
	return &PatchUserRoutingskillsBulkParams{
		Context: ctx,
	}
}

// NewPatchUserRoutingskillsBulkParamsWithHTTPClient creates a new PatchUserRoutingskillsBulkParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchUserRoutingskillsBulkParamsWithHTTPClient(client *http.Client) *PatchUserRoutingskillsBulkParams {
	return &PatchUserRoutingskillsBulkParams{
		HTTPClient: client,
	}
}

/*
PatchUserRoutingskillsBulkParams contains all the parameters to send to the API endpoint

	for the patch user routingskills bulk operation.

	Typically these are written to a http.Request.
*/
type PatchUserRoutingskillsBulkParams struct {

	/* Body.

	   Skill
	*/
	Body []*models.UserRoutingSkillPost

	/* UserID.

	   User ID
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch user routingskills bulk params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchUserRoutingskillsBulkParams) WithDefaults() *PatchUserRoutingskillsBulkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch user routingskills bulk params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchUserRoutingskillsBulkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch user routingskills bulk params
func (o *PatchUserRoutingskillsBulkParams) WithTimeout(timeout time.Duration) *PatchUserRoutingskillsBulkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch user routingskills bulk params
func (o *PatchUserRoutingskillsBulkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch user routingskills bulk params
func (o *PatchUserRoutingskillsBulkParams) WithContext(ctx context.Context) *PatchUserRoutingskillsBulkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch user routingskills bulk params
func (o *PatchUserRoutingskillsBulkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch user routingskills bulk params
func (o *PatchUserRoutingskillsBulkParams) WithHTTPClient(client *http.Client) *PatchUserRoutingskillsBulkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch user routingskills bulk params
func (o *PatchUserRoutingskillsBulkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch user routingskills bulk params
func (o *PatchUserRoutingskillsBulkParams) WithBody(body []*models.UserRoutingSkillPost) *PatchUserRoutingskillsBulkParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch user routingskills bulk params
func (o *PatchUserRoutingskillsBulkParams) SetBody(body []*models.UserRoutingSkillPost) {
	o.Body = body
}

// WithUserID adds the userID to the patch user routingskills bulk params
func (o *PatchUserRoutingskillsBulkParams) WithUserID(userID string) *PatchUserRoutingskillsBulkParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the patch user routingskills bulk params
func (o *PatchUserRoutingskillsBulkParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchUserRoutingskillsBulkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
