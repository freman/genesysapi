// Code generated by go-swagger; DO NOT EDIT.

package external_contacts

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

// NewPostExternalcontactsBulkContactsRemoveParams creates a new PostExternalcontactsBulkContactsRemoveParams object
// with the default values initialized.
func NewPostExternalcontactsBulkContactsRemoveParams() *PostExternalcontactsBulkContactsRemoveParams {
	var ()
	return &PostExternalcontactsBulkContactsRemoveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostExternalcontactsBulkContactsRemoveParamsWithTimeout creates a new PostExternalcontactsBulkContactsRemoveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostExternalcontactsBulkContactsRemoveParamsWithTimeout(timeout time.Duration) *PostExternalcontactsBulkContactsRemoveParams {
	var ()
	return &PostExternalcontactsBulkContactsRemoveParams{

		timeout: timeout,
	}
}

// NewPostExternalcontactsBulkContactsRemoveParamsWithContext creates a new PostExternalcontactsBulkContactsRemoveParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostExternalcontactsBulkContactsRemoveParamsWithContext(ctx context.Context) *PostExternalcontactsBulkContactsRemoveParams {
	var ()
	return &PostExternalcontactsBulkContactsRemoveParams{

		Context: ctx,
	}
}

// NewPostExternalcontactsBulkContactsRemoveParamsWithHTTPClient creates a new PostExternalcontactsBulkContactsRemoveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostExternalcontactsBulkContactsRemoveParamsWithHTTPClient(client *http.Client) *PostExternalcontactsBulkContactsRemoveParams {
	var ()
	return &PostExternalcontactsBulkContactsRemoveParams{
		HTTPClient: client,
	}
}

/*PostExternalcontactsBulkContactsRemoveParams contains all the parameters to send to the API endpoint
for the post externalcontacts bulk contacts remove operation typically these are written to a http.Request
*/
type PostExternalcontactsBulkContactsRemoveParams struct {

	/*Body
	  Contact ids

	*/
	Body *models.BulkIdsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post externalcontacts bulk contacts remove params
func (o *PostExternalcontactsBulkContactsRemoveParams) WithTimeout(timeout time.Duration) *PostExternalcontactsBulkContactsRemoveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post externalcontacts bulk contacts remove params
func (o *PostExternalcontactsBulkContactsRemoveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post externalcontacts bulk contacts remove params
func (o *PostExternalcontactsBulkContactsRemoveParams) WithContext(ctx context.Context) *PostExternalcontactsBulkContactsRemoveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post externalcontacts bulk contacts remove params
func (o *PostExternalcontactsBulkContactsRemoveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post externalcontacts bulk contacts remove params
func (o *PostExternalcontactsBulkContactsRemoveParams) WithHTTPClient(client *http.Client) *PostExternalcontactsBulkContactsRemoveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post externalcontacts bulk contacts remove params
func (o *PostExternalcontactsBulkContactsRemoveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post externalcontacts bulk contacts remove params
func (o *PostExternalcontactsBulkContactsRemoveParams) WithBody(body *models.BulkIdsRequest) *PostExternalcontactsBulkContactsRemoveParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post externalcontacts bulk contacts remove params
func (o *PostExternalcontactsBulkContactsRemoveParams) SetBody(body *models.BulkIdsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostExternalcontactsBulkContactsRemoveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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