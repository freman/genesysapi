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
)

// NewGetExternalcontactsRelationshipParams creates a new GetExternalcontactsRelationshipParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetExternalcontactsRelationshipParams() *GetExternalcontactsRelationshipParams {
	return &GetExternalcontactsRelationshipParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetExternalcontactsRelationshipParamsWithTimeout creates a new GetExternalcontactsRelationshipParams object
// with the ability to set a timeout on a request.
func NewGetExternalcontactsRelationshipParamsWithTimeout(timeout time.Duration) *GetExternalcontactsRelationshipParams {
	return &GetExternalcontactsRelationshipParams{
		timeout: timeout,
	}
}

// NewGetExternalcontactsRelationshipParamsWithContext creates a new GetExternalcontactsRelationshipParams object
// with the ability to set a context for a request.
func NewGetExternalcontactsRelationshipParamsWithContext(ctx context.Context) *GetExternalcontactsRelationshipParams {
	return &GetExternalcontactsRelationshipParams{
		Context: ctx,
	}
}

// NewGetExternalcontactsRelationshipParamsWithHTTPClient creates a new GetExternalcontactsRelationshipParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetExternalcontactsRelationshipParamsWithHTTPClient(client *http.Client) *GetExternalcontactsRelationshipParams {
	return &GetExternalcontactsRelationshipParams{
		HTTPClient: client,
	}
}

/*
GetExternalcontactsRelationshipParams contains all the parameters to send to the API endpoint

	for the get externalcontacts relationship operation.

	Typically these are written to a http.Request.
*/
type GetExternalcontactsRelationshipParams struct {

	/* Expand.

	   which fields, if any, to expand
	*/
	Expand *string

	/* RelationshipID.

	   Relationship Id
	*/
	RelationshipID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get externalcontacts relationship params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExternalcontactsRelationshipParams) WithDefaults() *GetExternalcontactsRelationshipParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get externalcontacts relationship params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExternalcontactsRelationshipParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get externalcontacts relationship params
func (o *GetExternalcontactsRelationshipParams) WithTimeout(timeout time.Duration) *GetExternalcontactsRelationshipParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get externalcontacts relationship params
func (o *GetExternalcontactsRelationshipParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get externalcontacts relationship params
func (o *GetExternalcontactsRelationshipParams) WithContext(ctx context.Context) *GetExternalcontactsRelationshipParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get externalcontacts relationship params
func (o *GetExternalcontactsRelationshipParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get externalcontacts relationship params
func (o *GetExternalcontactsRelationshipParams) WithHTTPClient(client *http.Client) *GetExternalcontactsRelationshipParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get externalcontacts relationship params
func (o *GetExternalcontactsRelationshipParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get externalcontacts relationship params
func (o *GetExternalcontactsRelationshipParams) WithExpand(expand *string) *GetExternalcontactsRelationshipParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get externalcontacts relationship params
func (o *GetExternalcontactsRelationshipParams) SetExpand(expand *string) {
	o.Expand = expand
}

// WithRelationshipID adds the relationshipID to the get externalcontacts relationship params
func (o *GetExternalcontactsRelationshipParams) WithRelationshipID(relationshipID string) *GetExternalcontactsRelationshipParams {
	o.SetRelationshipID(relationshipID)
	return o
}

// SetRelationshipID adds the relationshipId to the get externalcontacts relationship params
func (o *GetExternalcontactsRelationshipParams) SetRelationshipID(relationshipID string) {
	o.RelationshipID = relationshipID
}

// WriteToRequest writes these params to a swagger request
func (o *GetExternalcontactsRelationshipParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Expand != nil {

		// query param expand
		var qrExpand string

		if o.Expand != nil {
			qrExpand = *o.Expand
		}
		qExpand := qrExpand
		if qExpand != "" {

			if err := r.SetQueryParam("expand", qExpand); err != nil {
				return err
			}
		}
	}

	// path param relationshipId
	if err := r.SetPathParam("relationshipId", o.RelationshipID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
