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

// NewPostAuthorizationDivisionRestoreParams creates a new PostAuthorizationDivisionRestoreParams object
// with the default values initialized.
func NewPostAuthorizationDivisionRestoreParams() *PostAuthorizationDivisionRestoreParams {
	var ()
	return &PostAuthorizationDivisionRestoreParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAuthorizationDivisionRestoreParamsWithTimeout creates a new PostAuthorizationDivisionRestoreParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAuthorizationDivisionRestoreParamsWithTimeout(timeout time.Duration) *PostAuthorizationDivisionRestoreParams {
	var ()
	return &PostAuthorizationDivisionRestoreParams{

		timeout: timeout,
	}
}

// NewPostAuthorizationDivisionRestoreParamsWithContext creates a new PostAuthorizationDivisionRestoreParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAuthorizationDivisionRestoreParamsWithContext(ctx context.Context) *PostAuthorizationDivisionRestoreParams {
	var ()
	return &PostAuthorizationDivisionRestoreParams{

		Context: ctx,
	}
}

// NewPostAuthorizationDivisionRestoreParamsWithHTTPClient creates a new PostAuthorizationDivisionRestoreParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAuthorizationDivisionRestoreParamsWithHTTPClient(client *http.Client) *PostAuthorizationDivisionRestoreParams {
	var ()
	return &PostAuthorizationDivisionRestoreParams{
		HTTPClient: client,
	}
}

/*PostAuthorizationDivisionRestoreParams contains all the parameters to send to the API endpoint
for the post authorization division restore operation typically these are written to a http.Request
*/
type PostAuthorizationDivisionRestoreParams struct {

	/*Body
	  Recreated division data

	*/
	Body *models.AuthzDivision
	/*DivisionID
	  Division ID

	*/
	DivisionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post authorization division restore params
func (o *PostAuthorizationDivisionRestoreParams) WithTimeout(timeout time.Duration) *PostAuthorizationDivisionRestoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post authorization division restore params
func (o *PostAuthorizationDivisionRestoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post authorization division restore params
func (o *PostAuthorizationDivisionRestoreParams) WithContext(ctx context.Context) *PostAuthorizationDivisionRestoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post authorization division restore params
func (o *PostAuthorizationDivisionRestoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post authorization division restore params
func (o *PostAuthorizationDivisionRestoreParams) WithHTTPClient(client *http.Client) *PostAuthorizationDivisionRestoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post authorization division restore params
func (o *PostAuthorizationDivisionRestoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post authorization division restore params
func (o *PostAuthorizationDivisionRestoreParams) WithBody(body *models.AuthzDivision) *PostAuthorizationDivisionRestoreParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post authorization division restore params
func (o *PostAuthorizationDivisionRestoreParams) SetBody(body *models.AuthzDivision) {
	o.Body = body
}

// WithDivisionID adds the divisionID to the post authorization division restore params
func (o *PostAuthorizationDivisionRestoreParams) WithDivisionID(divisionID string) *PostAuthorizationDivisionRestoreParams {
	o.SetDivisionID(divisionID)
	return o
}

// SetDivisionID adds the divisionId to the post authorization division restore params
func (o *PostAuthorizationDivisionRestoreParams) SetDivisionID(divisionID string) {
	o.DivisionID = divisionID
}

// WriteToRequest writes these params to a swagger request
func (o *PostAuthorizationDivisionRestoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param divisionId
	if err := r.SetPathParam("divisionId", o.DivisionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}