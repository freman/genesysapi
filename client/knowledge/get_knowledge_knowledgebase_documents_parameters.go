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

// NewGetKnowledgeKnowledgebaseDocumentsParams creates a new GetKnowledgeKnowledgebaseDocumentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetKnowledgeKnowledgebaseDocumentsParams() *GetKnowledgeKnowledgebaseDocumentsParams {
	return &GetKnowledgeKnowledgebaseDocumentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetKnowledgeKnowledgebaseDocumentsParamsWithTimeout creates a new GetKnowledgeKnowledgebaseDocumentsParams object
// with the ability to set a timeout on a request.
func NewGetKnowledgeKnowledgebaseDocumentsParamsWithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseDocumentsParams {
	return &GetKnowledgeKnowledgebaseDocumentsParams{
		timeout: timeout,
	}
}

// NewGetKnowledgeKnowledgebaseDocumentsParamsWithContext creates a new GetKnowledgeKnowledgebaseDocumentsParams object
// with the ability to set a context for a request.
func NewGetKnowledgeKnowledgebaseDocumentsParamsWithContext(ctx context.Context) *GetKnowledgeKnowledgebaseDocumentsParams {
	return &GetKnowledgeKnowledgebaseDocumentsParams{
		Context: ctx,
	}
}

// NewGetKnowledgeKnowledgebaseDocumentsParamsWithHTTPClient creates a new GetKnowledgeKnowledgebaseDocumentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetKnowledgeKnowledgebaseDocumentsParamsWithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseDocumentsParams {
	return &GetKnowledgeKnowledgebaseDocumentsParams{
		HTTPClient: client,
	}
}

/*
GetKnowledgeKnowledgebaseDocumentsParams contains all the parameters to send to the API endpoint

	for the get knowledge knowledgebase documents operation.

	Typically these are written to a http.Request.
*/
type GetKnowledgeKnowledgebaseDocumentsParams struct {

	/* After.

	   The cursor that points to the end of the set of entities that has been returned.
	*/
	After *string

	/* Before.

	   The cursor that points to the start of the set of entities that has been returned.
	*/
	Before *string

	/* CategoryID.

	   If specified, retrieves documents associated with category ids, comma separated values expected.
	*/
	CategoryID []string

	/* DocumentID.

	   Retrieves the specified documents, comma separated values expected.
	*/
	DocumentID []string

	/* Expand.

	   The specified entity attributes will be filled. Comma separated values expected.
	*/
	Expand []string

	/* IncludeDrafts.

	   If includeDrafts is true, Documents in the draft state are also returned in the response.
	*/
	IncludeDrafts *bool

	/* IncludeSubcategories.

	   Works along with 'categoryId' query parameter. If specified, retrieves documents associated with category ids and its children categories.
	*/
	IncludeSubcategories *bool

	/* Interval.

	   Retrieves the documents modified in specified date and time range. If the after and before cursor parameters are within this interval, it would return valid data, otherwise it throws an error.The dates in the interval are represented in ISO-8601 format: YYYY-MM-DDThh:mm:ssZ/YYYY-MM-DDThh:mm:ssZ

	   Format: interval
	*/
	Interval *string

	/* KnowledgeBaseID.

	   Knowledge base ID
	*/
	KnowledgeBaseID string

	/* LabelIds.

	   If specified, retrieves documents associated with label ids, comma separated values expected.
	*/
	LabelIds []string

	/* PageSize.

	   Number of entities to return. Maximum of 200.
	*/
	PageSize *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get knowledge knowledgebase documents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKnowledgeKnowledgebaseDocumentsParams) WithDefaults() *GetKnowledgeKnowledgebaseDocumentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get knowledge knowledgebase documents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKnowledgeKnowledgebaseDocumentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) WithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseDocumentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) WithContext(ctx context.Context) *GetKnowledgeKnowledgebaseDocumentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) WithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseDocumentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAfter adds the after to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) WithAfter(after *string) *GetKnowledgeKnowledgebaseDocumentsParams {
	o.SetAfter(after)
	return o
}

