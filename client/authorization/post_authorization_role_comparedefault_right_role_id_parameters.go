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

// NewPostAuthorizationRoleComparedefaultRightRoleIDParams creates a new PostAuthorizationRoleComparedefaultRightRoleIDParams object
// with the default values initialized.
func NewPostAuthorizationRoleComparedefaultRightRoleIDParams() *PostAuthorizationRoleComparedefaultRightRoleIDParams {
	var ()
	return &PostAuthorizationRoleComparedefaultRightRoleIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAuthorizationRoleComparedefaultRightRoleIDParamsWithTimeout creates a new PostAuthorizationRoleComparedefaultRightRoleIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAuthorizationRoleComparedefaultRightRoleIDParamsWithTimeout(timeout time.Duration) *PostAuthorizationRoleComparedefaultRightRoleIDParams {
	var ()
	return &PostAuthorizationRoleComparedefaultRightRoleIDParams{

		timeout: timeout,
	}
}

// NewPostAuthorizationRoleComparedefaultRightRoleIDParamsWithContext creates a new PostAuthorizationRoleComparedefaultRightRoleIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAuthorizationRoleComparedefaultRightRoleIDParamsWithContext(ctx context.Context) *PostAuthorizationRoleComparedefaultRightRoleIDParams {
	var ()
	return &PostAuthorizationRoleComparedefaultRightRoleIDParams{

		Context: ctx,
	}
}

// NewPostAuthorizationRoleComparedefaultRightRoleIDParamsWithHTTPClient creates a new PostAuthorizationRoleComparedefaultRightRoleIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAuthorizationRoleComparedefaultRightRoleIDParamsWithHTTPClient(client *http.Client) *PostAuthorizationRoleComparedefaultRightRoleIDParams {
	var ()
	return &PostAuthorizationRoleComparedefaultRightRoleIDParams{
		HTTPClient: client,
	}
}

/*PostAuthorizationRoleComparedefaultRightRoleIDParams contains all the parameters to send to the API endpoint
for the post authorization role comparedefault right role Id operation typically these are written to a http.Request
*/
type PostAuthorizationRoleComparedefaultRightRoleIDParams struct {

	/*Body
	  Organization role

	*/
	Body *models.DomainOrganizationRole
	/*LeftRoleID
	  Left Role ID

	*/
	LeftRoleID string
	/*RightRoleID
	  Right Role id

	*/
	RightRoleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post authorization role comparedefault right role Id params
func (o *PostAuthorizationRoleComparedefaultRightRoleIDParams) WithTimeout(timeout time.Duration) *PostAuthorizationRoleComparedefaultRightRoleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post authorization role comparedefault right role Id params
func (o *PostAuthorizationRoleComparedefaultRightRoleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post authorization role comparedefault right role Id params
func (o *PostAuthorizationRoleComparedefaultRightRoleIDParams) WithContext(ctx context.Context) *PostAuthorizationRoleComparedefaultRightRoleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post authorization role comparedefault right role Id params
func (o *PostAuthorizationRoleComparedefaultRightRoleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post authorization role comparedefault right role Id params
func (o *PostAuthorizationRoleComparedefaultRightRoleIDParams) WithHTTPClient(client *http.Client) *PostAuthorizationRoleComparedefaultRightRoleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post authorization role comparedefault right role Id params
func (o *PostAuthorizationRoleComparedefaultRightRoleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post authorization role comparedefault right role Id params
func (o *PostAuthorizationRoleComparedefaultRightRoleIDParams) WithBody(body *models.DomainOrganizationRole) *PostAuthorizationRoleComparedefaultRightRoleIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post authorization role comparedefault right role Id params
func (o *PostAuthorizationRoleComparedefaultRightRoleIDParams) SetBody(body *models.DomainOrganizationRole) {
	o.Body = body
}

// WithLeftRoleID adds the leftRoleID to the post authorization role comparedefault right role Id params
func (o *PostAuthorizationRoleComparedefaultRightRoleIDParams) WithLeftRoleID(leftRoleID string) *PostAuthorizationRoleComparedefaultRightRoleIDParams {
	o.SetLeftRoleID(leftRoleID)
	return o
}

// SetLeftRoleID adds the leftRoleId to the post authorization role comparedefault right role Id params
func (o *PostAuthorizationRoleComparedefaultRightRoleIDParams) SetLeftRoleID(leftRoleID string) {
	o.LeftRoleID = leftRoleID
}

// WithRightRoleID adds the rightRoleID to the post authorization role comparedefault right role Id params
func (o *PostAuthorizationRoleComparedefaultRightRoleIDParams) WithRightRoleID(rightRoleID string) *PostAuthorizationRoleComparedefaultRightRoleIDParams {
	o.SetRightRoleID(rightRoleID)
	return o
}

// SetRightRoleID adds the rightRoleId to the post authorization role comparedefault right role Id params
func (o *PostAuthorizationRoleComparedefaultRightRoleIDParams) SetRightRoleID(rightRoleID string) {
	o.RightRoleID = rightRoleID
}

// WriteToRequest writes these params to a swagger request
func (o *PostAuthorizationRoleComparedefaultRightRoleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param leftRoleId
	if err := r.SetPathParam("leftRoleId", o.LeftRoleID); err != nil {
		return err
	}

	// path param rightRoleId
	if err := r.SetPathParam("rightRoleId", o.RightRoleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}