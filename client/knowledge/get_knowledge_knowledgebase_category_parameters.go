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

// NewGetKnowledgeKnowledgebaseCategoryParams creates a new GetKnowledgeKnowledgebaseCategoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetKnowledgeKnowledgebaseCategoryParams() *GetKnowledgeKnowledgebaseCategoryParams {
	return &GetKnowledgeKnowledgebaseCategoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetKnowledgeKnowledgebaseCategoryParamsWithTimeout creates a new GetKnowledgeKnowledgebaseCategoryParams object
// with the ability to set a timeout on a request.
func NewGetKnowledgeKnowledgebaseCategoryParamsWithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseCategoryParams {
	return &GetKnowledgeKnowledgebaseCategoryParams{
		timeout: timeout,
	}
}

// NewGetKnowledgeKnowledgebaseCategoryParamsWithContext creates a new GetKnowledgeKnowledgebaseCategoryParams object
// with the ability to set a context for a request.
func NewGetKnowledgeKnowledgebaseCategoryParamsWithContext(ctx context.Context) *GetKnowledgeKnowledgebaseCategoryParams {
	return &GetKnowledgeKnowledgebaseCategoryParams{
		Context: ctx,
	}
}

// NewGetKnowledgeKnowledgebaseCategoryParamsWithHTTPClient creates a new GetKnowledgeKnowledgebaseCategoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetKnowledgeKnowledgebaseCategoryParamsWithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseCategoryParams {
	return &GetKnowledgeKnowledgebaseCategoryParams{
		HTTPClient: client,
	}
}

/*
GetKnowledgeKnowledgebaseCategoryParams contains all the parameters to send to the API endpoint

	for the get knowledge knowledgebase category operation.

	Typically these are written to a http.Request.
*/
type GetKnowledgeKnowledgebaseCategoryParams struct {

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

// WithDefaults hydrates default values in the get knowledge knowledgebase category params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKnowledgeKnowledgebaseCategoryParams) WithDefaults() *GetKnowledgeKnowledgebaseCategoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get knowledge knowledgebase category params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKnowledgeKnowledgebaseCategoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get knowledge knowledgebase category params
func (o *GetKnowledgeKnowledgebaseCategoryParams) WithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseCategoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get knowledge knowledgebase category params
func (o *GetKnowledgeKnowledgebaseCategoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get knowledge knowledgebase category params
func (o *GetKnowledgeKnowledgebaseCategoryParams) WithContext(ctx context.Context) *GetKnowledgeKnowledgebaseCategoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get knowledge knowledgebase category params
func (o *GetKnowledgeKnowledgebaseCategoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get knowledge knowledgebase category params
func (o *GetKnowledgeKnowledgebaseCategoryParams) WithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseCategoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get knowledge knowledgebase category params
func (o *GetKnowledgeKnowledgebaseCategoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategoryID adds the categoryID to the get knowledge knowledgebase category params
func (o *GetKnowledgeKnowledgebaseCategoryParams) WithCategoryID(categoryID string) *GetKnowledgeKnowledgebaseCategoryParams {
	o.SetCategoryID(categoryID)
	return o
}

// SetCategoryID adds the categoryId to the get knowledge knowledgebase category params
func (o *GetKnowledgeKnowledgebaseCategoryParams) SetCategoryID(categoryID string) {
	o.CategoryID = categoryID
}

// WithKnowledgeBaseID adds the knowledgeBaseID to the get knowledge knowledgebase category params
func (o *GetKnowledgeKnowledgebaseCategoryParams) WithKnowledgeBaseID(knowledgeBaseID string) *GetKnowledgeKnowledgebaseCategoryParams {
	o.SetKnowledgeBaseID(knowledgeBaseID)
	return o
}

// SetKnowledgeBaseID adds the knowledgeBaseId to the get knowledge knowledgebase category params
func (o *GetKnowledgeKnowledgebaseCategoryParams) SetKnowledgeBaseID(knowledgeBaseID string) {
	o.KnowledgeBaseID = knowledgeBaseID
}

// WriteToRequest writes these params to a swagger request
func (o *GetKnowledgeKnowledgebaseCategoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
