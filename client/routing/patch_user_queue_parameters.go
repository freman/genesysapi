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

// NewPatchUserQueueParams creates a new PatchUserQueueParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchUserQueueParams() *PatchUserQueueParams {
	return &PatchUserQueueParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchUserQueueParamsWithTimeout creates a new PatchUserQueueParams object
// with the ability to set a timeout on a request.
func NewPatchUserQueueParamsWithTimeout(timeout time.Duration) *PatchUserQueueParams {
	return &PatchUserQueueParams{
		timeout: timeout,
	}
}

// NewPatchUserQueueParamsWithContext creates a new PatchUserQueueParams object
// with the ability to set a context for a request.
func NewPatchUserQueueParamsWithContext(ctx context.Context) *PatchUserQueueParams {
	return &PatchUserQueueParams{
		Context: ctx,
	}
}

// NewPatchUserQueueParamsWithHTTPClient creates a new PatchUserQueueParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchUserQueueParamsWithHTTPClient(client *http.Client) *PatchUserQueueParams {
	return &PatchUserQueueParams{
		HTTPClient: client,
	}
}

/*
PatchUserQueueParams contains all the parameters to send to the API endpoint

	for the patch user queue operation.

	Typically these are written to a http.Request.
*/
type PatchUserQueueParams struct {

	/* Body.

	   Queue Member
	*/
	Body *models.UserQueue

	/* QueueID.

	   Queue ID
	*/
	QueueID string

	/* UserID.

	   User ID
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch user queue params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchUserQueueParams) WithDefaults() *PatchUserQueueParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch user queue params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchUserQueueParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch user queue params
func (o *PatchUserQueueParams) WithTimeout(timeout time.Duration) *PatchUserQueueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch user queue params
func (o *PatchUserQueueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch user queue params
func (o *PatchUserQueueParams) WithContext(ctx context.Context) *PatchUserQueueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch user queue params
func (o *PatchUserQueueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch user queue params
func (o *PatchUserQueueParams) WithHTTPClient(client *http.Client) *PatchUserQueueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch user queue params
func (o *PatchUserQueueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch user queue params
func (o *PatchUserQueueParams) WithBody(body *models.UserQueue) *PatchUserQueueParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch user queue params
func (o *PatchUserQueueParams) SetBody(body *models.UserQueue) {
	o.Body = body
}

// WithQueueID adds the queueID to the patch user queue params
func (o *PatchUserQueueParams) WithQueueID(queueID string) *PatchUserQueueParams {
	o.SetQueueID(queueID)
	return o
}

// SetQueueID adds the queueId to the patch user queue params
func (o *PatchUserQueueParams) SetQueueID(queueID string) {
	o.QueueID = queueID
}

// WithUserID adds the userID to the patch user queue params
func (o *PatchUserQueueParams) WithUserID(userID string) *PatchUserQueueParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the patch user queue params
func (o *PatchUserQueueParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchUserQueueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param queueId
	if err := r.SetPathParam("queueId", o.QueueID); err != nil {
		return err
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
