// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// KnowledgeDocumentSearchRequest knowledge document search request
//
// swagger:model KnowledgeDocumentSearchRequest
type KnowledgeDocumentSearchRequest struct {

	// The client application details from which search request was sent.
	Application *KnowledgeSearchClientApplication `json:"application,omitempty"`

	// Conversation context information if the search is initiated in the context of a conversation.
	ConversationContext *KnowledgeConversationContext `json:"conversationContext,omitempty"`

	// Filter for the document search.
	Filter *DocumentQuery `json:"filter,omitempty"`

	// Indicates whether the search results would also include draft documents.
	IncludeDraftDocuments bool `json:"includeDraftDocuments"`

	// Retrieves the documents created/modified/published in specified date and time range.
	Interval *DocumentQueryInterval `json:"interval,omitempty"`

	// Number of pages returned in the result calculated according to the pageSize and the total
	// Read Only: true
	PageCount int32 `json:"pageCount,omitempty"`

	// Page number of the returned results.
	PageNumber int32 `json:"pageNumber,omitempty"`

	// Page size of the returned results.
	PageSize int32 `json:"pageSize,omitempty"`

	// Query to search content in the knowledge base. Maximum of 30 records per query can be fetched.
	// Required: true
	// Max Length: 2147483647
	// Min Length: 3
	Query *string `json:"query"`

	// The type of the query that initiates the search.
	// Enum: [AutoSearch ManualSearch Suggestion]
	QueryType string `json:"queryType,omitempty"`

	// The globally unique identifier for the search.
	// Read Only: true
	SearchID string `json:"searchId,omitempty"`

	// The field in the documents that you want to sort the search results by.
	// Enum: [ConfidenceScore DateCreated DateModified CategoryName LabelName]
	SortBy string `json:"sortBy,omitempty"`

	// The sort order for search results.
	// Enum: [Asc Desc]
	SortOrder string `json:"sortOrder,omitempty"`

	// The total number of documents matching the query.
	// Read Only: true
	Total int32 `json:"total,omitempty"`
}

// Validate validates this knowledge document search request
func (m *KnowledgeDocumentSearchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConversationContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueryType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSortBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSortOrder(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KnowledgeDocumentSearchRequest) validateApplication(formats strfmt.Registry) error {
	if swag.IsZero(m.Application) { // not required
		return nil
	}

	if m.Application != nil {
		if err := m.Application.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("application")
			}
			return err
		}
	}

	return nil
}

func (m *KnowledgeDocumentSearchRequest) validateConversationContext(formats strfmt.Registry) error {
	if swag.IsZero(m.ConversationContext) { // not required
		return nil
	}

	if m.ConversationContext != nil {
		if err := m.ConversationContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conversationContext")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("conversationContext")
			}
			return err
		}
	}

	return nil
}

func (m *KnowledgeDocumentSearchRequest) validateFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.Filter) { // not required
		return nil
	}

	if m.Filter != nil {
		if err := m.Filter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filter")
			}
			return err
		}
	}

	return nil
}

func (m *KnowledgeDocumentSearchRequest) validateInterval(formats strfmt.Registry) error {
	if swag.IsZero(m.Interval) { // not required
		return nil
	}

	if m.Interval != nil {
		if err := m.Interval.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interval")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("interval")
			}
			return err
		}
	}

	return nil
}

func (m *KnowledgeDocumentSearchRequest) validateQuery(formats strfmt.Registry) error {

	if err := validate.Required("query", "body", m.Query); err != nil {
		return err
	}

	if err := validate.MinLength("query", "body", *m.Query, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("query", "body", *m.Query, 2147483647); err != nil {
		return err
	}

	return nil
}

var knowledgeDocumentSearchRequestTypeQueryTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AutoSearch","ManualSearch","Suggestion"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		knowledgeDocumentSearchRequestTypeQueryTypePropEnum = append(knowledgeDocumentSearchRequestTypeQueryTypePropEnum, v)
	}
}

const (

	// KnowledgeDocumentSearchRequestQueryTypeAutoSearch captures enum value "AutoSearch"
	KnowledgeDocumentSearchRequestQueryTypeAutoSearch string = "AutoSearch"

	// KnowledgeDocumentSearchRequestQueryTypeManualSearch captures enum value "ManualSearch"
	KnowledgeDocumentSearchRequestQueryTypeManualSearch string = "ManualSearch"

	// KnowledgeDocumentSearchRequestQueryTypeSuggestion captures enum value "Suggestion"
	KnowledgeDocumentSearchRequestQueryTypeSuggestion string = "Suggestion"
)

