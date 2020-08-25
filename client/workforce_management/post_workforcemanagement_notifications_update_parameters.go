// Code generated by go-swagger; DO NOT EDIT.

package workforce_management

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

// NewPostWorkforcemanagementNotificationsUpdateParams creates a new PostWorkforcemanagementNotificationsUpdateParams object
// with the default values initialized.
func NewPostWorkforcemanagementNotificationsUpdateParams() *PostWorkforcemanagementNotificationsUpdateParams {
	var ()
	return &PostWorkforcemanagementNotificationsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkforcemanagementNotificationsUpdateParamsWithTimeout creates a new PostWorkforcemanagementNotificationsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostWorkforcemanagementNotificationsUpdateParamsWithTimeout(timeout time.Duration) *PostWorkforcemanagementNotificationsUpdateParams {
	var ()
	return &PostWorkforcemanagementNotificationsUpdateParams{

		timeout: timeout,
	}
}

// NewPostWorkforcemanagementNotificationsUpdateParamsWithContext creates a new PostWorkforcemanagementNotificationsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostWorkforcemanagementNotificationsUpdateParamsWithContext(ctx context.Context) *PostWorkforcemanagementNotificationsUpdateParams {
	var ()
	return &PostWorkforcemanagementNotificationsUpdateParams{

		Context: ctx,
	}
}

// NewPostWorkforcemanagementNotificationsUpdateParamsWithHTTPClient creates a new PostWorkforcemanagementNotificationsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostWorkforcemanagementNotificationsUpdateParamsWithHTTPClient(client *http.Client) *PostWorkforcemanagementNotificationsUpdateParams {
	var ()
	return &PostWorkforcemanagementNotificationsUpdateParams{
		HTTPClient: client,
	}
}

/*PostWorkforcemanagementNotificationsUpdateParams contains all the parameters to send to the API endpoint
for the post workforcemanagement notifications update operation typically these are written to a http.Request
*/
type PostWorkforcemanagementNotificationsUpdateParams struct {

	/*Body
	  body

	*/
	Body *models.UpdateNotificationsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post workforcemanagement notifications update params
func (o *PostWorkforcemanagementNotificationsUpdateParams) WithTimeout(timeout time.Duration) *PostWorkforcemanagementNotificationsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workforcemanagement notifications update params
func (o *PostWorkforcemanagementNotificationsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workforcemanagement notifications update params
func (o *PostWorkforcemanagementNotificationsUpdateParams) WithContext(ctx context.Context) *PostWorkforcemanagementNotificationsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workforcemanagement notifications update params
func (o *PostWorkforcemanagementNotificationsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workforcemanagement notifications update params
func (o *PostWorkforcemanagementNotificationsUpdateParams) WithHTTPClient(client *http.Client) *PostWorkforcemanagementNotificationsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workforcemanagement notifications update params
func (o *PostWorkforcemanagementNotificationsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post workforcemanagement notifications update params
func (o *PostWorkforcemanagementNotificationsUpdateParams) WithBody(body *models.UpdateNotificationsRequest) *PostWorkforcemanagementNotificationsUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post workforcemanagement notifications update params
func (o *PostWorkforcemanagementNotificationsUpdateParams) SetBody(body *models.UpdateNotificationsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkforcemanagementNotificationsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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