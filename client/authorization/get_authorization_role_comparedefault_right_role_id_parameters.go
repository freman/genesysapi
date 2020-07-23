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

// NewGetAuthorizationRoleComparedefaultRightRoleIDParams creates a new GetAuthorizationRoleComparedefaultRightRoleIDParams object
// with the default values initialized.
func NewGetAuthorizationRoleComparedefaultRightRoleIDParams() *GetAuthorizationRoleComparedefaultRightRoleIDParams {
	var ()
	return &GetAuthorizationRoleComparedefaultRightRoleIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthorizationRoleComparedefaultRightRoleIDParamsWithTimeout creates a new GetAuthorizationRoleComparedefaultRightRoleIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAuthorizationRoleComparedefaultRightRoleIDParamsWithTimeout(timeout time.Duration) *GetAuthorizationRoleComparedefaultRightRoleIDParams {
	var ()
	return &GetAuthorizationRoleComparedefaultRightRoleIDParams{

		timeout: timeout,
	}
}

// NewGetAuthorizationRoleComparedefaultRightRoleIDParamsWithContext creates a new GetAuthorizationRoleComparedefaultRightRoleIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAuthorizationRoleComparedefaultRightRoleIDParamsWithContext(ctx context.Context) *GetAuthorizationRoleComparedefaultRightRoleIDParams {
	var ()
	return &GetAuthorizationRoleComparedefaultRightRoleIDParams{

		Context: ctx,
	}
}

// NewGetAuthorizationRoleComparedefaultRightRoleIDParamsWithHTTPClient creates a new GetAuthorizationRoleComparedefaultRightRoleIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAuthorizationRoleComparedefaultRightRoleIDParamsWithHTTPClient(client *http.Client) *GetAuthorizationRoleComparedefaultRightRoleIDParams {
	var ()
	return &GetAuthorizationRoleComparedefaultRightRoleIDParams{
		HTTPClient: client,
	}
}

/*GetAuthorizationRoleComparedefaultRightRoleIDParams contains all the parameters to send to the API endpoint
for the get authorization role comparedefault right role Id operation typically these are written to a http.Request
*/
type GetAuthorizationRoleComparedefaultRightRoleIDParams struct {

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

// WithTimeout adds the timeout to the get authorization role comparedefault right role Id params
func (o *GetAuthorizationRoleComparedefaultRightRoleIDParams) WithTimeout(timeout time.Duration) *GetAuthorizationRoleComparedefaultRightRoleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get authorization role comparedefault right role Id params
func (o *GetAuthorizationRoleComparedefaultRightRoleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get authorization role comparedefault right role Id params
func (o *GetAuthorizationRoleComparedefaultRightRoleIDParams) WithContext(ctx context.Context) *GetAuthorizationRoleComparedefaultRightRoleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get authorization role comparedefault right role Id params
func (o *GetAuthorizationRoleComparedefaultRightRoleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get authorization role comparedefault right role Id params
func (o *GetAuthorizationRoleComparedefaultRightRoleIDParams) WithHTTPClient(client *http.Client) *GetAuthorizationRoleComparedefaultRightRoleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get authorization role comparedefault right role Id params
func (o *GetAuthorizationRoleComparedefaultRightRoleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLeftRoleID adds the leftRoleID to the get authorization role comparedefault right role Id params
func (o *GetAuthorizationRoleComparedefaultRightRoleIDParams) WithLeftRoleID(leftRoleID string) *GetAuthorizationRoleComparedefaultRightRoleIDParams {
	o.SetLeftRoleID(leftRoleID)
	return o
}

// SetLeftRoleID adds the leftRoleId to the get authorization role comparedefault right role Id params
func (o *GetAuthorizationRoleComparedefaultRightRoleIDParams) SetLeftRoleID(leftRoleID string) {
	o.LeftRoleID = leftRoleID
}

// WithRightRoleID adds the rightRoleID to the get authorization role comparedefault right role Id params
func (o *GetAuthorizationRoleComparedefaultRightRoleIDParams) WithRightRoleID(rightRoleID string) *GetAuthorizationRoleComparedefaultRightRoleIDParams {
	o.SetRightRoleID(rightRoleID)
	return o
}

// SetRightRoleID adds the rightRoleId to the get authorization role comparedefault right role Id params
func (o *GetAuthorizationRoleComparedefaultRightRoleIDParams) SetRightRoleID(rightRoleID string) {
	o.RightRoleID = rightRoleID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthorizationRoleComparedefaultRightRoleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
