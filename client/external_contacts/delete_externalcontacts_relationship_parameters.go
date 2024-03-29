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

// NewDeleteExternalcontactsRelationshipParams creates a new DeleteExternalcontactsRelationshipParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteExternalcontactsRelationshipParams() *DeleteExternalcontactsRelationshipParams {
	return &DeleteExternalcontactsRelationshipParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteExternalcontactsRelationshipParamsWithTimeout creates a new DeleteExternalcontactsRelationshipParams object
// with the ability to set a timeout on a request.
func NewDeleteExternalcontactsRelationshipParamsWithTimeout(timeout time.Duration) *DeleteExternalcontactsRelationshipParams {
	return &DeleteExternalcontactsRelationshipParams{
		timeout: timeout,
	}
}

// NewDeleteExternalcontactsRelationshipParamsWithContext creates a new DeleteExternalcontactsRelationshipParams object
// with the ability to set a context for a request.
func NewDeleteExternalcontactsRelationshipParamsWithContext(ctx context.Context) *DeleteExternalcontactsRelationshipParams {
	return &DeleteExternalcontactsRelationshipParams{
		Context: ctx,
	}
}

// NewDeleteExternalcontactsRelationshipParamsWithHTTPClient creates a new DeleteExternalcontactsRelationshipParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteExternalcontactsRelationshipParamsWithHTTPClient(client *http.Client) *DeleteExternalcontactsRelationshipParams {
	return &DeleteExternalcontactsRelationshipParams{
		HTTPClient: client,
	}
}

/*
DeleteExternalcontactsRelationshipParams contains all the parameters to send to the API endpoint

	for the delete externalcontacts relationship operation.

	Typically these are written to a http.Request.
*/
type DeleteExternalcontactsRelationshipParams struct {

	/* RelationshipID.

	   Relationship Id
	*/
	RelationshipID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete externalcontacts relationship params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteExternalcontactsRelationshipParams) WithDefaults() *DeleteExternalcontactsRelationshipParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete externalcontacts relationship params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteExternalcontactsRelationshipParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete externalcontacts relationship params
func (o *DeleteExternalcontactsRelationshipParams) WithTimeout(timeout time.Duration) *DeleteExternalcontactsRelationshipParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete externalcontacts relationship params
func (o *DeleteExternalcontactsRelationshipParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete externalcontacts relationship params
func (o *DeleteExternalcontactsRelationshipParams) WithContext(ctx context.Context) *DeleteExternalcontactsRelationshipParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete externalcontacts relationship params
func (o *DeleteExternalcontactsRelationshipParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete externalcontacts relationship params
func (o *DeleteExternalcontactsRelationshipParams) WithHTTPClient(client *http.Client) *DeleteExternalcontactsRelationshipParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete externalcontacts relationship params
func (o *DeleteExternalcontactsRelationshipParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRelationshipID adds the relationshipID to the delete externalcontacts relationship params
func (o *DeleteExternalcontactsRelationshipParams) WithRelationshipID(relationshipID string) *DeleteExternalcontactsRelationshipParams {
	o.SetRelationshipID(relationshipID)
	return o
}

// SetRelationshipID adds the relationshipId to the delete externalcontacts relationship params
func (o *DeleteExternalcontactsRelationshipParams) SetRelationshipID(relationshipID string) {
	o.RelationshipID = relationshipID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteExternalcontactsRelationshipParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param relationshipId
	if err := r.SetPathParam("relationshipId", o.RelationshipID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
