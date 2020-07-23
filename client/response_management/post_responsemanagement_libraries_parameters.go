// Code generated by go-swagger; DO NOT EDIT.

package response_management

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

// NewPostResponsemanagementLibrariesParams creates a new PostResponsemanagementLibrariesParams object
// with the default values initialized.
func NewPostResponsemanagementLibrariesParams() *PostResponsemanagementLibrariesParams {
	var ()
	return &PostResponsemanagementLibrariesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostResponsemanagementLibrariesParamsWithTimeout creates a new PostResponsemanagementLibrariesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostResponsemanagementLibrariesParamsWithTimeout(timeout time.Duration) *PostResponsemanagementLibrariesParams {
	var ()
	return &PostResponsemanagementLibrariesParams{

		timeout: timeout,
	}
}

// NewPostResponsemanagementLibrariesParamsWithContext creates a new PostResponsemanagementLibrariesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostResponsemanagementLibrariesParamsWithContext(ctx context.Context) *PostResponsemanagementLibrariesParams {
	var ()
	return &PostResponsemanagementLibrariesParams{

		Context: ctx,
	}
}

// NewPostResponsemanagementLibrariesParamsWithHTTPClient creates a new PostResponsemanagementLibrariesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostResponsemanagementLibrariesParamsWithHTTPClient(client *http.Client) *PostResponsemanagementLibrariesParams {
	var ()
	return &PostResponsemanagementLibrariesParams{
		HTTPClient: client,
	}
}

/*PostResponsemanagementLibrariesParams contains all the parameters to send to the API endpoint
for the post responsemanagement libraries operation typically these are written to a http.Request
*/
type PostResponsemanagementLibrariesParams struct {

	/*Body
	  Library

	*/
	Body *models.Library

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post responsemanagement libraries params
func (o *PostResponsemanagementLibrariesParams) WithTimeout(timeout time.Duration) *PostResponsemanagementLibrariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post responsemanagement libraries params
func (o *PostResponsemanagementLibrariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post responsemanagement libraries params
func (o *PostResponsemanagementLibrariesParams) WithContext(ctx context.Context) *PostResponsemanagementLibrariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post responsemanagement libraries params
func (o *PostResponsemanagementLibrariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post responsemanagement libraries params
func (o *PostResponsemanagementLibrariesParams) WithHTTPClient(client *http.Client) *PostResponsemanagementLibrariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post responsemanagement libraries params
func (o *PostResponsemanagementLibrariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post responsemanagement libraries params
func (o *PostResponsemanagementLibrariesParams) WithBody(body *models.Library) *PostResponsemanagementLibrariesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post responsemanagement libraries params
func (o *PostResponsemanagementLibrariesParams) SetBody(body *models.Library) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostResponsemanagementLibrariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
