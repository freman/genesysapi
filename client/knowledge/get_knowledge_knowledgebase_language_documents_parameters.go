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

// NewGetKnowledgeKnowledgebaseLanguageDocumentsParams creates a new GetKnowledgeKnowledgebaseLanguageDocumentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetKnowledgeKnowledgebaseLanguageDocumentsParams() *GetKnowledgeKnowledgebaseLanguageDocumentsParams {
	return &GetKnowledgeKnowledgebaseLanguageDocumentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetKnowledgeKnowledgebaseLanguageDocumentsParamsWithTimeout creates a new GetKnowledgeKnowledgebaseLanguageDocumentsParams object
// with the ability to set a timeout on a request.
func NewGetKnowledgeKnowledgebaseLanguageDocumentsParamsWithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseLanguageDocumentsParams {
	return &GetKnowledgeKnowledgebaseLanguageDocumentsParams{
		timeout: timeout,
	}
}

// NewGetKnowledgeKnowledgebaseLanguageDocumentsParamsWithContext creates a new GetKnowledgeKnowledgebaseLanguageDocumentsParams object
// with the ability to set a context for a request.
func NewGetKnowledgeKnowledgebaseLanguageDocumentsParamsWithContext(ctx context.Context) *GetKnowledgeKnowledgebaseLanguageDocumentsParams {
	return &GetKnowledgeKnowledgebaseLanguageDocumentsParams{
		Context: ctx,
	}
}

// NewGetKnowledgeKnowledgebaseLanguageDocumentsParamsWithHTTPClient creates a new GetKnowledgeKnowledgebaseLanguageDocumentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetKnowledgeKnowledgebaseLanguageDocumentsParamsWithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseLanguageDocumentsParams {
	return &GetKnowledgeKnowledgebaseLanguageDocumentsParams{
		HTTPClient: client,
	}
}

/*
GetKnowledgeKnowledgebaseLanguageDocumentsParams contains all the parameters to send to the API endpoint

	for the get knowledge knowledgebase language documents operation.

	Typically these are written to a http.Request.
*/
type GetKnowledgeKnowledgebaseLanguageDocumentsParams struct {

	/* After.

	   The cursor that points to the end of the set of entities that has been returned.
	*/
	After *string

	/* Before.

	   The cursor that points to the start of the set of entities that has been returned.
	*/
	Before *string

	/* Categories.

	   Filter by categories ids, comma separated values expected.
	*/
	Categories *string

	/* DocumentIds.

	   Comma-separated list of document identifiers to fetch by.
	*/
	DocumentIds []string

	/* KnowledgeBaseID.

	   Knowledge base ID
	*/
	KnowledgeBaseID string

	/* LanguageCode.

	   Language code, format: iso2-LOCALE
	*/
	LanguageCode string

	/* Limit.

	   Number of entities to return. Maximum of 200. Deprecated in favour of pageSize
	*/
	Limit *string

	/* PageSize.

	   Number of entities to return. Maximum of 200.
	*/
	PageSize *string

	/* SortBy.

	   Sort by.
	*/
	SortBy *string

	/* SortOrder.

	   Sort Order.
	*/
	SortOrder *string

	/* Title.

	   Filter by document title.
	*/
	Title *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get knowledge knowledgebase language documents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) WithDefaults() *GetKnowledgeKnowledgebaseLanguageDocumentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get knowledge knowledgebase language documents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) WithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseLanguageDocumentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) WithContext(ctx context.Context) *GetKnowledgeKnowledgebaseLanguageDocumentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) WithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseLanguageDocumentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAfter adds the after to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) WithAfter(after *string) *GetKnowledgeKnowledgebaseLanguageDocumentsParams {
	o.SetAfter(after)
	return o
}

// SetAfter adds the after to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) SetAfter(after *string) {
	o.After = after
}

// WithBefore adds the before to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) WithBefore(before *string) *GetKnowledgeKnowledgebaseLanguageDocumentsParams {
	o.SetBefore(before)
	return o
}

// SetBefore adds the before to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) SetBefore(before *string) {
	o.Before = before
}

// WithCategories adds the categories to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) WithCategories(categories *string) *GetKnowledgeKnowledgebaseLanguageDocumentsParams {
	o.SetCategories(categories)
	return o
}

// SetCategories adds the categories to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) SetCategories(categories *string) {
	o.Categories = categories
}

// WithDocumentIds adds the documentIds to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) WithDocumentIds(documentIds []string) *GetKnowledgeKnowledgebaseLanguageDocumentsParams {
	o.SetDocumentIds(documentIds)
	return o
}

// SetDocumentIds adds the documentIds to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) SetDocumentIds(documentIds []string) {
	o.DocumentIds = documentIds
}

