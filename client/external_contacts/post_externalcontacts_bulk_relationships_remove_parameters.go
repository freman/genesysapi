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

// NewPostExternalcontactsBulkRelationshipsRemoveParams creates a new PostExternalcontactsBulkRelationshipsRemoveParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostExternalcontactsBulkRelationshipsRemoveParams() *PostExternalcontactsBulkRelationshipsRemoveParams {
	return &PostExternalcontactsBulkRelationshipsRemoveParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostExternalcontactsBulkRelationshipsRemoveParamsWithTimeout creates a new PostExternalcontactsBulkRelationshipsRemoveParams object
// with the ability to set a timeout on a request.
func NewPostExternalcontactsBulkRelationshipsRemoveParamsWithTimeout(timeout time.Duration) *PostExternalcontactsBulkRelationshipsRemoveParams {
	return &PostExternalcontactsBulkRelationshipsRemoveParams{
		timeout: timeout,
	}
}

// NewPostExternalcontactsBulkRelationshipsRemoveParamsWithContext creates a new PostExternalcontactsBulkRelationshipsRemoveParams object
// with the ability to set a context for a request.
func NewPostExternalcontactsBulkRelationshipsRemoveParamsWithContext(ctx context.Context) *PostExternalcontactsBulkRelationshipsRemoveParams {
	return &PostExternalcontactsBulkRelationshipsRemoveParams{
		Context: ctx,
	}
}

// NewPostExternalcontactsBulkRelationshipsRemoveParamsWithHTTPClient creates a new PostExternalcontactsBulkRelationshipsRemoveParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostExternalcontactsBulkRelationshipsRemoveParamsWithHTTPClient(client *http.Client) *PostExternalcontactsBulkRelationshipsRemoveParams {
	return &PostExternalcontactsBulkRelationshipsRemoveParams{
		HTTPClient: client,
	}
}

/*
PostExternalcontactsBulkRelationshipsRemoveParams contains all the parameters to send to the API endpoint

	for the post externalcontacts bulk relationships remove operation.

	Typically these are written to a http.Request.
*/
type PostExternalcontactsBulkRelationshipsRemoveParams struct {

	/* Body.

	   Relationships ids
	*/
	Body *models.BulkIdsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post externalcontacts bulk relationships remove params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostExternalcontactsBulkRelationshipsRemoveParams) WithDefaults() *PostExternalcontactsBulkRelationshipsRemoveParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post externalcontacts bulk relationships remove params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostExternalcontactsBulkRelationshipsRemoveParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post externalcontacts bulk relationships remove params
func (o *PostExternalcontactsBulkRelationshipsRemoveParams) WithTimeout(timeout time.Duration) *PostExternalcontactsBulkRelationshipsRemoveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post externalcontacts bulk relationships remove params
func (o *PostExternalcontactsBulkRelationshipsRemoveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post externalcontacts bulk relationships remove params
func (o *PostExternalcontactsBulkRelationshipsRemoveParams) WithContext(ctx context.Context) *PostExternalcontactsBulkRelationshipsRemoveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post externalcontacts bulk relationships remove params
func (o *PostExternalcontactsBulkRelationshipsRemoveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post externalcontacts bulk relationships remove params
func (o *PostExternalcontactsBulkRelationshipsRemoveParams) WithHTTPClient(client *http.Client) *PostExternalcontactsBulkRelationshipsRemoveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post externalcontacts bulk relationships remove params
func (o *PostExternalcontactsBulkRelationshipsRemoveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post externalcontacts bulk relationships remove params
func (o *PostExternalcontactsBulkRelationshipsRemoveParams) WithBody(body *models.BulkIdsRequest) *PostExternalcontactsBulkRelationshipsRemoveParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post externalcontacts bulk relationships remove params
func (o *PostExternalcontactsBulkRelationshipsRemoveParams) SetBody(body *models.BulkIdsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostExternalcontactsBulkRelationshipsRemoveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
