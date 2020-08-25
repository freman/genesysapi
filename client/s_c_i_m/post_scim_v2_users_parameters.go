// Code generated by go-swagger; DO NOT EDIT.

package s_c_i_m

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

// NewPostScimV2UsersParams creates a new PostScimV2UsersParams object
// with the default values initialized.
func NewPostScimV2UsersParams() *PostScimV2UsersParams {
	var ()
	return &PostScimV2UsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostScimV2UsersParamsWithTimeout creates a new PostScimV2UsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostScimV2UsersParamsWithTimeout(timeout time.Duration) *PostScimV2UsersParams {
	var ()
	return &PostScimV2UsersParams{

		timeout: timeout,
	}
}

// NewPostScimV2UsersParamsWithContext creates a new PostScimV2UsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostScimV2UsersParamsWithContext(ctx context.Context) *PostScimV2UsersParams {
	var ()
	return &PostScimV2UsersParams{

		Context: ctx,
	}
}

// NewPostScimV2UsersParamsWithHTTPClient creates a new PostScimV2UsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostScimV2UsersParamsWithHTTPClient(client *http.Client) *PostScimV2UsersParams {
	var ()
	return &PostScimV2UsersParams{
		HTTPClient: client,
	}
}

/*PostScimV2UsersParams contains all the parameters to send to the API endpoint
for the post scim v2 users operation typically these are written to a http.Request
*/
type PostScimV2UsersParams struct {

	/*Body
	  The information used to create a user.

	*/
	Body *models.ScimV2CreateUser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post scim v2 users params
func (o *PostScimV2UsersParams) WithTimeout(timeout time.Duration) *PostScimV2UsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post scim v2 users params
func (o *PostScimV2UsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post scim v2 users params
func (o *PostScimV2UsersParams) WithContext(ctx context.Context) *PostScimV2UsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post scim v2 users params
func (o *PostScimV2UsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post scim v2 users params
func (o *PostScimV2UsersParams) WithHTTPClient(client *http.Client) *PostScimV2UsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post scim v2 users params
func (o *PostScimV2UsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post scim v2 users params
func (o *PostScimV2UsersParams) WithBody(body *models.ScimV2CreateUser) *PostScimV2UsersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post scim v2 users params
func (o *PostScimV2UsersParams) SetBody(body *models.ScimV2CreateUser) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostScimV2UsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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