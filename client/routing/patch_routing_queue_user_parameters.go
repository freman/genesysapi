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

// NewPatchRoutingQueueUserParams creates a new PatchRoutingQueueUserParams object
// with the default values initialized.
func NewPatchRoutingQueueUserParams() *PatchRoutingQueueUserParams {
	var ()
	return &PatchRoutingQueueUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchRoutingQueueUserParamsWithTimeout creates a new PatchRoutingQueueUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchRoutingQueueUserParamsWithTimeout(timeout time.Duration) *PatchRoutingQueueUserParams {
	var ()
	return &PatchRoutingQueueUserParams{

		timeout: timeout,
	}
}

// NewPatchRoutingQueueUserParamsWithContext creates a new PatchRoutingQueueUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchRoutingQueueUserParamsWithContext(ctx context.Context) *PatchRoutingQueueUserParams {
	var ()
	return &PatchRoutingQueueUserParams{

		Context: ctx,
	}
}

// NewPatchRoutingQueueUserParamsWithHTTPClient creates a new PatchRoutingQueueUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchRoutingQueueUserParamsWithHTTPClient(client *http.Client) *PatchRoutingQueueUserParams {
	var ()
	return &PatchRoutingQueueUserParams{
		HTTPClient: client,
	}
}

/*PatchRoutingQueueUserParams contains all the parameters to send to the API endpoint
for the patch routing queue user operation typically these are written to a http.Request
*/
type PatchRoutingQueueUserParams struct {

	/*Body
	  Queue Member

	*/
	Body *models.QueueMember
	/*MemberID
	  Member ID

	*/
	MemberID string
	/*QueueID
	  Queue ID

	*/
	QueueID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch routing queue user params
func (o *PatchRoutingQueueUserParams) WithTimeout(timeout time.Duration) *PatchRoutingQueueUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch routing queue user params
func (o *PatchRoutingQueueUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch routing queue user params
func (o *PatchRoutingQueueUserParams) WithContext(ctx context.Context) *PatchRoutingQueueUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch routing queue user params
func (o *PatchRoutingQueueUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch routing queue user params
func (o *PatchRoutingQueueUserParams) WithHTTPClient(client *http.Client) *PatchRoutingQueueUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch routing queue user params
func (o *PatchRoutingQueueUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch routing queue user params
func (o *PatchRoutingQueueUserParams) WithBody(body *models.QueueMember) *PatchRoutingQueueUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch routing queue user params
func (o *PatchRoutingQueueUserParams) SetBody(body *models.QueueMember) {
	o.Body = body
}

// WithMemberID adds the memberID to the patch routing queue user params
func (o *PatchRoutingQueueUserParams) WithMemberID(memberID string) *PatchRoutingQueueUserParams {
	o.SetMemberID(memberID)
	return o
}

// SetMemberID adds the memberId to the patch routing queue user params
func (o *PatchRoutingQueueUserParams) SetMemberID(memberID string) {
	o.MemberID = memberID
}

// WithQueueID adds the queueID to the patch routing queue user params
func (o *PatchRoutingQueueUserParams) WithQueueID(queueID string) *PatchRoutingQueueUserParams {
	o.SetQueueID(queueID)
	return o
}

// SetQueueID adds the queueId to the patch routing queue user params
func (o *PatchRoutingQueueUserParams) SetQueueID(queueID string) {
	o.QueueID = queueID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchRoutingQueueUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param memberId
	if err := r.SetPathParam("memberId", o.MemberID); err != nil {
		return err
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
