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
	"github.com/go-openapi/swag"
)

// NewGetKnowledgeGuestSessionDocumentsParams creates a new GetKnowledgeGuestSessionDocumentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetKnowledgeGuestSessionDocumentsParams() *GetKnowledgeGuestSessionDocumentsParams {
	return &GetKnowledgeGuestSessionDocumentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetKnowledgeGuestSessionDocumentsParamsWithTimeout creates a new GetKnowledgeGuestSessionDocumentsParams object
// with the ability to set a timeout on a request.
func NewGetKnowledgeGuestSessionDocumentsParamsWithTimeout(timeout time.Duration) *GetKnowledgeGuestSessionDocumentsParams {
	return &GetKnowledgeGuestSessionDocumentsParams{
		timeout: timeout,
	}
}

// NewGetKnowledgeGuestSessionDocumentsParamsWithContext creates a new GetKnowledgeGuestSessionDocumentsParams object
// with the ability to set a context for a request.
func NewGetKnowledgeGuestSessionDocumentsParamsWithContext(ctx context.Context) *GetKnowledgeGuestSessionDocumentsParams {
	return &GetKnowledgeGuestSessionDocumentsParams{
		Context: ctx,
	}
}

// NewGetKnowledgeGuestSessionDocumentsParamsWithHTTPClient creates a new GetKnowledgeGuestSessionDocumentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetKnowledgeGuestSessionDocumentsParamsWithHTTPClient(client *http.Client) *GetKnowledgeGuestSessionDocumentsParams {
	return &GetKnowledgeGuestSessionDocumentsParams{
		HTTPClient: client,
	}
}

/*
GetKnowledgeGuestSessionDocumentsParams contains all the parameters to send to the API endpoint

	for the get knowledge guest session documents operation.

	Typically these are written to a http.Request.
*/
type GetKnowledgeGuestSessionDocumentsParams struct {

	/* CategoryID.

	   If specified, retrieves documents associated with category ids, comma separated values expected.
	*/
	CategoryID []string

	/* PageSize.

	   Number of entities to return. Maximum of 200.

	   Format: int32
	*/
	PageSize *int32

	/* SessionID.

	   Knowledge guest session ID.
	*/
	SessionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get knowledge guest session documents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKnowledgeGuestSessionDocumentsParams) WithDefaults() *GetKnowledgeGuestSessionDocumentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get knowledge guest session documents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKnowledgeGuestSessionDocumentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get knowledge guest session documents params
func (o *GetKnowledgeGuestSessionDocumentsParams) WithTimeout(timeout time.Duration) *GetKnowledgeGuestSessionDocumentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get knowledge guest session documents params
func (o *GetKnowledgeGuestSessionDocumentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get knowledge guest session documents params
func (o *GetKnowledgeGuestSessionDocumentsParams) WithContext(ctx context.Context) *GetKnowledgeGuestSessionDocumentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get knowledge guest session documents params
func (o *GetKnowledgeGuestSessionDocumentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get knowledge guest session documents params
func (o *GetKnowledgeGuestSessionDocumentsParams) WithHTTPClient(client *http.Client) *GetKnowledgeGuestSessionDocumentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get knowledge guest session documents params
func (o *GetKnowledgeGuestSessionDocumentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategoryID adds the categoryID to the get knowledge guest session documents params
func (o *GetKnowledgeGuestSessionDocumentsParams) WithCategoryID(categoryID []string) *GetKnowledgeGuestSessionDocumentsParams {
	o.SetCategoryID(categoryID)
	return o
}

// SetCategoryID adds the categoryId to the get knowledge guest session documents params
func (o *GetKnowledgeGuestSessionDocumentsParams) SetCategoryID(categoryID []string) {
	o.CategoryID = categoryID
}

// WithPageSize adds the pageSize to the get knowledge guest session documents params
func (o *GetKnowledgeGuestSessionDocumentsParams) WithPageSize(pageSize *int32) *GetKnowledgeGuestSessionDocumentsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get knowledge guest session documents params
func (o *GetKnowledgeGuestSessionDocumentsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSessionID adds the sessionID to the get knowledge guest session documents params
func (o *GetKnowledgeGuestSessionDocumentsParams) WithSessionID(sessionID string) *GetKnowledgeGuestSessionDocumentsParams {
	o.SetSessionID(sessionID)
	return o
}

// SetSessionID adds the sessionId to the get knowledge guest session documents params
func (o *GetKnowledgeGuestSessionDocumentsParams) SetSessionID(sessionID string) {
	o.SessionID = sessionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetKnowledgeGuestSessionDocumentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CategoryID != nil {

		// binding items for categoryId
		joinedCategoryID := o.bindParamCategoryID(reg)

		// query array param categoryId
		if err := r.SetQueryParam("categoryId", joinedCategoryID...); err != nil {
			return err
		}
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int32

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}
	}

	// path param sessionId
	if err := r.SetPathParam("sessionId", o.SessionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetKnowledgeGuestSessionDocuments binds the parameter categoryId
func (o *GetKnowledgeGuestSessionDocumentsParams) bindParamCategoryID(formats strfmt.Registry) []string {
	categoryIDIR := o.CategoryID

	var categoryIDIC []string
	for _, categoryIDIIR := range categoryIDIR { // explode []string

		categoryIDIIV := categoryIDIIR // string as string
		categoryIDIC = append(categoryIDIC, categoryIDIIV)
	}

	// items.CollectionFormat: "multi"
	categoryIDIS := swag.JoinByFormat(categoryIDIC, "multi")

	return categoryIDIS
}
