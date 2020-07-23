// Code generated by go-swagger; DO NOT EDIT.

package architect

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

// NewGetFlowsParams creates a new GetFlowsParams object
// with the default values initialized.
func NewGetFlowsParams() *GetFlowsParams {
	var (
		deletedDefault        = bool(false)
		includeSchemasDefault = bool(false)
		pageNumberDefault     = int32(1)
		pageSizeDefault       = int32(25)
		sortByDefault         = string("id")
		sortOrderDefault      = string("asc")
	)
	return &GetFlowsParams{
		Deleted:        &deletedDefault,
		IncludeSchemas: &includeSchemasDefault,
		PageNumber:     &pageNumberDefault,
		PageSize:       &pageSizeDefault,
		SortBy:         &sortByDefault,
		SortOrder:      &sortOrderDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFlowsParamsWithTimeout creates a new GetFlowsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFlowsParamsWithTimeout(timeout time.Duration) *GetFlowsParams {
	var (
		deletedDefault        = bool(false)
		includeSchemasDefault = bool(false)
		pageNumberDefault     = int32(1)
		pageSizeDefault       = int32(25)
		sortByDefault         = string("id")
		sortOrderDefault      = string("asc")
	)
	return &GetFlowsParams{
		Deleted:        &deletedDefault,
		IncludeSchemas: &includeSchemasDefault,
		PageNumber:     &pageNumberDefault,
		PageSize:       &pageSizeDefault,
		SortBy:         &sortByDefault,
		SortOrder:      &sortOrderDefault,

		timeout: timeout,
	}
}

// NewGetFlowsParamsWithContext creates a new GetFlowsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFlowsParamsWithContext(ctx context.Context) *GetFlowsParams {
	var (
		deletedDefault        = bool(false)
		includeSchemasDefault = bool(false)
		pageNumberDefault     = int32(1)
		pageSizeDefault       = int32(25)
		sortByDefault         = string("id")
		sortOrderDefault      = string("asc")
	)
	return &GetFlowsParams{
		Deleted:        &deletedDefault,
		IncludeSchemas: &includeSchemasDefault,
		PageNumber:     &pageNumberDefault,
		PageSize:       &pageSizeDefault,
		SortBy:         &sortByDefault,
		SortOrder:      &sortOrderDefault,

		Context: ctx,
	}
}

// NewGetFlowsParamsWithHTTPClient creates a new GetFlowsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFlowsParamsWithHTTPClient(client *http.Client) *GetFlowsParams {
	var (
		deletedDefault        = bool(false)
		includeSchemasDefault = bool(false)
		pageNumberDefault     = int32(1)
		pageSizeDefault       = int32(25)
		sortByDefault         = string("id")
		sortOrderDefault      = string("asc")
	)
	return &GetFlowsParams{
		Deleted:        &deletedDefault,
		IncludeSchemas: &includeSchemasDefault,
		PageNumber:     &pageNumberDefault,
		PageSize:       &pageSizeDefault,
		SortBy:         &sortByDefault,
		SortOrder:      &sortOrderDefault,
		HTTPClient:     client,
	}
}