// SetAfter adds the after to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) SetAfter(after *string) {
	o.After = after
}

// WithBefore adds the before to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) WithBefore(before *string) *GetKnowledgeKnowledgebaseDocumentsParams {
	o.SetBefore(before)
	return o
}

// SetBefore adds the before to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) SetBefore(before *string) {
	o.Before = before
}

// WithCategoryID adds the categoryID to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) WithCategoryID(categoryID []string) *GetKnowledgeKnowledgebaseDocumentsParams {
	o.SetCategoryID(categoryID)
	return o
}

// SetCategoryID adds the categoryId to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) SetCategoryID(categoryID []string) {
	o.CategoryID = categoryID
}

// WithDocumentID adds the documentID to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) WithDocumentID(documentID []string) *GetKnowledgeKnowledgebaseDocumentsParams {
	o.SetDocumentID(documentID)
	return o
}

// SetDocumentID adds the documentId to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) SetDocumentID(documentID []string) {
	o.DocumentID = documentID
}

// WithExpand adds the expand to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) WithExpand(expand []string) *GetKnowledgeKnowledgebaseDocumentsParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithIncludeDrafts adds the includeDrafts to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) WithIncludeDrafts(includeDrafts *bool) *GetKnowledgeKnowledgebaseDocumentsParams {
	o.SetIncludeDrafts(includeDrafts)
	return o
}

// SetIncludeDrafts adds the includeDrafts to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) SetIncludeDrafts(includeDrafts *bool) {
	o.IncludeDrafts = includeDrafts
}

// WithIncludeSubcategories adds the includeSubcategories to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) WithIncludeSubcategories(includeSubcategories *bool) *GetKnowledgeKnowledgebaseDocumentsParams {
	o.SetIncludeSubcategories(includeSubcategories)
	return o
}

// SetIncludeSubcategories adds the includeSubcategories to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) SetIncludeSubcategories(includeSubcategories *bool) {
	o.IncludeSubcategories = includeSubcategories
}

// WithInterval adds the interval to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) WithInterval(interval *string) *GetKnowledgeKnowledgebaseDocumentsParams {
	o.SetInterval(interval)
	return o
}

// SetInterval adds the interval to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) SetInterval(interval *string) {
	o.Interval = interval
}

// WithKnowledgeBaseID adds the knowledgeBaseID to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) WithKnowledgeBaseID(knowledgeBaseID string) *GetKnowledgeKnowledgebaseDocumentsParams {
	o.SetKnowledgeBaseID(knowledgeBaseID)
	return o
}

// SetKnowledgeBaseID adds the knowledgeBaseId to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) SetKnowledgeBaseID(knowledgeBaseID string) {
	o.KnowledgeBaseID = knowledgeBaseID
}

// WithLabelIds adds the labelIds to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) WithLabelIds(labelIds []string) *GetKnowledgeKnowledgebaseDocumentsParams {
	o.SetLabelIds(labelIds)
	return o
}

// SetLabelIds adds the labelIds to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) SetLabelIds(labelIds []string) {
	o.LabelIds = labelIds
}