// WithKnowledgeBaseID adds the knowledgeBaseID to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) WithKnowledgeBaseID(knowledgeBaseID string) *GetKnowledgeKnowledgebaseLanguageDocumentsParams {
	o.SetKnowledgeBaseID(knowledgeBaseID)
	return o
}

// SetKnowledgeBaseID adds the knowledgeBaseId to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) SetKnowledgeBaseID(knowledgeBaseID string) {
	o.KnowledgeBaseID = knowledgeBaseID
}

// WithLanguageCode adds the languageCode to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) WithLanguageCode(languageCode string) *GetKnowledgeKnowledgebaseLanguageDocumentsParams {
	o.SetLanguageCode(languageCode)
	return o
}

// SetLanguageCode adds the languageCode to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) SetLanguageCode(languageCode string) {
	o.LanguageCode = languageCode
}

// WithLimit adds the limit to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) WithLimit(limit *string) *GetKnowledgeKnowledgebaseLanguageDocumentsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithPageSize adds the pageSize to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) WithPageSize(pageSize *string) *GetKnowledgeKnowledgebaseLanguageDocumentsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) SetPageSize(pageSize *string) {
	o.PageSize = pageSize
}

// WithSortBy adds the sortBy to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) WithSortBy(sortBy *string) *GetKnowledgeKnowledgebaseLanguageDocumentsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) WithSortOrder(sortOrder *string) *GetKnowledgeKnowledgebaseLanguageDocumentsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WithTitle adds the title to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) WithTitle(title *string) *GetKnowledgeKnowledgebaseLanguageDocumentsParams {
	o.SetTitle(title)
	return o
}

// SetTitle adds the title to the get knowledge knowledgebase language documents params
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) SetTitle(title *string) {
	o.Title = title
}

// WriteToRequest writes these params to a swagger request
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.After != nil {

		// query param after
		var qrAfter string

		if o.After != nil {
			qrAfter = *o.After
		}
		qAfter := qrAfter
		if qAfter != "" {

			if err := r.SetQueryParam("after", qAfter); err != nil {
				return err
			}
		}
	}

	if o.Before != nil {

		// query param before
		var qrBefore string

		if o.Before != nil {
			qrBefore = *o.Before
		}
		qBefore := qrBefore
		if qBefore != "" {

			if err := r.SetQueryParam("before", qBefore); err != nil {
				return err
			}
		}
	}

	if o.Categories != nil {

		// query param categories
		var qrCategories string

		if o.Categories != nil {
			qrCategories = *o.Categories
		}
		qCategories := qrCategories
		if qCategories != "" {

			if err := r.SetQueryParam("categories", qCategories); err != nil {
				return err
			}
		}
	}

	if o.DocumentIds != nil {

		// binding items for documentIds
		joinedDocumentIds := o.bindParamDocumentIds(reg)

		// query array param documentIds
		if err := r.SetQueryParam("documentIds", joinedDocumentIds...); err != nil {
			return err
		}
	}

	// path param knowledgeBaseId
	if err := r.SetPathParam("knowledgeBaseId", o.KnowledgeBaseID); err != nil {
		return err
	}

	// path param languageCode
	if err := r.SetPathParam("languageCode", o.LanguageCode); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit string

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := qrLimit
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize string

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := qrPageSize
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.SortBy != nil {

		// query param sortBy
		var qrSortBy string

		if o.SortBy != nil {
			qrSortBy = *o.SortBy
		}
		qSortBy := qrSortBy
		if qSortBy != "" {

			if err := r.SetQueryParam("sortBy", qSortBy); err != nil {
				return err
			}
		}
	}

	if o.SortOrder != nil {

		// query param sortOrder
		var qrSortOrder string

		if o.SortOrder != nil {
			qrSortOrder = *o.SortOrder
		}
		qSortOrder := qrSortOrder
		if qSortOrder != "" {

			if err := r.SetQueryParam("sortOrder", qSortOrder); err != nil {
				return err
			}
		}
	}

	if o.Title != nil {

		// query param title
		var qrTitle string

		if o.Title != nil {
			qrTitle = *o.Title
		}
		qTitle := qrTitle
		if qTitle != "" {

			if err := r.SetQueryParam("title", qTitle); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetKnowledgeKnowledgebaseLanguageDocuments binds the parameter documentIds
func (o *GetKnowledgeKnowledgebaseLanguageDocumentsParams) bindParamDocumentIds(formats strfmt.Registry) []string {
	documentIdsIR := o.DocumentIds

	var documentIdsIC []string
	for _, documentIdsIIR := range documentIdsIR { // explode []string

		documentIdsIIV := documentIdsIIR // string as string
		documentIdsIC = append(documentIdsIC, documentIdsIIV)
	}

	// items.CollectionFormat: "multi"
	documentIdsIS := swag.JoinByFormat(documentIdsIC, "multi")

	return documentIdsIS
}
