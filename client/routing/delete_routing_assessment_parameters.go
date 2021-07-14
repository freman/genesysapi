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
)

// NewDeleteRoutingAssessmentParams creates a new DeleteRoutingAssessmentParams object
// with the default values initialized.
func NewDeleteRoutingAssessmentParams() *DeleteRoutingAssessmentParams {
	var ()
	return &DeleteRoutingAssessmentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRoutingAssessmentParamsWithTimeout creates a new DeleteRoutingAssessmentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRoutingAssessmentParamsWithTimeout(timeout time.Duration) *DeleteRoutingAssessmentParams {
	var ()
	return &DeleteRoutingAssessmentParams{

		timeout: timeout,
	}
}

// NewDeleteRoutingAssessmentParamsWithContext creates a new DeleteRoutingAssessmentParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRoutingAssessmentParamsWithContext(ctx context.Context) *DeleteRoutingAssessmentParams {
	var ()
	return &DeleteRoutingAssessmentParams{

		Context: ctx,
	}
}

// NewDeleteRoutingAssessmentParamsWithHTTPClient creates a new DeleteRoutingAssessmentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRoutingAssessmentParamsWithHTTPClient(client *http.Client) *DeleteRoutingAssessmentParams {
	var ()
	return &DeleteRoutingAssessmentParams{
		HTTPClient: client,
	}
}

/*DeleteRoutingAssessmentParams contains all the parameters to send to the API endpoint
for the delete routing assessment operation typically these are written to a http.Request
*/
type DeleteRoutingAssessmentParams struct {

	/*AssessmentID
	  Benefit Assessment ID

	*/
	AssessmentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete routing assessment params
func (o *DeleteRoutingAssessmentParams) WithTimeout(timeout time.Duration) *DeleteRoutingAssessmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete routing assessment params
func (o *DeleteRoutingAssessmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete routing assessment params
func (o *DeleteRoutingAssessmentParams) WithContext(ctx context.Context) *DeleteRoutingAssessmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete routing assessment params
func (o *DeleteRoutingAssessmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete routing assessment params
func (o *DeleteRoutingAssessmentParams) WithHTTPClient(client *http.Client) *DeleteRoutingAssessmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete routing assessment params
func (o *DeleteRoutingAssessmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssessmentID adds the assessmentID to the delete routing assessment params
func (o *DeleteRoutingAssessmentParams) WithAssessmentID(assessmentID string) *DeleteRoutingAssessmentParams {
	o.SetAssessmentID(assessmentID)
	return o
}

// SetAssessmentID adds the assessmentId to the delete routing assessment params
func (o *DeleteRoutingAssessmentParams) SetAssessmentID(assessmentID string) {
	o.AssessmentID = assessmentID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRoutingAssessmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param assessmentId
	if err := r.SetPathParam("assessmentId", o.AssessmentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}