// WithPageSize adds the pageSize to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) WithPageSize(pageSize *string) *GetKnowledgeKnowledgebaseDocumentsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get knowledge knowledgebase documents params
func (o *GetKnowledgeKnowledgebaseDocumentsParams) SetPageSize(pageSize *string) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetKnowledgeKnowledgebaseDocumentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.CategoryID != nil {

		// binding items for categoryId
		joinedCategoryID := o.bindParamCategoryID(reg)

		// query array param categoryId
		if err := r.SetQueryParam("categoryId", joinedCategoryID...); err != nil {
			return err
		}
	}

	if o.DocumentID != nil {

		// binding items for documentId
		joinedDocumentID := o.bindParamDocumentID(reg)

		// query array param documentId
		if err := r.SetQueryParam("documentId", joinedDocumentID...); err != nil {
			return err
		}
	}

	if o.Expand != nil {

		// binding items for expand
		joinedExpand := o.bindParamExpand(reg)

		// query array param expand
		if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
			return err
		}
	}

	if o.IncludeDrafts != nil {

		// query param includeDrafts
		var qrIncludeDrafts bool

		if o.IncludeDrafts != nil {
			qrIncludeDrafts = *o.IncludeDrafts
		}
		qIncludeDrafts := swag.FormatBool(qrIncludeDrafts)
		if qIncludeDrafts != "" {

			if err := r.SetQueryParam("includeDrafts", qIncludeDrafts); err != nil {
				return err
			}
		}
	}

	if o.IncludeSubcategories != nil {

		// query param includeSubcategories
		var qrIncludeSubcategories bool

		if o.IncludeSubcategories != nil {
			qrIncludeSubcategories = *o.IncludeSubcategories
		}
		qIncludeSubcategories := swag.FormatBool(qrIncludeSubcategories)
		if qIncludeSubcategories != "" {

			if err := r.SetQueryParam("includeSubcategories", qIncludeSubcategories); err != nil {
				return err
			}
		}
	}

	if o.Interval != nil {

		// query param interval
		var qrInterval string

		if o.Interval != nil {
			qrInterval = *o.Interval
		}
		qInterval := qrInterval
		if qInterval != "" {

			if err := r.SetQueryParam("interval", qInterval); err != nil {
				return err
			}
		}
	}

	// path param knowledgeBaseId
	if err := r.SetPathParam("knowledgeBaseId", o.KnowledgeBaseID); err != nil {
		return err
	}

	if o.LabelIds != nil {

		// binding items for labelIds
		joinedLabelIds := o.bindParamLabelIds(reg)

		// query array param labelIds
		if err := r.SetQueryParam("labelIds", joinedLabelIds...); err != nil {
			return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetKnowledgeKnowledgebaseDocuments binds the parameter categoryId
func (o *GetKnowledgeKnowledgebaseDocumentsParams) bindParamCategoryID(formats strfmt.Registry) []string {
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

// bindParamGetKnowledgeKnowledgebaseDocuments binds the parameter documentId
func (o *GetKnowledgeKnowledgebaseDocumentsParams) bindParamDocumentID(formats strfmt.Registry) []string {
	documentIDIR := o.DocumentID

	var documentIDIC []string
	for _, documentIDIIR := range documentIDIR { // explode []string

		documentIDIIV := documentIDIIR // string as string
		documentIDIC = append(documentIDIC, documentIDIIV)
	}

	// items.CollectionFormat: "multi"
	documentIDIS := swag.JoinByFormat(documentIDIC, "multi")

	return documentIDIS
}

// bindParamGetKnowledgeKnowledgebaseDocuments binds the parameter expand
func (o *GetKnowledgeKnowledgebaseDocumentsParams) bindParamExpand(formats strfmt.Registry) []string {
	expandIR := o.Expand

	var expandIC []string
	for _, expandIIR := range expandIR { // explode []string

		expandIIV := expandIIR // string as string
		expandIC = append(expandIC, expandIIV)
	}

	// items.CollectionFormat: "multi"
	expandIS := swag.JoinByFormat(expandIC, "multi")

	return expandIS
}

// bindParamGetKnowledgeKnowledgebaseDocuments binds the parameter labelIds
func (o *GetKnowledgeKnowledgebaseDocumentsParams) bindParamLabelIds(formats strfmt.Registry) []string {
	labelIdsIR := o.LabelIds

	var labelIdsIC []string
	for _, labelIdsIIR := range labelIdsIR { // explode []string

		labelIdsIIV := labelIdsIIR // string as string
		labelIdsIC = append(labelIdsIC, labelIdsIIV)
	}

	// items.CollectionFormat: "multi"
	labelIdsIS := swag.JoinByFormat(labelIdsIC, "multi")

	return labelIdsIS
}