/*GetFlowsParams contains all the parameters to send to the API endpoint
for the get flows operation typically these are written to a http.Request
*/
type GetFlowsParams struct {

	/*Deleted
	  Include deleted

	*/
	Deleted *bool
	/*Description
	  Description

	*/
	Description *string
	/*DivisionID
	  division ID(s)

	*/
	DivisionID []string
	/*EditableBy
	  Editable by

	*/
	EditableBy *string
	/*ID
	  ID

	*/
	ID []string
	/*IncludeSchemas
	  Include variable schemas

	*/
	IncludeSchemas *bool
	/*LockedBy
	  Locked by

	*/
	LockedBy *string
	/*LockedByClientID
	  Locked by client ID

	*/
	LockedByClientID *string
	/*Name
	  Name

	*/
	Name *string
	/*NameOrDescription
	  Name or description

	*/
	NameOrDescription *string
	/*PageNumber
	  Page number

	*/
	PageNumber *int32
	/*PageSize
	  Page size

	*/
	PageSize *int32
	/*PublishVersionID
	  Publish version ID

	*/
	PublishVersionID *string
	/*PublishedAfter
	  Published after

	*/
	PublishedAfter *string
	/*PublishedBefore
	  Published before

	*/
	PublishedBefore *string
	/*Secure
	  Secure

	*/
	Secure *string
	/*SortBy
	  Sort by

	*/
	SortBy *string
	/*SortOrder
	  Sort order

	*/
	SortOrder *string
	/*Type
	  Type

	*/
	Type []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get flows params
func (o *GetFlowsParams) WithTimeout(timeout time.Duration) *GetFlowsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get flows params
func (o *GetFlowsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get flows params
func (o *GetFlowsParams) WithContext(ctx context.Context) *GetFlowsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get flows params
func (o *GetFlowsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get flows params
func (o *GetFlowsParams) WithHTTPClient(client *http.Client) *GetFlowsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get flows params
func (o *GetFlowsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeleted adds the deleted to the get flows params
func (o *GetFlowsParams) WithDeleted(deleted *bool) *GetFlowsParams {
	o.SetDeleted(deleted)
	return o
}

// SetDeleted adds the deleted to the get flows params
func (o *GetFlowsParams) SetDeleted(deleted *bool) {
	o.Deleted = deleted
}

// WithDescription adds the description to the get flows params
func (o *GetFlowsParams) WithDescription(description *string) *GetFlowsParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the get flows params
func (o *GetFlowsParams) SetDescription(description *string) {
	o.Description = description
}

// WithDivisionID adds the divisionID to the get flows params
func (o *GetFlowsParams) WithDivisionID(divisionID []string) *GetFlowsParams {
	o.SetDivisionID(divisionID)
	return o
}

// SetDivisionID adds the divisionId to the get flows params
func (o *GetFlowsParams) SetDivisionID(divisionID []string) {
	o.DivisionID = divisionID
}

// WithEditableBy adds the editableBy to the get flows params
func (o *GetFlowsParams) WithEditableBy(editableBy *string) *GetFlowsParams {
	o.SetEditableBy(editableBy)
	return o
}

// SetEditableBy adds the editableBy to the get flows params
func (o *GetFlowsParams) SetEditableBy(editableBy *string) {
	o.EditableBy = editableBy
}

// WithID adds the id to the get flows params
func (o *GetFlowsParams) WithID(id []string) *GetFlowsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get flows params
func (o *GetFlowsParams) SetID(id []string) {
	o.ID = id
}

// WithIncludeSchemas adds the includeSchemas to the get flows params
func (o *GetFlowsParams) WithIncludeSchemas(includeSchemas *bool) *GetFlowsParams {
	o.SetIncludeSchemas(includeSchemas)
	return o
}

// SetIncludeSchemas adds the includeSchemas to the get flows params
func (o *GetFlowsParams) SetIncludeSchemas(includeSchemas *bool) {
	o.IncludeSchemas = includeSchemas
}

// WithLockedBy adds the lockedBy to the get flows params
func (o *GetFlowsParams) WithLockedBy(lockedBy *string) *GetFlowsParams {
	o.SetLockedBy(lockedBy)
	return o
}

// SetLockedBy adds the lockedBy to the get flows params
func (o *GetFlowsParams) SetLockedBy(lockedBy *string) {
	o.LockedBy = lockedBy
}

// WithLockedByClientID adds the lockedByClientID to the get flows params
func (o *GetFlowsParams) WithLockedByClientID(lockedByClientID *string) *GetFlowsParams {
	o.SetLockedByClientID(lockedByClientID)
	return o
}

// SetLockedByClientID adds the lockedByClientId to the get flows params
func (o *GetFlowsParams) SetLockedByClientID(lockedByClientID *string) {
	o.LockedByClientID = lockedByClientID
}

// WithName adds the name to the get flows params
func (o *GetFlowsParams) WithName(name *string) *GetFlowsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get flows params
func (o *GetFlowsParams) SetName(name *string) {
	o.Name = name
}

// WithNameOrDescription adds the nameOrDescription to the get flows params
func (o *GetFlowsParams) WithNameOrDescription(nameOrDescription *string) *GetFlowsParams {
	o.SetNameOrDescription(nameOrDescription)
	return o
}

// SetNameOrDescription adds the nameOrDescription to the get flows params
func (o *GetFlowsParams) SetNameOrDescription(nameOrDescription *string) {
	o.NameOrDescription = nameOrDescription
}

// WithPageNumber adds the pageNumber to the get flows params
func (o *GetFlowsParams) WithPageNumber(pageNumber *int32) *GetFlowsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get flows params
func (o *GetFlowsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get flows params
func (o *GetFlowsParams) WithPageSize(pageSize *int32) *GetFlowsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get flows params
func (o *GetFlowsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithPublishVersionID adds the publishVersionID to the get flows params
func (o *GetFlowsParams) WithPublishVersionID(publishVersionID *string) *GetFlowsParams {
	o.SetPublishVersionID(publishVersionID)
	return o
}

// SetPublishVersionID adds the publishVersionId to the get flows params
func (o *GetFlowsParams) SetPublishVersionID(publishVersionID *string) {
	o.PublishVersionID = publishVersionID
}

// WithPublishedAfter adds the publishedAfter to the get flows params
func (o *GetFlowsParams) WithPublishedAfter(publishedAfter *string) *GetFlowsParams {
	o.SetPublishedAfter(publishedAfter)
	return o
}

// SetPublishedAfter adds the publishedAfter to the get flows params
func (o *GetFlowsParams) SetPublishedAfter(publishedAfter *string) {
	o.PublishedAfter = publishedAfter
}

// WithPublishedBefore adds the publishedBefore to the get flows params
func (o *GetFlowsParams) WithPublishedBefore(publishedBefore *string) *GetFlowsParams {
	o.SetPublishedBefore(publishedBefore)
	return o
}

// SetPublishedBefore adds the publishedBefore to the get flows params
func (o *GetFlowsParams) SetPublishedBefore(publishedBefore *string) {
	o.PublishedBefore = publishedBefore
}

// WithSecure adds the secure to the get flows params
func (o *GetFlowsParams) WithSecure(secure *string) *GetFlowsParams {
	o.SetSecure(secure)
	return o
}

// SetSecure adds the secure to the get flows params
func (o *GetFlowsParams) SetSecure(secure *string) {
	o.Secure = secure
}

// WithSortBy adds the sortBy to the get flows params
func (o *GetFlowsParams) WithSortBy(sortBy *string) *GetFlowsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get flows params
func (o *GetFlowsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get flows params
func (o *GetFlowsParams) WithSortOrder(sortOrder *string) *GetFlowsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get flows params
func (o *GetFlowsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WithType adds the typeVar to the get flows params
func (o *GetFlowsParams) WithType(typeVar []string) *GetFlowsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get flows params
func (o *GetFlowsParams) SetType(typeVar []string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetFlowsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Deleted != nil {

		// query param deleted
		var qrDeleted bool
		if o.Deleted != nil {
			qrDeleted = *o.Deleted
		}
		qDeleted := swag.FormatBool(qrDeleted)
		if qDeleted != "" {
			if err := r.SetQueryParam("deleted", qDeleted); err != nil {
				return err
			}
		}

	}

	if o.Description != nil {

		// query param description
		var qrDescription string
		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {
			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}

	}

	valuesDivisionID := o.DivisionID

	joinedDivisionID := swag.JoinByFormat(valuesDivisionID, "multi")
	// query array param divisionId
	if err := r.SetQueryParam("divisionId", joinedDivisionID...); err != nil {
		return err
	}

	if o.EditableBy != nil {

		// query param editableBy
		var qrEditableBy string
		if o.EditableBy != nil {
			qrEditableBy = *o.EditableBy
		}
		qEditableBy := qrEditableBy
		if qEditableBy != "" {
			if err := r.SetQueryParam("editableBy", qEditableBy); err != nil {
				return err
			}
		}

	}

	valuesID := o.ID

	joinedID := swag.JoinByFormat(valuesID, "multi")
	// query array param id
	if err := r.SetQueryParam("id", joinedID...); err != nil {
		return err
	}

	if o.IncludeSchemas != nil {

		// query param includeSchemas
		var qrIncludeSchemas bool
		if o.IncludeSchemas != nil {
			qrIncludeSchemas = *o.IncludeSchemas
		}
		qIncludeSchemas := swag.FormatBool(qrIncludeSchemas)
		if qIncludeSchemas != "" {
			if err := r.SetQueryParam("includeSchemas", qIncludeSchemas); err != nil {
				return err
			}
		}

	}

	if o.LockedBy != nil {

		// query param lockedBy
		var qrLockedBy string
		if o.LockedBy != nil {
			qrLockedBy = *o.LockedBy
		}
		qLockedBy := qrLockedBy
		if qLockedBy != "" {
			if err := r.SetQueryParam("lockedBy", qLockedBy); err != nil {
				return err
			}
		}

	}

	if o.LockedByClientID != nil {

		// query param lockedByClientId
		var qrLockedByClientID string
		if o.LockedByClientID != nil {
			qrLockedByClientID = *o.LockedByClientID
		}
		qLockedByClientID := qrLockedByClientID
		if qLockedByClientID != "" {
			if err := r.SetQueryParam("lockedByClientId", qLockedByClientID); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.NameOrDescription != nil {

		// query param nameOrDescription
		var qrNameOrDescription string
		if o.NameOrDescription != nil {
			qrNameOrDescription = *o.NameOrDescription
		}
		qNameOrDescription := qrNameOrDescription
		if qNameOrDescription != "" {
			if err := r.SetQueryParam("nameOrDescription", qNameOrDescription); err != nil {
				return err
			}
		}

	}

	if o.PageNumber != nil {

		// query param pageNumber
		var qrPageNumber int32
		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := swag.FormatInt32(qrPageNumber)
		if qPageNumber != "" {
			if err := r.SetQueryParam("pageNumber", qPageNumber); err != nil {
				return err
			}
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

	if o.PublishVersionID != nil {

		// query param publishVersionId
		var qrPublishVersionID string
		if o.PublishVersionID != nil {
			qrPublishVersionID = *o.PublishVersionID
		}
		qPublishVersionID := qrPublishVersionID
		if qPublishVersionID != "" {
			if err := r.SetQueryParam("publishVersionId", qPublishVersionID); err != nil {
				return err
			}
		}

	}

	if o.PublishedAfter != nil {

		// query param publishedAfter
		var qrPublishedAfter string
		if o.PublishedAfter != nil {
			qrPublishedAfter = *o.PublishedAfter
		}
		qPublishedAfter := qrPublishedAfter
		if qPublishedAfter != "" {
			if err := r.SetQueryParam("publishedAfter", qPublishedAfter); err != nil {
				return err
			}
		}

	}

	if o.PublishedBefore != nil {

		// query param publishedBefore
		var qrPublishedBefore string
		if o.PublishedBefore != nil {
			qrPublishedBefore = *o.PublishedBefore
		}
		qPublishedBefore := qrPublishedBefore
		if qPublishedBefore != "" {
			if err := r.SetQueryParam("publishedBefore", qPublishedBefore); err != nil {
				return err
			}
		}

	}

	if o.Secure != nil {

		// query param secure
		var qrSecure string
		if o.Secure != nil {
			qrSecure = *o.Secure
		}
		qSecure := qrSecure
		if qSecure != "" {
			if err := r.SetQueryParam("secure", qSecure); err != nil {
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

	valuesType := o.Type

	joinedType := swag.JoinByFormat(valuesType, "multi")
	// query array param type
	if err := r.SetQueryParam("type", joinedType...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
