// Code generated by go-swagger; DO NOT EDIT.

package authorization

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

// NewPutUserRolesParams creates a new PutUserRolesParams object
// with the default values initialized.
func NewPutUserRolesParams() *PutUserRolesParams {
	var ()
	return &PutUserRolesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutUserRolesParamsWithTimeout creates a new PutUserRolesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutUserRolesParamsWithTimeout(timeout time.Duration) *PutUserRolesParams {
	var ()
	return &PutUserRolesParams{

		timeout: timeout,
	}
}

// NewPutUserRolesParamsWithContext creates a new PutUserRolesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutUserRolesParamsWithContext(ctx context.Context) *PutUserRolesParams {
	var ()
	return &PutUserRolesParams{

		Context: ctx,
	}
}

// NewPutUserRolesParamsWithHTTPClient creates a new PutUserRolesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutUserRolesParamsWithHTTPClient(client *http.Client) *PutUserRolesParams {
	var ()
	return &PutUserRolesParams{
		HTTPClient: client,
	}
}

/*PutUserRolesParams contains all the parameters to send to the API endpoint
for the put user roles operation typically these are written to a http.Request
*/
type PutUserRolesParams struct {

	/*Body
	  List of roles

	*/
	Body []string
	/*SubjectID
	  User ID

	*/
	SubjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put user roles params
func (o *PutUserRolesParams) WithTimeout(timeout time.Duration) *PutUserRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put user roles params
func (o *PutUserRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put user roles params
func (o *PutUserRolesParams) WithContext(ctx context.Context) *PutUserRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put user roles params
func (o *PutUserRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put user roles params
func (o *PutUserRolesParams) WithHTTPClient(client *http.Client) *PutUserRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put user roles params
func (o *PutUserRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put user roles params
func (o *PutUserRolesParams) WithBody(body []string) *PutUserRolesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put user roles params
func (o *PutUserRolesParams) SetBody(body []string) {
	o.Body = body
}

// WithSubjectID adds the subjectID to the put user roles params
func (o *PutUserRolesParams) WithSubjectID(subjectID string) *PutUserRolesParams {
	o.SetSubjectID(subjectID)
	return o
}

// SetSubjectID adds the subjectId to the put user roles params
func (o *PutUserRolesParams) SetSubjectID(subjectID string) {
	o.SubjectID = subjectID
}

// WriteToRequest writes these params to a swagger request
func (o *PutUserRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param subjectId
	if err := r.SetPathParam("subjectId", o.SubjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
