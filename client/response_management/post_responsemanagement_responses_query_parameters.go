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

// NewPostResponsemanagementResponsesQueryParams creates a new PostResponsemanagementResponsesQueryParams object
// with the default values initialized.
func NewPostResponsemanagementResponsesQueryParams() *PostResponsemanagementResponsesQueryParams {
	var ()
	return &PostResponsemanagementResponsesQueryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostResponsemanagementResponsesQueryParamsWithTimeout creates a new PostResponsemanagementResponsesQueryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostResponsemanagementResponsesQueryParamsWithTimeout(timeout time.Duration) *PostResponsemanagementResponsesQueryParams {
	var ()
	return &PostResponsemanagementResponsesQueryParams{

		timeout: timeout,
	}
}

// NewPostResponsemanagementResponsesQueryParamsWithContext creates a new PostResponsemanagementResponsesQueryParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostResponsemanagementResponsesQueryParamsWithContext(ctx context.Context) *PostResponsemanagementResponsesQueryParams {
	var ()
	return &PostResponsemanagementResponsesQueryParams{

		Context: ctx,
	}
}

// NewPostResponsemanagementResponsesQueryParamsWithHTTPClient creates a new PostResponsemanagementResponsesQueryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostResponsemanagementResponsesQueryParamsWithHTTPClient(client *http.Client) *PostResponsemanagementResponsesQueryParams {
	var ()
	return &PostResponsemanagementResponsesQueryParams{
		HTTPClient: client,
	}
}

/*PostResponsemanagementResponsesQueryParams contains all the parameters to send to the API endpoint
for the post responsemanagement responses query operation typically these are written to a http.Request
*/
type PostResponsemanagementResponsesQueryParams struct {

	/*Body
	  Response

	*/
	Body *models.ResponseQueryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post responsemanagement responses query params
func (o *PostResponsemanagementResponsesQueryParams) WithTimeout(timeout time.Duration) *PostResponsemanagementResponsesQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post responsemanagement responses query params
func (o *PostResponsemanagementResponsesQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post responsemanagement responses query params
func (o *PostResponsemanagementResponsesQueryParams) WithContext(ctx context.Context) *PostResponsemanagementResponsesQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post responsemanagement responses query params
func (o *PostResponsemanagementResponsesQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post responsemanagement responses query params
func (o *PostResponsemanagementResponsesQueryParams) WithHTTPClient(client *http.Client) *PostResponsemanagementResponsesQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post responsemanagement responses query params
func (o *PostResponsemanagementResponsesQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post responsemanagement responses query params
func (o *PostResponsemanagementResponsesQueryParams) WithBody(body *models.ResponseQueryRequest) *PostResponsemanagementResponsesQueryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post responsemanagement responses query params
func (o *PostResponsemanagementResponsesQueryParams) SetBody(body *models.ResponseQueryRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostResponsemanagementResponsesQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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