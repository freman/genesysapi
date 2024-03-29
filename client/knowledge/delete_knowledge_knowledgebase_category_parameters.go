// Code generated by go-swagger; DO NOT EDIT.

package knowledge

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

// NewDeleteKnowledgeKnowledgebaseCategoryParams creates a new DeleteKnowledgeKnowledgebaseCategoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteKnowledgeKnowledgebaseCategoryParams() *DeleteKnowledgeKnowledgebaseCategoryParams {
	return &DeleteKnowledgeKnowledgebaseCategoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteKnowledgeKnowledgebaseCategoryParamsWithTimeout creates a new DeleteKnowledgeKnowledgebaseCategoryParams object
// with the ability to set a timeout on a request.
func NewDeleteKnowledgeKnowledgebaseCategoryParamsWithTimeout(timeout time.Duration) *DeleteKnowledgeKnowledgebaseCategoryParams {
	return &DeleteKnowledgeKnowledgebaseCategoryParams{
		timeout: timeout,
	}
}

// NewDeleteKnowledgeKnowledgebaseCategoryParamsWithContext creates a new DeleteKnowledgeKnowledgebaseCategoryParams object
// with the ability to set a context for a request.
func NewDeleteKnowledgeKnowledgebaseCategoryParamsWithContext(ctx context.Context) *DeleteKnowledgeKnowledgebaseCategoryParams {
	return &DeleteKnowledgeKnowledgebaseCategoryParams{
		Context: ctx,
	}
}

// NewDeleteKnowledgeKnowledgebaseCategoryParamsWithHTTPClient creates a new DeleteKnowledgeKnowledgebaseCategoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteKnowledgeKnowledgebaseCategoryParamsWithHTTPClient(client *http.Client) *DeleteKnowledgeKnowledgebaseCategoryParams {
	return &DeleteKnowledgeKnowledgebaseCategoryParams{
		HTTPClient: client,
	}
}

/*
DeleteKnowledgeKnowledgebaseCategoryParams contains all the parameters to send to the API endpoint

	for the delete knowledge knowledgebase category operation.

	Typically these are written to a http.Request.
*/
type DeleteKnowledgeKnowledgebaseCategoryParams struct {

	/* CategoryID.

	   Category ID
	*/
	CategoryID string

	/* KnowledgeBaseID.

	   Knowledge base ID
	*/
	KnowledgeBaseID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete knowledge knowledgebase category params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteKnowledgeKnowledgebaseCategoryParams) WithDefaults() *DeleteKnowledgeKnowledgebaseCategoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete knowledge knowledgebase category params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteKnowledgeKnowledgebaseCategoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete knowledge knowledgebase category params
func (o *DeleteKnowledgeKnowledgebaseCategoryParams) WithTimeout(timeout time.Duration) *DeleteKnowledgeKnowledgebaseCategoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete knowledge knowledgebase category params
func (o *DeleteKnowledgeKnowledgebaseCategoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete knowledge knowledgebase category params
func (o *DeleteKnowledgeKnowledgebaseCategoryParams) WithContext(ctx context.Context) *DeleteKnowledgeKnowledgebaseCategoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete knowledge knowledgebase category params
func (o *DeleteKnowledgeKnowledgebaseCategoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete knowledge knowledgebase category params
func (o *DeleteKnowledgeKnowledgebaseCategoryParams) WithHTTPClient(client *http.Client) *DeleteKnowledgeKnowledgebaseCategoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete knowledge knowledgebase category params
func (o *DeleteKnowledgeKnowledgebaseCategoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategoryID adds the categoryID to the delete knowledge knowledgebase category params
func (o *DeleteKnowledgeKnowledgebaseCategoryParams) WithCategoryID(categoryID string) *DeleteKnowledgeKnowledgebaseCategoryParams {
	o.SetCategoryID(categoryID)
	return o
}

// SetCategoryID adds the categoryId to the delete knowledge knowledgebase category params
func (o *DeleteKnowledgeKnowledgebaseCategoryParams) SetCategoryID(categoryID string) {
	o.CategoryID = categoryID
}

// WithKnowledgeBaseID adds the knowledgeBaseID to the delete knowledge knowledgebase category params
func (o *DeleteKnowledgeKnowledgebaseCategoryParams) WithKnowledgeBaseID(knowledgeBaseID string) *DeleteKnowledgeKnowledgebaseCategoryParams {
	o.SetKnowledgeBaseID(knowledgeBaseID)
	return o
}

// SetKnowledgeBaseID adds the knowledgeBaseId to the delete knowledge knowledgebase category params
func (o *DeleteKnowledgeKnowledgebaseCategoryParams) SetKnowledgeBaseID(knowledgeBaseID string) {
	o.KnowledgeBaseID = knowledgeBaseID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteKnowledgeKnowledgebaseCategoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param categoryId
	if err := r.SetPathParam("categoryId", o.CategoryID); err != nil {
		return err
	}

	// path param knowledgeBaseId
	if err := r.SetPathParam("knowledgeBaseId", o.KnowledgeBaseID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
