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

// NewPostExternalcontactsBulkRelationshipsAddParams creates a new PostExternalcontactsBulkRelationshipsAddParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostExternalcontactsBulkRelationshipsAddParams() *PostExternalcontactsBulkRelationshipsAddParams {
	return &PostExternalcontactsBulkRelationshipsAddParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostExternalcontactsBulkRelationshipsAddParamsWithTimeout creates a new PostExternalcontactsBulkRelationshipsAddParams object
// with the ability to set a timeout on a request.
func NewPostExternalcontactsBulkRelationshipsAddParamsWithTimeout(timeout time.Duration) *PostExternalcontactsBulkRelationshipsAddParams {
	return &PostExternalcontactsBulkRelationshipsAddParams{
		timeout: timeout,
	}
}

// NewPostExternalcontactsBulkRelationshipsAddParamsWithContext creates a new PostExternalcontactsBulkRelationshipsAddParams object
// with the ability to set a context for a request.
func NewPostExternalcontactsBulkRelationshipsAddParamsWithContext(ctx context.Context) *PostExternalcontactsBulkRelationshipsAddParams {
	return &PostExternalcontactsBulkRelationshipsAddParams{
		Context: ctx,
	}
}

// NewPostExternalcontactsBulkRelationshipsAddParamsWithHTTPClient creates a new PostExternalcontactsBulkRelationshipsAddParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostExternalcontactsBulkRelationshipsAddParamsWithHTTPClient(client *http.Client) *PostExternalcontactsBulkRelationshipsAddParams {
	return &PostExternalcontactsBulkRelationshipsAddParams{
		HTTPClient: client,
	}
}

/*
PostExternalcontactsBulkRelationshipsAddParams contains all the parameters to send to the API endpoint

	for the post externalcontacts bulk relationships add operation.

	Typically these are written to a http.Request.
*/
type PostExternalcontactsBulkRelationshipsAddParams struct {

	/* Body.

	   Relationships
	*/
	Body *models.BulkRelationshipsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post externalcontacts bulk relationships add params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostExternalcontactsBulkRelationshipsAddParams) WithDefaults() *PostExternalcontactsBulkRelationshipsAddParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post externalcontacts bulk relationships add params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostExternalcontactsBulkRelationshipsAddParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post externalcontacts bulk relationships add params
func (o *PostExternalcontactsBulkRelationshipsAddParams) WithTimeout(timeout time.Duration) *PostExternalcontactsBulkRelationshipsAddParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post externalcontacts bulk relationships add params
func (o *PostExternalcontactsBulkRelationshipsAddParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post externalcontacts bulk relationships add params
func (o *PostExternalcontactsBulkRelationshipsAddParams) WithContext(ctx context.Context) *PostExternalcontactsBulkRelationshipsAddParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post externalcontacts bulk relationships add params
func (o *PostExternalcontactsBulkRelationshipsAddParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post externalcontacts bulk relationships add params
func (o *PostExternalcontactsBulkRelationshipsAddParams) WithHTTPClient(client *http.Client) *PostExternalcontactsBulkRelationshipsAddParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post externalcontacts bulk relationships add params
func (o *PostExternalcontactsBulkRelationshipsAddParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post externalcontacts bulk relationships add params
func (o *PostExternalcontactsBulkRelationshipsAddParams) WithBody(body *models.BulkRelationshipsRequest) *PostExternalcontactsBulkRelationshipsAddParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post externalcontacts bulk relationships add params
func (o *PostExternalcontactsBulkRelationshipsAddParams) SetBody(body *models.BulkRelationshipsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostExternalcontactsBulkRelationshipsAddParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
