// Code generated by go-swagger; DO NOT EDIT.

package architect

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

// NewGetArchitectIvrParams creates a new GetArchitectIvrParams object
// with the default values initialized.
func NewGetArchitectIvrParams() *GetArchitectIvrParams {
	var ()
	return &GetArchitectIvrParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetArchitectIvrParamsWithTimeout creates a new GetArchitectIvrParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetArchitectIvrParamsWithTimeout(timeout time.Duration) *GetArchitectIvrParams {
	var ()
	return &GetArchitectIvrParams{

		timeout: timeout,
	}
}

// NewGetArchitectIvrParamsWithContext creates a new GetArchitectIvrParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetArchitectIvrParamsWithContext(ctx context.Context) *GetArchitectIvrParams {
	var ()
	return &GetArchitectIvrParams{

		Context: ctx,
	}
}

// NewGetArchitectIvrParamsWithHTTPClient creates a new GetArchitectIvrParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetArchitectIvrParamsWithHTTPClient(client *http.Client) *GetArchitectIvrParams {
	var ()
	return &GetArchitectIvrParams{
		HTTPClient: client,
	}
}

/*GetArchitectIvrParams contains all the parameters to send to the API endpoint
for the get architect ivr operation typically these are written to a http.Request
*/
type GetArchitectIvrParams struct {

	/*IvrID
	  IVR id

	*/
	IvrID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get architect ivr params
func (o *GetArchitectIvrParams) WithTimeout(timeout time.Duration) *GetArchitectIvrParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get architect ivr params
func (o *GetArchitectIvrParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get architect ivr params
func (o *GetArchitectIvrParams) WithContext(ctx context.Context) *GetArchitectIvrParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get architect ivr params
func (o *GetArchitectIvrParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get architect ivr params
func (o *GetArchitectIvrParams) WithHTTPClient(client *http.Client) *GetArchitectIvrParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get architect ivr params
func (o *GetArchitectIvrParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIvrID adds the ivrID to the get architect ivr params
func (o *GetArchitectIvrParams) WithIvrID(ivrID string) *GetArchitectIvrParams {
	o.SetIvrID(ivrID)
	return o
}

// SetIvrID adds the ivrId to the get architect ivr params
func (o *GetArchitectIvrParams) SetIvrID(ivrID string) {
	o.IvrID = ivrID
}

// WriteToRequest writes these params to a swagger request
func (o *GetArchitectIvrParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ivrId
	if err := r.SetPathParam("ivrId", o.IvrID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}