// Code generated by go-swagger; DO NOT EDIT.

package search

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

// NewPostUsersSearchParams creates a new PostUsersSearchParams object
// with the default values initialized.
func NewPostUsersSearchParams() *PostUsersSearchParams {
	var ()
	return &PostUsersSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostUsersSearchParamsWithTimeout creates a new PostUsersSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostUsersSearchParamsWithTimeout(timeout time.Duration) *PostUsersSearchParams {
	var ()
	return &PostUsersSearchParams{

		timeout: timeout,
	}
}

// NewPostUsersSearchParamsWithContext creates a new PostUsersSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostUsersSearchParamsWithContext(ctx context.Context) *PostUsersSearchParams {
	var ()
	return &PostUsersSearchParams{

		Context: ctx,
	}
}

// NewPostUsersSearchParamsWithHTTPClient creates a new PostUsersSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostUsersSearchParamsWithHTTPClient(client *http.Client) *PostUsersSearchParams {
	var ()
	return &PostUsersSearchParams{
		HTTPClient: client,
	}
}

/*PostUsersSearchParams contains all the parameters to send to the API endpoint
for the post users search operation typically these are written to a http.Request
*/
type PostUsersSearchParams struct {

	/*Body
	  Search request options

	*/
	Body *models.UserSearchRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post users search params
func (o *PostUsersSearchParams) WithTimeout(timeout time.Duration) *PostUsersSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post users search params
func (o *PostUsersSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post users search params
func (o *PostUsersSearchParams) WithContext(ctx context.Context) *PostUsersSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post users search params
func (o *PostUsersSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post users search params
func (o *PostUsersSearchParams) WithHTTPClient(client *http.Client) *PostUsersSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post users search params
func (o *PostUsersSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post users search params
func (o *PostUsersSearchParams) WithBody(body *models.UserSearchRequest) *PostUsersSearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post users search params
func (o *PostUsersSearchParams) SetBody(body *models.UserSearchRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostUsersSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
