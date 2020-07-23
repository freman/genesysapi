// Code generated by go-swagger; DO NOT EDIT.

package outbound

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

// NewPostOutboundDnclistsParams creates a new PostOutboundDnclistsParams object
// with the default values initialized.
func NewPostOutboundDnclistsParams() *PostOutboundDnclistsParams {
	var ()
	return &PostOutboundDnclistsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostOutboundDnclistsParamsWithTimeout creates a new PostOutboundDnclistsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostOutboundDnclistsParamsWithTimeout(timeout time.Duration) *PostOutboundDnclistsParams {
	var ()
	return &PostOutboundDnclistsParams{

		timeout: timeout,
	}
}

// NewPostOutboundDnclistsParamsWithContext creates a new PostOutboundDnclistsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostOutboundDnclistsParamsWithContext(ctx context.Context) *PostOutboundDnclistsParams {
	var ()
	return &PostOutboundDnclistsParams{

		Context: ctx,
	}
}

// NewPostOutboundDnclistsParamsWithHTTPClient creates a new PostOutboundDnclistsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostOutboundDnclistsParamsWithHTTPClient(client *http.Client) *PostOutboundDnclistsParams {
	var ()
	return &PostOutboundDnclistsParams{
		HTTPClient: client,
	}
}

/*PostOutboundDnclistsParams contains all the parameters to send to the API endpoint
for the post outbound dnclists operation typically these are written to a http.Request
*/
type PostOutboundDnclistsParams struct {

	/*Body
	  DncList

	*/
	Body *models.DncListCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post outbound dnclists params
func (o *PostOutboundDnclistsParams) WithTimeout(timeout time.Duration) *PostOutboundDnclistsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post outbound dnclists params
func (o *PostOutboundDnclistsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post outbound dnclists params
func (o *PostOutboundDnclistsParams) WithContext(ctx context.Context) *PostOutboundDnclistsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post outbound dnclists params
func (o *PostOutboundDnclistsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post outbound dnclists params
func (o *PostOutboundDnclistsParams) WithHTTPClient(client *http.Client) *PostOutboundDnclistsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post outbound dnclists params
func (o *PostOutboundDnclistsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post outbound dnclists params
func (o *PostOutboundDnclistsParams) WithBody(body *models.DncListCreate) *PostOutboundDnclistsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post outbound dnclists params
func (o *PostOutboundDnclistsParams) SetBody(body *models.DncListCreate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostOutboundDnclistsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
