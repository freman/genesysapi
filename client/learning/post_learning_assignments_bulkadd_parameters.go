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

// NewPostLearningAssignmentsBulkaddParams creates a new PostLearningAssignmentsBulkaddParams object
// with the default values initialized.
func NewPostLearningAssignmentsBulkaddParams() *PostLearningAssignmentsBulkaddParams {
	var ()
	return &PostLearningAssignmentsBulkaddParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLearningAssignmentsBulkaddParamsWithTimeout creates a new PostLearningAssignmentsBulkaddParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLearningAssignmentsBulkaddParamsWithTimeout(timeout time.Duration) *PostLearningAssignmentsBulkaddParams {
	var ()
	return &PostLearningAssignmentsBulkaddParams{

		timeout: timeout,
	}
}

// NewPostLearningAssignmentsBulkaddParamsWithContext creates a new PostLearningAssignmentsBulkaddParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLearningAssignmentsBulkaddParamsWithContext(ctx context.Context) *PostLearningAssignmentsBulkaddParams {
	var ()
	return &PostLearningAssignmentsBulkaddParams{

		Context: ctx,
	}
}

// NewPostLearningAssignmentsBulkaddParamsWithHTTPClient creates a new PostLearningAssignmentsBulkaddParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLearningAssignmentsBulkaddParamsWithHTTPClient(client *http.Client) *PostLearningAssignmentsBulkaddParams {
	var ()
	return &PostLearningAssignmentsBulkaddParams{
		HTTPClient: client,
	}
}

/*PostLearningAssignmentsBulkaddParams contains all the parameters to send to the API endpoint
for the post learning assignments bulkadd operation typically these are written to a http.Request
*/
type PostLearningAssignmentsBulkaddParams struct {

	/*Body
	  The learning assignments to be created

	*/
	Body []*models.LearningAssignmentItem

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post learning assignments bulkadd params
func (o *PostLearningAssignmentsBulkaddParams) WithTimeout(timeout time.Duration) *PostLearningAssignmentsBulkaddParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post learning assignments bulkadd params
func (o *PostLearningAssignmentsBulkaddParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post learning assignments bulkadd params
func (o *PostLearningAssignmentsBulkaddParams) WithContext(ctx context.Context) *PostLearningAssignmentsBulkaddParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post learning assignments bulkadd params
func (o *PostLearningAssignmentsBulkaddParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post learning assignments bulkadd params
func (o *PostLearningAssignmentsBulkaddParams) WithHTTPClient(client *http.Client) *PostLearningAssignmentsBulkaddParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post learning assignments bulkadd params
func (o *PostLearningAssignmentsBulkaddParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post learning assignments bulkadd params
func (o *PostLearningAssignmentsBulkaddParams) WithBody(body []*models.LearningAssignmentItem) *PostLearningAssignmentsBulkaddParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post learning assignments bulkadd params
func (o *PostLearningAssignmentsBulkaddParams) SetBody(body []*models.LearningAssignmentItem) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostLearningAssignmentsBulkaddParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
