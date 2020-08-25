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

	"github.com/freman/genesysapi/models"
)

// NewPutAuthorizationRolesDefaultParams creates a new PutAuthorizationRolesDefaultParams object
// with the default values initialized.
func NewPutAuthorizationRolesDefaultParams() *PutAuthorizationRolesDefaultParams {
	var ()
	return &PutAuthorizationRolesDefaultParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAuthorizationRolesDefaultParamsWithTimeout creates a new PutAuthorizationRolesDefaultParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAuthorizationRolesDefaultParamsWithTimeout(timeout time.Duration) *PutAuthorizationRolesDefaultParams {
	var ()
	return &PutAuthorizationRolesDefaultParams{

		timeout: timeout,
	}
}

// NewPutAuthorizationRolesDefaultParamsWithContext creates a new PutAuthorizationRolesDefaultParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutAuthorizationRolesDefaultParamsWithContext(ctx context.Context) *PutAuthorizationRolesDefaultParams {
	var ()
	return &PutAuthorizationRolesDefaultParams{

		Context: ctx,
	}
}

// NewPutAuthorizationRolesDefaultParamsWithHTTPClient creates a new PutAuthorizationRolesDefaultParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutAuthorizationRolesDefaultParamsWithHTTPClient(client *http.Client) *PutAuthorizationRolesDefaultParams {
	var ()
	return &PutAuthorizationRolesDefaultParams{
		HTTPClient: client,
	}
}

/*PutAuthorizationRolesDefaultParams contains all the parameters to send to the API endpoint
for the put authorization roles default operation typically these are written to a http.Request
*/
type PutAuthorizationRolesDefaultParams struct {

	/*Body
	  Organization roles list

	*/
	Body []*models.DomainOrganizationRole

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put authorization roles default params
func (o *PutAuthorizationRolesDefaultParams) WithTimeout(timeout time.Duration) *PutAuthorizationRolesDefaultParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put authorization roles default params
func (o *PutAuthorizationRolesDefaultParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put authorization roles default params
func (o *PutAuthorizationRolesDefaultParams) WithContext(ctx context.Context) *PutAuthorizationRolesDefaultParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put authorization roles default params
func (o *PutAuthorizationRolesDefaultParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put authorization roles default params
func (o *PutAuthorizationRolesDefaultParams) WithHTTPClient(client *http.Client) *PutAuthorizationRolesDefaultParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put authorization roles default params
func (o *PutAuthorizationRolesDefaultParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put authorization roles default params
func (o *PutAuthorizationRolesDefaultParams) WithBody(body []*models.DomainOrganizationRole) *PutAuthorizationRolesDefaultParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put authorization roles default params
func (o *PutAuthorizationRolesDefaultParams) SetBody(body []*models.DomainOrganizationRole) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutAuthorizationRolesDefaultParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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