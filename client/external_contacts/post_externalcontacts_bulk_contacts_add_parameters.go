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

// NewPostExternalcontactsBulkContactsAddParams creates a new PostExternalcontactsBulkContactsAddParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostExternalcontactsBulkContactsAddParams() *PostExternalcontactsBulkContactsAddParams {
	return &PostExternalcontactsBulkContactsAddParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostExternalcontactsBulkContactsAddParamsWithTimeout creates a new PostExternalcontactsBulkContactsAddParams object
// with the ability to set a timeout on a request.
func NewPostExternalcontactsBulkContactsAddParamsWithTimeout(timeout time.Duration) *PostExternalcontactsBulkContactsAddParams {
	return &PostExternalcontactsBulkContactsAddParams{
		timeout: timeout,
	}
}

// NewPostExternalcontactsBulkContactsAddParamsWithContext creates a new PostExternalcontactsBulkContactsAddParams object
// with the ability to set a context for a request.
func NewPostExternalcontactsBulkContactsAddParamsWithContext(ctx context.Context) *PostExternalcontactsBulkContactsAddParams {
	return &PostExternalcontactsBulkContactsAddParams{
		Context: ctx,
	}
}

// NewPostExternalcontactsBulkContactsAddParamsWithHTTPClient creates a new PostExternalcontactsBulkContactsAddParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostExternalcontactsBulkContactsAddParamsWithHTTPClient(client *http.Client) *PostExternalcontactsBulkContactsAddParams {
	return &PostExternalcontactsBulkContactsAddParams{
		HTTPClient: client,
	}
}

/*
PostExternalcontactsBulkContactsAddParams contains all the parameters to send to the API endpoint

	for the post externalcontacts bulk contacts add operation.

	Typically these are written to a http.Request.
*/
type PostExternalcontactsBulkContactsAddParams struct {

	/* Body.

	   Contacts
	*/
	Body *models.BulkContactsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post externalcontacts bulk contacts add params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostExternalcontactsBulkContactsAddParams) WithDefaults() *PostExternalcontactsBulkContactsAddParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post externalcontacts bulk contacts add params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostExternalcontactsBulkContactsAddParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post externalcontacts bulk contacts add params
func (o *PostExternalcontactsBulkContactsAddParams) WithTimeout(timeout time.Duration) *PostExternalcontactsBulkContactsAddParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post externalcontacts bulk contacts add params
func (o *PostExternalcontactsBulkContactsAddParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post externalcontacts bulk contacts add params
func (o *PostExternalcontactsBulkContactsAddParams) WithContext(ctx context.Context) *PostExternalcontactsBulkContactsAddParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post externalcontacts bulk contacts add params
func (o *PostExternalcontactsBulkContactsAddParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post externalcontacts bulk contacts add params
func (o *PostExternalcontactsBulkContactsAddParams) WithHTTPClient(client *http.Client) *PostExternalcontactsBulkContactsAddParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post externalcontacts bulk contacts add params
func (o *PostExternalcontactsBulkContactsAddParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post externalcontacts bulk contacts add params
func (o *PostExternalcontactsBulkContactsAddParams) WithBody(body *models.BulkContactsRequest) *PostExternalcontactsBulkContactsAddParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post externalcontacts bulk contacts add params
func (o *PostExternalcontactsBulkContactsAddParams) SetBody(body *models.BulkContactsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostExternalcontactsBulkContactsAddParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
