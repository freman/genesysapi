// Code generated by go-swagger; DO NOT EDIT.

package content_management

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

// NewPostContentmanagementSharesParams creates a new PostContentmanagementSharesParams object
// with the default values initialized.
func NewPostContentmanagementSharesParams() *PostContentmanagementSharesParams {
	var ()
	return &PostContentmanagementSharesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostContentmanagementSharesParamsWithTimeout creates a new PostContentmanagementSharesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostContentmanagementSharesParamsWithTimeout(timeout time.Duration) *PostContentmanagementSharesParams {
	var ()
	return &PostContentmanagementSharesParams{

		timeout: timeout,
	}
}

// NewPostContentmanagementSharesParamsWithContext creates a new PostContentmanagementSharesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostContentmanagementSharesParamsWithContext(ctx context.Context) *PostContentmanagementSharesParams {
	var ()
	return &PostContentmanagementSharesParams{

		Context: ctx,
	}
}

// NewPostContentmanagementSharesParamsWithHTTPClient creates a new PostContentmanagementSharesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostContentmanagementSharesParamsWithHTTPClient(client *http.Client) *PostContentmanagementSharesParams {
	var ()
	return &PostContentmanagementSharesParams{
		HTTPClient: client,
	}
}

/*PostContentmanagementSharesParams contains all the parameters to send to the API endpoint
for the post contentmanagement shares operation typically these are written to a http.Request
*/
type PostContentmanagementSharesParams struct {

	/*Body
	  CreateShareRequest - entity id and type and a single member or list of members are required

	*/
	Body *models.CreateShareRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post contentmanagement shares params
func (o *PostContentmanagementSharesParams) WithTimeout(timeout time.Duration) *PostContentmanagementSharesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post contentmanagement shares params
func (o *PostContentmanagementSharesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post contentmanagement shares params
func (o *PostContentmanagementSharesParams) WithContext(ctx context.Context) *PostContentmanagementSharesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post contentmanagement shares params
func (o *PostContentmanagementSharesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post contentmanagement shares params
func (o *PostContentmanagementSharesParams) WithHTTPClient(client *http.Client) *PostContentmanagementSharesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post contentmanagement shares params
func (o *PostContentmanagementSharesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post contentmanagement shares params
func (o *PostContentmanagementSharesParams) WithBody(body *models.CreateShareRequest) *PostContentmanagementSharesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post contentmanagement shares params
func (o *PostContentmanagementSharesParams) SetBody(body *models.CreateShareRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostContentmanagementSharesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
