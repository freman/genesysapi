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

// NewPatchLearningAssignmentRescheduleParams creates a new PatchLearningAssignmentRescheduleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchLearningAssignmentRescheduleParams() *PatchLearningAssignmentRescheduleParams {
	return &PatchLearningAssignmentRescheduleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchLearningAssignmentRescheduleParamsWithTimeout creates a new PatchLearningAssignmentRescheduleParams object
// with the ability to set a timeout on a request.
func NewPatchLearningAssignmentRescheduleParamsWithTimeout(timeout time.Duration) *PatchLearningAssignmentRescheduleParams {
	return &PatchLearningAssignmentRescheduleParams{
		timeout: timeout,
	}
}

// NewPatchLearningAssignmentRescheduleParamsWithContext creates a new PatchLearningAssignmentRescheduleParams object
// with the ability to set a context for a request.
func NewPatchLearningAssignmentRescheduleParamsWithContext(ctx context.Context) *PatchLearningAssignmentRescheduleParams {
	return &PatchLearningAssignmentRescheduleParams{
		Context: ctx,
	}
}

// NewPatchLearningAssignmentRescheduleParamsWithHTTPClient creates a new PatchLearningAssignmentRescheduleParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchLearningAssignmentRescheduleParamsWithHTTPClient(client *http.Client) *PatchLearningAssignmentRescheduleParams {
	return &PatchLearningAssignmentRescheduleParams{
		HTTPClient: client,
	}
}

/*
PatchLearningAssignmentRescheduleParams contains all the parameters to send to the API endpoint

	for the patch learning assignment reschedule operation.

	Typically these are written to a http.Request.
*/
type PatchLearningAssignmentRescheduleParams struct {

	/* AssignmentID.

	   The ID of Learning Assignment
	*/
	AssignmentID string

	/* Body.

	   The Learning assignment reschedule model
	*/
	Body *models.LearningAssignmentReschedule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch learning assignment reschedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchLearningAssignmentRescheduleParams) WithDefaults() *PatchLearningAssignmentRescheduleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch learning assignment reschedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchLearningAssignmentRescheduleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch learning assignment reschedule params
func (o *PatchLearningAssignmentRescheduleParams) WithTimeout(timeout time.Duration) *PatchLearningAssignmentRescheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch learning assignment reschedule params
func (o *PatchLearningAssignmentRescheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch learning assignment reschedule params
func (o *PatchLearningAssignmentRescheduleParams) WithContext(ctx context.Context) *PatchLearningAssignmentRescheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch learning assignment reschedule params
func (o *PatchLearningAssignmentRescheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch learning assignment reschedule params
func (o *PatchLearningAssignmentRescheduleParams) WithHTTPClient(client *http.Client) *PatchLearningAssignmentRescheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch learning assignment reschedule params
func (o *PatchLearningAssignmentRescheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssignmentID adds the assignmentID to the patch learning assignment reschedule params
func (o *PatchLearningAssignmentRescheduleParams) WithAssignmentID(assignmentID string) *PatchLearningAssignmentRescheduleParams {
	o.SetAssignmentID(assignmentID)
	return o
}

// SetAssignmentID adds the assignmentId to the patch learning assignment reschedule params
func (o *PatchLearningAssignmentRescheduleParams) SetAssignmentID(assignmentID string) {
	o.AssignmentID = assignmentID
}

// WithBody adds the body to the patch learning assignment reschedule params
func (o *PatchLearningAssignmentRescheduleParams) WithBody(body *models.LearningAssignmentReschedule) *PatchLearningAssignmentRescheduleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch learning assignment reschedule params
func (o *PatchLearningAssignmentRescheduleParams) SetBody(body *models.LearningAssignmentReschedule) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchLearningAssignmentRescheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param assignmentId
	if err := r.SetPathParam("assignmentId", o.AssignmentID); err != nil {
		return err
	}
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
