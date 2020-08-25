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

// NewPatchRoutingQueueUsersParams creates a new PatchRoutingQueueUsersParams object
// with the default values initialized.
func NewPatchRoutingQueueUsersParams() *PatchRoutingQueueUsersParams {
	var ()
	return &PatchRoutingQueueUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchRoutingQueueUsersParamsWithTimeout creates a new PatchRoutingQueueUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchRoutingQueueUsersParamsWithTimeout(timeout time.Duration) *PatchRoutingQueueUsersParams {
	var ()
	return &PatchRoutingQueueUsersParams{

		timeout: timeout,
	}
}

// NewPatchRoutingQueueUsersParamsWithContext creates a new PatchRoutingQueueUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchRoutingQueueUsersParamsWithContext(ctx context.Context) *PatchRoutingQueueUsersParams {
	var ()
	return &PatchRoutingQueueUsersParams{

		Context: ctx,
	}
}

// NewPatchRoutingQueueUsersParamsWithHTTPClient creates a new PatchRoutingQueueUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchRoutingQueueUsersParamsWithHTTPClient(client *http.Client) *PatchRoutingQueueUsersParams {
	var ()
	return &PatchRoutingQueueUsersParams{
		HTTPClient: client,
	}
}

/*PatchRoutingQueueUsersParams contains all the parameters to send to the API endpoint
for the patch routing queue users operation typically these are written to a http.Request
*/
type PatchRoutingQueueUsersParams struct {

	/*Body
	  Queue Members

	*/
	Body []*models.QueueMember
	/*QueueID
	  Queue ID

	*/
	QueueID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch routing queue users params
func (o *PatchRoutingQueueUsersParams) WithTimeout(timeout time.Duration) *PatchRoutingQueueUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch routing queue users params
func (o *PatchRoutingQueueUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch routing queue users params
func (o *PatchRoutingQueueUsersParams) WithContext(ctx context.Context) *PatchRoutingQueueUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch routing queue users params
func (o *PatchRoutingQueueUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch routing queue users params
func (o *PatchRoutingQueueUsersParams) WithHTTPClient(client *http.Client) *PatchRoutingQueueUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch routing queue users params
func (o *PatchRoutingQueueUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch routing queue users params
func (o *PatchRoutingQueueUsersParams) WithBody(body []*models.QueueMember) *PatchRoutingQueueUsersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch routing queue users params
func (o *PatchRoutingQueueUsersParams) SetBody(body []*models.QueueMember) {
	o.Body = body
}

// WithQueueID adds the queueID to the patch routing queue users params
func (o *PatchRoutingQueueUsersParams) WithQueueID(queueID string) *PatchRoutingQueueUsersParams {
	o.SetQueueID(queueID)
	return o
}

// SetQueueID adds the queueId to the patch routing queue users params
func (o *PatchRoutingQueueUsersParams) SetQueueID(queueID string) {
	o.QueueID = queueID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchRoutingQueueUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}