// Code generated by go-swagger; DO NOT EDIT.

package usage

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

// NewPostUsageQueryParams creates a new PostUsageQueryParams object
// with the default values initialized.
func NewPostUsageQueryParams() *PostUsageQueryParams {
	var ()
	return &PostUsageQueryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostUsageQueryParamsWithTimeout creates a new PostUsageQueryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostUsageQueryParamsWithTimeout(timeout time.Duration) *PostUsageQueryParams {
	var ()
	return &PostUsageQueryParams{

		timeout: timeout,
	}
}

// NewPostUsageQueryParamsWithContext creates a new PostUsageQueryParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostUsageQueryParamsWithContext(ctx context.Context) *PostUsageQueryParams {
	var ()
	return &PostUsageQueryParams{

		Context: ctx,
	}
}

// NewPostUsageQueryParamsWithHTTPClient creates a new PostUsageQueryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostUsageQueryParamsWithHTTPClient(client *http.Client) *PostUsageQueryParams {
	var ()
	return &PostUsageQueryParams{
		HTTPClient: client,
	}
}

/*PostUsageQueryParams contains all the parameters to send to the API endpoint
for the post usage query operation typically these are written to a http.Request
*/
type PostUsageQueryParams struct {

	/*Body
	  Query

	*/
	Body *models.APIUsageQuery

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post usage query params
func (o *PostUsageQueryParams) WithTimeout(timeout time.Duration) *PostUsageQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post usage query params
func (o *PostUsageQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post usage query params
func (o *PostUsageQueryParams) WithContext(ctx context.Context) *PostUsageQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post usage query params
func (o *PostUsageQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post usage query params
func (o *PostUsageQueryParams) WithHTTPClient(client *http.Client) *PostUsageQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post usage query params
func (o *PostUsageQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post usage query params
func (o *PostUsageQueryParams) WithBody(body *models.APIUsageQuery) *PostUsageQueryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post usage query params
func (o *PostUsageQueryParams) SetBody(body *models.APIUsageQuery) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostUsageQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
