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

// NewPostExternalcontactsBulkNotesUpdateParams creates a new PostExternalcontactsBulkNotesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostExternalcontactsBulkNotesUpdateParams() *PostExternalcontactsBulkNotesUpdateParams {
	return &PostExternalcontactsBulkNotesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostExternalcontactsBulkNotesUpdateParamsWithTimeout creates a new PostExternalcontactsBulkNotesUpdateParams object
// with the ability to set a timeout on a request.
func NewPostExternalcontactsBulkNotesUpdateParamsWithTimeout(timeout time.Duration) *PostExternalcontactsBulkNotesUpdateParams {
	return &PostExternalcontactsBulkNotesUpdateParams{
		timeout: timeout,
	}
}

// NewPostExternalcontactsBulkNotesUpdateParamsWithContext creates a new PostExternalcontactsBulkNotesUpdateParams object
// with the ability to set a context for a request.
func NewPostExternalcontactsBulkNotesUpdateParamsWithContext(ctx context.Context) *PostExternalcontactsBulkNotesUpdateParams {
	return &PostExternalcontactsBulkNotesUpdateParams{
		Context: ctx,
	}
}

// NewPostExternalcontactsBulkNotesUpdateParamsWithHTTPClient creates a new PostExternalcontactsBulkNotesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostExternalcontactsBulkNotesUpdateParamsWithHTTPClient(client *http.Client) *PostExternalcontactsBulkNotesUpdateParams {
	return &PostExternalcontactsBulkNotesUpdateParams{
		HTTPClient: client,
	}
}

/*
PostExternalcontactsBulkNotesUpdateParams contains all the parameters to send to the API endpoint

	for the post externalcontacts bulk notes update operation.

	Typically these are written to a http.Request.
*/
type PostExternalcontactsBulkNotesUpdateParams struct {

	/* Body.

	   Notes
	*/
	Body *models.BulkNotesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post externalcontacts bulk notes update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostExternalcontactsBulkNotesUpdateParams) WithDefaults() *PostExternalcontactsBulkNotesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post externalcontacts bulk notes update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostExternalcontactsBulkNotesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post externalcontacts bulk notes update params
func (o *PostExternalcontactsBulkNotesUpdateParams) WithTimeout(timeout time.Duration) *PostExternalcontactsBulkNotesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post externalcontacts bulk notes update params
func (o *PostExternalcontactsBulkNotesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post externalcontacts bulk notes update params
func (o *PostExternalcontactsBulkNotesUpdateParams) WithContext(ctx context.Context) *PostExternalcontactsBulkNotesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post externalcontacts bulk notes update params
func (o *PostExternalcontactsBulkNotesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post externalcontacts bulk notes update params
func (o *PostExternalcontactsBulkNotesUpdateParams) WithHTTPClient(client *http.Client) *PostExternalcontactsBulkNotesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post externalcontacts bulk notes update params
func (o *PostExternalcontactsBulkNotesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post externalcontacts bulk notes update params
func (o *PostExternalcontactsBulkNotesUpdateParams) WithBody(body *models.BulkNotesRequest) *PostExternalcontactsBulkNotesUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post externalcontacts bulk notes update params
func (o *PostExternalcontactsBulkNotesUpdateParams) SetBody(body *models.BulkNotesRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostExternalcontactsBulkNotesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