// prop value enum
func (m *KnowledgeDocumentSearchRequest) validateQueryTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, knowledgeDocumentSearchRequestTypeQueryTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KnowledgeDocumentSearchRequest) validateQueryType(formats strfmt.Registry) error {
	if swag.IsZero(m.QueryType) { // not required
		return nil
	}

	// value enum
	if err := m.validateQueryTypeEnum("queryType", "body", m.QueryType); err != nil {
		return err
	}

	return nil
}

var knowledgeDocumentSearchRequestTypeSortByPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ConfidenceScore","DateCreated","DateModified","CategoryName","LabelName"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		knowledgeDocumentSearchRequestTypeSortByPropEnum = append(knowledgeDocumentSearchRequestTypeSortByPropEnum, v)
	}
}

const (

	// KnowledgeDocumentSearchRequestSortByConfidenceScore captures enum value "ConfidenceScore"
	KnowledgeDocumentSearchRequestSortByConfidenceScore string = "ConfidenceScore"

	// KnowledgeDocumentSearchRequestSortByDateCreated captures enum value "DateCreated"
	KnowledgeDocumentSearchRequestSortByDateCreated string = "DateCreated"

	// KnowledgeDocumentSearchRequestSortByDateModified captures enum value "DateModified"
	KnowledgeDocumentSearchRequestSortByDateModified string = "DateModified"

	// KnowledgeDocumentSearchRequestSortByCategoryName captures enum value "CategoryName"
	KnowledgeDocumentSearchRequestSortByCategoryName string = "CategoryName"

	// KnowledgeDocumentSearchRequestSortByLabelName captures enum value "LabelName"
	KnowledgeDocumentSearchRequestSortByLabelName string = "LabelName"
)

// prop value enum
func (m *KnowledgeDocumentSearchRequest) validateSortByEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, knowledgeDocumentSearchRequestTypeSortByPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KnowledgeDocumentSearchRequest) validateSortBy(formats strfmt.Registry) error {
	if swag.IsZero(m.SortBy) { // not required
		return nil
	}

	// value enum
	if err := m.validateSortByEnum("sortBy", "body", m.SortBy); err != nil {
		return err
	}

	return nil
}

var knowledgeDocumentSearchRequestTypeSortOrderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Asc","Desc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		knowledgeDocumentSearchRequestTypeSortOrderPropEnum = append(knowledgeDocumentSearchRequestTypeSortOrderPropEnum, v)
	}
}

const (

	// KnowledgeDocumentSearchRequestSortOrderAsc captures enum value "Asc"
	KnowledgeDocumentSearchRequestSortOrderAsc string = "Asc"

	// KnowledgeDocumentSearchRequestSortOrderDesc captures enum value "Desc"
	KnowledgeDocumentSearchRequestSortOrderDesc string = "Desc"
)

// prop value enum
func (m *KnowledgeDocumentSearchRequest) validateSortOrderEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, knowledgeDocumentSearchRequestTypeSortOrderPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KnowledgeDocumentSearchRequest) validateSortOrder(formats strfmt.Registry) error {
	if swag.IsZero(m.SortOrder) { // not required
		return nil
	}

	// value enum
	if err := m.validateSortOrderEnum("sortOrder", "body", m.SortOrder); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this knowledge document search request based on the context it is used
func (m *KnowledgeDocumentSearchRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConversationContext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterval(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePageCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSearchID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KnowledgeDocumentSearchRequest) contextValidateApplication(ctx context.Context, formats strfmt.Registry) error {

	if m.Application != nil {
		if err := m.Application.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("application")
			}
			return err
		}
	}

	return nil
}

func (m *KnowledgeDocumentSearchRequest) contextValidateConversationContext(ctx context.Context, formats strfmt.Registry) error {

	if m.ConversationContext != nil {
		if err := m.ConversationContext.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conversationContext")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("conversationContext")
			}
			return err
		}
	}

	return nil
}

func (m *KnowledgeDocumentSearchRequest) contextValidateFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.Filter != nil {
		if err := m.Filter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filter")
			}
			return err
		}
	}

	return nil
}

func (m *KnowledgeDocumentSearchRequest) contextValidateInterval(ctx context.Context, formats strfmt.Registry) error {

	if m.Interval != nil {
		if err := m.Interval.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interval")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("interval")
			}
			return err
		}
	}

	return nil
}

func (m *KnowledgeDocumentSearchRequest) contextValidatePageCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "pageCount", "body", int32(m.PageCount)); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeDocumentSearchRequest) contextValidateSearchID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "searchId", "body", string(m.SearchID)); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeDocumentSearchRequest) contextValidateTotal(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "total", "body", int32(m.Total)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KnowledgeDocumentSearchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KnowledgeDocumentSearchRequest) UnmarshalBinary(b []byte) error {
	var res KnowledgeDocumentSearchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
